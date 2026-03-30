// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemMetadataOutput] class.
var (
	_AVPlayerItemMetadataOutputClass     AVPlayerItemMetadataOutputClass
	_AVPlayerItemMetadataOutputClassOnce sync.Once
)

func getAVPlayerItemMetadataOutputClass() AVPlayerItemMetadataOutputClass {
	_AVPlayerItemMetadataOutputClassOnce.Do(func() {
		_AVPlayerItemMetadataOutputClass = AVPlayerItemMetadataOutputClass{class: objc.GetClass("AVPlayerItemMetadataOutput")}
	})
	return _AVPlayerItemMetadataOutputClass
}

// GetAVPlayerItemMetadataOutputClass returns the class object for AVPlayerItemMetadataOutput.
func GetAVPlayerItemMetadataOutputClass() AVPlayerItemMetadataOutputClass {
	return getAVPlayerItemMetadataOutputClass()
}

type AVPlayerItemMetadataOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemMetadataOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemMetadataOutputClass) Alloc() AVPlayerItemMetadataOutput {
	rv := objc.Send[AVPlayerItemMetadataOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that vends collections of metadata items that a player item’s
// tracks carry.
//
// # Overview
//
// # Creating a metadata output
//
//   - [AVPlayerItemMetadataOutput.InitWithIdentifiers]: Creates an instance of AVPlayerItemMetadataOutput.
//
// # Configuring the delegate
//
//   - [AVPlayerItemMetadataOutput.AdvanceIntervalForDelegateInvocation]: The time interval, in seconds, the player item metadata output object messages its delegate earlier than normal.
//   - [AVPlayerItemMetadataOutput.SetAdvanceIntervalForDelegateInvocation]
//   - [AVPlayerItemMetadataOutput.Delegate]: The delegate object.
//   - [AVPlayerItemMetadataOutput.DelegateQueue]: The dispatch queue on which messages are sent to the delegate.
//   - [AVPlayerItemMetadataOutput.SetDelegateQueue]: Sets the delegate and a dispatch queue on which the delegate is called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput
type AVPlayerItemMetadataOutput struct {
	AVPlayerItemOutput
}

// AVPlayerItemMetadataOutputFromID constructs a [AVPlayerItemMetadataOutput] from an objc.ID.
//
// An object that vends collections of metadata items that a player item’s
// tracks carry.
func AVPlayerItemMetadataOutputFromID(id objc.ID) AVPlayerItemMetadataOutput {
	return AVPlayerItemMetadataOutput{AVPlayerItemOutput: AVPlayerItemOutputFromID(id)}
}

// NOTE: AVPlayerItemMetadataOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemMetadataOutput] class.
//
// # Creating a metadata output
//
//   - [IAVPlayerItemMetadataOutput.InitWithIdentifiers]: Creates an instance of AVPlayerItemMetadataOutput.
//
// # Configuring the delegate
//
//   - [IAVPlayerItemMetadataOutput.AdvanceIntervalForDelegateInvocation]: The time interval, in seconds, the player item metadata output object messages its delegate earlier than normal.
//   - [IAVPlayerItemMetadataOutput.SetAdvanceIntervalForDelegateInvocation]
//   - [IAVPlayerItemMetadataOutput.Delegate]: The delegate object.
//   - [IAVPlayerItemMetadataOutput.DelegateQueue]: The dispatch queue on which messages are sent to the delegate.
//   - [IAVPlayerItemMetadataOutput.SetDelegateQueue]: Sets the delegate and a dispatch queue on which the delegate is called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput
type IAVPlayerItemMetadataOutput interface {
	IAVPlayerItemOutput

	// Topic: Creating a metadata output

	// Creates an instance of AVPlayerItemMetadataOutput.
	InitWithIdentifiers(identifiers []string) AVPlayerItemMetadataOutput

	// Topic: Configuring the delegate

	// The time interval, in seconds, the player item metadata output object messages its delegate earlier than normal.
	AdvanceIntervalForDelegateInvocation() float64
	SetAdvanceIntervalForDelegateInvocation(value float64)
	// The delegate object.
	Delegate() AVPlayerItemMetadataOutputPushDelegate
	// The dispatch queue on which messages are sent to the delegate.
	DelegateQueue() dispatch.Queue
	// Sets the delegate and a dispatch queue on which the delegate is called.
	SetDelegateQueue(delegate AVPlayerItemMetadataOutputPushDelegate, delegateQueue dispatch.Queue)
}

// Init initializes the instance.
func (p AVPlayerItemMetadataOutput) Init() AVPlayerItemMetadataOutput {
	rv := objc.Send[AVPlayerItemMetadataOutput](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemMetadataOutput) Autorelease() AVPlayerItemMetadataOutput {
	rv := objc.Send[AVPlayerItemMetadataOutput](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemMetadataOutput creates a new AVPlayerItemMetadataOutput instance.
func NewAVPlayerItemMetadataOutput() AVPlayerItemMetadataOutput {
	class := getAVPlayerItemMetadataOutputClass()
	rv := objc.Send[AVPlayerItemMetadataOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an instance of AVPlayerItemMetadataOutput.
//
// identifiers: A array of metadata identifiers indicating the metadata items that the
// output should provide.
//
// # Return Value
//
// An AVPlayerItemMetadataOutput instance.
//
// # Discussion
//
// Pass `nil` to receive all of the timed metadata from all enabled
// [AVPlayerItemTracks] that carry timed metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/init(identifiers:)
func NewPlayerItemMetadataOutputWithIdentifiers(identifiers []string) AVPlayerItemMetadataOutput {
	instance := getAVPlayerItemMetadataOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifiers:"), objectivec.StringSliceToNSArray(identifiers))
	return AVPlayerItemMetadataOutputFromID(rv)
}

// Creates an instance of AVPlayerItemMetadataOutput.
//
// identifiers: A array of metadata identifiers indicating the metadata items that the
// output should provide.
//
// # Return Value
//
// An AVPlayerItemMetadataOutput instance.
//
// # Discussion
//
// Pass `nil` to receive all of the timed metadata from all enabled
// [AVPlayerItemTracks] that carry timed metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/init(identifiers:)
func (p AVPlayerItemMetadataOutput) InitWithIdentifiers(identifiers []string) AVPlayerItemMetadataOutput {
	rv := objc.Send[AVPlayerItemMetadataOutput](p.ID, objc.Sel("initWithIdentifiers:"), objectivec.StringSliceToNSArray(identifiers))
	return rv
}

// Sets the delegate and a dispatch queue on which the delegate is called.
//
// delegate: An object conforming to [AVPlayerItemMetadataOutputPushDelegate] protocol.
//
// delegateQueue: A dispatch queue on which all delegate methods will be called.
//
// # Discussion
//
// You specify the metadata delegate, and a dispatch queue on which it will be
// called, to be notified as new metadata is encountered in the source media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/setDelegate(_:queue:)
func (p AVPlayerItemMetadataOutput) SetDelegateQueue(delegate AVPlayerItemMetadataOutputPushDelegate, delegateQueue dispatch.Queue) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateQueue.Handle()))
}

// The time interval, in seconds, the player item metadata output object
// messages its delegate earlier than normal.
//
// # Discussion
//
// If possible, an [AVPlayerItemMetadataOutput] will message its delegate
// `advanceIntervalForDelegateInvocation` seconds earlier than otherwise. If
// the value you provide is large, effectively requesting provision of samples
// earlier than the [AVPlayerItemMetadataOutput] is prepared to act on them,
// the delegate will be invoked as soon as possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/advanceIntervalForDelegateInvocation
func (p AVPlayerItemMetadataOutput) AdvanceIntervalForDelegateInvocation() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}
func (p AVPlayerItemMetadataOutput) SetAdvanceIntervalForDelegateInvocation(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}

// The delegate object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/delegate
func (p AVPlayerItemMetadataOutput) Delegate() AVPlayerItemMetadataOutputPushDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return AVPlayerItemMetadataOutputPushDelegateObjectFromID(rv)
}

// The dispatch queue on which messages are sent to the delegate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMetadataOutput/delegateQueue
func (p AVPlayerItemMetadataOutput) DelegateQueue() dispatch.Queue {
	rv := objc.Send[uintptr](p.ID, objc.Sel("delegateQueue"))
	return dispatch.QueueFromHandle(rv)
}
