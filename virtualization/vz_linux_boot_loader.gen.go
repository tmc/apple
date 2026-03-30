// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZLinuxBootLoader] class.
var (
	_VZLinuxBootLoaderClass     VZLinuxBootLoaderClass
	_VZLinuxBootLoaderClassOnce sync.Once
)

func getVZLinuxBootLoaderClass() VZLinuxBootLoaderClass {
	_VZLinuxBootLoaderClassOnce.Do(func() {
		_VZLinuxBootLoaderClass = VZLinuxBootLoaderClass{class: objc.GetClass("VZLinuxBootLoader")}
	})
	return _VZLinuxBootLoaderClass
}

// GetVZLinuxBootLoaderClass returns the class object for VZLinuxBootLoader.
func GetVZLinuxBootLoaderClass() VZLinuxBootLoaderClass {
	return getVZLinuxBootLoaderClass()
}

type VZLinuxBootLoaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinuxBootLoaderClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxBootLoaderClass) Alloc() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that loads and configures a Linux kernel as the guest system of
// your VM.
//
// # Overview
//
// Create and configure a [VZLinuxBootLoader] object during the initial
// configuration of your VM. Use this object to specify the location of the
// Linux kernel that serves as the guest operating system. You can also
// specify additional information to use during the boot process, such as
// command-line parameters to pass to the kernel. Assign the
// [VZLinuxBootLoader] object to the [BootLoader] property of your
// [VZVirtualMachineConfiguration] object. A configuration with
// [VZLinuxBootLoader] is only valid if used with
// [VZGenericPlatformConfiguration].
//
// # Creating the Linux boot loader
//
//   - [VZLinuxBootLoader.InitWithKernelURL]: Creates a boot loader that launches the Linux kernel at the specified URL.
//
// # Configuring the boot parameters
//
//   - [VZLinuxBootLoader.CommandLine]: The command-line parameters to pass to the Linux kernel at boot time.
//   - [VZLinuxBootLoader.SetCommandLine]
//   - [VZLinuxBootLoader.InitialRamdiskURL]: The location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.
//   - [VZLinuxBootLoader.SetInitialRamdiskURL]
//
// # Getting the kernel file
//
//   - [VZLinuxBootLoader.KernelURL]: The URL of the Linux kernel file.
//   - [VZLinuxBootLoader.SetKernelURL]
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader
type VZLinuxBootLoader struct {
	VZBootLoader
}

// VZLinuxBootLoaderFromID constructs a [VZLinuxBootLoader] from an objc.ID.
//
// An object that loads and configures a Linux kernel as the guest system of
// your VM.
func VZLinuxBootLoaderFromID(id objc.ID) VZLinuxBootLoader {
	return VZLinuxBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}

// NOTE: VZLinuxBootLoader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZLinuxBootLoader] class.
//
// # Creating the Linux boot loader
//
//   - [IVZLinuxBootLoader.InitWithKernelURL]: Creates a boot loader that launches the Linux kernel at the specified URL.
//
// # Configuring the boot parameters
//
//   - [IVZLinuxBootLoader.CommandLine]: The command-line parameters to pass to the Linux kernel at boot time.
//   - [IVZLinuxBootLoader.SetCommandLine]
//   - [IVZLinuxBootLoader.InitialRamdiskURL]: The location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.
//   - [IVZLinuxBootLoader.SetInitialRamdiskURL]
//
// # Getting the kernel file
//
//   - [IVZLinuxBootLoader.KernelURL]: The URL of the Linux kernel file.
//   - [IVZLinuxBootLoader.SetKernelURL]
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader
type IVZLinuxBootLoader interface {
	IVZBootLoader

	// Topic: Creating the Linux boot loader

	// Creates a boot loader that launches the Linux kernel at the specified URL.
	InitWithKernelURL(kernelURL foundation.INSURL) VZLinuxBootLoader

	// Topic: Configuring the boot parameters

	// The command-line parameters to pass to the Linux kernel at boot time.
	CommandLine() string
	SetCommandLine(value string)
	// The location of an optional RAM disk, which the boot loader maps into memory before it boots the Linux kernel.
	InitialRamdiskURL() foundation.INSURL
	SetInitialRamdiskURL(value foundation.INSURL)

	// Topic: Getting the kernel file

	// The URL of the Linux kernel file.
	KernelURL() foundation.INSURL
	SetKernelURL(value foundation.INSURL)

	// The guest system to boot when the VM starts.
	BootLoader() IVZBootLoader
	SetBootLoader(value IVZBootLoader)
}

// Init initializes the instance.
func (l VZLinuxBootLoader) Init() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l VZLinuxBootLoader) Autorelease() VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxBootLoader creates a new VZLinuxBootLoader instance.
func NewVZLinuxBootLoader() VZLinuxBootLoader {
	class := getVZLinuxBootLoaderClass()
	rv := objc.Send[VZLinuxBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a boot loader that launches the Linux kernel at the specified URL.
//
// kernelURL: The location of a Linux kernel on the local file system.
//
// # Return Value
//
// A boot loader object for the specified Linux kernel.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/init(kernelURL:)
func NewLinuxBootLoaderWithKernelURL(kernelURL foundation.INSURL) VZLinuxBootLoader {
	instance := getVZLinuxBootLoaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKernelURL:"), kernelURL)
	return VZLinuxBootLoaderFromID(rv)
}

// Creates a boot loader that launches the Linux kernel at the specified URL.
//
// kernelURL: The location of a Linux kernel on the local file system.
//
// # Return Value
//
// A boot loader object for the specified Linux kernel.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/init(kernelURL:)
func (l VZLinuxBootLoader) InitWithKernelURL(kernelURL foundation.INSURL) VZLinuxBootLoader {
	rv := objc.Send[VZLinuxBootLoader](l.ID, objc.Sel("initWithKernelURL:"), kernelURL)
	return rv
}

// The command-line parameters to pass to the Linux kernel at boot time.
//
// # Discussion
//
// For information about the parameters you can pass to a Linux kernel, see
// “[The kernel’s command-line parameters]”.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/commandLine
//
// [The kernel’s command-line parameters]: https://www.kernel.org/doc/html/latest/admin-guide/kernel-parameters.html
func (l VZLinuxBootLoader) CommandLine() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("commandLine"))
	return foundation.NSStringFromID(rv).String()
}
func (l VZLinuxBootLoader) SetCommandLine(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setCommandLine:"), objc.String(value))
}

// The location of an optional RAM disk, which the boot loader maps into
// memory before it boots the Linux kernel.
//
// # Discussion
//
// The default value of this property is `nil`. If you want specific files to
// be available when your Linux kernel boots, provide a URL to a valid RAM
// disk file in this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/initialRamdiskURL
func (l VZLinuxBootLoader) InitialRamdiskURL() foundation.INSURL {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("initialRamdiskURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (l VZLinuxBootLoader) SetInitialRamdiskURL(value foundation.INSURL) {
	objc.Send[struct{}](l.ID, objc.Sel("setInitialRamdiskURL:"), value)
}

// The URL of the Linux kernel file.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxBootLoader/kernelURL
func (l VZLinuxBootLoader) KernelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("kernelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (l VZLinuxBootLoader) SetKernelURL(value foundation.INSURL) {
	objc.Send[struct{}](l.ID, objc.Sel("setKernelURL:"), value)
}

// The guest system to boot when the VM starts.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/bootloader
func (l VZLinuxBootLoader) BootLoader() IVZBootLoader {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("bootLoader"))
	return VZBootLoaderFromID(objc.ID(rv))
}
func (l VZLinuxBootLoader) SetBootLoader(value IVZBootLoader) {
	objc.Send[struct{}](l.ID, objc.Sel("setBootLoader:"), value)
}
