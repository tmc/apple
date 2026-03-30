// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods implemented by delegates of [NSDatePickerCell](<doc://com.apple.appkit/documentation/AppKit/NSDatePickerCell>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCellDelegate
type NSDatePickerCellDelegate interface {
	objectivec.IObject
}

// NSDatePickerCellDelegateObject wraps an existing Objective-C object that conforms to the NSDatePickerCellDelegate protocol.
type NSDatePickerCellDelegateObject struct {
	objectivec.Object
}

func (o NSDatePickerCellDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSDatePickerCellDelegateObjectFromID constructs a [NSDatePickerCellDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSDatePickerCellDelegateObjectFromID(id objc.ID) NSDatePickerCellDelegateObject {
	return NSDatePickerCellDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The delegate receives this message each time the user attempts to change
// the receiver’s value, allowing the delegate the opportunity to override
// the change.
//
// datePickerCell: The cell that sent the message.
//
// proposedDateValue: On input, contains the proposed new date. The delegate may change this
// value before returning.
//
// proposedTimeInterval: On input, contains the proposed new time interval. The delegate may change
// this value before returning.
//
// # Discussion
//
// When returning a new `proposedDateValue`, the [NSDate] instance should be
// autoreleased, and the `proposedDateValue` should not be released by the
// delegate.
//
// The `proposedDateValue` and `proposedTimeInterval` are guaranteed to lie
// between the dates returned by [MinDate] and [MaxDate]. If you modify these
// values, you should ensure that the new values are within the appropriate
// range.
//
// See: https://developer.apple.com/documentation/AppKit/NSDatePickerCellDelegate/datePickerCell(_:validateProposedDateValue:timeInterval:)
func (o NSDatePickerCellDelegateObject) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell INSDatePickerCell, proposedDateValue foundation.INSDate, proposedTimeInterval unsafe.Pointer) {
	objc.Send[struct{}](o.ID, objc.Sel("datePickerCell:validateProposedDateValue:timeInterval:"), datePickerCell, proposedDateValue, proposedTimeInterval)
}

// NSDatePickerCellDelegateConfig holds optional typed callbacks for [NSDatePickerCellDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsdatepickercelldelegate
type NSDatePickerCellDelegateConfig struct {

	// Content Validation
	// DatePickerCellValidateProposedDateValueTimeInterval — The delegate receives this message each time the user attempts to change the receiver’s value, allowing the delegate the opportunity to override the change.
	DatePickerCellValidateProposedDateValueTimeInterval func(datePickerCell NSDatePickerCell, proposedDateValue foundation.NSDate, proposedTimeInterval *float64)
}

// NewNSDatePickerCellDelegate creates an Objective-C object implementing the [NSDatePickerCellDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSDatePickerCellDelegateObject] satisfies the [NSDatePickerCellDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsdatepickercelldelegate
func NewNSDatePickerCellDelegate(config NSDatePickerCellDelegateConfig) NSDatePickerCellDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSDatePickerCellDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DatePickerCellValidateProposedDateValueTimeInterval != nil {
		fn := config.DatePickerCellValidateProposedDateValueTimeInterval
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("datePickerCell:validateProposedDateValue:timeInterval:"),
			Fn: func(self objc.ID, _cmd objc.SEL, datePickerCellID objc.ID, proposedDateValueID objc.ID, proposedTimeInterval *float64) {
				datePickerCell := NSDatePickerCellFromID(datePickerCellID)
				proposedDateValue := foundation.NSDateFromID(proposedDateValueID)
				fn(datePickerCell, proposedDateValue, proposedTimeInterval)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSDatePickerCellDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSDatePickerCellDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSDatePickerCellDelegateObjectFromID(instance)
}
