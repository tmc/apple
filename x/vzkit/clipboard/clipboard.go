package clipboard

import (
	"fmt"

	"github.com/tmc/apple/objc"
	vz "github.com/tmc/apple/virtualization"
)

// NewConfig creates a Virtio console device configured for SPICE agent
// clipboard sharing. The returned device should be added to the VM's
// console device list.
//
// Clipboard sharing only functions when the guest has spice-vdagent running.
// For macOS guests this requires macOS 15+ with SPICE guest tools installed.
func NewConfig() (vz.VZVirtioConsoleDeviceConfiguration, error) {
	spiceAgent := vz.NewVZSpiceAgentPortAttachment()
	if spiceAgent.ID == 0 {
		return vz.VZVirtioConsoleDeviceConfiguration{}, fmt.Errorf("create SPICE agent port attachment")
	}
	spiceAgent.SetSharesClipboard(true)

	portName := vz.GetVZSpiceAgentPortAttachmentClass().SpiceAgentPortName()

	spicePort := vz.NewVZVirtioConsolePortConfiguration()
	if spicePort.ID == 0 {
		return vz.VZVirtioConsoleDeviceConfiguration{}, fmt.Errorf("create console port configuration")
	}
	spicePort.SetName(portName)
	spicePort.SetAttachment(&spiceAgent.VZSerialPortAttachment)

	consoleDevice := vz.NewVZVirtioConsoleDeviceConfiguration()
	if consoleDevice.ID == 0 {
		return vz.VZVirtioConsoleDeviceConfiguration{}, fmt.Errorf("create console device configuration")
	}

	// Assign the SPICE port at index 0.
	// Generated bindings lack the indexed setter, so use raw objc.Send.
	ports := consoleDevice.Ports()
	objc.Send[struct{}](ports.GetID(), objc.Sel("setObject:atIndexedSubscript:"), spicePort.ID, uint(0))

	return consoleDevice, nil
}
