//  header block begin
// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
type QXmlStreamAttributes struct {
	*qtrt.CObject
}

func (this *QXmlStreamAttributes) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qxmlstream.h:164
// index:0
// inline
// void QXmlStreamAttributes()
func NewQXmlStreamAttributes() *QXmlStreamAttributes {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QXmlStreamAttributesC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributesFromPointer(cthis)
	return gothis
}
func NewQXmlStreamAttributesFromPointer(cthis unsafe.Pointer) *QXmlStreamAttributes {
	return &QXmlStreamAttributes{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:165
// index:0
// QStringRef value(const class QString &, const class QString &)
func (this *QXmlStreamAttributes) Value(namespaceUri unsafe.Pointer, name unsafe.Pointer) {
	// 0: (, namespaceUri const QString &, name const QString &), (namespaceUri, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:166
// index:1
// QStringRef value(const class QString &, class QLatin1String)
func (this *QXmlStreamAttributes) Value_1(namespaceUri unsafe.Pointer, name unsafe.Pointer) {
	// 1: (, namespaceUri const QString &, name QLatin1String), (namespaceUri, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QString13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:167
// index:2
// QStringRef value(class QLatin1String, class QLatin1String)
func (this *QXmlStreamAttributes) Value_2(namespaceUri unsafe.Pointer, name unsafe.Pointer) {
	// 2: (, namespaceUri QLatin1String, name QLatin1String), (namespaceUri, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueE13QLatin1StringS0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:168
// index:3
// QStringRef value(const class QString &)
func (this *QXmlStreamAttributes) Value_3(qualifiedName unsafe.Pointer) {
	// 3: (, qualifiedName const QString &), (qualifiedName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:169
// index:4
// QStringRef value(class QLatin1String)
func (this *QXmlStreamAttributes) Value_4(qualifiedName unsafe.Pointer) {
	// 4: (, qualifiedName QLatin1String), (qualifiedName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes5valueE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:173
// index:0
// inline
// bool hasAttribute(const class QString &)
func (this *QXmlStreamAttributes) HasAttribute(qualifiedName unsafe.Pointer) {
	// 0: (, qualifiedName const QString &), (qualifiedName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:178
// index:1
// inline
// bool hasAttribute(class QLatin1String)
func (this *QXmlStreamAttributes) HasAttribute_1(qualifiedName unsafe.Pointer) {
	// 1: (, qualifiedName QLatin1String), (qualifiedName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.GetCthis(), qualifiedName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:183
// index:2
// inline
// bool hasAttribute(const class QString &, const class QString &)
func (this *QXmlStreamAttributes) HasAttribute_2(namespaceUri unsafe.Pointer, name unsafe.Pointer) {
	// 2: (, namespaceUri const QString &, name const QString &), (namespaceUri, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QXmlStreamAttributes12hasAttributeERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), namespaceUri, name)
	gopp.ErrPrint(err, rv)
}

//  body block end