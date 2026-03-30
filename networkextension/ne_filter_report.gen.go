// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterReport] class.
var (
	_NEFilterReportClass     NEFilterReportClass
	_NEFilterReportClassOnce sync.Once
)

func getNEFilterReportClass() NEFilterReportClass {
	_NEFilterReportClassOnce.Do(func() {
		_NEFilterReportClass = NEFilterReportClass{class: objc.GetClass("NEFilterReport")}
	})
	return _NEFilterReportClass
}

// GetNEFilterReportClass returns the class object for NEFilterReport.
func GetNEFilterReportClass() NEFilterReportClass {
	return getNEFilterReportClass()
}

type NEFilterReportClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterReportClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterReportClass) Alloc() NEFilterReport {
	rv := objc.Send[NEFilterReport](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The report of the data provider’s action on a flow.
//
// # Overview
//
// The system issues a report by calling your control provider’s
// [HandleReport] method with a report instance when the data provider issues
// a verdict whose [NEFilterReport.ShouldReport] property is set to true.
//
// # Getting report properties
//
//   - [NEFilterReport.Flow]: The flow on which the associated action was taken.
//   - [NEFilterReport.Action]: The action taken on the reported flow.
//   - [NEFilterReport.Event]: The type of event indicated by this report.
//   - [NEFilterReport.BytesInboundCount]: The number of inbound bytes received from the flow.
//   - [NEFilterReport.BytesOutboundCount]: The number of outbound bytes sent on the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport
type NEFilterReport struct {
	objectivec.Object
}

// NEFilterReportFromID constructs a [NEFilterReport] from an objc.ID.
//
// The report of the data provider’s action on a flow.
func NEFilterReportFromID(id objc.ID) NEFilterReport {
	return NEFilterReport{objectivec.Object{ID: id}}
}

// NOTE: NEFilterReport adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterReport] class.
//
// # Getting report properties
//
//   - [INEFilterReport.Flow]: The flow on which the associated action was taken.
//   - [INEFilterReport.Action]: The action taken on the reported flow.
//   - [INEFilterReport.Event]: The type of event indicated by this report.
//   - [INEFilterReport.BytesInboundCount]: The number of inbound bytes received from the flow.
//   - [INEFilterReport.BytesOutboundCount]: The number of outbound bytes sent on the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport
type INEFilterReport interface {
	objectivec.IObject

	// Topic: Getting report properties

	// The flow on which the associated action was taken.
	Flow() INEFilterFlow
	// The action taken on the reported flow.
	Action() NEFilterAction
	// The type of event indicated by this report.
	Event() NEFilterReportEvent
	// The number of inbound bytes received from the flow.
	BytesInboundCount() uint
	// The number of outbound bytes sent on the flow.
	BytesOutboundCount() uint

	// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
	ShouldReport() bool
	SetShouldReport(value bool)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NEFilterReport) Init() NEFilterReport {
	rv := objc.Send[NEFilterReport](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterReport) Autorelease() NEFilterReport {
	rv := objc.Send[NEFilterReport](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterReport creates a new NEFilterReport instance.
func NewNEFilterReport() NEFilterReport {
	class := getNEFilterReportClass()
	rv := objc.Send[NEFilterReport](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (f NEFilterReport) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The flow on which the associated action was taken.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/flow
func (f NEFilterReport) Flow() INEFilterFlow {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("flow"))
	return NEFilterFlowFromID(objc.ID(rv))
}

// The action taken on the reported flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/action
func (f NEFilterReport) Action() NEFilterAction {
	rv := objc.Send[NEFilterAction](f.ID, objc.Sel("action"))
	return NEFilterAction(rv)
}

// The type of event indicated by this report.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/event-swift.property
func (f NEFilterReport) Event() NEFilterReportEvent {
	rv := objc.Send[NEFilterReportEvent](f.ID, objc.Sel("event"))
	return NEFilterReportEvent(rv)
}

// The number of inbound bytes received from the flow.
//
// # Discussion
//
// This property is only non-zero when the report [Event] is
// [NEFilterReportEventFlowClosed].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/bytesInboundCount
func (f NEFilterReport) BytesInboundCount() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("bytesInboundCount"))
	return rv
}

// The number of outbound bytes sent on the flow.
//
// # Discussion
//
// This property is only non-zero when the report [Event] is
// [NEFilterReportEventFlowClosed].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterReport/bytesOutboundCount
func (f NEFilterReport) BytesOutboundCount() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("bytesOutboundCount"))
	return rv
}

// A Boolean value that indicates whether to send a report to the control
// provider when processing this verdict.
//
// See: https://developer.apple.com/documentation/networkextension/nefilterverdict/shouldreport
func (f NEFilterReport) ShouldReport() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("shouldReport"))
	return rv
}
func (f NEFilterReport) SetShouldReport(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setShouldReport:"), value)
}
