// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSActionCell] class.
var (
	_NSActionCellClass     NSActionCellClass
	_NSActionCellClassOnce sync.Once
)

func getNSActionCellClass() NSActionCellClass {
	_NSActionCellClassOnce.Do(func() {
		_NSActionCellClass = NSActionCellClass{class: objc.GetClass("NSActionCell")}
	})
	return _NSActionCellClass
}

// GetNSActionCellClass returns the class object for NSActionCell.
func GetNSActionCellClass() NSActionCellClass {
	return getNSActionCellClass()
}

type NSActionCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSActionCellClass) Alloc() NSActionCell {
	rv := objc.Send[NSActionCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An active area inside a control.
//
// # Overview
// 
// An [NSActionCell] does three things: it displays text or an icon; it
// provides the target object and action method used by its [NSControl]
// object; and it handles mouse (cursor) tracking by properly highlighting its
// area and sending action messages to its target based on cursor movement.
// 
// The [NSActionCell.ControlView] of an [NSActionCell] is the view in which the receiver
// was last drawn.
// 
// # Obtaining and Setting Cell Values
// 
// The [NSActionCell.FloatValue], [NSActionCell.IntValue], and [NSActionCell.IntegerValue] methods return the value
// with their corresponding types after validating any editing of cell
// content. If the cell is not a text-type cell or the cell value is not
// scannable to the appropriate type, these return 0.
// 
// The [NSActionCell.StringValue] method returns the receiver’s value as a string object
// as converted by the cell’s formatter, if one exists. If no formatter
// exists and the value is an [NSString], returns the value as a plain,
// attributed, or localized formatted string. If the value is not an
// [NSString] or cannot be converted to one, returns an empty string. The
// method supplements the [NSCell] implementation by validating and retaining
// any editing changes being made to cell text.
// 
// Calling `` discards any editing of the receiver’s text and sets its
// object value to the specified object. After doing so, if the object value
// is different from what it was before the method was invoked, the method
// marks the receiver as needing redisplay.
// 
// # Configuring an NSActionCell Object
// 
// The [NSActionCell] implementation of [setFloatingPointFormat:left:right:]
// supplements the [NSCell] implementation by marking the receiver as needing
// redisplay after discarding any editing changes that were being made to cell
// text.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [setFloatingPointFormat:left:right:]: https://developer.apple.com/documentation/AppKit/NSCell/setFloatingPointFormat:left:right:
//
// See: https://developer.apple.com/documentation/AppKit/NSActionCell
type NSActionCell struct {
	NSCell
}

// NSActionCellFromID constructs a [NSActionCell] from an objc.ID.
//
// An active area inside a control.
func NSActionCellFromID(id objc.ID) NSActionCell {
	return NSActionCell{NSCell: NSCellFromID(id)}
}
// NOTE: NSActionCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSActionCell] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSActionCell
type INSActionCell interface {
	INSCell

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (a NSActionCell) Init() NSActionCell {
	rv := objc.Send[NSActionCell](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSActionCell) Autorelease() NSActionCell {
	rv := objc.Send[NSActionCell](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSActionCell creates a new NSActionCell instance.
func NewNSActionCell() NSActionCell {
	class := getNSActionCellClass()
	rv := objc.Send[NSActionCell](objc.ID(class.class), objc.Sel("new"))
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
func NewActionCellImageCell(image INSImage) NSActionCell {
	instance := getNSActionCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSActionCellFromID(rv)
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
func NewActionCellTextCell(string_ string) NSActionCell {
	instance := getNSActionCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSActionCellFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewActionCellWithCoder(coder foundation.INSCoder) NSActionCell {
	instance := getNSActionCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSActionCellFromID(rv)
}






func (a NSActionCell) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}







































