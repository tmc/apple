// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseBroadcastableTransposes] class.
var (
	_EspressoPassFuseBroadcastableTransposesClass     EspressoPassFuseBroadcastableTransposesClass
	_EspressoPassFuseBroadcastableTransposesClassOnce sync.Once
)

func getEspressoPassFuseBroadcastableTransposesClass() EspressoPassFuseBroadcastableTransposesClass {
	_EspressoPassFuseBroadcastableTransposesClassOnce.Do(func() {
		_EspressoPassFuseBroadcastableTransposesClass = EspressoPassFuseBroadcastableTransposesClass{class: objc.GetClass("EspressoPass_fuse_broadcastable_transposes")}
	})
	return _EspressoPassFuseBroadcastableTransposesClass
}

// GetEspressoPassFuseBroadcastableTransposesClass returns the class object for EspressoPass_fuse_broadcastable_transposes.
func GetEspressoPassFuseBroadcastableTransposesClass() EspressoPassFuseBroadcastableTransposesClass {
	return getEspressoPassFuseBroadcastableTransposesClass()
}

type EspressoPassFuseBroadcastableTransposesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseBroadcastableTransposesClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseBroadcastableTransposesClass) Alloc() EspressoPassFuseBroadcastableTransposes {
	rv := objc.Send[EspressoPassFuseBroadcastableTransposes](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFuseBroadcastableTransposes struct {
	EspressoCustomPass
}

// EspressoPassFuseBroadcastableTransposesFromID constructs a [EspressoPassFuseBroadcastableTransposes] from an objc.ID.
func EspressoPassFuseBroadcastableTransposesFromID(id objc.ID) EspressoPassFuseBroadcastableTransposes {
	return EspressoPassFuseBroadcastableTransposes{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_broadcastable_transposesFromID is an alias for [EspressoPassFuseBroadcastableTransposesFromID] for cross-framework compatibility.
func EspressoPass_fuse_broadcastable_transposesFromID(id objc.ID) EspressoPassFuseBroadcastableTransposes {
	return EspressoPassFuseBroadcastableTransposesFromID(id)
}

// Ensure EspressoPassFuseBroadcastableTransposes implements IEspressoPassFuseBroadcastableTransposes.
var _ IEspressoPassFuseBroadcastableTransposes = EspressoPassFuseBroadcastableTransposes{}

// An interface definition for the [EspressoPassFuseBroadcastableTransposes] class.
type IEspressoPassFuseBroadcastableTransposes interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseBroadcastableTransposes) Init() EspressoPassFuseBroadcastableTransposes {
	rv := objc.Send[EspressoPassFuseBroadcastableTransposes](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseBroadcastableTransposes) Autorelease() EspressoPassFuseBroadcastableTransposes {
	rv := objc.Send[EspressoPassFuseBroadcastableTransposes](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseBroadcastableTransposes creates a new EspressoPassFuseBroadcastableTransposes instance.
func NewEspressoPassFuseBroadcastableTransposes() EspressoPassFuseBroadcastableTransposes {
	class := getEspressoPassFuseBroadcastableTransposesClass()
	rv := objc.Send[EspressoPassFuseBroadcastableTransposes](objc.ID(class.class), objc.Sel("new"))
	return rv
}
