// Code generated from Apple documentation for Foundation. DO NOT EDIT.
//go:build ios
// +build ios

package foundation

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// Returns a human-readable string describing the data that SiriKit displays
// to the user when you handle an intent.
//
// # Discussion
//
// If you provide an Intents UI app extension, you can customize all or some
// of the interface that SiriKit displays to the user for a given intent. The
// information displayed by SiriKit is different for each intent, and can
// change in the future. During development, use this method to retrieve a
// human-readable description of the contents of the [INParameter] objects
// that SiriKit intends to display for the current intent. Use that
// information to plan your custom interface.
//
// For information about customizing the Siri and Maps interfaces, see
// [Creating an Intents App Extension].
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/interfaceParametersDescription()
//
// [Creating an Intents App Extension]: https://developer.apple.com/documentation/SiriKit/creating-an-intents-app-extension
// [INParameter]: https://developer.apple.com/documentation/Intents/INParameter
func (e NSExtensionContext) InterfaceParametersDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("interfaceParametersDescription"))
	return NSStringFromID(rv).String()
}

// Metadata for populating your share extensions interface.
//
// # Discussion
//
// When the user selects an app from the list of suggested apps in iOS’s
// share sheet, this property contains metadata that you can use to populate
// your share extensions interface. The source for the metadata is the
// [INSendMessageIntent] of your messaging app.
//
// This property is `nil` if your app’s share extension wasn’t launched
// from the list of suggested apps.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/intent
//
// [INSendMessageIntent]: https://developer.apple.com/documentation/Intents/INSendMessageIntent
func (e NSExtensionContext) Intent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("intent"))
	return rv
}

// The minimum size for a Siri hosted view.
//
// # Discussion
//
// Apps can customize the Siri interface using an Intents UI extension. The
// extension vends a view controller whose view contains the custom content
// that you want Siri to display. The size of that view controller’s view
// must be at least as large as the size value in this property.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/hostedViewMinimumAllowedSize
func (e NSExtensionContext) HostedViewMinimumAllowedSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](e.ID, objc.Sel("hostedViewMinimumAllowedSize"))
	return corefoundation.CGSize(rv)
}

// The maximum size for a Siri hosted view.
//
// # Discussion
//
// Apps can customize the Siri interface using an Intents UI extension. The
// extension vends a view controller whose view contains the custom content
// that you want Siri to display. The size of that view controller’s view
// must be no larger than the size value in this property.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/hostedViewMaximumAllowedSize
func (e NSExtensionContext) HostedViewMaximumAllowedSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](e.ID, objc.Sel("hostedViewMaximumAllowedSize"))
	return corefoundation.CGSize(rv)
}

// The active display mode of the widget.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/widgetActiveDisplayMode
func (e NSExtensionContext) WidgetActiveDisplayMode() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("widgetActiveDisplayMode"))
	return rv
}

// The largest display mode the widget supports.
//
// # Discussion
//
// The default value of this property is [NCWidgetDisplayMode.compact]. At any
// time, you can change the largest display mode your widget supports by
// changing the value of this property. For example, you can update the
// property value as more or less content is available to display in your
// widget.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/widgetLargestAvailableDisplayMode
//
// [NCWidgetDisplayMode.compact]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetDisplayMode/compact
func (e NSExtensionContext) WidgetLargestAvailableDisplayMode() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("widgetLargestAvailableDisplayMode"))
	return rv
}
func (e NSExtensionContext) SetWidgetLargestAvailableDisplayMode(value uint) {
	objc.Send[struct{}](e.ID, objc.Sel("setWidgetLargestAvailableDisplayMode:"), value)
}
