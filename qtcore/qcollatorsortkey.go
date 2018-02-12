package qtcore

// /usr/include/qt/QtCore/qcollator.h
// #include <qcollator.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 105
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QCollatorSortKey struct {
	*qtrt.CObject
}

func (this *QCollatorSortKey) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCollatorSortKey) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCollatorSortKeyFromPointer(cthis unsafe.Pointer) *QCollatorSortKey {
	return &QCollatorSortKey{&qtrt.CObject{cthis}}
}
func (*QCollatorSortKey) NewFromPointer(cthis unsafe.Pointer) *QCollatorSortKey {
	return NewQCollatorSortKeyFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcollator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCollatorSortKey()
func DeleteQCollatorSortKey(this *QCollatorSortKey) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCollatorSortKeyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcollator.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCollatorSortKey &)
func (this *QCollatorSortKey) Swap(other *QCollatorSortKey) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCollatorSortKey4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcollator.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int compare(const QCollatorSortKey &)
func (this *QCollatorSortKey) Compare(key *QCollatorSortKey) int {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCollatorSortKey7compareERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
