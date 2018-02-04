package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 221
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QCharRef struct {
	*qtrt.CObject
}

func (this *QCharRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCharRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCharRefFromPointer(cthis unsafe.Pointer) *QCharRef {
	return &QCharRef{&qtrt.CObject{cthis}}
}
func (*QCharRef) NewFromPointer(cthis unsafe.Pointer) *QCharRef {
	return NewQCharRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:1054
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QCharRef) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1055
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPrint()
func (this *QCharRef) IsPrint() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isPrintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1056
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPunct()
func (this *QCharRef) IsPunct() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isPunctEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1057
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSpace()
func (this *QCharRef) IsSpace() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isSpaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1058
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isMark()
func (this *QCharRef) IsMark() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef6isMarkEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1059
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLetter()
func (this *QCharRef) IsLetter() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8isLetterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1060
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNumber()
func (this *QCharRef) IsNumber() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8isNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1061
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLetterOrNumber()
func (this *QCharRef) IsLetterOrNumber() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef16isLetterOrNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1062
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDigit()
func (this *QCharRef) IsDigit() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isDigitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1063
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLower()
func (this *QCharRef) IsLower() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1064
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUpper()
func (this *QCharRef) IsUpper() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7isUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1065
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTitleCase()
func (this *QCharRef) IsTitleCase() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11isTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1067
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int digitValue()
func (this *QCharRef) DigitValue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef10digitValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1068
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toLower()
func (this *QCharRef) ToLower() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7toLowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1069
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toUpper()
func (this *QCharRef) ToUpper() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7toUpperEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1070
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toTitleCase()
func (this *QCharRef) ToTitleCase() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11toTitleCaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1072
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Category category()
func (this *QCharRef) Category() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8categoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1073
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Direction direction()
func (this *QCharRef) Direction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1074
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::JoiningType joiningType()
func (this *QCharRef) JoiningType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11joiningTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1076
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Joining joining()
func (this *QCharRef) Joining() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7joiningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1089
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasMirrored()
func (this *QCharRef) HasMirrored() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef11hasMirroredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1090
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar mirroredChar()
func (this *QCharRef) MirroredChar() *QChar /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef12mirroredCharEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1091
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString decomposition()
func (this *QCharRef) Decomposition() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef13decompositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1092
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Decomposition decompositionTag()
func (this *QCharRef) DecompositionTag() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef16decompositionTagEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1093
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar combiningClass()
func (this *QCharRef) CombiningClass() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef14combiningClassEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1095
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Script script()
func (this *QCharRef) Script() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef6scriptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1097
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::UnicodeVersion unicodeVersion()
func (this *QCharRef) UnicodeVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef14unicodeVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1099
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar cell()
func (this *QCharRef) Cell() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef4cellEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1100
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar row()
func (this *QCharRef) Row() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1101
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCell(uchar)
func (this *QCharRef) SetCell(cell byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef7setCellEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), cell)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRow(uchar)
func (this *QCharRef) SetRow(row byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef6setRowEh", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1107
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char toLatin1()
func (this *QCharRef) ToLatin1() byte {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef8toLatin1Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1108
// index:0
// Public inline Visibility=Default Availability=Available
// [2] ushort unicode()
func (this *QCharRef) Unicode() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QCharRef7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1109
// index:1
// Public inline Visibility=Default Availability=Available
// [2] ushort & unicode()
func (this *QCharRef) Unicode_1() uint16 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRef7unicodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint16(rv)
}

func DeleteQCharRef(this *QCharRef) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QCharRefD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end