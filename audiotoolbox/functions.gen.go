// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremidi"
	"github.com/tmc/apple/coreservices"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/os"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("AudioToolbox: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("AudioToolbox: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("AudioToolbox: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("AudioToolbox: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _aUEventListenerAddEventType func(inListener AUEventListenerRef, inObject unsafe.Pointer, inEvent *AudioUnitEvent) int32
var _aUEventListenerAddEventTypeErr error

func tryAUEventListenerAddEventType(inListener AUEventListenerRef, inObject unsafe.Pointer, inEvent *AudioUnitEvent) (int32, error) {
	if _aUEventListenerAddEventType == nil {
		return 0, symbolCallError("AUEventListenerAddEventType", "10.3", _aUEventListenerAddEventTypeErr)
	}
	return _aUEventListenerAddEventType(inListener, inObject, inEvent), nil
}

// AUEventListenerAddEventType.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerAddEventType(_:_:_:)
func AUEventListenerAddEventType(inListener AUEventListenerRef, inObject unsafe.Pointer, inEvent *AudioUnitEvent) int32 {
	result, callErr := tryAUEventListenerAddEventType(inListener, inObject, inEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUEventListenerCreate func(inProc AUEventListenerProc, inUserData unsafe.Pointer, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inNotificationInterval float32, inValueChangeGranularity float32, outListener *AUEventListenerRef) int32
var _aUEventListenerCreateErr error

func tryAUEventListenerCreate(inProc AUEventListenerProc, inUserData unsafe.Pointer, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inNotificationInterval float32, inValueChangeGranularity float32, outListener *AUEventListenerRef) (int32, error) {
	if _aUEventListenerCreate == nil {
		return 0, symbolCallError("AUEventListenerCreate", "10.3", _aUEventListenerCreateErr)
	}
	return _aUEventListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, inValueChangeGranularity, outListener), nil
}

// AUEventListenerCreate.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerCreate(_:_:_:_:_:_:_:)
func AUEventListenerCreate(inProc AUEventListenerProc, inUserData unsafe.Pointer, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inNotificationInterval float32, inValueChangeGranularity float32, outListener *AUEventListenerRef) int32 {
	result, callErr := tryAUEventListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, inValueChangeGranularity, outListener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUEventListenerCreateWithDispatchQueue func(outListener *AUEventListenerRef, inNotificationInterval float32, inValueChangeGranularity float32, inDispatchQueue uintptr, inBlock unsafe.Pointer) int32
var _aUEventListenerCreateWithDispatchQueueErr error

func tryAUEventListenerCreateWithDispatchQueue(outListener *AUEventListenerRef, inNotificationInterval float32, inValueChangeGranularity float32, inDispatchQueue dispatch.Queue, inBlock AUEventListenerBlock) (int32, error) {
	if _aUEventListenerCreateWithDispatchQueue == nil {
		return 0, symbolCallError("AUEventListenerCreateWithDispatchQueue", "10.6", _aUEventListenerCreateWithDispatchQueueErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 unsafe.Pointer, blockArg1 *AudioUnitEvent, blockArg2 uint64, blockArg3 float32) {
		inBlock(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _aUEventListenerCreateWithDispatchQueue(outListener, inNotificationInterval, inValueChangeGranularity, uintptr(inDispatchQueue.Handle()), _block0), nil
}

// AUEventListenerCreateWithDispatchQueue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerCreateWithDispatchQueue(_:_:_:_:_:)
func AUEventListenerCreateWithDispatchQueue(outListener *AUEventListenerRef, inNotificationInterval float32, inValueChangeGranularity float32, inDispatchQueue dispatch.Queue, inBlock AUEventListenerBlock) int32 {
	result, callErr := tryAUEventListenerCreateWithDispatchQueue(outListener, inNotificationInterval, inValueChangeGranularity, inDispatchQueue, inBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUEventListenerNotify func(inSendingListener AUEventListenerRef, inSendingObject unsafe.Pointer, inEvent *AudioUnitEvent) int32
var _aUEventListenerNotifyErr error

func tryAUEventListenerNotify(inSendingListener AUEventListenerRef, inSendingObject unsafe.Pointer, inEvent *AudioUnitEvent) (int32, error) {
	if _aUEventListenerNotify == nil {
		return 0, symbolCallError("AUEventListenerNotify", "10.3", _aUEventListenerNotifyErr)
	}
	return _aUEventListenerNotify(inSendingListener, inSendingObject, inEvent), nil
}

// AUEventListenerNotify.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerNotify(_:_:_:)
func AUEventListenerNotify(inSendingListener AUEventListenerRef, inSendingObject unsafe.Pointer, inEvent *AudioUnitEvent) int32 {
	result, callErr := tryAUEventListenerNotify(inSendingListener, inSendingObject, inEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUEventListenerRemoveEventType func(inListener AUEventListenerRef, inObject unsafe.Pointer, inEvent *AudioUnitEvent) int32
var _aUEventListenerRemoveEventTypeErr error

func tryAUEventListenerRemoveEventType(inListener AUEventListenerRef, inObject unsafe.Pointer, inEvent *AudioUnitEvent) (int32, error) {
	if _aUEventListenerRemoveEventType == nil {
		return 0, symbolCallError("AUEventListenerRemoveEventType", "10.3", _aUEventListenerRemoveEventTypeErr)
	}
	return _aUEventListenerRemoveEventType(inListener, inObject, inEvent), nil
}

// AUEventListenerRemoveEventType.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUEventListenerRemoveEventType(_:_:_:)
func AUEventListenerRemoveEventType(inListener AUEventListenerRef, inObject unsafe.Pointer, inEvent *AudioUnitEvent) int32 {
	result, callErr := tryAUEventListenerRemoveEventType(inListener, inObject, inEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphAddNode func(inGraph AUGraph, inDescription *AudioComponentDescription, outNode *AUNode) int32
var _aUGraphAddNodeErr error

func tryAUGraphAddNode(inGraph AUGraph, inDescription *AudioComponentDescription, outNode *AUNode) (int32, error) {
	if _aUGraphAddNode == nil {
		return 0, symbolCallError("AUGraphAddNode", "10.5", _aUGraphAddNodeErr)
	}
	return _aUGraphAddNode(inGraph, inDescription, outNode), nil
}

// AUGraphAddNode adds a node to an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphAddNode(_:_:_:)
func AUGraphAddNode(inGraph AUGraph, inDescription *AudioComponentDescription, outNode *AUNode) int32 {
	result, callErr := tryAUGraphAddNode(inGraph, inDescription, outNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphAddRenderNotify func(inGraph AUGraph, inCallback AURenderCallback, inRefCon unsafe.Pointer) int32
var _aUGraphAddRenderNotifyErr error

func tryAUGraphAddRenderNotify(inGraph AUGraph, inCallback AURenderCallback, inRefCon unsafe.Pointer) (int32, error) {
	if _aUGraphAddRenderNotify == nil {
		return 0, symbolCallError("AUGraphAddRenderNotify", "10.2", _aUGraphAddRenderNotifyErr)
	}
	return _aUGraphAddRenderNotify(inGraph, inCallback, inRefCon), nil
}

// AUGraphAddRenderNotify adds a render notification callback to an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphAddRenderNotify(_:_:_:)
func AUGraphAddRenderNotify(inGraph AUGraph, inCallback AURenderCallback, inRefCon unsafe.Pointer) int32 {
	result, callErr := tryAUGraphAddRenderNotify(inGraph, inCallback, inRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphClearConnections func(inGraph AUGraph) int32
var _aUGraphClearConnectionsErr error

func tryAUGraphClearConnections(inGraph AUGraph) (int32, error) {
	if _aUGraphClearConnections == nil {
		return 0, symbolCallError("AUGraphClearConnections", "10.0", _aUGraphClearConnectionsErr)
	}
	return _aUGraphClearConnections(inGraph), nil
}

// AUGraphClearConnections clears all of the interactions in an audio unit processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphClearConnections(_:)
func AUGraphClearConnections(inGraph AUGraph) int32 {
	result, callErr := tryAUGraphClearConnections(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphClose func(inGraph AUGraph) int32
var _aUGraphCloseErr error

func tryAUGraphClose(inGraph AUGraph) (int32, error) {
	if _aUGraphClose == nil {
		return 0, symbolCallError("AUGraphClose", "10.0", _aUGraphCloseErr)
	}
	return _aUGraphClose(inGraph), nil
}

// AUGraphClose closes an audio unit processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphClose(_:)
func AUGraphClose(inGraph AUGraph) int32 {
	result, callErr := tryAUGraphClose(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphConnectNodeInput func(inGraph AUGraph, inSourceNode AUNode, inSourceOutputNumber uint32, inDestNode AUNode, inDestInputNumber uint32) int32
var _aUGraphConnectNodeInputErr error

func tryAUGraphConnectNodeInput(inGraph AUGraph, inSourceNode AUNode, inSourceOutputNumber uint32, inDestNode AUNode, inDestInputNumber uint32) (int32, error) {
	if _aUGraphConnectNodeInput == nil {
		return 0, symbolCallError("AUGraphConnectNodeInput", "10.0", _aUGraphConnectNodeInputErr)
	}
	return _aUGraphConnectNodeInput(inGraph, inSourceNode, inSourceOutputNumber, inDestNode, inDestInputNumber), nil
}

// AUGraphConnectNodeInput connects one node’s output to another node’s input.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphConnectNodeInput(_:_:_:_:_:)
func AUGraphConnectNodeInput(inGraph AUGraph, inSourceNode AUNode, inSourceOutputNumber uint32, inDestNode AUNode, inDestInputNumber uint32) int32 {
	result, callErr := tryAUGraphConnectNodeInput(inGraph, inSourceNode, inSourceOutputNumber, inDestNode, inDestInputNumber)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphCountNodeConnections func(inGraph AUGraph, inNode AUNode, outNumConnections *uint32) int32
var _aUGraphCountNodeConnectionsErr error

func tryAUGraphCountNodeConnections(inGraph AUGraph, inNode AUNode, outNumConnections *uint32) (int32, error) {
	if _aUGraphCountNodeConnections == nil {
		return 0, symbolCallError("AUGraphCountNodeConnections", "10.3", _aUGraphCountNodeConnectionsErr)
	}
	return _aUGraphCountNodeConnections(inGraph, inNode, outNumConnections), nil
}

// AUGraphCountNodeConnections deprecated in OS X v10.5.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphCountNodeConnections
func AUGraphCountNodeConnections(inGraph AUGraph, inNode AUNode, outNumConnections *uint32) int32 {
	result, callErr := tryAUGraphCountNodeConnections(inGraph, inNode, outNumConnections)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphCountNodeInteractions func(inGraph AUGraph, inNode AUNode, outNumInteractions *uint32) int32
var _aUGraphCountNodeInteractionsErr error

func tryAUGraphCountNodeInteractions(inGraph AUGraph, inNode AUNode, outNumInteractions *uint32) (int32, error) {
	if _aUGraphCountNodeInteractions == nil {
		return 0, symbolCallError("AUGraphCountNodeInteractions", "10.5", _aUGraphCountNodeInteractionsErr)
	}
	return _aUGraphCountNodeInteractions(inGraph, inNode, outNumInteractions), nil
}

// AUGraphCountNodeInteractions retrieves the number of interactions of an audio processing graph’s node.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphCountNodeInteractions(_:_:_:)
func AUGraphCountNodeInteractions(inGraph AUGraph, inNode AUNode, outNumInteractions *uint32) int32 {
	result, callErr := tryAUGraphCountNodeInteractions(inGraph, inNode, outNumInteractions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphDisconnectNodeInput func(inGraph AUGraph, inDestNode AUNode, inDestInputNumber uint32) int32
var _aUGraphDisconnectNodeInputErr error

func tryAUGraphDisconnectNodeInput(inGraph AUGraph, inDestNode AUNode, inDestInputNumber uint32) (int32, error) {
	if _aUGraphDisconnectNodeInput == nil {
		return 0, symbolCallError("AUGraphDisconnectNodeInput", "10.0", _aUGraphDisconnectNodeInputErr)
	}
	return _aUGraphDisconnectNodeInput(inGraph, inDestNode, inDestInputNumber), nil
}

// AUGraphDisconnectNodeInput disconnects a node’s input.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphDisconnectNodeInput(_:_:_:)
func AUGraphDisconnectNodeInput(inGraph AUGraph, inDestNode AUNode, inDestInputNumber uint32) int32 {
	result, callErr := tryAUGraphDisconnectNodeInput(inGraph, inDestNode, inDestInputNumber)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetCPULoad func(inGraph AUGraph, outAverageCPULoad *float32) int32
var _aUGraphGetCPULoadErr error

func tryAUGraphGetCPULoad(inGraph AUGraph, outAverageCPULoad *float32) (int32, error) {
	if _aUGraphGetCPULoad == nil {
		return 0, symbolCallError("AUGraphGetCPULoad", "10.1", _aUGraphGetCPULoadErr)
	}
	return _aUGraphGetCPULoad(inGraph, outAverageCPULoad), nil
}

// AUGraphGetCPULoad obtains the short-term running average of the current CPU load of the audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetCPULoad(_:_:)
func AUGraphGetCPULoad(inGraph AUGraph, outAverageCPULoad *float32) int32 {
	result, callErr := tryAUGraphGetCPULoad(inGraph, outAverageCPULoad)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetConnectionInfo func(inGraph AUGraph, inConnectionIndex uint32, outSourceNode *AUNode, outSourceOutputNumber *uint32, outDestNode *AUNode, outDestInputNumber *uint32) int32
var _aUGraphGetConnectionInfoErr error

func tryAUGraphGetConnectionInfo(inGraph AUGraph, inConnectionIndex uint32, outSourceNode *AUNode, outSourceOutputNumber *uint32, outDestNode *AUNode, outDestInputNumber *uint32) (int32, error) {
	if _aUGraphGetConnectionInfo == nil {
		return 0, symbolCallError("AUGraphGetConnectionInfo", "10.1", _aUGraphGetConnectionInfoErr)
	}
	return _aUGraphGetConnectionInfo(inGraph, inConnectionIndex, outSourceNode, outSourceOutputNumber, outDestNode, outDestInputNumber), nil
}

// AUGraphGetConnectionInfo deprecated in OS X v10.5.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetConnectionInfo
func AUGraphGetConnectionInfo(inGraph AUGraph, inConnectionIndex uint32, outSourceNode *AUNode, outSourceOutputNumber *uint32, outDestNode *AUNode, outDestInputNumber *uint32) int32 {
	result, callErr := tryAUGraphGetConnectionInfo(inGraph, inConnectionIndex, outSourceNode, outSourceOutputNumber, outDestNode, outDestInputNumber)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetIndNode func(inGraph AUGraph, inIndex uint32, outNode *AUNode) int32
var _aUGraphGetIndNodeErr error

func tryAUGraphGetIndNode(inGraph AUGraph, inIndex uint32, outNode *AUNode) (int32, error) {
	if _aUGraphGetIndNode == nil {
		return 0, symbolCallError("AUGraphGetIndNode", "10.0", _aUGraphGetIndNodeErr)
	}
	return _aUGraphGetIndNode(inGraph, inIndex, outNode), nil
}

// AUGraphGetIndNode gets the audio processing graph node at a given index.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetIndNode(_:_:_:)
func AUGraphGetIndNode(inGraph AUGraph, inIndex uint32, outNode *AUNode) int32 {
	result, callErr := tryAUGraphGetIndNode(inGraph, inIndex, outNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetInteractionInfo func(inGraph AUGraph, inInteractionIndex uint32, outInteraction *AUNodeInteraction) int32
var _aUGraphGetInteractionInfoErr error

func tryAUGraphGetInteractionInfo(inGraph AUGraph, inInteractionIndex uint32, outInteraction *AUNodeInteraction) (int32, error) {
	if _aUGraphGetInteractionInfo == nil {
		return 0, symbolCallError("AUGraphGetInteractionInfo", "10.5", _aUGraphGetInteractionInfoErr)
	}
	return _aUGraphGetInteractionInfo(inGraph, inInteractionIndex, outInteraction), nil
}

// AUGraphGetInteractionInfo retrieves information about a particular interaction in an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetInteractionInfo(_:_:_:)
func AUGraphGetInteractionInfo(inGraph AUGraph, inInteractionIndex uint32, outInteraction *AUNodeInteraction) int32 {
	result, callErr := tryAUGraphGetInteractionInfo(inGraph, inInteractionIndex, outInteraction)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetMaxCPULoad func(inGraph AUGraph, outMaxLoad *float32) int32
var _aUGraphGetMaxCPULoadErr error

func tryAUGraphGetMaxCPULoad(inGraph AUGraph, outMaxLoad *float32) (int32, error) {
	if _aUGraphGetMaxCPULoad == nil {
		return 0, symbolCallError("AUGraphGetMaxCPULoad", "10.3", _aUGraphGetMaxCPULoadErr)
	}
	return _aUGraphGetMaxCPULoad(inGraph, outMaxLoad), nil
}

// AUGraphGetMaxCPULoad obtains the maximum CPU load of an audio processing graph since this call was last made or since the graph was last started.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetMaxCPULoad(_:_:)
func AUGraphGetMaxCPULoad(inGraph AUGraph, outMaxLoad *float32) int32 {
	result, callErr := tryAUGraphGetMaxCPULoad(inGraph, outMaxLoad)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetNodeConnections func(inGraph AUGraph, inNode AUNode, outConnections *AudioUnitNodeConnection, ioNumConnections *uint32) int32
var _aUGraphGetNodeConnectionsErr error

func tryAUGraphGetNodeConnections(inGraph AUGraph, inNode AUNode, outConnections *AudioUnitNodeConnection, ioNumConnections *uint32) (int32, error) {
	if _aUGraphGetNodeConnections == nil {
		return 0, symbolCallError("AUGraphGetNodeConnections", "10.3", _aUGraphGetNodeConnectionsErr)
	}
	return _aUGraphGetNodeConnections(inGraph, inNode, outConnections, ioNumConnections), nil
}

// AUGraphGetNodeConnections deprecated in OS X v10.5.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeConnections
func AUGraphGetNodeConnections(inGraph AUGraph, inNode AUNode, outConnections *AudioUnitNodeConnection, ioNumConnections *uint32) int32 {
	result, callErr := tryAUGraphGetNodeConnections(inGraph, inNode, outConnections, ioNumConnections)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetNodeCount func(inGraph AUGraph, outNumberOfNodes *uint32) int32
var _aUGraphGetNodeCountErr error

func tryAUGraphGetNodeCount(inGraph AUGraph, outNumberOfNodes *uint32) (int32, error) {
	if _aUGraphGetNodeCount == nil {
		return 0, symbolCallError("AUGraphGetNodeCount", "10.0", _aUGraphGetNodeCountErr)
	}
	return _aUGraphGetNodeCount(inGraph, outNumberOfNodes), nil
}

// AUGraphGetNodeCount the number of nodes in an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeCount(_:_:)
func AUGraphGetNodeCount(inGraph AUGraph, outNumberOfNodes *uint32) int32 {
	result, callErr := tryAUGraphGetNodeCount(inGraph, outNumberOfNodes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetNodeInfo func(inGraph AUGraph, inNode AUNode, outDescription *coreservices.ComponentDescription, outClassDataSize *uint32, outClassData unsafe.Pointer, outAudioUnit *AudioUnit) int32
var _aUGraphGetNodeInfoErr error

func tryAUGraphGetNodeInfo(inGraph AUGraph, inNode AUNode, outDescription *coreservices.ComponentDescription, outClassDataSize *uint32, outClassData unsafe.Pointer, outAudioUnit *AudioUnit) (int32, error) {
	if _aUGraphGetNodeInfo == nil {
		return 0, symbolCallError("AUGraphGetNodeInfo", "10.0", _aUGraphGetNodeInfoErr)
	}
	return _aUGraphGetNodeInfo(inGraph, inNode, outDescription, outClassDataSize, outClassData, outAudioUnit), nil
}

// AUGraphGetNodeInfo deprecated in OS X v10.5.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeInfo
func AUGraphGetNodeInfo(inGraph AUGraph, inNode AUNode, outDescription *coreservices.ComponentDescription, outClassDataSize *uint32, outClassData unsafe.Pointer, outAudioUnit *AudioUnit) int32 {
	result, callErr := tryAUGraphGetNodeInfo(inGraph, inNode, outDescription, outClassDataSize, outClassData, outAudioUnit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetNodeInfoSubGraph func(inGraph AUGraph, inNode AUNode, outSubGraph *AUGraph) int32
var _aUGraphGetNodeInfoSubGraphErr error

func tryAUGraphGetNodeInfoSubGraph(inGraph AUGraph, inNode AUNode, outSubGraph *AUGraph) (int32, error) {
	if _aUGraphGetNodeInfoSubGraph == nil {
		return 0, symbolCallError("AUGraphGetNodeInfoSubGraph", "10.2", _aUGraphGetNodeInfoSubGraphErr)
	}
	return _aUGraphGetNodeInfoSubGraph(inGraph, inNode, outSubGraph), nil
}

// AUGraphGetNodeInfoSubGraph gets the audio processing subgraph object represented by a node.
//
// Deprecated: Deprecated since macOS 27.0. no longer supported
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeInfoSubGraph(_:_:_:)
func AUGraphGetNodeInfoSubGraph(inGraph AUGraph, inNode AUNode, outSubGraph *AUGraph) int32 {
	result, callErr := tryAUGraphGetNodeInfoSubGraph(inGraph, inNode, outSubGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetNodeInteractions func(inGraph AUGraph, inNode AUNode, ioNumInteractions *uint32, outInteractions *AUNodeInteraction) int32
var _aUGraphGetNodeInteractionsErr error

func tryAUGraphGetNodeInteractions(inGraph AUGraph, inNode AUNode, ioNumInteractions *uint32, outInteractions *AUNodeInteraction) (int32, error) {
	if _aUGraphGetNodeInteractions == nil {
		return 0, symbolCallError("AUGraphGetNodeInteractions", "10.5", _aUGraphGetNodeInteractionsErr)
	}
	return _aUGraphGetNodeInteractions(inGraph, inNode, ioNumInteractions, outInteractions), nil
}

// AUGraphGetNodeInteractions retrieves information about the interactions in an audio processing graph for a given node.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNodeInteractions(_:_:_:_:)
func AUGraphGetNodeInteractions(inGraph AUGraph, inNode AUNode, ioNumInteractions *uint32, outInteractions *AUNodeInteraction) int32 {
	result, callErr := tryAUGraphGetNodeInteractions(inGraph, inNode, ioNumInteractions, outInteractions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetNumberOfConnections func(inGraph AUGraph, outNumConnections *uint32) int32
var _aUGraphGetNumberOfConnectionsErr error

func tryAUGraphGetNumberOfConnections(inGraph AUGraph, outNumConnections *uint32) (int32, error) {
	if _aUGraphGetNumberOfConnections == nil {
		return 0, symbolCallError("AUGraphGetNumberOfConnections", "10.1", _aUGraphGetNumberOfConnectionsErr)
	}
	return _aUGraphGetNumberOfConnections(inGraph, outNumConnections), nil
}

// AUGraphGetNumberOfConnections deprecated in OS X v10.5.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNumberOfConnections
func AUGraphGetNumberOfConnections(inGraph AUGraph, outNumConnections *uint32) int32 {
	result, callErr := tryAUGraphGetNumberOfConnections(inGraph, outNumConnections)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphGetNumberOfInteractions func(inGraph AUGraph, outNumInteractions *uint32) int32
var _aUGraphGetNumberOfInteractionsErr error

func tryAUGraphGetNumberOfInteractions(inGraph AUGraph, outNumInteractions *uint32) (int32, error) {
	if _aUGraphGetNumberOfInteractions == nil {
		return 0, symbolCallError("AUGraphGetNumberOfInteractions", "10.5", _aUGraphGetNumberOfInteractionsErr)
	}
	return _aUGraphGetNumberOfInteractions(inGraph, outNumInteractions), nil
}

// AUGraphGetNumberOfInteractions retrieves the number of interactions for an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphGetNumberOfInteractions(_:_:)
func AUGraphGetNumberOfInteractions(inGraph AUGraph, outNumInteractions *uint32) int32 {
	result, callErr := tryAUGraphGetNumberOfInteractions(inGraph, outNumInteractions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphInitialize func(inGraph AUGraph) int32
var _aUGraphInitializeErr error

func tryAUGraphInitialize(inGraph AUGraph) (int32, error) {
	if _aUGraphInitialize == nil {
		return 0, symbolCallError("AUGraphInitialize", "10.0", _aUGraphInitializeErr)
	}
	return _aUGraphInitialize(inGraph), nil
}

// AUGraphInitialize initializes an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphInitialize(_:)
func AUGraphInitialize(inGraph AUGraph) int32 {
	result, callErr := tryAUGraphInitialize(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphIsInitialized func(inGraph AUGraph, outIsInitialized *bool) int32
var _aUGraphIsInitializedErr error

func tryAUGraphIsInitialized(inGraph AUGraph, outIsInitialized *bool) (int32, error) {
	if _aUGraphIsInitialized == nil {
		return 0, symbolCallError("AUGraphIsInitialized", "10.0", _aUGraphIsInitializedErr)
	}
	return _aUGraphIsInitialized(inGraph, outIsInitialized), nil
}

// AUGraphIsInitialized determines whether an audio processing graph is initialized.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphIsInitialized(_:_:)
func AUGraphIsInitialized(inGraph AUGraph, outIsInitialized *bool) int32 {
	result, callErr := tryAUGraphIsInitialized(inGraph, outIsInitialized)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphIsNodeSubGraph func(inGraph AUGraph, inNode AUNode, outFlag *bool) int32
var _aUGraphIsNodeSubGraphErr error

func tryAUGraphIsNodeSubGraph(inGraph AUGraph, inNode AUNode, outFlag *bool) (int32, error) {
	if _aUGraphIsNodeSubGraph == nil {
		return 0, symbolCallError("AUGraphIsNodeSubGraph", "10.2", _aUGraphIsNodeSubGraphErr)
	}
	return _aUGraphIsNodeSubGraph(inGraph, inNode, outFlag), nil
}

// AUGraphIsNodeSubGraph determines whether a node object represent an audio processing graph or an audio unit.
//
// Deprecated: Deprecated since macOS 27.0. no longer supported
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphIsNodeSubGraph(_:_:_:)
func AUGraphIsNodeSubGraph(inGraph AUGraph, inNode AUNode, outFlag *bool) int32 {
	result, callErr := tryAUGraphIsNodeSubGraph(inGraph, inNode, outFlag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphIsOpen func(inGraph AUGraph, outIsOpen *bool) int32
var _aUGraphIsOpenErr error

func tryAUGraphIsOpen(inGraph AUGraph, outIsOpen *bool) (int32, error) {
	if _aUGraphIsOpen == nil {
		return 0, symbolCallError("AUGraphIsOpen", "10.0", _aUGraphIsOpenErr)
	}
	return _aUGraphIsOpen(inGraph, outIsOpen), nil
}

// AUGraphIsOpen determines whether an audio processing graph is open.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphIsOpen(_:_:)
func AUGraphIsOpen(inGraph AUGraph, outIsOpen *bool) int32 {
	result, callErr := tryAUGraphIsOpen(inGraph, outIsOpen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphIsRunning func(inGraph AUGraph, outIsRunning *bool) int32
var _aUGraphIsRunningErr error

func tryAUGraphIsRunning(inGraph AUGraph, outIsRunning *bool) (int32, error) {
	if _aUGraphIsRunning == nil {
		return 0, symbolCallError("AUGraphIsRunning", "10.0", _aUGraphIsRunningErr)
	}
	return _aUGraphIsRunning(inGraph, outIsRunning), nil
}

// AUGraphIsRunning determines whether an audio processing graph running.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphIsRunning(_:_:)
func AUGraphIsRunning(inGraph AUGraph, outIsRunning *bool) int32 {
	result, callErr := tryAUGraphIsRunning(inGraph, outIsRunning)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphNewNode func(inGraph AUGraph, inDescription *coreservices.ComponentDescription, inClassDataSize uint32, inClassData unsafe.Pointer, outNode *AUNode) int32
var _aUGraphNewNodeErr error

func tryAUGraphNewNode(inGraph AUGraph, inDescription *coreservices.ComponentDescription, inClassDataSize uint32, inClassData unsafe.Pointer, outNode *AUNode) (int32, error) {
	if _aUGraphNewNode == nil {
		return 0, symbolCallError("AUGraphNewNode", "10.0", _aUGraphNewNodeErr)
	}
	return _aUGraphNewNode(inGraph, inDescription, inClassDataSize, inClassData, outNode), nil
}

// AUGraphNewNode deprecated in OS X v10.5.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphNewNode
func AUGraphNewNode(inGraph AUGraph, inDescription *coreservices.ComponentDescription, inClassDataSize uint32, inClassData unsafe.Pointer, outNode *AUNode) int32 {
	result, callErr := tryAUGraphNewNode(inGraph, inDescription, inClassDataSize, inClassData, outNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphNewNodeSubGraph func(inGraph AUGraph, outNode *AUNode) int32
var _aUGraphNewNodeSubGraphErr error

func tryAUGraphNewNodeSubGraph(inGraph AUGraph, outNode *AUNode) (int32, error) {
	if _aUGraphNewNodeSubGraph == nil {
		return 0, symbolCallError("AUGraphNewNodeSubGraph", "10.2", _aUGraphNewNodeSubGraphErr)
	}
	return _aUGraphNewNodeSubGraph(inGraph, outNode), nil
}

// AUGraphNewNodeSubGraph creates a node object to represent a subgraph.
//
// Deprecated: Deprecated since macOS 27.0. no longer supported
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphNewNodeSubGraph(_:_:)
func AUGraphNewNodeSubGraph(inGraph AUGraph, outNode *AUNode) int32 {
	result, callErr := tryAUGraphNewNodeSubGraph(inGraph, outNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphNodeInfo func(inGraph AUGraph, inNode AUNode, outDescription *AudioComponentDescription, outAudioUnit *AudioUnit) int32
var _aUGraphNodeInfoErr error

func tryAUGraphNodeInfo(inGraph AUGraph, inNode AUNode, outDescription *AudioComponentDescription, outAudioUnit *AudioUnit) (int32, error) {
	if _aUGraphNodeInfo == nil {
		return 0, symbolCallError("AUGraphNodeInfo", "10.5", _aUGraphNodeInfoErr)
	}
	return _aUGraphNodeInfo(inGraph, inNode, outDescription, outAudioUnit), nil
}

// AUGraphNodeInfo returns information about a node object.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphNodeInfo(_:_:_:_:)
func AUGraphNodeInfo(inGraph AUGraph, inNode AUNode, outDescription *AudioComponentDescription, outAudioUnit *AudioUnit) int32 {
	result, callErr := tryAUGraphNodeInfo(inGraph, inNode, outDescription, outAudioUnit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphOpen func(inGraph AUGraph) int32
var _aUGraphOpenErr error

func tryAUGraphOpen(inGraph AUGraph) (int32, error) {
	if _aUGraphOpen == nil {
		return 0, symbolCallError("AUGraphOpen", "10.0", _aUGraphOpenErr)
	}
	return _aUGraphOpen(inGraph), nil
}

// AUGraphOpen opens an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphOpen(_:)
func AUGraphOpen(inGraph AUGraph) int32 {
	result, callErr := tryAUGraphOpen(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphRemoveNode func(inGraph AUGraph, inNode AUNode) int32
var _aUGraphRemoveNodeErr error

func tryAUGraphRemoveNode(inGraph AUGraph, inNode AUNode) (int32, error) {
	if _aUGraphRemoveNode == nil {
		return 0, symbolCallError("AUGraphRemoveNode", "10.0", _aUGraphRemoveNodeErr)
	}
	return _aUGraphRemoveNode(inGraph, inNode), nil
}

// AUGraphRemoveNode removes a node from an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphRemoveNode(_:_:)
func AUGraphRemoveNode(inGraph AUGraph, inNode AUNode) int32 {
	result, callErr := tryAUGraphRemoveNode(inGraph, inNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphRemoveRenderNotify func(inGraph AUGraph, inCallback AURenderCallback, inRefCon unsafe.Pointer) int32
var _aUGraphRemoveRenderNotifyErr error

func tryAUGraphRemoveRenderNotify(inGraph AUGraph, inCallback AURenderCallback, inRefCon unsafe.Pointer) (int32, error) {
	if _aUGraphRemoveRenderNotify == nil {
		return 0, symbolCallError("AUGraphRemoveRenderNotify", "10.2", _aUGraphRemoveRenderNotifyErr)
	}
	return _aUGraphRemoveRenderNotify(inGraph, inCallback, inRefCon), nil
}

// AUGraphRemoveRenderNotify removes a notification callback from an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphRemoveRenderNotify(_:_:_:)
func AUGraphRemoveRenderNotify(inGraph AUGraph, inCallback AURenderCallback, inRefCon unsafe.Pointer) int32 {
	result, callErr := tryAUGraphRemoveRenderNotify(inGraph, inCallback, inRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphSetNodeInputCallback func(inGraph AUGraph, inDestNode AUNode, inDestInputNumber uint32, inInputCallback *AURenderCallbackStruct) int32
var _aUGraphSetNodeInputCallbackErr error

func tryAUGraphSetNodeInputCallback(inGraph AUGraph, inDestNode AUNode, inDestInputNumber uint32, inInputCallback *AURenderCallbackStruct) (int32, error) {
	if _aUGraphSetNodeInputCallback == nil {
		return 0, symbolCallError("AUGraphSetNodeInputCallback", "10.5", _aUGraphSetNodeInputCallbackErr)
	}
	return _aUGraphSetNodeInputCallback(inGraph, inDestNode, inDestInputNumber, inInputCallback), nil
}

// AUGraphSetNodeInputCallback sets an input callback function for a node.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphSetNodeInputCallback(_:_:_:_:)
func AUGraphSetNodeInputCallback(inGraph AUGraph, inDestNode AUNode, inDestInputNumber uint32, inInputCallback *AURenderCallbackStruct) int32 {
	result, callErr := tryAUGraphSetNodeInputCallback(inGraph, inDestNode, inDestInputNumber, inInputCallback)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphStart func(inGraph AUGraph) int32
var _aUGraphStartErr error

func tryAUGraphStart(inGraph AUGraph) (int32, error) {
	if _aUGraphStart == nil {
		return 0, symbolCallError("AUGraphStart", "10.0", _aUGraphStartErr)
	}
	return _aUGraphStart(inGraph), nil
}

// AUGraphStart starts an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphStart(_:)
func AUGraphStart(inGraph AUGraph) int32 {
	result, callErr := tryAUGraphStart(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphStop func(inGraph AUGraph) int32
var _aUGraphStopErr error

func tryAUGraphStop(inGraph AUGraph) (int32, error) {
	if _aUGraphStop == nil {
		return 0, symbolCallError("AUGraphStop", "10.0", _aUGraphStopErr)
	}
	return _aUGraphStop(inGraph), nil
}

// AUGraphStop stops an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphStop(_:)
func AUGraphStop(inGraph AUGraph) int32 {
	result, callErr := tryAUGraphStop(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphUninitialize func(inGraph AUGraph) int32
var _aUGraphUninitializeErr error

func tryAUGraphUninitialize(inGraph AUGraph) (int32, error) {
	if _aUGraphUninitialize == nil {
		return 0, symbolCallError("AUGraphUninitialize", "10.0", _aUGraphUninitializeErr)
	}
	return _aUGraphUninitialize(inGraph), nil
}

// AUGraphUninitialize uninitializes an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphUninitialize(_:)
func AUGraphUninitialize(inGraph AUGraph) int32 {
	result, callErr := tryAUGraphUninitialize(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUGraphUpdate func(inGraph AUGraph, outIsUpdated *bool) int32
var _aUGraphUpdateErr error

func tryAUGraphUpdate(inGraph AUGraph, outIsUpdated *bool) (int32, error) {
	if _aUGraphUpdate == nil {
		return 0, symbolCallError("AUGraphUpdate", "10.0", _aUGraphUpdateErr)
	}
	return _aUGraphUpdate(inGraph, outIsUpdated), nil
}

// AUGraphUpdate updates the state of a running audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUGraphUpdate(_:_:)
func AUGraphUpdate(inGraph AUGraph, outIsUpdated *bool) int32 {
	result, callErr := tryAUGraphUpdate(inGraph, outIsUpdated)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUListenerAddParameter func(inListener AUParameterListenerRef, inObject unsafe.Pointer, inParameter *AudioUnitParameter) int32
var _aUListenerAddParameterErr error

func tryAUListenerAddParameter(inListener AUParameterListenerRef, inObject unsafe.Pointer, inParameter *AudioUnitParameter) (int32, error) {
	if _aUListenerAddParameter == nil {
		return 0, symbolCallError("AUListenerAddParameter", "10.2", _aUListenerAddParameterErr)
	}
	return _aUListenerAddParameter(inListener, inObject, inParameter), nil
}

// AUListenerAddParameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUListenerAddParameter(_:_:_:)
func AUListenerAddParameter(inListener AUParameterListenerRef, inObject unsafe.Pointer, inParameter *AudioUnitParameter) int32 {
	result, callErr := tryAUListenerAddParameter(inListener, inObject, inParameter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUListenerCreate func(inProc AUParameterListenerProc, inUserData unsafe.Pointer, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inNotificationInterval float32, outListener *AUParameterListenerRef) int32
var _aUListenerCreateErr error

func tryAUListenerCreate(inProc AUParameterListenerProc, inUserData unsafe.Pointer, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inNotificationInterval float32, outListener *AUParameterListenerRef) (int32, error) {
	if _aUListenerCreate == nil {
		return 0, symbolCallError("AUListenerCreate", "10.2", _aUListenerCreateErr)
	}
	return _aUListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, outListener), nil
}

// AUListenerCreate.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUListenerCreate(_:_:_:_:_:_:)
func AUListenerCreate(inProc AUParameterListenerProc, inUserData unsafe.Pointer, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inNotificationInterval float32, outListener *AUParameterListenerRef) int32 {
	result, callErr := tryAUListenerCreate(inProc, inUserData, inRunLoop, inRunLoopMode, inNotificationInterval, outListener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUListenerCreateWithDispatchQueue func(outListener *AUParameterListenerRef, inNotificationInterval float32, inDispatchQueue uintptr, inBlock unsafe.Pointer) int32
var _aUListenerCreateWithDispatchQueueErr error

func tryAUListenerCreateWithDispatchQueue(outListener *AUParameterListenerRef, inNotificationInterval float32, inDispatchQueue dispatch.Queue, inBlock AUParameterListenerBlock) (int32, error) {
	if _aUListenerCreateWithDispatchQueue == nil {
		return 0, symbolCallError("AUListenerCreateWithDispatchQueue", "10.6", _aUListenerCreateWithDispatchQueueErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 unsafe.Pointer, blockArg1 *AudioUnitParameter, blockArg2 float32) {
		inBlock(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _aUListenerCreateWithDispatchQueue(outListener, inNotificationInterval, uintptr(inDispatchQueue.Handle()), _block0), nil
}

// AUListenerCreateWithDispatchQueue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUListenerCreateWithDispatchQueue(_:_:_:_:)
func AUListenerCreateWithDispatchQueue(outListener *AUParameterListenerRef, inNotificationInterval float32, inDispatchQueue dispatch.Queue, inBlock AUParameterListenerBlock) int32 {
	result, callErr := tryAUListenerCreateWithDispatchQueue(outListener, inNotificationInterval, inDispatchQueue, inBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUListenerDispose func(inListener AUParameterListenerRef) int32
var _aUListenerDisposeErr error

func tryAUListenerDispose(inListener AUParameterListenerRef) (int32, error) {
	if _aUListenerDispose == nil {
		return 0, symbolCallError("AUListenerDispose", "10.2", _aUListenerDisposeErr)
	}
	return _aUListenerDispose(inListener), nil
}

// AUListenerDispose.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUListenerDispose(_:)
func AUListenerDispose(inListener AUParameterListenerRef) int32 {
	result, callErr := tryAUListenerDispose(inListener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUListenerRemoveParameter func(inListener AUParameterListenerRef, inObject unsafe.Pointer, inParameter *AudioUnitParameter) int32
var _aUListenerRemoveParameterErr error

func tryAUListenerRemoveParameter(inListener AUParameterListenerRef, inObject unsafe.Pointer, inParameter *AudioUnitParameter) (int32, error) {
	if _aUListenerRemoveParameter == nil {
		return 0, symbolCallError("AUListenerRemoveParameter", "10.2", _aUListenerRemoveParameterErr)
	}
	return _aUListenerRemoveParameter(inListener, inObject, inParameter), nil
}

// AUListenerRemoveParameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUListenerRemoveParameter(_:_:_:)
func AUListenerRemoveParameter(inListener AUParameterListenerRef, inObject unsafe.Pointer, inParameter *AudioUnitParameter) int32 {
	result, callErr := tryAUListenerRemoveParameter(inListener, inObject, inParameter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUParameterFormatValue func(inParameterValue float64, inParameter *AudioUnitParameter, inTextBuffer *byte, inDigits uint32) *byte
var _aUParameterFormatValueErr error

func tryAUParameterFormatValue(inParameterValue float64, inParameter *AudioUnitParameter, inTextBuffer *byte, inDigits uint32) (*byte, error) {
	if _aUParameterFormatValue == nil {
		return nil, symbolCallError("AUParameterFormatValue", "10.2", _aUParameterFormatValueErr)
	}
	return _aUParameterFormatValue(inParameterValue, inParameter, inTextBuffer, inDigits), nil
}

// AUParameterFormatValue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterFormatValue(_:_:_:_:)
func AUParameterFormatValue(inParameterValue float64, inParameter *AudioUnitParameter, inTextBuffer *byte, inDigits uint32) *byte {
	result, callErr := tryAUParameterFormatValue(inParameterValue, inParameter, inTextBuffer, inDigits)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUParameterListenerNotify func(inSendingListener AUParameterListenerRef, inSendingObject unsafe.Pointer, inParameter *AudioUnitParameter) int32
var _aUParameterListenerNotifyErr error

func tryAUParameterListenerNotify(inSendingListener AUParameterListenerRef, inSendingObject unsafe.Pointer, inParameter *AudioUnitParameter) (int32, error) {
	if _aUParameterListenerNotify == nil {
		return 0, symbolCallError("AUParameterListenerNotify", "10.2", _aUParameterListenerNotifyErr)
	}
	return _aUParameterListenerNotify(inSendingListener, inSendingObject, inParameter), nil
}

// AUParameterListenerNotify.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterListenerNotify(_:_:_:)
func AUParameterListenerNotify(inSendingListener AUParameterListenerRef, inSendingObject unsafe.Pointer, inParameter *AudioUnitParameter) int32 {
	result, callErr := tryAUParameterListenerNotify(inSendingListener, inSendingObject, inParameter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUParameterSet func(inSendingListener AUParameterListenerRef, inSendingObject unsafe.Pointer, inParameter *AudioUnitParameter, inValue AudioUnitParameterValue, inBufferOffsetInFrames uint32) int32
var _aUParameterSetErr error

func tryAUParameterSet(inSendingListener AUParameterListenerRef, inSendingObject unsafe.Pointer, inParameter *AudioUnitParameter, inValue AudioUnitParameterValue, inBufferOffsetInFrames uint32) (int32, error) {
	if _aUParameterSet == nil {
		return 0, symbolCallError("AUParameterSet", "10.2", _aUParameterSetErr)
	}
	return _aUParameterSet(inSendingListener, inSendingObject, inParameter, inValue, inBufferOffsetInFrames), nil
}

// AUParameterSet.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterSet(_:_:_:_:_:)
func AUParameterSet(inSendingListener AUParameterListenerRef, inSendingObject unsafe.Pointer, inParameter *AudioUnitParameter, inValue AudioUnitParameterValue, inBufferOffsetInFrames uint32) int32 {
	result, callErr := tryAUParameterSet(inSendingListener, inSendingObject, inParameter, inValue, inBufferOffsetInFrames)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUParameterValueFromLinear func(inLinearValue float32, inParameter *AudioUnitParameter) AudioUnitParameterValue
var _aUParameterValueFromLinearErr error

func tryAUParameterValueFromLinear(inLinearValue float32, inParameter *AudioUnitParameter) (AudioUnitParameterValue, error) {
	if _aUParameterValueFromLinear == nil {
		return *new(AudioUnitParameterValue), symbolCallError("AUParameterValueFromLinear", "10.2", _aUParameterValueFromLinearErr)
	}
	return _aUParameterValueFromLinear(inLinearValue, inParameter), nil
}

// AUParameterValueFromLinear.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterValueFromLinear(_:_:)
func AUParameterValueFromLinear(inLinearValue float32, inParameter *AudioUnitParameter) AudioUnitParameterValue {
	result, callErr := tryAUParameterValueFromLinear(inLinearValue, inParameter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aUParameterValueToLinear func(inParameterValue AudioUnitParameterValue, inParameter *AudioUnitParameter) float32
var _aUParameterValueToLinearErr error

func tryAUParameterValueToLinear(inParameterValue AudioUnitParameterValue, inParameter *AudioUnitParameter) (float32, error) {
	if _aUParameterValueToLinear == nil {
		return 0.0, symbolCallError("AUParameterValueToLinear", "10.2", _aUParameterValueToLinearErr)
	}
	return _aUParameterValueToLinear(inParameterValue, inParameter), nil
}

// AUParameterValueToLinear.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUParameterValueToLinear(_:_:)
func AUParameterValueToLinear(inParameterValue AudioUnitParameterValue, inParameter *AudioUnitParameter) float32 {
	result, callErr := tryAUParameterValueToLinear(inParameterValue, inParameter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecAppendInputBufferList func(inCodec AudioCodec, inBufferList *coreaudiotypes.AudioBufferList, ioNumberPackets *uint32, inPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outBytesConsumed *uint32) int32
var _audioCodecAppendInputBufferListErr error

func tryAudioCodecAppendInputBufferList(inCodec AudioCodec, inBufferList *coreaudiotypes.AudioBufferList, ioNumberPackets *uint32, inPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outBytesConsumed *uint32) (int32, error) {
	if _audioCodecAppendInputBufferList == nil {
		return 0, symbolCallError("AudioCodecAppendInputBufferList", "10.7", _audioCodecAppendInputBufferListErr)
	}
	return _audioCodecAppendInputBufferList(inCodec, inBufferList, ioNumberPackets, inPacketDescription, outBytesConsumed), nil
}

// AudioCodecAppendInputBufferList.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputBufferList(_:_:_:_:_:)
func AudioCodecAppendInputBufferList(inCodec AudioCodec, inBufferList *coreaudiotypes.AudioBufferList, ioNumberPackets *uint32, inPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outBytesConsumed *uint32) int32 {
	result, callErr := tryAudioCodecAppendInputBufferList(inCodec, inBufferList, ioNumberPackets, inPacketDescription, outBytesConsumed)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecAppendInputData func(inCodec AudioCodec, inInputData unsafe.Pointer, ioInputDataByteSize *uint32, ioNumberPackets *uint32, inPacketDescription *coreaudiotypes.AudioStreamPacketDescription) int32
var _audioCodecAppendInputDataErr error

func tryAudioCodecAppendInputData(inCodec AudioCodec, inInputData unsafe.Pointer, ioInputDataByteSize *uint32, ioNumberPackets *uint32, inPacketDescription *coreaudiotypes.AudioStreamPacketDescription) (int32, error) {
	if _audioCodecAppendInputData == nil {
		return 0, symbolCallError("AudioCodecAppendInputData", "10.2", _audioCodecAppendInputDataErr)
	}
	return _audioCodecAppendInputData(inCodec, inInputData, ioInputDataByteSize, ioNumberPackets, inPacketDescription), nil
}

// AudioCodecAppendInputData appends audio data to the codec’s input buffer.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecAppendInputData(_:_:_:_:_:)
func AudioCodecAppendInputData(inCodec AudioCodec, inInputData unsafe.Pointer, ioInputDataByteSize *uint32, ioNumberPackets *uint32, inPacketDescription *coreaudiotypes.AudioStreamPacketDescription) int32 {
	result, callErr := tryAudioCodecAppendInputData(inCodec, inInputData, ioInputDataByteSize, ioNumberPackets, inPacketDescription)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecGetProperty func(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioCodecGetPropertyErr error

func tryAudioCodecGetProperty(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioCodecGetProperty == nil {
		return 0, symbolCallError("AudioCodecGetProperty", "10.2", _audioCodecGetPropertyErr)
	}
	return _audioCodecGetProperty(inCodec, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioCodecGetProperty retrieves the value of a codec property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetProperty(_:_:_:_:)
func AudioCodecGetProperty(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioCodecGetProperty(inCodec, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecGetPropertyInfo func(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, outSize *uint32, outWritable *bool) int32
var _audioCodecGetPropertyInfoErr error

func tryAudioCodecGetPropertyInfo(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, outSize *uint32, outWritable *bool) (int32, error) {
	if _audioCodecGetPropertyInfo == nil {
		return 0, symbolCallError("AudioCodecGetPropertyInfo", "10.2", _audioCodecGetPropertyInfoErr)
	}
	return _audioCodecGetPropertyInfo(inCodec, inPropertyID, outSize, outWritable), nil
}

// AudioCodecGetPropertyInfo retrieves information about a codec property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecGetPropertyInfo(_:_:_:_:)
func AudioCodecGetPropertyInfo(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioCodecGetPropertyInfo(inCodec, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecInitialize func(inCodec AudioCodec, inInputFormat *coreaudiotypes.AudioStreamBasicDescription, inOutputFormat *coreaudiotypes.AudioStreamBasicDescription, inMagicCookie unsafe.Pointer, inMagicCookieByteSize uint32) int32
var _audioCodecInitializeErr error

func tryAudioCodecInitialize(inCodec AudioCodec, inInputFormat *coreaudiotypes.AudioStreamBasicDescription, inOutputFormat *coreaudiotypes.AudioStreamBasicDescription, inMagicCookie unsafe.Pointer, inMagicCookieByteSize uint32) (int32, error) {
	if _audioCodecInitialize == nil {
		return 0, symbolCallError("AudioCodecInitialize", "10.2", _audioCodecInitializeErr)
	}
	return _audioCodecInitialize(inCodec, inInputFormat, inOutputFormat, inMagicCookie, inMagicCookieByteSize), nil
}

// AudioCodecInitialize sets up the specified codec to perform a data format translation.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecInitialize(_:_:_:_:_:)
func AudioCodecInitialize(inCodec AudioCodec, inInputFormat *coreaudiotypes.AudioStreamBasicDescription, inOutputFormat *coreaudiotypes.AudioStreamBasicDescription, inMagicCookie unsafe.Pointer, inMagicCookieByteSize uint32) int32 {
	result, callErr := tryAudioCodecInitialize(inCodec, inInputFormat, inOutputFormat, inMagicCookie, inMagicCookieByteSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecProduceOutputBufferList func(inCodec AudioCodec, ioBufferList *coreaudiotypes.AudioBufferList, ioNumberPackets *uint32, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outStatus *uint32) int32
var _audioCodecProduceOutputBufferListErr error

func tryAudioCodecProduceOutputBufferList(inCodec AudioCodec, ioBufferList *coreaudiotypes.AudioBufferList, ioNumberPackets *uint32, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outStatus *uint32) (int32, error) {
	if _audioCodecProduceOutputBufferList == nil {
		return 0, symbolCallError("AudioCodecProduceOutputBufferList", "10.7", _audioCodecProduceOutputBufferListErr)
	}
	return _audioCodecProduceOutputBufferList(inCodec, ioBufferList, ioNumberPackets, outPacketDescription, outStatus), nil
}

// AudioCodecProduceOutputBufferList.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputBufferList(_:_:_:_:_:)
func AudioCodecProduceOutputBufferList(inCodec AudioCodec, ioBufferList *coreaudiotypes.AudioBufferList, ioNumberPackets *uint32, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outStatus *uint32) int32 {
	result, callErr := tryAudioCodecProduceOutputBufferList(inCodec, ioBufferList, ioNumberPackets, outPacketDescription, outStatus)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecProduceOutputPackets func(inCodec AudioCodec, outOutputData unsafe.Pointer, ioOutputDataByteSize *uint32, ioNumberPackets *uint32, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outStatus *uint32) int32
var _audioCodecProduceOutputPacketsErr error

func tryAudioCodecProduceOutputPackets(inCodec AudioCodec, outOutputData unsafe.Pointer, ioOutputDataByteSize *uint32, ioNumberPackets *uint32, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outStatus *uint32) (int32, error) {
	if _audioCodecProduceOutputPackets == nil {
		return 0, symbolCallError("AudioCodecProduceOutputPackets", "10.2", _audioCodecProduceOutputPacketsErr)
	}
	return _audioCodecProduceOutputPackets(inCodec, outOutputData, ioOutputDataByteSize, ioNumberPackets, outPacketDescription, outStatus), nil
}

// AudioCodecProduceOutputPackets retrieves output data from a codec.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecProduceOutputPackets(_:_:_:_:_:_:)
func AudioCodecProduceOutputPackets(inCodec AudioCodec, outOutputData unsafe.Pointer, ioOutputDataByteSize *uint32, ioNumberPackets *uint32, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription, outStatus *uint32) int32 {
	result, callErr := tryAudioCodecProduceOutputPackets(inCodec, outOutputData, ioOutputDataByteSize, ioNumberPackets, outPacketDescription, outStatus)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecReset func(inCodec AudioCodec) int32
var _audioCodecResetErr error

func tryAudioCodecReset(inCodec AudioCodec) (int32, error) {
	if _audioCodecReset == nil {
		return 0, symbolCallError("AudioCodecReset", "10.2", _audioCodecResetErr)
	}
	return _audioCodecReset(inCodec), nil
}

// AudioCodecReset flushes all the audio data in the codec and clears the input buffer.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecReset(_:)
func AudioCodecReset(inCodec AudioCodec) int32 {
	result, callErr := tryAudioCodecReset(inCodec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecSetProperty func(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioCodecSetPropertyErr error

func tryAudioCodecSetProperty(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioCodecSetProperty == nil {
		return 0, symbolCallError("AudioCodecSetProperty", "10.2", _audioCodecSetPropertyErr)
	}
	return _audioCodecSetProperty(inCodec, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioCodecSetProperty sets the value of a codec property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecSetProperty(_:_:_:_:)
func AudioCodecSetProperty(inCodec AudioCodec, inPropertyID AudioCodecPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioCodecSetProperty(inCodec, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioCodecUninitialize func(inCodec AudioCodec) int32
var _audioCodecUninitializeErr error

func tryAudioCodecUninitialize(inCodec AudioCodec) (int32, error) {
	if _audioCodecUninitialize == nil {
		return 0, symbolCallError("AudioCodecUninitialize", "10.2", _audioCodecUninitializeErr)
	}
	return _audioCodecUninitialize(inCodec), nil
}

// AudioCodecUninitialize moves the codec from the initialized state back to the uninitialized state.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioCodecUninitialize(_:)
func AudioCodecUninitialize(inCodec AudioCodec) int32 {
	result, callErr := tryAudioCodecUninitialize(inCodec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentCopyConfigurationInfo func(inComponent AudioComponent, outConfigurationInfo *corefoundation.CFDictionaryRef) int32
var _audioComponentCopyConfigurationInfoErr error

func tryAudioComponentCopyConfigurationInfo(inComponent AudioComponent, outConfigurationInfo *corefoundation.CFDictionaryRef) (int32, error) {
	if _audioComponentCopyConfigurationInfo == nil {
		return 0, symbolCallError("AudioComponentCopyConfigurationInfo", "10.7", _audioComponentCopyConfigurationInfoErr)
	}
	return _audioComponentCopyConfigurationInfo(inComponent, outConfigurationInfo), nil
}

// AudioComponentCopyConfigurationInfo.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyConfigurationInfo(_:_:)
func AudioComponentCopyConfigurationInfo(inComponent AudioComponent, outConfigurationInfo *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryAudioComponentCopyConfigurationInfo(inComponent, outConfigurationInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentCopyIcon func(comp AudioComponent) *objc.ID
var _audioComponentCopyIconErr error

func tryAudioComponentCopyIcon(comp AudioComponent) (*objc.ID, error) {
	if _audioComponentCopyIcon == nil {
		return nil, symbolCallError("AudioComponentCopyIcon", "11.0", _audioComponentCopyIconErr)
	}
	return _audioComponentCopyIcon(comp), nil
}

// AudioComponentCopyIcon.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyIcon(_:)
func AudioComponentCopyIcon(comp AudioComponent) *objc.ID {
	result, callErr := tryAudioComponentCopyIcon(comp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentCopyName func(inComponent AudioComponent, outName *corefoundation.CFStringRef) int32
var _audioComponentCopyNameErr error

func tryAudioComponentCopyName(inComponent AudioComponent, outName *corefoundation.CFStringRef) (int32, error) {
	if _audioComponentCopyName == nil {
		return 0, symbolCallError("AudioComponentCopyName", "10.6", _audioComponentCopyNameErr)
	}
	return _audioComponentCopyName(inComponent, outName), nil
}

// AudioComponentCopyName returns the generic name of an audio component.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCopyName(_:_:)
func AudioComponentCopyName(inComponent AudioComponent, outName *corefoundation.CFStringRef) int32 {
	result, callErr := tryAudioComponentCopyName(inComponent, outName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentCount func(inDesc *AudioComponentDescription) uint32
var _audioComponentCountErr error

func tryAudioComponentCount(inDesc *AudioComponentDescription) (uint32, error) {
	if _audioComponentCount == nil {
		return 0, symbolCallError("AudioComponentCount", "10.6", _audioComponentCountErr)
	}
	return _audioComponentCount(inDesc), nil
}

// AudioComponentCount returns the number of audio components that match a specified [AudioComponentDescription] structure.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentCount(_:)
func AudioComponentCount(inDesc *AudioComponentDescription) uint32 {
	result, callErr := tryAudioComponentCount(inDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentFindNext func(inComponent AudioComponent, inDesc *AudioComponentDescription) AudioComponent
var _audioComponentFindNextErr error

func tryAudioComponentFindNext(inComponent AudioComponent, inDesc *AudioComponentDescription) (AudioComponent, error) {
	if _audioComponentFindNext == nil {
		return *new(AudioComponent), symbolCallError("AudioComponentFindNext", "10.6", _audioComponentFindNextErr)
	}
	return _audioComponentFindNext(inComponent, inDesc), nil
}

// AudioComponentFindNext finds the next component that matches a specified [AudioComponentDescription] structure after a specified audio component.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentFindNext(_:_:)
func AudioComponentFindNext(inComponent AudioComponent, inDesc *AudioComponentDescription) AudioComponent {
	result, callErr := tryAudioComponentFindNext(inComponent, inDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentGetDescription func(inComponent AudioComponent, outDesc *AudioComponentDescription) int32
var _audioComponentGetDescriptionErr error

func tryAudioComponentGetDescription(inComponent AudioComponent, outDesc *AudioComponentDescription) (int32, error) {
	if _audioComponentGetDescription == nil {
		return 0, symbolCallError("AudioComponentGetDescription", "10.6", _audioComponentGetDescriptionErr)
	}
	return _audioComponentGetDescription(inComponent, outDesc), nil
}

// AudioComponentGetDescription gets the class description, as an [AudioComponentDescription] structure, of an audio component.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetDescription(_:_:)
func AudioComponentGetDescription(inComponent AudioComponent, outDesc *AudioComponentDescription) int32 {
	result, callErr := tryAudioComponentGetDescription(inComponent, outDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentGetIcon func(comp AudioComponent) *objc.ID
var _audioComponentGetIconErr error

func tryAudioComponentGetIcon(comp AudioComponent) (*objc.ID, error) {
	if _audioComponentGetIcon == nil {
		return nil, symbolCallError("AudioComponentGetIcon", "10.11", _audioComponentGetIconErr)
	}
	return _audioComponentGetIcon(comp), nil
}

// AudioComponentGetIcon the UIImage of the audio component’s icon.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetIcon(_:_:)
func AudioComponentGetIcon(comp AudioComponent) *objc.ID {
	result, callErr := tryAudioComponentGetIcon(comp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentGetVersion func(inComponent AudioComponent, outVersion *uint32) int32
var _audioComponentGetVersionErr error

func tryAudioComponentGetVersion(inComponent AudioComponent, outVersion *uint32) (int32, error) {
	if _audioComponentGetVersion == nil {
		return 0, symbolCallError("AudioComponentGetVersion", "10.6", _audioComponentGetVersionErr)
	}
	return _audioComponentGetVersion(inComponent, outVersion), nil
}

// AudioComponentGetVersion gets the version of an audio component in hexadecimal form as `0xMMMMmmDD` (major, minor, dot).
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentGetVersion(_:_:)
func AudioComponentGetVersion(inComponent AudioComponent, outVersion *uint32) int32 {
	result, callErr := tryAudioComponentGetVersion(inComponent, outVersion)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentInstanceCanDo func(inInstance AudioComponentInstance, inSelectorID int16) bool
var _audioComponentInstanceCanDoErr error

func tryAudioComponentInstanceCanDo(inInstance AudioComponentInstance, inSelectorID int16) (bool, error) {
	if _audioComponentInstanceCanDo == nil {
		return false, symbolCallError("AudioComponentInstanceCanDo", "10.6", _audioComponentInstanceCanDoErr)
	}
	return _audioComponentInstanceCanDo(inInstance, inSelectorID), nil
}

// AudioComponentInstanceCanDo determines if an audio component instance implements a particular function.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceCanDo(_:_:)
func AudioComponentInstanceCanDo(inInstance AudioComponentInstance, inSelectorID int16) bool {
	result, callErr := tryAudioComponentInstanceCanDo(inInstance, inSelectorID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentInstanceDispose func(inInstance AudioComponentInstance) int32
var _audioComponentInstanceDisposeErr error

func tryAudioComponentInstanceDispose(inInstance AudioComponentInstance) (int32, error) {
	if _audioComponentInstanceDispose == nil {
		return 0, symbolCallError("AudioComponentInstanceDispose", "10.6", _audioComponentInstanceDisposeErr)
	}
	return _audioComponentInstanceDispose(inInstance), nil
}

// AudioComponentInstanceDispose disposes of an audio component instance.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceDispose(_:)
func AudioComponentInstanceDispose(inInstance AudioComponentInstance) int32 {
	result, callErr := tryAudioComponentInstanceDispose(inInstance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentInstanceGetComponent func(inInstance AudioComponentInstance) AudioComponent
var _audioComponentInstanceGetComponentErr error

func tryAudioComponentInstanceGetComponent(inInstance AudioComponentInstance) (AudioComponent, error) {
	if _audioComponentInstanceGetComponent == nil {
		return *new(AudioComponent), symbolCallError("AudioComponentInstanceGetComponent", "10.6", _audioComponentInstanceGetComponentErr)
	}
	return _audioComponentInstanceGetComponent(inInstance), nil
}

// AudioComponentInstanceGetComponent retrieves a reference to an audio component from an instance of that audio component.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceGetComponent(_:)
func AudioComponentInstanceGetComponent(inInstance AudioComponentInstance) AudioComponent {
	result, callErr := tryAudioComponentInstanceGetComponent(inInstance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentInstanceNew func(inComponent AudioComponent, outInstance *AudioComponentInstance) int32
var _audioComponentInstanceNewErr error

func tryAudioComponentInstanceNew(inComponent AudioComponent, outInstance *AudioComponentInstance) (int32, error) {
	if _audioComponentInstanceNew == nil {
		return 0, symbolCallError("AudioComponentInstanceNew", "10.6", _audioComponentInstanceNewErr)
	}
	return _audioComponentInstanceNew(inComponent, outInstance), nil
}

// AudioComponentInstanceNew creates a new instance of an audio component.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstanceNew(_:_:)
func AudioComponentInstanceNew(inComponent AudioComponent, outInstance *AudioComponentInstance) int32 {
	result, callErr := tryAudioComponentInstanceNew(inComponent, outInstance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentInstantiate func(inComponent AudioComponent, inOptions AudioComponentInstantiationOptions)
var _audioComponentInstantiateErr error

func tryAudioComponentInstantiate(inComponent AudioComponent, inOptions AudioComponentInstantiationOptions) error {
	if _audioComponentInstantiate == nil {
		return symbolCallError("AudioComponentInstantiate", "10.11", _audioComponentInstantiateErr)
	}
	_audioComponentInstantiate(inComponent, inOptions)
	return nil
}

// AudioComponentInstantiate.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentInstantiate(_:_:_:)
func AudioComponentInstantiate(inComponent AudioComponent, inOptions AudioComponentInstantiationOptions) {
	if callErr := tryAudioComponentInstantiate(inComponent, inOptions); callErr != nil {
		panic(callErr)
	}
}

var _audioComponentRegister func(inDesc *AudioComponentDescription, inName corefoundation.CFStringRef, inVersion uint32, inFactory AudioComponentFactoryFunction) AudioComponent
var _audioComponentRegisterErr error

func tryAudioComponentRegister(inDesc *AudioComponentDescription, inName corefoundation.CFStringRef, inVersion uint32, inFactory AudioComponentFactoryFunction) (AudioComponent, error) {
	if _audioComponentRegister == nil {
		return *new(AudioComponent), symbolCallError("AudioComponentRegister", "10.7", _audioComponentRegisterErr)
	}
	return _audioComponentRegister(inDesc, inName, inVersion, inFactory), nil
}

// AudioComponentRegister.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentRegister(_:_:_:_:)
func AudioComponentRegister(inDesc *AudioComponentDescription, inName corefoundation.CFStringRef, inVersion uint32, inFactory AudioComponentFactoryFunction) AudioComponent {
	result, callErr := tryAudioComponentRegister(inDesc, inName, inVersion, inFactory)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentValidate func(inComponent AudioComponent, inValidationParameters corefoundation.CFDictionaryRef, outValidationResult *AudioComponentValidationResult) int32
var _audioComponentValidateErr error

func tryAudioComponentValidate(inComponent AudioComponent, inValidationParameters corefoundation.CFDictionaryRef, outValidationResult *AudioComponentValidationResult) (int32, error) {
	if _audioComponentValidate == nil {
		return 0, symbolCallError("AudioComponentValidate", "10.7", _audioComponentValidateErr)
	}
	return _audioComponentValidate(inComponent, inValidationParameters, outValidationResult), nil
}

// AudioComponentValidate.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidate(_:_:_:)
func AudioComponentValidate(inComponent AudioComponent, inValidationParameters corefoundation.CFDictionaryRef, outValidationResult *AudioComponentValidationResult) int32 {
	result, callErr := tryAudioComponentValidate(inComponent, inValidationParameters, outValidationResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioComponentValidateWithResults func(inComponent AudioComponent, inValidationParameters corefoundation.CFDictionaryRef) int32
var _audioComponentValidateWithResultsErr error

func tryAudioComponentValidateWithResults(inComponent AudioComponent, inValidationParameters corefoundation.CFDictionaryRef) (int32, error) {
	if _audioComponentValidateWithResults == nil {
		return 0, symbolCallError("AudioComponentValidateWithResults", "13.0", _audioComponentValidateWithResultsErr)
	}
	return _audioComponentValidateWithResults(inComponent, inValidationParameters), nil
}

// AudioComponentValidateWithResults.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioComponentValidateWithResults(_:_:_:)
func AudioComponentValidateWithResults(inComponent AudioComponent, inValidationParameters corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryAudioComponentValidateWithResults(inComponent, inValidationParameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterConvertBuffer func(inAudioConverter AudioConverterRef, inInputDataSize uint32, inInputData unsafe.Pointer, ioOutputDataSize *uint32, outOutputData unsafe.Pointer) int32
var _audioConverterConvertBufferErr error

func tryAudioConverterConvertBuffer(inAudioConverter AudioConverterRef, inInputDataSize uint32, inInputData unsafe.Pointer, ioOutputDataSize *uint32, outOutputData unsafe.Pointer) (int32, error) {
	if _audioConverterConvertBuffer == nil {
		return 0, symbolCallError("AudioConverterConvertBuffer", "10.1", _audioConverterConvertBufferErr)
	}
	return _audioConverterConvertBuffer(inAudioConverter, inInputDataSize, inInputData, ioOutputDataSize, outOutputData), nil
}

// AudioConverterConvertBuffer converts audio data from one linear PCM format to another.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterConvertBuffer(_:_:_:_:_:)
func AudioConverterConvertBuffer(inAudioConverter AudioConverterRef, inInputDataSize uint32, inInputData unsafe.Pointer, ioOutputDataSize *uint32, outOutputData unsafe.Pointer) int32 {
	result, callErr := tryAudioConverterConvertBuffer(inAudioConverter, inInputDataSize, inInputData, ioOutputDataSize, outOutputData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterConvertComplexBuffer func(inAudioConverter AudioConverterRef, inNumberPCMFrames uint32, inInputData *coreaudiotypes.AudioBufferList, outOutputData *coreaudiotypes.AudioBufferList) int32
var _audioConverterConvertComplexBufferErr error

func tryAudioConverterConvertComplexBuffer(inAudioConverter AudioConverterRef, inNumberPCMFrames uint32, inInputData *coreaudiotypes.AudioBufferList, outOutputData *coreaudiotypes.AudioBufferList) (int32, error) {
	if _audioConverterConvertComplexBuffer == nil {
		return 0, symbolCallError("AudioConverterConvertComplexBuffer", "10.7", _audioConverterConvertComplexBufferErr)
	}
	return _audioConverterConvertComplexBuffer(inAudioConverter, inNumberPCMFrames, inInputData, outOutputData), nil
}

// AudioConverterConvertComplexBuffer converts audio data from one linear PCM format to another, where both use the same sample rate.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterConvertComplexBuffer(_:_:_:_:)
func AudioConverterConvertComplexBuffer(inAudioConverter AudioConverterRef, inNumberPCMFrames uint32, inInputData *coreaudiotypes.AudioBufferList, outOutputData *coreaudiotypes.AudioBufferList) int32 {
	result, callErr := tryAudioConverterConvertComplexBuffer(inAudioConverter, inNumberPCMFrames, inInputData, outOutputData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterDispose func(inAudioConverter AudioConverterRef) int32
var _audioConverterDisposeErr error

func tryAudioConverterDispose(inAudioConverter AudioConverterRef) (int32, error) {
	if _audioConverterDispose == nil {
		return 0, symbolCallError("AudioConverterDispose", "10.1", _audioConverterDisposeErr)
	}
	return _audioConverterDispose(inAudioConverter), nil
}

// AudioConverterDispose disposes of an audio converter object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterDispose(_:)
func AudioConverterDispose(inAudioConverter AudioConverterRef) int32 {
	result, callErr := tryAudioConverterDispose(inAudioConverter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterFillComplexBuffer func(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription) int32
var _audioConverterFillComplexBufferErr error

func tryAudioConverterFillComplexBuffer(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription) (int32, error) {
	if _audioConverterFillComplexBuffer == nil {
		return 0, symbolCallError("AudioConverterFillComplexBuffer", "10.2", _audioConverterFillComplexBufferErr)
	}
	return _audioConverterFillComplexBuffer(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription), nil
}

// AudioConverterFillComplexBuffer converts audio data supplied by a callback function, supporting non-interleaved and packetized formats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBuffer(_:_:_:_:_:_:)
func AudioConverterFillComplexBuffer(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription) int32 {
	result, callErr := tryAudioConverterFillComplexBuffer(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterFillComplexBufferRealtimeSafe func(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProcRealtimeSafe, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription) int32
var _audioConverterFillComplexBufferRealtimeSafeErr error

func tryAudioConverterFillComplexBufferRealtimeSafe(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProcRealtimeSafe, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription) (int32, error) {
	if _audioConverterFillComplexBufferRealtimeSafe == nil {
		return 0, symbolCallError("AudioConverterFillComplexBufferRealtimeSafe", "26.0", _audioConverterFillComplexBufferRealtimeSafeErr)
	}
	return _audioConverterFillComplexBufferRealtimeSafe(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription), nil
}

// AudioConverterFillComplexBufferRealtimeSafe.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBufferRealtimeSafe(_:_:_:_:_:_:)
func AudioConverterFillComplexBufferRealtimeSafe(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProcRealtimeSafe, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescription *coreaudiotypes.AudioStreamPacketDescription) int32 {
	result, callErr := tryAudioConverterFillComplexBufferRealtimeSafe(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescription)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterFillComplexBufferWithPacketDependencies func(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, outPacketDependencies *coreaudiotypes.AudioStreamPacketDependencyDescription) int32
var _audioConverterFillComplexBufferWithPacketDependenciesErr error

func tryAudioConverterFillComplexBufferWithPacketDependencies(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, outPacketDependencies *coreaudiotypes.AudioStreamPacketDependencyDescription) (int32, error) {
	if _audioConverterFillComplexBufferWithPacketDependencies == nil {
		return 0, symbolCallError("AudioConverterFillComplexBufferWithPacketDependencies", "26.0", _audioConverterFillComplexBufferWithPacketDependenciesErr)
	}
	return _audioConverterFillComplexBufferWithPacketDependencies(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescriptions, outPacketDependencies), nil
}

// AudioConverterFillComplexBufferWithPacketDependencies.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterFillComplexBufferWithPacketDependencies(_:_:_:_:_:_:_:)
func AudioConverterFillComplexBufferWithPacketDependencies(inAudioConverter AudioConverterRef, inInputDataProc AudioConverterComplexInputDataProc, inInputDataProcUserData unsafe.Pointer, ioOutputDataPacketSize *uint32, outOutputData *coreaudiotypes.AudioBufferList, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, outPacketDependencies *coreaudiotypes.AudioStreamPacketDependencyDescription) int32 {
	result, callErr := tryAudioConverterFillComplexBufferWithPacketDependencies(inAudioConverter, inInputDataProc, inInputDataProcUserData, ioOutputDataPacketSize, outOutputData, outPacketDescriptions, outPacketDependencies)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterGetProperty func(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioConverterGetPropertyErr error

func tryAudioConverterGetProperty(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioConverterGetProperty == nil {
		return 0, symbolCallError("AudioConverterGetProperty", "10.1", _audioConverterGetPropertyErr)
	}
	return _audioConverterGetProperty(inAudioConverter, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioConverterGetProperty gets an audio converter property value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterGetProperty(_:_:_:_:)
func AudioConverterGetProperty(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioConverterGetProperty(inAudioConverter, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterGetPropertyInfo func(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, outSize *uint32, outWritable *bool) int32
var _audioConverterGetPropertyInfoErr error

func tryAudioConverterGetPropertyInfo(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, outSize *uint32, outWritable *bool) (int32, error) {
	if _audioConverterGetPropertyInfo == nil {
		return 0, symbolCallError("AudioConverterGetPropertyInfo", "10.1", _audioConverterGetPropertyInfoErr)
	}
	return _audioConverterGetPropertyInfo(inAudioConverter, inPropertyID, outSize, outWritable), nil
}

// AudioConverterGetPropertyInfo gets information about an audio converter property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterGetPropertyInfo(_:_:_:_:)
func AudioConverterGetPropertyInfo(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioConverterGetPropertyInfo(inAudioConverter, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterNew func(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, outAudioConverter *AudioConverterRef) int32
var _audioConverterNewErr error

func tryAudioConverterNew(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, outAudioConverter *AudioConverterRef) (int32, error) {
	if _audioConverterNew == nil {
		return 0, symbolCallError("AudioConverterNew", "10.1", _audioConverterNewErr)
	}
	return _audioConverterNew(inSourceFormat, inDestinationFormat, outAudioConverter), nil
}

// AudioConverterNew creates a new audio converter object based on specified audio formats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNew(_:_:_:)
func AudioConverterNew(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, outAudioConverter *AudioConverterRef) int32 {
	result, callErr := tryAudioConverterNew(inSourceFormat, inDestinationFormat, outAudioConverter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterNewSpecific func(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, inNumberClassDescriptions uint32, inClassDescriptions *coreaudiotypes.AudioClassDescription, outAudioConverter *AudioConverterRef) int32
var _audioConverterNewSpecificErr error

func tryAudioConverterNewSpecific(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, inNumberClassDescriptions uint32, inClassDescriptions *coreaudiotypes.AudioClassDescription, outAudioConverter *AudioConverterRef) (int32, error) {
	if _audioConverterNewSpecific == nil {
		return 0, symbolCallError("AudioConverterNewSpecific", "10.4", _audioConverterNewSpecificErr)
	}
	return _audioConverterNewSpecific(inSourceFormat, inDestinationFormat, inNumberClassDescriptions, inClassDescriptions, outAudioConverter), nil
}

// AudioConverterNewSpecific creates a new audio converter object using a specified codec.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNewSpecific(_:_:_:_:_:)
func AudioConverterNewSpecific(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, inNumberClassDescriptions uint32, inClassDescriptions *coreaudiotypes.AudioClassDescription, outAudioConverter *AudioConverterRef) int32 {
	result, callErr := tryAudioConverterNewSpecific(inSourceFormat, inDestinationFormat, inNumberClassDescriptions, inClassDescriptions, outAudioConverter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterNewWithOptions func(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, inOptions AudioConverterOptions, outAudioConverter *AudioConverterRef) int32
var _audioConverterNewWithOptionsErr error

func tryAudioConverterNewWithOptions(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, inOptions AudioConverterOptions, outAudioConverter *AudioConverterRef) (int32, error) {
	if _audioConverterNewWithOptions == nil {
		return 0, symbolCallError("AudioConverterNewWithOptions", "15.0", _audioConverterNewWithOptionsErr)
	}
	return _audioConverterNewWithOptions(inSourceFormat, inDestinationFormat, inOptions, outAudioConverter), nil
}

// AudioConverterNewWithOptions.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterNewWithOptions(_:_:_:_:)
func AudioConverterNewWithOptions(inSourceFormat *coreaudiotypes.AudioStreamBasicDescription, inDestinationFormat *coreaudiotypes.AudioStreamBasicDescription, inOptions AudioConverterOptions, outAudioConverter *AudioConverterRef) int32 {
	result, callErr := tryAudioConverterNewWithOptions(inSourceFormat, inDestinationFormat, inOptions, outAudioConverter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterPrepare func(inFlags uint32, ioReserved unsafe.Pointer)
var _audioConverterPrepareErr error

func tryAudioConverterPrepare(inFlags uint32, ioReserved unsafe.Pointer) error {
	if _audioConverterPrepare == nil {
		return symbolCallError("AudioConverterPrepare", "15.0", _audioConverterPrepareErr)
	}
	_audioConverterPrepare(inFlags, ioReserved)
	return nil
}

// AudioConverterPrepare.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterPrepare(_:_:_:)
func AudioConverterPrepare(inFlags uint32, ioReserved unsafe.Pointer) {
	if callErr := tryAudioConverterPrepare(inFlags, ioReserved); callErr != nil {
		panic(callErr)
	}
}

var _audioConverterReset func(inAudioConverter AudioConverterRef) int32
var _audioConverterResetErr error

func tryAudioConverterReset(inAudioConverter AudioConverterRef) (int32, error) {
	if _audioConverterReset == nil {
		return 0, symbolCallError("AudioConverterReset", "10.1", _audioConverterResetErr)
	}
	return _audioConverterReset(inAudioConverter), nil
}

// AudioConverterReset resets an audio converter object, clearing and flushing its buffers.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterReset(_:)
func AudioConverterReset(inAudioConverter AudioConverterRef) int32 {
	result, callErr := tryAudioConverterReset(inAudioConverter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConverterSetProperty func(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioConverterSetPropertyErr error

func tryAudioConverterSetProperty(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioConverterSetProperty == nil {
		return 0, symbolCallError("AudioConverterSetProperty", "10.1", _audioConverterSetPropertyErr)
	}
	return _audioConverterSetProperty(inAudioConverter, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioConverterSetProperty sets the value of an audio converter object property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioConverterSetProperty(_:_:_:_:)
func AudioConverterSetProperty(inAudioConverter AudioConverterRef, inPropertyID AudioConverterPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioConverterSetProperty(inAudioConverter, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileClose func(inAudioFile AudioFileID) int32
var _audioFileCloseErr error

func tryAudioFileClose(inAudioFile AudioFileID) (int32, error) {
	if _audioFileClose == nil {
		return 0, symbolCallError("AudioFileClose", "10.2", _audioFileCloseErr)
	}
	return _audioFileClose(inAudioFile), nil
}

// AudioFileClose closes an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileClose(_:)
func AudioFileClose(inAudioFile AudioFileID) int32 {
	result, callErr := tryAudioFileClose(inAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentCloseFile func(inComponent AudioFileComponent) int32
var _audioFileComponentCloseFileErr error

func tryAudioFileComponentCloseFile(inComponent AudioFileComponent) (int32, error) {
	if _audioFileComponentCloseFile == nil {
		return 0, symbolCallError("AudioFileComponentCloseFile", "10.4", _audioFileComponentCloseFileErr)
	}
	return _audioFileComponentCloseFile(inComponent), nil
}

// AudioFileComponentCloseFile.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCloseFile(_:)
func AudioFileComponentCloseFile(inComponent AudioFileComponent) int32 {
	result, callErr := tryAudioFileComponentCloseFile(inComponent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentCountUserData func(inComponent AudioFileComponent, inUserDataID uint32, outNumberItems *uint32) int32
var _audioFileComponentCountUserDataErr error

func tryAudioFileComponentCountUserData(inComponent AudioFileComponent, inUserDataID uint32, outNumberItems *uint32) (int32, error) {
	if _audioFileComponentCountUserData == nil {
		return 0, symbolCallError("AudioFileComponentCountUserData", "10.4", _audioFileComponentCountUserDataErr)
	}
	return _audioFileComponentCountUserData(inComponent, inUserDataID, outNumberItems), nil
}

// AudioFileComponentCountUserData.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCountUserData(_:_:_:)
func AudioFileComponentCountUserData(inComponent AudioFileComponent, inUserDataID uint32, outNumberItems *uint32) int32 {
	result, callErr := tryAudioFileComponentCountUserData(inComponent, inUserDataID, outNumberItems)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentCreate func(inComponent AudioFileComponent, inParentRef *coreservices.FSRef, inFileName corefoundation.CFStringRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, outNewFileRef *coreservices.FSRef) int32
var _audioFileComponentCreateErr error

func tryAudioFileComponentCreate(inComponent AudioFileComponent, inParentRef *coreservices.FSRef, inFileName corefoundation.CFStringRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, outNewFileRef *coreservices.FSRef) (int32, error) {
	if _audioFileComponentCreate == nil {
		return 0, symbolCallError("AudioFileComponentCreate", "10.4", _audioFileComponentCreateErr)
	}
	return _audioFileComponentCreate(inComponent, inParentRef, inFileName, inFormat, inFlags, outNewFileRef), nil
}

// AudioFileComponentCreate.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreate
func AudioFileComponentCreate(inComponent AudioFileComponent, inParentRef *coreservices.FSRef, inFileName corefoundation.CFStringRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, outNewFileRef *coreservices.FSRef) int32 {
	result, callErr := tryAudioFileComponentCreate(inComponent, inParentRef, inFileName, inFormat, inFlags, outNewFileRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentCreateURL func(inComponent AudioFileComponent, inFileRef corefoundation.CFURLRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) int32
var _audioFileComponentCreateURLErr error

func tryAudioFileComponentCreateURL(inComponent AudioFileComponent, inFileRef corefoundation.CFURLRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) (int32, error) {
	if _audioFileComponentCreateURL == nil {
		return 0, symbolCallError("AudioFileComponentCreateURL", "10.5", _audioFileComponentCreateURLErr)
	}
	return _audioFileComponentCreateURL(inComponent, inFileRef, inFormat, inFlags), nil
}

// AudioFileComponentCreateURL.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentCreateURL(_:_:_:_:)
func AudioFileComponentCreateURL(inComponent AudioFileComponent, inFileRef corefoundation.CFURLRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) int32 {
	result, callErr := tryAudioFileComponentCreateURL(inComponent, inFileRef, inFormat, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentDataIsThisFormat func(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, outResult *uint32) int32
var _audioFileComponentDataIsThisFormatErr error

func tryAudioFileComponentDataIsThisFormat(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, outResult *uint32) (int32, error) {
	if _audioFileComponentDataIsThisFormat == nil {
		return 0, symbolCallError("AudioFileComponentDataIsThisFormat", "10.4", _audioFileComponentDataIsThisFormatErr)
	}
	return _audioFileComponentDataIsThisFormat(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, outResult), nil
}

// AudioFileComponentDataIsThisFormat.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentDataIsThisFormat
func AudioFileComponentDataIsThisFormat(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, outResult *uint32) int32 {
	result, callErr := tryAudioFileComponentDataIsThisFormat(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentExtensionIsThisFormat func(inComponent AudioFileComponent, inExtension corefoundation.CFStringRef, outResult *uint32) int32
var _audioFileComponentExtensionIsThisFormatErr error

func tryAudioFileComponentExtensionIsThisFormat(inComponent AudioFileComponent, inExtension corefoundation.CFStringRef, outResult *uint32) (int32, error) {
	if _audioFileComponentExtensionIsThisFormat == nil {
		return 0, symbolCallError("AudioFileComponentExtensionIsThisFormat", "10.4", _audioFileComponentExtensionIsThisFormatErr)
	}
	return _audioFileComponentExtensionIsThisFormat(inComponent, inExtension, outResult), nil
}

// AudioFileComponentExtensionIsThisFormat.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentExtensionIsThisFormat(_:_:_:)
func AudioFileComponentExtensionIsThisFormat(inComponent AudioFileComponent, inExtension corefoundation.CFStringRef, outResult *uint32) int32 {
	result, callErr := tryAudioFileComponentExtensionIsThisFormat(inComponent, inExtension, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentFileDataIsThisFormat func(inComponent AudioFileComponent, inDataByteSize uint32, inData unsafe.Pointer, outResult *uint32) int32
var _audioFileComponentFileDataIsThisFormatErr error

func tryAudioFileComponentFileDataIsThisFormat(inComponent AudioFileComponent, inDataByteSize uint32, inData unsafe.Pointer, outResult *uint32) (int32, error) {
	if _audioFileComponentFileDataIsThisFormat == nil {
		return 0, symbolCallError("AudioFileComponentFileDataIsThisFormat", "10.4", _audioFileComponentFileDataIsThisFormatErr)
	}
	return _audioFileComponentFileDataIsThisFormat(inComponent, inDataByteSize, inData, outResult), nil
}

// AudioFileComponentFileDataIsThisFormat.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileDataIsThisFormat(_:_:_:_:)
func AudioFileComponentFileDataIsThisFormat(inComponent AudioFileComponent, inDataByteSize uint32, inData unsafe.Pointer, outResult *uint32) int32 {
	result, callErr := tryAudioFileComponentFileDataIsThisFormat(inComponent, inDataByteSize, inData, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentFileIsThisFormat func(inComponent AudioFileComponent, inFileRefNum int16, outResult *uint32) int32
var _audioFileComponentFileIsThisFormatErr error

func tryAudioFileComponentFileIsThisFormat(inComponent AudioFileComponent, inFileRefNum int16, outResult *uint32) (int32, error) {
	if _audioFileComponentFileIsThisFormat == nil {
		return 0, symbolCallError("AudioFileComponentFileIsThisFormat", "10.4", _audioFileComponentFileIsThisFormatErr)
	}
	return _audioFileComponentFileIsThisFormat(inComponent, inFileRefNum, outResult), nil
}

// AudioFileComponentFileIsThisFormat.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentFileIsThisFormat
func AudioFileComponentFileIsThisFormat(inComponent AudioFileComponent, inFileRefNum int16, outResult *uint32) int32 {
	result, callErr := tryAudioFileComponentFileIsThisFormat(inComponent, inFileRefNum, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetGlobalInfo func(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioFileComponentGetGlobalInfoErr error

func tryAudioFileComponentGetGlobalInfo(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileComponentGetGlobalInfo == nil {
		return 0, symbolCallError("AudioFileComponentGetGlobalInfo", "10.4", _audioFileComponentGetGlobalInfoErr)
	}
	return _audioFileComponentGetGlobalInfo(inComponent, inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData), nil
}

// AudioFileComponentGetGlobalInfo.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfo(_:_:_:_:_:_:)
func AudioFileComponentGetGlobalInfo(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentGetGlobalInfo(inComponent, inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetGlobalInfoSize func(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertySize *uint32) int32
var _audioFileComponentGetGlobalInfoSizeErr error

func tryAudioFileComponentGetGlobalInfoSize(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertySize *uint32) (int32, error) {
	if _audioFileComponentGetGlobalInfoSize == nil {
		return 0, symbolCallError("AudioFileComponentGetGlobalInfoSize", "10.4", _audioFileComponentGetGlobalInfoSizeErr)
	}
	return _audioFileComponentGetGlobalInfoSize(inComponent, inPropertyID, inSpecifierSize, inSpecifier, outPropertySize), nil
}

// AudioFileComponentGetGlobalInfoSize.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetGlobalInfoSize(_:_:_:_:_:)
func AudioFileComponentGetGlobalInfoSize(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertySize *uint32) int32 {
	result, callErr := tryAudioFileComponentGetGlobalInfoSize(inComponent, inPropertyID, inSpecifierSize, inSpecifier, outPropertySize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetProperty func(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioFileComponentGetPropertyErr error

func tryAudioFileComponentGetProperty(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileComponentGetProperty == nil {
		return 0, symbolCallError("AudioFileComponentGetProperty", "10.4", _audioFileComponentGetPropertyErr)
	}
	return _audioFileComponentGetProperty(inComponent, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioFileComponentGetProperty.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetProperty(_:_:_:_:)
func AudioFileComponentGetProperty(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentGetProperty(inComponent, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetPropertyInfo func(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, outPropertySize *uint32, outWritable *uint32) int32
var _audioFileComponentGetPropertyInfoErr error

func tryAudioFileComponentGetPropertyInfo(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, outPropertySize *uint32, outWritable *uint32) (int32, error) {
	if _audioFileComponentGetPropertyInfo == nil {
		return 0, symbolCallError("AudioFileComponentGetPropertyInfo", "10.4", _audioFileComponentGetPropertyInfoErr)
	}
	return _audioFileComponentGetPropertyInfo(inComponent, inPropertyID, outPropertySize, outWritable), nil
}

// AudioFileComponentGetPropertyInfo.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetPropertyInfo(_:_:_:_:)
func AudioFileComponentGetPropertyInfo(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, outPropertySize *uint32, outWritable *uint32) int32 {
	result, callErr := tryAudioFileComponentGetPropertyInfo(inComponent, inPropertyID, outPropertySize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetUserData func(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32
var _audioFileComponentGetUserDataErr error

func tryAudioFileComponentGetUserData(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, ioUserDataSize *uint32, outUserData unsafe.Pointer) (int32, error) {
	if _audioFileComponentGetUserData == nil {
		return 0, symbolCallError("AudioFileComponentGetUserData", "10.4", _audioFileComponentGetUserDataErr)
	}
	return _audioFileComponentGetUserData(inComponent, inUserDataID, inIndex, ioUserDataSize, outUserData), nil
}

// AudioFileComponentGetUserData.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserData(_:_:_:_:_:)
func AudioFileComponentGetUserData(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentGetUserData(inComponent, inUserDataID, inIndex, ioUserDataSize, outUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetUserDataAtOffset func(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, inOffset int64, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32
var _audioFileComponentGetUserDataAtOffsetErr error

func tryAudioFileComponentGetUserDataAtOffset(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, inOffset int64, ioUserDataSize *uint32, outUserData unsafe.Pointer) (int32, error) {
	if _audioFileComponentGetUserDataAtOffset == nil {
		return 0, symbolCallError("AudioFileComponentGetUserDataAtOffset", "14.0", _audioFileComponentGetUserDataAtOffsetErr)
	}
	return _audioFileComponentGetUserDataAtOffset(inComponent, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData), nil
}

// AudioFileComponentGetUserDataAtOffset.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataAtOffset(_:_:_:_:_:_:)
func AudioFileComponentGetUserDataAtOffset(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, inOffset int64, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentGetUserDataAtOffset(inComponent, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetUserDataSize func(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, outUserDataSize *uint32) int32
var _audioFileComponentGetUserDataSizeErr error

func tryAudioFileComponentGetUserDataSize(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, outUserDataSize *uint32) (int32, error) {
	if _audioFileComponentGetUserDataSize == nil {
		return 0, symbolCallError("AudioFileComponentGetUserDataSize", "10.4", _audioFileComponentGetUserDataSizeErr)
	}
	return _audioFileComponentGetUserDataSize(inComponent, inUserDataID, inIndex, outUserDataSize), nil
}

// AudioFileComponentGetUserDataSize.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize(_:_:_:_:)
func AudioFileComponentGetUserDataSize(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, outUserDataSize *uint32) int32 {
	result, callErr := tryAudioFileComponentGetUserDataSize(inComponent, inUserDataID, inIndex, outUserDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentGetUserDataSize64 func(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, outUserDataSize *uint64) int32
var _audioFileComponentGetUserDataSize64Err error

func tryAudioFileComponentGetUserDataSize64(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, outUserDataSize *uint64) (int32, error) {
	if _audioFileComponentGetUserDataSize64 == nil {
		return 0, symbolCallError("AudioFileComponentGetUserDataSize64", "14.0", _audioFileComponentGetUserDataSize64Err)
	}
	return _audioFileComponentGetUserDataSize64(inComponent, inUserDataID, inIndex, outUserDataSize), nil
}

// AudioFileComponentGetUserDataSize64.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentGetUserDataSize64(_:_:_:_:)
func AudioFileComponentGetUserDataSize64(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, outUserDataSize *uint64) int32 {
	result, callErr := tryAudioFileComponentGetUserDataSize64(inComponent, inUserDataID, inIndex, outUserDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentInitialize func(inComponent AudioFileComponent, inFileRef *coreservices.FSRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) int32
var _audioFileComponentInitializeErr error

func tryAudioFileComponentInitialize(inComponent AudioFileComponent, inFileRef *coreservices.FSRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) (int32, error) {
	if _audioFileComponentInitialize == nil {
		return 0, symbolCallError("AudioFileComponentInitialize", "10.4", _audioFileComponentInitializeErr)
	}
	return _audioFileComponentInitialize(inComponent, inFileRef, inFormat, inFlags), nil
}

// AudioFileComponentInitialize.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitialize
func AudioFileComponentInitialize(inComponent AudioFileComponent, inFileRef *coreservices.FSRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) int32 {
	result, callErr := tryAudioFileComponentInitialize(inComponent, inFileRef, inFormat, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentInitializeWithCallbacks func(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType uint32, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) int32
var _audioFileComponentInitializeWithCallbacksErr error

func tryAudioFileComponentInitializeWithCallbacks(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType uint32, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) (int32, error) {
	if _audioFileComponentInitializeWithCallbacks == nil {
		return 0, symbolCallError("AudioFileComponentInitializeWithCallbacks", "10.4", _audioFileComponentInitializeWithCallbacksErr)
	}
	return _audioFileComponentInitializeWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags), nil
}

// AudioFileComponentInitializeWithCallbacks.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentInitializeWithCallbacks(_:_:_:_:_:_:_:_:_:)
func AudioFileComponentInitializeWithCallbacks(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType uint32, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32) int32 {
	result, callErr := tryAudioFileComponentInitializeWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentOpenFile func(inComponent AudioFileComponent, inFileRef *coreservices.FSRef, inPermissions int8, inRefNum int16) int32
var _audioFileComponentOpenFileErr error

func tryAudioFileComponentOpenFile(inComponent AudioFileComponent, inFileRef *coreservices.FSRef, inPermissions int8, inRefNum int16) (int32, error) {
	if _audioFileComponentOpenFile == nil {
		return 0, symbolCallError("AudioFileComponentOpenFile", "10.4", _audioFileComponentOpenFileErr)
	}
	return _audioFileComponentOpenFile(inComponent, inFileRef, inPermissions, inRefNum), nil
}

// AudioFileComponentOpenFile.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenFile
func AudioFileComponentOpenFile(inComponent AudioFileComponent, inFileRef *coreservices.FSRef, inPermissions int8, inRefNum int16) int32 {
	result, callErr := tryAudioFileComponentOpenFile(inComponent, inFileRef, inPermissions, inRefNum)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentOpenURL func(inComponent AudioFileComponent, inFileRef corefoundation.CFURLRef, inPermissions int8, inFileDescriptor int32) int32
var _audioFileComponentOpenURLErr error

func tryAudioFileComponentOpenURL(inComponent AudioFileComponent, inFileRef corefoundation.CFURLRef, inPermissions int8, inFileDescriptor int32) (int32, error) {
	if _audioFileComponentOpenURL == nil {
		return 0, symbolCallError("AudioFileComponentOpenURL", "10.5", _audioFileComponentOpenURLErr)
	}
	return _audioFileComponentOpenURL(inComponent, inFileRef, inPermissions, inFileDescriptor), nil
}

// AudioFileComponentOpenURL.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenURL(_:_:_:_:)
func AudioFileComponentOpenURL(inComponent AudioFileComponent, inFileRef corefoundation.CFURLRef, inPermissions int8, inFileDescriptor int32) int32 {
	result, callErr := tryAudioFileComponentOpenURL(inComponent, inFileRef, inPermissions, inFileDescriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentOpenWithCallbacks func(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc) int32
var _audioFileComponentOpenWithCallbacksErr error

func tryAudioFileComponentOpenWithCallbacks(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc) (int32, error) {
	if _audioFileComponentOpenWithCallbacks == nil {
		return 0, symbolCallError("AudioFileComponentOpenWithCallbacks", "10.4", _audioFileComponentOpenWithCallbacksErr)
	}
	return _audioFileComponentOpenWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc), nil
}

// AudioFileComponentOpenWithCallbacks.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOpenWithCallbacks(_:_:_:_:_:_:)
func AudioFileComponentOpenWithCallbacks(inComponent AudioFileComponent, inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc) int32 {
	result, callErr := tryAudioFileComponentOpenWithCallbacks(inComponent, inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentOptimize func(inComponent AudioFileComponent) int32
var _audioFileComponentOptimizeErr error

func tryAudioFileComponentOptimize(inComponent AudioFileComponent) (int32, error) {
	if _audioFileComponentOptimize == nil {
		return 0, symbolCallError("AudioFileComponentOptimize", "10.4", _audioFileComponentOptimizeErr)
	}
	return _audioFileComponentOptimize(inComponent), nil
}

// AudioFileComponentOptimize.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentOptimize(_:)
func AudioFileComponentOptimize(inComponent AudioFileComponent) int32 {
	result, callErr := tryAudioFileComponentOptimize(inComponent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentReadBytes func(inComponent AudioFileComponent, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, outBuffer unsafe.Pointer) int32
var _audioFileComponentReadBytesErr error

func tryAudioFileComponentReadBytes(inComponent AudioFileComponent, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, outBuffer unsafe.Pointer) (int32, error) {
	if _audioFileComponentReadBytes == nil {
		return 0, symbolCallError("AudioFileComponentReadBytes", "10.4", _audioFileComponentReadBytesErr)
	}
	return _audioFileComponentReadBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, outBuffer), nil
}

// AudioFileComponentReadBytes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadBytes(_:_:_:_:_:)
func AudioFileComponentReadBytes(inComponent AudioFileComponent, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, outBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentReadBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, outBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentReadPacketData func(inComponent AudioFileComponent, inUseCache bool, ioNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) int32
var _audioFileComponentReadPacketDataErr error

func tryAudioFileComponentReadPacketData(inComponent AudioFileComponent, inUseCache bool, ioNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) (int32, error) {
	if _audioFileComponentReadPacketData == nil {
		return 0, symbolCallError("AudioFileComponentReadPacketData", "10.4", _audioFileComponentReadPacketDataErr)
	}
	return _audioFileComponentReadPacketData(inComponent, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer), nil
}

// AudioFileComponentReadPacketData.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPacketData(_:_:_:_:_:_:_:)
func AudioFileComponentReadPacketData(inComponent AudioFileComponent, inUseCache bool, ioNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentReadPacketData(inComponent, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentReadPackets func(inComponent AudioFileComponent, inUseCache bool, outNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) int32
var _audioFileComponentReadPacketsErr error

func tryAudioFileComponentReadPackets(inComponent AudioFileComponent, inUseCache bool, outNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) (int32, error) {
	if _audioFileComponentReadPackets == nil {
		return 0, symbolCallError("AudioFileComponentReadPackets", "10.4", _audioFileComponentReadPacketsErr)
	}
	return _audioFileComponentReadPackets(inComponent, inUseCache, outNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer), nil
}

// AudioFileComponentReadPackets.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentReadPackets(_:_:_:_:_:_:_:)
func AudioFileComponentReadPackets(inComponent AudioFileComponent, inUseCache bool, outNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentReadPackets(inComponent, inUseCache, outNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentRemoveUserData func(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32) int32
var _audioFileComponentRemoveUserDataErr error

func tryAudioFileComponentRemoveUserData(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32) (int32, error) {
	if _audioFileComponentRemoveUserData == nil {
		return 0, symbolCallError("AudioFileComponentRemoveUserData", "10.5", _audioFileComponentRemoveUserDataErr)
	}
	return _audioFileComponentRemoveUserData(inComponent, inUserDataID, inIndex), nil
}

// AudioFileComponentRemoveUserData.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentRemoveUserData(_:_:_:)
func AudioFileComponentRemoveUserData(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32) int32 {
	result, callErr := tryAudioFileComponentRemoveUserData(inComponent, inUserDataID, inIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentSetProperty func(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioFileComponentSetPropertyErr error

func tryAudioFileComponentSetProperty(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileComponentSetProperty == nil {
		return 0, symbolCallError("AudioFileComponentSetProperty", "10.4", _audioFileComponentSetPropertyErr)
	}
	return _audioFileComponentSetProperty(inComponent, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioFileComponentSetProperty.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetProperty(_:_:_:_:)
func AudioFileComponentSetProperty(inComponent AudioFileComponent, inPropertyID AudioFileComponentPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentSetProperty(inComponent, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentSetUserData func(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, inUserDataSize uint32, inUserData unsafe.Pointer) int32
var _audioFileComponentSetUserDataErr error

func tryAudioFileComponentSetUserData(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, inUserDataSize uint32, inUserData unsafe.Pointer) (int32, error) {
	if _audioFileComponentSetUserData == nil {
		return 0, symbolCallError("AudioFileComponentSetUserData", "10.4", _audioFileComponentSetUserDataErr)
	}
	return _audioFileComponentSetUserData(inComponent, inUserDataID, inIndex, inUserDataSize, inUserData), nil
}

// AudioFileComponentSetUserData.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentSetUserData(_:_:_:_:_:)
func AudioFileComponentSetUserData(inComponent AudioFileComponent, inUserDataID uint32, inIndex uint32, inUserDataSize uint32, inUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentSetUserData(inComponent, inUserDataID, inIndex, inUserDataSize, inUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentWriteBytes func(inComponent AudioFileComponent, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, inBuffer unsafe.Pointer) int32
var _audioFileComponentWriteBytesErr error

func tryAudioFileComponentWriteBytes(inComponent AudioFileComponent, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, inBuffer unsafe.Pointer) (int32, error) {
	if _audioFileComponentWriteBytes == nil {
		return 0, symbolCallError("AudioFileComponentWriteBytes", "10.4", _audioFileComponentWriteBytesErr)
	}
	return _audioFileComponentWriteBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, inBuffer), nil
}

// AudioFileComponentWriteBytes.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWriteBytes(_:_:_:_:_:)
func AudioFileComponentWriteBytes(inComponent AudioFileComponent, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, inBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentWriteBytes(inComponent, inUseCache, inStartingByte, ioNumBytes, inBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileComponentWritePackets func(inComponent AudioFileComponent, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) int32
var _audioFileComponentWritePacketsErr error

func tryAudioFileComponentWritePackets(inComponent AudioFileComponent, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) (int32, error) {
	if _audioFileComponentWritePackets == nil {
		return 0, symbolCallError("AudioFileComponentWritePackets", "10.4", _audioFileComponentWritePacketsErr)
	}
	return _audioFileComponentWritePackets(inComponent, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer), nil
}

// AudioFileComponentWritePackets.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileComponentWritePackets(_:_:_:_:_:_:_:)
func AudioFileComponentWritePackets(inComponent AudioFileComponent, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileComponentWritePackets(inComponent, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileCountUserData func(inAudioFile AudioFileID, inUserDataID uint32, outNumberItems *uint32) int32
var _audioFileCountUserDataErr error

func tryAudioFileCountUserData(inAudioFile AudioFileID, inUserDataID uint32, outNumberItems *uint32) (int32, error) {
	if _audioFileCountUserData == nil {
		return 0, symbolCallError("AudioFileCountUserData", "10.4", _audioFileCountUserDataErr)
	}
	return _audioFileCountUserData(inAudioFile, inUserDataID, outNumberItems), nil
}

// AudioFileCountUserData gets the number of user data items with a specified ID in a file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileCountUserData(_:_:_:)
func AudioFileCountUserData(inAudioFile AudioFileID, inUserDataID uint32, outNumberItems *uint32) int32 {
	result, callErr := tryAudioFileCountUserData(inAudioFile, inUserDataID, outNumberItems)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileCreateWithURL func(inFileRef corefoundation.CFURLRef, inFileType AudioFileTypeID, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags AudioFileFlags, outAudioFile *AudioFileID) int32
var _audioFileCreateWithURLErr error

func tryAudioFileCreateWithURL(inFileRef corefoundation.CFURLRef, inFileType AudioFileTypeID, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags AudioFileFlags, outAudioFile *AudioFileID) (int32, error) {
	if _audioFileCreateWithURL == nil {
		return 0, symbolCallError("AudioFileCreateWithURL", "10.5", _audioFileCreateWithURLErr)
	}
	return _audioFileCreateWithURL(inFileRef, inFileType, inFormat, inFlags, outAudioFile), nil
}

// AudioFileCreateWithURL creates a new audio file, or initializes an existing file, specified by a URL.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileCreateWithURL(_:_:_:_:_:)
func AudioFileCreateWithURL(inFileRef corefoundation.CFURLRef, inFileType AudioFileTypeID, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags AudioFileFlags, outAudioFile *AudioFileID) int32 {
	result, callErr := tryAudioFileCreateWithURL(inFileRef, inFileType, inFormat, inFlags, outAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetGlobalInfo func(inPropertyID AudioFilePropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioFileGetGlobalInfoErr error

func tryAudioFileGetGlobalInfo(inPropertyID AudioFilePropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileGetGlobalInfo == nil {
		return 0, symbolCallError("AudioFileGetGlobalInfo", "10.3", _audioFileGetGlobalInfoErr)
	}
	return _audioFileGetGlobalInfo(inPropertyID, inSpecifierSize, inSpecifier, ioDataSize, outPropertyData), nil
}

// AudioFileGetGlobalInfo copies the value of a global property into a buffer.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetGlobalInfo(_:_:_:_:_:)
func AudioFileGetGlobalInfo(inPropertyID AudioFilePropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileGetGlobalInfo(inPropertyID, inSpecifierSize, inSpecifier, ioDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetGlobalInfoSize func(inPropertyID AudioFilePropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outDataSize *uint32) int32
var _audioFileGetGlobalInfoSizeErr error

func tryAudioFileGetGlobalInfoSize(inPropertyID AudioFilePropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outDataSize *uint32) (int32, error) {
	if _audioFileGetGlobalInfoSize == nil {
		return 0, symbolCallError("AudioFileGetGlobalInfoSize", "10.3", _audioFileGetGlobalInfoSizeErr)
	}
	return _audioFileGetGlobalInfoSize(inPropertyID, inSpecifierSize, inSpecifier, outDataSize), nil
}

// AudioFileGetGlobalInfoSize gets the size of a global audio file property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetGlobalInfoSize(_:_:_:_:)
func AudioFileGetGlobalInfoSize(inPropertyID AudioFilePropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outDataSize *uint32) int32 {
	result, callErr := tryAudioFileGetGlobalInfoSize(inPropertyID, inSpecifierSize, inSpecifier, outDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetProperty func(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, ioDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioFileGetPropertyErr error

func tryAudioFileGetProperty(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, ioDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileGetProperty == nil {
		return 0, symbolCallError("AudioFileGetProperty", "10.2", _audioFileGetPropertyErr)
	}
	return _audioFileGetProperty(inAudioFile, inPropertyID, ioDataSize, outPropertyData), nil
}

// AudioFileGetProperty gets the value of an audio file property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetProperty(_:_:_:_:)
func AudioFileGetProperty(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, ioDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileGetProperty(inAudioFile, inPropertyID, ioDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetPropertyInfo func(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, outDataSize *uint32, isWritable *uint32) int32
var _audioFileGetPropertyInfoErr error

func tryAudioFileGetPropertyInfo(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, outDataSize *uint32, isWritable *uint32) (int32, error) {
	if _audioFileGetPropertyInfo == nil {
		return 0, symbolCallError("AudioFileGetPropertyInfo", "10.2", _audioFileGetPropertyInfoErr)
	}
	return _audioFileGetPropertyInfo(inAudioFile, inPropertyID, outDataSize, isWritable), nil
}

// AudioFileGetPropertyInfo gets information about an audio file property, including the size of the property value and whether the value is writable.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetPropertyInfo(_:_:_:_:)
func AudioFileGetPropertyInfo(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, outDataSize *uint32, isWritable *uint32) int32 {
	result, callErr := tryAudioFileGetPropertyInfo(inAudioFile, inPropertyID, outDataSize, isWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetUserData func(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32
var _audioFileGetUserDataErr error

func tryAudioFileGetUserData(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, ioUserDataSize *uint32, outUserData unsafe.Pointer) (int32, error) {
	if _audioFileGetUserData == nil {
		return 0, symbolCallError("AudioFileGetUserData", "10.4", _audioFileGetUserDataErr)
	}
	return _audioFileGetUserData(inAudioFile, inUserDataID, inIndex, ioUserDataSize, outUserData), nil
}

// AudioFileGetUserData gets a chunk from an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserData(_:_:_:_:_:)
func AudioFileGetUserData(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileGetUserData(inAudioFile, inUserDataID, inIndex, ioUserDataSize, outUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetUserDataAtOffset func(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, inOffset int64, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32
var _audioFileGetUserDataAtOffsetErr error

func tryAudioFileGetUserDataAtOffset(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, inOffset int64, ioUserDataSize *uint32, outUserData unsafe.Pointer) (int32, error) {
	if _audioFileGetUserDataAtOffset == nil {
		return 0, symbolCallError("AudioFileGetUserDataAtOffset", "14.0", _audioFileGetUserDataAtOffsetErr)
	}
	return _audioFileGetUserDataAtOffset(inAudioFile, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData), nil
}

// AudioFileGetUserDataAtOffset gets part of the data from a chunk in an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataAtOffset(_:_:_:_:_:_:)
func AudioFileGetUserDataAtOffset(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, inOffset int64, ioUserDataSize *uint32, outUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileGetUserDataAtOffset(inAudioFile, inUserDataID, inIndex, inOffset, ioUserDataSize, outUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetUserDataSize func(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, outUserDataSize *uint32) int32
var _audioFileGetUserDataSizeErr error

func tryAudioFileGetUserDataSize(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, outUserDataSize *uint32) (int32, error) {
	if _audioFileGetUserDataSize == nil {
		return 0, symbolCallError("AudioFileGetUserDataSize", "10.4", _audioFileGetUserDataSizeErr)
	}
	return _audioFileGetUserDataSize(inAudioFile, inUserDataID, inIndex, outUserDataSize), nil
}

// AudioFileGetUserDataSize gets the size of a user data item in an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataSize(_:_:_:_:)
func AudioFileGetUserDataSize(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, outUserDataSize *uint32) int32 {
	result, callErr := tryAudioFileGetUserDataSize(inAudioFile, inUserDataID, inIndex, outUserDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileGetUserDataSize64 func(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, outUserDataSize *uint64) int32
var _audioFileGetUserDataSize64Err error

func tryAudioFileGetUserDataSize64(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, outUserDataSize *uint64) (int32, error) {
	if _audioFileGetUserDataSize64 == nil {
		return 0, symbolCallError("AudioFileGetUserDataSize64", "14.0", _audioFileGetUserDataSize64Err)
	}
	return _audioFileGetUserDataSize64(inAudioFile, inUserDataID, inIndex, outUserDataSize), nil
}

// AudioFileGetUserDataSize64 gets the size of a user data item in an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileGetUserDataSize64(_:_:_:_:)
func AudioFileGetUserDataSize64(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, outUserDataSize *uint64) int32 {
	result, callErr := tryAudioFileGetUserDataSize64(inAudioFile, inUserDataID, inIndex, outUserDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileInitializeWithCallbacks func(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType AudioFileTypeID, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags AudioFileFlags, outAudioFile *AudioFileID) int32
var _audioFileInitializeWithCallbacksErr error

func tryAudioFileInitializeWithCallbacks(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType AudioFileTypeID, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags AudioFileFlags, outAudioFile *AudioFileID) (int32, error) {
	if _audioFileInitializeWithCallbacks == nil {
		return 0, symbolCallError("AudioFileInitializeWithCallbacks", "10.3", _audioFileInitializeWithCallbacksErr)
	}
	return _audioFileInitializeWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags, outAudioFile), nil
}

// AudioFileInitializeWithCallbacks deletes the content of an existing file and assigns callbacks to the audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileInitializeWithCallbacks(_:_:_:_:_:_:_:_:_:)
func AudioFileInitializeWithCallbacks(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileType AudioFileTypeID, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags AudioFileFlags, outAudioFile *AudioFileID) int32 {
	result, callErr := tryAudioFileInitializeWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileType, inFormat, inFlags, outAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileOpenURL func(inFileRef corefoundation.CFURLRef, inPermissions AudioFilePermissions, inFileTypeHint AudioFileTypeID, outAudioFile *AudioFileID) int32
var _audioFileOpenURLErr error

func tryAudioFileOpenURL(inFileRef corefoundation.CFURLRef, inPermissions AudioFilePermissions, inFileTypeHint AudioFileTypeID, outAudioFile *AudioFileID) (int32, error) {
	if _audioFileOpenURL == nil {
		return 0, symbolCallError("AudioFileOpenURL", "10.5", _audioFileOpenURLErr)
	}
	return _audioFileOpenURL(inFileRef, inPermissions, inFileTypeHint, outAudioFile), nil
}

// AudioFileOpenURL open an existing audio file specified by a URL.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpenURL(_:_:_:_:)
func AudioFileOpenURL(inFileRef corefoundation.CFURLRef, inPermissions AudioFilePermissions, inFileTypeHint AudioFileTypeID, outAudioFile *AudioFileID) int32 {
	result, callErr := tryAudioFileOpenURL(inFileRef, inPermissions, inFileTypeHint, outAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileOpenWithCallbacks func(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileTypeHint AudioFileTypeID, outAudioFile *AudioFileID) int32
var _audioFileOpenWithCallbacksErr error

func tryAudioFileOpenWithCallbacks(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileTypeHint AudioFileTypeID, outAudioFile *AudioFileID) (int32, error) {
	if _audioFileOpenWithCallbacks == nil {
		return 0, symbolCallError("AudioFileOpenWithCallbacks", "10.3", _audioFileOpenWithCallbacksErr)
	}
	return _audioFileOpenWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileTypeHint, outAudioFile), nil
}

// AudioFileOpenWithCallbacks opens an existing file with callbacks you provide.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileOpenWithCallbacks(_:_:_:_:_:_:_:)
func AudioFileOpenWithCallbacks(inClientData unsafe.Pointer, inReadFunc AudioFile_ReadProc, inWriteFunc AudioFile_WriteProc, inGetSizeFunc AudioFile_GetSizeProc, inSetSizeFunc AudioFile_SetSizeProc, inFileTypeHint AudioFileTypeID, outAudioFile *AudioFileID) int32 {
	result, callErr := tryAudioFileOpenWithCallbacks(inClientData, inReadFunc, inWriteFunc, inGetSizeFunc, inSetSizeFunc, inFileTypeHint, outAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileOptimize func(inAudioFile AudioFileID) int32
var _audioFileOptimizeErr error

func tryAudioFileOptimize(inAudioFile AudioFileID) (int32, error) {
	if _audioFileOptimize == nil {
		return 0, symbolCallError("AudioFileOptimize", "10.2", _audioFileOptimizeErr)
	}
	return _audioFileOptimize(inAudioFile), nil
}

// AudioFileOptimize consolidates audio data and performs other internal optimizations of the file structure.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileOptimize(_:)
func AudioFileOptimize(inAudioFile AudioFileID) int32 {
	result, callErr := tryAudioFileOptimize(inAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileReadBytes func(inAudioFile AudioFileID, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, outBuffer unsafe.Pointer) int32
var _audioFileReadBytesErr error

func tryAudioFileReadBytes(inAudioFile AudioFileID, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, outBuffer unsafe.Pointer) (int32, error) {
	if _audioFileReadBytes == nil {
		return 0, symbolCallError("AudioFileReadBytes", "10.2", _audioFileReadBytesErr)
	}
	return _audioFileReadBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, outBuffer), nil
}

// AudioFileReadBytes reads bytes of audio data from an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadBytes(_:_:_:_:_:)
func AudioFileReadBytes(inAudioFile AudioFileID, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, outBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileReadBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, outBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileReadPacketData func(inAudioFile AudioFileID, inUseCache bool, ioNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) int32
var _audioFileReadPacketDataErr error

func tryAudioFileReadPacketData(inAudioFile AudioFileID, inUseCache bool, ioNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) (int32, error) {
	if _audioFileReadPacketData == nil {
		return 0, symbolCallError("AudioFileReadPacketData", "10.6", _audioFileReadPacketDataErr)
	}
	return _audioFileReadPacketData(inAudioFile, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer), nil
}

// AudioFileReadPacketData reads packets of audio data from an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileReadPacketData(_:_:_:_:_:_:_:)
func AudioFileReadPacketData(inAudioFile AudioFileID, inUseCache bool, ioNumBytes *uint32, outPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, outBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileReadPacketData(inAudioFile, inUseCache, ioNumBytes, outPacketDescriptions, inStartingPacket, ioNumPackets, outBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileRemoveUserData func(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32) int32
var _audioFileRemoveUserDataErr error

func tryAudioFileRemoveUserData(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32) (int32, error) {
	if _audioFileRemoveUserData == nil {
		return 0, symbolCallError("AudioFileRemoveUserData", "10.5", _audioFileRemoveUserDataErr)
	}
	return _audioFileRemoveUserData(inAudioFile, inUserDataID, inIndex), nil
}

// AudioFileRemoveUserData removes a user data item from an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileRemoveUserData(_:_:_:)
func AudioFileRemoveUserData(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32) int32 {
	result, callErr := tryAudioFileRemoveUserData(inAudioFile, inUserDataID, inIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileSetProperty func(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, inDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioFileSetPropertyErr error

func tryAudioFileSetProperty(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, inDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileSetProperty == nil {
		return 0, symbolCallError("AudioFileSetProperty", "10.2", _audioFileSetPropertyErr)
	}
	return _audioFileSetProperty(inAudioFile, inPropertyID, inDataSize, inPropertyData), nil
}

// AudioFileSetProperty sets the value of an audio file property
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileSetProperty(_:_:_:_:)
func AudioFileSetProperty(inAudioFile AudioFileID, inPropertyID AudioFilePropertyID, inDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileSetProperty(inAudioFile, inPropertyID, inDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileSetUserData func(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, inUserDataSize uint32, inUserData unsafe.Pointer) int32
var _audioFileSetUserDataErr error

func tryAudioFileSetUserData(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, inUserDataSize uint32, inUserData unsafe.Pointer) (int32, error) {
	if _audioFileSetUserData == nil {
		return 0, symbolCallError("AudioFileSetUserData", "10.4", _audioFileSetUserDataErr)
	}
	return _audioFileSetUserData(inAudioFile, inUserDataID, inIndex, inUserDataSize, inUserData), nil
}

// AudioFileSetUserData sets a user data item in an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileSetUserData(_:_:_:_:_:)
func AudioFileSetUserData(inAudioFile AudioFileID, inUserDataID uint32, inIndex uint32, inUserDataSize uint32, inUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileSetUserData(inAudioFile, inUserDataID, inIndex, inUserDataSize, inUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileStreamClose func(inAudioFileStream AudioFileStreamID) int32
var _audioFileStreamCloseErr error

func tryAudioFileStreamClose(inAudioFileStream AudioFileStreamID) (int32, error) {
	if _audioFileStreamClose == nil {
		return 0, symbolCallError("AudioFileStreamClose", "10.5", _audioFileStreamCloseErr)
	}
	return _audioFileStreamClose(inAudioFileStream), nil
}

// AudioFileStreamClose closes and deallocates the specified audio file stream parser.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamClose(_:)
func AudioFileStreamClose(inAudioFileStream AudioFileStreamID) int32 {
	result, callErr := tryAudioFileStreamClose(inAudioFileStream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileStreamGetProperty func(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioFileStreamGetPropertyErr error

func tryAudioFileStreamGetProperty(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileStreamGetProperty == nil {
		return 0, symbolCallError("AudioFileStreamGetProperty", "10.5", _audioFileStreamGetPropertyErr)
	}
	return _audioFileStreamGetProperty(inAudioFileStream, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioFileStreamGetProperty retrieves the value of the specified property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamGetProperty(_:_:_:_:)
func AudioFileStreamGetProperty(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileStreamGetProperty(inAudioFileStream, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileStreamGetPropertyInfo func(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, outPropertyDataSize *uint32, outWritable *bool) int32
var _audioFileStreamGetPropertyInfoErr error

func tryAudioFileStreamGetPropertyInfo(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, outPropertyDataSize *uint32, outWritable *bool) (int32, error) {
	if _audioFileStreamGetPropertyInfo == nil {
		return 0, symbolCallError("AudioFileStreamGetPropertyInfo", "10.5", _audioFileStreamGetPropertyInfoErr)
	}
	return _audioFileStreamGetPropertyInfo(inAudioFileStream, inPropertyID, outPropertyDataSize, outWritable), nil
}

// AudioFileStreamGetPropertyInfo retrieves information about a property value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamGetPropertyInfo(_:_:_:_:)
func AudioFileStreamGetPropertyInfo(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, outPropertyDataSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioFileStreamGetPropertyInfo(inAudioFileStream, inPropertyID, outPropertyDataSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileStreamOpen func(inClientData unsafe.Pointer, inPropertyListenerProc AudioFileStream_PropertyListenerProc, inPacketsProc AudioFileStream_PacketsProc, inFileTypeHint AudioFileTypeID, outAudioFileStream *AudioFileStreamID) int32
var _audioFileStreamOpenErr error

func tryAudioFileStreamOpen(inClientData unsafe.Pointer, inPropertyListenerProc AudioFileStream_PropertyListenerProc, inPacketsProc AudioFileStream_PacketsProc, inFileTypeHint AudioFileTypeID, outAudioFileStream *AudioFileStreamID) (int32, error) {
	if _audioFileStreamOpen == nil {
		return 0, symbolCallError("AudioFileStreamOpen", "10.5", _audioFileStreamOpenErr)
	}
	return _audioFileStreamOpen(inClientData, inPropertyListenerProc, inPacketsProc, inFileTypeHint, outAudioFileStream), nil
}

// AudioFileStreamOpen creates and opens a new audio file stream parser.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamOpen(_:_:_:_:_:)
func AudioFileStreamOpen(inClientData unsafe.Pointer, inPropertyListenerProc AudioFileStream_PropertyListenerProc, inPacketsProc AudioFileStream_PacketsProc, inFileTypeHint AudioFileTypeID, outAudioFileStream *AudioFileStreamID) int32 {
	result, callErr := tryAudioFileStreamOpen(inClientData, inPropertyListenerProc, inPacketsProc, inFileTypeHint, outAudioFileStream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileStreamParseBytes func(inAudioFileStream AudioFileStreamID, inDataByteSize uint32, inData unsafe.Pointer, inFlags AudioFileStreamParseFlags) int32
var _audioFileStreamParseBytesErr error

func tryAudioFileStreamParseBytes(inAudioFileStream AudioFileStreamID, inDataByteSize uint32, inData unsafe.Pointer, inFlags AudioFileStreamParseFlags) (int32, error) {
	if _audioFileStreamParseBytes == nil {
		return 0, symbolCallError("AudioFileStreamParseBytes", "10.5", _audioFileStreamParseBytesErr)
	}
	return _audioFileStreamParseBytes(inAudioFileStream, inDataByteSize, inData, inFlags), nil
}

// AudioFileStreamParseBytes passes audio file stream data to the parser.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamParseBytes(_:_:_:_:)
func AudioFileStreamParseBytes(inAudioFileStream AudioFileStreamID, inDataByteSize uint32, inData unsafe.Pointer, inFlags AudioFileStreamParseFlags) int32 {
	result, callErr := tryAudioFileStreamParseBytes(inAudioFileStream, inDataByteSize, inData, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileStreamSeek func(inAudioFileStream AudioFileStreamID, inPacketOffset int64, outDataByteOffset *int64, ioFlags *AudioFileStreamSeekFlags) int32
var _audioFileStreamSeekErr error

func tryAudioFileStreamSeek(inAudioFileStream AudioFileStreamID, inPacketOffset int64, outDataByteOffset *int64, ioFlags *AudioFileStreamSeekFlags) (int32, error) {
	if _audioFileStreamSeek == nil {
		return 0, symbolCallError("AudioFileStreamSeek", "10.5", _audioFileStreamSeekErr)
	}
	return _audioFileStreamSeek(inAudioFileStream, inPacketOffset, outDataByteOffset, ioFlags), nil
}

// AudioFileStreamSeek provides a byte offset for a specified packet in the data stream.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSeek(_:_:_:_:)
func AudioFileStreamSeek(inAudioFileStream AudioFileStreamID, inPacketOffset int64, outDataByteOffset *int64, ioFlags *AudioFileStreamSeekFlags) int32 {
	result, callErr := tryAudioFileStreamSeek(inAudioFileStream, inPacketOffset, outDataByteOffset, ioFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileStreamSetProperty func(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioFileStreamSetPropertyErr error

func tryAudioFileStreamSetProperty(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioFileStreamSetProperty == nil {
		return 0, symbolCallError("AudioFileStreamSetProperty", "10.5", _audioFileStreamSetPropertyErr)
	}
	return _audioFileStreamSetProperty(inAudioFileStream, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioFileStreamSetProperty sets the value of the specified property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileStreamSetProperty(_:_:_:_:)
func AudioFileStreamSetProperty(inAudioFileStream AudioFileStreamID, inPropertyID AudioFileStreamPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFileStreamSetProperty(inAudioFileStream, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileWriteBytes func(inAudioFile AudioFileID, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, inBuffer unsafe.Pointer) int32
var _audioFileWriteBytesErr error

func tryAudioFileWriteBytes(inAudioFile AudioFileID, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, inBuffer unsafe.Pointer) (int32, error) {
	if _audioFileWriteBytes == nil {
		return 0, symbolCallError("AudioFileWriteBytes", "10.2", _audioFileWriteBytesErr)
	}
	return _audioFileWriteBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, inBuffer), nil
}

// AudioFileWriteBytes writes bytes of audio data to an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileWriteBytes(_:_:_:_:_:)
func AudioFileWriteBytes(inAudioFile AudioFileID, inUseCache bool, inStartingByte int64, ioNumBytes *uint32, inBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileWriteBytes(inAudioFile, inUseCache, inStartingByte, ioNumBytes, inBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileWritePackets func(inAudioFile AudioFileID, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) int32
var _audioFileWritePacketsErr error

func tryAudioFileWritePackets(inAudioFile AudioFileID, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) (int32, error) {
	if _audioFileWritePackets == nil {
		return 0, symbolCallError("AudioFileWritePackets", "10.2", _audioFileWritePacketsErr)
	}
	return _audioFileWritePackets(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer), nil
}

// AudioFileWritePackets writes packets of audio data to an audio data file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileWritePackets(_:_:_:_:_:_:_:)
func AudioFileWritePackets(inAudioFile AudioFileID, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileWritePackets(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inStartingPacket, ioNumPackets, inBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFileWritePacketsWithDependencies func(inAudioFile AudioFileID, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inPacketDependencies *coreaudiotypes.AudioStreamPacketDependencyDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) int32
var _audioFileWritePacketsWithDependenciesErr error

func tryAudioFileWritePacketsWithDependencies(inAudioFile AudioFileID, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inPacketDependencies *coreaudiotypes.AudioStreamPacketDependencyDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) (int32, error) {
	if _audioFileWritePacketsWithDependencies == nil {
		return 0, symbolCallError("AudioFileWritePacketsWithDependencies", "26.0", _audioFileWritePacketsWithDependenciesErr)
	}
	return _audioFileWritePacketsWithDependencies(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inPacketDependencies, inStartingPacket, ioNumPackets, inBuffer), nil
}

// AudioFileWritePacketsWithDependencies.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFileWritePacketsWithDependencies(_:_:_:_:_:_:_:_:)
func AudioFileWritePacketsWithDependencies(inAudioFile AudioFileID, inUseCache bool, inNumBytes uint32, inPacketDescriptions *coreaudiotypes.AudioStreamPacketDescription, inPacketDependencies *coreaudiotypes.AudioStreamPacketDependencyDescription, inStartingPacket int64, ioNumPackets *uint32, inBuffer unsafe.Pointer) int32 {
	result, callErr := tryAudioFileWritePacketsWithDependencies(inAudioFile, inUseCache, inNumBytes, inPacketDescriptions, inPacketDependencies, inStartingPacket, ioNumPackets, inBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFormatGetProperty func(inPropertyID AudioFormatPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioFormatGetPropertyErr error

func tryAudioFormatGetProperty(inPropertyID AudioFormatPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioFormatGetProperty == nil {
		return 0, symbolCallError("AudioFormatGetProperty", "10.3", _audioFormatGetPropertyErr)
	}
	return _audioFormatGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData), nil
}

// AudioFormatGetProperty gets the value of an audio format property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFormatGetProperty(_:_:_:_:_:)
func AudioFormatGetProperty(inPropertyID AudioFormatPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioFormatGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioFormatGetPropertyInfo func(inPropertyID AudioFormatPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertyDataSize *uint32) int32
var _audioFormatGetPropertyInfoErr error

func tryAudioFormatGetPropertyInfo(inPropertyID AudioFormatPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertyDataSize *uint32) (int32, error) {
	if _audioFormatGetPropertyInfo == nil {
		return 0, symbolCallError("AudioFormatGetPropertyInfo", "10.3", _audioFormatGetPropertyInfoErr)
	}
	return _audioFormatGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize), nil
}

// AudioFormatGetPropertyInfo gets information about an audio format property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioFormatGetPropertyInfo(_:_:_:_:)
func AudioFormatGetPropertyInfo(inPropertyID AudioFormatPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertyDataSize *uint32) int32 {
	result, callErr := tryAudioFormatGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioOutputUnitStart func(ci AudioUnit) int32
var _audioOutputUnitStartErr error

func tryAudioOutputUnitStart(ci AudioUnit) (int32, error) {
	if _audioOutputUnitStart == nil {
		return 0, symbolCallError("AudioOutputUnitStart", "10.0", _audioOutputUnitStartErr)
	}
	return _audioOutputUnitStart(ci), nil
}

// AudioOutputUnitStart starts an I/O audio unit, which in turn starts the audio unit processing graph that it is connected to.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStart(_:)
func AudioOutputUnitStart(ci AudioUnit) int32 {
	result, callErr := tryAudioOutputUnitStart(ci)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioOutputUnitStop func(ci AudioUnit) int32
var _audioOutputUnitStopErr error

func tryAudioOutputUnitStop(ci AudioUnit) (int32, error) {
	if _audioOutputUnitStop == nil {
		return 0, symbolCallError("AudioOutputUnitStop", "10.0", _audioOutputUnitStopErr)
	}
	return _audioOutputUnitStop(ci), nil
}

// AudioOutputUnitStop stops an I/O audio unit, which in turn stops the audio unit processing graph that it is connected to.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioOutputUnitStop(_:)
func AudioOutputUnitStop(ci AudioUnit) int32 {
	result, callErr := tryAudioOutputUnitStop(ci)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueAddPropertyListener func(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) int32
var _audioQueueAddPropertyListenerErr error

func tryAudioQueueAddPropertyListener(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) (int32, error) {
	if _audioQueueAddPropertyListener == nil {
		return 0, symbolCallError("AudioQueueAddPropertyListener", "10.5", _audioQueueAddPropertyListenerErr)
	}
	return _audioQueueAddPropertyListener(inAQ, inID, inProc, inUserData), nil
}

// AudioQueueAddPropertyListener adds a property listener callback to an audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAddPropertyListener(_:_:_:_:)
func AudioQueueAddPropertyListener(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioQueueAddPropertyListener(inAQ, inID, inProc, inUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueAllocateBuffer func(inAQ AudioQueueRef, inBufferByteSize uint32, outBuffer *AudioQueueBufferRef) int32
var _audioQueueAllocateBufferErr error

func tryAudioQueueAllocateBuffer(inAQ AudioQueueRef, inBufferByteSize uint32, outBuffer *AudioQueueBufferRef) (int32, error) {
	if _audioQueueAllocateBuffer == nil {
		return 0, symbolCallError("AudioQueueAllocateBuffer", "10.5", _audioQueueAllocateBufferErr)
	}
	return _audioQueueAllocateBuffer(inAQ, inBufferByteSize, outBuffer), nil
}

// AudioQueueAllocateBuffer asks an audio queue object to allocate an audio queue buffer.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAllocateBuffer(_:_:_:)
func AudioQueueAllocateBuffer(inAQ AudioQueueRef, inBufferByteSize uint32, outBuffer *AudioQueueBufferRef) int32 {
	result, callErr := tryAudioQueueAllocateBuffer(inAQ, inBufferByteSize, outBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueAllocateBufferWithPacketDescriptions func(inAQ AudioQueueRef, inBufferByteSize uint32, inNumberPacketDescriptions uint32, outBuffer *AudioQueueBufferRef) int32
var _audioQueueAllocateBufferWithPacketDescriptionsErr error

func tryAudioQueueAllocateBufferWithPacketDescriptions(inAQ AudioQueueRef, inBufferByteSize uint32, inNumberPacketDescriptions uint32, outBuffer *AudioQueueBufferRef) (int32, error) {
	if _audioQueueAllocateBufferWithPacketDescriptions == nil {
		return 0, symbolCallError("AudioQueueAllocateBufferWithPacketDescriptions", "10.6", _audioQueueAllocateBufferWithPacketDescriptionsErr)
	}
	return _audioQueueAllocateBufferWithPacketDescriptions(inAQ, inBufferByteSize, inNumberPacketDescriptions, outBuffer), nil
}

// AudioQueueAllocateBufferWithPacketDescriptions asks an audio queue object to allocate an audio queue buffer with space for packet descriptions.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueAllocateBufferWithPacketDescriptions(_:_:_:_:)
func AudioQueueAllocateBufferWithPacketDescriptions(inAQ AudioQueueRef, inBufferByteSize uint32, inNumberPacketDescriptions uint32, outBuffer *AudioQueueBufferRef) int32 {
	result, callErr := tryAudioQueueAllocateBufferWithPacketDescriptions(inAQ, inBufferByteSize, inNumberPacketDescriptions, outBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueCreateTimeline func(inAQ AudioQueueRef, outTimeline *AudioQueueTimelineRef) int32
var _audioQueueCreateTimelineErr error

func tryAudioQueueCreateTimeline(inAQ AudioQueueRef, outTimeline *AudioQueueTimelineRef) (int32, error) {
	if _audioQueueCreateTimeline == nil {
		return 0, symbolCallError("AudioQueueCreateTimeline", "10.5", _audioQueueCreateTimelineErr)
	}
	return _audioQueueCreateTimeline(inAQ, outTimeline), nil
}

// AudioQueueCreateTimeline creates a timeline object for an audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueCreateTimeline(_:_:)
func AudioQueueCreateTimeline(inAQ AudioQueueRef, outTimeline *AudioQueueTimelineRef) int32 {
	result, callErr := tryAudioQueueCreateTimeline(inAQ, outTimeline)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueDeviceGetCurrentTime func(inAQ AudioQueueRef, outTimeStamp *coreaudiotypes.AudioTimeStamp) int32
var _audioQueueDeviceGetCurrentTimeErr error

func tryAudioQueueDeviceGetCurrentTime(inAQ AudioQueueRef, outTimeStamp *coreaudiotypes.AudioTimeStamp) (int32, error) {
	if _audioQueueDeviceGetCurrentTime == nil {
		return 0, symbolCallError("AudioQueueDeviceGetCurrentTime", "10.5", _audioQueueDeviceGetCurrentTimeErr)
	}
	return _audioQueueDeviceGetCurrentTime(inAQ, outTimeStamp), nil
}

// AudioQueueDeviceGetCurrentTime gets the current time of the audio hardware device associated with an audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceGetCurrentTime(_:_:)
func AudioQueueDeviceGetCurrentTime(inAQ AudioQueueRef, outTimeStamp *coreaudiotypes.AudioTimeStamp) int32 {
	result, callErr := tryAudioQueueDeviceGetCurrentTime(inAQ, outTimeStamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueDeviceGetNearestStartTime func(inAQ AudioQueueRef, ioRequestedStartTime *coreaudiotypes.AudioTimeStamp, inFlags uint32) int32
var _audioQueueDeviceGetNearestStartTimeErr error

func tryAudioQueueDeviceGetNearestStartTime(inAQ AudioQueueRef, ioRequestedStartTime *coreaudiotypes.AudioTimeStamp, inFlags uint32) (int32, error) {
	if _audioQueueDeviceGetNearestStartTime == nil {
		return 0, symbolCallError("AudioQueueDeviceGetNearestStartTime", "10.5", _audioQueueDeviceGetNearestStartTimeErr)
	}
	return _audioQueueDeviceGetNearestStartTime(inAQ, ioRequestedStartTime, inFlags), nil
}

// AudioQueueDeviceGetNearestStartTime gets the start time, for an audio hardware device, that is closest to a requested start time.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceGetNearestStartTime(_:_:_:)
func AudioQueueDeviceGetNearestStartTime(inAQ AudioQueueRef, ioRequestedStartTime *coreaudiotypes.AudioTimeStamp, inFlags uint32) int32 {
	result, callErr := tryAudioQueueDeviceGetNearestStartTime(inAQ, ioRequestedStartTime, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueDeviceTranslateTime func(inAQ AudioQueueRef, inTime *coreaudiotypes.AudioTimeStamp, outTime *coreaudiotypes.AudioTimeStamp) int32
var _audioQueueDeviceTranslateTimeErr error

func tryAudioQueueDeviceTranslateTime(inAQ AudioQueueRef, inTime *coreaudiotypes.AudioTimeStamp, outTime *coreaudiotypes.AudioTimeStamp) (int32, error) {
	if _audioQueueDeviceTranslateTime == nil {
		return 0, symbolCallError("AudioQueueDeviceTranslateTime", "10.5", _audioQueueDeviceTranslateTimeErr)
	}
	return _audioQueueDeviceTranslateTime(inAQ, inTime, outTime), nil
}

// AudioQueueDeviceTranslateTime converts the time for an audio queue’s associated audio hardware device from one time base representation to another.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDeviceTranslateTime(_:_:_:)
func AudioQueueDeviceTranslateTime(inAQ AudioQueueRef, inTime *coreaudiotypes.AudioTimeStamp, outTime *coreaudiotypes.AudioTimeStamp) int32 {
	result, callErr := tryAudioQueueDeviceTranslateTime(inAQ, inTime, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueDispose func(inAQ AudioQueueRef, inImmediate bool) int32
var _audioQueueDisposeErr error

func tryAudioQueueDispose(inAQ AudioQueueRef, inImmediate bool) (int32, error) {
	if _audioQueueDispose == nil {
		return 0, symbolCallError("AudioQueueDispose", "10.5", _audioQueueDisposeErr)
	}
	return _audioQueueDispose(inAQ, inImmediate), nil
}

// AudioQueueDispose disposes of an audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDispose(_:_:)
func AudioQueueDispose(inAQ AudioQueueRef, inImmediate bool) int32 {
	result, callErr := tryAudioQueueDispose(inAQ, inImmediate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueDisposeTimeline func(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef) int32
var _audioQueueDisposeTimelineErr error

func tryAudioQueueDisposeTimeline(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef) (int32, error) {
	if _audioQueueDisposeTimeline == nil {
		return 0, symbolCallError("AudioQueueDisposeTimeline", "10.5", _audioQueueDisposeTimelineErr)
	}
	return _audioQueueDisposeTimeline(inAQ, inTimeline), nil
}

// AudioQueueDisposeTimeline disposes of an audio queue’s timeline object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueDisposeTimeline(_:_:)
func AudioQueueDisposeTimeline(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef) int32 {
	result, callErr := tryAudioQueueDisposeTimeline(inAQ, inTimeline)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueEnqueueBuffer func(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs uint32, inPacketDescs *coreaudiotypes.AudioStreamPacketDescription) int32
var _audioQueueEnqueueBufferErr error

func tryAudioQueueEnqueueBuffer(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs uint32, inPacketDescs *coreaudiotypes.AudioStreamPacketDescription) (int32, error) {
	if _audioQueueEnqueueBuffer == nil {
		return 0, symbolCallError("AudioQueueEnqueueBuffer", "10.5", _audioQueueEnqueueBufferErr)
	}
	return _audioQueueEnqueueBuffer(inAQ, inBuffer, inNumPacketDescs, inPacketDescs), nil
}

// AudioQueueEnqueueBuffer adds a buffer to the buffer queue of a recording or playback audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueEnqueueBuffer(_:_:_:_:)
func AudioQueueEnqueueBuffer(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs uint32, inPacketDescs *coreaudiotypes.AudioStreamPacketDescription) int32 {
	result, callErr := tryAudioQueueEnqueueBuffer(inAQ, inBuffer, inNumPacketDescs, inPacketDescs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueEnqueueBufferWithParameters func(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs uint32, inPacketDescs *coreaudiotypes.AudioStreamPacketDescription, inTrimFramesAtStart uint32, inTrimFramesAtEnd uint32, inNumParamValues uint32, inParamValues *AudioQueueParameterEvent, inStartTime *coreaudiotypes.AudioTimeStamp, outActualStartTime *coreaudiotypes.AudioTimeStamp) int32
var _audioQueueEnqueueBufferWithParametersErr error

func tryAudioQueueEnqueueBufferWithParameters(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs uint32, inPacketDescs *coreaudiotypes.AudioStreamPacketDescription, inTrimFramesAtStart uint32, inTrimFramesAtEnd uint32, inNumParamValues uint32, inParamValues *AudioQueueParameterEvent, inStartTime *coreaudiotypes.AudioTimeStamp, outActualStartTime *coreaudiotypes.AudioTimeStamp) (int32, error) {
	if _audioQueueEnqueueBufferWithParameters == nil {
		return 0, symbolCallError("AudioQueueEnqueueBufferWithParameters", "10.5", _audioQueueEnqueueBufferWithParametersErr)
	}
	return _audioQueueEnqueueBufferWithParameters(inAQ, inBuffer, inNumPacketDescs, inPacketDescs, inTrimFramesAtStart, inTrimFramesAtEnd, inNumParamValues, inParamValues, inStartTime, outActualStartTime), nil
}

// AudioQueueEnqueueBufferWithParameters adds a buffer to the buffer queue of a playback audio queue object, specifying start time and other settings.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueEnqueueBufferWithParameters(_:_:_:_:_:_:_:_:_:_:)
func AudioQueueEnqueueBufferWithParameters(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef, inNumPacketDescs uint32, inPacketDescs *coreaudiotypes.AudioStreamPacketDescription, inTrimFramesAtStart uint32, inTrimFramesAtEnd uint32, inNumParamValues uint32, inParamValues *AudioQueueParameterEvent, inStartTime *coreaudiotypes.AudioTimeStamp, outActualStartTime *coreaudiotypes.AudioTimeStamp) int32 {
	result, callErr := tryAudioQueueEnqueueBufferWithParameters(inAQ, inBuffer, inNumPacketDescs, inPacketDescs, inTrimFramesAtStart, inTrimFramesAtEnd, inNumParamValues, inParamValues, inStartTime, outActualStartTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueFlush func(inAQ AudioQueueRef) int32
var _audioQueueFlushErr error

func tryAudioQueueFlush(inAQ AudioQueueRef) (int32, error) {
	if _audioQueueFlush == nil {
		return 0, symbolCallError("AudioQueueFlush", "10.5", _audioQueueFlushErr)
	}
	return _audioQueueFlush(inAQ), nil
}

// AudioQueueFlush resets an audio queue’s decoder state.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueFlush(_:)
func AudioQueueFlush(inAQ AudioQueueRef) int32 {
	result, callErr := tryAudioQueueFlush(inAQ)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueFreeBuffer func(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef) int32
var _audioQueueFreeBufferErr error

func tryAudioQueueFreeBuffer(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef) (int32, error) {
	if _audioQueueFreeBuffer == nil {
		return 0, symbolCallError("AudioQueueFreeBuffer", "10.5", _audioQueueFreeBufferErr)
	}
	return _audioQueueFreeBuffer(inAQ, inBuffer), nil
}

// AudioQueueFreeBuffer asks an audio queue to dispose of an audio queue buffer.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueFreeBuffer(_:_:)
func AudioQueueFreeBuffer(inAQ AudioQueueRef, inBuffer AudioQueueBufferRef) int32 {
	result, callErr := tryAudioQueueFreeBuffer(inAQ, inBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueGetCurrentTime func(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef, outTimeStamp *coreaudiotypes.AudioTimeStamp, outTimelineDiscontinuity *bool) int32
var _audioQueueGetCurrentTimeErr error

func tryAudioQueueGetCurrentTime(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef, outTimeStamp *coreaudiotypes.AudioTimeStamp, outTimelineDiscontinuity *bool) (int32, error) {
	if _audioQueueGetCurrentTime == nil {
		return 0, symbolCallError("AudioQueueGetCurrentTime", "10.5", _audioQueueGetCurrentTimeErr)
	}
	return _audioQueueGetCurrentTime(inAQ, inTimeline, outTimeStamp, outTimelineDiscontinuity), nil
}

// AudioQueueGetCurrentTime gets the current audio queue time.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetCurrentTime(_:_:_:_:)
func AudioQueueGetCurrentTime(inAQ AudioQueueRef, inTimeline AudioQueueTimelineRef, outTimeStamp *coreaudiotypes.AudioTimeStamp, outTimelineDiscontinuity *bool) int32 {
	result, callErr := tryAudioQueueGetCurrentTime(inAQ, inTimeline, outTimeStamp, outTimelineDiscontinuity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueGetParameter func(inAQ AudioQueueRef, inParamID AudioQueueParameterID, outValue *AudioQueueParameterValue) int32
var _audioQueueGetParameterErr error

func tryAudioQueueGetParameter(inAQ AudioQueueRef, inParamID AudioQueueParameterID, outValue *AudioQueueParameterValue) (int32, error) {
	if _audioQueueGetParameter == nil {
		return 0, symbolCallError("AudioQueueGetParameter", "10.5", _audioQueueGetParameterErr)
	}
	return _audioQueueGetParameter(inAQ, inParamID, outValue), nil
}

// AudioQueueGetParameter gets an audio queue parameter value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetParameter(_:_:_:)
func AudioQueueGetParameter(inAQ AudioQueueRef, inParamID AudioQueueParameterID, outValue *AudioQueueParameterValue) int32 {
	result, callErr := tryAudioQueueGetParameter(inAQ, inParamID, outValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueGetProperty func(inAQ AudioQueueRef, inID AudioQueuePropertyID, outData unsafe.Pointer, ioDataSize *uint32) int32
var _audioQueueGetPropertyErr error

func tryAudioQueueGetProperty(inAQ AudioQueueRef, inID AudioQueuePropertyID, outData unsafe.Pointer, ioDataSize *uint32) (int32, error) {
	if _audioQueueGetProperty == nil {
		return 0, symbolCallError("AudioQueueGetProperty", "10.5", _audioQueueGetPropertyErr)
	}
	return _audioQueueGetProperty(inAQ, inID, outData, ioDataSize), nil
}

// AudioQueueGetProperty gets an audio queue property value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetProperty(_:_:_:_:)
func AudioQueueGetProperty(inAQ AudioQueueRef, inID AudioQueuePropertyID, outData unsafe.Pointer, ioDataSize *uint32) int32 {
	result, callErr := tryAudioQueueGetProperty(inAQ, inID, outData, ioDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueGetPropertySize func(inAQ AudioQueueRef, inID AudioQueuePropertyID, outDataSize *uint32) int32
var _audioQueueGetPropertySizeErr error

func tryAudioQueueGetPropertySize(inAQ AudioQueueRef, inID AudioQueuePropertyID, outDataSize *uint32) (int32, error) {
	if _audioQueueGetPropertySize == nil {
		return 0, symbolCallError("AudioQueueGetPropertySize", "10.5", _audioQueueGetPropertySizeErr)
	}
	return _audioQueueGetPropertySize(inAQ, inID, outDataSize), nil
}

// AudioQueueGetPropertySize gets the size of the value of an audio queue property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueGetPropertySize(_:_:_:)
func AudioQueueGetPropertySize(inAQ AudioQueueRef, inID AudioQueuePropertyID, outDataSize *uint32) int32 {
	result, callErr := tryAudioQueueGetPropertySize(inAQ, inID, outDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueNewInput func(inFormat *coreaudiotypes.AudioStreamBasicDescription, inCallbackProc AudioQueueInputCallback, inUserData unsafe.Pointer, inCallbackRunLoop corefoundation.CFRunLoopRef, inCallbackRunLoopMode corefoundation.CFStringRef, inFlags uint32, outAQ *AudioQueueRef) int32
var _audioQueueNewInputErr error

func tryAudioQueueNewInput(inFormat *coreaudiotypes.AudioStreamBasicDescription, inCallbackProc AudioQueueInputCallback, inUserData unsafe.Pointer, inCallbackRunLoop corefoundation.CFRunLoopRef, inCallbackRunLoopMode corefoundation.CFStringRef, inFlags uint32, outAQ *AudioQueueRef) (int32, error) {
	if _audioQueueNewInput == nil {
		return 0, symbolCallError("AudioQueueNewInput", "10.5", _audioQueueNewInputErr)
	}
	return _audioQueueNewInput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ), nil
}

// AudioQueueNewInput creates a new recording audio queue object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewInput(_:_:_:_:_:_:_:)
func AudioQueueNewInput(inFormat *coreaudiotypes.AudioStreamBasicDescription, inCallbackProc AudioQueueInputCallback, inUserData unsafe.Pointer, inCallbackRunLoop corefoundation.CFRunLoopRef, inCallbackRunLoopMode corefoundation.CFStringRef, inFlags uint32, outAQ *AudioQueueRef) int32 {
	result, callErr := tryAudioQueueNewInput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueNewInputWithDispatchQueue func(outAQ *AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, inCallbackDispatchQueue uintptr, inCallbackBlock unsafe.Pointer) int32
var _audioQueueNewInputWithDispatchQueueErr error

func tryAudioQueueNewInputWithDispatchQueue(outAQ *AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, inCallbackDispatchQueue dispatch.Queue, inCallbackBlock AudioQueueInputCallbackBlock) (int32, error) {
	if _audioQueueNewInputWithDispatchQueue == nil {
		return 0, symbolCallError("AudioQueueNewInputWithDispatchQueue", "10.6", _audioQueueNewInputWithDispatchQueueErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *uintptr, blockArg1 *AudioQueueBuffer, blockArg2 unsafe.Pointer, blockArg3 uint32, blockArg4 unsafe.Pointer) {
		inCallbackBlock(blockArg0, blockArg1, (*coreaudiotypes.AudioTimeStamp)(blockArg2), blockArg3, (*coreaudiotypes.AudioStreamPacketDescription)(blockArg4))
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _audioQueueNewInputWithDispatchQueue(outAQ, inFormat, inFlags, uintptr(inCallbackDispatchQueue.Handle()), _block0), nil
}

// AudioQueueNewInputWithDispatchQueue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewInputWithDispatchQueue(_:_:_:_:_:)
func AudioQueueNewInputWithDispatchQueue(outAQ *AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, inCallbackDispatchQueue dispatch.Queue, inCallbackBlock AudioQueueInputCallbackBlock) int32 {
	result, callErr := tryAudioQueueNewInputWithDispatchQueue(outAQ, inFormat, inFlags, inCallbackDispatchQueue, inCallbackBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueNewOutput func(inFormat *coreaudiotypes.AudioStreamBasicDescription, inCallbackProc AudioQueueOutputCallback, inUserData unsafe.Pointer, inCallbackRunLoop corefoundation.CFRunLoopRef, inCallbackRunLoopMode corefoundation.CFStringRef, inFlags uint32, outAQ *AudioQueueRef) int32
var _audioQueueNewOutputErr error

func tryAudioQueueNewOutput(inFormat *coreaudiotypes.AudioStreamBasicDescription, inCallbackProc AudioQueueOutputCallback, inUserData unsafe.Pointer, inCallbackRunLoop corefoundation.CFRunLoopRef, inCallbackRunLoopMode corefoundation.CFStringRef, inFlags uint32, outAQ *AudioQueueRef) (int32, error) {
	if _audioQueueNewOutput == nil {
		return 0, symbolCallError("AudioQueueNewOutput", "10.5", _audioQueueNewOutputErr)
	}
	return _audioQueueNewOutput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ), nil
}

// AudioQueueNewOutput creates a new playback audio queue object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewOutput(_:_:_:_:_:_:_:)
func AudioQueueNewOutput(inFormat *coreaudiotypes.AudioStreamBasicDescription, inCallbackProc AudioQueueOutputCallback, inUserData unsafe.Pointer, inCallbackRunLoop corefoundation.CFRunLoopRef, inCallbackRunLoopMode corefoundation.CFStringRef, inFlags uint32, outAQ *AudioQueueRef) int32 {
	result, callErr := tryAudioQueueNewOutput(inFormat, inCallbackProc, inUserData, inCallbackRunLoop, inCallbackRunLoopMode, inFlags, outAQ)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueNewOutputWithDispatchQueue func(outAQ *AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, inCallbackDispatchQueue uintptr, inCallbackBlock unsafe.Pointer) int32
var _audioQueueNewOutputWithDispatchQueueErr error

func tryAudioQueueNewOutputWithDispatchQueue(outAQ *AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, inCallbackDispatchQueue dispatch.Queue, inCallbackBlock AudioQueueOutputCallbackBlock) (int32, error) {
	if _audioQueueNewOutputWithDispatchQueue == nil {
		return 0, symbolCallError("AudioQueueNewOutputWithDispatchQueue", "10.6", _audioQueueNewOutputWithDispatchQueueErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *uintptr, blockArg1 *AudioQueueBuffer) {
		inCallbackBlock(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _audioQueueNewOutputWithDispatchQueue(outAQ, inFormat, inFlags, uintptr(inCallbackDispatchQueue.Handle()), _block0), nil
}

// AudioQueueNewOutputWithDispatchQueue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewOutputWithDispatchQueue(_:_:_:_:_:)
func AudioQueueNewOutputWithDispatchQueue(outAQ *AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inFlags uint32, inCallbackDispatchQueue dispatch.Queue, inCallbackBlock AudioQueueOutputCallbackBlock) int32 {
	result, callErr := tryAudioQueueNewOutputWithDispatchQueue(outAQ, inFormat, inFlags, inCallbackDispatchQueue, inCallbackBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueOfflineRender func(inAQ AudioQueueRef, inTimestamp *coreaudiotypes.AudioTimeStamp, ioBuffer AudioQueueBufferRef, inNumberFrames uint32) int32
var _audioQueueOfflineRenderErr error

func tryAudioQueueOfflineRender(inAQ AudioQueueRef, inTimestamp *coreaudiotypes.AudioTimeStamp, ioBuffer AudioQueueBufferRef, inNumberFrames uint32) (int32, error) {
	if _audioQueueOfflineRender == nil {
		return 0, symbolCallError("AudioQueueOfflineRender", "10.5", _audioQueueOfflineRenderErr)
	}
	return _audioQueueOfflineRender(inAQ, inTimestamp, ioBuffer, inNumberFrames), nil
}

// AudioQueueOfflineRender exports audio to a buffer, instead of to a device, using a playback audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueOfflineRender(_:_:_:_:)
func AudioQueueOfflineRender(inAQ AudioQueueRef, inTimestamp *coreaudiotypes.AudioTimeStamp, ioBuffer AudioQueueBufferRef, inNumberFrames uint32) int32 {
	result, callErr := tryAudioQueueOfflineRender(inAQ, inTimestamp, ioBuffer, inNumberFrames)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueuePause func(inAQ AudioQueueRef) int32
var _audioQueuePauseErr error

func tryAudioQueuePause(inAQ AudioQueueRef) (int32, error) {
	if _audioQueuePause == nil {
		return 0, symbolCallError("AudioQueuePause", "10.5", _audioQueuePauseErr)
	}
	return _audioQueuePause(inAQ), nil
}

// AudioQueuePause pauses audio playback or recording.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePause(_:)
func AudioQueuePause(inAQ AudioQueueRef) int32 {
	result, callErr := tryAudioQueuePause(inAQ)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueuePrime func(inAQ AudioQueueRef, inNumberOfFramesToPrepare uint32, outNumberOfFramesPrepared *uint32) int32
var _audioQueuePrimeErr error

func tryAudioQueuePrime(inAQ AudioQueueRef, inNumberOfFramesToPrepare uint32, outNumberOfFramesPrepared *uint32) (int32, error) {
	if _audioQueuePrime == nil {
		return 0, symbolCallError("AudioQueuePrime", "10.5", _audioQueuePrimeErr)
	}
	return _audioQueuePrime(inAQ, inNumberOfFramesToPrepare, outNumberOfFramesPrepared), nil
}

// AudioQueuePrime decodes enqueued buffers in preparation for playback.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueuePrime(_:_:_:)
func AudioQueuePrime(inAQ AudioQueueRef, inNumberOfFramesToPrepare uint32, outNumberOfFramesPrepared *uint32) int32 {
	result, callErr := tryAudioQueuePrime(inAQ, inNumberOfFramesToPrepare, outNumberOfFramesPrepared)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueProcessingTapDispose func(inAQTap AudioQueueProcessingTapRef) int32
var _audioQueueProcessingTapDisposeErr error

func tryAudioQueueProcessingTapDispose(inAQTap AudioQueueProcessingTapRef) (int32, error) {
	if _audioQueueProcessingTapDispose == nil {
		return 0, symbolCallError("AudioQueueProcessingTapDispose", "10.7", _audioQueueProcessingTapDisposeErr)
	}
	return _audioQueueProcessingTapDispose(inAQTap), nil
}

// AudioQueueProcessingTapDispose.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapDispose(_:)
func AudioQueueProcessingTapDispose(inAQTap AudioQueueProcessingTapRef) int32 {
	result, callErr := tryAudioQueueProcessingTapDispose(inAQTap)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueProcessingTapGetQueueTime func(inAQTap AudioQueueProcessingTapRef, outQueueSampleTime *float64, outQueueFrameCount *uint32) int32
var _audioQueueProcessingTapGetQueueTimeErr error

func tryAudioQueueProcessingTapGetQueueTime(inAQTap AudioQueueProcessingTapRef, outQueueSampleTime *float64, outQueueFrameCount *uint32) (int32, error) {
	if _audioQueueProcessingTapGetQueueTime == nil {
		return 0, symbolCallError("AudioQueueProcessingTapGetQueueTime", "10.8", _audioQueueProcessingTapGetQueueTimeErr)
	}
	return _audioQueueProcessingTapGetQueueTime(inAQTap, outQueueSampleTime, outQueueFrameCount), nil
}

// AudioQueueProcessingTapGetQueueTime.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapGetQueueTime(_:_:_:)
func AudioQueueProcessingTapGetQueueTime(inAQTap AudioQueueProcessingTapRef, outQueueSampleTime *float64, outQueueFrameCount *uint32) int32 {
	result, callErr := tryAudioQueueProcessingTapGetQueueTime(inAQTap, outQueueSampleTime, outQueueFrameCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueProcessingTapGetSourceAudio func(inAQTap AudioQueueProcessingTapRef, inNumberFrames uint32, ioTimeStamp *coreaudiotypes.AudioTimeStamp, outFlags *AudioQueueProcessingTapFlags, outNumberFrames *uint32, ioData *coreaudiotypes.AudioBufferList) int32
var _audioQueueProcessingTapGetSourceAudioErr error

func tryAudioQueueProcessingTapGetSourceAudio(inAQTap AudioQueueProcessingTapRef, inNumberFrames uint32, ioTimeStamp *coreaudiotypes.AudioTimeStamp, outFlags *AudioQueueProcessingTapFlags, outNumberFrames *uint32, ioData *coreaudiotypes.AudioBufferList) (int32, error) {
	if _audioQueueProcessingTapGetSourceAudio == nil {
		return 0, symbolCallError("AudioQueueProcessingTapGetSourceAudio", "10.7", _audioQueueProcessingTapGetSourceAudioErr)
	}
	return _audioQueueProcessingTapGetSourceAudio(inAQTap, inNumberFrames, ioTimeStamp, outFlags, outNumberFrames, ioData), nil
}

// AudioQueueProcessingTapGetSourceAudio.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapGetSourceAudio(_:_:_:_:_:_:)
func AudioQueueProcessingTapGetSourceAudio(inAQTap AudioQueueProcessingTapRef, inNumberFrames uint32, ioTimeStamp *coreaudiotypes.AudioTimeStamp, outFlags *AudioQueueProcessingTapFlags, outNumberFrames *uint32, ioData *coreaudiotypes.AudioBufferList) int32 {
	result, callErr := tryAudioQueueProcessingTapGetSourceAudio(inAQTap, inNumberFrames, ioTimeStamp, outFlags, outNumberFrames, ioData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueProcessingTapNew func(inAQ AudioQueueRef, inCallback AudioQueueProcessingTapCallback, inClientData unsafe.Pointer, inFlags AudioQueueProcessingTapFlags, outMaxFrames *uint32, outProcessingFormat *coreaudiotypes.AudioStreamBasicDescription, outAQTap *AudioQueueProcessingTapRef) int32
var _audioQueueProcessingTapNewErr error

func tryAudioQueueProcessingTapNew(inAQ AudioQueueRef, inCallback AudioQueueProcessingTapCallback, inClientData unsafe.Pointer, inFlags AudioQueueProcessingTapFlags, outMaxFrames *uint32, outProcessingFormat *coreaudiotypes.AudioStreamBasicDescription, outAQTap *AudioQueueProcessingTapRef) (int32, error) {
	if _audioQueueProcessingTapNew == nil {
		return 0, symbolCallError("AudioQueueProcessingTapNew", "10.7", _audioQueueProcessingTapNewErr)
	}
	return _audioQueueProcessingTapNew(inAQ, inCallback, inClientData, inFlags, outMaxFrames, outProcessingFormat, outAQTap), nil
}

// AudioQueueProcessingTapNew.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueProcessingTapNew(_:_:_:_:_:_:_:)
func AudioQueueProcessingTapNew(inAQ AudioQueueRef, inCallback AudioQueueProcessingTapCallback, inClientData unsafe.Pointer, inFlags AudioQueueProcessingTapFlags, outMaxFrames *uint32, outProcessingFormat *coreaudiotypes.AudioStreamBasicDescription, outAQTap *AudioQueueProcessingTapRef) int32 {
	result, callErr := tryAudioQueueProcessingTapNew(inAQ, inCallback, inClientData, inFlags, outMaxFrames, outProcessingFormat, outAQTap)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueRemovePropertyListener func(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) int32
var _audioQueueRemovePropertyListenerErr error

func tryAudioQueueRemovePropertyListener(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) (int32, error) {
	if _audioQueueRemovePropertyListener == nil {
		return 0, symbolCallError("AudioQueueRemovePropertyListener", "10.5", _audioQueueRemovePropertyListenerErr)
	}
	return _audioQueueRemovePropertyListener(inAQ, inID, inProc, inUserData), nil
}

// AudioQueueRemovePropertyListener removes a property listener callback from an audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueRemovePropertyListener(_:_:_:_:)
func AudioQueueRemovePropertyListener(inAQ AudioQueueRef, inID AudioQueuePropertyID, inProc AudioQueuePropertyListenerProc, inUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioQueueRemovePropertyListener(inAQ, inID, inProc, inUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueReset func(inAQ AudioQueueRef) int32
var _audioQueueResetErr error

func tryAudioQueueReset(inAQ AudioQueueRef) (int32, error) {
	if _audioQueueReset == nil {
		return 0, symbolCallError("AudioQueueReset", "10.5", _audioQueueResetErr)
	}
	return _audioQueueReset(inAQ), nil
}

// AudioQueueReset resets an audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueReset(_:)
func AudioQueueReset(inAQ AudioQueueRef) int32 {
	result, callErr := tryAudioQueueReset(inAQ)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueSetOfflineRenderFormat func(inAQ AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inLayout *coreaudiotypes.AudioChannelLayout) int32
var _audioQueueSetOfflineRenderFormatErr error

func tryAudioQueueSetOfflineRenderFormat(inAQ AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inLayout *coreaudiotypes.AudioChannelLayout) (int32, error) {
	if _audioQueueSetOfflineRenderFormat == nil {
		return 0, symbolCallError("AudioQueueSetOfflineRenderFormat", "10.5", _audioQueueSetOfflineRenderFormatErr)
	}
	return _audioQueueSetOfflineRenderFormat(inAQ, inFormat, inLayout), nil
}

// AudioQueueSetOfflineRenderFormat sets the rendering mode and audio format for a playback audio queue.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetOfflineRenderFormat(_:_:_:)
func AudioQueueSetOfflineRenderFormat(inAQ AudioQueueRef, inFormat *coreaudiotypes.AudioStreamBasicDescription, inLayout *coreaudiotypes.AudioChannelLayout) int32 {
	result, callErr := tryAudioQueueSetOfflineRenderFormat(inAQ, inFormat, inLayout)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueSetParameter func(inAQ AudioQueueRef, inParamID AudioQueueParameterID, inValue AudioQueueParameterValue) int32
var _audioQueueSetParameterErr error

func tryAudioQueueSetParameter(inAQ AudioQueueRef, inParamID AudioQueueParameterID, inValue AudioQueueParameterValue) (int32, error) {
	if _audioQueueSetParameter == nil {
		return 0, symbolCallError("AudioQueueSetParameter", "10.5", _audioQueueSetParameterErr)
	}
	return _audioQueueSetParameter(inAQ, inParamID, inValue), nil
}

// AudioQueueSetParameter sets a playback audio queue parameter value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetParameter(_:_:_:)
func AudioQueueSetParameter(inAQ AudioQueueRef, inParamID AudioQueueParameterID, inValue AudioQueueParameterValue) int32 {
	result, callErr := tryAudioQueueSetParameter(inAQ, inParamID, inValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueSetProperty func(inAQ AudioQueueRef, inID AudioQueuePropertyID, inData unsafe.Pointer, inDataSize uint32) int32
var _audioQueueSetPropertyErr error

func tryAudioQueueSetProperty(inAQ AudioQueueRef, inID AudioQueuePropertyID, inData unsafe.Pointer, inDataSize uint32) (int32, error) {
	if _audioQueueSetProperty == nil {
		return 0, symbolCallError("AudioQueueSetProperty", "10.5", _audioQueueSetPropertyErr)
	}
	return _audioQueueSetProperty(inAQ, inID, inData, inDataSize), nil
}

// AudioQueueSetProperty sets an audio queue property value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueSetProperty(_:_:_:_:)
func AudioQueueSetProperty(inAQ AudioQueueRef, inID AudioQueuePropertyID, inData unsafe.Pointer, inDataSize uint32) int32 {
	result, callErr := tryAudioQueueSetProperty(inAQ, inID, inData, inDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueStart func(inAQ AudioQueueRef, inStartTime *coreaudiotypes.AudioTimeStamp) int32
var _audioQueueStartErr error

func tryAudioQueueStart(inAQ AudioQueueRef, inStartTime *coreaudiotypes.AudioTimeStamp) (int32, error) {
	if _audioQueueStart == nil {
		return 0, symbolCallError("AudioQueueStart", "10.5", _audioQueueStartErr)
	}
	return _audioQueueStart(inAQ, inStartTime), nil
}

// AudioQueueStart begins playing or recording audio.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueStart(_:_:)
func AudioQueueStart(inAQ AudioQueueRef, inStartTime *coreaudiotypes.AudioTimeStamp) int32 {
	result, callErr := tryAudioQueueStart(inAQ, inStartTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioQueueStop func(inAQ AudioQueueRef, inImmediate bool) int32
var _audioQueueStopErr error

func tryAudioQueueStop(inAQ AudioQueueRef, inImmediate bool) (int32, error) {
	if _audioQueueStop == nil {
		return 0, symbolCallError("AudioQueueStop", "10.5", _audioQueueStopErr)
	}
	return _audioQueueStop(inAQ, inImmediate), nil
}

// AudioQueueStop stops playing or recording audio.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioQueueStop(_:_:)
func AudioQueueStop(inAQ AudioQueueRef, inImmediate bool) int32 {
	result, callErr := tryAudioQueueStop(inAQ, inImmediate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioServicesAddSystemSoundCompletion func(inSystemSoundID SystemSoundID, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inCompletionRoutine AudioServicesSystemSoundCompletionProc, inClientData unsafe.Pointer) int32
var _audioServicesAddSystemSoundCompletionErr error

func tryAudioServicesAddSystemSoundCompletion(inSystemSoundID SystemSoundID, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inCompletionRoutine AudioServicesSystemSoundCompletionProc, inClientData unsafe.Pointer) (int32, error) {
	if _audioServicesAddSystemSoundCompletion == nil {
		return 0, symbolCallError("AudioServicesAddSystemSoundCompletion", "10.5", _audioServicesAddSystemSoundCompletionErr)
	}
	return _audioServicesAddSystemSoundCompletion(inSystemSoundID, inRunLoop, inRunLoopMode, inCompletionRoutine, inClientData), nil
}

// AudioServicesAddSystemSoundCompletion registers a callback function that is invoked when a specified system sound finishes playing.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesAddSystemSoundCompletion(_:_:_:_:_:)
func AudioServicesAddSystemSoundCompletion(inSystemSoundID SystemSoundID, inRunLoop corefoundation.CFRunLoopRef, inRunLoopMode corefoundation.CFStringRef, inCompletionRoutine AudioServicesSystemSoundCompletionProc, inClientData unsafe.Pointer) int32 {
	result, callErr := tryAudioServicesAddSystemSoundCompletion(inSystemSoundID, inRunLoop, inRunLoopMode, inCompletionRoutine, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioServicesCreateSystemSoundID func(inFileURL corefoundation.CFURLRef, outSystemSoundID *SystemSoundID) int32
var _audioServicesCreateSystemSoundIDErr error

func tryAudioServicesCreateSystemSoundID(inFileURL corefoundation.CFURLRef, outSystemSoundID *SystemSoundID) (int32, error) {
	if _audioServicesCreateSystemSoundID == nil {
		return 0, symbolCallError("AudioServicesCreateSystemSoundID", "10.5", _audioServicesCreateSystemSoundIDErr)
	}
	return _audioServicesCreateSystemSoundID(inFileURL, outSystemSoundID), nil
}

// AudioServicesCreateSystemSoundID creates a system sound object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesCreateSystemSoundID(_:_:)
func AudioServicesCreateSystemSoundID(inFileURL corefoundation.CFURLRef, outSystemSoundID *SystemSoundID) int32 {
	result, callErr := tryAudioServicesCreateSystemSoundID(inFileURL, outSystemSoundID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioServicesDisposeSystemSoundID func(inSystemSoundID SystemSoundID) int32
var _audioServicesDisposeSystemSoundIDErr error

func tryAudioServicesDisposeSystemSoundID(inSystemSoundID SystemSoundID) (int32, error) {
	if _audioServicesDisposeSystemSoundID == nil {
		return 0, symbolCallError("AudioServicesDisposeSystemSoundID", "10.5", _audioServicesDisposeSystemSoundIDErr)
	}
	return _audioServicesDisposeSystemSoundID(inSystemSoundID), nil
}

// AudioServicesDisposeSystemSoundID disposes of a system sound object and associated resources.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesDisposeSystemSoundID(_:)
func AudioServicesDisposeSystemSoundID(inSystemSoundID SystemSoundID) int32 {
	result, callErr := tryAudioServicesDisposeSystemSoundID(inSystemSoundID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioServicesGetProperty func(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioServicesGetPropertyErr error

func tryAudioServicesGetProperty(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioServicesGetProperty == nil {
		return 0, symbolCallError("AudioServicesGetProperty", "10.5", _audioServicesGetPropertyErr)
	}
	return _audioServicesGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData), nil
}

// AudioServicesGetProperty gets a specified System Sound Services property value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesGetProperty(_:_:_:_:_:)
func AudioServicesGetProperty(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioServicesGetProperty(inPropertyID, inSpecifierSize, inSpecifier, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioServicesGetPropertyInfo func(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertyDataSize *uint32, outWritable *bool) int32
var _audioServicesGetPropertyInfoErr error

func tryAudioServicesGetPropertyInfo(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertyDataSize *uint32, outWritable *bool) (int32, error) {
	if _audioServicesGetPropertyInfo == nil {
		return 0, symbolCallError("AudioServicesGetPropertyInfo", "10.5", _audioServicesGetPropertyInfoErr)
	}
	return _audioServicesGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize, outWritable), nil
}

// AudioServicesGetPropertyInfo gets information about a System Sound Services property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesGetPropertyInfo(_:_:_:_:_:)
func AudioServicesGetPropertyInfo(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, outPropertyDataSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioServicesGetPropertyInfo(inPropertyID, inSpecifierSize, inSpecifier, outPropertyDataSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioServicesPlayAlertSound func(inSystemSoundID SystemSoundID)
var _audioServicesPlayAlertSoundErr error

func tryAudioServicesPlayAlertSound(inSystemSoundID SystemSoundID) error {
	if _audioServicesPlayAlertSound == nil {
		return symbolCallError("AudioServicesPlayAlertSound", "10.5", _audioServicesPlayAlertSoundErr)
	}
	_audioServicesPlayAlertSound(inSystemSoundID)
	return nil
}

// AudioServicesPlayAlertSound plays a system sound as an alert.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSound(_:)
func AudioServicesPlayAlertSound(inSystemSoundID SystemSoundID) {
	if callErr := tryAudioServicesPlayAlertSound(inSystemSoundID); callErr != nil {
		panic(callErr)
	}
}

var _audioServicesPlayAlertSoundWithCompletion func(inSystemSoundID SystemSoundID)
var _audioServicesPlayAlertSoundWithCompletionErr error

func tryAudioServicesPlayAlertSoundWithCompletion(inSystemSoundID SystemSoundID) error {
	if _audioServicesPlayAlertSoundWithCompletion == nil {
		return symbolCallError("AudioServicesPlayAlertSoundWithCompletion", "10.11", _audioServicesPlayAlertSoundWithCompletionErr)
	}
	_audioServicesPlayAlertSoundWithCompletion(inSystemSoundID)
	return nil
}

// AudioServicesPlayAlertSoundWithCompletion.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlayAlertSoundWithCompletion(_:_:)
func AudioServicesPlayAlertSoundWithCompletion(inSystemSoundID SystemSoundID) {
	if callErr := tryAudioServicesPlayAlertSoundWithCompletion(inSystemSoundID); callErr != nil {
		panic(callErr)
	}
}

var _audioServicesPlaySystemSound func(inSystemSoundID SystemSoundID)
var _audioServicesPlaySystemSoundErr error

func tryAudioServicesPlaySystemSound(inSystemSoundID SystemSoundID) error {
	if _audioServicesPlaySystemSound == nil {
		return symbolCallError("AudioServicesPlaySystemSound", "10.5", _audioServicesPlaySystemSoundErr)
	}
	_audioServicesPlaySystemSound(inSystemSoundID)
	return nil
}

// AudioServicesPlaySystemSound plays a system sound object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSound(_:)
func AudioServicesPlaySystemSound(inSystemSoundID SystemSoundID) {
	if callErr := tryAudioServicesPlaySystemSound(inSystemSoundID); callErr != nil {
		panic(callErr)
	}
}

var _audioServicesPlaySystemSoundWithCompletion func(inSystemSoundID SystemSoundID)
var _audioServicesPlaySystemSoundWithCompletionErr error

func tryAudioServicesPlaySystemSoundWithCompletion(inSystemSoundID SystemSoundID) error {
	if _audioServicesPlaySystemSoundWithCompletion == nil {
		return symbolCallError("AudioServicesPlaySystemSoundWithCompletion", "10.11", _audioServicesPlaySystemSoundWithCompletionErr)
	}
	_audioServicesPlaySystemSoundWithCompletion(inSystemSoundID)
	return nil
}

// AudioServicesPlaySystemSoundWithCompletion.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesPlaySystemSoundWithCompletion(_:_:)
func AudioServicesPlaySystemSoundWithCompletion(inSystemSoundID SystemSoundID) {
	if callErr := tryAudioServicesPlaySystemSoundWithCompletion(inSystemSoundID); callErr != nil {
		panic(callErr)
	}
}

var _audioServicesRemoveSystemSoundCompletion func(inSystemSoundID SystemSoundID)
var _audioServicesRemoveSystemSoundCompletionErr error

func tryAudioServicesRemoveSystemSoundCompletion(inSystemSoundID SystemSoundID) error {
	if _audioServicesRemoveSystemSoundCompletion == nil {
		return symbolCallError("AudioServicesRemoveSystemSoundCompletion", "10.5", _audioServicesRemoveSystemSoundCompletionErr)
	}
	_audioServicesRemoveSystemSoundCompletion(inSystemSoundID)
	return nil
}

// AudioServicesRemoveSystemSoundCompletion unregisters any completion callback functions that were registered for a specified system sound.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesRemoveSystemSoundCompletion(_:)
func AudioServicesRemoveSystemSoundCompletion(inSystemSoundID SystemSoundID) {
	if callErr := tryAudioServicesRemoveSystemSoundCompletion(inSystemSoundID); callErr != nil {
		panic(callErr)
	}
}

var _audioServicesSetProperty func(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioServicesSetPropertyErr error

func tryAudioServicesSetProperty(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioServicesSetProperty == nil {
		return 0, symbolCallError("AudioServicesSetProperty", "10.5", _audioServicesSetPropertyErr)
	}
	return _audioServicesSetProperty(inPropertyID, inSpecifierSize, inSpecifier, inPropertyDataSize, inPropertyData), nil
}

// AudioServicesSetProperty sets the value for a specified System Sound Services property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioServicesSetProperty(_:_:_:_:_:)
func AudioServicesSetProperty(inPropertyID AudioServicesPropertyID, inSpecifierSize uint32, inSpecifier unsafe.Pointer, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioServicesSetProperty(inPropertyID, inSpecifierSize, inSpecifier, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitAddPropertyListener func(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) int32
var _audioUnitAddPropertyListenerErr error

func tryAudioUnitAddPropertyListener(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) (int32, error) {
	if _audioUnitAddPropertyListener == nil {
		return 0, symbolCallError("AudioUnitAddPropertyListener", "10.0", _audioUnitAddPropertyListenerErr)
	}
	return _audioUnitAddPropertyListener(inUnit, inID, inProc, inProcUserData), nil
}

// AudioUnitAddPropertyListener registers a callback to receive audio unit property change notifications.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddPropertyListener(_:_:_:_:)
func AudioUnitAddPropertyListener(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioUnitAddPropertyListener(inUnit, inID, inProc, inProcUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitAddRenderNotify func(inUnit AudioUnit, inProc AURenderCallback, inProcUserData unsafe.Pointer) int32
var _audioUnitAddRenderNotifyErr error

func tryAudioUnitAddRenderNotify(inUnit AudioUnit, inProc AURenderCallback, inProcUserData unsafe.Pointer) (int32, error) {
	if _audioUnitAddRenderNotify == nil {
		return 0, symbolCallError("AudioUnitAddRenderNotify", "10.2", _audioUnitAddRenderNotifyErr)
	}
	return _audioUnitAddRenderNotify(inUnit, inProc, inProcUserData), nil
}

// AudioUnitAddRenderNotify registers a callback to receive audio unit render notifications.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitAddRenderNotify(_:_:_:)
func AudioUnitAddRenderNotify(inUnit AudioUnit, inProc AURenderCallback, inProcUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioUnitAddRenderNotify(inUnit, inProc, inProcUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitExtensionCopyComponentList func(extensionIdentifier corefoundation.CFStringRef) corefoundation.CFArrayRef
var _audioUnitExtensionCopyComponentListErr error

func tryAudioUnitExtensionCopyComponentList(extensionIdentifier corefoundation.CFStringRef) (corefoundation.CFArrayRef, error) {
	if _audioUnitExtensionCopyComponentList == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("AudioUnitExtensionCopyComponentList", "10.13", _audioUnitExtensionCopyComponentListErr)
	}
	return _audioUnitExtensionCopyComponentList(extensionIdentifier), nil
}

// AudioUnitExtensionCopyComponentList returns the component registrations for a given audio unit extension.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExtensionCopyComponentList(_:)
func AudioUnitExtensionCopyComponentList(extensionIdentifier corefoundation.CFStringRef) corefoundation.CFArrayRef {
	result, callErr := tryAudioUnitExtensionCopyComponentList(extensionIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitExtensionSetComponentList func(extensionIdentifier corefoundation.CFStringRef, audioComponentInfo corefoundation.CFArrayRef) int32
var _audioUnitExtensionSetComponentListErr error

func tryAudioUnitExtensionSetComponentList(extensionIdentifier corefoundation.CFStringRef, audioComponentInfo corefoundation.CFArrayRef) (int32, error) {
	if _audioUnitExtensionSetComponentList == nil {
		return 0, symbolCallError("AudioUnitExtensionSetComponentList", "10.13", _audioUnitExtensionSetComponentListErr)
	}
	return _audioUnitExtensionSetComponentList(extensionIdentifier, audioComponentInfo), nil
}

// AudioUnitExtensionSetComponentList allows the implementor of an audio unit extension to dynamically modify the list of component registrations for the extension.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitExtensionSetComponentList(_:_:)
func AudioUnitExtensionSetComponentList(extensionIdentifier corefoundation.CFStringRef, audioComponentInfo corefoundation.CFArrayRef) int32 {
	result, callErr := tryAudioUnitExtensionSetComponentList(extensionIdentifier, audioComponentInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitGetParameter func(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, outValue *AudioUnitParameterValue) int32
var _audioUnitGetParameterErr error

func tryAudioUnitGetParameter(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, outValue *AudioUnitParameterValue) (int32, error) {
	if _audioUnitGetParameter == nil {
		return 0, symbolCallError("AudioUnitGetParameter", "10.0", _audioUnitGetParameterErr)
	}
	return _audioUnitGetParameter(inUnit, inID, inScope, inElement, outValue), nil
}

// AudioUnitGetParameter gets the value of an audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetParameter(_:_:_:_:_:)
func AudioUnitGetParameter(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, outValue *AudioUnitParameterValue) int32 {
	result, callErr := tryAudioUnitGetParameter(inUnit, inID, inScope, inElement, outValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitGetProperty func(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outData unsafe.Pointer, ioDataSize *uint32) int32
var _audioUnitGetPropertyErr error

func tryAudioUnitGetProperty(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outData unsafe.Pointer, ioDataSize *uint32) (int32, error) {
	if _audioUnitGetProperty == nil {
		return 0, symbolCallError("AudioUnitGetProperty", "10.0", _audioUnitGetPropertyErr)
	}
	return _audioUnitGetProperty(inUnit, inID, inScope, inElement, outData, ioDataSize), nil
}

// AudioUnitGetProperty gets the value of an audio unit property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetProperty(_:_:_:_:_:_:)
func AudioUnitGetProperty(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outData unsafe.Pointer, ioDataSize *uint32) int32 {
	result, callErr := tryAudioUnitGetProperty(inUnit, inID, inScope, inElement, outData, ioDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitGetPropertyInfo func(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outDataSize *uint32, outWritable *bool) int32
var _audioUnitGetPropertyInfoErr error

func tryAudioUnitGetPropertyInfo(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outDataSize *uint32, outWritable *bool) (int32, error) {
	if _audioUnitGetPropertyInfo == nil {
		return 0, symbolCallError("AudioUnitGetPropertyInfo", "10.0", _audioUnitGetPropertyInfoErr)
	}
	return _audioUnitGetPropertyInfo(inUnit, inID, inScope, inElement, outDataSize, outWritable), nil
}

// AudioUnitGetPropertyInfo gets information about an audio unit property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitGetPropertyInfo(_:_:_:_:_:_:)
func AudioUnitGetPropertyInfo(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, outDataSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioUnitGetPropertyInfo(inUnit, inID, inScope, inElement, outDataSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitInitialize func(inUnit AudioUnit) int32
var _audioUnitInitializeErr error

func tryAudioUnitInitialize(inUnit AudioUnit) (int32, error) {
	if _audioUnitInitialize == nil {
		return 0, symbolCallError("AudioUnitInitialize", "10.0", _audioUnitInitializeErr)
	}
	return _audioUnitInitialize(inUnit), nil
}

// AudioUnitInitialize initializes an audio unit
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitInitialize(_:)
func AudioUnitInitialize(inUnit AudioUnit) int32 {
	result, callErr := tryAudioUnitInitialize(inUnit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitProcess func(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32
var _audioUnitProcessErr error

func tryAudioUnitProcess(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) (int32, error) {
	if _audioUnitProcess == nil {
		return 0, symbolCallError("AudioUnitProcess", "10.7", _audioUnitProcessErr)
	}
	return _audioUnitProcess(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, ioData), nil
}

// AudioUnitProcess.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcess(_:_:_:_:_:)
func AudioUnitProcess(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32 {
	result, callErr := tryAudioUnitProcess(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, ioData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitProcessMultiple func(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inNumberFrames uint32, inNumberInputBufferLists uint32, inInputBufferLists *objc.ID, inNumberOutputBufferLists uint32, ioOutputBufferLists *objc.ID) int32
var _audioUnitProcessMultipleErr error

func tryAudioUnitProcessMultiple(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inNumberFrames uint32, inNumberInputBufferLists uint32, inInputBufferLists *objc.ID, inNumberOutputBufferLists uint32, ioOutputBufferLists *objc.ID) (int32, error) {
	if _audioUnitProcessMultiple == nil {
		return 0, symbolCallError("AudioUnitProcessMultiple", "10.7", _audioUnitProcessMultipleErr)
	}
	return _audioUnitProcessMultiple(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, inNumberInputBufferLists, inInputBufferLists, inNumberOutputBufferLists, ioOutputBufferLists), nil
}

// AudioUnitProcessMultiple.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitProcessMultiple(_:_:_:_:_:_:_:_:)
func AudioUnitProcessMultiple(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inNumberFrames uint32, inNumberInputBufferLists uint32, inInputBufferLists *objc.ID, inNumberOutputBufferLists uint32, ioOutputBufferLists *objc.ID) int32 {
	result, callErr := tryAudioUnitProcessMultiple(inUnit, ioActionFlags, inTimeStamp, inNumberFrames, inNumberInputBufferLists, inInputBufferLists, inNumberOutputBufferLists, ioOutputBufferLists)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitRemovePropertyListenerWithUserData func(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) int32
var _audioUnitRemovePropertyListenerWithUserDataErr error

func tryAudioUnitRemovePropertyListenerWithUserData(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) (int32, error) {
	if _audioUnitRemovePropertyListenerWithUserData == nil {
		return 0, symbolCallError("AudioUnitRemovePropertyListenerWithUserData", "10.5", _audioUnitRemovePropertyListenerWithUserDataErr)
	}
	return _audioUnitRemovePropertyListenerWithUserData(inUnit, inID, inProc, inProcUserData), nil
}

// AudioUnitRemovePropertyListenerWithUserData unregisters a previously-registered property listener callback function.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemovePropertyListenerWithUserData(_:_:_:_:)
func AudioUnitRemovePropertyListenerWithUserData(inUnit AudioUnit, inID AudioUnitPropertyID, inProc AudioUnitPropertyListenerProc, inProcUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioUnitRemovePropertyListenerWithUserData(inUnit, inID, inProc, inProcUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitRemoveRenderNotify func(inUnit AudioUnit, inProc AURenderCallback, inProcUserData unsafe.Pointer) int32
var _audioUnitRemoveRenderNotifyErr error

func tryAudioUnitRemoveRenderNotify(inUnit AudioUnit, inProc AURenderCallback, inProcUserData unsafe.Pointer) (int32, error) {
	if _audioUnitRemoveRenderNotify == nil {
		return 0, symbolCallError("AudioUnitRemoveRenderNotify", "10.2", _audioUnitRemoveRenderNotifyErr)
	}
	return _audioUnitRemoveRenderNotify(inUnit, inProc, inProcUserData), nil
}

// AudioUnitRemoveRenderNotify unregisters a previously-registered render listener callback function.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRemoveRenderNotify(_:_:_:)
func AudioUnitRemoveRenderNotify(inUnit AudioUnit, inProc AURenderCallback, inProcUserData unsafe.Pointer) int32 {
	result, callErr := tryAudioUnitRemoveRenderNotify(inUnit, inProc, inProcUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitRender func(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inOutputBusNumber uint32, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32
var _audioUnitRenderErr error

func tryAudioUnitRender(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inOutputBusNumber uint32, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) (int32, error) {
	if _audioUnitRender == nil {
		return 0, symbolCallError("AudioUnitRender", "10.2", _audioUnitRenderErr)
	}
	return _audioUnitRender(inUnit, ioActionFlags, inTimeStamp, inOutputBusNumber, inNumberFrames, ioData), nil
}

// AudioUnitRender initiates a rendering cycle for an audio unit.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitRender(_:_:_:_:_:_:)
func AudioUnitRender(inUnit AudioUnit, ioActionFlags *AudioUnitRenderActionFlags, inTimeStamp *coreaudiotypes.AudioTimeStamp, inOutputBusNumber uint32, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32 {
	result, callErr := tryAudioUnitRender(inUnit, ioActionFlags, inTimeStamp, inOutputBusNumber, inNumberFrames, ioData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitReset func(inUnit AudioUnit, inScope AudioUnitScope, inElement AudioUnitElement) int32
var _audioUnitResetErr error

func tryAudioUnitReset(inUnit AudioUnit, inScope AudioUnitScope, inElement AudioUnitElement) (int32, error) {
	if _audioUnitReset == nil {
		return 0, symbolCallError("AudioUnitReset", "10.0", _audioUnitResetErr)
	}
	return _audioUnitReset(inUnit, inScope, inElement), nil
}

// AudioUnitReset resets an audio unit’s render state.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitReset(_:_:_:)
func AudioUnitReset(inUnit AudioUnit, inScope AudioUnitScope, inElement AudioUnitElement) int32 {
	result, callErr := tryAudioUnitReset(inUnit, inScope, inElement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitScheduleParameters func(inUnit AudioUnit, inParameterEvent *AudioUnitParameterEvent, inNumParamEvents uint32) int32
var _audioUnitScheduleParametersErr error

func tryAudioUnitScheduleParameters(inUnit AudioUnit, inParameterEvent *AudioUnitParameterEvent, inNumParamEvents uint32) (int32, error) {
	if _audioUnitScheduleParameters == nil {
		return 0, symbolCallError("AudioUnitScheduleParameters", "10.2", _audioUnitScheduleParametersErr)
	}
	return _audioUnitScheduleParameters(inUnit, inParameterEvent, inNumParamEvents), nil
}

// AudioUnitScheduleParameters schedules changes to the value of an audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitScheduleParameters(_:_:_:)
func AudioUnitScheduleParameters(inUnit AudioUnit, inParameterEvent *AudioUnitParameterEvent, inNumParamEvents uint32) int32 {
	result, callErr := tryAudioUnitScheduleParameters(inUnit, inParameterEvent, inNumParamEvents)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitSetParameter func(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, inValue AudioUnitParameterValue, inBufferOffsetInFrames uint32) int32
var _audioUnitSetParameterErr error

func tryAudioUnitSetParameter(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, inValue AudioUnitParameterValue, inBufferOffsetInFrames uint32) (int32, error) {
	if _audioUnitSetParameter == nil {
		return 0, symbolCallError("AudioUnitSetParameter", "10.0", _audioUnitSetParameterErr)
	}
	return _audioUnitSetParameter(inUnit, inID, inScope, inElement, inValue, inBufferOffsetInFrames), nil
}

// AudioUnitSetParameter sets the value of an audio unit parameter.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetParameter(_:_:_:_:_:_:)
func AudioUnitSetParameter(inUnit AudioUnit, inID AudioUnitParameterID, inScope AudioUnitScope, inElement AudioUnitElement, inValue AudioUnitParameterValue, inBufferOffsetInFrames uint32) int32 {
	result, callErr := tryAudioUnitSetParameter(inUnit, inID, inScope, inElement, inValue, inBufferOffsetInFrames)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitSetProperty func(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, inData unsafe.Pointer, inDataSize uint32) int32
var _audioUnitSetPropertyErr error

func tryAudioUnitSetProperty(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, inData unsafe.Pointer, inDataSize uint32) (int32, error) {
	if _audioUnitSetProperty == nil {
		return 0, symbolCallError("AudioUnitSetProperty", "10.0", _audioUnitSetPropertyErr)
	}
	return _audioUnitSetProperty(inUnit, inID, inScope, inElement, inData, inDataSize), nil
}

// AudioUnitSetProperty sets the value of an audio unit property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitSetProperty(_:_:_:_:_:_:)
func AudioUnitSetProperty(inUnit AudioUnit, inID AudioUnitPropertyID, inScope AudioUnitScope, inElement AudioUnitElement, inData unsafe.Pointer, inDataSize uint32) int32 {
	result, callErr := tryAudioUnitSetProperty(inUnit, inID, inScope, inElement, inData, inDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioUnitUninitialize func(inUnit AudioUnit) int32
var _audioUnitUninitializeErr error

func tryAudioUnitUninitialize(inUnit AudioUnit) (int32, error) {
	if _audioUnitUninitialize == nil {
		return 0, symbolCallError("AudioUnitUninitialize", "10.0", _audioUnitUninitializeErr)
	}
	return _audioUnitUninitialize(inUnit), nil
}

// AudioUnitUninitialize uninitializes an audio unit.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioUnitUninitialize(_:)
func AudioUnitUninitialize(inUnit AudioUnit) int32 {
	result, callErr := tryAudioUnitUninitialize(inUnit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioWorkIntervalCreate func(name string, clock os.Os_clockid_t, attr os.Os_workgroup_attr_t) os.Os_workgroup_interval_t
var _audioWorkIntervalCreateErr error

func tryAudioWorkIntervalCreate(name string, clock os.Os_clockid_t, attr os.Os_workgroup_attr_t) (os.Os_workgroup_interval_t, error) {
	if _audioWorkIntervalCreate == nil {
		return *new(os.Os_workgroup_interval_t), symbolCallError("AudioWorkIntervalCreate", "11.0", _audioWorkIntervalCreateErr)
	}
	return _audioWorkIntervalCreate(name, clock, attr), nil
}

// AudioWorkIntervalCreate creates a new interval workgroup for managing real-time audio threads.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AudioWorkIntervalCreate
func AudioWorkIntervalCreate(name string, clock os.Os_clockid_t, attr os.Os_workgroup_attr_t) os.Os_workgroup_interval_t {
	result, callErr := tryAudioWorkIntervalCreate(name, clock, attr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockAddListener func(inCAClock CAClockRef, inListenerProc CAClockListenerProc, inUserData unsafe.Pointer) int32
var _cAClockAddListenerErr error

func tryCAClockAddListener(inCAClock CAClockRef, inListenerProc CAClockListenerProc, inUserData unsafe.Pointer) (int32, error) {
	if _cAClockAddListener == nil {
		return 0, symbolCallError("CAClockAddListener", "10.4", _cAClockAddListenerErr)
	}
	return _cAClockAddListener(inCAClock, inListenerProc, inUserData), nil
}

// CAClockAddListener.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockAddListener(_:_:_:)
func CAClockAddListener(inCAClock CAClockRef, inListenerProc CAClockListenerProc, inUserData unsafe.Pointer) int32 {
	result, callErr := tryCAClockAddListener(inCAClock, inListenerProc, inUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockArm func(inCAClock CAClockRef) int32
var _cAClockArmErr error

func tryCAClockArm(inCAClock CAClockRef) (int32, error) {
	if _cAClockArm == nil {
		return 0, symbolCallError("CAClockArm", "10.4", _cAClockArmErr)
	}
	return _cAClockArm(inCAClock), nil
}

// CAClockArm.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockArm(_:)
func CAClockArm(inCAClock CAClockRef) int32 {
	result, callErr := tryCAClockArm(inCAClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockBarBeatTimeToBeats func(inCAClock CAClockRef, inBarBeatTime *CABarBeatTime, outBeats *CAClockBeats) int32
var _cAClockBarBeatTimeToBeatsErr error

func tryCAClockBarBeatTimeToBeats(inCAClock CAClockRef, inBarBeatTime *CABarBeatTime, outBeats *CAClockBeats) (int32, error) {
	if _cAClockBarBeatTimeToBeats == nil {
		return 0, symbolCallError("CAClockBarBeatTimeToBeats", "10.4", _cAClockBarBeatTimeToBeatsErr)
	}
	return _cAClockBarBeatTimeToBeats(inCAClock, inBarBeatTime, outBeats), nil
}

// CAClockBarBeatTimeToBeats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockBarBeatTimeToBeats(_:_:_:)
func CAClockBarBeatTimeToBeats(inCAClock CAClockRef, inBarBeatTime *CABarBeatTime, outBeats *CAClockBeats) int32 {
	result, callErr := tryCAClockBarBeatTimeToBeats(inCAClock, inBarBeatTime, outBeats)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockBeatsToBarBeatTime func(inCAClock CAClockRef, inBeats CAClockBeats, inSubbeatDivisor uint16, outBarBeatTime *CABarBeatTime) int32
var _cAClockBeatsToBarBeatTimeErr error

func tryCAClockBeatsToBarBeatTime(inCAClock CAClockRef, inBeats CAClockBeats, inSubbeatDivisor uint16, outBarBeatTime *CABarBeatTime) (int32, error) {
	if _cAClockBeatsToBarBeatTime == nil {
		return 0, symbolCallError("CAClockBeatsToBarBeatTime", "10.4", _cAClockBeatsToBarBeatTimeErr)
	}
	return _cAClockBeatsToBarBeatTime(inCAClock, inBeats, inSubbeatDivisor, outBarBeatTime), nil
}

// CAClockBeatsToBarBeatTime.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockBeatsToBarBeatTime(_:_:_:_:)
func CAClockBeatsToBarBeatTime(inCAClock CAClockRef, inBeats CAClockBeats, inSubbeatDivisor uint16, outBarBeatTime *CABarBeatTime) int32 {
	result, callErr := tryCAClockBeatsToBarBeatTime(inCAClock, inBeats, inSubbeatDivisor, outBarBeatTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockDisarm func(inCAClock CAClockRef) int32
var _cAClockDisarmErr error

func tryCAClockDisarm(inCAClock CAClockRef) (int32, error) {
	if _cAClockDisarm == nil {
		return 0, symbolCallError("CAClockDisarm", "10.4", _cAClockDisarmErr)
	}
	return _cAClockDisarm(inCAClock), nil
}

// CAClockDisarm.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockDisarm(_:)
func CAClockDisarm(inCAClock CAClockRef) int32 {
	result, callErr := tryCAClockDisarm(inCAClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockDispose func(inCAClock CAClockRef) int32
var _cAClockDisposeErr error

func tryCAClockDispose(inCAClock CAClockRef) (int32, error) {
	if _cAClockDispose == nil {
		return 0, symbolCallError("CAClockDispose", "10.4", _cAClockDisposeErr)
	}
	return _cAClockDispose(inCAClock), nil
}

// CAClockDispose.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockDispose(_:)
func CAClockDispose(inCAClock CAClockRef) int32 {
	result, callErr := tryCAClockDispose(inCAClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockGetCurrentTempo func(inCAClock CAClockRef, outTempo *CAClockTempo, outTimestamp *CAClockTime) int32
var _cAClockGetCurrentTempoErr error

func tryCAClockGetCurrentTempo(inCAClock CAClockRef, outTempo *CAClockTempo, outTimestamp *CAClockTime) (int32, error) {
	if _cAClockGetCurrentTempo == nil {
		return 0, symbolCallError("CAClockGetCurrentTempo", "10.4", _cAClockGetCurrentTempoErr)
	}
	return _cAClockGetCurrentTempo(inCAClock, outTempo, outTimestamp), nil
}

// CAClockGetCurrentTempo.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockGetCurrentTempo(_:_:_:)
func CAClockGetCurrentTempo(inCAClock CAClockRef, outTempo *CAClockTempo, outTimestamp *CAClockTime) int32 {
	result, callErr := tryCAClockGetCurrentTempo(inCAClock, outTempo, outTimestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockGetCurrentTime func(inCAClock CAClockRef, inTimeFormat CAClockTimeFormat, outTime *CAClockTime) int32
var _cAClockGetCurrentTimeErr error

func tryCAClockGetCurrentTime(inCAClock CAClockRef, inTimeFormat CAClockTimeFormat, outTime *CAClockTime) (int32, error) {
	if _cAClockGetCurrentTime == nil {
		return 0, symbolCallError("CAClockGetCurrentTime", "10.4", _cAClockGetCurrentTimeErr)
	}
	return _cAClockGetCurrentTime(inCAClock, inTimeFormat, outTime), nil
}

// CAClockGetCurrentTime.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockGetCurrentTime(_:_:_:)
func CAClockGetCurrentTime(inCAClock CAClockRef, inTimeFormat CAClockTimeFormat, outTime *CAClockTime) int32 {
	result, callErr := tryCAClockGetCurrentTime(inCAClock, inTimeFormat, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockGetPlayRate func(inCAClock CAClockRef, outPlayRate *float64) int32
var _cAClockGetPlayRateErr error

func tryCAClockGetPlayRate(inCAClock CAClockRef, outPlayRate *float64) (int32, error) {
	if _cAClockGetPlayRate == nil {
		return 0, symbolCallError("CAClockGetPlayRate", "10.4", _cAClockGetPlayRateErr)
	}
	return _cAClockGetPlayRate(inCAClock, outPlayRate), nil
}

// CAClockGetPlayRate.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockGetPlayRate(_:_:)
func CAClockGetPlayRate(inCAClock CAClockRef, outPlayRate *float64) int32 {
	result, callErr := tryCAClockGetPlayRate(inCAClock, outPlayRate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockGetProperty func(inCAClock CAClockRef, inPropertyID CAClockPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _cAClockGetPropertyErr error

func tryCAClockGetProperty(inCAClock CAClockRef, inPropertyID CAClockPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _cAClockGetProperty == nil {
		return 0, symbolCallError("CAClockGetProperty", "10.4", _cAClockGetPropertyErr)
	}
	return _cAClockGetProperty(inCAClock, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// CAClockGetProperty.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockGetProperty(_:_:_:_:)
func CAClockGetProperty(inCAClock CAClockRef, inPropertyID CAClockPropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryCAClockGetProperty(inCAClock, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockGetPropertyInfo func(inCAClock CAClockRef, inPropertyID CAClockPropertyID, outSize *uint32, outWritable *bool) int32
var _cAClockGetPropertyInfoErr error

func tryCAClockGetPropertyInfo(inCAClock CAClockRef, inPropertyID CAClockPropertyID, outSize *uint32, outWritable *bool) (int32, error) {
	if _cAClockGetPropertyInfo == nil {
		return 0, symbolCallError("CAClockGetPropertyInfo", "10.4", _cAClockGetPropertyInfoErr)
	}
	return _cAClockGetPropertyInfo(inCAClock, inPropertyID, outSize, outWritable), nil
}

// CAClockGetPropertyInfo.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockGetPropertyInfo(_:_:_:_:)
func CAClockGetPropertyInfo(inCAClock CAClockRef, inPropertyID CAClockPropertyID, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryCAClockGetPropertyInfo(inCAClock, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockGetStartTime func(inCAClock CAClockRef, inTimeFormat CAClockTimeFormat, outTime *CAClockTime) int32
var _cAClockGetStartTimeErr error

func tryCAClockGetStartTime(inCAClock CAClockRef, inTimeFormat CAClockTimeFormat, outTime *CAClockTime) (int32, error) {
	if _cAClockGetStartTime == nil {
		return 0, symbolCallError("CAClockGetStartTime", "10.4", _cAClockGetStartTimeErr)
	}
	return _cAClockGetStartTime(inCAClock, inTimeFormat, outTime), nil
}

// CAClockGetStartTime.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockGetStartTime(_:_:_:)
func CAClockGetStartTime(inCAClock CAClockRef, inTimeFormat CAClockTimeFormat, outTime *CAClockTime) int32 {
	result, callErr := tryCAClockGetStartTime(inCAClock, inTimeFormat, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockNew func(inReservedFlags uint32, outCAClock *CAClockRef) int32
var _cAClockNewErr error

func tryCAClockNew(inReservedFlags uint32, outCAClock *CAClockRef) (int32, error) {
	if _cAClockNew == nil {
		return 0, symbolCallError("CAClockNew", "10.4", _cAClockNewErr)
	}
	return _cAClockNew(inReservedFlags, outCAClock), nil
}

// CAClockNew.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockNew(_:_:)
func CAClockNew(inReservedFlags uint32, outCAClock *CAClockRef) int32 {
	result, callErr := tryCAClockNew(inReservedFlags, outCAClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockParseMIDI func(inCAClock CAClockRef, inMIDIPacketList *coremidi.MIDIPacketList) int32
var _cAClockParseMIDIErr error

func tryCAClockParseMIDI(inCAClock CAClockRef, inMIDIPacketList *coremidi.MIDIPacketList) (int32, error) {
	if _cAClockParseMIDI == nil {
		return 0, symbolCallError("CAClockParseMIDI", "10.5", _cAClockParseMIDIErr)
	}
	return _cAClockParseMIDI(inCAClock, inMIDIPacketList), nil
}

// CAClockParseMIDI.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockParseMIDI(_:_:)
func CAClockParseMIDI(inCAClock CAClockRef, inMIDIPacketList *coremidi.MIDIPacketList) int32 {
	result, callErr := tryCAClockParseMIDI(inCAClock, inMIDIPacketList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockRemoveListener func(inCAClock CAClockRef, inListenerProc CAClockListenerProc, inUserData unsafe.Pointer) int32
var _cAClockRemoveListenerErr error

func tryCAClockRemoveListener(inCAClock CAClockRef, inListenerProc CAClockListenerProc, inUserData unsafe.Pointer) (int32, error) {
	if _cAClockRemoveListener == nil {
		return 0, symbolCallError("CAClockRemoveListener", "10.4", _cAClockRemoveListenerErr)
	}
	return _cAClockRemoveListener(inCAClock, inListenerProc, inUserData), nil
}

// CAClockRemoveListener.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockRemoveListener(_:_:_:)
func CAClockRemoveListener(inCAClock CAClockRef, inListenerProc CAClockListenerProc, inUserData unsafe.Pointer) int32 {
	result, callErr := tryCAClockRemoveListener(inCAClock, inListenerProc, inUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockSMPTETimeToSeconds func(inCAClock CAClockRef, inSMPTETime *coreaudiotypes.SMPTETime, outSeconds *CAClockSeconds) int32
var _cAClockSMPTETimeToSecondsErr error

func tryCAClockSMPTETimeToSeconds(inCAClock CAClockRef, inSMPTETime *coreaudiotypes.SMPTETime, outSeconds *CAClockSeconds) (int32, error) {
	if _cAClockSMPTETimeToSeconds == nil {
		return 0, symbolCallError("CAClockSMPTETimeToSeconds", "10.4", _cAClockSMPTETimeToSecondsErr)
	}
	return _cAClockSMPTETimeToSeconds(inCAClock, inSMPTETime, outSeconds), nil
}

// CAClockSMPTETimeToSeconds.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSMPTETimeToSeconds(_:_:_:)
func CAClockSMPTETimeToSeconds(inCAClock CAClockRef, inSMPTETime *coreaudiotypes.SMPTETime, outSeconds *CAClockSeconds) int32 {
	result, callErr := tryCAClockSMPTETimeToSeconds(inCAClock, inSMPTETime, outSeconds)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockSecondsToSMPTETime func(inCAClock CAClockRef, inSeconds CAClockSeconds, inSubframeDivisor uint16, outSMPTETime *coreaudiotypes.SMPTETime) int32
var _cAClockSecondsToSMPTETimeErr error

func tryCAClockSecondsToSMPTETime(inCAClock CAClockRef, inSeconds CAClockSeconds, inSubframeDivisor uint16, outSMPTETime *coreaudiotypes.SMPTETime) (int32, error) {
	if _cAClockSecondsToSMPTETime == nil {
		return 0, symbolCallError("CAClockSecondsToSMPTETime", "10.4", _cAClockSecondsToSMPTETimeErr)
	}
	return _cAClockSecondsToSMPTETime(inCAClock, inSeconds, inSubframeDivisor, outSMPTETime), nil
}

// CAClockSecondsToSMPTETime.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSecondsToSMPTETime(_:_:_:_:)
func CAClockSecondsToSMPTETime(inCAClock CAClockRef, inSeconds CAClockSeconds, inSubframeDivisor uint16, outSMPTETime *coreaudiotypes.SMPTETime) int32 {
	result, callErr := tryCAClockSecondsToSMPTETime(inCAClock, inSeconds, inSubframeDivisor, outSMPTETime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockSetCurrentTempo func(inCAClock CAClockRef, inTempo CAClockTempo, inTimestamp *CAClockTime) int32
var _cAClockSetCurrentTempoErr error

func tryCAClockSetCurrentTempo(inCAClock CAClockRef, inTempo CAClockTempo, inTimestamp *CAClockTime) (int32, error) {
	if _cAClockSetCurrentTempo == nil {
		return 0, symbolCallError("CAClockSetCurrentTempo", "10.4", _cAClockSetCurrentTempoErr)
	}
	return _cAClockSetCurrentTempo(inCAClock, inTempo, inTimestamp), nil
}

// CAClockSetCurrentTempo.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSetCurrentTempo(_:_:_:)
func CAClockSetCurrentTempo(inCAClock CAClockRef, inTempo CAClockTempo, inTimestamp *CAClockTime) int32 {
	result, callErr := tryCAClockSetCurrentTempo(inCAClock, inTempo, inTimestamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockSetCurrentTime func(inCAClock CAClockRef, inTime *CAClockTime) int32
var _cAClockSetCurrentTimeErr error

func tryCAClockSetCurrentTime(inCAClock CAClockRef, inTime *CAClockTime) (int32, error) {
	if _cAClockSetCurrentTime == nil {
		return 0, symbolCallError("CAClockSetCurrentTime", "10.4", _cAClockSetCurrentTimeErr)
	}
	return _cAClockSetCurrentTime(inCAClock, inTime), nil
}

// CAClockSetCurrentTime.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSetCurrentTime(_:_:)
func CAClockSetCurrentTime(inCAClock CAClockRef, inTime *CAClockTime) int32 {
	result, callErr := tryCAClockSetCurrentTime(inCAClock, inTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockSetPlayRate func(inCAClock CAClockRef, inPlayRate float64) int32
var _cAClockSetPlayRateErr error

func tryCAClockSetPlayRate(inCAClock CAClockRef, inPlayRate float64) (int32, error) {
	if _cAClockSetPlayRate == nil {
		return 0, symbolCallError("CAClockSetPlayRate", "10.4", _cAClockSetPlayRateErr)
	}
	return _cAClockSetPlayRate(inCAClock, inPlayRate), nil
}

// CAClockSetPlayRate.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSetPlayRate(_:_:)
func CAClockSetPlayRate(inCAClock CAClockRef, inPlayRate float64) int32 {
	result, callErr := tryCAClockSetPlayRate(inCAClock, inPlayRate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockSetProperty func(inCAClock CAClockRef, inPropertyID CAClockPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _cAClockSetPropertyErr error

func tryCAClockSetProperty(inCAClock CAClockRef, inPropertyID CAClockPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _cAClockSetProperty == nil {
		return 0, symbolCallError("CAClockSetProperty", "10.4", _cAClockSetPropertyErr)
	}
	return _cAClockSetProperty(inCAClock, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// CAClockSetProperty.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockSetProperty(_:_:_:_:)
func CAClockSetProperty(inCAClock CAClockRef, inPropertyID CAClockPropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryCAClockSetProperty(inCAClock, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockStart func(inCAClock CAClockRef) int32
var _cAClockStartErr error

func tryCAClockStart(inCAClock CAClockRef) (int32, error) {
	if _cAClockStart == nil {
		return 0, symbolCallError("CAClockStart", "10.4", _cAClockStartErr)
	}
	return _cAClockStart(inCAClock), nil
}

// CAClockStart.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockStart(_:)
func CAClockStart(inCAClock CAClockRef) int32 {
	result, callErr := tryCAClockStart(inCAClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockStop func(inCAClock CAClockRef) int32
var _cAClockStopErr error

func tryCAClockStop(inCAClock CAClockRef) (int32, error) {
	if _cAClockStop == nil {
		return 0, symbolCallError("CAClockStop", "10.4", _cAClockStopErr)
	}
	return _cAClockStop(inCAClock), nil
}

// CAClockStop.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockStop(_:)
func CAClockStop(inCAClock CAClockRef) int32 {
	result, callErr := tryCAClockStop(inCAClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAClockTranslateTime func(inCAClock CAClockRef, inTime *CAClockTime, inOutputTimeFormat CAClockTimeFormat, outTime *CAClockTime) int32
var _cAClockTranslateTimeErr error

func tryCAClockTranslateTime(inCAClock CAClockRef, inTime *CAClockTime, inOutputTimeFormat CAClockTimeFormat, outTime *CAClockTime) (int32, error) {
	if _cAClockTranslateTime == nil {
		return 0, symbolCallError("CAClockTranslateTime", "10.4", _cAClockTranslateTimeErr)
	}
	return _cAClockTranslateTime(inCAClock, inTime, inOutputTimeFormat, outTime), nil
}

// CAClockTranslateTime.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAClockTranslateTime(_:_:_:_:)
func CAClockTranslateTime(inCAClock CAClockRef, inTime *CAClockTime, inOutputTimeFormat CAClockTimeFormat, outTime *CAClockTime) int32 {
	result, callErr := tryCAClockTranslateTime(inCAClock, inTime, inOutputTimeFormat, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cAShow func(inObject unsafe.Pointer)
var _cAShowErr error

func tryCAShow(inObject unsafe.Pointer) error {
	if _cAShow == nil {
		return symbolCallError("CAShow", "10.2", _cAShowErr)
	}
	_cAShow(inObject)
	return nil
}

// CAShow prints the internal state of an object to `stdio`.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAShow(_:)
func CAShow(inObject unsafe.Pointer) {
	if callErr := tryCAShow(inObject); callErr != nil {
		panic(callErr)
	}
}

var _cAShowFile func(inObject unsafe.Pointer, inFile unsafe.Pointer)
var _cAShowFileErr error

func tryCAShowFile(inObject unsafe.Pointer, inFile unsafe.Pointer) error {
	if _cAShowFile == nil {
		return symbolCallError("CAShowFile", "10.2", _cAShowFileErr)
	}
	_cAShowFile(inObject, inFile)
	return nil
}

// CAShowFile prints the internal state of an object to a file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CAShowFile(_:_:)
func CAShowFile(inObject unsafe.Pointer, inFile unsafe.Pointer) {
	if callErr := tryCAShowFile(inObject, inFile); callErr != nil {
		panic(callErr)
	}
}

var _copyInstrumentInfoFromSoundBank func(inURL corefoundation.CFURLRef, outInstrumentInfo *corefoundation.CFArrayRef) int32
var _copyInstrumentInfoFromSoundBankErr error

func tryCopyInstrumentInfoFromSoundBank(inURL corefoundation.CFURLRef, outInstrumentInfo *corefoundation.CFArrayRef) (int32, error) {
	if _copyInstrumentInfoFromSoundBank == nil {
		return 0, symbolCallError("CopyInstrumentInfoFromSoundBank", "10.8", _copyInstrumentInfoFromSoundBankErr)
	}
	return _copyInstrumentInfoFromSoundBank(inURL, outInstrumentInfo), nil
}

// CopyInstrumentInfoFromSoundBank.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CopyInstrumentInfoFromSoundBank(_:_:)
func CopyInstrumentInfoFromSoundBank(inURL corefoundation.CFURLRef, outInstrumentInfo *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCopyInstrumentInfoFromSoundBank(inURL, outInstrumentInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _copyNameFromSoundBank func(inURL corefoundation.CFURLRef, outName *corefoundation.CFStringRef) int32
var _copyNameFromSoundBankErr error

func tryCopyNameFromSoundBank(inURL corefoundation.CFURLRef, outName *corefoundation.CFStringRef) (int32, error) {
	if _copyNameFromSoundBank == nil {
		return 0, symbolCallError("CopyNameFromSoundBank", "10.5", _copyNameFromSoundBankErr)
	}
	return _copyNameFromSoundBank(inURL, outName), nil
}

// CopyNameFromSoundBank copies the name of a sound bank from a sound bank file at a specified URL.
//
// See: https://developer.apple.com/documentation/AudioToolbox/CopyNameFromSoundBank(_:_:)
func CopyNameFromSoundBank(inURL corefoundation.CFURLRef, outName *corefoundation.CFStringRef) int32 {
	result, callErr := tryCopyNameFromSoundBank(inURL, outName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _disposeAUGraph func(inGraph AUGraph) int32
var _disposeAUGraphErr error

func tryDisposeAUGraph(inGraph AUGraph) (int32, error) {
	if _disposeAUGraph == nil {
		return 0, symbolCallError("DisposeAUGraph", "10.0", _disposeAUGraphErr)
	}
	return _disposeAUGraph(inGraph), nil
}

// DisposeAUGraph disposes of an audio processing graph.
//
// Deprecated: Deprecated since macOS 27.0. AUGraph is deprecated in favor of AVAudioEngine
//
// See: https://developer.apple.com/documentation/AudioToolbox/DisposeAUGraph(_:)
func DisposeAUGraph(inGraph AUGraph) int32 {
	result, callErr := tryDisposeAUGraph(inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _disposeMusicEventIterator func(inIterator MusicEventIterator) int32
var _disposeMusicEventIteratorErr error

func tryDisposeMusicEventIterator(inIterator MusicEventIterator) (int32, error) {
	if _disposeMusicEventIterator == nil {
		return 0, symbolCallError("DisposeMusicEventIterator", "10.0", _disposeMusicEventIteratorErr)
	}
	return _disposeMusicEventIterator(inIterator), nil
}

// DisposeMusicEventIterator disposes of a music event iterator.
//
// See: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicEventIterator(_:)
func DisposeMusicEventIterator(inIterator MusicEventIterator) int32 {
	result, callErr := tryDisposeMusicEventIterator(inIterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _disposeMusicPlayer func(inPlayer MusicPlayer) int32
var _disposeMusicPlayerErr error

func tryDisposeMusicPlayer(inPlayer MusicPlayer) (int32, error) {
	if _disposeMusicPlayer == nil {
		return 0, symbolCallError("DisposeMusicPlayer", "10.0", _disposeMusicPlayerErr)
	}
	return _disposeMusicPlayer(inPlayer), nil
}

// DisposeMusicPlayer disposes of a music player.
//
// See: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicPlayer(_:)
func DisposeMusicPlayer(inPlayer MusicPlayer) int32 {
	result, callErr := tryDisposeMusicPlayer(inPlayer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _disposeMusicSequence func(inSequence MusicSequence) int32
var _disposeMusicSequenceErr error

func tryDisposeMusicSequence(inSequence MusicSequence) (int32, error) {
	if _disposeMusicSequence == nil {
		return 0, symbolCallError("DisposeMusicSequence", "10.0", _disposeMusicSequenceErr)
	}
	return _disposeMusicSequence(inSequence), nil
}

// DisposeMusicSequence disposes of a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/DisposeMusicSequence(_:)
func DisposeMusicSequence(inSequence MusicSequence) int32 {
	result, callErr := tryDisposeMusicSequence(inSequence)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileCreateNew func(inParentDir *coreservices.FSRef, inFileName corefoundation.CFStringRef, inFileType AudioFileTypeID, inStreamDesc *coreaudiotypes.AudioStreamBasicDescription, inChannelLayout *coreaudiotypes.AudioChannelLayout, outExtAudioFile *ExtAudioFileRef) int32
var _extAudioFileCreateNewErr error

func tryExtAudioFileCreateNew(inParentDir *coreservices.FSRef, inFileName corefoundation.CFStringRef, inFileType AudioFileTypeID, inStreamDesc *coreaudiotypes.AudioStreamBasicDescription, inChannelLayout *coreaudiotypes.AudioChannelLayout, outExtAudioFile *ExtAudioFileRef) (int32, error) {
	if _extAudioFileCreateNew == nil {
		return 0, symbolCallError("ExtAudioFileCreateNew", "10.4", _extAudioFileCreateNewErr)
	}
	return _extAudioFileCreateNew(inParentDir, inFileName, inFileType, inStreamDesc, inChannelLayout, outExtAudioFile), nil
}

// ExtAudioFileCreateNew deprecated.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileCreateNew
func ExtAudioFileCreateNew(inParentDir *coreservices.FSRef, inFileName corefoundation.CFStringRef, inFileType AudioFileTypeID, inStreamDesc *coreaudiotypes.AudioStreamBasicDescription, inChannelLayout *coreaudiotypes.AudioChannelLayout, outExtAudioFile *ExtAudioFileRef) int32 {
	result, callErr := tryExtAudioFileCreateNew(inParentDir, inFileName, inFileType, inStreamDesc, inChannelLayout, outExtAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileCreateWithURL func(inURL corefoundation.CFURLRef, inFileType AudioFileTypeID, inStreamDesc *coreaudiotypes.AudioStreamBasicDescription, inChannelLayout *coreaudiotypes.AudioChannelLayout, inFlags uint32, outExtAudioFile *ExtAudioFileRef) int32
var _extAudioFileCreateWithURLErr error

func tryExtAudioFileCreateWithURL(inURL corefoundation.CFURLRef, inFileType AudioFileTypeID, inStreamDesc *coreaudiotypes.AudioStreamBasicDescription, inChannelLayout *coreaudiotypes.AudioChannelLayout, inFlags uint32, outExtAudioFile *ExtAudioFileRef) (int32, error) {
	if _extAudioFileCreateWithURL == nil {
		return 0, symbolCallError("ExtAudioFileCreateWithURL", "10.5", _extAudioFileCreateWithURLErr)
	}
	return _extAudioFileCreateWithURL(inURL, inFileType, inStreamDesc, inChannelLayout, inFlags, outExtAudioFile), nil
}

// ExtAudioFileCreateWithURL creates a new audio file and associates it with a new extended audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileCreateWithURL(_:_:_:_:_:_:)
func ExtAudioFileCreateWithURL(inURL corefoundation.CFURLRef, inFileType AudioFileTypeID, inStreamDesc *coreaudiotypes.AudioStreamBasicDescription, inChannelLayout *coreaudiotypes.AudioChannelLayout, inFlags uint32, outExtAudioFile *ExtAudioFileRef) int32 {
	result, callErr := tryExtAudioFileCreateWithURL(inURL, inFileType, inStreamDesc, inChannelLayout, inFlags, outExtAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileDispose func(inExtAudioFile ExtAudioFileRef) int32
var _extAudioFileDisposeErr error

func tryExtAudioFileDispose(inExtAudioFile ExtAudioFileRef) (int32, error) {
	if _extAudioFileDispose == nil {
		return 0, symbolCallError("ExtAudioFileDispose", "10.4", _extAudioFileDisposeErr)
	}
	return _extAudioFileDispose(inExtAudioFile), nil
}

// ExtAudioFileDispose disposes of an extended audio file object and closes the associated file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileDispose(_:)
func ExtAudioFileDispose(inExtAudioFile ExtAudioFileRef) int32 {
	result, callErr := tryExtAudioFileDispose(inExtAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileGetProperty func(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _extAudioFileGetPropertyErr error

func tryExtAudioFileGetProperty(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _extAudioFileGetProperty == nil {
		return 0, symbolCallError("ExtAudioFileGetProperty", "10.4", _extAudioFileGetPropertyErr)
	}
	return _extAudioFileGetProperty(inExtAudioFile, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// ExtAudioFileGetProperty gets a property value from an extended audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileGetProperty(_:_:_:_:)
func ExtAudioFileGetProperty(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryExtAudioFileGetProperty(inExtAudioFile, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileGetPropertyInfo func(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, outSize *uint32, outWritable *bool) int32
var _extAudioFileGetPropertyInfoErr error

func tryExtAudioFileGetPropertyInfo(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, outSize *uint32, outWritable *bool) (int32, error) {
	if _extAudioFileGetPropertyInfo == nil {
		return 0, symbolCallError("ExtAudioFileGetPropertyInfo", "10.4", _extAudioFileGetPropertyInfoErr)
	}
	return _extAudioFileGetPropertyInfo(inExtAudioFile, inPropertyID, outSize, outWritable), nil
}

// ExtAudioFileGetPropertyInfo gets information about an extended audio file object property.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileGetPropertyInfo(_:_:_:_:)
func ExtAudioFileGetPropertyInfo(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryExtAudioFileGetPropertyInfo(inExtAudioFile, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileOpen func(inFSRef *coreservices.FSRef, outExtAudioFile *ExtAudioFileRef) int32
var _extAudioFileOpenErr error

func tryExtAudioFileOpen(inFSRef *coreservices.FSRef, outExtAudioFile *ExtAudioFileRef) (int32, error) {
	if _extAudioFileOpen == nil {
		return 0, symbolCallError("ExtAudioFileOpen", "10.4", _extAudioFileOpenErr)
	}
	return _extAudioFileOpen(inFSRef, outExtAudioFile), nil
}

// ExtAudioFileOpen deprecated.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileOpen
func ExtAudioFileOpen(inFSRef *coreservices.FSRef, outExtAudioFile *ExtAudioFileRef) int32 {
	result, callErr := tryExtAudioFileOpen(inFSRef, outExtAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileOpenURL func(inURL corefoundation.CFURLRef, outExtAudioFile *ExtAudioFileRef) int32
var _extAudioFileOpenURLErr error

func tryExtAudioFileOpenURL(inURL corefoundation.CFURLRef, outExtAudioFile *ExtAudioFileRef) (int32, error) {
	if _extAudioFileOpenURL == nil {
		return 0, symbolCallError("ExtAudioFileOpenURL", "10.5", _extAudioFileOpenURLErr)
	}
	return _extAudioFileOpenURL(inURL, outExtAudioFile), nil
}

// ExtAudioFileOpenURL opens an existing audio file for reading, and associates it with a new extended audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileOpenURL(_:_:)
func ExtAudioFileOpenURL(inURL corefoundation.CFURLRef, outExtAudioFile *ExtAudioFileRef) int32 {
	result, callErr := tryExtAudioFileOpenURL(inURL, outExtAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileRead func(inExtAudioFile ExtAudioFileRef, ioNumberFrames *uint32, ioData *coreaudiotypes.AudioBufferList) int32
var _extAudioFileReadErr error

func tryExtAudioFileRead(inExtAudioFile ExtAudioFileRef, ioNumberFrames *uint32, ioData *coreaudiotypes.AudioBufferList) (int32, error) {
	if _extAudioFileRead == nil {
		return 0, symbolCallError("ExtAudioFileRead", "10.4", _extAudioFileReadErr)
	}
	return _extAudioFileRead(inExtAudioFile, ioNumberFrames, ioData), nil
}

// ExtAudioFileRead performs a synchronous, sequential read operation on an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileRead(_:_:_:)
func ExtAudioFileRead(inExtAudioFile ExtAudioFileRef, ioNumberFrames *uint32, ioData *coreaudiotypes.AudioBufferList) int32 {
	result, callErr := tryExtAudioFileRead(inExtAudioFile, ioNumberFrames, ioData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileSeek func(inExtAudioFile ExtAudioFileRef, inFrameOffset int64) int32
var _extAudioFileSeekErr error

func tryExtAudioFileSeek(inExtAudioFile ExtAudioFileRef, inFrameOffset int64) (int32, error) {
	if _extAudioFileSeek == nil {
		return 0, symbolCallError("ExtAudioFileSeek", "10.4", _extAudioFileSeekErr)
	}
	return _extAudioFileSeek(inExtAudioFile, inFrameOffset), nil
}

// ExtAudioFileSeek seeks to a specified frame in a file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileSeek(_:_:)
func ExtAudioFileSeek(inExtAudioFile ExtAudioFileRef, inFrameOffset int64) int32 {
	result, callErr := tryExtAudioFileSeek(inExtAudioFile, inFrameOffset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileSetProperty func(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _extAudioFileSetPropertyErr error

func tryExtAudioFileSetProperty(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _extAudioFileSetProperty == nil {
		return 0, symbolCallError("ExtAudioFileSetProperty", "10.4", _extAudioFileSetPropertyErr)
	}
	return _extAudioFileSetProperty(inExtAudioFile, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// ExtAudioFileSetProperty sets a property value for an extended audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileSetProperty(_:_:_:_:)
func ExtAudioFileSetProperty(inExtAudioFile ExtAudioFileRef, inPropertyID ExtAudioFilePropertyID, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryExtAudioFileSetProperty(inExtAudioFile, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileTell func(inExtAudioFile ExtAudioFileRef, outFrameOffset *int64) int32
var _extAudioFileTellErr error

func tryExtAudioFileTell(inExtAudioFile ExtAudioFileRef, outFrameOffset *int64) (int32, error) {
	if _extAudioFileTell == nil {
		return 0, symbolCallError("ExtAudioFileTell", "10.4", _extAudioFileTellErr)
	}
	return _extAudioFileTell(inExtAudioFile, outFrameOffset), nil
}

// ExtAudioFileTell gets an audio file’s read/write position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileTell(_:_:)
func ExtAudioFileTell(inExtAudioFile ExtAudioFileRef, outFrameOffset *int64) int32 {
	result, callErr := tryExtAudioFileTell(inExtAudioFile, outFrameOffset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileWrapAudioFileID func(inFileID AudioFileID, inForWriting bool, outExtAudioFile *ExtAudioFileRef) int32
var _extAudioFileWrapAudioFileIDErr error

func tryExtAudioFileWrapAudioFileID(inFileID AudioFileID, inForWriting bool, outExtAudioFile *ExtAudioFileRef) (int32, error) {
	if _extAudioFileWrapAudioFileID == nil {
		return 0, symbolCallError("ExtAudioFileWrapAudioFileID", "10.4", _extAudioFileWrapAudioFileIDErr)
	}
	return _extAudioFileWrapAudioFileID(inFileID, inForWriting, outExtAudioFile), nil
}

// ExtAudioFileWrapAudioFileID wraps an audio file object in an extended audio file object.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWrapAudioFileID(_:_:_:)
func ExtAudioFileWrapAudioFileID(inFileID AudioFileID, inForWriting bool, outExtAudioFile *ExtAudioFileRef) int32 {
	result, callErr := tryExtAudioFileWrapAudioFileID(inFileID, inForWriting, outExtAudioFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileWrite func(inExtAudioFile ExtAudioFileRef, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32
var _extAudioFileWriteErr error

func tryExtAudioFileWrite(inExtAudioFile ExtAudioFileRef, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) (int32, error) {
	if _extAudioFileWrite == nil {
		return 0, symbolCallError("ExtAudioFileWrite", "10.4", _extAudioFileWriteErr)
	}
	return _extAudioFileWrite(inExtAudioFile, inNumberFrames, ioData), nil
}

// ExtAudioFileWrite performs a synchronous, sequential write operation on an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWrite(_:_:_:)
func ExtAudioFileWrite(inExtAudioFile ExtAudioFileRef, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32 {
	result, callErr := tryExtAudioFileWrite(inExtAudioFile, inNumberFrames, ioData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _extAudioFileWriteAsync func(inExtAudioFile ExtAudioFileRef, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32
var _extAudioFileWriteAsyncErr error

func tryExtAudioFileWriteAsync(inExtAudioFile ExtAudioFileRef, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) (int32, error) {
	if _extAudioFileWriteAsync == nil {
		return 0, symbolCallError("ExtAudioFileWriteAsync", "10.4", _extAudioFileWriteAsyncErr)
	}
	return _extAudioFileWriteAsync(inExtAudioFile, inNumberFrames, ioData), nil
}

// ExtAudioFileWriteAsync perform an asynchronous, sequential write operation on an audio file.
//
// See: https://developer.apple.com/documentation/AudioToolbox/ExtAudioFileWriteAsync(_:_:_:)
func ExtAudioFileWriteAsync(inExtAudioFile ExtAudioFileRef, inNumberFrames uint32, ioData *coreaudiotypes.AudioBufferList) int32 {
	result, callErr := tryExtAudioFileWriteAsync(inExtAudioFile, inNumberFrames, ioData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getNameFromSoundBank func(inSoundBankRef *coreservices.FSRef, outName *corefoundation.CFStringRef) int32
var _getNameFromSoundBankErr error

func tryGetNameFromSoundBank(inSoundBankRef *coreservices.FSRef, outName *corefoundation.CFStringRef) (int32, error) {
	if _getNameFromSoundBank == nil {
		return 0, symbolCallError("GetNameFromSoundBank", "10.2", _getNameFromSoundBankErr)
	}
	return _getNameFromSoundBank(inSoundBankRef, outName), nil
}

// GetNameFromSoundBank gets the name of a sound bank from a sound bank file.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/GetNameFromSoundBank
func GetNameFromSoundBank(inSoundBankRef *coreservices.FSRef, outName *corefoundation.CFStringRef) int32 {
	result, callErr := tryGetNameFromSoundBank(inSoundBankRef, outName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicDeviceMIDIEvent func(inUnit MusicDeviceComponent, inStatus uint32, inData1 uint32, inData2 uint32, inOffsetSampleFrame uint32) int32
var _musicDeviceMIDIEventErr error

func tryMusicDeviceMIDIEvent(inUnit MusicDeviceComponent, inStatus uint32, inData1 uint32, inData2 uint32, inOffsetSampleFrame uint32) (int32, error) {
	if _musicDeviceMIDIEvent == nil {
		return 0, symbolCallError("MusicDeviceMIDIEvent", "10.0", _musicDeviceMIDIEventErr)
	}
	return _musicDeviceMIDIEvent(inUnit, inStatus, inData1, inData2, inOffsetSampleFrame), nil
}

// MusicDeviceMIDIEvent.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEvent(_:_:_:_:_:)
func MusicDeviceMIDIEvent(inUnit MusicDeviceComponent, inStatus uint32, inData1 uint32, inData2 uint32, inOffsetSampleFrame uint32) int32 {
	result, callErr := tryMusicDeviceMIDIEvent(inUnit, inStatus, inData1, inData2, inOffsetSampleFrame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicDeviceMIDIEventList func(inUnit MusicDeviceComponent, inOffsetSampleFrame uint32, evtList *coremidi.MIDIEventList) int32
var _musicDeviceMIDIEventListErr error

func tryMusicDeviceMIDIEventList(inUnit MusicDeviceComponent, inOffsetSampleFrame uint32, evtList *coremidi.MIDIEventList) (int32, error) {
	if _musicDeviceMIDIEventList == nil {
		return 0, symbolCallError("MusicDeviceMIDIEventList", "12.0", _musicDeviceMIDIEventListErr)
	}
	return _musicDeviceMIDIEventList(inUnit, inOffsetSampleFrame, evtList), nil
}

// MusicDeviceMIDIEventList.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceMIDIEventList(_:_:_:)
func MusicDeviceMIDIEventList(inUnit MusicDeviceComponent, inOffsetSampleFrame uint32, evtList *coremidi.MIDIEventList) int32 {
	result, callErr := tryMusicDeviceMIDIEventList(inUnit, inOffsetSampleFrame, evtList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicDevicePrepareInstrument func(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) int32
var _musicDevicePrepareInstrumentErr error

func tryMusicDevicePrepareInstrument(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) (int32, error) {
	if _musicDevicePrepareInstrument == nil {
		return 0, symbolCallError("MusicDevicePrepareInstrument", "10.0", _musicDevicePrepareInstrumentErr)
	}
	return _musicDevicePrepareInstrument(inUnit, inInstrument), nil
}

// MusicDevicePrepareInstrument.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicDevicePrepareInstrument
func MusicDevicePrepareInstrument(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) int32 {
	result, callErr := tryMusicDevicePrepareInstrument(inUnit, inInstrument)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicDeviceReleaseInstrument func(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) int32
var _musicDeviceReleaseInstrumentErr error

func tryMusicDeviceReleaseInstrument(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) (int32, error) {
	if _musicDeviceReleaseInstrument == nil {
		return 0, symbolCallError("MusicDeviceReleaseInstrument", "10.0", _musicDeviceReleaseInstrumentErr)
	}
	return _musicDeviceReleaseInstrument(inUnit, inInstrument), nil
}

// MusicDeviceReleaseInstrument.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceReleaseInstrument
func MusicDeviceReleaseInstrument(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID) int32 {
	result, callErr := tryMusicDeviceReleaseInstrument(inUnit, inInstrument)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicDeviceStartNote func(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID, inGroupID MusicDeviceGroupID, outNoteInstanceID *NoteInstanceID, inOffsetSampleFrame uint32, inParams *MusicDeviceNoteParams) int32
var _musicDeviceStartNoteErr error

func tryMusicDeviceStartNote(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID, inGroupID MusicDeviceGroupID, outNoteInstanceID *NoteInstanceID, inOffsetSampleFrame uint32, inParams *MusicDeviceNoteParams) (int32, error) {
	if _musicDeviceStartNote == nil {
		return 0, symbolCallError("MusicDeviceStartNote", "10.0", _musicDeviceStartNoteErr)
	}
	return _musicDeviceStartNote(inUnit, inInstrument, inGroupID, outNoteInstanceID, inOffsetSampleFrame, inParams), nil
}

// MusicDeviceStartNote.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStartNote(_:_:_:_:_:_:)
func MusicDeviceStartNote(inUnit MusicDeviceComponent, inInstrument MusicDeviceInstrumentID, inGroupID MusicDeviceGroupID, outNoteInstanceID *NoteInstanceID, inOffsetSampleFrame uint32, inParams *MusicDeviceNoteParams) int32 {
	result, callErr := tryMusicDeviceStartNote(inUnit, inInstrument, inGroupID, outNoteInstanceID, inOffsetSampleFrame, inParams)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicDeviceStopNote func(inUnit MusicDeviceComponent, inGroupID MusicDeviceGroupID, inNoteInstanceID NoteInstanceID, inOffsetSampleFrame uint32) int32
var _musicDeviceStopNoteErr error

func tryMusicDeviceStopNote(inUnit MusicDeviceComponent, inGroupID MusicDeviceGroupID, inNoteInstanceID NoteInstanceID, inOffsetSampleFrame uint32) (int32, error) {
	if _musicDeviceStopNote == nil {
		return 0, symbolCallError("MusicDeviceStopNote", "10.0", _musicDeviceStopNoteErr)
	}
	return _musicDeviceStopNote(inUnit, inGroupID, inNoteInstanceID, inOffsetSampleFrame), nil
}

// MusicDeviceStopNote.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceStopNote(_:_:_:_:)
func MusicDeviceStopNote(inUnit MusicDeviceComponent, inGroupID MusicDeviceGroupID, inNoteInstanceID NoteInstanceID, inOffsetSampleFrame uint32) int32 {
	result, callErr := tryMusicDeviceStopNote(inUnit, inGroupID, inNoteInstanceID, inOffsetSampleFrame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicDeviceSysEx func(inUnit MusicDeviceComponent, inData *byte, inLength uint32) int32
var _musicDeviceSysExErr error

func tryMusicDeviceSysEx(inUnit MusicDeviceComponent, inData []byte, inLength uint32) (int32, error) {
	if _musicDeviceSysEx == nil {
		return 0, symbolCallError("MusicDeviceSysEx", "10.0", _musicDeviceSysExErr)
	}
	return _musicDeviceSysEx(inUnit, unsafe.SliceData(inData), inLength), nil
}

// MusicDeviceSysEx.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicDeviceSysEx(_:_:_:)
func MusicDeviceSysEx(inUnit MusicDeviceComponent, inData []byte, inLength uint32) int32 {
	result, callErr := tryMusicDeviceSysEx(inUnit, inData, inLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorDeleteEvent func(inIterator MusicEventIterator) int32
var _musicEventIteratorDeleteEventErr error

func tryMusicEventIteratorDeleteEvent(inIterator MusicEventIterator) (int32, error) {
	if _musicEventIteratorDeleteEvent == nil {
		return 0, symbolCallError("MusicEventIteratorDeleteEvent", "10.0", _musicEventIteratorDeleteEventErr)
	}
	return _musicEventIteratorDeleteEvent(inIterator), nil
}

// MusicEventIteratorDeleteEvent deletes the event at a music event iterator’s current position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorDeleteEvent(_:)
func MusicEventIteratorDeleteEvent(inIterator MusicEventIterator) int32 {
	result, callErr := tryMusicEventIteratorDeleteEvent(inIterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorGetEventInfo func(inIterator MusicEventIterator, outTimeStamp *MusicTimeStamp, outEventType *MusicEventType, outEventData unsafe.Pointer, outEventDataSize *uint32) int32
var _musicEventIteratorGetEventInfoErr error

func tryMusicEventIteratorGetEventInfo(inIterator MusicEventIterator, outTimeStamp *MusicTimeStamp, outEventType *MusicEventType, outEventData unsafe.Pointer, outEventDataSize *uint32) (int32, error) {
	if _musicEventIteratorGetEventInfo == nil {
		return 0, symbolCallError("MusicEventIteratorGetEventInfo", "10.0", _musicEventIteratorGetEventInfoErr)
	}
	return _musicEventIteratorGetEventInfo(inIterator, outTimeStamp, outEventType, outEventData, outEventDataSize), nil
}

// MusicEventIteratorGetEventInfo gets information about the event at a music event iterator’s current position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorGetEventInfo(_:_:_:_:_:)
func MusicEventIteratorGetEventInfo(inIterator MusicEventIterator, outTimeStamp *MusicTimeStamp, outEventType *MusicEventType, outEventData unsafe.Pointer, outEventDataSize *uint32) int32 {
	result, callErr := tryMusicEventIteratorGetEventInfo(inIterator, outTimeStamp, outEventType, outEventData, outEventDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorHasCurrentEvent func(inIterator MusicEventIterator, outHasCurEvent *bool) int32
var _musicEventIteratorHasCurrentEventErr error

func tryMusicEventIteratorHasCurrentEvent(inIterator MusicEventIterator, outHasCurEvent *bool) (int32, error) {
	if _musicEventIteratorHasCurrentEvent == nil {
		return 0, symbolCallError("MusicEventIteratorHasCurrentEvent", "10.2", _musicEventIteratorHasCurrentEventErr)
	}
	return _musicEventIteratorHasCurrentEvent(inIterator, outHasCurEvent), nil
}

// MusicEventIteratorHasCurrentEvent indicates whether or not a music track contains an event at the music event iterator’s current position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasCurrentEvent(_:_:)
func MusicEventIteratorHasCurrentEvent(inIterator MusicEventIterator, outHasCurEvent *bool) int32 {
	result, callErr := tryMusicEventIteratorHasCurrentEvent(inIterator, outHasCurEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorHasNextEvent func(inIterator MusicEventIterator, outHasNextEvent *bool) int32
var _musicEventIteratorHasNextEventErr error

func tryMusicEventIteratorHasNextEvent(inIterator MusicEventIterator, outHasNextEvent *bool) (int32, error) {
	if _musicEventIteratorHasNextEvent == nil {
		return 0, symbolCallError("MusicEventIteratorHasNextEvent", "10.0", _musicEventIteratorHasNextEventErr)
	}
	return _musicEventIteratorHasNextEvent(inIterator, outHasNextEvent), nil
}

// MusicEventIteratorHasNextEvent indicates whether or not a music track contains an event beyond the music event iterator’s current position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasNextEvent(_:_:)
func MusicEventIteratorHasNextEvent(inIterator MusicEventIterator, outHasNextEvent *bool) int32 {
	result, callErr := tryMusicEventIteratorHasNextEvent(inIterator, outHasNextEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorHasPreviousEvent func(inIterator MusicEventIterator, outHasPrevEvent *bool) int32
var _musicEventIteratorHasPreviousEventErr error

func tryMusicEventIteratorHasPreviousEvent(inIterator MusicEventIterator, outHasPrevEvent *bool) (int32, error) {
	if _musicEventIteratorHasPreviousEvent == nil {
		return 0, symbolCallError("MusicEventIteratorHasPreviousEvent", "10.0", _musicEventIteratorHasPreviousEventErr)
	}
	return _musicEventIteratorHasPreviousEvent(inIterator, outHasPrevEvent), nil
}

// MusicEventIteratorHasPreviousEvent indicates whether or not a music track contains an event before the music event iterator’s current position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorHasPreviousEvent(_:_:)
func MusicEventIteratorHasPreviousEvent(inIterator MusicEventIterator, outHasPrevEvent *bool) int32 {
	result, callErr := tryMusicEventIteratorHasPreviousEvent(inIterator, outHasPrevEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorNextEvent func(inIterator MusicEventIterator) int32
var _musicEventIteratorNextEventErr error

func tryMusicEventIteratorNextEvent(inIterator MusicEventIterator) (int32, error) {
	if _musicEventIteratorNextEvent == nil {
		return 0, symbolCallError("MusicEventIteratorNextEvent", "10.0", _musicEventIteratorNextEventErr)
	}
	return _musicEventIteratorNextEvent(inIterator), nil
}

// MusicEventIteratorNextEvent positions a music event iterator at the next event on a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorNextEvent(_:)
func MusicEventIteratorNextEvent(inIterator MusicEventIterator) int32 {
	result, callErr := tryMusicEventIteratorNextEvent(inIterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorPreviousEvent func(inIterator MusicEventIterator) int32
var _musicEventIteratorPreviousEventErr error

func tryMusicEventIteratorPreviousEvent(inIterator MusicEventIterator) (int32, error) {
	if _musicEventIteratorPreviousEvent == nil {
		return 0, symbolCallError("MusicEventIteratorPreviousEvent", "10.0", _musicEventIteratorPreviousEventErr)
	}
	return _musicEventIteratorPreviousEvent(inIterator), nil
}

// MusicEventIteratorPreviousEvent positions a music event iterator at the previous event on a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorPreviousEvent(_:)
func MusicEventIteratorPreviousEvent(inIterator MusicEventIterator) int32 {
	result, callErr := tryMusicEventIteratorPreviousEvent(inIterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorSeek func(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) int32
var _musicEventIteratorSeekErr error

func tryMusicEventIteratorSeek(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) (int32, error) {
	if _musicEventIteratorSeek == nil {
		return 0, symbolCallError("MusicEventIteratorSeek", "10.0", _musicEventIteratorSeekErr)
	}
	return _musicEventIteratorSeek(inIterator, inTimeStamp), nil
}

// MusicEventIteratorSeek positions a music event iterator at a specified timestamp, in beats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSeek(_:_:)
func MusicEventIteratorSeek(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) int32 {
	result, callErr := tryMusicEventIteratorSeek(inIterator, inTimeStamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorSetEventInfo func(inIterator MusicEventIterator, inEventType MusicEventType, inEventData unsafe.Pointer) int32
var _musicEventIteratorSetEventInfoErr error

func tryMusicEventIteratorSetEventInfo(inIterator MusicEventIterator, inEventType MusicEventType, inEventData unsafe.Pointer) (int32, error) {
	if _musicEventIteratorSetEventInfo == nil {
		return 0, symbolCallError("MusicEventIteratorSetEventInfo", "10.2", _musicEventIteratorSetEventInfoErr)
	}
	return _musicEventIteratorSetEventInfo(inIterator, inEventType, inEventData), nil
}

// MusicEventIteratorSetEventInfo sets information for the event at a music event iterator’s current position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSetEventInfo(_:_:_:)
func MusicEventIteratorSetEventInfo(inIterator MusicEventIterator, inEventType MusicEventType, inEventData unsafe.Pointer) int32 {
	result, callErr := tryMusicEventIteratorSetEventInfo(inIterator, inEventType, inEventData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicEventIteratorSetEventTime func(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) int32
var _musicEventIteratorSetEventTimeErr error

func tryMusicEventIteratorSetEventTime(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) (int32, error) {
	if _musicEventIteratorSetEventTime == nil {
		return 0, symbolCallError("MusicEventIteratorSetEventTime", "10.0", _musicEventIteratorSetEventTimeErr)
	}
	return _musicEventIteratorSetEventTime(inIterator, inTimeStamp), nil
}

// MusicEventIteratorSetEventTime sets the timestamp for the event at a music event iterator’s current position.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicEventIteratorSetEventTime(_:_:)
func MusicEventIteratorSetEventTime(inIterator MusicEventIterator, inTimeStamp MusicTimeStamp) int32 {
	result, callErr := tryMusicEventIteratorSetEventTime(inIterator, inTimeStamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerGetBeatsForHostTime func(inPlayer MusicPlayer, inHostTime uint64, outBeats *MusicTimeStamp) int32
var _musicPlayerGetBeatsForHostTimeErr error

func tryMusicPlayerGetBeatsForHostTime(inPlayer MusicPlayer, inHostTime uint64, outBeats *MusicTimeStamp) (int32, error) {
	if _musicPlayerGetBeatsForHostTime == nil {
		return 0, symbolCallError("MusicPlayerGetBeatsForHostTime", "10.2", _musicPlayerGetBeatsForHostTimeErr)
	}
	return _musicPlayerGetBeatsForHostTime(inPlayer, inHostTime, outBeats), nil
}

// MusicPlayerGetBeatsForHostTime gets the beat number associated a specified host time.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetBeatsForHostTime(_:_:_:)
func MusicPlayerGetBeatsForHostTime(inPlayer MusicPlayer, inHostTime uint64, outBeats *MusicTimeStamp) int32 {
	result, callErr := tryMusicPlayerGetBeatsForHostTime(inPlayer, inHostTime, outBeats)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerGetHostTimeForBeats func(inPlayer MusicPlayer, inBeats MusicTimeStamp, outHostTime *uint64) int32
var _musicPlayerGetHostTimeForBeatsErr error

func tryMusicPlayerGetHostTimeForBeats(inPlayer MusicPlayer, inBeats MusicTimeStamp, outHostTime *uint64) (int32, error) {
	if _musicPlayerGetHostTimeForBeats == nil {
		return 0, symbolCallError("MusicPlayerGetHostTimeForBeats", "10.2", _musicPlayerGetHostTimeForBeatsErr)
	}
	return _musicPlayerGetHostTimeForBeats(inPlayer, inBeats, outHostTime), nil
}

// MusicPlayerGetHostTimeForBeats gets the host time associated with a specified beat.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetHostTimeForBeats(_:_:_:)
func MusicPlayerGetHostTimeForBeats(inPlayer MusicPlayer, inBeats MusicTimeStamp, outHostTime *uint64) int32 {
	result, callErr := tryMusicPlayerGetHostTimeForBeats(inPlayer, inBeats, outHostTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerGetPlayRateScalar func(inPlayer MusicPlayer, outScaleRate *float64) int32
var _musicPlayerGetPlayRateScalarErr error

func tryMusicPlayerGetPlayRateScalar(inPlayer MusicPlayer, outScaleRate *float64) (int32, error) {
	if _musicPlayerGetPlayRateScalar == nil {
		return 0, symbolCallError("MusicPlayerGetPlayRateScalar", "10.3", _musicPlayerGetPlayRateScalarErr)
	}
	return _musicPlayerGetPlayRateScalar(inPlayer, outScaleRate), nil
}

// MusicPlayerGetPlayRateScalar gets the playback rate multiplier for a music player.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetPlayRateScalar(_:_:)
func MusicPlayerGetPlayRateScalar(inPlayer MusicPlayer, outScaleRate *float64) int32 {
	result, callErr := tryMusicPlayerGetPlayRateScalar(inPlayer, outScaleRate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerGetSequence func(inPlayer MusicPlayer, outSequence *MusicSequence) int32
var _musicPlayerGetSequenceErr error

func tryMusicPlayerGetSequence(inPlayer MusicPlayer, outSequence *MusicSequence) (int32, error) {
	if _musicPlayerGetSequence == nil {
		return 0, symbolCallError("MusicPlayerGetSequence", "10.3", _musicPlayerGetSequenceErr)
	}
	return _musicPlayerGetSequence(inPlayer, outSequence), nil
}

// MusicPlayerGetSequence gets the music sequence associated with a music player.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetSequence(_:_:)
func MusicPlayerGetSequence(inPlayer MusicPlayer, outSequence *MusicSequence) int32 {
	result, callErr := tryMusicPlayerGetSequence(inPlayer, outSequence)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerGetTime func(inPlayer MusicPlayer, outTime *MusicTimeStamp) int32
var _musicPlayerGetTimeErr error

func tryMusicPlayerGetTime(inPlayer MusicPlayer, outTime *MusicTimeStamp) (int32, error) {
	if _musicPlayerGetTime == nil {
		return 0, symbolCallError("MusicPlayerGetTime", "10.0", _musicPlayerGetTimeErr)
	}
	return _musicPlayerGetTime(inPlayer, outTime), nil
}

// MusicPlayerGetTime gets the playback point for a music player, in beats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerGetTime(_:_:)
func MusicPlayerGetTime(inPlayer MusicPlayer, outTime *MusicTimeStamp) int32 {
	result, callErr := tryMusicPlayerGetTime(inPlayer, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerIsPlaying func(inPlayer MusicPlayer, outIsPlaying *bool) int32
var _musicPlayerIsPlayingErr error

func tryMusicPlayerIsPlaying(inPlayer MusicPlayer, outIsPlaying *bool) (int32, error) {
	if _musicPlayerIsPlaying == nil {
		return 0, symbolCallError("MusicPlayerIsPlaying", "10.2", _musicPlayerIsPlayingErr)
	}
	return _musicPlayerIsPlaying(inPlayer, outIsPlaying), nil
}

// MusicPlayerIsPlaying indicates whether or not a music player is playing.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerIsPlaying(_:_:)
func MusicPlayerIsPlaying(inPlayer MusicPlayer, outIsPlaying *bool) int32 {
	result, callErr := tryMusicPlayerIsPlaying(inPlayer, outIsPlaying)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerPreroll func(inPlayer MusicPlayer) int32
var _musicPlayerPrerollErr error

func tryMusicPlayerPreroll(inPlayer MusicPlayer) (int32, error) {
	if _musicPlayerPreroll == nil {
		return 0, symbolCallError("MusicPlayerPreroll", "10.0", _musicPlayerPrerollErr)
	}
	return _musicPlayerPreroll(inPlayer), nil
}

// MusicPlayerPreroll prepares a music player to play.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerPreroll(_:)
func MusicPlayerPreroll(inPlayer MusicPlayer) int32 {
	result, callErr := tryMusicPlayerPreroll(inPlayer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerSetPlayRateScalar func(inPlayer MusicPlayer, inScaleRate float64) int32
var _musicPlayerSetPlayRateScalarErr error

func tryMusicPlayerSetPlayRateScalar(inPlayer MusicPlayer, inScaleRate float64) (int32, error) {
	if _musicPlayerSetPlayRateScalar == nil {
		return 0, symbolCallError("MusicPlayerSetPlayRateScalar", "10.3", _musicPlayerSetPlayRateScalarErr)
	}
	return _musicPlayerSetPlayRateScalar(inPlayer, inScaleRate), nil
}

// MusicPlayerSetPlayRateScalar sets a playback rate multiplier for a music player.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetPlayRateScalar(_:_:)
func MusicPlayerSetPlayRateScalar(inPlayer MusicPlayer, inScaleRate float64) int32 {
	result, callErr := tryMusicPlayerSetPlayRateScalar(inPlayer, inScaleRate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerSetSequence func(inPlayer MusicPlayer, inSequence MusicSequence) int32
var _musicPlayerSetSequenceErr error

func tryMusicPlayerSetSequence(inPlayer MusicPlayer, inSequence MusicSequence) (int32, error) {
	if _musicPlayerSetSequence == nil {
		return 0, symbolCallError("MusicPlayerSetSequence", "10.0", _musicPlayerSetSequenceErr)
	}
	return _musicPlayerSetSequence(inPlayer, inSequence), nil
}

// MusicPlayerSetSequence sets the music sequence for the music player to play.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetSequence(_:_:)
func MusicPlayerSetSequence(inPlayer MusicPlayer, inSequence MusicSequence) int32 {
	result, callErr := tryMusicPlayerSetSequence(inPlayer, inSequence)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerSetTime func(inPlayer MusicPlayer, inTime MusicTimeStamp) int32
var _musicPlayerSetTimeErr error

func tryMusicPlayerSetTime(inPlayer MusicPlayer, inTime MusicTimeStamp) (int32, error) {
	if _musicPlayerSetTime == nil {
		return 0, symbolCallError("MusicPlayerSetTime", "10.0", _musicPlayerSetTimeErr)
	}
	return _musicPlayerSetTime(inPlayer, inTime), nil
}

// MusicPlayerSetTime sets the playback point for a music player, in beats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerSetTime(_:_:)
func MusicPlayerSetTime(inPlayer MusicPlayer, inTime MusicTimeStamp) int32 {
	result, callErr := tryMusicPlayerSetTime(inPlayer, inTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerStart func(inPlayer MusicPlayer) int32
var _musicPlayerStartErr error

func tryMusicPlayerStart(inPlayer MusicPlayer) (int32, error) {
	if _musicPlayerStart == nil {
		return 0, symbolCallError("MusicPlayerStart", "10.0", _musicPlayerStartErr)
	}
	return _musicPlayerStart(inPlayer), nil
}

// MusicPlayerStart starts playback of a music player.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerStart(_:)
func MusicPlayerStart(inPlayer MusicPlayer) int32 {
	result, callErr := tryMusicPlayerStart(inPlayer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicPlayerStop func(inPlayer MusicPlayer) int32
var _musicPlayerStopErr error

func tryMusicPlayerStop(inPlayer MusicPlayer) (int32, error) {
	if _musicPlayerStop == nil {
		return 0, symbolCallError("MusicPlayerStop", "10.0", _musicPlayerStopErr)
	}
	return _musicPlayerStop(inPlayer), nil
}

// MusicPlayerStop stops playback of a music player.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicPlayerStop(_:)
func MusicPlayerStop(inPlayer MusicPlayer) int32 {
	result, callErr := tryMusicPlayerStop(inPlayer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceBarBeatTimeToBeats func(inSequence MusicSequence, inBarBeatTime *CABarBeatTime, outBeats *MusicTimeStamp) int32
var _musicSequenceBarBeatTimeToBeatsErr error

func tryMusicSequenceBarBeatTimeToBeats(inSequence MusicSequence, inBarBeatTime *CABarBeatTime, outBeats *MusicTimeStamp) (int32, error) {
	if _musicSequenceBarBeatTimeToBeats == nil {
		return 0, symbolCallError("MusicSequenceBarBeatTimeToBeats", "10.5", _musicSequenceBarBeatTimeToBeatsErr)
	}
	return _musicSequenceBarBeatTimeToBeats(inSequence, inBarBeatTime, outBeats), nil
}

// MusicSequenceBarBeatTimeToBeats formats a music sequence’s bar-beat time to its beat time.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceBarBeatTimeToBeats(_:_:_:)
func MusicSequenceBarBeatTimeToBeats(inSequence MusicSequence, inBarBeatTime *CABarBeatTime, outBeats *MusicTimeStamp) int32 {
	result, callErr := tryMusicSequenceBarBeatTimeToBeats(inSequence, inBarBeatTime, outBeats)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceBeatsToBarBeatTime func(inSequence MusicSequence, inBeats MusicTimeStamp, inSubbeatDivisor uint32, outBarBeatTime *CABarBeatTime) int32
var _musicSequenceBeatsToBarBeatTimeErr error

func tryMusicSequenceBeatsToBarBeatTime(inSequence MusicSequence, inBeats MusicTimeStamp, inSubbeatDivisor uint32, outBarBeatTime *CABarBeatTime) (int32, error) {
	if _musicSequenceBeatsToBarBeatTime == nil {
		return 0, symbolCallError("MusicSequenceBeatsToBarBeatTime", "10.5", _musicSequenceBeatsToBarBeatTimeErr)
	}
	return _musicSequenceBeatsToBarBeatTime(inSequence, inBeats, inSubbeatDivisor, outBarBeatTime), nil
}

// MusicSequenceBeatsToBarBeatTime formats a music sequence’s beat time to its bar-beat time.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceBeatsToBarBeatTime(_:_:_:_:)
func MusicSequenceBeatsToBarBeatTime(inSequence MusicSequence, inBeats MusicTimeStamp, inSubbeatDivisor uint32, outBarBeatTime *CABarBeatTime) int32 {
	result, callErr := tryMusicSequenceBeatsToBarBeatTime(inSequence, inBeats, inSubbeatDivisor, outBarBeatTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceDisposeTrack func(inSequence MusicSequence, inTrack MusicTrack) int32
var _musicSequenceDisposeTrackErr error

func tryMusicSequenceDisposeTrack(inSequence MusicSequence, inTrack MusicTrack) (int32, error) {
	if _musicSequenceDisposeTrack == nil {
		return 0, symbolCallError("MusicSequenceDisposeTrack", "10.0", _musicSequenceDisposeTrackErr)
	}
	return _musicSequenceDisposeTrack(inSequence, inTrack), nil
}

// MusicSequenceDisposeTrack removes a music track from a music sequence, and disposes of the track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceDisposeTrack(_:_:)
func MusicSequenceDisposeTrack(inSequence MusicSequence, inTrack MusicTrack) int32 {
	result, callErr := tryMusicSequenceDisposeTrack(inSequence, inTrack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceFileCreate func(inSequence MusicSequence, inFileRef corefoundation.CFURLRef, inFileType MusicSequenceFileTypeID, inFlags MusicSequenceFileFlags, inResolution int16) int32
var _musicSequenceFileCreateErr error

func tryMusicSequenceFileCreate(inSequence MusicSequence, inFileRef corefoundation.CFURLRef, inFileType MusicSequenceFileTypeID, inFlags MusicSequenceFileFlags, inResolution int16) (int32, error) {
	if _musicSequenceFileCreate == nil {
		return 0, symbolCallError("MusicSequenceFileCreate", "10.5", _musicSequenceFileCreateErr)
	}
	return _musicSequenceFileCreate(inSequence, inFileRef, inFileType, inFlags, inResolution), nil
}

// MusicSequenceFileCreate creates a MIDI file from the events in a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileCreate(_:_:_:_:_:)
func MusicSequenceFileCreate(inSequence MusicSequence, inFileRef corefoundation.CFURLRef, inFileType MusicSequenceFileTypeID, inFlags MusicSequenceFileFlags, inResolution int16) int32 {
	result, callErr := tryMusicSequenceFileCreate(inSequence, inFileRef, inFileType, inFlags, inResolution)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceFileCreateData func(inSequence MusicSequence, inFileType MusicSequenceFileTypeID, inFlags MusicSequenceFileFlags, inResolution int16, outData *corefoundation.CFDataRef) int32
var _musicSequenceFileCreateDataErr error

func tryMusicSequenceFileCreateData(inSequence MusicSequence, inFileType MusicSequenceFileTypeID, inFlags MusicSequenceFileFlags, inResolution int16, outData *corefoundation.CFDataRef) (int32, error) {
	if _musicSequenceFileCreateData == nil {
		return 0, symbolCallError("MusicSequenceFileCreateData", "10.5", _musicSequenceFileCreateDataErr)
	}
	return _musicSequenceFileCreateData(inSequence, inFileType, inFlags, inResolution, outData), nil
}

// MusicSequenceFileCreateData creates a data object containing the events from a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileCreateData(_:_:_:_:_:)
func MusicSequenceFileCreateData(inSequence MusicSequence, inFileType MusicSequenceFileTypeID, inFlags MusicSequenceFileFlags, inResolution int16, outData *corefoundation.CFDataRef) int32 {
	result, callErr := tryMusicSequenceFileCreateData(inSequence, inFileType, inFlags, inResolution, outData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceFileLoad func(inSequence MusicSequence, inFileRef corefoundation.CFURLRef, inFileTypeHint MusicSequenceFileTypeID, inFlags MusicSequenceLoadFlags) int32
var _musicSequenceFileLoadErr error

func tryMusicSequenceFileLoad(inSequence MusicSequence, inFileRef corefoundation.CFURLRef, inFileTypeHint MusicSequenceFileTypeID, inFlags MusicSequenceLoadFlags) (int32, error) {
	if _musicSequenceFileLoad == nil {
		return 0, symbolCallError("MusicSequenceFileLoad", "10.5", _musicSequenceFileLoadErr)
	}
	return _musicSequenceFileLoad(inSequence, inFileRef, inFileTypeHint, inFlags), nil
}

// MusicSequenceFileLoad loads data into a music sequence from a URL reference.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileLoad(_:_:_:_:)
func MusicSequenceFileLoad(inSequence MusicSequence, inFileRef corefoundation.CFURLRef, inFileTypeHint MusicSequenceFileTypeID, inFlags MusicSequenceLoadFlags) int32 {
	result, callErr := tryMusicSequenceFileLoad(inSequence, inFileRef, inFileTypeHint, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceFileLoadData func(inSequence MusicSequence, inData corefoundation.CFDataRef, inFileTypeHint MusicSequenceFileTypeID, inFlags MusicSequenceLoadFlags) int32
var _musicSequenceFileLoadDataErr error

func tryMusicSequenceFileLoadData(inSequence MusicSequence, inData corefoundation.CFDataRef, inFileTypeHint MusicSequenceFileTypeID, inFlags MusicSequenceLoadFlags) (int32, error) {
	if _musicSequenceFileLoadData == nil {
		return 0, symbolCallError("MusicSequenceFileLoadData", "10.5", _musicSequenceFileLoadDataErr)
	}
	return _musicSequenceFileLoadData(inSequence, inData, inFileTypeHint, inFlags), nil
}

// MusicSequenceFileLoadData load data into a music sequence from a data reference.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceFileLoadData(_:_:_:_:)
func MusicSequenceFileLoadData(inSequence MusicSequence, inData corefoundation.CFDataRef, inFileTypeHint MusicSequenceFileTypeID, inFlags MusicSequenceLoadFlags) int32 {
	result, callErr := tryMusicSequenceFileLoadData(inSequence, inData, inFileTypeHint, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetAUGraph func(inSequence MusicSequence, outGraph *AUGraph) int32
var _musicSequenceGetAUGraphErr error

func tryMusicSequenceGetAUGraph(inSequence MusicSequence, outGraph *AUGraph) (int32, error) {
	if _musicSequenceGetAUGraph == nil {
		return 0, symbolCallError("MusicSequenceGetAUGraph", "10.0", _musicSequenceGetAUGraphErr)
	}
	return _musicSequenceGetAUGraph(inSequence, outGraph), nil
}

// MusicSequenceGetAUGraph gets the audio processing graph associated with a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetAUGraph(_:_:)
func MusicSequenceGetAUGraph(inSequence MusicSequence, outGraph *AUGraph) int32 {
	result, callErr := tryMusicSequenceGetAUGraph(inSequence, outGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetBeatsForSeconds func(inSequence MusicSequence, inSeconds float64, outBeats *MusicTimeStamp) int32
var _musicSequenceGetBeatsForSecondsErr error

func tryMusicSequenceGetBeatsForSeconds(inSequence MusicSequence, inSeconds float64, outBeats *MusicTimeStamp) (int32, error) {
	if _musicSequenceGetBeatsForSeconds == nil {
		return 0, symbolCallError("MusicSequenceGetBeatsForSeconds", "10.2", _musicSequenceGetBeatsForSecondsErr)
	}
	return _musicSequenceGetBeatsForSeconds(inSequence, inSeconds, outBeats), nil
}

// MusicSequenceGetBeatsForSeconds calculates the number of beats that correspond to a number of seconds.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetBeatsForSeconds(_:_:_:)
func MusicSequenceGetBeatsForSeconds(inSequence MusicSequence, inSeconds float64, outBeats *MusicTimeStamp) int32 {
	result, callErr := tryMusicSequenceGetBeatsForSeconds(inSequence, inSeconds, outBeats)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetIndTrack func(inSequence MusicSequence, inTrackIndex uint32, outTrack *MusicTrack) int32
var _musicSequenceGetIndTrackErr error

func tryMusicSequenceGetIndTrack(inSequence MusicSequence, inTrackIndex uint32, outTrack *MusicTrack) (int32, error) {
	if _musicSequenceGetIndTrack == nil {
		return 0, symbolCallError("MusicSequenceGetIndTrack", "10.0", _musicSequenceGetIndTrackErr)
	}
	return _musicSequenceGetIndTrack(inSequence, inTrackIndex, outTrack), nil
}

// MusicSequenceGetIndTrack gets the music track at the specified track index.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetIndTrack(_:_:_:)
func MusicSequenceGetIndTrack(inSequence MusicSequence, inTrackIndex uint32, outTrack *MusicTrack) int32 {
	result, callErr := tryMusicSequenceGetIndTrack(inSequence, inTrackIndex, outTrack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetInfoDictionary func(inSequence MusicSequence) corefoundation.CFDictionaryRef
var _musicSequenceGetInfoDictionaryErr error

func tryMusicSequenceGetInfoDictionary(inSequence MusicSequence) (corefoundation.CFDictionaryRef, error) {
	if _musicSequenceGetInfoDictionary == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("MusicSequenceGetInfoDictionary", "10.5", _musicSequenceGetInfoDictionaryErr)
	}
	return _musicSequenceGetInfoDictionary(inSequence), nil
}

// MusicSequenceGetInfoDictionary returns a dictionary containing music sequence information.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetInfoDictionary(_:)
func MusicSequenceGetInfoDictionary(inSequence MusicSequence) corefoundation.CFDictionaryRef {
	result, callErr := tryMusicSequenceGetInfoDictionary(inSequence)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetSecondsForBeats func(inSequence MusicSequence, inBeats MusicTimeStamp, outSeconds *float64) int32
var _musicSequenceGetSecondsForBeatsErr error

func tryMusicSequenceGetSecondsForBeats(inSequence MusicSequence, inBeats MusicTimeStamp, outSeconds *float64) (int32, error) {
	if _musicSequenceGetSecondsForBeats == nil {
		return 0, symbolCallError("MusicSequenceGetSecondsForBeats", "10.2", _musicSequenceGetSecondsForBeatsErr)
	}
	return _musicSequenceGetSecondsForBeats(inSequence, inBeats, outSeconds), nil
}

// MusicSequenceGetSecondsForBeats calculates the number of seconds that correspond to a number of beats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetSecondsForBeats(_:_:_:)
func MusicSequenceGetSecondsForBeats(inSequence MusicSequence, inBeats MusicTimeStamp, outSeconds *float64) int32 {
	result, callErr := tryMusicSequenceGetSecondsForBeats(inSequence, inBeats, outSeconds)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetSequenceType func(inSequence MusicSequence, outType *MusicSequenceType) int32
var _musicSequenceGetSequenceTypeErr error

func tryMusicSequenceGetSequenceType(inSequence MusicSequence, outType *MusicSequenceType) (int32, error) {
	if _musicSequenceGetSequenceType == nil {
		return 0, symbolCallError("MusicSequenceGetSequenceType", "10.5", _musicSequenceGetSequenceTypeErr)
	}
	return _musicSequenceGetSequenceType(inSequence, outType), nil
}

// MusicSequenceGetSequenceType gets the sequence type for a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetSequenceType(_:_:)
func MusicSequenceGetSequenceType(inSequence MusicSequence, outType *MusicSequenceType) int32 {
	result, callErr := tryMusicSequenceGetSequenceType(inSequence, outType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetTempoTrack func(inSequence MusicSequence, outTrack *MusicTrack) int32
var _musicSequenceGetTempoTrackErr error

func tryMusicSequenceGetTempoTrack(inSequence MusicSequence, outTrack *MusicTrack) (int32, error) {
	if _musicSequenceGetTempoTrack == nil {
		return 0, symbolCallError("MusicSequenceGetTempoTrack", "10.1", _musicSequenceGetTempoTrackErr)
	}
	return _musicSequenceGetTempoTrack(inSequence, outTrack), nil
}

// MusicSequenceGetTempoTrack gets the tempo track for a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTempoTrack(_:_:)
func MusicSequenceGetTempoTrack(inSequence MusicSequence, outTrack *MusicTrack) int32 {
	result, callErr := tryMusicSequenceGetTempoTrack(inSequence, outTrack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetTrackCount func(inSequence MusicSequence, outNumberOfTracks *uint32) int32
var _musicSequenceGetTrackCountErr error

func tryMusicSequenceGetTrackCount(inSequence MusicSequence, outNumberOfTracks *uint32) (int32, error) {
	if _musicSequenceGetTrackCount == nil {
		return 0, symbolCallError("MusicSequenceGetTrackCount", "10.0", _musicSequenceGetTrackCountErr)
	}
	return _musicSequenceGetTrackCount(inSequence, outNumberOfTracks), nil
}

// MusicSequenceGetTrackCount gets the number of music tracks owned by a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTrackCount(_:_:)
func MusicSequenceGetTrackCount(inSequence MusicSequence, outNumberOfTracks *uint32) int32 {
	result, callErr := tryMusicSequenceGetTrackCount(inSequence, outNumberOfTracks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceGetTrackIndex func(inSequence MusicSequence, inTrack MusicTrack, outTrackIndex *uint32) int32
var _musicSequenceGetTrackIndexErr error

func tryMusicSequenceGetTrackIndex(inSequence MusicSequence, inTrack MusicTrack, outTrackIndex *uint32) (int32, error) {
	if _musicSequenceGetTrackIndex == nil {
		return 0, symbolCallError("MusicSequenceGetTrackIndex", "10.0", _musicSequenceGetTrackIndexErr)
	}
	return _musicSequenceGetTrackIndex(inSequence, inTrack, outTrackIndex), nil
}

// MusicSequenceGetTrackIndex gets the index number for a specified music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceGetTrackIndex(_:_:_:)
func MusicSequenceGetTrackIndex(inSequence MusicSequence, inTrack MusicTrack, outTrackIndex *uint32) int32 {
	result, callErr := tryMusicSequenceGetTrackIndex(inSequence, inTrack, outTrackIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceLoadSMFDataWithFlags func(inSequence MusicSequence, inData corefoundation.CFDataRef, inFlags MusicSequenceLoadFlags) int32
var _musicSequenceLoadSMFDataWithFlagsErr error

func tryMusicSequenceLoadSMFDataWithFlags(inSequence MusicSequence, inData corefoundation.CFDataRef, inFlags MusicSequenceLoadFlags) (int32, error) {
	if _musicSequenceLoadSMFDataWithFlags == nil {
		return 0, symbolCallError("MusicSequenceLoadSMFDataWithFlags", "10.3", _musicSequenceLoadSMFDataWithFlagsErr)
	}
	return _musicSequenceLoadSMFDataWithFlags(inSequence, inData, inFlags), nil
}

// MusicSequenceLoadSMFDataWithFlags.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadSMFDataWithFlags
func MusicSequenceLoadSMFDataWithFlags(inSequence MusicSequence, inData corefoundation.CFDataRef, inFlags MusicSequenceLoadFlags) int32 {
	result, callErr := tryMusicSequenceLoadSMFDataWithFlags(inSequence, inData, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceLoadSMFWithFlags func(inSequence MusicSequence, inFileRef *coreservices.FSRef, inFlags MusicSequenceLoadFlags) int32
var _musicSequenceLoadSMFWithFlagsErr error

func tryMusicSequenceLoadSMFWithFlags(inSequence MusicSequence, inFileRef *coreservices.FSRef, inFlags MusicSequenceLoadFlags) (int32, error) {
	if _musicSequenceLoadSMFWithFlags == nil {
		return 0, symbolCallError("MusicSequenceLoadSMFWithFlags", "10.3", _musicSequenceLoadSMFWithFlagsErr)
	}
	return _musicSequenceLoadSMFWithFlags(inSequence, inFileRef, inFlags), nil
}

// MusicSequenceLoadSMFWithFlags.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceLoadSMFWithFlags
func MusicSequenceLoadSMFWithFlags(inSequence MusicSequence, inFileRef *coreservices.FSRef, inFlags MusicSequenceLoadFlags) int32 {
	result, callErr := tryMusicSequenceLoadSMFWithFlags(inSequence, inFileRef, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceNewTrack func(inSequence MusicSequence, outTrack *MusicTrack) int32
var _musicSequenceNewTrackErr error

func tryMusicSequenceNewTrack(inSequence MusicSequence, outTrack *MusicTrack) (int32, error) {
	if _musicSequenceNewTrack == nil {
		return 0, symbolCallError("MusicSequenceNewTrack", "10.0", _musicSequenceNewTrackErr)
	}
	return _musicSequenceNewTrack(inSequence, outTrack), nil
}

// MusicSequenceNewTrack add a new, empty music track to a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceNewTrack(_:_:)
func MusicSequenceNewTrack(inSequence MusicSequence, outTrack *MusicTrack) int32 {
	result, callErr := tryMusicSequenceNewTrack(inSequence, outTrack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceReverse func(inSequence MusicSequence) int32
var _musicSequenceReverseErr error

func tryMusicSequenceReverse(inSequence MusicSequence) (int32, error) {
	if _musicSequenceReverse == nil {
		return 0, symbolCallError("MusicSequenceReverse", "10.0", _musicSequenceReverseErr)
	}
	return _musicSequenceReverse(inSequence), nil
}

// MusicSequenceReverse reverses the MIDI and tempo events in a music sequence, so the start becomes the end.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceReverse(_:)
func MusicSequenceReverse(inSequence MusicSequence) int32 {
	result, callErr := tryMusicSequenceReverse(inSequence)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceSaveMIDIFile func(inSequence MusicSequence, inParentDirectory *coreservices.FSRef, inFileName corefoundation.CFStringRef, inResolution uint16, inFlags uint32) int32
var _musicSequenceSaveMIDIFileErr error

func tryMusicSequenceSaveMIDIFile(inSequence MusicSequence, inParentDirectory *coreservices.FSRef, inFileName corefoundation.CFStringRef, inResolution uint16, inFlags uint32) (int32, error) {
	if _musicSequenceSaveMIDIFile == nil {
		return 0, symbolCallError("MusicSequenceSaveMIDIFile", "10.4", _musicSequenceSaveMIDIFileErr)
	}
	return _musicSequenceSaveMIDIFile(inSequence, inParentDirectory, inFileName, inResolution, inFlags), nil
}

// MusicSequenceSaveMIDIFile.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSaveMIDIFile
func MusicSequenceSaveMIDIFile(inSequence MusicSequence, inParentDirectory *coreservices.FSRef, inFileName corefoundation.CFStringRef, inResolution uint16, inFlags uint32) int32 {
	result, callErr := tryMusicSequenceSaveMIDIFile(inSequence, inParentDirectory, inFileName, inResolution, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceSaveSMFData func(inSequence MusicSequence, outData *corefoundation.CFDataRef, inResolution uint16) int32
var _musicSequenceSaveSMFDataErr error

func tryMusicSequenceSaveSMFData(inSequence MusicSequence, outData *corefoundation.CFDataRef, inResolution uint16) (int32, error) {
	if _musicSequenceSaveSMFData == nil {
		return 0, symbolCallError("MusicSequenceSaveSMFData", "10.2", _musicSequenceSaveSMFDataErr)
	}
	return _musicSequenceSaveSMFData(inSequence, outData, inResolution), nil
}

// MusicSequenceSaveSMFData.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSaveSMFData
func MusicSequenceSaveSMFData(inSequence MusicSequence, outData *corefoundation.CFDataRef, inResolution uint16) int32 {
	result, callErr := tryMusicSequenceSaveSMFData(inSequence, outData, inResolution)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceSetAUGraph func(inSequence MusicSequence, inGraph AUGraph) int32
var _musicSequenceSetAUGraphErr error

func tryMusicSequenceSetAUGraph(inSequence MusicSequence, inGraph AUGraph) (int32, error) {
	if _musicSequenceSetAUGraph == nil {
		return 0, symbolCallError("MusicSequenceSetAUGraph", "10.0", _musicSequenceSetAUGraphErr)
	}
	return _musicSequenceSetAUGraph(inSequence, inGraph), nil
}

// MusicSequenceSetAUGraph associates an audio processing graph with a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetAUGraph(_:_:)
func MusicSequenceSetAUGraph(inSequence MusicSequence, inGraph AUGraph) int32 {
	result, callErr := tryMusicSequenceSetAUGraph(inSequence, inGraph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceSetMIDIEndpoint func(inSequence MusicSequence, inEndpoint uint32) int32
var _musicSequenceSetMIDIEndpointErr error

func tryMusicSequenceSetMIDIEndpoint(inSequence MusicSequence, inEndpoint uint32) (int32, error) {
	if _musicSequenceSetMIDIEndpoint == nil {
		return 0, symbolCallError("MusicSequenceSetMIDIEndpoint", "10.1", _musicSequenceSetMIDIEndpointErr)
	}
	return _musicSequenceSetMIDIEndpoint(inSequence, inEndpoint), nil
}

// MusicSequenceSetMIDIEndpoint associates a specified MIDI endpoint with all music tracks in a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetMIDIEndpoint(_:_:)
func MusicSequenceSetMIDIEndpoint(inSequence MusicSequence, inEndpoint uint32) int32 {
	result, callErr := tryMusicSequenceSetMIDIEndpoint(inSequence, inEndpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceSetSequenceType func(inSequence MusicSequence, inType MusicSequenceType) int32
var _musicSequenceSetSequenceTypeErr error

func tryMusicSequenceSetSequenceType(inSequence MusicSequence, inType MusicSequenceType) (int32, error) {
	if _musicSequenceSetSequenceType == nil {
		return 0, symbolCallError("MusicSequenceSetSequenceType", "10.5", _musicSequenceSetSequenceTypeErr)
	}
	return _musicSequenceSetSequenceType(inSequence, inType), nil
}

// MusicSequenceSetSequenceType sets the sequence type for a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetSequenceType(_:_:)
func MusicSequenceSetSequenceType(inSequence MusicSequence, inType MusicSequenceType) int32 {
	result, callErr := tryMusicSequenceSetSequenceType(inSequence, inType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicSequenceSetUserCallback func(inSequence MusicSequence, inCallback MusicSequenceUserCallback, inClientData unsafe.Pointer) int32
var _musicSequenceSetUserCallbackErr error

func tryMusicSequenceSetUserCallback(inSequence MusicSequence, inCallback MusicSequenceUserCallback, inClientData unsafe.Pointer) (int32, error) {
	if _musicSequenceSetUserCallback == nil {
		return 0, symbolCallError("MusicSequenceSetUserCallback", "10.3", _musicSequenceSetUserCallbackErr)
	}
	return _musicSequenceSetUserCallback(inSequence, inCallback, inClientData), nil
}

// MusicSequenceSetUserCallback registers a user callback function with a music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicSequenceSetUserCallback(_:_:_:)
func MusicSequenceSetUserCallback(inSequence MusicSequence, inCallback MusicSequenceUserCallback, inClientData unsafe.Pointer) int32 {
	result, callErr := tryMusicSequenceSetUserCallback(inSequence, inCallback, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackClear func(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) int32
var _musicTrackClearErr error

func tryMusicTrackClear(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) (int32, error) {
	if _musicTrackClear == nil {
		return 0, symbolCallError("MusicTrackClear", "10.0", _musicTrackClearErr)
	}
	return _musicTrackClear(inTrack, inStartTime, inEndTime), nil
}

// MusicTrackClear removes a specified range of music track events.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackClear(_:_:_:)
func MusicTrackClear(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) int32 {
	result, callErr := tryMusicTrackClear(inTrack, inStartTime, inEndTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackCopyInsert func(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) int32
var _musicTrackCopyInsertErr error

func tryMusicTrackCopyInsert(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) (int32, error) {
	if _musicTrackCopyInsert == nil {
		return 0, symbolCallError("MusicTrackCopyInsert", "10.0", _musicTrackCopyInsertErr)
	}
	return _musicTrackCopyInsert(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime), nil
}

// MusicTrackCopyInsert copies a range of events from one music track and inserts them into another music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackCopyInsert(_:_:_:_:_:)
func MusicTrackCopyInsert(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) int32 {
	result, callErr := tryMusicTrackCopyInsert(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackCut func(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) int32
var _musicTrackCutErr error

func tryMusicTrackCut(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) (int32, error) {
	if _musicTrackCut == nil {
		return 0, symbolCallError("MusicTrackCut", "10.0", _musicTrackCutErr)
	}
	return _musicTrackCut(inTrack, inStartTime, inEndTime), nil
}

// MusicTrackCut removes a specified range of music track events, and shifts later events toward the start of the track to fill in the gap.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackCut(_:_:_:)
func MusicTrackCut(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp) int32 {
	result, callErr := tryMusicTrackCut(inTrack, inStartTime, inEndTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackGetDestMIDIEndpoint func(inTrack MusicTrack, outEndpoint *coremidi.MIDIEndpointRef) int32
var _musicTrackGetDestMIDIEndpointErr error

func tryMusicTrackGetDestMIDIEndpoint(inTrack MusicTrack, outEndpoint *coremidi.MIDIEndpointRef) (int32, error) {
	if _musicTrackGetDestMIDIEndpoint == nil {
		return 0, symbolCallError("MusicTrackGetDestMIDIEndpoint", "10.1", _musicTrackGetDestMIDIEndpointErr)
	}
	return _musicTrackGetDestMIDIEndpoint(inTrack, outEndpoint), nil
}

// MusicTrackGetDestMIDIEndpoint gets the MIDI endpoint that is the event target for a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetDestMIDIEndpoint(_:_:)
func MusicTrackGetDestMIDIEndpoint(inTrack MusicTrack, outEndpoint *coremidi.MIDIEndpointRef) int32 {
	result, callErr := tryMusicTrackGetDestMIDIEndpoint(inTrack, outEndpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackGetDestNode func(inTrack MusicTrack, outNode *AUNode) int32
var _musicTrackGetDestNodeErr error

func tryMusicTrackGetDestNode(inTrack MusicTrack, outNode *AUNode) (int32, error) {
	if _musicTrackGetDestNode == nil {
		return 0, symbolCallError("MusicTrackGetDestNode", "10.1", _musicTrackGetDestNodeErr)
	}
	return _musicTrackGetDestNode(inTrack, outNode), nil
}

// MusicTrackGetDestNode gets the audio unit node that is the event target for a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetDestNode(_:_:)
func MusicTrackGetDestNode(inTrack MusicTrack, outNode *AUNode) int32 {
	result, callErr := tryMusicTrackGetDestNode(inTrack, outNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackGetProperty func(inTrack MusicTrack, inPropertyID uint32, outData unsafe.Pointer, ioLength *uint32) int32
var _musicTrackGetPropertyErr error

func tryMusicTrackGetProperty(inTrack MusicTrack, inPropertyID uint32, outData unsafe.Pointer, ioLength *uint32) (int32, error) {
	if _musicTrackGetProperty == nil {
		return 0, symbolCallError("MusicTrackGetProperty", "10.0", _musicTrackGetPropertyErr)
	}
	return _musicTrackGetProperty(inTrack, inPropertyID, outData, ioLength), nil
}

// MusicTrackGetProperty gets a music track property value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetProperty(_:_:_:_:)
func MusicTrackGetProperty(inTrack MusicTrack, inPropertyID uint32, outData unsafe.Pointer, ioLength *uint32) int32 {
	result, callErr := tryMusicTrackGetProperty(inTrack, inPropertyID, outData, ioLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackGetSequence func(inTrack MusicTrack, outSequence *MusicSequence) int32
var _musicTrackGetSequenceErr error

func tryMusicTrackGetSequence(inTrack MusicTrack, outSequence *MusicSequence) (int32, error) {
	if _musicTrackGetSequence == nil {
		return 0, symbolCallError("MusicTrackGetSequence", "10.0", _musicTrackGetSequenceErr)
	}
	return _musicTrackGetSequence(inTrack, outSequence), nil
}

// MusicTrackGetSequence gets the music sequence that the music track is a member of.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackGetSequence(_:_:)
func MusicTrackGetSequence(inTrack MusicTrack, outSequence *MusicSequence) int32 {
	result, callErr := tryMusicTrackGetSequence(inTrack, outSequence)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackMerge func(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) int32
var _musicTrackMergeErr error

func tryMusicTrackMerge(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) (int32, error) {
	if _musicTrackMerge == nil {
		return 0, symbolCallError("MusicTrackMerge", "10.0", _musicTrackMergeErr)
	}
	return _musicTrackMerge(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime), nil
}

// MusicTrackMerge copies a range of events from one music track and merges them into another music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackMerge(_:_:_:_:_:)
func MusicTrackMerge(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, inDestTrack MusicTrack, inDestInsertTime MusicTimeStamp) int32 {
	result, callErr := tryMusicTrackMerge(inSourceTrack, inSourceStartTime, inSourceEndTime, inDestTrack, inDestInsertTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackMoveEvents func(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp, inMoveTime MusicTimeStamp) int32
var _musicTrackMoveEventsErr error

func tryMusicTrackMoveEvents(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp, inMoveTime MusicTimeStamp) (int32, error) {
	if _musicTrackMoveEvents == nil {
		return 0, symbolCallError("MusicTrackMoveEvents", "10.0", _musicTrackMoveEventsErr)
	}
	return _musicTrackMoveEvents(inTrack, inStartTime, inEndTime, inMoveTime), nil
}

// MusicTrackMoveEvents shifts music track events forward or backward in time, in terms of beats.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackMoveEvents(_:_:_:_:)
func MusicTrackMoveEvents(inTrack MusicTrack, inStartTime MusicTimeStamp, inEndTime MusicTimeStamp, inMoveTime MusicTimeStamp) int32 {
	result, callErr := tryMusicTrackMoveEvents(inTrack, inStartTime, inEndTime, inMoveTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewAUPresetEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inPresetEvent *AUPresetEvent) int32
var _musicTrackNewAUPresetEventErr error

func tryMusicTrackNewAUPresetEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inPresetEvent *AUPresetEvent) (int32, error) {
	if _musicTrackNewAUPresetEvent == nil {
		return 0, symbolCallError("MusicTrackNewAUPresetEvent", "10.3", _musicTrackNewAUPresetEventErr)
	}
	return _musicTrackNewAUPresetEvent(inTrack, inTimeStamp, inPresetEvent), nil
}

// MusicTrackNewAUPresetEvent adds an event of type [AUPresetEvent] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewAUPresetEvent(_:_:_:)
func MusicTrackNewAUPresetEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inPresetEvent *AUPresetEvent) int32 {
	result, callErr := tryMusicTrackNewAUPresetEvent(inTrack, inTimeStamp, inPresetEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewExtendedControlEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ExtendedControlEvent) int32
var _musicTrackNewExtendedControlEventErr error

func tryMusicTrackNewExtendedControlEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ExtendedControlEvent) (int32, error) {
	if _musicTrackNewExtendedControlEvent == nil {
		return 0, symbolCallError("MusicTrackNewExtendedControlEvent", "10.0", _musicTrackNewExtendedControlEventErr)
	}
	return _musicTrackNewExtendedControlEvent(inTrack, inTimeStamp, inInfo), nil
}

// MusicTrackNewExtendedControlEvent.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedControlEvent
func MusicTrackNewExtendedControlEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ExtendedControlEvent) int32 {
	result, callErr := tryMusicTrackNewExtendedControlEvent(inTrack, inTimeStamp, inInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewExtendedNoteEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ExtendedNoteOnEvent) int32
var _musicTrackNewExtendedNoteEventErr error

func tryMusicTrackNewExtendedNoteEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ExtendedNoteOnEvent) (int32, error) {
	if _musicTrackNewExtendedNoteEvent == nil {
		return 0, symbolCallError("MusicTrackNewExtendedNoteEvent", "10.0", _musicTrackNewExtendedNoteEventErr)
	}
	return _musicTrackNewExtendedNoteEvent(inTrack, inTimeStamp, inInfo), nil
}

// MusicTrackNewExtendedNoteEvent adds an event of type [ExtendedNoteOnEvent] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedNoteEvent(_:_:_:)
func MusicTrackNewExtendedNoteEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ExtendedNoteOnEvent) int32 {
	result, callErr := tryMusicTrackNewExtendedNoteEvent(inTrack, inTimeStamp, inInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewExtendedTempoEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inBPM float64) int32
var _musicTrackNewExtendedTempoEventErr error

func tryMusicTrackNewExtendedTempoEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inBPM float64) (int32, error) {
	if _musicTrackNewExtendedTempoEvent == nil {
		return 0, symbolCallError("MusicTrackNewExtendedTempoEvent", "10.0", _musicTrackNewExtendedTempoEventErr)
	}
	return _musicTrackNewExtendedTempoEvent(inTrack, inTimeStamp, inBPM), nil
}

// MusicTrackNewExtendedTempoEvent adds a tempo to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewExtendedTempoEvent(_:_:_:)
func MusicTrackNewExtendedTempoEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inBPM float64) int32 {
	result, callErr := tryMusicTrackNewExtendedTempoEvent(inTrack, inTimeStamp, inBPM)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewMIDIChannelEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage *MIDIChannelMessage) int32
var _musicTrackNewMIDIChannelEventErr error

func tryMusicTrackNewMIDIChannelEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage *MIDIChannelMessage) (int32, error) {
	if _musicTrackNewMIDIChannelEvent == nil {
		return 0, symbolCallError("MusicTrackNewMIDIChannelEvent", "10.0", _musicTrackNewMIDIChannelEventErr)
	}
	return _musicTrackNewMIDIChannelEvent(inTrack, inTimeStamp, inMessage), nil
}

// MusicTrackNewMIDIChannelEvent adds an event of type [MIDIChannelMessage] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDIChannelEvent(_:_:_:)
func MusicTrackNewMIDIChannelEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage *MIDIChannelMessage) int32 {
	result, callErr := tryMusicTrackNewMIDIChannelEvent(inTrack, inTimeStamp, inMessage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewMIDINoteEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage *MIDINoteMessage) int32
var _musicTrackNewMIDINoteEventErr error

func tryMusicTrackNewMIDINoteEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage *MIDINoteMessage) (int32, error) {
	if _musicTrackNewMIDINoteEvent == nil {
		return 0, symbolCallError("MusicTrackNewMIDINoteEvent", "10.0", _musicTrackNewMIDINoteEventErr)
	}
	return _musicTrackNewMIDINoteEvent(inTrack, inTimeStamp, inMessage), nil
}

// MusicTrackNewMIDINoteEvent adds an event of type [MIDINoteMessage] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDINoteEvent(_:_:_:)
func MusicTrackNewMIDINoteEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMessage *MIDINoteMessage) int32 {
	result, callErr := tryMusicTrackNewMIDINoteEvent(inTrack, inTimeStamp, inMessage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewMIDIRawDataEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inRawData *MIDIRawData) int32
var _musicTrackNewMIDIRawDataEventErr error

func tryMusicTrackNewMIDIRawDataEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inRawData *MIDIRawData) (int32, error) {
	if _musicTrackNewMIDIRawDataEvent == nil {
		return 0, symbolCallError("MusicTrackNewMIDIRawDataEvent", "10.0", _musicTrackNewMIDIRawDataEventErr)
	}
	return _musicTrackNewMIDIRawDataEvent(inTrack, inTimeStamp, inRawData), nil
}

// MusicTrackNewMIDIRawDataEvent adds an event of type [MIDIRawData] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMIDIRawDataEvent(_:_:_:)
func MusicTrackNewMIDIRawDataEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inRawData *MIDIRawData) int32 {
	result, callErr := tryMusicTrackNewMIDIRawDataEvent(inTrack, inTimeStamp, inRawData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewMetaEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMetaEvent *MIDIMetaEvent) int32
var _musicTrackNewMetaEventErr error

func tryMusicTrackNewMetaEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMetaEvent *MIDIMetaEvent) (int32, error) {
	if _musicTrackNewMetaEvent == nil {
		return 0, symbolCallError("MusicTrackNewMetaEvent", "10.0", _musicTrackNewMetaEventErr)
	}
	return _musicTrackNewMetaEvent(inTrack, inTimeStamp, inMetaEvent), nil
}

// MusicTrackNewMetaEvent adds an event of type [MIDIMetaEvent] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewMetaEvent(_:_:_:)
func MusicTrackNewMetaEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inMetaEvent *MIDIMetaEvent) int32 {
	result, callErr := tryMusicTrackNewMetaEvent(inTrack, inTimeStamp, inMetaEvent)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewParameterEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ParameterEvent) int32
var _musicTrackNewParameterEventErr error

func tryMusicTrackNewParameterEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ParameterEvent) (int32, error) {
	if _musicTrackNewParameterEvent == nil {
		return 0, symbolCallError("MusicTrackNewParameterEvent", "10.2", _musicTrackNewParameterEventErr)
	}
	return _musicTrackNewParameterEvent(inTrack, inTimeStamp, inInfo), nil
}

// MusicTrackNewParameterEvent adds an event of type [ParameterEvent] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewParameterEvent(_:_:_:)
func MusicTrackNewParameterEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inInfo *ParameterEvent) int32 {
	result, callErr := tryMusicTrackNewParameterEvent(inTrack, inTimeStamp, inInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackNewUserEvent func(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inUserData *MusicEventUserData) int32
var _musicTrackNewUserEventErr error

func tryMusicTrackNewUserEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inUserData *MusicEventUserData) (int32, error) {
	if _musicTrackNewUserEvent == nil {
		return 0, symbolCallError("MusicTrackNewUserEvent", "10.0", _musicTrackNewUserEventErr)
	}
	return _musicTrackNewUserEvent(inTrack, inTimeStamp, inUserData), nil
}

// MusicTrackNewUserEvent adds an event of type [MusicEventUserData] to a music track.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackNewUserEvent(_:_:_:)
func MusicTrackNewUserEvent(inTrack MusicTrack, inTimeStamp MusicTimeStamp, inUserData *MusicEventUserData) int32 {
	result, callErr := tryMusicTrackNewUserEvent(inTrack, inTimeStamp, inUserData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackSetDestMIDIEndpoint func(inTrack MusicTrack, inEndpoint uint32) int32
var _musicTrackSetDestMIDIEndpointErr error

func tryMusicTrackSetDestMIDIEndpoint(inTrack MusicTrack, inEndpoint uint32) (int32, error) {
	if _musicTrackSetDestMIDIEndpoint == nil {
		return 0, symbolCallError("MusicTrackSetDestMIDIEndpoint", "10.1", _musicTrackSetDestMIDIEndpointErr)
	}
	return _musicTrackSetDestMIDIEndpoint(inTrack, inEndpoint), nil
}

// MusicTrackSetDestMIDIEndpoint sets the music track’s event target to a MIDI endpoint.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetDestMIDIEndpoint(_:_:)
func MusicTrackSetDestMIDIEndpoint(inTrack MusicTrack, inEndpoint uint32) int32 {
	result, callErr := tryMusicTrackSetDestMIDIEndpoint(inTrack, inEndpoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackSetDestNode func(inTrack MusicTrack, inNode AUNode) int32
var _musicTrackSetDestNodeErr error

func tryMusicTrackSetDestNode(inTrack MusicTrack, inNode AUNode) (int32, error) {
	if _musicTrackSetDestNode == nil {
		return 0, symbolCallError("MusicTrackSetDestNode", "10.0", _musicTrackSetDestNodeErr)
	}
	return _musicTrackSetDestNode(inTrack, inNode), nil
}

// MusicTrackSetDestNode sets the music track’s event target to an audio unit node.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetDestNode(_:_:)
func MusicTrackSetDestNode(inTrack MusicTrack, inNode AUNode) int32 {
	result, callErr := tryMusicTrackSetDestNode(inTrack, inNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _musicTrackSetProperty func(inTrack MusicTrack, inPropertyID uint32, inData unsafe.Pointer, inLength uint32) int32
var _musicTrackSetPropertyErr error

func tryMusicTrackSetProperty(inTrack MusicTrack, inPropertyID uint32, inData unsafe.Pointer, inLength uint32) (int32, error) {
	if _musicTrackSetProperty == nil {
		return 0, symbolCallError("MusicTrackSetProperty", "10.0", _musicTrackSetPropertyErr)
	}
	return _musicTrackSetProperty(inTrack, inPropertyID, inData, inLength), nil
}

// MusicTrackSetProperty sets a music track property value.
//
// See: https://developer.apple.com/documentation/AudioToolbox/MusicTrackSetProperty(_:_:_:_:)
func MusicTrackSetProperty(inTrack MusicTrack, inPropertyID uint32, inData unsafe.Pointer, inLength uint32) int32 {
	result, callErr := tryMusicTrackSetProperty(inTrack, inPropertyID, inData, inLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newMusicEventIterator func(inTrack MusicTrack, outIterator *MusicEventIterator) int32
var _newMusicEventIteratorErr error

func tryNewMusicEventIterator(inTrack MusicTrack, outIterator *MusicEventIterator) (int32, error) {
	if _newMusicEventIterator == nil {
		return 0, symbolCallError("NewMusicEventIterator", "10.0", _newMusicEventIteratorErr)
	}
	return _newMusicEventIterator(inTrack, outIterator), nil
}

// NewMusicEventIterator creates a new music event iterator.
//
// See: https://developer.apple.com/documentation/AudioToolbox/NewMusicEventIterator(_:_:)
func NewMusicEventIterator(inTrack MusicTrack, outIterator *MusicEventIterator) int32 {
	result, callErr := tryNewMusicEventIterator(inTrack, outIterator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newMusicPlayer func(outPlayer *MusicPlayer) int32
var _newMusicPlayerErr error

func tryNewMusicPlayer(outPlayer *MusicPlayer) (int32, error) {
	if _newMusicPlayer == nil {
		return 0, symbolCallError("NewMusicPlayer", "10.0", _newMusicPlayerErr)
	}
	return _newMusicPlayer(outPlayer), nil
}

// NewMusicPlayer creates a new music player.
//
// See: https://developer.apple.com/documentation/AudioToolbox/NewMusicPlayer(_:)
func NewMusicPlayer(outPlayer *MusicPlayer) int32 {
	result, callErr := tryNewMusicPlayer(outPlayer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newMusicSequence func(outSequence *MusicSequence) int32
var _newMusicSequenceErr error

func tryNewMusicSequence(outSequence *MusicSequence) (int32, error) {
	if _newMusicSequence == nil {
		return 0, symbolCallError("NewMusicSequence", "10.0", _newMusicSequenceErr)
	}
	return _newMusicSequence(outSequence), nil
}

// NewMusicSequence creates a new empty music sequence.
//
// See: https://developer.apple.com/documentation/AudioToolbox/NewMusicSequence(_:)
func NewMusicSequence(outSequence *MusicSequence) int32 {
	result, callErr := tryNewMusicSequence(outSequence)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newMusicTrackFrom func(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, outNewTrack *MusicTrack) int32
var _newMusicTrackFromErr error

func tryNewMusicTrackFrom(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, outNewTrack *MusicTrack) (int32, error) {
	if _newMusicTrackFrom == nil {
		return 0, symbolCallError("NewMusicTrackFrom", "10.0", _newMusicTrackFromErr)
	}
	return _newMusicTrackFrom(inSourceTrack, inSourceStartTime, inSourceEndTime, outNewTrack), nil
}

// NewMusicTrackFrom.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AudioToolbox/NewMusicTrackFrom
func NewMusicTrackFrom(inSourceTrack MusicTrack, inSourceStartTime MusicTimeStamp, inSourceEndTime MusicTimeStamp, outNewTrack *MusicTrack) int32 {
	result, callErr := tryNewMusicTrackFrom(inSourceTrack, inSourceStartTime, inSourceEndTime, outNewTrack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_aUEventListenerAddEventType, &_aUEventListenerAddEventTypeErr, frameworkHandle, "AUEventListenerAddEventType", "10.3")
	registerFunc(&_aUEventListenerCreate, &_aUEventListenerCreateErr, frameworkHandle, "AUEventListenerCreate", "10.3")
	registerFunc(&_aUEventListenerCreateWithDispatchQueue, &_aUEventListenerCreateWithDispatchQueueErr, frameworkHandle, "AUEventListenerCreateWithDispatchQueue", "10.6")
	registerFunc(&_aUEventListenerNotify, &_aUEventListenerNotifyErr, frameworkHandle, "AUEventListenerNotify", "10.3")
	registerFunc(&_aUEventListenerRemoveEventType, &_aUEventListenerRemoveEventTypeErr, frameworkHandle, "AUEventListenerRemoveEventType", "10.3")
	registerFunc(&_aUGraphAddNode, &_aUGraphAddNodeErr, frameworkHandle, "AUGraphAddNode", "10.5")
	registerFunc(&_aUGraphAddRenderNotify, &_aUGraphAddRenderNotifyErr, frameworkHandle, "AUGraphAddRenderNotify", "10.2")
	registerFunc(&_aUGraphClearConnections, &_aUGraphClearConnectionsErr, frameworkHandle, "AUGraphClearConnections", "10.0")
	registerFunc(&_aUGraphClose, &_aUGraphCloseErr, frameworkHandle, "AUGraphClose", "10.0")
	registerFunc(&_aUGraphConnectNodeInput, &_aUGraphConnectNodeInputErr, frameworkHandle, "AUGraphConnectNodeInput", "10.0")
	registerFunc(&_aUGraphCountNodeConnections, &_aUGraphCountNodeConnectionsErr, frameworkHandle, "AUGraphCountNodeConnections", "10.3")
	registerFunc(&_aUGraphCountNodeInteractions, &_aUGraphCountNodeInteractionsErr, frameworkHandle, "AUGraphCountNodeInteractions", "10.5")
	registerFunc(&_aUGraphDisconnectNodeInput, &_aUGraphDisconnectNodeInputErr, frameworkHandle, "AUGraphDisconnectNodeInput", "10.0")
	registerFunc(&_aUGraphGetCPULoad, &_aUGraphGetCPULoadErr, frameworkHandle, "AUGraphGetCPULoad", "10.1")
	registerFunc(&_aUGraphGetConnectionInfo, &_aUGraphGetConnectionInfoErr, frameworkHandle, "AUGraphGetConnectionInfo", "10.1")
	registerFunc(&_aUGraphGetIndNode, &_aUGraphGetIndNodeErr, frameworkHandle, "AUGraphGetIndNode", "10.0")
	registerFunc(&_aUGraphGetInteractionInfo, &_aUGraphGetInteractionInfoErr, frameworkHandle, "AUGraphGetInteractionInfo", "10.5")
	registerFunc(&_aUGraphGetMaxCPULoad, &_aUGraphGetMaxCPULoadErr, frameworkHandle, "AUGraphGetMaxCPULoad", "10.3")
	registerFunc(&_aUGraphGetNodeConnections, &_aUGraphGetNodeConnectionsErr, frameworkHandle, "AUGraphGetNodeConnections", "10.3")
	registerFunc(&_aUGraphGetNodeCount, &_aUGraphGetNodeCountErr, frameworkHandle, "AUGraphGetNodeCount", "10.0")
	registerFunc(&_aUGraphGetNodeInfo, &_aUGraphGetNodeInfoErr, frameworkHandle, "AUGraphGetNodeInfo", "10.0")
	registerFunc(&_aUGraphGetNodeInfoSubGraph, &_aUGraphGetNodeInfoSubGraphErr, frameworkHandle, "AUGraphGetNodeInfoSubGraph", "10.2")
	registerFunc(&_aUGraphGetNodeInteractions, &_aUGraphGetNodeInteractionsErr, frameworkHandle, "AUGraphGetNodeInteractions", "10.5")
	registerFunc(&_aUGraphGetNumberOfConnections, &_aUGraphGetNumberOfConnectionsErr, frameworkHandle, "AUGraphGetNumberOfConnections", "10.1")
	registerFunc(&_aUGraphGetNumberOfInteractions, &_aUGraphGetNumberOfInteractionsErr, frameworkHandle, "AUGraphGetNumberOfInteractions", "10.5")
	registerFunc(&_aUGraphInitialize, &_aUGraphInitializeErr, frameworkHandle, "AUGraphInitialize", "10.0")
	registerFunc(&_aUGraphIsInitialized, &_aUGraphIsInitializedErr, frameworkHandle, "AUGraphIsInitialized", "10.0")
	registerFunc(&_aUGraphIsNodeSubGraph, &_aUGraphIsNodeSubGraphErr, frameworkHandle, "AUGraphIsNodeSubGraph", "10.2")
	registerFunc(&_aUGraphIsOpen, &_aUGraphIsOpenErr, frameworkHandle, "AUGraphIsOpen", "10.0")
	registerFunc(&_aUGraphIsRunning, &_aUGraphIsRunningErr, frameworkHandle, "AUGraphIsRunning", "10.0")
	registerFunc(&_aUGraphNewNode, &_aUGraphNewNodeErr, frameworkHandle, "AUGraphNewNode", "10.0")
	registerFunc(&_aUGraphNewNodeSubGraph, &_aUGraphNewNodeSubGraphErr, frameworkHandle, "AUGraphNewNodeSubGraph", "10.2")
	registerFunc(&_aUGraphNodeInfo, &_aUGraphNodeInfoErr, frameworkHandle, "AUGraphNodeInfo", "10.5")
	registerFunc(&_aUGraphOpen, &_aUGraphOpenErr, frameworkHandle, "AUGraphOpen", "10.0")
	registerFunc(&_aUGraphRemoveNode, &_aUGraphRemoveNodeErr, frameworkHandle, "AUGraphRemoveNode", "10.0")
	registerFunc(&_aUGraphRemoveRenderNotify, &_aUGraphRemoveRenderNotifyErr, frameworkHandle, "AUGraphRemoveRenderNotify", "10.2")
	registerFunc(&_aUGraphSetNodeInputCallback, &_aUGraphSetNodeInputCallbackErr, frameworkHandle, "AUGraphSetNodeInputCallback", "10.5")
	registerFunc(&_aUGraphStart, &_aUGraphStartErr, frameworkHandle, "AUGraphStart", "10.0")
	registerFunc(&_aUGraphStop, &_aUGraphStopErr, frameworkHandle, "AUGraphStop", "10.0")
	registerFunc(&_aUGraphUninitialize, &_aUGraphUninitializeErr, frameworkHandle, "AUGraphUninitialize", "10.0")
	registerFunc(&_aUGraphUpdate, &_aUGraphUpdateErr, frameworkHandle, "AUGraphUpdate", "10.0")
	registerFunc(&_aUListenerAddParameter, &_aUListenerAddParameterErr, frameworkHandle, "AUListenerAddParameter", "10.2")
	registerFunc(&_aUListenerCreate, &_aUListenerCreateErr, frameworkHandle, "AUListenerCreate", "10.2")
	registerFunc(&_aUListenerCreateWithDispatchQueue, &_aUListenerCreateWithDispatchQueueErr, frameworkHandle, "AUListenerCreateWithDispatchQueue", "10.6")
	registerFunc(&_aUListenerDispose, &_aUListenerDisposeErr, frameworkHandle, "AUListenerDispose", "10.2")
	registerFunc(&_aUListenerRemoveParameter, &_aUListenerRemoveParameterErr, frameworkHandle, "AUListenerRemoveParameter", "10.2")
	registerFunc(&_aUParameterFormatValue, &_aUParameterFormatValueErr, frameworkHandle, "AUParameterFormatValue", "10.2")
	registerFunc(&_aUParameterListenerNotify, &_aUParameterListenerNotifyErr, frameworkHandle, "AUParameterListenerNotify", "10.2")
	registerFunc(&_aUParameterSet, &_aUParameterSetErr, frameworkHandle, "AUParameterSet", "10.2")
	registerFunc(&_aUParameterValueFromLinear, &_aUParameterValueFromLinearErr, frameworkHandle, "AUParameterValueFromLinear", "10.2")
	registerFunc(&_aUParameterValueToLinear, &_aUParameterValueToLinearErr, frameworkHandle, "AUParameterValueToLinear", "10.2")
	registerFunc(&_audioCodecAppendInputBufferList, &_audioCodecAppendInputBufferListErr, frameworkHandle, "AudioCodecAppendInputBufferList", "10.7")
	registerFunc(&_audioCodecAppendInputData, &_audioCodecAppendInputDataErr, frameworkHandle, "AudioCodecAppendInputData", "10.2")
	registerFunc(&_audioCodecGetProperty, &_audioCodecGetPropertyErr, frameworkHandle, "AudioCodecGetProperty", "10.2")
	registerFunc(&_audioCodecGetPropertyInfo, &_audioCodecGetPropertyInfoErr, frameworkHandle, "AudioCodecGetPropertyInfo", "10.2")
	registerFunc(&_audioCodecInitialize, &_audioCodecInitializeErr, frameworkHandle, "AudioCodecInitialize", "10.2")
	registerFunc(&_audioCodecProduceOutputBufferList, &_audioCodecProduceOutputBufferListErr, frameworkHandle, "AudioCodecProduceOutputBufferList", "10.7")
	registerFunc(&_audioCodecProduceOutputPackets, &_audioCodecProduceOutputPacketsErr, frameworkHandle, "AudioCodecProduceOutputPackets", "10.2")
	registerFunc(&_audioCodecReset, &_audioCodecResetErr, frameworkHandle, "AudioCodecReset", "10.2")
	registerFunc(&_audioCodecSetProperty, &_audioCodecSetPropertyErr, frameworkHandle, "AudioCodecSetProperty", "10.2")
	registerFunc(&_audioCodecUninitialize, &_audioCodecUninitializeErr, frameworkHandle, "AudioCodecUninitialize", "10.2")
	registerFunc(&_audioComponentCopyConfigurationInfo, &_audioComponentCopyConfigurationInfoErr, frameworkHandle, "AudioComponentCopyConfigurationInfo", "10.7")
	registerFunc(&_audioComponentCopyIcon, &_audioComponentCopyIconErr, frameworkHandle, "AudioComponentCopyIcon", "11.0")
	registerFunc(&_audioComponentCopyName, &_audioComponentCopyNameErr, frameworkHandle, "AudioComponentCopyName", "10.6")
	registerFunc(&_audioComponentCount, &_audioComponentCountErr, frameworkHandle, "AudioComponentCount", "10.6")
	registerFunc(&_audioComponentFindNext, &_audioComponentFindNextErr, frameworkHandle, "AudioComponentFindNext", "10.6")
	registerFunc(&_audioComponentGetDescription, &_audioComponentGetDescriptionErr, frameworkHandle, "AudioComponentGetDescription", "10.6")
	registerFunc(&_audioComponentGetIcon, &_audioComponentGetIconErr, frameworkHandle, "AudioComponentGetIcon", "10.11")
	registerFunc(&_audioComponentGetVersion, &_audioComponentGetVersionErr, frameworkHandle, "AudioComponentGetVersion", "10.6")
	registerFunc(&_audioComponentInstanceCanDo, &_audioComponentInstanceCanDoErr, frameworkHandle, "AudioComponentInstanceCanDo", "10.6")
	registerFunc(&_audioComponentInstanceDispose, &_audioComponentInstanceDisposeErr, frameworkHandle, "AudioComponentInstanceDispose", "10.6")
	registerFunc(&_audioComponentInstanceGetComponent, &_audioComponentInstanceGetComponentErr, frameworkHandle, "AudioComponentInstanceGetComponent", "10.6")
	registerFunc(&_audioComponentInstanceNew, &_audioComponentInstanceNewErr, frameworkHandle, "AudioComponentInstanceNew", "10.6")
	registerFunc(&_audioComponentInstantiate, &_audioComponentInstantiateErr, frameworkHandle, "AudioComponentInstantiate", "10.11")
	registerFunc(&_audioComponentRegister, &_audioComponentRegisterErr, frameworkHandle, "AudioComponentRegister", "10.7")
	registerFunc(&_audioComponentValidate, &_audioComponentValidateErr, frameworkHandle, "AudioComponentValidate", "10.7")
	registerFunc(&_audioComponentValidateWithResults, &_audioComponentValidateWithResultsErr, frameworkHandle, "AudioComponentValidateWithResults", "13.0")
	registerFunc(&_audioConverterConvertBuffer, &_audioConverterConvertBufferErr, frameworkHandle, "AudioConverterConvertBuffer", "10.1")
	registerFunc(&_audioConverterConvertComplexBuffer, &_audioConverterConvertComplexBufferErr, frameworkHandle, "AudioConverterConvertComplexBuffer", "10.7")
	registerFunc(&_audioConverterDispose, &_audioConverterDisposeErr, frameworkHandle, "AudioConverterDispose", "10.1")
	registerFunc(&_audioConverterFillComplexBuffer, &_audioConverterFillComplexBufferErr, frameworkHandle, "AudioConverterFillComplexBuffer", "10.2")
	registerFunc(&_audioConverterFillComplexBufferRealtimeSafe, &_audioConverterFillComplexBufferRealtimeSafeErr, frameworkHandle, "AudioConverterFillComplexBufferRealtimeSafe", "26.0")
	registerFunc(&_audioConverterFillComplexBufferWithPacketDependencies, &_audioConverterFillComplexBufferWithPacketDependenciesErr, frameworkHandle, "AudioConverterFillComplexBufferWithPacketDependencies", "26.0")
	registerFunc(&_audioConverterGetProperty, &_audioConverterGetPropertyErr, frameworkHandle, "AudioConverterGetProperty", "10.1")
	registerFunc(&_audioConverterGetPropertyInfo, &_audioConverterGetPropertyInfoErr, frameworkHandle, "AudioConverterGetPropertyInfo", "10.1")
	registerFunc(&_audioConverterNew, &_audioConverterNewErr, frameworkHandle, "AudioConverterNew", "10.1")
	registerFunc(&_audioConverterNewSpecific, &_audioConverterNewSpecificErr, frameworkHandle, "AudioConverterNewSpecific", "10.4")
	registerFunc(&_audioConverterNewWithOptions, &_audioConverterNewWithOptionsErr, frameworkHandle, "AudioConverterNewWithOptions", "15.0")
	registerFunc(&_audioConverterPrepare, &_audioConverterPrepareErr, frameworkHandle, "AudioConverterPrepare", "15.0")
	registerFunc(&_audioConverterReset, &_audioConverterResetErr, frameworkHandle, "AudioConverterReset", "10.1")
	registerFunc(&_audioConverterSetProperty, &_audioConverterSetPropertyErr, frameworkHandle, "AudioConverterSetProperty", "10.1")
	registerFunc(&_audioFileClose, &_audioFileCloseErr, frameworkHandle, "AudioFileClose", "10.2")
	registerFunc(&_audioFileComponentCloseFile, &_audioFileComponentCloseFileErr, frameworkHandle, "AudioFileComponentCloseFile", "10.4")
	registerFunc(&_audioFileComponentCountUserData, &_audioFileComponentCountUserDataErr, frameworkHandle, "AudioFileComponentCountUserData", "10.4")
	registerFunc(&_audioFileComponentCreate, &_audioFileComponentCreateErr, frameworkHandle, "AudioFileComponentCreate", "10.4")
	registerFunc(&_audioFileComponentCreateURL, &_audioFileComponentCreateURLErr, frameworkHandle, "AudioFileComponentCreateURL", "10.5")
	registerFunc(&_audioFileComponentDataIsThisFormat, &_audioFileComponentDataIsThisFormatErr, frameworkHandle, "AudioFileComponentDataIsThisFormat", "10.4")
	registerFunc(&_audioFileComponentExtensionIsThisFormat, &_audioFileComponentExtensionIsThisFormatErr, frameworkHandle, "AudioFileComponentExtensionIsThisFormat", "10.4")
	registerFunc(&_audioFileComponentFileDataIsThisFormat, &_audioFileComponentFileDataIsThisFormatErr, frameworkHandle, "AudioFileComponentFileDataIsThisFormat", "10.4")
	registerFunc(&_audioFileComponentFileIsThisFormat, &_audioFileComponentFileIsThisFormatErr, frameworkHandle, "AudioFileComponentFileIsThisFormat", "10.4")
	registerFunc(&_audioFileComponentGetGlobalInfo, &_audioFileComponentGetGlobalInfoErr, frameworkHandle, "AudioFileComponentGetGlobalInfo", "10.4")
	registerFunc(&_audioFileComponentGetGlobalInfoSize, &_audioFileComponentGetGlobalInfoSizeErr, frameworkHandle, "AudioFileComponentGetGlobalInfoSize", "10.4")
	registerFunc(&_audioFileComponentGetProperty, &_audioFileComponentGetPropertyErr, frameworkHandle, "AudioFileComponentGetProperty", "10.4")
	registerFunc(&_audioFileComponentGetPropertyInfo, &_audioFileComponentGetPropertyInfoErr, frameworkHandle, "AudioFileComponentGetPropertyInfo", "10.4")
	registerFunc(&_audioFileComponentGetUserData, &_audioFileComponentGetUserDataErr, frameworkHandle, "AudioFileComponentGetUserData", "10.4")
	registerFunc(&_audioFileComponentGetUserDataAtOffset, &_audioFileComponentGetUserDataAtOffsetErr, frameworkHandle, "AudioFileComponentGetUserDataAtOffset", "14.0")
	registerFunc(&_audioFileComponentGetUserDataSize, &_audioFileComponentGetUserDataSizeErr, frameworkHandle, "AudioFileComponentGetUserDataSize", "10.4")
	registerFunc(&_audioFileComponentGetUserDataSize64, &_audioFileComponentGetUserDataSize64Err, frameworkHandle, "AudioFileComponentGetUserDataSize64", "14.0")
	registerFunc(&_audioFileComponentInitialize, &_audioFileComponentInitializeErr, frameworkHandle, "AudioFileComponentInitialize", "10.4")
	registerFunc(&_audioFileComponentInitializeWithCallbacks, &_audioFileComponentInitializeWithCallbacksErr, frameworkHandle, "AudioFileComponentInitializeWithCallbacks", "10.4")
	registerFunc(&_audioFileComponentOpenFile, &_audioFileComponentOpenFileErr, frameworkHandle, "AudioFileComponentOpenFile", "10.4")
	registerFunc(&_audioFileComponentOpenURL, &_audioFileComponentOpenURLErr, frameworkHandle, "AudioFileComponentOpenURL", "10.5")
	registerFunc(&_audioFileComponentOpenWithCallbacks, &_audioFileComponentOpenWithCallbacksErr, frameworkHandle, "AudioFileComponentOpenWithCallbacks", "10.4")
	registerFunc(&_audioFileComponentOptimize, &_audioFileComponentOptimizeErr, frameworkHandle, "AudioFileComponentOptimize", "10.4")
	registerFunc(&_audioFileComponentReadBytes, &_audioFileComponentReadBytesErr, frameworkHandle, "AudioFileComponentReadBytes", "10.4")
	registerFunc(&_audioFileComponentReadPacketData, &_audioFileComponentReadPacketDataErr, frameworkHandle, "AudioFileComponentReadPacketData", "10.4")
	registerFunc(&_audioFileComponentReadPackets, &_audioFileComponentReadPacketsErr, frameworkHandle, "AudioFileComponentReadPackets", "10.4")
	registerFunc(&_audioFileComponentRemoveUserData, &_audioFileComponentRemoveUserDataErr, frameworkHandle, "AudioFileComponentRemoveUserData", "10.5")
	registerFunc(&_audioFileComponentSetProperty, &_audioFileComponentSetPropertyErr, frameworkHandle, "AudioFileComponentSetProperty", "10.4")
	registerFunc(&_audioFileComponentSetUserData, &_audioFileComponentSetUserDataErr, frameworkHandle, "AudioFileComponentSetUserData", "10.4")
	registerFunc(&_audioFileComponentWriteBytes, &_audioFileComponentWriteBytesErr, frameworkHandle, "AudioFileComponentWriteBytes", "10.4")
	registerFunc(&_audioFileComponentWritePackets, &_audioFileComponentWritePacketsErr, frameworkHandle, "AudioFileComponentWritePackets", "10.4")
	registerFunc(&_audioFileCountUserData, &_audioFileCountUserDataErr, frameworkHandle, "AudioFileCountUserData", "10.4")
	registerFunc(&_audioFileCreateWithURL, &_audioFileCreateWithURLErr, frameworkHandle, "AudioFileCreateWithURL", "10.5")
	registerFunc(&_audioFileGetGlobalInfo, &_audioFileGetGlobalInfoErr, frameworkHandle, "AudioFileGetGlobalInfo", "10.3")
	registerFunc(&_audioFileGetGlobalInfoSize, &_audioFileGetGlobalInfoSizeErr, frameworkHandle, "AudioFileGetGlobalInfoSize", "10.3")
	registerFunc(&_audioFileGetProperty, &_audioFileGetPropertyErr, frameworkHandle, "AudioFileGetProperty", "10.2")
	registerFunc(&_audioFileGetPropertyInfo, &_audioFileGetPropertyInfoErr, frameworkHandle, "AudioFileGetPropertyInfo", "10.2")
	registerFunc(&_audioFileGetUserData, &_audioFileGetUserDataErr, frameworkHandle, "AudioFileGetUserData", "10.4")
	registerFunc(&_audioFileGetUserDataAtOffset, &_audioFileGetUserDataAtOffsetErr, frameworkHandle, "AudioFileGetUserDataAtOffset", "14.0")
	registerFunc(&_audioFileGetUserDataSize, &_audioFileGetUserDataSizeErr, frameworkHandle, "AudioFileGetUserDataSize", "10.4")
	registerFunc(&_audioFileGetUserDataSize64, &_audioFileGetUserDataSize64Err, frameworkHandle, "AudioFileGetUserDataSize64", "14.0")
	registerFunc(&_audioFileInitializeWithCallbacks, &_audioFileInitializeWithCallbacksErr, frameworkHandle, "AudioFileInitializeWithCallbacks", "10.3")
	registerFunc(&_audioFileOpenURL, &_audioFileOpenURLErr, frameworkHandle, "AudioFileOpenURL", "10.5")
	registerFunc(&_audioFileOpenWithCallbacks, &_audioFileOpenWithCallbacksErr, frameworkHandle, "AudioFileOpenWithCallbacks", "10.3")
	registerFunc(&_audioFileOptimize, &_audioFileOptimizeErr, frameworkHandle, "AudioFileOptimize", "10.2")
	registerFunc(&_audioFileReadBytes, &_audioFileReadBytesErr, frameworkHandle, "AudioFileReadBytes", "10.2")
	registerFunc(&_audioFileReadPacketData, &_audioFileReadPacketDataErr, frameworkHandle, "AudioFileReadPacketData", "10.6")
	registerFunc(&_audioFileRemoveUserData, &_audioFileRemoveUserDataErr, frameworkHandle, "AudioFileRemoveUserData", "10.5")
	registerFunc(&_audioFileSetProperty, &_audioFileSetPropertyErr, frameworkHandle, "AudioFileSetProperty", "10.2")
	registerFunc(&_audioFileSetUserData, &_audioFileSetUserDataErr, frameworkHandle, "AudioFileSetUserData", "10.4")
	registerFunc(&_audioFileStreamClose, &_audioFileStreamCloseErr, frameworkHandle, "AudioFileStreamClose", "10.5")
	registerFunc(&_audioFileStreamGetProperty, &_audioFileStreamGetPropertyErr, frameworkHandle, "AudioFileStreamGetProperty", "10.5")
	registerFunc(&_audioFileStreamGetPropertyInfo, &_audioFileStreamGetPropertyInfoErr, frameworkHandle, "AudioFileStreamGetPropertyInfo", "10.5")
	registerFunc(&_audioFileStreamOpen, &_audioFileStreamOpenErr, frameworkHandle, "AudioFileStreamOpen", "10.5")
	registerFunc(&_audioFileStreamParseBytes, &_audioFileStreamParseBytesErr, frameworkHandle, "AudioFileStreamParseBytes", "10.5")
	registerFunc(&_audioFileStreamSeek, &_audioFileStreamSeekErr, frameworkHandle, "AudioFileStreamSeek", "10.5")
	registerFunc(&_audioFileStreamSetProperty, &_audioFileStreamSetPropertyErr, frameworkHandle, "AudioFileStreamSetProperty", "10.5")
	registerFunc(&_audioFileWriteBytes, &_audioFileWriteBytesErr, frameworkHandle, "AudioFileWriteBytes", "10.2")
	registerFunc(&_audioFileWritePackets, &_audioFileWritePacketsErr, frameworkHandle, "AudioFileWritePackets", "10.2")
	registerFunc(&_audioFileWritePacketsWithDependencies, &_audioFileWritePacketsWithDependenciesErr, frameworkHandle, "AudioFileWritePacketsWithDependencies", "26.0")
	registerFunc(&_audioFormatGetProperty, &_audioFormatGetPropertyErr, frameworkHandle, "AudioFormatGetProperty", "10.3")
	registerFunc(&_audioFormatGetPropertyInfo, &_audioFormatGetPropertyInfoErr, frameworkHandle, "AudioFormatGetPropertyInfo", "10.3")
	registerFunc(&_audioOutputUnitStart, &_audioOutputUnitStartErr, frameworkHandle, "AudioOutputUnitStart", "10.0")
	registerFunc(&_audioOutputUnitStop, &_audioOutputUnitStopErr, frameworkHandle, "AudioOutputUnitStop", "10.0")
	registerFunc(&_audioQueueAddPropertyListener, &_audioQueueAddPropertyListenerErr, frameworkHandle, "AudioQueueAddPropertyListener", "10.5")
	registerFunc(&_audioQueueAllocateBuffer, &_audioQueueAllocateBufferErr, frameworkHandle, "AudioQueueAllocateBuffer", "10.5")
	registerFunc(&_audioQueueAllocateBufferWithPacketDescriptions, &_audioQueueAllocateBufferWithPacketDescriptionsErr, frameworkHandle, "AudioQueueAllocateBufferWithPacketDescriptions", "10.6")
	registerFunc(&_audioQueueCreateTimeline, &_audioQueueCreateTimelineErr, frameworkHandle, "AudioQueueCreateTimeline", "10.5")
	registerFunc(&_audioQueueDeviceGetCurrentTime, &_audioQueueDeviceGetCurrentTimeErr, frameworkHandle, "AudioQueueDeviceGetCurrentTime", "10.5")
	registerFunc(&_audioQueueDeviceGetNearestStartTime, &_audioQueueDeviceGetNearestStartTimeErr, frameworkHandle, "AudioQueueDeviceGetNearestStartTime", "10.5")
	registerFunc(&_audioQueueDeviceTranslateTime, &_audioQueueDeviceTranslateTimeErr, frameworkHandle, "AudioQueueDeviceTranslateTime", "10.5")
	registerFunc(&_audioQueueDispose, &_audioQueueDisposeErr, frameworkHandle, "AudioQueueDispose", "10.5")
	registerFunc(&_audioQueueDisposeTimeline, &_audioQueueDisposeTimelineErr, frameworkHandle, "AudioQueueDisposeTimeline", "10.5")
	registerFunc(&_audioQueueEnqueueBuffer, &_audioQueueEnqueueBufferErr, frameworkHandle, "AudioQueueEnqueueBuffer", "10.5")
	registerFunc(&_audioQueueEnqueueBufferWithParameters, &_audioQueueEnqueueBufferWithParametersErr, frameworkHandle, "AudioQueueEnqueueBufferWithParameters", "10.5")
	registerFunc(&_audioQueueFlush, &_audioQueueFlushErr, frameworkHandle, "AudioQueueFlush", "10.5")
	registerFunc(&_audioQueueFreeBuffer, &_audioQueueFreeBufferErr, frameworkHandle, "AudioQueueFreeBuffer", "10.5")
	registerFunc(&_audioQueueGetCurrentTime, &_audioQueueGetCurrentTimeErr, frameworkHandle, "AudioQueueGetCurrentTime", "10.5")
	registerFunc(&_audioQueueGetParameter, &_audioQueueGetParameterErr, frameworkHandle, "AudioQueueGetParameter", "10.5")
	registerFunc(&_audioQueueGetProperty, &_audioQueueGetPropertyErr, frameworkHandle, "AudioQueueGetProperty", "10.5")
	registerFunc(&_audioQueueGetPropertySize, &_audioQueueGetPropertySizeErr, frameworkHandle, "AudioQueueGetPropertySize", "10.5")
	registerFunc(&_audioQueueNewInput, &_audioQueueNewInputErr, frameworkHandle, "AudioQueueNewInput", "10.5")
	registerFunc(&_audioQueueNewInputWithDispatchQueue, &_audioQueueNewInputWithDispatchQueueErr, frameworkHandle, "AudioQueueNewInputWithDispatchQueue", "10.6")
	registerFunc(&_audioQueueNewOutput, &_audioQueueNewOutputErr, frameworkHandle, "AudioQueueNewOutput", "10.5")
	registerFunc(&_audioQueueNewOutputWithDispatchQueue, &_audioQueueNewOutputWithDispatchQueueErr, frameworkHandle, "AudioQueueNewOutputWithDispatchQueue", "10.6")
	registerFunc(&_audioQueueOfflineRender, &_audioQueueOfflineRenderErr, frameworkHandle, "AudioQueueOfflineRender", "10.5")
	registerFunc(&_audioQueuePause, &_audioQueuePauseErr, frameworkHandle, "AudioQueuePause", "10.5")
	registerFunc(&_audioQueuePrime, &_audioQueuePrimeErr, frameworkHandle, "AudioQueuePrime", "10.5")
	registerFunc(&_audioQueueProcessingTapDispose, &_audioQueueProcessingTapDisposeErr, frameworkHandle, "AudioQueueProcessingTapDispose", "10.7")
	registerFunc(&_audioQueueProcessingTapGetQueueTime, &_audioQueueProcessingTapGetQueueTimeErr, frameworkHandle, "AudioQueueProcessingTapGetQueueTime", "10.8")
	registerFunc(&_audioQueueProcessingTapGetSourceAudio, &_audioQueueProcessingTapGetSourceAudioErr, frameworkHandle, "AudioQueueProcessingTapGetSourceAudio", "10.7")
	registerFunc(&_audioQueueProcessingTapNew, &_audioQueueProcessingTapNewErr, frameworkHandle, "AudioQueueProcessingTapNew", "10.7")
	registerFunc(&_audioQueueRemovePropertyListener, &_audioQueueRemovePropertyListenerErr, frameworkHandle, "AudioQueueRemovePropertyListener", "10.5")
	registerFunc(&_audioQueueReset, &_audioQueueResetErr, frameworkHandle, "AudioQueueReset", "10.5")
	registerFunc(&_audioQueueSetOfflineRenderFormat, &_audioQueueSetOfflineRenderFormatErr, frameworkHandle, "AudioQueueSetOfflineRenderFormat", "10.5")
	registerFunc(&_audioQueueSetParameter, &_audioQueueSetParameterErr, frameworkHandle, "AudioQueueSetParameter", "10.5")
	registerFunc(&_audioQueueSetProperty, &_audioQueueSetPropertyErr, frameworkHandle, "AudioQueueSetProperty", "10.5")
	registerFunc(&_audioQueueStart, &_audioQueueStartErr, frameworkHandle, "AudioQueueStart", "10.5")
	registerFunc(&_audioQueueStop, &_audioQueueStopErr, frameworkHandle, "AudioQueueStop", "10.5")
	registerFunc(&_audioServicesAddSystemSoundCompletion, &_audioServicesAddSystemSoundCompletionErr, frameworkHandle, "AudioServicesAddSystemSoundCompletion", "10.5")
	registerFunc(&_audioServicesCreateSystemSoundID, &_audioServicesCreateSystemSoundIDErr, frameworkHandle, "AudioServicesCreateSystemSoundID", "10.5")
	registerFunc(&_audioServicesDisposeSystemSoundID, &_audioServicesDisposeSystemSoundIDErr, frameworkHandle, "AudioServicesDisposeSystemSoundID", "10.5")
	registerFunc(&_audioServicesGetProperty, &_audioServicesGetPropertyErr, frameworkHandle, "AudioServicesGetProperty", "10.5")
	registerFunc(&_audioServicesGetPropertyInfo, &_audioServicesGetPropertyInfoErr, frameworkHandle, "AudioServicesGetPropertyInfo", "10.5")
	registerFunc(&_audioServicesPlayAlertSound, &_audioServicesPlayAlertSoundErr, frameworkHandle, "AudioServicesPlayAlertSound", "10.5")
	registerFunc(&_audioServicesPlayAlertSoundWithCompletion, &_audioServicesPlayAlertSoundWithCompletionErr, frameworkHandle, "AudioServicesPlayAlertSoundWithCompletion", "10.11")
	registerFunc(&_audioServicesPlaySystemSound, &_audioServicesPlaySystemSoundErr, frameworkHandle, "AudioServicesPlaySystemSound", "10.5")
	registerFunc(&_audioServicesPlaySystemSoundWithCompletion, &_audioServicesPlaySystemSoundWithCompletionErr, frameworkHandle, "AudioServicesPlaySystemSoundWithCompletion", "10.11")
	registerFunc(&_audioServicesRemoveSystemSoundCompletion, &_audioServicesRemoveSystemSoundCompletionErr, frameworkHandle, "AudioServicesRemoveSystemSoundCompletion", "10.5")
	registerFunc(&_audioServicesSetProperty, &_audioServicesSetPropertyErr, frameworkHandle, "AudioServicesSetProperty", "10.5")
	registerFunc(&_audioUnitAddPropertyListener, &_audioUnitAddPropertyListenerErr, frameworkHandle, "AudioUnitAddPropertyListener", "10.0")
	registerFunc(&_audioUnitAddRenderNotify, &_audioUnitAddRenderNotifyErr, frameworkHandle, "AudioUnitAddRenderNotify", "10.2")
	registerFunc(&_audioUnitExtensionCopyComponentList, &_audioUnitExtensionCopyComponentListErr, frameworkHandle, "AudioUnitExtensionCopyComponentList", "10.13")
	registerFunc(&_audioUnitExtensionSetComponentList, &_audioUnitExtensionSetComponentListErr, frameworkHandle, "AudioUnitExtensionSetComponentList", "10.13")
	registerFunc(&_audioUnitGetParameter, &_audioUnitGetParameterErr, frameworkHandle, "AudioUnitGetParameter", "10.0")
	registerFunc(&_audioUnitGetProperty, &_audioUnitGetPropertyErr, frameworkHandle, "AudioUnitGetProperty", "10.0")
	registerFunc(&_audioUnitGetPropertyInfo, &_audioUnitGetPropertyInfoErr, frameworkHandle, "AudioUnitGetPropertyInfo", "10.0")
	registerFunc(&_audioUnitInitialize, &_audioUnitInitializeErr, frameworkHandle, "AudioUnitInitialize", "10.0")
	registerFunc(&_audioUnitProcess, &_audioUnitProcessErr, frameworkHandle, "AudioUnitProcess", "10.7")
	registerFunc(&_audioUnitProcessMultiple, &_audioUnitProcessMultipleErr, frameworkHandle, "AudioUnitProcessMultiple", "10.7")
	registerFunc(&_audioUnitRemovePropertyListenerWithUserData, &_audioUnitRemovePropertyListenerWithUserDataErr, frameworkHandle, "AudioUnitRemovePropertyListenerWithUserData", "10.5")
	registerFunc(&_audioUnitRemoveRenderNotify, &_audioUnitRemoveRenderNotifyErr, frameworkHandle, "AudioUnitRemoveRenderNotify", "10.2")
	registerFunc(&_audioUnitRender, &_audioUnitRenderErr, frameworkHandle, "AudioUnitRender", "10.2")
	registerFunc(&_audioUnitReset, &_audioUnitResetErr, frameworkHandle, "AudioUnitReset", "10.0")
	registerFunc(&_audioUnitScheduleParameters, &_audioUnitScheduleParametersErr, frameworkHandle, "AudioUnitScheduleParameters", "10.2")
	registerFunc(&_audioUnitSetParameter, &_audioUnitSetParameterErr, frameworkHandle, "AudioUnitSetParameter", "10.0")
	registerFunc(&_audioUnitSetProperty, &_audioUnitSetPropertyErr, frameworkHandle, "AudioUnitSetProperty", "10.0")
	registerFunc(&_audioUnitUninitialize, &_audioUnitUninitializeErr, frameworkHandle, "AudioUnitUninitialize", "10.0")
	registerFunc(&_audioWorkIntervalCreate, &_audioWorkIntervalCreateErr, frameworkHandle, "AudioWorkIntervalCreate", "11.0")
	registerFunc(&_cAClockAddListener, &_cAClockAddListenerErr, frameworkHandle, "CAClockAddListener", "10.4")
	registerFunc(&_cAClockArm, &_cAClockArmErr, frameworkHandle, "CAClockArm", "10.4")
	registerFunc(&_cAClockBarBeatTimeToBeats, &_cAClockBarBeatTimeToBeatsErr, frameworkHandle, "CAClockBarBeatTimeToBeats", "10.4")
	registerFunc(&_cAClockBeatsToBarBeatTime, &_cAClockBeatsToBarBeatTimeErr, frameworkHandle, "CAClockBeatsToBarBeatTime", "10.4")
	registerFunc(&_cAClockDisarm, &_cAClockDisarmErr, frameworkHandle, "CAClockDisarm", "10.4")
	registerFunc(&_cAClockDispose, &_cAClockDisposeErr, frameworkHandle, "CAClockDispose", "10.4")
	registerFunc(&_cAClockGetCurrentTempo, &_cAClockGetCurrentTempoErr, frameworkHandle, "CAClockGetCurrentTempo", "10.4")
	registerFunc(&_cAClockGetCurrentTime, &_cAClockGetCurrentTimeErr, frameworkHandle, "CAClockGetCurrentTime", "10.4")
	registerFunc(&_cAClockGetPlayRate, &_cAClockGetPlayRateErr, frameworkHandle, "CAClockGetPlayRate", "10.4")
	registerFunc(&_cAClockGetProperty, &_cAClockGetPropertyErr, frameworkHandle, "CAClockGetProperty", "10.4")
	registerFunc(&_cAClockGetPropertyInfo, &_cAClockGetPropertyInfoErr, frameworkHandle, "CAClockGetPropertyInfo", "10.4")
	registerFunc(&_cAClockGetStartTime, &_cAClockGetStartTimeErr, frameworkHandle, "CAClockGetStartTime", "10.4")
	registerFunc(&_cAClockNew, &_cAClockNewErr, frameworkHandle, "CAClockNew", "10.4")
	registerFunc(&_cAClockParseMIDI, &_cAClockParseMIDIErr, frameworkHandle, "CAClockParseMIDI", "10.5")
	registerFunc(&_cAClockRemoveListener, &_cAClockRemoveListenerErr, frameworkHandle, "CAClockRemoveListener", "10.4")
	registerFunc(&_cAClockSMPTETimeToSeconds, &_cAClockSMPTETimeToSecondsErr, frameworkHandle, "CAClockSMPTETimeToSeconds", "10.4")
	registerFunc(&_cAClockSecondsToSMPTETime, &_cAClockSecondsToSMPTETimeErr, frameworkHandle, "CAClockSecondsToSMPTETime", "10.4")
	registerFunc(&_cAClockSetCurrentTempo, &_cAClockSetCurrentTempoErr, frameworkHandle, "CAClockSetCurrentTempo", "10.4")
	registerFunc(&_cAClockSetCurrentTime, &_cAClockSetCurrentTimeErr, frameworkHandle, "CAClockSetCurrentTime", "10.4")
	registerFunc(&_cAClockSetPlayRate, &_cAClockSetPlayRateErr, frameworkHandle, "CAClockSetPlayRate", "10.4")
	registerFunc(&_cAClockSetProperty, &_cAClockSetPropertyErr, frameworkHandle, "CAClockSetProperty", "10.4")
	registerFunc(&_cAClockStart, &_cAClockStartErr, frameworkHandle, "CAClockStart", "10.4")
	registerFunc(&_cAClockStop, &_cAClockStopErr, frameworkHandle, "CAClockStop", "10.4")
	registerFunc(&_cAClockTranslateTime, &_cAClockTranslateTimeErr, frameworkHandle, "CAClockTranslateTime", "10.4")
	registerFunc(&_cAShow, &_cAShowErr, frameworkHandle, "CAShow", "10.2")
	registerFunc(&_cAShowFile, &_cAShowFileErr, frameworkHandle, "CAShowFile", "10.2")
	registerFunc(&_copyInstrumentInfoFromSoundBank, &_copyInstrumentInfoFromSoundBankErr, frameworkHandle, "CopyInstrumentInfoFromSoundBank", "10.8")
	registerFunc(&_copyNameFromSoundBank, &_copyNameFromSoundBankErr, frameworkHandle, "CopyNameFromSoundBank", "10.5")
	registerFunc(&_disposeAUGraph, &_disposeAUGraphErr, frameworkHandle, "DisposeAUGraph", "10.0")
	registerFunc(&_disposeMusicEventIterator, &_disposeMusicEventIteratorErr, frameworkHandle, "DisposeMusicEventIterator", "10.0")
	registerFunc(&_disposeMusicPlayer, &_disposeMusicPlayerErr, frameworkHandle, "DisposeMusicPlayer", "10.0")
	registerFunc(&_disposeMusicSequence, &_disposeMusicSequenceErr, frameworkHandle, "DisposeMusicSequence", "10.0")
	registerFunc(&_extAudioFileCreateNew, &_extAudioFileCreateNewErr, frameworkHandle, "ExtAudioFileCreateNew", "10.4")
	registerFunc(&_extAudioFileCreateWithURL, &_extAudioFileCreateWithURLErr, frameworkHandle, "ExtAudioFileCreateWithURL", "10.5")
	registerFunc(&_extAudioFileDispose, &_extAudioFileDisposeErr, frameworkHandle, "ExtAudioFileDispose", "10.4")
	registerFunc(&_extAudioFileGetProperty, &_extAudioFileGetPropertyErr, frameworkHandle, "ExtAudioFileGetProperty", "10.4")
	registerFunc(&_extAudioFileGetPropertyInfo, &_extAudioFileGetPropertyInfoErr, frameworkHandle, "ExtAudioFileGetPropertyInfo", "10.4")
	registerFunc(&_extAudioFileOpen, &_extAudioFileOpenErr, frameworkHandle, "ExtAudioFileOpen", "10.4")
	registerFunc(&_extAudioFileOpenURL, &_extAudioFileOpenURLErr, frameworkHandle, "ExtAudioFileOpenURL", "10.5")
	registerFunc(&_extAudioFileRead, &_extAudioFileReadErr, frameworkHandle, "ExtAudioFileRead", "10.4")
	registerFunc(&_extAudioFileSeek, &_extAudioFileSeekErr, frameworkHandle, "ExtAudioFileSeek", "10.4")
	registerFunc(&_extAudioFileSetProperty, &_extAudioFileSetPropertyErr, frameworkHandle, "ExtAudioFileSetProperty", "10.4")
	registerFunc(&_extAudioFileTell, &_extAudioFileTellErr, frameworkHandle, "ExtAudioFileTell", "10.4")
	registerFunc(&_extAudioFileWrapAudioFileID, &_extAudioFileWrapAudioFileIDErr, frameworkHandle, "ExtAudioFileWrapAudioFileID", "10.4")
	registerFunc(&_extAudioFileWrite, &_extAudioFileWriteErr, frameworkHandle, "ExtAudioFileWrite", "10.4")
	registerFunc(&_extAudioFileWriteAsync, &_extAudioFileWriteAsyncErr, frameworkHandle, "ExtAudioFileWriteAsync", "10.4")
	registerFunc(&_getNameFromSoundBank, &_getNameFromSoundBankErr, frameworkHandle, "GetNameFromSoundBank", "10.2")
	registerFunc(&_musicDeviceMIDIEvent, &_musicDeviceMIDIEventErr, frameworkHandle, "MusicDeviceMIDIEvent", "10.0")
	registerFunc(&_musicDeviceMIDIEventList, &_musicDeviceMIDIEventListErr, frameworkHandle, "MusicDeviceMIDIEventList", "12.0")
	registerFunc(&_musicDevicePrepareInstrument, &_musicDevicePrepareInstrumentErr, frameworkHandle, "MusicDevicePrepareInstrument", "10.0")
	registerFunc(&_musicDeviceReleaseInstrument, &_musicDeviceReleaseInstrumentErr, frameworkHandle, "MusicDeviceReleaseInstrument", "10.0")
	registerFunc(&_musicDeviceStartNote, &_musicDeviceStartNoteErr, frameworkHandle, "MusicDeviceStartNote", "10.0")
	registerFunc(&_musicDeviceStopNote, &_musicDeviceStopNoteErr, frameworkHandle, "MusicDeviceStopNote", "10.0")
	registerFunc(&_musicDeviceSysEx, &_musicDeviceSysExErr, frameworkHandle, "MusicDeviceSysEx", "10.0")
	registerFunc(&_musicEventIteratorDeleteEvent, &_musicEventIteratorDeleteEventErr, frameworkHandle, "MusicEventIteratorDeleteEvent", "10.0")
	registerFunc(&_musicEventIteratorGetEventInfo, &_musicEventIteratorGetEventInfoErr, frameworkHandle, "MusicEventIteratorGetEventInfo", "10.0")
	registerFunc(&_musicEventIteratorHasCurrentEvent, &_musicEventIteratorHasCurrentEventErr, frameworkHandle, "MusicEventIteratorHasCurrentEvent", "10.2")
	registerFunc(&_musicEventIteratorHasNextEvent, &_musicEventIteratorHasNextEventErr, frameworkHandle, "MusicEventIteratorHasNextEvent", "10.0")
	registerFunc(&_musicEventIteratorHasPreviousEvent, &_musicEventIteratorHasPreviousEventErr, frameworkHandle, "MusicEventIteratorHasPreviousEvent", "10.0")
	registerFunc(&_musicEventIteratorNextEvent, &_musicEventIteratorNextEventErr, frameworkHandle, "MusicEventIteratorNextEvent", "10.0")
	registerFunc(&_musicEventIteratorPreviousEvent, &_musicEventIteratorPreviousEventErr, frameworkHandle, "MusicEventIteratorPreviousEvent", "10.0")
	registerFunc(&_musicEventIteratorSeek, &_musicEventIteratorSeekErr, frameworkHandle, "MusicEventIteratorSeek", "10.0")
	registerFunc(&_musicEventIteratorSetEventInfo, &_musicEventIteratorSetEventInfoErr, frameworkHandle, "MusicEventIteratorSetEventInfo", "10.2")
	registerFunc(&_musicEventIteratorSetEventTime, &_musicEventIteratorSetEventTimeErr, frameworkHandle, "MusicEventIteratorSetEventTime", "10.0")
	registerFunc(&_musicPlayerGetBeatsForHostTime, &_musicPlayerGetBeatsForHostTimeErr, frameworkHandle, "MusicPlayerGetBeatsForHostTime", "10.2")
	registerFunc(&_musicPlayerGetHostTimeForBeats, &_musicPlayerGetHostTimeForBeatsErr, frameworkHandle, "MusicPlayerGetHostTimeForBeats", "10.2")
	registerFunc(&_musicPlayerGetPlayRateScalar, &_musicPlayerGetPlayRateScalarErr, frameworkHandle, "MusicPlayerGetPlayRateScalar", "10.3")
	registerFunc(&_musicPlayerGetSequence, &_musicPlayerGetSequenceErr, frameworkHandle, "MusicPlayerGetSequence", "10.3")
	registerFunc(&_musicPlayerGetTime, &_musicPlayerGetTimeErr, frameworkHandle, "MusicPlayerGetTime", "10.0")
	registerFunc(&_musicPlayerIsPlaying, &_musicPlayerIsPlayingErr, frameworkHandle, "MusicPlayerIsPlaying", "10.2")
	registerFunc(&_musicPlayerPreroll, &_musicPlayerPrerollErr, frameworkHandle, "MusicPlayerPreroll", "10.0")
	registerFunc(&_musicPlayerSetPlayRateScalar, &_musicPlayerSetPlayRateScalarErr, frameworkHandle, "MusicPlayerSetPlayRateScalar", "10.3")
	registerFunc(&_musicPlayerSetSequence, &_musicPlayerSetSequenceErr, frameworkHandle, "MusicPlayerSetSequence", "10.0")
	registerFunc(&_musicPlayerSetTime, &_musicPlayerSetTimeErr, frameworkHandle, "MusicPlayerSetTime", "10.0")
	registerFunc(&_musicPlayerStart, &_musicPlayerStartErr, frameworkHandle, "MusicPlayerStart", "10.0")
	registerFunc(&_musicPlayerStop, &_musicPlayerStopErr, frameworkHandle, "MusicPlayerStop", "10.0")
	registerFunc(&_musicSequenceBarBeatTimeToBeats, &_musicSequenceBarBeatTimeToBeatsErr, frameworkHandle, "MusicSequenceBarBeatTimeToBeats", "10.5")
	registerFunc(&_musicSequenceBeatsToBarBeatTime, &_musicSequenceBeatsToBarBeatTimeErr, frameworkHandle, "MusicSequenceBeatsToBarBeatTime", "10.5")
	registerFunc(&_musicSequenceDisposeTrack, &_musicSequenceDisposeTrackErr, frameworkHandle, "MusicSequenceDisposeTrack", "10.0")
	registerFunc(&_musicSequenceFileCreate, &_musicSequenceFileCreateErr, frameworkHandle, "MusicSequenceFileCreate", "10.5")
	registerFunc(&_musicSequenceFileCreateData, &_musicSequenceFileCreateDataErr, frameworkHandle, "MusicSequenceFileCreateData", "10.5")
	registerFunc(&_musicSequenceFileLoad, &_musicSequenceFileLoadErr, frameworkHandle, "MusicSequenceFileLoad", "10.5")
	registerFunc(&_musicSequenceFileLoadData, &_musicSequenceFileLoadDataErr, frameworkHandle, "MusicSequenceFileLoadData", "10.5")
	registerFunc(&_musicSequenceGetAUGraph, &_musicSequenceGetAUGraphErr, frameworkHandle, "MusicSequenceGetAUGraph", "10.0")
	registerFunc(&_musicSequenceGetBeatsForSeconds, &_musicSequenceGetBeatsForSecondsErr, frameworkHandle, "MusicSequenceGetBeatsForSeconds", "10.2")
	registerFunc(&_musicSequenceGetIndTrack, &_musicSequenceGetIndTrackErr, frameworkHandle, "MusicSequenceGetIndTrack", "10.0")
	registerFunc(&_musicSequenceGetInfoDictionary, &_musicSequenceGetInfoDictionaryErr, frameworkHandle, "MusicSequenceGetInfoDictionary", "10.5")
	registerFunc(&_musicSequenceGetSecondsForBeats, &_musicSequenceGetSecondsForBeatsErr, frameworkHandle, "MusicSequenceGetSecondsForBeats", "10.2")
	registerFunc(&_musicSequenceGetSequenceType, &_musicSequenceGetSequenceTypeErr, frameworkHandle, "MusicSequenceGetSequenceType", "10.5")
	registerFunc(&_musicSequenceGetTempoTrack, &_musicSequenceGetTempoTrackErr, frameworkHandle, "MusicSequenceGetTempoTrack", "10.1")
	registerFunc(&_musicSequenceGetTrackCount, &_musicSequenceGetTrackCountErr, frameworkHandle, "MusicSequenceGetTrackCount", "10.0")
	registerFunc(&_musicSequenceGetTrackIndex, &_musicSequenceGetTrackIndexErr, frameworkHandle, "MusicSequenceGetTrackIndex", "10.0")
	registerFunc(&_musicSequenceLoadSMFDataWithFlags, &_musicSequenceLoadSMFDataWithFlagsErr, frameworkHandle, "MusicSequenceLoadSMFDataWithFlags", "10.3")
	registerFunc(&_musicSequenceLoadSMFWithFlags, &_musicSequenceLoadSMFWithFlagsErr, frameworkHandle, "MusicSequenceLoadSMFWithFlags", "10.3")
	registerFunc(&_musicSequenceNewTrack, &_musicSequenceNewTrackErr, frameworkHandle, "MusicSequenceNewTrack", "10.0")
	registerFunc(&_musicSequenceReverse, &_musicSequenceReverseErr, frameworkHandle, "MusicSequenceReverse", "10.0")
	registerFunc(&_musicSequenceSaveMIDIFile, &_musicSequenceSaveMIDIFileErr, frameworkHandle, "MusicSequenceSaveMIDIFile", "10.4")
	registerFunc(&_musicSequenceSaveSMFData, &_musicSequenceSaveSMFDataErr, frameworkHandle, "MusicSequenceSaveSMFData", "10.2")
	registerFunc(&_musicSequenceSetAUGraph, &_musicSequenceSetAUGraphErr, frameworkHandle, "MusicSequenceSetAUGraph", "10.0")
	registerFunc(&_musicSequenceSetMIDIEndpoint, &_musicSequenceSetMIDIEndpointErr, frameworkHandle, "MusicSequenceSetMIDIEndpoint", "10.1")
	registerFunc(&_musicSequenceSetSequenceType, &_musicSequenceSetSequenceTypeErr, frameworkHandle, "MusicSequenceSetSequenceType", "10.5")
	registerFunc(&_musicSequenceSetUserCallback, &_musicSequenceSetUserCallbackErr, frameworkHandle, "MusicSequenceSetUserCallback", "10.3")
	registerFunc(&_musicTrackClear, &_musicTrackClearErr, frameworkHandle, "MusicTrackClear", "10.0")
	registerFunc(&_musicTrackCopyInsert, &_musicTrackCopyInsertErr, frameworkHandle, "MusicTrackCopyInsert", "10.0")
	registerFunc(&_musicTrackCut, &_musicTrackCutErr, frameworkHandle, "MusicTrackCut", "10.0")
	registerFunc(&_musicTrackGetDestMIDIEndpoint, &_musicTrackGetDestMIDIEndpointErr, frameworkHandle, "MusicTrackGetDestMIDIEndpoint", "10.1")
	registerFunc(&_musicTrackGetDestNode, &_musicTrackGetDestNodeErr, frameworkHandle, "MusicTrackGetDestNode", "10.1")
	registerFunc(&_musicTrackGetProperty, &_musicTrackGetPropertyErr, frameworkHandle, "MusicTrackGetProperty", "10.0")
	registerFunc(&_musicTrackGetSequence, &_musicTrackGetSequenceErr, frameworkHandle, "MusicTrackGetSequence", "10.0")
	registerFunc(&_musicTrackMerge, &_musicTrackMergeErr, frameworkHandle, "MusicTrackMerge", "10.0")
	registerFunc(&_musicTrackMoveEvents, &_musicTrackMoveEventsErr, frameworkHandle, "MusicTrackMoveEvents", "10.0")
	registerFunc(&_musicTrackNewAUPresetEvent, &_musicTrackNewAUPresetEventErr, frameworkHandle, "MusicTrackNewAUPresetEvent", "10.3")
	registerFunc(&_musicTrackNewExtendedControlEvent, &_musicTrackNewExtendedControlEventErr, frameworkHandle, "MusicTrackNewExtendedControlEvent", "10.0")
	registerFunc(&_musicTrackNewExtendedNoteEvent, &_musicTrackNewExtendedNoteEventErr, frameworkHandle, "MusicTrackNewExtendedNoteEvent", "10.0")
	registerFunc(&_musicTrackNewExtendedTempoEvent, &_musicTrackNewExtendedTempoEventErr, frameworkHandle, "MusicTrackNewExtendedTempoEvent", "10.0")
	registerFunc(&_musicTrackNewMIDIChannelEvent, &_musicTrackNewMIDIChannelEventErr, frameworkHandle, "MusicTrackNewMIDIChannelEvent", "10.0")
	registerFunc(&_musicTrackNewMIDINoteEvent, &_musicTrackNewMIDINoteEventErr, frameworkHandle, "MusicTrackNewMIDINoteEvent", "10.0")
	registerFunc(&_musicTrackNewMIDIRawDataEvent, &_musicTrackNewMIDIRawDataEventErr, frameworkHandle, "MusicTrackNewMIDIRawDataEvent", "10.0")
	registerFunc(&_musicTrackNewMetaEvent, &_musicTrackNewMetaEventErr, frameworkHandle, "MusicTrackNewMetaEvent", "10.0")
	registerFunc(&_musicTrackNewParameterEvent, &_musicTrackNewParameterEventErr, frameworkHandle, "MusicTrackNewParameterEvent", "10.2")
	registerFunc(&_musicTrackNewUserEvent, &_musicTrackNewUserEventErr, frameworkHandle, "MusicTrackNewUserEvent", "10.0")
	registerFunc(&_musicTrackSetDestMIDIEndpoint, &_musicTrackSetDestMIDIEndpointErr, frameworkHandle, "MusicTrackSetDestMIDIEndpoint", "10.1")
	registerFunc(&_musicTrackSetDestNode, &_musicTrackSetDestNodeErr, frameworkHandle, "MusicTrackSetDestNode", "10.0")
	registerFunc(&_musicTrackSetProperty, &_musicTrackSetPropertyErr, frameworkHandle, "MusicTrackSetProperty", "10.0")
	registerFunc(&_newMusicEventIterator, &_newMusicEventIteratorErr, frameworkHandle, "NewMusicEventIterator", "10.0")
	registerFunc(&_newMusicPlayer, &_newMusicPlayerErr, frameworkHandle, "NewMusicPlayer", "10.0")
	registerFunc(&_newMusicSequence, &_newMusicSequenceErr, frameworkHandle, "NewMusicSequence", "10.0")
	registerFunc(&_newMusicTrackFrom, &_newMusicTrackFromErr, frameworkHandle, "NewMusicTrackFrom", "10.0")
}
