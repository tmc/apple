// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
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
	rv := objc.SendIfResponds[EspressoMxnetToolsImageBinaryRecordReader](objc.ID(ec.class), objc.Sel("alloc"))
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
type IEspressoMxnetToolsImageBinaryRecordReader interface {
	objectivec.IObject

	// Topic: Methods

	CurrentOffset() uint64
	SetCurrentOffset(value uint64)
	ImageData() objectivec.IObject
	ImageHeader() appleneuralengine.MxnetToolsImageHeaderT
	SetImageHeader(value appleneuralengine.MxnetToolsImageHeaderT)
	ImageID() appleneuralengine.MxnetToolsImageIDT
	Labels() objectivec.IObject
	LabelsPrivate() foundation.INSArray
	SetLabelsPrivate(value foundation.INSArray)
	NextRecordAndError() (bool, error)
	RecFileHandle() foundation.FileHandle
	SetRecFileHandle(value foundation.FileHandle)
	RecordHeader() appleneuralengine.MxnetToolsRecordHeaderT
	SetRecordHeader(value appleneuralengine.MxnetToolsRecordHeaderT)
	SeekRecordWithIDError(id *MxnetToolsImageIDT) (bool, error)
	InitWithRecFileError(file objectivec.IObject) (EspressoMxnetToolsImageBinaryRecordReader, error)
}

// Init initializes the instance.
func (e EspressoMxnetToolsImageBinaryRecordReader) Init() EspressoMxnetToolsImageBinaryRecordReader {
	rv := objc.SendIfResponds[EspressoMxnetToolsImageBinaryRecordReader](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoMxnetToolsImageBinaryRecordReader) Autorelease() EspressoMxnetToolsImageBinaryRecordReader {
	rv := objc.SendIfResponds[EspressoMxnetToolsImageBinaryRecordReader](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoMxnetToolsImageBinaryRecordReader creates a new EspressoMxnetToolsImageBinaryRecordReader instance.
func NewEspressoMxnetToolsImageBinaryRecordReader() EspressoMxnetToolsImageBinaryRecordReader {
	class := getEspressoMxnetToolsImageBinaryRecordReaderClass()
	rv := objc.SendIfResponds[EspressoMxnetToolsImageBinaryRecordReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewEspresso_mxnetTools_ImageBinaryRecordReaderWithRecFileError(file objectivec.IObject) (EspressoMxnetToolsImageBinaryRecordReader, error) {
	var errorPtr objc.ID
	instance := getEspressoMxnetToolsImageBinaryRecordReaderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRecFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return EspressoMxnetToolsImageBinaryRecordReader{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return EspressoMxnetToolsImageBinaryRecordReader{}, objc.ErrInitFailed
	}
	return EspressoMxnetToolsImageBinaryRecordReaderFromID(rv), nil
}

func (e EspressoMxnetToolsImageBinaryRecordReader) ImageData() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("imageData"))
	return objectivec.Object{ID: rv}
}
func (e EspressoMxnetToolsImageBinaryRecordReader) ImageID() appleneuralengine.MxnetToolsImageIDT {
	rv := objc.SendIfResponds[appleneuralengine.MxnetToolsImageIDT](e.ID, objc.Sel("imageID"))
	return appleneuralengine.MxnetToolsImageIDT(rv)
}
func (e EspressoMxnetToolsImageBinaryRecordReader) Labels() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("labels"))
	return objectivec.Object{ID: rv}
}
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
func (e EspressoMxnetToolsImageBinaryRecordReader) SeekRecordWithIDError(id *MxnetToolsImageIDT) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("seekRecordWithID:error:"), unsafe.Pointer(id), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("seekRecordWithID:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e EspressoMxnetToolsImageBinaryRecordReader) InitWithRecFileError(file objectivec.IObject) (EspressoMxnetToolsImageBinaryRecordReader, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithRecFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(EspressoMxnetToolsImageBinaryRecordReader), foundation.NSErrorFrom(errorPtr)
	}
	return EspressoMxnetToolsImageBinaryRecordReaderFromID(rv), nil

}

func (e EspressoMxnetToolsImageBinaryRecordReader) CurrentOffset() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("currentOffset"))
	return rv
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetCurrentOffset(value uint64) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setCurrentOffset:"), value)
}
func (e EspressoMxnetToolsImageBinaryRecordReader) ImageHeader() appleneuralengine.MxnetToolsImageHeaderT {
	rv := objc.SendIfResponds[appleneuralengine.MxnetToolsImageHeaderT](e.ID, objc.Sel("imageHeader"))
	return appleneuralengine.MxnetToolsImageHeaderT(rv)
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetImageHeader(value appleneuralengine.MxnetToolsImageHeaderT) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setImageHeader:"), value)
}
func (e EspressoMxnetToolsImageBinaryRecordReader) LabelsPrivate() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("labelsPrivate"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetLabelsPrivate(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setLabelsPrivate:"), value)
}
func (e EspressoMxnetToolsImageBinaryRecordReader) RecFileHandle() foundation.FileHandle {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("recFileHandle"))
	return foundation.FileHandleFromID(objc.ID(rv))
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetRecFileHandle(value foundation.FileHandle) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setRecFileHandle:"), value)
}
func (e EspressoMxnetToolsImageBinaryRecordReader) RecordHeader() appleneuralengine.MxnetToolsRecordHeaderT {
	rv := objc.SendIfResponds[appleneuralengine.MxnetToolsRecordHeaderT](e.ID, objc.Sel("recordHeader"))
	return appleneuralengine.MxnetToolsRecordHeaderT(rv)
}
func (e EspressoMxnetToolsImageBinaryRecordReader) SetRecordHeader(value appleneuralengine.MxnetToolsRecordHeaderT) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setRecordHeader:"), value)
}
