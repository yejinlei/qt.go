//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 49
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QTextCharFormat struct {
	*QTextFormat
}

func (this *QTextCharFormat) GetCthis() unsafe.Pointer {
	return this.QTextFormat.GetCthis()
}

// /usr/include/qt/QtGui/qtextformat.h:412
// index:0
// void QTextCharFormat()
func NewQTextCharFormat() *QTextCharFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCharFormatFromPointer(cthis)
	return gothis
}
func NewQTextCharFormatFromPointer(cthis unsafe.Pointer) *QTextCharFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextCharFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:557
// index:1
// void QTextCharFormat(const class QTextFormat &)
func NewQTextCharFormat_1(fmt unsafe.Pointer) *QTextCharFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, fmt)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCharFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:414
// index:0
// inline
// bool isValid()
func (this *QTextCharFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:420
// index:0
// void setFont(const class QFont &, enum QTextCharFormat::FontPropertiesInheritanceBehavior)
func (this *QTextCharFormat) SetFont(font unsafe.Pointer, behavior int) {
	// 0: (, font const QFont &, behavior QTextCharFormat::FontPropertiesInheritanceBehavior), (font, &behavior)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat7setFontERK5QFontNS_33FontPropertiesInheritanceBehaviorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font, &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:421
// index:1
// void setFont(const class QFont &)
func (this *QTextCharFormat) SetFont_1(font unsafe.Pointer) {
	// 1: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat7setFontERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:422
// index:0
// QFont font()
func (this *QTextCharFormat) Font() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat4fontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:424
// index:0
// inline
// void setFontFamily(const class QString &)
func (this *QTextCharFormat) SetFontFamily(family unsafe.Pointer) {
	// 0: (, family const QString &), (family)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontFamilyERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:426
// index:0
// inline
// QString fontFamily()
func (this *QTextCharFormat) FontFamily() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontFamilyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:429
// index:0
// inline
// void setFontPointSize(qreal)
func (this *QTextCharFormat) SetFontPointSize(size float64) {
	// 0: (, size qreal), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontPointSizeEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:431
// index:0
// inline
// qreal fontPointSize()
func (this *QTextCharFormat) FontPointSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontPointSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:434
// index:0
// inline
// void setFontWeight(int)
func (this *QTextCharFormat) SetFontWeight(weight int) {
	// 0: (, weight int), (&weight)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontWeightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &weight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:436
// index:0
// inline
// int fontWeight()
func (this *QTextCharFormat) FontWeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontWeightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:438
// index:0
// inline
// void setFontItalic(_Bool)
func (this *QTextCharFormat) SetFontItalic(italic bool) {
	// 0: (, italic bool), (&italic)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontItalicEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &italic)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:440
// index:0
// inline
// bool fontItalic()
func (this *QTextCharFormat) FontItalic() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontItalicEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:442
// index:0
// inline
// void setFontCapitalization(class QFont::Capitalization)
func (this *QTextCharFormat) SetFontCapitalization(capitalization int) {
	// 0: (, capitalization QFont::Capitalization), (&capitalization)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat21setFontCapitalizationEN5QFont14CapitalizationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &capitalization)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:444
// index:0
// inline
// QFont::Capitalization fontCapitalization()
func (this *QTextCharFormat) FontCapitalization() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat18fontCapitalizationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:446
// index:0
// inline
// void setFontLetterSpacingType(class QFont::SpacingType)
func (this *QTextCharFormat) SetFontLetterSpacingType(letterSpacingType int) {
	// 0: (, letterSpacingType QFont::SpacingType), (&letterSpacingType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat24setFontLetterSpacingTypeEN5QFont11SpacingTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &letterSpacingType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:448
// index:0
// inline
// QFont::SpacingType fontLetterSpacingType()
func (this *QTextCharFormat) FontLetterSpacingType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat21fontLetterSpacingTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:450
// index:0
// inline
// void setFontLetterSpacing(qreal)
func (this *QTextCharFormat) SetFontLetterSpacing(spacing float64) {
	// 0: (, spacing qreal), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat20setFontLetterSpacingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:452
// index:0
// inline
// qreal fontLetterSpacing()
func (this *QTextCharFormat) FontLetterSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat17fontLetterSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:454
// index:0
// inline
// void setFontWordSpacing(qreal)
func (this *QTextCharFormat) SetFontWordSpacing(spacing float64) {
	// 0: (, spacing qreal), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat18setFontWordSpacingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:456
// index:0
// inline
// qreal fontWordSpacing()
func (this *QTextCharFormat) FontWordSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat15fontWordSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:459
// index:0
// inline
// void setFontUnderline(_Bool)
func (this *QTextCharFormat) SetFontUnderline(underline bool) {
	// 0: (, underline bool), (&underline)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontUnderlineEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &underline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:461
// index:0
// bool fontUnderline()
func (this *QTextCharFormat) FontUnderline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontUnderlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:463
// index:0
// inline
// void setFontOverline(_Bool)
func (this *QTextCharFormat) SetFontOverline(overline bool) {
	// 0: (, overline bool), (&overline)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat15setFontOverlineEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &overline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:465
// index:0
// inline
// bool fontOverline()
func (this *QTextCharFormat) FontOverline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat12fontOverlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:468
// index:0
// inline
// void setFontStrikeOut(_Bool)
func (this *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	// 0: (, strikeOut bool), (&strikeOut)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontStrikeOutEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &strikeOut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:470
// index:0
// inline
// bool fontStrikeOut()
func (this *QTextCharFormat) FontStrikeOut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontStrikeOutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:473
// index:0
// inline
// void setUnderlineColor(const class QColor &)
func (this *QTextCharFormat) SetUnderlineColor(color unsafe.Pointer) {
	// 0: (, color const QColor &), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat17setUnderlineColorERK6QColor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:475
// index:0
// inline
// QColor underlineColor()
func (this *QTextCharFormat) UnderlineColor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat14underlineColorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:478
// index:0
// inline
// void setFontFixedPitch(_Bool)
func (this *QTextCharFormat) SetFontFixedPitch(fixedPitch bool) {
	// 0: (, fixedPitch bool), (&fixedPitch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat17setFontFixedPitchEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &fixedPitch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:480
// index:0
// inline
// bool fontFixedPitch()
func (this *QTextCharFormat) FontFixedPitch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat14fontFixedPitchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:483
// index:0
// inline
// void setFontStretch(int)
func (this *QTextCharFormat) SetFontStretch(factor int) {
	// 0: (, factor int), (&factor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setFontStretchEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &factor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:485
// index:0
// inline
// int fontStretch()
func (this *QTextCharFormat) FontStretch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat11fontStretchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:488
// index:0
// inline
// void setFontStyleHint(class QFont::StyleHint, class QFont::StyleStrategy)
func (this *QTextCharFormat) SetFontStyleHint(hint int, strategy int) {
	// 0: (, hint QFont::StyleHint, strategy QFont::StyleStrategy), (&hint, &strategy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontStyleHintEN5QFont9StyleHintENS0_13StyleStrategyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hint, &strategy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:490
// index:0
// inline
// void setFontStyleStrategy(class QFont::StyleStrategy)
func (this *QTextCharFormat) SetFontStyleStrategy(strategy int) {
	// 0: (, strategy QFont::StyleStrategy), (&strategy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat20setFontStyleStrategyEN5QFont13StyleStrategyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &strategy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:492
// index:0
// inline
// QFont::StyleHint fontStyleHint()
func (this *QTextCharFormat) FontStyleHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontStyleHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:494
// index:0
// inline
// QFont::StyleStrategy fontStyleStrategy()
func (this *QTextCharFormat) FontStyleStrategy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat17fontStyleStrategyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:497
// index:0
// inline
// void setFontHintingPreference(class QFont::HintingPreference)
func (this *QTextCharFormat) SetFontHintingPreference(hintingPreference int) {
	// 0: (, hintingPreference QFont::HintingPreference), (&hintingPreference)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat24setFontHintingPreferenceEN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:502
// index:0
// inline
// QFont::HintingPreference fontHintingPreference()
func (this *QTextCharFormat) FontHintingPreference() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat21fontHintingPreferenceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:507
// index:0
// inline
// void setFontKerning(_Bool)
func (this *QTextCharFormat) SetFontKerning(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setFontKerningEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:509
// index:0
// inline
// bool fontKerning()
func (this *QTextCharFormat) FontKerning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat11fontKerningEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:512
// index:0
// void setUnderlineStyle(enum QTextCharFormat::UnderlineStyle)
func (this *QTextCharFormat) SetUnderlineStyle(style int) {
	// 0: (, style QTextCharFormat::UnderlineStyle), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat17setUnderlineStyleENS_14UnderlineStyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:513
// index:0
// inline
// QTextCharFormat::UnderlineStyle underlineStyle()
func (this *QTextCharFormat) UnderlineStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat14underlineStyleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:516
// index:0
// inline
// void setVerticalAlignment(enum QTextCharFormat::VerticalAlignment)
func (this *QTextCharFormat) SetVerticalAlignment(alignment int) {
	// 0: (, alignment QTextCharFormat::VerticalAlignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat20setVerticalAlignmentENS_17VerticalAlignmentE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:518
// index:0
// inline
// QTextCharFormat::VerticalAlignment verticalAlignment()
func (this *QTextCharFormat) VerticalAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat17verticalAlignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:521
// index:0
// inline
// void setTextOutline(const class QPen &)
func (this *QTextCharFormat) SetTextOutline(pen unsafe.Pointer) {
	// 0: (, pen const QPen &), (pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setTextOutlineERK4QPen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:523
// index:0
// inline
// QPen textOutline()
func (this *QTextCharFormat) TextOutline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat11textOutlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:526
// index:0
// inline
// void setToolTip(const class QString &)
func (this *QTextCharFormat) SetToolTip(tip unsafe.Pointer) {
	// 0: (, tip const QString &), (tip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), tip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:528
// index:0
// inline
// QString toolTip()
func (this *QTextCharFormat) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat7toolTipEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:531
// index:0
// inline
// void setAnchor(_Bool)
func (this *QTextCharFormat) SetAnchor(anchor bool) {
	// 0: (, anchor bool), (&anchor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat9setAnchorEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:533
// index:0
// inline
// bool isAnchor()
func (this *QTextCharFormat) IsAnchor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat8isAnchorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:536
// index:0
// inline
// void setAnchorHref(const class QString &)
func (this *QTextCharFormat) SetAnchorHref(value unsafe.Pointer) {
	// 0: (, value const QString &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setAnchorHrefERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:538
// index:0
// inline
// QString anchorHref()
func (this *QTextCharFormat) AnchorHref() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10anchorHrefEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:541
// index:0
// inline
// void setAnchorName(const class QString &)
func (this *QTextCharFormat) SetAnchorName(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setAnchorNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:543
// index:0
// QString anchorName()
func (this *QTextCharFormat) AnchorName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10anchorNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:545
// index:0
// inline
// void setAnchorNames(const class QStringList &)
func (this *QTextCharFormat) SetAnchorNames(names unsafe.Pointer) {
	// 0: (, names const QStringList &), (names)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setAnchorNamesERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), names)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:547
// index:0
// QStringList anchorNames()
func (this *QTextCharFormat) AnchorNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat11anchorNamesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:549
// index:0
// inline
// void setTableCellRowSpan(int)
func (this *QTextCharFormat) SetTableCellRowSpan(tableCellRowSpan int) {
	// 0: (, tableCellRowSpan int), (&tableCellRowSpan)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat19setTableCellRowSpanEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &tableCellRowSpan)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:550
// index:0
// inline
// int tableCellRowSpan()
func (this *QTextCharFormat) TableCellRowSpan() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat16tableCellRowSpanEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:552
// index:0
// inline
// void setTableCellColumnSpan(int)
func (this *QTextCharFormat) SetTableCellColumnSpan(tableCellColumnSpan int) {
	// 0: (, tableCellColumnSpan int), (&tableCellColumnSpan)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat22setTableCellColumnSpanEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &tableCellColumnSpan)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:553
// index:0
// inline
// int tableCellColumnSpan()
func (this *QTextCharFormat) TableCellColumnSpan() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat19tableCellColumnSpanEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end