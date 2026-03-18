// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4CommandBufferOptions] class.
var (
	_MTL4CommandBufferOptionsClass     MTL4CommandBufferOptionsClass
	_MTL4CommandBufferOptionsClassOnce sync.Once
)

func getMTL4CommandBufferOptionsClass() MTL4CommandBufferOptionsClass {
	_MTL4CommandBufferOptionsClassOnce.Do(func() {
		_MTL4CommandBufferOptionsClass = MTL4CommandBufferOptionsClass{class: objc.GetClass("MTL4CommandBufferOptions")}
	})
	return _MTL4CommandBufferOptionsClass
}

// GetMTL4CommandBufferOptionsClass returns the class object for MTL4CommandBufferOptions.
func GetMTL4CommandBufferOptionsClass() MTL4CommandBufferOptionsClass {
	return getMTL4CommandBufferOptionsClass()
}

type MTL4CommandBufferOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4CommandBufferOptionsClass) Alloc() MTL4CommandBufferOptions {
	rv := objc.Send[MTL4CommandBufferOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Options to configure a command buffer before encoding work into it.
//
// # Instance Properties
//
//   - [MTL4CommandBufferOptions.LogState]: Contains information related to shader logging.
//   - [MTL4CommandBufferOptions.SetLogState]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBufferOptions
type MTL4CommandBufferOptions struct {
	objectivec.Object
}

// MTL4CommandBufferOptionsFromID constructs a [MTL4CommandBufferOptions] from an objc.ID.
//
// Options to configure a command buffer before encoding work into it.
func MTL4CommandBufferOptionsFromID(id objc.ID) MTL4CommandBufferOptions {
	return MTL4CommandBufferOptions{objectivec.Object{ID: id}}
}
// NOTE: MTL4CommandBufferOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4CommandBufferOptions] class.
//
// # Instance Properties
//
//   - [IMTL4CommandBufferOptions.LogState]: Contains information related to shader logging.
//   - [IMTL4CommandBufferOptions.SetLogState]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBufferOptions
type IMTL4CommandBufferOptions interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Contains information related to shader logging.
	LogState() MTLLogState
	SetLogState(value MTLLogState)

	MTL4CommandQueueErrorDomain() string
}

// Init initializes the instance.
func (m MTL4CommandBufferOptions) Init() MTL4CommandBufferOptions {
	rv := objc.Send[MTL4CommandBufferOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4CommandBufferOptions) Autorelease() MTL4CommandBufferOptions {
	rv := objc.Send[MTL4CommandBufferOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandBufferOptions creates a new MTL4CommandBufferOptions instance.
func NewMTL4CommandBufferOptions() MTL4CommandBufferOptions {
	class := getMTL4CommandBufferOptionsClass()
	rv := objc.Send[MTL4CommandBufferOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Contains information related to shader logging.
//
// # Discussion
// 
// To enable shader logging, call [BeginCommandBufferWithAllocatorOptions]
// with an instance of [MTL4CommandBufferOptions] that contains a non-`nil`
// [MTLLogState] instance in this property.
// 
// Shader functions log messages until the command buffer ends.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBufferOptions/logState
func (m MTL4CommandBufferOptions) LogState() MTLLogState {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("logState"))
	return MTLLogStateObjectFromID(rv)
}
func (m MTL4CommandBufferOptions) SetLogState(value MTLLogState) {
	objc.Send[struct{}](m.ID, objc.Sel("setLogState:"), value)
}

// See: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m MTL4CommandBufferOptions) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

