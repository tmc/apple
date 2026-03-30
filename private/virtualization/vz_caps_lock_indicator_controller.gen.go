// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCapsLockIndicatorController] class.
var (
	_VZCapsLockIndicatorControllerClass     VZCapsLockIndicatorControllerClass
	_VZCapsLockIndicatorControllerClassOnce sync.Once
)

func getVZCapsLockIndicatorControllerClass() VZCapsLockIndicatorControllerClass {
	_VZCapsLockIndicatorControllerClassOnce.Do(func() {
		_VZCapsLockIndicatorControllerClass = VZCapsLockIndicatorControllerClass{class: objc.GetClass("VZCapsLockIndicatorController")}
	})
	return _VZCapsLockIndicatorControllerClass
}

// GetVZCapsLockIndicatorControllerClass returns the class object for VZCapsLockIndicatorController.
func GetVZCapsLockIndicatorControllerClass() VZCapsLockIndicatorControllerClass {
	return getVZCapsLockIndicatorControllerClass()
}

type VZCapsLockIndicatorControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCapsLockIndicatorControllerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCapsLockIndicatorControllerClass) Alloc() VZCapsLockIndicatorController {
	rv := objc.Send[VZCapsLockIndicatorController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZCapsLockIndicatorController
type VZCapsLockIndicatorController struct {
	objectivec.Object
}

// VZCapsLockIndicatorControllerFromID constructs a [VZCapsLockIndicatorController] from an objc.ID.
func VZCapsLockIndicatorControllerFromID(id objc.ID) VZCapsLockIndicatorController {
	return VZCapsLockIndicatorController{objectivec.Object{ID: id}}
}

// Ensure VZCapsLockIndicatorController implements IVZCapsLockIndicatorController.
var _ IVZCapsLockIndicatorController = VZCapsLockIndicatorController{}

// An interface definition for the [VZCapsLockIndicatorController] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZCapsLockIndicatorController
type IVZCapsLockIndicatorController interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c VZCapsLockIndicatorController) Init() VZCapsLockIndicatorController {
	rv := objc.Send[VZCapsLockIndicatorController](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VZCapsLockIndicatorController) Autorelease() VZCapsLockIndicatorController {
	rv := objc.Send[VZCapsLockIndicatorController](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCapsLockIndicatorController creates a new VZCapsLockIndicatorController instance.
func NewVZCapsLockIndicatorController() VZCapsLockIndicatorController {
	class := getVZCapsLockIndicatorControllerClass()
	rv := objc.Send[VZCapsLockIndicatorController](objc.ID(class.class), objc.Sel("new"))
	return rv
}
