// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
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

// Class returns the underlying Objective-C class pointer.
func (vc VZEFIVariableStoreClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFIVariableStoreClass) Alloc() VZEFIVariableStore {
	rv := objc.SendIfResponds[VZEFIVariableStore](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEFIVariableStore.DisableSecureBootWithError]
//   - [VZEFIVariableStore.EnableSecureBootUsingDefaultPlatformKeyWithError]
//   - [VZEFIVariableStore.EnableSecureBootWithPlatformKeyError]
//   - [VZEFIVariableStore.EnrollDefaultSecureBootSignaturesWithError]
//   - [VZEFIVariableStore.EnrollSecureBootSignaturesError]
//   - [VZEFIVariableStore.GetEnrolledSecureBootSignaturesWithError]
//   - [VZEFIVariableStore.GetSecureBootEnabledError]
//   - [VZEFIVariableStore.ResetSecureBootWithError]
type VZEFIVariableStore struct {
	objectivec.Object
}

// VZEFIVariableStoreFromID constructs a [VZEFIVariableStore] from an objc.ID.
func VZEFIVariableStoreFromID(id objc.ID) VZEFIVariableStore {
	return VZEFIVariableStore{objectivec.Object{ID: id}}
}

// Ensure VZEFIVariableStore implements IVZEFIVariableStore.
var _ IVZEFIVariableStore = VZEFIVariableStore{}

// An interface definition for the [VZEFIVariableStore] class.
//
// # Methods
//
//   - [IVZEFIVariableStore.DisableSecureBootWithError]
//   - [IVZEFIVariableStore.EnableSecureBootUsingDefaultPlatformKeyWithError]
//   - [IVZEFIVariableStore.EnableSecureBootWithPlatformKeyError]
//   - [IVZEFIVariableStore.EnrollDefaultSecureBootSignaturesWithError]
//   - [IVZEFIVariableStore.EnrollSecureBootSignaturesError]
//   - [IVZEFIVariableStore.GetEnrolledSecureBootSignaturesWithError]
//   - [IVZEFIVariableStore.GetSecureBootEnabledError]
//   - [IVZEFIVariableStore.ResetSecureBootWithError]
type IVZEFIVariableStore interface {
	objectivec.IObject

	// Topic: Methods

	DisableSecureBootWithError() (bool, error)
	EnableSecureBootUsingDefaultPlatformKeyWithError() (bool, error)
	EnableSecureBootWithPlatformKeyError(key security.SecCertificateRef) (bool, error)
	EnrollDefaultSecureBootSignaturesWithError() (bool, error)
	EnrollSecureBootSignaturesError(signatures objectivec.IObject) (bool, error)
	GetEnrolledSecureBootSignaturesWithError() (objectivec.IObject, error)
	GetSecureBootEnabledError() (bool, error)
	ResetSecureBootWithError() (bool, error)
}

// Init initializes the instance.
func (v VZEFIVariableStore) Init() VZEFIVariableStore {
	rv := objc.SendIfResponds[VZEFIVariableStore](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEFIVariableStore) Autorelease() VZEFIVariableStore {
	rv := objc.SendIfResponds[VZEFIVariableStore](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFIVariableStore creates a new VZEFIVariableStore instance.
func NewVZEFIVariableStore() VZEFIVariableStore {
	class := getVZEFIVariableStoreClass()
	rv := objc.SendIfResponds[VZEFIVariableStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZEFIVariableStore) DisableSecureBootWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("disableSecureBootWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("disableSecureBootWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (v VZEFIVariableStore) EnableSecureBootUsingDefaultPlatformKeyWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("enableSecureBootUsingDefaultPlatformKeyWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enableSecureBootUsingDefaultPlatformKeyWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (v VZEFIVariableStore) EnableSecureBootWithPlatformKeyError(key security.SecCertificateRef) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("enableSecureBootWithPlatformKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enableSecureBootWithPlatformKey:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (v VZEFIVariableStore) EnrollDefaultSecureBootSignaturesWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("enrollDefaultSecureBootSignaturesWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enrollDefaultSecureBootSignaturesWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (v VZEFIVariableStore) EnrollSecureBootSignaturesError(signatures objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("enrollSecureBootSignatures:error:"), signatures, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enrollSecureBootSignatures:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (v VZEFIVariableStore) GetEnrolledSecureBootSignaturesWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getEnrolledSecureBootSignaturesWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (v VZEFIVariableStore) GetSecureBootEnabledError() (bool, error) {
	var enabled bool
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("getSecureBootEnabled:error:"), unsafe.Pointer(&enabled), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getSecureBootEnabled:error: returned NO with nil NSError")
	}
	return enabled, nil
}
func (v VZEFIVariableStore) ResetSecureBootWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("resetSecureBootWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resetSecureBootWithError: returned NO with nil NSError")
	}
	return rv, nil

}
