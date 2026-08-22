// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLResourceSPI protocol.
type MTLResourceSPI interface {
	objectivec.IObject

	// AllocationID protocol.
	AllocationID() uint64

	// DoesAliasAllResourcesCount protocol.
	DoesAliasAllResourcesCount(resources []objectivec.IObject, count uint64) bool

	// DoesAliasAnyResourcesCount protocol.
	DoesAliasAnyResourcesCount(resources []objectivec.IObject, count uint64) bool

	// DoesAliasResource protocol.
	DoesAliasResource(resource objectivec.IObject) bool

	// IsComplete protocol.
	IsComplete() bool

	// IsPurgeable protocol.
	IsPurgeable() bool

	// IsWriteComplete protocol.
	IsWriteComplete() bool

	// ProtectionOptions protocol.
	ProtectionOptions() uint64

	// ResponsibleProcess protocol.
	ResponsibleProcess() int

	// SetOwnerWithIdentity protocol.
	SetOwnerWithIdentity(identity uint32) int

	// SetResponsibleProcess protocol.
	SetResponsibleProcess(process int)

	// UnfilteredResourceOptions protocol.
	UnfilteredResourceOptions() uint64

	// WaitUntilComplete protocol.
	WaitUntilComplete()
}

// MTLResourceSPIObject wraps an existing Objective-C object that conforms to the MTLResourceSPI protocol.
type MTLResourceSPIObject struct {
	objectivec.Object
}

func (o MTLResourceSPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLResourceSPIObjectFromID constructs a [MTLResourceSPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLResourceSPIObjectFromID(id objc.ID) MTLResourceSPIObject {
	return MTLResourceSPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLResourceSPIObject) AllocationID() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("allocationID"))
	return rv
}
func (o MTLResourceSPIObject) DoesAliasAllResourcesCount(resources []objectivec.IObject, count uint64) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("doesAliasAllResources:count:"), objc.CArray(resources), count)
	return rv
}
func (o MTLResourceSPIObject) DoesAliasAnyResourcesCount(resources []objectivec.IObject, count uint64) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("doesAliasAnyResources:count:"), objc.CArray(resources), count)
	return rv
}
func (o MTLResourceSPIObject) DoesAliasResource(resource objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("doesAliasResource:"), resource)
	return rv
}
func (o MTLResourceSPIObject) IsComplete() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isComplete"))
	return rv
}
func (o MTLResourceSPIObject) IsPurgeable() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isPurgeable"))
	return rv
}
func (o MTLResourceSPIObject) IsWriteComplete() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isWriteComplete"))
	return rv
}
func (o MTLResourceSPIObject) ProtectionOptions() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("protectionOptions"))
	return rv
}
func (o MTLResourceSPIObject) ResponsibleProcess() int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("responsibleProcess"))
	return rv
}
func (o MTLResourceSPIObject) SetOwnerWithIdentity(identity uint32) int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("setOwnerWithIdentity:"), identity)
	return rv
}
func (o MTLResourceSPIObject) SetResponsibleProcess(process int) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setResponsibleProcess:"), process)
}
func (o MTLResourceSPIObject) UnfilteredResourceOptions() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("unfilteredResourceOptions"))
	return rv
}
func (o MTLResourceSPIObject) WaitUntilComplete() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("waitUntilComplete"))
}
