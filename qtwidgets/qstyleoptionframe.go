package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QStyleOptionFrame struct {
	*QStyleOption
}

func (this *QStyleOptionFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionFrame) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionFrameFromPointer(cthis unsafe.Pointer) *QStyleOptionFrame {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionFrame{bcthis0}
}
func (*QStyleOptionFrame) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionFrame {
	return NewQStyleOptionFrameFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionFrame()
func NewQStyleOptionFrame() *QStyleOptionFrame {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QStyleOptionFrameC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFrame)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:147
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionFrame(int)
func NewQStyleOptionFrame_1(version int) *QStyleOptionFrame {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QStyleOptionFrameC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFrame)
	return gothis
}

func DeleteQStyleOptionFrame(this *QStyleOptionFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QStyleOptionFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionFrame__StyleOptionType = int

const QStyleOptionFrame__Type QStyleOptionFrame__StyleOptionType = 5

type QStyleOptionFrame__StyleOptionVersion = int

const QStyleOptionFrame__Version QStyleOptionFrame__StyleOptionVersion = 3

type QStyleOptionFrame__FrameFeature = int

const QStyleOptionFrame__None QStyleOptionFrame__FrameFeature = 0
const QStyleOptionFrame__Flat QStyleOptionFrame__FrameFeature = 1
const QStyleOptionFrame__Rounded QStyleOptionFrame__FrameFeature = 2

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
