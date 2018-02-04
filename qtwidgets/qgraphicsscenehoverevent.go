package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

type QGraphicsSceneHoverEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneHoverEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneHoverEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneHoverEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHoverEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneHoverEvent{bcthis0}
}
func (*QGraphicsSceneHoverEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHoverEvent {
	return NewQGraphicsSceneHoverEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneHoverEvent(enum QEvent::Type)
func NewQGraphicsSceneHoverEvent(type_ int) *QGraphicsSceneHoverEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneHoverEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneHoverEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:202
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneHoverEvent()
func DeleteQGraphicsSceneHoverEvent(this *QGraphicsSceneHoverEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:204
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos()
func (this *QGraphicsSceneHoverEvent) Pos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:207
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos()
func (this *QGraphicsSceneHoverEvent) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent8scenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScenePos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:210
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint screenPos()
func (this *QGraphicsSceneHoverEvent) ScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreenPos(const QPoint &)
func (this *QGraphicsSceneHoverEvent) SetScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:213
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastPos()
func (this *QGraphicsSceneHoverEvent) LastPos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent7lastPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastPos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetLastPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent10setLastPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:216
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF lastScenePos()
func (this *QGraphicsSceneHoverEvent) LastScenePos() *qtcore.QPointF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent12lastScenePosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastScenePos(const QPointF &)
func (this *QGraphicsSceneHoverEvent) SetLastScenePos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent15setLastScenePosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint lastScreenPos()
func (this *QGraphicsSceneHoverEvent) LastScreenPos() *qtcore.QPoint /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent13lastScreenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLastScreenPos(const QPoint &)
func (this *QGraphicsSceneHoverEvent) SetLastScreenPos(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent16setLastScreenPosERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:222
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneHoverEvent) Modifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifiers(Qt::KeyboardModifiers)
func (this *QGraphicsSceneHoverEvent) SetModifiers(modifiers int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent12setModifiersE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	gopp.ErrPrint(err, rv)
}

//  body block end