// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionMessagePort] class.
var (
	_WKWebExtensionMessagePortClass     WKWebExtensionMessagePortClass
	_WKWebExtensionMessagePortClassOnce sync.Once
)

func getWKWebExtensionMessagePortClass() WKWebExtensionMessagePortClass {
	_WKWebExtensionMessagePortClassOnce.Do(func() {
		_WKWebExtensionMessagePortClass = WKWebExtensionMessagePortClass{class: objc.GetClass("WKWebExtensionMessagePort")}
	})
	return _WKWebExtensionMessagePortClass
}

// GetWKWebExtensionMessagePortClass returns the class object for WKWebExtensionMessagePort.
func GetWKWebExtensionMessagePortClass() WKWebExtensionMessagePortClass {
	return getWKWebExtensionMessagePortClass()
}

type WKWebExtensionMessagePortClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionMessagePortClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionMessagePortClass) Alloc() WKWebExtensionMessagePort {
	rv := objc.Send[WKWebExtensionMessagePort](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages message-based communication with a web extension.
//
// # Overview
//
// Contains properties and methods to handle message exchanges with a web
// extension.
//
// # Instance Properties
//
//   - [WKWebExtensionMessagePort.ApplicationIdentifier]: The unique identifier for the app to which this port should be connected.
//   - [WKWebExtensionMessagePort.DisconnectHandler]: The block to be executed when the port disconnects.
//   - [WKWebExtensionMessagePort.SetDisconnectHandler]
//   - [WKWebExtensionMessagePort.Disconnected]: Indicates whether the message port is disconnected.
//   - [WKWebExtensionMessagePort.MessageHandler]: The block to be executed when a message is received from the web extension.
//   - [WKWebExtensionMessagePort.SetMessageHandler]
//
// # Instance Methods
//
//   - [WKWebExtensionMessagePort.Disconnect]: Disconnects the port, terminating all further messages.
//   - [WKWebExtensionMessagePort.DisconnectWithError]: Disconnects the port, terminating all further messages with an optional error.
//   - [WKWebExtensionMessagePort.SendMessageCompletionHandler]: Sends a message to the connected web extension.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort
type WKWebExtensionMessagePort struct {
	objectivec.Object
}

// WKWebExtensionMessagePortFromID constructs a [WKWebExtensionMessagePort] from an objc.ID.
//
// An object that manages message-based communication with a web extension.
func WKWebExtensionMessagePortFromID(id objc.ID) WKWebExtensionMessagePort {
	return WKWebExtensionMessagePort{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionMessagePort adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionMessagePort] class.
//
// # Instance Properties
//
//   - [IWKWebExtensionMessagePort.ApplicationIdentifier]: The unique identifier for the app to which this port should be connected.
//   - [IWKWebExtensionMessagePort.DisconnectHandler]: The block to be executed when the port disconnects.
//   - [IWKWebExtensionMessagePort.SetDisconnectHandler]
//   - [IWKWebExtensionMessagePort.Disconnected]: Indicates whether the message port is disconnected.
//   - [IWKWebExtensionMessagePort.MessageHandler]: The block to be executed when a message is received from the web extension.
//   - [IWKWebExtensionMessagePort.SetMessageHandler]
//
// # Instance Methods
//
//   - [IWKWebExtensionMessagePort.Disconnect]: Disconnects the port, terminating all further messages.
//   - [IWKWebExtensionMessagePort.DisconnectWithError]: Disconnects the port, terminating all further messages with an optional error.
//   - [IWKWebExtensionMessagePort.SendMessageCompletionHandler]: Sends a message to the connected web extension.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort
type IWKWebExtensionMessagePort interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The unique identifier for the app to which this port should be connected.
	ApplicationIdentifier() string
	// The block to be executed when the port disconnects.
	DisconnectHandler() ErrorHandler
	SetDisconnectHandler(value ErrorHandler)
	// Indicates whether the message port is disconnected.
	Disconnected() bool
	// The block to be executed when a message is received from the web extension.
	MessageHandler() ObjectErrorHandler
	SetMessageHandler(value ObjectErrorHandler)

	// Topic: Instance Methods

	// Disconnects the port, terminating all further messages.
	Disconnect()
	// Disconnects the port, terminating all further messages with an optional error.
	DisconnectWithError(error_ foundation.INSError)
	// Sends a message to the connected web extension.
	SendMessageCompletionHandler(message objectivec.IObject, completionHandler ErrorHandler)
}

// Init initializes the instance.
func (w WKWebExtensionMessagePort) Init() WKWebExtensionMessagePort {
	rv := objc.Send[WKWebExtensionMessagePort](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionMessagePort) Autorelease() WKWebExtensionMessagePort {
	rv := objc.Send[WKWebExtensionMessagePort](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionMessagePort creates a new WKWebExtensionMessagePort instance.
func NewWKWebExtensionMessagePort() WKWebExtensionMessagePort {
	class := getWKWebExtensionMessagePortClass()
	rv := objc.Send[WKWebExtensionMessagePort](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Disconnects the port, terminating all further messages.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/disconnect()
func (w WKWebExtensionMessagePort) Disconnect() {
	objc.Send[objc.ID](w.ID, objc.Sel("disconnect"))
}

// Disconnects the port, terminating all further messages with an optional
// error.
//
// error: An optional error indicating the reason for disconnection.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/disconnect(throwing:)
func (w WKWebExtensionMessagePort) DisconnectWithError(error_ foundation.INSError) {
	objc.Send[objc.ID](w.ID, objc.Sel("disconnectWithError:"), error_)
}

// Sends a message to the connected web extension.
//
// message: The JSON-serializable message to be sent.
//
// completionHandler: An optional block to be invoked after the message is sent, taking an
// optional error.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/sendMessage(_:completionHandler:)
func (w WKWebExtensionMessagePort) SendMessageCompletionHandler(message objectivec.IObject, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("sendMessage:completionHandler:"), message, _block1)
}

// The unique identifier for the app to which this port should be connected.
//
// # Discussion
//
// This identifier is provided by the web extension and may or may not be used
// by the app. It’s up to the app to decide how to interpret this
// identifier.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/applicationIdentifier
func (w WKWebExtensionMessagePort) ApplicationIdentifier() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("applicationIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The block to be executed when the port disconnects.
//
// # Discussion
//
// An optional block to be invoked when the port disconnects, taking an
// optional error that indicates if the disconnection was caused by an error.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/disconnectHandler
func (w WKWebExtensionMessagePort) DisconnectHandler() ErrorHandler {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("disconnectHandler"))
	_ = rv
	return nil
}
func (w WKWebExtensionMessagePort) SetDisconnectHandler(value ErrorHandler) {
	block, cleanup := NewErrorBlock(value)
	defer cleanup()
	objc.Send[struct{}](w.ID, objc.Sel("setDisconnectHandler:"), block)
}

// Indicates whether the message port is disconnected.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/isDisconnected
func (w WKWebExtensionMessagePort) Disconnected() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isDisconnected"))
	return rv
}

// The block to be executed when a message is received from the web extension.
//
// # Discussion
//
// An optional block to be invoked when a message is received, taking two
// parameters: the message and an optional error.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/messageHandler
func (w WKWebExtensionMessagePort) MessageHandler() ObjectErrorHandler {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("messageHandler"))
	_ = rv
	return nil
}
func (w WKWebExtensionMessagePort) SetMessageHandler(value ObjectErrorHandler) {
	block, cleanup := NewObjectErrorBlock(value)
	defer cleanup()
	objc.Send[struct{}](w.ID, objc.Sel("setMessageHandler:"), block)
}

// SendMessage is a synchronous wrapper around [WKWebExtensionMessagePort.SendMessageCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebExtensionMessagePort) SendMessage(ctx context.Context, message objectivec.IObject) error {
	done := make(chan error, 1)
	w.SendMessageCompletionHandler(message, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
