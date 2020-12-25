package qtrt

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -ldl

extern void* vtablehook_hook(void* instance, void* hook, int offset);
extern void test_vthook(void* app);
extern void* find_f_address(void* aClass, void* memfnptr);
extern void dump_vtable_entry(void* elfvtptr);
extern int fillin_vtable_entry(void* elfvtptr, void** arr);

void
cxa_demangle_asmcc(struct {void* fnptr; const char *mangled; char *buf; size_t *len; int *status;} *ax) {
   void* (*fnptr2)() = ax->fnptr;
   char* res = fnptr2(ax->mangled, ax->buf, ax->len, ax->status);
   // printf("%p ***%s***\n", res, ax->mangled);
   *(ax->len) = res == NULL ? 0 : strlen(res);
}

*/
import "C"
import (
	"fmt"
	"go/scanner"
	"go/token"
	"log"
	"strings"
	"sync"

	"github.com/kitech/dl/asmcgocall"
)

func TestVthook(app Voidptr) {
	C.test_vthook(app)
}

// 这种方式似乎对dlsym取到的memfnptr不适用？
func FindVtmAddr(obj Voidptr, memfnptr Voidptr) Voidptr {
	log.Println(obj, memfnptr)
	return C.find_f_address(obj, memfnptr)
}

///// 这种方式似乎对linux, .so有效
func DumpVtableEntry(clsname string) {
	vtname := fmt.Sprintf("_ZTV%d%s", len(clsname), clsname)
	ptr := GetQtSymAddrRaw(vtname)
	log.Println(vtname, ptr)
	C.dump_vtable_entry(ptr)
}

func GetVtableEntry(clsname string) []string {
	vtname := fmt.Sprintf("_ZTV%d%s", len(clsname), clsname)
	ptr := GetQtSymAddrRaw(vtname)
	// log.Println(vtname, ptr)

	items := make([]Voidptr, 128)
	cnt := C.fillin_vtable_entry(ptr, (*Voidptr)(&items[0]))
	// log.Println(cnt)
	items = items[:int(cnt)]

	items2 := make([]string, int(cnt))
	for idx, item := range items {
		s := C.GoString((*C.char)(item))
		if false {
			log.Println(cnt, idx, s)
		}
		items2[idx] = s
	}

	return items2
}

// thread safe???
// lazy fillin for requested class
type virtualTable struct {
	mu     sync.RWMutex
	tables map[string]*vtableEntry // clsname =>
}

type vtableEntry struct {
	clsname string
	items   []string              // mangled name
	names   []string              // just pure func name like `mouseMoveEvent`
	hooks   map[string]*hookEntry // mthname =>
}

type hookEntry struct {
	mgname  string
	rawname string
	index   int
	oldfn   Voidptr
	newfn   Voidptr
	vmclos  *CppvmClosure
}

var virtab = newVirtualTable()

func newVirtualTable() *virtualTable {
	this := &virtualTable{}
	this.tables = map[string]*vtableEntry{}
	return this
}

// fn for under closure types resolve
func (this *virtualTable) hookit(clsname, mthname string, obj Voidptr, fn interface{}) {
	if this.hookExist(clsname, mthname) {
		return
	}
	idx := this.getVTMIndex(clsname, mthname)
	if idx == -1 {
		log.Println("not found", clsname, mthname)
		return
	}

	vmclos := NewCppvmClosure(clsname, mthname, fn)
	fnptr := vmclos.Func()

	//hookimpl()
	oldfn := C.vtablehook_hook(obj, fnptr, C.int(idx))
	he := &hookEntry{}
	he.mgname = ""
	he.rawname = mthname
	he.index = idx
	he.oldfn = oldfn
	he.newfn = fnptr
	he.vmclos = vmclos

	vte, ok := this.tables[clsname]
	if !ok { //wtt
	}
	vte.hooks[mthname] = he
}

func (this *virtualTable) unhookit(clsname, mthname string, obj Voidptr) {
	if !this.hookExist(clsname, mthname) {
		return
	}

	vte, ok := this.tables[clsname]
	if !ok { // wtt
	}
	he, ok := vte.hooks[mthname]
	if !ok { // wtt
	}
	oldfn := C.vtablehook_hook(obj, he.oldfn, C.int(he.index))
	if oldfn != he.newfn {
		// wtt
	}
	delete(vte.hooks, mthname)
}

func (this *virtualTable) hookExist(clsname, mthname string) bool {
	vte, ok := this.tables[clsname]
	if ok {
		_, ok := vte.hooks[mthname]
		if ok {
			return true
		}
	}
	return false
}

func (this *virtualTable) getVTMIndex(clsname, mthname string) int {
	vte, ok := this.tables[clsname]
	if !ok {
		err := this.fillinClassVtable(clsname)
		ErrPrint(err, clsname)
	}
	vte, ok = this.tables[clsname]
	if !ok {
		return -1
	}

	// TODO overloaded resolution???
	for idx, name := range vte.names {
		if name == mthname {
			return idx
		}
	}
	return -1
}

func (this *virtualTable) fillinClassVtable(clsname string) error {
	_, ok := this.tables[clsname]
	if ok {
		return nil
	}
	items := GetVtableEntry(clsname)

	e := &vtableEntry{}
	e.clsname = clsname
	e.items = items
	e.hooks = make(map[string]*hookEntry, len(items))
	e.names = make([]string, len(items))

	for idx, item := range items {
		// log.Println(idx, item)
		cmo := cppfilt(item)
		e.names[idx] = cmo.mth
	}

	this.tables[clsname] = e

	return nil
}

type cppmethod struct {
	cls     string
	mth     string
	types   []string
	names   []string
	isconst bool
}

func (this *cppmethod) isctor() bool {
	return this.cls != "" && (this.cls == this.mth)
}

func (this *cppmethod) isdtor() bool {
	return strings.HasPrefix(this.mth, "~")
}

func (this *cppmethod) ismth() bool {
	return this.cls != ""
}

func (this *cppmethod) argc() int { return len(this.types) }
func (this *cppmethod) argisref(idx int) bool {
	return strings.Count(this.types[idx], "&") == 1
}
func (this *cppmethod) argisptr(idx int) bool {
	return strings.Count(this.types[idx], "*") == 1
}

// char* qil_cxa_demangle(const char *mangled, char *buf, size_t *len, int *status)

func cxa_demangle(mgname string) (string, error) {
	mgname2 := mgname + "\x00"
	mgnameref := CStringRefRaw(&mgname2)

	blen := len(mgname)*3 + 3
	buf := make([]byte, blen)
	status := 0

	var argv = struct {
		fnptr   Voidptr
		mangled Voidptr
		buf     Voidptr
		len     *C.size_t
		status  *C.int
	}{cxa_demangle_fnptr, mgnameref, Voidptr(&buf[0]),
		(*C.size_t)(Voidptr(&blen)), (*C.int)(Voidptr(&status))}
	asmcgocall.Asmcc(C.cxa_demangle_asmcc, Voidptr(&argv))
	if status != 0 {
		return "", fmt.Errorf("demangle error %d", status)
	}
	res := string(buf[:blen])
	hintpfx := "non-virtual thunk to " // _ZThn
	if strings.HasPrefix(res, hintpfx) {
		res = res[len(hintpfx):]
	}
	hintpfx = "typeinfo for "
	if strings.HasPrefix(res, hintpfx) {
		res = res[len(hintpfx):]
	}
	return res, nil
}

// seems make execute binary size increase 600KB
// the demangle package too big
func cppfilt(mgname string) *cppmethod {
	cmo := &cppmethod{}
	// line, err := demangle.ToString(mgname)
	line, err := cxa_demangle(mgname)
	ErrPrint(err, mgname)
	// log.Println(mgname, line)

	{
		// src := []byte("cos(x) + 1i*sin(x) // Euler")
		src := []byte(line)

		// Initialize the scanner.
		var s scanner.Scanner
		fset := token.NewFileSet()                      // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
		s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

		var idts []string
		var colon = 0
		// Repeated calls to Scan yield the token sequence found in the input.
		for {
			pos, tok, lit := s.Scan()
			if tok == token.EOF {
				break
			}
			// fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
			switch tok {
			case token.CONST:
			case token.IDENT:
				idts = append(idts, lit)
			case token.COLON:
				colon += 1
				if colon == 2 {
					cmo.cls = idts[len(idts)-1]
				}
			case token.LPAREN:
				cmo.mth = idts[len(idts)-1]
				idts = []string{}
			case token.RPAREN:
				if len(idts) == 0 {
					break
				}
				tyfld := strings.Join(idts, " ")
				cmo.types = append(cmo.types, tyfld)
				idts = []string{}
			case token.MUL:
				idts = append(idts, tok.String())
			case token.AND:
				idts = append(idts, tok.String())
			case token.COMMA:
				if len(idts) == 0 {
					break
				}
				tyfld := strings.Join(idts, " ")
				cmo.types = append(cmo.types, tyfld)
				idts = []string{}
			case token.SEMICOLON:
			case token.LSS: // TODO template<>
				idts = append(idts, tok.String())
			case token.GTR:
				idts = append(idts, tok.String())
			default:
				if lit == "~" {
					idts = append(idts, tok.String())
				} else {
					log.Printf("wtt [%v] [%v] [%v]%v [%v] [%v]\n",
						fset.Position(pos), tok, lit, lit[0], mgname, line)
				}
			}
		}
		// log.Printf("%#v\n", cmo)
	}

	return cmo
}
