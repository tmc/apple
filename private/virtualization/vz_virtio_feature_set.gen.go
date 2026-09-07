// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioFeatureSet] class.
var (
	_VZVirtioFeatureSetClass     VZVirtioFeatureSetClass
	_VZVirtioFeatureSetClassOnce sync.Once
)

func getVZVirtioFeatureSetClass() VZVirtioFeatureSetClass {
	_VZVirtioFeatureSetClassOnce.Do(func() {
		_VZVirtioFeatureSetClass = VZVirtioFeatureSetClass{class: objc.GetClass("VZVirtioFeatureSet")}
	})
	return _VZVirtioFeatureSetClass
}

// GetVZVirtioFeatureSetClass returns the class object for VZVirtioFeatureSet.
func GetVZVirtioFeatureSetClass() VZVirtioFeatureSetClass {
	return getVZVirtioFeatureSetClass()
}

type VZVirtioFeatureSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioFeatureSetClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioFeatureSetClass) Alloc() VZVirtioFeatureSet {
	rv := objc.SendIfResponds[VZVirtioFeatureSet](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioFeatureSet.Subset0]
//   - [VZVirtioFeatureSet.SetSubset0]
//   - [VZVirtioFeatureSet.Subset1]
//   - [VZVirtioFeatureSet.SetSubset1]
type VZVirtioFeatureSet struct {
	objectivec.Object
}

// VZVirtioFeatureSetFromID constructs a [VZVirtioFeatureSet] from an objc.ID.
func VZVirtioFeatureSetFromID(id objc.ID) VZVirtioFeatureSet {
	return VZVirtioFeatureSet{objectivec.Object{ID: id}}
}

// Ensure VZVirtioFeatureSet implements IVZVirtioFeatureSet.
var _ IVZVirtioFeatureSet = VZVirtioFeatureSet{}

// An interface definition for the [VZVirtioFeatureSet] class.
//
// # Methods
//
//   - [IVZVirtioFeatureSet.Subset0]
//   - [IVZVirtioFeatureSet.SetSubset0]
//   - [IVZVirtioFeatureSet.Subset1]
//   - [IVZVirtioFeatureSet.SetSubset1]
type IVZVirtioFeatureSet interface {
	objectivec.IObject

	// Topic: Methods

	Subset0() uint32
	SetSubset0(value uint32)
	Subset1() uint32
	SetSubset1(value uint32)
}

// Init initializes the instance.
func (v VZVirtioFeatureSet) Init() VZVirtioFeatureSet {
	rv := objc.SendIfResponds[VZVirtioFeatureSet](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioFeatureSet) Autorelease() VZVirtioFeatureSet {
	rv := objc.SendIfResponds[VZVirtioFeatureSet](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioFeatureSet creates a new VZVirtioFeatureSet instance.
func NewVZVirtioFeatureSet() VZVirtioFeatureSet {
	class := getVZVirtioFeatureSetClass()
	rv := objc.SendIfResponds[VZVirtioFeatureSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtioFeatureSet) Subset0() uint32 {
	rv := objc.SendIfResponds[uint32](v.ID, objc.Sel("subset0"))
	return rv
}
func (v VZVirtioFeatureSet) SetSubset0(value uint32) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setSubset0:"), value)
}
func (v VZVirtioFeatureSet) Subset1() uint32 {
	rv := objc.SendIfResponds[uint32](v.ID, objc.Sel("subset1"))
	return rv
}
func (v VZVirtioFeatureSet) SetSubset1(value uint32) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setSubset1:"), value)
}
