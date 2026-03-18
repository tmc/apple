// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLCaptureManager] class.
var (
	_MTLCaptureManagerClass     MTLCaptureManagerClass
	_MTLCaptureManagerClassOnce sync.Once
)

func getMTLCaptureManagerClass() MTLCaptureManagerClass {
	_MTLCaptureManagerClassOnce.Do(func() {
		_MTLCaptureManagerClass = MTLCaptureManagerClass{class: objc.GetClass("MTLCaptureManager")}
	})
	return _MTLCaptureManagerClass
}

// GetMTLCaptureManagerClass returns the class object for MTLCaptureManager.
func GetMTLCaptureManagerClass() MTLCaptureManagerClass {
	return getMTLCaptureManagerClass()
}

type MTLCaptureManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLCaptureManagerClass) Alloc() MTLCaptureManager {
	rv := objc.Send[MTLCaptureManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An instance you use to capture Metal command data in your app.
//
// # Overview
// 
// A capture manager works with the frame capture feature to:
// 
// - Capture data about Metal commands programmatically. See [Capturing a
// Metal workload programmatically]. - Only capture commands that apply to a
// specific [MTLDevice], command queue, or [MTLCaptureScope] instance. -
// Assign a default [MTLCaptureScope] instance for captures you create in
// Xcode by clicking the Capture GPU workload button in the debug bar, which
// has an icon with the Metal logo.
// 
// The Metal debugger requires you to enable GPU Frame Capture in your project
// settings; see [Capturing a Metal workload in Xcode].
// 
// For more information about Metal frame capture, see [Metal debugger].
//
// [Capturing a Metal workload in Xcode]: https://developer.apple.com/documentation/Xcode/Capturing-a-Metal-workload-in-Xcode
// [Capturing a Metal workload programmatically]: https://developer.apple.com/documentation/Xcode/Capturing-a-Metal-workload-programmatically
// [Metal debugger]: https://developer.apple.com/documentation/Xcode/Metal-debugger
//
// # Querying support for a capture destination
//
//   - [MTLCaptureManager.SupportsDestination]: Checks to see whether a particular capture destination is supported.
//
// # Creating a capture scope
//
//   - [MTLCaptureManager.NewCaptureScopeWithDevice]: Creates a capture scope for commands submitted to a specific device object.
//   - [MTLCaptureManager.NewCaptureScopeWithCommandQueue]: Creates a capture scope for commands submitted to a specific command queue.
//   - [MTLCaptureManager.DefaultCaptureScope]: The capture scope to use when a capture is initiated in Xcode.
//   - [MTLCaptureManager.SetDefaultCaptureScope]
//
// # Starting capture
//
//   - [MTLCaptureManager.StartCaptureWithDescriptorError]: Starts capturing any of your app’s Metal commands, with the capture session defined by a descriptor object.
//
// # Stopping capture
//
//   - [MTLCaptureManager.StopCapture]: Stops capturing Metal commands.
//
// # Monitoring capture
//
//   - [MTLCaptureManager.IsCapturing]: A Boolean value that indicates whether Metal commands are being captured.
//
// # Instance Methods
//
//   - [MTLCaptureManager.NewCaptureScopeWithMTL4CommandQueue]
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager
type MTLCaptureManager struct {
	objectivec.Object
}

// MTLCaptureManagerFromID constructs a [MTLCaptureManager] from an objc.ID.
//
// An instance you use to capture Metal command data in your app.
func MTLCaptureManagerFromID(id objc.ID) MTLCaptureManager {
	return MTLCaptureManager{objectivec.Object{ID: id}}
}
// NOTE: MTLCaptureManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLCaptureManager] class.
//
// # Querying support for a capture destination
//
//   - [IMTLCaptureManager.SupportsDestination]: Checks to see whether a particular capture destination is supported.
//
// # Creating a capture scope
//
//   - [IMTLCaptureManager.NewCaptureScopeWithDevice]: Creates a capture scope for commands submitted to a specific device object.
//   - [IMTLCaptureManager.NewCaptureScopeWithCommandQueue]: Creates a capture scope for commands submitted to a specific command queue.
//   - [IMTLCaptureManager.DefaultCaptureScope]: The capture scope to use when a capture is initiated in Xcode.
//   - [IMTLCaptureManager.SetDefaultCaptureScope]
//
// # Starting capture
//
//   - [IMTLCaptureManager.StartCaptureWithDescriptorError]: Starts capturing any of your app’s Metal commands, with the capture session defined by a descriptor object.
//
// # Stopping capture
//
//   - [IMTLCaptureManager.StopCapture]: Stops capturing Metal commands.
//
// # Monitoring capture
//
//   - [IMTLCaptureManager.IsCapturing]: A Boolean value that indicates whether Metal commands are being captured.
//
// # Instance Methods
//
//   - [IMTLCaptureManager.NewCaptureScopeWithMTL4CommandQueue]
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager
type IMTLCaptureManager interface {
	objectivec.IObject

	// Topic: Querying support for a capture destination

	// Checks to see whether a particular capture destination is supported.
	SupportsDestination(destination MTLCaptureDestination) bool

	// Topic: Creating a capture scope

	// Creates a capture scope for commands submitted to a specific device object.
	NewCaptureScopeWithDevice(device MTLDevice) MTLCaptureScope
	// Creates a capture scope for commands submitted to a specific command queue.
	NewCaptureScopeWithCommandQueue(commandQueue MTLCommandQueue) MTLCaptureScope
	// The capture scope to use when a capture is initiated in Xcode.
	DefaultCaptureScope() MTLCaptureScope
	SetDefaultCaptureScope(value MTLCaptureScope)

	// Topic: Starting capture

	// Starts capturing any of your app’s Metal commands, with the capture session defined by a descriptor object.
	StartCaptureWithDescriptorError(descriptor IMTLCaptureDescriptor) (bool, error)

	// Topic: Stopping capture

	// Stops capturing Metal commands.
	StopCapture()

	// Topic: Monitoring capture

	// A Boolean value that indicates whether Metal commands are being captured.
	IsCapturing() bool

	// Topic: Instance Methods

	NewCaptureScopeWithMTL4CommandQueue(commandQueue MTL4CommandQueue) MTLCaptureScope
}

// Init initializes the instance.
func (c MTLCaptureManager) Init() MTLCaptureManager {
	rv := objc.Send[MTLCaptureManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLCaptureManager) Autorelease() MTLCaptureManager {
	rv := objc.Send[MTLCaptureManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLCaptureManager creates a new MTLCaptureManager instance.
func NewMTLCaptureManager() MTLCaptureManager {
	class := getMTLCaptureManagerClass()
	rv := objc.Send[MTLCaptureManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Checks to see whether a particular capture destination is supported.
//
// destination: The destination to test for.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/supportsDestination(_:)
func (c MTLCaptureManager) SupportsDestination(destination MTLCaptureDestination) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportsDestination:"), destination)
	return rv
}

// Creates a capture scope for commands submitted to a specific device object.
//
// device: The device object whose commands you want to capture.
//
// # Discussion
// 
// The capture scope captures commands in command buffers created on any
// command queues created by the device object.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/makeCaptureScope(device:)
func (c MTLCaptureManager) NewCaptureScopeWithDevice(device MTLDevice) MTLCaptureScope {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("newCaptureScopeWithDevice:"), device)
	return MTLCaptureScopeObjectFromID(rv)
}

// Creates a capture scope for commands submitted to a specific command queue.
//
// commandQueue: The command queue whose commands you want to capture.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/makeCaptureScope(commandQueue:)-1rozd
func (c MTLCaptureManager) NewCaptureScopeWithCommandQueue(commandQueue MTLCommandQueue) MTLCaptureScope {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("newCaptureScopeWithCommandQueue:"), commandQueue)
	return MTLCaptureScopeObjectFromID(rv)
}

// Starts capturing any of your app’s Metal commands, with the capture
// session defined by a descriptor object.
//
// descriptor: A description of the capture session to create.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/startCapture(with:)
func (c MTLCaptureManager) StartCaptureWithDescriptorError(descriptor IMTLCaptureDescriptor) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("startCaptureWithDescriptor:error:"), descriptor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startCaptureWithDescriptor:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Stops capturing Metal commands.
//
// # Discussion
// 
// Calling this method stops a capture that was started manually in Xcode or
// programmatically by calling one of the methods on [MTLCaptureManager].
// 
// When using a custom capture scope, calling this function preempts any
// [EndScope] demarcations of the capture scope.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/stopCapture()
func (c MTLCaptureManager) StopCapture() {
	objc.Send[objc.ID](c.ID, objc.Sel("stopCapture"))
}

//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/makeCaptureScope(commandQueue:)-9wie3
func (c MTLCaptureManager) NewCaptureScopeWithMTL4CommandQueue(commandQueue MTL4CommandQueue) MTLCaptureScope {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("newCaptureScopeWithMTL4CommandQueue:"), commandQueue)
	return MTLCaptureScopeObjectFromID(rv)
}

// Provides the shared capture manager for your Metal app.
//
// # Discussion
// 
// There is only one capture manager per process.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/shared()
func (_MTLCaptureManagerClass MTLCaptureManagerClass) SharedCaptureManager() MTLCaptureManager {
	rv := objc.Send[objc.ID](objc.ID(_MTLCaptureManagerClass.class), objc.Sel("sharedCaptureManager"))
	return MTLCaptureManagerFromID(rv)
}

// The capture scope to use when a capture is initiated in Xcode.
//
// # Discussion
// 
// Use this property to specify a default capture scope for Xcode to use when
// the user presses the capture button. You can still long-press the button to
// select a different capture scope.
// 
// The default value is `nil.` When the value is `nil`, the capture scope is
// defined by drawable presentation boundaries; such as those created by calls
// to [PresentDrawable] or [Present].
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/defaultCaptureScope
func (c MTLCaptureManager) DefaultCaptureScope() MTLCaptureScope {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("defaultCaptureScope"))
	return MTLCaptureScopeObjectFromID(rv)
}
func (c MTLCaptureManager) SetDefaultCaptureScope(value MTLCaptureScope) {
	objc.Send[struct{}](c.ID, objc.Sel("setDefaultCaptureScope:"), value)
}

// A Boolean value that indicates whether Metal commands are being captured.
//
// See: https://developer.apple.com/documentation/Metal/MTLCaptureManager/isCapturing
func (c MTLCaptureManager) IsCapturing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCapturing"))
	return rv
}

