// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that delegates of a URL connection created with Newsstand Kit implement to receive data associated with a download.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDownloadDelegate
type NSURLConnectionDownloadDelegate interface {
	objectivec.IObject
	NSURLConnectionDelegate

	// Sent to the delegate when the URL connection has successfully downloaded the URL asset to a destination file.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDownloadDelegate/connectionDidFinishDownloading(_:destinationURL:)
	ConnectionDidFinishDownloadingDestinationURL(connection INSURLConnection, destinationURL INSURL)
}

// NSURLConnectionDownloadDelegateObject wraps an existing Objective-C object that conforms to the NSURLConnectionDownloadDelegate protocol.
type NSURLConnectionDownloadDelegateObject struct {
	objectivec.Object
}
func (o NSURLConnectionDownloadDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSURLConnectionDownloadDelegateObjectFromID constructs a [NSURLConnectionDownloadDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSURLConnectionDownloadDelegateObjectFromID(id objc.ID) NSURLConnectionDownloadDelegateObject {
	return NSURLConnectionDownloadDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sent to the delegate when the URL connection has successfully downloaded
// the URL asset to a destination file.
//
// connection: The URL connection object that downloaded the asset.
//
// destinationURL: A file URL specifying a destination in the file system. For iOS
// applications, this is a location in the application sandbox.
//
// # Discussion
// 
// This method will be called once after a successful download. The file
// downloaded to `destinationURL` is guaranteed to exist there only for the
// duration of this method implementation; the delegate should copy or move
// the file to a more persistent and appropriate location.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDownloadDelegate/connectionDidFinishDownloading(_:destinationURL:)

func (o NSURLConnectionDownloadDelegateObject) ConnectionDidFinishDownloadingDestinationURL(connection INSURLConnection, destinationURL INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connectionDidFinishDownloading:destinationURL:"), connection, destinationURL)
	}

// Sent to the delegate to deliver progress information for a download of a
// URL asset to a destination file.
//
// connection: The URL connection object downloading the asset.
//
// bytesWritten: The number of bytes written since the last call of this method.
//
// totalBytesWritten: The total number of bytes of the downloading asset that have been written
// to the file.
//
// expectedTotalBytes: The total number of bytes of the URL asset once it is completely downloaded
// and written to a file. This parameter can be zero if the total number of
// bytes is not known.
//
// # Discussion
// 
// This method is invoked repeatedly during the download of a URL asset to the
// destination file. The delegate typically uses the values of the three
// “bytes” parameters to update a progress indicator in the
// application’s user interface.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDownloadDelegate/connection(_:didWriteData:totalBytesWritten:expectedTotalBytes:)

func (o NSURLConnectionDownloadDelegateObject) ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection INSURLConnection, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connection:didWriteData:totalBytesWritten:expectedTotalBytes:"), connection, bytesWritten, totalBytesWritten, expectedTotalBytes)
	}

// Sent to the delegate when an URL connection resumes downloading a URL asset
// that was earlier suspended.
//
// connection: The URL connection object downloading the asset.
//
// totalBytesWritten: The total number of bytes of the downloading asset that have been written
// to the destination file.
//
// expectedTotalBytes: The total number of bytes of the URL asset once it is completely downloaded
// and written to a file.
//
// # Discussion
// 
// This method is invoked once a suspended download of a URL asset resumes
// downloading. In response, the delegate can display a progress indicator,
// setting the initial value of the indicator to where it was when downloading
// was suspended. After the URL-connection object sends this message, it sends
// one or more [ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes] to
// the delegate until the download concludes.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDownloadDelegate/connectionDidResumeDownloading(_:totalBytesWritten:expectedTotalBytes:)

func (o NSURLConnectionDownloadDelegateObject) ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection INSURLConnection, totalBytesWritten int64, expectedTotalBytes int64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connectionDidResumeDownloading:totalBytesWritten:expectedTotalBytes:"), connection, totalBytesWritten, expectedTotalBytes)
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
// [connection(_:canAuthenticateAgainstProtectionSpace:)]: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:canAuthenticateAgainstProtectionSpace:)
// [connection(_:didReceive:)]: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:didReceive:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connection(_:willSendRequestFor:)

func (o NSURLConnectionDownloadDelegateObject) ConnectionWillSendRequestForAuthenticationChallenge(connection INSURLConnection, challenge INSURLAuthenticationChallenge) {
	
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
// If you return [false], the connection does not consult the credential
// storage automatically, and does not store credentials. However, in your
// connection:didReceiveAuthenticationChallenge: method, you can consult the
// credential storage yourself and store credentials yourself, as needed.
// 
// Not implementing this method is the same as returning [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURLConnectionDelegate/connectionShouldUseCredentialStorage(_:)

func (o NSURLConnectionDownloadDelegateObject) ConnectionShouldUseCredentialStorage(connection INSURLConnection) bool {
	
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

func (o NSURLConnectionDownloadDelegateObject) ConnectionDidFailWithError(connection INSURLConnection, error_ INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("connection:didFailWithError:"), connection, error_)
	}

// NSURLConnectionDownloadDelegateConfig holds optional typed callbacks for [NSURLConnectionDownloadDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate
type NSURLConnectionDownloadDelegateConfig struct {

	// Managing Downloads of URL Assets
	// ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes — Sent to the delegate to deliver progress information for a download of a URL asset to a destination file.
	ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes func(connection NSURLConnection, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64)
	// ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes — Sent to the delegate when an URL connection resumes downloading a URL asset that was earlier suspended.
	ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes func(connection NSURLConnection, totalBytesWritten int64, expectedTotalBytes int64)
	// ConnectionDidFinishDownloadingDestinationURL — Sent to the delegate when the URL connection has successfully downloaded the URL asset to a destination file.
	ConnectionDidFinishDownloadingDestinationURL func(connection NSURLConnection, destinationURL objectivec.Object)
}

// NewNSURLConnectionDownloadDelegate creates an Objective-C object implementing the [NSURLConnectionDownloadDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSURLConnectionDownloadDelegateObject] satisfies the [NSURLConnectionDownloadDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsurlconnectiondownloaddelegate
func NewNSURLConnectionDownloadDelegate(config NSURLConnectionDownloadDelegateConfig) NSURLConnectionDownloadDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSURLConnectionDownloadDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ConnectionDidFinishDownloadingDestinationURL != nil {
		fn := config.ConnectionDidFinishDownloadingDestinationURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connectionDidFinishDownloading:destinationURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, destinationURLID objc.ID) {
				connection := NSURLConnectionFromID(connectionID)
				destinationURL := objectivec.ObjectFromID(destinationURLID)
				fn(connection, destinationURL)
			},
		})
	}

	if config.ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes != nil {
		fn := config.ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connection:didWriteData:totalBytesWritten:expectedTotalBytes:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, bytesWritten int64, totalBytesWritten int64, expectedTotalBytes int64) {
				connection := NSURLConnectionFromID(connectionID)
				fn(connection, bytesWritten, totalBytesWritten, expectedTotalBytes)
			},
		})
	}

	if config.ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes != nil {
		fn := config.ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connectionDidResumeDownloading:totalBytesWritten:expectedTotalBytes:"),
			Fn: func(self objc.ID, _cmd objc.SEL, connectionID objc.ID, totalBytesWritten int64, expectedTotalBytes int64) {
				connection := NSURLConnectionFromID(connectionID)
				fn(connection, totalBytesWritten, expectedTotalBytes)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSURLConnectionDownloadDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSURLConnectionDownloadDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSURLConnectionDownloadDelegateObjectFromID(instance)
}

