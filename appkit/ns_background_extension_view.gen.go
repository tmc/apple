// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSBackgroundExtensionView] class.
var (
	_NSBackgroundExtensionViewClass     NSBackgroundExtensionViewClass
	_NSBackgroundExtensionViewClassOnce sync.Once
)

func getNSBackgroundExtensionViewClass() NSBackgroundExtensionViewClass {
	_NSBackgroundExtensionViewClassOnce.Do(func() {
		_NSBackgroundExtensionViewClass = NSBackgroundExtensionViewClass{class: objc.GetClass("NSBackgroundExtensionView")}
	})
	return _NSBackgroundExtensionViewClass
}

// GetNSBackgroundExtensionViewClass returns the class object for NSBackgroundExtensionView.
func GetNSBackgroundExtensionViewClass() NSBackgroundExtensionViewClass {
	return getNSBackgroundExtensionViewClass()
}

type NSBackgroundExtensionViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBackgroundExtensionViewClass) Alloc() NSBackgroundExtensionView {
	rv := objc.Send[NSBackgroundExtensionView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A view that extends content to fill its own bounds.
//
// # Overview
// 
// A background extension view can be laid out to extend outside the safe
// area, such as under the titlebar, sidebar, or inspector. By default it lays
// out its content to stay within the safe area, and uses modifications of the
// content along the edges to fill the container view.
//
// # Instance Properties
//
//   - [NSBackgroundExtensionView.AutomaticallyPlacesContentView]: Controls the automatic safe area placement of the `contentView` within the container.
//   - [NSBackgroundExtensionView.SetAutomaticallyPlacesContentView]
//   - [NSBackgroundExtensionView.ContentView]: The content view to extend to fill the [NSBackgroundExtensionView].
//   - [NSBackgroundExtensionView.SetContentView]
//
// See: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView
type NSBackgroundExtensionView struct {
	NSView
}

// NSBackgroundExtensionViewFromID constructs a [NSBackgroundExtensionView] from an objc.ID.
//
// A view that extends content to fill its own bounds.
func NSBackgroundExtensionViewFromID(id objc.ID) NSBackgroundExtensionView {
	return NSBackgroundExtensionView{NSView: NSViewFromID(id)}
}
// NOTE: NSBackgroundExtensionView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSBackgroundExtensionView] class.
//
// # Instance Properties
//
//   - [INSBackgroundExtensionView.AutomaticallyPlacesContentView]: Controls the automatic safe area placement of the `contentView` within the container.
//   - [INSBackgroundExtensionView.SetAutomaticallyPlacesContentView]
//   - [INSBackgroundExtensionView.ContentView]: The content view to extend to fill the [NSBackgroundExtensionView].
//   - [INSBackgroundExtensionView.SetContentView]
//
// See: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView
type INSBackgroundExtensionView interface {
	INSView

	// Topic: Instance Properties

	// Controls the automatic safe area placement of the `contentView` within the container.
	AutomaticallyPlacesContentView() bool
	SetAutomaticallyPlacesContentView(value bool)
	// The content view to extend to fill the [NSBackgroundExtensionView].
	ContentView() INSView
	SetContentView(value INSView)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (b NSBackgroundExtensionView) Init() NSBackgroundExtensionView {
	rv := objc.Send[NSBackgroundExtensionView](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBackgroundExtensionView) Autorelease() NSBackgroundExtensionView {
	rv := objc.Send[NSBackgroundExtensionView](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBackgroundExtensionView creates a new NSBackgroundExtensionView instance.
func NewNSBackgroundExtensionView() NSBackgroundExtensionView {
	class := getNSBackgroundExtensionViewClass()
	rv := objc.Send[NSBackgroundExtensionView](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewBackgroundExtensionViewWithCoder(coder foundation.INSCoder) NSBackgroundExtensionView {
	instance := getNSBackgroundExtensionViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSBackgroundExtensionViewFromID(rv)
}


// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewBackgroundExtensionViewWithFrame(frameRect corefoundation.CGRect) NSBackgroundExtensionView {
	instance := getNSBackgroundExtensionViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSBackgroundExtensionViewFromID(rv)
}






func (b NSBackgroundExtensionView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}












// Controls the automatic safe area placement of the `contentView` within the
// container.
//
// # Discussion
// 
// When [NO], the frame of the content view must be explicitly set or
// constraints added. The extension effect will be used to fill the container
// view around the content.
// 
// Defaults to [YES].
//
// See: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/automaticallyPlacesContentView
func (b NSBackgroundExtensionView) AutomaticallyPlacesContentView() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("automaticallyPlacesContentView"))
	return rv
}
func (b NSBackgroundExtensionView) SetAutomaticallyPlacesContentView(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setAutomaticallyPlacesContentView:"), value)
}



// The content view to extend to fill the [NSBackgroundExtensionView].
//
// # Discussion
// 
// The content view will be added as a subview of the extension view and
// placed within the safe area by default. See
// `automaticallyPlacesContentView` to customize the layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSBackgroundExtensionView/contentView
func (b NSBackgroundExtensionView) ContentView() INSView {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("contentView"))
	return NSViewFromID(objc.ID(rv))
}
func (b NSBackgroundExtensionView) SetContentView(value INSView) {
	objc.Send[struct{}](b.ID, objc.Sel("setContentView:"), value)
}



































