package framebuffer

import (
	"fmt"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
)

// Display wraps a private graphics display. Screenshot capture lives in
// screenshot.go (CaptureImage/CapturePNG), which drives the corrected
// two-argument completion ABI, void(^)(CGImageRef, NSError*).
type Display struct {
	raw pvz.VZGraphicsDisplay
}

// FromGraphicsDisplay wraps a private graphics display.
func FromGraphicsDisplay(display pvz.VZGraphicsDisplay) Display {
	return Display{raw: display}
}

// FromGraphicsDisplayID wraps a private graphics display from an objc.ID.
func FromGraphicsDisplayID(id objectivec.IObject) Display {
	return Display{raw: pvz.VZGraphicsDisplayFromID(id.GetID())}
}

// Raw returns the underlying display.
func (d Display) Raw() pvz.VZGraphicsDisplay {
	return d.raw
}

// GraphicsOrientation returns the display orientation.
func (d Display) GraphicsOrientation() int64 {
	v, _ := d.raw.GraphicsOrientation()
	return v
}

// Uuid returns the display UUID object.
func (d Display) Uuid() objectivec.IObject {
	v, _ := d.raw.Uuid()
	return v
}

// Configuration returns the private display configuration object.
func (d Display) Configuration() objectivec.IObject {
	v, _ := d.raw.Configuration()
	return v
}

// GraphicsDevice returns the underlying graphics device object.
func (d Display) GraphicsDevice() objectivec.IObject {
	v, _ := d.raw.GraphicsDevice()
	return v
}

// LinearFramebufferConfig describes a private linear framebuffer device.
type LinearFramebufferConfig struct {
	Width  int
	Height int
}

// NewLinearFramebufferGraphicsDeviceConfiguration creates a private linear
// framebuffer graphics device configuration.
func NewLinearFramebufferGraphicsDeviceConfiguration(c LinearFramebufferConfig) (vz.VZGraphicsDeviceConfiguration, error) {
	if c.Width <= 0 || c.Height <= 0 {
		return vz.VZGraphicsDeviceConfiguration{}, fmt.Errorf("width and height must be positive")
	}
	if pvz.GetVZLinearFramebufferGraphicsDeviceConfigurationClass().Class() == 0 {
		return vz.VZGraphicsDeviceConfiguration{}, fmt.Errorf("linear framebuffer graphics configuration is unavailable")
	}
	graphics := pvz.NewVZLinearFramebufferGraphicsDeviceConfigurationWithBackingStoreSize(corefoundation.CGSize{
		Width:  float64(c.Width),
		Height: float64(c.Height),
	})
	if graphics.ID == 0 {
		return vz.VZGraphicsDeviceConfiguration{}, fmt.Errorf("create linear framebuffer graphics configuration")
	}
	graphics.Retain()
	return vz.VZGraphicsDeviceConfigurationFromID(graphics.ID), nil
}

// SetLinearFramebufferGraphicsDevice sets one private linear framebuffer
// graphics device on config.
func SetLinearFramebufferGraphicsDevice(config vz.VZVirtualMachineConfiguration, c LinearFramebufferConfig) error {
	device, err := NewLinearFramebufferGraphicsDeviceConfiguration(c)
	if err != nil {
		return err
	}
	objc.Send[struct{}](config.ID, objc.Sel("setGraphicsDevices:"), objectivec.IObjectSliceToNSArray([]vz.VZGraphicsDeviceConfiguration{device}))
	return nil
}
