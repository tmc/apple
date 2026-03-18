// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ISA] class.
var (
	_ISAClass     ISAClass
	_ISAClassOnce sync.Once
)

func getISAClass() ISAClass {
	_ISAClassOnce.Do(func() {
		_ISAClass = ISAClass{class: objc.GetClass("isa")}
	})
	return _ISAClass
}

// GetISAClass returns the class object for isa.
func GetISAClass() ISAClass {
	return getISAClass()
}

type ISAClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ic ISAClass) Alloc() ISA {
	rv := objc.Send[ISA](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSProxy/isa
type ISA struct {
	objectivec.Object
}

// ISAFromID constructs a [ISA] from an objc.ID.
func ISAFromID(id objc.ID) ISA {
	return ISA{objectivec.Object{ID: id}}
}
// Ensure ISA implements IISA.
var _ IISA = ISA{}

// An interface definition for the [ISA] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy/isa
type IISA interface {
	objectivec.IObject
}

// Init initializes the instance.
func (i ISA) Init() ISA {
	rv := objc.Send[ISA](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i ISA) Autorelease() ISA {
	rv := objc.Send[ISA](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewISA creates a new ISA instance.
func NewISA() ISA {
	class := getISAClass()
	rv := objc.Send[ISA](objc.ID(class.class), objc.Sel("new"))
	return rv
}

