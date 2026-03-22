// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods and constants that a haptic feedback performer implements.
//
// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackPerformer
type NSHapticFeedbackPerformer interface {
	objectivec.IObject

	// Initiates a specific pattern of haptic feedback to the user.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackPerformer/perform(_:performanceTime:)
	PerformFeedbackPatternPerformanceTime(pattern NSHapticFeedbackPattern, performanceTime NSHapticFeedbackPerformanceTime)
}

// NSHapticFeedbackPerformerObject wraps an existing Objective-C object that conforms to the NSHapticFeedbackPerformer protocol.
type NSHapticFeedbackPerformerObject struct {
	objectivec.Object
}
func (o NSHapticFeedbackPerformerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSHapticFeedbackPerformerObjectFromID constructs a [NSHapticFeedbackPerformerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSHapticFeedbackPerformerObjectFromID(id objc.ID) NSHapticFeedbackPerformerObject {
	return NSHapticFeedbackPerformerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Initiates a specific pattern of haptic feedback to the user.
//
// pattern: A pattern of feedback to be provided to the user. For possible values, see
// [NSHapticFeedbackManager.FeedbackPattern].
// //
// [NSHapticFeedbackManager.FeedbackPattern]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern
//
// performanceTime: The time when the feedback should be provided to the user. For possible
// values, see [NSHapticFeedbackManager.PerformanceTime].
// //
// [NSHapticFeedbackManager.PerformanceTime]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime
//
// # Discussion
// 
// In some cases, the system may override a call to this method. For example,
// a Force Touch trackpad won’t provide haptic feedback if the user isn’t
// touching the trackpad.
//
// See: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackPerformer/perform(_:performanceTime:)
func (o NSHapticFeedbackPerformerObject) PerformFeedbackPatternPerformanceTime(pattern NSHapticFeedbackPattern, performanceTime NSHapticFeedbackPerformanceTime) {
	objc.Send[struct{}](o.ID, objc.Sel("performFeedbackPattern:performanceTime:"), pattern, performanceTime)
	}

