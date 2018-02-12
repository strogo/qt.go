package qtwidgets

// /usr/include/qt/QtWidgets/qkeyeventtransition.h
// #include <qkeyeventtransition.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

// void onTransition(class QEvent *)
func (this *QKeyEventTransition) InheritOnTransition(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onTransition", f)
}

// bool eventTest(class QEvent *)
func (this *QKeyEventTransition) InheritEventTest(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventTest", f)
}

type QKeyEventTransition struct {
	*qtcore.QEventTransition
}

func (this *QKeyEventTransition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEventTransition.GetCthis()
	}
}
func (this *QKeyEventTransition) SetCthis(cthis unsafe.Pointer) {
	this.QEventTransition = qtcore.NewQEventTransitionFromPointer(cthis)
}
func NewQKeyEventTransitionFromPointer(cthis unsafe.Pointer) *QKeyEventTransition {
	bcthis0 := qtcore.NewQEventTransitionFromPointer(cthis)
	return &QKeyEventTransition{bcthis0}
}
func (*QKeyEventTransition) NewFromPointer(cthis unsafe.Pointer) *QKeyEventTransition {
	return NewQKeyEventTransitionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QKeyEventTransition) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QKeyEventTransition10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeyEventTransition(QState *)
func NewQKeyEventTransition(sourceState *qtcore.QState /*777 QState **/) *QKeyEventTransition {
	var convArg0 = sourceState.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeyEventTransition(QObject *, QEvent::Type, int, QState *)
func NewQKeyEventTransition_1(object *qtcore.QObject /*777 QObject **/, type_ int, key int, sourceState *qtcore.QState /*777 QState **/) *QKeyEventTransition {
	var convArg0 = object.GetCthis()
	var convArg3 = sourceState.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP7QObjectN6QEvent4TypeEiP6QState", qtrt.FFI_TYPE_POINTER, convArg0, type_, key, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QKeyEventTransition()
func DeleteQKeyEventTransition(this *QKeyEventTransition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QKeyEventTransitionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:62
// index:0
// Public Visibility=Default Availability=Available
// [4] int key()
func (this *QKeyEventTransition) Key() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QKeyEventTransition3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKey(int)
func (this *QKeyEventTransition) SetKey(key int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QKeyEventTransition6setKeyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifierMask()
func (this *QKeyEventTransition) ModifierMask() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QKeyEventTransition12modifierMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifierMask(Qt::KeyboardModifiers)
func (this *QKeyEventTransition) SetModifierMask(modifiers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QKeyEventTransition15setModifierMaskE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:69
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)
func (this *QKeyEventTransition) OnTransition(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QKeyEventTransition12onTransitionEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)
func (this *QKeyEventTransition) EventTest(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QKeyEventTransition9eventTestEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
