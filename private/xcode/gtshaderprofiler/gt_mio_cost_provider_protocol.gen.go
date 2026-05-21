// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioCostProvider protocol.
type GTMioCostProvider interface {
	objectivec.IObject

	// CostCount protocol.
	CostCount() uint64

	// CostForScopeScopeIdentifierCost protocol.
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost GTMioCostInfo) bool

	// Costs protocol.
	Costs() unsafe.Pointer

	// InstructionCountForScopeScopeIdentifierDataMaster protocol.
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64

	// TotalCostForScopeScopeIdentifierDataMaster protocol.
	TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64
}

// GTMioCostProviderObject wraps an existing Objective-C object that conforms to the GTMioCostProvider protocol.
type GTMioCostProviderObject struct {
	objectivec.Object
}

func (o GTMioCostProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTMioCostProviderObjectFromID constructs a [GTMioCostProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTMioCostProviderObjectFromID(id objc.ID) GTMioCostProviderObject {
	return GTMioCostProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTMioCostProviderObject) CostCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("costCount"))
	return rv
}
func (o GTMioCostProviderObject) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost GTMioCostInfo) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}
func (o GTMioCostProviderObject) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("costs"))
	return rv
}
func (o GTMioCostProviderObject) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
func (o GTMioCostProviderObject) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
