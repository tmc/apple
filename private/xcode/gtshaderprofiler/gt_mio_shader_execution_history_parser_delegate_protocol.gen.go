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
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate
type GTMioShaderExecutionHistoryParserDelegate interface {
	objectivec.IObject

	// CliqueExecutionHistoryStyle protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/cliqueExecutionHistoryStyle
	CliqueExecutionHistoryStyle() uint32

	// TimestampNextInstructionCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/timestamp:next:instructionCount:
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

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/cacheKey
func (o GTMioShaderExecutionHistoryParserDelegateObject) CacheKey() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("cacheKey"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/cacheObject
func (o GTMioShaderExecutionHistoryParserDelegateObject) CacheObject() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("cacheObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/cliqueExecutionHistoryBegin:usc:
func (o GTMioShaderExecutionHistoryParserDelegateObject) CliqueExecutionHistoryBeginUsc(begin *GTMioUSCCliqueMetadataRef, usc objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("cliqueExecutionHistoryBegin:usc:"), begin, usc)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/cliqueExecutionHistoryEnd:usc:
func (o GTMioShaderExecutionHistoryParserDelegateObject) CliqueExecutionHistoryEndUsc(end *GTMioUSCCliqueMetadataRef, usc objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("cliqueExecutionHistoryEnd:usc:"), end, usc)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/cliqueExecutionHistoryStyle
func (o GTMioShaderExecutionHistoryParserDelegateObject) CliqueExecutionHistoryStyle() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("cliqueExecutionHistoryStyle"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/handleCachedObject:
func (o GTMioShaderExecutionHistoryParserDelegateObject) HandleCachedObject(object objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("handleCachedObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/loopBack:instructionEnd:loopCount:currentLoopCount:binary:
func (o GTMioShaderExecutionHistoryParserDelegateObject) LoopBackInstructionEndLoopCountCurrentLoopCountBinary(back uint32, end uint32, count uint32, count2 uint32, binary objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("loopBack:instructionEnd:loopCount:currentLoopCount:binary:"), back, end, count, count2, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/popFunction:binaryRange:binary:
func (o GTMioShaderExecutionHistoryParserDelegateObject) PopFunctionBinaryRangeBinary(function unsafe.Pointer, range_ unsafe.Pointer, binary objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("popFunction:binaryRange:binary:"), function, range_, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/popLoop:instructionEnd:loopCount:binary:
func (o GTMioShaderExecutionHistoryParserDelegateObject) PopLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("popLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/processInstruction:binaryRange:binary:numHits:
func (o GTMioShaderExecutionHistoryParserDelegateObject) ProcessInstructionBinaryRangeBinaryNumHits(instruction uint32, range_ unsafe.Pointer, binary objectivec.IObject, hits uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("processInstruction:binaryRange:binary:numHits:"), instruction, range_, binary, hits)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/pushFunction:binaryRangeIndex:inlined:binary:callerLocation:callerBinaryRangeIndex:callerBinary:
func (o GTMioShaderExecutionHistoryParserDelegateObject) PushFunctionBinaryRangeIndexInlinedBinaryCallerLocationCallerBinaryRangeIndexCallerBinary(function unsafe.Pointer, index uint32, inlined bool, binary objectivec.IObject, location unsafe.Pointer, index2 uint32, binary2 objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pushFunction:binaryRangeIndex:inlined:binary:callerLocation:callerBinaryRangeIndex:callerBinary:"), function, index, inlined, binary, location, index2, binary2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/pushLoop:instructionEnd:loopCount:binary:
func (o GTMioShaderExecutionHistoryParserDelegateObject) PushLoopInstructionEndLoopCountBinary(loop uint32, end uint32, count uint32, binary objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pushLoop:instructionEnd:loopCount:binary:"), loop, end, count, binary)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryParserDelegate/timestamp:next:instructionCount:
func (o GTMioShaderExecutionHistoryParserDelegateObject) TimestampNextInstructionCount(timestamp uint64, next uint64, count uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("timestamp:next:instructionCount:"), timestamp, next, count)
}

// GTMioShaderExecutionHistoryParserDelegateConfig holds optional typed callbacks for [GTMioShaderExecutionHistoryParserDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/gtshaderprofiler/gtmioshaderexecutionhistoryparserdelegate
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
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/gtshaderprofiler/gtmioshaderexecutionhistoryparserdelegate
func NewGTMioShaderExecutionHistoryParserDelegate(config GTMioShaderExecutionHistoryParserDelegateConfig) GTMioShaderExecutionHistoryParserDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoGTMioShaderExecutionHistoryParserDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CliqueExecutionHistoryStyle != nil {
		fn := config.CliqueExecutionHistoryStyle
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("cliqueExecutionHistoryStyle"),
			Fn: func(self objc.ID, _cmd objc.SEL) uint32 {
				return fn()
			},
		})
	}

	if config.TimestampNextInstructionCount != nil {
		fn := config.TimestampNextInstructionCount
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("timestamp:next:instructionCount:"),
			Fn: func(self objc.ID, _cmd objc.SEL, timestamp uint64, next uint64, count uint32) {
				fn(timestamp, next, count)
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
