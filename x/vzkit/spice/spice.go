package spice

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	"github.com/tmc/apple/x/vzkit/vm"
)

// PortAttachment wraps the private SPICE agent port attachment.
type PortAttachment struct {
	raw pvz.VZSpiceAgentPortAttachment
}

// NewPortAttachment creates a SPICE agent port attachment.
func NewPortAttachment(sharesClipboard bool) PortAttachment {
	attachment := pvz.NewVZSpiceAgentPortAttachment()
	setSharesClipboard(attachment, sharesClipboard)
	return PortAttachment{raw: attachment}
}

// Raw returns the underlying attachment.
func (p PortAttachment) Raw() pvz.VZSpiceAgentPortAttachment {
	return p.raw
}

// SharesClipboard reports whether host clipboard sharing is enabled.
func (p PortAttachment) SharesClipboard() bool {
	return objc.Send[bool](p.raw.ID, objc.Sel("sharesClipboard"))
}

// SetSharesClipboard toggles host clipboard sharing.
func (p PortAttachment) SetSharesClipboard(sharesClipboard bool) {
	setSharesClipboard(p.raw, sharesClipboard)
}

func setSharesClipboard(attachment pvz.VZSpiceAgentPortAttachment, sharesClipboard bool) {
	objc.Send[struct{}](attachment.ID, objc.Sel("setSharesClipboard:"), sharesClipboard)
}

// PortName returns the console port name for the SPICE agent attachment.
func PortName() string {
	class := pvz.GetVZSpiceAgentPortAttachmentClass()
	rv := objc.Send[objc.ID](objc.ID(class.Class()), objc.Sel("spiceAgentPortName"))
	return foundation.NSStringFromID(rv).String()
}

// ConsoleConfig creates a virtio console device configured for SPICE agent use.
func ConsoleConfig(sharesClipboard bool) (pvz.VZVirtioConsoleDeviceConfiguration, error) {
	attachment := pvz.NewVZSpiceAgentPortAttachment()
	if attachment.ID == 0 {
		return pvz.VZVirtioConsoleDeviceConfiguration{}, fmt.Errorf("create SPICE agent port attachment")
	}
	setSharesClipboard(attachment, sharesClipboard)

	portName := PortName()
	port := pvz.NewVZVirtioConsolePortConfiguration()
	if port.ID == 0 {
		return pvz.VZVirtioConsoleDeviceConfiguration{}, fmt.Errorf("create console port configuration")
	}
	objc.Send[struct{}](port.ID, objc.Sel("setName:"), objc.String(portName))
	objc.Send[struct{}](port.ID, objc.Sel("setAttachment:"), attachment.VZSerialPortAttachment)

	console := pvz.NewVZVirtioConsoleDeviceConfiguration()
	if console.ID == 0 {
		return pvz.VZVirtioConsoleDeviceConfiguration{}, fmt.Errorf("create console device configuration")
	}
	ports := objc.Send[objc.ID](console.ID, objc.Sel("ports"))
	objc.Send[struct{}](ports, objc.Sel("setObject:atIndexedSubscript:"), port.ID, uint(0))
	return console, nil
}

// Capabilities wraps SPICE agent core capabilities.
type Capabilities struct {
	raw pvz.VZSpiceAgentCoreCapabilities
}

// NewCapabilities creates a default SPICE capability set.
func NewCapabilities(clipboard bool) Capabilities {
	caps := pvz.NewVZSpiceAgentCoreCapabilities()
	caps.SetClipboard(clipboard)
	return Capabilities{raw: caps}
}

// Raw returns the underlying capabilities object.
func (c Capabilities) Raw() pvz.VZSpiceAgentCoreCapabilities {
	return c.raw
}

// Session wraps a SPICE session object.
type Session struct {
	raw pvz.VZSpiceAgentCoreSession
}

// NewSession creates a session wrapper.
func NewSession() Session {
	return Session{raw: pvz.NewVZSpiceAgentCoreSession()}
}

// Core wraps the private SPICE core object.
type Core struct {
	raw pvz.VZSpiceAgentCore
}

// NewCore creates a SPICE core object.
func NewCore(pasteboard objectivec.IObject, queue *vm.Queue, caps Capabilities, input, output unsafe.Pointer) Core {
	var q pvz.DispatchQueue
	if queue != nil {
		id := uintptr(queue.Queue().ID())
		q = *(*pvz.DispatchQueue)(unsafe.Pointer(&id))
	}
	core := pvz.NewVZSpiceAgentCoreWithPasteboardQueueCapabilitiesInputOutput(
		pasteboard, q, caps.Raw(), input, output,
	)
	return Core{raw: core}
}

// IsValid reports whether the core is currently valid.
func (c Core) IsValid() bool {
	return c.raw.IsValid()
}

// Pause pauses the SPICE core.
func (c Core) Pause() {
	c.raw.Pause()
}

// Resume resumes the SPICE core.
func (c Core) Resume() {
	c.raw.Resume()
}

// Stop stops the SPICE core.
func (c Core) Stop() {
	c.raw.Stop()
}

// PortAttachmentForClipboard returns a clipboard-enabled port attachment.
func PortAttachmentForClipboard() PortAttachment {
	return NewPortAttachment(true)
}

// MakePasteboardItemState returns a new pasteboard item state object.
func MakePasteboardItemState() pvz.VZSpiceAgentCorePasteboardItemState {
	return pvz.NewVZSpiceAgentCorePasteboardItemState()
}

// MakePasteboardState returns a new pasteboard state object.
func MakePasteboardState() pvz.VZSpiceAgentCorePasteboardState {
	return pvz.NewVZSpiceAgentCorePasteboardState()
}

// MakeClipboardCapabilities returns capabilities with clipboard sharing enabled.
func MakeClipboardCapabilities() Capabilities {
	return NewCapabilities(true)
}

// ClipboardConfig is a compatibility helper for the current clipboard package.
func ClipboardConfig() (pvz.VZVirtioConsoleDeviceConfiguration, error) {
	return ConsoleConfig(true)
}

// BuildPasswordObject converts a password to an Objective-C string object.
func BuildPasswordObject(password string) objectivec.IObject {
	return foundation.NewNSString().InitWithString(password)
}
