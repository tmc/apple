// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4CommandQueueDescriptor] class.
var (
	_MTL4CommandQueueDescriptorClass     MTL4CommandQueueDescriptorClass
	_MTL4CommandQueueDescriptorClassOnce sync.Once
)

func getMTL4CommandQueueDescriptorClass() MTL4CommandQueueDescriptorClass {
	_MTL4CommandQueueDescriptorClassOnce.Do(func() {
		_MTL4CommandQueueDescriptorClass = MTL4CommandQueueDescriptorClass{class: objc.GetClass("MTL4CommandQueueDescriptor")}
	})
	return _MTL4CommandQueueDescriptorClass
}

// GetMTL4CommandQueueDescriptorClass returns the class object for MTL4CommandQueueDescriptor.
func GetMTL4CommandQueueDescriptorClass() MTL4CommandQueueDescriptorClass {
	return getMTL4CommandQueueDescriptorClass()
}

type MTL4CommandQueueDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4CommandQueueDescriptorClass) Alloc() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Groups together parameters for the creation of a new command queue.
//
// # Instance Properties
//
//   - [MTL4CommandQueueDescriptor.FeedbackQueue]: Assigns a dispatch queue to which Metal submits feedback notification blocks.
//   - [MTL4CommandQueueDescriptor.SetFeedbackQueue]
//   - [MTL4CommandQueueDescriptor.Label]: Assigns an optional label to the command queue instance for debugging purposes.
//   - [MTL4CommandQueueDescriptor.SetLabel]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor
type MTL4CommandQueueDescriptor struct {
	objectivec.Object
}

// MTL4CommandQueueDescriptorFromID constructs a [MTL4CommandQueueDescriptor] from an objc.ID.
//
// Groups together parameters for the creation of a new command queue.
func MTL4CommandQueueDescriptorFromID(id objc.ID) MTL4CommandQueueDescriptor {
	return MTL4CommandQueueDescriptor{objectivec.Object{id}}
}
// NOTE: MTL4CommandQueueDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4CommandQueueDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4CommandQueueDescriptor.FeedbackQueue]: Assigns a dispatch queue to which Metal submits feedback notification blocks.
//   - [IMTL4CommandQueueDescriptor.SetFeedbackQueue]
//   - [IMTL4CommandQueueDescriptor.Label]: Assigns an optional label to the command queue instance for debugging purposes.
//   - [IMTL4CommandQueueDescriptor.SetLabel]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor
type IMTL4CommandQueueDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Assigns a dispatch queue to which Metal submits feedback notification blocks.
	FeedbackQueue() dispatch.Queue
	SetFeedbackQueue(value dispatch.Queue)
	// Assigns an optional label to the command queue instance for debugging purposes.
	Label() string
	SetLabel(value string)

	MTL4CommandQueueErrorDomain() string
}





// Init initializes the instance.
func (m MTL4CommandQueueDescriptor) Init() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4CommandQueueDescriptor) Autorelease() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandQueueDescriptor creates a new MTL4CommandQueueDescriptor instance.
func NewMTL4CommandQueueDescriptor() MTL4CommandQueueDescriptor {
	class := getMTL4CommandQueueDescriptorClass()
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Assigns a dispatch queue to which Metal submits feedback notification
// blocks.
//
// # Discussion
// 
// When you assign a dispatch queue via this method, Metal requires that the
// queue parameter you provide is a serial queue.
// 
// If you set the value of property to `nil`, the default, Metal allocates an
// internal dispatch queue to service feedback notifications.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor/feedbackQueue
func (m MTL4CommandQueueDescriptor) FeedbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](m.ID, objc.Sel("feedbackQueue"))
	return dispatch.QueueFromHandle(rv)
}
func (m MTL4CommandQueueDescriptor) SetFeedbackQueue(value dispatch.Queue) {
	objc.Send[struct{}](m.ID, objc.Sel("setFeedbackQueue:"), uintptr(value.Handle()))
}



// Assigns an optional label to the command queue instance for debugging
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor/label
func (m MTL4CommandQueueDescriptor) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4CommandQueueDescriptor) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m MTL4CommandQueueDescriptor) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}
























