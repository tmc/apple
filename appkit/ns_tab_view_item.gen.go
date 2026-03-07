// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTabViewItem] class.
var (
	_NSTabViewItemClass     NSTabViewItemClass
	_NSTabViewItemClassOnce sync.Once
)

func getNSTabViewItemClass() NSTabViewItemClass {
	_NSTabViewItemClassOnce.Do(func() {
		_NSTabViewItemClass = NSTabViewItemClass{class: objc.GetClass("NSTabViewItem")}
	})
	return _NSTabViewItemClass
}

// GetNSTabViewItemClass returns the class object for NSTabViewItem.
func GetNSTabViewItemClass() NSTabViewItemClass {
	return getNSTabViewItemClass()
}

type NSTabViewItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTabViewItemClass) Alloc() NSTabViewItem {
	rv := objc.Send[NSTabViewItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An item in a tab view.
//
// # Overview
// 
// An [NSTabViewItem] is a convenient way for presenting information in
// multiple pages. A tab view is usually distinguished by a row of tabs that
// give the visual appearance of folder tabs. When the user clicks a tab, the
// tab view displays a view page provided by your application. A tab view
// keeps a zero-based array of tab view items, one for each tab in the view.
//
// # Creating a Tab View Item
//
//   - [NSTabViewItem.InitWithIdentifier]: Performs default initialization for the receiver.
//
// # Working with Labels
//
//   - [NSTabViewItem.DrawLabelInRect]: Draws the receiver’s label in `tabRect`, which is the area between the curved end caps.
//   - [NSTabViewItem.Label]: Sets the label text for the receiver to `label`.
//   - [NSTabViewItem.SetLabel]
//   - [NSTabViewItem.SizeOfLabel]: Calculates the size of the receiver’s label.
//
// # Checking the Tab Display State
//
//   - [NSTabViewItem.TabState]: Returns the current display state of the tab associated with the receiver.
//
// # Assigning an Identifier Object
//
//   - [NSTabViewItem.Identifier]: Sets the receiver’s optional identifier object to `identifier`.
//   - [NSTabViewItem.SetIdentifier]
//
// # Setting the Color
//
//   - [NSTabViewItem.Color]: Sets the background color for content in the view.
//   - [NSTabViewItem.SetColor]
//
// # Assigning a View
//
//   - [NSTabViewItem.View]: Sets the view associated with the receiver to `view`.
//   - [NSTabViewItem.SetView]
//
// # Setting the Initial First Responder
//
//   - [NSTabViewItem.InitialFirstResponder]: Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to `view`.
//   - [NSTabViewItem.SetInitialFirstResponder]
//
// # Accessing the Parent Tab View
//
//   - [NSTabViewItem.TabView]: Returns the parent tab view for the receiver.
//
// # Getting and Setting Tooltips
//
//   - [NSTabViewItem.ToolTip]: Sets the tooltip displayed for the tab view item.
//   - [NSTabViewItem.SetToolTip]
//
// # Instance Properties
//
//   - [NSTabViewItem.Image]
//   - [NSTabViewItem.SetImage]
//   - [NSTabViewItem.ViewController]
//   - [NSTabViewItem.SetViewController]
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem
type NSTabViewItem struct {
	objectivec.Object
}

// NSTabViewItemFromID constructs a [NSTabViewItem] from an objc.ID.
//
// An item in a tab view.
func NSTabViewItemFromID(id objc.ID) NSTabViewItem {
	return NSTabViewItem{objectivec.Object{id}}
}
// NOTE: NSTabViewItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTabViewItem] class.
//
// # Creating a Tab View Item
//
//   - [INSTabViewItem.InitWithIdentifier]: Performs default initialization for the receiver.
//
// # Working with Labels
//
//   - [INSTabViewItem.DrawLabelInRect]: Draws the receiver’s label in `tabRect`, which is the area between the curved end caps.
//   - [INSTabViewItem.Label]: Sets the label text for the receiver to `label`.
//   - [INSTabViewItem.SetLabel]
//   - [INSTabViewItem.SizeOfLabel]: Calculates the size of the receiver’s label.
//
// # Checking the Tab Display State
//
//   - [INSTabViewItem.TabState]: Returns the current display state of the tab associated with the receiver.
//
// # Assigning an Identifier Object
//
//   - [INSTabViewItem.Identifier]: Sets the receiver’s optional identifier object to `identifier`.
//   - [INSTabViewItem.SetIdentifier]
//
// # Setting the Color
//
//   - [INSTabViewItem.Color]: Sets the background color for content in the view.
//   - [INSTabViewItem.SetColor]
//
// # Assigning a View
//
//   - [INSTabViewItem.View]: Sets the view associated with the receiver to `view`.
//   - [INSTabViewItem.SetView]
//
// # Setting the Initial First Responder
//
//   - [INSTabViewItem.InitialFirstResponder]: Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to `view`.
//   - [INSTabViewItem.SetInitialFirstResponder]
//
// # Accessing the Parent Tab View
//
//   - [INSTabViewItem.TabView]: Returns the parent tab view for the receiver.
//
// # Getting and Setting Tooltips
//
//   - [INSTabViewItem.ToolTip]: Sets the tooltip displayed for the tab view item.
//   - [INSTabViewItem.SetToolTip]
//
// # Instance Properties
//
//   - [INSTabViewItem.Image]
//   - [INSTabViewItem.SetImage]
//   - [INSTabViewItem.ViewController]
//   - [INSTabViewItem.SetViewController]
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem
type INSTabViewItem interface {
	objectivec.IObject

	// Topic: Creating a Tab View Item

	// Performs default initialization for the receiver.
	InitWithIdentifier(identifier objectivec.IObject) NSTabViewItem

	// Topic: Working with Labels

	// Draws the receiver’s label in `tabRect`, which is the area between the curved end caps.
	DrawLabelInRect(shouldTruncateLabel bool, labelRect corefoundation.CGRect)
	// Sets the label text for the receiver to `label`.
	Label() string
	SetLabel(value string)
	// Calculates the size of the receiver’s label.
	SizeOfLabel(computeMin bool) corefoundation.CGSize

	// Topic: Checking the Tab Display State

	// Returns the current display state of the tab associated with the receiver.
	TabState() NSTabState

	// Topic: Assigning an Identifier Object

	// Sets the receiver’s optional identifier object to `identifier`.
	Identifier() objectivec.IObject
	SetIdentifier(value objectivec.IObject)

	// Topic: Setting the Color

	// Sets the background color for content in the view.
	Color() INSColor
	SetColor(value INSColor)

	// Topic: Assigning a View

	// Sets the view associated with the receiver to `view`.
	View() INSView
	SetView(value INSView)

	// Topic: Setting the Initial First Responder

	// Sets the initial first responder for the view associated with the receiver (the view that is displayed when a user clicks on the tab) to `view`.
	InitialFirstResponder() INSView
	SetInitialFirstResponder(value INSView)

	// Topic: Accessing the Parent Tab View

	// Returns the parent tab view for the receiver.
	TabView() INSTabView

	// Topic: Getting and Setting Tooltips

	// Sets the tooltip displayed for the tab view item.
	ToolTip() string
	SetToolTip(value string)

	// Topic: Instance Properties

	Image() INSImage
	SetImage(value INSImage)
	ViewController() INSViewController
	SetViewController(value INSViewController)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTabViewItem) Init() NSTabViewItem {
	rv := objc.Send[NSTabViewItem](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTabViewItem) Autorelease() NSTabViewItem {
	rv := objc.Send[NSTabViewItem](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTabViewItem creates a new NSTabViewItem instance.
func NewNSTabViewItem() NSTabViewItem {
	class := getNSTabViewItemClass()
	rv := objc.Send[NSTabViewItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Performs default initialization for the receiver.
//
// # Discussion
// 
// Sets the receiver’s identifier object to `identifier`, if it is not
// `nil`. Use this method when creating tab view items programmatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(identifier:)
func NewTabViewItemWithIdentifier(identifier objectivec.IObject) NSTabViewItem {
	instance := getNSTabViewItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	return NSTabViewItemFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(viewController:)
func NewTabViewItemWithViewController(viewController INSViewController) NSTabViewItem {
	rv := objc.Send[objc.ID](objc.ID(getNSTabViewItemClass().class), objc.Sel("tabViewItemWithViewController:"), viewController)
	return NSTabViewItemFromID(rv)
}







// Performs default initialization for the receiver.
//
// # Discussion
// 
// Sets the receiver’s identifier object to `identifier`, if it is not
// `nil`. Use this method when creating tab view items programmatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/init(identifier:)
func (t NSTabViewItem) InitWithIdentifier(identifier objectivec.IObject) NSTabViewItem {
	rv := objc.Send[NSTabViewItem](t.ID, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Draws the receiver’s label in `tabRect`, which is the area between the
// curved end caps.
//
// # Discussion
// 
// If `shouldTruncateLabel` is [false], draws the full label in the rectangle
// specified by `tabRect`. If `shouldTruncateLabel` is [true], draws the
// truncated label. You can override this method to perform customized label
// drawing. For example, you might want to add an icon to each tab in the
// view.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/drawLabel(_:in:)
func (t NSTabViewItem) DrawLabelInRect(shouldTruncateLabel bool, labelRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawLabel:inRect:"), shouldTruncateLabel, labelRect)
}

// Calculates the size of the receiver’s label.
//
// # Discussion
// 
// If `shouldTruncateLabel` is [false], returns the size of the receiver’s
// full label. If `shouldTruncateLabel` is [true], returns the truncated size.
// If your application does anything to change the size of tab labels, such as
// overriding the [DrawLabelInRect] method to add an icon to each tab, you
// should override [SizeOfLabel] too so the NSTabView knows the correct size
// for the tab label.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/sizeOfLabel(_:)
func (t NSTabViewItem) SizeOfLabel(computeMin bool) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("sizeOfLabel:"), computeMin)
	return corefoundation.CGSize(rv)
}
func (t NSTabViewItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}












// Sets the label text for the receiver to `label`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/label
func (t NSTabViewItem) Label() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTabViewItem) SetLabel(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLabel:"), objc.String(value))
}



// Returns the current display state of the tab associated with the receiver.
//
// # Discussion
// 
// The possible values are [NSSelectedTab], [NSBackgroundTab], or
// [NSPressedTab]. Your application does not directly set the tab state.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabState
func (t NSTabViewItem) TabState() NSTabState {
	rv := objc.Send[NSTabState](t.ID, objc.Sel("tabState"))
	return NSTabState(rv)
}



// Sets the receiver’s optional identifier object to `identifier`.
//
// # Discussion
// 
// To customize how your application works with tabs, you can specify an
// identifier object for each tab view item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/identifier
func (t NSTabViewItem) Identifier() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return objectivec.Object{ID: rv}
}
func (t NSTabViewItem) SetIdentifier(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setIdentifier:"), value)
}



// Sets the background color for content in the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/color
func (t NSTabViewItem) Color() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("color"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTabViewItem) SetColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setColor:"), value)
}



// Sets the view associated with the receiver to `view`.
//
// # Discussion
// 
// This is the view displayed when a user clicks the tab. When you set a new
// view, the old view is released.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/view
func (t NSTabViewItem) View() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSTabViewItem) SetView(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setView:"), value)
}



// Sets the initial first responder for the view associated with the receiver
// (the view that is displayed when a user clicks on the tab) to `view`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/initialFirstResponder
func (t NSTabViewItem) InitialFirstResponder() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initialFirstResponder"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSTabViewItem) SetInitialFirstResponder(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setInitialFirstResponder:"), value)
}



// Returns the parent tab view for the receiver.
//
// # Discussion
// 
// Note that this is the tab view itself, not the view displayed when a user
// clicks the tab.
// 
// A tab view item normally learns about its parent tab view when it is
// inserted into the view’s array of items. The NSTabView methods
// [AddTabViewItem] and [InsertTabViewItemAtIndex] set the tab view for the
// added or inserted item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/tabView
func (t NSTabViewItem) TabView() INSTabView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tabView"))
	return NSTabViewFromID(objc.ID(rv))
}



// Sets the tooltip displayed for the tab view item.
//
// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/toolTip
func (t NSTabViewItem) ToolTip() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("toolTip"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTabViewItem) SetToolTip(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setToolTip:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/image
func (t NSTabViewItem) Image() INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (t NSTabViewItem) SetImage(value INSImage) {
	objc.Send[struct{}](t.ID, objc.Sel("setImage:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSTabViewItem/viewController
func (t NSTabViewItem) ViewController() INSViewController {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("viewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (t NSTabViewItem) SetViewController(value INSViewController) {
	objc.Send[struct{}](t.ID, objc.Sel("setViewController:"), value)
}
























