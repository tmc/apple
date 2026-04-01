// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
)

// The class instance for the [WKWebView] class.
var (
	_WKWebViewClass     WKWebViewClass
	_WKWebViewClassOnce sync.Once
)

func getWKWebViewClass() WKWebViewClass {
	_WKWebViewClassOnce.Do(func() {
		_WKWebViewClass = WKWebViewClass{class: objc.GetClass("WKWebView")}
	})
	return _WKWebViewClass
}

// GetWKWebViewClass returns the class object for WKWebView.
func GetWKWebViewClass() WKWebViewClass {
	return getWKWebViewClass()
}

type WKWebViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebViewClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebViewClass) Alloc() WKWebView {
	rv := objc.Send[WKWebView](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that displays interactive web content, such as for an in-app
// browser.
//
// # Overview
//
// A [WKWebView] object is a platform-native view that you use to incorporate
// web content seamlessly into your app’s UI. A web view supports a full
// web-browsing experience, and presents HTML, CSS, and JavaScript content
// alongside your app’s native views. Use it when web technologies satisfy
// your app’s layout and styling requirements more readily than native
// views. For example, you might use it when your app’s content changes
// frequently.
//
// A web view offers control over the navigation and user experience through
// delegate objects. Use the navigation delegate to react when the user clicks
// links in your web content, or interacts with the content in a way that
// affects navigation. For example, you might prevent the user from navigating
// to new content unless specific conditions are met. Use the UI delegate to
// present native UI elements, such as alerts or contextual menus, in response
// to interactions with your web content.
//
// Embed a [WKWebView] object programmatically into your view hierarchy, or
// add it using Interface Builder. Interface Builder supports many
// customizations, such as configuring data detectors, media playback, and
// interaction behaviors. For more extensive customizations, create your web
// view programmatically using a [WKWebViewConfiguration] object. For example,
// use a web view configuration object to specify handlers for custom URL
// schemes, manage cookies, and customize preferences for your web content.
//
// Before your web view appears onscreen, load content from a web server using
// a [URLRequest] structure or load content directly from a local file or HTML
// string. The web view automatically loads embedded resources such as images
// or videos as part of the initial load request. It then renders your content
// and displays the results inside the view’s bounds rectangle. The
// following code example shows a view controller that replaces its default
// view with a custom [WKWebView] object.
//
// A web view automatically converts telephone numbers that appear in web
// content to Phone links. When the user taps a Phone link, the Phone app
// launches and dials the number. Use the [WKWebViewConfiguration] object to
// change the default data detector behavior.
//
// You can also use [WKWebView.SetMagnificationCenteredAtPoint] to programmatically set
// the scale of web content the first time it appears in a web view.
// Thereafter, the user can change the scale using gestures.
//
// # Manage the navigation through your web content
//
// [WKWebView] provides a complete browsing experience, including the ability
// to navigate between different webpages using links, forward and back
// buttons, and more. When the user clicks a link in your content, the web
// view acts like a browser and displays the content at that link. To disallow
// navigation, or to customize your web view’s navigation behavior, provide
// your web view with a navigation delegate — an object that conforms to the
// [WKNavigationDelegate] protocol. Use your navigation delegate to modify the
// web view’s navigation behavior, or to track the loading progress of new
// content.
//
// You can also use the methods of [WKWebView] to navigate programmatically
// through your content, or to trigger navigation from other parts of your
// app’s interface. For example, if your UI includes forward and back
// buttons, connect those buttons to the [WKWebView.GoBack] and [WKWebView.GoForward] methods of
// your web view to trigger the corresponding web navigation. Use the
// [WKWebView.CanGoBack] and [WKWebView.CanGoForward] properties to determine when to enable or
// disable your buttons.
//
// # Provide sharing options
//
// People may want to share the contents of your web view with an app or other
// people. Use a [UIActivityViewController] to present a share sheet offering
// all the ways people can share the web content.
//
// If your app has the [com.apple.developer.web-browser] entitlement, the iOS
// share sheet can offer Add to Home Screen for an `http` or `https` webpage,
// creating a convenient link to a web app or bookmark. To allow someone to
// add the current webpage to the Home Screen, include the [WKWebView]
// instance in the `activityItems` array when you call
// [init(activityItems:applicationActivities:)] to create the
// [UIActivityViewController]. For more information about building a browser
// app, see [Preparing your app to be the default web browser].
//
// # Creating a web view
//
//   - [WKWebView.InitWithFrameConfiguration]: Creates a web view and initializes it with the specified frame and configuration data.
//   - [WKWebView.Configuration]: The object that contains the configuration details for the web view.
//
// # Displaying native user interface elements
//
//   - [WKWebView.UIDelegate]: The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//   - [WKWebView.SetUIDelegate]
//
// # Managing navigation between webpages
//
//   - [WKWebView.NavigationDelegate]: The object you use to manage navigation behavior for the web view.
//   - [WKWebView.SetNavigationDelegate]
//
// # Loading web content
//
//   - [WKWebView.LoadRequest]: Loads the web content that the specified URL request object references and navigates to that content.
//   - [WKWebView.LoadDataMIMETypeCharacterEncodingNameBaseURL]: Loads the content of the specified data object and navigates to it.
//   - [WKWebView.LoadHTMLStringBaseURL]: Loads the contents of the specified HTML string and navigates to it.
//   - [WKWebView.LoadFileRequestAllowingReadAccessToURL]: Loads the web content from the file the URL request object specifies and navigates to that content.
//   - [WKWebView.LoadFileURLAllowingReadAccessToURL]: Loads the web content from the specified file and navigates to it.
//   - [WKWebView.LoadSimulatedRequestResponseResponseData]: Loads the web content from the data you provide as if the data were the response to the request.
//   - [WKWebView.LoadSimulatedRequestResponseHTMLString]: Loads the web content from the HTML you provide as if the HTML were the response to the request.
//   - [WKWebView.Loading]: A Boolean value that indicates whether the view is currently loading content.
//   - [WKWebView.EstimatedProgress]: An estimate of what fraction of the current navigation has been loaded.
//
// # Managing the loading process
//
//   - [WKWebView.Reload]: Reloads the current webpage.
//   - [WKWebView.ReloadWithSender]: Reloads the current webpage.
//   - [WKWebView.ReloadFromOrigin]: Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//   - [WKWebView.ReloadFromOriginWithSender]: Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//   - [WKWebView.StopLoading]: Stops loading all resources on the current page.
//   - [WKWebView.StopLoadingWithSender]: Stops loading all resources on the current page.
//
// # Managing downloads
//
//   - [WKWebView.StartDownloadUsingRequestCompletionHandler]: Starts to download the resource at the URL in the request.
//   - [WKWebView.ResumeDownloadFromResumeDataCompletionHandler]: Resumes a failed or canceled download.
//
// # Making web content inspectable
//
//   - [WKWebView.Inspectable]: A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//   - [WKWebView.SetInspectable]
//
// # Inspecting the view information
//
//   - [WKWebView.Title]: The page title.
//   - [WKWebView.URL]: The URL for the current webpage.
//   - [WKWebView.MediaType]: The media type for the contents of the web view.
//   - [WKWebView.SetMediaType]
//   - [WKWebView.CustomUserAgent]: The custom user agent string.
//   - [WKWebView.SetCustomUserAgent]
//   - [WKWebView.ServerTrust]: The trust management object you use to evaluate trust for the current webpage.
//   - [WKWebView.HasOnlySecureContent]: A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//   - [WKWebView.ThemeColor]: The theme color that the system gets from the first valid meta tag in the webpage.
//   - [WKWebView.UnderPageBackgroundColor]: The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
//   - [WKWebView.SetUnderPageBackgroundColor]
//
// # Scaling content
//
//   - [WKWebView.PageZoom]: The scale factor by which the web view scales content relative to its bounds.
//   - [WKWebView.SetPageZoom]
//   - [WKWebView.AllowsMagnification]: A Boolean value that indicates whether magnify gestures change the web view’s magnification.
//   - [WKWebView.SetAllowsMagnification]
//   - [WKWebView.Magnification]: The factor by which the page content is currently scaled.
//   - [WKWebView.SetMagnification]
//   - [WKWebView.SetMagnificationCenteredAtPoint]: Scales the page content and centers the result on the specified point.
//
// # Interacting with media
//
//   - [WKWebView.PauseAllMediaPlaybackWithCompletionHandler]: Pauses playback of all media in the web view.
//   - [WKWebView.RequestMediaPlaybackStateWithCompletionHandler]: Requests the playback status of media in the web view.
//   - [WKWebView.SetAllMediaPlaybackSuspendedCompletionHandler]: Changes whether the webpage is suspending playback of all media in the page.
//   - [WKWebView.CloseAllMediaPresentationsWithCompletionHandler]: Closes all media the web view is presenting, including picture-in-picture video and fullscreen video.
//
// # Managing the microphone and camera
//
//   - [WKWebView.CameraCaptureState]: An enumeration case that indicates whether the webpage is using the camera to capture images or video.
//   - [WKWebView.MicrophoneCaptureState]: An enumeration case that indicates whether the webpage is using the microphone to capture audio.
//   - [WKWebView.SetCameraCaptureStateCompletionHandler]: Changes whether the webpage is using the camera to capture images or video.
//   - [WKWebView.SetMicrophoneCaptureStateCompletionHandler]: Changes whether the webpage is using the microphone to capture audio.
//
// # Navigating between webpages
//
//   - [WKWebView.AllowsBackForwardNavigationGestures]: A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//   - [WKWebView.SetAllowsBackForwardNavigationGestures]
//   - [WKWebView.BackForwardList]: The web view’s back-forward list.
//   - [WKWebView.GoBackWithSender]: Navigates to the back item in the back-forward list.
//   - [WKWebView.GoBack]: Navigates to the back item in the back-forward list.
//   - [WKWebView.GoForwardWithSender]: Navigates to the forward item in the back-forward list.
//   - [WKWebView.GoForward]: Navigates to the forward item in the back-forward list.
//   - [WKWebView.GoToBackForwardListItem]: Navigates to an item from the back-forward list and sets it as the current item.
//   - [WKWebView.CanGoBack]: A Boolean value that indicates whether there is a valid back item in the back-forward list.
//   - [WKWebView.CanGoForward]: A Boolean value that indicates whether there is a valid forward item in the back-forward list.
//   - [WKWebView.AllowsLinkPreview]: A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//   - [WKWebView.SetAllowsLinkPreview]
//   - [WKWebView.InteractionState]: An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
//   - [WKWebView.SetInteractionState]
//
// # Executing JavaScript
//
//   - [WKWebView.EvaluateJavaScriptCompletionHandler]: Evaluates the specified JavaScript string.
//
// # Capturing the web view’s content
//
//   - [WKWebView.TakeSnapshotWithConfigurationCompletionHandler]: Generates a platform-native image from the web view’s contents asynchronously.
//   - [WKWebView.PrintOperationWithPrintInfo]: Returns the print operation object to use when printing the contents of the web view.
//
// # Handling full-screen transitions
//
//   - [WKWebView.FullscreenState]
//
// # Configuring viewport insets
//
//   - [WKWebView.SetMinimumViewportInsetMaximumViewportInset]
//   - [WKWebView.MinimumViewportInset]
//   - [WKWebView.MaximumViewportInset]
//
// # Instance Properties
//
//   - [WKWebView.IsBlockedByScreenTime]
//   - [WKWebView.WritingToolsActive]
//   - [WKWebView.ObscuredContentInsets]
//   - [WKWebView.SetObscuredContentInsets]
//
// # Instance Methods
//
//   - [WKWebView.FetchDataOfTypesCompletionHandler]
//   - [WKWebView.RestoreDataCompletionHandler]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView
//
// [Preparing your app to be the default web browser]: https://developer.apple.com/documentation/Xcode/preparing-your-app-to-be-the-default-browser
// [UIActivityViewController]: https://developer.apple.com/documentation/UIKit/UIActivityViewController
// [URLRequest]: https://developer.apple.com/documentation/Foundation/URLRequest
// [com.apple.developer.web-browser]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.web-browser
// [init(activityItems:applicationActivities:)]: https://developer.apple.com/documentation/UIKit/UIActivityViewController/init(activityItems:applicationActivities:)
type WKWebView struct {
	appkit.NSView
}

// WKWebViewFromID constructs a [WKWebView] from an objc.ID.
//
// An object that displays interactive web content, such as for an in-app
// browser.
func WKWebViewFromID(id objc.ID) WKWebView {
	return WKWebView{NSView: appkit.NSViewFromID(id)}
}

// NOTE: WKWebView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebView] class.
//
// # Creating a web view
//
//   - [IWKWebView.InitWithFrameConfiguration]: Creates a web view and initializes it with the specified frame and configuration data.
//   - [IWKWebView.Configuration]: The object that contains the configuration details for the web view.
//
// # Displaying native user interface elements
//
//   - [IWKWebView.UIDelegate]: The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
//   - [IWKWebView.SetUIDelegate]
//
// # Managing navigation between webpages
//
//   - [IWKWebView.NavigationDelegate]: The object you use to manage navigation behavior for the web view.
//   - [IWKWebView.SetNavigationDelegate]
//
// # Loading web content
//
//   - [IWKWebView.LoadRequest]: Loads the web content that the specified URL request object references and navigates to that content.
//   - [IWKWebView.LoadDataMIMETypeCharacterEncodingNameBaseURL]: Loads the content of the specified data object and navigates to it.
//   - [IWKWebView.LoadHTMLStringBaseURL]: Loads the contents of the specified HTML string and navigates to it.
//   - [IWKWebView.LoadFileRequestAllowingReadAccessToURL]: Loads the web content from the file the URL request object specifies and navigates to that content.
//   - [IWKWebView.LoadFileURLAllowingReadAccessToURL]: Loads the web content from the specified file and navigates to it.
//   - [IWKWebView.LoadSimulatedRequestResponseResponseData]: Loads the web content from the data you provide as if the data were the response to the request.
//   - [IWKWebView.LoadSimulatedRequestResponseHTMLString]: Loads the web content from the HTML you provide as if the HTML were the response to the request.
//   - [IWKWebView.Loading]: A Boolean value that indicates whether the view is currently loading content.
//   - [IWKWebView.EstimatedProgress]: An estimate of what fraction of the current navigation has been loaded.
//
// # Managing the loading process
//
//   - [IWKWebView.Reload]: Reloads the current webpage.
//   - [IWKWebView.ReloadWithSender]: Reloads the current webpage.
//   - [IWKWebView.ReloadFromOrigin]: Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//   - [IWKWebView.ReloadFromOriginWithSender]: Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
//   - [IWKWebView.StopLoading]: Stops loading all resources on the current page.
//   - [IWKWebView.StopLoadingWithSender]: Stops loading all resources on the current page.
//
// # Managing downloads
//
//   - [IWKWebView.StartDownloadUsingRequestCompletionHandler]: Starts to download the resource at the URL in the request.
//   - [IWKWebView.ResumeDownloadFromResumeDataCompletionHandler]: Resumes a failed or canceled download.
//
// # Making web content inspectable
//
//   - [IWKWebView.Inspectable]: A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
//   - [IWKWebView.SetInspectable]
//
// # Inspecting the view information
//
//   - [IWKWebView.Title]: The page title.
//   - [IWKWebView.URL]: The URL for the current webpage.
//   - [IWKWebView.MediaType]: The media type for the contents of the web view.
//   - [IWKWebView.SetMediaType]
//   - [IWKWebView.CustomUserAgent]: The custom user agent string.
//   - [IWKWebView.SetCustomUserAgent]
//   - [IWKWebView.ServerTrust]: The trust management object you use to evaluate trust for the current webpage.
//   - [IWKWebView.HasOnlySecureContent]: A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
//   - [IWKWebView.ThemeColor]: The theme color that the system gets from the first valid meta tag in the webpage.
//   - [IWKWebView.UnderPageBackgroundColor]: The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
//   - [IWKWebView.SetUnderPageBackgroundColor]
//
// # Scaling content
//
//   - [IWKWebView.PageZoom]: The scale factor by which the web view scales content relative to its bounds.
//   - [IWKWebView.SetPageZoom]
//   - [IWKWebView.AllowsMagnification]: A Boolean value that indicates whether magnify gestures change the web view’s magnification.
//   - [IWKWebView.SetAllowsMagnification]
//   - [IWKWebView.Magnification]: The factor by which the page content is currently scaled.
//   - [IWKWebView.SetMagnification]
//   - [IWKWebView.SetMagnificationCenteredAtPoint]: Scales the page content and centers the result on the specified point.
//
// # Interacting with media
//
//   - [IWKWebView.PauseAllMediaPlaybackWithCompletionHandler]: Pauses playback of all media in the web view.
//   - [IWKWebView.RequestMediaPlaybackStateWithCompletionHandler]: Requests the playback status of media in the web view.
//   - [IWKWebView.SetAllMediaPlaybackSuspendedCompletionHandler]: Changes whether the webpage is suspending playback of all media in the page.
//   - [IWKWebView.CloseAllMediaPresentationsWithCompletionHandler]: Closes all media the web view is presenting, including picture-in-picture video and fullscreen video.
//
// # Managing the microphone and camera
//
//   - [IWKWebView.CameraCaptureState]: An enumeration case that indicates whether the webpage is using the camera to capture images or video.
//   - [IWKWebView.MicrophoneCaptureState]: An enumeration case that indicates whether the webpage is using the microphone to capture audio.
//   - [IWKWebView.SetCameraCaptureStateCompletionHandler]: Changes whether the webpage is using the camera to capture images or video.
//   - [IWKWebView.SetMicrophoneCaptureStateCompletionHandler]: Changes whether the webpage is using the microphone to capture audio.
//
// # Navigating between webpages
//
//   - [IWKWebView.AllowsBackForwardNavigationGestures]: A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
//   - [IWKWebView.SetAllowsBackForwardNavigationGestures]
//   - [IWKWebView.BackForwardList]: The web view’s back-forward list.
//   - [IWKWebView.GoBackWithSender]: Navigates to the back item in the back-forward list.
//   - [IWKWebView.GoBack]: Navigates to the back item in the back-forward list.
//   - [IWKWebView.GoForwardWithSender]: Navigates to the forward item in the back-forward list.
//   - [IWKWebView.GoForward]: Navigates to the forward item in the back-forward list.
//   - [IWKWebView.GoToBackForwardListItem]: Navigates to an item from the back-forward list and sets it as the current item.
//   - [IWKWebView.CanGoBack]: A Boolean value that indicates whether there is a valid back item in the back-forward list.
//   - [IWKWebView.CanGoForward]: A Boolean value that indicates whether there is a valid forward item in the back-forward list.
//   - [IWKWebView.AllowsLinkPreview]: A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
//   - [IWKWebView.SetAllowsLinkPreview]
//   - [IWKWebView.InteractionState]: An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
//   - [IWKWebView.SetInteractionState]
//
// # Executing JavaScript
//
//   - [IWKWebView.EvaluateJavaScriptCompletionHandler]: Evaluates the specified JavaScript string.
//
// # Capturing the web view’s content
//
//   - [IWKWebView.TakeSnapshotWithConfigurationCompletionHandler]: Generates a platform-native image from the web view’s contents asynchronously.
//   - [IWKWebView.PrintOperationWithPrintInfo]: Returns the print operation object to use when printing the contents of the web view.
//
// # Handling full-screen transitions
//
//   - [IWKWebView.FullscreenState]
//
// # Configuring viewport insets
//
//   - [IWKWebView.SetMinimumViewportInsetMaximumViewportInset]
//   - [IWKWebView.MinimumViewportInset]
//   - [IWKWebView.MaximumViewportInset]
//
// # Instance Properties
//
//   - [IWKWebView.IsBlockedByScreenTime]
//   - [IWKWebView.WritingToolsActive]
//   - [IWKWebView.ObscuredContentInsets]
//   - [IWKWebView.SetObscuredContentInsets]
//
// # Instance Methods
//
//   - [IWKWebView.FetchDataOfTypesCompletionHandler]
//   - [IWKWebView.RestoreDataCompletionHandler]
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView
type IWKWebView interface {
	appkit.INSView

	// Topic: Creating a web view

	// Creates a web view and initializes it with the specified frame and configuration data.
	InitWithFrameConfiguration(frame corefoundation.CGRect, configuration IWKWebViewConfiguration) WKWebView
	// The object that contains the configuration details for the web view.
	Configuration() IWKWebViewConfiguration

	// Topic: Displaying native user interface elements

	// The object you use to integrate custom user interface elements, such as contextual menus or panels, into web view interactions.
	UIDelegate() WKUIDelegate
	SetUIDelegate(value WKUIDelegate)

	// Topic: Managing navigation between webpages

	// The object you use to manage navigation behavior for the web view.
	NavigationDelegate() WKNavigationDelegate
	SetNavigationDelegate(value WKNavigationDelegate)

	// Topic: Loading web content

	// Loads the web content that the specified URL request object references and navigates to that content.
	LoadRequest(request foundation.NSURLRequest) IWKNavigation
	// Loads the content of the specified data object and navigates to it.
	LoadDataMIMETypeCharacterEncodingNameBaseURL(data foundation.INSData, MIMEType string, characterEncodingName string, baseURL foundation.INSURL) IWKNavigation
	// Loads the contents of the specified HTML string and navigates to it.
	LoadHTMLStringBaseURL(string_ string, baseURL foundation.INSURL) IWKNavigation
	// Loads the web content from the file the URL request object specifies and navigates to that content.
	LoadFileRequestAllowingReadAccessToURL(request foundation.NSURLRequest, readAccessURL foundation.INSURL) IWKNavigation
	// Loads the web content from the specified file and navigates to it.
	LoadFileURLAllowingReadAccessToURL(URL foundation.INSURL, readAccessURL foundation.INSURL) IWKNavigation
	// Loads the web content from the data you provide as if the data were the response to the request.
	LoadSimulatedRequestResponseResponseData(request foundation.NSURLRequest, response foundation.NSURLResponse, data foundation.INSData) IWKNavigation
	// Loads the web content from the HTML you provide as if the HTML were the response to the request.
	LoadSimulatedRequestResponseHTMLString(request foundation.NSURLRequest, string_ string) IWKNavigation
	// A Boolean value that indicates whether the view is currently loading content.
	Loading() bool
	// An estimate of what fraction of the current navigation has been loaded.
	EstimatedProgress() float64

	// Topic: Managing the loading process

	// Reloads the current webpage.
	Reload() IWKNavigation
	// Reloads the current webpage.
	ReloadWithSender(sender objectivec.IObject)
	// Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
	ReloadFromOrigin() IWKNavigation
	// Reloads the current webpage, and performs end-to-end revalidation of the content using cache-validating conditionals, if possible.
	ReloadFromOriginWithSender(sender objectivec.IObject)
	// Stops loading all resources on the current page.
	StopLoading()
	// Stops loading all resources on the current page.
	StopLoadingWithSender(sender objectivec.IObject)

	// Topic: Managing downloads

	// Starts to download the resource at the URL in the request.
	StartDownloadUsingRequestCompletionHandler(request foundation.NSURLRequest, completionHandler WKDownloadHandler)
	// Resumes a failed or canceled download.
	ResumeDownloadFromResumeDataCompletionHandler(resumeData foundation.INSData, completionHandler WKDownloadHandler)

	// Topic: Making web content inspectable

	// A Boolean value that indicates whether you can inspect the view with Safari Web Inspector.
	Inspectable() bool
	SetInspectable(value bool)

	// Topic: Inspecting the view information

	// The page title.
	Title() string
	// The URL for the current webpage.
	URL() foundation.INSURL
	// The media type for the contents of the web view.
	MediaType() string
	SetMediaType(value string)
	// The custom user agent string.
	CustomUserAgent() string
	SetCustomUserAgent(value string)
	// The trust management object you use to evaluate trust for the current webpage.
	ServerTrust() security.SecTrustRef
	// A Boolean value that indicates whether the web view loaded all resources on the page through securely encrypted connections.
	HasOnlySecureContent() bool
	// The theme color that the system gets from the first valid meta tag in the webpage.
	ThemeColor() appkit.NSColor
	// The color the web view displays behind the active page, visible when the user scrolls beyond the bounds of the page.
	UnderPageBackgroundColor() appkit.NSColor
	SetUnderPageBackgroundColor(value appkit.NSColor)

	// Topic: Scaling content

	// The scale factor by which the web view scales content relative to its bounds.
	PageZoom() float64
	SetPageZoom(value float64)
	// A Boolean value that indicates whether magnify gestures change the web view’s magnification.
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	// The factor by which the page content is currently scaled.
	Magnification() float64
	SetMagnification(value float64)
	// Scales the page content and centers the result on the specified point.
	SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint)

	// Topic: Interacting with media

	// Pauses playback of all media in the web view.
	PauseAllMediaPlaybackWithCompletionHandler(completionHandler VoidHandler)
	// Requests the playback status of media in the web view.
	RequestMediaPlaybackStateWithCompletionHandler(completionHandler WKMediaPlaybackStateHandler)
	// Changes whether the webpage is suspending playback of all media in the page.
	SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler VoidHandler)
	// Closes all media the web view is presenting, including picture-in-picture video and fullscreen video.
	CloseAllMediaPresentationsWithCompletionHandler(completionHandler VoidHandler)

	// Topic: Managing the microphone and camera

	// An enumeration case that indicates whether the webpage is using the camera to capture images or video.
	CameraCaptureState() WKMediaCaptureState
	// An enumeration case that indicates whether the webpage is using the microphone to capture audio.
	MicrophoneCaptureState() WKMediaCaptureState
	// Changes whether the webpage is using the camera to capture images or video.
	SetCameraCaptureStateCompletionHandler(state WKMediaCaptureState, completionHandler VoidHandler)
	// Changes whether the webpage is using the microphone to capture audio.
	SetMicrophoneCaptureStateCompletionHandler(state WKMediaCaptureState, completionHandler VoidHandler)

	// Topic: Navigating between webpages

	// A Boolean value that indicates whether horizontal swipe gestures trigger backward and forward page navigation.
	AllowsBackForwardNavigationGestures() bool
	SetAllowsBackForwardNavigationGestures(value bool)
	// The web view’s back-forward list.
	BackForwardList() IWKBackForwardList
	// Navigates to the back item in the back-forward list.
	GoBackWithSender(sender objectivec.IObject)
	// Navigates to the back item in the back-forward list.
	GoBack() IWKNavigation
	// Navigates to the forward item in the back-forward list.
	GoForwardWithSender(sender objectivec.IObject)
	// Navigates to the forward item in the back-forward list.
	GoForward() IWKNavigation
	// Navigates to an item from the back-forward list and sets it as the current item.
	GoToBackForwardListItem(item IWKBackForwardListItem) IWKNavigation
	// A Boolean value that indicates whether there is a valid back item in the back-forward list.
	CanGoBack() bool
	// A Boolean value that indicates whether there is a valid forward item in the back-forward list.
	CanGoForward() bool
	// A Boolean value that determines whether pressing a link displays a preview of the destination for the link.
	AllowsLinkPreview() bool
	SetAllowsLinkPreview(value bool)
	// An object you use to capture the current state of interaction in a web view so that you can restore that state later to another web view.
	InteractionState() objectivec.IObject
	SetInteractionState(value objectivec.IObject)

	// Topic: Executing JavaScript

	// Evaluates the specified JavaScript string.
	EvaluateJavaScriptCompletionHandler(javaScriptString string, completionHandler ObjectErrorHandler)

	// Topic: Capturing the web view’s content

	// Generates a platform-native image from the web view’s contents asynchronously.
	TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration IWKSnapshotConfiguration, completionHandler ImageErrorHandler)
	// Returns the print operation object to use when printing the contents of the web view.
	PrintOperationWithPrintInfo(printInfo appkit.NSPrintInfo) appkit.NSPrintOperation

	// Topic: Handling full-screen transitions

	FullscreenState() WKFullscreenState

	// Topic: Configuring viewport insets

	SetMinimumViewportInsetMaximumViewportInset(minimumViewportInset foundation.NSEdgeInsets, maximumViewportInset foundation.NSEdgeInsets)
	MinimumViewportInset() foundation.NSEdgeInsets
	MaximumViewportInset() foundation.NSEdgeInsets

	// Topic: Instance Properties

	IsBlockedByScreenTime() bool
	WritingToolsActive() bool
	ObscuredContentInsets() foundation.NSEdgeInsets
	SetObscuredContentInsets(value foundation.NSEdgeInsets)

	// Topic: Instance Methods

	FetchDataOfTypesCompletionHandler(dataTypes WKWebViewDataType, completionHandler DataErrorHandler)
	RestoreDataCompletionHandler(data foundation.INSData, completionHandler ErrorHandler)

	// Generates PDF data from the web view’s contents asynchronously.
	CreatePDFWithConfigurationCompletionHandler(pdfConfiguration IWKPDFConfiguration, completionHandler DataErrorHandler)
	// Creates a web archive of the web view’s contents asynchronously.
	CreateWebArchiveDataWithCompletionHandler(completionHandler DataErrorHandler)
	// Searches for the specified string in the web view’s content.
	FindStringWithConfigurationCompletionHandler(string_ string, configuration IWKFindConfiguration, completionHandler WKFindResultHandler)
}

// Init initializes the instance.
func (w WKWebView) Init() WKWebView {
	rv := objc.Send[WKWebView](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebView) Autorelease() WKWebView {
	rv := objc.Send[WKWebView](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebView creates a new WKWebView instance.
func NewWKWebView() WKWebView {
	class := getWKWebViewClass()
	rv := objc.Send[WKWebView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an object initialized from data in the specified coder object.
//
// coder: The coder object that contains the object’s data.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/init(coder:)
func NewWebViewWithCoder(coder foundation.INSCoder) WKWebView {
	instance := getWKWebViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return WKWebViewFromID(rv)
}

// Creates a web view and initializes it with the specified frame and
// configuration data.
//
// frame: The frame rectangle for the new web view.
//
// configuration: The configuration of the new web view. This method saves a copy of your
// configuration object. Changes you make to your original object after
// calling this method have no effect on the web view’s configuration. For a
// list of configuration options and their default values, see
// [WKWebViewConfiguration].
//
// # Return Value
//
// An initialized web view, or `nil` if the view couldn’t be initialized.
//
// # Discussion
//
// This method is the designated initializer for the class. Use this method to
// create a web view that requires custom configuration. For example, use it
// when you need to specify custom cookies or content filters for the web
// content.
//
// To create a web view with default configuration values, call the inherited
// [init(frame:)] method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/init(frame:configuration:)
//
// [init(frame:)]: https://developer.apple.com/documentation/UIKit/UIView/init(frame:)
func NewWebViewWithFrameConfiguration(frame corefoundation.CGRect, configuration IWKWebViewConfiguration) WKWebView {
	instance := getWKWebViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:configuration:"), frame, configuration)
	return WKWebViewFromID(rv)
}

// Creates a web view and initializes it with the specified frame and
// configuration data.
//
// frame: The frame rectangle for the new web view.
//
// configuration: The configuration of the new web view. This method saves a copy of your
// configuration object. Changes you make to your original object after
// calling this method have no effect on the web view’s configuration. For a
// list of configuration options and their default values, see
// [WKWebViewConfiguration].
//
// # Return Value
//
// An initialized web view, or `nil` if the view couldn’t be initialized.
//
// # Discussion
//
// This method is the designated initializer for the class. Use this method to
// create a web view that requires custom configuration. For example, use it
// when you need to specify custom cookies or content filters for the web
// content.
//
// To create a web view with default configuration values, call the inherited
// [init(frame:)] method.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/init(frame:configuration:)
//
// [init(frame:)]: https://developer.apple.com/documentation/UIKit/UIView/init(frame:)
func (w WKWebView) InitWithFrameConfiguration(frame corefoundation.CGRect, configuration IWKWebViewConfiguration) WKWebView {
	rv := objc.Send[WKWebView](w.ID, objc.Sel("initWithFrame:configuration:"), frame, configuration)
	return rv
}

// Loads the web content that the specified URL request object references and
// navigates to that content.
//
// request: A URL request that specifies the resource to display.
//
// # Return Value
//
// A new navigation object that you use to track the loading progress of the
// request.
//
// # Discussion
//
// Use this method to load a page from a local or network-based URL. For
// example, you might use this method to navigate to a network-based webpage.
//
// Provide the source of this load request for app activity data by setting
// the [attribution] parameter on your request.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/load(_:)
//
// [attribution]: https://developer.apple.com/documentation/Foundation/URLRequest/attribution-swift.property
func (w WKWebView) LoadRequest(request foundation.NSURLRequest) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("loadRequest:"), request)
	return WKNavigationFromID(rv)
}

// Loads the content of the specified data object and navigates to it.
//
// data: The data to use as the contents of the webpage.
//
// MIMEType: The MIME type of the information in the `data` parameter. This parameter
// must not contain an empty string.
//
// characterEncodingName: The data’s character encoding name.
//
// baseURL: A URL that you use to resolve relative URLs within the document.
//
// # Return Value
//
// A new navigation object for tracking the request.
//
// # Discussion
//
// Use this method to navigate to a webpage that you loaded yourself and saved
// in a data object. For example, if you previously wrote HTML content to a
// data object, use this method to navigate to that content.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/load(_:mimeType:characterEncodingName:baseURL:)
func (w WKWebView) LoadDataMIMETypeCharacterEncodingNameBaseURL(data foundation.INSData, MIMEType string, characterEncodingName string, baseURL foundation.INSURL) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("loadData:MIMEType:characterEncodingName:baseURL:"), data, objc.String(MIMEType), objc.String(characterEncodingName), baseURL)
	return WKNavigationFromID(rv)
}

// Loads the contents of the specified HTML string and navigates to it.
//
// string: The string to use as the contents of the webpage.
//
// baseURL: The base URL to use when the system resolves relative URLs within the HTML
// string.
//
// # Return Value
//
// A new navigation object you use to track the loading progress of the
// request.
//
// # Discussion
//
// Use this method to navigate to a webpage that you loaded or created
// yourself. For example, you might use this method to load HTML content that
// your app generates programmatically.
//
// This method sets the source of this load request for app activity data to
// [NSURLRequest.Attribution.developer].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/loadHTMLString(_:baseURL:)
//
// [NSURLRequest.Attribution.developer]: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum/developer
func (w WKWebView) LoadHTMLStringBaseURL(string_ string, baseURL foundation.INSURL) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("loadHTMLString:baseURL:"), objc.String(string_), baseURL)
	return WKNavigationFromID(rv)
}

// Loads the web content from the file the URL request object specifies and
// navigates to that content.
//
// request: A URL request that specifies the file to display. The URL in this request
// must be a file-based URL.
//
// readAccessURL: The URL of a file or directory containing web content that you grant the
// system permission to read. This URL must be a file-based URL and must not
// be empty. To prevent WebKit from reading any other content, specify the
// same value as the URL parameter. To read additional files related to the
// content file, specify a directory.
//
// # Return Value
//
// A new navigation object you use to track the loading progress of the
// request.
//
// # Discussion
//
// Provide the source of this load request for app activity data by setting
// the [attribution] parameter on your request.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/loadFileRequest(_:allowingReadAccessTo:)
//
// [attribution]: https://developer.apple.com/documentation/Foundation/URLRequest/attribution-swift.property
func (w WKWebView) LoadFileRequestAllowingReadAccessToURL(request foundation.NSURLRequest, readAccessURL foundation.INSURL) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("loadFileRequest:allowingReadAccessToURL:"), request, readAccessURL)
	return WKNavigationFromID(rv)
}

// Loads the web content from the specified file and navigates to it.
//
// URL: The URL of a file that contains web content. This URL must be a file-based
// URL.
//
// readAccessURL: The URL of a file or directory containing web content that you grant the
// system permission to read. This URL must be a file-based URL and must not
// be empty. To prevent WebKit from reading any other content, specify the
// same value as the URL parameter. To read additional files related to the
// content file, specify a directory.
//
// # Return Value
//
// A new navigation object you use to track the loading progress of the
// request.
//
// # Discussion
//
// This method sets the source of this load request for app activity data to
// [NSURLRequest.Attribution.developer]. To specify the source of this load,
// use [LoadFileRequestAllowingReadAccessToURL] instead.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/loadFileURL(_:allowingReadAccessTo:)
//
// [NSURLRequest.Attribution.developer]: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum/developer
func (w WKWebView) LoadFileURLAllowingReadAccessToURL(URL foundation.INSURL, readAccessURL foundation.INSURL) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("loadFileURL:allowingReadAccessToURL:"), URL, readAccessURL)
	return WKNavigationFromID(rv)
}

// Loads the web content from the data you provide as if the data were the
// response to the request.
//
// request: A URL request that specifies the base URL and other loading details the
// system uses to interpret the data you provide.
//
// response: A response the system uses to interpret the data you provide.
//
// data: The data to use as the contents of the webpage.
//
// # Return Value
//
// A new navigation object you use to track the loading progress of the
// request.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:response:responseData:)
func (w WKWebView) LoadSimulatedRequestResponseResponseData(request foundation.NSURLRequest, response foundation.NSURLResponse, data foundation.INSData) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("loadSimulatedRequest:response:responseData:"), request, response, data)
	return WKNavigationFromID(rv)
}

// Loads the web content from the HTML you provide as if the HTML were the
// response to the request.
//
// request: A URL request that specifies the base URL and other loading details the
// system uses to interpret the HTML you provide.
//
// string: The HTML code you provide in a string to use as the contents of the
// webpage.
//
// # Return Value
//
// A new navigation object you use to track the loading progress of the
// request.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/loadSimulatedRequest(_:responseHTML:)
func (w WKWebView) LoadSimulatedRequestResponseHTMLString(request foundation.NSURLRequest, string_ string) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("loadSimulatedRequest:responseHTMLString:"), request, objc.String(string_))
	return WKNavigationFromID(rv)
}

// Reloads the current webpage.
//
// # Return Value
//
// A new navigation object to represent the reload operation.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/reload()
func (w WKWebView) Reload() IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("reload"))
	return WKNavigationFromID(rv)
}

// Reloads the current webpage.
//
// sender: The object that sent the message.
//
// # Discussion
//
// Make this method the action of any controls that trigger a reload of your
// web content. Connect your controls to this method programmatically or in
// Interface Builder.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/reload(_:)
func (w WKWebView) ReloadWithSender(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("reload:"), sender)
}

// Reloads the current webpage, and performs end-to-end revalidation of the
// content using cache-validating conditionals, if possible.
//
// # Return Value
//
// A new navigation that represents the reloaded webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/reloadFromOrigin()
func (w WKWebView) ReloadFromOrigin() IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("reloadFromOrigin"))
	return WKNavigationFromID(rv)
}

// Reloads the current webpage, and performs end-to-end revalidation of the
// content using cache-validating conditionals, if possible.
//
// sender: The object that sent the message.
//
// # Discussion
//
// Make this method the action of any controls that trigger a reload of your
// web content. Connect your controls to this method programmatically or in
// Interface Builder.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/reloadFromOrigin(_:)
func (w WKWebView) ReloadFromOriginWithSender(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("reloadFromOrigin:"), sender)
}

// Stops loading all resources on the current page.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/stopLoading()
func (w WKWebView) StopLoading() {
	objc.Send[objc.ID](w.ID, objc.Sel("stopLoading"))
}

// Stops loading all resources on the current page.
//
// sender: The object that sent the message.
//
// # Discussion
//
// Make this method the action of any controls that stop the loading of your
// web content. Connect your controls to this method programmatically or in
// Interface Builder.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/stopLoading(_:)
func (w WKWebView) StopLoadingWithSender(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("stopLoading:"), sender)
}

// Starts to download the resource at the URL in the request.
//
// request: An object that encapsulates a URL and other parameters that you need to
// download a resource from a webpage.
//
// completionHandler: A closure the system executes when it has started to download the resource.
//
// # Discussion
//
// To receive progress updates, set the delegate of the download object in the
// completion handler.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/startDownload(using:completionHandler:)
func (w WKWebView) StartDownloadUsingRequestCompletionHandler(request foundation.NSURLRequest, completionHandler WKDownloadHandler) {
	_block1, _ := NewWKDownloadBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("startDownloadUsingRequest:completionHandler:"), request, _block1)
}

// Resumes a failed or canceled download.
//
// resumeData: An object with data that you use to resume a failed or canceled download.
//
// completionHandler: A closure the system executes when it has resumed a download.
//
// # Discussion
//
// To receive progress updates, set the delegate of the download object in the
// completion handler.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/resumeDownload(fromResumeData:completionHandler:)
func (w WKWebView) ResumeDownloadFromResumeDataCompletionHandler(resumeData foundation.INSData, completionHandler WKDownloadHandler) {
	_block1, _ := NewWKDownloadBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("resumeDownloadFromResumeData:completionHandler:"), resumeData, _block1)
}

// Scales the page content and centers the result on the specified point.
//
// magnification: The factor by which to scale the content.
//
// point: The point (in the web view’s bounds) at which to center magnification.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/setMagnification(_:centeredAt:)
func (w WKWebView) SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint) {
	objc.Send[objc.ID](w.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}

// Pauses playback of all media in the web view.
//
// completionHandler: A closure the system executes after the web view pauses media playback.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/pauseAllMediaPlayback(completionHandler:)
func (w WKWebView) PauseAllMediaPlaybackWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("pauseAllMediaPlaybackWithCompletionHandler:"), _block0)
}

// Requests the playback status of media in the web view.
//
// completionHandler: A closure the system executes after the web view determines the current
// state of media playback.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/requestMediaPlaybackState(completionHandler:)
func (w WKWebView) RequestMediaPlaybackStateWithCompletionHandler(completionHandler WKMediaPlaybackStateHandler) {
	_block0, _ := NewWKMediaPlaybackStateBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("requestMediaPlaybackStateWithCompletionHandler:"), _block0)
}

// Changes whether the webpage is suspending playback of all media in the
// page.
//
// suspended: A Boolean value that indicates whether the webpage should suspend media
// playback.
//
// completionHandler: A closure the system executes after it completes changing the media
// playback suspension status.
//
// # Discussion
//
// Pass `true` to pause all media the web view is playing. Neither the user
// nor the webpage can resume playback until you call this method again with
// `false`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/setAllMediaPlaybackSuspended(_:completionHandler:)
func (w WKWebView) SetAllMediaPlaybackSuspendedCompletionHandler(suspended bool, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("setAllMediaPlaybackSuspended:completionHandler:"), suspended, _block1)
}

// Closes all media the web view is presenting, including picture-in-picture
// video and fullscreen video.
//
// completionHandler: A closure the system executes after it completes closing all media
// presentations.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/closeAllMediaPresentations(completionHandler:)
func (w WKWebView) CloseAllMediaPresentationsWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("closeAllMediaPresentationsWithCompletionHandler:"), _block0)
}

// Changes whether the webpage is using the camera to capture images or video.
//
// state: An enumeration case that indicates whether the webpage should use the
// camera to capture images or video.
//
// completionHandler: A closure the system executes after changing whether the webpage is using
// the camera to capture images or video.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/setCameraCaptureState(_:completionHandler:)
func (w WKWebView) SetCameraCaptureStateCompletionHandler(state WKMediaCaptureState, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("setCameraCaptureState:completionHandler:"), state, _block1)
}

// Changes whether the webpage is using the microphone to capture audio.
//
// state: An enumeration case that indicates whether the webpage should use the
// microphone to capture audio.
//
// completionHandler: A closure the system executes after changing whether the webpage is using
// the microphone to capture audio.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/setMicrophoneCaptureState(_:completionHandler:)
func (w WKWebView) SetMicrophoneCaptureStateCompletionHandler(state WKMediaCaptureState, completionHandler VoidHandler) {
	_block1, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("setMicrophoneCaptureState:completionHandler:"), state, _block1)
}

// Navigates to the back item in the back-forward list.
//
// sender: The object sending this message.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/goBack(_:)
func (w WKWebView) GoBackWithSender(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("goBack:"), sender)
}

// Navigates to the back item in the back-forward list.
//
// # Return Value
//
// A new navigation to the requested item, or `nil` if there is no back item
// in the back-forward list.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/goBack()
func (w WKWebView) GoBack() IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("goBack"))
	return WKNavigationFromID(rv)
}

// Navigates to the forward item in the back-forward list.
//
// sender: The object sending this message.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/goForward(_:)
func (w WKWebView) GoForwardWithSender(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("goForward:"), sender)
}

// Navigates to the forward item in the back-forward list.
//
// # Return Value
//
// A new navigation to the requested item, or `nil` if there is no forward
// item in the back-forward list.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/goForward()
func (w WKWebView) GoForward() IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("goForward"))
	return WKNavigationFromID(rv)
}

// Navigates to an item from the back-forward list and sets it as the current
// item.
//
// item: The item to navigate to. The item must be in the web view’s back-forward
// list.
//
// # Return Value
//
// A new navigation to the requested item, or `nil` if it is already the
// current item or is not part of the web view’s back-forward list.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/go(to:)
func (w WKWebView) GoToBackForwardListItem(item IWKBackForwardListItem) IWKNavigation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("goToBackForwardListItem:"), item)
	return WKNavigationFromID(rv)
}

// Evaluates the specified JavaScript string.
//
// javaScriptString: The JavaScript string to evaluate.
//
// completionHandler: A handler block to execute when script evaluation finishes. The method
// calls your block whether script evaluation completes successfully or fails.
// The block has no return value and takes the following parameters:
//
// object: The result of the script evaluation, or `nil` if an error occurred.
// error: `nil` on success, or an error object with information about the
// problem.
//
// # Discussion
//
// After evaluating the script, this method executes the completion handler
// block with either the result of the script evaluation or an error. The
// completion handler always runs on the app’s main thread.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/evaluateJavaScript(_:completionHandler:)
func (w WKWebView) EvaluateJavaScriptCompletionHandler(javaScriptString string, completionHandler ObjectErrorHandler) {
	_block1, _ := NewObjectErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("evaluateJavaScript:completionHandler:"), objc.String(javaScriptString), _block1)
}

// Generates a platform-native image from the web view’s contents
// asynchronously.
//
// snapshotConfiguration: The object that specifies the portion of the web view to capture, and other
// capture-related behaviors.
//
// completionHandler: The completion handler to call when the image is ready. This block has no
// return value and takes the following parameters:
//
// snapshotImage: A platform-native image that contains the specified portion
// of the web view. error: An error object if a problem occurred, or `nil` on
// success.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/takeSnapshot(with:completionHandler:)
func (w WKWebView) TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration IWKSnapshotConfiguration, completionHandler ImageErrorHandler) {
	_block1, _ := NewImageErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("takeSnapshotWithConfiguration:completionHandler:"), snapshotConfiguration, _block1)
}

// Returns the print operation object to use when printing the contents of the
// web view.
//
// printInfo: The printer information object to use when configuring the print operation.
//
// # Return Value
//
// The print operation object to use when printing the web view, or `nil` if
// printing is not supported.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/printOperation(with:)
func (w WKWebView) PrintOperationWithPrintInfo(printInfo appkit.NSPrintInfo) appkit.NSPrintOperation {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("printOperationWithPrintInfo:"), printInfo)
	return appkit.NSPrintOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/setMinimumViewportInset(_:maximumViewportInset:)
func (w WKWebView) SetMinimumViewportInsetMaximumViewportInset(minimumViewportInset foundation.NSEdgeInsets, maximumViewportInset foundation.NSEdgeInsets) {
	objc.Send[objc.ID](w.ID, objc.Sel("setMinimumViewportInset:maximumViewportInset:"), minimumViewportInset, maximumViewportInset)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/fetchData(of:completionHandler:)
func (w WKWebView) FetchDataOfTypesCompletionHandler(dataTypes WKWebViewDataType, completionHandler DataErrorHandler) {
	_block1, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("fetchDataOfTypes:completionHandler:"), dataTypes, _block1)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/restoreData(_:completionHandler:)
func (w WKWebView) RestoreDataCompletionHandler(data foundation.INSData, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("restoreData:completionHandler:"), data, _block1)
}

// Generates PDF data from the web view’s contents asynchronously.
//
// pdfConfiguration: The object that specifies the portion of the web view to capture as PDF
// data.
//
// completionHandler: The completion handler to call when the data is ready. This block has no
// return value and takes the following parameters:
//
// pdfDocumentData: A data object that contains the PDF data to use for
// rendering the contents of the web view. error: An error object if a problem
// occurred, or `nil` on success.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/createPDFWithConfiguration:completionHandler:
func (w WKWebView) CreatePDFWithConfigurationCompletionHandler(pdfConfiguration IWKPDFConfiguration, completionHandler DataErrorHandler) {
	_block1, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("createPDFWithConfiguration:completionHandler:"), pdfConfiguration, _block1)
}

// Creates a web archive of the web view’s contents asynchronously.
//
// completionHandler: The completion handler block to call when the web archive data is ready.
// This block has no return value and takes the following parameters:
//
// data: A data object that contains the web archive. error: An error object
// if an error occurs, or `nil` on success.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/createWebArchiveDataWithCompletionHandler:
func (w WKWebView) CreateWebArchiveDataWithCompletionHandler(completionHandler DataErrorHandler) {
	_block0, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("createWebArchiveDataWithCompletionHandler:"), _block0)
}

// Searches for the specified string in the web view’s content.
//
// string: The search string to use.
//
// configuration: The search parameters. Use this object to specify whether the search is
// case-sensitive, moves forward or backward, and wraps when it reaches the
// end of the page.
//
// completionHandler: The completion handler to call with the results of the search. This handler
// has no return value and takes the following parameter:
//
// result: The object that contains the results of the search.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/findString:withConfiguration:completionHandler:
func (w WKWebView) FindStringWithConfigurationCompletionHandler(string_ string, configuration IWKFindConfiguration, completionHandler WKFindResultHandler) {
	_block2, _ := NewWKFindResultBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("findString:withConfiguration:completionHandler:"), objc.String(string_), configuration, _block2)
}

// Returns a Boolean value that indicates whether WebKit natively supports
// resources with the specified URL scheme.
//
// urlScheme: The URL scheme associated with the resource.
//
// # Return Value
//
// true if WebKit provides native support for the URL scheme, or false if it
// doesn’t.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/handlesURLScheme(_:)
func (_WKWebViewClass WKWebViewClass) HandlesURLScheme(urlScheme string) bool {
	rv := objc.Send[bool](objc.ID(_WKWebViewClass.class), objc.Sel("handlesURLScheme:"), objc.String(urlScheme))
	return rv
}

// The object that contains the configuration details for the web view.
//
// # Discussion
//
// Use the object in this property to obtain information about your web
// view’s configuration. Because this property returns a copy of the
// configuration object, changes you make to that object don’t affect the
// web view’s configuration.
//
// If you didn’t create your web view using the [InitWithFrameConfiguration]
// method, this property contains a default configuration object.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/configuration
func (w WKWebView) Configuration() IWKWebViewConfiguration {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("configuration"))
	return WKWebViewConfigurationFromID(objc.ID(rv))
}

// The object you use to integrate custom user interface elements, such as
// contextual menus or panels, into web view interactions.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/uiDelegate
func (w WKWebView) UIDelegate() WKUIDelegate {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("UIDelegate"))
	return WKUIDelegateObjectFromID(rv)
}
func (w WKWebView) SetUIDelegate(value WKUIDelegate) {
	objc.Send[struct{}](w.ID, objc.Sel("setUIDelegate:"), value)
}

// The object you use to manage navigation behavior for the web view.
//
// # Discussion
//
// Provide a delegate object when you want to manage or restrict navigation in
// your web content, track the progress of navigation requests, and handle
// authentication challenges for any new content. The object you specify must
// conform to the [WKNavigationDelegate] protocol.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/navigationDelegate
func (w WKWebView) NavigationDelegate() WKNavigationDelegate {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("navigationDelegate"))
	return WKNavigationDelegateObjectFromID(rv)
}
func (w WKWebView) SetNavigationDelegate(value WKNavigationDelegate) {
	objc.Send[struct{}](w.ID, objc.Sel("setNavigationDelegate:"), value)
}

// A Boolean value that indicates whether the view is currently loading
// content.
//
// # Discussion
//
// Set to `true` if the receiver is still loading content; otherwise, `false.`
// The [WKWebView] class is key-value observing (KVO) compliant for this
// property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/isLoading
func (w WKWebView) Loading() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isLoading"))
	return rv
}

// An estimate of what fraction of the current navigation has been loaded.
//
// # Discussion
//
// This value ranges from `0.0` to `1.0` based on the total number of bytes
// received, including the main document and all of its potential
// subresources. After navigation loading completes, the `estimatedProgress`
// value remains at `1.0` until a new navigation starts, at which point the
// `estimatedProgress` value resets to `0.0`. The [WKWebView] class is
// key-value observing (KVO) compliant for this property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/estimatedProgress
func (w WKWebView) EstimatedProgress() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("estimatedProgress"))
	return rv
}

// A Boolean value that indicates whether you can inspect the view with Safari
// Web Inspector.
//
// # Discussion
//
// Defaults to `false`.
//
// Set to `true` at any point in the view’s lifetime to allow Safari Web
// Inspector access to inspect the view’s content. Then, select your view in
// Safari’s Develop menu for either your computer or an attached device to
// inspect it.
//
// If you set this value to `false` during inspection, the system immediately
// closes Safari Web Inspector and does not provide any further information
// about the web content.
//
// For more information, see [Enabling the Inspection of Web Content in Apps].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/isInspectable
//
// [Enabling the Inspection of Web Content in Apps]: https://webkit.org/blog/13936/enabling-the-inspection-of-web-content-in-apps/
func (w WKWebView) Inspectable() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isInspectable"))
	return rv
}
func (w WKWebView) SetInspectable(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setInspectable:"), value)
}

// The page title.
//
// # Discussion
//
// [WKWebView] is key-value observing (KVO) compliant for this property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/title
func (w WKWebView) Title() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// The URL for the current webpage.
//
// # Discussion
//
// This property contains the URL for the webpage that the web view currently
// displays. Use this URL in places where you reflect the webpage address in
// your app’s user interface.
//
// [WKWebView] is key-value observing (KVO) compliant for this property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/url
func (w WKWebView) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The media type for the contents of the web view.
//
// # Discussion
//
// When the value of this property is `nil`, the web view derives the current
// media type from the CSS media property of its content. If you assign a
// value other than `nil` to this property, the web view uses the value you
// provide instead. The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/mediaType
func (w WKWebView) MediaType() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("mediaType"))
	return foundation.NSStringFromID(rv).String()
}
func (w WKWebView) SetMediaType(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setMediaType:"), objc.String(value))
}

// The custom user agent string.
//
// # Discussion
//
// Use this property to specify a custom user agent string for your web view.
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/customUserAgent
func (w WKWebView) CustomUserAgent() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("customUserAgent"))
	return foundation.NSStringFromID(rv).String()
}
func (w WKWebView) SetCustomUserAgent(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setCustomUserAgent:"), objc.String(value))
}

// The trust management object you use to evaluate trust for the current
// webpage.
//
// # Discussion
//
// Use the object in this property to validate the webpage’s certificate and
// associated credentials. [WKWebView] is key-value observing (KVO) compliant
// for this property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/serverTrust
func (w WKWebView) ServerTrust() security.SecTrustRef {
	rv := objc.Send[security.SecTrustRef](w.ID, objc.Sel("serverTrust"))
	return security.SecTrustRef(rv)
}

// A Boolean value that indicates whether the web view loaded all resources on
// the page through securely encrypted connections.
//
// # Discussion
//
// [WKWebView] is key-value observing (KVO) compliant for this property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/hasOnlySecureContent
func (w WKWebView) HasOnlySecureContent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hasOnlySecureContent"))
	return rv
}

// The theme color that the system gets from the first valid meta tag in the
// webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/themeColor
func (w WKWebView) ThemeColor() appkit.NSColor {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("themeColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}

// The color the web view displays behind the active page, visible when the
// user scrolls beyond the bounds of the page.
//
// # Discussion
//
// The web view derives the default value of this property from the content of
// the page, using the background colors of the “ and “ elements with the
// background color of the web view. To override the default color, set this
// property to a new color.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/underPageBackgroundColor
func (w WKWebView) UnderPageBackgroundColor() appkit.NSColor {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("underPageBackgroundColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (w WKWebView) SetUnderPageBackgroundColor(value appkit.NSColor) {
	objc.Send[struct{}](w.ID, objc.Sel("setUnderPageBackgroundColor:"), value)
}

// The scale factor by which the web view scales content relative to its
// bounds.
//
// # Discussion
//
// The default value of this property is `1.0`, which displays the content
// without any scaling. Changing the value of this property is equivalent to
// setting the CSS `zoom` property on all page content.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/pageZoom
func (w WKWebView) PageZoom() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("pageZoom"))
	return rv
}
func (w WKWebView) SetPageZoom(value float64) {
	objc.Send[struct{}](w.ID, objc.Sel("setPageZoom:"), value)
}

// A Boolean value that indicates whether magnify gestures change the web
// view’s magnification.
//
// # Discussion
//
// The default value is false. You can set the `magnification` property even
// if `allowsMagnification` is set to false.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/allowsMagnification
func (w WKWebView) AllowsMagnification() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsMagnification"))
	return rv
}
func (w WKWebView) SetAllowsMagnification(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsMagnification:"), value)
}

// The factor by which the page content is currently scaled.
//
// # Discussion
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/magnification
func (w WKWebView) Magnification() float64 {
	rv := objc.Send[float64](w.ID, objc.Sel("magnification"))
	return rv
}
func (w WKWebView) SetMagnification(value float64) {
	objc.Send[struct{}](w.ID, objc.Sel("setMagnification:"), value)
}

// An enumeration case that indicates whether the webpage is using the camera
// to capture images or video.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/cameraCaptureState
func (w WKWebView) CameraCaptureState() WKMediaCaptureState {
	rv := objc.Send[WKMediaCaptureState](w.ID, objc.Sel("cameraCaptureState"))
	return WKMediaCaptureState(rv)
}

// An enumeration case that indicates whether the webpage is using the
// microphone to capture audio.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/microphoneCaptureState
func (w WKWebView) MicrophoneCaptureState() WKMediaCaptureState {
	rv := objc.Send[WKMediaCaptureState](w.ID, objc.Sel("microphoneCaptureState"))
	return WKMediaCaptureState(rv)
}

// A Boolean value that indicates whether horizontal swipe gestures trigger
// backward and forward page navigation.
//
// # Discussion
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/allowsBackForwardNavigationGestures
func (w WKWebView) AllowsBackForwardNavigationGestures() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsBackForwardNavigationGestures"))
	return rv
}
func (w WKWebView) SetAllowsBackForwardNavigationGestures(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsBackForwardNavigationGestures:"), value)
}

// The web view’s back-forward list.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/backForwardList
func (w WKWebView) BackForwardList() IWKBackForwardList {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("backForwardList"))
	return WKBackForwardListFromID(objc.ID(rv))
}

// A Boolean value that indicates whether there is a valid back item in the
// back-forward list.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/canGoBack
func (w WKWebView) CanGoBack() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("canGoBack"))
	return rv
}

// A Boolean value that indicates whether there is a valid forward item in the
// back-forward list.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/canGoForward
func (w WKWebView) CanGoForward() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("canGoForward"))
	return rv
}

// A Boolean value that determines whether pressing a link displays a preview
// of the destination for the link.
//
// # Discussion
//
// In iOS, this property is available on devices that support 3D Touch. In iOS
// 10 and later, the default value is true; in previous versions of iOS, the
// default value is false.
//
// If you set this property’s value to true, an iOS user can press links to
// preview link destinations and detected data such as addresses and phone
// numbers. Such previews are known to users as . If a user presses deeper on
// a link preview, the preview navigates (or , in user terminology) to the
// destination. Because pop navigation switches the user from your app to
// Safari, it is opt-in for iOS apps.
//
// If you want to support link preview in iOS but also want to keep users
// within your app, you can switch from using the [WKWebView] class to the
// [SFSafariViewController] class. If you are using a web view as an in-app
// browser, making this change is best practice. The Safari view controller
// class automatically supports link previews.
//
// In macOS, this property is available on devices with a Force Touch
// trackpad. The default value is true.
//
// With this property set to its default value of true, a macOS user can force
// click a link to display a preview window showing the link’s destination.
// If the user then clicks the preview, the destination opens in Safari.
//
// On both platforms, all types of detected data respond to presses when this
// property is set to true. That is, in iOS 9 and OS X 10.11, you cannot
// specify which types of data are detected in the [WKWebView] class.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/allowsLinkPreview
//
// [SFSafariViewController]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController
func (w WKWebView) AllowsLinkPreview() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsLinkPreview"))
	return rv
}
func (w WKWebView) SetAllowsLinkPreview(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsLinkPreview:"), value)
}

// An object you use to capture the current state of interaction in a web view
// so that you can restore that state later to another web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/interactionState
func (w WKWebView) InteractionState() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("interactionState"))
	return objectivec.Object{ID: rv}
}
func (w WKWebView) SetInteractionState(value objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setInteractionState:"), value)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/fullscreenState-swift.property
func (w WKWebView) FullscreenState() WKFullscreenState {
	rv := objc.Send[WKFullscreenState](w.ID, objc.Sel("fullscreenState"))
	return WKFullscreenState(rv)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/minimumViewportInset
func (w WKWebView) MinimumViewportInset() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](w.ID, objc.Sel("minimumViewportInset"))
	return foundation.NSEdgeInsets(rv)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/maximumViewportInset
func (w WKWebView) MaximumViewportInset() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](w.ID, objc.Sel("maximumViewportInset"))
	return foundation.NSEdgeInsets(rv)
}

// # Discussion
//
// A Boolean value indicating whether Screen Time blocking has occurred.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/isBlockedByScreenTime
func (w WKWebView) IsBlockedByScreenTime() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isBlockedByScreenTime"))
	return rv
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/isWritingToolsActive
func (w WKWebView) WritingToolsActive() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isWritingToolsActive"))
	return rv
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/obscuredContentInsets
func (w WKWebView) ObscuredContentInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](w.ID, objc.Sel("obscuredContentInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (w WKWebView) SetObscuredContentInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](w.ID, objc.Sel("setObscuredContentInsets:"), value)
}

// StartDownloadUsingRequest is a synchronous wrapper around [WKWebView.StartDownloadUsingRequestCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) StartDownloadUsingRequest(ctx context.Context, request foundation.NSURLRequest) (*WKDownload, error) {
	done := make(chan *WKDownload, 1)
	w.StartDownloadUsingRequestCompletionHandler(request, func(val *WKDownload) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ResumeDownloadFromResumeData is a synchronous wrapper around [WKWebView.ResumeDownloadFromResumeDataCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) ResumeDownloadFromResumeData(ctx context.Context, resumeData foundation.INSData) (*WKDownload, error) {
	done := make(chan *WKDownload, 1)
	w.ResumeDownloadFromResumeDataCompletionHandler(resumeData, func(val *WKDownload) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// PauseAllMediaPlayback is a synchronous wrapper around [WKWebView.PauseAllMediaPlaybackWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) PauseAllMediaPlayback(ctx context.Context) error {
	done := make(chan struct{}, 1)
	w.PauseAllMediaPlaybackWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestMediaPlaybackState is a synchronous wrapper around [WKWebView.RequestMediaPlaybackStateWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) RequestMediaPlaybackState(ctx context.Context) (int, error) {
	done := make(chan int, 1)
	w.RequestMediaPlaybackStateWithCompletionHandler(func(val int) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// SetAllMediaPlaybackSuspended is a synchronous wrapper around [WKWebView.SetAllMediaPlaybackSuspendedCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) SetAllMediaPlaybackSuspended(ctx context.Context, suspended bool) error {
	done := make(chan struct{}, 1)
	w.SetAllMediaPlaybackSuspendedCompletionHandler(suspended, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CloseAllMediaPresentations is a synchronous wrapper around [WKWebView.CloseAllMediaPresentationsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) CloseAllMediaPresentations(ctx context.Context) error {
	done := make(chan struct{}, 1)
	w.CloseAllMediaPresentationsWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetCameraCaptureState is a synchronous wrapper around [WKWebView.SetCameraCaptureStateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) SetCameraCaptureState(ctx context.Context, state WKMediaCaptureState) error {
	done := make(chan struct{}, 1)
	w.SetCameraCaptureStateCompletionHandler(state, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetMicrophoneCaptureState is a synchronous wrapper around [WKWebView.SetMicrophoneCaptureStateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) SetMicrophoneCaptureState(ctx context.Context, state WKMediaCaptureState) error {
	done := make(chan struct{}, 1)
	w.SetMicrophoneCaptureStateCompletionHandler(state, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EvaluateJavaScript is a synchronous wrapper around [WKWebView.EvaluateJavaScriptCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) EvaluateJavaScript(ctx context.Context, javaScriptString string) (objectivec.IObject, error) {
	type result struct {
		val objectivec.IObject
		err error
	}
	done := make(chan result, 1)
	w.EvaluateJavaScriptCompletionHandler(javaScriptString, func(val objectivec.IObject, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// TakeSnapshotWithConfiguration is a synchronous wrapper around [WKWebView.TakeSnapshotWithConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) TakeSnapshotWithConfiguration(ctx context.Context, snapshotConfiguration IWKSnapshotConfiguration) (*appkit.NSImage, error) {
	type result struct {
		val *appkit.NSImage
		err error
	}
	done := make(chan result, 1)
	w.TakeSnapshotWithConfigurationCompletionHandler(snapshotConfiguration, func(val *appkit.NSImage, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchDataOfTypes is a synchronous wrapper around [WKWebView.FetchDataOfTypesCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) FetchDataOfTypes(ctx context.Context, dataTypes WKWebViewDataType) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	w.FetchDataOfTypesCompletionHandler(dataTypes, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RestoreData is a synchronous wrapper around [WKWebView.RestoreDataCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) RestoreData(ctx context.Context, data foundation.INSData) error {
	done := make(chan error, 1)
	w.RestoreDataCompletionHandler(data, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CreatePDFWithConfiguration is a synchronous wrapper around [WKWebView.CreatePDFWithConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) CreatePDFWithConfiguration(ctx context.Context, pdfConfiguration IWKPDFConfiguration) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	w.CreatePDFWithConfigurationCompletionHandler(pdfConfiguration, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// CreateWebArchiveData is a synchronous wrapper around [WKWebView.CreateWebArchiveDataWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) CreateWebArchiveData(ctx context.Context) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	w.CreateWebArchiveDataWithCompletionHandler(func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FindStringWithConfiguration is a synchronous wrapper around [WKWebView.FindStringWithConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WKWebView) FindStringWithConfiguration(ctx context.Context, string_ string, configuration IWKFindConfiguration) (*WKFindResult, error) {
	done := make(chan *WKFindResult, 1)
	w.FindStringWithConfigurationCompletionHandler(string_, configuration, func(val *WKFindResult) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
