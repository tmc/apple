// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioPlayerNode] class.
var (
	_AVAudioPlayerNodeClass     AVAudioPlayerNodeClass
	_AVAudioPlayerNodeClassOnce sync.Once
)

func getAVAudioPlayerNodeClass() AVAudioPlayerNodeClass {
	_AVAudioPlayerNodeClassOnce.Do(func() {
		_AVAudioPlayerNodeClass = AVAudioPlayerNodeClass{class: objc.GetClass("AVAudioPlayerNode")}
	})
	return _AVAudioPlayerNodeClass
}

// GetAVAudioPlayerNodeClass returns the class object for AVAudioPlayerNode.
func GetAVAudioPlayerNodeClass() AVAudioPlayerNodeClass {
	return getAVAudioPlayerNodeClass()
}

type AVAudioPlayerNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioPlayerNodeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioPlayerNodeClass) Alloc() AVAudioPlayerNode {
	rv := objc.Send[AVAudioPlayerNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object for scheduling the playback of buffers or segments of audio
// files.
//
// # Overview
// 
// This audio node supports scheduling the playback of [AVAudioPCMBuffer]
// instances, or segments of audio files that you open through [AVAudioFile].
// You can schedule buffers and segments to play at specific points in time or
// to play immediately following preceding segments.
// 
// Generally, you want to configure the node’s output format with the same
// number of channels as in the files and buffers. Otherwise, the node drops
// or adds channels as necessary. It’s usually preferable to use an
// [AVAudioMixerNode] for this configuration.
// 
// Similarly, when playing file segments, the node makes sample rate
// conversions, if necessary. It’s preferable to configure the node’s
// output sample rate to match that of the files, and to use a mixer to
// perform the rate conversion.
// 
// When playing buffers, there’s an implicit assumption that the buffers are
// at the same sample rate as the node’s output format.
// 
// The [AVAudioPlayerNode.Stop] method unschedules all previously scheduled buffers and file
// segments, and returns the player timeline to sample time `0`.
// 
// # Player Timeline
// 
// The usual [AVAudioNode] sample times, which [AVAudioPlayerNode.LastRenderTime] observes, have
// an arbitrary zero point. The [AVAudioPlayerNode] class superimposes a
// second player timeline on top of this to reflect when the player starts and
// intervals when it pauses. The methods [AVAudioPlayerNode.NodeTimeForPlayerTime] and
// [AVAudioPlayerNode.PlayerTimeForNodeTime] convert between the two.
// 
// # Scheduling Playback Time
// 
// The [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionHandler],
// [AVAudioPlayerNode.ScheduleFileAtTimeCompletionHandler], and
// [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler] methods
// take an [AVAudioTime] `when` parameter, and you interpret it as follows:
// 
// - If the `when` parameter is `nil`: - If there are previous commands, the
// new one plays immediately following the last one. - Otherwise, if the node
// is in a playing state, the event plays in the very near future. -
// Otherwise, the command plays at sample time `0`. - If the `when` parameter
// is a sample time, the parameter interprets it as such. - If the `when`
// parameter is a host time, the system ignores it unless the sample time is
// invalid when the engine is rendering to an audio device.
// 
// The scheduling methods fail if:
// 
// - A buffer’s channel count doesn’t match that of the node’s output
// format. - The system can’t access a file. - An [AVAudioTime] doesn’t
// specify a valid sample time or a host time. - A segment’s start frame or
// frame count is a negative value.
// 
// # Handling Buffer or File Completion
// 
// The buffer of file completion handlers are a means to schedule more data if
// available on the player node. For more information on the different
// completion callback types, see [AVAudioPlayerNodeCompletionCallbackType].
// 
// # Rendering Offline
// 
// When you use a player node with the engine operating in manual rendering
// mode, you use the buffer or file completion handlers — [AVAudioPlayerNode.LastRenderTime],
// [AVAudioPlayerNode.Latency], and [AVAudioPlayerNode.OutputPresentationLatency] — to track how much data the
// player rendered and how much remains to render.
//
// [AVAudioPlayerNodeCompletionCallbackType]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNodeCompletionCallbackType
//
// # Scheduling Playback
//
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionHandler]: Schedules the playing of an entire audio file.
//   - [AVAudioPlayerNode.ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler]: Schedules the playing of an entire audio file with a callback option you specify.
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler]: Schedules the playing of an audio file segment.
//   - [AVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler]: Schedules the playing of an audio file segment with a callback option you specify.
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionHandler]: Schedules the playing samples from an audio buffer at the time and playback options you specify.
//   - [AVAudioPlayerNode.ScheduleBufferCompletionHandler]: Schedules the playing samples from an audio buffer.
//   - [AVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler]: Schedules the playing samples from an audio buffer with the playback options you specify.
//   - [AVAudioPlayerNode.ScheduleBufferCompletionCallbackTypeCompletionHandler]: Schedules the playing samples from an audio buffer with the callback option you specify.
//
// # Converting Node and Player Times
//
//   - [AVAudioPlayerNode.NodeTimeForPlayerTime]: Converts from player time to node time.
//   - [AVAudioPlayerNode.PlayerTimeForNodeTime]: Converts from node time to player time.
//
// # Controlling Playback
//
//   - [AVAudioPlayerNode.PrepareWithFrameCount]: Prepares the file regions or buffers you schedule for playback.
//   - [AVAudioPlayerNode.Play]: Starts or resumes playback immediately.
//   - [AVAudioPlayerNode.PlayAtTime]: Starts or resumes playback at a time you specify.
//   - [AVAudioPlayerNode.Playing]: A Boolean value that indicates whether the player is playing.
//   - [AVAudioPlayerNode.Pause]: Pauses the node’s playback.
//   - [AVAudioPlayerNode.Stop]: Clears all of the node’s events you schedule and stops playback.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode
type AVAudioPlayerNode struct {
	AVAudioNode
}

// AVAudioPlayerNodeFromID constructs a [AVAudioPlayerNode] from an objc.ID.
//
// An object for scheduling the playback of buffers or segments of audio
// files.
func AVAudioPlayerNodeFromID(id objc.ID) AVAudioPlayerNode {
	return AVAudioPlayerNode{AVAudioNode: AVAudioNodeFromID(id)}
}
// NOTE: AVAudioPlayerNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioPlayerNode] class.
//
// # Scheduling Playback
//
//   - [IAVAudioPlayerNode.ScheduleFileAtTimeCompletionHandler]: Schedules the playing of an entire audio file.
//   - [IAVAudioPlayerNode.ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler]: Schedules the playing of an entire audio file with a callback option you specify.
//   - [IAVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler]: Schedules the playing of an audio file segment.
//   - [IAVAudioPlayerNode.ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler]: Schedules the playing of an audio file segment with a callback option you specify.
//   - [IAVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionHandler]: Schedules the playing samples from an audio buffer at the time and playback options you specify.
//   - [IAVAudioPlayerNode.ScheduleBufferCompletionHandler]: Schedules the playing samples from an audio buffer.
//   - [IAVAudioPlayerNode.ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler]: Schedules the playing samples from an audio buffer with the playback options you specify.
//   - [IAVAudioPlayerNode.ScheduleBufferCompletionCallbackTypeCompletionHandler]: Schedules the playing samples from an audio buffer with the callback option you specify.
//
// # Converting Node and Player Times
//
//   - [IAVAudioPlayerNode.NodeTimeForPlayerTime]: Converts from player time to node time.
//   - [IAVAudioPlayerNode.PlayerTimeForNodeTime]: Converts from node time to player time.
//
// # Controlling Playback
//
//   - [IAVAudioPlayerNode.PrepareWithFrameCount]: Prepares the file regions or buffers you schedule for playback.
//   - [IAVAudioPlayerNode.Play]: Starts or resumes playback immediately.
//   - [IAVAudioPlayerNode.PlayAtTime]: Starts or resumes playback at a time you specify.
//   - [IAVAudioPlayerNode.Playing]: A Boolean value that indicates whether the player is playing.
//   - [IAVAudioPlayerNode.Pause]: Pauses the node’s playback.
//   - [IAVAudioPlayerNode.Stop]: Clears all of the node’s events you schedule and stops playback.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode
type IAVAudioPlayerNode interface {
	IAVAudioNode
	AVAudio3DMixing
	AVAudioMixing
	AVAudioStereoMixing

	// Topic: Scheduling Playback

	// Schedules the playing of an entire audio file.
	ScheduleFileAtTimeCompletionHandler(file IAVAudioFile, when IAVAudioTime, completionHandler ErrorHandler)
	// Schedules the playing of an entire audio file with a callback option you specify.
	ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler(file IAVAudioFile, when IAVAudioTime, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler)
	// Schedules the playing of an audio file segment.
	ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler(file IAVAudioFile, startFrame AVAudioFramePosition, numberFrames AVAudioFrameCount, when IAVAudioTime, completionHandler ErrorHandler)
	// Schedules the playing of an audio file segment with a callback option you specify.
	ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler(file IAVAudioFile, startFrame AVAudioFramePosition, numberFrames AVAudioFrameCount, when IAVAudioTime, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler)
	// Schedules the playing samples from an audio buffer at the time and playback options you specify.
	ScheduleBufferAtTimeOptionsCompletionHandler(buffer IAVAudioPCMBuffer, when IAVAudioTime, options AVAudioPlayerNodeBufferOptions, completionHandler ErrorHandler)
	// Schedules the playing samples from an audio buffer.
	ScheduleBufferCompletionHandler(buffer IAVAudioPCMBuffer, completionHandler ErrorHandler)
	// Schedules the playing samples from an audio buffer with the playback options you specify.
	ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler(buffer IAVAudioPCMBuffer, when IAVAudioTime, options AVAudioPlayerNodeBufferOptions, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler)
	// Schedules the playing samples from an audio buffer with the callback option you specify.
	ScheduleBufferCompletionCallbackTypeCompletionHandler(buffer IAVAudioPCMBuffer, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler)

	// Topic: Converting Node and Player Times

	// Converts from player time to node time.
	NodeTimeForPlayerTime(playerTime IAVAudioTime) IAVAudioTime
	// Converts from node time to player time.
	PlayerTimeForNodeTime(nodeTime IAVAudioTime) IAVAudioTime

	// Topic: Controlling Playback

	// Prepares the file regions or buffers you schedule for playback.
	PrepareWithFrameCount(frameCount AVAudioFrameCount)
	// Starts or resumes playback immediately.
	Play()
	// Starts or resumes playback at a time you specify.
	PlayAtTime(when IAVAudioTime)
	// A Boolean value that indicates whether the player is playing.
	Playing() bool
	// Pauses the node’s playback.
	Pause()
	// Clears all of the node’s events you schedule and stops playback.
	Stop()
}

// Init initializes the instance.
func (a AVAudioPlayerNode) Init() AVAudioPlayerNode {
	rv := objc.Send[AVAudioPlayerNode](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPlayerNode) Autorelease() AVAudioPlayerNode {
	rv := objc.Send[AVAudioPlayerNode](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPlayerNode creates a new AVAudioPlayerNode instance.
func NewAVAudioPlayerNode() AVAudioPlayerNode {
	class := getAVAudioPlayerNodeClass()
	rv := objc.Send[AVAudioPlayerNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Schedules the playing of an entire audio file.
//
// file: The URL of the file to play.
//
// when: The time the buffer plays. For more information, see [AVAudioPlayerNode].
//
// completionHandler: The handler the system calls after the player schedules the file for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleFile(_:at:completionHandler:)
func (a AVAudioPlayerNode) ScheduleFileAtTimeCompletionHandler(file IAVAudioFile, when IAVAudioTime, completionHandler ErrorHandler) {
_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleFile:atTime:completionHandler:"), file, when, _block2)
}
// Schedules the playing of an entire audio file with a callback option you
// specify.
//
// file: The file to play.
//
// when: The time the file plays.
//
// callbackType: The option to specify when the system must call the completion handler.
//
// completionHandler: The handler the system calls after the player schedules the file for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleFile(_:at:completionCallbackType:completionHandler:)
func (a AVAudioPlayerNode) ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler(file IAVAudioFile, when IAVAudioTime, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler) {
_block3, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleFile:atTime:completionCallbackType:completionHandler:"), file, when, callbackType, _block3)
}
// Schedules the playing of an audio file segment.
//
// file: The URL of the file to play.
//
// startFrame: The starting frame position in the stream.
//
// numberFrames: The number of frames to play.
//
// when: The time the buffer plays. For more information, see [AVAudioPlayerNode].
//
// completionHandler: The handler the system calls after the player schedules the segment for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleSegment(_:startingFrame:frameCount:at:completionHandler:)
func (a AVAudioPlayerNode) ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler(file IAVAudioFile, startFrame AVAudioFramePosition, numberFrames AVAudioFrameCount, when IAVAudioTime, completionHandler ErrorHandler) {
_block4, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleSegment:startingFrame:frameCount:atTime:completionHandler:"), file, startFrame, numberFrames, when, _block4)
}
// Schedules the playing of an audio file segment with a callback option you
// specify.
//
// file: The file to play.
//
// startFrame: The starting frame position in the stream.
//
// numberFrames: The number of frames to play.
//
// when: The time the region plays.
//
// callbackType: The option to specify when the system must call the completion handler.
//
// completionHandler: The handler the system calls after the player schedules the segment for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleSegment(_:startingFrame:frameCount:at:completionCallbackType:completionHandler:)
func (a AVAudioPlayerNode) ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler(file IAVAudioFile, startFrame AVAudioFramePosition, numberFrames AVAudioFrameCount, when IAVAudioTime, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler) {
_block5, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleSegment:startingFrame:frameCount:atTime:completionCallbackType:completionHandler:"), file, startFrame, numberFrames, when, callbackType, _block5)
}
// Schedules the playing samples from an audio buffer at the time and playback
// options you specify.
//
// buffer: The buffer to play.
//
// when: The time the buffer plays. For more information, see [AVAudioPlayerNode].
//
// options: The playback options that control buffer scheduling.
//
// completionHandler: The handler the system calls after the player schedules the buffer for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:at:options:completionHandler:)
func (a AVAudioPlayerNode) ScheduleBufferAtTimeOptionsCompletionHandler(buffer IAVAudioPCMBuffer, when IAVAudioTime, options AVAudioPlayerNodeBufferOptions, completionHandler ErrorHandler) {
_block3, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleBuffer:atTime:options:completionHandler:"), buffer, when, options, _block3)
}
// Schedules the playing samples from an audio buffer.
//
// buffer: The buffer to play.
//
// completionHandler: The handler the system calls after the player schedules the buffer for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:completionHandler:)
func (a AVAudioPlayerNode) ScheduleBufferCompletionHandler(buffer IAVAudioPCMBuffer, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleBuffer:completionHandler:"), buffer, _block1)
}
// Schedules the playing samples from an audio buffer with the playback
// options you specify.
//
// buffer: The buffer to play.
//
// when: The time the buffer plays.
//
// options: The playback options that control buffer scheduling.
//
// callbackType: The option to specify when the system must call the completion handler.
//
// completionHandler: The handler the system calls after the player schedules the buffer for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:at:options:completionCallbackType:completionHandler:)
func (a AVAudioPlayerNode) ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler(buffer IAVAudioPCMBuffer, when IAVAudioTime, options AVAudioPlayerNodeBufferOptions, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler) {
_block4, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleBuffer:atTime:options:completionCallbackType:completionHandler:"), buffer, when, options, callbackType, _block4)
}
// Schedules the playing samples from an audio buffer with the callback option
// you specify.
//
// buffer: The buffer to play.
//
// callbackType: The option to specify when the system must call the completion handler.
//
// completionHandler: The handler the system calls after the player schedules the buffer for
// playback on the render thread, or the player stops.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:completionCallbackType:completionHandler:)
func (a AVAudioPlayerNode) ScheduleBufferCompletionCallbackTypeCompletionHandler(buffer IAVAudioPCMBuffer, callbackType AVAudioPlayerNodeCompletionCallbackType, completionHandler ErrorHandler) {
_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("scheduleBuffer:completionCallbackType:completionHandler:"), buffer, callbackType, _block2)
}
// Converts from player time to node time.
//
// playerTime: A time relative to the player’s start time.
//
// # Return Value
// 
// A node time, or `nil` if the player isn’t playing.
//
// # Discussion
// 
// For more information about this method and its inverse
// [PlayerTimeForNodeTime], see [AVAudioPlayerNode].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/nodeTime(forPlayerTime:)
func (a AVAudioPlayerNode) NodeTimeForPlayerTime(playerTime IAVAudioTime) IAVAudioTime {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nodeTimeForPlayerTime:"), playerTime)
	return AVAudioTimeFromID(rv)
}
// Converts from node time to player time.
//
// nodeTime: The node time.
//
// # Return Value
// 
// A time relative to the player’s start time, or `nil` if the player
// isn’t playing.
//
// # Discussion
// 
// For more information about this method and its inverse
// [NodeTimeForPlayerTime], see [AVAudioPlayerNode].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/playerTime(forNodeTime:)
func (a AVAudioPlayerNode) PlayerTimeForNodeTime(nodeTime IAVAudioTime) IAVAudioTime {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("playerTimeForNodeTime:"), nodeTime)
	return AVAudioTimeFromID(rv)
}
// Prepares the file regions or buffers you schedule for playback.
//
// frameCount: The number of sample frames of data to prepare.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/prepare(withFrameCount:)
func (a AVAudioPlayerNode) PrepareWithFrameCount(frameCount AVAudioFrameCount) {
	objc.Send[objc.ID](a.ID, objc.Sel("prepareWithFrameCount:"), frameCount)
}
// Starts or resumes playback immediately.
//
// # Discussion
// 
// This is equivalent to [PlayAtTime] with a value of `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/play()
func (a AVAudioPlayerNode) Play() {
	objc.Send[objc.ID](a.ID, objc.Sel("play"))
}
// Starts or resumes playback at a time you specify.
//
// when: The node time to start or resume playback. Passing `nil` starts playback
// immediately.
//
// # Discussion
// 
// This node is initially in a paused state. The framework enqueues your
// requests to play buffers or file segments, and any necessary decoding
// begins immediately. Playback doesn’t begin until the player starts
// playing through this method.
// 
// The following example code shows how to start a player `0.5` seconds in the
// future:
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/play(at:)
func (a AVAudioPlayerNode) PlayAtTime(when IAVAudioTime) {
	objc.Send[objc.ID](a.ID, objc.Sel("playAtTime:"), when)
}
// Pauses the node’s playback.
//
// # Discussion
// 
// The player’s sample time doesn’t advance while the node is in a paused
// state.
// 
// Pausing or stopping all of the players you connect to an engine doesn’t
// pause or stop the engine or the underlying hardware. You must explicitly
// pause or stop the engine for the hardware to stop. When your app doesn’t
// need to play audio, pause or stop the engine to minimize power consumption.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/pause()
func (a AVAudioPlayerNode) Pause() {
	objc.Send[objc.ID](a.ID, objc.Sel("pause"))
}
// Clears all of the node’s events you schedule and stops playback.
//
// # Discussion
// 
// Clears all events you schedule, including any events in the middle of
// playing. It resets the node’s sample time to `0`, and doesn’t proceed
// until the node starts again through [Play] or [PlayAtTime].
// 
// Pausing or stopping all of the players you connect to an engine doesn’t
// pause or stop the engine or the underlying hardware. You must explicitly
// pause or stop the engine for the hardware to stop. When your app doesn’t
// need to play audio, pause or stop the engine to minimize power consumption.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/stop()
func (a AVAudioPlayerNode) Stop() {
	objc.Send[objc.ID](a.ID, objc.Sel("stop"))
}
// Gets the audio mixing destination object that corresponds to the specified
// mixer node and input bus.
//
// mixer: The mixer to get destination details for.
//
// bus: The input bus.
//
// # Return Value
// 
// Returns `self` if the specified mixer or input bus matches its connection
// point. If the mixer or input bus doesn’t match its connection point, or
// if the source node isn’t in a connected state to the mixer or input bus,
// the method returns `nil.`
//
// # Discussion
// 
// When you connect a source node to multiple mixers downstream, setting
// [AVAudioMixing] properties directly on the source node applies the change
// to all of them. Use this method to get the corresponding
// [AVAudioMixingDestination] for a specific mixer. Properties set on
// individual destination instances don’t reflect at the source node level.
// 
// If there’s any disconnection between the source and mixer nodes, the
// return value can be invalid. Fetch the return value every time you want to
// set or get properties on a specific mixer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/destination(forMixer:bus:)
func (a AVAudioPlayerNode) DestinationForMixerBus(mixer IAVAudioNode, bus AVAudioNodeBus) IAVAudioMixingDestination {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("destinationForMixer:bus:"), mixer, bus)
	return AVAudioMixingDestinationFromID(rv)
}

// A Boolean value that indicates whether the player is playing.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/isPlaying
func (a AVAudioPlayerNode) Playing() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPlaying"))
	return rv
}
// A value that simulates filtering of the direct path of sound due to an
// obstacle.
//
// # Discussion
// 
// The value of `obstruction` is in decibels. The system blocks only the
// direct path of sound between the source and listener.
// 
// The default value is `0.0`, and the range of valid values is `-100` to `0`.
// Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/obstruction
func (a AVAudioPlayerNode) Obstruction() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("obstruction"))
	return rv
}
func (a AVAudioPlayerNode) SetObstruction(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setObstruction:"), value)
}
// A value that simulates filtering of the direct and reverb paths of sound
// due to an obstacle.
//
// # Discussion
// 
// The value of `obstruction` is in decibels. The system blocks the direct and
// reverb paths of sound between the source and listener.
// 
// The default value is `0.0`, and the range of valid values is `-100` to `0`.
// Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/occlusion
func (a AVAudioPlayerNode) Occlusion() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("occlusion"))
	return rv
}
func (a AVAudioPlayerNode) SetOcclusion(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setOcclusion:"), value)
}
// The bus’s stereo pan.
//
// # Discussion
// 
// The default value is `0.0`, and the range of valid values is `-1.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioStereoMixing/pan
func (a AVAudioPlayerNode) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}
func (a AVAudioPlayerNode) SetPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPan:"), value)
}
// The in-head mode for a point source.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/pointSourceInHeadMode
func (a AVAudioPlayerNode) PointSourceInHeadMode() AVAudio3DMixingPointSourceInHeadMode {
	rv := objc.Send[AVAudio3DMixingPointSourceInHeadMode](a.ID, objc.Sel("pointSourceInHeadMode"))
	return AVAudio3DMixingPointSourceInHeadMode(rv)
}
func (a AVAudioPlayerNode) SetPointSourceInHeadMode(value AVAudio3DMixingPointSourceInHeadMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setPointSourceInHeadMode:"), value)
}
// The location of the source in the 3D environment.
//
// # Discussion
// 
// The system specifies the coordinates in meters. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/position
func (a AVAudioPlayerNode) Position() AVAudio3DPoint {
	rv := objc.Send[AVAudio3DPoint](a.ID, objc.Sel("position"))
	return AVAudio3DPoint(rv)
}
func (a AVAudioPlayerNode) SetPosition(value AVAudio3DPoint) {
	objc.Send[struct{}](a.ID, objc.Sel("setPosition:"), value)
}
// A value that changes the playback rate of the input signal.
//
// # Discussion
// 
// A value of `2.0` results in the output audio playing one octave higher. A
// value of `0.5` results in the output audio playing one octave lower.
// 
// The default value is `1.0`, and the range of valid values is `0.5` to
// `2.0`. Only the [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/rate
func (a AVAudioPlayerNode) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioPlayerNode) SetRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRate:"), value)
}
// The type of rendering algorithm the mixer uses.
//
// # Discussion
// 
// Depending on the current output format of the [AVAudioEnvironmentNode]
// instance, the system may only support a subset of the rendering algorithms.
// You can retrieve an array of valid rendering algorithms by calling the
// [ApplicableRenderingAlgorithms] function of the [AVAudioEnvironmentNode]
// instance.
// 
// The default rendering algorithm is
// [Audio3DMixingRenderingAlgorithmEqualPowerPanning]. Only the
// [AVAudioEnvironmentNode] class implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/renderingAlgorithm
func (a AVAudioPlayerNode) RenderingAlgorithm() AVAudio3DMixingRenderingAlgorithm {
	rv := objc.Send[AVAudio3DMixingRenderingAlgorithm](a.ID, objc.Sel("renderingAlgorithm"))
	return AVAudio3DMixingRenderingAlgorithm(rv)
}
func (a AVAudioPlayerNode) SetRenderingAlgorithm(value AVAudio3DMixingRenderingAlgorithm) {
	objc.Send[struct{}](a.ID, objc.Sel("setRenderingAlgorithm:"), value)
}
// A value that controls the blend of dry and reverb processed audio.
//
// # Discussion
// 
// This property controls the amount of the source’s audio that the
// [AVAudioEnvironmentNode] instance processes. A value of `0.5` results in an
// equal blend of dry and processed (wet) audio.
// 
// The default is `0.0`, and the range of valid values is `0.0` (completely
// dry) to `1.0` (completely wet). Only the [AVAudioEnvironmentNode] class
// implements this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/reverbBlend
func (a AVAudioPlayerNode) ReverbBlend() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("reverbBlend"))
	return rv
}
func (a AVAudioPlayerNode) SetReverbBlend(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setReverbBlend:"), value)
}
// The source mode for the input bus of the audio environment node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudio3DMixing/sourceMode
func (a AVAudioPlayerNode) SourceMode() AVAudio3DMixingSourceMode {
	rv := objc.Send[AVAudio3DMixingSourceMode](a.ID, objc.Sel("sourceMode"))
	return AVAudio3DMixingSourceMode(rv)
}
func (a AVAudioPlayerNode) SetSourceMode(value AVAudio3DMixingSourceMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setSourceMode:"), value)
}
// The bus’s input volume.
//
// # Discussion
// 
// The default value is `1.0`, and the range of valid values is `0.0` to
// `1.0`. Only the [AVAudioEnvironmentNode] and the [AVAudioMixerNode]
// implement this property.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixing/volume
func (a AVAudioPlayerNode) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}
func (a AVAudioPlayerNode) SetVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVolume:"), value)
}

			// Protocol methods for AVAudioMixing
			

