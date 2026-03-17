// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Espresso_mxnetTools_ImageBinaryRecordReader] class.
var (
	_Espresso_mxnetTools_ImageBinaryRecordReaderClass     Espresso_mxnetTools_ImageBinaryRecordReaderClass
	_Espresso_mxnetTools_ImageBinaryRecordReaderClassOnce sync.Once
)

func getEspresso_mxnetTools_ImageBinaryRecordReaderClass() Espresso_mxnetTools_ImageBinaryRecordReaderClass {
	_Espresso_mxnetTools_ImageBinaryRecordReaderClassOnce.Do(func() {
		_Espresso_mxnetTools_ImageBinaryRecordReaderClass = Espresso_mxnetTools_ImageBinaryRecordReaderClass{class: objc.GetClass("Espresso_mxnetTools_ImageBinaryRecordReader")}
	})
	return _Espresso_mxnetTools_ImageBinaryRecordReaderClass
}

// GetEspresso_mxnetTools_ImageBinaryRecordReaderClass returns the class object for Espresso_mxnetTools_ImageBinaryRecordReader.
func GetEspresso_mxnetTools_ImageBinaryRecordReaderClass() Espresso_mxnetTools_ImageBinaryRecordReaderClass {
	return getEspresso_mxnetTools_ImageBinaryRecordReaderClass()
}

type Espresso_mxnetTools_ImageBinaryRecordReaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec Espresso_mxnetTools_ImageBinaryRecordReaderClass) Alloc() Espresso_mxnetTools_ImageBinaryRecordReader {
	rv := objc.Send[Espresso_mxnetTools_ImageBinaryRecordReader](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.CurrentOffset]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.SetCurrentOffset]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.ImageData]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.ImageHeader]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.SetImageHeader]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.ImageID]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.Labels]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.LabelsPrivate]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.SetLabelsPrivate]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.NextRecordAndError]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.RecFileHandle]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.SetRecFileHandle]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.RecordHeader]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.SetRecordHeader]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.SeekRecordWithIDError]
//   - [Espresso_mxnetTools_ImageBinaryRecordReader.InitWithRecFileError]
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader
type Espresso_mxnetTools_ImageBinaryRecordReader struct {
	objectivec.Object
}

// Espresso_mxnetTools_ImageBinaryRecordReaderFromID constructs a [Espresso_mxnetTools_ImageBinaryRecordReader] from an objc.ID.
func Espresso_mxnetTools_ImageBinaryRecordReaderFromID(id objc.ID) Espresso_mxnetTools_ImageBinaryRecordReader {
	return Espresso_mxnetTools_ImageBinaryRecordReader{objectivec.Object{ID: id}}
}
// Ensure Espresso_mxnetTools_ImageBinaryRecordReader implements IEspresso_mxnetTools_ImageBinaryRecordReader.
var _ IEspresso_mxnetTools_ImageBinaryRecordReader = Espresso_mxnetTools_ImageBinaryRecordReader{}

// An interface definition for the [Espresso_mxnetTools_ImageBinaryRecordReader] class.
//
// # Methods
//
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.CurrentOffset]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.SetCurrentOffset]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.ImageData]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.ImageHeader]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.SetImageHeader]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.ImageID]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.Labels]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.LabelsPrivate]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.SetLabelsPrivate]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.NextRecordAndError]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.RecFileHandle]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.SetRecFileHandle]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.RecordHeader]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.SetRecordHeader]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.SeekRecordWithIDError]
//   - [IEspresso_mxnetTools_ImageBinaryRecordReader.InitWithRecFileError]
//
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader
type IEspresso_mxnetTools_ImageBinaryRecordReader interface {
	objectivec.IObject

	// Topic: Methods

	CurrentOffset() uint64
	SetCurrentOffset(value uint64)
	ImageData() objectivec.IObject
	ImageHeader() objectivec.IObject
	SetImageHeader(value objectivec.IObject)
	ImageID() objectivec.IObject
	Labels() objectivec.IObject
	LabelsPrivate() foundation.INSArray
	SetLabelsPrivate(value foundation.INSArray)
	NextRecordAndError() (bool, error)
	RecFileHandle() *foundation.NSFileHandle
	SetRecFileHandle(value *foundation.NSFileHandle)
	RecordHeader() objectivec.IObject
	SetRecordHeader(value objectivec.IObject)
	SeekRecordWithIDError(id objectivec.IObject) (bool, error)
	InitWithRecFileError(file objectivec.IObject) (Espresso_mxnetTools_ImageBinaryRecordReader, error)
}

// Init initializes the instance.
func (e Espresso_mxnetTools_ImageBinaryRecordReader) Init() Espresso_mxnetTools_ImageBinaryRecordReader {
	rv := objc.Send[Espresso_mxnetTools_ImageBinaryRecordReader](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e Espresso_mxnetTools_ImageBinaryRecordReader) Autorelease() Espresso_mxnetTools_ImageBinaryRecordReader {
	rv := objc.Send[Espresso_mxnetTools_ImageBinaryRecordReader](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspresso_mxnetTools_ImageBinaryRecordReader creates a new Espresso_mxnetTools_ImageBinaryRecordReader instance.
func NewEspresso_mxnetTools_ImageBinaryRecordReader() Espresso_mxnetTools_ImageBinaryRecordReader {
	class := getEspresso_mxnetTools_ImageBinaryRecordReaderClass()
	rv := objc.Send[Espresso_mxnetTools_ImageBinaryRecordReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/initWithRecFile:error:
func NewEspresso_mxnetTools_ImageBinaryRecordReaderWithRecFileError(file objectivec.IObject) (Espresso_mxnetTools_ImageBinaryRecordReader, error) {
	var errorPtr objc.ID
	instance := getEspresso_mxnetTools_ImageBinaryRecordReaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return Espresso_mxnetTools_ImageBinaryRecordReaderFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return Espresso_mxnetTools_ImageBinaryRecordReaderFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/imageData
func (e Espresso_mxnetTools_ImageBinaryRecordReader) ImageData() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imageData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/imageID
func (e Espresso_mxnetTools_ImageBinaryRecordReader) ImageID() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imageID"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/labels
func (e Espresso_mxnetTools_ImageBinaryRecordReader) Labels() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("labels"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/nextRecordAndError:
func (e Espresso_mxnetTools_ImageBinaryRecordReader) NextRecordAndError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("nextRecordAndError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("nextRecordAndError: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/seekRecordWithID:error:
func (e Espresso_mxnetTools_ImageBinaryRecordReader) SeekRecordWithIDError(id objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("seekRecordWithID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("seekRecordWithID:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/initWithRecFile:error:
func (e Espresso_mxnetTools_ImageBinaryRecordReader) InitWithRecFileError(file objectivec.IObject) (Espresso_mxnetTools_ImageBinaryRecordReader, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithRecFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return Espresso_mxnetTools_ImageBinaryRecordReader{}, foundation.NSErrorFrom(errorPtr)
	}
	return Espresso_mxnetTools_ImageBinaryRecordReaderFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/currentOffset
func (e Espresso_mxnetTools_ImageBinaryRecordReader) CurrentOffset() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("currentOffset"))
	return rv
}
func (e Espresso_mxnetTools_ImageBinaryRecordReader) SetCurrentOffset(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setCurrentOffset:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/imageHeader
func (e Espresso_mxnetTools_ImageBinaryRecordReader) ImageHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imageHeader"))
	return objectivec.Object{ID: rv}
}
func (e Espresso_mxnetTools_ImageBinaryRecordReader) SetImageHeader(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setImageHeader:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/labelsPrivate
func (e Espresso_mxnetTools_ImageBinaryRecordReader) LabelsPrivate() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("labelsPrivate"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e Espresso_mxnetTools_ImageBinaryRecordReader) SetLabelsPrivate(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setLabelsPrivate:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/recFileHandle
func (e Espresso_mxnetTools_ImageBinaryRecordReader) RecFileHandle() *foundation.NSFileHandle {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("recFileHandle"))
	if rv == 0 {
		return nil
	}
	val := foundation.NSFileHandleFromID(objc.ID(rv))
	return &val
}
func (e Espresso_mxnetTools_ImageBinaryRecordReader) SetRecFileHandle(value *foundation.NSFileHandle) {
	if value == nil {
		objc.Send[struct{}](e.ID, objc.Sel("setRecFileHandle:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](e.ID, objc.Sel("setRecFileHandle:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/recordHeader
func (e Espresso_mxnetTools_ImageBinaryRecordReader) RecordHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("recordHeader"))
	return objectivec.Object{ID: rv}
}
func (e Espresso_mxnetTools_ImageBinaryRecordReader) SetRecordHeader(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setRecordHeader:"), value)
}

