// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterVerdict] class.
var (
	_NEFilterVerdictClass     NEFilterVerdictClass
	_NEFilterVerdictClassOnce sync.Once
)

func getNEFilterVerdictClass() NEFilterVerdictClass {
	_NEFilterVerdictClassOnce.Do(func() {
		_NEFilterVerdictClass = NEFilterVerdictClass{class: objc.GetClass("NEFilterVerdict")}
	})
	return _NEFilterVerdictClass
}

// GetNEFilterVerdictClass returns the class object for NEFilterVerdict.
func GetNEFilterVerdictClass() NEFilterVerdictClass {
	return getNEFilterVerdictClass()
}

type NEFilterVerdictClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterVerdictClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterVerdictClass) Alloc() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class for filter verdict classes.
//
// # Overview
//
// Filter providers use instances this class to inform the system about how to
// handle flows of network data.
//
// # Configuring report generation
//
//   - [NEFilterVerdict.ShouldReport]: A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
//   - [NEFilterVerdict.SetShouldReport]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterVerdict
type NEFilterVerdict struct {
	objectivec.Object
}

// NEFilterVerdictFromID constructs a [NEFilterVerdict] from an objc.ID.
//
// The abstract base class for filter verdict classes.
func NEFilterVerdictFromID(id objc.ID) NEFilterVerdict {
	return NEFilterVerdict{objectivec.Object{ID: id}}
}

// NOTE: NEFilterVerdict adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterVerdict] class.
//
// # Configuring report generation
//
//   - [INEFilterVerdict.ShouldReport]: A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
//   - [INEFilterVerdict.SetShouldReport]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterVerdict
type INEFilterVerdict interface {
	objectivec.IObject

	// Topic: Configuring report generation

	// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
	ShouldReport() bool
	SetShouldReport(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NEFilterVerdict) Init() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterVerdict) Autorelease() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterVerdict creates a new NEFilterVerdict instance.
func NewNEFilterVerdict() NEFilterVerdict {
	class := getNEFilterVerdictClass()
	rv := objc.Send[NEFilterVerdict](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (f NEFilterVerdict) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether to send a report to the control
// provider when processing this verdict.
//
// # Discussion
//
// If the property is equal to true, the system sends a report to the control
// provider’s [HandleReport] method when processing this verdict in the data
// provider. This property has no effect if the verdict originates in the
// control provider.
//
// The data provider doesn’t need to wait for a response from the control
// provider before continuing to process the flow. Therefore, calling the
// [HandleReport] method is a more efficient way to report a flow to the
// control provider than returning a [NeedRulesVerdict] verdict.
//
// This property applies when the action taken on a flow is
// [NEFilterActionAllow], [NEFilterActionDrop], [NEFilterActionRemediate], or
// [NEFilterActionFilterData] (the last of which is only for new flows).
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterVerdict/shouldReport
func (f NEFilterVerdict) ShouldReport() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("shouldReport"))
	return rv
}
func (f NEFilterVerdict) SetShouldReport(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setShouldReport:"), value)
}
