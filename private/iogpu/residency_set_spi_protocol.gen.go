// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLResidencySetSPI protocol.
type MTLResidencySetSPI interface {
	objectivec.IObject

	// AllCommittedAllocations protocol.
	AllCommittedAllocations() objectivec.IObject

	// CurrentGeneration protocol.
	CurrentGeneration() uint64

	// ExpiredGeneration protocol.
	ExpiredGeneration() uint64

	// GenerationForAllocation protocol.
	GenerationForAllocation(allocation objectivec.IObject) uint64

	// SetCurrentGeneration protocol.
	SetCurrentGeneration(generation uint64)

	// SetExpiredGeneration protocol.
	SetExpiredGeneration(generation uint64)
}

// MTLResidencySetSPIObject wraps an existing Objective-C object that conforms to the MTLResidencySetSPI protocol.
type MTLResidencySetSPIObject struct {
	objectivec.Object
}

func (o MTLResidencySetSPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLResidencySetSPIObjectFromID constructs a [MTLResidencySetSPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLResidencySetSPIObjectFromID(id objc.ID) MTLResidencySetSPIObject {
	return MTLResidencySetSPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLResidencySetSPIObject) AllCommittedAllocations() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("allCommittedAllocations"))
	return objectivec.Object{ID: rv}
}
func (o MTLResidencySetSPIObject) CurrentGeneration() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("currentGeneration"))
	return rv
}
func (o MTLResidencySetSPIObject) ExpiredGeneration() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("expiredGeneration"))
	return rv
}
func (o MTLResidencySetSPIObject) GenerationForAllocation(allocation objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("generationForAllocation:"), allocation)
	return rv
}
func (o MTLResidencySetSPIObject) SetCurrentGeneration(generation uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setCurrentGeneration:"), generation)
}
func (o MTLResidencySetSPIObject) SetExpiredGeneration(generation uint64) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setExpiredGeneration:"), generation)
}
