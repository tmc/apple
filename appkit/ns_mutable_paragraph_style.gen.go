// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMutableParagraphStyle] class.
var (
	_NSMutableParagraphStyleClass     NSMutableParagraphStyleClass
	_NSMutableParagraphStyleClassOnce sync.Once
)

func getNSMutableParagraphStyleClass() NSMutableParagraphStyleClass {
	_NSMutableParagraphStyleClassOnce.Do(func() {
		_NSMutableParagraphStyleClass = NSMutableParagraphStyleClass{class: objc.GetClass("NSMutableParagraphStyle")}
	})
	return _NSMutableParagraphStyleClass
}

// GetNSMutableParagraphStyleClass returns the class object for NSMutableParagraphStyle.
func GetNSMutableParagraphStyleClass() NSMutableParagraphStyleClass {
	return getNSMutableParagraphStyleClass()
}

type NSMutableParagraphStyleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMutableParagraphStyleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMutableParagraphStyleClass) Alloc() NSMutableParagraphStyle {
	rv := objc.Send[NSMutableParagraphStyle](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object for changing the values of the subattributes in a paragraph style
// attribute.
//
// # Overview
// 
// The [NSMutableParagraphStyle] class adds methods to its superclass,
// [NSParagraphStyle], for changing the values of the subattributes in a
// paragraph style attribute. For more information, see [NSParagraphStyle] and
// [NSAttributedString].
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// # Specifying tab information
//
//   - [NSMutableParagraphStyle.AddTabStop]: Adds the specified tab stop to the paragraph.
//   - [NSMutableParagraphStyle.RemoveTabStop]: Removes the first text tab with a location and type equal to the specified tab stop.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle
type NSMutableParagraphStyle struct {
	NSParagraphStyle
}

// NSMutableParagraphStyleFromID constructs a [NSMutableParagraphStyle] from an objc.ID.
//
// An object for changing the values of the subattributes in a paragraph style
// attribute.
func NSMutableParagraphStyleFromID(id objc.ID) NSMutableParagraphStyle {
	return NSMutableParagraphStyle{NSParagraphStyle: NSParagraphStyleFromID(id)}
}
// NOTE: NSMutableParagraphStyle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMutableParagraphStyle] class.
//
// # Specifying tab information
//
//   - [INSMutableParagraphStyle.AddTabStop]: Adds the specified tab stop to the paragraph.
//   - [INSMutableParagraphStyle.RemoveTabStop]: Removes the first text tab with a location and type equal to the specified tab stop.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle
type INSMutableParagraphStyle interface {
	INSParagraphStyle

	// Topic: Specifying tab information

	// Adds the specified tab stop to the paragraph.
	AddTabStop(anObject INSTextTab)
	// Removes the first text tab with a location and type equal to the specified tab stop.
	RemoveTabStop(anObject INSTextTab)
}

// Init initializes the instance.
func (m NSMutableParagraphStyle) Init() NSMutableParagraphStyle {
	rv := objc.Send[NSMutableParagraphStyle](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMutableParagraphStyle) Autorelease() NSMutableParagraphStyle {
	rv := objc.Send[NSMutableParagraphStyle](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMutableParagraphStyle creates a new NSMutableParagraphStyle instance.
func NewNSMutableParagraphStyle() NSMutableParagraphStyle {
	class := getNSMutableParagraphStyleClass()
	rv := objc.Send[NSMutableParagraphStyle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds the specified tab stop to the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/addTabStop(_:)
func (m NSMutableParagraphStyle) AddTabStop(anObject INSTextTab) {
	objc.Send[objc.ID](m.ID, objc.Sel("addTabStop:"), anObject)
}
// Removes the first text tab with a location and type equal to the specified
// tab stop.
//
// See: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle/removeTabStop(_:)
func (m NSMutableParagraphStyle) RemoveTabStop(anObject INSTextTab) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTabStop:"), anObject)
}

