// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataPoint] class.
var (
	_ETDataPointClass     ETDataPointClass
	_ETDataPointClassOnce sync.Once
)

func getETDataPointClass() ETDataPointClass {
	_ETDataPointClassOnce.Do(func() {
		_ETDataPointClass = ETDataPointClass{class: objc.GetClass("ETDataPoint")}
	})
	return _ETDataPointClass
}

// GetETDataPointClass returns the class object for ETDataPoint.
func GetETDataPointClass() ETDataPointClass {
	return getETDataPointClass()
}

type ETDataPointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataPointClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataPointClass) Alloc() ETDataPoint {
	rv := objc.Send[ETDataPoint](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETDataPoint.BufferWithKey]
//   - [ETDataPoint.GetSampleData]
//   - [ETDataPoint.ImageWithKey]
//   - [ETDataPoint.IterateBuffersByKey]
//   - [ETDataPoint.SetDataSizeForKeyFreeWhenDone]
//   - [ETDataPoint.SetImageForKey]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataPoint
type ETDataPoint struct {
	objectivec.Object
}

// ETDataPointFromID constructs a [ETDataPoint] from an objc.ID.
func ETDataPointFromID(id objc.ID) ETDataPoint {
	return ETDataPoint{objectivec.Object{ID: id}}
}

// Ensure ETDataPoint implements IETDataPoint.
var _ IETDataPoint = ETDataPoint{}

// An interface definition for the [ETDataPoint] class.
//
// # Methods
//
//   - [IETDataPoint.BufferWithKey]
//   - [IETDataPoint.GetSampleData]
//   - [IETDataPoint.ImageWithKey]
//   - [IETDataPoint.IterateBuffersByKey]
//   - [IETDataPoint.SetDataSizeForKeyFreeWhenDone]
//   - [IETDataPoint.SetImageForKey]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataPoint
type IETDataPoint interface {
	objectivec.IObject

	// Topic: Methods

	BufferWithKey(key objectivec.IObject) unsafe.Pointer
	GetSampleData() objectivec.IObject
	ImageWithKey(key objectivec.IObject) unsafe.Pointer
	IterateBuffersByKey(key VoidHandler)
	SetDataSizeForKeyFreeWhenDone(data unsafe.Pointer, size uint64, key objectivec.IObject, done bool)
	SetImageForKey(image unsafe.Pointer, key objectivec.IObject)
}

// Init initializes the instance.
func (e ETDataPoint) Init() ETDataPoint {
	rv := objc.Send[ETDataPoint](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataPoint) Autorelease() ETDataPoint {
	rv := objc.Send[ETDataPoint](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataPoint creates a new ETDataPoint instance.
func NewETDataPoint() ETDataPoint {
	class := getETDataPointClass()
	rv := objc.Send[ETDataPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataPoint/bufferWithKey:
func (e ETDataPoint) BufferWithKey(key objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("bufferWithKey:"), key)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataPoint/getSampleData
func (e ETDataPoint) GetSampleData() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getSampleData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETDataPoint/imageWithKey:
func (e ETDataPoint) ImageWithKey(key objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("imageWithKey:"), key)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataPoint/iterateBuffersByKey:
func (e ETDataPoint) IterateBuffersByKey(key VoidHandler) {
	_block0, _ := NewVoidBlock(key)
	objc.Send[objc.ID](e.ID, objc.Sel("iterateBuffersByKey:"), _block0)
}

// See: https://developer.apple.com/documentation/Espresso/ETDataPoint/setData:size:forKey:freeWhenDone:
func (e ETDataPoint) SetDataSizeForKeyFreeWhenDone(data unsafe.Pointer, size uint64, key objectivec.IObject, done bool) {
	objc.Send[objc.ID](e.ID, objc.Sel("setData:size:forKey:freeWhenDone:"), data, size, key, done)
}

// See: https://developer.apple.com/documentation/Espresso/ETDataPoint/setImage:forKey:
func (e ETDataPoint) SetImageForKey(image unsafe.Pointer, key objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("setImage:forKey:"), image, key)
}

// IterateBuffersByKeySync is a synchronous wrapper around [ETDataPoint.IterateBuffersByKey].
// It blocks until the completion handler fires or the context is cancelled.
func (e ETDataPoint) IterateBuffersByKeySync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	e.IterateBuffersByKey(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
