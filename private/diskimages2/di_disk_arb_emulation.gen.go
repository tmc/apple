// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [DIDiskArbEmulation] class.
var (
	_DIDiskArbEmulationClass     DIDiskArbEmulationClass
	_DIDiskArbEmulationClassOnce sync.Once
)

func getDIDiskArbEmulationClass() DIDiskArbEmulationClass {
	_DIDiskArbEmulationClassOnce.Do(func() {
		_DIDiskArbEmulationClass = DIDiskArbEmulationClass{class: objc.GetClass("DIDiskArbEmulation")}
	})
	return _DIDiskArbEmulationClass
}

// GetDIDiskArbEmulationClass returns the class object for DIDiskArbEmulation.
func GetDIDiskArbEmulationClass() DIDiskArbEmulationClass {
	return getDIDiskArbEmulationClass()
}

type DIDiskArbEmulationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIDiskArbEmulationClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIDiskArbEmulationClass) Alloc() DIDiskArbEmulation {
	rv := objc.Send[DIDiskArbEmulation](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArbEmulation
type DIDiskArbEmulation struct {
	DIDiskArb
}

// DIDiskArbEmulationFromID constructs a [DIDiskArbEmulation] from an objc.ID.
func DIDiskArbEmulationFromID(id objc.ID) DIDiskArbEmulation {
	return DIDiskArbEmulation{DIDiskArb: DIDiskArbFromID(id)}
}

// Ensure DIDiskArbEmulation implements IDIDiskArbEmulation.
var _ IDIDiskArbEmulation = DIDiskArbEmulation{}

// An interface definition for the [DIDiskArbEmulation] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArbEmulation
type IDIDiskArbEmulation interface {
	IDIDiskArb
}

// Init initializes the instance.
func (d DIDiskArbEmulation) Init() DIDiskArbEmulation {
	rv := objc.Send[DIDiskArbEmulation](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIDiskArbEmulation) Autorelease() DIDiskArbEmulation {
	rv := objc.Send[DIDiskArbEmulation](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIDiskArbEmulation creates a new DIDiskArbEmulation instance.
func NewDIDiskArbEmulation() DIDiskArbEmulation {
	class := getDIDiskArbEmulationClass()
	rv := objc.Send[DIDiskArbEmulation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDiskArb/initWithError:
func NewDIDiskArbEmulationWithError() (DIDiskArbEmulation, error) {
	var errorPtr objc.ID
	instance := getDIDiskArbEmulationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIDiskArbEmulation{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIDiskArbEmulationFromID(rv), nil
}
