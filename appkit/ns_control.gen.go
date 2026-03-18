// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSControl] class.
var (
	_NSControlClass     NSControlClass
	_NSControlClassOnce sync.Once
)

func getNSControlClass() NSControlClass {
	_NSControlClassOnce.Do(func() {
		_NSControlClass = NSControlClass{class: objc.GetClass("NSControl")}
	})
	return _NSControlClass
}

// GetNSControlClass returns the class object for NSControl.
func GetNSControlClass() NSControlClass {
	return getNSControlClass()
}

type NSControlClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSControlClass) Alloc() NSControl {
	rv := objc.Send[NSControl](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specialized view, such as a button or text field, that notifies your app
// of relevant events using the target-action design pattern.
//
// # Overview
// 
// The [NSControl] class is abstract and must be subclassed to be used.
// Although you can subclass it yourself, more often you use one of the
// subclasses already defined by AppKit. A control draws content on the
// screen, automatically handles user interactions with that content, and
// calls the action method of its target object for any significant user
// interactions.
// 
// # About delegate methods
// 
// The [NSControl] class provides several delegate methods for its subclasses
// that allow text editing, such as [NSTextField] and [NSMatrix]. These
// include: [controlTextDidBeginEditing:], [controlTextDidChange:], and
// [controlTextDidEndEditing:].
// 
// Note that although [NSControl] defines delegate methods, it doesn’t
// itself have a delegate. Any subclass that uses these methods must have a
// delegate and the methods to get and set it. In addition, a formal delegate
// protocol [NSControlTextEditingDelegate] also defines delegate methods used
// by control delegates.
// 
// # Responding to mouse events
// 
// When the mouse button is pressed while the cursor is within the bounds of
// the receiver, the system calls [NSControl.MouseDown]. This method highlights the
// receiver’s cell and sends it a [TrackMouseInRectOfViewUntilMouseUp]
// message. Whenever the cell finishes tracking the mouse (for example,
// because the cursor has left the cell’s bounds), the cell is
// unhighlighted. If the mouse button is still down and the cursor reenters
// the bounds, the cell is again highlighted and a new
// [TrackMouseInRectOfViewUntilMouseUp] message is sent. This behavior repeats
// until the mouse button goes up. If it goes up with the cursor in the
// control, the state of the control is changed, and the action message is
// sent to the target. If the mouse button goes up when the cursor is outside
// the control, no action message is sent.
//
// [controlTextDidBeginEditing:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidBeginEditing:
// [controlTextDidChange:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidChange:
// [controlTextDidEndEditing:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidEndEditing:
//
// # Enabling and Disabling the Control
//
//   - [NSControl.Enabled]: A Boolean value that indicates whether the receiver reacts to mouse events.
//   - [NSControl.SetEnabled]
//
// # Accessing the Control’s Value
//
//   - [NSControl.DoubleValue]: The value of the receiver’s cell as a double-precision floating-point number.
//   - [NSControl.SetDoubleValue]
//   - [NSControl.FloatValue]: The value of the receiver’s cell as a single-precision floating-point number.
//   - [NSControl.SetFloatValue]
//   - [NSControl.IntValue]: The value of the receiver’s cell as an integer.
//   - [NSControl.SetIntValue]
//   - [NSControl.IntegerValue]: The value of the receiver’s cell as an integer value.
//   - [NSControl.SetIntegerValue]
//   - [NSControl.ObjectValue]: The value of the receiver’s cell as an Objective-C object.
//   - [NSControl.SetObjectValue]
//   - [NSControl.StringValue]: The value of the receiver’s cell as an [NSString] object.
//   - [NSControl.SetStringValue]
//   - [NSControl.AttributedStringValue]: The value of the receiver’s cell as an attributed string.
//   - [NSControl.SetAttributedStringValue]
//
// # Interacting with Other Controls
//
//   - [NSControl.TakeDoubleValueFrom]: Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//   - [NSControl.TakeFloatValueFrom]: Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//   - [NSControl.TakeIntValueFrom]: Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//   - [NSControl.TakeIntegerValueFrom]: Sets the value of the receiver’s cell to an [NSInteger] value obtained from the specified object.
//   - [NSControl.TakeObjectValueFrom]: Sets the value of the receiver’s cell to the object value obtained from the specified object.
//   - [NSControl.TakeStringValueFrom]: Sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// # Formatting Text
//
//   - [NSControl.Alignment]: The alignment mode of the text in the receiver’s cell.
//   - [NSControl.SetAlignment]
//   - [NSControl.Font]: The font used to draw text in the receiver’s cell.
//   - [NSControl.SetFont]
//   - [NSControl.LineBreakMode]: The line break mode to use for text in the control’s cell.
//   - [NSControl.SetLineBreakMode]
//   - [NSControl.UsesSingleLineMode]: A Boolean value that indicates whether the text in the control’s cell uses single line mode.
//   - [NSControl.SetUsesSingleLineMode]
//   - [NSControl.Formatter]: The receiver’s formatter.
//   - [NSControl.SetFormatter]
//   - [NSControl.BaseWritingDirection]: The initial writing direction used to determine the actual writing direction for text.
//   - [NSControl.SetBaseWritingDirection]
//
// # Managing Expansion Tool Tips
//
//   - [NSControl.DrawWithExpansionFrameInView]: Performs custom expansion tool tip drawing.
//   - [NSControl.AllowsExpansionToolTips]: A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//   - [NSControl.SetAllowsExpansionToolTips]
//   - [NSControl.ExpansionFrameWithFrame]: The frame in which a tool tip can be displayed, if needed.
//
// # Managing the Field Editor
//
//   - [NSControl.AbortEditing]: Terminates the current editing operation and discards any edited text.
//   - [NSControl.CurrentEditor]: Returns the current field editor for the control.
//   - [NSControl.ValidateEditing]: Validates changes to any user-typed text.
//   - [NSControl.EditWithFrameEditorDelegateEvent]: Begins editing of the receiver’s text using the specified field editor.
//   - [NSControl.EndEditing]: Ends the editing of text in the receiver using the specified field editor.
//   - [NSControl.SelectWithFrameEditorDelegateStartLength]: Selects the specified text range in the receiver’s field editor.
//
// # Resizing the Control
//
//   - [NSControl.ControlSize]: The size of the control.
//   - [NSControl.SetControlSize]
//   - [NSControl.SizeThatFits]: Asks the control to calculate and return the size that best fits the specified size.
//   - [NSControl.SizeToFit]: Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
//
// # Displaying a Cell
//
//   - [NSControl.Highlighted]: A Boolean value that indicates whether the cell is highlighted.
//   - [NSControl.SetHighlighted]
//
// # Implementing the Target-Action Mechanism
//
//   - [NSControl.Action]: The default action-message selector associated with the control.
//   - [NSControl.SetAction]
//   - [NSControl.Target]: The target object that receives action messages from the cell.
//   - [NSControl.SetTarget]
//   - [NSControl.Continuous]: A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//   - [NSControl.SetContinuous]
//   - [NSControl.SendActionTo]: Causes the specified action to be sent to the target.
//   - [NSControl.SendActionOn]: Sets the conditions on which the receiver sends action messages to its target.
//
// # Activating from the Keyboard
//
//   - [NSControl.PerformClick]: Simulates a single mouse click on the receiver.
//   - [NSControl.RefusesFirstResponder]: A Boolean value indicating whether the receiver refuses the first responder role.
//   - [NSControl.SetRefusesFirstResponder]
//
// # Tracking the Mouse
//
//   - [NSControl.IgnoresMultiClick]: A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//   - [NSControl.SetIgnoresMultiClick]
//
// # Supporting Constraint-Based Layout
//
//   - [NSControl.InvalidateIntrinsicContentSizeForCell]: Notifies the control that the intrinsic content size for its cell is no longer valid.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl
type NSControl struct {
	NSView
}

// NSControlFromID constructs a [NSControl] from an objc.ID.
//
// A specialized view, such as a button or text field, that notifies your app
// of relevant events using the target-action design pattern.
func NSControlFromID(id objc.ID) NSControl {
	return NSControl{NSView: NSViewFromID(id)}
}
// NOTE: NSControl adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSControl] class.
//
// # Enabling and Disabling the Control
//
//   - [INSControl.Enabled]: A Boolean value that indicates whether the receiver reacts to mouse events.
//   - [INSControl.SetEnabled]
//
// # Accessing the Control’s Value
//
//   - [INSControl.DoubleValue]: The value of the receiver’s cell as a double-precision floating-point number.
//   - [INSControl.SetDoubleValue]
//   - [INSControl.FloatValue]: The value of the receiver’s cell as a single-precision floating-point number.
//   - [INSControl.SetFloatValue]
//   - [INSControl.IntValue]: The value of the receiver’s cell as an integer.
//   - [INSControl.SetIntValue]
//   - [INSControl.IntegerValue]: The value of the receiver’s cell as an integer value.
//   - [INSControl.SetIntegerValue]
//   - [INSControl.ObjectValue]: The value of the receiver’s cell as an Objective-C object.
//   - [INSControl.SetObjectValue]
//   - [INSControl.StringValue]: The value of the receiver’s cell as an [NSString] object.
//   - [INSControl.SetStringValue]
//   - [INSControl.AttributedStringValue]: The value of the receiver’s cell as an attributed string.
//   - [INSControl.SetAttributedStringValue]
//
// # Interacting with Other Controls
//
//   - [INSControl.TakeDoubleValueFrom]: Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
//   - [INSControl.TakeFloatValueFrom]: Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
//   - [INSControl.TakeIntValueFrom]: Sets the value of the receiver’s cell to an integer value obtained from the specified object.
//   - [INSControl.TakeIntegerValueFrom]: Sets the value of the receiver’s cell to an [NSInteger] value obtained from the specified object.
//   - [INSControl.TakeObjectValueFrom]: Sets the value of the receiver’s cell to the object value obtained from the specified object.
//   - [INSControl.TakeStringValueFrom]: Sets the value of the receiver’s cell to the string value obtained from the specified object.
//
// # Formatting Text
//
//   - [INSControl.Alignment]: The alignment mode of the text in the receiver’s cell.
//   - [INSControl.SetAlignment]
//   - [INSControl.Font]: The font used to draw text in the receiver’s cell.
//   - [INSControl.SetFont]
//   - [INSControl.LineBreakMode]: The line break mode to use for text in the control’s cell.
//   - [INSControl.SetLineBreakMode]
//   - [INSControl.UsesSingleLineMode]: A Boolean value that indicates whether the text in the control’s cell uses single line mode.
//   - [INSControl.SetUsesSingleLineMode]
//   - [INSControl.Formatter]: The receiver’s formatter.
//   - [INSControl.SetFormatter]
//   - [INSControl.BaseWritingDirection]: The initial writing direction used to determine the actual writing direction for text.
//   - [INSControl.SetBaseWritingDirection]
//
// # Managing Expansion Tool Tips
//
//   - [INSControl.DrawWithExpansionFrameInView]: Performs custom expansion tool tip drawing.
//   - [INSControl.AllowsExpansionToolTips]: A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
//   - [INSControl.SetAllowsExpansionToolTips]
//   - [INSControl.ExpansionFrameWithFrame]: The frame in which a tool tip can be displayed, if needed.
//
// # Managing the Field Editor
//
//   - [INSControl.AbortEditing]: Terminates the current editing operation and discards any edited text.
//   - [INSControl.CurrentEditor]: Returns the current field editor for the control.
//   - [INSControl.ValidateEditing]: Validates changes to any user-typed text.
//   - [INSControl.EditWithFrameEditorDelegateEvent]: Begins editing of the receiver’s text using the specified field editor.
//   - [INSControl.EndEditing]: Ends the editing of text in the receiver using the specified field editor.
//   - [INSControl.SelectWithFrameEditorDelegateStartLength]: Selects the specified text range in the receiver’s field editor.
//
// # Resizing the Control
//
//   - [INSControl.ControlSize]: The size of the control.
//   - [INSControl.SetControlSize]
//   - [INSControl.SizeThatFits]: Asks the control to calculate and return the size that best fits the specified size.
//   - [INSControl.SizeToFit]: Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
//
// # Displaying a Cell
//
//   - [INSControl.Highlighted]: A Boolean value that indicates whether the cell is highlighted.
//   - [INSControl.SetHighlighted]
//
// # Implementing the Target-Action Mechanism
//
//   - [INSControl.Action]: The default action-message selector associated with the control.
//   - [INSControl.SetAction]
//   - [INSControl.Target]: The target object that receives action messages from the cell.
//   - [INSControl.SetTarget]
//   - [INSControl.Continuous]: A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//   - [INSControl.SetContinuous]
//   - [INSControl.SendActionTo]: Causes the specified action to be sent to the target.
//   - [INSControl.SendActionOn]: Sets the conditions on which the receiver sends action messages to its target.
//
// # Activating from the Keyboard
//
//   - [INSControl.PerformClick]: Simulates a single mouse click on the receiver.
//   - [INSControl.RefusesFirstResponder]: A Boolean value indicating whether the receiver refuses the first responder role.
//   - [INSControl.SetRefusesFirstResponder]
//
// # Tracking the Mouse
//
//   - [INSControl.IgnoresMultiClick]: A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
//   - [INSControl.SetIgnoresMultiClick]
//
// # Supporting Constraint-Based Layout
//
//   - [INSControl.InvalidateIntrinsicContentSizeForCell]: Notifies the control that the intrinsic content size for its cell is no longer valid.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl
type INSControl interface {
	INSView

	// Topic: Enabling and Disabling the Control

	// A Boolean value that indicates whether the receiver reacts to mouse events.
	Enabled() bool
	SetEnabled(value bool)

	// Topic: Accessing the Control’s Value

	// The value of the receiver’s cell as a double-precision floating-point number.
	DoubleValue() float64
	SetDoubleValue(value float64)
	// The value of the receiver’s cell as a single-precision floating-point number.
	FloatValue() float32
	SetFloatValue(value float32)
	// The value of the receiver’s cell as an integer.
	IntValue() int
	SetIntValue(value int)
	// The value of the receiver’s cell as an integer value.
	IntegerValue() int
	SetIntegerValue(value int)
	// The value of the receiver’s cell as an Objective-C object.
	ObjectValue() objectivec.IObject
	SetObjectValue(value objectivec.IObject)
	// The value of the receiver’s cell as an [NSString] object.
	StringValue() string
	SetStringValue(value string)
	// The value of the receiver’s cell as an attributed string.
	AttributedStringValue() foundation.NSAttributedString
	SetAttributedStringValue(value foundation.NSAttributedString)

	// Topic: Interacting with Other Controls

	// Sets the value of the receiver’s cell to a double-precision floating-point value obtained from the specified object.
	TakeDoubleValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to a single-precision floating-point value obtained from the specified object.
	TakeFloatValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to an integer value obtained from the specified object.
	TakeIntValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to an [NSInteger] value obtained from the specified object.
	TakeIntegerValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to the object value obtained from the specified object.
	TakeObjectValueFrom(sender objectivec.IObject)
	// Sets the value of the receiver’s cell to the string value obtained from the specified object.
	TakeStringValueFrom(sender objectivec.IObject)

	// Topic: Formatting Text

	// The alignment mode of the text in the receiver’s cell.
	Alignment() NSTextAlignment
	SetAlignment(value NSTextAlignment)
	// The font used to draw text in the receiver’s cell.
	Font() NSFont
	SetFont(value NSFont)
	// The line break mode to use for text in the control’s cell.
	LineBreakMode() NSLineBreakMode
	SetLineBreakMode(value NSLineBreakMode)
	// A Boolean value that indicates whether the text in the control’s cell uses single line mode.
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	// The receiver’s formatter.
	Formatter() foundation.NSFormatter
	SetFormatter(value foundation.NSFormatter)
	// The initial writing direction used to determine the actual writing direction for text.
	BaseWritingDirection() NSWritingDirection
	SetBaseWritingDirection(value NSWritingDirection)

	// Topic: Managing Expansion Tool Tips

	// Performs custom expansion tool tip drawing.
	DrawWithExpansionFrameInView(contentFrame corefoundation.CGRect, view INSView)
	// A Boolean value that indicates whether expansion tool tips are shown when the control is hovered over.
	AllowsExpansionToolTips() bool
	SetAllowsExpansionToolTips(value bool)
	// The frame in which a tool tip can be displayed, if needed.
	ExpansionFrameWithFrame(contentFrame corefoundation.CGRect) corefoundation.CGRect

	// Topic: Managing the Field Editor

	// Terminates the current editing operation and discards any edited text.
	AbortEditing() bool
	// Returns the current field editor for the control.
	CurrentEditor() INSText
	// Validates changes to any user-typed text.
	ValidateEditing()
	// Begins editing of the receiver’s text using the specified field editor.
	EditWithFrameEditorDelegateEvent(rect corefoundation.CGRect, textObj INSText, delegate objectivec.IObject, event INSEvent)
	// Ends the editing of text in the receiver using the specified field editor.
	EndEditing(textObj INSText)
	// Selects the specified text range in the receiver’s field editor.
	SelectWithFrameEditorDelegateStartLength(rect corefoundation.CGRect, textObj INSText, delegate objectivec.IObject, selStart int, selLength int)

	// Topic: Resizing the Control

	// The size of the control.
	ControlSize() NSControlSize
	SetControlSize(value NSControlSize)
	// Asks the control to calculate and return the size that best fits the specified size.
	SizeThatFits(size corefoundation.CGSize) corefoundation.CGSize
	// Resizes the receiver’s frame so that it’s the minimum size needed to contain its cell.
	SizeToFit()

	// Topic: Displaying a Cell

	// A Boolean value that indicates whether the cell is highlighted.
	Highlighted() bool
	SetHighlighted(value bool)

	// Topic: Implementing the Target-Action Mechanism

	// The default action-message selector associated with the control.
	Action() objc.SEL
	SetAction(value objc.SEL)
	// The target object that receives action messages from the cell.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
	Continuous() bool
	SetContinuous(value bool)
	// Causes the specified action to be sent to the target.
	SendActionTo(action objc.SEL, target objectivec.IObject) bool
	// Sets the conditions on which the receiver sends action messages to its target.
	SendActionOn(mask NSEventMask) int

	// Topic: Activating from the Keyboard

	// Simulates a single mouse click on the receiver.
	PerformClick(sender objectivec.IObject)
	// A Boolean value indicating whether the receiver refuses the first responder role.
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)

	// Topic: Tracking the Mouse

	// A Boolean value indicating whether the receiver ignores multiple clicks made in rapid succession.
	IgnoresMultiClick() bool
	SetIgnoresMultiClick(value bool)

	// Topic: Supporting Constraint-Based Layout

	// Notifies the control that the intrinsic content size for its cell is no longer valid.
	InvalidateIntrinsicContentSizeForCell(cell INSCell)

	// The receiver’s cell object.
	Cell() INSCell
	SetCell(value INSCell)
	// Draws the specified cell, as long as it belongs to the receiver.
	DrawCell(cell INSCell)
	// Draws the inside of the receiver’s cell (the area within the bezel or border)
	DrawCellInside(cell INSCell)
	// Selects the specified cell and redraws the control as needed.
	SelectCell(cell INSCell)
	// Returns the receiver’s selected cell.
	SelectedCell() INSCell
	// Returns the tag of the receiver’s selected cell.
	SelectedTag() int
	// Marks the specified cell as in need of redrawing.
	UpdateCell(cell INSCell)
	// Marks the inside of the specified cell as in need of redrawing.
	UpdateCellInside(cell INSCell)
}

// Init initializes the instance.
func (c NSControl) Init() NSControl {
	rv := objc.Send[NSControl](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSControl) Autorelease() NSControl {
	rv := objc.Send[NSControl](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSControl creates a new NSControl instance.
func NewNSControl() NSControl {
	class := getNSControlClass()
	rv := objc.Send[NSControl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewControlWithCoder(coder foundation.INSCoder) NSControl {
	instance := getNSControlClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSControlFromID(rv)
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
func NewControlWithFrame(frameRect corefoundation.CGRect) NSControl {
	instance := getNSControlClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSControlFromID(rv)
}

// Sets the value of the receiver’s cell to a double-precision
// floating-point value obtained from the specified object.
//
// sender: The object from which to take the value. This object must respond to the
// [DoubleValue] property.
//
// # Discussion
// 
// You can use this method to link action messages between controls. It
// permits one control or cell (`sender`) to affect the value of another
// control (the receiver) by invoking this method in an action message to the
// receiver. For example, a text field can be made the target of a slider.
// Whenever the slider is moved, it sends this message to the text field. The
// text field then obtains the slider’s value, turns it into a text string,
// and displays it.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/takeDoubleValueFrom(_:)
func (c NSControl) TakeDoubleValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeDoubleValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to a single-precision
// floating-point value obtained from the specified object.
//
// sender: The object from which to take the value. This object must respond to the
// [FloatValue] property.
//
// # Discussion
// 
// You can use this method to link action messages between controls. It
// permits one control or cell (`sender`) to affect the value of another
// control (the receiver) by invoking this method in an action message to the
// receiver. For example, a text field can be made the target of a slider.
// Whenever the slider is moved, it sends this message to the text field. The
// text field then obtains the slider’s value, turns it into a text string,
// and displays it.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/takeFloatValueFrom(_:)
func (c NSControl) TakeFloatValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeFloatValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an integer value obtained from
// the specified object.
//
// sender: The object from which to take the value. This object must respond to the
// [IntValue] property.
//
// # Discussion
// 
// You can use this method to link action messages between controls. It
// permits one control or cell (`sender`) to affect the value of another
// control (the receiver) by invoking this method in an action message to the
// receiver. For example, a text field can be made the target of a slider.
// Whenever the slider is moved, it sends this message to the text field. The
// text field then obtains the slider’s value, turns it into a text string,
// and displays it.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/takeIntValueFrom(_:)
func (c NSControl) TakeIntValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeIntValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to an [NSInteger] value obtained
// from the specified object.
//
// sender: The object from which to take the value. This object must respond to the
// [FloatValue] message.
//
// # Discussion
// 
// You can use this method to link action messages between controls. It
// permits one control or cell (`sender`) to affect the value of another
// control (the receiver) by invoking this method in an action message to the
// receiver. For example, a text field can be made the target of a slider.
// Whenever the slider is moved, it sends this message to the text field. The
// text field then obtains the slider’s value, turns it into a text string,
// and displays it.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/takeIntegerValueFrom(_:)
func (c NSControl) TakeIntegerValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeIntegerValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to the object value obtained from
// the specified object.
//
// sender: The object from which to take the value. This object must respond to the
// [ObjectValue] property.
//
// # Discussion
// 
// You can use this method to link action messages between controls. It
// permits one control or cell (`sender`) to affect the value of another
// control (the receiver) by invoking this method in an action message to the
// receiver. For example, a text field can be made the target of a slider.
// Whenever the slider is moved, it sends this message to the text field. The
// text field then obtains the slider’s value, turns it into a text string,
// and displays it.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/takeObjectValueFrom(_:)
func (c NSControl) TakeObjectValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeObjectValueFrom:"), sender)
}

// Sets the value of the receiver’s cell to the string value obtained from
// the specified object.
//
// sender: The object from which to take the value. This object must respond to the
// [StringValue] property.
//
// # Discussion
// 
// You can use this method to link action messages between controls. It
// permits one control or cell (`sender`) to affect the value of another
// control (the receiver) by invoking this method in an action message to the
// receiver. For example, a text field can be made the target of a slider.
// Whenever the slider is moved, it sends this message to the text field. The
// text field then obtains the slider’s value, turns it into a text string,
// and displays it.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/takeStringValueFrom(_:)
func (c NSControl) TakeStringValueFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("takeStringValueFrom:"), sender)
}

// Performs custom expansion tool tip drawing.
//
// contentFrame: The frame in which to draw.
//
// view: The view in which to draw.
//
// # Discussion
// 
// Note that the view may be different from the original view in which the
// text appeared.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/draw(withExpansionFrame:in:)
func (c NSControl) DrawWithExpansionFrameInView(contentFrame corefoundation.CGRect, view INSView) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawWithExpansionFrame:inView:"), contentFrame, view)
}

// The frame in which a tool tip can be displayed, if needed.
//
// contentFrame: The frame of the control.
//
// # Return Value
// 
// The frame in which the tool tip should be displayed, or [NSZeroRect] by
// default.
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// # Discussion
// 
// This method lets the control return an expansion tool tip frame if
// `contentFrame` is too small for the entire contents in the view. When the
// pointer hovers over the text in certain controls, the full contents will be
// shown in a special floating tool tip view. If the frame is big enough to
// display the contents, return an empty rect from this method and no
// expansion tool tip view will be shown. Note that some subclasses, such as
// [NSTextField], return the proper frame when required.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/expansionFrame(withFrame:)
func (c NSControl) ExpansionFrameWithFrame(contentFrame corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("expansionFrameWithFrame:"), contentFrame)
	return corefoundation.CGRect(rv)
}

// Terminates the current editing operation and discards any edited text.
//
// # Return Value
// 
// [true] if there was a field editor associated with the control; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If there was a field editor, this method removes the field editor’s
// delegate.
// 
// Because the control discards any edits, it doesn’t call
// [ControlTextDidEndEditing].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/abortEditing()
func (c NSControl) AbortEditing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("abortEditing"))
	return rv
}

// Returns the current field editor for the control.
//
// # Return Value
// 
// The field editor for the current control, or `nil` if the receiver does not
// have a field editor.
//
// # Discussion
// 
// When the receiver is a control displaying editable text (for example, a
// text field) and it is the first responder, it has a field editor, which is
// returned by this method. The field editor is a single [NSTextView] object
// that is shared among all the controls in a window for light text-editing
// needs. It is automatically instantiated when needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/currentEditor()
func (c NSControl) CurrentEditor() INSText {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("currentEditor"))
	return NSTextFromID(rv)
}

// Validates changes to any user-typed text.
//
// # Discussion
// 
// Validation sets the object value of the cell to the current contents of the
// cell’s editor (the [NSText] object used for editing), storing it as a
// simple [NSString] or an attributed string object based on the attributes of
// the editor.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/validateEditing()
func (c NSControl) ValidateEditing() {
	objc.Send[objc.ID](c.ID, objc.Sel("validateEditing"))
}

// Begins editing of the receiver’s text using the specified field editor.
//
// rect: The bounding rectangle of the control’s cell.
//
// textObj: The field editor to use.
//
// delegate: The object to use as a delegate for the field editor. This delegate object
// receives various [NSText] delegation and notification methods during the
// course of editing the cell’s contents.
//
// event: The [NSLeftMouseDown] event that initiated the editing behavior.
// //
// [NSLeftMouseDown]: https://developer.apple.com/documentation/AppKit/NSLeftMouseDown
//
// # Discussion
// 
// For a receiver that is a control with editable text (such as an
// [NSTextField] object), the field editor is sized to `aRect` and is then
// activated and editing begins. It’s the responsibility of the delegate to
// end editing when responding to [ControlTextShouldEndEditing]. Upon ending
// the editing session, the delegate should remove any data from the field
// editor.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/edit(withFrame:editor:delegate:event:)
func (c NSControl) EditWithFrameEditorDelegateEvent(rect corefoundation.CGRect, textObj INSText, delegate objectivec.IObject, event INSEvent) {
	objc.Send[objc.ID](c.ID, objc.Sel("editWithFrame:editor:delegate:event:"), rect, textObj, delegate, event)
}

// Ends the editing of text in the receiver using the specified field editor.
//
// textObj: The field editor currently handling the editing of the cell’s content.
//
// # Discussion
// 
// Ends any editing of text that began with a call to
// [EditWithFrameEditorDelegateEvent] or
// [SelectWithFrameEditorDelegateStartLength].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/endEditing(_:)
func (c NSControl) EndEditing(textObj INSText) {
	objc.Send[objc.ID](c.ID, objc.Sel("endEditing:"), textObj)
}

// Selects the specified text range in the receiver’s field editor.
//
// rect: The bounding rectangle of the control’s cell.
//
// textObj: The field editor to use.
//
// delegate: The object to use as a delegate for the field editor. This delegate object
// receives various [NSText] delegation and notification methods during the
// course of editing the cell’s contents.
//
// selStart: The start of the text selection.
//
// selLength: The length of the text range.
//
// # Discussion
// 
// This method is similar to [EditWithFrameEditorDelegateEvent], except that
// it can be invoked in any situation, not only on a mouse-down event. This
// method returns without doing anything if `textObj` or the receiver is
// `nil`, or if the receiver has no font set for it.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/select(withFrame:editor:delegate:start:length:)
func (c NSControl) SelectWithFrameEditorDelegateStartLength(rect corefoundation.CGRect, textObj INSText, delegate objectivec.IObject, selStart int, selLength int) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectWithFrame:editor:delegate:start:length:"), rect, textObj, delegate, selStart, selLength)
}

// Asks the control to calculate and return the size that best fits the
// specified size.
//
// size: The size for which the control should calculate its best-fitting size.
//
// # Return Value
// 
// A new size that fits the receiver’s subviews.
//
// # Discussion
// 
// By default, this method returns the [intrinsicContentSize] of the receiver.
//
// [intrinsicContentSize]: https://developer.apple.com/documentation/AppKit/NSView/intrinsicContentSize
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/sizeThatFits(_:)
func (c NSControl) SizeThatFits(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("sizeThatFits:"), size)
	return corefoundation.CGSize(rv)
}

// Resizes the receiver’s frame so that it’s the minimum size needed to
// contain its cell.
//
// # Discussion
// 
// If you want a multiple-cell custom subclass of [NSControl] to size itself
// to fit its cells, you must override this method. This method neither
// redisplays the receiver nor marks it as needing display. You must do this
// yourself with either the[Display] or [setNeedsDisplay()] method.
//
// [setNeedsDisplay()]: https://developer.apple.com/documentation/AppKit/NSControl/setNeedsDisplay()
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/sizeToFit()
func (c NSControl) SizeToFit() {
	objc.Send[objc.ID](c.ID, objc.Sel("sizeToFit"))
}

// Causes the specified action to be sent to the target.
//
// action: The selector to invoke on the target. If the selector is [NULL], no message
// is sent.
//
// target: The target object to receive the message. If the object is `nil`, the
// application searches the responder chain for an object capable of handling
// the message. For more information on dispatching actions, see the class
// description for [NSActionCell].
//
// # Return Value
// 
// [true] if the message was successfully sent; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method uses the [SendActionToFrom] method of [NSApplication] to invoke
// the specified method on an object. The receiver is passed as the parameter
// to the action message. This method is invoked primarily by the
// [TrackMouseInRectOfViewUntilMouseUp] method of [NSCell].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(_:to:)
func (c NSControl) SendActionTo(action objc.SEL, target objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("sendAction:to:"), action, target)
	return rv
}

// Sets the conditions on which the receiver sends action messages to its
// target.
//
// mask: A bit mask containing the conditions for sending the action. The only
// conditions that are actually checked are associated with the
// [NSLeftMouseDownMask], [NSLeftMouseUpMask], [NSLeftMouseDraggedMask], and
// [NSPeriodicMask] bits.
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
// The default implementation of this method simply invokes the [SendActionOn]
// method of its associated cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/sendAction(on:)
func (c NSControl) SendActionOn(mask NSEventMask) int {
	rv := objc.Send[int](c.ID, objc.Sel("sendActionOn:"), mask)
	return rv
}

// Simulates a single mouse click on the receiver.
//
// sender: The object requesting the action. This parameter is ignored.
//
// # Discussion
// 
// This method calls the [PerformClick] method of the receiver’s cell with
// the sender being the control itself. This method raises an exception if the
// action message cannot be successfully sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/performClick(_:)
func (c NSControl) PerformClick(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("performClick:"), sender)
}

// Notifies the control that the intrinsic content size for its cell is no
// longer valid.
//
// cell: The cell whose intrinsic content size has changed.
//
// # Discussion
// 
// Controls determine their intrinsic content size based on the cell size for
// a given bounds returned by their cell. When the content of the cell changes
// in a way that would change the return value of [CellSizeForBounds], the
// cell needs to call this method to notify its control that its intrinsic
// size is no longer valid.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/invalidateIntrinsicContentSize(for:)
func (c NSControl) InvalidateIntrinsicContentSizeForCell(cell INSCell) {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateIntrinsicContentSizeForCell:"), cell)
}

// Draws the specified cell, as long as it belongs to the receiver.
//
// cell: The cell to draw. If the cell does not belong to the receiver, this method
// does nothing.
//
// # Discussion
// 
// This method is provided primarily to support a consistent set of methods
// between [NSControl] objects with single and multiple cells, because a
// control with multiple cells needs to be able to draw individual cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/drawCell(_:)
func (c NSControl) DrawCell(cell INSCell) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawCell:"), cell)
}

// Draws the inside of the receiver’s cell (the area within the bezel or
// border)
//
// cell: The cell to draw. If the cell does not belong to the receiver, this method
// does nothing.
//
// # Discussion
// 
// If the receiver is transparent, the method causes the superview to draw
// itself. This method invokes the [DrawInteriorWithFrameInView] method of
// NSCell. This method has no effect on controls (such as [NSMatrix] and
// [NSForm]) that have multiple cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/drawCellInside(_:)
func (c NSControl) DrawCellInside(cell INSCell) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawCellInside:"), cell)
}

// Selects the specified cell and redraws the control as needed.
//
// cell: The cell to select. The cell must belong to the receiver.
//
// # Discussion
// 
// If the cell is already selected (or does not belong to the receiver), this
// method does nothing. If the cell belongs to the receiver and is not
// selected, this method changes its state to [NSOnState] and redraws the
// cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/selectCell(_:)
func (c NSControl) SelectCell(cell INSCell) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectCell:"), cell)
}

// Returns the receiver’s selected cell.
//
// # Return Value
// 
// The selected cell object.
//
// # Discussion
// 
// The default implementation of this method simply returns the control’s
// associated cell (or `nil` if no cell has been set). Subclasses of
// [NSControl] that manage multiple cells (such as [NSMatrix] and [NSForm])
// must override this method to return the cell selected by the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/selectedCell()
func (c NSControl) SelectedCell() INSCell {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("selectedCell"))
	return NSCellFromID(rv)
}

// Returns the tag of the receiver’s selected cell.
//
// # Return Value
// 
// The tag of the selected cell, or `-1` if no cell is selected.
//
// # Discussion
// 
// When you set the tag of a control with a single cell in Interface Builder,
// it sets the tags of both the control and the cell with the same value as a
// convenience.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/selectedTag()
func (c NSControl) SelectedTag() int {
	rv := objc.Send[int](c.ID, objc.Sel("selectedTag"))
	return rv
}

// Marks the specified cell as in need of redrawing.
//
// cell: The cell to redraw.
//
// # Discussion
// 
// If the cell currently has the focus, this method updates the cell’s focus
// ring; otherwise, the entire cell is marked as needing redisplay. The cell
// is redrawn during the next update cycle.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/updateCell(_:)
func (c NSControl) UpdateCell(cell INSCell) {
	objc.Send[objc.ID](c.ID, objc.Sel("updateCell:"), cell)
}

// Marks the inside of the specified cell as in need of redrawing.
//
// cell: The cell to redraw.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/updateCellInside(_:)
func (c NSControl) UpdateCellInside(cell INSCell) {
	objc.Send[objc.ID](c.ID, objc.Sel("updateCellInside:"), cell)
}

// A Boolean value that indicates whether the receiver reacts to mouse events.
//
// # Discussion
// 
// The value of this property is [true] if the receiver responds to mouse
// events; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/isEnabled
func (c NSControl) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c NSControl) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}

// The value of the receiver’s cell as a double-precision floating-point
// number.
//
// # Discussion
// 
// If the control contains many cells (for example, [NSMatrix]), then this
// property contains the value of the currently selected cell. If the control
// is in the process of editing the affected cell, then it invokes the
// [ValidateEditing] method before getting the value.
// 
// If the cell is being edited, setting this property aborts all editing
// before setting the value. If the cell does not inherit from [NSActionCell],
// setting this property marks the cell’s interior as needing to be
// redisplayed; [NSActionCell] performs its own updating of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/doubleValue
func (c NSControl) DoubleValue() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("doubleValue"))
	return rv
}
func (c NSControl) SetDoubleValue(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setDoubleValue:"), value)
}

// The value of the receiver’s cell as a single-precision floating-point
// number.
//
// # Discussion
// 
// If the control contains many cells (for example, [NSMatrix]), then this
// property contains the value of the currently selected cell. If the control
// is in the process of editing the affected cell, then it invokes the
// [ValidateEditing] method before getting the value.
// 
// If the cell is being edited, setting this property aborts all editing
// before setting the value. If the cell does not inherit from [NSActionCell],
// setting this property marks the cell’s interior as needing to be
// redisplayed; [NSActionCell] performs its own updating of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/floatValue
func (c NSControl) FloatValue() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("floatValue"))
	return rv
}
func (c NSControl) SetFloatValue(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setFloatValue:"), value)
}

// The value of the receiver’s cell as an integer.
//
// # Discussion
// 
// If the control contains many cells (for example, [NSMatrix]), then this
// property contains the value of the currently selected cell. If the control
// is in the process of editing the affected cell, then it invokes the
// [ValidateEditing] method before getting the value.
// 
// If the cell is being edited, setting this property aborts all editing
// before setting the value. If the cell does not inherit from [NSActionCell],
// setting this property marks the cell’s interior as needing to be
// redisplayed; [NSActionCell] performs its own updating of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/intValue
func (c NSControl) IntValue() int {
	rv := objc.Send[int](c.ID, objc.Sel("intValue"))
	return rv
}
func (c NSControl) SetIntValue(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setIntValue:"), value)
}

// The value of the receiver’s cell as an integer value.
//
// # Discussion
// 
// If the control contains many cells (for example, [NSMatrix]), then this
// property contains the value of the currently selected cell. If the control
// is in the process of editing the affected cell, then it invokes the
// [ValidateEditing] method before getting the value.
// 
// If the cell is being edited, setting this property aborts all editing
// before setting the value. If the cell does not inherit from [NSActionCell],
// setting this property marks the cell’s interior as needing to be
// redisplayed; [NSActionCell] performs its own updating of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/integerValue
func (c NSControl) IntegerValue() int {
	rv := objc.Send[int](c.ID, objc.Sel("integerValue"))
	return rv
}
func (c NSControl) SetIntegerValue(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setIntegerValue:"), value)
}

// The value of the receiver’s cell as an Objective-C object.
//
// # Discussion
// 
// If the control contains many cells (for example, [NSMatrix]), then this
// property contains the value of the currently selected cell. If the control
// is in the process of editing the affected cell, then it invokes the
// [ValidateEditing] method before getting the value.
// 
// If the cell is being edited, setting this property aborts all editing
// before setting the value. If the cell does not inherit from [NSActionCell],
// setting this property marks the cell’s interior as needing to be
// redisplayed; [NSActionCell] performs its own updating of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/objectValue
func (c NSControl) ObjectValue() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectValue"))
	return objectivec.Object{ID: rv}
}
func (c NSControl) SetObjectValue(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setObjectValue:"), value)
}

// The value of the receiver’s cell as an [NSString] object.
//
// # Discussion
// 
// If the control contains many cells (for example, [NSMatrix]), then this
// property contains the value of the currently selected cell. If the control
// is in the process of editing the affected cell, then it invokes the
// [ValidateEditing] method before getting the value.
// 
// If the cell is being edited, setting this property aborts all editing
// before setting the value. If the cell does not inherit from [NSActionCell],
// setting this property marks the cell’s interior as needing to be
// redisplayed; [NSActionCell] performs its own updating of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/stringValue
func (c NSControl) StringValue() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (c NSControl) SetStringValue(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setStringValue:"), objc.String(value))
}

// The value of the receiver’s cell as an attributed string.
//
// # Discussion
// 
// If the control contains many cells (for example, [NSMatrix]), this property
// contains the value of the currently selected cell. If the control is in the
// process of editing the affected cell, then it invokes the [ValidateEditing]
// method before getting the value.
// 
// If the cell is being edited, setting this property aborts all editing
// before setting the value. If the cell does not inherit from [NSActionCell],
// setting the property marks the cell’s interior as needing to be
// redisplayed; [NSActionCell] performs its own updating of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/attributedStringValue
func (c NSControl) AttributedStringValue() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("attributedStringValue"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (c NSControl) SetAttributedStringValue(value foundation.NSAttributedString) {
	objc.Send[struct{}](c.ID, objc.Sel("setAttributedStringValue:"), value)
}

// The alignment mode of the text in the receiver’s cell.
//
// # Discussion
// 
// The value of this property can be one of the following constants:
// [NSLeftTextAlignment], [NSRightTextAlignment],[NSCenterTextAlignment],
// [NSJustifiedTextAlignment], or [NSNaturalTextAlignment]. The default value
// is [NSNaturalTextAlignment]. Setting this property while the cell is
// currently being edited aborts the edits to change the alignment.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/alignment
func (c NSControl) Alignment() NSTextAlignment {
	rv := objc.Send[NSTextAlignment](c.ID, objc.Sel("alignment"))
	return NSTextAlignment(rv)
}
func (c NSControl) SetAlignment(value NSTextAlignment) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlignment:"), value)
}

// The font used to draw text in the receiver’s cell.
//
// # Discussion
// 
// If the cell is being edited, setting this property causes the text in the
// cell to be redrawn in the new font, and the cell’s editor (the [NSText]
// object used globally for editing) is updated with the new font object.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/font
func (c NSControl) Font() NSFont {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("font"))
	return NSFontFromID(objc.ID(rv))
}
func (c NSControl) SetFont(value NSFont) {
	objc.Send[struct{}](c.ID, objc.Sel("setFont:"), value)
}

// The line break mode to use for text in the control’s cell.
//
// # Discussion
// 
// For a list of possible values, see [NSLineBreakMode].
//
// [NSLineBreakMode]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/lineBreakMode
func (c NSControl) LineBreakMode() NSLineBreakMode {
	rv := objc.Send[NSLineBreakMode](c.ID, objc.Sel("lineBreakMode"))
	return NSLineBreakMode(rv)
}
func (c NSControl) SetLineBreakMode(value NSLineBreakMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setLineBreakMode:"), value)
}

// A Boolean value that indicates whether the text in the control’s cell
// uses single line mode.
//
// # Discussion
// 
// See [UsesSingleLineMode] for details.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/usesSingleLineMode
func (c NSControl) UsesSingleLineMode() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("usesSingleLineMode"))
	return rv
}
func (c NSControl) SetUsesSingleLineMode(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setUsesSingleLineMode:"), value)
}

// The receiver’s formatter.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/formatter
func (c NSControl) Formatter() foundation.NSFormatter {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("formatter"))
	return foundation.NSFormatterFromID(objc.ID(rv))
}
func (c NSControl) SetFormatter(value foundation.NSFormatter) {
	objc.Send[struct{}](c.ID, objc.Sel("setFormatter:"), value)
}

// The initial writing direction used to determine the actual writing
// direction for text.
//
// # Discussion
// 
// This property can have one of the following values:
// [NSWritingDirectionNatural], [NSWritingDirectionLeftToRight], or
// [NSWritingDirectionRightToLeft]. The default value is
// [NSWritingDirectionNatural]. The text system uses this value as a hint for
// calculating the actual direction for displaying Unicode characters. You
// should not need to access this value directly. If you know the base writing
// direction of the text you are rendering, you can set this property to
// specify that direction to the text system.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/baseWritingDirection
func (c NSControl) BaseWritingDirection() NSWritingDirection {
	rv := objc.Send[NSWritingDirection](c.ID, objc.Sel("baseWritingDirection"))
	return NSWritingDirection(rv)
}
func (c NSControl) SetBaseWritingDirection(value NSWritingDirection) {
	objc.Send[struct{}](c.ID, objc.Sel("setBaseWritingDirection:"), value)
}

// A Boolean value that indicates whether expansion tool tips are shown when
// the control is hovered over.
//
// # Discussion
// 
// When the value of this property is [true], the expansion tool tip will
// expand; [false] means the tool tip won’t expand. The default value is
// [false].
// 
// Expansion tooltips are shown when the cell cannot show the full content and
// the user hovers the pointer over the control. This is controlled by the
// [NSCell] class method [ExpansionFrameWithFrameInView] and is drawn by
// [DrawWithExpansionFrameInView]. This value is encoded along with the
// control.
// 
// In general, it is recommended to turn this on for [NSTextField] instances
// in a view-based [NSTableView].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/allowsExpansionToolTips
func (c NSControl) AllowsExpansionToolTips() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsExpansionToolTips"))
	return rv
}
func (c NSControl) SetAllowsExpansionToolTips(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsExpansionToolTips:"), value)
}

// The size of the control.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/controlSize-swift.property
func (c NSControl) ControlSize() NSControlSize {
	rv := objc.Send[NSControlSize](c.ID, objc.Sel("controlSize"))
	return NSControlSize(rv)
}
func (c NSControl) SetControlSize(value NSControlSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setControlSize:"), value)
}

// A Boolean value that indicates whether the cell is highlighted.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/isHighlighted
func (c NSControl) Highlighted() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isHighlighted"))
	return rv
}
func (c NSControl) SetHighlighted(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHighlighted:"), value)
}

// The default action-message selector associated with the control.
//
// # Discussion
// 
// This property contains the action message selector of the receiver’s
// cell. Controls that support multiple cells (such as [NSMatrix] and
// [NSForm]) must supply the appropriate action-message selector in this
// property. Specify [NULL] to prevent action messages from being sent to the
// receiver’s target.
// 
// If you want the action-message selector for a control that has multiple
// cells, it is better to get the selector directly from the cell’s own
// `action` property.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/action
func (c NSControl) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c.ID, objc.Sel("action"))
	return rv
}
func (c NSControl) SetAction(value objc.SEL) {
	objc.Send[struct{}](c.ID, objc.Sel("setAction:"), value)
}

// The target object that receives action messages from the cell.
//
// # Discussion
// 
// When the value of this property is `nil`, the application follows the
// responder chain looking for an object that can respond to the message. See
// the description of the [NSActionCell] class for details.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/target
func (c NSControl) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (c NSControl) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setTarget:"), value)
}

// A Boolean value indicating whether the receiver’s cell sends its action
// message continuously to its target during mouse tracking.
//
// # Discussion
// 
// The value of this property is [true] if the action message is sent
// continuously; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/isContinuous
func (c NSControl) Continuous() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isContinuous"))
	return rv
}
func (c NSControl) SetContinuous(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setContinuous:"), value)
}

// A Boolean value indicating whether the receiver refuses the first responder
// role.
//
// # Discussion
// 
// The value of this property is [true] if the receiver refuses the first
// responder role; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/refusesFirstResponder
func (c NSControl) RefusesFirstResponder() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("refusesFirstResponder"))
	return rv
}
func (c NSControl) SetRefusesFirstResponder(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setRefusesFirstResponder:"), value)
}

// A Boolean value indicating whether the receiver ignores multiple clicks
// made in rapid succession.
//
// # Discussion
// 
// The value of this property is [true] if the receiver ignores multiple
// clicks; otherwise, [false].
// 
// By default, controls treat double clicks as two distinct clicks, triple
// clicks as three distinct clicks, and so on. However, if you set this
// propery to [true], additional clicks (within a predetermined interval after
// the first) occurring after the first click are not processed by the
// receiver, and are instead passed on to `super`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/ignoresMultiClick
func (c NSControl) IgnoresMultiClick() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("ignoresMultiClick"))
	return rv
}
func (c NSControl) SetIgnoresMultiClick(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setIgnoresMultiClick:"), value)
}

// The receiver’s cell object.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/cell
func (c NSControl) Cell() INSCell {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("cell"))
	return NSCellFromID(objc.ID(rv))
}
func (c NSControl) SetCell(value INSCell) {
	objc.Send[struct{}](c.ID, objc.Sel("setCell:"), value)
}

// Returns the type of cell used by the receiver.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/cellclass
func (_NSControlClass NSControlClass) CellClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSControlClass.class), objc.Sel("cellClass"))
	return rv
}
func (_NSControlClass NSControlClass) SetCellClass(value objc.Class) {
	objc.Send[struct{}](objc.ID(_NSControlClass.class), objc.Sel("setCellClass:"), value)
}

