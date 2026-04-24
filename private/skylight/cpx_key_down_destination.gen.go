// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXKeyDownDestination] class.
var (
	_CPXKeyDownDestinationClass     CPXKeyDownDestinationClass
	_CPXKeyDownDestinationClassOnce sync.Once
)

func getCPXKeyDownDestinationClass() CPXKeyDownDestinationClass {
	_CPXKeyDownDestinationClassOnce.Do(func() {
		_CPXKeyDownDestinationClass = CPXKeyDownDestinationClass{class: objc.GetClass("_CPXKeyDownDestination")}
	})
	return _CPXKeyDownDestinationClass
}

// GetCPXKeyDownDestinationClass returns the class object for _CPXKeyDownDestination.
func GetCPXKeyDownDestinationClass() CPXKeyDownDestinationClass {
	return getCPXKeyDownDestinationClass()
}

type CPXKeyDownDestinationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXKeyDownDestinationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXKeyDownDestinationClass) Alloc() CPXKeyDownDestination {
	rv := objc.Send[CPXKeyDownDestination](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_CPXKeyDownDestination
type CPXKeyDownDestination struct {
	objectivec.Object
}

// CPXKeyDownDestinationFromID constructs a [CPXKeyDownDestination] from an objc.ID.
func CPXKeyDownDestinationFromID(id objc.ID) CPXKeyDownDestination {
	return CPXKeyDownDestination{objectivec.Object{ID: id}}
}

// Ensure CPXKeyDownDestination implements ICPXKeyDownDestination.
var _ ICPXKeyDownDestination = CPXKeyDownDestination{}

// An interface definition for the [CPXKeyDownDestination] class.
//
// See: https://developer.apple.com/documentation/SkyLight/_CPXKeyDownDestination
type ICPXKeyDownDestination interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CPXKeyDownDestination) Init() CPXKeyDownDestination {
	rv := objc.Send[CPXKeyDownDestination](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXKeyDownDestination) Autorelease() CPXKeyDownDestination {
	rv := objc.Send[CPXKeyDownDestination](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXKeyDownDestination creates a new CPXKeyDownDestination instance.
func NewCPXKeyDownDestination() CPXKeyDownDestination {
	class := getCPXKeyDownDestinationClass()
	rv := objc.Send[CPXKeyDownDestination](objc.ID(class.class), objc.Sel("new"))
	return rv
}
