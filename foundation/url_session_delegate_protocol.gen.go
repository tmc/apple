// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A protocol that defines methods that URL session instances call on their delegates to handle session-level events, like session life cycle changes.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDelegate
type NSURLSessionDelegate interface {
	objectivec.IObject
}



// NSURLSessionDelegateObject wraps an existing Objective-C object that conforms to the NSURLSessionDelegate protocol.
type NSURLSessionDelegateObject struct {
	objectivec.Object
}
func (o NSURLSessionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSURLSessionDelegateObjectFromID constructs a [NSURLSessionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLSessionDelegateObjectFromID(id objc.ID) NSURLSessionDelegateObject {
	return NSURLSessionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the URL session that the session has been invalidated.
//
// session: The session object that was invalidated.
//
// error: The error that caused invalidation, or `nil` if the invalidation was
// explicit.
//
// # Discussion
// 
// If you invalidate a session by calling its [FinishTasksAndInvalidate]
// method, the session waits until after the final task in the session
// finishes or fails before calling this delegate method. If you call the
// [InvalidateAndCancel] method, the session calls this delegate method
// immediately.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDelegate/urlSession(_:didBecomeInvalidWithError:)

func (o NSURLSessionDelegateObject) URLSessionDidBecomeInvalidWithError(session INSURLSession, error_ INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:didBecomeInvalidWithError:"), session, error_)
	}

// Tells the delegate that all messages enqueued for a session have been
// delivered.
//
// session: The session that no longer has any outstanding requests.
//
// # Discussion
// 
// In iOS, when a background transfer completes or requires credentials, if
// your app is no longer running, your app is automatically relaunched in the
// background, and the app’s [UIApplicationDelegate] is sent an
// [application(_:handleEventsForBackgroundURLSession:completionHandler:)]
// message. This call contains the identifier of the session that caused your
// app to be launched. You should then store that completion handler before
// creating a background configuration object with the same identifier, and
// creating a session with that configuration. The newly created session is
// automatically reassociated with ongoing background activity.
// 
// When your app later receives a
// [URLSessionDidFinishEventsForBackgroundURLSession] message, this indicates
// that all messages previously enqueued for this session have been delivered,
// and that it is now safe to invoke the previously stored completion handler
// or to begin any internal updates that may result in invoking the completion
// handler.
//
// [application(_:handleEventsForBackgroundURLSession:completionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:handleEventsForBackgroundURLSession:completionHandler:)
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDelegate/urlSessionDidFinishEvents(forBackgroundURLSession:)

func (o NSURLSessionDelegateObject) URLSessionDidFinishEventsForBackgroundURLSession(session INSURLSession) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSessionDidFinishEventsForBackgroundURLSession:"), session)
	}

// Requests credentials from the delegate in response to a session-level
// authentication request from the remote server.
//
// session: The session containing the task that requested authentication.
//
// challenge: An object that contains the request for authentication.
//
// completionHandler: A handler that your delegate method must call. This completion handler
// takes the following parameters::
// 
// - `disposition`—One of several constants that describes how the challenge
// should be handled. - `credential`—The credential that should be used for
// authentication if disposition is [NSURLSessionAuthChallengeUseCredential],
// otherwise [NULL].
//
// # Discussion
// 
// This method is called in two situations:
// 
// - When a remote server asks for client certificates or Windows NT LAN
// Manager (NTLM) authentication, to allow your app to provide appropriate
// credentials - When a session first establishes a connection to a remote
// server that uses SSL or TLS, to allow your app to verify the server’s
// certificate chain
// 
// If you do not implement this method, the session calls its delegate’s
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method instead.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionDelegate/urlSession(_:didReceive:completionHandler:)

func (o NSURLSessionDelegateObject) URLSessionDidReceiveChallengeCompletionHandler(session INSURLSession, challenge INSURLAuthenticationChallenge, completionHandler URLSessionAuthChallengeDispositionURLCredentialHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("URLSession:didReceiveChallenge:completionHandler:"), session, challenge, completionHandler)
	}





// NSURLSessionDelegateConfig holds optional typed callbacks for [NSURLSessionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate
type NSURLSessionDelegateConfig struct {

	// Other Methods
	// URLSessionDidBecomeInvalidWithError — Tells the URL session that the session has been invalidated.
	URLSessionDidBecomeInvalidWithError func(session NSURLSession, error_ objectivec.Object)
	// URLSessionDidFinishEventsForBackgroundURLSession — Tells the delegate that all messages enqueued for a session have been delivered.
	URLSessionDidFinishEventsForBackgroundURLSession func(session NSURLSession)
	// URLSessionDidReceiveChallengeCompletionHandler — Requests credentials from the delegate in response to a session-level authentication request from the remote server.
	URLSessionDidReceiveChallengeCompletionHandler func(session NSURLSession, challenge NSURLAuthenticationChallenge, completionHandler objc.ID)
}

// NewNSURLSessionDelegate creates an Objective-C object implementing the [NSURLSessionDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSURLSessionDelegateObject] satisfies the [NSURLSessionDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlsessiondelegate
func NewNSURLSessionDelegate(config NSURLSessionDelegateConfig) NSURLSessionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSURLSessionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.URLSessionDidBecomeInvalidWithError != nil {
		fn := config.URLSessionDidBecomeInvalidWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:didBecomeInvalidWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, error_ID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				error_ := objectivec.ObjectFromID(error_ID)
				fn(session, error_)
			},
		})
	}

	if config.URLSessionDidFinishEventsForBackgroundURLSession != nil {
		fn := config.URLSessionDidFinishEventsForBackgroundURLSession
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSessionDidFinishEventsForBackgroundURLSession:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				fn(session)
			},
		})
	}

	if config.URLSessionDidReceiveChallengeCompletionHandler != nil {
		fn := config.URLSessionDidReceiveChallengeCompletionHandler
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("URLSession:didReceiveChallenge:completionHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sessionID objc.ID, challengeID objc.ID, completionHandlerID objc.ID) {
				session := NSURLSessionFromID(sessionID)
				challenge := NSURLAuthenticationChallengeFromID(challengeID)
				completionHandler := completionHandlerID
				fn(session, challenge, completionHandler)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSURLSessionDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSURLSessionDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSURLSessionDelegateObjectFromID(instance)
}





