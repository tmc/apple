// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NEFilterNewFlowVerdict] class.
var (
	_NEFilterNewFlowVerdictClass     NEFilterNewFlowVerdictClass
	_NEFilterNewFlowVerdictClassOnce sync.Once
)

func getNEFilterNewFlowVerdictClass() NEFilterNewFlowVerdictClass {
	_NEFilterNewFlowVerdictClassOnce.Do(func() {
		_NEFilterNewFlowVerdictClass = NEFilterNewFlowVerdictClass{class: objc.GetClass("NEFilterNewFlowVerdict")}
	})
	return _NEFilterNewFlowVerdictClass
}

// GetNEFilterNewFlowVerdictClass returns the class object for NEFilterNewFlowVerdict.
func GetNEFilterNewFlowVerdictClass() NEFilterNewFlowVerdictClass {
	return getNEFilterNewFlowVerdictClass()
}

type NEFilterNewFlowVerdictClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterNewFlowVerdictClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterNewFlowVerdictClass) Alloc() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The result from a filter data provder after the initial examination of a
// flow.
//
// # Inspecting new flow verdict properties
//
//   - [NEFilterNewFlowVerdict.StatisticsReportFrequency]: The frequency at which the data provider receives reports.
//   - [NEFilterNewFlowVerdict.SetStatisticsReportFrequency]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict
type NEFilterNewFlowVerdict struct {
	NEFilterVerdict
}

// NEFilterNewFlowVerdictFromID constructs a [NEFilterNewFlowVerdict] from an objc.ID.
//
// The result from a filter data provder after the initial examination of a
// flow.
func NEFilterNewFlowVerdictFromID(id objc.ID) NEFilterNewFlowVerdict {
	return NEFilterNewFlowVerdict{NEFilterVerdict: NEFilterVerdictFromID(id)}
}

// NOTE: NEFilterNewFlowVerdict adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterNewFlowVerdict] class.
//
// # Inspecting new flow verdict properties
//
//   - [INEFilterNewFlowVerdict.StatisticsReportFrequency]: The frequency at which the data provider receives reports.
//   - [INEFilterNewFlowVerdict.SetStatisticsReportFrequency]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict
type INEFilterNewFlowVerdict interface {
	INEFilterVerdict

	// Topic: Inspecting new flow verdict properties

	// The frequency at which the data provider receives reports.
	StatisticsReportFrequency() NEFilterReportFrequency
	SetStatisticsReportFrequency(value NEFilterReportFrequency)
}

// Init initializes the instance.
func (f NEFilterNewFlowVerdict) Init() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterNewFlowVerdict) Autorelease() NEFilterNewFlowVerdict {
	rv := objc.Send[NEFilterNewFlowVerdict](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterNewFlowVerdict creates a new NEFilterNewFlowVerdict instance.
func NewNEFilterNewFlowVerdict() NEFilterNewFlowVerdict {
	class := getNEFilterNewFlowVerdictClass()
	rv := objc.Send[NEFilterNewFlowVerdict](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create a verdict that indicates to the system that the all of the new
// flow’s data should be allowed to pass to its final destination.
//
// # Return Value
//
// A [NEFilterNewFlowVerdict] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/allow()
func (_NEFilterNewFlowVerdictClass NEFilterNewFlowVerdictClass) AllowVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterNewFlowVerdictClass.class), objc.Sel("allowVerdict"))
	return NEFilterNewFlowVerdictFromID(rv)
}

// Create a verdict that indicates to the system that all of the new flow’s
// data should dropped, and the user should not be given the opportunity to
// request access.
//
// # Return Value
//
// A [NEFilterNewFlowVerdict] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/drop()
func (_NEFilterNewFlowVerdictClass NEFilterNewFlowVerdictClass) DropVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterNewFlowVerdictClass.class), objc.Sel("dropVerdict"))
	return NEFilterNewFlowVerdictFromID(rv)
}

// Creates a verdict that tells the system to pause the flow.
//
// # Discussion
//
// Once paused, the system doesn’t call any of the data provider’s handler
// callbacks until you resume the flow by calling [ResumeFlowWithVerdict].
//
// You can pause TCP flows indefinitely. You can pause UDP flows for up to 10
// seconds, after which the system drops the flow. Pausing a flow that’s
// already paused is an invalid operation.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/pause()
func (_NEFilterNewFlowVerdictClass NEFilterNewFlowVerdictClass) PauseVerdict() NEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterNewFlowVerdictClass.class), objc.Sel("pauseVerdict"))
	return NEFilterNewFlowVerdictFromID(rv)
}

// Create a verdict that indicates to the system that the filter needs to make
// a decision about a new flow after seeing a portion of the flow’s data.
//
// filterInbound: A Boolean indicating whether or not the filter needs to see inbound data
// for the flow.
//
// peekInboundBytes: The number of inbound bytes that the filter needs to see in the subsequent
// call to -[NEFilterDataProvider] “].
//
// filterOutbound: A Boolean indicating whether or not the filter needs to see outbound data
// for the flow.
//
// peekOutboundBytes: The number of outbound bytes that the filter needs to see in the subsequent
// call to -[NEFilterDataProvider] `readBytes`:].
//
// # Return Value
//
// A [NEFilterNewFlowVerdict] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/filterDataVerdict(withFilterInbound:peekInboundBytes:filterOutbound:peekOutboundBytes:)
func (_NEFilterNewFlowVerdictClass NEFilterNewFlowVerdictClass) FilterDataVerdictWithFilterInboundPeekInboundBytesFilterOutboundPeekOutboundBytes(filterInbound bool, peekInboundBytes uint, filterOutbound bool, peekOutboundBytes uint) NEFilterNewFlowVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterNewFlowVerdictClass.class), objc.Sel("filterDataVerdictWithFilterInbound:peekInboundBytes:filterOutbound:peekOutboundBytes:"), filterInbound, peekInboundBytes, filterOutbound, peekOutboundBytes)
	return NEFilterNewFlowVerdictFromID(rv)
}

// The frequency at which the data provider receives reports.
//
// # Discussion
//
// This property determines the frequency at which the system calls the data
// provider’s [HandleReport] method with an [NEFilterReport] instance that
// contains an [NEFilterReportEventStatistics] [Event].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterNewFlowVerdict/statisticsReportFrequency
func (f NEFilterNewFlowVerdict) StatisticsReportFrequency() NEFilterReportFrequency {
	rv := objc.Send[NEFilterReportFrequency](f.ID, objc.Sel("statisticsReportFrequency"))
	return NEFilterReportFrequency(rv)
}
func (f NEFilterNewFlowVerdict) SetStatisticsReportFrequency(value NEFilterReportFrequency) {
	objc.Send[struct{}](f.ID, objc.Sel("setStatisticsReportFrequency:"), value)
}
