// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETBufferDataSource] class.
var (
	_ETBufferDataSourceClass     ETBufferDataSourceClass
	_ETBufferDataSourceClassOnce sync.Once
)

func getETBufferDataSourceClass() ETBufferDataSourceClass {
	_ETBufferDataSourceClassOnce.Do(func() {
		_ETBufferDataSourceClass = ETBufferDataSourceClass{class: objc.GetClass("_ETBufferDataSource")}
	})
	return _ETBufferDataSourceClass
}

// GetETBufferDataSourceClass returns the class object for _ETBufferDataSource.
func GetETBufferDataSourceClass() ETBufferDataSourceClass {
	return getETBufferDataSourceClass()
}

type ETBufferDataSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETBufferDataSourceClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETBufferDataSourceClass) Alloc() ETBufferDataSource {
	rv := objc.SendIfResponds[ETBufferDataSource](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETBufferDataSource.BatchSize]
//   - [ETBufferDataSource.SetBatchSize]
//   - [ETBufferDataSource.BlobShapes]
//   - [ETBufferDataSource.SetBlobShapes]
//   - [ETBufferDataSource.DataAtIndexKey]
//   - [ETBufferDataSource.DataPointAtIndexError]
//   - [ETBufferDataSource.DataStorage]
//   - [ETBufferDataSource.SetDataStorage]
//   - [ETBufferDataSource.NonBatchBlobNames]
//   - [ETBufferDataSource.SetNonBatchBlobNames]
//   - [ETBufferDataSource.NumberOfDataPoints]
//   - [ETBufferDataSource.Number_of_data_points]
//   - [ETBufferDataSource.SetNumber_of_data_points]
//   - [ETBufferDataSource.InitWithBlobShapesNumberOfDataPointsBatchSizeError]
//   - [ETBufferDataSource.DebugDescription]
//   - [ETBufferDataSource.Description]
//   - [ETBufferDataSource.Hash]
//   - [ETBufferDataSource.Superclass]
type ETBufferDataSource struct {
	objectivec.Object
}

// ETBufferDataSourceFromID constructs a [ETBufferDataSource] from an objc.ID.
func ETBufferDataSourceFromID(id objc.ID) ETBufferDataSource {
	return ETBufferDataSource{objectivec.Object{ID: id}}
}

// Ensure ETBufferDataSource implements IETBufferDataSource.
var _ IETBufferDataSource = ETBufferDataSource{}

// An interface definition for the [ETBufferDataSource] class.
//
// # Methods
//
//   - [IETBufferDataSource.BatchSize]
//   - [IETBufferDataSource.SetBatchSize]
//   - [IETBufferDataSource.BlobShapes]
//   - [IETBufferDataSource.SetBlobShapes]
//   - [IETBufferDataSource.DataAtIndexKey]
//   - [IETBufferDataSource.DataPointAtIndexError]
//   - [IETBufferDataSource.DataStorage]
//   - [IETBufferDataSource.SetDataStorage]
//   - [IETBufferDataSource.NonBatchBlobNames]
//   - [IETBufferDataSource.SetNonBatchBlobNames]
//   - [IETBufferDataSource.NumberOfDataPoints]
//   - [IETBufferDataSource.Number_of_data_points]
//   - [IETBufferDataSource.SetNumber_of_data_points]
//   - [IETBufferDataSource.InitWithBlobShapesNumberOfDataPointsBatchSizeError]
//   - [IETBufferDataSource.DebugDescription]
//   - [IETBufferDataSource.Description]
//   - [IETBufferDataSource.Hash]
//   - [IETBufferDataSource.Superclass]
type IETBufferDataSource interface {
	objectivec.IObject

	// Topic: Methods

	BatchSize() uint64
	SetBatchSize(value uint64)
	BlobShapes() unsafe.Pointer
	SetBlobShapes(value unsafe.Pointer)
	DataAtIndexKey(index uint64, key unsafe.Pointer) unsafe.Pointer
	DataPointAtIndexError(index uint64) (objectivec.IObject, error)
	DataStorage() unsafe.Pointer
	SetDataStorage(value unsafe.Pointer)
	NonBatchBlobNames() unsafe.Pointer
	SetNonBatchBlobNames(value unsafe.Pointer)
	NumberOfDataPoints() uint64
	Number_of_data_points() uint64
	SetNumber_of_data_points(value uint64)
	InitWithBlobShapesNumberOfDataPointsBatchSizeError(shapes unsafe.Pointer, points uint64, size uint64) (ETBufferDataSource, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (e ETBufferDataSource) Init() ETBufferDataSource {
	rv := objc.SendIfResponds[ETBufferDataSource](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETBufferDataSource) Autorelease() ETBufferDataSource {
	rv := objc.SendIfResponds[ETBufferDataSource](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETBufferDataSource creates a new ETBufferDataSource instance.
func NewETBufferDataSource() ETBufferDataSource {
	class := getETBufferDataSourceClass()
	rv := objc.SendIfResponds[ETBufferDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewETBufferDataSourceWithBlobShapesNumberOfDataPointsBatchSizeError(shapes unsafe.Pointer, points uint64, size uint64) (ETBufferDataSource, error) {
	var errorPtr objc.ID
	instance := getETBufferDataSourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBlobShapes:numberOfDataPoints:batchSize:error:"), shapes, points, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETBufferDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return ETBufferDataSource{}, objc.ErrInitFailed
	}
	return ETBufferDataSourceFromID(rv), nil
}

func (e ETBufferDataSource) DataAtIndexKey(index uint64, key unsafe.Pointer) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("dataAtIndex:key:"), index, key)
	return rv
}
func (e ETBufferDataSource) DataPointAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataPointAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (e ETBufferDataSource) NumberOfDataPoints() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
func (e ETBufferDataSource) InitWithBlobShapesNumberOfDataPointsBatchSizeError(shapes unsafe.Pointer, points uint64, size uint64) (ETBufferDataSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithBlobShapes:numberOfDataPoints:batchSize:error:"), shapes, points, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(ETBufferDataSource), foundation.NSErrorFrom(errorPtr)
	}
	return ETBufferDataSourceFromID(rv), nil

}

func (e ETBufferDataSource) BatchSize() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("batchSize"))
	return rv
}
func (e ETBufferDataSource) SetBatchSize(value uint64) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setBatchSize:"), value)
}
func (e ETBufferDataSource) BlobShapes() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("blobShapes"))
	return rv
}
func (e ETBufferDataSource) SetBlobShapes(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setBlobShapes:"), value)
}
func (e ETBufferDataSource) DataStorage() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("dataStorage"))
	return rv
}
func (e ETBufferDataSource) SetDataStorage(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setDataStorage:"), value)
}
func (e ETBufferDataSource) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETBufferDataSource) Description() string {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETBufferDataSource) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("hash"))
	return rv
}
func (e ETBufferDataSource) NonBatchBlobNames() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("nonBatchBlobNames"))
	return rv
}
func (e ETBufferDataSource) SetNonBatchBlobNames(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setNonBatchBlobNames:"), value)
}
func (e ETBufferDataSource) Number_of_data_points() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("number_of_data_points"))
	return rv
}
func (e ETBufferDataSource) SetNumber_of_data_points(value uint64) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setNumber_of_data_points:"), value)
}
func (e ETBufferDataSource) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](e.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
