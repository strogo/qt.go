//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTextTableCellFormat struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextformat.h:945
// index:0
// void QTextTableCellFormat()
func NewQTextTableCellFormat() *QTextTableCellFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextTableCellFormat{cthis}
}

// /usr/include/qt/QtGui/qtextformat.h:947
// index:0
// inline
// bool isValid()
func (this *QTextTableCellFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:949
// index:0
// inline
// void setTopPadding(qreal)
func (this *QTextTableCellFormat) SetTopPadding(padding float64) {
	// 0: (, qreal padding), (&padding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat13setTopPaddingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:950
// index:0
// inline
// qreal topPadding()
func (this *QTextTableCellFormat) TopPadding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat10topPaddingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:952
// index:0
// inline
// void setBottomPadding(qreal)
func (this *QTextTableCellFormat) SetBottomPadding(padding float64) {
	// 0: (, qreal padding), (&padding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat16setBottomPaddingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:953
// index:0
// inline
// qreal bottomPadding()
func (this *QTextTableCellFormat) BottomPadding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat13bottomPaddingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:955
// index:0
// inline
// void setLeftPadding(qreal)
func (this *QTextTableCellFormat) SetLeftPadding(padding float64) {
	// 0: (, qreal padding), (&padding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat14setLeftPaddingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:956
// index:0
// inline
// qreal leftPadding()
func (this *QTextTableCellFormat) LeftPadding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat11leftPaddingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:958
// index:0
// inline
// void setRightPadding(qreal)
func (this *QTextTableCellFormat) SetRightPadding(padding float64) {
	// 0: (, qreal padding), (&padding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat15setRightPaddingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:959
// index:0
// inline
// qreal rightPadding()
func (this *QTextTableCellFormat) RightPadding() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat12rightPaddingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:961
// index:0
// inline
// void setPadding(qreal)
func (this *QTextTableCellFormat) SetPadding(padding float64) {
	// 0: (, qreal padding), (&padding)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat10setPaddingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &padding)
	gopp.ErrPrint(err, rv)
}

//  body block end