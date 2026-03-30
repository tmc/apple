// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods implemented by the delegate of an [NSAlert](<doc://com.apple.appkit/documentation/AppKit/NSAlert>) object to respond to a user’s request for help.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlertDelegate
type NSAlertDelegate interface {
	objectivec.IObject
}

// NSAlertDelegateObject wraps an existing Objective-C object that conforms to the NSAlertDelegate protocol.
type NSAlertDelegateObject struct {
	objectivec.Object
}

func (o NSAlertDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAlertDelegateObjectFromID constructs a [NSAlertDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAlertDelegateObjectFromID(id objc.ID) NSAlertDelegateObject {
	return NSAlertDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sent to the delegate when the user clicks the alert’s help button. The
// delegate causes help to be displayed for an alert, directly or indirectly.
//
// # Return Value
//
// true when the delegate displayed help directly, false otherwise. When false
// and the alert has a help anchor ([HelpAnchor]), the application’s help
// manager displays help using the help anchor.
//
// # Discussion
//
// The delegate implements this method only to override the help-anchor lookup
// behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlertDelegate/alertShowHelp(_:)
func (o NSAlertDelegateObject) AlertShowHelp(alert NSAlert) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("alertShowHelp:"), alert)
	return rv
}

// NSAlertDelegateConfig holds optional typed callbacks for [NSAlertDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsalertdelegate
type NSAlertDelegateConfig struct {

	// Displaying Help
	// AlertShowHelp — Sent to the delegate when the user clicks the alert’s help button. The delegate causes help to be displayed for an alert, directly or indirectly.
	AlertShowHelp func(alert NSAlert) bool
}

// NewNSAlertDelegate creates an Objective-C object implementing the [NSAlertDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSAlertDelegateObject] satisfies the [NSAlertDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsalertdelegate
func NewNSAlertDelegate(config NSAlertDelegateConfig) NSAlertDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSAlertDelegate_%d", n)

	var methods []objc.MethodDef

	if config.AlertShowHelp != nil {
		fn := config.AlertShowHelp
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("alertShowHelp:"),
			Fn: func(self objc.ID, _cmd objc.SEL, alertID objc.ID) bool {
				alert := NSAlertFromID(alertID)
				return fn(alert)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSAlertDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSAlertDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSAlertDelegateObjectFromID(instance)
}
