// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCMetricsManager] class.
var (
	_AVVCMetricsManagerClass     AVVCMetricsManagerClass
	_AVVCMetricsManagerClassOnce sync.Once
)

func getAVVCMetricsManagerClass() AVVCMetricsManagerClass {
	_AVVCMetricsManagerClassOnce.Do(func() {
		_AVVCMetricsManagerClass = AVVCMetricsManagerClass{class: objc.GetClass("AVVCMetricsManager")}
	})
	return _AVVCMetricsManagerClass
}

// GetAVVCMetricsManagerClass returns the class object for AVVCMetricsManager.
func GetAVVCMetricsManagerClass() AVVCMetricsManagerClass {
	return getAVVCMetricsManagerClass()
}

type AVVCMetricsManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCMetricsManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCMetricsManagerClass) Alloc() AVVCMetricsManager {
	rv := objc.Send[AVVCMetricsManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCMetricsManager._disposeADAM]
//   - [AVVCMetricsManager.AdamAnalyzeBufferNumFramesTimeStampShouldAnalyze]
//   - [AVVCMetricsManager.AudioIssueDetectorAnalyzeBufferNumFramesTimeStampShouldAnalyze]
//   - [AVVCMetricsManager.AvvcProfilingInfoDictionary]
//   - [AVVCMetricsManager.SetAvvcProfilingInfoDictionary]
//   - [AVVCMetricsManager.CallToStartRecordHostTime]
//   - [AVVCMetricsManager.SetCallToStartRecordHostTime]
//   - [AVVCMetricsManager.CheckAndUpdateReporterID]
//   - [AVVCMetricsManager.DisposeADAM]
//   - [AVVCMetricsManager.GetStringDate]
//   - [AVVCMetricsManager.LogABCMetricCategoryTypeReporterID]
//   - [AVVCMetricsManager.LogProfileMetrics]
//   - [AVVCMetricsManager.LogRecordAudioFormatReporterID]
//   - [AVVCMetricsManager.LogRecordRouteAndPlaybackRouteReporterID]
//   - [AVVCMetricsManager.LogSessionMetricValueCategoryTypeContextReporterID]
//   - [AVVCMetricsManager.LogSessionMetricValueCategoryTypeReporterID]
//   - [AVVCMetricsManager.LogSingleMetricValueCategoryTypeReporterID]
//   - [AVVCMetricsManager.MeasureElapseTimeForMetricBlock]
//   - [AVVCMetricsManager.PublicMetrics]
//   - [AVVCMetricsManager.SetPublicMetrics]
//   - [AVVCMetricsManager.ReporterID]
//   - [AVVCMetricsManager.ResetAudioIssueDetector]
//   - [AVVCMetricsManager.ResetProfileMetrics]
//   - [AVVCMetricsManager.RetrieveMetrics]
//   - [AVVCMetricsManager.RetrieveProfileMetrics]
//   - [AVVCMetricsManager.SetADAMFormatNumFrames]
//   - [AVVCMetricsManager.SetAudioIssueDetectorFormatNumFrames]
//   - [AVVCMetricsManager.UpdateWithReporterID]
//   - [AVVCMetricsManager.VoiceTriggerStartHostTime]
//   - [AVVCMetricsManager.SetVoiceTriggerStartHostTime]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager
type AVVCMetricsManager struct {
	objectivec.Object
}

// AVVCMetricsManagerFromID constructs a [AVVCMetricsManager] from an objc.ID.
func AVVCMetricsManagerFromID(id objc.ID) AVVCMetricsManager {
	return AVVCMetricsManager{objectivec.Object{ID: id}}
}
// Ensure AVVCMetricsManager implements IAVVCMetricsManager.
var _ IAVVCMetricsManager = AVVCMetricsManager{}

// An interface definition for the [AVVCMetricsManager] class.
//
// # Methods
//
//   - [IAVVCMetricsManager._disposeADAM]
//   - [IAVVCMetricsManager.AdamAnalyzeBufferNumFramesTimeStampShouldAnalyze]
//   - [IAVVCMetricsManager.AudioIssueDetectorAnalyzeBufferNumFramesTimeStampShouldAnalyze]
//   - [IAVVCMetricsManager.AvvcProfilingInfoDictionary]
//   - [IAVVCMetricsManager.SetAvvcProfilingInfoDictionary]
//   - [IAVVCMetricsManager.CallToStartRecordHostTime]
//   - [IAVVCMetricsManager.SetCallToStartRecordHostTime]
//   - [IAVVCMetricsManager.CheckAndUpdateReporterID]
//   - [IAVVCMetricsManager.DisposeADAM]
//   - [IAVVCMetricsManager.GetStringDate]
//   - [IAVVCMetricsManager.LogABCMetricCategoryTypeReporterID]
//   - [IAVVCMetricsManager.LogProfileMetrics]
//   - [IAVVCMetricsManager.LogRecordAudioFormatReporterID]
//   - [IAVVCMetricsManager.LogRecordRouteAndPlaybackRouteReporterID]
//   - [IAVVCMetricsManager.LogSessionMetricValueCategoryTypeContextReporterID]
//   - [IAVVCMetricsManager.LogSessionMetricValueCategoryTypeReporterID]
//   - [IAVVCMetricsManager.LogSingleMetricValueCategoryTypeReporterID]
//   - [IAVVCMetricsManager.MeasureElapseTimeForMetricBlock]
//   - [IAVVCMetricsManager.PublicMetrics]
//   - [IAVVCMetricsManager.SetPublicMetrics]
//   - [IAVVCMetricsManager.ReporterID]
//   - [IAVVCMetricsManager.ResetAudioIssueDetector]
//   - [IAVVCMetricsManager.ResetProfileMetrics]
//   - [IAVVCMetricsManager.RetrieveMetrics]
//   - [IAVVCMetricsManager.RetrieveProfileMetrics]
//   - [IAVVCMetricsManager.SetADAMFormatNumFrames]
//   - [IAVVCMetricsManager.SetAudioIssueDetectorFormatNumFrames]
//   - [IAVVCMetricsManager.UpdateWithReporterID]
//   - [IAVVCMetricsManager.VoiceTriggerStartHostTime]
//   - [IAVVCMetricsManager.SetVoiceTriggerStartHostTime]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager
type IAVVCMetricsManager interface {
	objectivec.IObject

	// Topic: Methods

	_disposeADAM() int
	AdamAnalyzeBufferNumFramesTimeStampShouldAnalyze(buffer objectivec.IObject, frames uint32, stamp unsafe.Pointer, analyze bool) int
	AudioIssueDetectorAnalyzeBufferNumFramesTimeStampShouldAnalyze(buffer objectivec.IObject, frames uint32, stamp unsafe.Pointer, analyze bool) int
	AvvcProfilingInfoDictionary() foundation.INSDictionary
	SetAvvcProfilingInfoDictionary(value foundation.INSDictionary)
	CallToStartRecordHostTime() uint64
	SetCallToStartRecordHostTime(value uint64)
	CheckAndUpdateReporterID(id int64)
	DisposeADAM() int
	GetStringDate(date objectivec.IObject) objectivec.IObject
	LogABCMetricCategoryTypeReporterID(aBCMetric objectivec.IObject, category uint32, type_ uint16, id int64)
	LogProfileMetrics(metrics objectivec.IObject)
	LogRecordAudioFormatReporterID(format objectivec.IObject, id int64)
	LogRecordRouteAndPlaybackRouteReporterID(route objectivec.IObject, route2 objectivec.IObject, id int64)
	LogSessionMetricValueCategoryTypeContextReporterID(metric objectivec.IObject, value objectivec.IObject, category uint32, type_ uint16, context objectivec.IObject, id int64)
	LogSessionMetricValueCategoryTypeReporterID(metric objectivec.IObject, value objectivec.IObject, category uint32, type_ uint16, id int64)
	LogSingleMetricValueCategoryTypeReporterID(metric objectivec.IObject, value objectivec.IObject, category uint32, type_ uint16, id int64)
	MeasureElapseTimeForMetricBlock(metric objectivec.IObject, block VoidHandler) bool
	PublicMetrics() foundation.INSDictionary
	SetPublicMetrics(value foundation.INSDictionary)
	ReporterID() int64
	ResetAudioIssueDetector() int
	ResetProfileMetrics()
	RetrieveMetrics() objectivec.IObject
	RetrieveProfileMetrics() objectivec.IObject
	SetADAMFormatNumFrames(aDAMFormat objectivec.IObject, frames uint32) int
	SetAudioIssueDetectorFormatNumFrames(format objectivec.IObject, frames uint32) int
	UpdateWithReporterID(id int64)
	VoiceTriggerStartHostTime() uint64
	SetVoiceTriggerStartHostTime(value uint64)
}

// Init initializes the instance.
func (v AVVCMetricsManager) Init() AVVCMetricsManager {
	rv := objc.Send[AVVCMetricsManager](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCMetricsManager) Autorelease() AVVCMetricsManager {
	rv := objc.Send[AVVCMetricsManager](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCMetricsManager creates a new AVVCMetricsManager instance.
func NewAVVCMetricsManager() AVVCMetricsManager {
	class := getAVVCMetricsManagerClass()
	rv := objc.Send[AVVCMetricsManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/_disposeADAM
func (v AVVCMetricsManager) _disposeADAM() int {
	rv := objc.Send[int](v.ID, objc.Sel("_disposeADAM"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/adamAnalyzeBuffer:numFrames:timeStamp:shouldAnalyze:
func (v AVVCMetricsManager) AdamAnalyzeBufferNumFramesTimeStampShouldAnalyze(buffer objectivec.IObject, frames uint32, stamp unsafe.Pointer, analyze bool) int {
	rv := objc.Send[int](v.ID, objc.Sel("adamAnalyzeBuffer:numFrames:timeStamp:shouldAnalyze:"), buffer, frames, stamp, analyze)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/audioIssueDetectorAnalyzeBuffer:numFrames:timeStamp:shouldAnalyze:
func (v AVVCMetricsManager) AudioIssueDetectorAnalyzeBufferNumFramesTimeStampShouldAnalyze(buffer objectivec.IObject, frames uint32, stamp unsafe.Pointer, analyze bool) int {
	rv := objc.Send[int](v.ID, objc.Sel("audioIssueDetectorAnalyzeBuffer:numFrames:timeStamp:shouldAnalyze:"), buffer, frames, stamp, analyze)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/checkAndUpdateReporterID:
func (v AVVCMetricsManager) CheckAndUpdateReporterID(id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("checkAndUpdateReporterID:"), id)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/disposeADAM
func (v AVVCMetricsManager) DisposeADAM() int {
	rv := objc.Send[int](v.ID, objc.Sel("disposeADAM"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/getStringDate:
func (v AVVCMetricsManager) GetStringDate(date objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getStringDate:"), date)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/logABCMetric:category:type:reporterID:
func (v AVVCMetricsManager) LogABCMetricCategoryTypeReporterID(aBCMetric objectivec.IObject, category uint32, type_ uint16, id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("logABCMetric:category:type:reporterID:"), aBCMetric, category, type_, id)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/logProfileMetrics:
func (v AVVCMetricsManager) LogProfileMetrics(metrics objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("logProfileMetrics:"), metrics)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/logRecordAudioFormat:reporterID:
func (v AVVCMetricsManager) LogRecordAudioFormatReporterID(format objectivec.IObject, id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("logRecordAudioFormat:reporterID:"), format, id)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/logRecordRoute:andPlaybackRoute:reporterID:
func (v AVVCMetricsManager) LogRecordRouteAndPlaybackRouteReporterID(route objectivec.IObject, route2 objectivec.IObject, id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("logRecordRoute:andPlaybackRoute:reporterID:"), route, route2, id)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/logSessionMetric:value:category:type:context:reporterID:
func (v AVVCMetricsManager) LogSessionMetricValueCategoryTypeContextReporterID(metric objectivec.IObject, value objectivec.IObject, category uint32, type_ uint16, context objectivec.IObject, id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("logSessionMetric:value:category:type:context:reporterID:"), metric, value, category, type_, context, id)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/logSessionMetric:value:category:type:reporterID:
func (v AVVCMetricsManager) LogSessionMetricValueCategoryTypeReporterID(metric objectivec.IObject, value objectivec.IObject, category uint32, type_ uint16, id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("logSessionMetric:value:category:type:reporterID:"), metric, value, category, type_, id)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/logSingleMetric:value:category:type:reporterID:
func (v AVVCMetricsManager) LogSingleMetricValueCategoryTypeReporterID(metric objectivec.IObject, value objectivec.IObject, category uint32, type_ uint16, id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("logSingleMetric:value:category:type:reporterID:"), metric, value, category, type_, id)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/measureElapseTimeForMetric:block:
func (v AVVCMetricsManager) MeasureElapseTimeForMetricBlock(metric objectivec.IObject, block VoidHandler) bool {
_block1, _ := NewVoidBlock(block)
	rv := objc.Send[bool](v.ID, objc.Sel("measureElapseTimeForMetric:block:"), metric, _block1)
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/reporterID
func (v AVVCMetricsManager) ReporterID() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("reporterID"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/resetAudioIssueDetector
func (v AVVCMetricsManager) ResetAudioIssueDetector() int {
	rv := objc.Send[int](v.ID, objc.Sel("resetAudioIssueDetector"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/resetProfileMetrics
func (v AVVCMetricsManager) ResetProfileMetrics() {
	objc.Send[objc.ID](v.ID, objc.Sel("resetProfileMetrics"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/retrieveMetrics
func (v AVVCMetricsManager) RetrieveMetrics() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("retrieveMetrics"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/retrieveProfileMetrics
func (v AVVCMetricsManager) RetrieveProfileMetrics() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("retrieveProfileMetrics"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/setADAMFormat:numFrames:
func (v AVVCMetricsManager) SetADAMFormatNumFrames(aDAMFormat objectivec.IObject, frames uint32) int {
	rv := objc.Send[int](v.ID, objc.Sel("setADAMFormat:numFrames:"), aDAMFormat, frames)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/setAudioIssueDetectorFormat:numFrames:
func (v AVVCMetricsManager) SetAudioIssueDetectorFormatNumFrames(format objectivec.IObject, frames uint32) int {
	rv := objc.Send[int](v.ID, objc.Sel("setAudioIssueDetectorFormat:numFrames:"), format, frames)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/updateWithReporterID:
func (v AVVCMetricsManager) UpdateWithReporterID(id int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("updateWithReporterID:"), id)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/createSharedManager
func (_AVVCMetricsManagerClass AVVCMetricsManagerClass) CreateSharedManager() {
	objc.Send[objc.ID](objc.ID(_AVVCMetricsManagerClass.class), objc.Sel("createSharedManager"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/destroySharedManager
func (_AVVCMetricsManagerClass AVVCMetricsManagerClass) DestroySharedManager() {
	objc.Send[objc.ID](objc.ID(_AVVCMetricsManagerClass.class), objc.Sel("destroySharedManager"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/getLock
func (_AVVCMetricsManagerClass AVVCMetricsManagerClass) GetLock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_AVVCMetricsManagerClass.class), objc.Sel("getLock"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/sharedManager
func (_AVVCMetricsManagerClass AVVCMetricsManagerClass) SharedManager() AVVCMetricsManager {
	rv := objc.Send[objc.ID](objc.ID(_AVVCMetricsManagerClass.class), objc.Sel("sharedManager"))
	return AVVCMetricsManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/avvcProfilingInfoDictionary
func (v AVVCMetricsManager) AvvcProfilingInfoDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("avvcProfilingInfoDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v AVVCMetricsManager) SetAvvcProfilingInfoDictionary(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setAvvcProfilingInfoDictionary:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/callToStartRecordHostTime
func (v AVVCMetricsManager) CallToStartRecordHostTime() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("callToStartRecordHostTime"))
	return rv
}
func (v AVVCMetricsManager) SetCallToStartRecordHostTime(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setCallToStartRecordHostTime:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/publicMetrics
func (v AVVCMetricsManager) PublicMetrics() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("publicMetrics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v AVVCMetricsManager) SetPublicMetrics(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setPublicMetrics:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCMetricsManager/voiceTriggerStartHostTime
func (v AVVCMetricsManager) VoiceTriggerStartHostTime() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("voiceTriggerStartHostTime"))
	return rv
}
func (v AVVCMetricsManager) SetVoiceTriggerStartHostTime(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setVoiceTriggerStartHostTime:"), value)
}

// MeasureElapseTimeForMetricBlockSync is a synchronous wrapper around [AVVCMetricsManager.MeasureElapseTimeForMetricBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVCMetricsManager) MeasureElapseTimeForMetricBlockSync(ctx context.Context, metric objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.MeasureElapseTimeForMetricBlock(metric, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

