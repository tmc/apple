// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassMergePytorchLayernorm] class.
var (
	_EspressoPassMergePytorchLayernormClass     EspressoPassMergePytorchLayernormClass
	_EspressoPassMergePytorchLayernormClassOnce sync.Once
)

func getEspressoPassMergePytorchLayernormClass() EspressoPassMergePytorchLayernormClass {
	_EspressoPassMergePytorchLayernormClassOnce.Do(func() {
		_EspressoPassMergePytorchLayernormClass = EspressoPassMergePytorchLayernormClass{class: objc.GetClass("EspressoPass_merge_pytorch_layernorm")}
	})
	return _EspressoPassMergePytorchLayernormClass
}

// GetEspressoPassMergePytorchLayernormClass returns the class object for EspressoPass_merge_pytorch_layernorm.
func GetEspressoPassMergePytorchLayernormClass() EspressoPassMergePytorchLayernormClass {
	return getEspressoPassMergePytorchLayernormClass()
}

type EspressoPassMergePytorchLayernormClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassMergePytorchLayernormClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassMergePytorchLayernormClass) Alloc() EspressoPassMergePytorchLayernorm {
	rv := objc.Send[EspressoPassMergePytorchLayernorm](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_pytorch_layernorm
type EspressoPassMergePytorchLayernorm struct {
	EspressoCustomPass
}

// EspressoPassMergePytorchLayernormFromID constructs a [EspressoPassMergePytorchLayernorm] from an objc.ID.
func EspressoPassMergePytorchLayernormFromID(id objc.ID) EspressoPassMergePytorchLayernorm {
	return EspressoPassMergePytorchLayernorm{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_merge_pytorch_layernormFromID is an alias for [EspressoPassMergePytorchLayernormFromID] for cross-framework compatibility.
func EspressoPass_merge_pytorch_layernormFromID(id objc.ID) EspressoPassMergePytorchLayernorm {
	return EspressoPassMergePytorchLayernormFromID(id)
}

// Ensure EspressoPassMergePytorchLayernorm implements IEspressoPassMergePytorchLayernorm.
var _ IEspressoPassMergePytorchLayernorm = EspressoPassMergePytorchLayernorm{}

// An interface definition for the [EspressoPassMergePytorchLayernorm] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_pytorch_layernorm
type IEspressoPassMergePytorchLayernorm interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassMergePytorchLayernorm) Init() EspressoPassMergePytorchLayernorm {
	rv := objc.Send[EspressoPassMergePytorchLayernorm](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassMergePytorchLayernorm) Autorelease() EspressoPassMergePytorchLayernorm {
	rv := objc.Send[EspressoPassMergePytorchLayernorm](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassMergePytorchLayernorm creates a new EspressoPassMergePytorchLayernorm instance.
func NewEspressoPassMergePytorchLayernorm() EspressoPassMergePytorchLayernorm {
	class := getEspressoPassMergePytorchLayernormClass()
	rv := objc.Send[EspressoPassMergePytorchLayernorm](objc.ID(class.class), objc.Sel("new"))
	return rv
}
