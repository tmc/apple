// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistoryRootNode] class.
var (
	_GTMioShaderExecutionHistoryRootNodeClass     GTMioShaderExecutionHistoryRootNodeClass
	_GTMioShaderExecutionHistoryRootNodeClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryRootNodeClass() GTMioShaderExecutionHistoryRootNodeClass {
	_GTMioShaderExecutionHistoryRootNodeClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryRootNodeClass = GTMioShaderExecutionHistoryRootNodeClass{class: objc.GetClass("GTMioShaderExecutionHistoryRootNode")}
	})
	return _GTMioShaderExecutionHistoryRootNodeClass
}

// GetGTMioShaderExecutionHistoryRootNodeClass returns the class object for GTMioShaderExecutionHistoryRootNode.
func GetGTMioShaderExecutionHistoryRootNodeClass() GTMioShaderExecutionHistoryRootNodeClass {
	return getGTMioShaderExecutionHistoryRootNodeClass()
}

type GTMioShaderExecutionHistoryRootNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryRootNodeClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryRootNodeClass) Alloc() GTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[GTMioShaderExecutionHistoryRootNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderExecutionHistoryRootNode._pushNewFunction]
//   - [GTMioShaderExecutionHistoryRootNode.CacheKey]
//   - [GTMioShaderExecutionHistoryRootNode.CacheObject]
//   - [GTMioShaderExecutionHistoryRootNode.CliqueExecutionHistoryBeginUsc]
//   - [GTMioShaderExecutionHistoryRootNode.CliqueExecutionHistoryEndUsc]
//   - [GTMioShaderExecutionHistoryRootNode.CliqueExecutionHistoryStyle]
//   - [GTMioShaderExecutionHistoryRootNode.Delegate]
//   - [GTMioShaderExecutionHistoryRootNode.SetDelegate]
//   - [GTMioShaderExecutionHistoryRootNode.DumpTree]
//   - [GTMioShaderExecutionHistoryRootNode.EnumerateFunctionCallSites]
//   - [GTMioShaderExecutionHistoryRootNode.FunctionCallSitesForIdentifier]
//   - [GTMioShaderExecutionHistoryRootNode.HandleCachedObject]
//   - [GTMioShaderExecutionHistoryRootNode.LoopBackInstructionEndLoopCountCurrentLoopCountBinary]
//   - [GTMioShaderExecutionHistoryRootNode.NumHitsForInstructionBinaryIndex]
//   - [GTMioShaderExecutionHistoryRootNode.Options]
//   - [GTMioShaderExecutionHistoryRootNode.PopFunctionBinaryRangeBinary]
//   - [GTMioShaderExecutionHistoryRootNode.PopLoopInstructionEndLoopCountBinary]
//   - [GTMioShaderExecutionHistoryRootNode.PopStack]
//   - [GTMioShaderExecutionHistoryRootNode.ProcessInstructionBinaryRangeBinaryNumHits]
//   - [GTMioShaderExecutionHistoryRootNode.PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary]
//   - [GTMioShaderExecutionHistoryRootNode.PushLoopInstructionEndLoopCountBinary]
//   - [GTMioShaderExecutionHistoryRootNode.PushStack]
//   - [GTMioShaderExecutionHistoryRootNode.Style]
//   - [GTMioShaderExecutionHistoryRootNode.TimestampNextInstructionCount]
//   - [GTMioShaderExecutionHistoryRootNode.InitWithStyleOptionsDelegate]
type GTMioShaderExecutionHistoryRootNode struct {
	GTMioShaderExecutionHistoryNode
}

// GTMioShaderExecutionHistoryRootNodeFromID constructs a [GTMioShaderExecutionHistoryRootNode] from an objc.ID.
func GTMioShaderExecutionHistoryRootNodeFromID(id objc.ID) GTMioShaderExecutionHistoryRootNode {
	return GTMioShaderExecutionHistoryRootNode{GTMioShaderExecutionHistoryNode: GTMioShaderExecutionHistoryNodeFromID(id)}
}

// Ensure GTMioShaderExecutionHistoryRootNode implements IGTMioShaderExecutionHistoryRootNode.
var _ IGTMioShaderExecutionHistoryRootNode = GTMioShaderExecutionHistoryRootNode{}

// An interface definition for the [GTMioShaderExecutionHistoryRootNode] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistoryRootNode._pushNewFunction]
//   - [IGTMioShaderExecutionHistoryRootNode.CacheKey]
//   - [IGTMioShaderExecutionHistoryRootNode.CacheObject]
//   - [IGTMioShaderExecutionHistoryRootNode.CliqueExecutionHistoryBeginUsc]
//   - [IGTMioShaderExecutionHistoryRootNode.CliqueExecutionHistoryEndUsc]
//   - [IGTMioShaderExecutionHistoryRootNode.CliqueExecutionHistoryStyle]
//   - [IGTMioShaderExecutionHistoryRootNode.Delegate]
//   - [IGTMioShaderExecutionHistoryRootNode.SetDelegate]
//   - [IGTMioShaderExecutionHistoryRootNode.DumpTree]
//   - [IGTMioShaderExecutionHistoryRootNode.EnumerateFunctionCallSites]
//   - [IGTMioShaderExecutionHistoryRootNode.FunctionCallSitesForIdentifier]
//   - [IGTMioShaderExecutionHistoryRootNode.HandleCachedObject]
//   - [IGTMioShaderExecutionHistoryRootNode.LoopBackInstructionEndLoopCountCurrentLoopCountBinary]
//   - [IGTMioShaderExecutionHistoryRootNode.NumHitsForInstructionBinaryIndex]
//   - [IGTMioShaderExecutionHistoryRootNode.Options]
//   - [IGTMioShaderExecutionHistoryRootNode.PopFunctionBinaryRangeBinary]
//   - [IGTMioShaderExecutionHistoryRootNode.PopLoopInstructionEndLoopCountBinary]
//   - [IGTMioShaderExecutionHistoryRootNode.PopStack]
//   - [IGTMioShaderExecutionHistoryRootNode.ProcessInstructionBinaryRangeBinaryNumHits]
//   - [IGTMioShaderExecutionHistoryRootNode.PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary]
//   - [IGTMioShaderExecutionHistoryRootNode.PushLoopInstructionEndLoopCountBinary]
//   - [IGTMioShaderExecutionHistoryRootNode.PushStack]
//   - [IGTMioShaderExecutionHistoryRootNode.Style]
//   - [IGTMioShaderExecutionHistoryRootNode.TimestampNextInstructionCount]
//   - [IGTMioShaderExecutionHistoryRootNode.InitWithStyleOptionsDelegate]
type IGTMioShaderExecutionHistoryRootNode interface {
	IGTMioShaderExecutionHistoryNode

	// Topic: Methods

	_pushNewFunction(function objectivec.IObject)
	CacheKey() objectivec.IObject
	CacheObject() objectivec.IObject
	CliqueExecutionHistoryBeginUsc(begin *GTMioUSCCliqueMetadata, usc objectivec.IObject)
	CliqueExecutionHistoryEndUsc(end *GTMioUSCCliqueMetadata, usc objectivec.IObject)
	CliqueExecutionHistoryStyle() uint32
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DumpTree(tree objectivec.IObject)
	EnumerateFunctionCallSites(sites VoidHandler)
	FunctionCallSitesForIdentifier(identifier uint64) objectivec.IObject
	HandleCachedObject(object objectivec.IObject) bool
	LoopBackInstructionEndLoopCountCurrentLoopCountBinary(back uint32, end uint32, count uint32, count2 uint32, binary objectivec.IObject)
	NumHitsForInstructionBinaryIndex(instruction uint32, index uint32) uint32
	Options() uint32
	PopFunctionBinaryRangeBinary(function *GTMioShaderBinaryDebugLocation, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject)
	PopLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject)
	PopStack()
	ProcessInstructionBinaryRangeBinaryNumHits(instruction uint32, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject, hits uint32)
	PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary(function *GTMioShaderBinaryDebugLocation, index uint32, inlined bool, binary objectivec.IObject, location *GTMioShaderBinaryDebugLocation, index2 uint32, binary2 objectivec.IObject)
	PushLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject)
	PushStack(stack objectivec.IObject)
	Style() uint32
	TimestampNextInstructionCount(timestamp uint64, next uint64, count uint32)
	InitWithStyleOptionsDelegate(style uint32, options uint32, delegate objectivec.IObject) GTMioShaderExecutionHistoryRootNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryRootNode) Init() GTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[GTMioShaderExecutionHistoryRootNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryRootNode) Autorelease() GTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[GTMioShaderExecutionHistoryRootNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryRootNode creates a new GTMioShaderExecutionHistoryRootNode instance.
func NewGTMioShaderExecutionHistoryRootNode() GTMioShaderExecutionHistoryRootNode {
	class := getGTMioShaderExecutionHistoryRootNodeClass()
	rv := objc.Send[GTMioShaderExecutionHistoryRootNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioShaderExecutionHistoryRootNodeWithStyleOptionsDelegate(style uint32, options uint32, delegate objectivec.IObject) GTMioShaderExecutionHistoryRootNode {
	instance := getGTMioShaderExecutionHistoryRootNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStyle:options:delegate:"), style, options, delegate)
	return GTMioShaderExecutionHistoryRootNodeFromID(rv)
}

func NewGTMioShaderExecutionHistoryRootNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryRootNode {
	instance := getGTMioShaderExecutionHistoryRootNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryRootNodeFromID(rv)
}

func (g GTMioShaderExecutionHistoryRootNode) _pushNewFunction(function objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_pushNewFunction:"), function)
}

// PushNewFunction is an exported wrapper for the private method _pushNewFunction.
func (g GTMioShaderExecutionHistoryRootNode) PushNewFunction(function objectivec.IObject) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_pushNewFunction:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_pushNewFunction:"}
		return err
	}
	g._pushNewFunction(function)
	return nil
}

// CanPushNewFunction reports whether the receiver responds to the private selector _pushNewFunction:.
func (g GTMioShaderExecutionHistoryRootNode) CanPushNewFunction() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_pushNewFunction:"))
}
func (g GTMioShaderExecutionHistoryRootNode) CacheKey() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cacheKey"))
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderExecutionHistoryRootNode) CacheObject() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cacheObject"))
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderExecutionHistoryRootNode) CliqueExecutionHistoryBeginUsc(begin *GTMioUSCCliqueMetadata, usc objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("cliqueExecutionHistoryBegin:usc:"), begin, usc)
}
func (g GTMioShaderExecutionHistoryRootNode) CliqueExecutionHistoryEndUsc(end *GTMioUSCCliqueMetadata, usc objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("cliqueExecutionHistoryEnd:usc:"), end, usc)
}
func (g GTMioShaderExecutionHistoryRootNode) CliqueExecutionHistoryStyle() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("cliqueExecutionHistoryStyle"))
	return rv
}
func (g GTMioShaderExecutionHistoryRootNode) DumpTree(tree objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("dumpTree:"), tree)
}
func (g GTMioShaderExecutionHistoryRootNode) EnumerateFunctionCallSites(sites VoidHandler) {
	_block0, _ := NewVoidBlock(sites)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateFunctionCallSites:"), _block0)
}
func (g GTMioShaderExecutionHistoryRootNode) FunctionCallSitesForIdentifier(identifier uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("functionCallSitesForIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (g GTMioShaderExecutionHistoryRootNode) HandleCachedObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("handleCachedObject:"), object)
	return rv
}
func (g GTMioShaderExecutionHistoryRootNode) LoopBackInstructionEndLoopCountCurrentLoopCountBinary(back uint32, end uint32, count uint32, count2 uint32, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("loopBack:instructionEnd:loopCount:currentLoopCount:binary:"), back, end, count, count2, binary)
}
func (g GTMioShaderExecutionHistoryRootNode) NumHitsForInstructionBinaryIndex(instruction uint32, index uint32) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("numHitsForInstruction:binaryIndex:"), instruction, index)
	return rv
}
func (g GTMioShaderExecutionHistoryRootNode) PopFunctionBinaryRangeBinary(function *GTMioShaderBinaryDebugLocation, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("popFunction:binaryRange:binary:"), function, range_, binary)
}
func (g GTMioShaderExecutionHistoryRootNode) PopLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("popLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}
func (g GTMioShaderExecutionHistoryRootNode) PopStack() {
	objc.Send[objc.ID](g.ID, objc.Sel("popStack"))
}
func (g GTMioShaderExecutionHistoryRootNode) ProcessInstructionBinaryRangeBinaryNumHits(instruction uint32, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject, hits uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("processInstruction:binaryRange:binary:numHits:"), instruction, range_, binary, hits)
}
func (g GTMioShaderExecutionHistoryRootNode) PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary(function *GTMioShaderBinaryDebugLocation, index uint32, inlined bool, binary objectivec.IObject, location *GTMioShaderBinaryDebugLocation, index2 uint32, binary2 objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("pushFunction:binaryRangeIndex:inlined:binary:callerLocation:callerBinaryRangeIndex:callerBinary:"), function, index, inlined, binary, location, index2, binary2)
}
func (g GTMioShaderExecutionHistoryRootNode) PushLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("pushLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}
func (g GTMioShaderExecutionHistoryRootNode) PushStack(stack objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("pushStack:"), stack)
}
func (g GTMioShaderExecutionHistoryRootNode) TimestampNextInstructionCount(timestamp uint64, next uint64, count uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("timestamp:next:instructionCount:"), timestamp, next, count)
}
func (g GTMioShaderExecutionHistoryRootNode) InitWithStyleOptionsDelegate(style uint32, options uint32, delegate objectivec.IObject) GTMioShaderExecutionHistoryRootNode {
	rv := objc.Send[GTMioShaderExecutionHistoryRootNode](g.ID, objc.Sel("initWithStyle:options:delegate:"), style, options, delegate)
	return rv
}

func (g GTMioShaderExecutionHistoryRootNode) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("delegate"))
	return rv
}
func (g GTMioShaderExecutionHistoryRootNode) SetDelegate(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setDelegate:"), value)
}
func (g GTMioShaderExecutionHistoryRootNode) Options() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("options"))
	return rv
}
func (g GTMioShaderExecutionHistoryRootNode) Style() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("style"))
	return rv
}

// EnumerateFunctionCallSitesSync is a synchronous wrapper around [GTMioShaderExecutionHistoryRootNode.EnumerateFunctionCallSites].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioShaderExecutionHistoryRootNode) EnumerateFunctionCallSitesSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateFunctionCallSites(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
