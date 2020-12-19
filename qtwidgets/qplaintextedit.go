// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qplaintextedit.h
// #include <qplaintextedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*
 */
// size 48
type QPlainTextEdit struct {
	*QAbstractScrollArea
}
type QPlainTextEdit_ITF interface {
	QAbstractScrollArea_ITF
	QPlainTextEdit_PTR() *QPlainTextEdit
}

func (ptr *QPlainTextEdit) QPlainTextEdit_PTR() *QPlainTextEdit { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QPlainTextEditFromptr(cthis unsafe.Pointer) *QPlainTextEdit {
	bcthis0 := QAbstractScrollAreaFromptr(cthis)
	return &QPlainTextEdit{bcthis0}
}
func (*QPlainTextEdit) Fromptr(cthis unsafe.Pointer) *QPlainTextEdit {
	return QPlainTextEditFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	return NewQPlainTextEdit(parent)
}
func NewQPlainTextEdit(parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(396516486, "_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInheritp() *QPlainTextEdit {
	return NewQPlainTextEditp()
}
func NewQPlainTextEditp() *QPlainTextEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(396516486, "_ZN14QPlainTextEditC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(const QString &, QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInherit1(text string, parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	return NewQPlainTextEdit1(text, parent)
}
func NewQPlainTextEdit1(text string, parent QWidget_ITF /*777 QWidget **/) *QPlainTextEdit {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2780232515, "_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:96
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPlainTextEdit(const QString &, QWidget *)

/*
 */
func (*QPlainTextEdit) NewForInherit1p(text string) *QPlainTextEdit {
	return NewQPlainTextEdit1p(text)
}
func NewQPlainTextEdit1p(text string) *QPlainTextEdit {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(2780232515, "_ZN14QPlainTextEditC2ERK7QStringP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := QPlainTextEditFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QPlainTextEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:102
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPlaceholderText(const QString &)

/*
 */
func (this *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	var tmpArg0 = qtcore.NewQString5(placeholderText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(3191527676, "_ZN14QPlainTextEdit18setPlaceholderTextERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:103
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString placeholderText() const

/*
 */
func (this *QPlainTextEdit) PlaceholderText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(388573047, "_ZNK14QPlainTextEdit15placeholderTextEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:109
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setReadOnly(bool)

/*
 */
func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	rv, err := qtrt.Qtcc1(3324855670, "_ZN14QPlainTextEdit11setReadOnlyEb", qtrt.FFITY_POINTER, this.GetCthis(), ro)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextInteractionFlags(Qt::TextInteractionFlags)

/*
 */
func (this *QPlainTextEdit) SetTextInteractionFlags(flags int) {
	rv, err := qtrt.Qtcc1(400835228, "_ZN14QPlainTextEdit23setTextInteractionFlagsE6QFlagsIN2Qt19TextInteractionFlagEE", qtrt.FFITY_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:143
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setBackgroundVisible(bool)

/*
 */
func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	rv, err := qtrt.Qtcc1(957632610, "_ZN14QPlainTextEdit20setBackgroundVisibleEb", qtrt.FFITY_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:157
// index:0
// Public inline Indirect Visibility=Default Availability=Available
// [8] QString toPlainText() const

/*
 */
func (this *QPlainTextEdit) ToPlainText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(4142553692, "_ZNK14QPlainTextEdit11toPlainTextEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:160
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void ensureCursorVisible()

/*
 */
func (this *QPlainTextEdit) EnsureCursorVisible() {
	rv, err := qtrt.Qtcc1(3070872959, "_ZN14QPlainTextEdit19ensureCursorVisibleEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:203
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPlainText(const QString &)

/*
 */
func (this *QPlainTextEdit) SetPlainText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(2634248857, "_ZN14QPlainTextEdit12setPlainTextERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qplaintextedit.h:228
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textChanged()

/*
 */
func (this *QPlainTextEdit) TextChanged() {
	rv, err := qtrt.Qtcc1(4250366701, "_ZN14QPlainTextEdit11textChangedEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQPlainTextEdit(this *QPlainTextEdit) {
	rv, err := qtrt.Qtcc1(0, "_ZN14QPlainTextEditD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QPlainTextEdit__LineWrapMode = int

//
const QPlainTextEdit__NoWrap QPlainTextEdit__LineWrapMode = 0

//
const QPlainTextEdit__WidgetWidth QPlainTextEdit__LineWrapMode = 1

func (this *QPlainTextEdit) LineWrapModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QPlainTextEdit_LineWrapModeItemName(val int) string {
	var nilthis *QPlainTextEdit
	return nilthis.LineWrapModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10143() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
