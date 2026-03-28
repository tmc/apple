// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDisplay] class.
var (
	_VZGraphicsDisplayClass     VZGraphicsDisplayClass
	_VZGraphicsDisplayClassOnce sync.Once
)

func getVZGraphicsDisplayClass() VZGraphicsDisplayClass {
	_VZGraphicsDisplayClassOnce.Do(func() {
		_VZGraphicsDisplayClass = VZGraphicsDisplayClass{class: objc.GetClass("VZGraphicsDisplay")}
	})
	return _VZGraphicsDisplayClass
}

// GetVZGraphicsDisplayClass returns the class object for VZGraphicsDisplay.
func GetVZGraphicsDisplayClass() VZGraphicsDisplayClass {
	return getVZGraphicsDisplayClass()
}

type VZGraphicsDisplayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGraphicsDisplayClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDisplayClass) Alloc() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZGraphicsDisplay._configuration]
//   - [VZGraphicsDisplay._graphicsDevice]
//   - [VZGraphicsDisplay._graphicsOrientation]
//   - [VZGraphicsDisplay._initDetached]
//   - [VZGraphicsDisplay._matchesConfiguration]
//   - [VZGraphicsDisplay._setGraphicsDevice]
//   - [VZGraphicsDisplay._takeScreenshotWithCompletionHandler]
//   - [VZGraphicsDisplay._uuid]
//   - [VZGraphicsDisplay.InitWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid]
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay
type VZGraphicsDisplay struct {
	objectivec.Object
}

// VZGraphicsDisplayFromID constructs a [VZGraphicsDisplay] from an objc.ID.
func VZGraphicsDisplayFromID(id objc.ID) VZGraphicsDisplay {
	return VZGraphicsDisplay{objectivec.Object{ID: id}}
}
// Ensure VZGraphicsDisplay implements IVZGraphicsDisplay.
var _ IVZGraphicsDisplay = VZGraphicsDisplay{}

// An interface definition for the [VZGraphicsDisplay] class.
//
// # Methods
//
//   - [IVZGraphicsDisplay._configuration]
//   - [IVZGraphicsDisplay._graphicsDevice]
//   - [IVZGraphicsDisplay._graphicsOrientation]
//   - [IVZGraphicsDisplay._initDetached]
//   - [IVZGraphicsDisplay._matchesConfiguration]
//   - [IVZGraphicsDisplay._setGraphicsDevice]
//   - [IVZGraphicsDisplay._takeScreenshotWithCompletionHandler]
//   - [IVZGraphicsDisplay._uuid]
//   - [IVZGraphicsDisplay.InitWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid]
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay
type IVZGraphicsDisplay interface {
	objectivec.IObject

	// Topic: Methods

	_configuration() objectivec.IObject
	_graphicsDevice() objectivec.IObject
	_graphicsOrientation() int64
	_initDetached() objectivec.IObject
	_matchesConfiguration(configuration objectivec.IObject) bool
	_setGraphicsDevice(device objectivec.IObject)
	_takeScreenshotWithCompletionHandler(handler ErrorHandler)
	_uuid() objectivec.IObject
	InitWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZGraphicsDisplay
}

// Init initializes the instance.
func (g VZGraphicsDisplay) Init() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDisplay) Autorelease() VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplay creates a new VZGraphicsDisplay instance.
func NewVZGraphicsDisplay() VZGraphicsDisplay {
	class := getVZGraphicsDisplayClass()
	rv := objc.Send[VZGraphicsDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:
func NewGraphicsDisplayWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZGraphicsDisplay {
	instance := getVZGraphicsDisplayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:"), machine, index, index2, uuid)
	return VZGraphicsDisplayFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_configuration
func (g VZGraphicsDisplay) _configuration() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_configuration"))
	return objectivec.Object{ID: rv}
}

// Configuration is an exported wrapper for the private method _configuration.
func (g VZGraphicsDisplay) Configuration() objectivec.IObject {
	return g._configuration()
}
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_graphicsDevice
func (g VZGraphicsDisplay) _graphicsDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_graphicsDevice"))
	return objectivec.Object{ID: rv}
}

// GraphicsDevice is an exported wrapper for the private method _graphicsDevice.
func (g VZGraphicsDisplay) GraphicsDevice() objectivec.IObject {
	return g._graphicsDevice()
}
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_graphicsOrientation
func (g VZGraphicsDisplay) _graphicsOrientation() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("_graphicsOrientation"))
	return rv
}

// GraphicsOrientation is an exported wrapper for the private method _graphicsOrientation.
func (g VZGraphicsDisplay) GraphicsOrientation() int64 {
	return g._graphicsOrientation()
}
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_initDetached
func (g VZGraphicsDisplay) _initDetached() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_initDetached"))
	return objectivec.Object{ID: rv}
}

// InitDetached is an exported wrapper for the private method _initDetached.
func (g VZGraphicsDisplay) InitDetached() objectivec.IObject {
	return g._initDetached()
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_matchesConfiguration:
func (g VZGraphicsDisplay) _matchesConfiguration(configuration objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_matchesConfiguration:"), configuration)
	return rv
}

// MatchesConfiguration is an exported wrapper for the private method _matchesConfiguration.
func (g VZGraphicsDisplay) MatchesConfiguration(configuration objectivec.IObject) bool {
	return g._matchesConfiguration(configuration)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_setGraphicsDevice:
func (g VZGraphicsDisplay) _setGraphicsDevice(device objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_setGraphicsDevice:"), device)
}

// SetGraphicsDevice is an exported wrapper for the private method _setGraphicsDevice.
func (g VZGraphicsDisplay) SetGraphicsDevice(device objectivec.IObject) {
	g._setGraphicsDevice(device)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_takeScreenshotWithCompletionHandler:
func (g VZGraphicsDisplay) _takeScreenshotWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](g.ID, objc.Sel("_takeScreenshotWithCompletionHandler:"), _block0)
}

// TakeScreenshotWithCompletionHandler is an exported wrapper for the private method _takeScreenshotWithCompletionHandler.
func (g VZGraphicsDisplay) TakeScreenshotWithCompletionHandler(handler ErrorHandler) {
	g._takeScreenshotWithCompletionHandler(handler)
}
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/_uuid
func (g VZGraphicsDisplay) _uuid() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_uuid"))
	return objectivec.Object{ID: rv}
}

// Uuid is an exported wrapper for the private method _uuid.
func (g VZGraphicsDisplay) Uuid() objectivec.IObject {
	return g._uuid()
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplay/initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:
func (g VZGraphicsDisplay) InitWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZGraphicsDisplay {
	rv := objc.Send[VZGraphicsDisplay](g.ID, objc.Sel("initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:"), machine, index, index2, uuid)
	return rv
}

// _takeScreenshot is a synchronous wrapper around [VZGraphicsDisplay._takeScreenshotWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (g VZGraphicsDisplay) _takeScreenshot(ctx context.Context) error {
	done := make(chan error, 1)
	g._takeScreenshotWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

