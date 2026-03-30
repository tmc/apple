// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioCostProvider protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider
type GTMioCostProvider interface {
	objectivec.IObject

	// CostCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/costCount
	CostCount() uint64

	// CostForScopeScopeIdentifierCost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/costForScope:scopeIdentifier:cost:
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool

	// Costs protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/costs
	Costs() unsafe.Pointer

	// InstructionCountForScopeScopeIdentifierDataMaster protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/instructionCountForScope:scopeIdentifier:dataMaster:
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64

	// TotalCostForScopeScopeIdentifierDataMaster protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/totalCostForScope:scopeIdentifier:dataMaster:
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

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/costCount
func (o GTMioCostProviderObject) CostCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("costCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/costForScope:scopeIdentifier:cost:
func (o GTMioCostProviderObject) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/costs
func (o GTMioCostProviderObject) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("costs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/instructionCountForScope:scopeIdentifier:dataMaster:
func (o GTMioCostProviderObject) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCostProvider/totalCostForScope:scopeIdentifier:dataMaster:
func (o GTMioCostProviderObject) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}
