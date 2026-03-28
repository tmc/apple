// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistoryNode] class.
var (
	_GTMioShaderExecutionHistoryNodeClass     GTMioShaderExecutionHistoryNodeClass
	_GTMioShaderExecutionHistoryNodeClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryNodeClass() GTMioShaderExecutionHistoryNodeClass {
	_GTMioShaderExecutionHistoryNodeClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryNodeClass = GTMioShaderExecutionHistoryNodeClass{class: objc.GetClass("GTMioShaderExecutionHistoryNode")}
	})
	return _GTMioShaderExecutionHistoryNodeClass
}

// GetGTMioShaderExecutionHistoryNodeClass returns the class object for GTMioShaderExecutionHistoryNode.
func GetGTMioShaderExecutionHistoryNodeClass() GTMioShaderExecutionHistoryNodeClass {
	return getGTMioShaderExecutionHistoryNodeClass()
}

type GTMioShaderExecutionHistoryNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryNodeClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryNodeClass) Alloc() GTMioShaderExecutionHistoryNode {
	rv := objc.Send[GTMioShaderExecutionHistoryNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioShaderExecutionHistoryNode._costForScopeScopeIdentifierCacheCost]
//   - [GTMioShaderExecutionHistoryNode._dfsEnumerator]
//   - [GTMioShaderExecutionHistoryNode.AddChild]
//   - [GTMioShaderExecutionHistoryNode.AddInstructionHitNumHits]
//   - [GTMioShaderExecutionHistoryNode.BeginTimestamp]
//   - [GTMioShaderExecutionHistoryNode.Binaries]
//   - [GTMioShaderExecutionHistoryNode.Children]
//   - [GTMioShaderExecutionHistoryNode.CliqueRoot]
//   - [GTMioShaderExecutionHistoryNode.ContainsInstructionBinaryIndex]
//   - [GTMioShaderExecutionHistoryNode.CostForScopeScopeIdentifierCost]
//   - [GTMioShaderExecutionHistoryNode.Dfs]
//   - [GTMioShaderExecutionHistoryNode.DurationPercentage]
//   - [GTMioShaderExecutionHistoryNode.EndTimestamp]
//   - [GTMioShaderExecutionHistoryNode.EndTimestamps]
//   - [GTMioShaderExecutionHistoryNode.EnumerateInstructions]
//   - [GTMioShaderExecutionHistoryNode.FunctionForIdentifier]
//   - [GTMioShaderExecutionHistoryNode.InstructionForInstructionIndex]
//   - [GTMioShaderExecutionHistoryNode.IsSameLocation]
//   - [GTMioShaderExecutionHistoryNode.LastInstructionIndex]
//   - [GTMioShaderExecutionHistoryNode.NormalizedInstructionCostScopeIdentifierBinaryCost]
//   - [GTMioShaderExecutionHistoryNode.OverlapsWithInstructionInstructionCountBinaryIndex]
//   - [GTMioShaderExecutionHistoryNode.Parent]
//   - [GTMioShaderExecutionHistoryNode.ParentFunction]
//   - [GTMioShaderExecutionHistoryNode.Root]
//   - [GTMioShaderExecutionHistoryNode.SelfCostForScopeScopeIdentifierCost]
//   - [GTMioShaderExecutionHistoryNode.StartTimestamps]
//   - [GTMioShaderExecutionHistoryNode.TimestampCount]
//   - [GTMioShaderExecutionHistoryNode.TopCostPercentage]
//   - [GTMioShaderExecutionHistoryNode.TotalDuration]
//   - [GTMioShaderExecutionHistoryNode.Type]
//   - [GTMioShaderExecutionHistoryNode.InitWithTypeParent]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode
type GTMioShaderExecutionHistoryNode struct {
	objectivec.Object
}

// GTMioShaderExecutionHistoryNodeFromID constructs a [GTMioShaderExecutionHistoryNode] from an objc.ID.
func GTMioShaderExecutionHistoryNodeFromID(id objc.ID) GTMioShaderExecutionHistoryNode {
	return GTMioShaderExecutionHistoryNode{objectivec.Object{ID: id}}
}
// Ensure GTMioShaderExecutionHistoryNode implements IGTMioShaderExecutionHistoryNode.
var _ IGTMioShaderExecutionHistoryNode = GTMioShaderExecutionHistoryNode{}

// An interface definition for the [GTMioShaderExecutionHistoryNode] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistoryNode._costForScopeScopeIdentifierCacheCost]
//   - [IGTMioShaderExecutionHistoryNode._dfsEnumerator]
//   - [IGTMioShaderExecutionHistoryNode.AddChild]
//   - [IGTMioShaderExecutionHistoryNode.AddInstructionHitNumHits]
//   - [IGTMioShaderExecutionHistoryNode.BeginTimestamp]
//   - [IGTMioShaderExecutionHistoryNode.Binaries]
//   - [IGTMioShaderExecutionHistoryNode.Children]
//   - [IGTMioShaderExecutionHistoryNode.CliqueRoot]
//   - [IGTMioShaderExecutionHistoryNode.ContainsInstructionBinaryIndex]
//   - [IGTMioShaderExecutionHistoryNode.CostForScopeScopeIdentifierCost]
//   - [IGTMioShaderExecutionHistoryNode.Dfs]
//   - [IGTMioShaderExecutionHistoryNode.DurationPercentage]
//   - [IGTMioShaderExecutionHistoryNode.EndTimestamp]
//   - [IGTMioShaderExecutionHistoryNode.EndTimestamps]
//   - [IGTMioShaderExecutionHistoryNode.EnumerateInstructions]
//   - [IGTMioShaderExecutionHistoryNode.FunctionForIdentifier]
//   - [IGTMioShaderExecutionHistoryNode.InstructionForInstructionIndex]
//   - [IGTMioShaderExecutionHistoryNode.IsSameLocation]
//   - [IGTMioShaderExecutionHistoryNode.LastInstructionIndex]
//   - [IGTMioShaderExecutionHistoryNode.NormalizedInstructionCostScopeIdentifierBinaryCost]
//   - [IGTMioShaderExecutionHistoryNode.OverlapsWithInstructionInstructionCountBinaryIndex]
//   - [IGTMioShaderExecutionHistoryNode.Parent]
//   - [IGTMioShaderExecutionHistoryNode.ParentFunction]
//   - [IGTMioShaderExecutionHistoryNode.Root]
//   - [IGTMioShaderExecutionHistoryNode.SelfCostForScopeScopeIdentifierCost]
//   - [IGTMioShaderExecutionHistoryNode.StartTimestamps]
//   - [IGTMioShaderExecutionHistoryNode.TimestampCount]
//   - [IGTMioShaderExecutionHistoryNode.TopCostPercentage]
//   - [IGTMioShaderExecutionHistoryNode.TotalDuration]
//   - [IGTMioShaderExecutionHistoryNode.Type]
//   - [IGTMioShaderExecutionHistoryNode.InitWithTypeParent]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode
type IGTMioShaderExecutionHistoryNode interface {
	objectivec.IObject

	// Topic: Methods

	_costForScopeScopeIdentifierCacheCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool
	_dfsEnumerator(_dfs uint32, enumerator VoidHandler)
	AddChild(child objectivec.IObject)
	AddInstructionHitNumHits(hit uint32, hits uint32)
	BeginTimestamp(timestamp uint64)
	Binaries() foundation.INSArray
	Children() foundation.INSArray
	CliqueRoot() objectivec.IObject
	ContainsInstructionBinaryIndex(instruction uint32, index uint32) bool
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool
	Dfs(dfs VoidHandler)
	DurationPercentage() float64
	EndTimestamp(timestamp uint64)
	EndTimestamps() unsafe.Pointer
	EnumerateInstructions(instructions VoidHandler)
	FunctionForIdentifier(identifier uint64) objectivec.IObject
	InstructionForInstructionIndex(index uint32) objectivec.IObject
	IsSameLocation(location objectivec.IObject) bool
	LastInstructionIndex() uint32
	NormalizedInstructionCostScopeIdentifierBinaryCost(cost uint16, identifier uint64, binary objectivec.IObject, cost2 unsafe.Pointer) bool
	OverlapsWithInstructionInstructionCountBinaryIndex(instruction uint32, count uint32, index uint32) bool
	Parent() IGTMioShaderExecutionHistoryNode
	ParentFunction() IGTMioShaderExecutionHistoryFunctionNode
	Root() IGTMioShaderExecutionHistoryRootNode
	SelfCostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool
	StartTimestamps() unsafe.Pointer
	TimestampCount() uint32
	TopCostPercentage() float64
	TotalDuration() uint64
	Type() uint32
	InitWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryNode) Init() GTMioShaderExecutionHistoryNode {
	rv := objc.Send[GTMioShaderExecutionHistoryNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryNode) Autorelease() GTMioShaderExecutionHistoryNode {
	rv := objc.Send[GTMioShaderExecutionHistoryNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryNode creates a new GTMioShaderExecutionHistoryNode instance.
func NewGTMioShaderExecutionHistoryNode() GTMioShaderExecutionHistoryNode {
	class := getGTMioShaderExecutionHistoryNodeClass()
	rv := objc.Send[GTMioShaderExecutionHistoryNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/initWithType:parent:
func NewGTMioShaderExecutionHistoryNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryNode {
	instance := getGTMioShaderExecutionHistoryNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/_costForScope:scopeIdentifier:cacheCost:
func (g GTMioShaderExecutionHistoryNode) _costForScopeScopeIdentifierCacheCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_costForScope:scopeIdentifier:cacheCost:"), scope, identifier, cost)
	return rv
}

// CostForScopeScopeIdentifierCacheCost is an exported wrapper for the private method _costForScopeScopeIdentifierCacheCost.
func (g GTMioShaderExecutionHistoryNode) CostForScopeScopeIdentifierCacheCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	return g._costForScopeScopeIdentifierCacheCost(scope, identifier, cost)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/_dfs:enumerator:
func (g GTMioShaderExecutionHistoryNode) _dfsEnumerator(_dfs uint32, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("_dfs:enumerator:"), _dfs, _block1)
}

// DfsEnumerator is an exported wrapper for the private method _dfsEnumerator.
func (g GTMioShaderExecutionHistoryNode) DfsEnumerator(_dfs uint32, enumerator VoidHandler) {
	g._dfsEnumerator(_dfs, enumerator)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/addChild:
func (g GTMioShaderExecutionHistoryNode) AddChild(child objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("addChild:"), child)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/addInstructionHit:numHits:
func (g GTMioShaderExecutionHistoryNode) AddInstructionHitNumHits(hit uint32, hits uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("addInstructionHit:numHits:"), hit, hits)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/beginTimestamp:
func (g GTMioShaderExecutionHistoryNode) BeginTimestamp(timestamp uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("beginTimestamp:"), timestamp)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/cliqueRoot
func (g GTMioShaderExecutionHistoryNode) CliqueRoot() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cliqueRoot"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/containsInstruction:binaryIndex:
func (g GTMioShaderExecutionHistoryNode) ContainsInstructionBinaryIndex(instruction uint32, index uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("containsInstruction:binaryIndex:"), instruction, index)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/costForScope:scopeIdentifier:cost:
func (g GTMioShaderExecutionHistoryNode) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/dfs:
func (g GTMioShaderExecutionHistoryNode) Dfs(dfs VoidHandler) {
_block0, _ := NewVoidBlock(dfs)
	objc.Send[objc.ID](g.ID, objc.Sel("dfs:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/endTimestamp:
func (g GTMioShaderExecutionHistoryNode) EndTimestamp(timestamp uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("endTimestamp:"), timestamp)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/enumerateInstructions:
func (g GTMioShaderExecutionHistoryNode) EnumerateInstructions(instructions VoidHandler) {
_block0, _ := NewVoidBlock(instructions)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateInstructions:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/functionForIdentifier:
func (g GTMioShaderExecutionHistoryNode) FunctionForIdentifier(identifier uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("functionForIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/instructionForInstructionIndex:
func (g GTMioShaderExecutionHistoryNode) InstructionForInstructionIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("instructionForInstructionIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/isSameLocation:
func (g GTMioShaderExecutionHistoryNode) IsSameLocation(location objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isSameLocation:"), location)
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/lastInstructionIndex
func (g GTMioShaderExecutionHistoryNode) LastInstructionIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("lastInstructionIndex"))
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/normalizedInstructionCost:scopeIdentifier:binary:cost:
func (g GTMioShaderExecutionHistoryNode) NormalizedInstructionCostScopeIdentifierBinaryCost(cost uint16, identifier uint64, binary objectivec.IObject, cost2 unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("normalizedInstructionCost:scopeIdentifier:binary:cost:"), cost, identifier, binary, cost2)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/overlapsWithInstruction:instructionCount:binaryIndex:
func (g GTMioShaderExecutionHistoryNode) OverlapsWithInstructionInstructionCountBinaryIndex(instruction uint32, count uint32, index uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("overlapsWithInstruction:instructionCount:binaryIndex:"), instruction, count, index)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/selfCostForScope:scopeIdentifier:cost:
func (g GTMioShaderExecutionHistoryNode) SelfCostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("selfCostForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/initWithType:parent:
func (g GTMioShaderExecutionHistoryNode) InitWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryNode {
	rv := objc.Send[GTMioShaderExecutionHistoryNode](g.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/binaries
func (g GTMioShaderExecutionHistoryNode) Binaries() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/children
func (g GTMioShaderExecutionHistoryNode) Children() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("children"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/durationPercentage
func (g GTMioShaderExecutionHistoryNode) DurationPercentage() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("durationPercentage"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/endTimestamps
func (g GTMioShaderExecutionHistoryNode) EndTimestamps() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("endTimestamps"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/parent
func (g GTMioShaderExecutionHistoryNode) Parent() IGTMioShaderExecutionHistoryNode {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("parent"))
	return GTMioShaderExecutionHistoryNodeFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/parentFunction
func (g GTMioShaderExecutionHistoryNode) ParentFunction() IGTMioShaderExecutionHistoryFunctionNode {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("parentFunction"))
	return GTMioShaderExecutionHistoryFunctionNodeFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/root
func (g GTMioShaderExecutionHistoryNode) Root() IGTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("root"))
	return GTMioShaderExecutionHistoryRootNodeFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/startTimestamps
func (g GTMioShaderExecutionHistoryNode) StartTimestamps() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("startTimestamps"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/timestampCount
func (g GTMioShaderExecutionHistoryNode) TimestampCount() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("timestampCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/topCostPercentage
func (g GTMioShaderExecutionHistoryNode) TopCostPercentage() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("topCostPercentage"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/totalDuration
func (g GTMioShaderExecutionHistoryNode) TotalDuration() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("totalDuration"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/type
func (g GTMioShaderExecutionHistoryNode) Type() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("type"))
	return rv
}

// _dfsEnumeratorSync is a synchronous wrapper around [GTMioShaderExecutionHistoryNode._dfsEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderExecutionHistoryNode) _dfsEnumeratorSync(ctx context.Context, _dfs uint32) error {
	done := make(chan struct{}, 1)
	g._dfsEnumerator(_dfs, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DfsSync is a synchronous wrapper around [GTMioShaderExecutionHistoryNode.Dfs].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderExecutionHistoryNode) DfsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.Dfs(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateInstructionsSync is a synchronous wrapper around [GTMioShaderExecutionHistoryNode.EnumerateInstructions].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderExecutionHistoryNode) EnumerateInstructionsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateInstructions(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

