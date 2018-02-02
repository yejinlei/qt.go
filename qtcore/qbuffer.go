package qtcore

// /usr/include/qt/QtCore/qbuffer.h
// #include <qbuffer.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 59
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
// void connectNotify(const class QMetaMethod &)
func (this *QBuffer) InheritConnectNotify(f func(arg0 *QMetaMethod)) {
	ffiqt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const class QMetaMethod &)
func (this *QBuffer) InheritDisconnectNotify(f func(arg0 *QMetaMethod)) {
	ffiqt.SetAllInheritCallback(this, "disconnectNotify", f)
}

// qint64 readData(char *, qint64)
func (this *QBuffer) InheritReadData(f func(data string, maxlen int64) int64) {
	ffiqt.SetAllInheritCallback(this, "readData", f)
}

// qint64 writeData(const char *, qint64)
func (this *QBuffer) InheritWriteData(f func(data string, len int64) int64) {
	ffiqt.SetAllInheritCallback(this, "writeData", f)
}

type QBuffer struct {
	*QIODevice
}

func (this *QBuffer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QBuffer) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = NewQIODeviceFromPointer(cthis)
}
func NewQBufferFromPointer(cthis unsafe.Pointer) *QBuffer {
	bcthis0 := NewQIODeviceFromPointer(cthis)
	return &QBuffer{bcthis0}
}
func (*QBuffer) NewFromPointer(cthis unsafe.Pointer) *QBuffer {
	return NewQBufferFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbuffer.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QBuffer) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBuffer(QObject *)
func NewQBuffer(parent *QObject /*777 QObject **/) *QBuffer {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QBuffer(QByteArray *, QObject *)
func NewQBuffer_1(buf *QByteArray /*777 QByteArray **/, parent *QObject /*777 QObject **/) *QBuffer {
	var convArg0 = buf.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP10QByteArrayP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QBuffer()
func DeleteQBuffer(this *QBuffer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qbuffer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & buffer()
func (this *QBuffer) Buffer() *QByteArray {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer6bufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:69
// index:1
// Public Visibility=Default Availability=Available
// [8] const QByteArray & buffer()
func (this *QBuffer) Buffer_1() *QByteArray {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer6bufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBuffer(QByteArray *)
func (this *QBuffer) SetBuffer(a *QByteArray /*777 QByteArray **/) {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer9setBufferEP10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QByteArray &)
func (this *QBuffer) SetData(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setData(const char *, int)
func (this *QBuffer) SetData_1(data string, len int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] const QByteArray & data()
func (this *QBuffer) Data() *QByteArray {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QBuffer) Open(openMode int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer4openE6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QBuffer) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 size()
func (this *QBuffer) Size() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 pos()
func (this *QBuffer) Pos() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool seek(qint64)
func (this *QBuffer) Seek(off int64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer4seekEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), off)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QBuffer) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine()
func (this *QBuffer) CanReadLine() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:87
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)
func (this *QBuffer) ConnectNotify(arg0 *QMetaMethod) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void disconnectNotify(const QMetaMethod &)
func (this *QBuffer) DisconnectNotify(arg0 *QMetaMethod) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
func (this *QBuffer) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QBuffer) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

//  body block end
