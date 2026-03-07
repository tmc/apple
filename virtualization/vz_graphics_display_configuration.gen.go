// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDisplayConfiguration] class.
var (
	_VZGraphicsDisplayConfigurationClass     VZGraphicsDisplayConfigurationClass
	_VZGraphicsDisplayConfigurationClassOnce sync.Once
)

func getVZGraphicsDisplayConfigurationClass() VZGraphicsDisplayConfigurationClass {
	_VZGraphicsDisplayConfigurationClassOnce.Do(func() {
		_VZGraphicsDisplayConfigurationClass = VZGraphicsDisplayConfigurationClass{class: objc.GetClass("VZGraphicsDisplayConfiguration")}
	})
	return _VZGraphicsDisplayConfigurationClass
}

// GetVZGraphicsDisplayConfigurationClass returns the class object for VZGraphicsDisplayConfiguration.
func GetVZGraphicsDisplayConfigurationClass() VZGraphicsDisplayConfigurationClass {
	return getVZGraphicsDisplayConfigurationClass()
}

type VZGraphicsDisplayConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDisplayConfigurationClass) Alloc() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The base class for a graphics display configuration.
//
// # Overview
// 
// Don’t instantiate [VZGraphicsDisplayConfiguration] directly. Use one of
// its subclasses instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration
type VZGraphicsDisplayConfiguration struct {
	objectivec.Object
}

// VZGraphicsDisplayConfigurationFromID constructs a [VZGraphicsDisplayConfiguration] from an objc.ID.
//
// The base class for a graphics display configuration.
func VZGraphicsDisplayConfigurationFromID(id objc.ID) VZGraphicsDisplayConfiguration {
	return VZGraphicsDisplayConfiguration{objectivec.Object{id}}
}
// NOTE: VZGraphicsDisplayConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZGraphicsDisplayConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDisplayConfiguration
type IVZGraphicsDisplayConfiguration interface {
	objectivec.IObject
}





// Init initializes the instance.
func (g VZGraphicsDisplayConfiguration) Init() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDisplayConfiguration) Autorelease() VZGraphicsDisplayConfiguration {
	rv := objc.Send[VZGraphicsDisplayConfiguration](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDisplayConfiguration creates a new VZGraphicsDisplayConfiguration instance.
func NewVZGraphicsDisplayConfiguration() VZGraphicsDisplayConfiguration {
	class := getVZGraphicsDisplayConfigurationClass()
	rv := objc.Send[VZGraphicsDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































