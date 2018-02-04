package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsEllipseItem) InheritSupportsExtension(f func(extension int) bool) {
	ffiqt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsEllipseItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant)) {
	ffiqt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
func (this *QGraphicsEllipseItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "extension", f)
}

type QGraphicsEllipseItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsEllipseItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractGraphicsShapeItem.GetCthis()
	}
}
func (this *QGraphicsEllipseItem) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractGraphicsShapeItem = NewQAbstractGraphicsShapeItemFromPointer(cthis)
}
func NewQGraphicsEllipseItemFromPointer(cthis unsafe.Pointer) *QGraphicsEllipseItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsEllipseItem{bcthis0}
}
func (*QGraphicsEllipseItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsEllipseItem {
	return NewQGraphicsEllipseItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:699
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEllipseItem(QGraphicsItem *)
func NewQGraphicsEllipseItem(parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsEllipseItem {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsEllipseItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:700
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEllipseItem(const QRectF &, QGraphicsItem *)
func NewQGraphicsEllipseItem_1(rect *qtcore.QRectF, parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsEllipseItem {
	var convArg0 = rect.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2ERK6QRectFP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsEllipseItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:701
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsEllipseItem(qreal, qreal, qreal, qreal, QGraphicsItem *)
func NewQGraphicsEllipseItem_2(x float64, y float64, w float64, h float64, parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsEllipseItem {
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemC2EddddP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, x, y, w, h, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsEllipseItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsEllipseItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:702
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsEllipseItem()
func DeleteQGraphicsEllipseItem(this *QGraphicsEllipseItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItemD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:704
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect()
func (this *QGraphicsEllipseItem) Rect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:705
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)
func (this *QGraphicsEllipseItem) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem7setRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:706
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsEllipseItem) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:708
// index:0
// Public Visibility=Default Availability=Available
// [4] int startAngle()
func (this *QGraphicsEllipseItem) StartAngle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem10startAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:709
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartAngle(int)
func (this *QGraphicsEllipseItem) SetStartAngle(angle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem13setStartAngleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:711
// index:0
// Public Visibility=Default Availability=Available
// [4] int spanAngle()
func (this *QGraphicsEllipseItem) SpanAngle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem9spanAngleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:712
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpanAngle(int)
func (this *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem12setSpanAngleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:714
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsEllipseItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:715
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape()
func (this *QGraphicsEllipseItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:716
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QGraphicsEllipseItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:718
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsEllipseItem) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:720
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QGraphicsEllipseItem) IsObscuredBy(item *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:721
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QGraphicsEllipseItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:724
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsEllipseItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:727
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsEllipseItem) SupportsExtension(extension int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:728
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsEllipseItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsEllipseItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:729
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &)
func (this *QGraphicsEllipseItem) Extension(variant *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsEllipseItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QGraphicsEllipseItem__ = int

const QGraphicsEllipseItem__Type QGraphicsEllipseItem__ = 4

//  body block end