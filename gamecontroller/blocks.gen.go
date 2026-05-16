// Code generated from Apple documentation. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
)

// GCDevicePhysicalInputHandler is the signature for a completion handler block.
type GCDevicePhysicalInputHandler = func(GCDevicePhysicalInput)

// NewGCDevicePhysicalInputBlock wraps a Go [GCDevicePhysicalInputHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCDevicePhysicalInputBlock(handler GCDevicePhysicalInputHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result GCDevicePhysicalInput
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = GCDevicePhysicalInputObjectFromID(resultID)
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCPhysicalInputProfileGCControllerElementHandler is the signature for a completion handler block.
type GCPhysicalInputProfileGCControllerElementHandler = func(*GCPhysicalInputProfile, *GCControllerElement)

// NewGCPhysicalInputProfileGCControllerElementBlock wraps a Go [GCPhysicalInputProfileGCControllerElementHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCPhysicalInputProfileGCControllerElementBlock(handler GCPhysicalInputProfileGCControllerElementHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *GCPhysicalInputProfile
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCPhysicalInputProfileFromID(resultID)
			result = &v
		}
		var extra0 *GCControllerElement
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := GCControllerElementFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles The block that the framework calls when it completes the request.
//
// Used by:
//   - [GCController.StartWirelessControllerDiscoveryWithCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GCController.StartWirelessControllerDiscoveryWithCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
