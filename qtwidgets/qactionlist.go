package qtwidgets

// /usr/include/qt/QtWidgets/qaction.h
// #include <qaction.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  keep block begin

func init_unused_10115() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end

//  body block begin
type QActionList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QActionList) Operator_equal0() *QActionList {
	// QActionList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QActionList) Operator_equal1() *QActionList {
	// QActionList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QActionList) Swap0() {
	// QActionList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QActionList) Operator_equal_equal0() bool {
	// QActionList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QActionList) Operator_not_equal0() bool {
	// QActionList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QActionList) Size0() int {
	// QActionList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QActionList) Detach0() {
	// QActionList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QActionList) DetachShared0() {
	// QActionList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QActionList) IsDetached0() bool {
	// QActionList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QActionList) SetSharable0() {
	// QActionList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QActionList) IsSharedWith0() bool {
	// QActionList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QActionList) IsEmpty0() bool {
	// QActionList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QActionList) Clear0() {
	// QActionList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QActionList) At0() *QAction {
	// QActionList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// const T & operator[](int)
func (this *QActionList) Operator_get_index0() *QAction {
	// QActionList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// T & operator[](int)
func (this *QActionList) Operator_get_index1() *QAction {
	// QActionList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// void reserve(int)
func (this *QActionList) Reserve0() {
	// QActionList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QActionList) Append0() {
	// QActionList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QActionList) Append1() {
	// QActionList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QActionList) Prepend0() {
	// QActionList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QActionList) Insert0() {
	// QActionList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QActionList) Replace0() {
	// QActionList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QActionList) RemoveAt0() {
	// QActionList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QActionList) RemoveAll0() int {
	// QActionList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QActionList) RemoveOne0() bool {
	// QActionList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QActionList) TakeAt0() *QAction {
	// QActionList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// T takeFirst()
func (this *QActionList) TakeFirst0() *QAction {
	// QActionList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// T takeLast()
func (this *QActionList) TakeLast0() *QAction {
	// QActionList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// void move(int, int)
func (this *QActionList) Move0() {
	// QActionList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QActionList) Swap1() {
	// QActionList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QActionList) IndexOf0() int {
	// QActionList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QActionList) LastIndexOf0() int {
	// QActionList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QActionList) Contains0() bool {
	// QActionList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QActionList) Count0() int {
	// QActionList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QActionList) Begin0() {
	// QActionList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QActionList) Begin1() {
	// QActionList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QActionList) Cbegin0() {
	// QActionList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QActionList) ConstBegin0() {
	// QActionList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QActionList) End0() {
	// QActionList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QActionList) End1() {
	// QActionList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QActionList) Cend0() {
	// QActionList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QActionList) ConstEnd0() {
	// QActionList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QActionList) Rbegin0() {
	// QActionList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QActionList) Rend0() {
	// QActionList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QActionList) Rbegin1() {
	// QActionList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QActionList) Rend1() {
	// QActionList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QActionList) Crbegin0() {
	// QActionList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QActionList) Crend0() {
	// QActionList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QActionList) Insert1() {
	// QActionList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QActionList) Erase0() {
	// QActionList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QActionList) Erase1() {
	// QActionList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QActionList) Count1() int {
	// QActionList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QActionList) Length0() int {
	// QActionList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QActionList) First0() *QAction {
	// QActionList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// const T & constFirst()
func (this *QActionList) ConstFirst0() *QAction {
	// QActionList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// const T & first()
func (this *QActionList) First1() *QAction {
	// QActionList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// T & last()
func (this *QActionList) Last0() *QAction {
	// QActionList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// const T & last()
func (this *QActionList) Last1() *QAction {
	// QActionList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// const T & constLast()
func (this *QActionList) ConstLast0() *QAction {
	// QActionList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// void removeFirst()
func (this *QActionList) RemoveFirst0() {
	// QActionList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QActionList) RemoveLast0() {
	// QActionList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QActionList) StartsWith0() bool {
	// QActionList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QActionList) EndsWith0() bool {
	// QActionList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QActionList) Mid0() *QActionList {
	// QActionList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QActionList) Value0() *QAction {
	// QActionList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// T value(int, const T &)
func (this *QActionList) Value1() *QAction {
	// QActionList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// void push_back(const T &)
func (this *QActionList) Push_back0() {
	// QActionList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QActionList) Push_front0() {
	// QActionList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QActionList) Front0() *QAction {
	// QActionList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// const T & front()
func (this *QActionList) Front1() *QAction {
	// QActionList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// T & back()
func (this *QActionList) Back0() *QAction {
	// QActionList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// const T & back()
func (this *QActionList) Back1() *QAction {
	// QActionList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAction{}
}

// void pop_front()
func (this *QActionList) Pop_front0() {
	// QActionList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QActionList) Pop_back0() {
	// QActionList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QActionList) Empty0() bool {
	// QActionList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QActionList) Operator_add_equal0() *QActionList {
	// QActionList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QActionList) Operator_add0() *QActionList {
	// QActionList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QActionList) Operator_add_equal1() *QActionList {
	// QActionList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QActionList) Operator_left_shift0() *QActionList {
	// QActionList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QActionList) Operator_left_shift1() *QActionList {
	// QActionList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QActionList) ToVector0() {
	// QActionList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QActionList) ToSet0() {
	// QActionList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QActionList) FromVector0() *QActionList {
	// QActionList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QActionList) FromSet0() *QActionList {
	// QActionList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QActionList) FromStdList0() *QActionList {
	// QActionList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QActionList) ToStdList0() {
	// QActionList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QActionList) Detach_helper_grow0() {
	// QActionList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QActionList) Detach_helper0() {
	// QActionList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QActionList) Detach_helper1() {
	// QActionList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QActionList) Dealloc0() {
	// QActionList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QActionList) Node_construct0() {
	// QActionList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QActionList) Node_destruct0() {
	// QActionList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QActionList) Node_copy0() {
	// QActionList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QActionList) Node_destruct1() {
	// QActionList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QActionList) IsValidIterator0() bool {
	// QActionList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QActionList) Op_eq_impl0() bool {
	// QActionList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QActionList) Op_eq_impl1() bool {
	// QActionList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QActionList) Contains_impl0() bool {
	// QActionList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QActionList) Contains_impl1() bool {
	// QActionList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QActionList) Count_impl0() int {
	// QActionList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QActionList) Count_impl1() int {
	// QActionList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QActionList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
