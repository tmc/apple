// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [URLSessionStreamTask] class.
var (
	_URLSessionStreamTaskClass     URLSessionStreamTaskClass
	_URLSessionStreamTaskClassOnce sync.Once
)

func getURLSessionStreamTaskClass() URLSessionStreamTaskClass {
	_URLSessionStreamTaskClassOnce.Do(func() {
		_URLSessionStreamTaskClass = URLSessionStreamTaskClass{class: objc.GetClass("NSURLSessionStreamTask")}
	})
	return _URLSessionStreamTaskClass
}

// GetURLSessionStreamTaskClass returns the class object for NSURLSessionStreamTask.
func GetURLSessionStreamTaskClass() URLSessionStreamTaskClass {
	return getURLSessionStreamTaskClass()
}

type URLSessionStreamTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLSessionStreamTaskClass) Alloc() URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}







// A URL session task that is stream-based.
//
// # Overview
// 
// [NSURLSessionStreamTask] is a concrete subclass of [NSURLSessionTask]. Many
// of the methods in the [NSURLSessionStreamTask] class are documented in
// [NSURLSessionTask].
// 
// The [NSURLSessionStreamTask] class provides an interface a TCP/IP
// connection created via [NSURLSession]. Tasks may be created from an
// [NSURLSession] using the [StreamTaskWithHostNamePort] and
// [StreamTaskWithNetService] methods. They may also be created as a result of
// an [NSURLSessionDataTask] being upgraded via the HTTP `Upgrade:` response
// header and appropriate use of the [HTTPShouldUsePipelining] option of
// [NSURLSessionConfiguration].
// 
// A [NSURLSessionStreamTask] object performs asynchronous reads and writes,
// which are enqueued and executed serially, calling a handler upon completion
// being on the session delegate queue. If the task is canceled, all enqueued
// reads and writes will call their completion handlers with an appropriate
// error.
// 
// When working with APIs that accept [NSStream] objects, you can create
// [NSInputStream] and [NSOutputStream] objects from an
// [NSURLSessionStreamTask] object by calling the [CaptureStreams] method.
//
// # Reading and writing data
//
//   - [URLSessionStreamTask.ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler]: Asynchronously reads a number of bytes from the stream, and calls a handler upon completion.
//   - [URLSessionStreamTask.WriteDataTimeoutCompletionHandler]: Asynchronously writes the specified data to the stream, and calls a handler upon completion.
//
// # Capturing streams
//
//   - [URLSessionStreamTask.CaptureStreams]: Completes any already enqueued reads and writes, and then invokes the [urlSession(_:streamTask:didBecome:outputStream:)](<doc://com.apple.foundation/documentation/Foundation/URLSessionStreamDelegate/urlSession(_:streamTask:didBecome:outputStream:)>) delegate message.
//
// # Closing read and write sockets
//
//   - [URLSessionStreamTask.CloseRead]: Completes any enqueued reads and writes, and then closes the read side of the underlying socket.
//   - [URLSessionStreamTask.CloseWrite]: Completes any enqueued reads and writes, and then closes the write side of the underlying socket.
//
// # Starting and stopping secure connections
//
//   - [URLSessionStreamTask.StartSecureConnection]: Completes any enqueued reads and writes, and establishes a secure connection.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask
type URLSessionStreamTask struct {
	NSURLSessionTask
}

// URLSessionStreamTaskFromID constructs a [URLSessionStreamTask] from an objc.ID.
//
// A URL session task that is stream-based.
func URLSessionStreamTaskFromID(id objc.ID) URLSessionStreamTask {
	return NSURLSessionStreamTask{NSURLSessionTask: NSURLSessionTaskFromID(id)}
}

// NSURLSessionStreamTaskFromID is an alias for [URLSessionStreamTaskFromID] for cross-framework compatibility.
func NSURLSessionStreamTaskFromID(id objc.ID) URLSessionStreamTask { return URLSessionStreamTaskFromID(id) }
// NOTE: URLSessionStreamTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [URLSessionStreamTask] class.
//
// # Reading and writing data
//
//   - [IURLSessionStreamTask.ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler]: Asynchronously reads a number of bytes from the stream, and calls a handler upon completion.
//   - [IURLSessionStreamTask.WriteDataTimeoutCompletionHandler]: Asynchronously writes the specified data to the stream, and calls a handler upon completion.
//
// # Capturing streams
//
//   - [IURLSessionStreamTask.CaptureStreams]: Completes any already enqueued reads and writes, and then invokes the [urlSession(_:streamTask:didBecome:outputStream:)](<doc://com.apple.foundation/documentation/Foundation/URLSessionStreamDelegate/urlSession(_:streamTask:didBecome:outputStream:)>) delegate message.
//
// # Closing read and write sockets
//
//   - [IURLSessionStreamTask.CloseRead]: Completes any enqueued reads and writes, and then closes the read side of the underlying socket.
//   - [IURLSessionStreamTask.CloseWrite]: Completes any enqueued reads and writes, and then closes the write side of the underlying socket.
//
// # Starting and stopping secure connections
//
//   - [IURLSessionStreamTask.StartSecureConnection]: Completes any enqueued reads and writes, and establishes a secure connection.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask
type IURLSessionStreamTask interface {
	INSURLSessionTask

	// Topic: Reading and writing data

	// Asynchronously reads a number of bytes from the stream, and calls a handler upon completion.
	ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler(minBytes uint, maxBytes uint, timeout float64, completionHandler DataErrorHandler)
	// Asynchronously writes the specified data to the stream, and calls a handler upon completion.
	WriteDataTimeoutCompletionHandler(data INSData, timeout float64, completionHandler ErrorHandler)

	// Topic: Capturing streams

	// Completes any already enqueued reads and writes, and then invokes the [urlSession(_:streamTask:didBecome:outputStream:)](<doc://com.apple.foundation/documentation/Foundation/URLSessionStreamDelegate/urlSession(_:streamTask:didBecome:outputStream:)>) delegate message.
	CaptureStreams()

	// Topic: Closing read and write sockets

	// Completes any enqueued reads and writes, and then closes the read side of the underlying socket.
	CloseRead()
	// Completes any enqueued reads and writes, and then closes the write side of the underlying socket.
	CloseWrite()

	// Topic: Starting and stopping secure connections

	// Completes any enqueued reads and writes, and establishes a secure connection.
	StartSecureConnection()

	// A Boolean value that determines whether the session should use HTTP pipelining.
	HttpShouldUsePipelining() bool
	SetHttpShouldUsePipelining(value bool)
}





// Init initializes the instance.
func (u URLSessionStreamTask) Init() URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLSessionStreamTask) Autorelease() URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionStreamTask creates a new URLSessionStreamTask instance.
func NewURLSessionStreamTask() URLSessionStreamTask {
	class := getURLSessionStreamTaskClass()
	rv := objc.Send[URLSessionStreamTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Asynchronously reads a number of bytes from the stream, and calls a handler
// upon completion.
//
// minBytes: The minimum number of bytes to read.
//
// maxBytes: The maximum number of bytes to read.
//
// timeout: A timeout for reading bytes. If the read is not completed within the
// specified interval, the read is canceled and the `completionHandler` is
// called with an error. Pass `0` to prevent a read from timing out.
//
// completionHandler: The completion handler to call when all bytes are read, or an error occurs.
// This handler is executed on the delegate queue.
// 
// This completion handler takes the following parameters:
// 
// `data`: The data read from the stream. `atEOF`: Whether or not the stream
// reached end-of-file (EOF), such that no more data can be read. `error`: An
// error object that indicates why the read failed, or `nil` if the read was
// successful.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/readData(ofMinLength:maxLength:timeout:completionHandler:)
func (u URLSessionStreamTask) ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler(minBytes uint, maxBytes uint, timeout float64, completionHandler DataErrorHandler) {
		_block3, _cleanup3 := NewDataErrorBlock(completionHandler)
	defer _cleanup3()
		objc.Send[objc.ID](u.ID, objc.Sel("readDataOfMinLength:maxLength:timeout:completionHandler:"), minBytes, maxBytes, timeout, _block3)
}

// Asynchronously writes the specified data to the stream, and calls a handler
// upon completion.
//
// data: The data to be written.
//
// timeout: A timeout for writing bytes. If the write is not completed within the
// specified interval, the write is canceled and the `completionHandler` is
// called with an error. Pass `0` to prevent a write from timing out.
//
// completionHandler: The completion handler to call when all bytes are written, or an error
// occurs. This handler is executed on the delegate queue.
// 
// This completion handler takes the following parameter:
// 
// `error`: An error object that indicates why the write failed, or `nil` if
// the write was successful.
//
// # Discussion
// 
// There is no guarantee that the remote side of the stream has received all
// of the written bytes at the time that `completionHandler` is called, only
// that all of the data has been written to the kernel.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/write(_:timeout:completionHandler:)
func (u URLSessionStreamTask) WriteDataTimeoutCompletionHandler(data INSData, timeout float64, completionHandler ErrorHandler) {
		_block2, _cleanup2 := NewErrorBlock(completionHandler)
	defer _cleanup2()
		objc.Send[objc.ID](u.ID, objc.Sel("writeData:timeout:completionHandler:"), data, timeout, _block2)
}

// Completes any already enqueued reads and writes, and then invokes the
// [URLSessionStreamTaskDidBecomeInputStreamOutputStream] delegate message.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/captureStreams()
func (u URLSessionStreamTask) CaptureStreams() {
	objc.Send[objc.ID](u.ID, objc.Sel("captureStreams"))
}

// Completes any enqueued reads and writes, and then closes the read side of
// the underlying socket.
//
// # Discussion
// 
// You may continue to write data using the
// [WriteDataTimeoutCompletionHandler] method after calling this method. Any
// calls to [ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler] after
// calling this method will result in an error.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/closeRead()
func (u URLSessionStreamTask) CloseRead() {
	objc.Send[objc.ID](u.ID, objc.Sel("closeRead"))
}

// Completes any enqueued reads and writes, and then closes the write side of
// the underlying socket.
//
// # Discussion
// 
// You may continue to read data using the
// [ReadDataOfMinLengthMaxLengthTimeoutCompletionHandler] method after calling
// this method. Any calls to [WriteDataTimeoutCompletionHandler] after calling
// this method will result in an error.
// 
// Because the server may continue to write bytes to the client, it is
// recommended that you continue reading until the stream reaches end-of-file
// (EOF).
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/closeWrite()
func (u URLSessionStreamTask) CloseWrite() {
	objc.Send[objc.ID](u.ID, objc.Sel("closeWrite"))
}

// Completes any enqueued reads and writes, and establishes a secure
// connection.
//
// # Discussion
// 
// Authentication callbacks are sent to the session’s delegate using the
// [URLSessionTaskDidReceiveChallengeCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/startSecureConnection()
func (u URLSessionStreamTask) StartSecureConnection() {
	objc.Send[objc.ID](u.ID, objc.Sel("startSecureConnection"))
}












// A Boolean value that determines whether the session should use HTTP
// pipelining.
//
// See: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldusepipelining
func (u URLSessionStreamTask) HttpShouldUsePipelining() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}
func (u URLSessionStreamTask) SetHttpShouldUsePipelining(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setHTTPShouldUsePipelining:"), value)
}






















// WriteDataTimeout is a synchronous wrapper around [URLSessionStreamTask.WriteDataTimeoutCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u URLSessionStreamTask) WriteDataTimeout(ctx context.Context, data INSData, timeout float64) error {
	done := make(chan error, 1)
	u.WriteDataTimeoutCompletionHandler(data, timeout, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}






