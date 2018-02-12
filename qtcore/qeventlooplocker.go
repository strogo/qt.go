package qtcore

// /usr/include/qt/QtCore/qeventloop.h
// #include <qeventloop.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QEventLoopLocker struct {
	*qtrt.CObject
}

func (this *QEventLoopLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QEventLoopLocker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQEventLoopLockerFromPointer(cthis unsafe.Pointer) *QEventLoopLocker {
	return &QEventLoopLocker{&qtrt.CObject{cthis}}
}
func (*QEventLoopLocker) NewFromPointer(cthis unsafe.Pointer) *QEventLoopLocker {
	return NewQEventLoopLockerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qeventloop.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEventLoopLocker()
func NewQEventLoopLocker() *QEventLoopLocker {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventLoopLockerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEventLoopLocker)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:94
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QEventLoopLocker(QEventLoop *)
func NewQEventLoopLocker_1(loop *QEventLoop /*777 QEventLoop **/) *QEventLoopLocker {
	var convArg0 = loop.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventLoopLockerC2EP10QEventLoop", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEventLoopLocker)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:95
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QEventLoopLocker(QThread *)
func NewQEventLoopLocker_2(thread *QThread /*777 QThread **/) *QEventLoopLocker {
	var convArg0 = thread.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventLoopLockerC2EP7QThread", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEventLoopLocker)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QEventLoopLocker()
func DeleteQEventLoopLocker(this *QEventLoopLocker) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QEventLoopLockerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
}

//  keep block end
