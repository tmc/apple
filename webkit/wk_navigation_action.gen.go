// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKNavigationAction] class.
var (
	_WKNavigationActionClass     WKNavigationActionClass
	_WKNavigationActionClassOnce sync.Once
)

func getWKNavigationActionClass() WKNavigationActionClass {
	_WKNavigationActionClassOnce.Do(func() {
		_WKNavigationActionClass = WKNavigationActionClass{class: objc.GetClass("WKNavigationAction")}
	})
	return _WKNavigationActionClass
}

// GetWKNavigationActionClass returns the class object for WKNavigationAction.
func GetWKNavigationActionClass() WKNavigationActionClass {
	return getWKNavigationActionClass()
}

type WKNavigationActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKNavigationActionClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKNavigationActionClass) Alloc() WKNavigationAction {
	rv := objc.Send[WKNavigationAction](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains information about an action that causes navigation
// to occur.
//
// # Overview
//
// Use a [WKNavigationAction] object to make policy decisions about whether to
// allow navigation within your app’s web view. You don’t create
// [WKNavigationAction] objects directly. Instead, the web view creates them
// and delivers them to the appropriate delegate objects. Use the methods of
// your delegate to analyze the action and determine whether to allow the
// resulting navigation to occur.
//
// # Getting the navigation type
//
//   - [WKNavigationAction.NavigationType]: The type of action that triggered the navigation.
//
// # Inspecting navigation information
//
//   - [WKNavigationAction.Request]: The URL request object associated with the navigation action.
//   - [WKNavigationAction.SourceFrame]: The frame that requested the navigation.
//   - [WKNavigationAction.TargetFrame]: The frame in which to display the new content.
//   - [WKNavigationAction.ShouldPerformDownload]: A Boolean value that indicates whether the web content provided an attribute that indicates a download.
//
// # Inspecting user actions
//
//   - [WKNavigationAction.ButtonNumber]: The number of the mouse button that caused the navigation request.
//   - [WKNavigationAction.ModifierFlags]: The modifier keys that were pressed at the time of the navigation request.
//
// # Instance Properties
//
//   - [WKNavigationAction.IsContentRuleListRedirect]
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction
type WKNavigationAction struct {
	objectivec.Object
}

// WKNavigationActionFromID constructs a [WKNavigationAction] from an objc.ID.
//
// An object that contains information about an action that causes navigation
// to occur.
func WKNavigationActionFromID(id objc.ID) WKNavigationAction {
	return WKNavigationAction{objectivec.Object{ID: id}}
}

// NOTE: WKNavigationAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKNavigationAction] class.
//
// # Getting the navigation type
//
//   - [IWKNavigationAction.NavigationType]: The type of action that triggered the navigation.
//
// # Inspecting navigation information
//
//   - [IWKNavigationAction.Request]: The URL request object associated with the navigation action.
//   - [IWKNavigationAction.SourceFrame]: The frame that requested the navigation.
//   - [IWKNavigationAction.TargetFrame]: The frame in which to display the new content.
//   - [IWKNavigationAction.ShouldPerformDownload]: A Boolean value that indicates whether the web content provided an attribute that indicates a download.
//
// # Inspecting user actions
//
//   - [IWKNavigationAction.ButtonNumber]: The number of the mouse button that caused the navigation request.
//   - [IWKNavigationAction.ModifierFlags]: The modifier keys that were pressed at the time of the navigation request.
//
// # Instance Properties
//
//   - [IWKNavigationAction.IsContentRuleListRedirect]
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction
type IWKNavigationAction interface {
	objectivec.IObject

	// Topic: Getting the navigation type

	// The type of action that triggered the navigation.
	NavigationType() WKNavigationType

	// Topic: Inspecting navigation information

	// The URL request object associated with the navigation action.
	Request() foundation.NSURLRequest
	// The frame that requested the navigation.
	SourceFrame() IWKFrameInfo
	// The frame in which to display the new content.
	TargetFrame() IWKFrameInfo
	// A Boolean value that indicates whether the web content provided an attribute that indicates a download.
	ShouldPerformDownload() bool

	// Topic: Inspecting user actions

	// The number of the mouse button that caused the navigation request.
	ButtonNumber() objectivec.IObject
	// The modifier keys that were pressed at the time of the navigation request.
	ModifierFlags() objectivec.IObject

	// Topic: Instance Properties

	IsContentRuleListRedirect() bool
}

// Init initializes the instance.
func (n WKNavigationAction) Init() WKNavigationAction {
	rv := objc.Send[WKNavigationAction](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n WKNavigationAction) Autorelease() WKNavigationAction {
	rv := objc.Send[WKNavigationAction](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKNavigationAction creates a new WKNavigationAction instance.
func NewWKNavigationAction() WKNavigationAction {
	class := getWKNavigationActionClass()
	rv := objc.Send[WKNavigationAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type of action that triggered the navigation.
//
// # Discussion
//
// The value is one of the constants of the enumerated type
// [WKNavigationType].
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/navigationType
//
// [WKNavigationType]: https://developer.apple.com/documentation/WebKit/WKNavigationType
func (n WKNavigationAction) NavigationType() WKNavigationType {
	rv := objc.Send[WKNavigationType](n.ID, objc.Sel("navigationType"))
	return WKNavigationType(rv)
}

// The URL request object associated with the navigation action.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/request
func (n WKNavigationAction) Request() foundation.NSURLRequest {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("request"))
	return foundation.NSURLRequestFromID(objc.ID(rv))
}

// The frame that requested the navigation.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/sourceFrame
func (n WKNavigationAction) SourceFrame() IWKFrameInfo {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("sourceFrame"))
	return WKFrameInfoFromID(objc.ID(rv))
}

// The frame in which to display the new content.
//
// # Discussion
//
// If the target of the navigation is a new window, this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/targetFrame
func (n WKNavigationAction) TargetFrame() IWKFrameInfo {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("targetFrame"))
	return WKFrameInfoFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the web content provided an
// attribute that indicates a download.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/shouldPerformDownload
func (n WKNavigationAction) ShouldPerformDownload() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("shouldPerformDownload"))
	return rv
}

// The number of the mouse button that caused the navigation request.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/buttonNumber
func (n WKNavigationAction) ButtonNumber() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("buttonNumber"))
	return objectivec.Object{ID: rv}
}

// The modifier keys that were pressed at the time of the navigation request.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/modifierFlags
func (n WKNavigationAction) ModifierFlags() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("modifierFlags"))
	return objectivec.Object{ID: rv}
}

// # Discussion
//
// Whether or not the navigation is a redirect from a content rule list.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationAction/isContentRuleListRedirect
func (n WKNavigationAction) IsContentRuleListRedirect() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isContentRuleListRedirect"))
	return rv
}
