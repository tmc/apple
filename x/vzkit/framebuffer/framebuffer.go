package framebuffer

import (
	"context"
	"fmt"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// Display wraps a private graphics display.
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
	return d.raw.GraphicsOrientation()
}

// Uuid returns the display UUID object.
func (d Display) Uuid() objectivec.IObject {
	return d.raw.Uuid()
}

// Configuration returns the private display configuration object.
func (d Display) Configuration() objectivec.IObject {
	return d.raw.Configuration()
}

// GraphicsDevice returns the underlying graphics device object.
func (d Display) GraphicsDevice() objectivec.IObject {
	return d.raw.GraphicsDevice()
}

// TakeScreenshot triggers a screenshot on the display and waits for completion.
func (d Display) TakeScreenshot(ctx context.Context) error {
	done := make(chan error, 1)
	d.raw.TakeScreenshotWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Framebuffer wraps a private framebuffer object.
type Framebuffer struct {
	raw pvz.VZFramebuffer
}

// NewFramebuffer creates a new framebuffer wrapper.
func NewFramebuffer() Framebuffer {
	return Framebuffer{raw: pvz.NewVZFramebuffer()}
}

// FromFramebuffer wraps an existing framebuffer.
func FromFramebuffer(raw pvz.VZFramebuffer) Framebuffer {
	return Framebuffer{raw: raw}
}

// Raw returns the underlying framebuffer.
func (f Framebuffer) Raw() pvz.VZFramebuffer {
	return f.raw
}

// TakeScreenshot triggers the private framebuffer screenshot path.
// The generated binding currently exposes only completion and conversion blocks
// without typed image payloads, so this helper just waits for completion.
func (f Framebuffer) TakeScreenshot(ctx context.Context) error {
	done := make(chan error, 1)
	f.raw.TakeScreenshotWithCompletionHandlerImageConversionBlock(func() {
		done <- nil
	}, func() {})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// View wraps a framebuffer view.
type View struct {
	raw pvz.VZFramebufferView
}

// NewView creates a new framebuffer view.
func NewView() View {
	return View{raw: pvz.NewVZFramebufferView()}
}

// NewViewWithFrame creates a framebuffer view for the given frame.
func NewViewWithFrame(frame corefoundation.CGRect) View {
	return View{raw: pvz.NewVZFramebufferViewWithFrame(frame)}
}

// Raw returns the underlying framebuffer view.
func (v View) Raw() pvz.VZFramebufferView {
	return v.raw
}

// Framebuffer returns the attached framebuffer.
func (v View) Framebuffer() *pvz.VZFramebuffer {
	return v.raw.Framebuffer()
}

// SetFramebuffer attaches a framebuffer to the view.
func (v View) SetFramebuffer(framebuffer *pvz.VZFramebuffer) {
	v.raw.SetFramebuffer(framebuffer)
}

// FramebufferSize returns the framebuffer size.
func (v View) FramebufferSize() corefoundation.CGSize {
	return v.raw.FramebufferSize()
}

// ShowsCursor reports whether the cursor is visible.
func (v View) ShowsCursor() bool {
	return v.raw.ShowsCursor()
}

// SetShowsCursor toggles cursor visibility.
func (v View) SetShowsCursor(shows bool) {
	v.raw.SetShowsCursor(shows)
}

// SuppressFrameUpdates reports whether frame updates are suppressed.
func (v View) SuppressFrameUpdates() bool {
	return v.raw.SuppressFrameUpdates()
}

// SetSuppressFrameUpdates toggles frame update suppression.
func (v View) SetSuppressFrameUpdates(suppress bool) {
	v.raw.SetSuppressFrameUpdates(suppress)
}

// DisplayProtectionOptions returns the current display protection options.
func (v View) DisplayProtectionOptions() objectivec.IObject {
	return v.raw.GetDisplayProtectionOptions()
}

// AttachFramebuffer allocates a framebuffer and attaches it to the view.
func (v View) AttachFramebuffer(framebuffer *pvz.VZFramebuffer) error {
	if framebuffer == nil || framebuffer.ID == 0 {
		return fmt.Errorf("framebuffer required")
	}
	v.SetFramebuffer(framebuffer)
	return nil
}
