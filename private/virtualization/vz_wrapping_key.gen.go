// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZWrappingKey] class.
var (
	_VZWrappingKeyClass     VZWrappingKeyClass
	_VZWrappingKeyClassOnce sync.Once
)

func getVZWrappingKeyClass() VZWrappingKeyClass {
	_VZWrappingKeyClassOnce.Do(func() {
		_VZWrappingKeyClass = VZWrappingKeyClass{class: objc.GetClass("_VZWrappingKey")}
	})
	return _VZWrappingKeyClass
}

// GetVZWrappingKeyClass returns the class object for _VZWrappingKey.
func GetVZWrappingKeyClass() VZWrappingKeyClass {
	return getVZWrappingKeyClass()
}

type VZWrappingKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZWrappingKeyClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZWrappingKeyClass) Alloc() VZWrappingKey {
	rv := objc.Send[VZWrappingKey](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZWrappingKey.InitWithAESKeyError]
//   - [VZWrappingKey.InitWithAsymmetricKeyError]
//   - [VZWrappingKey.InitWithPasswordError]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey
type VZWrappingKey struct {
	objectivec.Object
}

// VZWrappingKeyFromID constructs a [VZWrappingKey] from an objc.ID.
func VZWrappingKeyFromID(id objc.ID) VZWrappingKey {
	return VZWrappingKey{objectivec.Object{ID: id}}
}

// Ensure VZWrappingKey implements IVZWrappingKey.
var _ IVZWrappingKey = VZWrappingKey{}

// An interface definition for the [VZWrappingKey] class.
//
// # Methods
//
//   - [IVZWrappingKey.InitWithAESKeyError]
//   - [IVZWrappingKey.InitWithAsymmetricKeyError]
//   - [IVZWrappingKey.InitWithPasswordError]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey
type IVZWrappingKey interface {
	objectivec.IObject

	// Topic: Methods

	InitWithAESKeyError(aESKey objectivec.IObject) (VZWrappingKey, error)
	InitWithAsymmetricKeyError(key string) (VZWrappingKey, error)
	InitWithPasswordError(password objectivec.IObject) (VZWrappingKey, error)
}

// Init initializes the instance.
func (v VZWrappingKey) Init() VZWrappingKey {
	rv := objc.Send[VZWrappingKey](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZWrappingKey) Autorelease() VZWrappingKey {
	rv := objc.Send[VZWrappingKey](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZWrappingKey creates a new VZWrappingKey instance.
func NewVZWrappingKey() VZWrappingKey {
	class := getVZWrappingKeyClass()
	rv := objc.Send[VZWrappingKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey/initWithAESKey:error:
func NewVZWrappingKeyWithAESKeyError(aESKey objectivec.IObject) (VZWrappingKey, error) {
	var errorPtr objc.ID
	instance := getVZWrappingKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAESKey:error:"), aESKey, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZWrappingKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZWrappingKeyFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey/initWithAsymmetricKey:error:
func NewVZWrappingKeyWithAsymmetricKeyError(key string) (VZWrappingKey, error) {
	var errorPtr objc.ID
	instance := getVZWrappingKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsymmetricKey:error:"), objc.String(key), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZWrappingKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZWrappingKeyFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey/initWithPassword:error:
func NewVZWrappingKeyWithPasswordError(password objectivec.IObject) (VZWrappingKey, error) {
	var errorPtr objc.ID
	instance := getVZWrappingKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPassword:error:"), password, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZWrappingKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZWrappingKeyFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey/initWithAESKey:error:
func (v VZWrappingKey) InitWithAESKeyError(aESKey objectivec.IObject) (VZWrappingKey, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithAESKey:error:"), aESKey, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZWrappingKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZWrappingKeyFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey/initWithAsymmetricKey:error:
func (v VZWrappingKey) InitWithAsymmetricKeyError(key string) (VZWrappingKey, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithAsymmetricKey:error:"), objc.String(key), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZWrappingKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZWrappingKeyFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZWrappingKey/initWithPassword:error:
func (v VZWrappingKey) InitWithPasswordError(password objectivec.IObject) (VZWrappingKey, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithPassword:error:"), password, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZWrappingKey{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZWrappingKeyFromID(rv), nil

}
