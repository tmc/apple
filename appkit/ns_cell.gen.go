// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCell] class.
var (
	_NSCellClass     NSCellClass
	_NSCellClassOnce sync.Once
)

func getNSCellClass() NSCellClass {
	_NSCellClassOnce.Do(func() {
		_NSCellClass = NSCellClass{class: objc.GetClass("NSCell")}
	})
	return _NSCellClass
}

// GetNSCellClass returns the class object for NSCell.
func GetNSCellClass() NSCellClass {
	return getNSCellClass()
}

type NSCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCellClass) Alloc() NSCell {
	rv := objc.Send[NSCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A mechanism for displaying text or images in a view object without the
// overhead of a full [NSView] subclass.
//
// # Overview
// 
// Cells are used by most of the [NSControl] classes to implement their
// internal workings.
// 
// # Designated Initializers
// 
// When subclassing [NSCell] you must implement all of the designated
// initializers. Those methods include [NSCell.Init], [NSCell.InitWithCoder],
// [NSCell.InitTextCell], and [NSCell.InitImageCell].
//
// # Initializing a Cell
//
//   - [NSCell.InitImageCell]: Returns an [NSCell] object initialized with the specified image and set to have the cell’s default menu.
//   - [NSCell.InitTextCell]: Returns an NSCell object initialized with the specified string and set to have the cell’s default menu.
//
// # Managing Cell Values
//
//   - [NSCell.ObjectValue]: The cell’s value as an Objective-C object.
//   - [NSCell.SetObjectValue]
//   - [NSCell.HasValidObjectValue]: A Boolean value that indicates whether the cell has a valid object value.
//   - [NSCell.IntValue]: The cell’s value as an integer.
//   - [NSCell.SetIntValue]
//   - [NSCell.IntegerValue]: The cell’s value as an integer value.
//   - [NSCell.SetIntegerValue]
//   - [NSCell.StringValue]: The cell’s value as a string.
//   - [NSCell.SetStringValue]
//   - [NSCell.DoubleValue]: The cell’s value as a double-precision floating-point number.
//   - [NSCell.SetDoubleValue]
//   - [NSCell.FloatValue]: The cell’s value as a single-precision floating-point number.
//   - [NSCell.SetFloatValue]
//
// # Managing Cell Attributes
//
//   - [NSCell.SetCellAttributeTo]: Sets the value for the specified cell attribute.
//   - [NSCell.CellAttribute]: Returns the value for the specified cell attribute.
//   - [NSCell.Type]: The type of the cell.
//   - [NSCell.SetType]
//   - [NSCell.Enabled]: A Boolean value indicating whether the cell is currently enabled.
//   - [NSCell.SetEnabled]
//   - [NSCell.AllowsUndo]: A Boolean value indicating whether the cell assumes responsibility for undo operations.
//   - [NSCell.SetAllowsUndo]
//
// # Managing Display Attributes
//
//   - [NSCell.Bezeled]: A Boolean value indicating whether the cell has a bezeled border.
//   - [NSCell.SetBezeled]
//   - [NSCell.Bordered]: A Boolean value indicating whether the cell draws itself outlined with a plain border.
//   - [NSCell.SetBordered]
//   - [NSCell.Opaque]: A Boolean value indicating whether the cell is completely opaque.
//   - [NSCell.BackgroundStyle]: The cell’s background style.
//   - [NSCell.SetBackgroundStyle]
//   - [NSCell.InteriorBackgroundStyle]: The cell’s interior background style.
//
// # Managing Cell State
//
//   - [NSCell.AllowsMixedState]: A Boolean value indicating whether the cell supports three states instead of two.
//   - [NSCell.SetAllowsMixedState]
//   - [NSCell.NextState]: The cell’s next state.
//   - [NSCell.SetNextState]: Changes cell’s state to the next value in the sequence.
//   - [NSCell.State]: The cell’s current state.
//   - [NSCell.SetState]
//
// # Modifying Textual Attributes
//
//   - [NSCell.Editable]: A Boolean value indicating whether the cell is editable.
//   - [NSCell.SetEditable]
//   - [NSCell.Selectable]: A Boolean value indicating whether the cell’s text can be selected.
//   - [NSCell.SetSelectable]
//   - [NSCell.Scrollable]: A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//   - [NSCell.SetScrollable]
//   - [NSCell.Alignment]: The alignment of the cell’s text.
//   - [NSCell.SetAlignment]
//   - [NSCell.Font]: The font that the cell uses to display text.
//   - [NSCell.SetFont]
//   - [NSCell.LineBreakMode]: The line break mode to use when drawing text in the cell.
//   - [NSCell.SetLineBreakMode]
//   - [NSCell.TruncatesLastVisibleLine]: A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
//   - [NSCell.SetTruncatesLastVisibleLine]
//   - [NSCell.Wraps]: A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
//   - [NSCell.SetWraps]
//   - [NSCell.BaseWritingDirection]: The initial writing direction used to determine the actual writing direction for text.
//   - [NSCell.SetBaseWritingDirection]
//   - [NSCell.AttributedStringValue]: The cell’s value as an attributed string.
//   - [NSCell.SetAttributedStringValue]
//   - [NSCell.AllowsEditingTextAttributes]: A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
//   - [NSCell.SetAllowsEditingTextAttributes]
//   - [NSCell.ImportsGraphics]: A Boolean value indicating whether the cell supports the importation of images into its text.
//   - [NSCell.SetImportsGraphics]
//   - [NSCell.SetUpFieldEditorAttributes]: Configures the textual and background attributes of the receiver’s field editor.
//   - [NSCell.Title]: The cell’s title text.
//   - [NSCell.SetTitle]
//
// # Managing the Target and Action
//
//   - [NSCell.Action]: The action performed by the cell.
//   - [NSCell.SetAction]
//   - [NSCell.Target]: The object that receives the cell’s action messages.
//   - [NSCell.SetTarget]
//   - [NSCell.Continuous]: A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//   - [NSCell.SetContinuous]
//   - [NSCell.SendActionOn]: Sets the conditions on which the receiver sends action messages to its target.
//
// # Managing the Image
//
//   - [NSCell.Image]: The image displayed by the cell, if any.
//   - [NSCell.SetImage]
//
// # Managing the Tag
//
//   - [NSCell.Tag]: A tag for identifying the cell.
//   - [NSCell.SetTag]
//
// # Formatting and Validating Data
//
//   - [NSCell.Formatter]: The cell’s formatter object.
//   - [NSCell.SetFormatter]
//
// # Managing Menus
//
//   - [NSCell.Menu]: The cell’s contextual menu.
//   - [NSCell.SetMenu]
//   - [NSCell.MenuForEventInRectOfView]: Returns the menu associated with the cell and related to the specified event and frame.
//
// # Comparing Cells
//
//   - [NSCell.Compare]: Compares the string values of the receiver another cell, disregarding case.
//
// # Respond to Keyboard Events
//
//   - [NSCell.AcceptsFirstResponder]: A Boolean value indicating whether the cell accepts first responder status.
//   - [NSCell.ShowsFirstResponder]: A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
//   - [NSCell.SetShowsFirstResponder]
//   - [NSCell.RefusesFirstResponder]: A Boolean value indicating whether the cell refuses the first responder status.
//   - [NSCell.SetRefusesFirstResponder]
//   - [NSCell.PerformClick]: Simulates a single mouse click on the receiver.
//
// # Deriving Values
//
//   - [NSCell.TakeObjectValueFrom]: Sets the value of the receiver’s cell to the object value obtained from the specified object.
//   - [NSCell.TakeIntegerValueFrom]: Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//   - [NSCell.TakeIntValueFrom]: Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//   - [NSCell.TakeStringValueFrom]: Sets the value of the receiver’s cell to the string value obtained from the specified object.
//   - [NSCell.TakeDoubleValueFrom]: Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//   - [NSCell.TakeFloatValueFrom]: Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// # Representing an Object
//
//   - [NSCell.RepresentedObject]: The object represented by the cell.
//   - [NSCell.SetRepresentedObject]
//
// # Tracking the Mouse
//
//   - [NSCell.TrackMouseInRectOfViewUntilMouseUp]: Initiates the mouse tracking behavior in a cell.
//   - [NSCell.StartTrackingAtInView]: Begins tracking mouse events within the receiver.
//   - [NSCell.ContinueTrackingAtInView]: Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell.
//   - [NSCell.StopTrackingAtInViewMouseIsUp]: Stops tracking mouse events within the receiver.
//   - [NSCell.MouseDownFlags]: The modifier flags for the last (left) mouse-down event.
//   - [NSCell.GetPeriodicDelayInterval]: Returns the initial delay and repeat values for continuous sending of action messages to target objects.
//
// # Hit Testing
//
//   - [NSCell.HitTestForEventInRectOfView]: Returns hit testing information for the receiver.
//
// # Managing the Cursor
//
//   - [NSCell.ResetCursorRectInView]: Sets the receiver to show the I-beam cursor while it tracks the mouse.
//
// # Handling Keyboard Alternatives
//
//   - [NSCell.KeyEquivalent]: The key equivalent associated with clicking the cell.
//
// # Dragging Cells
//
//   - [NSCell.DraggingImageComponentsWithFrameInView]: Generates dragging image components with the specified frame in the view.
//
// # Managing Focus Rings
//
//   - [NSCell.DrawFocusRingMaskWithFrameInView]: Draws the focus ring for the control.
//   - [NSCell.FocusRingMaskBoundsForFrameInView]: Returns the bounds of the focus ring mask.
//   - [NSCell.FocusRingType]: The type of focus ring to use with the associated view.
//   - [NSCell.SetFocusRingType]
//
// # Determining Cell Size
//
//   - [NSCell.CalcDrawInfo]: Recalculates the cell geometry.
//   - [NSCell.CellSize]: The minimum size needed to display the cell.
//   - [NSCell.CellSizeForBounds]: Returns the minimum size needed to display the receiver, constraining it to the specified rectangle.
//   - [NSCell.DrawingRectForBounds]: Returns the rectangle within which the receiver draws itself
//   - [NSCell.ImageRectForBounds]: Returns the rectangle in which the receiver draws its image.
//   - [NSCell.TitleRectForBounds]: Returns the rectangle in which the receiver draws its title text.
//   - [NSCell.ControlSize]: The size of the cell.
//   - [NSCell.SetControlSize]
//
// # Drawing and Highlighting
//
//   - [NSCell.DrawWithFrameInView]: Draws the receiver’s border and then draws the interior of the cell.
//   - [NSCell.HighlightColorWithFrameInView]: Returns the color the receiver uses when drawing the selection highlight.
//   - [NSCell.DrawInteriorWithFrameInView]: Draws the interior portion of the receiver, which includes the image or text portion but does not include the border.
//   - [NSCell.ControlView]: The view associated with the cell.
//   - [NSCell.SetControlView]
//   - [NSCell.HighlightWithFrameInView]: Redraws the receiver with the specified highlight setting.
//   - [NSCell.Highlighted]: A Boolean value indicating whether the cell has a highlighted appearance.
//   - [NSCell.SetHighlighted]
//
// # Editing and Selecting Text
//
//   - [NSCell.EditWithFrameInViewEditorDelegateEvent]: Begins editing of the receiver’s text using the specified field editor.
//   - [NSCell.SelectWithFrameInViewEditorDelegateStartLength]: Selects the specified text range in the cell’s field editor.
//   - [NSCell.SendsActionOnEndEditing]: A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
//   - [NSCell.SetSendsActionOnEndEditing]
//   - [NSCell.EndEditing]: Ends the editing of text in the receiver using the specified field editor.
//   - [NSCell.WantsNotificationForMarkedText]: A Boolean value indicating whether the cell’s field editor should post text change notifications.
//   - [NSCell.FieldEditorForView]: Returns a custom field editor for editing in the view.
//   - [NSCell.UsesSingleLineMode]: A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
//   - [NSCell.SetUsesSingleLineMode]
//
// # Managing Expansion Frames
//
//   - [NSCell.ExpansionFrameWithFrameInView]: Returns the expansion cell frame for the receiver.
//   - [NSCell.DrawWithExpansionFrameInView]: Instructs the receiver to draw in an expansion frame.
//
// # User Interface Layout Direction
//
//   - [NSCell.UserInterfaceLayoutDirection]: The layout direction of the user interface.
//   - [NSCell.SetUserInterfaceLayoutDirection]
//
// # Initializers
//
//   - [NSCell.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSCell
type NSCell struct {
	objectivec.Object
}

// NSCellFromID constructs a [NSCell] from an objc.ID.
//
// A mechanism for displaying text or images in a view object without the
// overhead of a full [NSView] subclass.
func NSCellFromID(id objc.ID) NSCell {
	return NSCell{objectivec.Object{ID: id}}
}
// NOTE: NSCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCell] class.
//
// # Initializing a Cell
//
//   - [INSCell.InitImageCell]: Returns an [NSCell] object initialized with the specified image and set to have the cell’s default menu.
//   - [INSCell.InitTextCell]: Returns an NSCell object initialized with the specified string and set to have the cell’s default menu.
//
// # Managing Cell Values
//
//   - [INSCell.ObjectValue]: The cell’s value as an Objective-C object.
//   - [INSCell.SetObjectValue]
//   - [INSCell.HasValidObjectValue]: A Boolean value that indicates whether the cell has a valid object value.
//   - [INSCell.IntValue]: The cell’s value as an integer.
//   - [INSCell.SetIntValue]
//   - [INSCell.IntegerValue]: The cell’s value as an integer value.
//   - [INSCell.SetIntegerValue]
//   - [INSCell.StringValue]: The cell’s value as a string.
//   - [INSCell.SetStringValue]
//   - [INSCell.DoubleValue]: The cell’s value as a double-precision floating-point number.
//   - [INSCell.SetDoubleValue]
//   - [INSCell.FloatValue]: The cell’s value as a single-precision floating-point number.
//   - [INSCell.SetFloatValue]
//
// # Managing Cell Attributes
//
//   - [INSCell.SetCellAttributeTo]: Sets the value for the specified cell attribute.
//   - [INSCell.CellAttribute]: Returns the value for the specified cell attribute.
//   - [INSCell.Type]: The type of the cell.
//   - [INSCell.SetType]
//   - [INSCell.Enabled]: A Boolean value indicating whether the cell is currently enabled.
//   - [INSCell.SetEnabled]
//   - [INSCell.AllowsUndo]: A Boolean value indicating whether the cell assumes responsibility for undo operations.
//   - [INSCell.SetAllowsUndo]
//
// # Managing Display Attributes
//
//   - [INSCell.Bezeled]: A Boolean value indicating whether the cell has a bezeled border.
//   - [INSCell.SetBezeled]
//   - [INSCell.Bordered]: A Boolean value indicating whether the cell draws itself outlined with a plain border.
//   - [INSCell.SetBordered]
//   - [INSCell.Opaque]: A Boolean value indicating whether the cell is completely opaque.
//   - [INSCell.BackgroundStyle]: The cell’s background style.
//   - [INSCell.SetBackgroundStyle]
//   - [INSCell.InteriorBackgroundStyle]: The cell’s interior background style.
//
// # Managing Cell State
//
//   - [INSCell.AllowsMixedState]: A Boolean value indicating whether the cell supports three states instead of two.
//   - [INSCell.SetAllowsMixedState]
//   - [INSCell.NextState]: The cell’s next state.
//   - [INSCell.SetNextState]: Changes cell’s state to the next value in the sequence.
//   - [INSCell.State]: The cell’s current state.
//   - [INSCell.SetState]
//
// # Modifying Textual Attributes
//
//   - [INSCell.Editable]: A Boolean value indicating whether the cell is editable.
//   - [INSCell.SetEditable]
//   - [INSCell.Selectable]: A Boolean value indicating whether the cell’s text can be selected.
//   - [INSCell.SetSelectable]
//   - [INSCell.Scrollable]: A Boolean value indicating whether excess text scrolls past the cell’s bounds.
//   - [INSCell.SetScrollable]
//   - [INSCell.Alignment]: The alignment of the cell’s text.
//   - [INSCell.SetAlignment]
//   - [INSCell.Font]: The font that the cell uses to display text.
//   - [INSCell.SetFont]
//   - [INSCell.LineBreakMode]: The line break mode to use when drawing text in the cell.
//   - [INSCell.SetLineBreakMode]
//   - [INSCell.TruncatesLastVisibleLine]: A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
//   - [INSCell.SetTruncatesLastVisibleLine]
//   - [INSCell.Wraps]: A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
//   - [INSCell.SetWraps]
//   - [INSCell.BaseWritingDirection]: The initial writing direction used to determine the actual writing direction for text.
//   - [INSCell.SetBaseWritingDirection]
//   - [INSCell.AttributedStringValue]: The cell’s value as an attributed string.
//   - [INSCell.SetAttributedStringValue]
//   - [INSCell.AllowsEditingTextAttributes]: A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
//   - [INSCell.SetAllowsEditingTextAttributes]
//   - [INSCell.ImportsGraphics]: A Boolean value indicating whether the cell supports the importation of images into its text.
//   - [INSCell.SetImportsGraphics]
//   - [INSCell.SetUpFieldEditorAttributes]: Configures the textual and background attributes of the receiver’s field editor.
//   - [INSCell.Title]: The cell’s title text.
//   - [INSCell.SetTitle]
//
// # Managing the Target and Action
//
//   - [INSCell.Action]: The action performed by the cell.
//   - [INSCell.SetAction]
//   - [INSCell.Target]: The object that receives the cell’s action messages.
//   - [INSCell.SetTarget]
//   - [INSCell.Continuous]: A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
//   - [INSCell.SetContinuous]
//   - [INSCell.SendActionOn]: Sets the conditions on which the receiver sends action messages to its target.
//
// # Managing the Image
//
//   - [INSCell.Image]: The image displayed by the cell, if any.
//   - [INSCell.SetImage]
//
// # Managing the Tag
//
//   - [INSCell.Tag]: A tag for identifying the cell.
//   - [INSCell.SetTag]
//
// # Formatting and Validating Data
//
//   - [INSCell.Formatter]: The cell’s formatter object.
//   - [INSCell.SetFormatter]
//
// # Managing Menus
//
//   - [INSCell.Menu]: The cell’s contextual menu.
//   - [INSCell.SetMenu]
//   - [INSCell.MenuForEventInRectOfView]: Returns the menu associated with the cell and related to the specified event and frame.
//
// # Comparing Cells
//
//   - [INSCell.Compare]: Compares the string values of the receiver another cell, disregarding case.
//
// # Respond to Keyboard Events
//
//   - [INSCell.AcceptsFirstResponder]: A Boolean value indicating whether the cell accepts first responder status.
//   - [INSCell.ShowsFirstResponder]: A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
//   - [INSCell.SetShowsFirstResponder]
//   - [INSCell.RefusesFirstResponder]: A Boolean value indicating whether the cell refuses the first responder status.
//   - [INSCell.SetRefusesFirstResponder]
//   - [INSCell.PerformClick]: Simulates a single mouse click on the receiver.
//
// # Deriving Values
//
//   - [INSCell.TakeObjectValueFrom]: Sets the value of the receiver’s cell to the object value obtained from the specified object.
//   - [INSCell.TakeIntegerValueFrom]: Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//   - [INSCell.TakeIntValueFrom]: Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//   - [INSCell.TakeStringValueFrom]: Sets the value of the receiver’s cell to the string value obtained from the specified object.
//   - [INSCell.TakeDoubleValueFrom]: Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//   - [INSCell.TakeFloatValueFrom]: Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//
// # Representing an Object
//
//   - [INSCell.RepresentedObject]: The object represented by the cell.
//   - [INSCell.SetRepresentedObject]
//
// # Tracking the Mouse
//
//   - [INSCell.TrackMouseInRectOfViewUntilMouseUp]: Initiates the mouse tracking behavior in a cell.
//   - [INSCell.StartTrackingAtInView]: Begins tracking mouse events within the receiver.
//   - [INSCell.ContinueTrackingAtInView]: Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell.
//   - [INSCell.StopTrackingAtInViewMouseIsUp]: Stops tracking mouse events within the receiver.
//   - [INSCell.MouseDownFlags]: The modifier flags for the last (left) mouse-down event.
//   - [INSCell.GetPeriodicDelayInterval]: Returns the initial delay and repeat values for continuous sending of action messages to target objects.
//
// # Hit Testing
//
//   - [INSCell.HitTestForEventInRectOfView]: Returns hit testing information for the receiver.
//
// # Managing the Cursor
//
//   - [INSCell.ResetCursorRectInView]: Sets the receiver to show the I-beam cursor while it tracks the mouse.
//
// # Handling Keyboard Alternatives
//
//   - [INSCell.KeyEquivalent]: The key equivalent associated with clicking the cell.
//
// # Dragging Cells
//
//   - [INSCell.DraggingImageComponentsWithFrameInView]: Generates dragging image components with the specified frame in the view.
//
// # Managing Focus Rings
//
//   - [INSCell.DrawFocusRingMaskWithFrameInView]: Draws the focus ring for the control.
//   - [INSCell.FocusRingMaskBoundsForFrameInView]: Returns the bounds of the focus ring mask.
//   - [INSCell.FocusRingType]: The type of focus ring to use with the associated view.
//   - [INSCell.SetFocusRingType]
//
// # Determining Cell Size
//
//   - [INSCell.CalcDrawInfo]: Recalculates the cell geometry.
//   - [INSCell.CellSize]: The minimum size needed to display the cell.
//   - [INSCell.CellSizeForBounds]: Returns the minimum size needed to display the receiver, constraining it to the specified rectangle.
//   - [INSCell.DrawingRectForBounds]: Returns the rectangle within which the receiver draws itself
//   - [INSCell.ImageRectForBounds]: Returns the rectangle in which the receiver draws its image.
//   - [INSCell.TitleRectForBounds]: Returns the rectangle in which the receiver draws its title text.
//   - [INSCell.ControlSize]: The size of the cell.
//   - [INSCell.SetControlSize]
//
// # Drawing and Highlighting
//
//   - [INSCell.DrawWithFrameInView]: Draws the receiver’s border and then draws the interior of the cell.
//   - [INSCell.HighlightColorWithFrameInView]: Returns the color the receiver uses when drawing the selection highlight.
//   - [INSCell.DrawInteriorWithFrameInView]: Draws the interior portion of the receiver, which includes the image or text portion but does not include the border.
//   - [INSCell.ControlView]: The view associated with the cell.
//   - [INSCell.SetControlView]
//   - [INSCell.HighlightWithFrameInView]: Redraws the receiver with the specified highlight setting.
//   - [INSCell.Highlighted]: A Boolean value indicating whether the cell has a highlighted appearance.
//   - [INSCell.SetHighlighted]
//
// # Editing and Selecting Text
//
//   - [INSCell.EditWithFrameInViewEditorDelegateEvent]: Begins editing of the receiver’s text using the specified field editor.
//   - [INSCell.SelectWithFrameInViewEditorDelegateStartLength]: Selects the specified text range in the cell’s field editor.
//   - [INSCell.SendsActionOnEndEditing]: A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
//   - [INSCell.SetSendsActionOnEndEditing]
//   - [INSCell.EndEditing]: Ends the editing of text in the receiver using the specified field editor.
//   - [INSCell.WantsNotificationForMarkedText]: A Boolean value indicating whether the cell’s field editor should post text change notifications.
//   - [INSCell.FieldEditorForView]: Returns a custom field editor for editing in the view.
//   - [INSCell.UsesSingleLineMode]: A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
//   - [INSCell.SetUsesSingleLineMode]
//
// # Managing Expansion Frames
//
//   - [INSCell.ExpansionFrameWithFrameInView]: Returns the expansion cell frame for the receiver.
//   - [INSCell.DrawWithExpansionFrameInView]: Instructs the receiver to draw in an expansion frame.
//
// # User Interface Layout Direction
//
//   - [INSCell.UserInterfaceLayoutDirection]: The layout direction of the user interface.
//   - [INSCell.SetUserInterfaceLayoutDirection]
//
// # Initializers
//
//   - [INSCell.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSCell
type INSCell interface {
	objectivec.IObject
	NSAccessibilityElementProtocol
	NSAccessibilityProtocol
	NSUserInterfaceItemIdentification

	// Topic: Initializing a Cell

	// Returns an [NSCell] object initialized with the specified image and set to have the cell’s default menu.
	InitImageCell(image INSImage) NSCell
	// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu.
	InitTextCell(string_ string) NSCell

	// Topic: Managing Cell Values

	// The cell’s value as an Objective-C object.
	ObjectValue() objectivec.IObject
	SetObjectValue(value objectivec.IObject)
	// A Boolean value that indicates whether the cell has a valid object value.
	HasValidObjectValue() bool
	// The cell’s value as an integer.
	IntValue() int
	SetIntValue(value int)
	// The cell’s value as an integer value.
	IntegerValue() int
	SetIntegerValue(value int)
	// The cell’s value as a string.
	StringValue() string
	SetStringValue(value string)
	// The cell’s value as a double-precision floating-point number.
	DoubleValue() float64
	SetDoubleValue(value float64)
	// The cell’s value as a single-precision floating-point number.
	FloatValue() float32
	SetFloatValue(value float32)

	// Topic: Managing Cell Attributes

	// Sets the value for the specified cell attribute.
	SetCellAttributeTo(parameter NSCellAttribute, value int)
	// Returns the value for the specified cell attribute.
	CellAttribute(parameter NSCellAttribute) int
	// The type of the cell.
	Type() NSCellType
	SetType(value NSCellType)
	// A Boolean value indicating whether the cell is currently enabled.
	Enabled() bool
	SetEnabled(value bool)
	// A Boolean value indicating whether the cell assumes responsibility for undo operations.
	AllowsUndo() bool
	SetAllowsUndo(value bool)

	// Topic: Managing Display Attributes

	// A Boolean value indicating whether the cell has a bezeled border.
	Bezeled() bool
	SetBezeled(value bool)
	// A Boolean value indicating whether the cell draws itself outlined with a plain border.
	Bordered() bool
	SetBordered(value bool)
	// A Boolean value indicating whether the cell is completely opaque.
	Opaque() bool
	// The cell’s background style.
	BackgroundStyle() NSBackgroundStyle
	SetBackgroundStyle(value NSBackgroundStyle)
	// The cell’s interior background style.
	InteriorBackgroundStyle() NSBackgroundStyle

	// Topic: Managing Cell State

	// A Boolean value indicating whether the cell supports three states instead of two.
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	// The cell’s next state.
	NextState() int
	// Changes cell’s state to the next value in the sequence.
	SetNextState()
	// The cell’s current state.
	State() NSControlStateValue
	SetState(value NSControlStateValue)

	// Topic: Modifying Textual Attributes

	// A Boolean value indicating whether the cell is editable.
	Editable() bool
	SetEditable(value bool)
	// A Boolean value indicating whether the cell’s text can be selected.
	Selectable() bool
	SetSelectable(value bool)
	// A Boolean value indicating whether excess text scrolls past the cell’s bounds.
	Scrollable() bool
	SetScrollable(value bool)
	// The alignment of the cell’s text.
	Alignment() NSTextAlignment
	SetAlignment(value NSTextAlignment)
	// The font that the cell uses to display text.
	Font() NSFont
	SetFont(value NSFont)
	// The line break mode to use when drawing text in the cell.
	LineBreakMode() NSLineBreakMode
	SetLineBreakMode(value NSLineBreakMode)
	// A Boolean value indicating whether the cell truncates text that does not fit within the cell’s bounds.
	TruncatesLastVisibleLine() bool
	SetTruncatesLastVisibleLine(value bool)
	// A Boolean value indicating whether the cell wraps text whose length that exceeds the cell’s frame.
	Wraps() bool
	SetWraps(value bool)
	// The initial writing direction used to determine the actual writing direction for text.
	BaseWritingDirection() NSWritingDirection
	SetBaseWritingDirection(value NSWritingDirection)
	// The cell’s value as an attributed string.
	AttributedStringValue() foundation.NSAttributedString
	SetAttributedStringValue(value foundation.NSAttributedString)
	// A Boolean value indicating whether the cell allows the editing of its content’s text attributes by the user.
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	// A Boolean value indicating whether the cell supports the importation of images into its text.
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	// Configures the textual and background attributes of the receiver’s field editor.
	SetUpFieldEditorAttributes(textObj INSText) INSText
	// The cell’s title text.
	Title() string
	SetTitle(value string)

	// Topic: Managing the Target and Action

	// The action performed by the cell.
	Action() objc.SEL
	SetAction(value objc.SEL)
	// The object that receives the cell’s action messages.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	// A Boolean value indicating whether the cell sends its action message continuously during mouse tracking.
	Continuous() bool
	SetContinuous(value bool)
	// Sets the conditions on which the receiver sends action messages to its target.
	SendActionOn(mask NSEventMask) int

	// Topic: Managing the Image

	// The image displayed by the cell, if any.
	Image() INSImage
	SetImage(value INSImage)

	// Topic: Managing the Tag

	// A tag for identifying the cell.
	Tag() int
	SetTag(value int)

	// Topic: Formatting and Validating Data

	// The cell’s formatter object.
	Formatter() foundation.NSFormatter
	SetFormatter(value foundation.NSFormatter)

	// Topic: Managing Menus

	// The cell’s contextual menu.
	Menu() INSMenu
	SetMenu(value INSMenu)
	// Returns the menu associated with the cell and related to the specified event and frame.
	MenuForEventInRectOfView(event INSEvent, cellFrame corefoundation.CGRect, view INSView) INSMenu

	// Topic: Comparing Cells

	// Compares the string values of the receiver another cell, disregarding case.
	Compare(otherCell objectivec.IObject) foundation.NSComparisonResult

	// Topic: Respond to Keyboard Events

	// A Boolean value indicating whether the cell accepts first responder status.
	AcceptsFirstResponder() bool
	// A Boolean value indicating whether the cell provides a visual indication that it is the first responder.
	ShowsFirstResponder() bool
	SetShowsFirstResponder(value bool)
	// A Boolean value indicating whether the cell refuses the first responder status.
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	// Simulates a single mouse click on the receiver.
	PerformClick(sender objectivec.IObject)

	// Topic: Deriving Values

	// Sets the value of the receiver’s cell to the object value obtained from the specified object.
	TakeObjectValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
	TakeIntegerValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
	TakeIntValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to the string value obtained from the specified object.
	TakeStringValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
	TakeDoubleValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
	TakeFloatValueFrom(sender objectivec.IObject)

	// Topic: Representing an Object

	// The object represented by the cell.
	RepresentedObject() objectivec.IObject
	SetRepresentedObject(value objectivec.IObject)

	// Topic: Tracking the Mouse

	// Initiates the mouse tracking behavior in a cell.
	TrackMouseInRectOfViewUntilMouseUp(event INSEvent, cellFrame corefoundation.CGRect, controlView INSView, flag bool) bool
	// Begins tracking mouse events within the receiver.
	StartTrackingAtInView(startPoint corefoundation.CGPoint, controlView INSView) bool
	// Returns a Boolean value that indicates whether mouse tracking should continue in the receiving cell.
	ContinueTrackingAtInView(lastPoint corefoundation.CGPoint, currentPoint corefoundation.CGPoint, controlView INSView) bool
	// Stops tracking mouse events within the receiver.
	StopTrackingAtInViewMouseIsUp(lastPoint corefoundation.CGPoint, stopPoint corefoundation.CGPoint, controlView INSView, flag bool)
	// The modifier flags for the last (left) mouse-down event.
	MouseDownFlags() int
	// Returns the initial delay and repeat values for continuous sending of action messages to target objects.
	GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer)

	// Topic: Hit Testing

	// Returns hit testing information for the receiver.
	HitTestForEventInRectOfView(event INSEvent, cellFrame corefoundation.CGRect, controlView INSView) NSCellHitResult

	// Topic: Managing the Cursor

	// Sets the receiver to show the I-beam cursor while it tracks the mouse.
	ResetCursorRectInView(cellFrame corefoundation.CGRect, controlView INSView)

	// Topic: Handling Keyboard Alternatives

	// The key equivalent associated with clicking the cell.
	KeyEquivalent() string

	// Topic: Dragging Cells

	// Generates dragging image components with the specified frame in the view.
	DraggingImageComponentsWithFrameInView(frame corefoundation.CGRect, view INSView) []NSDraggingImageComponent

	// Topic: Managing Focus Rings

	// Draws the focus ring for the control.
	DrawFocusRingMaskWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// Returns the bounds of the focus ring mask.
	FocusRingMaskBoundsForFrameInView(cellFrame corefoundation.CGRect, controlView INSView) corefoundation.CGRect
	// The type of focus ring to use with the associated view.
	FocusRingType() NSFocusRingType
	SetFocusRingType(value NSFocusRingType)

	// Topic: Determining Cell Size

	// Recalculates the cell geometry.
	CalcDrawInfo(rect corefoundation.CGRect)
	// The minimum size needed to display the cell.
	CellSize() corefoundation.CGSize
	// Returns the minimum size needed to display the receiver, constraining it to the specified rectangle.
	CellSizeForBounds(rect corefoundation.CGRect) corefoundation.CGSize
	// Returns the rectangle within which the receiver draws itself
	DrawingRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect
	// Returns the rectangle in which the receiver draws its image.
	ImageRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect
	// Returns the rectangle in which the receiver draws its title text.
	TitleRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect
	// The size of the cell.
	ControlSize() NSControlSize
	SetControlSize(value NSControlSize)

	// Topic: Drawing and Highlighting

	// Draws the receiver’s border and then draws the interior of the cell.
	DrawWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// Returns the color the receiver uses when drawing the selection highlight.
	HighlightColorWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) INSColor
	// Draws the interior portion of the receiver, which includes the image or text portion but does not include the border.
	DrawInteriorWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// The view associated with the cell.
	ControlView() INSView
	SetControlView(value INSView)
	// Redraws the receiver with the specified highlight setting.
	HighlightWithFrameInView(flag bool, cellFrame corefoundation.CGRect, controlView INSView)
	// A Boolean value indicating whether the cell has a highlighted appearance.
	Highlighted() bool
	SetHighlighted(value bool)

	// Topic: Editing and Selecting Text

	// Begins editing of the receiver’s text using the specified field editor.
	EditWithFrameInViewEditorDelegateEvent(rect corefoundation.CGRect, controlView INSView, textObj INSText, delegate objectivec.IObject, event INSEvent)
	// Selects the specified text range in the cell’s field editor.
	SelectWithFrameInViewEditorDelegateStartLength(rect corefoundation.CGRect, controlView INSView, textObj INSText, delegate objectivec.IObject, selStart int, selLength int)
	// A Boolean value indicating whether the cell’s control object sends its action message when the user finishes editing the cell’s text.
	SendsActionOnEndEditing() bool
	SetSendsActionOnEndEditing(value bool)
	// Ends the editing of text in the receiver using the specified field editor.
	EndEditing(textObj INSText)
	// A Boolean value indicating whether the cell’s field editor should post text change notifications.
	WantsNotificationForMarkedText() bool
	// Returns a custom field editor for editing in the view.
	FieldEditorForView(controlView INSView) INSTextView
	// A Boolean value indicating whether the cell restricts layout and rendering of text to a single line.
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)

	// Topic: Managing Expansion Frames

	// Returns the expansion cell frame for the receiver.
	ExpansionFrameWithFrameInView(cellFrame corefoundation.CGRect, view INSView) corefoundation.CGRect
	// Instructs the receiver to draw in an expansion frame.
	DrawWithExpansionFrameInView(cellFrame corefoundation.CGRect, view INSView)

	// Topic: User Interface Layout Direction

	// The layout direction of the user interface.
	UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) NSCell

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c NSCell) Init() NSCell {
	rv := objc.Send[NSCell](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCell) Autorelease() NSCell {
	rv := objc.Send[NSCell](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCell creates a new NSCell instance.
func NewNSCell() NSCell {
	class := getNSCellClass()
	rv := objc.Send[NSCell](objc.ID(class.class), objc.Sel("new"))
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
func NewCellImageCell(image INSImage) NSCell {
	instance := getNSCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSCellFromID(rv)
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
func NewCellTextCell(string_ string) NSCell {
	instance := getNSCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewCellWithCoder(coder foundation.INSCoder) NSCell {
	instance := getNSCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCellFromID(rv)
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
func (c NSCell) InitImageCell(image INSImage) NSCell {
	rv := objc.Send[NSCell](c.ID, objc.Sel("initImageCell:"), image)
	return rv
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
func (c NSCell) InitTextCell(string_ string) NSCell {
	rv := objc.Send[NSCell](c.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return rv
}

// Sets the value for the specified cell attribute.
//
// parameter: The cell attribute whose value you want to set. Attributes include the
// receiver’s current state and whether it is disabled, editable, or
// highlighted.
//
// value: The new value for the attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/setCellAttribute(_:to:)
func (c NSCell) SetCellAttributeTo(parameter NSCellAttribute, value int) {
	objc.Send[objc.ID](c.ID, objc.Sel("setCellAttribute:to:"), parameter, value)
}

// Returns the value for the specified cell attribute.
//
// parameter: The cell attribute whose value you want to get. Attributes include the
// receiver’s current state and whether it is disabled, editable, or
// highlighted.
//
// # Return Value
// 
// The value for the cell attribute specified by `aParameter`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/cellAttribute(_:)
func (c NSCell) CellAttribute(parameter NSCellAttribute) int {
	rv := objc.Send[int](c.ID, objc.Sel("cellAttribute:"), parameter)
	return rv
}

// Changes cell’s state to the next value in the sequence.
//
// # Discussion
// 
// If a cell has three states, it cycles in this order: on, off, mixed, on,
// off, and so forth. If the cell has only two states, it toggles between
// them.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/setNextState()
func (c NSCell) SetNextState() {
	objc.Send[objc.ID](c.ID, objc.Sel("setNextState"))
}

// Configures the textual and background attributes of the receiver’s field
// editor.
//
// textObj: The field editor to configure.
//
// # Return Value
// 
// The configured field editor.
//
// # Discussion
// 
// If the receiver is disabled, this method sets the text color to dark gray;
// otherwise the method sets it to the default color. If the receiver has a
// bezeled border, this method sets the background to the default color for
// text backgrounds; otherwise, the method sets it to the color of the
// receiver’s [NSControl] object.
// 
// You should not use this method to substitute a new field editor.
// [SetUpFieldEditorAttributes] is intended to modify the attributes of the
// text object (that is, the field editor) passed into it and return that text
// object. If you want to substitute your own field editor, use the
// [FieldEditorForObject] method or the [WindowWillReturnFieldEditorToObject]
// delegate method of [NSWindow].
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/setUpFieldEditorAttributes(_:)
func (c NSCell) SetUpFieldEditorAttributes(textObj INSText) INSText {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("setUpFieldEditorAttributes:"), textObj)
	return NSTextFromID(rv)
}

// Sets the conditions on which the receiver sends action messages to its
// target.
//
// mask: A bit mask containing the conditions for sending the action. The only
// conditions that are actually checked are associated with the
// [NSLeftMouseDownMask], [NSLeftMouseUpMask], [NSLeftMouseDraggedMask], and
// [NSPeriodicMask] bits.
// //
// [NSLeftMouseDownMask]: https://developer.apple.com/documentation/AppKit/NSLeftMouseDownMask
// [NSLeftMouseDraggedMask]: https://developer.apple.com/documentation/AppKit/NSLeftMouseDraggedMask
// [NSLeftMouseUpMask]: https://developer.apple.com/documentation/AppKit/NSLeftMouseUpMask
// [NSPeriodicMask]: https://developer.apple.com/documentation/AppKit/NSPeriodicMask
//
// # Return Value
// 
// A bit mask containing the previous settings. This bit mask uses the same
// values as specified in the `mask` parameter.
//
// # Discussion
// 
// You use this method during mouse tracking when the mouse button changes
// state, the mouse moves, or if the cell is marked to send its action
// continuously while tracking. Because of this, the only bits checked in
// `mask` are [NSLeftMouseDownMask], [NSLeftMouseUpMask],
// [NSLeftMouseDraggedMask], and [NSPeriodicMask], which are declared in the
// [NSEvent] class reference.
// 
// You can use the [Continuous] property to turn on the flag corresponding to
// [NSPeriodicMask] or [NSLeftMouseDraggedMask], whichever is appropriate to
// the given subclass of [NSCell].
//
// [NSLeftMouseDownMask]: https://developer.apple.com/documentation/AppKit/NSLeftMouseDownMask
// [NSLeftMouseDraggedMask]: https://developer.apple.com/documentation/AppKit/NSLeftMouseDraggedMask
// [NSLeftMouseUpMask]: https://developer.apple.com/documentation/AppKit/NSLeftMouseUpMask
// [NSPeriodicMask]: https://developer.apple.com/documentation/AppKit/NSPeriodicMask
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/sendAction(on:)
func (c NSCell) SendActionOn(mask NSEventMask) int {
	rv := objc.Send[int](c.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}

// Returns the menu associated with the cell and related to the specified
// event and frame.
//
// event: The event used to find the menu.
//
// cellFrame: The cell’s rectangle. This rectangle indicates the region containing the
// cursor.
//
// view: The view that manages the cell. This is usually the control object that
// owns the cell.
//
// # Return Value
// 
// The menu associated with the cell and event parameters, or `nil` if no menu
// is set.
//
// # Discussion
// 
// This method is usually invoked by the [NSControl] object (`aView`) managing
// the receiver. The default implementation gets the value of the [Menu]
// property and returns `nil` if no menu has been set. Subclasses can override
// to customize the returned menu according to the event received and the area
// in which the mouse event occurs.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/menu(for:in:of:)
func (c NSCell) MenuForEventInRectOfView(event INSEvent, cellFrame corefoundation.CGRect, view INSView) INSMenu {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("menuForEvent:inRect:ofView:"), event, cellFrame, view)
	return NSMenuFromID(rv)
}

// Compares the string values of the receiver another cell, disregarding case.
//
// otherCell: The cell to compare against the receiver. This parameter must be of type
// [NSCell]; if it is not, this method raises [NSBadComparisonException].
// 
// This value must not be `nil`. If the value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// # Return Value
// 
// [NSOrderedAscending] if the string value of the receiver precedes the
// string value of `otherCell` in lexical ordering, [NSOrderedSame] if the
// string values are equivalent in lexical value, and [NSOrderedDescending]
// string value of the receiver follows the string value of `otherCell` in
// lexical ordering.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/compare(_:)
func (c NSCell) Compare(otherCell objectivec.IObject) foundation.NSComparisonResult {
	rv := objc.Send[foundation.NSComparisonResult](c.ID, objc.Sel("compare:"), otherCell)
	return foundation.NSComparisonResult(rv)
}

// Simulates a single mouse click on the receiver.
//
// sender: The object to use as the sender of the event (if the receiver’s control
// view is not valid). This object must be a subclass of [NSView].
//
// # Discussion
// 
// This method performs the receiver’s action on its target. The receiver
// must be enabled to perform the action. If the receiver’s control view is
// valid, that view is used as the sender; otherwise, the value in `sender` is
// used.
// 
// The receiver of this message must be a cell of type [NSActionCell]. This
// method raises an exception if the action message cannot be successfully
// sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/performClick(_:)
func (c NSCell) PerformClick(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("performClick:"), sender)
}

// Sets the value of the receiver’s cell to the object value obtained from
// the specified object.
//
// sender: The object from which to take the value. This object must support the
// [ObjectValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/takeObjectValueFrom(_:)
func (c NSCell) TakeObjectValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeObjectValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an integer value obtained from
// the specified object.
//
// sender: The object from which to take the value. This object must implement the
// [IntegerValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/takeIntegerValueFrom(_:)
func (c NSCell) TakeIntegerValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an integer value obtained from
// the specified object.
//
// sender: The object from which to take the value. This object must implement the
// [IntValue] property.
//
// # Discussion
// 
// Use the [TakeIntegerValueFrom] method instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/takeIntValueFrom(_:)
func (c NSCell) TakeIntValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeIntValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to the string value obtained from
// the specified object.
//
// sender: The object from which to take the value. This object must implement the
// [StringValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/takeStringValueFrom(_:)
func (c NSCell) TakeStringValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeStringValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to a double-precision
// floating-point value obtained from the specified object.
//
// sender: The object from which to take the value. This object must implement the
// [DoubleValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/takeDoubleValueFrom(_:)
func (c NSCell) TakeDoubleValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to a single-precision
// floating-point value obtained from the specified object.
//
// sender: The object from which to take the value. This object must implement the
// [FloatValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/takeFloatValueFrom(_:)
func (c NSCell) TakeFloatValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeFloatValueFrom:"), sender)
}

// Initiates the mouse tracking behavior in a cell.
//
// event: The event that caused the mouse tracking to occur.
//
// cellFrame: The receiver’s frame rectangle.
//
// controlView: The view containing the receiver. This is usually an [NSControl] object.
//
// flag: If [true], mouse tracking continues until the user releases the mouse
// button. If [false], tracking continues until the cursor leaves the tracking
// rectangle, specified by the `cellFrame` parameter, regardless of the mouse
// button state. See the discussion for more information.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if the mouse tracking conditions are met, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is generally not overridden because the default implementation
// invokes other [NSCell] methods that can be overridden to handle specific
// events in a dragging session. This method’s return value depends on the
// `untilMouseUp` flag. If `untilMouseUp` is set to [true], this method
// returns [true] if the mouse button goes up while the cursor is anywhere;
// [false], otherwise. If `untilMouseUp` is set to [false], this method
// returns [true] if the mouse button goes up while the cursor is within
// `cellFrame`; [false], otherwise.
// 
// This method first invokes [StartTrackingAtInView]. If that method returns
// [true], then as mouse-dragged events are intercepted,
// [ContinueTrackingAtInView] is invoked until either the method returns
// [false] or the mouse is released. Finally, [StopTrackingAtInViewMouseIsUp]
// is invoked if the mouse is released. If `untilMouseUp` is [true], it’s
// invoked when the mouse button goes up while the cursor is anywhere. If
// `untilMouseUp` is [false], it’s invoked when the mouse button goes up
// while the cursor is within `cellFrame`. You usually override one or more of
// these methods to respond to specific mouse events.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/trackMouse(with:in:of:untilMouseUp:)
func (c NSCell) TrackMouseInRectOfViewUntilMouseUp(event INSEvent, cellFrame corefoundation.CGRect, controlView INSView, flag bool) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), event, cellFrame, controlView, flag)
	return rv
}

// Begins tracking mouse events within the receiver.
//
// startPoint: The initial location of the cursor.
//
// controlView: The [NSControl] object managing the receiver.
//
// # Return Value
// 
// [true] if the receiver is set to respond continuously or set to respond
// when the mouse is dragged, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The [NSCell] implementation of [TrackMouseInRectOfViewUntilMouseUp] invokes
// this method when tracking begins. Subclasses can override this method to
// implement special mouse-tracking behavior at the beginning of mouse
// tracking—for example, displaying a special cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/startTracking(at:in:)
func (c NSCell) StartTrackingAtInView(startPoint corefoundation.CGPoint, controlView INSView) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("startTrackingAt:inView:"), startPoint, controlView)
	return rv
}

// Returns a Boolean value that indicates whether mouse tracking should
// continue in the receiving cell.
//
// lastPoint: Contains either the initial location of the cursor when tracking began or
// the previous current point.
//
// currentPoint: The current location of the cursor.
//
// controlView: The [NSControl] object managing the receiver.
//
// # Return Value
// 
// [true] if mouse tracking should continue, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked in [TrackMouseInRectOfViewUntilMouseUp]. The default
// implementation returns [true] if the cell is set to continuously send
// action messages to its target when the mouse button is down or the mouse is
// being dragged. Subclasses can override this method to provide more
// sophisticated tracking behavior.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/continueTracking(last:current:in:)
func (c NSCell) ContinueTrackingAtInView(lastPoint corefoundation.CGPoint, currentPoint corefoundation.CGPoint, controlView INSView) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("continueTracking:at:inView:"), lastPoint, currentPoint, controlView)
	return rv
}

// Stops tracking mouse events within the receiver.
//
// lastPoint: Contains the previous position of the cursor.
//
// stopPoint: The current location of the cursor.
//
// controlView: The [NSControl] object managing the receiver.
//
// flag: If [true], this method was invoked because the user released the mouse
// button; otherwise, if [false], the cursor left the designated tracking
// rectangle.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The default [NSCell] implementation of [TrackMouseInRectOfViewUntilMouseUp]
// invokes this method when the cursor has left the bounds of the receiver or
// the mouse button goes up. The default [NSCell] implementation of this
// method does nothing. Subclasses often override this method to provide
// customized tracking behavior. The following example increments the state of
// a tristate cell when the mouse button is clicked:
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/stopTracking(last:current:in:mouseIsUp:)
func (c NSCell) StopTrackingAtInViewMouseIsUp(lastPoint corefoundation.CGPoint, stopPoint corefoundation.CGPoint, controlView INSView, flag bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, controlView, flag)
}

// Returns the initial delay and repeat values for continuous sending of
// action messages to target objects.
//
// delay: On input, a pointer to a floating-point variable. On output, the variable
// contains the current delay (measured in seconds) before messages are sent.
// This parameter must not be [NULL].
//
// interval: On input, a pointer to a floating point variable. On output, the variable
// contains the interval (measured in seconds) at which messages are sent.
// This parameter must not be [NULL].
//
// # Discussion
// 
// The default implementation returns a delay of `0.2` and an interval of
// `0.025` seconds. Subclasses can override this method to supply their own
// delay and interval values.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/getPeriodicDelay(_:interval:)
func (c NSCell) GetPeriodicDelayInterval(delay unsafe.Pointer, interval unsafe.Pointer) {
	objc.Send[objc.ID](c.ID, objc.Sel("getPeriodicDelay:interval:"), delay, interval)
}

// Returns hit testing information for the receiver.
//
// event: The current event.
//
// cellFrame: The cell’s frame.
//
// controlView: The control object in which the cell is located.
//
// # Return Value
// 
// A constant that specifies the type of area in which the event
// occurred—see [NSCell.HitResult] for values.
//
// [NSCell.HitResult]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult
//
// # Discussion
// 
// You can use a bit-wise mask to look for a specific value when calling this
// method—see [NSCell.HitResult] for values.
// 
// Generally, this method should be overridden by custom [NSCell] subclasses
// to return the correct result. Currently, it is called by some multi-cell
// views, such as [NSTableView].
// 
// By default, [NSCell] looks at the cell type and does the following:
// 
// - [NSImageCellType]: If the image exists and the event point is in the
// image returns [NSCellHitContentArea], otherwise [NSCellHitNone]. -
// [NSTextCellType] (also applies to [NSTextFieldCell]):
// 
// If there is text: If the event point hits in the text, return
// [NSCellHitContentArea]. Additionally, if the cell is enabled return
// [NSCellHitContentArea] `|` [NSCellHitEditableTextArea].
// 
// If there is not text: return [NSCellHitNone].
// 
// - [NSNullCellType] (this is the default that applies to non text or image
// cells who don’t override [HitTestForEventInRectOfView]):
// 
// Return [NSCellHitContentArea] by default;
// 
// If the cell not disabled, and it would track, return [NSCellHitContentArea]
// `|` [NSCellHitTrackableArea].
//
// [NSCell.HitResult]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/hitTest(for:in:of:)
func (c NSCell) HitTestForEventInRectOfView(event INSEvent, cellFrame corefoundation.CGRect, controlView INSView) NSCellHitResult {
	rv := objc.Send[NSCellHitResult](c.ID, objc.Sel("hitTestForEvent:inRect:ofView:"), event, cellFrame, controlView)
	return NSCellHitResult(rv)
}

// Sets the receiver to show the I-beam cursor while it tracks the mouse.
//
// cellFrame: The rectangle in which to display the I-beam cursor.
//
// controlView: The control that manages the cell.
//
// # Discussion
// 
// The receiver must be an enabled and selectable (or editable) text-type
// cell.
// 
// This method is invoked by [ResetCursorRects] and in general you do not need
// to call this method unless you have a custom [NSView] that uses a cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/resetCursorRect(_:in:)
func (c NSCell) ResetCursorRectInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](c.ID, objc.Sel("resetCursorRect:inView:"), cellFrame, controlView)
}

// Generates dragging image components with the specified frame in the view.
//
// frame: The bounding rectangle of the receiver.
//
// view: The view that manages the cell.
//
// # Return Value
// 
// An array of [NSDraggingImageComponent] objects representing the cell.
//
// # Discussion
// 
// The default implementation generates an image from the cell and return two
// components: one for [label] and another for [icon]. This is done by
// capturing the portion from the [TitleRectForBounds] and
// [ImageRectForBounds] methods respectively.
// 
// This method can be subclassed and overridden to provide a custom set of
// [NSDraggingImageComponent] to create the drag image for the cell. This
// method is generally used by NSTableView/NSOutlineView.
//
// [icon]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/ImageComponentKey/icon
// [label]: https://developer.apple.com/documentation/AppKit/NSDraggingItem/ImageComponentKey/label
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/draggingImageComponents(withFrame:in:)
func (c NSCell) DraggingImageComponentsWithFrameInView(frame corefoundation.CGRect, view INSView) []NSDraggingImageComponent {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("draggingImageComponentsWithFrame:inView:"), frame, view)
	return objc.ConvertSlice(rv, func(id objc.ID) NSDraggingImageComponent {
		return NSDraggingImageComponentFromID(id)
	})
}

// Draws the focus ring for the control.
//
// cellFrame: The bounding rectangle of the receiver, or a portion of the bounding
// rectangle.
//
// controlView: The view that manages the cell.
//
// # Discussion
// 
// Implemented by NSCell subclasses to draw the appropriate focus ring for the
// control The default implementation does nothing.
// 
// The Application Kit automatically invokes this method when appropriate, to
// render the view’s focus ring. The view must be eligible to become its
// window’s [FirstResponder], by overriding [AcceptsFirstResponder]
// returning [true].
// 
// The focus ring is rendered using the style specified by the [FocusRingType]
// property, which must not be [NSFocusRingType.none] unless you want to
// suppress drawing of the focus ring. An implementation of `` can assume that
// it is drawing in the view’s bounds, and that the fill and stroke colors
// have been set to an arbitrary fully opaque color. It needs only draw the
// desired focus ring shape or an image or other object whose outline defines
// the focus ring mask.
// 
// This new OS X v10.7 focus ring API should no longer call [set()]() and
// perform it’s own drawing.
//
// [NSFocusRingType.none]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/none
// [set()]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/set()
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/drawFocusRingMask(withFrame:in:)
func (c NSCell) DrawFocusRingMaskWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawFocusRingMaskWithFrame:inView:"), cellFrame, controlView)
}

// Returns the bounds of the focus ring mask.
//
// cellFrame: The bounding rectangle of the receiver, or a portion of the bounding
// rectangle.
//
// controlView: The view that manages the cell.
//
// # Return Value
// 
// Returns a rectangle encompassing the focus ring bounds in the `controlView`
// coordinate space.
//
// # Discussion
// 
// Implemented by [NSCell] subclasses to allow the cell to provide the
// rectangular bounds of the focus ring mask for the cell.
// 
// The default implementation returns an empty value. Subclasses are expected
// to implement this method if they intend to draw a focus ring.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/focusRingMaskBounds(forFrame:in:)
func (c NSCell) FocusRingMaskBoundsForFrameInView(cellFrame corefoundation.CGRect, controlView INSView) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("focusRingMaskBoundsForFrame:inView:"), cellFrame, controlView)
	return corefoundation.CGRect(rv)
}

// Recalculates the cell geometry.
//
// rect: The reference rectangle to use when calculating the cell information.
//
// # Discussion
// 
// Objects (such as controls) that manage [NSCell] objects generally maintain
// a flag that informs them if any of their cells have been modified in such a
// way that the location or size of the cell should be recomputed. If so,
// [calcSize()] method of [NSControl] is automatically invoked prior to the
// display of the cell, and that method invokes the [CalcDrawInfo] method of
// the cell.
// 
// The default implementation of this method does nothing.
//
// [calcSize()]: https://developer.apple.com/documentation/AppKit/NSControl/calcSize()
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/calcDrawInfo(_:)
func (c NSCell) CalcDrawInfo(rect corefoundation.CGRect) {
	objc.Send[objc.ID](c.ID, objc.Sel("calcDrawInfo:"), rect)
}

// Returns the minimum size needed to display the receiver, constraining it to
// the specified rectangle.
//
// rect: The size of the cell, or the size of the `aRect` parameter if the cell is
// not a text or image cell. If the cell is an image cell but no image has
// been set, returns [NSZeroSize].
//
// # Discussion
// 
// This method takes into account of the size of the image or text within a
// certain offset determined by the border type of the cell. If the receiver
// is of text type, the text is resized to fit within `aRect` (as much as
// `aRect` is within the bounds of the cell).
// 
// To support constraint-based layout, when the content of a custom cell
// changes in such a way that the return value of this method would change,
// the cell needs to notify its control of the change by calling the
// control’s [InvalidateIntrinsicContentSizeForCell] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/cellSize(forBounds:)
func (c NSCell) CellSizeForBounds(rect corefoundation.CGRect) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("cellSizeForBounds:"), rect)
	return corefoundation.CGSize(rv)
}

// Returns the rectangle within which the receiver draws itself
//
// rect: The bounding rectangle of the receiver.
//
// # Return Value
// 
// The rectangle in which the receiver draws itself. This rectangle is
// slightly inset from the one in `theRect`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/drawingRect(forBounds:)
func (c NSCell) DrawingRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("drawingRectForBounds:"), rect)
	return corefoundation.CGRect(rv)
}

// Returns the rectangle in which the receiver draws its image.
//
// rect: The bounding rectangle of the receiver.
//
// # Return Value
// 
// The rectangle in which the receiver draws its image. This rectangle is
// slightly offset from the one in `theRect`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/imageRect(forBounds:)
func (c NSCell) ImageRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("imageRectForBounds:"), rect)
	return corefoundation.CGRect(rv)
}

// Returns the rectangle in which the receiver draws its title text.
//
// rect: The bounding rectangle of the receiver.
//
// # Return Value
// 
// The rectangle in which the receiver draws its title text.
//
// # Discussion
// 
// If the receiver is a text-type cell, this method resizes the drawing
// rectangle for the title (`theRect`) inward by a small offset to accommodate
// the cell border. If the receiver is not a text-type cell, the method does
// nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/titleRect(forBounds:)
func (c NSCell) TitleRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("titleRectForBounds:"), rect)
	return corefoundation.CGRect(rv)
}

// Draws the receiver’s border and then draws the interior of the cell.
//
// cellFrame: The bounding rectangle of the receiver.
//
// controlView: The control that manages the cell.
//
// # Discussion
// 
// This method draws the cell in the currently focused view, which can be
// different from the `controlView` passed in. Taking advantage of this
// behavior is not recommended, however.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/draw(withFrame:in:)
func (c NSCell) DrawWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawWithFrame:inView:"), cellFrame, controlView)
}

// Returns the color the receiver uses when drawing the selection highlight.
//
// cellFrame: The bounding rectangle of the receiver.
//
// controlView: The control that manages the cell.
//
// # Return Value
// 
// The color the receiver uses when drawing the selection highlight.
//
// # Discussion
// 
// You should not assume that a cell would necessarily want to draw itself
// with the value returned from [selectedControlColor]. A cell may wish to
// draw with different a selection highlight color depending on such things as
// the key state of its `controlView`.
//
// [selectedControlColor]: https://developer.apple.com/documentation/AppKit/NSColor/selectedControlColor
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/highlightColor(withFrame:in:)
func (c NSCell) HighlightColorWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("highlightColorWithFrame:inView:"), cellFrame, controlView)
	return NSColorFromID(rv)
}

// Draws the interior portion of the receiver, which includes the image or
// text portion but does not include the border.
//
// cellFrame: The bounding rectangle of the receiver, or a portion of the bounding
// rectangle.
//
// controlView: The control that manages the cell.
//
// # Discussion
// 
// Text-type [NSCell] objects display their contents in a rectangle slightly
// inset from `cellFrame` using a global [NSText] object. Image-type [NSCell]
// objects display their contents centered within `cellFrame`. If the proper
// attributes are set, this method also displays the dotted-line rectangle to
// indicate if the control is the first responder and highlights the cell.
// This method is invoked from the [DrawCellInside] method of [NSControl] to
// visually update what the cell displays when its contents change. The
// drawing done by the [NSCell] implementation is minimal and becomes more
// complex in objects such as [NSButtonCell] and [NSSliderCell].
// 
// This method draws the cell in the currently focused view, which can be
// different from the `controlView` passed in. Taking advantage of this is not
// recommended.
// 
// Subclasses often override this method to provide more sophisticated drawing
// of cell contents. Because [DrawWithFrameInView] invokes
// [DrawInteriorWithFrameInView] after it draws the cell’s border, do not
// invoke [DrawWithFrameInView] in your override implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/drawInterior(withFrame:in:)
func (c NSCell) DrawInteriorWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawInteriorWithFrame:inView:"), cellFrame, controlView)
}

// Redraws the receiver with the specified highlight setting.
//
// flag: If [true], the cell is redrawn with a highlight; otherwise, if [false], the
// highlight is removed.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// cellFrame: The bounding rectangle of the receiver.
//
// controlView: The control that manages the cell.
//
// # Discussion
// 
// Note that the [NSCell] highlighting does not appear when highlighted cells
// are printed (although instances of [NSTextFieldCell], [NSButtonCell], and
// others can print themselves highlighted). Generally, you cannot depend on
// highlighting being printed because implementations of this method may
// choose (or not choose) to use transparency.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/highlight(_:withFrame:in:)
func (c NSCell) HighlightWithFrameInView(flag bool, cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](c.ID, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}

// Begins editing of the receiver’s text using the specified field editor.
//
// rect: The bounding rectangle of the cell.
//
// controlView: The control that manages the cell.
//
// textObj: The field editor to use for editing the cell.
//
// delegate: The object to use as a delegate for the field editor (`textObj` parameter).
// This delegate object receives various [NSText] delegation and notification
// methods during the course of editing the cell’s contents.
//
// event: The [NSLeftMouseDown] event that initiated the editing behavior.
//
// # Discussion
// 
// If the receiver isn’t a text-type [NSCell] object, no editing is
// performed. Otherwise, the field editor (`textObj`) is sized to `aRect` and
// its superview is set to `controlView`, so it exactly covers the receiver.
// The field editor is then activated and editing begins. It’s the
// responsibility of the delegate to end editing when responding to
// [TextShouldEndEditing]. Upon ending the editing session, the delegate
// should remove any data from the field editor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/edit(withFrame:in:editor:delegate:event:)
func (c NSCell) EditWithFrameInViewEditorDelegateEvent(rect corefoundation.CGRect, controlView INSView, textObj INSText, delegate objectivec.IObject, event INSEvent) {
	objc.Send[objc.ID](c.ID, objc.Sel("editWithFrame:inView:editor:delegate:event:"), rect, controlView, textObj, delegate, event)
}

// Selects the specified text range in the cell’s field editor.
//
// rect: The bounding rectangle of the cell.
//
// controlView: The control that manages the cell.
//
// textObj: The field editor to use for editing the cell.
//
// delegate: The object to use as a delegate for the field editor (`textObj` parameter).
// This delegate object receives various [NSText] delegation and notification
// methods during the course of editing the cell’s contents.
//
// selStart: The start of the text selection.
//
// selLength: The length of the text range.
//
// # Discussion
// 
// This method is similar to [EditWithFrameInViewEditorDelegateEvent], except
// that it can be invoked in any situation, not only on a mouse-down event.
// This method returns without doing anything if `controlView`, `textObj`, or
// the receiver is `nil`, or if the receiver has no font set for it.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/select(withFrame:in:editor:delegate:start:length:)
func (c NSCell) SelectWithFrameInViewEditorDelegateStartLength(rect corefoundation.CGRect, controlView INSView, textObj INSText, delegate objectivec.IObject, selStart int, selLength int) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectWithFrame:inView:editor:delegate:start:length:"), rect, controlView, textObj, delegate, selStart, selLength)
}

// Ends the editing of text in the receiver using the specified field editor.
//
// textObj: The field editor currently handling the editing of the cell’s content.
//
// # Discussion
// 
// Ends any editing of text that began with a call to
// [EditWithFrameInViewEditorDelegateEvent] or
// [SelectWithFrameInViewEditorDelegateStartLength].
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/endEditing(_:)
func (c NSCell) EndEditing(textObj INSText) {
	objc.Send[objc.ID](c.ID, objc.Sel("endEditing:"), textObj)
}

// Returns a custom field editor for editing in the view.
//
// controlView: The view containing cells that require a custom field editor.
//
// # Return Value
// 
// A custom field editor. The field editor must have [FieldEditor] set to
// [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This is an override point for [NSCell] subclasses designed to use their own
// custom field editors. This message is sent to the selected cell of
// `aControlView` using the [NSWindow] method in [FieldEditorForObject].
// 
// Returning non-`nil` from this method indicates skipping the standard field
// editor querying processes including [WindowWillReturnFieldEditorToObject]
// delegation.
// 
// The default implementation returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/fieldEditor(for:)
func (c NSCell) FieldEditorForView(controlView INSView) INSTextView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fieldEditorForView:"), controlView)
	return NSTextViewFromID(rv)
}

// Returns the expansion cell frame for the receiver.
//
// cellFrame: The frame for the receiver.
//
// view: The view in which the receiver will be drawn.
//
// # Return Value
// 
// The expansion cell frame for the receiver. If the frame is not too small,
// return an empty rect ([NSZeroRect]), and no expansion tool tip view will be
// shown.
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// # Discussion
// 
// This method allows the cell to return an expansion cell frame if
// `cellFrame` is too small for the entire contents in the view. When the
// mouse is hovered over the cell in certain controls, the full cell contents
// are shown in a special floating tool tip view. By default, [NSCell] returns
// [NSZeroRect], while some subclasses (such as [NSTextFieldCell]) will return
// the proper frame when required.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/expansionFrame(withFrame:in:)
func (c NSCell) ExpansionFrameWithFrameInView(cellFrame corefoundation.CGRect, view INSView) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("expansionFrameWithFrame:inView:"), cellFrame, view)
	return corefoundation.CGRect(rv)
}

// Instructs the receiver to draw in an expansion frame.
//
// cellFrame: The frame in which to draw.
//
// view: The view in which to draw. This view may be different from the original
// view that the cell appeared in.
//
// # Discussion
// 
// This method allows the cell to perform custom expansion tool tip drawing.
// By default, [NSCell] simply calls [DrawWithFrameInView].
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/draw(withExpansionFrame:in:)
func (c NSCell) DrawWithExpansionFrameInView(cellFrame corefoundation.CGRect, view INSView) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawWithExpansionFrame:inView:"), cellFrame, view)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func (c NSCell) InitWithCoder(coder foundation.INSCoder) NSCell {
	rv := objc.Send[NSCell](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c NSCell) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The cell’s value as an Objective-C object.
//
// # Discussion
// 
// To be valid object value, the cell must have a formatter capable of
// converting the object to and from its textual representation. The value of
// this property is `nil` if an object has not been assigned to the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/objectValue
func (c NSCell) ObjectValue() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectValue"))
	return objectivec.Object{ID: rv}
}
func (c NSCell) SetObjectValue(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setObjectValue:"), value)
}

// A Boolean value that indicates whether the cell has a valid object value.
//
// # Discussion
// 
// The value of this property is [true] if the cell has a valid object value
// or [false] if it does not. A valid object value is one that the cell’s
// formatter can “understand.” Objects are always assumed to be valid
// unless they are rejected by the formatter. Invalid objects can still be
// accepted by the delegate of the cell’s [NSControl] object (using the
// [ControlDidFailToFormatStringErrorDescription] delegate method).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/hasValidObjectValue
func (c NSCell) HasValidObjectValue() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasValidObjectValue"))
	return rv
}

// The cell’s value as an integer.
//
// # Discussion
// 
// Use the [IntegerValue] property instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/intValue
func (c NSCell) IntValue() int {
	rv := objc.Send[int](c.ID, objc.Sel("intValue"))
	return rv
}
func (c NSCell) SetIntValue(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setIntValue:"), value)
}

// The cell’s value as an integer value.
//
// # Discussion
// 
// This property uses the [ObjectValue] property to access the actual value.
// If the cell is not a text-type cell or its contents are not scannable, the
// value in this property is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/integerValue
func (c NSCell) IntegerValue() int {
	rv := objc.Send[int](c.ID, objc.Sel("integerValue"))
	return rv
}
func (c NSCell) SetIntegerValue(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setIntegerValue:"), value)
}

// The cell’s value as a string.
//
// # Discussion
// 
// This property uses the [ObjectValue] property to access the actual value.
// If no formatter is assigned to the cell or if the formatter cannot
// “translate” a new string appropriately, the cell is flagged as having
// an invalid object. If the cell’s object is not an [NSString] object or
// cannot be converted to one, this property contains an empty string. If the
// cell is not a text-type cell, this method converts it to one before setting
// the object value.
// 
// If you use a class that has an [AttributedStringValue] property, the cell
// gets the string from that property instead of this one.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/stringValue
func (c NSCell) StringValue() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (c NSCell) SetStringValue(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setStringValue:"), objc.String(value))
}

// The cell’s value as a double-precision floating-point number.
//
// # Discussion
// 
// This property uses the [ObjectValue] property to access the actual value.
// If the cell is not a text-type cell or the cell’s value is not scannable,
// this property contains the value `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/doubleValue
func (c NSCell) DoubleValue() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("doubleValue"))
	return rv
}
func (c NSCell) SetDoubleValue(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setDoubleValue:"), value)
}

// The cell’s value as a single-precision floating-point number.
//
// # Discussion
// 
// This property uses the [ObjectValue] property to access the actual value.
// If the cell is not a text-type cell or the cell’s value is not scannable,
// this property contains the value `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/floatValue
func (c NSCell) FloatValue() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("floatValue"))
	return rv
}
func (c NSCell) SetFloatValue(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setFloatValue:"), value)
}

// The type of the cell.
//
// # Discussion
// 
// When you change the cell type to [NSCell.CellType.textCellType], the cell
// converts gives itself a default title and sets the font to the system font
// at the default size. When you change the cell type to
// [NSCell.CellType.imageCellType], the cell type does not change until you
// assign a new non-`nil` image to it.
// 
// For a list of possible cell types, see [NSCell.CellType].
//
// [NSCell.CellType.imageCellType]: https://developer.apple.com/documentation/AppKit/NSCell/CellType/imageCellType
// [NSCell.CellType.textCellType]: https://developer.apple.com/documentation/AppKit/NSCell/CellType/textCellType
// [NSCell.CellType]: https://developer.apple.com/documentation/AppKit/NSCell/CellType
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/type
func (c NSCell) Type() NSCellType {
	rv := objc.Send[NSCellType](c.ID, objc.Sel("type"))
	return NSCellType(rv)
}
func (c NSCell) SetType(value NSCellType) {
	objc.Send[struct{}](c.ID, objc.Sel("setType:"), value)
}

// A Boolean value indicating whether the cell is currently enabled.
//
// # Discussion
// 
// The value of this property is [true] when the cell is enabled or [false]
// when it is disabled. The text of disabled cells is gray. If a cell is
// disabled, it cannot be highlighted, does not support mouse tracking (and
// thus cannot participate in target/action functionality), and cannot be
// edited. However, you can still alter many attributes of a disabled cell
// programmatically. (The [State] property, for instance, still works.)
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isEnabled
func (c NSCell) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c NSCell) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}

// A Boolean value indicating whether the cell assumes responsibility for undo
// operations.
//
// # Discussion
// 
// The value of this property is [true] if the cell handles undo operations
// itself or [false] if the app’s custom undo manager must handle undo
// behavior. Cell subclasses set the value of this property to indicate their
// preference for handling undo operations. For example, the [NSTextFieldCell]
// class uses sets this property to indicate it handles undo operations for
// edited text, and other controls set a value that is appropriate for their
// implementation. Do not change the value of this property otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/allowsUndo
func (c NSCell) AllowsUndo() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsUndo"))
	return rv
}
func (c NSCell) SetAllowsUndo(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsUndo:"), value)
}

// A Boolean value indicating whether the cell has a bezeled border.
//
// # Discussion
// 
// The value of this property is [true] if the cell has a bezeled border or
// [false] if it does not. This property is mutually exclusive with the
// [Bordered] property—that is, a cell’s border can be plain or bezeled
// but not both. Changing the value of this property automatically removes any
// border that has been set, regardless of the value you assign to the
// property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isBezeled
func (c NSCell) Bezeled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBezeled"))
	return rv
}
func (c NSCell) SetBezeled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setBezeled:"), value)
}

// A Boolean value indicating whether the cell draws itself outlined with a
// plain border.
//
// # Discussion
// 
// The value of this property is [true] if the cell has a plain border or
// [false] if it does not. This property is mutually exclusive with the
// [Bezeled] property—that is, a cell’s border can be plain or bezeled but
// not both. Changing the value of this property automatically removes any
// bezel that has been set, regardless of the value you assign to the
// property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isBordered
func (c NSCell) Bordered() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBordered"))
	return rv
}
func (c NSCell) SetBordered(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setBordered:"), value)
}

// A Boolean value indicating whether the cell is completely opaque.
//
// # Discussion
// 
// The value of this property is [true] if the cell is completely opaque or
// [false] if it contains some transparency.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isOpaque
func (c NSCell) Opaque() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isOpaque"))
	return rv
}

// The cell’s background style.
//
// # Discussion
// 
// The background describes the surface the cell is drawn onto in the
// [DrawWithFrameInView] method. A control typically sets the value of this
// property before it asks the cell to draw. A cell may draw differently based
// on background characteristics. For example, a table view drawing a cell in
// a selected row might set the value to [dark]. A text cell might decide to
// render its text white as a result. A rating-style level indicator might
// draw its stars white instead of gray.
//
// [dark]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle/dark
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/backgroundStyle
func (c NSCell) BackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](c.ID, objc.Sel("backgroundStyle"))
	return NSBackgroundStyle(rv)
}
func (c NSCell) SetBackgroundStyle(value NSBackgroundStyle) {
	objc.Send[struct{}](c.ID, objc.Sel("setBackgroundStyle:"), value)
}

// The cell’s interior background style.
//
// # Discussion
// 
// The interior background style describes the surface drawn onto in the
// [DrawInteriorWithFrameInView] method. This is often the same as the
// [BackgroundStyle], but a button that draws a bezel would use a different
// value for this property.
// 
// In a custom button with a custom bezel you can override this property and
// return a different value to describe that surface. A cell that has custom
// interior drawing might use the value of this property to pick an image that
// looks good on the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/interiorBackgroundStyle
func (c NSCell) InteriorBackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](c.ID, objc.Sel("interiorBackgroundStyle"))
	return NSBackgroundStyle(rv)
}

// A Boolean value indicating whether the cell supports three states instead
// of two.
//
// # Discussion
// 
// When the value of this property is [true], the cell supports three states:
// on, off, and mixed. When the value is [false], the cell supports only the
// on and off states.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/allowsMixedState
func (c NSCell) AllowsMixedState() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsMixedState"))
	return rv
}
func (c NSCell) SetAllowsMixedState(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsMixedState:"), value)
}

// The cell’s next state.
//
// # Discussion
// 
// If a cell has three states, it cycles in this order: on, off, mixed, on,
// off, and so forth. If the cell has only two states, it toggles between
// them.
// 
// For a list of constants representing the possible cell states, see
// [NSCell.StateValue].
//
// [NSCell.StateValue]: https://developer.apple.com/documentation/AppKit/NSCell/StateValue
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/nextState
func (c NSCell) NextState() int {
	rv := objc.Send[int](c.ID, objc.Sel("nextState"))
	return rv
}

// The cell’s current state.
//
// # Discussion
// 
// The [NSOffState] state indicates the normal or unpressed state. The
// [NSOnState] state indicates the alternate or pressed state. The
// [NSMixedState] state indicates that the feature represented by the control
// is in effect somewhere.
// 
// Although using the enumerated constants is preferred, you can also assign
// an integer value to this property. If the cell has two states, `0` is
// treated as [NSOffState], and a nonzero value is treated as [NSOnState]. If
// the cell has three states, `0` is treated as [NSOffState], a negative value
// is treated as [NSMixedState], and a positive value is treated as
// [NSOnState].
//
// [NSMixedState]: https://developer.apple.com/documentation/AppKit/NSMixedState
// [NSOffState]: https://developer.apple.com/documentation/AppKit/NSOffState
// [NSOnState]: https://developer.apple.com/documentation/AppKit/NSOnState
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/state
func (c NSCell) State() NSControlStateValue {
	rv := objc.Send[NSControlStateValue](c.ID, objc.Sel("state"))
	return NSControlStateValue(rv)
}
func (c NSCell) SetState(value NSControlStateValue) {
	objc.Send[struct{}](c.ID, objc.Sel("setState:"), value)
}

// A Boolean value indicating whether the cell is editable.
//
// # Discussion
// 
// When the value of this property is [true], the cell is editable. When the
// value is [true], the text of the cell is also made selectable. When you
// change the value of the property to [false], the [Selectable] property is
// restored to the value it had before the cell was made editable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isEditable
func (c NSCell) Editable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEditable"))
	return rv
}
func (c NSCell) SetEditable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEditable:"), value)
}

// A Boolean value indicating whether the cell’s text can be selected.
//
// # Discussion
// 
// When the value of this property is [true], the cell’s text is selectable.
// Setting the value of this property to [false] also sets the [Editable]
// property to [false]. If the value of this cell is [true], the value in the
// [Editable] property is not affected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isSelectable
func (c NSCell) Selectable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSelectable"))
	return rv
}
func (c NSCell) SetSelectable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSelectable:"), value)
}

// A Boolean value indicating whether excess text scrolls past the cell’s
// bounds.
//
// # Discussion
// 
// When the value of this property is [true], text can be scrolled past the
// cell’s bound. When the value is [false], the cell wraps its text.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isScrollable
func (c NSCell) Scrollable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isScrollable"))
	return rv
}
func (c NSCell) SetScrollable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setScrollable:"), value)
}

// The alignment of the cell’s text.
//
// # Discussion
// 
// The default value of this property is [NSNaturalTextAlignment]. For a list
// of possible text alignments, see [NSTextAlignment].
//
// [NSTextAlignment]: https://developer.apple.com/documentation/AppKit/NSTextAlignment
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/alignment
func (c NSCell) Alignment() NSTextAlignment {
	rv := objc.Send[NSTextAlignment](c.ID, objc.Sel("alignment"))
	return NSTextAlignment(rv)
}
func (c NSCell) SetAlignment(value NSTextAlignment) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlignment:"), value)
}

// The font that the cell uses to display text.
//
// # Discussion
// 
// The value of this property is `nil` if the cell is not a text-type cell.
// Assigning a new font object to this property converts the cell to a
// text-type cell if it is not already one.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/font
func (c NSCell) Font() NSFont {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("font"))
	return NSFontFromID(objc.ID(rv))
}
func (c NSCell) SetFont(value NSFont) {
	objc.Send[struct{}](c.ID, objc.Sel("setFont:"), value)
}

// The line break mode to use when drawing text in the cell.
//
// # Discussion
// 
// The value in this property can also be modified when you change the value
// of the [Wraps] property. For a list of supported line break modes, see
// NSLineBreakMode.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/lineBreakMode
func (c NSCell) LineBreakMode() NSLineBreakMode {
	rv := objc.Send[NSLineBreakMode](c.ID, objc.Sel("lineBreakMode"))
	return NSLineBreakMode(rv)
}
func (c NSCell) SetLineBreakMode(value NSLineBreakMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setLineBreakMode:"), value)
}

// A Boolean value indicating whether the cell truncates text that does not
// fit within the cell’s bounds.
//
// # Discussion
// 
// When the value of this property is [true], the cell truncates text and adds
// an ellipsis character to the last visible line when the text does not fit.
// The value in the [LineBreakMode] property must be
// [NSLineBreakMode.byWordWrapping] or [NSLineBreakMode.byCharWrapping] for
// this option to have any effect.
//
// [NSLineBreakMode.byCharWrapping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byCharWrapping
// [NSLineBreakMode.byWordWrapping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byWordWrapping
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/truncatesLastVisibleLine
func (c NSCell) TruncatesLastVisibleLine() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("truncatesLastVisibleLine"))
	return rv
}
func (c NSCell) SetTruncatesLastVisibleLine(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setTruncatesLastVisibleLine:"), value)
}

// A Boolean value indicating whether the cell wraps text whose length that
// exceeds the cell’s frame.
//
// # Discussion
// 
// When the value of this property is [true], the cell wraps text and makes
// the cell non-scrollable. If the text of the cell is an attributed string
// value, you must explicitly set the paragraph style line break mode. Setting
// the value of this property to [true] is equivalent to setting the
// [LineBreakMode] property to [NSLineBreakMode.byWordWrapping].
//
// [NSLineBreakMode.byWordWrapping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byWordWrapping
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/wraps
func (c NSCell) Wraps() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("wraps"))
	return rv
}
func (c NSCell) SetWraps(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setWraps:"), value)
}

// The initial writing direction used to determine the actual writing
// direction for text.
//
// # Discussion
// 
// The default value of this property is [NSWritingDirection.natural]. The
// Text system uses this value as a hint for calculating the actual direction
// for displaying Unicode characters. If you know the base writing direction
// of the text you are rendering, you can set the value of this property to
// the correct direction to help the text system.
//
// [NSWritingDirection.natural]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/natural
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/baseWritingDirection
func (c NSCell) BaseWritingDirection() NSWritingDirection {
	rv := objc.Send[NSWritingDirection](c.ID, objc.Sel("baseWritingDirection"))
	return NSWritingDirection(rv)
}
func (c NSCell) SetBaseWritingDirection(value NSWritingDirection) {
	objc.Send[struct{}](c.ID, objc.Sel("setBaseWritingDirection:"), value)
}

// The cell’s value as an attributed string.
//
// # Discussion
// 
// Use this property to get the value of the cell interpreted as an attributed
// string. The textual attributes included in the string are the default
// paragraph style, the cell’s font and alignment, and whether the cell is
// enabled and scrollable.
// 
// When setting the value of this property, if the cell has a formatter, but
// the formatter does not understand the attributed string, the formatter
// marks the cell’s object as invalid. If the receiver is not a text-type
// cell, it is converted to one before the value is set.
// 
// If you use a class that has an [AttributedStringValue] property, the cell
// gets the string from that property instead of using the [StringValue]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/attributedStringValue
func (c NSCell) AttributedStringValue() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("attributedStringValue"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (c NSCell) SetAttributedStringValue(value foundation.NSAttributedString) {
	objc.Send[struct{}](c.ID, objc.Sel("setAttributedStringValue:"), value)
}

// A Boolean value indicating whether the cell allows the editing of its
// content’s text attributes by the user.
//
// # Discussion
// 
// When the value of this property is [true], the user can modify the font and
// other textual attributes of the cell. When the value is [false], the user
// cannot edit the text or import graphics, which effectively means the cell
// cannot support RTFD text.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/allowsEditingTextAttributes
func (c NSCell) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}
func (c NSCell) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}

// A Boolean value indicating whether the cell supports the importation of
// images into its text.
//
// # Discussion
// 
// When the value of this property is [true], the cell can import images into
// its text and support the RTFD text format.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/importsGraphics
func (c NSCell) ImportsGraphics() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("importsGraphics"))
	return rv
}
func (c NSCell) SetImportsGraphics(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setImportsGraphics:"), value)
}

// The cell’s title text.
//
// # Discussion
// 
// This property contains the string to display in the cell. Subclasses (such
// as [NSButtonCell]) may use the title for a different purpose.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/title
func (c NSCell) Title() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (c NSCell) SetTitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The action performed by the cell.
//
// # Discussion
// 
// The value of this property is the selector to call on the cell’s [Target]
// object. Set the value of this property to `nil` to stop the delivery of
// action messages.
// 
// The default value of this property is `nil`. Setting the value of this
// property raises with [internalInconsistencyException]. Subclasses are
// expected to override this property as part of their target/action
// implementation.
//
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/action
func (c NSCell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c.ID, objc.Sel("action"))
	return rv
}
func (c NSCell) SetAction(value objc.SEL) {
	objc.Send[struct{}](c.ID, objc.Sel("setAction:"), value)
}

// The object that receives the cell’s action messages.
//
// # Discussion
// 
// The value of this property is the object that implements the selector
// specified by the [Action] property. Set the value of this property to `nil`
// to stop the delivery of action messages.
// 
// The default value of this property is `nil`. Setting the value of this
// property raises with [internalInconsistencyException]. Subclasses are
// expected to override this property as part of their target/action
// implementation.
//
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/target
func (c NSCell) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (c NSCell) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setTarget:"), value)
}

// A Boolean value indicating whether the cell sends its action message
// continuously during mouse tracking.
//
// # Discussion
// 
// When the value of this property is YES, the cell’s action message is sent
// continuously during mouse tracking. In practice, the continuous delivery of
// action messages has meaning only for [NSActionCell] and its subclasses,
// which implement the target/action mechanism. Some [NSControl] subclasses,
// notably [NSMatrix], send a default action to a default target when a cell
// doesn’t provide a target or action.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isContinuous
func (c NSCell) Continuous() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isContinuous"))
	return rv
}
func (c NSCell) SetContinuous(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setContinuous:"), value)
}

// The image displayed by the cell, if any.
//
// # Discussion
// 
// Setting the value of this property converts the cell to an image-type cell,
// if it is not one already. The value of this property is `nil` if the cell
// is not an image-type cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/image
func (c NSCell) Image() INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (c NSCell) SetImage(value INSImage) {
	objc.Send[struct{}](c.ID, objc.Sel("setImage:"), value)
}

// A tag for identifying the cell.
//
// # Discussion
// 
// Tags allow you to identify particular cells. Tag values are not used
// internally by AppKit. You typically set tag values in Interface Builder and
// use them at runtime in your application. When you set the tag of a control
// with a single cell in Interface Builder, it sets the tags of both the
// control and the cell to the same value as a convenience.
// 
// The default value of this property is `-1`. Setting the value of this
// property raises with [internalInconsistencyException]. Subclasses are
// expected to override this property if they support tags. The [NSActionCell]
// class implements this property and stores the integer you specify.
//
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/tag
func (c NSCell) Tag() int {
	rv := objc.Send[int](c.ID, objc.Sel("tag"))
	return rv
}
func (c NSCell) SetTag(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setTag:"), value)
}

// The cell’s formatter object.
//
// # Discussion
// 
// A formatter handles the translation of the receiver’s contents between
// its onscreen representation and its object value. Cells use a formatter
// object to format the textual representation of their object value and to
// validate cell input and convert that input to an object value. When
// assigning a new formatter to a cell, the formatter attempts to interpret
// the cell’s current value. If it cannot do so, the formatter converts the
// current value to a string object.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/formatter
func (c NSCell) Formatter() foundation.NSFormatter {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("formatter"))
	return foundation.NSFormatterFromID(objc.ID(rv))
}
func (c NSCell) SetFormatter(value foundation.NSFormatter) {
	objc.Send[struct{}](c.ID, objc.Sel("setFormatter:"), value)
}

// The cell’s contextual menu.
//
// # Discussion
// 
// Use this property to specify a menu containing contextual commands
// associated with the cell. If the cell does not have a menu, set this
// property to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/menu
func (c NSCell) Menu() INSMenu {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("menu"))
	return NSMenuFromID(objc.ID(rv))
}
func (c NSCell) SetMenu(value INSMenu) {
	objc.Send[struct{}](c.ID, objc.Sel("setMenu:"), value)
}

// A Boolean value indicating whether the cell accepts first responder status.
//
// # Discussion
// 
// When the value of this property is [true], the cell is able to become the
// first responder. The default value of this property is [true] when the cell
// is enabled. Subclasses may override this method to return a different
// value.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/acceptsFirstResponder
func (c NSCell) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}

// A Boolean value indicating whether the cell provides a visual indication
// that it is the first responder.
//
// # Discussion
// 
// When the value of this property is [true] and the cell becomes the first
// responder, the cell performs additional drawing to indicate that it is the
// first responder. The [NSCell] class itself does not draw a first-responder
// indicator. Subclasses may use the value in this property to determine
// whether or not they should draw one.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/showsFirstResponder
func (c NSCell) ShowsFirstResponder() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("showsFirstResponder"))
	return rv
}
func (c NSCell) SetShowsFirstResponder(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShowsFirstResponder:"), value)
}

// A Boolean value indicating whether the cell refuses the first responder
// status.
//
// # Discussion
// 
// Set the value of this property to [true] to prevent the cell from becoming
// the first responder. To determine whether the cell can become first
// responder right now, get the value of the [AcceptsFirstResponder] property.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/refusesFirstResponder
func (c NSCell) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("refusesFirstResponder"))
	return rv
}
func (c NSCell) SetRefusesFirstResponder(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setRefusesFirstResponder:"), value)
}

// The object represented by the cell.
//
// # Discussion
// 
// Use this property to link the cell an appropriate object. For example, in a
// pop-up list of color names, the represented object of each cell could be
// the appropriate [NSColor] object.
// 
// # Special Considerations
// 
// When you copy an [NSCell] object, the value of this property is set to
// `nil` in the copy.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/representedObject
func (c NSCell) RepresentedObject() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("representedObject"))
	return objectivec.Object{ID: rv}
}
func (c NSCell) SetRepresentedObject(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setRepresentedObject:"), value)
}

// The modifier flags for the last (left) mouse-down event.
//
// # Discussion
// 
// The value of this property is the value of the modifier flags from the most
// recent [NSEvent] object representing a mouse-down event. If tracking has
// not yet occurred or the event contained no modifier keys, the value of this
// property is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/mouseDownFlags
func (c NSCell) MouseDownFlags() int {
	rv := objc.Send[int](c.ID, objc.Sel("mouseDownFlags"))
	return rv
}

// The key equivalent associated with clicking the cell.
//
// # Discussion
// 
// Subclasses can override this property and return a string with a valid
// character for the key equivalent. The default implementation of this
// property returns an empty string.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/keyEquivalent
func (c NSCell) KeyEquivalent() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keyEquivalent"))
	return foundation.NSStringFromID(rv).String()
}

// The type of focus ring to use with the associated view.
//
// # Discussion
// 
// You can disable a view’s focus ring drawing by setting this property to
// [NSFocusRingType.none]. The only times you should disable focus ring
// drawing are when you want to draw your own focus ring or when there is
// insufficient space to display a focus ring in the default location.
//
// [NSFocusRingType.none]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/none
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/focusRingType
func (c NSCell) FocusRingType() NSFocusRingType {
	rv := objc.Send[NSFocusRingType](c.ID, objc.Sel("focusRingType"))
	return NSFocusRingType(rv)
}
func (c NSCell) SetFocusRingType(value NSFocusRingType) {
	objc.Send[struct{}](c.ID, objc.Sel("setFocusRingType:"), value)
}

// The minimum size needed to display the cell.
//
// # Discussion
// 
// This property contains the smallest cell size (in points) required to draw
// its contents. If the cell is not a text or image cell, the cell size is set
// to (`10000`, `10000`). If an image cell does not yet have an associated
// image, the cell size is [NSZeroSize].This method takes into account of the
// size of the image or text in the cell along with any margin areas required
// by the cell’s border, if any.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/cellSize
func (c NSCell) CellSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("cellSize"))
	return corefoundation.CGSize(rv)
}

// The size of the cell.
//
// # Discussion
// 
// Use this property to change the rendered size of the cell and its control.
// For a list of possible values, see [NSControl.ControlSize].
// 
// Changing the cell’s control size does not change the font used by the
// cell. Use the [SystemFontSizeForControlSize] class method of [NSFont] to
// obtain an appropriate font based on the new control size.
//
// [NSControl.ControlSize]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/controlSize
func (c NSCell) ControlSize() NSControlSize {
	rv := objc.Send[NSControlSize](c.ID, objc.Sel("controlSize"))
	return NSControlSize(rv)
}
func (c NSCell) SetControlSize(value NSControlSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setControlSize:"), value)
}

// The view associated with the cell.
//
// # Discussion
// 
// This property contains the view associated with the cell. This view is
// normally an [NSControl] object. The default value of this property is
// `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/controlView
func (c NSCell) ControlView() INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("controlView"))
	return NSViewFromID(objc.ID(rv))
}
func (c NSCell) SetControlView(value INSView) {
	objc.Send[struct{}](c.ID, objc.Sel("setControlView:"), value)
}

// A Boolean value indicating whether the cell has a highlighted appearance.
//
// # Discussion
// 
// When the value of this property is [true], the cell draws itself with a
// highlighted appearance. The default value of this property is [false].
// 
// Assigning a new value to this property has no effect by default. Subclasses
// can override the property to provide a highlighting behavior. For example,
// the [NSButtonCell] class overrides this property, so that when the value is
// [true] the button draws the button with a highlight appearance specified by
// [NSCell.Attribute.cellLightsByBackground],
// [NSCell.Attribute.cellLightsByContents], or
// [NSCell.Attribute.cellLightsByGray].
//
// [NSCell.Attribute.cellLightsByBackground]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellLightsByBackground
// [NSCell.Attribute.cellLightsByContents]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellLightsByContents
// [NSCell.Attribute.cellLightsByGray]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellLightsByGray
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/isHighlighted
func (c NSCell) Highlighted() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isHighlighted"))
	return rv
}
func (c NSCell) SetHighlighted(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHighlighted:"), value)
}

// A Boolean value indicating whether the cell’s control object sends its
// action message when the user finishes editing the cell’s text.
//
// # Discussion
// 
// When the value of this property is [true], the control sends its action
// message when editing is complete. Editing is complete when the user does
// one of the following:
// 
// - Presses the Return key - Presses the Tab key to move out of the field -
// Clicks another text field
// 
// When the value is [false], the control sends its action message only when
// the user presses the Return key.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/sendsActionOnEndEditing
func (c NSCell) SendsActionOnEndEditing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("sendsActionOnEndEditing"))
	return rv
}
func (c NSCell) SetSendsActionOnEndEditing(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSendsActionOnEndEditing:"), value)
}

// A Boolean value indicating whether the cell’s field editor should post
// text change notifications.
//
// # Discussion
// 
// When the value of this property is [true], the field editor should post
// text notification changes while editing marked text. When the value is NO,
// the field editor should delay notifications until the marked text is
// confirmed.
// 
// The default value of this property is [false]. Subclasses can override the
// property and change the value as appropriate.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/wantsNotificationForMarkedText
func (c NSCell) WantsNotificationForMarkedText() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("wantsNotificationForMarkedText"))
	return rv
}

// A Boolean value indicating whether the cell restricts layout and rendering
// of text to a single line.
//
// # Discussion
// 
// When the value of this property is [true], text layout and rendering is
// restricted to a single line. In addition, the cell ignores the return value
// from [Wraps], interprets [NSLineBreakMode.byWordWrapping] and
// [NSLineBreakMode.byCharWrapping] returned by [LineBreakMode] as
// [NSLineBreakMode.byClipping], and configures the field editor to ignore key
// binding commands that insert paragraph and line separators.
// 
// The field editor bound to a single-line cell filters out paragraph and line
// separator insertion from user actions. Cells in single-line mode use the
// fixed baseline layout. The text baseline position is determined solely by
// the control size regardless of content font style or size.
//
// [NSLineBreakMode.byCharWrapping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byCharWrapping
// [NSLineBreakMode.byClipping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byClipping
// [NSLineBreakMode.byWordWrapping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byWordWrapping
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/usesSingleLineMode
func (c NSCell) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("usesSingleLineMode"))
	return rv
}
func (c NSCell) SetUsesSingleLineMode(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setUsesSingleLineMode:"), value)
}

// The layout direction of the user interface.
//
// # Discussion
// 
// This property specifies the general user interface layout flow directions.
// For subclasses that have multiple visual components in a single cell
// instance, this property should specify the directionality or flow of
// components.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/userInterfaceLayoutDirection
func (c NSCell) UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection {
	rv := objc.Send[NSUserInterfaceLayoutDirection](c.ID, objc.Sel("userInterfaceLayoutDirection"))
	return NSUserInterfaceLayoutDirection(rv)
}
func (c NSCell) SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
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
func (c NSCell) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (c NSCell) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](c.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}

// Returns the default menu for instances of the cell.
//
// # Return Value
// 
// The default menu. The [NSCell] implementation of this method returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/defaultMenu
func (_NSCellClass NSCellClass) DefaultMenu() NSMenu {
	rv := objc.Send[objc.ID](objc.ID(_NSCellClass.class), objc.Sel("defaultMenu"))
	return NSMenuFromID(objc.ID(rv))
}

// Returns a Boolean value that indicates whether tracking stops when the
// cursor leaves the cell.
//
// # Return Value
// 
// [true] if tracking stops when the cursor leaves the cell, otherwise
// [false].
// 
// # Discussion
// 
// The default implementation returns [false]. Subclasses may override this
// method to return a different value.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/prefersTrackingUntilMouseUp
func (_NSCellClass NSCellClass) PrefersTrackingUntilMouseUp() bool {
	rv := objc.Send[bool](objc.ID(_NSCellClass.class), objc.Sel("prefersTrackingUntilMouseUp"))
	return rv
}

// Returns the default type of focus ring for the receiver.
//
// # Return Value
// 
// The default type of focus ring for the receiver (one of the values listed
// in [NSFocusRingType]).
//
// [NSFocusRingType]: https://developer.apple.com/documentation/AppKit/NSFocusRingType
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/defaultFocusRingType
func (_NSCellClass NSCellClass) DefaultFocusRingType() NSFocusRingType {
	rv := objc.Send[NSFocusRingType](objc.ID(_NSCellClass.class), objc.Sel("defaultFocusRingType"))
	return NSFocusRingType(rv)
}

// Sent after the user changes control tint preference.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/currentcontroltintdidchangenotification
func (_NSCellClass NSCellClass) CurrentControlTintDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSCellClass.class), objc.Sel("NSControlTintDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

			// Protocol methods for NSAccessibilityElementProtocol
			

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()

func (o NSCell) AccessibilityFrame() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
// 
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()

func (o NSCell) AccessibilityParent() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}

// Returns the accessibility element’s identity.
//
// # Return Value
// 
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()

func (o NSCell) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()

func (o NSCell) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

			// Protocol methods for NSAccessibilityProtocol
			

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityelement()

func (o NSCell) IsAccessibilityElement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityelement(_:)

func (o NSCell) SetAccessibilityElement(accessibilityElement bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
	}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityenabled()

func (o NSCell) IsAccessibilityEnabled() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityenabled(_:)

func (o NSCell) SetAccessibilityEnabled(accessibilityEnabled bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
	}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityframe(_:)

func (o NSCell) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
	}

// Returns the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhelp()

func (o NSCell) AccessibilityHelp() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhelp(_:)

func (o NSCell) SetAccessibilityHelp(accessibilityHelp string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
	}

// Returns a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabel()

func (o NSCell) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabel(_:)

func (o NSCell) SetAccessibilityLabel(accessibilityLabel string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
	}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitle()

func (o NSCell) AccessibilityTitle() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the title of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitle(_:)

func (o NSCell) SetAccessibilityTitle(accessibilityTitle string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
	}

// Returns the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvalue()

func (o NSCell) AccessibilityValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvalue(_:)

func (o NSCell) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
	}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
// 
// [true], if accessibility clients can call the selector; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)

func (o NSCell) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
	}

// Returns the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycontents()

func (o NSCell) AccessibilityContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycontents(_:)

func (o NSCell) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
	}

// Returns the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycriticalvalue()

func (o NSCell) AccessibilityCriticalValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycriticalvalue(_:)

func (o NSCell) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
	}

// Sets the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityidentifier(_:)

func (o NSCell) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
	}

// Returns the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymaxvalue()

func (o NSCell) AccessibilityMaxValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymaxvalue(_:)

func (o NSCell) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
	}

// Returns the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminvalue()

func (o NSCell) AccessibilityMinValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminvalue(_:)

func (o NSCell) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
	}

// Returns the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityorientation()

func (o NSCell) AccessibilityOrientation() NSAccessibilityOrientation {
	
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
	}

// Sets the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorientation(_:)

func (o NSCell) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
	}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityprotectedcontent()

func (o NSCell) IsAccessibilityProtectedContent() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityprotectedcontent(_:)

func (o NSCell) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
	}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityselected()

func (o NSCell) IsAccessibilitySelected() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselected(_:)

func (o NSCell) SetAccessibilitySelected(accessibilitySelected bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
	}

// Returns the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityurl()

func (o NSCell) AccessibilityURL() foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
	}

// Sets the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityurl(_:)

func (o NSCell) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
	}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvaluedescription()

func (o NSCell) AccessibilityValueDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvaluedescription(_:)

func (o NSCell) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
	}

// Returns the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywarningvalue()

func (o NSCell) AccessibilityWarningValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywarningvalue(_:)

func (o NSCell) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
	}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildren()

func (o NSCell) AccessibilityChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildren(_:)

func (o NSCell) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
	}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildreninnavigationorder()

func (o NSCell) AccessibilityChildrenInNavigationOrder() unsafe.Pointer {
	
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return rv
	}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildreninnavigationorder(_:)

func (o NSCell) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
	}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityparent(_:)

func (o NSCell) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
	}

// Returns the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedchildren()

func (o NSCell) AccessibilitySelectedChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedchildren(_:)

func (o NSCell) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
	}

// Returns the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytopleveluielement()

func (o NSCell) AccessibilityTopLevelUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytopleveluielement(_:)

func (o NSCell) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
	}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblechildren()

func (o NSCell) AccessibilityVisibleChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblechildren(_:)

func (o NSCell) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
	}

// Returns the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityapplicationfocuseduielement()

func (o NSCell) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityapplicationfocuseduielement(_:)

func (o NSCell) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
	}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocused(_:)

func (o NSCell) SetAccessibilityFocused(accessibilityFocused bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
	}

// Returns the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfocusedwindow()

func (o NSCell) AccessibilityFocusedWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocusedwindow(_:)

func (o NSCell) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
	}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedfocuselements()

func (o NSCell) AccessibilitySharedFocusElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedfocuselements(_:)

func (o NSCell) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
	}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityrequired()

func (o NSCell) IsAccessibilityRequired() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrequired(_:)

func (o NSCell) SetAccessibilityRequired(accessibilityRequired bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
	}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrole()

func (o NSCell) AccessibilityRole() NSAccessibilityRole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
	}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrole(_:)

func (o NSCell) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
	}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityroledescription()

func (o NSCell) AccessibilityRoleDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityroledescription(_:)

func (o NSCell) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
	}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysubrole()

func (o NSCell) AccessibilitySubrole() NSAccessibilitySubrole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
	}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysubrole(_:)

func (o NSCell) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
	}

// Returns the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomactions()

func (o NSCell) AccessibilityCustomActions() INSAccessibilityCustomAction {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return NSAccessibilityCustomActionFromID(rv)
	}

// Sets the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomactions(_:)

func (o NSCell) SetAccessibilityCustomActions(accessibilityCustomActions foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
	}

// Returns the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomrotors()

func (o NSCell) AccessibilityCustomRotors() INSAccessibilityCustomRotor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return NSAccessibilityCustomRotorFromID(rv)
	}

// Sets the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomrotors(_:)

func (o NSCell) SetAccessibilityCustomRotors(accessibilityCustomRotors foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
	}

// Returns the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityinsertionpointlinenumber()

func (o NSCell) AccessibilityInsertionPointLineNumber() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
	}

// Sets the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityinsertionpointlinenumber(_:)

func (o NSCell) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
	}

// Returns the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynumberofcharacters()

func (o NSCell) AccessibilityNumberOfCharacters() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
	}

// Sets the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynumberofcharacters(_:)

func (o NSCell) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
	}

// Returns the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityplaceholdervalue()

func (o NSCell) AccessibilityPlaceholderValue() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityplaceholdervalue(_:)

func (o NSCell) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
	}

// Returns the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtext()

func (o NSCell) AccessibilitySelectedText() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtext(_:)

func (o NSCell) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
	}

// Returns the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextrange()

func (o NSCell) AccessibilitySelectedTextRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
	}

// Sets the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextrange(_:)

func (o NSCell) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
	}

// Returns an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextranges()

func (o NSCell) AccessibilitySelectedTextRanges() foundation.NSValue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSValueFromID(rv)
	}

// Sets an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextranges(_:)

func (o NSCell) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
	}

// Returns the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedcharacterrange()

func (o NSCell) AccessibilitySharedCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
	}

// Sets the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedcharacterrange(_:)

func (o NSCell) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
	}

// Returns the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedtextuielements()

func (o NSCell) AccessibilitySharedTextUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedtextuielements(_:)

func (o NSCell) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
	}

// Returns the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecharacterrange()

func (o NSCell) AccessibilityVisibleCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
	}

// Sets the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecharacterrange(_:)

func (o NSCell) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
	}

// Returns the substring for the specified range.
//
// range: A range of characters contained by the element.
//
// # Return Value
// 
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)

func (o NSCell) AccessibilityStringForRange(range_ foundation.NSRange) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
	}

// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)

func (o NSCell) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
	}

// Returns the rich text format (RTF) data that describes the specified range
// of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// A data object containing an RTF representation of the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)

func (o NSCell) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
	}

// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// The rectangle that encloses the specified characters.
//
// # Discussion
// 
// If the range crosses a line boundary, the returned rectangle fully encloses
// all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)

func (o NSCell) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
	}

// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
// 
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)

func (o NSCell) AccessibilityLineForIndex(index int) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
	}

// Returns the range of characters for the glyph that includes the specified
// character.
//
// index: The specified character.
//
// # Return Value
// 
// The range of characters for the glyph.
//
// # Discussion
// 
// This value always includes the specified character but may include
// additional characters if that character is part of a multicharacter glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3

func (o NSCell) AccessibilityRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
	}

// Returns a range of characters that all have the same style as the specified
// character.
//
// index: The index of the specified character.
//
// # Return Value
// 
// A range of characters with the same style as the specified character.
//
// # Discussion
// 
// This method returns a range of characters that meet two conditions: The
// range must include the specified character, and all the other characters in
// the range must match the specified character’s style. If none of the
// adjacent characters match the specified character’s style, the method
// returns only the specified character.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)

func (o NSCell) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
	}

// Returns the range of characters in the specified line.
//
// line: The line number to be examined.
//
// # Return Value
// 
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)

func (o NSCell) AccessibilityRangeForLine(line int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
	}

// Returns the range of characters for the glyph at the specified point.
//
// point: A point in screen coordinates.
//
// # Return Value
// 
// The range of characters that make up the glyph at the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm

func (o NSCell) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
	}

// Returns the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityactivationpoint()

func (o NSCell) AccessibilityActivationPoint() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
	}

// Sets the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityactivationpoint(_:)

func (o NSCell) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
	}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityalternateuivisible()

func (o NSCell) IsAccessibilityAlternateUIVisible() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
	}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityalternateuivisible(_:)

func (o NSCell) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
	}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycancelbutton()

func (o NSCell) AccessibilityCancelButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycancelbutton(_:)

func (o NSCell) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
	}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclosebutton()

func (o NSCell) AccessibilityCloseButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclosebutton(_:)

func (o NSCell) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
	}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydefaultbutton()

func (o NSCell) AccessibilityDefaultButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydefaultbutton(_:)

func (o NSCell) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
	}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfullscreenbutton()

func (o NSCell) AccessibilityFullScreenButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfullscreenbutton(_:)

func (o NSCell) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
	}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitygrowarea()

func (o NSCell) AccessibilityGrowArea() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitygrowarea(_:)

func (o NSCell) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
	}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymain()

func (o NSCell) IsAccessibilityMain() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
	}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymain(_:)

func (o NSCell) SetAccessibilityMain(accessibilityMain bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
	}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminimizebutton()

func (o NSCell) AccessibilityMinimizeButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimizebutton(_:)

func (o NSCell) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
	}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityminimized()

func (o NSCell) IsAccessibilityMinimized() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
	}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimized(_:)

func (o NSCell) SetAccessibilityMinimized(accessibilityMinimized bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
	}

// Returns a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymodal()

func (o NSCell) IsAccessibilityModal() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
	}

// Sets a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymodal(_:)

func (o NSCell) SetAccessibilityModal(accessibilityModal bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
	}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityproxy()

func (o NSCell) AccessibilityProxy() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityproxy(_:)

func (o NSCell) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
	}

// Returns the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityshownmenu()

func (o NSCell) AccessibilityShownMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityshownmenu(_:)

func (o NSCell) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
	}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytoolbarbutton()

func (o NSCell) AccessibilityToolbarButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytoolbarbutton(_:)

func (o NSCell) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
	}

// Returns the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindow()

func (o NSCell) AccessibilityWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindow(_:)

func (o NSCell) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
	}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityzoombutton()

func (o NSCell) AccessibilityZoomButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityzoombutton(_:)

func (o NSCell) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
	}

// Returns the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityextrasmenubar()

func (o NSCell) AccessibilityExtrasMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityextrasmenubar(_:)

func (o NSCell) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
	}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfrontmost()

func (o NSCell) IsAccessibilityFrontmost() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
	}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfrontmost(_:)

func (o NSCell) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
	}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityhidden()

func (o NSCell) IsAccessibilityHidden() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
	}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhidden(_:)

func (o NSCell) SetAccessibilityHidden(accessibilityHidden bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
	}

// Returns the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymainwindow()

func (o NSCell) AccessibilityMainWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymainwindow(_:)

func (o NSCell) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
	}

// Returns the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymenubar()

func (o NSCell) AccessibilityMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymenubar(_:)

func (o NSCell) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
	}

// Returns an array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindows()

func (o NSCell) AccessibilityWindows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindows(_:)

func (o NSCell) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
	}

// Returns the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumncount()

func (o NSCell) AccessibilityColumnCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
	}

// Sets the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumncount(_:)

func (o NSCell) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
	}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityorderedbyrow()

func (o NSCell) IsAccessibilityOrderedByRow() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
	}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorderedbyrow(_:)

func (o NSCell) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
	}

// Returns the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowcount()

func (o NSCell) AccessibilityRowCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
	}

// Sets the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowcount(_:)

func (o NSCell) SetAccessibilityRowCount(accessibilityRowCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
	}

// Returns the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalscrollbar()

func (o NSCell) AccessibilityHorizontalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalscrollbar(_:)

func (o NSCell) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
	}

// Returns the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalscrollbar()

func (o NSCell) AccessibilityVerticalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalscrollbar(_:)

func (o NSCell) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
	}

// Returns the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnheaderuielements()

func (o NSCell) AccessibilityColumnHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnheaderuielements(_:)

func (o NSCell) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
	}

// Returns the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumns()

func (o NSCell) AccessibilityColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumns(_:)

func (o NSCell) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
	}

// Returns the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumntitles()

func (o NSCell) AccessibilityColumnTitles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return objectivec.Object{ID: rv}
	}

// Sets the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumntitles(_:)

func (o NSCell) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
	}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityexpanded()

func (o NSCell) IsAccessibilityExpanded() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
	}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityexpanded(_:)

func (o NSCell) SetAccessibilityExpanded(accessibilityExpanded bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
	}

// Returns the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityheader()

func (o NSCell) AccessibilityHeader() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
	}

// Sets the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityheader(_:)

func (o NSCell) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
	}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityindex()

func (o NSCell) AccessibilityIndex() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
	}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityindex(_:)

func (o NSCell) SetAccessibilityIndex(accessibilityIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
	}

// Returns the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowheaderuielements()

func (o NSCell) AccessibilityRowHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowheaderuielements(_:)

func (o NSCell) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
	}

// Returns the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrows()

func (o NSCell) AccessibilityRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrows(_:)

func (o NSCell) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
	}

// Returns the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcolumns()

func (o NSCell) AccessibilitySelectedColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcolumns(_:)

func (o NSCell) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
	}

// Returns the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedrows()

func (o NSCell) AccessibilitySelectedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedrows(_:)

func (o NSCell) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
	}

// Returns the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysortdirection()

func (o NSCell) AccessibilitySortDirection() NSAccessibilitySortDirection {
	
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
	}

// Sets the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysortdirection(_:)

func (o NSCell) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
	}

// Returns the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecolumns()

func (o NSCell) AccessibilityVisibleColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecolumns(_:)

func (o NSCell) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
	}

// Returns the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblerows()

func (o NSCell) AccessibilityVisibleRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblerows(_:)

func (o NSCell) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
	}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitydisclosed()

func (o NSCell) IsAccessibilityDisclosed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
	}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosed(_:)

func (o NSCell) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
	}

// Returns the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedbyrow()

func (o NSCell) AccessibilityDisclosedByRow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
	}

// Sets the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedbyrow(_:)

func (o NSCell) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
	}

// Returns the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedrows()

func (o NSCell) AccessibilityDisclosedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedrows(_:)

func (o NSCell) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
	}

// Returns the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosurelevel()

func (o NSCell) AccessibilityDisclosureLevel() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
	}

// Sets the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosurelevel(_:)

func (o NSCell) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
	}

// Returns the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnindexrange()

func (o NSCell) AccessibilityColumnIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
	}

// Sets the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnindexrange(_:)

func (o NSCell) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
	}

// Returns the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowindexrange()

func (o NSCell) AccessibilityRowIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
	}

// Sets the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowindexrange(_:)

func (o NSCell) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
	}

// Returns the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcells()

func (o NSCell) AccessibilitySelectedCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcells(_:)

func (o NSCell) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
	}

// Returns the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecells()

func (o NSCell) AccessibilityVisibleCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecells(_:)

func (o NSCell) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
	}

// Returns the cell at the specified column and row.
//
// column: The column index.
//
// row: The row index.
//
// # Return Value
// 
// The cell specified by the column and row indexes.
//
// # Discussion
// 
// This property is required for all elements that function as cell-based
// tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCell(forColumn:row:)

func (o NSCell) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
	}

// Returns the drag handle elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhandles()

func (o NSCell) AccessibilityHandles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return objectivec.Object{ID: rv}
	}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhandles(_:)

func (o NSCell) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
	}

// Returns the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunits()

func (o NSCell) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
	}

// Sets the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunits(_:)

func (o NSCell) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
	}

// Returns the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunitdescription()

func (o NSCell) AccessibilityHorizontalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunitdescription(_:)

func (o NSCell) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
	}

// Returns the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunits()

func (o NSCell) AccessibilityVerticalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
	}

// Sets the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunits(_:)

func (o NSCell) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
	}

// Returns the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunitdescription()

func (o NSCell) AccessibilityVerticalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunitdescription(_:)

func (o NSCell) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
	}

// Converts the provided point in screen coordinates to a point in the layout
// area’s coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
// 
// A point in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)

func (o NSCell) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
	}

// Converts the provided size in screen coordinates to a size in the layout
// area’s coordinate system.
//
// size: A size in the screen’s coordinate system.
//
// # Return Value
// 
// A size in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)

func (o NSCell) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
	}

// Converts the provided point in the layout area’s coordinates to a point
// in the screen’s coordinate system.
//
// point: A point in the layout area’s coordinate system.
//
// # Return Value
// 
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)

func (o NSCell) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
	}

// Converts the provided size in the layout area’s coordinates to a size in
// the screen’s coordinate system.
//
// size: A size in the layout area’s coordinate system.
//
// # Return Value
// 
// A size in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)

func (o NSCell) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
	}

// Returns the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityallowedvalues()

func (o NSCell) AccessibilityAllowedValues() foundation.NSNumber {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSNumberFromID(rv)
	}

// Sets the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityallowedvalues(_:)

func (o NSCell) SetAccessibilityAllowedValues(accessibilityAllowedValues foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
	}

// Returns the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabeluielements()

func (o NSCell) AccessibilityLabelUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabeluielements(_:)

func (o NSCell) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
	}

// Returns the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabelvalue()

func (o NSCell) AccessibilityLabelValue() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
	}

// Sets the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabelvalue(_:)

func (o NSCell) SetAccessibilityLabelValue(accessibilityLabelValue float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
	}

// Returns the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynextcontents()

func (o NSCell) AccessibilityNextContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynextcontents(_:)

func (o NSCell) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
	}

// Returns the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitypreviouscontents()

func (o NSCell) AccessibilityPreviousContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitypreviouscontents(_:)

func (o NSCell) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
	}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysplitters()

func (o NSCell) AccessibilitySplitters() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysplitters(_:)

func (o NSCell) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
	}

// Returns the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityoverflowbutton()

func (o NSCell) AccessibilityOverflowButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityoverflowbutton(_:)

func (o NSCell) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
	}

// Returns the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytabs()

func (o NSCell) AccessibilityTabs() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return objectivec.Object{ID: rv}
	}

// Sets the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytabs(_:)

func (o NSCell) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
	}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkergroupuielement()

func (o NSCell) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkergroupuielement(_:)

func (o NSCell) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
	}

// Returns the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkertypedescription()

func (o NSCell) AccessibilityMarkerTypeDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkertypedescription(_:)

func (o NSCell) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
	}

// Returns the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkeruielements()

func (o NSCell) AccessibilityMarkerUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkeruielements(_:)

func (o NSCell) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
	}

// Returns the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkervalues()

func (o NSCell) AccessibilityMarkerValues() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
	}

// Sets the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkervalues(_:)

func (o NSCell) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
	}

// Returns the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrulermarkertype()

func (o NSCell) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
	}

// Sets the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrulermarkertype(_:)

func (o NSCell) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
	}

// Returns the units for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunits()

func (o NSCell) AccessibilityUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
	}

// Sets the units used for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunits(_:)

func (o NSCell) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
	}

// Returns the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunitdescription()

func (o NSCell) AccessibilityUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunitdescription(_:)

func (o NSCell) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
	}

// Returns the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydocument()

func (o NSCell) AccessibilityDocument() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydocument(_:)

func (o NSCell) SetAccessibilityDocument(accessibilityDocument string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
	}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityedited()

func (o NSCell) IsAccessibilityEdited() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
	}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityedited(_:)

func (o NSCell) SetAccessibilityEdited(accessibilityEdited bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
	}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfilename()

func (o NSCell) AccessibilityFilename() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the filename for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfilename(_:)

func (o NSCell) SetAccessibilityFilename(accessibilityFilename string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
	}

// Returns the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylinkeduielements()

func (o NSCell) AccessibilityLinkedUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylinkeduielements(_:)

func (o NSCell) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
	}

// Returns the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityservesastitleforuielements()

func (o NSCell) AccessibilityServesAsTitleForUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityservesastitleforuielements(_:)

func (o NSCell) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
	}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitleuielement()

func (o NSCell) AccessibilityTitleUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitleuielement(_:)

func (o NSCell) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
	}

// Returns the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclearbutton()

func (o NSCell) AccessibilityClearButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclearbutton(_:)

func (o NSCell) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
	}

// Returns the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchbutton()

func (o NSCell) AccessibilitySearchButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchbutton(_:)

func (o NSCell) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
	}

// Returns the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchmenu()

func (o NSCell) AccessibilitySearchMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchmenu(_:)

func (o NSCell) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
	}

// Cancels the current operation.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()

func (o NSCell) AccessibilityPerformCancel() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
	}

// Simulates pressing Return in the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()

func (o NSCell) AccessibilityPerformConfirm() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
	}

// Selects the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()

func (o NSCell) AccessibilityPerformPick() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
	}

// Simulates clicking the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()

func (o NSCell) AccessibilityPerformPress() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
	}

// Displays the accessibility element’s alternative UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()

func (o NSCell) AccessibilityPerformShowAlternateUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
	}

// Returns to the accessibility element’s original UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()

func (o NSCell) AccessibilityPerformShowDefaultUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
	}

// Displays the menu accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()

func (o NSCell) AccessibilityPerformShowMenu() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
	}

// Brings the window to the front.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()

func (o NSCell) AccessibilityPerformRaise() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
	}

// Returns the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityincrementbutton()

func (o NSCell) AccessibilityIncrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityincrementbutton(_:)

func (o NSCell) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
	}

// Returns the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydecrementbutton()

func (o NSCell) AccessibilityDecrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydecrementbutton(_:)

func (o NSCell) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
	}

// Increments the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()

func (o NSCell) AccessibilityPerformIncrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
	}

// Decrements the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()

func (o NSCell) AccessibilityPerformDecrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
	}

// Deletes the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()

func (o NSCell) AccessibilityPerformDelete() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
	}

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityattributeduserinputlabels()

func (o NSCell) AccessibilityAttributedUserInputLabels() foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSAttributedStringFromID(rv)
	}

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityuserinputlabels()

func (o NSCell) AccessibilityUserInputLabels() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSStringFromID(rv).String()
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityattributeduserinputlabels(_:)

func (o NSCell) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityuserinputlabels(_:)

func (o NSCell) SetAccessibilityUserInputLabels(accessibilityUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
	}

			// Protocol methods for NSUserInterfaceItemIdentification
			

