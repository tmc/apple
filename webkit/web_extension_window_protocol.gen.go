// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol with methods that represent a window to web extensions.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow
type WKWebExtensionWindow interface {
	objectivec.IObject
}

// WKWebExtensionWindowObject wraps an existing Objective-C object that conforms to the WKWebExtensionWindow protocol.
type WKWebExtensionWindowObject struct {
	objectivec.Object
}

func (o WKWebExtensionWindowObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKWebExtensionWindowObjectFromID constructs a [WKWebExtensionWindowObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKWebExtensionWindowObjectFromID(id objc.ID) WKWebExtensionWindowObject {
	return WKWebExtensionWindowObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when the active tab is needed for the window.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to `nil` if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/activeTab(for:)
func (o WKWebExtensionWindowObject) ActiveTabForWebExtensionContext(context IWKWebExtensionContext) WKWebExtensionTab {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("activeTabForWebExtensionContext:"), context)
	return WKWebExtensionTabObjectFromID(rv)
}

// Called to close the window.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/close(for:completionHandler:)
func (o WKWebExtensionWindowObject) CloseForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("closeForWebExtensionContext:completionHandler:"), context, completionHandler)
}

// Called to focus the window.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// No action is performed if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/focus(for:completionHandler:)
func (o WKWebExtensionWindowObject) FocusForWebExtensionContextCompletionHandler(context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("focusForWebExtensionContext:completionHandler:"), context, completionHandler)
}

// Called when the frame of the window is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [CGRectNull] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/frame(for:)
//
// [CGRectNull]: https://developer.apple.com/documentation/CoreGraphics/CGRectNull
func (o WKWebExtensionWindowObject) FrameForWebExtensionContext(context IWKWebExtensionContext) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("frameForWebExtensionContext:"), context)
	return rv
}

// Called when the private state of the window is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to [NO] if not implemented. This value is cached and will not
// change for the duration of the window or its contained tabs.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/isPrivate(for:)
func (o WKWebExtensionWindowObject) IsPrivateForWebExtensionContext(context IWKWebExtensionContext) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPrivateForWebExtensionContext:"), context)
	return rv
}

// Called when the screen frame containing the window is needed.
//
// context: The context associated with the running web extension.
//
// # Discussion
//
// Defaults to [CGRectNull] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/screenFrame(for:)
//
// [CGRectNull]: https://developer.apple.com/documentation/CoreGraphics/CGRectNull
func (o WKWebExtensionWindowObject) ScreenFrameForWebExtensionContext(context IWKWebExtensionContext) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("screenFrameForWebExtensionContext:"), context)
	return rv
}

// Called to set the frame of the window.
//
// frame: The new frame of the window, in screen coordinates.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// On macOS, the implementation of both [FrameForWebExtensionContext] and
// [ScreenFrameForWebExtensionContext] are prerequisites. On iOS, iPadOS, and
// visionOS, only [FrameForWebExtensionContext] is a prerequisite. Without the
// respective method(s), this method will not be called.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/setFrame(_:for:completionHandler:)
func (o WKWebExtensionWindowObject) SetFrameForWebExtensionContextCompletionHandler(frame corefoundation.CGRect, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setFrame:forWebExtensionContext:completionHandler:"), frame, context, completionHandler)
}

// Called to set the state of the window.
//
// state: The new state of the window.
//
// context: The context in which the web extension is running.
//
// completionHandler: A block that must be called upon completion. It takes a single error
// argument, which should be provided if any errors occurred.
//
// # Discussion
//
// The implementation of [WindowStateForWebExtensionContext] is a
// prerequisite.
//
// Without it, this method will not be called.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/setWindowState(_:for:completionHandler:)
func (o WKWebExtensionWindowObject) SetWindowStateForWebExtensionContextCompletionHandler(state WKWebExtensionWindowState, context IWKWebExtensionContext, completionHandler ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setWindowState:forWebExtensionContext:completionHandler:"), state, context, completionHandler)
}

// Called when the tabs are needed for the window.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to an empty array if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/tabs(for:)
func (o WKWebExtensionWindowObject) TabsForWebExtensionContext(context IWKWebExtensionContext) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("tabsForWebExtensionContext:"), context)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Called when the state of the window is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to[WKWebExtensionWindowStateNormal] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/windowState(for:)
func (o WKWebExtensionWindowObject) WindowStateForWebExtensionContext(context IWKWebExtensionContext) WKWebExtensionWindowState {
	rv := objc.Send[WKWebExtensionWindowState](o.ID, objc.Sel("windowStateForWebExtensionContext:"), context)
	return rv
}

// Called when the type of the window is needed.
//
// context: The context in which the web extension is running.
//
// # Discussion
//
// Defaults to[WKWebExtensionWindowTypeNormal] if not implemented.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionWindow/windowType(for:)
func (o WKWebExtensionWindowObject) WindowTypeForWebExtensionContext(context IWKWebExtensionContext) WKWebExtensionWindowType {
	rv := objc.Send[WKWebExtensionWindowType](o.ID, objc.Sel("windowTypeForWebExtensionContext:"), context)
	return rv
}
