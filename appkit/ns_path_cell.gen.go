// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [NSPathCell] class.
var (
	_NSPathCellClass     NSPathCellClass
	_NSPathCellClassOnce sync.Once
)

func getNSPathCellClass() NSPathCellClass {
	_NSPathCellClassOnce.Do(func() {
		_NSPathCellClass = NSPathCellClass{class: objc.GetClass("NSPathCell")}
	})
	return _NSPathCellClass
}

// GetNSPathCellClass returns the class object for NSPathCell.
func GetNSPathCellClass() NSPathCellClass {
	return getNSPathCellClass()
}

type NSPathCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPathCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPathCellClass) Alloc() NSPathCell {
	rv := objc.Send[NSPathCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The user interface of a path control object.
//
// # Overview
//
// [NSPathCell] maintains a collection of [NSPathComponentCell] objects that
// represent a particular path to be displayed to the user.
//
// The path shown can be set with the [NSPathCell.ClickedPathComponentCell] method. Doing
// so removes all displayed [NSPathComponentCell] objects and automatically
// fills the control with [NSPathComponentCell] objects set to have the
// appropriate icons, display titles, and [NSURL] values for the particular
// path component they represent. Alternatively, you can fill the control
// manually by setting the cell array or directly modifying existing cells.
//
// Both an action and double-click action can be set for the path control. To
// find out what path component cell was clicked in the action, you can read
// the value of [NSPathCell.ClickedPathComponentCell]. When the style is set to
// [NSPathStylePopUp], the action is still sent, and the
// [NSPathCell.ClickedPathComponentCell] value for the represented menu item is correctly
// set. The [NSPathCell.ClickedPathComponentCell] value is valid only when the action is
// being sent. It is also valid when the keyboard is used to invoke the
// action.
//
// Automatic animated expansion of partially hidden [NSPathComponentCell]
// objects happens if you correctly call [MouseEntered] and [MouseExited] for
// each [NSPathComponentCell] in the [NSPathCell] object. This is not required
// if the [NSPathCell.PathStyle] is set to [NSPathStylePopUp], or if you wish to not have
// the animation.
//
// [NSPathCell] supports several path display styles. [NSPathStyleStandard]
// has a light blue background with arrows indicating the path.
// [NSPathStyleNavigationBar] has more defined arrows (chevrons) and looks a
// little like a segmented button. [NSPathStylePopUp] looks and works like an
// [NSPopUpButton] object to display the full path, or, if the cell is
// editable, select a new path.
//
// If the cell’s [Editable] method returns true (the default), you can drag
// and drop into the cell to change the value. You can constrain what can be
// dropped using UTIs (Uniform Type Identifiers) with [NSPathCell.AllowedTypes] or the
// appropriate delegate methods on [NSPathControl].
//
// If the cell’s [Selectable] method returns true (the default), the
// cell’s contents can automatically be dragged out. The proper UTI,
// filename, and URL are placed on the pasteboard. You can further control or
// limit this by using the appropriate delegate methods on [NSPathControl].
//
// If the cell is editable and has the path style set to [NSPathStylePopUp],
// an additional item in the pop-up menu allows selecting another location. By
// default, an [NSOpenPanel] object is configured based on the allowed types.
// The [NSOpenPanel] object can be customized with a delegate method.
//
// # Setting the control size
//
// When setting the [NSPathCell.ControlSize] property, [NSPathCell] properly respects the
// control size for the [NSPathStyleStandard] and [NSPathStylePopUp] styles.
// When the control size is set, the new size is propagated to subcells. When
// the path style is set to [NSPathStyleNavigationBar], you cannot change the
// control size, and it is always set to [NSSmallControlSize]. Attempting to
// change the control size when the path style is [NSPathStyleNavigationBar]
// causes an assertion. Setting the path style to [NSPathStyleNavigationBar]
// forces the control size to be [NSSmallControlSize].
//
// # Displaying Hidden Components
//
//   - [NSPathCell.MouseEnteredWithFrameInView]: Displays the cell component over which the mouse is hovering.
//   - [NSPathCell.MouseExitedWithFrameInView]: Hides the cell component over which the mouse is hovering.
//
// # Setting the Allowed Types
//
//   - [NSPathCell.AllowedTypes]: Sets the component types allowed in the path when the cell is editable.
//   - [NSPathCell.SetAllowedTypes]
//
// # Setting the Control Style
//
//   - [NSPathCell.PathStyle]: Sets the receiver’s path style.
//   - [NSPathCell.SetPathStyle]
//
// # Setting Cell Appearance
//
//   - [NSPathCell.PlaceholderAttributedString]: Sets the value of the placeholder attributed string.
//   - [NSPathCell.SetPlaceholderAttributedString]
//   - [NSPathCell.PlaceholderString]: Returns the placeholder string.
//   - [NSPathCell.SetPlaceholderString]
//   - [NSPathCell.BackgroundColor]: Returns the current background color of the receiver.
//   - [NSPathCell.SetBackgroundColor]
//
// # Managing Path Components
//
//   - [NSPathCell.RectOfPathComponentCellWithFrameInView]: Returns the current rectangle being displayed for a given path component cell, with respect to a given frame in a given view.
//   - [NSPathCell.PathComponentCellAtPointWithFrameInView]: Returns the cell located at the given point within the given frame of the given view.
//   - [NSPathCell.ClickedPathComponentCell]: Sets the value of the path displayed by the receiver.
//   - [NSPathCell.PathComponentCells]: Sets the array of [NSPathComponentCell] objects currently being displayed.
//   - [NSPathCell.SetPathComponentCells]
//
// # Setting the Double-Click Action
//
//   - [NSPathCell.DoubleAction]: Sets the receiver’s double-click action.
//   - [NSPathCell.SetDoubleAction]
//
// # Setting the Path
//
//   - [NSPathCell.URL]: Returns the path displayed by the receiver.
//   - [NSPathCell.SetURL]
//
// # Setting the Delegate
//
//   - [NSPathCell.Delegate]: Sets the receiver’s delegate.
//   - [NSPathCell.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell
//
// [NSSmallControlSize]: https://developer.apple.com/documentation/AppKit/NSSmallControlSize
type NSPathCell struct {
	NSActionCell
}

// NSPathCellFromID constructs a [NSPathCell] from an objc.ID.
//
// The user interface of a path control object.
func NSPathCellFromID(id objc.ID) NSPathCell {
	return NSPathCell{NSActionCell: NSActionCellFromID(id)}
}

// NOTE: NSPathCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPathCell] class.
//
// # Displaying Hidden Components
//
//   - [INSPathCell.MouseEnteredWithFrameInView]: Displays the cell component over which the mouse is hovering.
//   - [INSPathCell.MouseExitedWithFrameInView]: Hides the cell component over which the mouse is hovering.
//
// # Setting the Allowed Types
//
//   - [INSPathCell.AllowedTypes]: Sets the component types allowed in the path when the cell is editable.
//   - [INSPathCell.SetAllowedTypes]
//
// # Setting the Control Style
//
//   - [INSPathCell.PathStyle]: Sets the receiver’s path style.
//   - [INSPathCell.SetPathStyle]
//
// # Setting Cell Appearance
//
//   - [INSPathCell.PlaceholderAttributedString]: Sets the value of the placeholder attributed string.
//   - [INSPathCell.SetPlaceholderAttributedString]
//   - [INSPathCell.PlaceholderString]: Returns the placeholder string.
//   - [INSPathCell.SetPlaceholderString]
//   - [INSPathCell.BackgroundColor]: Returns the current background color of the receiver.
//   - [INSPathCell.SetBackgroundColor]
//
// # Managing Path Components
//
//   - [INSPathCell.RectOfPathComponentCellWithFrameInView]: Returns the current rectangle being displayed for a given path component cell, with respect to a given frame in a given view.
//   - [INSPathCell.PathComponentCellAtPointWithFrameInView]: Returns the cell located at the given point within the given frame of the given view.
//   - [INSPathCell.ClickedPathComponentCell]: Sets the value of the path displayed by the receiver.
//   - [INSPathCell.PathComponentCells]: Sets the array of [NSPathComponentCell] objects currently being displayed.
//   - [INSPathCell.SetPathComponentCells]
//
// # Setting the Double-Click Action
//
//   - [INSPathCell.DoubleAction]: Sets the receiver’s double-click action.
//   - [INSPathCell.SetDoubleAction]
//
// # Setting the Path
//
//   - [INSPathCell.URL]: Returns the path displayed by the receiver.
//   - [INSPathCell.SetURL]
//
// # Setting the Delegate
//
//   - [INSPathCell.Delegate]: Sets the receiver’s delegate.
//   - [INSPathCell.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell
type INSPathCell interface {
	INSActionCell
	NSMenuItemValidation

	// Topic: Displaying Hidden Components

	// Displays the cell component over which the mouse is hovering.
	MouseEnteredWithFrameInView(event INSEvent, frame corefoundation.CGRect, view INSView)
	// Hides the cell component over which the mouse is hovering.
	MouseExitedWithFrameInView(event INSEvent, frame corefoundation.CGRect, view INSView)

	// Topic: Setting the Allowed Types

	// Sets the component types allowed in the path when the cell is editable.
	AllowedTypes() []string
	SetAllowedTypes(value []string)

	// Topic: Setting the Control Style

	// Sets the receiver’s path style.
	PathStyle() NSPathStyle
	SetPathStyle(value NSPathStyle)

	// Topic: Setting Cell Appearance

	// Sets the value of the placeholder attributed string.
	PlaceholderAttributedString() foundation.NSAttributedString
	SetPlaceholderAttributedString(value foundation.NSAttributedString)
	// Returns the placeholder string.
	PlaceholderString() string
	SetPlaceholderString(value string)
	// Returns the current background color of the receiver.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)

	// Topic: Managing Path Components

	// Returns the current rectangle being displayed for a given path component cell, with respect to a given frame in a given view.
	RectOfPathComponentCellWithFrameInView(cell INSPathComponentCell, frame corefoundation.CGRect, view INSView) corefoundation.CGRect
	// Returns the cell located at the given point within the given frame of the given view.
	PathComponentCellAtPointWithFrameInView(point corefoundation.CGPoint, frame corefoundation.CGRect, view INSView) INSPathComponentCell
	// Sets the value of the path displayed by the receiver.
	ClickedPathComponentCell() INSPathComponentCell
	// Sets the array of [NSPathComponentCell] objects currently being displayed.
	PathComponentCells() []NSPathComponentCell
	SetPathComponentCells(value []NSPathComponentCell)

	// Topic: Setting the Double-Click Action

	// Sets the receiver’s double-click action.
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)

	// Topic: Setting the Path

	// Returns the path displayed by the receiver.
	URL() foundation.INSURL
	SetURL(value foundation.INSURL)

	// Topic: Setting the Delegate

	// Sets the receiver’s delegate.
	Delegate() NSPathCellDelegate
	SetDelegate(value NSPathCellDelegate)

	// A Boolean value indicating whether the cell is editable.
	IsEditable() bool
	SetIsEditable(value bool)
	// A Boolean value indicating whether the cell’s text can be selected.
	IsSelectable() bool
	SetIsSelectable(value bool)
	// Tells the delegate that the user changed the selected directory to the directory located at the specified URL.
	PanelDidChangeToDirectoryURL(sender objectivec.IObject, url foundation.INSURL)
	// [NSSavePanel]: Optional — Sent when the user changes the current type. [NSOpenPanel]: Not sent.
	PanelDidSelectType(sender objectivec.IObject, type_ uniformtypeidentifiers.UTType)
	// [NSSavePanel]: Optional — Sent when the content type popup is displayed and the save panel needs the display name for a type. If `nil` is returned, the save panel will display type’s `localizedDescription`. [NSOpenPanel]: Not sent.
	PanelDisplayNameForType(sender objectivec.IObject, type_ uniformtypeidentifiers.UTType) string
	// Tells the delegate that the user changed the selection in the specified Save panel.
	PanelSelectionDidChange(sender objectivec.IObject)
	// Asks the delegate whether the specified URL should be enabled in the Open panel.
	PanelShouldEnableURL(sender objectivec.IObject, url foundation.INSURL) bool
	// Tells the delegate that the user confirmed a filename choice by clicking Save in a Save panel.
	PanelUserEnteredFilenameConfirmed(sender objectivec.IObject, filename string, okFlag bool) string
	// Asks the delegate to validate the URL for a file that the user selected.
	PanelValidateURLError(sender objectivec.IObject, url foundation.INSURL) (bool, error)
	// Tells the delegate that the Save panel is about to expand or collapse because the user clicked the disclosure triangle that displays or hides the file browser.
	PanelWillExpand(sender objectivec.IObject, expanding bool)
}

// Init initializes the instance.
func (p NSPathCell) Init() NSPathCell {
	rv := objc.Send[NSPathCell](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPathCell) Autorelease() NSPathCell {
	rv := objc.Send[NSPathCell](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPathCell creates a new NSPathCell instance.
func NewNSPathCell() NSPathCell {
	class := getNSPathCellClass()
	rv := objc.Send[NSPathCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSCell] object initialized with the specified image and set to
// have the cell’s default menu.
//
// image: The image to use for the cell. If this parameter is `nil`, no image is set.
//
// # Return Value
//
// An initialized [NSCell] object, or `nil` if the cell could not be
// initialized.
//
// # Discussion
//
// This is one of four designated initializers you must implement when
// subclassing. See [NSCell] for the complete list.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewPathCellImageCell(image INSImage) NSPathCell {
	instance := getNSPathCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSPathCellFromID(rv)
}

// Returns an NSCell object initialized with the specified string and set to
// have the cell’s default menu.
//
// string: The initial string to use for the cell.
//
// # Return Value
//
// An initialized [NSCell] object, or `nil` if the cell could not be
// initialized.
//
// # Discussion
//
// If no field editor (a shared [NSText] object) has been created for all
// [NSCell] objects, one is created.
//
// This is one of four designated initializers you must implement when
// subclassing. See [NSCell] for the complete list.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(textCell:)
func NewPathCellTextCell(string_ string) NSPathCell {
	instance := getNSPathCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSPathCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewPathCellWithCoder(coder foundation.INSCoder) NSPathCell {
	instance := getNSPathCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPathCellFromID(rv)
}

// Displays the cell component over which the mouse is hovering.
//
// event: The mouse-entered event.
//
// frame: The frame in which the cell is located.
//
// view: The view in which the cell is located.
//
// # Discussion
//
// The [NSPathCell] object dynamically animates to display the component that
// the mouse is hovering over using mouse-entered and mouse-exited events. The
// control should call these methods to correctly display the hovered
// component to the user. The control can acquire rectangles to track using
// [RectOfPathComponentCellWithFrameInView].
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/mouseEntered(with:frame:in:)
func (p NSPathCell) MouseEnteredWithFrameInView(event INSEvent, frame corefoundation.CGRect, view INSView) {
	objc.Send[objc.ID](p.ID, objc.Sel("mouseEntered:withFrame:inView:"), event, frame, view)
}

// Hides the cell component over which the mouse is hovering.
//
// event: The mouse-exited event.
//
// frame: The frame in which the cell is located.
//
// view: The view in which the cell is located.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/mouseExited(with:frame:in:)
func (p NSPathCell) MouseExitedWithFrameInView(event INSEvent, frame corefoundation.CGRect, view INSView) {
	objc.Send[objc.ID](p.ID, objc.Sel("mouseExited:withFrame:inView:"), event, frame, view)
}

// Returns the current rectangle being displayed for a given path component
// cell, with respect to a given frame in a given view.
//
// cell: The path component cell.
//
// frame: The frame of the view in which the cell appears.
//
// view: The view in which the cell appears.
//
// # Return Value
//
// The rectangle occupied by the path component cell. [NSZeroRect] is returned
// if `cell` is not found or is not currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/rect(of:withFrame:in:)
func (p NSPathCell) RectOfPathComponentCellWithFrameInView(cell INSPathComponentCell, frame corefoundation.CGRect, view INSView) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("rectOfPathComponentCell:withFrame:inView:"), cell, frame, view)
	return corefoundation.CGRect(rv)
}

// Returns the cell located at the given point within the given frame of the
// given view.
//
// point: The point within the returned cell.
//
// frame: The frame within which the point is located.
//
// view: The view within which the frame is located.
//
// # Return Value
//
// The component cell within which the given point is located, or `nil` if no
// cell exists at that location.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCell(at:withFrame:in:)
func (p NSPathCell) PathComponentCellAtPointWithFrameInView(point corefoundation.CGPoint, frame corefoundation.CGRect, view INSView) INSPathComponentCell {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pathComponentCellAtPoint:withFrame:inView:"), point, frame, view)
	return NSPathComponentCellFromID(rv)
}

// Tells the delegate that the user changed the selected directory to the
// directory located at the specified URL.
//
// sender: The panel whose directory changed.
//
// url: The URL of the new directory, or `nil` if it can’t be represented by an
// [NSURL] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:didChangeToDirectoryURL:)
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
func (p NSPathCell) PanelDidChangeToDirectoryURL(sender objectivec.IObject, url foundation.INSURL) {
	objc.Send[objc.ID](p.ID, objc.Sel("panel:didChangeToDirectoryURL:"), sender, url)
}

// [NSSavePanel]: Optional — Sent when the user changes the current type.
// [NSOpenPanel]: Not sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:didSelect:)
func (p NSPathCell) PanelDidSelectType(sender objectivec.IObject, type_ uniformtypeidentifiers.UTType) {
	objc.Send[objc.ID](p.ID, objc.Sel("panel:didSelectType:"), sender, type_)
}

// [NSSavePanel]: Optional — Sent when the content type popup is displayed
// and the save panel needs the display name for a type. If `nil` is returned,
// the save panel will display type’s `localizedDescription`. [NSOpenPanel]:
// Not sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:displayNameFor:)
func (p NSPathCell) PanelDisplayNameForType(sender objectivec.IObject, type_ uniformtypeidentifiers.UTType) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("panel:displayNameForType:"), sender, type_)
	return foundation.NSStringFromID(rv).String()
}

// Tells the delegate that the user changed the selection in the specified
// Save panel.
//
// sender: The panel whose selection changed.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panelSelectionDidChange(_:)
func (p NSPathCell) PanelSelectionDidChange(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("panelSelectionDidChange:"), sender)
}

// Asks the delegate whether the specified URL should be enabled in the Open
// panel.
//
// sender: The panel that asks whether the URL should be enabled.
//
// url: The URL for you to check.
//
// # Return Value
//
// true if you want the panel to display the item at the specifed URL as
// enabled, or false to display it as disabled.
//
// # Discussion
//
// Save panels do not call this method; they always disable URLs.
// Implementations of this method should be fast to avoid stalling the user
// interface. Use [PanelValidateURLError] instead if processing will take a
// long time.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:shouldEnable:)
func (p NSPathCell) PanelShouldEnableURL(sender objectivec.IObject, url foundation.INSURL) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("panel:shouldEnableURL:"), sender, url)
	return rv
}

// Tells the delegate that the user confirmed a filename choice by clicking
// Save in a Save panel.
//
// sender: The panel that reports the user’s confirmation of a filename choice.
//
// filename: The user’s filename choice.
//
// okFlag: If true, the user clicked the Save button; if false, the user did not.
//
// # Return Value
//
// The filename selected by the user, or `nil` if you want to cancel the save
// operation and leave the Save panel onscreen.
//
// # Discussion
//
// The Save panel calls this method before appending any required filename
// extension information, and before it asks the user whether to replace an
// existing file, if a file with the specified name already exists in the
// given location.
//
// The panel may call this method multiple times as the user types. When it
// does, the `okFlag` parameter is false. When the use confirms their choice,
// the value in the `okFlag` is true. If your delegate does extensive
// validation or puts up alerts, do so only when `okFlag` is true.
//
// In macOS 10.15 and later, you cannot change the filename that the user
// selects. Prior to macOS 10.15, you could sanitize the app’s filename to
// remove undesirable characters or limit its length only if your app wasn’t
// running in a sandbox.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:userEnteredFilename:confirmed:)
func (p NSPathCell) PanelUserEnteredFilenameConfirmed(sender objectivec.IObject, filename string, okFlag bool) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("panel:userEnteredFilename:confirmed:"), sender, objc.String(filename), okFlag)
	return foundation.NSStringFromID(rv).String()
}

// Asks the delegate to validate the URL for a file that the user selected.
//
// sender: The panel that requests URL validation.
//
// url: The URL for you to validate.
//
// # Discussion
//
// Save panels call this method when the user clicks the Save button. Open
// panels call it when the user clicks the Open button. An Open panel calls
// this method once for each selected filename or directory.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:validate:)
func (p NSPathCell) PanelValidateURLError(sender objectivec.IObject, url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("panel:validateURL:error:"), sender, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("panel:validateURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Tells the delegate that the Save panel is about to expand or collapse
// because the user clicked the disclosure triangle that displays or hides the
// file browser.
//
// sender: The panel that is about to expand or collapse.
//
// expanding: true specifies that the panel is expanding; false specifies that it is
// collapsing.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:willExpand:)
func (p NSPathCell) PanelWillExpand(sender objectivec.IObject, expanding bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("panel:willExpand:"), sender, expanding)
}

// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
//
// true to enable `menuItem`, false to disable it.
//
// # Discussion
//
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
//
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (p NSPathCell) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}

// Sets the component types allowed in the path when the cell is editable.
//
// # Discussion
//
// The `allowedTypes` array can contain file extensions (without the period
// that begins the extension) or UTIs. To allow folders, include the
// `public.Folder()` identifier. To allow any type, use `nil`. If the value of
// `allowedTypes` is an empty array, nothing is allowed. The default value is
// `nil`, allowing all types.
//
// If the cell is editable and its type is [NSPathStylePopUp], a Choose item
// is included to enable selection of a different path by invoking an Open
// panel. The allowed types are passed to the Open panel to filter out other
// types. The allowed types are also used with drag and drop to indicate if a
// drop is allowed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/allowedTypes
func (p NSPathCell) AllowedTypes() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("allowedTypes"))
	return objc.ConvertSliceToStrings(rv)
}
func (p NSPathCell) SetAllowedTypes(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowedTypes:"), objectivec.StringSliceToNSArray(value))
}

// Sets the receiver’s path style.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/pathStyle
func (p NSPathCell) PathStyle() NSPathStyle {
	rv := objc.Send[NSPathStyle](p.ID, objc.Sel("pathStyle"))
	return NSPathStyle(rv)
}
func (p NSPathCell) SetPathStyle(value NSPathStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setPathStyle:"), value)
}

// Sets the value of the placeholder attributed string.
//
// # Discussion
//
// If the [NSPathCell] object contains no [NSPathComponentCell] objects, the
// placeholder attributed string is drawn in their place, if it is not `nil`.
// If the placeholder attributed string is `nil`, the (non-attributed)
// placeholder string is drawn with default attributes, if it is not `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/placeholderAttributedString
func (p NSPathCell) PlaceholderAttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("placeholderAttributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (p NSPathCell) SetPlaceholderAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](p.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}

// Returns the placeholder string.
//
// # Return Value
//
// The placeholder string.
//
// # Discussion
//
// If the [NSPathCell] object contains no [NSPathComponentCell] objects, the
// placeholder attributed string is drawn in their place, if it is not `nil`.
// If the placeholder attributed string is `nil`, the (non-attributed)
// placeholder string is drawn with default attributes, if it is not `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/placeholderString
func (p NSPathCell) PlaceholderString() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("placeholderString"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPathCell) SetPlaceholderString(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}

// Returns the current background color of the receiver.
//
// # Return Value
//
// The background color.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/backgroundColor
func (p NSPathCell) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (p NSPathCell) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setBackgroundColor:"), value)
}

// Sets the array of [NSPathComponentCell] objects currently being displayed.
//
// # Discussion
//
// Each item in the array must be an instance of [NSPathComponentCell] or a
// subclass thereof. You cannot set this value to `nil`, but you can set it to
// an empty array using, for example, `[NSArray array]`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCells
func (p NSPathCell) PathComponentCells() []NSPathComponentCell {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("pathComponentCells"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPathComponentCell {
		return NSPathComponentCellFromID(id)
	})
}
func (p NSPathCell) SetPathComponentCells(value []NSPathComponentCell) {
	objc.Send[struct{}](p.ID, objc.Sel("setPathComponentCells:"), objectivec.IObjectSliceToNSArray(value))
}

// Sets the receiver’s double-click action.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/doubleAction
func (p NSPathCell) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](p.ID, objc.Sel("doubleAction"))
	return rv
}
func (p NSPathCell) SetDoubleAction(value objc.SEL) {
	objc.Send[struct{}](p.ID, objc.Sel("setDoubleAction:"), value)
}

// Returns the path displayed by the receiver.
//
// # Return Value
//
// The path value.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/url
func (p NSPathCell) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p NSPathCell) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}

// Sets the value of the path displayed by the receiver.
//
// # Discussion
//
// When setting, an array of [NSPathComponentCell] objects is automatically
// set, based on the path in `url`. The type of [NSPathComponentCell] objects
// created can be controlled by subclassing [NSPathCell] and overriding
// [PathComponentCellClass].
//
// If `url` is a file URL (returns true from [isFileURL]), the images are
// automatically filled with file icons, if the path exists. The URL value
// itself is stored in the `objectValue` property of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/clickedPathComponentCell
//
// [isFileURL]: https://developer.apple.com/documentation/Foundation/NSURL/isFileURL
func (p NSPathCell) ClickedPathComponentCell() INSPathComponentCell {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("clickedPathComponentCell"))
	return NSPathComponentCellFromID(objc.ID(rv))
}

// Sets the receiver’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/delegate
func (p NSPathCell) Delegate() NSPathCellDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return NSPathCellDelegateObjectFromID(rv)
}
func (p NSPathCell) SetDelegate(value NSPathCellDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value indicating whether the cell is editable.
//
// See: https://developer.apple.com/documentation/appkit/nscell/iseditable
func (p NSPathCell) IsEditable() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("editable"))
	return rv
}
func (p NSPathCell) SetIsEditable(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setEditable:"), value)
}

// A Boolean value indicating whether the cell’s text can be selected.
//
// See: https://developer.apple.com/documentation/appkit/nscell/isselectable
func (p NSPathCell) IsSelectable() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("selectable"))
	return rv
}
func (p NSPathCell) SetIsSelectable(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setSelectable:"), value)
}

// Returns the class used to create `pathComponentCell` objects when
// automatically filling up the control.
//
// # Return Value
//
// The class used to create [NSPathComponentCell] objects.
//
// # Discussion
//
// Subclasses can override this method to return a custom cell class that is
// automatically used. By default, the method returns `[NSPathComponentCell
// class]`, or a specialized subclass thereof.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathCell/pathComponentCellClass
func (_NSPathCellClass NSPathCellClass) PathComponentCellClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSPathCellClass.class), objc.Sel("pathComponentCellClass"))
	return rv
}

// Protocol methods for NSMenuItemValidation
