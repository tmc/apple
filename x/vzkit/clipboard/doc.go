// Package clipboard configures host-guest clipboard sharing via SPICE agent.
//
// It creates a VZVirtioConsoleDeviceConfiguration with a SPICE agent port
// attachment that enables bidirectional clipboard sync (text and images).
// Requires macOS 13+ on the host and spice-vdagent running in the guest.
//
// Basic usage:
//
//	consoleDevice, err := clipboard.NewConfig()
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Add consoleDevice to your VM configuration's console devices.
package clipboard
