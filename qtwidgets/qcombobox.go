// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qcombobox.h
// #include <qcombobox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QComboBox struct {
	*QWidget
}
type QComboBox_ITF interface {
	QWidget_ITF
	QComboBox_PTR() *QComboBox
}

func (ptr *QComboBox) QComboBox_PTR() *QComboBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QComboBoxFromptr(cthis unsafe.Pointer) *QComboBox {
	bcthis0 := QWidgetFromptr(cthis)
	return &QComboBox{bcthis0}
}
func (*QComboBox) Fromptr(cthis unsafe.Pointer) *QComboBox {
	return QComboBoxFromptr(cthis)
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QComboBox(QWidget *)

/*
 */
func (*QComboBox) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QComboBox {
	return NewQComboBox(parent)
}
func NewQComboBox(parent QWidget_ITF /*777 QWidget **/) *QComboBox {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(400444357, "_ZN9QComboBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QComboBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QComboBox(QWidget *)

/*
 */
func (*QComboBox) NewForInheritp() *QComboBox {
	return NewQComboBoxp()
}
func NewQComboBoxp() *QComboBox {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	cthis := qtrt.Malloc(48)
	rv, err := qtrt.Qtcc1(400444357, "_ZN9QComboBoxC2EP7QWidget", qtrt.FFITY_POINTER, cthis, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := QComboBoxFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QComboBox")
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:91
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maxVisibleItems() const

/*
 */
func (this *QComboBox) MaxVisibleItems() int {
	rv, err := qtrt.Qtcc1(2452855331, "_ZNK9QComboBox15maxVisibleItemsEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaxVisibleItems(int)

/*
 */
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	rv, err := qtrt.Qtcc1(2605451080, "_ZN9QComboBox18setMaxVisibleItemsEi", qtrt.FFITY_POINTER, this.GetCthis(), maxItems)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:94
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int count() const

/*
 */
func (this *QComboBox) Count() int {
	rv, err := qtrt.Qtcc1(484839320, "_ZNK9QComboBox5countEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:95
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setMaxCount(int)

/*
 */
func (this *QComboBox) SetMaxCount(max int) {
	rv, err := qtrt.Qtcc1(804866272, "_ZN9QComboBox11setMaxCountEi", qtrt.FFITY_POINTER, this.GetCthis(), max)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int maxCount() const

/*
 */
func (this *QComboBox) MaxCount() int {
	rv, err := qtrt.Qtcc1(3494713860, "_ZNK9QComboBox8maxCountEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:111
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool duplicatesEnabled() const

/*
 */
func (this *QComboBox) DuplicatesEnabled() bool {
	rv, err := qtrt.Qtcc1(3571501781, "_ZNK9QComboBox17duplicatesEnabledEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcombobox.h:112
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setDuplicatesEnabled(bool)

/*
 */
func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	rv, err := qtrt.Qtcc1(1134441101, "_ZN9QComboBox20setDuplicatesEnabledEb", qtrt.FFITY_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:159
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setEditable(bool)

/*
 */
func (this *QComboBox) SetEditable(editable bool) {
	rv, err := qtrt.Qtcc1(1670976916, "_ZN9QComboBox11setEditableEb", qtrt.FFITY_POINTER, this.GetCthis(), editable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int currentIndex() const

/*
 */
func (this *QComboBox) CurrentIndex() int {
	rv, err := qtrt.Qtcc1(3220285656, "_ZNK9QComboBox12currentIndexEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString currentText() const

/*
 */
func (this *QComboBox) CurrentText() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(1212789540, "_ZNK9QComboBox11currentTextEv", qtrt.FFITY_POINTER, sretobj, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcombobox.h:188
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString itemText(int) const

/*
 */
func (this *QComboBox) ItemText(index int) string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc1(2984951329, "_ZNK9QComboBox8itemTextEi", qtrt.FFITY_POINTER, sretobj, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void addItem(const QString &, const QVariant &)

/*
 */
func (this *QComboBox) AddItem(text string, userData qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg1 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(931319215, "_ZN9QComboBox7addItemERK7QStringRK8QVariant", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void addItem(const QString &, const QVariant &)

/*
 */
func (this *QComboBox) AddItemp(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg1 = qtcore.NewQVariant()
	rv, err := qtrt.Qtcc1(931319215, "_ZN9QComboBox7addItemERK7QStringRK8QVariant", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:198
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &, const QVariant &)

/*
 */
func (this *QComboBox) InsertItem(index int, text string, userData qtcore.QVariant_ITF) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if userData != nil && userData.QVariant_PTR() != nil {
		convArg2 = userData.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc1(3465040126, "_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:198
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void insertItem(int, const QString &, const QVariant &)

/*
 */
func (this *QComboBox) InsertItemp(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg2 = qtcore.NewQVariant()
	rv, err := qtrt.Qtcc1(3465040126, "_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:204
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void removeItem(int)

/*
 */
func (this *QComboBox) RemoveItem(index int) {
	rv, err := qtrt.Qtcc1(3369550016, "_ZN9QComboBox10removeItemEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:206
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setItemText(int, const QString &)

/*
 */
func (this *QComboBox) SetItemText(index int, text string) {
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.Qtcc1(198468933, "_ZN9QComboBox11setItemTextEiRK7QString", qtrt.FFITY_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:210
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QAbstractItemView * view() const

/*
 */
func (this *QComboBox) View() *QAbstractItemView /*777 QAbstractItemView **/ {
	rv, err := qtrt.Qtcc1(3310046969, "_ZNK9QComboBox4viewEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ QAbstractItemViewFromptr(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcombobox.h:224
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clear()

/*
 */
func (this *QComboBox) Clear() {
	rv, err := qtrt.Qtcc1(3991980318, "_ZN9QComboBox5clearEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:225
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void clearEditText()

/*
 */
func (this *QComboBox) ClearEditText() {
	rv, err := qtrt.Qtcc1(2293827830, "_ZN9QComboBox13clearEditTextEv", qtrt.FFITY_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:226
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setEditText(const QString &)

/*
 */
func (this *QComboBox) SetEditText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(3736805172, "_ZN9QComboBox11setEditTextERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:227
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)

/*
 */
func (this *QComboBox) SetCurrentIndex(index int) {
	rv, err := qtrt.Qtcc1(934993843, "_ZN9QComboBox15setCurrentIndexEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:228
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCurrentText(const QString &)

/*
 */
func (this *QComboBox) SetCurrentText(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(3375923067, "_ZN9QComboBox14setCurrentTextERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:231
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void editTextChanged(const QString &)

/*
 */
func (this *QComboBox) EditTextChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(681558797, "_ZN9QComboBox15editTextChangedERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:232
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void activated(int)

/*
 */
func (this *QComboBox) Activated(index int) {
	rv, err := qtrt.Qtcc1(3913087343, "_ZN9QComboBox9activatedEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:233
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textActivated(const QString &)

/*
 */
func (this *QComboBox) TextActivated(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(1528827286, "_ZN9QComboBox13textActivatedERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:234
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void highlighted(int)

/*
 */
func (this *QComboBox) Highlighted(index int) {
	rv, err := qtrt.Qtcc1(1991475282, "_ZN9QComboBox11highlightedEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:235
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void textHighlighted(const QString &)

/*
 */
func (this *QComboBox) TextHighlighted(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(1845413289, "_ZN9QComboBox15textHighlightedERK7QString", qtrt.FFITY_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:236
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void currentIndexChanged(int)

/*
 */
func (this *QComboBox) CurrentIndexChanged(index int) {
	rv, err := qtrt.Qtcc1(3753237607, "_ZN9QComboBox19currentIndexChangedEi", qtrt.FFITY_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

func DeleteQComboBox(this *QComboBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN9QComboBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QComboBox__InsertPolicy = int

//
const QComboBox__NoInsert QComboBox__InsertPolicy = 0

//
const QComboBox__InsertAtTop QComboBox__InsertPolicy = 1

//
const QComboBox__InsertAtCurrent QComboBox__InsertPolicy = 2

//
const QComboBox__InsertAtBottom QComboBox__InsertPolicy = 3

//
const QComboBox__InsertAfterCurrent QComboBox__InsertPolicy = 4

//
const QComboBox__InsertBeforeCurrent QComboBox__InsertPolicy = 5

//
const QComboBox__InsertAlphabetically QComboBox__InsertPolicy = 6

func (this *QComboBox) InsertPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QComboBox_InsertPolicyItemName(val int) string {
	var nilthis *QComboBox
	return nilthis.InsertPolicyItemName(val)
}

/*


 */
type QComboBox__SizeAdjustPolicy = int

//
const QComboBox__AdjustToContents QComboBox__SizeAdjustPolicy = 0

//
const QComboBox__AdjustToContentsOnFirstShow QComboBox__SizeAdjustPolicy = 1

//
const QComboBox__AdjustToMinimumContentsLength QComboBox__SizeAdjustPolicy = 2

//
const QComboBox__AdjustToMinimumContentsLengthWithIcon QComboBox__SizeAdjustPolicy = 3

func (this *QComboBox) SizeAdjustPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QComboBox_SizeAdjustPolicyItemName(val int) string {
	var nilthis *QComboBox
	return nilthis.SizeAdjustPolicyItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10123() {
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
