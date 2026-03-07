// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMemoryBalloonDeviceConfiguration] class.
var (
	_VZMemoryBalloonDeviceConfigurationClass     VZMemoryBalloonDeviceConfigurationClass
	_VZMemoryBalloonDeviceConfigurationClassOnce sync.Once
)

func getVZMemoryBalloonDeviceConfigurationClass() VZMemoryBalloonDeviceConfigurationClass {
	_VZMemoryBalloonDeviceConfigurationClassOnce.Do(func() {
		_VZMemoryBalloonDeviceConfigurationClass = VZMemoryBalloonDeviceConfigurationClass{class: objc.GetClass("VZMemoryBalloonDeviceConfiguration")}
	})
	return _VZMemoryBalloonDeviceConfigurationClass
}

// GetVZMemoryBalloonDeviceConfigurationClass returns the class object for VZMemoryBalloonDeviceConfiguration.
func GetVZMemoryBalloonDeviceConfigurationClass() VZMemoryBalloonDeviceConfigurationClass {
	return getVZMemoryBalloonDeviceConfigurationClass()
}

type VZMemoryBalloonDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMemoryBalloonDeviceConfigurationClass) Alloc() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The common configuration traits for memory balloon devices.
//
// # Overview
// 
// Don’t instantiate this abstract class directly. Instead, instantiate one
// of its subclasses such as
// [VZVirtioTraditionalMemoryBalloonDeviceConfiguration].
//
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration
type VZMemoryBalloonDeviceConfiguration struct {
	objectivec.Object
}

// VZMemoryBalloonDeviceConfigurationFromID constructs a [VZMemoryBalloonDeviceConfiguration] from an objc.ID.
//
// The common configuration traits for memory balloon devices.
func VZMemoryBalloonDeviceConfigurationFromID(id objc.ID) VZMemoryBalloonDeviceConfiguration {
	return VZMemoryBalloonDeviceConfiguration{objectivec.Object{id}}
}
// NOTE: VZMemoryBalloonDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMemoryBalloonDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDeviceConfiguration
type IVZMemoryBalloonDeviceConfiguration interface {
	objectivec.IObject
}





// Init initializes the instance.
func (m VZMemoryBalloonDeviceConfiguration) Init() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMemoryBalloonDeviceConfiguration) Autorelease() VZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDeviceConfiguration creates a new VZMemoryBalloonDeviceConfiguration instance.
func NewVZMemoryBalloonDeviceConfiguration() VZMemoryBalloonDeviceConfiguration {
	class := getVZMemoryBalloonDeviceConfigurationClass()
	rv := objc.Send[VZMemoryBalloonDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































