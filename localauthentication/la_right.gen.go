// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LARight] class.
var (
	_LARightClass     LARightClass
	_LARightClassOnce sync.Once
)

func getLARightClass() LARightClass {
	_LARightClassOnce.Do(func() {
		_LARightClass = LARightClass{class: objc.GetClass("LARight")}
	})
	return _LARightClass
}

// GetLARightClass returns the class object for LARight.
func GetLARightClass() LARightClass {
	return getLARightClass()
}

type LARightClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LARightClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LARightClass) Alloc() LARight {
	rv := objc.Send[LARight](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// A grouped set of requirements that gate access to a resource or operation.
//
// # Overview
//
// Use [LARight] instances to protect access to portions of your app that may
// contain sensitive information. By default, [LARight] instances require
// people to authenticate with Face ID, Touch ID, Apple Watch, or the device
// passcode. The following creates an [LARight] with the default
// authentication requirements:
//
// # Authorizing a right
//
//   - [LARight.InitWithRequirement]: Creates a right with the authentication requirements you supply.
//   - [LARight.Tag]: An integer you use to identify a right.
//   - [LARight.SetTag]
//   - [LARight.AuthorizeWithLocalizedReasonCompletion]: Performs an authorization on the right.
//   - [LARight.AuthorizeWithLocalizedReasonInPresentationContextCompletion]: Performs an authorization on the right with a window context you supply.
//
// # Deauthorizing a right
//
//   - [LARight.DeauthorizeWithCompletion]: Invalidates a previously authorized right.
//
// # Monitoring authorization status
//
//   - [LARight.CheckCanAuthorizeWithCompletion]: Checks whether the right has permission to perform authorization.
//   - [LARight.State]: The current authorization state for a right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight
type LARight struct {
	objectivec.Object
}

// LARightFromID constructs a [LARight] from an objc.ID.
//
// A grouped set of requirements that gate access to a resource or operation.
func LARightFromID(id objc.ID) LARight {
	return LARight{objectivec.Object{ID: id}}
}

// NOTE: LARight adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LARight] class.
//
// # Authorizing a right
//
//   - [ILARight.InitWithRequirement]: Creates a right with the authentication requirements you supply.
//   - [ILARight.Tag]: An integer you use to identify a right.
//   - [ILARight.SetTag]
//   - [ILARight.AuthorizeWithLocalizedReasonCompletion]: Performs an authorization on the right.
//   - [ILARight.AuthorizeWithLocalizedReasonInPresentationContextCompletion]: Performs an authorization on the right with a window context you supply.
//
// # Deauthorizing a right
//
//   - [ILARight.DeauthorizeWithCompletion]: Invalidates a previously authorized right.
//
// # Monitoring authorization status
//
//   - [ILARight.CheckCanAuthorizeWithCompletion]: Checks whether the right has permission to perform authorization.
//   - [ILARight.State]: The current authorization state for a right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight
type ILARight interface {
	objectivec.IObject

	// Topic: Authorizing a right

	// Creates a right with the authentication requirements you supply.
	InitWithRequirement(requirement ILAAuthenticationRequirement) LARight
	// An integer you use to identify a right.
	Tag() int
	SetTag(value int)
	// Performs an authorization on the right.
	AuthorizeWithLocalizedReasonCompletion(localizedReason string, handler ErrorHandler)
	// Performs an authorization on the right with a window context you supply.
	AuthorizeWithLocalizedReasonInPresentationContextCompletion(localizedReason string, presentationContext unsafe.Pointer, handler ErrorHandler)

	// Topic: Deauthorizing a right

	// Invalidates a previously authorized right.
	DeauthorizeWithCompletion(handler VoidHandler)

	// Topic: Monitoring authorization status

	// Checks whether the right has permission to perform authorization.
	CheckCanAuthorizeWithCompletion(handler ErrorHandler)
	// The current authorization state for a right.
	State() LARightState
}

// Init initializes the instance.
func (r LARight) Init() LARight {
	rv := objc.Send[LARight](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r LARight) Autorelease() LARight {
	rv := objc.Send[LARight](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewLARight creates a new LARight instance.
func NewLARight() LARight {
	class := getLARightClass()
	rv := objc.Send[LARight](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a right with the authentication requirements you supply.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/init(requirement:)
func NewRightWithRequirement(requirement ILAAuthenticationRequirement) LARight {
	instance := getLARightClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRequirement:"), requirement)
	return LARightFromID(rv)
}

// Creates a right with the authentication requirements you supply.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/init(requirement:)
func (r LARight) InitWithRequirement(requirement ILAAuthenticationRequirement) LARight {
	rv := objc.Send[LARight](r.ID, objc.Sel("initWithRequirement:"), requirement)
	return rv
}

// Performs an authorization on the right.
//
// localizedReason: A reason for the authorization that the system displays to the user.
//
// handler: A completion handler called at the end of the authorization process.
//
// `error`: If `nil`, the authorization is successful. Otherwise, the error
// contains information about the failure reason.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/authorize(localizedReason:completion:)
func (r LARight) AuthorizeWithLocalizedReasonCompletion(localizedReason string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("authorizeWithLocalizedReason:completion:"), objc.String(localizedReason), _block1)
}

// Performs an authorization on the right with a window context you supply.
//
// presentationContext is a [localauthenticationembeddedui.LAPresentationContext].
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/authorize(localizedReason:in:completion:)
func (r LARight) AuthorizeWithLocalizedReasonInPresentationContextCompletion(localizedReason string, presentationContext unsafe.Pointer, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("authorizeWithLocalizedReason:inPresentationContext:completion:"), objc.String(localizedReason), presentationContext, _block2)
}

// Invalidates a previously authorized right.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/deauthorize(completion:)
func (r LARight) DeauthorizeWithCompletion(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("deauthorizeWithCompletion:"), _block0)
}

// Checks whether the right has permission to perform authorization.
//
// handler: A completion handler called when the authorization check finishes.
//
// `error`: If `nil`, the right can be authorized.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/checkCanAuthorize(completion:)
func (r LARight) CheckCanAuthorizeWithCompletion(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](r.ID, objc.Sel("checkCanAuthorizeWithCompletion:"), _block0)
}

// An integer you use to identify a right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/tag
func (r LARight) Tag() int {
	rv := objc.Send[int](r.ID, objc.Sel("tag"))
	return rv
}
func (r LARight) SetTag(value int) {
	objc.Send[struct{}](r.ID, objc.Sel("setTag:"), value)
}

// The current authorization state for a right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/state-swift.property
func (r LARight) State() LARightState {
	rv := objc.Send[LARightState](r.ID, objc.Sel("state"))
	return LARightState(rv)
}

// AuthorizeWithLocalizedReasonCompletionSync is a synchronous wrapper around [LARight.AuthorizeWithLocalizedReasonCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARight) AuthorizeWithLocalizedReasonCompletionSync(ctx context.Context, localizedReason string) error {
	done := make(chan error, 1)
	r.AuthorizeWithLocalizedReasonCompletion(localizedReason, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// AuthorizeWithLocalizedReasonInPresentationContextCompletionSync is a synchronous wrapper around [LARight.AuthorizeWithLocalizedReasonInPresentationContextCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARight) AuthorizeWithLocalizedReasonInPresentationContextCompletionSync(ctx context.Context, localizedReason string, presentationContext unsafe.Pointer) error {
	done := make(chan error, 1)
	r.AuthorizeWithLocalizedReasonInPresentationContextCompletion(localizedReason, presentationContext, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Deauthorize is a synchronous wrapper around [LARight.DeauthorizeWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARight) Deauthorize(ctx context.Context) error {
	done := make(chan struct{}, 1)
	r.DeauthorizeWithCompletion(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CheckCanAuthorize is a synchronous wrapper around [LARight.CheckCanAuthorizeWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (r LARight) CheckCanAuthorize(ctx context.Context) error {
	done := make(chan error, 1)
	r.CheckCanAuthorizeWithCompletion(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
