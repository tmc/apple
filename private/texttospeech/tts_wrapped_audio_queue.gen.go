// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
	AqRef() OpaqueAudioQueueRef
	SetAqRef(value OpaqueAudioQueueRef)
	AudioDevice() uint32
	SetAudioDevice(value uint32)
	AudioQueueActive() bool
	AudioQueueFlags() uint32
	SetAudioQueueFlags(value uint32)
	BufferCallback(callback avfaudio.AudioQueueBuffer)
	CachedAudioConverter() avfaudio.AVAudioConverter
	SetCachedAudioConverter(value avfaudio.AVAudioConverter)
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
	QueueStreamDescription() coreaudiotypes.AudioStreamBasicDescription
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

func (t TTSWrappedAudioQueue) _attemptQueueStart() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_attemptQueueStart"))
	return rv
}

// AttemptQueueStart is an exported wrapper for the private method _attemptQueueStart.
func (t TTSWrappedAudioQueue) AttemptQueueStart() (bool, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_attemptQueueStart")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_attemptQueueStart"}
		return false, err
	}
	return t._attemptQueueStart(), nil
}

// CanAttemptQueueStart reports whether the receiver responds to the private selector _attemptQueueStart.
func (t TTSWrappedAudioQueue) CanAttemptQueueStart() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_attemptQueueStart"))
}
func (t TTSWrappedAudioQueue) _buildAudioQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("_buildAudioQueue"))
}

// BuildAudioQueue is an exported wrapper for the private method _buildAudioQueue.
func (t TTSWrappedAudioQueue) BuildAudioQueue() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_buildAudioQueue")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_buildAudioQueue"}
		return err
	}
	t._buildAudioQueue()
	return nil
}

// CanBuildAudioQueue reports whether the receiver responds to the private selector _buildAudioQueue.
func (t TTSWrappedAudioQueue) CanBuildAudioQueue() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_buildAudioQueue"))
}
func (t TTSWrappedAudioQueue) _configureEffects() {
	objc.Send[objc.ID](t.ID, objc.Sel("_configureEffects"))
}

// ConfigureEffects is an exported wrapper for the private method _configureEffects.
func (t TTSWrappedAudioQueue) ConfigureEffects() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_configureEffects")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_configureEffects"}
		return err
	}
	t._configureEffects()
	return nil
}

// CanConfigureEffects reports whether the receiver responds to the private selector _configureEffects.
func (t TTSWrappedAudioQueue) CanConfigureEffects() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_configureEffects"))
}
func (t TTSWrappedAudioQueue) _initializeDSPGraphAU() {
	objc.Send[objc.ID](t.ID, objc.Sel("_initializeDSPGraphAU"))
}

// InitializeDSPGraphAU is an exported wrapper for the private method _initializeDSPGraphAU.
func (t TTSWrappedAudioQueue) InitializeDSPGraphAU() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_initializeDSPGraphAU")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initializeDSPGraphAU"}
		return err
	}
	t._initializeDSPGraphAU()
	return nil
}

// CanInitializeDSPGraphAU reports whether the receiver responds to the private selector _initializeDSPGraphAU.
func (t TTSWrappedAudioQueue) CanInitializeDSPGraphAU() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_initializeDSPGraphAU"))
}
func (t TTSWrappedAudioQueue) _minimumBufferByteSize() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("_minimumBufferByteSize"))
	return rv
}

// MinimumBufferByteSize is an exported wrapper for the private method _minimumBufferByteSize.
func (t TTSWrappedAudioQueue) MinimumBufferByteSize() (uint64, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_minimumBufferByteSize")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_minimumBufferByteSize"}
		return 0, err
	}
	return t._minimumBufferByteSize(), nil
}

// CanMinimumBufferByteSize reports whether the receiver responds to the private selector _minimumBufferByteSize.
func (t TTSWrappedAudioQueue) CanMinimumBufferByteSize() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_minimumBufferByteSize"))
}
func (t TTSWrappedAudioQueue) _rebuildAudioQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("_rebuildAudioQueue"))
}

// RebuildAudioQueue is an exported wrapper for the private method _rebuildAudioQueue.
func (t TTSWrappedAudioQueue) RebuildAudioQueue() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_rebuildAudioQueue")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_rebuildAudioQueue"}
		return err
	}
	t._rebuildAudioQueue()
	return nil
}

// CanRebuildAudioQueue reports whether the receiver responds to the private selector _rebuildAudioQueue.
func (t TTSWrappedAudioQueue) CanRebuildAudioQueue() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_rebuildAudioQueue"))
}
func (t TTSWrappedAudioQueue) _reconfigureQueueFormatForMultiChannelOutputIfNecessary() {
	objc.Send[objc.ID](t.ID, objc.Sel("_reconfigureQueueFormatForMultiChannelOutputIfNecessary"))
}

// ReconfigureQueueFormatForMultiChannelOutputIfNecessary is an exported wrapper for the private method _reconfigureQueueFormatForMultiChannelOutputIfNecessary.
func (t TTSWrappedAudioQueue) ReconfigureQueueFormatForMultiChannelOutputIfNecessary() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_reconfigureQueueFormatForMultiChannelOutputIfNecessary")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_reconfigureQueueFormatForMultiChannelOutputIfNecessary"}
		return err
	}
	t._reconfigureQueueFormatForMultiChannelOutputIfNecessary()
	return nil
}

// CanReconfigureQueueFormatForMultiChannelOutputIfNecessary reports whether the receiver responds to the private selector _reconfigureQueueFormatForMultiChannelOutputIfNecessary.
func (t TTSWrappedAudioQueue) CanReconfigureQueueFormatForMultiChannelOutputIfNecessary() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_reconfigureQueueFormatForMultiChannelOutputIfNecessary"))
}
func (t TTSWrappedAudioQueue) _startQueueWithRetry() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_startQueueWithRetry"))
	return rv
}

// StartQueueWithRetry is an exported wrapper for the private method _startQueueWithRetry.
func (t TTSWrappedAudioQueue) StartQueueWithRetry() (bool, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_startQueueWithRetry")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_startQueueWithRetry"}
		return false, err
	}
	return t._startQueueWithRetry(), nil
}

// CanStartQueueWithRetry reports whether the receiver responds to the private selector _startQueueWithRetry.
func (t TTSWrappedAudioQueue) CanStartQueueWithRetry() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_startQueueWithRetry"))
}
func (t TTSWrappedAudioQueue) _syncGraphParameters() {
	objc.Send[objc.ID](t.ID, objc.Sel("_syncGraphParameters"))
}

// SyncGraphParameters is an exported wrapper for the private method _syncGraphParameters.
func (t TTSWrappedAudioQueue) SyncGraphParameters() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_syncGraphParameters")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_syncGraphParameters"}
		return err
	}
	t._syncGraphParameters()
	return nil
}

// CanSyncGraphParameters reports whether the receiver responds to the private selector _syncGraphParameters.
func (t TTSWrappedAudioQueue) CanSyncGraphParameters() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_syncGraphParameters"))
}
func (t TTSWrappedAudioQueue) _syncGraphProperties() {
	objc.Send[objc.ID](t.ID, objc.Sel("_syncGraphProperties"))
}

// SyncGraphProperties is an exported wrapper for the private method _syncGraphProperties.
func (t TTSWrappedAudioQueue) SyncGraphProperties() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_syncGraphProperties")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_syncGraphProperties"}
		return err
	}
	t._syncGraphProperties()
	return nil
}

// CanSyncGraphProperties reports whether the receiver responds to the private selector _syncGraphProperties.
func (t TTSWrappedAudioQueue) CanSyncGraphProperties() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_syncGraphProperties"))
}
func (t TTSWrappedAudioQueue) _tearDownAudioQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("_tearDownAudioQueue"))
}

// TearDownAudioQueue is an exported wrapper for the private method _tearDownAudioQueue.
func (t TTSWrappedAudioQueue) TearDownAudioQueue() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_tearDownAudioQueue")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_tearDownAudioQueue"}
		return err
	}
	t._tearDownAudioQueue()
	return nil
}

// CanTearDownAudioQueue reports whether the receiver responds to the private selector _tearDownAudioQueue.
func (t TTSWrappedAudioQueue) CanTearDownAudioQueue() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_tearDownAudioQueue"))
}
func (t TTSWrappedAudioQueue) _tearDownDSPGraphAU() {
	objc.Send[objc.ID](t.ID, objc.Sel("_tearDownDSPGraphAU"))
}

// TearDownDSPGraphAU is an exported wrapper for the private method _tearDownDSPGraphAU.
func (t TTSWrappedAudioQueue) TearDownDSPGraphAU() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_tearDownDSPGraphAU")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_tearDownDSPGraphAU"}
		return err
	}
	t._tearDownDSPGraphAU()
	return nil
}

// CanTearDownDSPGraphAU reports whether the receiver responds to the private selector _tearDownDSPGraphAU.
func (t TTSWrappedAudioQueue) CanTearDownDSPGraphAU() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_tearDownDSPGraphAU"))
}
func (t TTSWrappedAudioQueue) BufferCallback(callback avfaudio.AudioQueueBuffer) {
	objc.Send[objc.ID](t.ID, objc.Sel("bufferCallback:"), callback)
}
func (t TTSWrappedAudioQueue) ConvertBufferIfNecessary(necessary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("convertBufferIfNecessary:"), necessary)
	return objectivec.Object{ID: rv}
}
func (t TTSWrappedAudioQueue) HandleMediaServicesReset() {
	objc.Send[objc.ID](t.ID, objc.Sel("handleMediaServicesReset"))
}
func (t TTSWrappedAudioQueue) IsRunning() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isRunning"))
	return rv
}
func (t TTSWrappedAudioQueue) Pause() {
	objc.Send[objc.ID](t.ID, objc.Sel("pause"))
}
func (t TTSWrappedAudioQueue) Play() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("play"))
	return rv
}
func (t TTSWrappedAudioQueue) PlayBufferCompletionHandler(buffer objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("playBuffer:completionHandler:"), buffer, _block1)
}
func (t TTSWrappedAudioQueue) QueueStreamDescription() coreaudiotypes.AudioStreamBasicDescription {
	rv := objc.Send[coreaudiotypes.AudioStreamBasicDescription](t.ID, objc.Sel("queueStreamDescription"))
	return coreaudiotypes.AudioStreamBasicDescription(rv)
}
func (t TTSWrappedAudioQueue) ScheduleBufferCompletionHandler(buffer objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("scheduleBuffer:completionHandler:"), buffer, _block1)
}
func (t TTSWrappedAudioQueue) ScheduleBufferCompletionHandlerLastBuffer(buffer objectivec.IObject, handler ErrorHandler, buffer2 bool) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("scheduleBuffer:completionHandler:lastBuffer:"), buffer, _block1, buffer2)
}
func (t TTSWrappedAudioQueue) Stop() {
	objc.Send[objc.ID](t.ID, objc.Sel("stop"))
}

func (t TTSWrappedAudioQueue) AqRef() OpaqueAudioQueueRef {
	rv := objc.Send[OpaqueAudioQueueRef](t.ID, objc.Sel("aqRef"))
	return OpaqueAudioQueueRef(rv)
}
func (t TTSWrappedAudioQueue) SetAqRef(value OpaqueAudioQueueRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setAqRef:"), value)
}
func (t TTSWrappedAudioQueue) AudioDevice() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioDevice"))
	return rv
}
func (t TTSWrappedAudioQueue) SetAudioDevice(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioDevice:"), value)
}
func (t TTSWrappedAudioQueue) AudioQueueActive() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("audioQueueActive"))
	return rv
}
func (t TTSWrappedAudioQueue) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TTSWrappedAudioQueue) SetAudioQueueFlags(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}
func (t TTSWrappedAudioQueue) CachedAudioConverter() avfaudio.AVAudioConverter {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("cachedAudioConverter"))
	return avfaudio.AVAudioConverterFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetCachedAudioConverter(value avfaudio.AVAudioConverter) {
	objc.Send[struct{}](t.ID, objc.Sel("setCachedAudioConverter:"), value)
}
func (t TTSWrappedAudioQueue) CallbackQueue() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("callbackQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetCallbackQueue(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setCallbackQueue:"), value)
}
func (t TTSWrappedAudioQueue) CurrentSilenceBufferCount() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("currentSilenceBufferCount"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetCurrentSilenceBufferCount(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setCurrentSilenceBufferCount:"), value)
}
func (t TTSWrappedAudioQueue) DspGraph() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dspGraph"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSWrappedAudioQueue) SetDspGraph(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setDspGraph:"), objc.String(value))
}
func (t TTSWrappedAudioQueue) GraphParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("graphParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetGraphParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setGraphParameters:"), value)
}
func (t TTSWrappedAudioQueue) GraphProperties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("graphProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetGraphProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setGraphProperties:"), value)
}
func (t TTSWrappedAudioQueue) InflightBuffers() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("inflightBuffers"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetInflightBuffers(value foundation.INSOrderedSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setInflightBuffers:"), value)
}
func (t TTSWrappedAudioQueue) NeedsParameterSync() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("needsParameterSync"))
	return rv
}
func (t TTSWrappedAudioQueue) SetNeedsParameterSync(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setNeedsParameterSync:"), value)
}
func (t TTSWrappedAudioQueue) NeedsPropertySync() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("needsPropertySync"))
	return rv
}
func (t TTSWrappedAudioQueue) SetNeedsPropertySync(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setNeedsPropertySync:"), value)
}
func (t TTSWrappedAudioQueue) OutputFormat() ITTSAudioFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("outputFormat"))
	return TTSAudioFormatFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetOutputFormat(value ITTSAudioFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setOutputFormat:"), value)
}
func (t TTSWrappedAudioQueue) ProcNodeRef() OpaqueATAudioProcessingNodeRef {
	rv := objc.Send[OpaqueATAudioProcessingNodeRef](t.ID, objc.Sel("procNodeRef"))
	return OpaqueATAudioProcessingNodeRef(rv)
}
func (t TTSWrappedAudioQueue) SetProcNodeRef(value OpaqueATAudioProcessingNodeRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setProcNodeRef:"), value)
}
func (t TTSWrappedAudioQueue) QueueFormat() avfaudio.AVAudioFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("queueFormat"))
	return avfaudio.AVAudioFormatFromID(objc.ID(rv))
}
func (t TTSWrappedAudioQueue) SetQueueFormat(value avfaudio.AVAudioFormat) {
	objc.Send[struct{}](t.ID, objc.Sel("setQueueFormat:"), value)
}
func (t TTSWrappedAudioQueue) ShouldRebuildAudioQueue() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldRebuildAudioQueue"))
	return rv
}
func (t TTSWrappedAudioQueue) SetShouldRebuildAudioQueue(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldRebuildAudioQueue:"), value)
}
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
