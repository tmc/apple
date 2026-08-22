// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlannedSegmentConfiguration] class.
var (
	_AVPlannedSegmentConfigurationClass     AVPlannedSegmentConfigurationClass
	_AVPlannedSegmentConfigurationClassOnce sync.Once
)

func getAVPlannedSegmentConfigurationClass() AVPlannedSegmentConfigurationClass {
	_AVPlannedSegmentConfigurationClassOnce.Do(func() {
		_AVPlannedSegmentConfigurationClass = AVPlannedSegmentConfigurationClass{class: objc.GetClass("AVPlannedSegmentConfiguration")}
	})
	return _AVPlannedSegmentConfigurationClass
}

// GetAVPlannedSegmentConfigurationClass returns the class object for AVPlannedSegmentConfiguration.
func GetAVPlannedSegmentConfigurationClass() AVPlannedSegmentConfigurationClass {
	return getAVPlannedSegmentConfigurationClass()
}

type AVPlannedSegmentConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlannedSegmentConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlannedSegmentConfigurationClass) Alloc() AVPlannedSegmentConfiguration {
	rv := objc.Send[AVPlannedSegmentConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// AVPlannedSegmentConfiguration describes the requirements for a planned
// segment in an incremental writing session executed by the
// AVAssetWritingPlanner. Subclasses of this type that are used from Swift
// must fulfill the requirements of a Sendable type.
//
// # Creating a segment configuration
//
//   - [AVPlannedSegmentConfiguration.InitWithDuration]: Creates an instance of AVPlannedSegmentConfiguration specifying the duration of the planned segment.
//
// # Inspecting the configuration
//
//   - [AVPlannedSegmentConfiguration.Duration]: The duration of this planned segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentConfiguration
type AVPlannedSegmentConfiguration struct {
	objectivec.Object
}

// AVPlannedSegmentConfigurationFromID constructs a [AVPlannedSegmentConfiguration] from an objc.ID.
//
// AVPlannedSegmentConfiguration describes the requirements for a planned
// segment in an incremental writing session executed by the
// AVAssetWritingPlanner. Subclasses of this type that are used from Swift
// must fulfill the requirements of a Sendable type.
func AVPlannedSegmentConfigurationFromID(id objc.ID) AVPlannedSegmentConfiguration {
	return AVPlannedSegmentConfiguration{objectivec.Object{ID: id}}
}

// NOTE: AVPlannedSegmentConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlannedSegmentConfiguration] class.
//
// # Creating a segment configuration
//
//   - [IAVPlannedSegmentConfiguration.InitWithDuration]: Creates an instance of AVPlannedSegmentConfiguration specifying the duration of the planned segment.
//
// # Inspecting the configuration
//
//   - [IAVPlannedSegmentConfiguration.Duration]: The duration of this planned segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentConfiguration
type IAVPlannedSegmentConfiguration interface {
	objectivec.IObject

	// Topic: Creating a segment configuration

	// Creates an instance of AVPlannedSegmentConfiguration specifying the duration of the planned segment.
	InitWithDuration(duration coremedia.CMTime) AVPlannedSegmentConfiguration

	// Topic: Inspecting the configuration

	// The duration of this planned segment.
	Duration() coremedia.CMTime
}

// Init initializes the instance.
func (p AVPlannedSegmentConfiguration) Init() AVPlannedSegmentConfiguration {
	rv := objc.Send[AVPlannedSegmentConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlannedSegmentConfiguration) Autorelease() AVPlannedSegmentConfiguration {
	rv := objc.Send[AVPlannedSegmentConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlannedSegmentConfiguration creates a new AVPlannedSegmentConfiguration instance.
func NewAVPlannedSegmentConfiguration() AVPlannedSegmentConfiguration {
	class := getAVPlannedSegmentConfigurationClass()
	rv := objc.Send[AVPlannedSegmentConfiguration](objc.ID(class.class), objc.Sel("new"))
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
func NewPlannedSegmentConfigurationWithDuration(duration coremedia.CMTime) AVPlannedSegmentConfiguration {
	instance := getAVPlannedSegmentConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDuration:"), duration)
	return AVPlannedSegmentConfigurationFromID(rv)
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
func (p AVPlannedSegmentConfiguration) InitWithDuration(duration coremedia.CMTime) AVPlannedSegmentConfiguration {
	rv := objc.Send[AVPlannedSegmentConfiguration](p.ID, objc.Sel("initWithDuration:"), duration)
	return rv
}

// The duration of this planned segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlannedSegmentConfiguration/duration
func (p AVPlannedSegmentConfiguration) Duration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("duration"))
	return coremedia.CMTime(rv)
}
