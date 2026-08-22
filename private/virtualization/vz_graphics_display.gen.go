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
	rv := objc.SendIfResponds[VZGraphicsDisplay](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

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
func (v VZGraphicsDisplay) Init() VZGraphicsDisplay {
	rv := objc.SendIfResponds[VZGraphicsDisplay](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGraphicsDisplay) Autorelease() VZGraphicsDisplay {
	rv := objc.SendIfResponds[VZGraphicsDisplay](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplay creates a new VZGraphicsDisplay instance.
func NewVZGraphicsDisplay() VZGraphicsDisplay {
	class := getVZGraphicsDisplayClass()
	rv := objc.SendIfResponds[VZGraphicsDisplay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGraphicsDisplayWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZGraphicsDisplay {
	instance := getVZGraphicsDisplayClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:"), machine, index, index2, uuid)
	return VZGraphicsDisplayFromID(rv)
}

func (v VZGraphicsDisplay) _configuration() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_configuration"))
	return objectivec.Object{ID: rv}
}

// Configuration is an exported wrapper for the private method _configuration.
func (v VZGraphicsDisplay) Configuration() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_configuration")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_configuration"}
		return nil, err
	}
	return v._configuration(), nil
}

// CanConfiguration reports whether the receiver responds to the private selector _configuration.
func (v VZGraphicsDisplay) CanConfiguration() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_configuration"))
}
func (v VZGraphicsDisplay) _graphicsDevice() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_graphicsDevice"))
	return objectivec.Object{ID: rv}
}

// GraphicsDevice is an exported wrapper for the private method _graphicsDevice.
func (v VZGraphicsDisplay) GraphicsDevice() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_graphicsDevice")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_graphicsDevice"}
		return nil, err
	}
	return v._graphicsDevice(), nil
}

// CanGraphicsDevice reports whether the receiver responds to the private selector _graphicsDevice.
func (v VZGraphicsDisplay) CanGraphicsDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_graphicsDevice"))
}
func (v VZGraphicsDisplay) _graphicsOrientation() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("_graphicsOrientation"))
	return rv
}

// GraphicsOrientation is an exported wrapper for the private method _graphicsOrientation.
func (v VZGraphicsDisplay) GraphicsOrientation() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_graphicsOrientation")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_graphicsOrientation"}
		return 0, err
	}
	return v._graphicsOrientation(), nil
}

// CanGraphicsOrientation reports whether the receiver responds to the private selector _graphicsOrientation.
func (v VZGraphicsDisplay) CanGraphicsOrientation() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_graphicsOrientation"))
}
func (v VZGraphicsDisplay) _initDetached() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_initDetached"))
	return objectivec.Object{ID: rv}
}

// InitDetached is an exported wrapper for the private method _initDetached.
func (v VZGraphicsDisplay) InitDetached() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initDetached")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initDetached"}
		return nil, err
	}
	return v._initDetached(), nil
}

// CanInitDetached reports whether the receiver responds to the private selector _initDetached.
func (v VZGraphicsDisplay) CanInitDetached() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initDetached"))
}
func (v VZGraphicsDisplay) _matchesConfiguration(configuration objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_matchesConfiguration:"), configuration)
	return rv
}

// MatchesConfiguration is an exported wrapper for the private method _matchesConfiguration.
func (v VZGraphicsDisplay) MatchesConfiguration(configuration objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_matchesConfiguration:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_matchesConfiguration:"}
		return false, err
	}
	return v._matchesConfiguration(configuration), nil
}

// CanMatchesConfiguration reports whether the receiver responds to the private selector _matchesConfiguration:.
func (v VZGraphicsDisplay) CanMatchesConfiguration() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_matchesConfiguration:"))
}
func (v VZGraphicsDisplay) _setGraphicsDevice(device objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setGraphicsDevice:"), device)
}

// SetGraphicsDevice is an exported wrapper for the private method _setGraphicsDevice.
func (v VZGraphicsDisplay) SetGraphicsDevice(device objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setGraphicsDevice:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setGraphicsDevice:"}
		return err
	}
	v._setGraphicsDevice(device)
	return nil
}

// CanSetGraphicsDevice reports whether the receiver responds to the private selector _setGraphicsDevice:.
func (v VZGraphicsDisplay) CanSetGraphicsDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setGraphicsDevice:"))
}
func (v VZGraphicsDisplay) _takeScreenshotWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_takeScreenshotWithCompletionHandler:"), _block0)
}

// TakeScreenshotWithCompletionHandler is an exported wrapper for the private method _takeScreenshotWithCompletionHandler.
func (v VZGraphicsDisplay) TakeScreenshotWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_takeScreenshotWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_takeScreenshotWithCompletionHandler:"}
		return err
	}
	v._takeScreenshotWithCompletionHandler(handler)
	return nil
}

// CanTakeScreenshotWithCompletionHandler reports whether the receiver responds to the private selector _takeScreenshotWithCompletionHandler:.
func (v VZGraphicsDisplay) CanTakeScreenshotWithCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_takeScreenshotWithCompletionHandler:"))
}
func (v VZGraphicsDisplay) _uuid() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_uuid"))
	return objectivec.Object{ID: rv}
}

// Uuid is an exported wrapper for the private method _uuid.
func (v VZGraphicsDisplay) Uuid() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_uuid")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_uuid"}
		return nil, err
	}
	return v._uuid(), nil
}

// CanUuid reports whether the receiver responds to the private selector _uuid.
func (v VZGraphicsDisplay) CanUuid() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_uuid"))
}
func (v VZGraphicsDisplay) InitWithVirtualMachineGraphicsDeviceIndexFramebufferIndexUuid(machine objectivec.IObject, index uint64, index2 uint64, uuid objectivec.IObject) VZGraphicsDisplay {
	rv := objc.SendIfResponds[VZGraphicsDisplay](v.ID, objc.Sel("initWithVirtualMachine:graphicsDeviceIndex:framebufferIndex:uuid:"), machine, index, index2, uuid)
	return rv
}

// _takeScreenshot is a synchronous wrapper around [VZGraphicsDisplay._takeScreenshotWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZGraphicsDisplay) _takeScreenshot(ctx context.Context) error {
	done := make(chan error, 1)
	v._takeScreenshotWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
