// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSClient] class.
var (
	_FSClientClass     FSClientClass
	_FSClientClassOnce sync.Once
)

func getFSClientClass() FSClientClass {
	_FSClientClassOnce.Do(func() {
		_FSClientClass = FSClientClass{class: objc.GetClass("FSClient")}
	})
	return _FSClientClass
}

// GetFSClientClass returns the class object for FSClient.
func GetFSClientClass() FSClientClass {
	return getFSClientClass()
}

type FSClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSClientClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSClientClass) Alloc() FSClient {
	rv := objc.Send[FSClient](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An interface for apps and daemons to interact with FSKit.
//
// # Overview
//
// FSClient is the primary management interface for FSKit. Use this class to
// discover FSKit extensions installed on the system, including your own.
//
// # Discovering installed extensions
//
//   - [FSClient.FetchInstalledExtensionsWithCompletionHandler]: Asynchronously retrieves an list of installed file system modules.
//
// See: https://developer.apple.com/documentation/FSKit/FSClient
type FSClient struct {
	objectivec.Object
}

// FSClientFromID constructs a [FSClient] from an objc.ID.
//
// An interface for apps and daemons to interact with FSKit.
func FSClientFromID(id objc.ID) FSClient {
	return FSClient{objectivec.Object{ID: id}}
}

// NOTE: FSClient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSClient] class.
//
// # Discovering installed extensions
//
//   - [IFSClient.FetchInstalledExtensionsWithCompletionHandler]: Asynchronously retrieves an list of installed file system modules.
//
// See: https://developer.apple.com/documentation/FSKit/FSClient
type IFSClient interface {
	objectivec.IObject

	// Topic: Discovering installed extensions

	// Asynchronously retrieves an list of installed file system modules.
	FetchInstalledExtensionsWithCompletionHandler(completionHandler FSModuleIdentityArrayErrorHandler)

	// Fetches installed FSKit modules using the FSClient XPC method.
	InstalledExtensions(handler FSModuleIdentityArrayErrorHandler)
	// Fetches installed FSKit modules using the FSClient synchronous XPC method.
	InstalledExtensionsSync(handler FSModuleIdentityArrayErrorHandler)
	// Sets the enabled state for an FSKit module.
	SetEnabledStateForIdentifier(identifier string, enabled bool, handler ErrorHandler)
}

// Init initializes the instance.
func (c FSClient) Init() FSClient {
	rv := objc.Send[FSClient](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c FSClient) Autorelease() FSClient {
	rv := objc.Send[FSClient](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSClient creates a new FSClient instance.
func NewFSClient() FSClient {
	class := getFSClientClass()
	rv := objc.Send[FSClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Asynchronously retrieves an list of installed file system modules.
//
// completionHandler: A block or closure that executes when FSKit finishes its fetch process. If
// the fetch succeeds, the first parameter contains an array of
// [FSModuleIdentity] instances that identify installed modules. If the fetch
// fails, the second parameter contains an error detailing the failure.
//
// # Discussion
//
// In Swift, you can either call this method and pass a completion handler
// closure, or get the value of the `installedExtensions` property with the
// `async` keyword.
//
// See: https://developer.apple.com/documentation/FSKit/FSClient/fetchInstalledExtensions(completionHandler:)
func (c FSClient) FetchInstalledExtensionsWithCompletionHandler(completionHandler FSModuleIdentityArrayErrorHandler) {
	_block0, _ := NewFSModuleIdentityArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchInstalledExtensionsWithCompletionHandler:"), _block0)
}

// Fetches installed FSKit modules using the FSClient XPC method. [Full Topic]
func (c FSClient) InstalledExtensions(handler FSModuleIdentityArrayErrorHandler) {
	_block0, _ := NewFSModuleIdentityArrayErrorBlock(handler)
	objc.Send[objc.ID](c.ID, objc.Sel("installedExtensions:"), _block0)
}

// Fetches installed FSKit modules using the FSClient synchronous XPC method. [Full Topic]
func (c FSClient) InstalledExtensionsSync(handler FSModuleIdentityArrayErrorHandler) {
	_block0, _ := NewFSModuleIdentityArrayErrorBlock(handler)
	objc.Send[objc.ID](c.ID, objc.Sel("installedExtensionsSync:"), _block0)
}

// Sets the enabled state for an FSKit module. [Full Topic]
func (c FSClient) SetEnabledStateForIdentifier(identifier string, enabled bool, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](c.ID, objc.Sel("setEnabledStateForIdentifier:newState:replyHandler:"), objc.String(identifier), enabled, _block2)
}

// The shared instance of the FSKit client class.
//
// See: https://developer.apple.com/documentation/FSKit/FSClient/shared
func (_FSClientClass FSClientClass) SharedInstance() FSClient {
	rv := objc.Send[objc.ID](objc.ID(_FSClientClass.class), objc.Sel("sharedInstance"))
	return FSClientFromID(objc.ID(rv))
}

// FetchInstalledExtensions is a synchronous wrapper around [FSClient.FetchInstalledExtensionsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c FSClient) FetchInstalledExtensions(ctx context.Context) ([]FSModuleIdentity, error) {
	type result struct {
		val []FSModuleIdentity
		err error
	}
	done := make(chan result, 1)
	c.FetchInstalledExtensionsWithCompletionHandler(func(val *[]FSModuleIdentity, err error) {
		var out []FSModuleIdentity
		if val != nil {
			out = append(out, (*val)...)
		}
		done <- result{out, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SetEnabledState is a synchronous wrapper around [FSClient.SetEnabledStateForIdentifier].
// It blocks until the completion handler fires or the context is cancelled.
func (c FSClient) SetEnabledState(ctx context.Context, identifier string, enabled bool) error {
	done := make(chan error, 1)
	c.SetEnabledStateForIdentifier(identifier, enabled, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
