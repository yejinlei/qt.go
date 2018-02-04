package qtcore

// /usr/include/qt/QtCore/qcommandlineoption.h
// #include <qcommandlineoption.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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

type QCommandLineOption struct {
	*qtrt.CObject
}

func (this *QCommandLineOption) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCommandLineOption) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCommandLineOptionFromPointer(cthis unsafe.Pointer) *QCommandLineOption {
	return &QCommandLineOption{&qtrt.CObject{cthis}}
}
func (*QCommandLineOption) NewFromPointer(cthis unsafe.Pointer) *QCommandLineOption {
	return NewQCommandLineOptionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QString &)
func NewQCommandLineOption(name *QString) *QCommandLineOption {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QStringList &)
func NewQCommandLineOption_1(names *QStringList) *QCommandLineOption {
	var convArg0 = names.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:63
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QString &, const QString &, const QString &, const QString &)
func NewQCommandLineOption_2(name *QString, description *QString, valueName *QString, defaultValue *QString) *QCommandLineOption {
	var convArg0 = name.GetCthis()
	var convArg1 = description.GetCthis()
	var convArg2 = valueName.GetCthis()
	var convArg3 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QStringList &, const QString &, const QString &, const QString &)
func NewQCommandLineOption_3(names *QStringList, description *QString, valueName *QString, defaultValue *QString) *QCommandLineOption {
	var convArg0 = names.GetCthis()
	var convArg1 = description.GetCthis()
	var convArg2 = valueName.GetCthis()
	var convArg3 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCommandLineOption()
func DeleteQCommandLineOption(this *QCommandLineOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCommandLineOption &)
func (this *QCommandLineOption) Swap(other *QCommandLineOption) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValueName(const QString &)
func (this *QCommandLineOption) SetValueName(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption12setValueNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString valueName()
func (this *QCommandLineOption) ValueName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption9valueNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineoption.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)
func (this *QCommandLineOption) SetDescription(description *QString) {
	var convArg0 = description.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption14setDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description()
func (this *QCommandLineOption) Description() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption11descriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineoption.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultValue(const QString &)
func (this *QCommandLineOption) SetDefaultValue(defaultValue *QString) {
	var convArg0 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption15setDefaultValueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultValues(const QStringList &)
func (this *QCommandLineOption) SetDefaultValues(defaultValues *QStringList) {
	var convArg0 = defaultValues.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption16setDefaultValuesERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QCommandLineOption::Flags flags()
func (this *QCommandLineOption) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QCommandLineOption::Flags)
func (this *QCommandLineOption) SetFlags(aflags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption8setFlagsE6QFlagsINS_4FlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), aflags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHidden(_Bool)
func (this *QCommandLineOption) SetHidden(hidden bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption9setHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hidden)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHidden()
func (this *QCommandLineOption) IsHidden() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption8isHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QCommandLineOption__Flag = int

const QCommandLineOption__HiddenFromHelp QCommandLineOption__Flag = 1
const QCommandLineOption__ShortOptionStyle QCommandLineOption__Flag = 2

//  body block end