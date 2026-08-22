// Code generated from Apple documentation. DO NOT EDIT.

package iousbhost

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// IOUSBHostDefaultControlCompletionTimeout is the default completion timeout for input/output requests.
	//
	// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDefaultControlCompletionTimeout
	IOUSBHostDefaultControlCompletionTimeout foundation.NSTimeInterval
)

var ()

var (
	// IOUSBHostErrorDomain is the error domain for the framework.
	//
	// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostErrorDomain
	IOUSBHostErrorDomain foundation.NSErrorDomain
)

var (
	// IOUSBHostInterfacePropertyKeyAlternateSetting is the USB interface’s current alternative setting value.
	//
	// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterfacePropertyKey/alternateSetting
	IOUSBHostInterfacePropertyKeyAlternateSetting IOUSBHostInterfacePropertyKey
)

var ()

var (
	// IOUSBHostPropertyKeyLocationID is the location ID of the USB host device.
	//
	// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPropertyKeyLocationID
	IOUSBHostPropertyKeyLocationID IOUSBHostPropertyKey
)

var (
	// IOUSBHostVersionNumber is the version number of the framework.
	//
	// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostVersionNumber
	IOUSBHostVersionNumber float64
)

var (
	// IOUSBHostVersionString is a string representation of the framework’s version number.
	//
	// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostVersionString
	IOUSBHostVersionString uint8
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostDefaultControlCompletionTimeout"); err == nil && ptr != 0 {
		IOUSBHostDefaultControlCompletionTimeout = objc.ValueAt[foundation.NSTimeInterval](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostDevicePropertyKeyContainerID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostDevicePropertyKeys.ContainerID = IOUSBHostDevicePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostDevicePropertyKeyCurrentConfiguration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostDevicePropertyKeys.CurrentConfiguration = IOUSBHostDevicePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostDevicePropertyKeySerialNumberString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostDevicePropertyKeys.SerialNumberString = IOUSBHostDevicePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostDevicePropertyKeyVendorString"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostDevicePropertyKeys.VendorString = IOUSBHostDevicePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostInterfacePropertyKeyAlternateSetting"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostInterfacePropertyKeyAlternateSetting = IOUSBHostInterfacePropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyConfigurationValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.ConfigurationValue = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyDeviceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.DeviceClass = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyDeviceProtocol"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.DeviceProtocol = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyDeviceReleaseNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.DeviceReleaseNumber = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyDeviceSubClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.DeviceSubClass = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyInterfaceClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.InterfaceClass = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyInterfaceNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.InterfaceNumber = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyInterfaceProtocol"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.InterfaceProtocol = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyInterfaceSubClass"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.InterfaceSubClass = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyProductID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.ProductID = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyProductIDArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.ProductIDArray = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyProductIDMask"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.ProductIDMask = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeySpeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.Speed = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostMatchingPropertyKeyVendorID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostMatchingPropertyKeys.VendorID = IOUSBHostMatchingPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostPropertyKeyLocationID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOUSBHostPropertyKeyLocationID = IOUSBHostPropertyKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostVersionNumber"); err == nil && ptr != 0 {
		IOUSBHostVersionNumber = objc.ValueAt[float64](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOUSBHostVersionString"); err == nil && ptr != 0 {
		IOUSBHostVersionString = objc.ValueAt[uint8](ptr)
	}

}

// IOUSBHostDevicePropertyKeys provides typed accessors for [IOUSBHostDevicePropertyKey] constants.
var IOUSBHostDevicePropertyKeys struct {
	// ContainerID: The device’s container ID.
	ContainerID IOUSBHostDevicePropertyKey
	// CurrentConfiguration: The device’s current configuration value.
	CurrentConfiguration IOUSBHostDevicePropertyKey
	// SerialNumberString: The device’s serial number as a string.
	SerialNumberString IOUSBHostDevicePropertyKey
	// VendorString: The device’s vendor name.
	VendorString IOUSBHostDevicePropertyKey
}

// IOUSBHostMatchingPropertyKeys provides typed accessors for [IOUSBHostMatchingPropertyKey] constants.
var IOUSBHostMatchingPropertyKeys struct {
	// ConfigurationValue: The matching property for the device’s current configuration value.
	ConfigurationValue IOUSBHostMatchingPropertyKey
	// DeviceClass: The matching property for the device’s class.
	DeviceClass IOUSBHostMatchingPropertyKey
	// DeviceProtocol: The matching property for the device’s protocol.
	DeviceProtocol IOUSBHostMatchingPropertyKey
	// DeviceReleaseNumber: The matching property for the device’s release number.
	DeviceReleaseNumber IOUSBHostMatchingPropertyKey
	// DeviceSubClass: The matching property for the device’s subclass.
	DeviceSubClass IOUSBHostMatchingPropertyKey
	// InterfaceClass: The matching property for the interface’s class ID.
	InterfaceClass IOUSBHostMatchingPropertyKey
	// InterfaceNumber: The matching property for the device’s interface number.
	InterfaceNumber IOUSBHostMatchingPropertyKey
	// InterfaceProtocol: The matching property for the interface’s protocol.
	InterfaceProtocol IOUSBHostMatchingPropertyKey
	// InterfaceSubClass: The matching property for the interface’s subclass ID.
	InterfaceSubClass IOUSBHostMatchingPropertyKey
	// ProductID: The matching property for the device’s product ID.
	ProductID IOUSBHostMatchingPropertyKey
	// ProductIDArray: The matching property on a list of product IDs.
	ProductIDArray IOUSBHostMatchingPropertyKey
	// ProductIDMask: The matching property on a mask of product IDs.
	ProductIDMask IOUSBHostMatchingPropertyKey
	// Speed: The matching property for the device’s enumeration speed.
	Speed IOUSBHostMatchingPropertyKey
	// VendorID: The matching property for the device’s vendor ID.
	VendorID IOUSBHostMatchingPropertyKey
}
