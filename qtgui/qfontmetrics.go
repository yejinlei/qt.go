package qtgui

// /usr/include/qt/QtGui/qfontmetrics.h
// #include <qfontmetrics.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QFontMetrics struct {
	*qtrt.CObject
}

func (this *QFontMetrics) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontMetrics) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontMetricsFromPointer(cthis unsafe.Pointer) *QFontMetrics {
	return &QFontMetrics{&qtrt.CObject{cthis}}
}
func (*QFontMetrics) NewFromPointer(cthis unsafe.Pointer) *QFontMetrics {
	return NewQFontMetricsFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontmetrics.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontMetrics(const QFont &)
func NewQFontMetrics(arg0 *QFont) *QFontMetrics {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFont", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetrics)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFontMetrics(const QFont &, QPaintDevice *)
func NewQFontMetrics_1(arg0 *QFont, pd *QPaintDevice /*777 QPaintDevice **/) *QFontMetrics {
	var convArg0 = arg0.GetCthis()
	var convArg1 = pd.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsC2ERK5QFontP12QPaintDevice", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetrics)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFontMetrics()
func DeleteQFontMetrics(this *QFontMetrics) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetricsD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qfontmetrics.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFontMetrics &)
func (this *QFontMetrics) Swap(other *QFontMetrics) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QFontMetrics4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int ascent()
func (this *QFontMetrics) Ascent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6ascentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int capHeight()
func (this *QFontMetrics) CapHeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9capHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int descent()
func (this *QFontMetrics) Descent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7descentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int height()
func (this *QFontMetrics) Height() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int leading()
func (this *QFontMetrics) Leading() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7leadingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineSpacing()
func (this *QFontMetrics) LineSpacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11lineSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int minLeftBearing()
func (this *QFontMetrics) MinLeftBearing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics14minLeftBearingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] int minRightBearing()
func (this *QFontMetrics) MinRightBearing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics15minRightBearingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxWidth()
func (this *QFontMetrics) MaxWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics8maxWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int xHeight()
func (this *QFontMetrics) XHeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics7xHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int averageCharWidth()
func (this *QFontMetrics) AverageCharWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics16averageCharWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFont(QChar)
func (this *QFontMetrics) InFont(arg0 *qtcore.QChar /*123*/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics6inFontE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFontUcs4(uint)
func (this *QFontMetrics) InFontUcs4(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics10inFontUcs4Ej", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int leftBearing(QChar)
func (this *QFontMetrics) LeftBearing(arg0 *qtcore.QChar /*123*/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11leftBearingE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int rightBearing(QChar)
func (this *QFontMetrics) RightBearing(arg0 *qtcore.QChar /*123*/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12rightBearingE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int width(const QString &, int)
func (this *QFontMetrics) Width(arg0 *qtcore.QString, len int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:94
// index:1
// Public Visibility=Default Availability=Available
// [4] int width(const QString &, int, int)
func (this *QFontMetrics) Width_1(arg0 *qtcore.QString, len int, flags int) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthERK7QStringii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:96
// index:2
// Public Visibility=Default Availability=Available
// [4] int width(QChar)
func (this *QFontMetrics) Width_2(arg0 *qtcore.QChar /*123*/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics5widthE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] int charWidth(const QString &, int)
func (this *QFontMetrics) CharWidth(str *qtcore.QString, pos int) int {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9charWidthERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pos)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:101
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(QChar)
func (this *QFontMetrics) BoundingRect(arg0 *qtcore.QChar /*123*/) *qtcore.QRect /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:103
// index:1
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QString &)
func (this *QFontMetrics) BoundingRect_1(text *qtcore.QString) *qtcore.QRect /*123*/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:104
// index:2
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect(const QRect &, int, const QString &, int, int *)
func (this *QFontMetrics) BoundingRect_2(r *qtcore.QRect, flags int, text *qtcore.QString, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QRect /*123*/ {
	var convArg0 = r.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, &tabarray)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:105
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QRect boundingRect(int, int, int, int, int, const QString &, int, int *)
func (this *QFontMetrics) BoundingRect_3(x int, y int, w int, h int, flags int, text *qtcore.QString, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QRect /*123*/ {
	var convArg5 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, flags, convArg5, tabstops, &tabarray)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size(int, const QString &, int, int *)
func (this *QFontMetrics) Size(flags int, str *qtcore.QString, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QSize /*123*/ {
	var convArg1 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics4sizeEiRK7QStringiPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, &tabarray)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:110
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect tightBoundingRect(const QString &)
func (this *QFontMetrics) TightBoundingRect(text *qtcore.QString) *qtcore.QRect /*123*/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics17tightBoundingRectERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elidedText(const QString &, Qt::TextElideMode, int, int)
func (this *QFontMetrics) ElidedText(text *qtcore.QString, mode int, width int, flags int) *qtcore.QString /*123*/ {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics10elidedTextERK7QStringN2Qt13TextElideModeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, width, flags)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int underlinePos()
func (this *QFontMetrics) UnderlinePos() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12underlinePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] int overlinePos()
func (this *QFontMetrics) OverlinePos() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics11overlinePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] int strikeOutPos()
func (this *QFontMetrics) StrikeOutPos() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics12strikeOutPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int lineWidth()
func (this *QFontMetrics) LineWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QFontMetrics9lineWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end