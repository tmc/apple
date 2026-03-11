//go:build darwin

package ane

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
)

// Probe queries the system for ANE hardware information.
func Probe() (DeviceInfo, error) {
	cls := appleneuralengine.GetANEDeviceInfoClass()

	info := DeviceInfo{
		HasANE:   cls.HasANE(),
		NumCores: cls.NumANECores(),
		IsVM:     cls.IsVirtualMachine(),
	}

	// Architecture, product, and build version come back as ObjC objects;
	// convert via NSString description.
	if obj := cls.AneArchitectureType(); obj != nil {
		info.Architecture = descriptionString(obj.GetID())
	}
	if obj := cls.ProductName(); obj != nil {
		info.Product = descriptionString(obj.GetID())
	}
	if obj := cls.BuildVersion(); obj != nil {
		info.BuildVersion = descriptionString(obj.GetID())
	}

	if !info.HasANE {
		return info, fmt.Errorf("%w", ErrNoANE)
	}
	return info, nil
}

// descriptionString calls -description on an ObjC object and returns the Go string.
func descriptionString(id objc.ID) string {
	if id == 0 {
		return ""
	}
	rv := objc.Send[objc.ID](id, objc.Sel("description"))
	if rv == 0 {
		return ""
	}
	return objc.Send[string](rv, objc.Sel("UTF8String"))
}
