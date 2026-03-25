// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEFIVariableStore] class.
var (
	_VZEFIVariableStoreClass     VZEFIVariableStoreClass
	_VZEFIVariableStoreClassOnce sync.Once
)

func getVZEFIVariableStoreClass() VZEFIVariableStoreClass {
	_VZEFIVariableStoreClassOnce.Do(func() {
		_VZEFIVariableStoreClass = VZEFIVariableStoreClass{class: objc.GetClass("VZEFIVariableStore")}
	})
	return _VZEFIVariableStoreClass
}

// GetVZEFIVariableStoreClass returns the class object for VZEFIVariableStore.
func GetVZEFIVariableStoreClass() VZEFIVariableStoreClass {
	return getVZEFIVariableStoreClass()
}

type VZEFIVariableStoreClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFIVariableStoreClass) Alloc() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the Extensible Firmware Interface (EFI) variable
// store that contains NVRAM variables the EFI exposes.
//
// # Creating the variable store
//
//   - [VZEFIVariableStore.InitCreatingVariableStoreAtURLOptionsError]: Creates a new EFI variable store at specified the URL on the filesystem, initialization options, and error-return variable.
//   - [VZEFIVariableStore.InitWithURL]: Initialize the variable store from the URL of an existing file.
//
// # Instance properties
//
//   - [VZEFIVariableStore.URL]: The URL of the variable store on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore
type VZEFIVariableStore struct {
	objectivec.Object
}

// VZEFIVariableStoreFromID constructs a [VZEFIVariableStore] from an objc.ID.
//
// An object that represents the Extensible Firmware Interface (EFI) variable
// store that contains NVRAM variables the EFI exposes.
func VZEFIVariableStoreFromID(id objc.ID) VZEFIVariableStore {
	return VZEFIVariableStore{objectivec.Object{ID: id}}
}
// NOTE: VZEFIVariableStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZEFIVariableStore] class.
//
// # Creating the variable store
//
//   - [IVZEFIVariableStore.InitCreatingVariableStoreAtURLOptionsError]: Creates a new EFI variable store at specified the URL on the filesystem, initialization options, and error-return variable.
//   - [IVZEFIVariableStore.InitWithURL]: Initialize the variable store from the URL of an existing file.
//
// # Instance properties
//
//   - [IVZEFIVariableStore.URL]: The URL of the variable store on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore
type IVZEFIVariableStore interface {
	objectivec.IObject

	// Topic: Creating the variable store

	// Creates a new EFI variable store at specified the URL on the filesystem, initialization options, and error-return variable.
	InitCreatingVariableStoreAtURLOptionsError(URL foundation.INSURL, options VZEFIVariableStoreInitializationOptions) (VZEFIVariableStore, error)
	// Initialize the variable store from the URL of an existing file.
	InitWithURL(URL foundation.INSURL) VZEFIVariableStore

	// Topic: Instance properties

	// The URL of the variable store on the local file system.
	URL() foundation.INSURL
}

// Init initializes the instance.
func (e VZEFIVariableStore) Init() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e VZEFIVariableStore) Autorelease() VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFIVariableStore creates a new VZEFIVariableStore instance.
func NewVZEFIVariableStore() VZEFIVariableStore {
	class := getVZEFIVariableStoreClass()
	rv := objc.Send[VZEFIVariableStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new EFI variable store at specified the URL on the filesystem,
// initialization options, and error-return variable.
//
// URL: A URL that specifies the location on disk at which to store the EFI
// information.
//
// options: An array of possible [VZEFIVariableStore.InitializationOptions].
// //
// [VZEFIVariableStore.InitializationOptions]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/InitializationOptions
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(creatingVariableStoreAt:options:)
func NewEFIVariableStoreCreatingVariableStoreAtURLOptionsError(URL foundation.INSURL, options VZEFIVariableStoreInitializationOptions) (VZEFIVariableStore, error) {
	var errorPtr objc.ID
	instance := getVZEFIVariableStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initCreatingVariableStoreAtURL:options:error:"), URL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZEFIVariableStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZEFIVariableStoreFromID(rv), nil
}

// Initialize the variable store from the URL of an existing file.
//
// URL: The URL of the location on disk that contains the stored EFI information.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(url:)
func NewEFIVariableStoreWithURL(URL foundation.INSURL) VZEFIVariableStore {
	instance := getVZEFIVariableStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return VZEFIVariableStoreFromID(rv)
}

// Creates a new EFI variable store at specified the URL on the filesystem,
// initialization options, and error-return variable.
//
// URL: A URL that specifies the location on disk at which to store the EFI
// information.
//
// options: An array of possible [VZEFIVariableStore.InitializationOptions].
// //
// [VZEFIVariableStore.InitializationOptions]: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/InitializationOptions
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(creatingVariableStoreAt:options:)
func (e VZEFIVariableStore) InitCreatingVariableStoreAtURLOptionsError(URL foundation.INSURL, options VZEFIVariableStoreInitializationOptions) (VZEFIVariableStore, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initCreatingVariableStoreAtURL:options:error:"), URL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZEFIVariableStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZEFIVariableStoreFromID(rv), nil

}
// Initialize the variable store from the URL of an existing file.
//
// URL: The URL of the location on disk that contains the stored EFI information.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/init(url:)
func (e VZEFIVariableStore) InitWithURL(URL foundation.INSURL) VZEFIVariableStore {
	rv := objc.Send[VZEFIVariableStore](e.ID, objc.Sel("initWithURL:"), URL)
	return rv
}

// The URL of the variable store on the local file system.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIVariableStore/url
func (e VZEFIVariableStore) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

