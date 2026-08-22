// Code generated from Apple documentation for kernel. DO NOT EDIT.

// Package kernel provides Go bindings for the kernel framework.
//
// Develop kernel-resident device drivers and kernel extensions.
//
// The Kernel Framework provides the APIs and support for kernel-resident
// device drivers and other kernel extensions. It defines the base class for
// I/O Kit device drivers ([IOService]), several helper classes, and the
// families that support many types of devices.
//
// # Kernel Extensions
//
//   - [Implementing drivers, system extensions, and kexts]: Create drivers and system extensions to communicate with hardware and provide low-level services, and only use kernel extensions for a few tasks.
//   - [Installing a custom kernel extension]: Install kernel extensions using a custom installer package, and help users understand the installation process.
//   - [Debugging a custom kernel extension]: Configure your system to enable the debugging of custom kernel extensions from a second Mac.
//   - [Generating a Non-Maskable Interrupt]: Interrupt the kernel on a target Mac and attach a remote debugger to it.
//
// # IOKit Drivers
//
//   - [IOKit Fundamentals]: Implement a driver for your custom hardware using a third-party kernel extension. ([IOService], [IORegistryEntry], [IORegistryIterator], [IOKitDiagnosticsParameters], [DriverDescription])
//   - [Hardware Families]: Add support for specific hardware protocols such as USB, and for standard network, serial, audio, and graphics interfaces.
//   - [Driver Support]: Explore the device registry and access power-management utilities and other shared driver features. ([IOACPIAddressSpaceDescriptor], [IOACPIAddressSpaceID], [IOPMPowerState])
//   - libkern: Access the runtime support and base classes of the kernel library. ([MD5_CTX])
//
// # BSD
//
//   - hfs: Access HFS file-system data structures. ([HFSCatalogFile], [HFSCatalogFolder], [HFSCatalogKey], [HFSCatalogThread], [HFSExtentDescriptor])
//   - net: Access network-related utilities. ([Ifnet_attach_proto_param], [Ifnet_attach_proto_param_v2], [Ifnet_demux_desc], [Ifnet_init_params])
//   - sys: Access general system utilities for time, file systems, and system information. ([Vnode_attr])
//
// # Mach
//
//   - [mach-o]: Access interfaces associated with the Mach-O runtime. ([Dyld_info_command], [Dyld_uuid_info_32], [Dyld_uuid_info_64], [Dyld_uuid_info_64_v2])
//
// # Utilities
//
//   - AppleDSP: Perform digital signal processing on data. ([IIRChannel])
//
// # Additional Reference
//
//   - [Kernel Functions] ([Nlist])
//   - [Kernel Structures] ([BTHeaderRec], [BTNodeDescriptor], [Boot_Video], [Boot_VideoV1], [CS_BlobIndex])
//   - [Kernel Data Types] ([AVIDType], [AbsoluteTime], [BDDiscInfo], [BDFeatures], [BDMediaType])
//   - [Kernel Enumerations] ([EFI_MEMORY_TYPE], [EFI_RESET_TYPE], [EXBrightMessageType], [EXDisplayPipeIndicator], [IOAudioDevicePowerState])
//   - [Kernel Constants] ([Arcade_upcall_subsystem], [Audit_triggers_subsystem], [Catch_exc_subsystem], [Catch_mach_exc_subsystem], [Clock_reply_subsystem])//
//
// [Debugging a custom kernel extension]: https://developer.apple.com/documentation/apple-silicon/debugging-a-custom-kernel-extension
// [Driver Support]: https://developer.apple.com/documentation/kernel/driver_support
// [Generating a Non-Maskable Interrupt]: https://developer.apple.com/documentation/kernel/generating_a_non-maskable_interrupt
// [Hardware Families]: https://developer.apple.com/documentation/kernel/hardware_families
// [IOKit Fundamentals]: https://developer.apple.com/documentation/kernel/iokit_fundamentals
// [Implementing drivers, system extensions, and kexts]: https://developer.apple.com/documentation/kernel/implementing_drivers_system_extensions_and_kexts
// [Installing a custom kernel extension]: https://developer.apple.com/documentation/apple-silicon/installing-a-custom-kernel-extension
// [Kernel Constants]: https://developer.apple.com/documentation/kernel/kernel_constants
// [Kernel Data Types]: https://developer.apple.com/documentation/kernel/kernel_data_types
// [Kernel Enumerations]: https://developer.apple.com/documentation/kernel/kernel_enumerations
// [Kernel Functions]: https://developer.apple.com/documentation/kernel/kernel_functions
// [Kernel Structures]: https://developer.apple.com/documentation/kernel/kernel_structures
// [mach-o]: https://developer.apple.com/documentation/kernel/mach-o
package kernel

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the kernel library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/usr/lib/libSystem.B.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: kernel: failed to load framework from any known path\n")
	}
}
