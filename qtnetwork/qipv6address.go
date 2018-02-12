package qtnetwork

// /usr/include/qt/QtNetwork/qhostaddress.h
// #include <qhostaddress.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QIPv6Address struct {
	*qtrt.CObject
}

func (this *QIPv6Address) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIPv6Address) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQIPv6AddressFromPointer(cthis unsafe.Pointer) *QIPv6Address {
	return &QIPv6Address{&qtrt.CObject{cthis}}
}
func (*QIPv6Address) NewFromPointer(cthis unsafe.Pointer) *QIPv6Address {
	return NewQIPv6AddressFromPointer(cthis)
}

func DeleteQIPv6Address(this *QIPv6Address) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QIPv6AddressD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
