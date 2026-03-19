// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCShareableContent] class.
var (
	_SCShareableContentClass     SCShareableContentClass
	_SCShareableContentClassOnce sync.Once
)

func getSCShareableContentClass() SCShareableContentClass {
	_SCShareableContentClassOnce.Do(func() {
		_SCShareableContentClass = SCShareableContentClass{class: objc.GetClass("SCShareableContent")}
	})
	return _SCShareableContentClass
}

// GetSCShareableContentClass returns the class object for SCShareableContent.
func GetSCShareableContentClass() SCShareableContentClass {
	return getSCShareableContentClass()
}

type SCShareableContentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCShareableContentClass) Alloc() SCShareableContent {
	rv := objc.Send[SCShareableContent](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An instance that represents a set of displays, apps, and windows that your
// app can capture.
//
// # Overview
// 
// Use the [SCShareableContent.Displays], [SCShareableContent.Windows], and [SCShareableContent.Applications] properties to create a
// [SCContentFilter] object that specifies what display content to capture.
// You apply the filter to an instance of [SCStream] to limit its output to
// only the content matching your filter.
//
// # Inspecting shareable content
//
//   - [SCShareableContent.Windows]: The windows available for capture.
//   - [SCShareableContent.Displays]: The displays available for capture.
//   - [SCShareableContent.Applications]: The apps available for capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent
type SCShareableContent struct {
	objectivec.Object
}

// SCShareableContentFromID constructs a [SCShareableContent] from an objc.ID.
//
// An instance that represents a set of displays, apps, and windows that your
// app can capture.
func SCShareableContentFromID(id objc.ID) SCShareableContent {
	return SCShareableContent{objectivec.Object{ID: id}}
}
// NOTE: SCShareableContent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCShareableContent] class.
//
// # Inspecting shareable content
//
//   - [ISCShareableContent.Windows]: The windows available for capture.
//   - [ISCShareableContent.Displays]: The displays available for capture.
//   - [ISCShareableContent.Applications]: The apps available for capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent
type ISCShareableContent interface {
	objectivec.IObject

	// Topic: Inspecting shareable content

	// The windows available for capture.
	Windows() []SCWindow
	// The displays available for capture.
	Displays() []SCDisplay
	// The apps available for capture.
	Applications() []SCRunningApplication
}

// Init initializes the instance.
func (s SCShareableContent) Init() SCShareableContent {
	rv := objc.Send[SCShareableContent](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SCShareableContent) Autorelease() SCShareableContent {
	rv := objc.Send[SCShareableContent](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCShareableContent creates a new SCShareableContent instance.
func NewSCShareableContent() SCShareableContent {
	class := getSCShareableContentClass()
	rv := objc.Send[SCShareableContent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the displays, apps, and windows that your app can capture.
//
// completionHandler: A callback the system invokes with the shareable content, or an error if a
// failure occurs.
//
// # Discussion
// 
// Use this method to retrieve the onscreen content that your app can capture.
// If the call is successful, the system returns the shareable content to the
// completion handler; otherwise, it returns an error that describes the
// failure.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getWithCompletionHandler(_:)
func (_SCShareableContentClass SCShareableContentClass) GetShareableContentWithCompletionHandler(completionHandler SCShareableContentErrorHandler) {
_block0, _cleanup0 := NewSCShareableContentErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](objc.ID(_SCShareableContentClass.class), objc.Sel("getShareableContentWithCompletionHandler:"), _block0)
}
// Retrieves the displays, apps, and windows that match your criteria.
//
// excludeDesktopWindows: A Boolean value that indicates whether to exclude desktop windows like
// Finder, Dock, and Desktop from the set of shareable content.
//
// onScreenWindowsOnly: A Boolean value that indicates whether to include only onscreen windows in
// the set of shareable content.
//
// completionHandler: A callback the system invokes with the shareable content, or an error if a
// failure occurs.
//
// # Discussion
// 
// Use this method to retrieve the onscreen content matching your filtering
// criteria. If the call is successful, the system passes an
// [SCShareableContent] instance to the completion handler; otherwise, it
// returns an error that describes the failure.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getExcludingDesktopWindows(_:onScreenWindowsOnly:completionHandler:)
func (_SCShareableContentClass SCShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler(excludeDesktopWindows bool, onScreenWindowsOnly bool, completionHandler SCShareableContentErrorHandler) {
_block2, _cleanup2 := NewSCShareableContentErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_SCShareableContentClass.class), objc.Sel("getShareableContentExcludingDesktopWindows:onScreenWindowsOnly:completionHandler:"), excludeDesktopWindows, onScreenWindowsOnly, _block2)
}
// Retrieves the displays, apps, and windows that are in front of the
// specified window.
//
// excludeDesktopWindows: A Boolean value that indicates whether to exclude desktop windows like
// Finder, Dock, and Desktop from the set of shareable content.
//
// window: The window above which to retrieve shareable content.
//
// completionHandler: A callback the system invokes with the shareable content, or an error if a
// failure occurs.
//
// # Discussion
// 
// Use this method to retrieve the onscreen content matching your filtering
// criteria. If the call is successful, the system passes an
// [SCShareableContent] instance to the completion handler; otherwise, it
// returns an error that describes the failure.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getExcludingDesktopWindows(_:onScreenWindowsOnlyAbove:completionHandler:)
func (_SCShareableContentClass SCShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindowCompletionHandler(excludeDesktopWindows bool, window ISCWindow, completionHandler SCShareableContentErrorHandler) {
_block2, _cleanup2 := NewSCShareableContentErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_SCShareableContentClass.class), objc.Sel("getShareableContentExcludingDesktopWindows:onScreenWindowsOnlyAboveWindow:completionHandler:"), excludeDesktopWindows, window, _block2)
}
// Retrieves the displays, apps, and windows that are behind the specified
// window.
//
// excludeDesktopWindows: A Boolean value that indicates whether to exclude desktop windows from the
// set of shareable content.
//
// window: The window above which to retrieve shareable content.
//
// completionHandler: A callback the system invokes with the shareable content, or an error if a
// failure occurs.
//
// # Discussion
// 
// Use this method to retrieve the onscreen content matching your filtering
// criteria. If the call is successful, the system passes an
// [SCShareableContent] instance to the completion handler; otherwise, it
// returns an error that describes the failure.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getExcludingDesktopWindows(_:onScreenWindowsOnlyBelow:completionHandler:)
func (_SCShareableContentClass SCShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindowCompletionHandler(excludeDesktopWindows bool, window ISCWindow, completionHandler SCShareableContentErrorHandler) {
_block2, _cleanup2 := NewSCShareableContentErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_SCShareableContentClass.class), objc.Sel("getShareableContentExcludingDesktopWindows:onScreenWindowsOnlyBelowWindow:completionHandler:"), excludeDesktopWindows, window, _block2)
}
// Retrieves any available sharable content information that matches the
// provided filter.
//
// filter: The filter to match current sharable content against.
//
// # Return Value
// 
// The sharable content matching the filter, or `nil` if none is found.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/info(for:)
func (_SCShareableContentClass SCShareableContentClass) InfoForFilter(filter ISCContentFilter) SCShareableContentInfo {
	rv := objc.Send[objc.ID](objc.ID(_SCShareableContentClass.class), objc.Sel("infoForFilter:"), filter)
	return SCShareableContentInfoFromID(rv)
}
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/getCurrentProcessShareableContent(completionHandler:)
func (_SCShareableContentClass SCShareableContentClass) GetCurrentProcessShareableContentWithCompletionHandler(completionHandler SCShareableContentErrorHandler) {
_block0, _cleanup0 := NewSCShareableContentErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](objc.ID(_SCShareableContentClass.class), objc.Sel("getCurrentProcessShareableContentWithCompletionHandler:"), _block0)
}

// The windows available for capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/windows
func (s SCShareableContent) Windows() []SCWindow {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("windows"))
	return objc.ConvertSlice(rv, func(id objc.ID) SCWindow {
		return SCWindowFromID(id)
	})
}
// The displays available for capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/displays
func (s SCShareableContent) Displays() []SCDisplay {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("displays"))
	return objc.ConvertSlice(rv, func(id objc.ID) SCDisplay {
		return SCDisplayFromID(id)
	})
}
// The apps available for capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContent/applications
func (s SCShareableContent) Applications() []SCRunningApplication {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("applications"))
	return objc.ConvertSlice(rv, func(id objc.ID) SCRunningApplication {
		return SCRunningApplicationFromID(id)
	})
}

// GetShareableContent is a synchronous wrapper around [SCShareableContent.GetShareableContentWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCShareableContentClass) GetShareableContent(ctx context.Context) (*SCShareableContent, error) {
	type result struct {
		val *SCShareableContent
		err error
	}
	done := make(chan result, 1)
	sc.GetShareableContentWithCompletionHandler(func(val *SCShareableContent, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnly is a synchronous wrapper around [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnly(ctx context.Context, excludeDesktopWindows bool, onScreenWindowsOnly bool) (*SCShareableContent, error) {
	type result struct {
		val *SCShareableContent
		err error
	}
	done := make(chan result, 1)
	sc.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyCompletionHandler(excludeDesktopWindows, onScreenWindowsOnly, func(val *SCShareableContent, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindow is a synchronous wrapper around [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindow(ctx context.Context, excludeDesktopWindows bool, window ISCWindow) (*SCShareableContent, error) {
	type result struct {
		val *SCShareableContent
		err error
	}
	done := make(chan result, 1)
	sc.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyAboveWindowCompletionHandler(excludeDesktopWindows, window, func(val *SCShareableContent, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindow is a synchronous wrapper around [SCShareableContent.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCShareableContentClass) GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindow(ctx context.Context, excludeDesktopWindows bool, window ISCWindow) (*SCShareableContent, error) {
	type result struct {
		val *SCShareableContent
		err error
	}
	done := make(chan result, 1)
	sc.GetShareableContentExcludingDesktopWindowsOnScreenWindowsOnlyBelowWindowCompletionHandler(excludeDesktopWindows, window, func(val *SCShareableContent, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// GetCurrentProcessShareableContent is a synchronous wrapper around [SCShareableContent.GetCurrentProcessShareableContentWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SCShareableContentClass) GetCurrentProcessShareableContent(ctx context.Context) (*SCShareableContent, error) {
	type result struct {
		val *SCShareableContent
		err error
	}
	done := make(chan result, 1)
	sc.GetCurrentProcessShareableContentWithCompletionHandler(func(val *SCShareableContent, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

