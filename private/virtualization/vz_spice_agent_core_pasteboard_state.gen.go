// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSpiceAgentCorePasteboardState] class.
var (
	_VZSpiceAgentCorePasteboardStateClass     VZSpiceAgentCorePasteboardStateClass
	_VZSpiceAgentCorePasteboardStateClassOnce sync.Once
)

func getVZSpiceAgentCorePasteboardStateClass() VZSpiceAgentCorePasteboardStateClass {
	_VZSpiceAgentCorePasteboardStateClassOnce.Do(func() {
		_VZSpiceAgentCorePasteboardStateClass = VZSpiceAgentCorePasteboardStateClass{class: objc.GetClass("_VZSpiceAgentCorePasteboardState")}
	})
	return _VZSpiceAgentCorePasteboardStateClass
}

// GetVZSpiceAgentCorePasteboardStateClass returns the class object for _VZSpiceAgentCorePasteboardState.
func GetVZSpiceAgentCorePasteboardStateClass() VZSpiceAgentCorePasteboardStateClass {
	return getVZSpiceAgentCorePasteboardStateClass()
}

type VZSpiceAgentCorePasteboardStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSpiceAgentCorePasteboardStateClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSpiceAgentCorePasteboardStateClass) Alloc() VZSpiceAgentCorePasteboardState {
	rv := objc.Send[VZSpiceAgentCorePasteboardState](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCorePasteboardState
type VZSpiceAgentCorePasteboardState struct {
	objectivec.Object
}

// VZSpiceAgentCorePasteboardStateFromID constructs a [VZSpiceAgentCorePasteboardState] from an objc.ID.
func VZSpiceAgentCorePasteboardStateFromID(id objc.ID) VZSpiceAgentCorePasteboardState {
	return VZSpiceAgentCorePasteboardState{objectivec.Object{ID: id}}
}

// Ensure VZSpiceAgentCorePasteboardState implements IVZSpiceAgentCorePasteboardState.
var _ IVZSpiceAgentCorePasteboardState = VZSpiceAgentCorePasteboardState{}

// An interface definition for the [VZSpiceAgentCorePasteboardState] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCorePasteboardState
type IVZSpiceAgentCorePasteboardState interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZSpiceAgentCorePasteboardState) Init() VZSpiceAgentCorePasteboardState {
	rv := objc.Send[VZSpiceAgentCorePasteboardState](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSpiceAgentCorePasteboardState) Autorelease() VZSpiceAgentCorePasteboardState {
	rv := objc.Send[VZSpiceAgentCorePasteboardState](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgentCorePasteboardState creates a new VZSpiceAgentCorePasteboardState instance.
func NewVZSpiceAgentCorePasteboardState() VZSpiceAgentCorePasteboardState {
	class := getVZSpiceAgentCorePasteboardStateClass()
	rv := objc.Send[VZSpiceAgentCorePasteboardState](objc.ID(class.class), objc.Sel("new"))
	return rv
}
