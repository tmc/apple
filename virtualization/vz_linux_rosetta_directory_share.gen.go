// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZLinuxRosettaDirectoryShare] class.
var (
	_VZLinuxRosettaDirectoryShareClass     VZLinuxRosettaDirectoryShareClass
	_VZLinuxRosettaDirectoryShareClassOnce sync.Once
)

func getVZLinuxRosettaDirectoryShareClass() VZLinuxRosettaDirectoryShareClass {
	_VZLinuxRosettaDirectoryShareClassOnce.Do(func() {
		_VZLinuxRosettaDirectoryShareClass = VZLinuxRosettaDirectoryShareClass{class: objc.GetClass("VZLinuxRosettaDirectoryShare")}
	})
	return _VZLinuxRosettaDirectoryShareClass
}

// GetVZLinuxRosettaDirectoryShareClass returns the class object for VZLinuxRosettaDirectoryShare.
func GetVZLinuxRosettaDirectoryShareClass() VZLinuxRosettaDirectoryShareClass {
	return getVZLinuxRosettaDirectoryShareClass()
}

type VZLinuxRosettaDirectoryShareClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinuxRosettaDirectoryShareClass) Alloc() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The Linux directory share for Rosetta.
//
// # Overview
// 
// This directory share exposes the Rosetta directory from the host file
// system to the guest. The example below shows the process of creating a
// [VZVirtualMachineConfiguration], and then associating the Rosetta directory
// share with the VM configuration.
// 
// For complete instructions on installing Rosetta see [Running Intel Binaries
// in Linux VMs with Rosetta], which includes additional information about
// checking for Rosetta availability, mounting the directory share, and
// registering the Rosetta runtime binary to run Intel binaries in a guest VM.
// 
// For information on using a custom kernel to enhance Rosetta performance,
// see [Accelerating the performance of Rosetta].
//
// [Accelerating the performance of Rosetta]: https://developer.apple.com/documentation/Virtualization/accelerating-the-performance-of-rosetta
// [Running Intel Binaries in Linux VMs with Rosetta]: https://developer.apple.com/documentation/Virtualization/running-intel-binaries-in-linux-vms-with-rosetta
//
// # Creating a Rosetta directory share
//
//   - [VZLinuxRosettaDirectoryShare.InitWithError]: Creates a new Rosetta directory share, or returns an error if Rosetta isn’t installed.
//
// # Setting the ahead of time (AOT) caching options
//
//   - [VZLinuxRosettaDirectoryShare.CachingOptions]: The value that enables translation caching and configures the socket communication type for Rosetta.
//   - [VZLinuxRosettaDirectoryShare.SetCachingOptions]
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare
type VZLinuxRosettaDirectoryShare struct {
	VZDirectoryShare
}

// VZLinuxRosettaDirectoryShareFromID constructs a [VZLinuxRosettaDirectoryShare] from an objc.ID.
//
// The Linux directory share for Rosetta.
func VZLinuxRosettaDirectoryShareFromID(id objc.ID) VZLinuxRosettaDirectoryShare {
	return VZLinuxRosettaDirectoryShare{VZDirectoryShare: VZDirectoryShareFromID(id)}
}
// NOTE: VZLinuxRosettaDirectoryShare adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZLinuxRosettaDirectoryShare] class.
//
// # Creating a Rosetta directory share
//
//   - [IVZLinuxRosettaDirectoryShare.InitWithError]: Creates a new Rosetta directory share, or returns an error if Rosetta isn’t installed.
//
// # Setting the ahead of time (AOT) caching options
//
//   - [IVZLinuxRosettaDirectoryShare.CachingOptions]: The value that enables translation caching and configures the socket communication type for Rosetta.
//   - [IVZLinuxRosettaDirectoryShare.SetCachingOptions]
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare
type IVZLinuxRosettaDirectoryShare interface {
	IVZDirectoryShare

	// Topic: Creating a Rosetta directory share

	// Creates a new Rosetta directory share, or returns an error if Rosetta isn’t installed.
	InitWithError() (VZLinuxRosettaDirectoryShare, error)

	// Topic: Setting the ahead of time (AOT) caching options

	// The value that enables translation caching and configures the socket communication type for Rosetta.
	CachingOptions() uint
	SetCachingOptions(value uint)

	// The value that enables translation caching and configures the socket communication type for Rosetta.
	Options() IVZLinuxRosettaCachingOptions
	SetOptions(value IVZLinuxRosettaCachingOptions)
}





// Init initializes the instance.
func (l VZLinuxRosettaDirectoryShare) Init() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l VZLinuxRosettaDirectoryShare) Autorelease() VZLinuxRosettaDirectoryShare {
	rv := objc.Send[VZLinuxRosettaDirectoryShare](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinuxRosettaDirectoryShare creates a new VZLinuxRosettaDirectoryShare instance.
func NewVZLinuxRosettaDirectoryShare() VZLinuxRosettaDirectoryShare {
	class := getVZLinuxRosettaDirectoryShareClass()
	rv := objc.Send[VZLinuxRosettaDirectoryShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a new Rosetta directory share, or returns an error if Rosetta
// isn’t installed.
//
// # Discussion
// 
// Check the status of Rosetta by examining the [Availability] class property
// before creating a new Rosetta directory share to ensure the capability is
// both supported and available on host Mac. For complete instructions on
// installing Rosetta see [Running Intel Binaries in Linux VMs with Rosetta].
//
// [Running Intel Binaries in Linux VMs with Rosetta]: https://developer.apple.com/documentation/Virtualization/running-intel-binaries-in-linux-vms-with-rosetta
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/init()
func NewLinuxRosettaDirectoryShareWithError() (VZLinuxRosettaDirectoryShare, error) {
	var errorPtr objc.ID
	instance := getVZLinuxRosettaDirectoryShareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZLinuxRosettaDirectoryShareFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VZLinuxRosettaDirectoryShareFromID(rv), nil
}







// Creates a new Rosetta directory share, or returns an error if Rosetta
// isn’t installed.
//
// # Discussion
// 
// Check the status of Rosetta by examining the [Availability] class property
// before creating a new Rosetta directory share to ensure the capability is
// both supported and available on host Mac. For complete instructions on
// installing Rosetta see [Running Intel Binaries in Linux VMs with Rosetta].
//
// [Running Intel Binaries in Linux VMs with Rosetta]: https://developer.apple.com/documentation/Virtualization/running-intel-binaries-in-linux-vms-with-rosetta
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/init()
func (l VZLinuxRosettaDirectoryShare) InitWithError() (VZLinuxRosettaDirectoryShare, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZLinuxRosettaDirectoryShare{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZLinuxRosettaDirectoryShareFromID(rv), nil

}





// Starts the installation of Rosetta.
//
// completionHandler: The completion handler the framework invokes after the request finishes
// processing.
//
// # Discussion
// 
// The completion handler returns an error object that contains information
// about a problem, or `nil` if the installation completed successfully.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/installRosetta(completionHandler:)
func (_VZLinuxRosettaDirectoryShareClass VZLinuxRosettaDirectoryShareClass) InstallRosettaWithCompletionHandler(completionHandler ErrorHandler) {
		_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
		objc.Send[objc.ID](objc.ID(_VZLinuxRosettaDirectoryShareClass.class), objc.Sel("installRosettaWithCompletionHandler:"), _block0)
}








// The value that enables translation caching and configures the socket
// communication type for Rosetta.
//
// See: https://developer.apple.com/documentation/virtualization/vzlinuxrosettadirectoryshare/cachingoptions-swift.property
func (l VZLinuxRosettaDirectoryShare) CachingOptions() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("cachingOptions"))
	return rv
}
func (l VZLinuxRosettaDirectoryShare) SetCachingOptions(value uint) {
	objc.Send[struct{}](l.ID, objc.Sel("setCachingOptions:"), value)
}



// The value that enables translation caching and configures the socket
// communication type for Rosetta.
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/options
func (l VZLinuxRosettaDirectoryShare) Options() IVZLinuxRosettaCachingOptions {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("options"))
	return VZLinuxRosettaCachingOptionsFromID(objc.ID(rv))
}
func (l VZLinuxRosettaDirectoryShare) SetOptions(value IVZLinuxRosettaCachingOptions) {
	objc.Send[struct{}](l.ID, objc.Sel("setOptions:"), value)
}







// A value that indicates the current state of Rosetta’s availability.
//
// # Discussion
// 
// The value is one of the possible [VZLinuxRosettaAvailability] values.
//
// [VZLinuxRosettaAvailability]: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaAvailability
//
// See: https://developer.apple.com/documentation/Virtualization/VZLinuxRosettaDirectoryShare/availability
func (_VZLinuxRosettaDirectoryShareClass VZLinuxRosettaDirectoryShareClass) Availability() VZLinuxRosettaAvailability {
	rv := objc.Send[VZLinuxRosettaAvailability](objc.ID(_VZLinuxRosettaDirectoryShareClass.class), objc.Sel("availability"))
	return VZLinuxRosettaAvailability(rv)
}














// InstallRosetta is a synchronous wrapper around [VZLinuxRosettaDirectoryShare.InstallRosettaWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (lc VZLinuxRosettaDirectoryShareClass) InstallRosetta(ctx context.Context) error {
	done := make(chan error, 1)
	lc.InstallRosettaWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}






