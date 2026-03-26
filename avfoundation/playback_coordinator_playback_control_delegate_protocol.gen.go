// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the method to implement to respond to playback commands from the playback coordinator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate
type AVPlaybackCoordinatorPlaybackControlDelegate interface {
	objectivec.IObject

	// Tells the delegate to match the playback rate to that of the group when the rate is nonzero.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-73p3a
	PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, playCommand IAVDelegatingPlaybackCoordinatorPlayCommand, completionHandler VoidHandler)

	// Tells the delegate to pause playback.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-56t01
	PlaybackCoordinatorDidIssuePauseCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, pauseCommand IAVDelegatingPlaybackCoordinatorPauseCommand, completionHandler VoidHandler)

	// Tells the delegate to seek to a new time.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-4fk8y
	PlaybackCoordinatorDidIssueSeekCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, seekCommand IAVDelegatingPlaybackCoordinatorSeekCommand, completionHandler VoidHandler)

	// Tells the delegate to expect playback soon and to start buffering media data in preparation.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-btle
	PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, bufferingCommand IAVDelegatingPlaybackCoordinatorBufferingCommand, completionHandler VoidHandler)
}

// AVPlaybackCoordinatorPlaybackControlDelegateObject wraps an existing Objective-C object that conforms to the AVPlaybackCoordinatorPlaybackControlDelegate protocol.
type AVPlaybackCoordinatorPlaybackControlDelegateObject struct {
	objectivec.Object
}
func (o AVPlaybackCoordinatorPlaybackControlDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlaybackCoordinatorPlaybackControlDelegateObjectFromID constructs a [AVPlaybackCoordinatorPlaybackControlDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlaybackCoordinatorPlaybackControlDelegateObjectFromID(id objc.ID) AVPlaybackCoordinatorPlaybackControlDelegateObject {
	return AVPlaybackCoordinatorPlaybackControlDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate to match the playback rate to that of the group when the
// rate is nonzero.
//
// coordinator: The playback coordinator that issues the command.
//
// playCommand: The command to execute. Before performing it, verify that its
// [ExpectedCurrentItemIdentifier] property value matches the item that
// you’re currently playing. If the command isn’t valid for the current
// item, ignore it and call the completion handler.
//
// completionHandler: A completion handler that your app must call when it finishes handling the
// command, either successfully or after beginning a suspension if it can’t
// handle the command currently.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-73p3a
func (o AVPlaybackCoordinatorPlaybackControlDelegateObject) PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, playCommand IAVDelegatingPlaybackCoordinatorPlayCommand, completionHandler VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("playbackCoordinator:didIssuePlayCommand:completionHandler:"), coordinator, playCommand, completionHandler)
	}
// Tells the delegate to pause playback.
//
// coordinator: The playback coordinator that issues the command.
//
// pauseCommand: The command to execute. Before performing it, verify that its
// [ExpectedCurrentItemIdentifier] property value matches the item that
// you’re currently playing. If the command isn’t valid for the current
// item, ignore it and call the completion handler.
//
// completionHandler: A completion handler that your app must call when it finishes handling the
// command, either successfully or after beginning a suspension if it can’t
// handle the command currently.
// 
// If the value of the command’s [ShouldBufferInAnticipationOfPlayback]
// property is [true], call the completion handler only after the player is
// ready for playback.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-56t01
func (o AVPlaybackCoordinatorPlaybackControlDelegateObject) PlaybackCoordinatorDidIssuePauseCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, pauseCommand IAVDelegatingPlaybackCoordinatorPauseCommand, completionHandler VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("playbackCoordinator:didIssuePauseCommand:completionHandler:"), coordinator, pauseCommand, completionHandler)
	}
// Tells the delegate to seek to a new time.
//
// coordinator: The playback coordinator that issues the command.
//
// seekCommand: The command to execute. Before performing it, verify that its
// [ExpectedCurrentItemIdentifier] property value matches the item that
// you’re currently playing. If the command isn’t valid for the current
// item, ignore it and call the completion handler.
//
// completionHandler: A completion handler that your app must call when it finishes handling the
// command, either successfully or after beginning a suspension if it can’t
// handle the command currently.
// 
// If the value of the command’s [ShouldBufferInAnticipationOfPlayback]
// property is [true], call the completion handler only after the player is
// ready for playback.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The coordinator issues this command to perform a seek in the item timeline,
// potentially pausing playback in the process.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-4fk8y
func (o AVPlaybackCoordinatorPlaybackControlDelegateObject) PlaybackCoordinatorDidIssueSeekCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, seekCommand IAVDelegatingPlaybackCoordinatorSeekCommand, completionHandler VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("playbackCoordinator:didIssueSeekCommand:completionHandler:"), coordinator, seekCommand, completionHandler)
	}
// Tells the delegate to expect playback soon and to start buffering media
// data in preparation.
//
// coordinator: The playback coordinator that issues the command.
//
// bufferingCommand: The command to execute. Before performing it, verify that its
// [ExpectedCurrentItemIdentifier] property value matches the item that
// you’re currently playing. If the command isn’t valid for the current
// item, ignore it and call the completion handler.
//
// completionHandler: A completion handler that your app must call when it finishes handling the
// command, either successfully or after beginning a suspension if it can’t
// handle the command currently.
// 
// If the value of the command’s [ShouldBufferInAnticipationOfPlayback]
// property is [true], call the completion handler only after the player is
// ready for playback.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The coordinator issues this command when playback is currently in a paused
// state and the coordinator is expecting playback to start soon. It provides
// an appropriate opportunity to update your player UI to indicate that
// playback is in a waiting state.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate/playbackCoordinator(_:didIssue:completionHandler:)-btle
func (o AVPlaybackCoordinatorPlaybackControlDelegateObject) PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, bufferingCommand IAVDelegatingPlaybackCoordinatorBufferingCommand, completionHandler VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("playbackCoordinator:didIssueBufferingCommand:completionHandler:"), coordinator, bufferingCommand, completionHandler)
	}

