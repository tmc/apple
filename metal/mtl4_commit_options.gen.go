// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4CommitOptions] class.
var (
	_MTL4CommitOptionsClass     MTL4CommitOptionsClass
	_MTL4CommitOptionsClassOnce sync.Once
)

func getMTL4CommitOptionsClass() MTL4CommitOptionsClass {
	_MTL4CommitOptionsClassOnce.Do(func() {
		_MTL4CommitOptionsClass = MTL4CommitOptionsClass{class: objc.GetClass("MTL4CommitOptions")}
	})
	return _MTL4CommitOptionsClass
}

// GetMTL4CommitOptionsClass returns the class object for MTL4CommitOptions.
func GetMTL4CommitOptionsClass() MTL4CommitOptionsClass {
	return getMTL4CommitOptionsClass()
}

type MTL4CommitOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4CommitOptionsClass) Alloc() MTL4CommitOptions {
	rv := objc.Send[MTL4CommitOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Represents options to configure a commit operation on a command queue.
//
// # Overview
// 
// You pass these options as a parameter when you call [CommitCountOptions].
// 
// - Note Instances of this class are not thread-safe. If your app modifies a
// shared commit options instance from multiple threads simultaneously, you
// are responsible for providing external synchronization.
//
// # Instance Methods
//
//   - [MTL4CommitOptions.AddFeedbackHandler]: Registers a commit feedback handler that Metal calls with feedback data when available.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitOptions
type MTL4CommitOptions struct {
	objectivec.Object
}

// MTL4CommitOptionsFromID constructs a [MTL4CommitOptions] from an objc.ID.
//
// Represents options to configure a commit operation on a command queue.
func MTL4CommitOptionsFromID(id objc.ID) MTL4CommitOptions {
	return MTL4CommitOptions{objectivec.Object{id}}
}
// NOTE: MTL4CommitOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4CommitOptions] class.
//
// # Instance Methods
//
//   - [IMTL4CommitOptions.AddFeedbackHandler]: Registers a commit feedback handler that Metal calls with feedback data when available.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitOptions
type IMTL4CommitOptions interface {
	objectivec.IObject

	// Topic: Instance Methods

	// Registers a commit feedback handler that Metal calls with feedback data when available.
	AddFeedbackHandler(block MTL4CommitFeedbackHandler)

	MTL4CommandQueueErrorDomain() string
}





// Init initializes the instance.
func (m MTL4CommitOptions) Init() MTL4CommitOptions {
	rv := objc.Send[MTL4CommitOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4CommitOptions) Autorelease() MTL4CommitOptions {
	rv := objc.Send[MTL4CommitOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommitOptions creates a new MTL4CommitOptions instance.
func NewMTL4CommitOptions() MTL4CommitOptions {
	class := getMTL4CommitOptionsClass()
	rv := objc.Send[MTL4CommitOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Registers a commit feedback handler that Metal calls with feedback data
// when available.
//
// block: [MTL4CommitFeedbackHandler] that Metal invokes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitOptions/addFeedbackHandler(_:)
func (m MTL4CommitOptions) AddFeedbackHandler(block MTL4CommitFeedbackHandler) {
	objc.Send[objc.ID](m.ID, objc.Sel("addFeedbackHandler:"), block)
}












// See: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m MTL4CommitOptions) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}























