// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NEFilterDataProvider] class.
var (
	_NEFilterDataProviderClass     NEFilterDataProviderClass
	_NEFilterDataProviderClassOnce sync.Once
)

func getNEFilterDataProviderClass() NEFilterDataProviderClass {
	_NEFilterDataProviderClassOnce.Do(func() {
		_NEFilterDataProviderClass = NEFilterDataProviderClass{class: objc.GetClass("NEFilterDataProvider")}
	})
	return _NEFilterDataProviderClass
}

// GetNEFilterDataProviderClass returns the class object for NEFilterDataProvider.
func GetNEFilterDataProviderClass() NEFilterDataProviderClass {
	return getNEFilterDataProviderClass()
}

type NEFilterDataProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterDataProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterDataProviderClass) Alloc() NEFilterDataProvider {
	rv := objc.Send[NEFilterDataProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The principal class for a filter data provider extension.
//
// # Overview
// 
// Network content is delivered to the Filter Data Provider in the form of
// [NEFilterFlow] objects. Each [NEFilterFlow] object corresponds to a network
// connection opened by an application running on the device. The Filter Data
// Provider can choose to pass or block the data when it receives a new flow,
// or it can ask the system to see more of the flow’s data in either the
// outbound or inbound direction before making a pass or block decision.
// 
// In addition to passing or blocking network data, the Filter Data Provider
// can tell the system that it needs more information before it can make a
// decision about a particular flow of data. The system will then ask the
// Filter Control Provider to update the current set of rules and place them
// in a location on disk that is readable from the Filter Data Provider
// extension.
// 
// When a [NEFilterFlow] object is originated from a WebKit browser object,
// the Filter Data Provider can affect the user experience in the following
// ways:
// 
// - If the Filter Data Provider chooses to block the web page, then a special
// “block” page is displayed in the WebKit browser object informing the
// user that their attempt to access the content was blocked. The Filter Data
// Provider can choose to add a link to this block page, giving the user the
// option of requesting access to the content. - If the Filter Data Provider
// chooses to allow the web page, then it can also specify that a string be
// appended to the web page URL. This allows the Filter Data Provider to
// direct the WebKit browser object to a “safe” version of the web page.
// 
// To protect the user’s privacy, the Filter Data Provider extension sandbox
// prevents the extension from moving network content outside of its address
// space.
// 
// # Creating a Filter Data Provider Extension
// 
// Filter Data Providers run as App Extensions for the
// `com.AppleXCUIElementTypeNetworkextensionXCUIElementTypeFilter()-data`
// extension point.
// 
// To create a Filter Data Provider extension, first create a new App
// Extension target in your project.
// 
// For an example of an Xcode build target for this app extension, see the
// [SimpleTunnel: Customized Networking Using the NetworkExtension Framework]
// sample code project.
// 
// Once you have a Filter Data Provider extension target, create a subclass of
// [NEFilterDataProvider]. Then set the [NSExtensionPrincipalClass] key in the
// the extension’s `Info.Plist()` to the name of your subclass.
// 
// If it is not done already, set the [NSExtensionPointIdentifier] key in the
// extension’s `Info.Plist()` to
// `com.AppleXCUIElementTypeNetworkextensionXCUIElementTypeFilter()-data`.
// 
// Here is an example of the [NSExtension] dictionary in a Filter Data
// Provider extension’s `Info.Plist()`:
// 
// Finally, add your Filter Data Provider extension target to your app’s
// Embed App Extensions build phase.
// 
// # Subclassing Notes
// 
// To create a Filter Data Provider extension, you must first create a
// subclass of [NEFilterDataProvider] and override the methods listed below.
// 
// # Methods to Override
// 
// - [NEFilterDataProvider.HandleNewFlow] -
// [NEFilterDataProvider.HandleInboundDataFromFlowReadBytesStartOffsetReadBytes] -
// [NEFilterDataProvider.HandleOutboundDataFromFlowReadBytesStartOffsetReadBytes] -
// [NEFilterDataProvider.HandleInboundDataCompleteForFlow] - [NEFilterDataProvider.HandleOutboundDataCompleteForFlow] -
// [NEFilterDataProvider.HandleRemediationForFlow] - [NEFilterDataProvider.HandleRulesChanged]
//
// [SimpleTunnel: Customized Networking Using the NetworkExtension Framework]: https://developer.apple.com/library/archive/samplecode/SimpleTunnel/Introduction/Intro.html#//apple_ref/doc/uid/TP40016140
//
// # Filtering network content
//
//   - [NEFilterDataProvider.HandleNewFlow]: Make a filtering decision for a newly-created flow of network content.
//   - [NEFilterDataProvider.HandleInboundDataFromFlowReadBytesStartOffsetReadBytes]: Make a filtering decision about a chunk of inbound data.
//   - [NEFilterDataProvider.HandleOutboundDataFromFlowReadBytesStartOffsetReadBytes]: Make a filtering decision about a chunk of outbound data.
//   - [NEFilterDataProvider.HandleInboundDataCompleteForFlow]: Make a filtering decision after seeing all of the inbound data for a flow.
//   - [NEFilterDataProvider.HandleOutboundDataCompleteForFlow]: Make a filtering decision after seeing all of the outbound data for a flow.
//
// # Changing filter settings
//
//   - [NEFilterDataProvider.ApplySettingsCompletionHandler]: Applies a set of filtering rules associated with the provider and changes the default filtering action.
//
// # Resuming data flows
//
//   - [NEFilterDataProvider.ResumeFlowWithVerdict]: Resumes a previously-paused flow.
//
// # Updating filter verdicts
//
//   - [NEFilterDataProvider.UpdateFlowUsingVerdictForDirection]: Updates the verdict for a flow outside the context of any filter data provider callback.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider
type NEFilterDataProvider struct {
	NEFilterProvider
}

// NEFilterDataProviderFromID constructs a [NEFilterDataProvider] from an objc.ID.
//
// The principal class for a filter data provider extension.
func NEFilterDataProviderFromID(id objc.ID) NEFilterDataProvider {
	return NEFilterDataProvider{NEFilterProvider: NEFilterProviderFromID(id)}
}
// NOTE: NEFilterDataProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterDataProvider] class.
//
// # Filtering network content
//
//   - [INEFilterDataProvider.HandleNewFlow]: Make a filtering decision for a newly-created flow of network content.
//   - [INEFilterDataProvider.HandleInboundDataFromFlowReadBytesStartOffsetReadBytes]: Make a filtering decision about a chunk of inbound data.
//   - [INEFilterDataProvider.HandleOutboundDataFromFlowReadBytesStartOffsetReadBytes]: Make a filtering decision about a chunk of outbound data.
//   - [INEFilterDataProvider.HandleInboundDataCompleteForFlow]: Make a filtering decision after seeing all of the inbound data for a flow.
//   - [INEFilterDataProvider.HandleOutboundDataCompleteForFlow]: Make a filtering decision after seeing all of the outbound data for a flow.
//
// # Changing filter settings
//
//   - [INEFilterDataProvider.ApplySettingsCompletionHandler]: Applies a set of filtering rules associated with the provider and changes the default filtering action.
//
// # Resuming data flows
//
//   - [INEFilterDataProvider.ResumeFlowWithVerdict]: Resumes a previously-paused flow.
//
// # Updating filter verdicts
//
//   - [INEFilterDataProvider.UpdateFlowUsingVerdictForDirection]: Updates the verdict for a flow outside the context of any filter data provider callback.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider
type INEFilterDataProvider interface {
	INEFilterProvider

	// Topic: Filtering network content

	// Make a filtering decision for a newly-created flow of network content.
	HandleNewFlow(flow INEFilterFlow) INEFilterNewFlowVerdict
	// Make a filtering decision about a chunk of inbound data.
	HandleInboundDataFromFlowReadBytesStartOffsetReadBytes(flow INEFilterFlow, offset uint, readBytes foundation.INSData) INEFilterDataVerdict
	// Make a filtering decision about a chunk of outbound data.
	HandleOutboundDataFromFlowReadBytesStartOffsetReadBytes(flow INEFilterFlow, offset uint, readBytes foundation.INSData) INEFilterDataVerdict
	// Make a filtering decision after seeing all of the inbound data for a flow.
	HandleInboundDataCompleteForFlow(flow INEFilterFlow) INEFilterDataVerdict
	// Make a filtering decision after seeing all of the outbound data for a flow.
	HandleOutboundDataCompleteForFlow(flow INEFilterFlow) INEFilterDataVerdict

	// Topic: Changing filter settings

	// Applies a set of filtering rules associated with the provider and changes the default filtering action.
	ApplySettingsCompletionHandler(settings INEFilterSettings, completionHandler ErrorHandler)

	// Topic: Resuming data flows

	// Resumes a previously-paused flow.
	ResumeFlowWithVerdict(flow INEFilterFlow, verdict INEFilterVerdict)

	// Topic: Updating filter verdicts

	// Updates the verdict for a flow outside the context of any filter data provider callback.
	UpdateFlowUsingVerdictForDirection(flow INEFilterSocketFlow, verdict INEFilterDataVerdict, direction NETrafficDirection)
}

// Init initializes the instance.
func (f NEFilterDataProvider) Init() NEFilterDataProvider {
	rv := objc.Send[NEFilterDataProvider](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterDataProvider) Autorelease() NEFilterDataProvider {
	rv := objc.Send[NEFilterDataProvider](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterDataProvider creates a new NEFilterDataProvider instance.
func NewNEFilterDataProvider() NEFilterDataProvider {
	class := getNEFilterDataProviderClass()
	rv := objc.Send[NEFilterDataProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Make a filtering decision for a newly-created flow of network content.
//
// flow: An [NEFilterFlow] object containing information about the new flow.
//
// # Return Value
// 
// An [NEFilterNewFlowVerdict] object indicating how the system should handle
// the flow.
//
// # Discussion
// 
// This function is called by the system when a filtering decision needs to be
// made about a new flow of network content.
// 
// [NEFilterDataProvider] subclasses must override this method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleNewFlow(_:)
func (f NEFilterDataProvider) HandleNewFlow(flow INEFilterFlow) INEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("handleNewFlow:"), flow)
	return NEFilterNewFlowVerdictFromID(rv)
}
// Make a filtering decision about a chunk of inbound data.
//
// flow: An [NEFilterFlow] object containing information about the flow.
//
// offset: An unsigned integer containing the offset of the data stored in
// `readBytes`. This offset is measured from the beginning of the flow’s
// inbound data.
//
// readBytes: An [NSData] object containing the data to filter. For non-UDP/TCP flows,
// since the data may optionally include the IP header, `readBytes` includes a
// 4-byte [NEFilterDataAttribute] field preceding the user data. Your handler
// must examine the [NEFilterDataAttribute] field and handle the data
// accordingly.
// //
// [NEFilterDataAttribute]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataAttribute
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// # Return Value
// 
// A [NEFilterDataVerdict] object indicating how the system should handle the
// chunk of data and all subsequent inbound data for the flow.
//
// # Discussion
// 
// [NEFilterDataProvider] subclasses must override this method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleInboundData(from:readBytesStartOffset:readBytes:)
func (f NEFilterDataProvider) HandleInboundDataFromFlowReadBytesStartOffsetReadBytes(flow INEFilterFlow, offset uint, readBytes foundation.INSData) INEFilterDataVerdict {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("handleInboundDataFromFlow:readBytesStartOffset:readBytes:"), flow, offset, readBytes)
	return NEFilterDataVerdictFromID(rv)
}
// Make a filtering decision about a chunk of outbound data.
//
// flow: An [NEFilterFlow] object containing information about the flow.
//
// offset: An unsigned integer containing the offset of the data stored in
// `readBytes`. This offset is measured from the beginning of the flow’s
// outbound data.
//
// readBytes: An [NSData] object containing the data to be filtered.
// //
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
//
// # Return Value
// 
// An [NEFilterDataVerdict] indicating how the system should handle the chunk
// of data and all subsequent outbound data for the flow.
//
// # Discussion
// 
// [NEFilterDataProvider] subclasses must override this method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleOutboundData(from:readBytesStartOffset:readBytes:)
func (f NEFilterDataProvider) HandleOutboundDataFromFlowReadBytesStartOffsetReadBytes(flow INEFilterFlow, offset uint, readBytes foundation.INSData) INEFilterDataVerdict {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("handleOutboundDataFromFlow:readBytesStartOffset:readBytes:"), flow, offset, readBytes)
	return NEFilterDataVerdictFromID(rv)
}
// Make a filtering decision after seeing all of the inbound data for a flow.
//
// flow: An [NEFilterFlow] object containing information about the flow.
//
// # Return Value
// 
// An [NEFilterDataVerdict] object indicating how the system should handle the
// flow of network content.
//
// # Discussion
// 
// The system calls this method after all of the inbound data for a flow of
// network content has been given to the Filter Data Provider.
// 
// [NEFilterDataProvider] subclasses must override this method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleInboundDataComplete(for:)
func (f NEFilterDataProvider) HandleInboundDataCompleteForFlow(flow INEFilterFlow) INEFilterDataVerdict {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("handleInboundDataCompleteForFlow:"), flow)
	return NEFilterDataVerdictFromID(rv)
}
// Make a filtering decision after seeing all of the outbound data for a flow.
//
// flow: An [NEFilterFlow] object containing information about the flow.
//
// # Return Value
// 
// An [NEFilterDataVerdict] object indicating how the system should handle the
// flow of network content.
//
// # Discussion
// 
// The system calls this method after all of the outbound data for a flow of
// network content has been given to the Filter Data Provider.
// 
// [NEFilterDataProvider] subclasses must override this method.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleOutboundDataComplete(for:)
func (f NEFilterDataProvider) HandleOutboundDataCompleteForFlow(flow INEFilterFlow) INEFilterDataVerdict {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("handleOutboundDataCompleteForFlow:"), flow)
	return NEFilterDataVerdictFromID(rv)
}
// Applies a set of filtering rules associated with the provider and changes
// the default filtering action.
//
// settings: A [NEFilterSettings] object containing the filter settings to apply to the
// system. Pass `nil` to revert to the default settings, which are an empty
// list of rules and a default action of [NEFilterAction.filterData].
// //
// [NEFilterAction.filterData]: https://developer.apple.com/documentation/NetworkExtension/NEFilterAction/filterData
//
// completionHandler: A Swift closure or ObjectiveC block that executes when the system finishes
// applying the settings. It receives an [NSError] parameter; a non-`nil`
// value that indicates there’s an error contidition.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/apply(_:completionHandler:)
func (f NEFilterDataProvider) ApplySettingsCompletionHandler(settings INEFilterSettings, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("applySettings:completionHandler:"), settings, _block1)
}
// Resumes a previously-paused flow.
//
// # Discussion
// 
// The provider calls this method to resume a flow that the provider
// previously paused by returning a pause verdict.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/resumeFlow(_:with:)
func (f NEFilterDataProvider) ResumeFlowWithVerdict(flow INEFilterFlow, verdict INEFilterVerdict) {
	objc.Send[objc.ID](f.ID, objc.Sel("resumeFlow:withVerdict:"), flow, verdict)
}
// Updates the verdict for a flow outside the context of any filter data
// provider callback.
//
// flow: The NEFilterSocketFlow to update the verdict for.
//
// verdict: An [NEFilterDataVerdict] instance. This must be an [AllowVerdict] or
// [DropVerdict] verdict, or a data verdict created with the Swift initializer
// or ObjectiveC type method, [DataVerdictWithPassBytesPeekBytes].
//
// direction: The direction to which the verdict applies. Pass [NETrafficDirection.any]
// to update the verdict for both the inbound and outbound directions. This
// parameter has no effect if the verdict is [DropVerdict].
// //
// [NETrafficDirection.any]: https://developer.apple.com/documentation/NetworkExtension/NETrafficDirection/any
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/update(_:using:for:)
func (f NEFilterDataProvider) UpdateFlowUsingVerdictForDirection(flow INEFilterSocketFlow, verdict INEFilterDataVerdict, direction NETrafficDirection) {
	objc.Send[objc.ID](f.ID, objc.Sel("updateFlow:usingVerdict:forDirection:"), flow, verdict, direction)
}

// ApplySettings is a synchronous wrapper around [NEFilterDataProvider.ApplySettingsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NEFilterDataProvider) ApplySettings(ctx context.Context, settings INEFilterSettings) error {
	done := make(chan error, 1)
	f.ApplySettingsCompletionHandler(settings, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

