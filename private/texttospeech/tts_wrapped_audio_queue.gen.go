// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/avfaudio"
)

// The class instance for the [TTSWrappedAudioQueue] class.
var (
	_TTSWrappedAudioQueueClass     TTSWrappedAudioQueueClass
	_TTSWrappedAudioQueueClassOnce sync.Once
)

func getTTSWrappedAudioQueueClass() TTSWrappedAudioQueueClass {
	_TTSWrappedAudioQueueClassOnce.Do(func() {
		_TTSWrappedAudioQueueClass = TTSWrappedAudioQueueClass{class: objc.GetClass("TTSWrappedAudioQueue")}
	})
	return _TTSWrappedAudioQueueClass
}

// GetTTSWrappedAudioQueueClass returns the class object for TTSWrappedAudioQueue.
func GetTTSWrappedAudioQueueClass() TTSWrappedAudioQueueClass {
	return getTTSWrappedAudioQueueClass()
}

type TTSWrappedAudioQueueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSWrappedAudioQueueClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSWrappedAudioQueueClass) Alloc() TTSWrappedAudioQueue {
	rv := objc.Send[TTSWrappedAudioQueue](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSWrappedAudioQueue._attemptQueueStart]
//   - [TTSWrappedAudioQueue._buildAudioQueue]
//   - [TTSWrappedAudioQueue._configureEffects]
//   - [TTSWrappedAudioQueue._initializeDSPGraphAU]
//   - [TTSWrappedAudioQueue._minimumBufferByteSize]
//   - [TTSWrappedAudioQueue._rebuildAudioQueue]
//   - [TTSWrappedAudioQueue._reconfigureQueueFormatForMultiChannelOutputIfNecessary]
//   - [TTSWrappedAudioQueue._startQueueWithRetry]
//   - [TTSWrappedAudioQueue._syncGraphParameters]
//   - [TTSWrappedAudioQueue._syncGraphProperties]
//   - [TTSWrappedAudioQueue._tearDownAudioQueue]
//   - [TTSWrappedAudioQueue._tearDownDSPGraphAU]
//   - [TTSWrappedAudioQueue.AqRef]
//   - [TTSWrappedAudioQueue.SetAqRef]
//   - [TTSWrappedAudioQueue.AudioDevice]
//   - [TTSWrappedAudioQueue.SetAudioDevice]
//   - [TTSWrappedAudioQueue.AudioQueueActive]
//   - [TTSWrappedAudioQueue.AudioQueueFlags]
//   - [TTSWrappedAudioQueue.SetAudioQueueFlags]
//   - [TTSWrappedAudioQueue.BufferCallback]
//   - [TTSWrappedAudioQueue.CachedAudioConverter]
//   - [TTSWrappedAudioQueue.SetCachedAudioConverter]
//   - [TTSWrappedAudioQueue.CallbackQueue]
//   - [TTSWrappedAudioQueue.SetCallbackQueue]
//   - [TTSWrappedAudioQueue.ConvertBufferIfNecessary]
//   - [TTSWrappedAudioQueue.CurrentSilenceBufferCount]
//   - [TTSWrappedAudioQueue.SetCurrentSilenceBufferCount]
//   - [TTSWrappedAudioQueue.DspGraph]
//   - [TTSWrappedAudioQueue.SetDspGraph]
//   - [TTSWrappedAudioQueue.GraphParameters]
//   - [TTSWrappedAudioQueue.SetGraphParameters]
//   - [TTSWrappedAudioQueue.GraphProperties]
//   - [TTSWrappedAudioQueue.SetGraphProperties]
//   - [TTSWrappedAudioQueue.HandleMediaServicesReset]
//   - [TTSWrappedAudioQueue.InflightBuffers]
//   - [TTSWrappedAudioQueue.SetInflightBuffers]
//   - [TTSWrappedAudioQueue.IsRunning]
//   - [TTSWrappedAudioQueue.NeedsParameterSync]
//   - [TTSWrappedAudioQueue.SetNeedsParameterSync]
//   - [TTSWrappedAudioQueue.NeedsPropertySync]
//   - [TTSWrappedAudioQueue.SetNeedsPropertySync]
//   - [TTSWrappedAudioQueue.OutputFormat]
//   - [TTSWrappedAudioQueue.SetOutputFormat]
//   - [TTSWrappedAudioQueue.Pause]
//   - [TTSWrappedAudioQueue.Play]
//   - [TTSWrappedAudioQueue.PlayBufferCompletionHandler]
//   - [TTSWrappedAudioQueue.ProcNodeRef]
//   - [TTSWrappedAudioQueue.SetProcNodeRef]
//   - [TTSWrappedAudioQueue.QueueFormat]
//   - [TTSWrappedAudioQueue.SetQueueFormat]
//   - [TTSWrappedAudioQueue.QueueStreamDescription]
//   - [TTSWrappedAudioQueue.ScheduleBufferCompletionHandler]
//   - [TTSWrappedAudioQueue.ScheduleBufferCompletionHandlerLastBuffer]
//   - [TTSWrappedAudioQueue.ShouldRebuildAudioQueue]
//   - [TTSWrappedAudioQueue.SetShouldRebuildAudioQueue]
//   - [TTSWrappedAudioQueue.State]
//   - [TTSWrappedAudioQueue.SetState]
//   - [TTSWrappedAudioQueue.Stop]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue
type TTSWrappedAudioQueue struct {
	objectivec.Object
}

// TTSWrappedAudioQueueFromID constructs a [TTSWrappedAudioQueue] from an objc.ID.
func TTSWrappedAudioQueueFromID(id objc.ID) TTSWrappedAudioQueue {
	return TTSWrappedAudioQueue{objectivec.Object{ID: id}}
}

// Ensure TTSWrappedAudioQueue implements ITTSWrappedAudioQueue.
var _ ITTSWrappedAudioQueue = TTSWrappedAudioQueue{}

// An interface definition for the [TTSWrappedAudioQueue] class.
//
// # Methods
//
//   - [ITTSWrappedAudioQueue._attemptQueueStart]
//   - [ITTSWrappedAudioQueue._buildAudioQueue]
//   - [ITTSWrappedAudioQueue._configureEffects]
//   - [ITTSWrappedAudioQueue._initializeDSPGraphAU]
//   - [ITTSWrappedAudioQueue._minimumBufferByteSize]
//   - [ITTSWrappedAudioQueue._rebuildAudioQueue]
//   - [ITTSWrappedAudioQueue._reconfigureQueueFormatForMultiChannelOutputIfNecessary]
//   - [ITTSWrappedAudioQueue._startQueueWithRetry]
//   - [ITTSWrappedAudioQueue._syncGraphParameters]
//   - [ITTSWrappedAudioQueue._syncGraphProperties]
//   - [ITTSWrappedAudioQueue._tearDownAudioQueue]
//   - [ITTSWrappedAudioQueue._tearDownDSPGraphAU]
//   - [ITTSWrappedAudioQueue.AqRef]
//   - [ITTSWrappedAudioQueue.SetAqRef]
//   - [ITTSWrappedAudioQueue.AudioDevice]
//   - [ITTSWrappedAudioQueue.SetAudioDevice]
//   - [ITTSWrappedAudioQueue.AudioQueueActive]
//   - [ITTSWrappedAudioQueue.AudioQueueFlags]
//   - [ITTSWrappedAudioQueue.SetAudioQueueFlags]
//   - [ITTSWrappedAudioQueue.BufferCallback]
//   - [ITTSWrappedAudioQueue.CachedAudioConverter]
//   - [ITTSWrappedAudioQueue.SetCachedAudioConverter]
//   - [ITTSWrappedAudioQueue.CallbackQueue]
//   - [ITTSWrappedAudioQueue.SetCallbackQueue]
//   - [ITTSWrappedAudioQueue.ConvertBufferIfNecessary]
//   - [ITTSWrappedAudioQueue.CurrentSilenceBufferCount]
//   - [ITTSWrappedAudioQueue.SetCurrentSilenceBufferCount]
//   - [ITTSWrappedAudioQueue.DspGraph]
//   - [ITTSWrappedAudioQueue.SetDspGraph]
//   - [ITTSWrappedAudioQueue.GraphParameters]
//   - [ITTSWrappedAudioQueue.SetGraphParameters]
//   - [ITTSWrappedAudioQueue.GraphProperties]
//   - [ITTSWrappedAudioQueue.SetGraphProperties]
//   - [ITTSWrappedAudioQueue.HandleMediaServicesReset]
//   - [ITTSWrappedAudioQueue.InflightBuffers]
//   - [ITTSWrappedAudioQueue.SetInflightBuffers]
//   - [ITTSWrappedAudioQueue.IsRunning]
//   - [ITTSWrappedAudioQueue.NeedsParameterSync]
//   - [ITTSWrappedAudioQueue.SetNeedsParameterSync]
//   - [ITTSWrappedAudioQueue.NeedsPropertySync]
//   - [ITTSWrappedAudioQueue.SetNeedsPropertySync]
//   - [ITTSWrappedAudioQueue.OutputFormat]
//   - [ITTSWrappedAudioQueue.SetOutputFormat]
//   - [ITTSWrappedAudioQueue.Pause]
//   - [ITTSWrappedAudioQueue.Play]
//   - [ITTSWrappedAudioQueue.PlayBufferCompletionHandler]
//   - [ITTSWrappedAudioQueue.ProcNodeRef]
//   - [ITTSWrappedAudioQueue.SetProcNodeRef]
//   - [ITTSWrappedAudioQueue.QueueFormat]
//   - [ITTSWrappedAudioQueue.SetQueueFormat]
//   - [ITTSWrappedAudioQueue.QueueStreamDescription]
//   - [ITTSWrappedAudioQueue.ScheduleBufferCompletionHandler]
//   - [ITTSWrappedAudioQueue.ScheduleBufferCompletionHandlerLastBuffer]
//   - [ITTSWrappedAudioQueue.ShouldRebuildAudioQueue]
//   - [ITTSWrappedAudioQueue.SetShouldRebuildAudioQueue]
//   - [ITTSWrappedAudioQueue.State]
//   - [ITTSWrappedAudioQueue.SetState]
//   - [ITTSWrappedAudioQueue.Stop]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue
type ITTSWrappedAudioQueue interface {
	objectivec.IObject

	// Topic: Methods

	_attemptQueueStart() bool
	_buildAudioQueue()
	_configureEffects()
	_initializeDSPGraphAU()
	_minimumBufferByteSize() uint64
	_rebuildAudioQueue()
	_reconfigureQueueFormatForMultiChannelOutputIfNecessary()
	_startQueueWithRetry() bool
	_syncGraphParameters()
	_syncGraphProperties()
	_tearDownAudioQueue()
	_tearDownDSPGraphAU()
	AqRef() avfaudio.OpaqueAudioQueueRef
	SetAqRef(value avfaudio.OpaqueAudioQueueRef)
	AudioDevice() uint32
	SetAudioDevice(value uint32)
	AudioQueueActive() bool
	AudioQueueFlags() uint32
	SetAudioQueueFlags(value uint32)
	BufferCallback(callback *avfaudio.AudioQueueBufferRef)
	CachedAudioConverter() unsafe.Pointer
	SetCachedAudioConverter(value unsafe.Pointer)
	CallbackQueue() objectivec.Object
	SetCallbackQueue(value objectivec.Object)
	ConvertBufferIfNecessary(necessary objectivec.IObject) objectivec.IObject
	CurrentSilenceBufferCount() foundation.NSNumber
	SetCurrentSilenceBufferCount(value foundation.NSNumber)
	DspGraph() string
	SetDspGraph(value string)
	GraphParameters() foundation.INSDictionary
	SetGraphParameters(value foundation.INSDictionary)
	GraphProperties() foundation.INSDictionary
	SetGraphProperties(value foundation.INSDictionary)
	HandleMediaServicesReset()
	InflightBuffers() foundation.INSOrderedSet
	SetInflightBuffers(value foundation.INSOrderedSet)
	IsRunning() bool
	NeedsParameterSync() bool
	SetNeedsParameterSync(value bool)
	NeedsPropertySync() bool
	SetNeedsPropertySync(value bool)
	OutputFormat() ITTSAudioFormat
	SetOutputFormat(value ITTSAudioFormat)
	Pause()
	Play() bool
	PlayBufferCompletionHandler(buffer objectivec.IObject, handler ErrorHandler)
	ProcNodeRef() OpaqueATAudioProcessingNodeRef
	SetProcNodeRef(value OpaqueATAudioProcessingNodeRef)
	QueueFormat() avfaudio.AVAudioFormat
	SetQueueFormat(value avfaudio.AVAudioFormat)
	QueueStreamDescription() objectivec.IObject
	ScheduleBufferCompletionHandler(buffer objectivec.IObject, handler ErrorHandler)
	ScheduleBufferCompletionHandlerLastBuffer(buffer objectivec.IObject, handler ErrorHandler, buffer2 bool)
	ShouldRebuildAudioQueue() bool
	SetShouldRebuildAudioQueue(value bool)
	State() uint64
	SetState(value uint64)
	Stop()
}

// Init initializes the instance.
func (t TTSWrappedAudioQueue) Init() TTSWrappedAudioQueue {
	rv := objc.Send[TTSWrappedAudioQueue](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSWrappedAudioQueue) Autorelease() TTSWrappedAudioQueue {
	rv := objc.Send[TTSWrappedAudioQueue](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSWrappedAudioQueue creates a new TTSWrappedAudioQueue instance.
func NewTTSWrappedAudioQueue() TTSWrappedAudioQueue {
	class := getTTSWrappedAudioQueueClass()
	rv := objc.Send[TTSWrappedAudioQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_attemptQueueStart
func (t TTSWrappedAudioQueue) _attemptQueueStart() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_attemptQueueStart"))
	return rv
}

// AttemptQueueStart is an exported wrapper for the private method _attemptQueueStart.
func (t TTSWrappedAudioQueue) AttemptQueueStart() bool {
	return t._attemptQueueStart()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_buildAudioQueue
func (t TTSWrappedAudioQueue) _buildAudioQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("_buildAudioQueue"))
}

// BuildAudioQueue is an exported wrapper for the private method _buildAudioQueue.
func (t TTSWrappedAudioQueue) BuildAudioQueue() {
	t._buildAudioQueue()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_configureEffects
func (t TTSWrappedAudioQueue) _configureEffects() {
	objc.Send[objc.ID](t.ID, objc.Sel("_configureEffects"))
}

// ConfigureEffects is an exported wrapper for the private method _configureEffects.
func (t TTSWrappedAudioQueue) ConfigureEffects() {
	t._configureEffects()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_initializeDSPGraphAU
func (t TTSWrappedAudioQueue) _initializeDSPGraphAU() {
	objc.Send[objc.ID](t.ID, objc.Sel("_initializeDSPGraphAU"))
}

// InitializeDSPGraphAU is an exported wrapper for the private method _initializeDSPGraphAU.
func (t TTSWrappedAudioQueue) InitializeDSPGraphAU() {
	t._initializeDSPGraphAU()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_minimumBufferByteSize
func (t TTSWrappedAudioQueue) _minimumBufferByteSize() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("_minimumBufferByteSize"))
	return rv
}

// MinimumBufferByteSize is an exported wrapper for the private method _minimumBufferByteSize.
func (t TTSWrappedAudioQueue) MinimumBufferByteSize() uint64 {
	return t._minimumBufferByteSize()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_rebuildAudioQueue
func (t TTSWrappedAudioQueue) _rebuildAudioQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("_rebuildAudioQueue"))
}

// RebuildAudioQueue is an exported wrapper for the private method _rebuildAudioQueue.
func (t TTSWrappedAudioQueue) RebuildAudioQueue() {
	t._rebuildAudioQueue()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_reconfigureQueueFormatForMultiChannelOutputIfNecessary
func (t TTSWrappedAudioQueue) _reconfigureQueueFormatForMultiChannelOutputIfNecessary() {
	objc.Send[objc.ID](t.ID, objc.Sel("_reconfigureQueueFormatForMultiChannelOutputIfNecessary"))
}

// ReconfigureQueueFormatForMultiChannelOutputIfNecessary is an exported wrapper for the private method _reconfigureQueueFormatForMultiChannelOutputIfNecessary.
func (t TTSWrappedAudioQueue) ReconfigureQueueFormatForMultiChannelOutputIfNecessary() {
	t._reconfigureQueueFormatForMultiChannelOutputIfNecessary()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_startQueueWithRetry
func (t TTSWrappedAudioQueue) _startQueueWithRetry() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_startQueueWithRetry"))
	return rv
}

// StartQueueWithRetry is an exported wrapper for the private method _startQueueWithRetry.
func (t TTSWrappedAudioQueue) StartQueueWithRetry() bool {
	return t._startQueueWithRetry()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_syncGraphParameters
func (t TTSWrappedAudioQueue) _syncGraphParameters() {
	objc.Send[objc.ID](t.ID, objc.Sel("_syncGraphParameters"))
}

// SyncGraphParameters is an exported wrapper for the private method _syncGraphParameters.
func (t TTSWrappedAudioQueue) SyncGraphParameters() {
	t._syncGraphParameters()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_syncGraphProperties
func (t TTSWrappedAudioQueue) _syncGraphProperties() {
	objc.Send[objc.ID](t.ID, objc.Sel("_syncGraphProperties"))
}

// SyncGraphProperties is an exported wrapper for the private method _syncGraphProperties.
func (t TTSWrappedAudioQueue) SyncGraphProperties() {
	t._syncGraphProperties()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_tearDownAudioQueue
func (t TTSWrappedAudioQueue) _tearDownAudioQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("_tearDownAudioQueue"))
}

// TearDownAudioQueue is an exported wrapper for the private method _tearDownAudioQueue.
func (t TTSWrappedAudioQueue) TearDownAudioQueue() {
	t._tearDownAudioQueue()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/_tearDownDSPGraphAU
func (t TTSWrappedAudioQueue) _tearDownDSPGraphAU() {
	objc.Send[objc.ID](t.ID, objc.Sel("_tearDownDSPGraphAU"))
}

// TearDownDSPGraphAU is an exported wrapper for the private method _tearDownDSPGraphAU.
func (t TTSWrappedAudioQueue) TearDownDSPGraphAU() {
	t._tearDownDSPGraphAU()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/bufferCallback:
func (t TTSWrappedAudioQueue) BufferCallback(callback *avfaudio.AudioQueueBufferRef) {
	objc.Send[objc.ID](t.ID, objc.Sel("bufferCallback:"), callback)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/convertBufferIfNecessary:
func (t TTSWrappedAudioQueue) ConvertBufferIfNecessary(necessary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("convertBufferIfNecessary:"), necessary)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/handleMediaServicesReset
func (t TTSWrappedAudioQueue) HandleMediaServicesReset() {
	objc.Send[objc.ID](t.ID, objc.Sel("handleMediaServicesReset"))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/isRunning
func (t TTSWrappedAudioQueue) IsRunning() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isRunning"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/pause
func (t TTSWrappedAudioQueue) Pause() {
	objc.Send[objc.ID](t.ID, objc.Sel("pause"))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/play
func (t TTSWrappedAudioQueue) Play() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("play"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/playBuffer:completionHandler:
func (t TTSWrappedAudioQueue) PlayBufferCompletionHandler(buffer objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("playBuffer:completionHandler:"), buffer, _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/queueStreamDescription
func (t TTSWrappedAudioQueue) QueueStreamDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("queueStreamDescription"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/scheduleBuffer:completionHandler:
func (t TTSWrappedAudioQueue) ScheduleBufferCompletionHandler(buffer objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("scheduleBuffer:completionHandler:"), buffer, _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/scheduleBuffer:completionHandler:lastBuffer:
func (t TTSWrappedAudioQueue) ScheduleBufferCompletionHandlerLastBuffer(buffer objectivec.IObject, handler ErrorHandler, buffer2 bool) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("scheduleBuffer:completionHandler:lastBuffer:"), buffer, _block1, buffer2)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/stop
func (t TTSWrappedAudioQueue) Stop() {
	objc.Send[objc.ID](t.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/aqRef
func (t TTSWrappedAudioQueue) AqRef() avfaudio.OpaqueAudioQueueRef {
	rv := objc.Send[avfaudio.OpaqueAudioQueueRef](t.ID, objc.Sel("aqRef"))
	return avfaudio.OpaqueAudioQueueRef(rv)
}
func (t TTSWrappedAudioQueue) SetAqRef(value avfaudio.OpaqueAudioQueueRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setAqRef:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/audioDevice
func (t TTSWrappedAudioQueue) AudioDevice() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioDevice"))
	return rv
}
func (t TTSWrappedAudioQueue) SetAudioDevice(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioDevice:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/audioQueueActive
func (t TTSWrappedAudioQueue) AudioQueueActive() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("audioQueueActive"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/audioQueueFlags
func (t TTSWrappedAudioQueue) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TTSWrappedAudioQueue) SetAudioQueueFlags(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/cachedAudioConverter
func (t TTSWrappedAudioQueue) CachedAudioConverter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("cachedAudioConverter"))
	return rv
}
func (t TTSWrappedAudioQueue) SetCachedAudioConverter(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setCachedAudioConverter:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/callbackQueue
func (t TTSWrappedAudioQueue) CallbackQueue() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("callbackQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetCallbackQueue(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setCallbackQueue:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/currentSilenceBufferCount
func (t TTSWrappedAudioQueue) CurrentSilenceBufferCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("currentSilenceBufferCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetCurrentSilenceBufferCount(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setCurrentSilenceBufferCount:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/dspGraph
func (t TTSWrappedAudioQueue) DspGraph() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dspGraph"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSWrappedAudioQueue) SetDspGraph(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setDspGraph:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/graphParameters
func (t TTSWrappedAudioQueue) GraphParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("graphParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetGraphParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setGraphParameters:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/graphProperties
func (t TTSWrappedAudioQueue) GraphProperties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("graphProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetGraphProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setGraphProperties:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/inflightBuffers
func (t TTSWrappedAudioQueue) InflightBuffers() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("inflightBuffers"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetInflightBuffers(value foundation.INSOrderedSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setInflightBuffers:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/needsParameterSync
func (t TTSWrappedAudioQueue) NeedsParameterSync() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("needsParameterSync"))
	return rv
}
func (t TTSWrappedAudioQueue) SetNeedsParameterSync(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setNeedsParameterSync:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/needsPropertySync
func (t TTSWrappedAudioQueue) NeedsPropertySync() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("needsPropertySync"))
	return rv
}
func (t TTSWrappedAudioQueue) SetNeedsPropertySync(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setNeedsPropertySync:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/outputFormat
func (t TTSWrappedAudioQueue) OutputFormat() ITTSAudioFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("outputFormat"))
	return TTSAudioFormatFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetOutputFormat(value ITTSAudioFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setOutputFormat:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/procNodeRef
func (t TTSWrappedAudioQueue) ProcNodeRef() OpaqueATAudioProcessingNodeRef {
	rv := objc.Send[OpaqueATAudioProcessingNodeRef](t.ID, objc.Sel("procNodeRef"))
	return OpaqueATAudioProcessingNodeRef(rv)
}
func (t TTSWrappedAudioQueue) SetProcNodeRef(value OpaqueATAudioProcessingNodeRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setProcNodeRef:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/queueFormat
func (t TTSWrappedAudioQueue) QueueFormat() avfaudio.AVAudioFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("queueFormat"))
	return avfaudio.AVAudioFormatFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetQueueFormat(value avfaudio.AVAudioFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setQueueFormat:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/shouldRebuildAudioQueue
func (t TTSWrappedAudioQueue) ShouldRebuildAudioQueue() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldRebuildAudioQueue"))
	return rv
}
func (t TTSWrappedAudioQueue) SetShouldRebuildAudioQueue(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldRebuildAudioQueue:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueue/state
func (t TTSWrappedAudioQueue) State() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("state"))
	return rv
}
func (t TTSWrappedAudioQueue) SetState(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setState:"), value)
}

// PlayBuffer is a synchronous wrapper around [TTSWrappedAudioQueue.PlayBufferCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSWrappedAudioQueue) PlayBuffer(ctx context.Context, buffer objectivec.IObject) error {
	done := make(chan error, 1)
	t.PlayBufferCompletionHandler(buffer, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ScheduleBuffer is a synchronous wrapper around [TTSWrappedAudioQueue.ScheduleBufferCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSWrappedAudioQueue) ScheduleBuffer(ctx context.Context, buffer objectivec.IObject) error {
	done := make(chan error, 1)
	t.ScheduleBufferCompletionHandler(buffer, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
