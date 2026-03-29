// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEFilterDataVerdict] class.
var (
	_NEFilterDataVerdictClass     NEFilterDataVerdictClass
	_NEFilterDataVerdictClassOnce sync.Once
)

func getNEFilterDataVerdictClass() NEFilterDataVerdictClass {
	_NEFilterDataVerdictClassOnce.Do(func() {
		_NEFilterDataVerdictClass = NEFilterDataVerdictClass{class: objc.GetClass("NEFilterDataVerdict")}
	})
	return _NEFilterDataVerdictClass
}

// GetNEFilterDataVerdictClass returns the class object for NEFilterDataVerdict.
func GetNEFilterDataVerdictClass() NEFilterDataVerdictClass {
	return getNEFilterDataVerdictClass()
}

type NEFilterDataVerdictClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterDataVerdictClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterDataVerdictClass) Alloc() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The result from a filter data provder for subsequent chunks of data on a
// flow.
//
// # Overview
// 
// Return this verdict type from the various methods of
// [NEFilterDataProvider].
//
// # Reporting statistics
//
//   - [NEFilterDataVerdict.StatisticsReportFrequency]: The frequencty at which to provide flow statistics to the data provider.
//   - [NEFilterDataVerdict.SetStatisticsReportFrequency]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict
type NEFilterDataVerdict struct {
	NEFilterVerdict
}

// NEFilterDataVerdictFromID constructs a [NEFilterDataVerdict] from an objc.ID.
//
// The result from a filter data provder for subsequent chunks of data on a
// flow.
func NEFilterDataVerdictFromID(id objc.ID) NEFilterDataVerdict {
	return NEFilterDataVerdict{NEFilterVerdict: NEFilterVerdictFromID(id)}
}
// NOTE: NEFilterDataVerdict adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterDataVerdict] class.
//
// # Reporting statistics
//
//   - [INEFilterDataVerdict.StatisticsReportFrequency]: The frequencty at which to provide flow statistics to the data provider.
//   - [INEFilterDataVerdict.SetStatisticsReportFrequency]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict
type INEFilterDataVerdict interface {
	INEFilterVerdict

	// Topic: Reporting statistics

	// The frequencty at which to provide flow statistics to the data provider.
	StatisticsReportFrequency() NEFilterReportFrequency
	SetStatisticsReportFrequency(value NEFilterReportFrequency)
}

// Init initializes the instance.
func (f NEFilterDataVerdict) Init() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterDataVerdict) Autorelease() NEFilterDataVerdict {
	rv := objc.Send[NEFilterDataVerdict](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterDataVerdict creates a new NEFilterDataVerdict instance.
func NewNEFilterDataVerdict() NEFilterDataVerdict {
	class := getNEFilterDataVerdictClass()
	rv := objc.Send[NEFilterDataVerdict](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a verdict that tells the system to pass a chunk of network data to
// its final destination, and specifies the next chunk of data to provide.
//
// passBytes: The number of bytes to pass to its final destination.
//
// peekBytes: The number of bytes after the end of the `passBytes` that the Filter Data
// Provider expects in the next call to
// [HandleOutboundDataFromFlowReadBytesStartOffsetReadBytes] or
// [HandleInboundDataFromFlowReadBytesStartOffsetReadBytes]. The Filter Data
// Provider uses this chunk of data to make its next filtering decision.
// 
// To see all subsequent bytes, set this parameter to [NEFilterFlowBytesMax].
// //
// [NEFilterFlowBytesMax]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlowBytesMax
//
// # Return Value
// 
// A [NEFilterDataVerdict] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/init(passBytes:peekBytes:)
func NewFilterDataVerdictWithPassBytesPeekBytes(passBytes uint, peekBytes uint) NEFilterDataVerdict {
	rv := objc.Send[objc.ID](objc.ID(getNEFilterDataVerdictClass().class), objc.Sel("dataVerdictWithPassBytes:peekBytes:"), passBytes, peekBytes)
	return NEFilterDataVerdictFromID(rv)
}

// Creates a verdict that tells the system to pass the current chunk of
// network data and all subsequent data for the current flow to its final
// destination.
//
// # Return Value
// 
// A [NEFilterDataVerdict] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/allow()
func (_NEFilterDataVerdictClass NEFilterDataVerdictClass) AllowVerdict() NEFilterDataVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterDataVerdictClass.class), objc.Sel("allowVerdict"))
	return NEFilterDataVerdictFromID(rv)
}
// Creates a verdict that tells the system to drop the current chunk of
// network data and all subsequent data for the current flow.
//
// # Return Value
// 
// A [NEFilterDataVerdict] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/drop()
func (_NEFilterDataVerdictClass NEFilterDataVerdictClass) DropVerdict() NEFilterDataVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterDataVerdictClass.class), objc.Sel("dropVerdict"))
	return NEFilterDataVerdictFromID(rv)
}
// Creates a verdict that tells the system to pause the flow.
//
// # Return Value
// 
// A [NEFilterDataVerdict] object.
//
// # Discussion
// 
// After pausing the flow, the system doesn’t call any of the data
// provider’s handler callbacks until you resume the flow by calling
// [ResumeFlowWithVerdict].
// 
// You can pause TCP flows indefinitely. You can pause UDP flows for up to 10
// seconds, after which the system drops the flow. Pausing a flow that’s
// already paused is an invalid operation.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/pause()
func (_NEFilterDataVerdictClass NEFilterDataVerdictClass) PauseVerdict() NEFilterDataVerdict {
	rv := objc.Send[objc.ID](objc.ID(_NEFilterDataVerdictClass.class), objc.Sel("pauseVerdict"))
	return NEFilterDataVerdictFromID(rv)
}

// The frequencty at which to provide flow statistics to the data provider.
//
// # Discussion
// 
// This property determines the frequency at which the provider receives a
// call to its [HandleReport] method with an [NEFilterReport.Event.statistics]
// event.
// 
// The default value of this property [NEFilterReport.Frequency.none], meaning
// that the provider receives no statistics by default.
//
// [NEFilterReport.Event.statistics]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Event-swift.enum/statistics
// [NEFilterReport.Frequency.none]: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/Frequency/none
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataVerdict/statisticsReportFrequency
func (f NEFilterDataVerdict) StatisticsReportFrequency() NEFilterReportFrequency {
	rv := objc.Send[NEFilterReportFrequency](f.ID, objc.Sel("statisticsReportFrequency"))
	return NEFilterReportFrequency(rv)
}
func (f NEFilterDataVerdict) SetStatisticsReportFrequency(value NEFilterReportFrequency) {
	objc.Send[struct{}](f.ID, objc.Sel("setStatisticsReportFrequency:"), value)
}

