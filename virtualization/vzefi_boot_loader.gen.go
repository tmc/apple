// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZEFIBootLoader] class.
var (
	_VZEFIBootLoaderClass     VZEFIBootLoaderClass
	_VZEFIBootLoaderClassOnce sync.Once
)

func getVZEFIBootLoaderClass() VZEFIBootLoaderClass {
	_VZEFIBootLoaderClassOnce.Do(func() {
		_VZEFIBootLoaderClass = VZEFIBootLoaderClass{class: objc.GetClass("VZEFIBootLoader")}
	})
	return _VZEFIBootLoaderClass
}

// GetVZEFIBootLoaderClass returns the class object for VZEFIBootLoader.
func GetVZEFIBootLoaderClass() VZEFIBootLoaderClass {
	return getVZEFIBootLoaderClass()
}

type VZEFIBootLoaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFIBootLoaderClass) Alloc() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The boot loader configuration the system uses to boot guest-operating
// systems that expect an Extensible Firmware Interface (EFI) ROM.
//
// # Instance properties
//
//   - [VZEFIBootLoader.VariableStore]: The boot loader’s EFI variable store.
//   - [VZEFIBootLoader.SetVariableStore]
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader
type VZEFIBootLoader struct {
	VZBootLoader
}

// VZEFIBootLoaderFromID constructs a [VZEFIBootLoader] from an objc.ID.
//
// The boot loader configuration the system uses to boot guest-operating
// systems that expect an Extensible Firmware Interface (EFI) ROM.
func VZEFIBootLoaderFromID(id objc.ID) VZEFIBootLoader {
	return VZEFIBootLoader{VZBootLoader: VZBootLoaderFromID(id)}
}
// NOTE: VZEFIBootLoader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZEFIBootLoader] class.
//
// # Instance properties
//
//   - [IVZEFIBootLoader.VariableStore]: The boot loader’s EFI variable store.
//   - [IVZEFIBootLoader.SetVariableStore]
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader
type IVZEFIBootLoader interface {
	IVZBootLoader

	// Topic: Instance properties

	// The boot loader’s EFI variable store.
	VariableStore() IVZEFIVariableStore
	SetVariableStore(value IVZEFIVariableStore)
}





// Init initializes the instance.
func (e VZEFIBootLoader) Init() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e VZEFIBootLoader) Autorelease() VZEFIBootLoader {
	rv := objc.Send[VZEFIBootLoader](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFIBootLoader creates a new VZEFIBootLoader instance.
func NewVZEFIBootLoader() VZEFIBootLoader {
	class := getVZEFIBootLoaderClass()
	rv := objc.Send[VZEFIBootLoader](objc.ID(class.class), objc.Sel("new"))
	return rv
}






















// The boot loader’s EFI variable store.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEFIBootLoader/variableStore
func (e VZEFIBootLoader) VariableStore() IVZEFIVariableStore {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variableStore"))
	return VZEFIVariableStoreFromID(objc.ID(rv))
}
func (e VZEFIBootLoader) SetVariableStore(value IVZEFIVariableStore) {
	objc.Send[struct{}](e.ID, objc.Sel("setVariableStore:"), value)
}
























