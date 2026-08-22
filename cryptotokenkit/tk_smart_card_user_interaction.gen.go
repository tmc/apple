// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCardUserInteraction] class.
var (
	_TKSmartCardUserInteractionClass     TKSmartCardUserInteractionClass
	_TKSmartCardUserInteractionClassOnce sync.Once
)

func getTKSmartCardUserInteractionClass() TKSmartCardUserInteractionClass {
	_TKSmartCardUserInteractionClassOnce.Do(func() {
		_TKSmartCardUserInteractionClass = TKSmartCardUserInteractionClass{class: objc.GetClass("TKSmartCardUserInteraction")}
	})
	return _TKSmartCardUserInteractionClass
}

// GetTKSmartCardUserInteractionClass returns the class object for TKSmartCardUserInteraction.
func GetTKSmartCardUserInteractionClass() TKSmartCardUserInteractionClass {
	return getTKSmartCardUserInteractionClass()
}

type TKSmartCardUserInteractionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardUserInteractionClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardUserInteractionClass) Alloc() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// The base class for encapsulating user interaction with a Smart Card reader.
//
// # Overview
//
// There are two types of user interactions: those for secure PIN change and
// those for secure PIN validation. These interactions are instances of the
// [TKSmartCardUserInteractionForSecurePINChange], or
// [TKSmartCardUserInteractionForSecurePINVerification] subclasses of
// [TKSmartCardUserInteractionForPINOperation], respectively.
// [TKSmartCardUserInteractionForPINOperation] is a subclass of
// [TKSmartCardUserInteraction].
//
// You interact with instances of one of the subclasses of
// [TKSmartCardUserInteractionForPINOperation]when calling the
// [TKSmartCard.UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset]
// and
// [TKSmartCard.UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset]
// methods on an [TKSmartCard] object.
//
// # Handling User Interaction Events
//
//   - [TKSmartCardUserInteraction.Delegate]: The delegate for observing events that occur during the user interaction.
//   - [TKSmartCardUserInteraction.SetDelegate]
//
// # Configuring Timeout
//
//   - [TKSmartCardUserInteraction.InitialTimeout]: The timeout, in seconds, for initial interaction. If set to `0`, the reader-defined default timeout is used. `0` by default.
//   - [TKSmartCardUserInteraction.SetInitialTimeout]
//   - [TKSmartCardUserInteraction.InteractionTimeout]: The timeout, in seconds, after the first key stroke. If set to `0`, the reader-defined default timeout is used. `0` by default.
//   - [TKSmartCardUserInteraction.SetInteractionTimeout]
//
// # Starting and Stopping
//
//   - [TKSmartCardUserInteraction.RunWithReply]: Runs the user interaction and asynchronously receives a reply.
//   - [TKSmartCardUserInteraction.Cancel]: Attempts to cancel an interaction started by calling [run(reply:)](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/run(reply:)>). For certain interactions, cancellation may not be available.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction
type TKSmartCardUserInteraction struct {
	objectivec.Object
}

// TKSmartCardUserInteractionFromID constructs a [TKSmartCardUserInteraction] from an objc.ID.
//
// The base class for encapsulating user interaction with a Smart Card reader.
func TKSmartCardUserInteractionFromID(id objc.ID) TKSmartCardUserInteraction {
	return TKSmartCardUserInteraction{objectivec.Object{ID: id}}
}

// NOTE: TKSmartCardUserInteraction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardUserInteraction] class.
//
// # Handling User Interaction Events
//
//   - [ITKSmartCardUserInteraction.Delegate]: The delegate for observing events that occur during the user interaction.
//   - [ITKSmartCardUserInteraction.SetDelegate]
//
// # Configuring Timeout
//
//   - [ITKSmartCardUserInteraction.InitialTimeout]: The timeout, in seconds, for initial interaction. If set to `0`, the reader-defined default timeout is used. `0` by default.
//   - [ITKSmartCardUserInteraction.SetInitialTimeout]
//   - [ITKSmartCardUserInteraction.InteractionTimeout]: The timeout, in seconds, after the first key stroke. If set to `0`, the reader-defined default timeout is used. `0` by default.
//   - [ITKSmartCardUserInteraction.SetInteractionTimeout]
//
// # Starting and Stopping
//
//   - [ITKSmartCardUserInteraction.RunWithReply]: Runs the user interaction and asynchronously receives a reply.
//   - [ITKSmartCardUserInteraction.Cancel]: Attempts to cancel an interaction started by calling [run(reply:)](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/run(reply:)>). For certain interactions, cancellation may not be available.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction
type ITKSmartCardUserInteraction interface {
	objectivec.IObject

	// Topic: Handling User Interaction Events

	// The delegate for observing events that occur during the user interaction.
	Delegate() TKSmartCardUserInteractionDelegate
	SetDelegate(value TKSmartCardUserInteractionDelegate)

	// Topic: Configuring Timeout

	// The timeout, in seconds, for initial interaction. If set to `0`, the reader-defined default timeout is used. `0` by default.
	InitialTimeout() foundation.NSTimeInterval
	SetInitialTimeout(value foundation.NSTimeInterval)
	// The timeout, in seconds, after the first key stroke. If set to `0`, the reader-defined default timeout is used. `0` by default.
	InteractionTimeout() foundation.NSTimeInterval
	SetInteractionTimeout(value foundation.NSTimeInterval)

	// Topic: Starting and Stopping

	// Runs the user interaction and asynchronously receives a reply.
	RunWithReply(reply BoolErrorHandler)
	// Attempts to cancel an interaction started by calling [run(reply:)](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/run(reply:)>). For certain interactions, cancellation may not be available.
	Cancel() bool
}

// Init initializes the instance.
func (t TKSmartCardUserInteraction) Init() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardUserInteraction) Autorelease() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteraction creates a new TKSmartCardUserInteraction instance.
func NewTKSmartCardUserInteraction() TKSmartCardUserInteraction {
	class := getTKSmartCardUserInteractionClass()
	rv := objc.Send[TKSmartCardUserInteraction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Runs the user interaction and asynchronously receives a reply.
//
// reply: success: Whether the user interaction was successful. error: Contains
// information about the the error that occurred during the user interaction.
//
// The [NSError] object is created in the [TKErrorDomain] domain with a code
// in the [TKError.Code] enumeration.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/run(reply:)
//
// [TKError.Code]: https://developer.apple.com/documentation/CryptoTokenKit/TKError/Code
// [TKErrorDomain]: https://developer.apple.com/documentation/CryptoTokenKit/TKErrorDomain
func (t TKSmartCardUserInteraction) RunWithReply(reply BoolErrorHandler) {
	_block0, _ := NewBoolErrorBlock(reply)
	objc.Send[objc.ID](t.ID, objc.Sel("runWithReply:"), _block0)
}

// Attempts to cancel an interaction started by calling
// [TKSmartCardUserInteraction.RunWithReply]. For certain interactions,
// cancellation may not be available.
//
// # Return Value
//
// Returns false if the operation is not running, or if cancelation is not
// supported.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/cancel()
func (t TKSmartCardUserInteraction) Cancel() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("cancel"))
	return rv
}

// The delegate for observing events that occur during the user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/delegate
func (t TKSmartCardUserInteraction) Delegate() TKSmartCardUserInteractionDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return TKSmartCardUserInteractionDelegateObjectFromID(rv)
}
func (t TKSmartCardUserInteraction) SetDelegate(value TKSmartCardUserInteractionDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}

// The timeout, in seconds, for initial interaction. If set to `0`, the
// reader-defined default timeout is used. `0` by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/initialTimeout
func (t TKSmartCardUserInteraction) InitialTimeout() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](t.ID, objc.Sel("initialTimeout"))
	return foundation.NSTimeInterval(rv)
}
func (t TKSmartCardUserInteraction) SetInitialTimeout(value foundation.NSTimeInterval) {
	objc.Send[struct{}](t.ID, objc.Sel("setInitialTimeout:"), value)
}

// The timeout, in seconds, after the first key stroke. If set to `0`, the
// reader-defined default timeout is used. `0` by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/interactionTimeout
func (t TKSmartCardUserInteraction) InteractionTimeout() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](t.ID, objc.Sel("interactionTimeout"))
	return foundation.NSTimeInterval(rv)
}
func (t TKSmartCardUserInteraction) SetInteractionTimeout(value foundation.NSTimeInterval) {
	objc.Send[struct{}](t.ID, objc.Sel("setInteractionTimeout:"), value)
}

// RunWithReplySync is a synchronous wrapper around [TKSmartCardUserInteraction.RunWithReply].
// It blocks until the completion handler fires or the context is cancelled.
func (t TKSmartCardUserInteraction) RunWithReplySync(ctx context.Context) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.RunWithReply(func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}
