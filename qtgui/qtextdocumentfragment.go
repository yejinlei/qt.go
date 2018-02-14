package qtgui

// /usr/include/qt/QtGui/qtextdocumentfragment.h
// #include <qtextdocumentfragment.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextDocumentFragment struct {
	*qtrt.CObject
}
type QTextDocumentFragment_ITF interface {
	QTextDocumentFragment_PTR() *QTextDocumentFragment
}

func (ptr *QTextDocumentFragment) QTextDocumentFragment_PTR() *QTextDocumentFragment { return ptr }

func (this *QTextDocumentFragment) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextDocumentFragment) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextDocumentFragmentFromPointer(cthis unsafe.Pointer) *QTextDocumentFragment {
	return &QTextDocumentFragment{&qtrt.CObject{cthis}}
}
func (*QTextDocumentFragment) NewFromPointer(cthis unsafe.Pointer) *QTextDocumentFragment {
	return NewQTextDocumentFragmentFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentFragment()
func NewQTextDocumentFragment() *QTextDocumentFragment {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentFragment(const QTextDocument *)
func NewQTextDocumentFragment_1(document QTextDocument_ITF /*777 const QTextDocument **/) *QTextDocumentFragment {
	var convArg0 unsafe.Pointer
	if document != nil && document.QTextDocument_PTR() != nil {
		convArg0 = document.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2EPK13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:59
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QTextDocumentFragment(const QTextCursor &)
func NewQTextDocumentFragment_2(range_ QTextCursor_ITF) *QTextDocumentFragment {
	var convArg0 unsafe.Pointer
	if range_ != nil && range_.QTextCursor_PTR() != nil {
		convArg0 = range_.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2ERK11QTextCursor", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextDocumentFragment)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocumentFragment & operator=(const QTextDocumentFragment &)
func (this *QTextDocumentFragment) Operator_equal(rhs QTextDocumentFragment_ITF) *QTextDocumentFragment {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextDocumentFragment_PTR() != nil {
		convArg0 = rhs.QTextDocumentFragment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragmentaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextDocumentFragment()
func DeleteQTextDocumentFragment(this *QTextDocumentFragment) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragmentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QTextDocumentFragment) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QTextDocumentFragment7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toPlainText()
func (this *QTextDocumentFragment) ToPlainText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QTextDocumentFragment11toPlainTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toHtml(const QByteArray &)
func (this *QTextDocumentFragment) ToHtml(encoding qtcore.QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if encoding != nil && encoding.QByteArray_PTR() != nil {
		convArg0 = encoding.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QTextDocumentFragment6toHtmlERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:71
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextDocumentFragment fromPlainText(const QString &)
func (this *QTextDocumentFragment) FromPlainText(plainText string) *QTextDocumentFragment /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(plainText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragment13fromPlainTextERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}
func QTextDocumentFragment_FromPlainText(plainText string) *QTextDocumentFragment /*123*/ {
	var nilthis *QTextDocumentFragment
	rv := nilthis.FromPlainText(plainText)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:73
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextDocumentFragment fromHtml(const QString &)
func (this *QTextDocumentFragment) FromHtml(html string) *QTextDocumentFragment /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}
func QTextDocumentFragment_FromHtml(html string) *QTextDocumentFragment /*123*/ {
	var nilthis *QTextDocumentFragment
	rv := nilthis.FromHtml(html)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:74
// index:1
// Public static Visibility=Default Availability=Available
// [8] QTextDocumentFragment fromHtml(const QString &, const QTextDocument *)
func (this *QTextDocumentFragment) FromHtml_1(html string, resourceProvider QTextDocument_ITF /*777 const QTextDocument **/) *QTextDocumentFragment /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(html)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if resourceProvider != nil && resourceProvider.QTextDocument_PTR() != nil {
		convArg1 = resourceProvider.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextDocumentFragmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextDocumentFragment)
	return rv2
}
func QTextDocumentFragment_FromHtml_1(html string, resourceProvider QTextDocument_ITF /*777 const QTextDocument **/) *QTextDocumentFragment /*123*/ {
	var nilthis *QTextDocumentFragment
	rv := nilthis.FromHtml_1(html, resourceProvider)
	return rv
}

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
