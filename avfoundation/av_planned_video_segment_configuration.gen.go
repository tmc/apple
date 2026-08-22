// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVPlannedVideoSegmentConfiguration] class.
var (
	_AVPlannedVideoSegmentConfigurationClass     AVPlannedVideoSegmentConfigurationClass
	_AVPlannedVideoSegmentConfigurationClassOnce sync.Once
)

func getAVPlannedVideoSegmentConfigurationClass() AVPlannedVideoSegmentConfigurationClass {
	_AVPlannedVideoSegmentConfigurationClassOnce.Do(func() {
		_AVPlannedVideoSegmentConfigurationClass = AVPlannedVideoSegmentConfigurationClass{class: objc.GetClass("AVPlannedVideoSegmentConfiguration")}
	})
	return _AVPlannedVideoSegmentConfigurationClass
}

// GetAVPlannedVideoSegmentConfigurationClass returns the class object for AVPlannedVideoSegmentConfiguration.
func GetAVPlannedVideoSegmentConfigurationClass() AVPlannedVideoSegmentConfigurationClass {
	return getAVPlannedVideoSegmentConfigurationClass()
}

type AVPlannedVideoSegmentConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlannedVideoSegmentConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlannedVideoSegmentConfigurationClass) Alloc() AVPlannedVideoSegmentConfiguration {
	rv := objc.Send[AVPlannedVideoSegmentConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVPlannedVideoSegmentConfiguration describes the requirements for a planned
// video segment in an incremental writing session executed by the
// AVAssetWritingPlanner.
//
// # Overview
//
// Use this class instead of the base class AVPlannedSegmentConfiguration if
// you are setting up AVAssetWriterInput to do video compression.
// AVAssetWritingPlanner will provide required video compression properties in
// its AVPlannedSegmentWritingRequest that are needed to prevent visual
// artifacts on segment boundaries.
//
// # Creating a video segment configuration
//
//   - [AVPlannedVideoSegmentConfiguration.InitWithNumberOfFramesDuration]: Creates an instance of AVPlannedVideoSegmentConfiguration specifying the number of frames in and total duration of the segment.
//
// # Inspecting the configuration
//
//   - [AVPlannedVideoSegmentConfiguration.FrameCount]: The number of frames in this planned video segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentConfiguration
type AVPlannedVideoSegmentConfiguration struct {
	AVPlannedSegmentConfiguration
}

// AVPlannedVideoSegmentConfigurationFromID constructs a [AVPlannedVideoSegmentConfiguration] from an objc.ID.
//
// AVPlannedVideoSegmentConfiguration describes the requirements for a planned
// video segment in an incremental writing session executed by the
// AVAssetWritingPlanner.
func AVPlannedVideoSegmentConfigurationFromID(id objc.ID) AVPlannedVideoSegmentConfiguration {
	return AVPlannedVideoSegmentConfiguration{AVPlannedSegmentConfiguration: AVPlannedSegmentConfigurationFromID(id)}
}

// NOTE: AVPlannedVideoSegmentConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlannedVideoSegmentConfiguration] class.
//
// # Creating a video segment configuration
//
//   - [IAVPlannedVideoSegmentConfiguration.InitWithNumberOfFramesDuration]: Creates an instance of AVPlannedVideoSegmentConfiguration specifying the number of frames in and total duration of the segment.
//
// # Inspecting the configuration
//
//   - [IAVPlannedVideoSegmentConfiguration.FrameCount]: The number of frames in this planned video segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentConfiguration
type IAVPlannedVideoSegmentConfiguration interface {
	IAVPlannedSegmentConfiguration

	// Topic: Creating a video segment configuration

	// Creates an instance of AVPlannedVideoSegmentConfiguration specifying the number of frames in and total duration of the segment.
	InitWithNumberOfFramesDuration(frameCount int, duration coremedia.CMTime) AVPlannedVideoSegmentConfiguration

	// Topic: Inspecting the configuration

	// The number of frames in this planned video segment.
	FrameCount() int
}

// Init initializes the instance.
func (p AVPlannedVideoSegmentConfiguration) Init() AVPlannedVideoSegmentConfiguration {
	rv := objc.Send[AVPlannedVideoSegmentConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlannedVideoSegmentConfiguration) Autorelease() AVPlannedVideoSegmentConfiguration {
	rv := objc.Send[AVPlannedVideoSegmentConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlannedVideoSegmentConfiguration creates a new AVPlannedVideoSegmentConfiguration instance.
func NewAVPlannedVideoSegmentConfiguration() AVPlannedVideoSegmentConfiguration {
	class := getAVPlannedVideoSegmentConfigurationClass()
	rv := objc.Send[AVPlannedVideoSegmentConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an instance of AVPlannedSegmentConfiguration specifying the
// duration of the planned segment.
//
// duration: The total duration of this planned segment. If an empty edit is included,
// this duration may be larger than the sum of the durations of the samples in
// this planned segment.
//
// # Return Value
//
// An instance of AVPlannedSegmentConfiguration, or nil if initialization
// fails.
//
// # Discussion
//
// The duration parameter must be numeric and greater than 0. Otherwise, the
// initializer throws NSInvalidArgumentException.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentConfiguration/init(duration:)
func NewPlannedVideoSegmentConfigurationWithDuration(duration coremedia.CMTime) AVPlannedVideoSegmentConfiguration {
	instance := getAVPlannedVideoSegmentConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDuration:"), duration)
	return AVPlannedVideoSegmentConfigurationFromID(rv)
}

// Creates an instance of AVPlannedVideoSegmentConfiguration specifying the
// number of frames in and total duration of the segment.
//
// frameCount: The number of frames in this planned video segment.
//
// duration: The duration of this planned video segment.
//
// # Return Value
//
// An instance of AVPlannedVideoSegmentConfiguration.
//
// # Discussion
//
// For best results, frameCount and duration should be greater or equal to the
// minimumFrameCount and minimumDuration of
// AVPlannedVideoSegmentBoundaryGuidelines respectively. This initializer
// throws NSInvalidArgumentException if frameCount is less than or equal to 0,
// or duration is not numeric, or duration is less than or equal to 0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentConfiguration/init(numberOfFrames:duration:)
func NewPlannedVideoSegmentConfigurationWithNumberOfFramesDuration(frameCount int, duration coremedia.CMTime) AVPlannedVideoSegmentConfiguration {
	instance := getAVPlannedVideoSegmentConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNumberOfFrames:duration:"), frameCount, duration)
	return AVPlannedVideoSegmentConfigurationFromID(rv)
}

// Creates an instance of AVPlannedVideoSegmentConfiguration specifying the
// number of frames in and total duration of the segment.
//
// frameCount: The number of frames in this planned video segment.
//
// duration: The duration of this planned video segment.
//
// # Return Value
//
// An instance of AVPlannedVideoSegmentConfiguration.
//
// # Discussion
//
// For best results, frameCount and duration should be greater or equal to the
// minimumFrameCount and minimumDuration of
// AVPlannedVideoSegmentBoundaryGuidelines respectively. This initializer
// throws NSInvalidArgumentException if frameCount is less than or equal to 0,
// or duration is not numeric, or duration is less than or equal to 0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentConfiguration/init(numberOfFrames:duration:)
func (p AVPlannedVideoSegmentConfiguration) InitWithNumberOfFramesDuration(frameCount int, duration coremedia.CMTime) AVPlannedVideoSegmentConfiguration {
	rv := objc.Send[AVPlannedVideoSegmentConfiguration](p.ID, objc.Sel("initWithNumberOfFrames:duration:"), frameCount, duration)
	return rv
}

// The number of frames in this planned video segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedVideoSegmentConfiguration/frameCount
func (p AVPlannedVideoSegmentConfiguration) FrameCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("frameCount"))
	return rv
}
