// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [InMemoryArchiveReader] class.
var (
	_InMemoryArchiveReaderClass     InMemoryArchiveReaderClass
	_InMemoryArchiveReaderClassOnce sync.Once
)

func getInMemoryArchiveReaderClass() InMemoryArchiveReaderClass {
	_InMemoryArchiveReaderClassOnce.Do(func() {
		_InMemoryArchiveReaderClass = InMemoryArchiveReaderClass{class: objc.GetClass("_InMemoryArchiveReader")}
	})
	return _InMemoryArchiveReaderClass
}

// GetInMemoryArchiveReaderClass returns the class object for _InMemoryArchiveReader.
func GetInMemoryArchiveReaderClass() InMemoryArchiveReaderClass {
	return getInMemoryArchiveReaderClass()
}

type InMemoryArchiveReaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic InMemoryArchiveReaderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic InMemoryArchiveReaderClass) Alloc() InMemoryArchiveReader {
	rv := objc.Send[InMemoryArchiveReader](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [InMemoryArchiveReader.CopyLayerShapesToContainer]
//   - [InMemoryArchiveReader.LayerInfos]
//   - [InMemoryArchiveReader.LoadUpdatableParams]
//   - [InMemoryArchiveReader.ModelPath]
//   - [InMemoryArchiveReader.TransformParams]
//   - [InMemoryArchiveReader.DebugDescription]
//   - [InMemoryArchiveReader.Description]
//   - [InMemoryArchiveReader.Hash]
//   - [InMemoryArchiveReader.Superclass]
type InMemoryArchiveReader struct {
	objectivec.Object
}

// InMemoryArchiveReaderFromID constructs a [InMemoryArchiveReader] from an objc.ID.
func InMemoryArchiveReaderFromID(id objc.ID) InMemoryArchiveReader {
	return InMemoryArchiveReader{objectivec.Object{ID: id}}
}

// Ensure InMemoryArchiveReader implements IInMemoryArchiveReader.
var _ IInMemoryArchiveReader = InMemoryArchiveReader{}

// An interface definition for the [InMemoryArchiveReader] class.
//
// # Methods
//
//   - [IInMemoryArchiveReader.CopyLayerShapesToContainer]
//   - [IInMemoryArchiveReader.LayerInfos]
//   - [IInMemoryArchiveReader.LoadUpdatableParams]
//   - [IInMemoryArchiveReader.ModelPath]
//   - [IInMemoryArchiveReader.TransformParams]
//   - [IInMemoryArchiveReader.DebugDescription]
//   - [IInMemoryArchiveReader.Description]
//   - [IInMemoryArchiveReader.Hash]
//   - [IInMemoryArchiveReader.Superclass]
type IInMemoryArchiveReader interface {
	objectivec.IObject

	// Topic: Methods

	CopyLayerShapesToContainer(container objectivec.IObject)
	LayerInfos() foundation.INSArray
	LoadUpdatableParams(params []objectivec.IObject) objectivec.IObject
	ModelPath() string
	TransformParams() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (i InMemoryArchiveReader) Init() InMemoryArchiveReader {
	rv := objc.Send[InMemoryArchiveReader](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i InMemoryArchiveReader) Autorelease() InMemoryArchiveReader {
	rv := objc.Send[InMemoryArchiveReader](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewInMemoryArchiveReader creates a new InMemoryArchiveReader instance.
func NewInMemoryArchiveReader() InMemoryArchiveReader {
	class := getInMemoryArchiveReaderClass()
	rv := objc.Send[InMemoryArchiveReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewInMemoryArchiveReaderWithNetwork(network unsafe.Pointer) InMemoryArchiveReader {
	instance := getInMemoryArchiveReaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return InMemoryArchiveReaderFromID(rv)
}

func (i InMemoryArchiveReader) CopyLayerShapesToContainer(container objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("copyLayerShapesToContainer:"), container)
}
func (i InMemoryArchiveReader) LoadUpdatableParams(params []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("loadUpdatableParams:"), objectivec.IObjectSliceToNSArray(params))
	return objectivec.Object{ID: rv}
}
func (i InMemoryArchiveReader) TransformParams() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("transformParams"))
	return objectivec.Object{ID: rv}
}

func (_InMemoryArchiveReaderClass InMemoryArchiveReaderClass) ModelName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_InMemoryArchiveReaderClass.class), objc.Sel("modelName"))
	return objectivec.Object{ID: rv}
}
func (_InMemoryArchiveReaderClass InMemoryArchiveReaderClass) ReaderFromArchiverError(archiver unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_InMemoryArchiveReaderClass.class), objc.Sel("readerFromArchiver:error:"), archiver, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (i InMemoryArchiveReader) DebugDescription() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (i InMemoryArchiveReader) Description() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (i InMemoryArchiveReader) Hash() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("hash"))
	return rv
}
func (i InMemoryArchiveReader) LayerInfos() foundation.INSArray {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("layerInfos"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (i InMemoryArchiveReader) ModelPath() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("modelPath"))
	return foundation.NSStringFromID(rv).String()
}
func (i InMemoryArchiveReader) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](i.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
