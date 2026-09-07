// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANECompileFlavorPolicy] class.
var (
	_ANECompileFlavorPolicyClass     ANECompileFlavorPolicyClass
	_ANECompileFlavorPolicyClassOnce sync.Once
)

func getANECompileFlavorPolicyClass() ANECompileFlavorPolicyClass {
	_ANECompileFlavorPolicyClassOnce.Do(func() {
		_ANECompileFlavorPolicyClass = ANECompileFlavorPolicyClass{class: objc.GetClass("_ANECompileFlavorPolicy")}
	})
	return _ANECompileFlavorPolicyClass
}

// GetANECompileFlavorPolicyClass returns the class object for _ANECompileFlavorPolicy.
func GetANECompileFlavorPolicyClass() ANECompileFlavorPolicyClass {
	return getANECompileFlavorPolicyClass()
}

type ANECompileFlavorPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANECompileFlavorPolicyClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANECompileFlavorPolicyClass) Alloc() ANECompileFlavorPolicy {
	rv := objc.SendIfResponds[ANECompileFlavorPolicy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

type ANECompileFlavorPolicy struct {
	objectivec.Object
}

// ANECompileFlavorPolicyFromID constructs a [ANECompileFlavorPolicy] from an objc.ID.
func ANECompileFlavorPolicyFromID(id objc.ID) ANECompileFlavorPolicy {
	return ANECompileFlavorPolicy{objectivec.Object{ID: id}}
}

// Ensure ANECompileFlavorPolicy implements IANECompileFlavorPolicy.
var _ IANECompileFlavorPolicy = ANECompileFlavorPolicy{}

// An interface definition for the [ANECompileFlavorPolicy] class.
type IANECompileFlavorPolicy interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANECompileFlavorPolicy) Init() ANECompileFlavorPolicy {
	rv := objc.SendIfResponds[ANECompileFlavorPolicy](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECompileFlavorPolicy) Autorelease() ANECompileFlavorPolicy {
	rv := objc.SendIfResponds[ANECompileFlavorPolicy](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECompileFlavorPolicy creates a new ANECompileFlavorPolicy instance.
func NewANECompileFlavorPolicy() ANECompileFlavorPolicy {
	class := getANECompileFlavorPolicyClass()
	rv := objc.SendIfResponds[ANECompileFlavorPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_ANECompileFlavorPolicyClass ANECompileFlavorPolicyClass) NonBondedCsIdentities() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANECompileFlavorPolicyClass.class), objc.Sel("nonBondedCsIdentities"))
	return objectivec.Object{ID: rv}
}
func (_ANECompileFlavorPolicyClass ANECompileFlavorPolicyClass) ShouldDisableBondedForCsIdentity(identity objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANECompileFlavorPolicyClass.class), objc.Sel("shouldDisableBondedForCsIdentity:"), identity)
	return rv
}
