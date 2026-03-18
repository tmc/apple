// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSPathComponentCell] class.
var (
	_NSPathComponentCellClass     NSPathComponentCellClass
	_NSPathComponentCellClassOnce sync.Once
)

func getNSPathComponentCellClass() NSPathComponentCellClass {
	_NSPathComponentCellClassOnce.Do(func() {
		_NSPathComponentCellClass = NSPathComponentCellClass{class: objc.GetClass("NSPathComponentCell")}
	})
	return _NSPathComponentCellClass
}

// GetNSPathComponentCellClass returns the class object for NSPathComponentCell.
func GetNSPathComponentCellClass() NSPathComponentCellClass {
	return getNSPathComponentCellClass()
}

type NSPathComponentCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPathComponentCellClass) Alloc() NSPathComponentCell {
	rv := objc.Send[NSPathComponentCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A component of a path.
//
// # Overview
// 
// An [NSPathCell] object manages a collection of [NSPathComponentCell]
// objects, in conjunction with an [NSPathControl] object, to represent a
// path.
//
// # Setting the Path
//
//   - [NSPathComponentCell.URL]: The portion of the path from the root through the component represented by the receiver.
//   - [NSPathComponentCell.SetURL]
//
// See: https://developer.apple.com/documentation/AppKit/NSPathComponentCell
type NSPathComponentCell struct {
	NSTextFieldCell
}

// NSPathComponentCellFromID constructs a [NSPathComponentCell] from an objc.ID.
//
// A component of a path.
func NSPathComponentCellFromID(id objc.ID) NSPathComponentCell {
	return NSPathComponentCell{NSTextFieldCell: NSTextFieldCellFromID(id)}
}
// NOTE: NSPathComponentCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPathComponentCell] class.
//
// # Setting the Path
//
//   - [INSPathComponentCell.URL]: The portion of the path from the root through the component represented by the receiver.
//   - [INSPathComponentCell.SetURL]
//
// See: https://developer.apple.com/documentation/AppKit/NSPathComponentCell
type INSPathComponentCell interface {
	INSTextFieldCell

	// Topic: Setting the Path

	// The portion of the path from the root through the component represented by the receiver.
	URL() foundation.INSURL
	SetURL(value foundation.INSURL)
}

// Init initializes the instance.
func (p NSPathComponentCell) Init() NSPathComponentCell {
	rv := objc.Send[NSPathComponentCell](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPathComponentCell) Autorelease() NSPathComponentCell {
	rv := objc.Send[NSPathComponentCell](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPathComponentCell creates a new NSPathComponentCell instance.
func NewNSPathComponentCell() NSPathComponentCell {
	class := getNSPathComponentCellClass()
	rv := objc.Send[NSPathComponentCell](objc.ID(class.class), objc.Sel("new"))
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
func NewPathComponentCellImageCell(image INSImage) NSPathComponentCell {
	instance := getNSPathComponentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSPathComponentCellFromID(rv)
}

// Initializes a text field cell that displays the specified string.
//
// string: The string that the text field cell displays.
//
// # Return Value
// 
// A text field cell that displays a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(textCell:)
func NewPathComponentCellTextCell(string_ string) NSPathComponentCell {
	instance := getNSPathComponentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSPathComponentCellFromID(rv)
}

// Initializes a text field cell from data in the provided unarchiver.
//
// coder: An unarchiver object.
//
// # Return Value
// 
// A text field cell that displays a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(coder:)
func NewPathComponentCellWithCoder(coder foundation.INSCoder) NSPathComponentCell {
	instance := getNSPathComponentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPathComponentCellFromID(rv)
}

// The portion of the path from the root through the component represented by
// the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSPathComponentCell/url
func (p NSPathComponentCell) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (p NSPathComponentCell) SetURL(value foundation.INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setURL:"), value)
}

