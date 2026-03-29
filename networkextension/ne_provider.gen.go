// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEProvider] class.
var (
	_NEProviderClass     NEProviderClass
	_NEProviderClassOnce sync.Once
)

func getNEProviderClass() NEProviderClass {
	_NEProviderClassOnce.Do(func() {
		_NEProviderClass = NEProviderClass{class: objc.GetClass("NEProvider")}
	})
	return _NEProviderClass
}

// GetNEProviderClass returns the class object for NEProvider.
func GetNEProviderClass() NEProviderClass {
	return getNEProviderClass()
}

type NEProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEProviderClass) Alloc() NEProvider {
	rv := objc.Send[NEProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for all NetworkExtension providers.
//
// # Overview
// 
// See the documentation for the [NEProvider] subclasses for details about how
// to create Network Extension Provider extensions.
// 
// The [NEProvider] class and its subclasses expose methods and properties
// that allow Network Extension Provider extensions to participate in and
// affect the network data path on iOS and macOS. For example, the `` method
// in [NEFilterDataProvider] allows Filter Data Provider extensions to make
// pass/block decisions on TCP connections as the connections are established
// on the system.
// 
// # Subclassing Notes
// 
// The [NEProvider] class should not be subclassed directly. Instead, you
// should create subclasses of [NEProvider] subclasses (and in some cases
// subsubclasses).
// 
// # Methods to Override
// 
// - [NEProvider.SleepWithCompletionHandler]
// - [NEProvider.Wake]
//
// # Handling sleep and wake
//
//   - [NEProvider.SleepWithCompletionHandler]: Handle a sleep event.
//   - [NEProvider.Wake]: Handle a wake event.
//
// # Monitoring the network state
//
//   - [NEProvider.DefaultPath]: The current default network path used for connections created by the provider.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProvider
type NEProvider struct {
	objectivec.Object
}

// NEProviderFromID constructs a [NEProvider] from an objc.ID.
//
// An abstract base class for all NetworkExtension providers.
func NEProviderFromID(id objc.ID) NEProvider {
	return NEProvider{objectivec.Object{ID: id}}
}
// NOTE: NEProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEProvider] class.
//
// # Handling sleep and wake
//
//   - [INEProvider.SleepWithCompletionHandler]: Handle a sleep event.
//   - [INEProvider.Wake]: Handle a wake event.
//
// # Monitoring the network state
//
//   - [INEProvider.DefaultPath]: The current default network path used for connections created by the provider.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProvider
type INEProvider interface {
	objectivec.IObject

	// Topic: Handling sleep and wake

	// Handle a sleep event.
	SleepWithCompletionHandler(completionHandler VoidHandler)
	// Handle a wake event.
	Wake()

	// Topic: Monitoring the network state

	// The current default network path used for connections created by the provider.
	DefaultPath() INWPath
}

// Init initializes the instance.
func (p NEProvider) Init() NEProvider {
	rv := objc.Send[NEProvider](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NEProvider) Autorelease() NEProvider {
	rv := objc.Send[NEProvider](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProvider creates a new NEProvider instance.
func NewNEProvider() NEProvider {
	class := getNEProviderClass()
	rv := objc.Send[NEProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Handle a sleep event.
//
// completionHandler: Implementations of this method must execute this block when the provider is
// finished handling the sleep event.
//
// # Discussion
// 
// This method is called by the system when the device is about to go to
// sleep.
// 
// [NEProvider] subclasses should override this method if the provider needs
// to perform any tasks before the device sleeps, such as disconnecting a
// tunnel connection.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProvider/sleep(completionHandler:)
func (p NEProvider) SleepWithCompletionHandler(completionHandler VoidHandler) {
_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("sleepWithCompletionHandler:"), _block0)
}
// Handle a wake event.
//
// # Discussion
// 
// This method is called by the system when the device wakes up from sleep
// mode.
// 
// [NEProvider] subclasses should override this method if the provider needs
// to perform any tasks when the device wakes up, such as reconnecting a
// tunnel connection.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProvider/wake()
func (p NEProvider) Wake() {
	objc.Send[objc.ID](p.ID, objc.Sel("wake"))
}

// Starts the Network Extension machinery from inside a System Extension.
//
// # Discussion
// 
// Call this method as early as possible after your system extension starts.
// 
// Once called, this class method causes your system extension to start
// handling requests from the Network Extension session manager daemon to
// instantiate appropriate [NEProvider] subclass instances. The system
// extension must declare a mapping of Network Extension extension points to
// [NEProvider] subclass instances in its `Info.Plist()`. The following
// example shows this mapping:
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProvider/startSystemExtensionMode()
func (_NEProviderClass NEProviderClass) StartSystemExtensionMode() {
	objc.Send[objc.ID](objc.ID(_NEProviderClass.class), objc.Sel("startSystemExtensionMode"))
}

// The current default network path used for connections created by the
// provider.
//
// # Discussion
// 
// This NWPath object contains information about which physical network
// interface will be used by connections opened by the Network Extension
// provider. You can determine when this physical interface changes by
// observing this property using KVO.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProvider/defaultPath
func (p NEProvider) DefaultPath() INWPath {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("defaultPath"))
	return NWPathFromID(objc.ID(rv))
}

// Sleep is a synchronous wrapper around [NEProvider.SleepWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p NEProvider) Sleep(ctx context.Context) error {
	done := make(chan struct{}, 1)
	p.SleepWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

