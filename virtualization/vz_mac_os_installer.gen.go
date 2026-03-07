// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacOSInstaller] class.
var (
	_VZMacOSInstallerClass     VZMacOSInstallerClass
	_VZMacOSInstallerClassOnce sync.Once
)

func getVZMacOSInstallerClass() VZMacOSInstallerClass {
	_VZMacOSInstallerClassOnce.Do(func() {
		_VZMacOSInstallerClass = VZMacOSInstallerClass{class: objc.GetClass("VZMacOSInstaller")}
	})
	return _VZMacOSInstallerClass
}

// GetVZMacOSInstallerClass returns the class object for VZMacOSInstaller.
func GetVZMacOSInstallerClass() VZMacOSInstallerClass {
	return getVZMacOSInstallerClass()
}

type VZMacOSInstallerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSInstallerClass) Alloc() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object you use to install macOS on the specified virtual machine.
//
// # Overview
// 
// Initialize a [VZMacOSInstaller] object with a [VZVirtualMachine] and a file
// URL that refers to a macOS restore image.
// 
// The following code example shows how to use a `VZMacOSInstaller:`
//
// # Creating a macOS Installer
//
//   - [VZMacOSInstaller.InitWithVirtualMachineRestoreImageURL]: Creates a macOS installer object.
//
// # Getting Information About an Installation
//
//   - [VZMacOSInstaller.Progress]: A progress object that you can use to observe or cancel an installation.
//   - [VZMacOSInstaller.RestoreImageURL]: The restore image URL used to initialize this installer.
//   - [VZMacOSInstaller.VirtualMachine]: The virtual machine used to initialize this installer.
//
// # Installing macOS
//
//   - [VZMacOSInstaller.InstallWithCompletionHandler]: Start installing macOS.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller
type VZMacOSInstaller struct {
	objectivec.Object
}

// VZMacOSInstallerFromID constructs a [VZMacOSInstaller] from an objc.ID.
//
// An object you use to install macOS on the specified virtual machine.
func VZMacOSInstallerFromID(id objc.ID) VZMacOSInstaller {
	return VZMacOSInstaller{objectivec.Object{id}}
}
// NOTE: VZMacOSInstaller adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMacOSInstaller] class.
//
// # Creating a macOS Installer
//
//   - [IVZMacOSInstaller.InitWithVirtualMachineRestoreImageURL]: Creates a macOS installer object.
//
// # Getting Information About an Installation
//
//   - [IVZMacOSInstaller.Progress]: A progress object that you can use to observe or cancel an installation.
//   - [IVZMacOSInstaller.RestoreImageURL]: The restore image URL used to initialize this installer.
//   - [IVZMacOSInstaller.VirtualMachine]: The virtual machine used to initialize this installer.
//
// # Installing macOS
//
//   - [IVZMacOSInstaller.InstallWithCompletionHandler]: Start installing macOS.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller
type IVZMacOSInstaller interface {
	objectivec.IObject

	// Topic: Creating a macOS Installer

	// Creates a macOS installer object.
	InitWithVirtualMachineRestoreImageURL(virtualMachine IVZVirtualMachine, restoreImageFileURL foundation.INSURL) VZMacOSInstaller

	// Topic: Getting Information About an Installation

	// A progress object that you can use to observe or cancel an installation.
	Progress() foundation.NSProgress
	// The restore image URL used to initialize this installer.
	RestoreImageURL() foundation.INSURL
	// The virtual machine used to initialize this installer.
	VirtualMachine() IVZVirtualMachine

	// Topic: Installing macOS

	// Start installing macOS.
	InstallWithCompletionHandler(completionHandler ErrorHandler)
}





// Init initializes the instance.
func (m VZMacOSInstaller) Init() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSInstaller) Autorelease() VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSInstaller creates a new VZMacOSInstaller instance.
func NewVZMacOSInstaller() VZMacOSInstaller {
	class := getVZMacOSInstallerClass()
	rv := objc.Send[VZMacOSInstaller](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a macOS installer object.
//
// virtualMachine: The virtual machine to install the operating system on.
//
// restoreImageFileURL: A file URL that indicates the macOS restore image to install.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/init(virtualMachine:restoringFromImageAt:)
func NewMacOSInstallerWithVirtualMachineRestoreImageURL(virtualMachine IVZVirtualMachine, restoreImageFileURL foundation.INSURL) VZMacOSInstaller {
	instance := getVZMacOSInstallerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:restoreImageURL:"), virtualMachine, restoreImageFileURL)
	return VZMacOSInstallerFromID(rv)
}







// Creates a macOS installer object.
//
// virtualMachine: The virtual machine to install the operating system on.
//
// restoreImageFileURL: A file URL that indicates the macOS restore image to install.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/init(virtualMachine:restoringFromImageAt:)
func (m VZMacOSInstaller) InitWithVirtualMachineRestoreImageURL(virtualMachine IVZVirtualMachine, restoreImageFileURL foundation.INSURL) VZMacOSInstaller {
	rv := objc.Send[VZMacOSInstaller](m.ID, objc.Sel("initWithVirtualMachine:restoreImageURL:"), virtualMachine, restoreImageFileURL)
	return rv
}

// Start installing macOS.
//
// # Discussion
// 
// This method starts the installation process. The VM must be in a stopped
// state. During the installation operation, pausing or stopping the VM
// results in an undefined behavior.
// 
// If you start the installation on the same [VZMacOSInstaller] object more
// than once, the framework raises an exception.
// 
// Call this method only on the VM’s queue.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/install()
func (m VZMacOSInstaller) InstallWithCompletionHandler(completionHandler ErrorHandler) {
		_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
		objc.Send[objc.ID](m.ID, objc.Sel("installWithCompletionHandler:"), _block0)
}












// A progress object that you can use to observe or cancel an installation.
//
// # Discussion
// 
// Canceling the progress object before starting an installation raises an
// exception.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/progress
func (m VZMacOSInstaller) Progress() foundation.NSProgress {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progress"))
	return foundation.NSProgressFromID(objc.ID(rv))
}



// The restore image URL used to initialize this installer.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/restoreImageURL
func (m VZMacOSInstaller) RestoreImageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("restoreImageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}



// The virtual machine used to initialize this installer.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSInstaller/virtualMachine
func (m VZMacOSInstaller) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}


















// Install is a synchronous wrapper around [VZMacOSInstaller.InstallWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m VZMacOSInstaller) Install(ctx context.Context) error {
	done := make(chan error, 1)
	m.InstallWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}






