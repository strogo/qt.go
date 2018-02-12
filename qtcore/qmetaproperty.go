package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QMetaProperty struct {
	*qtrt.CObject
}

func (this *QMetaProperty) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaProperty) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMetaPropertyFromPointer(cthis unsafe.Pointer) *QMetaProperty {
	return &QMetaProperty{&qtrt.CObject{cthis}}
}
func (*QMetaProperty) NewFromPointer(cthis unsafe.Pointer) *QMetaProperty {
	return NewQMetaPropertyFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMetaProperty()
func NewQMetaProperty() *QMetaProperty {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMetaPropertyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMetaProperty)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:250
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * name()
func (this *QMetaProperty) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:251
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * typeName()
func (this *QMetaProperty) TypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8typeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:252
// index:0
// Public Visibility=Default Availability=Available
// [4] QVariant::Type type()
func (this *QMetaProperty) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int userType()
func (this *QMetaProperty) UserType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8userTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:254
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyIndex()
func (this *QMetaProperty) PropertyIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty13propertyIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:256
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable()
func (this *QMetaProperty) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:257
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable()
func (this *QMetaProperty) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:258
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isResettable()
func (this *QMetaProperty) IsResettable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isResettableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:259
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDesignable(const QObject *)
func (this *QMetaProperty) IsDesignable(obj *QObject /*777 const QObject **/) bool {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isDesignableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScriptable(const QObject *)
func (this *QMetaProperty) IsScriptable(obj *QObject /*777 const QObject **/) bool {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isScriptableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStored(const QObject *)
func (this *QMetaProperty) IsStored(obj *QObject /*777 const QObject **/) bool {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8isStoredEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:262
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEditable(const QObject *)
func (this *QMetaProperty) IsEditable(obj *QObject /*777 const QObject **/) bool {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isEditableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:263
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUser(const QObject *)
func (this *QMetaProperty) IsUser(obj *QObject /*777 const QObject **/) bool {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty6isUserEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:264
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isConstant()
func (this *QMetaProperty) IsConstant() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isConstantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:265
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinal()
func (this *QMetaProperty) IsFinal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty7isFinalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:267
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlagType()
func (this *QMetaProperty) IsFlagType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isFlagTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:268
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnumType()
func (this *QMetaProperty) IsEnumType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isEnumTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:269
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaEnum enumerator()
func (this *QMetaProperty) Enumerator() *QMetaEnum /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10enumeratorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaEnumFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaEnum)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:271
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNotifySignal()
func (this *QMetaProperty) HasNotifySignal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty15hasNotifySignalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:272
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaMethod notifySignal()
func (this *QMetaProperty) NotifySignal() *QMetaMethod /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12notifySignalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaMethod)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:273
// index:0
// Public Visibility=Default Availability=Available
// [4] int notifySignalIndex()
func (this *QMetaProperty) NotifySignalIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty17notifySignalIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:275
// index:0
// Public Visibility=Default Availability=Available
// [4] int revision()
func (this *QMetaProperty) Revision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8revisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:277
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant read(const QObject *)
func (this *QMetaProperty) Read(obj *QObject /*777 const QObject **/) *QVariant /*123*/ {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty4readEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:278
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write(QObject *, const QVariant &)
func (this *QMetaProperty) Write(obj *QObject /*777 QObject **/, value *QVariant) bool {
	var convArg0 = obj.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty5writeEP7QObjectRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:279
// index:0
// Public Visibility=Default Availability=Available
// [1] bool reset(QObject *)
func (this *QMetaProperty) Reset(obj *QObject /*777 QObject **/) bool {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty5resetEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:281
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant readOnGadget(const void *)
func (this *QMetaProperty) ReadOnGadget(gadget unsafe.Pointer /*666*/) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12readOnGadgetEPKv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:282
// index:0
// Public Visibility=Default Availability=Available
// [1] bool writeOnGadget(void *, const QVariant &)
func (this *QMetaProperty) WriteOnGadget(gadget unsafe.Pointer /*666*/, value *QVariant) bool {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:283
// index:0
// Public Visibility=Default Availability=Available
// [1] bool resetOnGadget(void *)
func (this *QMetaProperty) ResetOnGadget(gadget unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty13resetOnGadgetEPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:285
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasStdCppSet()
func (this *QMetaProperty) HasStdCppSet() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12hasStdCppSetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:286
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QMetaProperty) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:287
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMetaObject * enclosingMetaObject()
func (this *QMetaProperty) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty19enclosingMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQMetaProperty(this *QMetaProperty) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMetaPropertyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
