// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoMxnetToolsImageBinaryRecordReader] class.
var (
	_EspressoMxnetToolsImageBinaryRecordReaderClass     EspressoMxnetToolsImageBinaryRecordReaderClass
	_EspressoMxnetToolsImageBinaryRecordReaderClassOnce sync.Once
)

func getEspressoMxnetToolsImageBinaryRecordReaderClass() EspressoMxnetToolsImageBinaryRecordReaderClass {
	_EspressoMxnetToolsImageBinaryRecordReaderClassOnce.Do(func() {
		_EspressoMxnetToolsImageBinaryRecordReaderClass = EspressoMxnetToolsImageBinaryRecordReaderClass{class: objc.GetClass("Espresso_mxnetTools_ImageBinaryRecordReader")}
	})
	return _EspressoMxnetToolsImageBinaryRecordReaderClass
}

// GetEspressoMxnetToolsImageBinaryRecordReaderClass returns the class object for Espresso_mxnetTools_ImageBinaryRecordReader.
func GetEspressoMxnetToolsImageBinaryRecordReaderClass() EspressoMxnetToolsImageBinaryRecordReaderClass {
	return getEspressoMxnetToolsImageBinaryRecordReaderClass()
}

type EspressoMxnetToolsImageBinaryRecordReaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoMxnetToolsImageBinaryRecordReaderClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoMxnetToolsImageBinaryRecordReaderClass) Alloc() EspressoMxnetToolsImageBinaryRecordReader {
	rv := objc.Send[EspressoMxnetToolsImageBinaryRecordReader](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoMxnetToolsImageBinaryRecordReader.CurrentOffset]
//   - [EspressoMxnetToolsImageBinaryRecordReader.SetCurrentOffset]
//   - [EspressoMxnetToolsImageBinaryRecordReader.ImageData]
//   - [EspressoMxnetToolsImageBinaryRecordReader.ImageHeader]
//   - [EspressoMxnetToolsImageBinaryRecordReader.SetImageHeader]
//   - [EspressoMxnetToolsImageBinaryRecordReader.ImageID]
//   - [EspressoMxnetToolsImageBinaryRecordReader.Labels]
//   - [EspressoMxnetToolsImageBinaryRecordReader.LabelsPrivate]
//   - [EspressoMxnetToolsImageBinaryRecordReader.SetLabelsPrivate]
//   - [EspressoMxnetToolsImageBinaryRecordReader.NextRecordAndError]
//   - [EspressoMxnetToolsImageBinaryRecordReader.RecFileHandle]
//   - [EspressoMxnetToolsImageBinaryRecordReader.SetRecFileHandle]
//   - [EspressoMxnetToolsImageBinaryRecordReader.RecordHeader]
//   - [EspressoMxnetToolsImageBinaryRecordReader.SetRecordHeader]
//   - [EspressoMxnetToolsImageBinaryRecordReader.SeekRecordWithIDError]
//   - [EspressoMxnetToolsImageBinaryRecordReader.InitWithRecFileError]
//
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader
type EspressoMxnetToolsImageBinaryRecordReader struct {
	objectivec.Object
}

// EspressoMxnetToolsImageBinaryRecordReaderFromID constructs a [EspressoMxnetToolsImageBinaryRecordReader] from an objc.ID.
func EspressoMxnetToolsImageBinaryRecordReaderFromID(id objc.ID) EspressoMxnetToolsImageBinaryRecordReader {
	return EspressoMxnetToolsImageBinaryRecordReader{objectivec.Object{ID: id}}
}

// Espresso_mxnetTools_ImageBinaryRecordReaderFromID is an alias for [EspressoMxnetToolsImageBinaryRecordReaderFromID] for cross-framework compatibility.
func Espresso_mxnetTools_ImageBinaryRecordReaderFromID(id objc.ID) EspressoMxnetToolsImageBinaryRecordReader {
	return EspressoMxnetToolsImageBinaryRecordReaderFromID(id)
}

// Ensure EspressoMxnetToolsImageBinaryRecordReader implements IEspressoMxnetToolsImageBinaryRecordReader.
var _ IEspressoMxnetToolsImageBinaryRecordReader = EspressoMxnetToolsImageBinaryRecordReader{}

// An interface definition for the [EspressoMxnetToolsImageBinaryRecordReader] class.
//
// # Methods
//
//   - [IEspressoMxnetToolsImageBinaryRecordReader.CurrentOffset]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.SetCurrentOffset]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.ImageData]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.ImageHeader]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.SetImageHeader]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.ImageID]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.Labels]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.LabelsPrivate]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.SetLabelsPrivate]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.NextRecordAndError]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.RecFileHandle]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.SetRecFileHandle]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.RecordHeader]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.SetRecordHeader]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.SeekRecordWithIDError]
//   - [IEspressoMxnetToolsImageBinaryRecordReader.InitWithRecFileError]
//
// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader
type IEspressoMxnetToolsImageBinaryRecordReader interface {
	objectivec.IObject

	// Topic: Methods

	CurrentOffset() uint64
	SetCurrentOffset(value uint64)
	ImageData() objectivec.IObject
	ImageHeader() MxnetToolsImageHeaderT
	SetImageHeader(value MxnetToolsImageHeaderT)
	ImageID() MxnetToolsImageIDT
	Labels() objectivec.IObject
	LabelsPrivate() foundation.INSArray
	SetLabelsPrivate(value foundation.INSArray)
	NextRecordAndError() (bool, error)
	RecFileHandle() foundation.NSFileHandle
	SetRecFileHandle(value foundation.NSFileHandle)
	RecordHeader() MxnetToolsRecordHeaderT
	SetRecordHeader(value MxnetToolsRecordHeaderT)
	SeekRecordWithIDError(id unsafe.Pointer) (bool, error)
	InitWithRecFileError(file objectivec.IObject) (EspressoMxnetToolsImageBinaryRecordReader, error)
}

// Init initializes the instance.
func (e EspressoMxnetToolsImageBinaryRecordReader) Init() EspressoMxnetToolsImageBinaryRecordReader {
	rv := objc.Send[EspressoMxnetToolsImageBinaryRecordReader](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoMxnetToolsImageBinaryRecordReader) Autorelease() EspressoMxnetToolsImageBinaryRecordReader {
	rv := objc.Send[EspressoMxnetToolsImageBinaryRecordReader](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoMxnetToolsImageBinaryRecordReader creates a new EspressoMxnetToolsImageBinaryRecordReader instance.
func NewEspressoMxnetToolsImageBinaryRecordReader() EspressoMxnetToolsImageBinaryRecordReader {
	class := getEspressoMxnetToolsImageBinaryRecordReaderClass()
	rv := objc.Send[EspressoMxnetToolsImageBinaryRecordReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/initWithRecFile:error:
func NewEspresso_mxnetTools_ImageBinaryRecordReaderWithRecFileError(file objectivec.IObject) (EspressoMxnetToolsImageBinaryRecordReader, error) {
	var errorPtr objc.ID
	instance := getEspressoMxnetToolsImageBinaryRecordReaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return EspressoMxnetToolsImageBinaryRecordReader{}, foundation.NSErrorFrom(errorPtr)
	}
	return EspressoMxnetToolsImageBinaryRecordReaderFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/imageData
func (e EspressoMxnetToolsImageBinaryRecordReader) ImageData() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imageData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/imageID
func (e EspressoMxnetToolsImageBinaryRecordReader) ImageID() MxnetToolsImageIDT {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imageID"))
	_ = rv
	return MxnetToolsImageIDT{}
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/labels
func (e EspressoMxnetToolsImageBinaryRecordReader) Labels() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("labels"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/nextRecordAndError:
func (e EspressoMxnetToolsImageBinaryRecordReader) NextRecordAndError() (bool, error) {
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

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/seekRecordWithID:error:
func (e EspressoMxnetToolsImageBinaryRecordReader) SeekRecordWithIDError(id unsafe.Pointer) (bool, error) {
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

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/initWithRecFile:error:
func (e EspressoMxnetToolsImageBinaryRecordReader) InitWithRecFileError(file objectivec.IObject) (EspressoMxnetToolsImageBinaryRecordReader, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithRecFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return EspressoMxnetToolsImageBinaryRecordReader{}, foundation.NSErrorFrom(errorPtr)
	}
	return EspressoMxnetToolsImageBinaryRecordReaderFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/currentOffset
func (e EspressoMxnetToolsImageBinaryRecordReader) CurrentOffset() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("currentOffset"))
	return rv
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetCurrentOffset(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setCurrentOffset:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/imageHeader
func (e EspressoMxnetToolsImageBinaryRecordReader) ImageHeader() MxnetToolsImageHeaderT {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imageHeader"))
	_ = rv
	return MxnetToolsImageHeaderT{}
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetImageHeader(value MxnetToolsImageHeaderT) {
	objc.Send[struct{}](e.ID, objc.Sel("setImageHeader:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/labelsPrivate
func (e EspressoMxnetToolsImageBinaryRecordReader) LabelsPrivate() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("labelsPrivate"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetLabelsPrivate(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setLabelsPrivate:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/recFileHandle
func (e EspressoMxnetToolsImageBinaryRecordReader) RecFileHandle() foundation.NSFileHandle {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("recFileHandle"))
	return foundation.NSFileHandleFromID(objc.ID(rv))
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetRecFileHandle(value foundation.NSFileHandle) {
	objc.Send[struct{}](e.ID, objc.Sel("setRecFileHandle:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/Espresso_mxnetTools_ImageBinaryRecordReader/recordHeader
func (e EspressoMxnetToolsImageBinaryRecordReader) RecordHeader() MxnetToolsRecordHeaderT {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("recordHeader"))
	_ = rv
	return MxnetToolsRecordHeaderT{}
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetRecordHeader(value MxnetToolsRecordHeaderT) {
	objc.Send[struct{}](e.ID, objc.Sel("setRecordHeader:"), value)
}
