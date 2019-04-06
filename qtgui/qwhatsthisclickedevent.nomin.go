//  header block begin

// +build !minimal

package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// /usr/include/qt/QtGui/qevent.h:718
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWhatsThisClickedEvent(const QString &)

/*

 */
func (*QWhatsThisClickedEvent) NewForInherit(href string) *QWhatsThisClickedEvent {
	return NewQWhatsThisClickedEvent(href)
}
func NewQWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	var tmpArg0 = qtcore.NewQString5(href)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWhatsThisClickedEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWhatsThisClickedEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:719
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWhatsThisClickedEvent()

/*

 */
func DeleteQWhatsThisClickedEvent(this *QWhatsThisClickedEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:721
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString href() const

/*

 */
func (this *QWhatsThisClickedEvent) Href() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWhatsThisClickedEvent4hrefEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end

//  keep block begin

func init_unused_10698() {
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
}

//  keep block end
