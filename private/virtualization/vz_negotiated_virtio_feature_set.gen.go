// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZNegotiatedVirtioFeatureSet] class.
var (
	_VZNegotiatedVirtioFeatureSetClass     VZNegotiatedVirtioFeatureSetClass
	_VZNegotiatedVirtioFeatureSetClassOnce sync.Once
)

func getVZNegotiatedVirtioFeatureSetClass() VZNegotiatedVirtioFeatureSetClass {
	_VZNegotiatedVirtioFeatureSetClassOnce.Do(func() {
		_VZNegotiatedVirtioFeatureSetClass = VZNegotiatedVirtioFeatureSetClass{class: objc.GetClass("VZNegotiatedVirtioFeatureSet")}
	})
	return _VZNegotiatedVirtioFeatureSetClass
}

// GetVZNegotiatedVirtioFeatureSetClass returns the class object for VZNegotiatedVirtioFeatureSet.
func GetVZNegotiatedVirtioFeatureSetClass() VZNegotiatedVirtioFeatureSetClass {
	return getVZNegotiatedVirtioFeatureSetClass()
}

type VZNegotiatedVirtioFeatureSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZNegotiatedVirtioFeatureSetClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNegotiatedVirtioFeatureSetClass) Alloc() VZNegotiatedVirtioFeatureSet {
	rv := objc.SendIfResponds[VZNegotiatedVirtioFeatureSet](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZNegotiatedVirtioFeatureSet.Subset0]
//   - [VZNegotiatedVirtioFeatureSet.Subset1]
type VZNegotiatedVirtioFeatureSet struct {
	objectivec.Object
}

// VZNegotiatedVirtioFeatureSetFromID constructs a [VZNegotiatedVirtioFeatureSet] from an objc.ID.
func VZNegotiatedVirtioFeatureSetFromID(id objc.ID) VZNegotiatedVirtioFeatureSet {
	return VZNegotiatedVirtioFeatureSet{objectivec.Object{ID: id}}
}

// Ensure VZNegotiatedVirtioFeatureSet implements IVZNegotiatedVirtioFeatureSet.
var _ IVZNegotiatedVirtioFeatureSet = VZNegotiatedVirtioFeatureSet{}

// An interface definition for the [VZNegotiatedVirtioFeatureSet] class.
//
// # Methods
//
//   - [IVZNegotiatedVirtioFeatureSet.Subset0]
//   - [IVZNegotiatedVirtioFeatureSet.Subset1]
type IVZNegotiatedVirtioFeatureSet interface {
	objectivec.IObject

	// Topic: Methods

	Subset0() uint32
	Subset1() uint32
}

// Init initializes the instance.
func (v VZNegotiatedVirtioFeatureSet) Init() VZNegotiatedVirtioFeatureSet {
	rv := objc.SendIfResponds[VZNegotiatedVirtioFeatureSet](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZNegotiatedVirtioFeatureSet) Autorelease() VZNegotiatedVirtioFeatureSet {
	rv := objc.SendIfResponds[VZNegotiatedVirtioFeatureSet](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNegotiatedVirtioFeatureSet creates a new VZNegotiatedVirtioFeatureSet instance.
func NewVZNegotiatedVirtioFeatureSet() VZNegotiatedVirtioFeatureSet {
	class := getVZNegotiatedVirtioFeatureSetClass()
	rv := objc.SendIfResponds[VZNegotiatedVirtioFeatureSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZNegotiatedVirtioFeatureSet) Subset0() uint32 {
	rv := objc.SendIfResponds[uint32](v.ID, objc.Sel("subset0"))
	return rv
}
func (v VZNegotiatedVirtioFeatureSet) Subset1() uint32 {
	rv := objc.SendIfResponds[uint32](v.ID, objc.Sel("subset1"))
	return rv
}
