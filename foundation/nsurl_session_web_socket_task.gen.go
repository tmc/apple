// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [URLSessionWebSocketTask] class.
var (
	_URLSessionWebSocketTaskClass     URLSessionWebSocketTaskClass
	_URLSessionWebSocketTaskClassOnce sync.Once
)

func getURLSessionWebSocketTaskClass() URLSessionWebSocketTaskClass {
	_URLSessionWebSocketTaskClassOnce.Do(func() {
		_URLSessionWebSocketTaskClass = URLSessionWebSocketTaskClass{class: objc.GetClass("NSURLSessionWebSocketTask")}
	})
	return _URLSessionWebSocketTaskClass
}

// GetURLSessionWebSocketTaskClass returns the class object for NSURLSessionWebSocketTask.
func GetURLSessionWebSocketTaskClass() URLSessionWebSocketTaskClass {
	return getURLSessionWebSocketTaskClass()
}

type URLSessionWebSocketTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionWebSocketTaskClass) Alloc() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A URL session task that communicates over the WebSockets protocol standard.
//
// # Overview
// 
// [NSURLSessionWebSocketTask] is a concrete subclass of [NSURLSessionTask]
// that provides a message-oriented transport protocol over TCP and TLS in the
// form of WebSocket framing. It follows the WebSocket Protocol defined in
// [RFC 6455].
// 
// You create a [NSURLSessionWebSocketTask] with either a `` or `` URL. When
// creating the task, you can also provide a list of protocols to advertise
// during the handshake phase. Once the handshake completes, your app receives
// notifications through the session’s [Delegate].
// 
// You send data with [send(_:completionHandler:)] and receive data with
// [receive(completionHandler:)]. The task performs reads and writes
// asynchronously, and allows you to send and receive messages that contain
// both binary frames and UTF-8 encoded text frames. The task enqueues any
// reads or writes you perform prior to the handshake’s completion, and
// executes them after the handshake completes.
// 
// [NSURLSessionWebSocketTask] supports redirection and authentication like
// other types of tasks do, using the methods in [NSURLSessionTaskDelegate].
// The WebSocket task calls the redirection and authentication delegate
// methods prior to completing the handshake. The WebSocket task also supports
// cookies, by storing cookies to the session configuration’s
// [HTTPCookieStorage], and attaches cookies to outgoing HTTP handshake
// requests.
//
// [RFC 6455]: https://tools.ietf.org/html/rfc6455
// [receive(completionHandler:)]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/receive(completionHandler:)
// [send(_:completionHandler:)]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/send(_:completionHandler:)
//
// # Sending and receiving data
//
//   - [URLSessionWebSocketTask.MaximumMessageSize]: The maximum number of bytes to buffer before the receive call fails with an error.
//   - [URLSessionWebSocketTask.SetMaximumMessageSize]
//
// # Sending ping frames
//
//   - [URLSessionWebSocketTask.SendPingWithPongReceiveHandler]: Sends a ping frame from the client side, with a closure to receive the pong from the server endpoint.
//
// # Closing the connection
//
//   - [URLSessionWebSocketTask.CancelWithCloseCodeReason]: Sends a close frame with the given close code and optional close reason.
//   - [URLSessionWebSocketTask.CloseCode]: A code that indicates the reason a connection closed.
//   - [URLSessionWebSocketTask.CloseReason]: A block of data that provides further information about why a connection closed.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask
type URLSessionWebSocketTask struct {
	NSURLSessionTask
}

// URLSessionWebSocketTaskFromID constructs a [URLSessionWebSocketTask] from an objc.ID.
//
// A URL session task that communicates over the WebSockets protocol standard.
func URLSessionWebSocketTaskFromID(id objc.ID) URLSessionWebSocketTask {
	return NSURLSessionWebSocketTask{NSURLSessionTask: NSURLSessionTaskFromID(id)}
}

// NSURLSessionWebSocketTaskFromID is an alias for [URLSessionWebSocketTaskFromID] for cross-framework compatibility.
func NSURLSessionWebSocketTaskFromID(id objc.ID) URLSessionWebSocketTask { return URLSessionWebSocketTaskFromID(id) }
// NOTE: URLSessionWebSocketTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLSessionWebSocketTask] class.
//
// # Sending and receiving data
//
//   - [IURLSessionWebSocketTask.MaximumMessageSize]: The maximum number of bytes to buffer before the receive call fails with an error.
//   - [IURLSessionWebSocketTask.SetMaximumMessageSize]
//
// # Sending ping frames
//
//   - [IURLSessionWebSocketTask.SendPingWithPongReceiveHandler]: Sends a ping frame from the client side, with a closure to receive the pong from the server endpoint.
//
// # Closing the connection
//
//   - [IURLSessionWebSocketTask.CancelWithCloseCodeReason]: Sends a close frame with the given close code and optional close reason.
//   - [IURLSessionWebSocketTask.CloseCode]: A code that indicates the reason a connection closed.
//   - [IURLSessionWebSocketTask.CloseReason]: A block of data that provides further information about why a connection closed.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask
type IURLSessionWebSocketTask interface {
	INSURLSessionTask

	// Topic: Sending and receiving data

	// The maximum number of bytes to buffer before the receive call fails with an error.
	MaximumMessageSize() int
	SetMaximumMessageSize(value int)

	// Topic: Sending ping frames

	// Sends a ping frame from the client side, with a closure to receive the pong from the server endpoint.
	SendPingWithPongReceiveHandler(pongReceiveHandler ErrorHandler)

	// Topic: Closing the connection

	// Sends a close frame with the given close code and optional close reason.
	CancelWithCloseCodeReason(closeCode NSURLSessionWebSocketCloseCode, reason INSData)
	// A code that indicates the reason a connection closed.
	CloseCode() NSURLSessionWebSocketCloseCode
	// A block of data that provides further information about why a connection closed.
	CloseReason() INSData

	// The cookie store for storing cookies within this session.
	HttpCookieStorage() INSHTTPCookieStorage
	SetHttpCookieStorage(value INSHTTPCookieStorage)
	// Reads a WebSocket message once all the frames of the message are available.
	ReceiveMessageWithCompletionHandler(completionHandler URLSessionWebSocketMessageErrorHandler)
	// Sends a WebSocket message, receiving the result in a completion handler.
	SendMessageCompletionHandler(message INSURLSessionWebSocketMessage, completionHandler ErrorHandler)
}

// Init initializes the instance.
func (u URLSessionWebSocketTask) Init() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionWebSocketTask) Autorelease() URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionWebSocketTask creates a new URLSessionWebSocketTask instance.
func NewURLSessionWebSocketTask() URLSessionWebSocketTask {
	class := getURLSessionWebSocketTaskClass()
	rv := objc.Send[URLSessionWebSocketTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sends a ping frame from the client side, with a closure to receive the pong
// from the server endpoint.
//
// pongReceiveHandler: A closure called by the task when it receives the pong from the server. The
// block/closure receives an [NSError] that indicates a lost connection or
// other problem, or `nil` if no error occurred.
//
// # Discussion
// 
// When sending multiple pings, the task always calls `pongReceiveHandler` in
// the order it sent the pings.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/sendPing(pongReceiveHandler:)
func (u URLSessionWebSocketTask) SendPingWithPongReceiveHandler(pongReceiveHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(pongReceiveHandler)
	defer _cleanup0()
	objc.Send[objc.ID](u.ID, objc.Sel("sendPingWithPongReceiveHandler:"), _block0)
}

// Sends a close frame with the given close code and optional close reason.
//
// closeCode: A [URLSessionWebSocketTask.CloseCode] that indicates the reason for closing
// the connection.
// //
// [URLSessionWebSocketTask.CloseCode]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum
//
// reason: Optional further information to explain the closing. The value of this
// parameter is defined by the endpoints, not by the standard.
//
// # Discussion
// 
// If you call [Cancel] on the task instead of this method, it sends a
// cancellation frame with no close code or reason.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/cancel(with:reason:)
func (u URLSessionWebSocketTask) CancelWithCloseCodeReason(closeCode NSURLSessionWebSocketCloseCode, reason INSData) {
	objc.Send[objc.ID](u.ID, objc.Sel("cancelWithCloseCode:reason:"), closeCode, reason)
}

// Reads a WebSocket message once all the frames of the message are available.
//
// completionHandler: A closure that receives two parameters: the WebSocket message, and an
// [NSError] that indicates an error encountered while receiving the message.
// The error is `nil` if no error occurred.
//
// # Discussion
// 
// If the task reaches the [MaximumMessageSize] while buffering the frames,
// this call fails with an error.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketTask/receiveMessageWithCompletionHandler:
func (u URLSessionWebSocketTask) ReceiveMessageWithCompletionHandler(completionHandler URLSessionWebSocketMessageErrorHandler) {
_block0, _cleanup0 := NewURLSessionWebSocketMessageErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](u.ID, objc.Sel("receiveMessageWithCompletionHandler:"), _block0)
}

// Sends a WebSocket message, receiving the result in a completion handler.
//
// message: The WebSocket message to send.
//
// completionHandler: A block that receives an [NSError] that indicates an error encountered
// while sending, or `nil` if no error occurred.
//
// # Discussion
// 
// If an error occurs while sending the message, any outstanding work also
// fails.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketTask/sendMessage:completionHandler:
func (u URLSessionWebSocketTask) SendMessageCompletionHandler(message INSURLSessionWebSocketMessage, completionHandler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](u.ID, objc.Sel("sendMessage:completionHandler:"), message, _block1)
}

// The maximum number of bytes to buffer before the receive call fails with an
// error.
//
// # Discussion
// 
// This value includes the sum of all bytes from continuation frames. Receive
// calls will fail once the task reaches this limit.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/maximumMessageSize
func (u URLSessionWebSocketTask) MaximumMessageSize() int {
	rv := objc.Send[int](u.ID, objc.Sel("maximumMessageSize"))
	return rv
}
func (u URLSessionWebSocketTask) SetMaximumMessageSize(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setMaximumMessageSize:"), value)
}

// A code that indicates the reason a connection closed.
//
// # Discussion
// 
// You can retrieve the close code at any time. When the task is not yet
// closed, this value is [URLSessionWebSocketCloseCodeInvalid].
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/closeCode-swift.property
func (u URLSessionWebSocketTask) CloseCode() NSURLSessionWebSocketCloseCode {
	rv := objc.Send[NSURLSessionWebSocketCloseCode](u.ID, objc.Sel("closeCode"))
	return NSURLSessionWebSocketCloseCode(rv)
}

// A block of data that provides further information about why a connection
// closed.
//
// # Discussion
// 
// The close reason provides further information about why a connection
// closed, beyond that provided by the [CloseCode]. The value of this property
// isn’t defined by [RFC 6455]; the endpoints define how it’s used.
// 
// You can retrieve the close reason at any time. When the task is not yet
// closed, this value is [URLSessionWebSocketCloseCodeInvalid].
//
// [RFC 6455]: https://tools.ietf.org/html/rfc6455
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/closeReason
func (u URLSessionWebSocketTask) CloseReason() INSData {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("closeReason"))
	return NSDataFromID(objc.ID(rv))
}

// The cookie store for storing cookies within this session.
//
// See: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookiestorage
func (u URLSessionWebSocketTask) HttpCookieStorage() INSHTTPCookieStorage {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("HTTPCookieStorage"))
	return NSHTTPCookieStorageFromID(objc.ID(rv))
}
func (u URLSessionWebSocketTask) SetHttpCookieStorage(value INSHTTPCookieStorage) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPCookieStorage:"), value)
}

// SendPingWithPongReceiveHandlerSync is a synchronous wrapper around [URLSessionWebSocketTask.SendPingWithPongReceiveHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSessionWebSocketTask) SendPingWithPongReceiveHandlerSync(ctx context.Context) error {
	done := make(chan error, 1)
	u.SendPingWithPongReceiveHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ReceiveMessage is a synchronous wrapper around [URLSessionWebSocketTask.ReceiveMessageWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSessionWebSocketTask) ReceiveMessage(ctx context.Context) (*NSURLSessionWebSocketMessage, error) {
	type result struct {
		val *NSURLSessionWebSocketMessage
		err error
	}
	done := make(chan result, 1)
	u.ReceiveMessageWithCompletionHandler(func(val *NSURLSessionWebSocketMessage, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SendMessage is a synchronous wrapper around [URLSessionWebSocketTask.SendMessageCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSessionWebSocketTask) SendMessage(ctx context.Context, message INSURLSessionWebSocketMessage) error {
	done := make(chan error, 1)
	u.SendMessageCompletionHandler(message, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

