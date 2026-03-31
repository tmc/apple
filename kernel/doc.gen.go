// Code generated from Apple documentation for kernel. DO NOT EDIT.

// Package kernel provides Go bindings for the kernel framework.
//
// Develop kernel-resident device drivers and kernel extensions.
//
// The Kernel Framework provides the APIs and support for kernel-resident device drivers and other kernel extensions. It defines the base class for I/O Kit device drivers ([IOService](<doc://com.apple.documentation/documentation/kernel/ioservice-5h>)), several helper classes, and the families that support many types of devices.
//
// # Kernel Extensions
//
//   - Implementing drivers, system extensions, and kexts: Create drivers and system extensions to communicate with hardware and provide low-level services, and only use kernel extensions for a few tasks.
//   - Installing a custom kernel extension: Install kernel extensions using a custom installer package, and help users understand the installation process.
//   - Debugging a custom kernel extension: Configure your system to enable the debugging of custom kernel extensions from a second Mac.
//   - Generating a Non-Maskable Interrupt: Interrupt the kernel on a target Mac and attach a remote debugger to it.
//
// # IOKit Drivers
//
//   - IOKit Fundamentals: Implement a driver for your custom hardware using a third-party kernel extension. ([IOService], [IOPlatformIO], [IORegistryEntry], [IORegistryIterator], [IOSharedDataQueue])
//   - Hardware Families: Add support for specific hardware protocols such as USB, and for standard network, serial, audio, and graphics interfaces.
//   - Driver Support: Explore the device registry and access power-management utilities and other shared driver features. ([ApplePlatformExpert], [IODTPlatformExpert], [IOPlatformExpert], [IOPlatformExpertDevice], [IOPlatformDevice])
//   - libkern: Access the runtime support and base classes of the kernel library. ([MD5_CTX])
//
// # BSD
//
//   - architecture: Access machine-level and architectural information about the current platform. ([Cpuid_arch_perf_leaf_t], [Cpuid_cache_desc_t], [Cpuid_mwait_leaf_t], [Cpuid_thermal_leaf_t], [Cpuid_tsc_leaf_t])
//   - bsm: Audit resource usage on the system. ([Au_asflgs_t], [Au_asid_t], [Au_class_t], [Au_ctlmode_t], [Au_emod_t])
//   - hfs: Access HFS file-system data structures. ([HFSCatalogFile], [HFSCatalogFolder], [HFSCatalogKey], [HFSCatalogThread], [HFSExtentDescriptor])
//   - kern: Access kernel-level interfaces including clock, task, kernel extension, lock, and compression utilities.
//   - Math: Perform mathematical operations and manipulate integer, float, and double values.
//   - miscfs: Access device nodes and other file-system entities.
//   - net: Access network-related utilities. ([Ifnet_attach_proto_param], [Ifnet_attach_proto_param_v2], [Ifnet_demux_desc], [Ifnet_family_t], [Ifnet_init_params])
//   - Strings: Compare, convert, and catenate strings and access the resulting content of those strings.
//   - sys: Access general system utilities for time, file systems, and system information. ([Vnode_attr], [Vnode_fsparam], [Vnode_t], [Vnodeopv_desc], [Vnodeopv_entry_desc])
//   - vfs: Access the virtual file-system interfaces.
//   - vm: Interact with the virtual memory system.
//
// # Mach
//
//   - mach: Access Mach interfaces including processor, memory, thread, and semaphore support. ([Mach_msg_audit_trailer_t], [Mach_msg_base_t], [Mach_msg_body_t], [Mach_msg_context_trailer_t], [Mach_msg_empty_rcv_t])
//   - mach-o: Access interfaces associated with the Mach-O runtime. ([Dyld_info_command], [Dyld_kernel_image_info_array_t], [Dyld_uuid_info_32], [Dyld_uuid_info_64], [Dyld_uuid_info_64_v2])
//
// # Utilities
//
//   - Debugging: Debug your kernel extensions using the kernel debugger, assertions, exceptions, backtraces, and logging.
//   - AppleDSP: Perform digital signal processing on data. ([IIRChannel])
//
// # Additional Reference
//
//   - Kernel Functions ([Nlist])
//   - Kernel Structures ([BTHeaderRec], [BTNodeDescriptor], [Boot_Video], [Boot_VideoV1], [CS_BlobIndex])
//   - Kernel Data Types ([AVIDType], [AbsoluteTime], [AsyncPendingTrans], [BDDiscInfo], [BDFeatures])
//   - Kernel Enumerations ([EFI_MEMORY_TYPE], [EFI_RESET_TYPE], [EXBrightMessageType], [EXDisplayPipeIndicator], [IOAudioDevicePowerState])
//   - Kernel Constants ([Arcade_upcall_subsystem], [Audit_triggers_subsystem], [Catch_exc_subsystem], [Catch_mach_exc_subsystem], [Clock_reply_subsystem])
//
// # Classes
//
//   - IOCatalogue: In-kernel database for IOKit driver personalities.
//   - IOEventLink
//   - IOEventLinkInterface
//   - IOGuardPageMemoryDescriptor
//   - IOHIDTranslationService
//   - IOServiceStateNotificationDispatchSource
//   - IOServiceStateNotificationDispatchSourceInterface
//   - IOWorkGroup
//   - IOWorkGroupInterface
//   - OSAction_IOHIDEventService__CopyEvent
//   - OSAction_IOHIDEventService__CopyEventInterface
//   - OSAction_IOHIDEventService__SetLED
//   - OSAction_IOHIDEventService__SetLEDInterface
//   - OSAction_IOHIDEventService__SetUserProperties
//   - OSAction_IOHIDEventService__SetUserPropertiesInterface
//
// [kernel Documentation]: https://developer.apple.com/documentation/kernel
package kernel

import (
	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the kernel library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/kernel.framework/kernel",
	"/usr/lib/libkernel.dylib",
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
}
