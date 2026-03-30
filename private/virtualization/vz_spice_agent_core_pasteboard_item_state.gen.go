// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSpiceAgentCorePasteboardItemState] class.
var (
	_VZSpiceAgentCorePasteboardItemStateClass     VZSpiceAgentCorePasteboardItemStateClass
	_VZSpiceAgentCorePasteboardItemStateClassOnce sync.Once
)

func getVZSpiceAgentCorePasteboardItemStateClass() VZSpiceAgentCorePasteboardItemStateClass {
	_VZSpiceAgentCorePasteboardItemStateClassOnce.Do(func() {
		_VZSpiceAgentCorePasteboardItemStateClass = VZSpiceAgentCorePasteboardItemStateClass{class: objc.GetClass("_VZSpiceAgentCorePasteboardItemState")}
	})
	return _VZSpiceAgentCorePasteboardItemStateClass
}

// GetVZSpiceAgentCorePasteboardItemStateClass returns the class object for _VZSpiceAgentCorePasteboardItemState.
func GetVZSpiceAgentCorePasteboardItemStateClass() VZSpiceAgentCorePasteboardItemStateClass {
	return getVZSpiceAgentCorePasteboardItemStateClass()
}

type VZSpiceAgentCorePasteboardItemStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSpiceAgentCorePasteboardItemStateClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSpiceAgentCorePasteboardItemStateClass) Alloc() VZSpiceAgentCorePasteboardItemState {
	rv := objc.Send[VZSpiceAgentCorePasteboardItemState](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCorePasteboardItemState
type VZSpiceAgentCorePasteboardItemState struct {
	objectivec.Object
}

// VZSpiceAgentCorePasteboardItemStateFromID constructs a [VZSpiceAgentCorePasteboardItemState] from an objc.ID.
func VZSpiceAgentCorePasteboardItemStateFromID(id objc.ID) VZSpiceAgentCorePasteboardItemState {
	return VZSpiceAgentCorePasteboardItemState{objectivec.Object{ID: id}}
}

// Ensure VZSpiceAgentCorePasteboardItemState implements IVZSpiceAgentCorePasteboardItemState.
var _ IVZSpiceAgentCorePasteboardItemState = VZSpiceAgentCorePasteboardItemState{}

// An interface definition for the [VZSpiceAgentCorePasteboardItemState] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCorePasteboardItemState
type IVZSpiceAgentCorePasteboardItemState interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZSpiceAgentCorePasteboardItemState) Init() VZSpiceAgentCorePasteboardItemState {
	rv := objc.Send[VZSpiceAgentCorePasteboardItemState](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSpiceAgentCorePasteboardItemState) Autorelease() VZSpiceAgentCorePasteboardItemState {
	rv := objc.Send[VZSpiceAgentCorePasteboardItemState](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgentCorePasteboardItemState creates a new VZSpiceAgentCorePasteboardItemState instance.
func NewVZSpiceAgentCorePasteboardItemState() VZSpiceAgentCorePasteboardItemState {
	class := getVZSpiceAgentCorePasteboardItemStateClass()
	rv := objc.Send[VZSpiceAgentCorePasteboardItemState](objc.ID(class.class), objc.Sel("new"))
	return rv
}
