// Code generated from Apple documentation. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
)

// GCControllerAxisInputFloat32Handler is the signature for a completion handler block.
type GCControllerAxisInputFloat32Handler = func(*GCControllerAxisInput, float32)

// NewGCControllerAxisInputFloat32Block wraps a Go [GCControllerAxisInputFloat32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerAxisInputFloat32Block(handler GCControllerAxisInputFloat32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float32) {
		var result *GCControllerAxisInput
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCControllerAxisInputFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerAxisValueChangedHandler handles The signature for the block that executes when the user changes the axis value.

// NewGCControllerAxisValueChangedHandlerBlock wraps a Go [GCControllerAxisValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerAxisValueChangedHandlerBlock(handler GCControllerAxisValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCControllerAxisInput, extra0 float32) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerButtonInputFloat32BoolBoolHandler is the signature for a completion handler block.
type GCControllerButtonInputFloat32BoolBoolHandler = func(*GCControllerButtonInput, float32, bool, bool)

// NewGCControllerButtonInputFloat32BoolBoolBlock wraps a Go [GCControllerButtonInputFloat32BoolBoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerButtonInputFloat32BoolBoolBlock(handler GCControllerButtonInputFloat32BoolBoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float32, extra1 bool, extra2 bool) {
		var result *GCControllerButtonInput
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCControllerButtonInputFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerButtonInputFloat32BoolHandler is the signature for a completion handler block.
type GCControllerButtonInputFloat32BoolHandler = func(*GCControllerButtonInput, float32, bool)

// NewGCControllerButtonInputFloat32BoolBlock wraps a Go [GCControllerButtonInputFloat32BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerButtonInputFloat32BoolBlock(handler GCControllerButtonInputFloat32BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float32, extra1 bool) {
		var result *GCControllerButtonInput
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCControllerButtonInputFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerButtonTouchedChangedHandler handles The signature for the block that executes when the user touches the button if the controller supports that feature.

// NewGCControllerButtonTouchedChangedHandlerBlock wraps a Go [GCControllerButtonTouchedChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerButtonTouchedChangedHandlerBlock(handler GCControllerButtonTouchedChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCControllerButtonInput, extra0 float32, extra1 bool, extra2 bool) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerButtonValueChangedHandler handles The signature for the block that executes when a button’s state changes.

// NewGCControllerButtonValueChangedHandlerBlock wraps a Go [GCControllerButtonValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerButtonValueChangedHandlerBlock(handler GCControllerButtonValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCControllerButtonInput, extra0 float32, extra1 bool) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerDirectionPadFloat32Float32Handler is the signature for a completion handler block.
type GCControllerDirectionPadFloat32Float32Handler = func(*GCControllerDirectionPad, float32, float32)

// NewGCControllerDirectionPadFloat32Float32Block wraps a Go [GCControllerDirectionPadFloat32Float32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerDirectionPadFloat32Float32Block(handler GCControllerDirectionPadFloat32Float32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float32, extra1 float32) {
		var result *GCControllerDirectionPad
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCControllerDirectionPadFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerDirectionPadValueChangedHandler handles The signature for the block that executes when either axis changes values.

// NewGCControllerDirectionPadValueChangedHandlerBlock wraps a Go [GCControllerDirectionPadValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerDirectionPadValueChangedHandlerBlock(handler GCControllerDirectionPadValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCControllerDirectionPad, extra0 float32, extra1 float32) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerTouchpadFloat32Float32Float32BoolHandler is the signature for a completion handler block.
type GCControllerTouchpadFloat32Float32Float32BoolHandler = func(*GCControllerTouchpad, float32, float32, float32, bool)

// NewGCControllerTouchpadFloat32Float32Float32BoolBlock wraps a Go [GCControllerTouchpadFloat32Float32Float32BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerTouchpadFloat32Float32Float32BoolBlock(handler GCControllerTouchpadFloat32Float32Float32BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float32, extra1 float32, extra2 float32, extra3 bool) {
		var result *GCControllerTouchpad
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCControllerTouchpadFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCControllerTouchpadHandler handles The signature for the block that executes when the user interacts with the touchpad.

// NewGCControllerTouchpadHandlerBlock wraps a Go [GCControllerTouchpadHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCControllerTouchpadHandlerBlock(handler GCControllerTouchpadHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCControllerTouchpad, extra0 float32, extra1 float32, extra2 float32, extra3 bool) {
		handler(primitive, extra0, extra1, extra2, extra3)
	})
	return objc.ID(block), func() { block.Release() }
}

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

// GCExtendedGamepadGCControllerElementHandler is the signature for a completion handler block.
type GCExtendedGamepadGCControllerElementHandler = func(*GCExtendedGamepad, *GCControllerElement)

// NewGCExtendedGamepadGCControllerElementBlock wraps a Go [GCExtendedGamepadGCControllerElementHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCExtendedGamepadGCControllerElementBlock(handler GCExtendedGamepadGCControllerElementHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *GCExtendedGamepad
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCExtendedGamepadFromID(resultID)
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

// GCExtendedGamepadValueChangedHandler handles The signature for the block that the profile calls when an element’s value changes.

// NewGCExtendedGamepadValueChangedHandlerBlock wraps a Go [GCExtendedGamepadValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCExtendedGamepadValueChangedHandlerBlock(handler GCExtendedGamepadValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCExtendedGamepad, extra0 GCControllerElement) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCGamepadValueChangedHandler handles Signature for the block executed if any element in the gamepad profile changes value.

// NewGCGamepadValueChangedHandlerBlock wraps a Go [GCGamepadValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCGamepadValueChangedHandlerBlock(handler GCGamepadValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *uintptr, extra0 GCControllerElement) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCKeyboardInputGCControllerButtonInputInt64BoolHandler is the signature for a completion handler block.
type GCKeyboardInputGCControllerButtonInputInt64BoolHandler = func(*GCKeyboardInput, *GCControllerButtonInput, int64, bool)

// NewGCKeyboardInputGCControllerButtonInputInt64BoolBlock wraps a Go [GCKeyboardInputGCControllerButtonInputInt64BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCKeyboardInputGCControllerButtonInputInt64BoolBlock(handler GCKeyboardInputGCControllerButtonInputInt64BoolHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1 int64, extra2 bool) {
		var result *GCKeyboardInput
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCKeyboardInputFromID(resultID)
			result = &v
		}
		var extra0 *GCControllerButtonInput
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := GCControllerButtonInputFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCKeyboardValueChangedHandler handles The signature for the block that the keyboard input profile calls when a key value changes.

// NewGCKeyboardValueChangedHandlerBlock wraps a Go [GCKeyboardValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCKeyboardValueChangedHandlerBlock(handler GCKeyboardValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCKeyboardInput, extra0 GCControllerButtonInput, extra1 int, extra2 bool) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCMicroGamepadGCControllerElementHandler is the signature for a completion handler block.
type GCMicroGamepadGCControllerElementHandler = func(*GCMicroGamepad, *GCControllerElement)

// NewGCMicroGamepadGCControllerElementBlock wraps a Go [GCMicroGamepadGCControllerElementHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCMicroGamepadGCControllerElementBlock(handler GCMicroGamepadGCControllerElementHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *GCMicroGamepad
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCMicroGamepadFromID(resultID)
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

// GCMicroGamepadValueChangedHandler handles Signature for the block that this profile calls when an element’s value changes.

// NewGCMicroGamepadValueChangedHandlerBlock wraps a Go [GCMicroGamepadValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCMicroGamepadValueChangedHandlerBlock(handler GCMicroGamepadValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCMicroGamepad, extra0 GCControllerElement) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCMotionHandler is the signature for a completion handler block.
type GCMotionHandler = func(*GCMotion)

// NewGCMotionBlock wraps a Go [GCMotionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCMotionBlock(handler GCMotionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *GCMotion
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCMotionFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCMotionValueChangedHandler handles The signature for the block that the profile calls when an element’s value changes.

// NewGCMotionValueChangedHandlerBlock wraps a Go [GCMotionValueChangedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCMotionValueChangedHandlerBlock(handler GCMotionValueChangedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal GCMotion) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCMouseInputFloat32Float32Handler is the signature for a completion handler block.
type GCMouseInputFloat32Float32Handler = func(*GCMouseInput, float32, float32)

// NewGCMouseInputFloat32Float32Block wraps a Go [GCMouseInputFloat32Float32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCMouseInputFloat32Float32Block(handler GCMouseInputFloat32Float32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float32, extra1 float32) {
		var result *GCMouseInput
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := GCMouseInputFromID(resultID)
			result = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// GCMouseMoved handles The signature for the block that the mouse input profile calls when the mouse moves.

// NewGCMouseMovedBlock wraps a Go [GCMouseMoved] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewGCMouseMovedBlock(handler GCMouseMoved) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive GCMouseInput, extra0 float32, extra1 float32) {
		handler(primitive, extra0, extra1)
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
