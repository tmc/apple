package serial

import (
	"fmt"

	privvz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
)

// Kind names an experimental serial port implementation.
type Kind string

const (
	PL011 Kind = "pl011"
	UART  Kind = "16550"
)

// Available reports whether kind is present in the host framework.
func Available(kind Kind) bool {
	switch kind {
	case PL011:
		return privvz.GetVZPL011SerialPortConfigurationClass().Class() != 0
	case UART:
		return privvz.GetVZ16550SerialPortConfigurationClass().Class() != 0
	default:
		return false
	}
}

// New creates an experimental serial port with attachment.
func New(kind Kind, attachment vz.VZSerialPortAttachment) (vz.VZSerialPortConfiguration, error) {
	if attachment.ID == 0 {
		return vz.VZSerialPortConfiguration{}, fmt.Errorf("serial attachment required")
	}
	switch kind {
	case PL011:
		if !Available(kind) {
			return vz.VZSerialPortConfiguration{}, fmt.Errorf("PL011 serial port configuration is unavailable")
		}
		config := privvz.NewVZPL011SerialPortConfiguration()
		if config.ID == 0 {
			return vz.VZSerialPortConfiguration{}, fmt.Errorf("create PL011 serial port configuration")
		}
		config.Retain()
		serial := vz.VZSerialPortConfigurationFromID(config.ID)
		serial.SetAttachment(&attachment)
		return serial, nil
	case UART:
		if !Available(kind) {
			return vz.VZSerialPortConfiguration{}, fmt.Errorf("16550 serial port configuration is unavailable")
		}
		config := privvz.NewVZ16550SerialPortConfiguration()
		if config.ID == 0 {
			return vz.VZSerialPortConfiguration{}, fmt.Errorf("create 16550 serial port configuration")
		}
		config.Retain()
		serial := vz.VZSerialPortConfigurationFromID(config.ID)
		serial.SetAttachment(&attachment)
		return serial, nil
	default:
		return vz.VZSerialPortConfiguration{}, fmt.Errorf("unsupported serial kind: %s", kind)
	}
}
