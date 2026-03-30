// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEFilterProvider] class.
var (
	_NEFilterProviderClass     NEFilterProviderClass
	_NEFilterProviderClassOnce sync.Once
)

func getNEFilterProviderClass() NEFilterProviderClass {
	_NEFilterProviderClassOnce.Do(func() {
		_NEFilterProviderClass = NEFilterProviderClass{class: objc.GetClass("NEFilterProvider")}
	})
	return _NEFilterProviderClass
}

// GetNEFilterProviderClass returns the class object for NEFilterProvider.
func GetNEFilterProviderClass() NEFilterProviderClass {
	return getNEFilterProviderClass()
}

type NEFilterProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterProviderClass) Alloc() NEFilterProvider {
	rv := objc.Send[NEFilterProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class shared by content filters.
//
// # Overview
//
// A Network Content Filter is made up of two Filter Provider extensions:
//
// The examines network content as it passes through the network stack on the
// device and decides if the network content should be blocked or allowed to
// pass on to its final destination.
//
// Because the Filter Data Provider extension has access to all of the network
// content flowing through the device, it runs in a very restrictive sandbox.
// The sandbox prevents the Filter Data Provider extension from moving network
// content outside of its address space by blocking all network access, IPC,
// and disk write operations.
//
// The Filter Data Provider extension is implemented by creating a custom
// subclass of the [NEFilterDataProvider] class.
//
// The is responsible for feeding information to the Filter Data Provider
// extension so that the Filter Data Provider extension can do its job.
//
// For example, the Filter Control Provider extension can be notified by the
// Filter Data Provider extension that it does not have enough information to
// make a decision about a particular flow of network content. The Filter
// Control Provider extension can then download more filtering rules from a
// server and write the rules to a location where the Filter Data Provider can
// access them.
//
// The Filter Control Provider extension is implemented by creating a custom
// subclass of the [NEFilterControlProvider] class.
//
// # Subclassing Notes
//
// [NEFilterProvider] should not be subclassed directly. Instead, you should
// create subclasses of `NEFilterProvider’s` subclasses and override the
// following methods:
//
// # Methods to Override
//
// - [NEFilterProvider.StartFilterWithCompletionHandler] -
// [NEFilterProvider.StopFilterWithReasonCompletionHandler]
//
// # Managing the filter life cycle
//
//   - [NEFilterProvider.StartFilterWithCompletionHandler]: Start the filter.
//   - [NEFilterProvider.StopFilterWithReasonCompletionHandler]: Stop the filter.
//
// # Getting the filter configuration
//
//   - [NEFilterProvider.FilterConfiguration]: An [NEFilterProviderConfiguration](<doc://com.apple.networkextension/documentation/NetworkExtension/NEFilterProviderConfiguration>) object containing the current filter configuration.
//
// # Receiving reports
//
//   - [NEFilterProvider.HandleReport]: Receives a report from the framework.
//
// # Handling errors
//
//   - [NEFilterProvider.NEFilterErrorDomain]: The domain for errors resulting from calls to the filter manager.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProvider
//
// [NEFilterControlProvider]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider
type NEFilterProvider struct {
	NEProvider
}

// NEFilterProviderFromID constructs a [NEFilterProvider] from an objc.ID.
//
// An abstract base class shared by content filters.
func NEFilterProviderFromID(id objc.ID) NEFilterProvider {
	return NEFilterProvider{NEProvider: NEProviderFromID(id)}
}

// NOTE: NEFilterProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterProvider] class.
//
// # Managing the filter life cycle
//
//   - [INEFilterProvider.StartFilterWithCompletionHandler]: Start the filter.
//   - [INEFilterProvider.StopFilterWithReasonCompletionHandler]: Stop the filter.
//
// # Getting the filter configuration
//
//   - [INEFilterProvider.FilterConfiguration]: An [NEFilterProviderConfiguration](<doc://com.apple.networkextension/documentation/NetworkExtension/NEFilterProviderConfiguration>) object containing the current filter configuration.
//
// # Receiving reports
//
//   - [INEFilterProvider.HandleReport]: Receives a report from the framework.
//
// # Handling errors
//
//   - [INEFilterProvider.NEFilterErrorDomain]: The domain for errors resulting from calls to the filter manager.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProvider
type INEFilterProvider interface {
	INEProvider

	// Topic: Managing the filter life cycle

	// Start the filter.
	StartFilterWithCompletionHandler(completionHandler ErrorHandler)
	// Stop the filter.
	StopFilterWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler)

	// Topic: Getting the filter configuration

	// An [NEFilterProviderConfiguration](<doc://com.apple.networkextension/documentation/NetworkExtension/NEFilterProviderConfiguration>) object containing the current filter configuration.
	FilterConfiguration() INEFilterProviderConfiguration

	// Topic: Receiving reports

	// Receives a report from the framework.
	HandleReport(report INEFilterReport)

	// Topic: Handling errors

	// The domain for errors resulting from calls to the filter manager.
	NEFilterErrorDomain() string
}

// Init initializes the instance.
func (f NEFilterProvider) Init() NEFilterProvider {
	rv := objc.Send[NEFilterProvider](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterProvider) Autorelease() NEFilterProvider {
	rv := objc.Send[NEFilterProvider](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterProvider creates a new NEFilterProvider instance.
func NewNEFilterProvider() NEFilterProvider {
	class := getNEFilterProviderClass()
	rv := objc.Send[NEFilterProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Start the filter.
//
// completionHandler: A block that must be executed when the filter is running and is ready to
// filter network content.
//
// # Discussion
//
// This method is called by the system to start the filter.
//
// [NEFilterProvider] subclasses must override this method.
//
// When this method is called, the Filter Provider should perform any steps
// necessary to initialize the filter and then execute the `completionHandler`
// block.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProvider/startFilter(completionHandler:)
func (f NEFilterProvider) StartFilterWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("startFilterWithCompletionHandler:"), _block0)
}

// Stop the filter.
//
// reason: An [NEProviderStopReason] code indicating why the filter is being stopped.
// For a list of possible codes, see [NEProvider].
//
// completionHandler: A block that must be executed when the filter is fully stopped.
//
// # Discussion
//
// This method is called by the system to stop the filter.
//
// [NEFilterProvider] subclasses must override this method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProvider/stopFilter(with:completionHandler:)
func (f NEFilterProvider) StopFilterWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("stopFilterWithReason:completionHandler:"), reason, _block1)
}

// Receives a report from the framework.
//
// report: The report delivered from the framework.
//
// # Discussion
//
// The framework calls this method when the data provider extension returns a
// verdict with the [ShouldReport] property set to true. Override this method
// in a subclass if you want to handle the flow report.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProvider/handle(_:)
func (f NEFilterProvider) HandleReport(report INEFilterReport) {
	objc.Send[objc.ID](f.ID, objc.Sel("handleReport:"), report)
}

// An [NEFilterProviderConfiguration] object containing the current filter
// configuration.
//
// # Discussion
//
// The Filter Provider can observe this property to be notified when the
// configuration changes, using KVO. See [Key-Value Observing Programming
// Guide].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterProvider/filterConfiguration
//
// [Key-Value Observing Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/KeyValueObserving/KeyValueObserving.html#//apple_ref/doc/uid/10000177i
func (f NEFilterProvider) FilterConfiguration() INEFilterProviderConfiguration {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("filterConfiguration"))
	return NEFilterProviderConfigurationFromID(objc.ID(rv))
}

// The domain for errors resulting from calls to the filter manager.
//
// See: https://developer.apple.com/documentation/networkextension/nefiltererrordomain
func (f NEFilterProvider) NEFilterErrorDomain() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NEFilterErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// StartFilter is a synchronous wrapper around [NEFilterProvider.StartFilterWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NEFilterProvider) StartFilter(ctx context.Context) error {
	done := make(chan error, 1)
	f.StartFilterWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopFilterWithReason is a synchronous wrapper around [NEFilterProvider.StopFilterWithReasonCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NEFilterProvider) StopFilterWithReason(ctx context.Context, reason NEProviderStopReason) error {
	done := make(chan struct{}, 1)
	f.StopFilterWithReasonCompletionHandler(reason, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
