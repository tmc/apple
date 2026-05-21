// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OnDiskArchiveReader] class.
var (
	_OnDiskArchiveReaderClass     OnDiskArchiveReaderClass
	_OnDiskArchiveReaderClassOnce sync.Once
)

func getOnDiskArchiveReaderClass() OnDiskArchiveReaderClass {
	_OnDiskArchiveReaderClassOnce.Do(func() {
		_OnDiskArchiveReaderClass = OnDiskArchiveReaderClass{class: objc.GetClass("_OnDiskArchiveReader")}
	})
	return _OnDiskArchiveReaderClass
}

// GetOnDiskArchiveReaderClass returns the class object for _OnDiskArchiveReader.
func GetOnDiskArchiveReaderClass() OnDiskArchiveReaderClass {
	return getOnDiskArchiveReaderClass()
}

type OnDiskArchiveReaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OnDiskArchiveReaderClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OnDiskArchiveReaderClass) Alloc() OnDiskArchiveReader {
	rv := objc.Send[OnDiskArchiveReader](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [OnDiskArchiveReader.CopyLayerShapesToContainer]
//   - [OnDiskArchiveReader.LayerInfos]
//   - [OnDiskArchiveReader.LoadUpdatableParams]
//   - [OnDiskArchiveReader.ModelPath]
//   - [OnDiskArchiveReader.NetJson]
//   - [OnDiskArchiveReader.ShapeJson]
//   - [OnDiskArchiveReader.TransformParams]
//   - [OnDiskArchiveReader.InitWithNetJsonShapeJsonModelPath]
//   - [OnDiskArchiveReader.DebugDescription]
//   - [OnDiskArchiveReader.Description]
//   - [OnDiskArchiveReader.Hash]
//   - [OnDiskArchiveReader.Superclass]
type OnDiskArchiveReader struct {
	objectivec.Object
}

// OnDiskArchiveReaderFromID constructs a [OnDiskArchiveReader] from an objc.ID.
func OnDiskArchiveReaderFromID(id objc.ID) OnDiskArchiveReader {
	return OnDiskArchiveReader{objectivec.Object{ID: id}}
}

// Ensure OnDiskArchiveReader implements IOnDiskArchiveReader.
var _ IOnDiskArchiveReader = OnDiskArchiveReader{}

// An interface definition for the [OnDiskArchiveReader] class.
//
// # Methods
//
//   - [IOnDiskArchiveReader.CopyLayerShapesToContainer]
//   - [IOnDiskArchiveReader.LayerInfos]
//   - [IOnDiskArchiveReader.LoadUpdatableParams]
//   - [IOnDiskArchiveReader.ModelPath]
//   - [IOnDiskArchiveReader.NetJson]
//   - [IOnDiskArchiveReader.ShapeJson]
//   - [IOnDiskArchiveReader.TransformParams]
//   - [IOnDiskArchiveReader.InitWithNetJsonShapeJsonModelPath]
//   - [IOnDiskArchiveReader.DebugDescription]
//   - [IOnDiskArchiveReader.Description]
//   - [IOnDiskArchiveReader.Hash]
//   - [IOnDiskArchiveReader.Superclass]
type IOnDiskArchiveReader interface {
	objectivec.IObject

	// Topic: Methods

	CopyLayerShapesToContainer(container objectivec.IObject)
	LayerInfos() foundation.INSArray
	LoadUpdatableParams(params []objectivec.IObject) objectivec.IObject
	ModelPath() string
	NetJson() foundation.INSDictionary
	ShapeJson() foundation.INSDictionary
	TransformParams() objectivec.IObject
	InitWithNetJsonShapeJsonModelPath(json objectivec.IObject, json2 objectivec.IObject, path objectivec.IObject) OnDiskArchiveReader
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (o OnDiskArchiveReader) Init() OnDiskArchiveReader {
	rv := objc.Send[OnDiskArchiveReader](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OnDiskArchiveReader) Autorelease() OnDiskArchiveReader {
	rv := objc.Send[OnDiskArchiveReader](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOnDiskArchiveReader creates a new OnDiskArchiveReader instance.
func NewOnDiskArchiveReader() OnDiskArchiveReader {
	class := getOnDiskArchiveReaderClass()
	rv := objc.Send[OnDiskArchiveReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewOnDiskArchiveReaderWithNetJsonShapeJsonModelPath(json objectivec.IObject, json2 objectivec.IObject, path objectivec.IObject) OnDiskArchiveReader {
	instance := getOnDiskArchiveReaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetJson:shapeJson:modelPath:"), json, json2, path)
	return OnDiskArchiveReaderFromID(rv)
}

func (o OnDiskArchiveReader) CopyLayerShapesToContainer(container objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("copyLayerShapesToContainer:"), container)
}
func (o OnDiskArchiveReader) LoadUpdatableParams(params []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("loadUpdatableParams:"), objectivec.IObjectSliceToNSArray(params))
	return objectivec.Object{ID: rv}
}
func (o OnDiskArchiveReader) TransformParams() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("transformParams"))
	return objectivec.Object{ID: rv}
}
func (o OnDiskArchiveReader) InitWithNetJsonShapeJsonModelPath(json objectivec.IObject, json2 objectivec.IObject, path objectivec.IObject) OnDiskArchiveReader {
	rv := objc.Send[OnDiskArchiveReader](o.ID, objc.Sel("initWithNetJson:shapeJson:modelPath:"), json, json2, path)
	return rv
}

func (_OnDiskArchiveReaderClass OnDiskArchiveReaderClass) ModelName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_OnDiskArchiveReaderClass.class), objc.Sel("modelName"))
	return objectivec.Object{ID: rv}
}
func (_OnDiskArchiveReaderClass OnDiskArchiveReaderClass) ModelNetFileName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_OnDiskArchiveReaderClass.class), objc.Sel("modelNetFileName"))
	return objectivec.Object{ID: rv}
}
func (_OnDiskArchiveReaderClass OnDiskArchiveReaderClass) ModelShapeFileName() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_OnDiskArchiveReaderClass.class), objc.Sel("modelShapeFileName"))
	return objectivec.Object{ID: rv}
}
func (_OnDiskArchiveReaderClass OnDiskArchiveReaderClass) ParseCompiledNetworkBlobWithNameArchiveError(name objectivec.IObject, archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_OnDiskArchiveReaderClass.class), objc.Sel("parseCompiledNetworkBlobWithName:archive:error:"), name, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_OnDiskArchiveReaderClass OnDiskArchiveReaderClass) ReaderFromArchiverError(archiver unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_OnDiskArchiveReaderClass.class), objc.Sel("readerFromArchiver:error:"), archiver, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (o OnDiskArchiveReader) DebugDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (o OnDiskArchiveReader) Description() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (o OnDiskArchiveReader) Hash() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("hash"))
	return rv
}
func (o OnDiskArchiveReader) LayerInfos() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("layerInfos"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o OnDiskArchiveReader) ModelPath() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modelPath"))
	return foundation.NSStringFromID(rv).String()
}
func (o OnDiskArchiveReader) NetJson() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("netJson"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (o OnDiskArchiveReader) ShapeJson() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("shapeJson"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (o OnDiskArchiveReader) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](o.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
