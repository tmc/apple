// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoBrickTensorCPU] class.
var (
	_EspressoBrickTensorCPUClass     EspressoBrickTensorCPUClass
	_EspressoBrickTensorCPUClassOnce sync.Once
)

func getEspressoBrickTensorCPUClass() EspressoBrickTensorCPUClass {
	_EspressoBrickTensorCPUClassOnce.Do(func() {
		_EspressoBrickTensorCPUClass = EspressoBrickTensorCPUClass{class: objc.GetClass("EspressoBrickTensorCPU")}
	})
	return _EspressoBrickTensorCPUClass
}

// GetEspressoBrickTensorCPUClass returns the class object for EspressoBrickTensorCPU.
func GetEspressoBrickTensorCPUClass() EspressoBrickTensorCPUClass {
	return getEspressoBrickTensorCPUClass()
}

type EspressoBrickTensorCPUClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoBrickTensorCPUClass) Alloc() EspressoBrickTensorCPU {
	rv := objc.Send[EspressoBrickTensorCPU](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [EspressoBrickTensorCPU.RawPointer]
//   - [EspressoBrickTensorCPU.SetRawPointer]
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorCPU
type EspressoBrickTensorCPU struct {
	EspressoBrickTensor
}

// EspressoBrickTensorCPUFromID constructs a [EspressoBrickTensorCPU] from an objc.ID.
func EspressoBrickTensorCPUFromID(id objc.ID) EspressoBrickTensorCPU {
	return EspressoBrickTensorCPU{EspressoBrickTensor: EspressoBrickTensorFromID(id)}
}
// Ensure EspressoBrickTensorCPU implements IEspressoBrickTensorCPU.
var _ IEspressoBrickTensorCPU = EspressoBrickTensorCPU{}





// An interface definition for the [EspressoBrickTensorCPU] class.
//
// # Methods
//
//   - [IEspressoBrickTensorCPU.RawPointer]
//   - [IEspressoBrickTensorCPU.SetRawPointer]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorCPU
type IEspressoBrickTensorCPU interface {
	IEspressoBrickTensor

	// Topic: Methods

	RawPointer() unsafe.Pointer
	SetRawPointer(value unsafe.Pointer)
}





// Init initializes the instance.
func (e EspressoBrickTensorCPU) Init() EspressoBrickTensorCPU {
	rv := objc.Send[EspressoBrickTensorCPU](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoBrickTensorCPU) Autorelease() EspressoBrickTensorCPU {
	rv := objc.Send[EspressoBrickTensorCPU](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoBrickTensorCPU creates a new EspressoBrickTensorCPU instance.
func NewEspressoBrickTensorCPU() EspressoBrickTensorCPU {
	class := getEspressoBrickTensorCPUClass()
	rv := objc.Send[EspressoBrickTensorCPU](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// See: https://developer.apple.com/documentation/Espresso/EspressoBrickTensorCPU/rawPointer
func (e EspressoBrickTensorCPU) RawPointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("rawPointer"))
	return rv
}
func (e EspressoBrickTensorCPU) SetRawPointer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setRawPointer:"), value)
}

















