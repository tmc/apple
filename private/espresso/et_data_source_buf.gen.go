// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataSourceBuf] class.
var (
	_ETDataSourceBufClass     ETDataSourceBufClass
	_ETDataSourceBufClassOnce sync.Once
)

func getETDataSourceBufClass() ETDataSourceBufClass {
	_ETDataSourceBufClassOnce.Do(func() {
		_ETDataSourceBufClass = ETDataSourceBufClass{class: objc.GetClass("ETDataSourceBuf")}
	})
	return _ETDataSourceBufClass
}

// GetETDataSourceBufClass returns the class object for ETDataSourceBuf.
func GetETDataSourceBufClass() ETDataSourceBufClass {
	return getETDataSourceBufClass()
}

type ETDataSourceBufClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataSourceBufClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataSourceBufClass) Alloc() ETDataSourceBuf {
	rv := objc.SendIfResponds[ETDataSourceBuf](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETDataSourceBuf.DataAtIndexKey]
//   - [ETDataSourceBuf.DataPointAtIndex]
//   - [ETDataSourceBuf.NumberOfDataPoints]
//   - [ETDataSourceBuf.SetBlobsNumberOfDataPointsNonBatches]
type ETDataSourceBuf struct {
	objectivec.Object
}

// ETDataSourceBufFromID constructs a [ETDataSourceBuf] from an objc.ID.
func ETDataSourceBufFromID(id objc.ID) ETDataSourceBuf {
	return ETDataSourceBuf{objectivec.Object{ID: id}}
}

// Ensure ETDataSourceBuf implements IETDataSourceBuf.
var _ IETDataSourceBuf = ETDataSourceBuf{}

// An interface definition for the [ETDataSourceBuf] class.
//
// # Methods
//
//   - [IETDataSourceBuf.DataAtIndexKey]
//   - [IETDataSourceBuf.DataPointAtIndex]
//   - [IETDataSourceBuf.NumberOfDataPoints]
//   - [IETDataSourceBuf.SetBlobsNumberOfDataPointsNonBatches]
type IETDataSourceBuf interface {
	objectivec.IObject

	// Topic: Methods

	DataAtIndexKey(index uint64, key unsafe.Pointer) unsafe.Pointer
	DataPointAtIndex(index int) objectivec.IObject
	NumberOfDataPoints() int
	SetBlobsNumberOfDataPointsNonBatches(blobs unsafe.Pointer, points int, batches unsafe.Pointer)
}

// Init initializes the instance.
func (e ETDataSourceBuf) Init() ETDataSourceBuf {
	rv := objc.SendIfResponds[ETDataSourceBuf](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataSourceBuf) Autorelease() ETDataSourceBuf {
	rv := objc.SendIfResponds[ETDataSourceBuf](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataSourceBuf creates a new ETDataSourceBuf instance.
func NewETDataSourceBuf() ETDataSourceBuf {
	class := getETDataSourceBufClass()
	rv := objc.SendIfResponds[ETDataSourceBuf](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (e ETDataSourceBuf) DataAtIndexKey(index uint64, key unsafe.Pointer) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("dataAtIndex:key:"), index, key)
	return rv
}
func (e ETDataSourceBuf) DataPointAtIndex(index int) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("dataPointAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (e ETDataSourceBuf) NumberOfDataPoints() int {
	rv := objc.SendIfResponds[int](e.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
func (e ETDataSourceBuf) SetBlobsNumberOfDataPointsNonBatches(blobs unsafe.Pointer, points int, batches unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("setBlobs:numberOfDataPoints:nonBatches:"), blobs, points, batches)
}
