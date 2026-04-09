// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
)

// The class instance for the [LAContext] class.
var (
	_LAContextClass     LAContextClass
	_LAContextClassOnce sync.Once
)

func getLAContextClass() LAContextClass {
	_LAContextClassOnce.Do(func() {
		_LAContextClass = LAContextClass{class: objc.GetClass("LAContext")}
	})
	return _LAContextClass
}

// GetLAContextClass returns the class object for LAContext.
func GetLAContextClass() LAContextClass {
	return getLAContextClass()
}

type LAContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAContextClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAContextClass) Alloc() LAContext {
	rv := objc.Send[LAContext](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// A mechanism for evaluating authentication policies and access controls.
//
// # Overview
//
// You use an authentication context to evaluate the user’s identity, either
// with biometrics like Touch ID or Face ID, or by supplying the device
// passcode. The context handles user interaction, and also interfaces to the
// Secure Enclave, the underlying hardware element that manages biometric
// data. You create and configure the context, and ask it to carry out the
// authentication. You then receive an asynchronous callback, which provides
// an indication of authentication success or failure, and an error instance
// that explains the reason for a failure, if any.
//
// # Checking availability
//
//   - [LAContext.CanEvaluatePolicyError]: Assesses whether authentication can proceed for a given policy.
//   - [LAContext.BiometryType]: The type of biometric authentication supported by the device.
//
// # Evaluating authentication policies
//
//   - [LAContext.EvaluatePolicyLocalizedReasonReply]: Evaluates the specified policy.
//   - [LAContext.EvaluatedPolicyDomainState]: The current state of the evaluated policy domain.
//
// # Evaluating access controls
//
//   - [LAContext.EvaluateAccessControlOperationLocalizedReasonReply]: Evaluates an access control for a given operation.
//   - [LAContext.InteractionNotAllowed]: A Boolean value indicating whether authentication can be interactive.
//   - [LAContext.SetInteractionNotAllowed]
//
// # Customizing authentication prompts
//
//   - [LAContext.LocalizedReason]: The localized explanation for authentication shown in the dialog presented to the user.
//   - [LAContext.SetLocalizedReason]
//   - [LAContext.LocalizedFallbackTitle]: The localized title for the fallback button in the dialog presented to the user during authentication.
//   - [LAContext.SetLocalizedFallbackTitle]
//   - [LAContext.LocalizedCancelTitle]: The localized title for the cancel button in the dialog presented to the user during authentication.
//   - [LAContext.SetLocalizedCancelTitle]
//
// # Reusing device unlock state
//
//   - [LAContext.TouchIDAuthenticationAllowableReuseDuration]: The duration for which Touch ID authentication reuse is allowable.
//   - [LAContext.SetTouchIDAuthenticationAllowableReuseDuration]
//   - [LAContext.LATouchIDAuthenticationMaximumAllowableReuseDuration]: The maximum allowable reuse duration.
//
// # Managing credentials
//
//   - [LAContext.SetCredentialType]: Sets an application-provided credential to be used when evaluating authentication.
//   - [LAContext.IsCredentialSet]: Returns a Boolean value indicating whether the specified credential type is set.
//
// # Invalidating the authentication context
//
//   - [LAContext.Invalidate]: Invalidates the authentication context.
//
// # Instance Properties
//
//   - [LAContext.DomainState]: Contains authentication domain state.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext
type LAContext struct {
	objectivec.Object
}

// LAContextFromID constructs a [LAContext] from an objc.ID.
//
// A mechanism for evaluating authentication policies and access controls.
func LAContextFromID(id objc.ID) LAContext {
	return LAContext{objectivec.Object{ID: id}}
}

// NOTE: LAContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAContext] class.
//
// # Checking availability
//
//   - [ILAContext.CanEvaluatePolicyError]: Assesses whether authentication can proceed for a given policy.
//   - [ILAContext.BiometryType]: The type of biometric authentication supported by the device.
//
// # Evaluating authentication policies
//
//   - [ILAContext.EvaluatePolicyLocalizedReasonReply]: Evaluates the specified policy.
//   - [ILAContext.EvaluatedPolicyDomainState]: The current state of the evaluated policy domain.
//
// # Evaluating access controls
//
//   - [ILAContext.EvaluateAccessControlOperationLocalizedReasonReply]: Evaluates an access control for a given operation.
//   - [ILAContext.InteractionNotAllowed]: A Boolean value indicating whether authentication can be interactive.
//   - [ILAContext.SetInteractionNotAllowed]
//
// # Customizing authentication prompts
//
//   - [ILAContext.LocalizedReason]: The localized explanation for authentication shown in the dialog presented to the user.
//   - [ILAContext.SetLocalizedReason]
//   - [ILAContext.LocalizedFallbackTitle]: The localized title for the fallback button in the dialog presented to the user during authentication.
//   - [ILAContext.SetLocalizedFallbackTitle]
//   - [ILAContext.LocalizedCancelTitle]: The localized title for the cancel button in the dialog presented to the user during authentication.
//   - [ILAContext.SetLocalizedCancelTitle]
//
// # Reusing device unlock state
//
//   - [ILAContext.TouchIDAuthenticationAllowableReuseDuration]: The duration for which Touch ID authentication reuse is allowable.
//   - [ILAContext.SetTouchIDAuthenticationAllowableReuseDuration]
//   - [ILAContext.LATouchIDAuthenticationMaximumAllowableReuseDuration]: The maximum allowable reuse duration.
//
// # Managing credentials
//
//   - [ILAContext.SetCredentialType]: Sets an application-provided credential to be used when evaluating authentication.
//   - [ILAContext.IsCredentialSet]: Returns a Boolean value indicating whether the specified credential type is set.
//
// # Invalidating the authentication context
//
//   - [ILAContext.Invalidate]: Invalidates the authentication context.
//
// # Instance Properties
//
//   - [ILAContext.DomainState]: Contains authentication domain state.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext
type ILAContext interface {
	objectivec.IObject

	// Topic: Checking availability

	// Assesses whether authentication can proceed for a given policy.
	CanEvaluatePolicyError(policy LAPolicy) (bool, error)
	// The type of biometric authentication supported by the device.
	BiometryType() LABiometryType

	// Topic: Evaluating authentication policies

	// Evaluates the specified policy.
	EvaluatePolicyLocalizedReasonReply(policy LAPolicy, localizedReason string, reply BoolErrorHandler)
	// The current state of the evaluated policy domain.
	EvaluatedPolicyDomainState() foundation.INSData

	// Topic: Evaluating access controls

	// Evaluates an access control for a given operation.
	EvaluateAccessControlOperationLocalizedReasonReply(accessControl security.SecAccessControlRef, operation LAAccessControlOperation, localizedReason string, reply BoolErrorHandler)
	// A Boolean value indicating whether authentication can be interactive.
	InteractionNotAllowed() bool
	SetInteractionNotAllowed(value bool)

	// Topic: Customizing authentication prompts

	// The localized explanation for authentication shown in the dialog presented to the user.
	LocalizedReason() string
	SetLocalizedReason(value string)
	// The localized title for the fallback button in the dialog presented to the user during authentication.
	LocalizedFallbackTitle() string
	SetLocalizedFallbackTitle(value string)
	// The localized title for the cancel button in the dialog presented to the user during authentication.
	LocalizedCancelTitle() string
	SetLocalizedCancelTitle(value string)

	// Topic: Reusing device unlock state

	// The duration for which Touch ID authentication reuse is allowable.
	TouchIDAuthenticationAllowableReuseDuration() float64
	SetTouchIDAuthenticationAllowableReuseDuration(value float64)
	// The maximum allowable reuse duration.
	LATouchIDAuthenticationMaximumAllowableReuseDuration() float64

	// Topic: Managing credentials

	// Sets an application-provided credential to be used when evaluating authentication.
	SetCredentialType(credential foundation.INSData, type_ LACredentialType) bool
	// Returns a Boolean value indicating whether the specified credential type is set.
	IsCredentialSet(type_ LACredentialType) bool

	// Topic: Invalidating the authentication context

	// Invalidates the authentication context.
	Invalidate()

	// Topic: Instance Properties

	// Contains authentication domain state.
	DomainState() ILADomainState
}

// Init initializes the instance.
func (c LAContext) Init() LAContext {
	rv := objc.Send[LAContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c LAContext) Autorelease() LAContext {
	rv := objc.Send[LAContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAContext creates a new LAContext instance.
func NewLAContext() LAContext {
	class := getLAContextClass()
	rv := objc.Send[LAContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Assesses whether authentication can proceed for a given policy.
//
// policy: The policy to evaluate. For possible values, see [LAPolicy].
//
// error: If the method fails, it uses this parameter to return an error detailing
// what went wrong. See [LAError.Code] for possible error codes.
//
// Specify `nil` for this parameter to ignore any errors.
//
// # Return Value
//
// true if the policy can be evaluated, otherwise false.
//
// # Discussion
//
// Some policies impose requirements that must be met before authentication
// can proceed. For example, a policy that requires biometrics can’t
// authenticate if Touch ID or Face ID is disabled. This method tests all the
// prerequisites for a given policy.
//
// Don’t store the return value from this method because it might change as
// a result of changes in the system. For example, a user might disable Touch
// ID after you call this method.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/canEvaluatePolicy(_:error:)
//
// [LAPolicy]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy
// [LAError.Code]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code
func (c LAContext) CanEvaluatePolicyError(policy LAPolicy) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("canEvaluatePolicy:error:"), policy, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("canEvaluatePolicy:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Evaluates the specified policy.
//
// policy: The policy to evaluate. For possible values, see [LAPolicy].
//
// localizedReason: The app-provided reason for requesting authentication, which displays in
// the authentication dialog presented to the user.
//
// reply: A closure that is executed when policy evaluation finishes. This is
// evaluated on a private queue internal to the framework in an unspecified
// threading context. You must not call [CanEvaluatePolicyError] in this
// block, because doing so could lead to deadlock.
//
// success: true if policy evaluation succeeded, otherwise false. error: `nil`
// if policy evaluation succeeded, an error object that should be presented to
// the user otherwise. See [LAError.Code] for possible error codes
//
// # Discussion
//
// This method asynchronously evaluates an authentication policy. Evaluating a
// policy may involve prompting the user for various kinds of interaction or
// authentication. The actual behavior is dependent on the evaluated policy
// and the device type. The behavior can also be affected by installed
// configuration profiles.
//
// In the localized string you present to the user in the authentication
// dialog, provide a clear reason for the authentication request, and describe
// the resulting action. Make the message short and clear, and provide it in
// the user’s language. Don’t include the app name, which already appears
// in the authentication dialog (in macOS, in the title of the dialog; in iOS,
// in the subtitle).
//
// Don’t assume that a previous successful policy evaluation means that
// future evaluations will also succeed. Policy evaluation can fail for
// various reasons, including cancellation by the user or the system.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/evaluatePolicy(_:localizedReason:reply:)
//
// [LAPolicy]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy
// [LAError.Code]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code
func (c LAContext) EvaluatePolicyLocalizedReasonReply(policy LAPolicy, localizedReason string, reply BoolErrorHandler) {
	_block2, _ := NewBoolErrorBlock(reply)
	objc.Send[objc.ID](c.ID, objc.Sel("evaluatePolicy:localizedReason:reply:"), policy, objc.String(localizedReason), _block2)
}

// Evaluates an access control for a given operation.
//
// accessControl: The access control to be evaluated.
//
// operation: The operation for the access control to be evaluated. For possible values,
// see [LAAccessControlOperation].
//
// localizedReason: The app-provided reason for requesting authentication, which displays in
// the authentication dialog presented to the user.
//
// reply: A block that is executed when access control evaluation finishes. This
// block is evaluated on a private queue internal to the framework in an
// unspecified threading context.
//
// success: `true` if policy evaluation succeeded, otherwise `false`. error:
// `nil` if policy evaluation succeeded, an error object that should be
// presented to the user otherwise. See [LAError.Code] for possible error
// codes
//
// # Discussion
//
// This method asynchronously evaluates an access control. Evaluating an
// access control may involve prompting the user for various kinds of
// interaction or authentication. The actual behavior is dependent on the
// access control and device type. It can also be affected by installed
// configuration profiles.
//
// The localized string you present to the user should provide a clear reason
// for why you are requesting they authenticate themselves, and what action
// you will be taking based on that authentication. This string should be
// provided in the user’s current language and should be short and clear. It
// should not contain the app name, because that appears elsewhere in the
// authentication dialog. In macOS this appears in the dialog title, and in
// iOS this appears in the dialog subtitle.
//
// You should not assume that a previous successful evaluation of an access
// control necessarily leads to a subsequent successful evaluation. Access
// control evaluation can fail for various reasons, including cancelation by
// the user or the system.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/evaluateAccessControl(_:operation:localizedReason:reply:)
//
// [LAAccessControlOperation]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation
// [LAError.Code]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code
func (c LAContext) EvaluateAccessControlOperationLocalizedReasonReply(accessControl security.SecAccessControlRef, operation LAAccessControlOperation, localizedReason string, reply BoolErrorHandler) {
	_block3, _ := NewBoolErrorBlock(reply)
	objc.Send[objc.ID](c.ID, objc.Sel("evaluateAccessControl:operation:localizedReason:reply:"), accessControl, operation, objc.String(localizedReason), _block3)
}

// Sets an application-provided credential to be used when evaluating
// authentication.
//
// credential: The credential to be used when evaluating the authentication context.
//
// Setting this parameter to `nil` removes any existing credential of the
// specified type.
//
// type: The type of the specified credential. For possible values, see
// [LACredentialType].
//
// # Return Value
//
// true if the credential was set, otherwise false.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/setCredential(_:type:)
//
// [LACredentialType]: https://developer.apple.com/documentation/LocalAuthentication/LACredentialType
func (c LAContext) SetCredentialType(credential foundation.INSData, type_ LACredentialType) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("setCredential:type:"), credential, type_)
	return rv
}

// Returns a Boolean value indicating whether the specified credential type is
// set.
//
// type: The type of the credential. For possible values, see [LACredentialType]
//
// # Return Value
//
// true if the credential is set, otherwise false.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/isCredentialSet(_:)
//
// [LACredentialType]: https://developer.apple.com/documentation/LocalAuthentication/LACredentialType
func (c LAContext) IsCredentialSet(type_ LACredentialType) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCredentialSet:"), type_)
	return rv
}

// Invalidates the authentication context.
//
// # Discussion
//
// Calling this method stops any pending policy evaluations, causing them to
// fail with the [LAErrorAppCancel] error code. Once an authentication context
// has been invalidated, it cannot be used for policy evaluation. Invalidating
// a context that has been already invalidated has no effect.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/invalidate()
func (c LAContext) Invalidate() {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidate"))
}

// The type of biometric authentication supported by the device.
//
// # Discussion
//
// Use the value of this property to ensure that any authentication-related
// user prompts you create match the biometric capabilities of the device. For
// example, if the value of this property is [LABiometryTypeFaceID], don’t
// refer to Touch ID in an authentication prompt.
//
// This property is set only after you call the [CanEvaluatePolicyError]
// method, and is set no matter what the call returns. The default value is
// [LABiometryTypeNone].
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/biometryType
func (c LAContext) BiometryType() LABiometryType {
	rv := objc.Send[LABiometryType](c.ID, objc.Sel("biometryType"))
	return LABiometryType(rv)
}

// The current state of the evaluated policy domain.
//
// # Discussion
//
// The value of this property is non-`nil` when the [CanEvaluatePolicyError]
// method succeeds for a biometric policy or the person successfully
// authenticates using biometrics, following a call to
// [EvaluatePolicyLocalizedReasonReply]. Otherwise, its value is `nil`.
//
// Compare the values you get from successive calls to this property to
// determine whether the authorized database changed. However, the value you
// get doesn’t describe the nature of a change; it only lets you detect if a
// change happens.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/evaluatedPolicyDomainState
func (c LAContext) EvaluatedPolicyDomainState() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("evaluatedPolicyDomainState"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A Boolean value indicating whether authentication can be interactive.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/interactionNotAllowed
func (c LAContext) InteractionNotAllowed() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("interactionNotAllowed"))
	return rv
}
func (c LAContext) SetInteractionNotAllowed(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setInteractionNotAllowed:"), value)
}

// The localized explanation for authentication shown in the dialog presented
// to the user.
//
// # Discussion
//
// This property is overwritten if an authentication reason is provided in
// [EvaluatePolicyLocalizedReasonReply].
//
// The localized string you present to the user should provide a clear reason
// for why you are requesting they authenticate themselves, and what action
// you will be taking based on that authentication. This string should be
// provided in the user’s current language and should be short and clear. It
// should not contain the app name, because that appears elsewhere in the
// authentication dialog. In macOS this appears in the dialog title, and in
// iOS this appears in the dialog subtitle.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedReason
func (c LAContext) LocalizedReason() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedReason"))
	return foundation.NSStringFromID(rv).String()
}
func (c LAContext) SetLocalizedReason(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLocalizedReason:"), objc.String(value))
}

// The localized title for the fallback button in the dialog presented to the
// user during authentication.
//
// # Discussion
//
// The system presents a fallback button when biometric authentication
// fails—for example, because the system doesn’t recognize the presented
// finger, or after several failed attempts to recognize the user’s face.
// Tapping the button lets the user revert to authentication using the device
// passcode or password instead.
//
// Use the [LocalizedFallbackTitle] property to set a title for the fallback
// button. When the property is `nil`—as it is by default—the user sees an
// appropriate default title, like “Use Passcode”. Otherwise, provide a
// localized string that’s short and clear.
//
// To eliminate the fallback option, set the fallback title to an empty
// string. This hides the button from the interface.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedFallbackTitle
func (c LAContext) LocalizedFallbackTitle() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedFallbackTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (c LAContext) SetLocalizedFallbackTitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLocalizedFallbackTitle:"), objc.String(value))
}

// The localized title for the cancel button in the dialog presented to the
// user during authentication.
//
// # Discussion
//
// The system presents a cancel button during biometric authentication to let
// the user abort the authentication attempt. The button appears every time
// the system asks the user to present a finger registered with Touch ID. For
// Face ID, the button only appears if authentication fails and the user is
// prompted to try again. Either way, the user can stop trying to authenticate
// by tapping the button.
//
// Use the [LocalizedCancelTitle] property to choose a title for the cancel
// button. If you set the property to `nil`—as it is by default—or assign
// an empty string, the system uses an appropriate default title, like
// “Cancel”. Otherwise, provide a localized string that’s short and
// clear.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedCancelTitle
func (c LAContext) LocalizedCancelTitle() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedCancelTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (c LAContext) SetLocalizedCancelTitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLocalizedCancelTitle:"), objc.String(value))
}

// The duration for which Touch ID authentication reuse is allowable.
//
// # Discussion
//
// If the user unlocks the device using Touch ID within the specified time
// interval, then authentication for the receiver succeeds automatically,
// without prompting the user for Touch ID. This bypasses a scenario where the
// user unlocks the device and then is almost immediately prompted for another
// fingerprint.
//
// The default value is `0`, meaning that Touch ID authentication isn’t
// reused.
//
// The maximum allowable duration for Touch ID authentication reuse is
// specified by the [LATouchIDAuthenticationMaximumAllowableReuseDuration]
// constant. You cannot specify a longer duration by setting this property to
// a value greater than this constant.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/touchIDAuthenticationAllowableReuseDuration
func (c LAContext) TouchIDAuthenticationAllowableReuseDuration() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("touchIDAuthenticationAllowableReuseDuration"))
	return rv
}
func (c LAContext) SetTouchIDAuthenticationAllowableReuseDuration(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setTouchIDAuthenticationAllowableReuseDuration:"), value)
}

// The maximum allowable reuse duration.
//
// See: https://developer.apple.com/documentation/localauthentication/latouchidauthenticationmaximumallowablereuseduration
func (c LAContext) LATouchIDAuthenticationMaximumAllowableReuseDuration() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("LATouchIDAuthenticationMaximumAllowableReuseDuration"))
	return rv
}

// Contains authentication domain state.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAContext/domainState
func (c LAContext) DomainState() ILADomainState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("domainState"))
	return LADomainStateFromID(objc.ID(rv))
}

// EvaluatePolicyLocalizedReasonReplySync is a synchronous wrapper around [LAContext.EvaluatePolicyLocalizedReasonReply].
// It blocks until the completion handler fires or the context is cancelled.
func (c LAContext) EvaluatePolicyLocalizedReasonReplySync(ctx context.Context, policy LAPolicy, localizedReason string) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	c.EvaluatePolicyLocalizedReasonReply(policy, localizedReason, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// EvaluateAccessControlOperationLocalizedReasonReplySync is a synchronous wrapper around [LAContext.EvaluateAccessControlOperationLocalizedReasonReply].
// It blocks until the completion handler fires or the context is cancelled.
func (c LAContext) EvaluateAccessControlOperationLocalizedReasonReplySync(ctx context.Context, accessControl security.SecAccessControlRef, operation LAAccessControlOperation, localizedReason string) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	c.EvaluateAccessControlOperationLocalizedReasonReply(accessControl, operation, localizedReason, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}
