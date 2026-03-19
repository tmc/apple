// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPathControl] class.
var (
	_NSPathControlClass     NSPathControlClass
	_NSPathControlClassOnce sync.Once
)

func getNSPathControlClass() NSPathControlClass {
	_NSPathControlClassOnce.Do(func() {
		_NSPathControlClass = NSPathControlClass{class: objc.GetClass("NSPathControl")}
	})
	return _NSPathControlClass
}

// GetNSPathControlClass returns the class object for NSPathControl.
func GetNSPathControlClass() NSPathControlClass {
	return getNSPathControlClass()
}

type NSPathControlClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPathControlClass) Alloc() NSPathControl {
	rv := objc.Send[NSPathControl](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A display of a file system path or virtual path information.
//
// # Overview
// 
// The [NSPathControl] class uses [NSPathCell] to implement its user
// interface. [NSPathControl] provides cover methods for most [NSPathCell]
// methods—the cover method simply invokes the corresponding cell method.
// See also [NSPathComponentCell], which represents individual components of
// the path, and two associated protocols: [NSPathCellDelegate] and
// [NSPathControlDelegate].
// 
// [NSPathControl] has three styles represented by the [NSPathControl.Style]
// enumeration constants [NSPathControl.Style.standard],
// [NSPathStyleNavigationBar], and [NSPathControl.Style.popUp]. The
// represented path can be a file system path or any other type of path
// leading through a sequence of nodes or components, as defined by the
// programmer.
// 
// [NSPathControl] automatically supports drag and drop, which can be further
// customized via delegate methods. To accept drag and drop, [NSPathControl]
// calls [RegisterForDraggedTypes] with [NSFilenamesPboardType] and
// [NSURLPboardType]. When the URL value in the [NSPathControl] object changes
// because of an automatic drag and drop operation or the user selecting a new
// path via the open panel, the action is sent. In OS X v10.5 the value
// returned by [clickedPathComponentCell()] is `nil`, in macOS 10.6 and later,
// [clickedPathComponentCell()] returns the clicked cell.
//
// [NSFilenamesPboardType]: https://developer.apple.com/documentation/AppKit/NSFilenamesPboardType
// [NSPathControl.Style.popUp]: https://developer.apple.com/documentation/AppKit/NSPathControl/Style/popUp
// [NSPathControl.Style.standard]: https://developer.apple.com/documentation/AppKit/NSPathControl/Style/standard
// [NSPathControl.Style]: https://developer.apple.com/documentation/AppKit/NSPathControl/Style
// [NSPathStyleNavigationBar]: https://developer.apple.com/documentation/AppKit/NSPathStyle/NSPathStyleNavigationBar
// [NSURLPboardType]: https://developer.apple.com/documentation/AppKit/NSURLPboardType
// [clickedPathComponentCell()]: https://developer.apple.com/documentation/AppKit/NSPathControl/clickedPathComponentCell()
//
// # Setting the Control Style
//
//   - [NSPathControl.PathStyle]: The receiver’s path style.
//   - [NSPathControl.SetPathStyle]
//
// # Setting the Background Color
//
//   - [NSPathControl.BackgroundColor]: The receiver’s background color.
//   - [NSPathControl.SetBackgroundColor]
//
// # Setting the Double-Click Action
//
//   - [NSPathControl.DoubleAction]: The receiver’s double-click action method.
//   - [NSPathControl.SetDoubleAction]
//
// # Setting the Path
//
//   - [NSPathControl.URL]: The path value displayed by the receiver.
//   - [NSPathControl.SetURL]
//
// # Setting the Delegate
//
//   - [NSPathControl.Delegate]: The receiver’s delegate.
//   - [NSPathControl.SetDelegate]
//
// # Setting the Drag Operation Mask
//
//   - [NSPathControl.SetDraggingSourceOperationMaskForLocal]: Configures the drag operation mask.
//
// # Instance Properties
//
//   - [NSPathControl.AllowedTypes]
//   - [NSPathControl.SetAllowedTypes]
//   - [NSPathControl.ClickedPathItem]
//   - [NSPathControl.Editable]
//   - [NSPathControl.SetEditable]
//   - [NSPathControl.PathItems]
//   - [NSPathControl.SetPathItems]
//   - [NSPathControl.PlaceholderAttributedString]
//   - [NSPathControl.SetPlaceholderAttributedString]
//   - [NSPathControl.PlaceholderString]
//   - [NSPathControl.SetPlaceholderString]
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl
type NSPathControl struct {
	NSControl
}

// NSPathControlFromID constructs a [NSPathControl] from an objc.ID.
//
// A display of a file system path or virtual path information.
func NSPathControlFromID(id objc.ID) NSPathControl {
	return NSPathControl{NSControl: NSControlFromID(id)}
}
// NOTE: NSPathControl adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPathControl] class.
//
// # Setting the Control Style
//
//   - [INSPathControl.PathStyle]: The receiver’s path style.
//   - [INSPathControl.SetPathStyle]
//
// # Setting the Background Color
//
//   - [INSPathControl.BackgroundColor]: The receiver’s background color.
//   - [INSPathControl.SetBackgroundColor]
//
// # Setting the Double-Click Action
//
//   - [INSPathControl.DoubleAction]: The receiver’s double-click action method.
//   - [INSPathControl.SetDoubleAction]
//
// # Setting the Path
//
//   - [INSPathControl.URL]: The path value displayed by the receiver.
//   - [INSPathControl.SetURL]
//
// # Setting the Delegate
//
//   - [INSPathControl.Delegate]: The receiver’s delegate.
//   - [INSPathControl.SetDelegate]
//
// # Setting the Drag Operation Mask
//
//   - [INSPathControl.SetDraggingSourceOperationMaskForLocal]: Configures the drag operation mask.
//
// # Instance Properties
//
//   - [INSPathControl.AllowedTypes]
//   - [INSPathControl.SetAllowedTypes]
//   - [INSPathControl.ClickedPathItem]
//   - [INSPathControl.Editable]
//   - [INSPathControl.SetEditable]
//   - [INSPathControl.PathItems]
//   - [INSPathControl.SetPathItems]
//   - [INSPathControl.PlaceholderAttributedString]
//   - [INSPathControl.SetPlaceholderAttributedString]
//   - [INSPathControl.PlaceholderString]
//   - [INSPathControl.SetPlaceholderString]
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl
type INSPathControl interface {
	INSControl

	// Topic: Setting the Control Style

	// The receiver’s path style.
	PathStyle() NSPathStyle
	SetPathStyle(value NSPathStyle)

	// Topic: Setting the Background Color

	// The receiver’s background color.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)

	// Topic: Setting the Double-Click Action

	// The receiver’s double-click action method.
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)

	// Topic: Setting the Path

	// The path value displayed by the receiver.
	URL() foundation.INSURL
	SetURL(value foundation.INSURL)

	// Topic: Setting the Delegate

	// The receiver’s delegate.
	Delegate() NSPathControlDelegate
	SetDelegate(value NSPathControlDelegate)

	// Topic: Setting the Drag Operation Mask

	// Configures the drag operation mask.
	SetDraggingSourceOperationMaskForLocal(mask NSDragOperation, isLocal bool)

	// Topic: Instance Properties

	AllowedTypes() []string
	SetAllowedTypes(value []string)
	ClickedPathItem() INSPathControlItem
	Editable() bool
	SetEditable(value bool)
	PathItems() []NSPathControlItem
	SetPathItems(value []NSPathControlItem)
	PlaceholderAttributedString() foundation.NSAttributedString
	SetPlaceholderAttributedString(value foundation.NSAttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
}

// Init initializes the instance.
func (p NSPathControl) Init() NSPathControl {
	rv := objc.Send[NSPathControl](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPathControl) Autorelease() NSPathControl {
	rv := objc.Send[NSPathControl](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPathControl creates a new NSPathControl instance.
func NewNSPathControl() NSPathControl {
	class := getNSPathControlClass()
	rv := objc.Send[NSPathControl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewPathControlWithCoder(coder foundation.INSCoder) NSPathControl {
	instance := getNSPathControlClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPathControlFromID(rv)
}

// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewPathControlWithFrame(frameRect corefoundation.CGRect) NSPathControl {
	instance := getNSPathControlClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSPathControlFromID(rv)
}

// Configures the drag operation mask.
//
// mask: The types of drag operations allowed.
//
// isLocal: If [true], `mask` applies when the drag destination object is in the same
// application as the receiver; if [false], `mask` applies when the
// destination object is outside the receiver’s application.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method configures the default return value of
// [draggingSourceOperationMaskForLocal:]. By default, this method returns
// [DragOperationEvery] when `isLocal` is [true] and [NSDragOperationNone]
// when `isLocal` is [false].
//
// [NSDragOperationNone]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
// [draggingSourceOperationMaskForLocal:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/draggingSourceOperationMaskForLocal:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/setDraggingSourceOperationMask(_:forLocal:)
func (p NSPathControl) SetDraggingSourceOperationMaskForLocal(mask NSDragOperation, isLocal bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

// The receiver’s path style.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/pathStyle
func (p NSPathControl) PathStyle() NSPathStyle {
	rv := objc.Send[NSPathStyle](p.ID, objc.Sel("pathStyle"))
	return NSPathStyle(rv)
}
func (p NSPathControl) SetPathStyle(value NSPathStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setPathStyle:"), value)
}
// The receiver’s background color.
//
// # Discussion
// 
// By default, the background is set to a light blue color for
// [NSPathStyleStandard] and `nil` for the other styles. You can use `[NSColor
// clearColor]` to make the background transparent.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/backgroundColor
func (p NSPathControl) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (p NSPathControl) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setBackgroundColor:"), value)
}
// The receiver’s double-click action method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/doubleAction
func (p NSPathControl) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](p.ID, objc.Sel("doubleAction"))
	return rv
}
func (p NSPathControl) SetDoubleAction(value objc.SEL) {
	objc.Send[struct{}](p.ID, objc.Sel("setDoubleAction:"), value)
}
// The path value displayed by the receiver.
//
// # Discussion
// 
// When setting, an array of [NSPathComponentCell] objects is automatically
// set based on the path in `url`. If `url` is a file URL (returns [true] from
// [isFileURL]), the images are automatically filled with file icons, if the
// path exists. The URL value itself is stored in the `objectValue` property
// of the cell.
//
// [isFileURL]: https://developer.apple.com/documentation/Foundation/NSURL/isFileURL
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/url
func (p NSPathControl) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p NSPathControl) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}
// The receiver’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/delegate
func (p NSPathControl) Delegate() NSPathControlDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return NSPathControlDelegateObjectFromID(rv)
}
func (p NSPathControl) SetDelegate(value NSPathControlDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/allowedTypes
func (p NSPathControl) AllowedTypes() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("allowedTypes"))
	return objc.ConvertSliceToStrings(rv)
}
func (p NSPathControl) SetAllowedTypes(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowedTypes:"), objectivec.StringSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/clickedPathItem
func (p NSPathControl) ClickedPathItem() INSPathControlItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("clickedPathItem"))
	return NSPathControlItemFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/isEditable
func (p NSPathControl) Editable() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isEditable"))
	return rv
}
func (p NSPathControl) SetEditable(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setEditable:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/pathItems
func (p NSPathControl) PathItems() []NSPathControlItem {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("pathItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPathControlItem {
		return NSPathControlItemFromID(id)
	})
}
func (p NSPathControl) SetPathItems(value []NSPathControlItem) {
	objc.Send[struct{}](p.ID, objc.Sel("setPathItems:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/placeholderAttributedString
func (p NSPathControl) PlaceholderAttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("placeholderAttributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (p NSPathControl) SetPlaceholderAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](p.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSPathControl/placeholderString
func (p NSPathControl) PlaceholderString() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("placeholderString"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPathControl) SetPlaceholderString(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}

