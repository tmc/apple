// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSTextAttachmentCell] class.
var (
	_NSTextAttachmentCellClass     NSTextAttachmentCellClass
	_NSTextAttachmentCellClassOnce sync.Once
)

func getNSTextAttachmentCellClass() NSTextAttachmentCellClass {
	_NSTextAttachmentCellClassOnce.Do(func() {
		_NSTextAttachmentCellClass = NSTextAttachmentCellClass{class: objc.GetClass("NSTextAttachmentCell")}
	})
	return _NSTextAttachmentCellClass
}

// GetNSTextAttachmentCellClass returns the class object for NSTextAttachmentCell.
func GetNSTextAttachmentCellClass() NSTextAttachmentCellClass {
	return getNSTextAttachmentCellClass()
}

type NSTextAttachmentCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextAttachmentCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextAttachmentCellClass) Alloc() NSTextAttachmentCell {
	rv := objc.Send[NSTextAttachmentCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that implements the functionality of the text attachment cell
// protocol.
//
// # Overview
//
// This specification describes only those methods whose implementations have
// features that are particular to this class. For a general discussion of the
// protocol’s methods, see [NSTextAttachmentCellProtocol].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCell-swift.class
type NSTextAttachmentCell struct {
	NSCell
}

// NSTextAttachmentCellFromID constructs a [NSTextAttachmentCell] from an objc.ID.
//
// An object that implements the functionality of the text attachment cell
// protocol.
func NSTextAttachmentCellFromID(id objc.ID) NSTextAttachmentCell {
	return NSTextAttachmentCell{NSCell: NSCellFromID(id)}
}

// NOTE: NSTextAttachmentCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextAttachmentCell] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCell-swift.class
type INSTextAttachmentCell interface {
	INSCell
}

// Init initializes the instance.
func (t NSTextAttachmentCell) Init() NSTextAttachmentCell {
	rv := objc.Send[NSTextAttachmentCell](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextAttachmentCell) Autorelease() NSTextAttachmentCell {
	rv := objc.Send[NSTextAttachmentCell](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextAttachmentCell creates a new NSTextAttachmentCell instance.
func NewNSTextAttachmentCell() NSTextAttachmentCell {
	class := getNSTextAttachmentCellClass()
	rv := objc.Send[NSTextAttachmentCell](objc.ID(class.class), objc.Sel("new"))
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
func NewTextAttachmentCellImageCell(image INSImage) NSTextAttachmentCell {
	instance := getNSTextAttachmentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSTextAttachmentCellFromID(rv)
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
func NewTextAttachmentCellTextCell(string_ string) NSTextAttachmentCell {
	instance := getNSTextAttachmentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSTextAttachmentCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewTextAttachmentCellWithCoder(coder foundation.INSCoder) NSTextAttachmentCell {
	instance := getNSTextAttachmentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextAttachmentCellFromID(rv)
}
