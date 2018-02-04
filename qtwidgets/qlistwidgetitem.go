package qtwidgets

// /usr/include/qt/QtWidgets/qlistwidget.h
// #include <qlistwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
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

type QListWidgetItem struct {
	*qtrt.CObject
}

func (this *QListWidgetItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QListWidgetItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQListWidgetItemFromPointer(cthis unsafe.Pointer) *QListWidgetItem {
	return &QListWidgetItem{&qtrt.CObject{cthis}}
}
func (*QListWidgetItem) NewFromPointer(cthis unsafe.Pointer) *QListWidgetItem {
	return NewQListWidgetItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(QListWidget *, int)
func NewQListWidgetItem(view *QListWidget /*777 QListWidget **/, type_ int) *QListWidgetItem {
	var convArg0 = view.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemC2EP11QListWidgeti", ffiqt.FFI_TYPE_POINTER, convArg0, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QString &, QListWidget *, int)
func NewQListWidgetItem_1(text *qtcore.QString, view *QListWidget /*777 QListWidget **/, type_ int) *QListWidgetItem {
	var convArg0 = text.GetCthis()
	var convArg1 = view.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QListWidgetItem(const QIcon &, const QString &, QListWidget *, int)
func NewQListWidgetItem_2(icon *qtgui.QIcon, text *qtcore.QString, view *QListWidget /*777 QListWidget **/, type_ int) *QListWidgetItem {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = view.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQListWidgetItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qlistwidget.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QListWidgetItem()
func DeleteQListWidgetItem(this *QListWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QListWidgetItem * clone()
func (this *QListWidgetItem) Clone() *QListWidgetItem /*777 QListWidgetItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem5cloneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQListWidgetItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QListWidget * listWidget()
func (this *QListWidgetItem) ListWidget() *QListWidget /*777 QListWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10listWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQListWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSelected(_Bool)
func (this *QListWidgetItem) SetSelected(select_ bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem11setSelectedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSelected()
func (this *QListWidgetItem) IsSelected() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10isSelectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHidden(_Bool)
func (this *QListWidgetItem) SetHidden(hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem9setHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isHidden()
func (this *QListWidgetItem) IsHidden() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem8isHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlistwidget.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags()
func (this *QListWidgetItem) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::ItemFlags)
func (this *QListWidgetItem) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem8setFlagsE6QFlagsIN2Qt8ItemFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:84
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text()
func (this *QListWidgetItem) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QListWidgetItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QListWidgetItem) Icon() *qtgui.QIcon /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4iconEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:90
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QListWidgetItem) SetIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString statusTip()
func (this *QListWidgetItem) StatusTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem9statusTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)
func (this *QListWidgetItem) SetStatusTip(statusTip *qtcore.QString) {
	var convArg0 = statusTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem12setStatusTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QListWidgetItem) ToolTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QListWidgetItem) SetToolTip(toolTip *qtcore.QString) {
	var convArg0 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:103
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis()
func (this *QListWidgetItem) WhatsThis() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem9whatsThisEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)
func (this *QListWidgetItem) SetWhatsThis(whatsThis *qtcore.QString) {
	var convArg0 = whatsThis.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QFont font()
func (this *QListWidgetItem) Font() *qtgui.QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:110
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QListWidgetItem) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int textAlignment()
func (this *QListWidgetItem) TextAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem13textAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlistwidget.h:114
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextAlignment(int)
func (this *QListWidgetItem) SetTextAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem16setTextAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor backgroundColor()
func (this *QListWidgetItem) BackgroundColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem15backgroundColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:119
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void setBackgroundColor(const QColor &)
func (this *QListWidgetItem) SetBackgroundColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem18setBackgroundColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background()
func (this *QListWidgetItem) Background() *qtgui.QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:124
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)
func (this *QListWidgetItem) SetBackground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor textColor()
func (this *QListWidgetItem) TextColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem9textColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextColor(const QColor &)
func (this *QListWidgetItem) SetTextColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem12setTextColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground()
func (this *QListWidgetItem) Foreground() *qtgui.QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)
func (this *QListWidgetItem) SetForeground(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem13setForegroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:137
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::CheckState checkState()
func (this *QListWidgetItem) CheckState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem10checkStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:139
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCheckState(Qt::CheckState)
func (this *QListWidgetItem) SetCheckState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem13setCheckStateEN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QListWidgetItem) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSizeHint(const QSize &)
func (this *QListWidgetItem) SetSizeHint(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem11setSizeHintERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:147
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(int)
func (this *QListWidgetItem) Data(role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4dataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qlistwidget.h:148
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(int, const QVariant &)
func (this *QListWidgetItem) SetData(role int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem7setDataEiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:153
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void read(QDataStream &)
func (this *QListWidgetItem) Read(in *qtcore.QDataStream) {
	var convArg0 = in.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QListWidgetItem4readER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void write(QDataStream &)
func (this *QListWidgetItem) Write(out *qtcore.QDataStream) {
	var convArg0 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem5writeER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlistwidget.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int type()
func (this *QListWidgetItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QListWidgetItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QListWidgetItem__ItemType = int

const QListWidgetItem__Type QListWidgetItem__ItemType = 0
const QListWidgetItem__UserType QListWidgetItem__ItemType = 1000

//  body block end