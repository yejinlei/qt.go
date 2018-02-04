package qtwidgets

// /usr/include/qt/QtWidgets/qtabwidget.h
// #include <qtabwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 81
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
// void tabInserted(int)
func (this *QTabWidget) InheritTabInserted(f func(index int)) {
	ffiqt.SetAllInheritCallback(this, "tabInserted", f)
}

// void tabRemoved(int)
func (this *QTabWidget) InheritTabRemoved(f func(index int)) {
	ffiqt.SetAllInheritCallback(this, "tabRemoved", f)
}

// void showEvent(class QShowEvent *)
func (this *QTabWidget) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "showEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QTabWidget) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QTabWidget) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QTabWidget) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void setTabBar(class QTabBar *)
func (this *QTabWidget) InheritSetTabBar(f func(arg0 *QTabBar /*777 QTabBar **/)) {
	ffiqt.SetAllInheritCallback(this, "setTabBar", f)
}

// void changeEvent(class QEvent *)
func (this *QTabWidget) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// bool event(class QEvent *)
func (this *QTabWidget) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// void initStyleOption(class QStyleOptionTabWidgetFrame *)
func (this *QTabWidget) InheritInitStyleOption(f func(option *QStyleOptionTabWidgetFrame /*777 QStyleOptionTabWidgetFrame **/)) {
	ffiqt.SetAllInheritCallback(this, "initStyleOption", f)
}

type QTabWidget struct {
	*QWidget
}

func (this *QTabWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QTabWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQTabWidgetFromPointer(cthis unsafe.Pointer) *QTabWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QTabWidget{bcthis0}
}
func (*QTabWidget) NewFromPointer(cthis unsafe.Pointer) *QTabWidget {
	return NewQTabWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTabWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTabWidget(QWidget *)
func NewQTabWidget(parent *QWidget /*777 QWidget **/) *QTabWidget {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidgetC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTabWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qtabwidget.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTabWidget()
func DeleteQTabWidget(this *QTabWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidgetD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int addTab(QWidget *, const QString &)
func (this *QTabWidget) AddTab(widget *QWidget /*777 QWidget **/, arg1 *qtcore.QString) int {
	var convArg0 = widget.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:75
// index:1
// Public Visibility=Default Availability=Available
// [4] int addTab(QWidget *, const QIcon &, const QString &)
func (this *QTabWidget) AddTab_1(widget *QWidget /*777 QWidget **/, icon *qtgui.QIcon, label *qtcore.QString) int {
	var convArg0 = widget.GetCthis()
	var convArg1 = icon.GetCthis()
	var convArg2 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, QWidget *, const QString &)
func (this *QTabWidget) InsertTab(index int, widget *QWidget /*777 QWidget **/, arg2 *qtcore.QString) int {
	var convArg1 = widget.GetCthis()
	var convArg2 = arg2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:78
// index:1
// Public Visibility=Default Availability=Available
// [4] int insertTab(int, QWidget *, const QIcon &, const QString &)
func (this *QTabWidget) InsertTab_1(index int, widget *QWidget /*777 QWidget **/, icon *qtgui.QIcon, label *qtcore.QString) int {
	var convArg1 = widget.GetCthis()
	var convArg2 = icon.GetCthis()
	var convArg3 = label.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeTab(int)
func (this *QTabWidget) RemoveTab(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9removeTabEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTabEnabled(int)
func (this *QTabWidget) IsTabEnabled(index int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12isTabEnabledEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabEnabled(int, _Bool)
func (this *QTabWidget) SetTabEnabled(index int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13setTabEnabledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabText(int)
func (this *QTabWidget) TabText(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7tabTextEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabText(int, const QString &)
func (this *QTabWidget) SetTabText(index int, arg1 *qtcore.QString) {
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setTabTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon tabIcon(int)
func (this *QTabWidget) TabIcon(index int) *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7tabIconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabIcon(int, const QIcon &)
func (this *QTabWidget) SetTabIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setTabIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabToolTip(int, const QString &)
func (this *QTabWidget) SetTabToolTip(index int, tip *qtcore.QString) {
	var convArg1 = tip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13setTabToolTipEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabToolTip(int)
func (this *QTabWidget) TabToolTip(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget10tabToolTipEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabWhatsThis(int, const QString &)
func (this *QTabWidget) SetTabWhatsThis(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setTabWhatsThisEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString tabWhatsThis(int)
func (this *QTabWidget) TabWhatsThis(index int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12tabWhatsThisEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex()
func (this *QTabWidget) CurrentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * currentWidget()
func (this *QTabWidget) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget13currentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int)
func (this *QTabWidget) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *)
func (this *QTabWidget) IndexOf(widget *QWidget /*777 QWidget **/) int {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QTabWidget) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabPosition tabPosition()
func (this *QTabWidget) TabPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget11tabPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabPosition(enum QTabWidget::TabPosition)
func (this *QTabWidget) SetTabPosition(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget14setTabPositionENS_11TabPositionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabsClosable()
func (this *QTabWidget) TabsClosable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12tabsClosableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabsClosable(_Bool)
func (this *QTabWidget) SetTabsClosable(closeable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setTabsClosableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), closeable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMovable()
func (this *QTabWidget) IsMovable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget9isMovableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMovable(_Bool)
func (this *QTabWidget) SetMovable(movable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10setMovableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), movable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] QTabWidget::TabShape tabShape()
func (this *QTabWidget) TabShape() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8tabShapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabShape(enum QTabWidget::TabShape)
func (this *QTabWidget) SetTabShape(s int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11setTabShapeENS_8TabShapeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:123
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QTabWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QTabWidget) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QTabWidget) HeightForWidth(width int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtabwidget.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QTabWidget) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *, Qt::Corner)
func (this *QTabWidget) SetCornerWidget(w *QWidget /*777 QWidget **/, corner int) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setCornerWidgetEP7QWidgetN2Qt6CornerE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, corner)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget(Qt::Corner)
func (this *QTabWidget) CornerWidget(corner int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12cornerWidgetEN2Qt6CornerE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), corner)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextElideMode elideMode()
func (this *QTabWidget) ElideMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget9elideModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setElideMode(Qt::TextElideMode)
func (this *QTabWidget) SetElideMode(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget12setElideModeEN2Qt13TextElideModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize iconSize()
func (this *QTabWidget) IconSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget8iconSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconSize(const QSize &)
func (this *QTabWidget) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool usesScrollButtons()
func (this *QTabWidget) UsesScrollButtons() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget17usesScrollButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUsesScrollButtons(_Bool)
func (this *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget20setUsesScrollButtonsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), useButtons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool documentMode()
func (this *QTabWidget) DocumentMode() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget12documentModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDocumentMode(_Bool)
func (this *QTabWidget) SetDocumentMode(set bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setDocumentModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), set)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tabBarAutoHide()
func (this *QTabWidget) TabBarAutoHide() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget14tabBarAutoHideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabBarAutoHide(_Bool)
func (this *QTabWidget) SetTabBarAutoHide(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget17setTabBarAutoHideEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QTabWidget) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:148
// index:0
// Public Visibility=Default Availability=Available
// [8] QTabBar * tabBar()
func (this *QTabWidget) TabBar() *QTabBar /*777 QTabBar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget6tabBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTabBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qtabwidget.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)
func (this *QTabWidget) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)
func (this *QTabWidget) SetCurrentWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)
func (this *QTabWidget) CurrentChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabCloseRequested(int)
func (this *QTabWidget) TabCloseRequested(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget17tabCloseRequestedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarClicked(int)
func (this *QTabWidget) TabBarClicked(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13tabBarClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabBarDoubleClicked(int)
func (this *QTabWidget) TabBarDoubleClicked(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget19tabBarDoubleClickedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:161
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabInserted(int)
func (this *QTabWidget) TabInserted(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11tabInsertedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:162
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabRemoved(int)
func (this *QTabWidget) TabRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10tabRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:164
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QTabWidget) ShowEvent(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:165
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QTabWidget) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:166
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QTabWidget) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:167
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QTabWidget) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:168
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setTabBar(QTabBar *)
func (this *QTabWidget) SetTabBar(arg0 *QTabBar /*777 QTabBar **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget9setTabBarEP7QTabBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:169
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QTabWidget) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtabwidget.h:170
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QTabWidget) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTabWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qtabwidget.h:171
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionTabWidgetFrame *)
func (this *QTabWidget) InitStyleOption(option *QStyleOptionTabWidgetFrame /*777 QStyleOptionTabWidgetFrame **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QTabWidget__TabPosition = int

const QTabWidget__North QTabWidget__TabPosition = 0
const QTabWidget__South QTabWidget__TabPosition = 1
const QTabWidget__West QTabWidget__TabPosition = 2
const QTabWidget__East QTabWidget__TabPosition = 3

type QTabWidget__TabShape = int

const QTabWidget__Rounded QTabWidget__TabShape = 0
const QTabWidget__Triangular QTabWidget__TabShape = 1

//  body block end