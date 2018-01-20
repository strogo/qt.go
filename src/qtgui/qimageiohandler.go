//  header block begin
// /usr/include/qt/QtGui/qimageiohandler.h
// #include <qimageiohandler.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QImageIOHandler struct {
	*qtrt.CObject
}

func (this *QImageIOHandler) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQImageIOHandlerFromPointer(cthis unsafe.Pointer) *QImageIOHandler {
	return &QImageIOHandler{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qimageiohandler.h:62
// index:0
// Public
// void QImageIOHandler()
func NewQImageIOHandler() *QImageIOHandler {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandlerC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageIOHandlerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:63
// index:0
// Public virtual
// void ~QImageIOHandler()
func DeleteQImageIOHandler(*QImageIOHandler) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandlerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:65
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QImageIOHandler) SetDevice(device unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), device)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:66
// index:0
// Public
// QIODevice * device()
func (this *QImageIOHandler) Device() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:68
// index:0
// Public
// void setFormat(const class QByteArray &)
func (this *QImageIOHandler) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:69
// index:1
// Public
// void setFormat(const class QByteArray &)
func (this *QImageIOHandler) SetFormat_1(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler9setFormatERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:70
// index:0
// Public
// QByteArray format()
func (this *QImageIOHandler) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:72
// index:0
// Public virtual
// QByteArray name()
func (this *QImageIOHandler) Name() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:74
// index:0
// Public pure virtual
// bool canRead()
func (this *QImageIOHandler) CanRead() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler7canReadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:75
// index:0
// Public pure virtual
// bool read(class QImage *)
func (this *QImageIOHandler) Read(image unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler4readEP6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), image)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:76
// index:0
// Public virtual
// bool write(const class QImage &)
func (this *QImageIOHandler) Write(image *QImage) interface{} {
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler5writeERK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:115
// index:0
// Public virtual
// QVariant option(enum QImageIOHandler::ImageOption)
func (this *QImageIOHandler) Option(option int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler6optionENS_11ImageOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:116
// index:0
// Public virtual
// void setOption(enum QImageIOHandler::ImageOption, const class QVariant &)
func (this *QImageIOHandler) SetOption(option int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler9setOptionENS_11ImageOptionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:117
// index:0
// Public virtual
// bool supportsOption(enum QImageIOHandler::ImageOption)
func (this *QImageIOHandler) SupportsOption(option int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler14supportsOptionENS_11ImageOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:120
// index:0
// Public virtual
// bool jumpToNextImage()
func (this *QImageIOHandler) JumpToNextImage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler15jumpToNextImageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:121
// index:0
// Public virtual
// bool jumpToImage(int)
func (this *QImageIOHandler) JumpToImage(imageNumber int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QImageIOHandler11jumpToImageEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &imageNumber)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:122
// index:0
// Public virtual
// int loopCount()
func (this *QImageIOHandler) LoopCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler9loopCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:123
// index:0
// Public virtual
// int imageCount()
func (this *QImageIOHandler) ImageCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler10imageCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:124
// index:0
// Public virtual
// int nextImageDelay()
func (this *QImageIOHandler) NextImageDelay() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler14nextImageDelayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:125
// index:0
// Public virtual
// int currentImageNumber()
func (this *QImageIOHandler) CurrentImageNumber() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler18currentImageNumberEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qimageiohandler.h:126
// index:0
// Public virtual
// QRect currentImageRect()
func (this *QImageIOHandler) CurrentImageRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QImageIOHandler16currentImageRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end