// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The [URLAuthenticationChallengeSender] protocol represents the interface that the sender of an authentication challenge must implement.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender
type NSURLAuthenticationChallengeSender interface {
	objectivec.IObject

	// Cancels a given authentication challenge.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/cancel(_:)
	CancelAuthenticationChallenge(challenge INSURLAuthenticationChallenge)

	// Attempt to continue downloading a request without providing a credential for a given challenge.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/continueWithoutCredential(for:)
	ContinueWithoutCredentialForAuthenticationChallenge(challenge INSURLAuthenticationChallenge)

	// Attempt to use a given credential for a given authentication challenge.
	//
	// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/use(_:for:)
	UseCredentialForAuthenticationChallenge(credential INSURLCredential, challenge INSURLAuthenticationChallenge)
}

// NSURLAuthenticationChallengeSenderObject wraps an existing Objective-C object that conforms to the NSURLAuthenticationChallengeSender protocol.
type NSURLAuthenticationChallengeSenderObject struct {
	objectivec.Object
}

func (o NSURLAuthenticationChallengeSenderObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSURLAuthenticationChallengeSenderObjectFromID constructs a [NSURLAuthenticationChallengeSenderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLAuthenticationChallengeSenderObjectFromID(id objc.ID) NSURLAuthenticationChallengeSenderObject {
	return NSURLAuthenticationChallengeSenderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Cancels a given authentication challenge.
//
// challenge: The authentication challenge to cancel.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/cancel(_:)
func (o NSURLAuthenticationChallengeSenderObject) CancelAuthenticationChallenge(challenge INSURLAuthenticationChallenge) {
	objc.Send[struct{}](o.ID, objc.Sel("cancelAuthenticationChallenge:"), challenge)
}

// Attempt to continue downloading a request without providing a credential
// for a given challenge.
//
// challenge: A challenge without authentication credentials.
//
// # Discussion
//
// This method has no effect if it is called with an authentication challenge
// that has already been handled.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/continueWithoutCredential(for:)
func (o NSURLAuthenticationChallengeSenderObject) ContinueWithoutCredentialForAuthenticationChallenge(challenge INSURLAuthenticationChallenge) {
	objc.Send[struct{}](o.ID, objc.Sel("continueWithoutCredentialForAuthenticationChallenge:"), challenge)
}

// Attempt to use a given credential for a given authentication challenge.
//
// credential: The credential to use for authentication.
//
// challenge: The challenge for which to use `credential`.
//
// # Discussion
//
// This method has no effect if it is called with an authentication challenge
// that has already been handled.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/use(_:for:)
func (o NSURLAuthenticationChallengeSenderObject) UseCredentialForAuthenticationChallenge(credential INSURLCredential, challenge INSURLAuthenticationChallenge) {
	objc.Send[struct{}](o.ID, objc.Sel("useCredential:forAuthenticationChallenge:"), credential, challenge)
}

// Causes the system-provided default behavior to be used.
//
// challenge: The challenge for which the default behavior should be used.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/performDefaultHandling(for:)
func (o NSURLAuthenticationChallengeSenderObject) PerformDefaultHandlingForAuthenticationChallenge(challenge INSURLAuthenticationChallenge) {
	objc.Send[struct{}](o.ID, objc.Sel("performDefaultHandlingForAuthenticationChallenge:"), challenge)
}

// Rejects the currently supplied protection space.
//
// challenge: The challenge that should be rejected.
//
// See: https://developer.apple.com/documentation/Foundation/URLAuthenticationChallengeSender/rejectProtectionSpaceAndContinue(with:)
func (o NSURLAuthenticationChallengeSenderObject) RejectProtectionSpaceAndContinueWithChallenge(challenge INSURLAuthenticationChallenge) {
	objc.Send[struct{}](o.ID, objc.Sel("rejectProtectionSpaceAndContinueWithChallenge:"), challenge)
}
