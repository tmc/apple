// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that delegates of a URL connection implement to receive status about and provide feedback to the connection object.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate
type NSURLConnectionDelegate interface {
	objectivec.IObject
}

// NSURLConnectionDelegateObject wraps an existing Objective-C object that conforms to the NSURLConnectionDelegate protocol.
type NSURLConnectionDelegateObject struct {
	objectivec.Object
}

func (o NSURLConnectionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSURLConnectionDelegateObjectFromID constructs a [NSURLConnectionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLConnectionDelegateObjectFromID(id objc.ID) NSURLConnectionDelegateObject {
	return NSURLConnectionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the connection will send a request for an
// authentication challenge.
//
// connection: The connection sending the message.
//
// challenge: The authentication challenge for which a request is being sent.
//
// # Discussion
//
// This method allows the delegate to make an informed decision about
// connection authentication at once. If the delegate implements this method,
// it has no need to implement
// [connection(_:canAuthenticateAgainstProtectionSpace:)] or
// [connection(_:didReceive:)]. In fact, those other methods are not invoked
// (except on older operating systems, where applicable).
//
// In this method,you invoke one of the challenge-responder methods
// ([NSURLAuthenticationChallengeSender] protocol):
//
// - [UseCredentialForAuthenticationChallenge] -
// [ContinueWithoutCredentialForAuthenticationChallenge] -
// [CancelAuthenticationChallenge] -
// [PerformDefaultHandlingForAuthenticationChallenge] -
// [RejectProtectionSpaceAndContinueWithChallenge]
//
// You might also want to analyze `challenge` for the authentication scheme
// and the proposed credential before calling a
// [NSURLAuthenticationChallengeSender] method. You should never assume that a
// proposed credential is present. You can either create your own credential
// and respond with that, or you can send the proposed credential back.
// (Because this object is immutable, if you want to change it you must copy
// it and then modify the copy.)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:willSendRequestFor:)
//
// [connection(_:canAuthenticateAgainstProtectionSpace:)]: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:canAuthenticateAgainstProtectionSpace:)
// [connection(_:didReceive:)]: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:didReceive:)
func (o NSURLConnectionDelegateObject) ConnectionWillSendRequestForAuthenticationChallenge(connection INSURLConnection, challenge INSURLAuthenticationChallenge) {
	objc.Send[struct{}](o.ID, objc.Sel("connection:willSendRequestForAuthenticationChallenge:"), connection, challenge)
}

// Sent to determine whether the URL loader should use the credential storage
// for authenticating the connection.
//
// connection: The connection sending the message.
//
// # Discussion
//
// This method is called before any attempt to authenticate is made.
//
// If you return false, the connection does not consult the credential storage
// automatically, and does not store credentials. However, in your
// connection:didReceiveAuthenticationChallenge: method, you can consult the
// credential storage yourself and store credentials yourself, as needed.
//
// Not implementing this method is the same as returning true.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connectionShouldUseCredentialStorage(_:)
func (o NSURLConnectionDelegateObject) ConnectionShouldUseCredentialStorage(connection INSURLConnection) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("connectionShouldUseCredentialStorage:"), connection)
	return rv
}

// Sent when a connection fails to load its request successfully.
//
// connection: The connection sending the message.
//
// error: An error object containing details of why the connection failed to load the
// request successfully.
//
// # Discussion
//
// Once the delegate receives this message, it will receive no further
// messages for `connection`.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:didFailWithError:)
func (o NSURLConnectionDelegateObject) ConnectionDidFailWithError(connection INSURLConnection, error_ INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("connection:didFailWithError:"), connection, error_)
}

// NSURLConnectionDelegateConfig holds optional typed callbacks for [NSURLConnectionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate
type NSURLConnectionDelegateConfig struct {

	// Connection Authentication
	// ConnectionShouldUseCredentialStorage — Sent to determine whether the URL loader should use the credential storage for authenticating the connection.
	ConnectionShouldUseCredentialStorage func(connection NSURLConnection) bool

	// Connection Completion
	// ConnectionDidFailWithError — Sent when a connection fails to load its request successfully.
	ConnectionDidFailWithError func(connection NSURLConnection, error_ objectivec.Object)

	// Other Methods
	// ConnectionWillSendRequestForAuthenticationChallenge — Tells the delegate that the connection will send a request for an authentication challenge.
	ConnectionWillSendRequestForAuthenticationChallenge func(connection NSURLConnection, challenge NSURLAuthenticationChallenge)
}

// NewNSURLConnectionDelegate creates an Objective-C object implementing the [NSURLConnectionDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSURLConnectionDelegateObject] satisfies the [NSURLConnectionDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlconnectiondelegate
func NewNSURLConnectionDelegate(config NSURLConnectionDelegateConfig) NSURLConnectionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSURLConnectionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ConnectionWillSendRequestForAuthenticationChallenge != nil {
		fn := config.ConnectionWillSendRequestForAuthenticationChallenge
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:willSendRequestForAuthenticationChallenge:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, challengeID objc.ID) {
				connection := NSURLConnectionFromID(connectionID)
				challenge := NSURLAuthenticationChallengeFromID(challengeID)
				fn(connection, challenge)
			},
		})
	}

	if config.ConnectionShouldUseCredentialStorage != nil {
		fn := config.ConnectionShouldUseCredentialStorage
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connectionShouldUseCredentialStorage:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID) bool {
				connection := NSURLConnectionFromID(connectionID)
				return fn(connection)
			},
		})
	}

	if config.ConnectionDidFailWithError != nil {
		fn := config.ConnectionDidFailWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:didFailWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, error_ID objc.ID) {
				connection := NSURLConnectionFromID(connectionID)
				error_ := objectivec.ObjectFromID(error_ID)
				fn(connection, error_)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSURLConnectionDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSURLConnectionDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSURLConnectionDelegateObjectFromID(instance)
}
