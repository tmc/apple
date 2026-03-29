// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioQueue] class.
var (
	_VZVirtioQueueClass     VZVirtioQueueClass
	_VZVirtioQueueClassOnce sync.Once
)

func getVZVirtioQueueClass() VZVirtioQueueClass {
	_VZVirtioQueueClassOnce.Do(func() {
		_VZVirtioQueueClass = VZVirtioQueueClass{class: objc.GetClass("_VZVirtioQueue")}
	})
	return _VZVirtioQueueClass
}

// GetVZVirtioQueueClass returns the class object for _VZVirtioQueue.
func GetVZVirtioQueueClass() VZVirtioQueueClass {
	return getVZVirtioQueueClass()
}

type VZVirtioQueueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioQueueClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioQueueClass) Alloc() VZVirtioQueue {
	rv := objc.Send[VZVirtioQueue](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtioQueue.NextElement]
//   - [VZVirtioQueue.QueueIndex]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueue
type VZVirtioQueue struct {
	objectivec.Object
}

// VZVirtioQueueFromID constructs a [VZVirtioQueue] from an objc.ID.
func VZVirtioQueueFromID(id objc.ID) VZVirtioQueue {
	return VZVirtioQueue{objectivec.Object{ID: id}}
}
// Ensure VZVirtioQueue implements IVZVirtioQueue.
var _ IVZVirtioQueue = VZVirtioQueue{}

// An interface definition for the [VZVirtioQueue] class.
//
// # Methods
//
//   - [IVZVirtioQueue.NextElement]
//   - [IVZVirtioQueue.QueueIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueue
type IVZVirtioQueue interface {
	objectivec.IObject

	// Topic: Methods

	NextElement() objectivec.IObject
	QueueIndex() uint16
}

// Init initializes the instance.
func (v VZVirtioQueue) Init() VZVirtioQueue {
	rv := objc.Send[VZVirtioQueue](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioQueue) Autorelease() VZVirtioQueue {
	rv := objc.Send[VZVirtioQueue](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioQueue creates a new VZVirtioQueue instance.
func NewVZVirtioQueue() VZVirtioQueue {
	class := getVZVirtioQueueClass()
	rv := objc.Send[VZVirtioQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueue/nextElement
func (v VZVirtioQueue) NextElement() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nextElement"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioQueue/queueIndex
func (v VZVirtioQueue) QueueIndex() uint16 {
	rv := objc.Send[uint16](v.ID, objc.Sel("queueIndex"))
	return rv
}

