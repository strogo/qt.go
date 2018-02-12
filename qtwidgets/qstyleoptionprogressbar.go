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

type QStyleOptionProgressBar struct {
	*QStyleOption
}

func (this *QStyleOptionProgressBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionProgressBar) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionProgressBarFromPointer(cthis unsafe.Pointer) *QStyleOptionProgressBar {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionProgressBar{bcthis0}
}
func (*QStyleOptionProgressBar) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionProgressBar {
	return NewQStyleOptionProgressBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:342
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionProgressBar()
func NewQStyleOptionProgressBar() *QStyleOptionProgressBar {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QStyleOptionProgressBarC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionProgressBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionProgressBar)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:346
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionProgressBar(int)
func NewQStyleOptionProgressBar_1(version int) *QStyleOptionProgressBar {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QStyleOptionProgressBarC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionProgressBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionProgressBar)
	return gothis
}

func DeleteQStyleOptionProgressBar(this *QStyleOptionProgressBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QStyleOptionProgressBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionProgressBar__StyleOptionType = int

const QStyleOptionProgressBar__Type QStyleOptionProgressBar__StyleOptionType = 6

type QStyleOptionProgressBar__StyleOptionVersion = int

const QStyleOptionProgressBar__Version QStyleOptionProgressBar__StyleOptionVersion = 2

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
