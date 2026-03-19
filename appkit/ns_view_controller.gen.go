// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSViewController] class.
var (
	_NSViewControllerClass     NSViewControllerClass
	_NSViewControllerClassOnce sync.Once
)

func getNSViewControllerClass() NSViewControllerClass {
	_NSViewControllerClassOnce.Do(func() {
		_NSViewControllerClass = NSViewControllerClass{class: objc.GetClass("NSViewController")}
	})
	return _NSViewControllerClass
}

// GetNSViewControllerClass returns the class object for NSViewController.
func GetNSViewControllerClass() NSViewControllerClass {
	return getNSViewControllerClass()
}

type NSViewControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSViewControllerClass) Alloc() NSViewController {
	rv := objc.Send[NSViewController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A controller that manages a view, typically loaded from a nib file.
//
// # Overview
// 
// View controller management includes:
// 
// - Memory management of top-level objects similar to that performed by the
// [NSWindowController] class, taking the same care to prevent reference
// cycles when controls are bound to the nib file’s owner. - Declaring a
// generic [View] property, to make it easy to establish bindings in the nib
// to an object that isn’t yet known at nib-loading time or readily
// available to the code that’s doing the nib loading. - Implementing the
// key-value binding NSEditor informal protocol, so that apps using a view
// controller can easily make bound controls in the views commit or discard
// changes by the user.
// 
// In macOS 10.10 and later, a view controller offers a full set of life cycle
// methods, allowing you to manage the content of a window in a way that is on
// a par with iOS view controller management. These methods, presented in
// order here to reflect a typical cycle, are:
// 
// - [NSViewController.ViewDidLoad]
// - [NSViewController.ViewWillAppear]
// - [NSViewController.ViewDidAppear]
// 
// - [NSViewController.UpdateViewConstraints] - [NSViewController.ViewWillLayout] - [NSViewController.ViewDidLayout] -
// [NSViewController.ViewWillDisappear] - [NSViewController.ViewDidDisappear]
// 
// In addition, in macOS 10.10 and later, a view controller participates in
// the responder chain. You can implement action methods directly in the view
// controller. Corresponding actions that originate in the view controller’s
// view proceed up the responder chain and are handled by those methods.
// 
// Prior to OS X v10.10, a typical usage pattern for loading a nib file was to
// subclass [NSViewController] and override its [NSViewController.LoadView] method to call
// `[super loadView]`. But in macOS 10.10 and later, the [NSViewController.LoadView] method
// automatically looks for a nib file with the same name as the view
// controller. To take advantage of this behavior, name a nib file after its
// corresponding view controller and pass `nil` to both parameters of the
// [NSViewController.InitWithNibNameBundle] method.
// 
// A view controller employs lazy loading of its view: Immediately after a
// view controller is loaded into memory, the value of its [NSViewController.ViewLoaded]
// property is [false]. The value changes to [true] after the [NSViewController.LoadView]
// method returns and just before the system calls the [NSViewController.ViewDidLoad] method.
// 
// A view controller is meant to be highly reusable, such as for dynamically
// representing various objects. For example, the [AddAccessoryController]
// methods of the [NSPageLayout] and [NSPrintPanel] classes take an
// [NSViewController] instance as the argument, and set the
// [NSViewController.RepresentedObject] property to the [NSPrintInfo] object that is to be
// shown to the user. This allows a developer to easily create new printing
// accessory views using bindings and the [NSPrintInfo] class’s key-value
// coding and key-value observing compliance. When the user dismisses a
// printing dialog, the [NSPageLayout] and [NSPrintPanel] classes each send
// NSEditor messages to each accessory view controller to ensure that the
// user’s changes have been committed or discarded properly. The titles of
// the accessories are retrieved from the view controllers and shown to the
// user in menus that the user can choose from.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating A View Controller
//
//   - [NSViewController.InitWithNibNameBundle]: Returns a view controller object initialized to the nib file in the specified bundle.
//   - [NSViewController.LoadView]: Instantiates a view from a nib file and sets the value of the [view](<doc://com.apple.appkit/documentation/AppKit/NSViewController/view>) property.
//
// # Represented Object
//
//   - [NSViewController.RepresentedObject]: The object whose value is presented in the receiver’s primary view.
//   - [NSViewController.SetRepresentedObject]
//
// # Nib Properties
//
//   - [NSViewController.NibBundle]: The nib bundle to be loaded to instantiate the receiver’s primary view.
//   - [NSViewController.NibName]: The name of the nib file to be loaded to instantiate the receiver’s primary view.
//
// # View Properties
//
//   - [NSViewController.View]: The view controller’s primary view.
//   - [NSViewController.SetView]
//   - [NSViewController.Title]: The localized title of the receiver’s primary view.
//   - [NSViewController.SetTitle]
//
// # NSEditor Conformance
//
//   - [NSViewController.CommitEditingWithDelegateDidCommitSelectorContextInfo]: Attempt to commit any currently edited results of the receiver.
//   - [NSViewController.CommitEditing]: Returns whether the receiver was able to commit any pending edits.
//   - [NSViewController.DiscardEditing]: Causes the receiver to discard any changes, restoring the previous values.
//
// # Using a Storyboard
//
//   - [NSViewController.Storyboard]: The storyboard from which the view controller was loaded.
//   - [NSViewController.DismissController]
//
// # Responding to View Events
//
//   - [NSViewController.ViewDidLoad]: Called after the view controller’s view has been loaded into memory.
//   - [NSViewController.LoadViewIfNeeded]
//   - [NSViewController.ViewLoaded]: A Boolean value indicating whether the view controller’s view is loaded into memory.
//   - [NSViewController.ViewIfLoaded]
//   - [NSViewController.ViewWillAppear]: Called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
//   - [NSViewController.ViewDidAppear]: Called when the view controller’s view is fully transitioned onto the screen.
//   - [NSViewController.ViewWillDisappear]: Called when the view controller’s view is about to be removed from the view hierarchy in the window.
//   - [NSViewController.ViewDidDisappear]: Called after the view controller’s view is removed from the view hierarchy in a window.
//
// # Managing View Layout
//
//   - [NSViewController.PreferredContentSize]: The desired size of the view controller’s view, in screen units.
//   - [NSViewController.SetPreferredContentSize]
//   - [NSViewController.UpdateViewConstraints]: Called during Auto Layout constraint updating to enable the view controller to mediate the process.
//   - [NSViewController.ViewWillLayout]: Called just before the [layout()](<doc://com.apple.appkit/documentation/AppKit/NSView/layout()>) method of the view controller’s view is called.
//   - [NSViewController.ViewDidLayout]: Called immediately after the [layout()](<doc://com.apple.appkit/documentation/AppKit/NSView/layout()>) method of the view controller’s view is called.
//
// # Managing Child View Controllers in a Custom Container
//
//   - [NSViewController.AddChildViewController]: A convenience method for adding a child view controller at the end of the [children](<doc://com.apple.appkit/documentation/AppKit/NSViewController/children>) array.
//   - [NSViewController.ChildViewControllers]: An array of view controllers that are hierarchical children of the view controller.
//   - [NSViewController.SetChildViewControllers]
//   - [NSViewController.TransitionFromViewControllerToViewControllerOptionsCompletionHandler]: Performs a transition between two sibling child view controllers of the view controller.
//   - [NSViewController.InsertChildViewControllerAtIndex]: Inserts a specified child view controller into the [children](<doc://com.apple.appkit/documentation/AppKit/NSViewController/children>) array at a specified position.
//   - [NSViewController.RemoveChildViewControllerAtIndex]: Removes a specified child controller from the view controller.
//   - [NSViewController.RemoveFromParentViewController]: Removes the called view controller from its parent view controller.
//   - [NSViewController.PreferredContentSizeDidChangeForViewController]: Called when there is a change in value of the [preferredContentSize](<doc://com.apple.appkit/documentation/AppKit/NSViewController/preferredContentSize>) property of a child view controller or a presented view controller.
//
// # Presenting Another View Controller’s Content
//
//   - [NSViewController.PresentViewControllerAnimator]: Presents another view controller using a specified, custom animator for presentation and dismissal.
//   - [NSViewController.DismissViewController]: Dismisses a presented view controller, using the same animator that presented it.
//   - [NSViewController.PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior]: Presents another view controller as a popover.
//   - [NSViewController.PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent]
//   - [NSViewController.PresentViewControllerAsModalWindow]: Presents another view controller as a modal window, also known as an alert.
//   - [NSViewController.PresentViewControllerAsSheet]: Presents another view controller as a sheet.
//
// # Getting Related View Controllers
//
//   - [NSViewController.ParentViewController]: The immediate ancestor view controller of the view controller.
//   - [NSViewController.PresentedViewControllers]: The view controllers, if any, that are currently presented by the view controller.
//   - [NSViewController.PresentingViewController]: The view controller that presented the view controller or that presented its farthest ancestor view controller.
//
// # Configuring an App Extension View Controller
//
//   - [NSViewController.PreferredScreenOrigin]: For a view controller that is part of an app extension, the preferred screen origin.
//   - [NSViewController.SetPreferredScreenOrigin]
//   - [NSViewController.PreferredMaximumSize]: For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.
//   - [NSViewController.PreferredMinimumSize]: For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
//   - [NSViewController.ViewWillTransitionToSize]: For a view controller that is part of an app extension, called when its view is about to be resized.
//   - [NSViewController.SourceItemView]
//   - [NSViewController.SetSourceItemView]
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController
type NSViewController struct {
	NSResponder
}

// NSViewControllerFromID constructs a [NSViewController] from an objc.ID.
//
// A controller that manages a view, typically loaded from a nib file.
func NSViewControllerFromID(id objc.ID) NSViewController {
	return NSViewController{NSResponder: NSResponderFromID(id)}
}
// NOTE: NSViewController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSViewController] class.
//
// # Creating A View Controller
//
//   - [INSViewController.InitWithNibNameBundle]: Returns a view controller object initialized to the nib file in the specified bundle.
//   - [INSViewController.LoadView]: Instantiates a view from a nib file and sets the value of the [view](<doc://com.apple.appkit/documentation/AppKit/NSViewController/view>) property.
//
// # Represented Object
//
//   - [INSViewController.RepresentedObject]: The object whose value is presented in the receiver’s primary view.
//   - [INSViewController.SetRepresentedObject]
//
// # Nib Properties
//
//   - [INSViewController.NibBundle]: The nib bundle to be loaded to instantiate the receiver’s primary view.
//   - [INSViewController.NibName]: The name of the nib file to be loaded to instantiate the receiver’s primary view.
//
// # View Properties
//
//   - [INSViewController.View]: The view controller’s primary view.
//   - [INSViewController.SetView]
//   - [INSViewController.Title]: The localized title of the receiver’s primary view.
//   - [INSViewController.SetTitle]
//
// # NSEditor Conformance
//
//   - [INSViewController.CommitEditingWithDelegateDidCommitSelectorContextInfo]: Attempt to commit any currently edited results of the receiver.
//   - [INSViewController.CommitEditing]: Returns whether the receiver was able to commit any pending edits.
//   - [INSViewController.DiscardEditing]: Causes the receiver to discard any changes, restoring the previous values.
//
// # Using a Storyboard
//
//   - [INSViewController.Storyboard]: The storyboard from which the view controller was loaded.
//   - [INSViewController.DismissController]
//
// # Responding to View Events
//
//   - [INSViewController.ViewDidLoad]: Called after the view controller’s view has been loaded into memory.
//   - [INSViewController.LoadViewIfNeeded]
//   - [INSViewController.ViewLoaded]: A Boolean value indicating whether the view controller’s view is loaded into memory.
//   - [INSViewController.ViewIfLoaded]
//   - [INSViewController.ViewWillAppear]: Called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
//   - [INSViewController.ViewDidAppear]: Called when the view controller’s view is fully transitioned onto the screen.
//   - [INSViewController.ViewWillDisappear]: Called when the view controller’s view is about to be removed from the view hierarchy in the window.
//   - [INSViewController.ViewDidDisappear]: Called after the view controller’s view is removed from the view hierarchy in a window.
//
// # Managing View Layout
//
//   - [INSViewController.PreferredContentSize]: The desired size of the view controller’s view, in screen units.
//   - [INSViewController.SetPreferredContentSize]
//   - [INSViewController.UpdateViewConstraints]: Called during Auto Layout constraint updating to enable the view controller to mediate the process.
//   - [INSViewController.ViewWillLayout]: Called just before the [layout()](<doc://com.apple.appkit/documentation/AppKit/NSView/layout()>) method of the view controller’s view is called.
//   - [INSViewController.ViewDidLayout]: Called immediately after the [layout()](<doc://com.apple.appkit/documentation/AppKit/NSView/layout()>) method of the view controller’s view is called.
//
// # Managing Child View Controllers in a Custom Container
//
//   - [INSViewController.AddChildViewController]: A convenience method for adding a child view controller at the end of the [children](<doc://com.apple.appkit/documentation/AppKit/NSViewController/children>) array.
//   - [INSViewController.ChildViewControllers]: An array of view controllers that are hierarchical children of the view controller.
//   - [INSViewController.SetChildViewControllers]
//   - [INSViewController.TransitionFromViewControllerToViewControllerOptionsCompletionHandler]: Performs a transition between two sibling child view controllers of the view controller.
//   - [INSViewController.InsertChildViewControllerAtIndex]: Inserts a specified child view controller into the [children](<doc://com.apple.appkit/documentation/AppKit/NSViewController/children>) array at a specified position.
//   - [INSViewController.RemoveChildViewControllerAtIndex]: Removes a specified child controller from the view controller.
//   - [INSViewController.RemoveFromParentViewController]: Removes the called view controller from its parent view controller.
//   - [INSViewController.PreferredContentSizeDidChangeForViewController]: Called when there is a change in value of the [preferredContentSize](<doc://com.apple.appkit/documentation/AppKit/NSViewController/preferredContentSize>) property of a child view controller or a presented view controller.
//
// # Presenting Another View Controller’s Content
//
//   - [INSViewController.PresentViewControllerAnimator]: Presents another view controller using a specified, custom animator for presentation and dismissal.
//   - [INSViewController.DismissViewController]: Dismisses a presented view controller, using the same animator that presented it.
//   - [INSViewController.PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior]: Presents another view controller as a popover.
//   - [INSViewController.PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent]
//   - [INSViewController.PresentViewControllerAsModalWindow]: Presents another view controller as a modal window, also known as an alert.
//   - [INSViewController.PresentViewControllerAsSheet]: Presents another view controller as a sheet.
//
// # Getting Related View Controllers
//
//   - [INSViewController.ParentViewController]: The immediate ancestor view controller of the view controller.
//   - [INSViewController.PresentedViewControllers]: The view controllers, if any, that are currently presented by the view controller.
//   - [INSViewController.PresentingViewController]: The view controller that presented the view controller or that presented its farthest ancestor view controller.
//
// # Configuring an App Extension View Controller
//
//   - [INSViewController.PreferredScreenOrigin]: For a view controller that is part of an app extension, the preferred screen origin.
//   - [INSViewController.SetPreferredScreenOrigin]
//   - [INSViewController.PreferredMaximumSize]: For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.
//   - [INSViewController.PreferredMinimumSize]: For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
//   - [INSViewController.ViewWillTransitionToSize]: For a view controller that is part of an app extension, called when its view is about to be resized.
//   - [INSViewController.SourceItemView]
//   - [INSViewController.SetSourceItemView]
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController
type INSViewController interface {
	INSResponder
	NSSeguePerforming
	NSUserInterfaceItemIdentification

	// Topic: Creating A View Controller

	// Returns a view controller object initialized to the nib file in the specified bundle.
	InitWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSViewController
	// Instantiates a view from a nib file and sets the value of the [view](<doc://com.apple.appkit/documentation/AppKit/NSViewController/view>) property.
	LoadView()

	// Topic: Represented Object

	// The object whose value is presented in the receiver’s primary view.
	RepresentedObject() objectivec.IObject
	SetRepresentedObject(value objectivec.IObject)

	// Topic: Nib Properties

	// The nib bundle to be loaded to instantiate the receiver’s primary view.
	NibBundle() foundation.NSBundle
	// The name of the nib file to be loaded to instantiate the receiver’s primary view.
	NibName() NSNibName

	// Topic: View Properties

	// The view controller’s primary view.
	View() INSView
	SetView(value INSView)
	// The localized title of the receiver’s primary view.
	Title() string
	SetTitle(value string)

	// Topic: NSEditor Conformance

	// Attempt to commit any currently edited results of the receiver.
	CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo uintptr)
	// Returns whether the receiver was able to commit any pending edits.
	CommitEditing() bool
	// Causes the receiver to discard any changes, restoring the previous values.
	DiscardEditing()

	// Topic: Using a Storyboard

	// The storyboard from which the view controller was loaded.
	Storyboard() INSStoryboard
	DismissController(sender objectivec.IObject)

	// Topic: Responding to View Events

	// Called after the view controller’s view has been loaded into memory.
	ViewDidLoad()
	LoadViewIfNeeded()
	// A Boolean value indicating whether the view controller’s view is loaded into memory.
	ViewLoaded() bool
	ViewIfLoaded() INSView
	// Called after the view controller’s view has been loaded into memory is about to be added to the view hierarchy in the window.
	ViewWillAppear()
	// Called when the view controller’s view is fully transitioned onto the screen.
	ViewDidAppear()
	// Called when the view controller’s view is about to be removed from the view hierarchy in the window.
	ViewWillDisappear()
	// Called after the view controller’s view is removed from the view hierarchy in a window.
	ViewDidDisappear()

	// Topic: Managing View Layout

	// The desired size of the view controller’s view, in screen units.
	PreferredContentSize() corefoundation.CGSize
	SetPreferredContentSize(value corefoundation.CGSize)
	// Called during Auto Layout constraint updating to enable the view controller to mediate the process.
	UpdateViewConstraints()
	// Called just before the [layout()](<doc://com.apple.appkit/documentation/AppKit/NSView/layout()>) method of the view controller’s view is called.
	ViewWillLayout()
	// Called immediately after the [layout()](<doc://com.apple.appkit/documentation/AppKit/NSView/layout()>) method of the view controller’s view is called.
	ViewDidLayout()

	// Topic: Managing Child View Controllers in a Custom Container

	// A convenience method for adding a child view controller at the end of the [children](<doc://com.apple.appkit/documentation/AppKit/NSViewController/children>) array.
	AddChildViewController(childViewController INSViewController)
	// An array of view controllers that are hierarchical children of the view controller.
	ChildViewControllers() []NSViewController
	SetChildViewControllers(value []NSViewController)
	// Performs a transition between two sibling child view controllers of the view controller.
	TransitionFromViewControllerToViewControllerOptionsCompletionHandler(fromViewController INSViewController, toViewController INSViewController, options NSViewControllerTransitionOptions, completion VoidHandler)
	// Inserts a specified child view controller into the [children](<doc://com.apple.appkit/documentation/AppKit/NSViewController/children>) array at a specified position.
	InsertChildViewControllerAtIndex(childViewController INSViewController, index int)
	// Removes a specified child controller from the view controller.
	RemoveChildViewControllerAtIndex(index int)
	// Removes the called view controller from its parent view controller.
	RemoveFromParentViewController()
	// Called when there is a change in value of the [preferredContentSize](<doc://com.apple.appkit/documentation/AppKit/NSViewController/preferredContentSize>) property of a child view controller or a presented view controller.
	PreferredContentSizeDidChangeForViewController(viewController INSViewController)

	// Topic: Presenting Another View Controller’s Content

	// Presents another view controller using a specified, custom animator for presentation and dismissal.
	PresentViewControllerAnimator(viewController INSViewController, animator NSViewControllerPresentationAnimator)
	// Dismisses a presented view controller, using the same animator that presented it.
	DismissViewController(viewController INSViewController)
	// Presents another view controller as a popover.
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController INSViewController, positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge, behavior NSPopoverBehavior)
	PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController INSViewController, positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge, behavior NSPopoverBehavior, hasFullSizeContent bool)
	// Presents another view controller as a modal window, also known as an alert.
	PresentViewControllerAsModalWindow(viewController INSViewController)
	// Presents another view controller as a sheet.
	PresentViewControllerAsSheet(viewController INSViewController)

	// Topic: Getting Related View Controllers

	// The immediate ancestor view controller of the view controller.
	ParentViewController() INSViewController
	// The view controllers, if any, that are currently presented by the view controller.
	PresentedViewControllers() []NSViewController
	// The view controller that presented the view controller or that presented its farthest ancestor view controller.
	PresentingViewController() INSViewController

	// Topic: Configuring an App Extension View Controller

	// For a view controller that is part of an app extension, the preferred screen origin.
	PreferredScreenOrigin() corefoundation.CGPoint
	SetPreferredScreenOrigin(value corefoundation.CGPoint)
	// For a view controller that is part of an app extension, the largest allowable size for the app extension’s primary view, in screen units.
	PreferredMaximumSize() corefoundation.CGSize
	// For a view controller that is part of an app extension, the smallest allowable size for the app extension’s primary view, in screen units.
	PreferredMinimumSize() corefoundation.CGSize
	// For a view controller that is part of an app extension, called when its view is about to be resized.
	ViewWillTransitionToSize(newSize corefoundation.CGSize)
	SourceItemView() INSView
	SetSourceItemView(value INSView)

	CommitEditingAndReturnError() (bool, error)
}

// Init initializes the instance.
func (v NSViewController) Init() NSViewController {
	rv := objc.Send[NSViewController](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NSViewController) Autorelease() NSViewController {
	rv := objc.Send[NSViewController](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSViewController creates a new NSViewController instance.
func NewNSViewController() NSViewController {
	class := getNSViewControllerClass()
	rv := objc.Send[NSViewController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewViewControllerWithCoder(coder foundation.INSCoder) NSViewController {
	instance := getNSViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSViewControllerFromID(rv)
}

// Returns a view controller object initialized to the nib file in the
// specified bundle.
//
// nibNameOrNil: The name of the nib file, without any leading path information.
//
// nibBundleOrNil: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
// 
// The initialized [NSViewController] object.
//
// # Discussion
// 
// The [NSViewController] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the Resources
// directory.
// 
// The specified nib file should typically have the class of the file’s
// owner set to [NSViewController], or a custom subclass, with the `view`
// outlet connected to a view.
// 
// If you pass in `nil` for `nibNameOrNil`, [NibName] returns `nil` and
// [LoadView] throws an exception; in this case you must set [View] before
// [View] is invoked, or override [LoadView].
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewViewControllerWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSViewController {
	instance := getNSViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return NSViewControllerFromID(rv)
}

// Returns a view controller object initialized to the nib file in the
// specified bundle.
//
// nibNameOrNil: The name of the nib file, without any leading path information.
//
// nibBundleOrNil: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
// 
// The initialized [NSViewController] object.
//
// # Discussion
// 
// The [NSViewController] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the Resources
// directory.
// 
// The specified nib file should typically have the class of the file’s
// owner set to [NSViewController], or a custom subclass, with the `view`
// outlet connected to a view.
// 
// If you pass in `nil` for `nibNameOrNil`, [NibName] returns `nil` and
// [LoadView] throws an exception; in this case you must set [View] before
// [View] is invoked, or override [LoadView].
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func (v NSViewController) InitWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSViewController {
	rv := objc.Send[NSViewController](v.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return rv
}
// Instantiates a view from a nib file and sets the value of the [View]
// property.
//
// # Discussion
// 
// This method connects an instantiated view from a nib file to the [View]
// property of the view controller. This method is called by the system, and
// is exposed in this class so you can override it to add behavior immediately
// before or after nib loading.
// 
// Do not call this method. If you require this method to be called, access
// the [View] property.
// 
// Do not invoke this method from other objects unless you take care to avoid
// redundant invocations. The default implementation of the [LoadView] method
// handles redundant invocations correctly, but a view controller subclass
// might not. To be safe, other objects should instead access a view
// controller’s [View] property.
// 
// The [LoadView] method first obtains the values of the view controller’s
// [NibName] and [NibBundle] properties. It then employs the [NSNib] class to
// instantiate the specified nib file via the
// [InstantiateWithOwnerTopLevelObjects] method, providing the view controller
// object as the `owner` parameter.
// 
// For this method to work correctly, you need to have specified the file’s
// owner of the nib file, in Interface Builder, to be [NSViewController]. You
// also need to have correctly connected the `view` outlet of the file’s
// owner to the intended view in the nib file. Then, at runtime, the nib
// loading machinery sets the value of the view controller’s [View] property
// to the nib file’s instantiated view.
// 
// Prior to OS X v10.10, the [LoadView] method did not provide well-defined
// behavior if the [NibName] property’s value was `nil`. In macOS 10.10 and
// later, however, you get correct behavior without specifying a nib name as
// long as the nib file’s name is the same as that of the view controller.
// For example, if you have a view controller subclass called
// [MyViewController] and a nib file with the same name, you can employ the
// convenient initialization pattern `[[MyViewController alloc] init]`.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/loadView()
func (v NSViewController) LoadView() {
	objc.Send[objc.ID](v.ID, objc.Sel("loadView"))
}
// Attempt to commit any currently edited results of the receiver.
//
// delegate: An object that can serve as the receiver’s delegate. It should implement
// the method specified by `didCommitSelector`.
//
// didCommitSelector: A selector that is invoked on delegate.
//
// contextInfo: Contextual information that is sent as the `contextInfo` argument to
// delegate when `didCommitSelector` is invoked.
//
// # Discussion
// 
// The receiver must have been registered as the editor of an object using
// [objectDidBeginEditing:], and has not yet been unregistered by a subsequent
// invocation of [objectDidEndEditing:]. When the committing has either
// succeeded or failed, send the `delegate` the message specified by
// `didCommitSelector`.
// 
// The `didCommitSelector` method must have the following method signature:.
// 
// If an error occurs while attempting to commit, for example if key-value
// coding validation fails, an implementation of this method should typically
// send the receiver’s view
// a[PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo] message,
// specifying the view’s containing window.
// 
// You may find this method useful in some situations when you want to ensure
// that pending changes are applied before a change in user interface state.
// For example, you may need to ensure that changes pending in a text field
// are applied before a window is closed. See also [CommitEditing] which
// performs a similar function but which allows you to handle any errors
// directly, although it provides no information beyond simple
// success/failure.
//
// [objectDidBeginEditing:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/objectDidBeginEditing:
// [objectDidEndEditing:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/objectDidEndEditing:
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/commitEditing(withDelegate:didCommit:contextInfo:)
func (v NSViewController) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](v.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}
// Returns whether the receiver was able to commit any pending edits.
//
// # Return Value
// 
// Returns [true] if the changes were successfully applied to the model,
// [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// A commit is denied if the receiver fails to apply the changes to the model
// object, perhaps due to a validation error.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/commitEditing()
func (v NSViewController) CommitEditing() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("commitEditing"))
	return rv
}
// Causes the receiver to discard any changes, restoring the previous values.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/discardEditing()
func (v NSViewController) DiscardEditing() {
	objc.Send[objc.ID](v.ID, objc.Sel("discardEditing"))
}
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-3n76y
func (v NSViewController) DismissController(sender objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("dismissController:"), sender)
}
// Called after the view controller’s view has been loaded into memory.
//
// # Discussion
// 
// You can override this method to perform tasks to immediately follow the
// setting of the [View] property.
// 
// Typically, your override would perform one-time instantiation and
// initialization of the contents of the view controller’s view. If you
// override this method, call this method on `super` at some point in your
// implementation in case a superclass also overrides this method.
// 
// For a view controller originating in a nib file, this method is called
// immediately after the [View] property is set. For a view controller created
// programmatically, this method is called immediately after the [LoadView]
// method completes.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidLoad()
func (v NSViewController) ViewDidLoad() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidLoad"))
}
// See: https://developer.apple.com/documentation/AppKit/NSViewController/loadViewIfNeeded()
func (v NSViewController) LoadViewIfNeeded() {
	objc.Send[objc.ID](v.ID, objc.Sel("loadViewIfNeeded"))
}
// Called after the view controller’s view has been loaded into memory is
// about to be added to the view hierarchy in the window.
//
// # Discussion
// 
// You can override this method to perform tasks prior to a view
// controller’s view getting added to view hierarchy, such as setting the
// view’s highlight color. This method is called when:
// 
// - The view is about to be added to the view hierarchy of the view
// controller - The view controller’s window is about to become visible,
// such as coming to the front or becoming unhidden
// 
// If you override this method, call this method on `super` at some point in
// your implementation in case a superclass also overrides this method.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillAppear()
func (v NSViewController) ViewWillAppear() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillAppear"))
}
// Called when the view controller’s view is fully transitioned onto the
// screen.
//
// # Discussion
// 
// This method is called after the completion of any drawing and animations
// involved in the initial appearance of the view. You can override this
// method to perform tasks appropriate for that time, such as work that should
// not interfere with the presentation animation, or starting an animation
// that you want to begin after the view appears.
// 
// If you override this method, call this method on `super` at some point in
// your implementation in case a superclass also overrides this method.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidAppear()
func (v NSViewController) ViewDidAppear() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidAppear"))
}
// Called when the view controller’s view is about to be removed from the
// view hierarchy in the window.
//
// # Discussion
// 
// You can override this method to perform tasks that are to precede the
// disappearance of the view controller’s view, such as stopping a
// continuous animation that you started in response to the [ViewDidAppear]
// method call. This method is called when:
// 
// - The view is about to be removed from the view hierarchy of the window -
// The view is about to be hidden or obscured, such as in the case of a view
// controller whose parent is a tab view controller and the user switched to
// another tab - The window is being closed
// 
// If you override this method, call this method on `super` at some point in
// your implementation in case a superclass also overrides this method.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillDisappear()
func (v NSViewController) ViewWillDisappear() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillDisappear"))
}
// Called after the view controller’s view is removed from the view
// hierarchy in a window.
//
// # Discussion
// 
// You can override this method to perform tasks associated with removing the
// view controller’s view from the window’s view hierarchy, such as
// releasing resources not needed when the view is not visible or no longer
// part of the window.
// 
// If you override this method, call this method on `super` at some point in
// your implementation in case a superclass also overrides this method.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidDisappear()
func (v NSViewController) ViewDidDisappear() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidDisappear"))
}
// Called during Auto Layout constraint updating to enable the view controller
// to mediate the process.
//
// # Discussion
// 
// This method gets called, for example, when the user interacts with a view
// in a way that causes the layout to change. When called, the default
// implementation of this method in turn calls the [UpdateConstraints] method
// on the view controller’s view.
// 
// You can override this method to update custom view constraints, as an
// alternative to subclassing the view controller’s view and overriding its
// [UpdateConstraints] method.
// 
// If you override this method, you must call this method on `super` at some
// point in your implementation or call the [UpdateConstraints] method on the
// view controller’s view.
// 
// This method is called only for apps that link against macOS 10.10 or later.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/updateViewConstraints()
func (v NSViewController) UpdateViewConstraints() {
	objc.Send[objc.ID](v.ID, objc.Sel("updateViewConstraints"))
}
// Called just before the [Layout] method of the view controller’s view is
// called.
//
// # Discussion
// 
// You can override this method to perform tasks to precede the layout of the
// view controller’s view, such as adjusting Auto Layout constraints. If you
// override this method, call this method on `super` at some point in your
// implementation in case a superclass also overrides this method.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillLayout()
func (v NSViewController) ViewWillLayout() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillLayout"))
}
// Called immediately after the [Layout] method of the view controller’s
// view is called.
//
// # Discussion
// 
// You can override this method to perform tasks to follow the completion of
// layout of the view controller’s view. If you override this method, call
// this method on `super` at some point in your implementation in case a
// superclass also overrides this method.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewDidLayout()
func (v NSViewController) ViewDidLayout() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidLayout"))
}
// A convenience method for adding a child view controller at the end of the
// [ChildViewControllers] array.
//
// childViewController: The view controller to be added to the end of the [ChildViewControllers]
// array.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/addChild(_:)
func (v NSViewController) AddChildViewController(childViewController INSViewController) {
	objc.Send[objc.ID](v.ID, objc.Sel("addChildViewController:"), childViewController)
}
// Performs a transition between two sibling child view controllers of the
// view controller.
//
// fromViewController: A child view controller whose view is visible in the view controller’s
// view hierarchy.
//
// toViewController: A child view controller whose view is not in the view controller’s view
// hierarchy.
//
// options: A bitmask of options that specify how you want to perform the transition
// animation. For the options, see the [NSViewController.TransitionOptions]
// enumeration.
// //
// [NSViewController.TransitionOptions]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions
//
// completion: A block called immediately after the transition animation completes.
//
// # Discussion
// 
// Use this method to transition between sibling child view controllers owned
// by a parent view controller (which is the receiver of this method).
// 
// This method adds the view in the `toViewController` view controller to the
// superview of the view in the `fromViewController` view controller.
// Likewise, this method removes the `fromViewController` view from the parent
// view controller’s view hierarchy at the appropriate time. It is important
// to allow this method to add and remove these views.
// 
// To create a parent/child relationship between view controllers, use the
// [AddChildViewController] method or the [InsertChildViewControllerAtIndex]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/transition(from:to:options:completionHandler:)
func (v NSViewController) TransitionFromViewControllerToViewControllerOptionsCompletionHandler(fromViewController INSViewController, toViewController INSViewController, options NSViewControllerTransitionOptions, completion VoidHandler) {
_block3, _cleanup3 := NewVoidBlock(completion)
	defer _cleanup3()
	objc.Send[objc.ID](v.ID, objc.Sel("transitionFromViewController:toViewController:options:completionHandler:"), fromViewController, toViewController, options, _block3)
}
// Inserts a specified child view controller into the [ChildViewControllers]
// array at a specified position.
//
// childViewController: The child view controller to add to the [ChildViewControllers] array.
//
// index: The index in the [ChildViewControllers] array at which to insert the child
// view controller. This value must not be greater than the count of elements
// in the array.
//
// # Discussion
// 
// You should instead use the [AddChildViewController] method unless you want
// to perform work on child view controllers as you add them. In that case,
// override this method to perform that work.
// 
// If a child view controller has a different parent when you call this
// method, the child is first be removed from its existing parent by calling
// the child’s [RemoveFromParentViewController] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/insertChild(_:at:)
func (v NSViewController) InsertChildViewControllerAtIndex(childViewController INSViewController, index int) {
	objc.Send[objc.ID](v.ID, objc.Sel("insertChildViewController:atIndex:"), childViewController, index)
}
// Removes a specified child controller from the view controller.
//
// index: The index in the [ChildViewControllers] array for the child view controller
// you want to remove.
//
// # Discussion
// 
// Override this method if you want to perform work during the removal of a
// child view controller. If you do override this method, in your
// implementation call this method on `super`.
// 
// If you just want to remove a child view controller, instead use the
// [RemoveFromParentViewController] method
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/removeChild(at:)
func (v NSViewController) RemoveChildViewControllerAtIndex(index int) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeChildViewControllerAtIndex:"), index)
}
// Removes the called view controller from its parent view controller.
//
// # Discussion
// 
// Use this method to remove a child view controller from its parent view
// controller, unless you want to perform work during the removal. In that
// case, instead override the [RemoveChildViewControllerAtIndex] method to
// perform that work and call that method.
// 
// This is a convenience method that calls the
// [RemoveChildViewControllerAtIndex] method, automatically supplying the
// appropriate index value as an argument.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/removeFromParent()
func (v NSViewController) RemoveFromParentViewController() {
	objc.Send[objc.ID](v.ID, objc.Sel("removeFromParentViewController"))
}
// Called when there is a change in value of the [PreferredContentSize]
// property of a child view controller or a presented view controller.
//
// viewController: The view controller whose [PreferredContentSize] property value changed.
//
// # Discussion
// 
// Override this method if you want to adjust layout when a child view
// controller or presented view controller changes its preferred content size.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/preferredContentSizeDidChange(for:)
func (v NSViewController) PreferredContentSizeDidChangeForViewController(viewController INSViewController) {
	objc.Send[objc.ID](v.ID, objc.Sel("preferredContentSizeDidChangeForViewController:"), viewController)
}
// Presents another view controller using a specified, custom animator for
// presentation and dismissal.
//
// viewController: The other view controller to present from the view controller.
//
// animator: The animation delegate to employ for presentation and dismissal of the
// other view controller. The animator that you specify is retained until the
// [DismissViewController] method is called and the dismissal animation
// completes.
//
// # Discussion
// 
// Do not call this method unless you want to use a custom animator. To use
// one of the standard animators to present another view controller, instead
// call one of the dedicated presentation methods:
// 
// - [PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior]
// - [PresentViewControllerAsModalWindow] - [PresentViewControllerAsSheet]
// 
// Each of these methods calls this method in turn. User interaction is
// blocked during presentation and dismissal.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:animator:)
func (v NSViewController) PresentViewControllerAnimator(viewController INSViewController, animator NSViewControllerPresentationAnimator) {
	objc.Send[objc.ID](v.ID, objc.Sel("presentViewController:animator:"), viewController, animator)
}
// Dismisses a presented view controller, using the same animator that
// presented it.
//
// viewController: The presented view controller that you are dismissing.
//
// # Discussion
// 
// In macOS, this is the universal way to dismiss a view controller, no matter
// how it was presented.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/dismiss(_:)-91my5
func (v NSViewController) DismissViewController(viewController INSViewController) {
	objc.Send[objc.ID](v.ID, objc.Sel("dismissViewController:"), viewController)
}
// Presents another view controller as a popover.
//
// viewController: The other view controller to present as a popover.
//
// positioningRect: The content size of the popover.
//
// positioningView: The view relative to which the popover should be positioned. Must not be
// `nil`, or else the view controller raises an [InvalidArgumentException]
// exception.
//
// preferredEdge: The edge of `positioningView` that the popover should prefer to be anchored
// to.
//
// behavior: The popover’s closing behavior. See the [NSPopover.Behavior] enumeration.
// //
// [NSPopover.Behavior]: https://developer.apple.com/documentation/AppKit/NSPopover/Behavior-swift.enum
//
// # Discussion
// 
// This method calls the [PresentViewControllerAnimator] method on `self` (the
// presenting view controller), and passes a popover animator to that method.
// 
// The presented view controller is the delegate and the content view
// controller of the popover. You can use [NSPopoverDelegate] methods to
// customize the popover.
// 
// To dismiss the popover, call the [DismissViewController] method on `self`
// (the presenting view controller).
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:)
func (v NSViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehavior(viewController INSViewController, positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge, behavior NSPopoverBehavior) {
	objc.Send[objc.ID](v.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:"), viewController, positioningRect, positioningView, preferredEdge, behavior)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/present(_:asPopoverRelativeTo:of:preferredEdge:behavior:hasFullSizeContent:)
func (v NSViewController) PresentViewControllerAsPopoverRelativeToRectOfViewPreferredEdgeBehaviorHasFullSizeContent(viewController INSViewController, positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge, behavior NSPopoverBehavior, hasFullSizeContent bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:hasFullSizeContent:"), viewController, positioningRect, positioningView, preferredEdge, behavior, hasFullSizeContent)
}
// Presents another view controller as a modal window, also known as an alert.
//
// viewController: The other view controller to present as a modal window.
//
// # Discussion
// 
// This method calls the [PresentViewControllerAnimator] method on `self` (the
// presenting view controller), and passes a modal window animator to that
// method.
// 
// The presented view controller is the delegate and the content view
// controller of its window. You can use [NSWindowDelegate] methods to prevent
// the closing of the modal window, if needed.
// 
// To dismiss the modal window, call the [DismissViewController] method on
// `self` (the presenting view controller).
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsModalWindow(_:)
func (v NSViewController) PresentViewControllerAsModalWindow(viewController INSViewController) {
	objc.Send[objc.ID](v.ID, objc.Sel("presentViewControllerAsModalWindow:"), viewController)
}
// Presents another view controller as a sheet.
//
// viewController: The other view controller to present as a sheet.
//
// # Discussion
// 
// This method calls the [PresentViewControllerAnimator] method on `self` (the
// presenting view controller), and passes a sheet animator to that method.
// 
// The presented view controller is the delegate and the content view
// controller of its sheet.
// 
// To dismiss the sheet, call the [DismissViewController] method on `self`
// (the presenting view controller).
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/presentAsSheet(_:)
func (v NSViewController) PresentViewControllerAsSheet(viewController INSViewController) {
	objc.Send[objc.ID](v.ID, objc.Sel("presentViewControllerAsSheet:"), viewController)
}
// For a view controller that is part of an app extension, called when its
// view is about to be resized.
//
// newSize: The new size for the view controller’s view.
//
// # Discussion
// 
// Override this method if you want to change layout in response to the change
// in size, potentially in an animated way.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewWillTransition(to:)
func (v NSViewController) ViewWillTransitionToSize(newSize corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillTransitionToSize:"), newSize)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSEditor/commitEditingWithoutPresentingError()
func (v NSViewController) CommitEditingAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("commitEditingAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("commitEditingAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
// Performs the specified segue.
//
// identifier: The string that uniquely identifies the segue in the storyboard file.
// 
// In Interface Builder, you can provide an identifier string to a segue using
// the inspector. Pass this string to this parameter.
//
// sender: The object that you want to use to initiate the segue. This parameter makes
// the object available to your implementation during the segue.
//
// # Discussion
// 
// Apps typically do not need to trigger segues programmatically. If needed,
// you can call this method to trigger a segue for an action that cannot be
// expressed in a storyboard file, such as a transition between scenes in
// different storyboards.
// 
// Typically, a segue is triggered by a user action, such as clicking a
// button. In Interface Builder, configure an object, such as a control
// embedded in the view controller’s view hierarchy, to trigger the segue.
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/performSegue(withIdentifier:sender:)
func (v NSViewController) PerformSegueWithIdentifierSender(identifier NSStoryboardSegueIdentifier, sender objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("performSegueWithIdentifier:sender:"), objc.String(string(identifier)), sender)
}
// Called when a segue is about to be performed.
//
// segue: The segue object containing information about the view controllers involved
// in the segue.
//
// sender: The object that initiated the segue. You might use this parameter to
// perform different actions based on which control (or other object)
// initiated the segue.
//
// # Discussion
// 
// The default implementation of this method does nothing; you can override it
// to pass relevant data to the new view controller or window controller,
// based on the context of the segue. The `segue` object describes the
// transition and includes references to both controllers involved in the
// segue.
// 
// Segues can be triggered from multiple sources, so use the information in
// the `segue` and `sender` parameters to disambiguate between different
// logical paths in your app. For example, if the segue originated from a
// table view, the `sender` parameter would identify the cell that the user
// clicked. You could use that information to set the data on the destination
// view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/prepare(for:sender:)
func (v NSViewController) PrepareForSegueSender(segue INSStoryboardSegue, sender objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("prepareForSegue:sender:"), segue, sender)
}
// Called immediately prior to the performance of a storyboard segue.
//
// identifier: The string that identifies the segue to be performed.
// 
// Using the Interface Builder inspector, provide a unique identifier string
// for each segue in a storyboard. The system provides a segue’s identifier
// to this parameter when it calls this method. The identifier string is used
// to locate the segue inside the storyboard file that contains the view
// controller.
//
// sender: The object that initiated the segue. This object is made available for
// informational purposes during the segue.
//
// # Return Value
// 
// [true] to allow a segue to proceed or [false] to stop it from proceeding.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Override this method to return [false] for cases where you want to prevent
// the performance of a segue. By default, invocation of a segue results in
// the segue being performed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/shouldPerformSegue(withIdentifier:sender:)
func (v NSViewController) ShouldPerformSegueWithIdentifierSender(identifier NSStoryboardSegueIdentifier, sender objectivec.IObject) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("shouldPerformSegueWithIdentifier:sender:"), objc.String(string(identifier)), sender)
	return rv
}

// The object whose value is presented in the receiver’s primary view.
//
// # Discussion
// 
// This property the object you provide to it; it does not it. In another
// words, a view controller has a relationship with its represented object and
// does not own it as an attribute.
// 
// The [RepresentedObject] property is key-value coding and key-value
// observing compliant. When you use the represented object as the file’s
// owner of a nib file, you can bind controls to the file’s owner using key
// paths that start with the string `representedObject`.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/representedObject
func (v NSViewController) RepresentedObject() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("representedObject"))
	return objectivec.Object{ID: rv}
}
func (v NSViewController) SetRepresentedObject(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setRepresentedObject:"), value)
}
// The nib bundle to be loaded to instantiate the receiver’s primary view.
//
// # Discussion
// 
// This property’s value is the bundle you provide to the `nibBundleOrNil`
// parameter in the [InitWithNibNameBundle] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/nibBundle
func (v NSViewController) NibBundle() foundation.NSBundle {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nibBundle"))
	return foundation.NSBundleFromID(objc.ID(rv))
}
// The name of the nib file to be loaded to instantiate the receiver’s
// primary view.
//
// # Discussion
// 
// This property’s value is the name you provide to the `nibNameOrNil`
// parameter in the [InitWithNibNameBundle] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/nibName
func (v NSViewController) NibName() NSNibName {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nibName"))
	return NSNibName(foundation.NSStringFromID(rv).String())
}
// The view controller’s primary view.
//
// # Discussion
// 
// If this property’s value is not already set when you access it, the view
// controller invokes the [LoadView] method. That method, in turn, sets the
// view from the nib file identified by the view controller’s [NibName] and
// [NibBundle] properties.
// 
// If you want to set a view controller’s view directly, set this
// property’s value immediately after creating the view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/view
func (v NSViewController) View() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSViewController) SetView(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setView:"), value)
}
// The localized title of the receiver’s primary view.
//
// # Discussion
// 
// You can employ the [Title] property as needed for your app’s user
// interface, such as to enable a user to choose among multiple named views in
// a menu or other affordance. The [NSViewController] class does not use this
// property for its own purposes.
// 
// The [Title] property is key-value coding and key-value observing compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/title
func (v NSViewController) Title() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (v NSViewController) SetTitle(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The storyboard from which the view controller was loaded.
//
// # Discussion
// 
// If the view controller was not loaded from a storyboard, the value of this
// property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/storyboard
func (v NSViewController) Storyboard() INSStoryboard {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("storyboard"))
	return NSStoryboardFromID(objc.ID(rv))
}
// A Boolean value indicating whether the view controller’s view is loaded
// into memory.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/isViewLoaded
func (v NSViewController) ViewLoaded() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isViewLoaded"))
	return rv
}
// See: https://developer.apple.com/documentation/AppKit/NSViewController/viewIfLoaded
func (v NSViewController) ViewIfLoaded() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("viewIfLoaded"))
	return NSViewFromID(objc.ID(rv))
}
// The desired size of the view controller’s view, in screen units.
//
// # Discussion
// 
// Set this property to express the desired size for a view controller’s
// view. A parent view controller can consult the value of this property when
// performing layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/preferredContentSize
func (v NSViewController) PreferredContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("preferredContentSize"))
	return corefoundation.CGSize(rv)
}
func (v NSViewController) SetPreferredContentSize(value corefoundation.CGSize) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreferredContentSize:"), value)
}
// An array of view controllers that are hierarchical children of the view
// controller.
//
// # Discussion
// 
// You can add or remove child view controllers by using this property. When
// you do, the [AddChildViewController] or [RemoveFromParentViewController]
// method gets called accordingly.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/children
func (v NSViewController) ChildViewControllers() []NSViewController {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("childViewControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSViewController {
		return NSViewControllerFromID(id)
	})
}
func (v NSViewController) SetChildViewControllers(value []NSViewController) {
	objc.Send[struct{}](v.ID, objc.Sel("setChildViewControllers:"), objectivec.IObjectSliceToNSArray(value))
}
// The immediate ancestor view controller of the view controller.
//
// # Discussion
// 
// The value of this property is `nil` if the view controller has no parent
// view controller, such as if the view controller is a window’s content
// view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/parent
func (v NSViewController) ParentViewController() INSViewController {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("parentViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
// The view controllers, if any, that are currently presented by the view
// controller.
//
// # Discussion
// 
// There is a one-to-many relationship between the view controller whose
// [PresentedViewControllers] property you are accessing, and the view
// controllers it is currently presenting.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/presentedViewControllers
func (v NSViewController) PresentedViewControllers() []NSViewController {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("presentedViewControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSViewController {
		return NSViewControllerFromID(id)
	})
}
// The view controller that presented the view controller or that presented
// its farthest ancestor view controller.
//
// # Discussion
// 
// The is the one that is ultimately responsible for presenting the view
// controller whose [PresentingViewController] property you are accessing.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/presentingViewController
func (v NSViewController) PresentingViewController() INSViewController {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("presentingViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
// For a view controller that is part of an app extension, the preferred
// screen origin.
//
// # Discussion
// 
// Set this property to position the lower-left corner of the app
// extension’s primary view in screen space. To specify the desired primary
// view size for an app extension’s view controller, use the
// [PreferredContentSize] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/preferredScreenOrigin
func (v NSViewController) PreferredScreenOrigin() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("preferredScreenOrigin"))
	return corefoundation.CGPoint(rv)
}
func (v NSViewController) SetPreferredScreenOrigin(value corefoundation.CGPoint) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreferredScreenOrigin:"), value)
}
// For a view controller that is part of an app extension, the largest
// allowable size for the app extension’s primary view, in screen units.
//
// # Discussion
// 
// An app extension should return the maximum dimensions that are potentially
// useful for its root view, based on the items the service has been sent. By
// default, the value of this property is a large or infinite size.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/preferredMaximumSize
func (v NSViewController) PreferredMaximumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("preferredMaximumSize"))
	return corefoundation.CGSize(rv)
}
// For a view controller that is part of an app extension, the smallest
// allowable size for the app extension’s primary view, in screen units.
//
// # Discussion
// 
// An app extension should return the minimum dimensions its primary view can
// accommodate, based on the items the app extension has been sent. By
// default, the value of this property is a small but non-empty size.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/preferredMinimumSize
func (v NSViewController) PreferredMinimumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("preferredMinimumSize"))
	return corefoundation.CGSize(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSViewController/sourceItemView
func (v NSViewController) SourceItemView() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sourceItemView"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSViewController) SetSourceItemView(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setSourceItemView:"), value)
}
// A string that identifies the user interface item.
//
// # Discussion
// 
// Identifiers are used during window restoration operations to uniquely
// identify the windows of the application. You can set the value of this
// string programmatically or in Interface Builder. If you create an item in
// Interface Builder and do not set a value for this string, a unique value is
// created for the item when the nib file is loaded. For programmatically
// created views, you typically set this value after creating the item but
// before adding it to a window.
// 
// You should not change the value of a window’s identifier after adding any
// views to the window. For views and controls in a window, the value you
// specify for this string must be unique on a per-window basis.
// 
// The slash (`/`), backslash (`\`), or colon (`:`) characters are reserved
// and must not be used in your custom identifiers. Similarly, Apple reserves
// all identifiers beginning with an underscore (`_`) character. Applications
// and frameworks should use a consistent prefix for their identifiers to
// avoid collisions with other frameworks. For a list of prefixes used by the
// system frameworks, see [OS X Frameworks] in [Mac Technology Overview].
// 
// If you are subclassing a class from one of the system frameworks, do not
// override the accessor methods of this protocol.
//
// [Mac Technology Overview]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/About/About.html#//apple_ref/doc/uid/TP40001067
// [OS X Frameworks]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/SystemFrameworks/SystemFrameworks.html#//apple_ref/doc/uid/TP40001067-CH210
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (v NSViewController) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (v NSViewController) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](v.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}

			// Protocol methods for NSSeguePerforming
			

			// Protocol methods for NSUserInterfaceItemIdentification
			

// TransitionFromViewControllerToViewControllerOptions is a synchronous wrapper around [NSViewController.TransitionFromViewControllerToViewControllerOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v NSViewController) TransitionFromViewControllerToViewControllerOptions(ctx context.Context, fromViewController INSViewController, toViewController INSViewController, options NSViewControllerTransitionOptions) error {
	done := make(chan struct{}, 1)
	v.TransitionFromViewControllerToViewControllerOptionsCompletionHandler(fromViewController, toViewController, options, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

