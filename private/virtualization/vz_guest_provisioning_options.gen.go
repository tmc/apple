// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGuestProvisioningOptions] class.
var (
	_VZGuestProvisioningOptionsClass     VZGuestProvisioningOptionsClass
	_VZGuestProvisioningOptionsClassOnce sync.Once
)

func getVZGuestProvisioningOptionsClass() VZGuestProvisioningOptionsClass {
	_VZGuestProvisioningOptionsClassOnce.Do(func() {
		_VZGuestProvisioningOptionsClass = VZGuestProvisioningOptionsClass{class: objc.GetClass("VZGuestProvisioningOptions")}
	})
	return _VZGuestProvisioningOptionsClass
}

// GetVZGuestProvisioningOptionsClass returns the class object for VZGuestProvisioningOptions.
func GetVZGuestProvisioningOptionsClass() VZGuestProvisioningOptionsClass {
	return getVZGuestProvisioningOptionsClass()
}

type VZGuestProvisioningOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGuestProvisioningOptionsClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGuestProvisioningOptionsClass) Alloc() VZGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZGuestProvisioningOptions](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGuestProvisioningOptions._createDataIfNeededWithError]
//   - [VZGuestProvisioningOptions._init]
//   - [VZGuestProvisioningOptions.ValidateWithError]
type VZGuestProvisioningOptions struct {
	objectivec.Object
}

// VZGuestProvisioningOptionsFromID constructs a [VZGuestProvisioningOptions] from an objc.ID.
func VZGuestProvisioningOptionsFromID(id objc.ID) VZGuestProvisioningOptions {
	return VZGuestProvisioningOptions{objectivec.Object{ID: id}}
}

// Ensure VZGuestProvisioningOptions implements IVZGuestProvisioningOptions.
var _ IVZGuestProvisioningOptions = VZGuestProvisioningOptions{}

// An interface definition for the [VZGuestProvisioningOptions] class.
//
// # Methods
//
//   - [IVZGuestProvisioningOptions._createDataIfNeededWithError]
//   - [IVZGuestProvisioningOptions._init]
//   - [IVZGuestProvisioningOptions.ValidateWithError]
type IVZGuestProvisioningOptions interface {
	objectivec.IObject

	// Topic: Methods

	_createDataIfNeededWithError() (bool, error)
	_init() objectivec.IObject
	ValidateWithError() (bool, error)
}

// Init initializes the instance.
func (v VZGuestProvisioningOptions) Init() VZGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZGuestProvisioningOptions](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGuestProvisioningOptions) Autorelease() VZGuestProvisioningOptions {
	rv := objc.SendIfResponds[VZGuestProvisioningOptions](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGuestProvisioningOptions creates a new VZGuestProvisioningOptions instance.
func NewVZGuestProvisioningOptions() VZGuestProvisioningOptions {
	class := getVZGuestProvisioningOptionsClass()
	rv := objc.SendIfResponds[VZGuestProvisioningOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZGuestProvisioningOptions) _createDataIfNeededWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("_createDataIfNeededWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_createDataIfNeededWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// CreateDataIfNeededWithError is an exported wrapper for the private method _createDataIfNeededWithError.
func (v VZGuestProvisioningOptions) CreateDataIfNeededWithError() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_createDataIfNeededWithError:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createDataIfNeededWithError:"}
		return false, err
	}
	return v._createDataIfNeededWithError()
}

// CanCreateDataIfNeededWithError reports whether the receiver responds to the private selector _createDataIfNeededWithError:.
func (v VZGuestProvisioningOptions) CanCreateDataIfNeededWithError() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_createDataIfNeededWithError:"))
}
func (v VZGuestProvisioningOptions) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZGuestProvisioningOptions) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}
