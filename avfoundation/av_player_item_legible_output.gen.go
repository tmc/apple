// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemLegibleOutput] class.
var (
	_AVPlayerItemLegibleOutputClass     AVPlayerItemLegibleOutputClass
	_AVPlayerItemLegibleOutputClassOnce sync.Once
)

func getAVPlayerItemLegibleOutputClass() AVPlayerItemLegibleOutputClass {
	_AVPlayerItemLegibleOutputClassOnce.Do(func() {
		_AVPlayerItemLegibleOutputClass = AVPlayerItemLegibleOutputClass{class: objc.GetClass("AVPlayerItemLegibleOutput")}
	})
	return _AVPlayerItemLegibleOutputClass
}

// GetAVPlayerItemLegibleOutputClass returns the class object for AVPlayerItemLegibleOutput.
func GetAVPlayerItemLegibleOutputClass() AVPlayerItemLegibleOutputClass {
	return getAVPlayerItemLegibleOutputClass()
}

type AVPlayerItemLegibleOutputClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemLegibleOutputClass) Alloc() AVPlayerItemLegibleOutput {
	rv := objc.Send[AVPlayerItemLegibleOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that vends attributed strings for media with a legible
// characteristic.
//
// # Creating a legible output
//
//   - [AVPlayerItemLegibleOutput.InitWithMediaSubtypesForNativeRepresentation]: Creates an initialized legible-output object.
//
// # Configuring text styling
//
//   - [AVPlayerItemLegibleOutput.TextStylingResolution]: A string identifier indicating the degree of text styling to be applied to attributed strings vended by the  object.
//   - [AVPlayerItemLegibleOutput.SetTextStylingResolution]
//
// # Configuring the delegate
//
//   - [AVPlayerItemLegibleOutput.Delegate]: The delegate of the output class.
//   - [AVPlayerItemLegibleOutput.SetDelegateQueue]: Sets the receiver’s delegate and a dispatch queue on which the delegate is called.
//   - [AVPlayerItemLegibleOutput.AdvanceIntervalForDelegateInvocation]: The time interval, in seconds, that a player item legible output object messages its delegate earlier than normal.
//   - [AVPlayerItemLegibleOutput.SetAdvanceIntervalForDelegateInvocation]
//   - [AVPlayerItemLegibleOutput.DelegateQueue]: The dispatch queue on which the delegate is called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput
type AVPlayerItemLegibleOutput struct {
	AVPlayerItemOutput
}

// AVPlayerItemLegibleOutputFromID constructs a [AVPlayerItemLegibleOutput] from an objc.ID.
//
// An object that vends attributed strings for media with a legible
// characteristic.
func AVPlayerItemLegibleOutputFromID(id objc.ID) AVPlayerItemLegibleOutput {
	return AVPlayerItemLegibleOutput{AVPlayerItemOutput: AVPlayerItemOutputFromID(id)}
}
// NOTE: AVPlayerItemLegibleOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemLegibleOutput] class.
//
// # Creating a legible output
//
//   - [IAVPlayerItemLegibleOutput.InitWithMediaSubtypesForNativeRepresentation]: Creates an initialized legible-output object.
//
// # Configuring text styling
//
//   - [IAVPlayerItemLegibleOutput.TextStylingResolution]: A string identifier indicating the degree of text styling to be applied to attributed strings vended by the  object.
//   - [IAVPlayerItemLegibleOutput.SetTextStylingResolution]
//
// # Configuring the delegate
//
//   - [IAVPlayerItemLegibleOutput.Delegate]: The delegate of the output class.
//   - [IAVPlayerItemLegibleOutput.SetDelegateQueue]: Sets the receiver’s delegate and a dispatch queue on which the delegate is called.
//   - [IAVPlayerItemLegibleOutput.AdvanceIntervalForDelegateInvocation]: The time interval, in seconds, that a player item legible output object messages its delegate earlier than normal.
//   - [IAVPlayerItemLegibleOutput.SetAdvanceIntervalForDelegateInvocation]
//   - [IAVPlayerItemLegibleOutput.DelegateQueue]: The dispatch queue on which the delegate is called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput
type IAVPlayerItemLegibleOutput interface {
	IAVPlayerItemOutput

	// Topic: Creating a legible output

	// Creates an initialized legible-output object.
	InitWithMediaSubtypesForNativeRepresentation(subtypes []foundation.NSNumber) AVPlayerItemLegibleOutput

	// Topic: Configuring text styling

	// A string identifier indicating the degree of text styling to be applied to attributed strings vended by the  object.
	TextStylingResolution() AVPlayerItemLegibleOutputTextStylingResolution
	SetTextStylingResolution(value AVPlayerItemLegibleOutputTextStylingResolution)

	// Topic: Configuring the delegate

	// The delegate of the output class.
	Delegate() AVPlayerItemLegibleOutputPushDelegate
	// Sets the receiver’s delegate and a dispatch queue on which the delegate is called.
	SetDelegateQueue(delegate AVPlayerItemLegibleOutputPushDelegate, delegateQueue dispatch.Queue)
	// The time interval, in seconds, that a player item legible output object messages its delegate earlier than normal.
	AdvanceIntervalForDelegateInvocation() float64
	SetAdvanceIntervalForDelegateInvocation(value float64)
	// The dispatch queue on which the delegate is called.
	DelegateQueue() dispatch.Queue
}

// Init initializes the instance.
func (p AVPlayerItemLegibleOutput) Init() AVPlayerItemLegibleOutput {
	rv := objc.Send[AVPlayerItemLegibleOutput](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemLegibleOutput) Autorelease() AVPlayerItemLegibleOutput {
	rv := objc.Send[AVPlayerItemLegibleOutput](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemLegibleOutput creates a new AVPlayerItemLegibleOutput instance.
func NewAVPlayerItemLegibleOutput() AVPlayerItemLegibleOutput {
	class := getAVPlayerItemLegibleOutputClass()
	rv := objc.Send[AVPlayerItemLegibleOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an initialized legible-output object.
//
// subtypes: An [NSArray] of [NSNumber] FourCC codes.
// //
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// # Return Value
// 
// An initialized instance of [AVPlayerItemLegibleOutput].
//
// # Discussion
// 
// When creating an instance you add media subtype FourCC codes as [NSNumber]
// objects to the `subtypes` array to elect to receive that type as a
// [CMSampleBuffer] instead of an attributed string. FourCC codes are
// converted to [NSNumber] objects as shown:
// 
// Initializing an [AVPlayerItemLegibleOutput] using the `init` method (which
// is preferred) is equivalent to calling this method with an empty `subtypes`
// array, which means that all legible data, regardless of media subtype, is
// delivered using [NSAttributedString] instances in a common format.
// 
// If a media subtype for which there is no legible data in the current player
// item is included in the media `subtypes` array, no error occurs. An
// [AVPlayerItemLegibleOutput] instance doesn’t vend closed caption data as
// a [CMSampleBuffer], so it is an error to include `'c608'` in the media
// subtypes array.
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/init(mediaSubtypesForNativeRepresentation:)
func NewPlayerItemLegibleOutputWithMediaSubtypesForNativeRepresentation(subtypes []foundation.NSNumber) AVPlayerItemLegibleOutput {
	instance := getAVPlayerItemLegibleOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMediaSubtypesForNativeRepresentation:"), objectivec.IObjectSliceToNSArray(subtypes))
	return AVPlayerItemLegibleOutputFromID(rv)
}

// Creates an initialized legible-output object.
//
// subtypes: An [NSArray] of [NSNumber] FourCC codes.
// //
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// # Return Value
// 
// An initialized instance of [AVPlayerItemLegibleOutput].
//
// # Discussion
// 
// When creating an instance you add media subtype FourCC codes as [NSNumber]
// objects to the `subtypes` array to elect to receive that type as a
// [CMSampleBuffer] instead of an attributed string. FourCC codes are
// converted to [NSNumber] objects as shown:
// 
// Initializing an [AVPlayerItemLegibleOutput] using the `init` method (which
// is preferred) is equivalent to calling this method with an empty `subtypes`
// array, which means that all legible data, regardless of media subtype, is
// delivered using [NSAttributedString] instances in a common format.
// 
// If a media subtype for which there is no legible data in the current player
// item is included in the media `subtypes` array, no error occurs. An
// [AVPlayerItemLegibleOutput] instance doesn’t vend closed caption data as
// a [CMSampleBuffer], so it is an error to include `'c608'` in the media
// subtypes array.
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/init(mediaSubtypesForNativeRepresentation:)
func (p AVPlayerItemLegibleOutput) InitWithMediaSubtypesForNativeRepresentation(subtypes []foundation.NSNumber) AVPlayerItemLegibleOutput {
	rv := objc.Send[AVPlayerItemLegibleOutput](p.ID, objc.Sel("initWithMediaSubtypesForNativeRepresentation:"), objectivec.IObjectSliceToNSArray(subtypes))
	return rv
}
// Sets the receiver’s delegate and a dispatch queue on which the delegate
// is called.
//
// delegate: An object conforming to the [AVPlayerItemLegibleOutputPushDelegate]
// protocol.
//
// delegateQueue: A dispatch queue on which all delegate methods will be called.
//
// # Discussion
// 
// Because the delegate is held using a zeroing-weak reference, it is safe to
// deallocate the delegate while the receiver still has a reference to it.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/setDelegate(_:queue:)
func (p AVPlayerItemLegibleOutput) SetDelegateQueue(delegate AVPlayerItemLegibleOutputPushDelegate, delegateQueue dispatch.Queue) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateQueue.Handle()))
}

// A string identifier indicating the degree of text styling to be applied to
// attributed strings vended by the object.
//
// # Discussion
// 
// Valid values are described in `Text Style Settings`. An exception
// ([InvalidArgumentException]) is raised if this property is set to any other
// value.
// 
// The default value is [default], which indicates that attributed strings
// vended by the receiver includes the same level of styling information that
// would be used if the text was rendered by an instance of [AVPlayerLayer].
//
// [default]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/TextStylingResolution-swift.struct/default
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/textStylingResolution-swift.property
func (p AVPlayerItemLegibleOutput) TextStylingResolution() AVPlayerItemLegibleOutputTextStylingResolution {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("textStylingResolution"))
	return AVPlayerItemLegibleOutputTextStylingResolution(foundation.NSStringFromID(rv).String())
}
func (p AVPlayerItemLegibleOutput) SetTextStylingResolution(value AVPlayerItemLegibleOutputTextStylingResolution) {
	objc.Send[struct{}](p.ID, objc.Sel("setTextStylingResolution:"), objc.String(string(value)))
}
// The delegate of the output class.
//
// # Discussion
// 
// Because the delegate is held using a zeroing-weak reference, this property
// has a value of `nil` after a delegate that was previously set has been
// deallocated.
// 
// This property does not support key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/delegate
func (p AVPlayerItemLegibleOutput) Delegate() AVPlayerItemLegibleOutputPushDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return AVPlayerItemLegibleOutputPushDelegateObjectFromID(rv)
}
// The time interval, in seconds, that a player item legible output object
// messages its delegate earlier than normal.
//
// # Discussion
// 
// If possible, an [AVPlayerItemLegibleOutput] instance messages its delegate
// `advanceIntervalForDelegateInvocation` seconds earlier than it otherwise
// would.
// 
// If the value provided is large, the delegate methods are invoked as soon as
// possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/advanceIntervalForDelegateInvocation
func (p AVPlayerItemLegibleOutput) AdvanceIntervalForDelegateInvocation() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}
func (p AVPlayerItemLegibleOutput) SetAdvanceIntervalForDelegateInvocation(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}
// The dispatch queue on which the delegate is called.
//
// # Discussion
// 
// This property does not support key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/delegateQueue
func (p AVPlayerItemLegibleOutput) DelegateQueue() dispatch.Queue {
	rv := objc.Send[uintptr](p.ID, objc.Sel("delegateQueue"))
	return dispatch.QueueFromHandle(rv)
}

