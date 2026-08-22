// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// GTMioShaderExecutionHistoryParserDelegate protocol.
type GTMioShaderExecutionHistoryParserDelegate interface {
	objectivec.IObject

	// CacheKey protocol.
	CacheKey() objectivec.IObject

	// CacheObject protocol.
	CacheObject() objectivec.IObject

	// CliqueExecutionHistoryBeginUsc protocol.
	CliqueExecutionHistoryBeginUsc(begin *GTMioUSCCliqueMetadata, usc objectivec.IObject)

	// CliqueExecutionHistoryEndUsc protocol.
	CliqueExecutionHistoryEndUsc(end *GTMioUSCCliqueMetadata, usc objectivec.IObject)

	// CliqueExecutionHistoryStyle protocol.
	CliqueExecutionHistoryStyle() uint32

	// HandleCachedObject protocol.
	HandleCachedObject(object objectivec.IObject) bool

	// LoopBackInstructionEndLoopCountCurrentLoopCountBinary protocol.
	LoopBackInstructionEndLoopCountCurrentLoopCountBinary(back uint32, end uint32, count uint32, count2 uint32, binary objectivec.IObject)

	// PopFunctionBinaryRangeBinary protocol.
	PopFunctionBinaryRangeBinary(function *GTMioShaderBinaryDebugLocation, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject)

	// PopLoopInstructionEndLoopCountBinary protocol.
	PopLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject)

	// ProcessInstructionBinaryRangeBinaryNumHits protocol.
	ProcessInstructionBinaryRangeBinaryNumHits(instruction uint32, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject, hits uint32)

	// PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary protocol.
	PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary(function *GTMioShaderBinaryDebugLocation, index uint32, inlined bool, binary objectivec.IObject, location *GTMioShaderBinaryDebugLocation, index2 uint32, binary2 objectivec.IObject)

	// PushLoopInstructionEndLoopCountBinary protocol.
	PushLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject)

	// TimestampNextInstructionCount protocol.
	TimestampNextInstructionCount(timestamp uint64, next uint64, count uint32)
}

// GTMioShaderExecutionHistoryParserDelegateObject wraps an existing Objective-C object that conforms to the GTMioShaderExecutionHistoryParserDelegate protocol.
type GTMioShaderExecutionHistoryParserDelegateObject struct {
	objectivec.Object
}

func (o GTMioShaderExecutionHistoryParserDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTMioShaderExecutionHistoryParserDelegateObjectFromID constructs a [GTMioShaderExecutionHistoryParserDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTMioShaderExecutionHistoryParserDelegateObjectFromID(id objc.ID) GTMioShaderExecutionHistoryParserDelegateObject {
	return GTMioShaderExecutionHistoryParserDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTMioShaderExecutionHistoryParserDelegateObject) CacheKey() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("cacheKey"))
	return objectivec.Object{ID: rv}
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) CacheObject() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("cacheObject"))
	return objectivec.Object{ID: rv}
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) CliqueExecutionHistoryBeginUsc(begin *GTMioUSCCliqueMetadata, usc objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("cliqueExecutionHistoryBegin:usc:"), unsafe.Pointer(begin), usc)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) CliqueExecutionHistoryEndUsc(end *GTMioUSCCliqueMetadata, usc objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("cliqueExecutionHistoryEnd:usc:"), unsafe.Pointer(end), usc)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) CliqueExecutionHistoryStyle() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("cliqueExecutionHistoryStyle"))
	return rv
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) HandleCachedObject(object objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("handleCachedObject:"), object)
	return rv
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) LoopBackInstructionEndLoopCountCurrentLoopCountBinary(back uint32, end uint32, count uint32, count2 uint32, binary objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("loopBack:instructionEnd:loopCount:currentLoopCount:binary:"), back, end, count, count2, binary)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) PopFunctionBinaryRangeBinary(function *GTMioShaderBinaryDebugLocation, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("popFunction:binaryRange:binary:"), unsafe.Pointer(function), unsafe.Pointer(range_), binary)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) PopLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("popLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) ProcessInstructionBinaryRangeBinaryNumHits(instruction uint32, range_ *GTMioShaderBinaryDebugBinaryRange, binary objectivec.IObject, hits uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("processInstruction:binaryRange:binary:numHits:"), instruction, unsafe.Pointer(range_), binary, hits)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary(function *GTMioShaderBinaryDebugLocation, index uint32, inlined bool, binary objectivec.IObject, location *GTMioShaderBinaryDebugLocation, index2 uint32, binary2 objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("pushFunction:binaryRangeIndex:inlined:binary:callerLocation:callerBinaryRangeIndex:callerBinary:"), unsafe.Pointer(function), index, inlined, binary, unsafe.Pointer(location), index2, binary2)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) PushLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("pushLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}
func (o GTMioShaderExecutionHistoryParserDelegateObject) TimestampNextInstructionCount(timestamp uint64, next uint64, count uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("timestamp:next:instructionCount:"), timestamp, next, count)
}

// GTMioShaderExecutionHistoryParserDelegateConfig holds optional typed callbacks for [GTMioShaderExecutionHistoryParserDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
type GTMioShaderExecutionHistoryParserDelegateConfig struct {

	// Other Methods
	CliqueExecutionHistoryStyle   func() uint32
	TimestampNextInstructionCount func(timestamp uint64, next uint64, count uint32)
}

// NewGTMioShaderExecutionHistoryParserDelegate creates an Objective-C object implementing the [GTMioShaderExecutionHistoryParserDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [GTMioShaderExecutionHistoryParserDelegateObject] satisfies the [GTMioShaderExecutionHistoryParserDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
func NewGTMioShaderExecutionHistoryParserDelegate(config GTMioShaderExecutionHistoryParserDelegateConfig) GTMioShaderExecutionHistoryParserDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoGTMioShaderExecutionHistoryParserDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CliqueExecutionHistoryStyle != nil {
		fn := config.CliqueExecutionHistoryStyle
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cliqueExecutionHistoryStyle"),
			Fn: func(self objc.ID, _cmd objc.SEL) uint32 {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("GTMioShaderExecutionHistoryParserDelegate", "cliqueExecutionHistoryStyle")
					}
				}()
				_delegateResult := fn()
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.TimestampNextInstructionCount != nil {
		fn := config.TimestampNextInstructionCount
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("timestamp:next:instructionCount:"),
			Fn: func(self objc.ID, _cmd objc.SEL, timestamp uint64, next uint64, count uint32) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("GTMioShaderExecutionHistoryParserDelegate", "timestamp:next:instructionCount:")
					}
				}()
				fn(timestamp, next, count)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("GTMioShaderExecutionHistoryParserDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewGTMioShaderExecutionHistoryParserDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return GTMioShaderExecutionHistoryParserDelegateObjectFromID(instance)
}
