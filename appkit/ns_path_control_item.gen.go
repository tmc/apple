// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPathControlItem] class.
var (
	_NSPathControlItemClass     NSPathControlItemClass
	_NSPathControlItemClassOnce sync.Once
)

func getNSPathControlItemClass() NSPathControlItemClass {
	_NSPathControlItemClassOnce.Do(func() {
		_NSPathControlItemClass = NSPathControlItemClass{class: objc.GetClass("NSPathControlItem")}
	})
	return _NSPathControlItemClass
}

// GetNSPathControlItemClass returns the class object for NSPathControlItem.
func GetNSPathControlItemClass() NSPathControlItemClass {
	return getNSPathControlItemClass()
}

type NSPathControlItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPathControlItemClass) Alloc() NSPathControlItem {
	rv := objc.Send[NSPathControlItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

//
// # Instance Properties
//
//   - [NSPathControlItem.AttributedTitle]
//   - [NSPathControlItem.SetAttributedTitle]
//   - [NSPathControlItem.Image]
//   - [NSPathControlItem.SetImage]
//   - [NSPathControlItem.Title]
//   - [NSPathControlItem.SetTitle]
//   - [NSPathControlItem.URL]
// See: https://developer.apple.com/documentation/AppKit/NSPathControlItem
type NSPathControlItem struct {
	objectivec.Object
}

// NSPathControlItemFromID constructs a [NSPathControlItem] from an objc.ID.
func NSPathControlItemFromID(id objc.ID) NSPathControlItem {
	return NSPathControlItem{objectivec.Object{ID: id}}
}
// NOTE: NSPathControlItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPathControlItem] class.
//
// # Instance Properties
//
//   - [INSPathControlItem.AttributedTitle]
//   - [INSPathControlItem.SetAttributedTitle]
//   - [INSPathControlItem.Image]
//   - [INSPathControlItem.SetImage]
//   - [INSPathControlItem.Title]
//   - [INSPathControlItem.SetTitle]
//   - [INSPathControlItem.URL]
//
// See: https://developer.apple.com/documentation/AppKit/NSPathControlItem
type INSPathControlItem interface {
	objectivec.IObject

	// Topic: Instance Properties

	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)
	Image() INSImage
	SetImage(value INSImage)
	Title() string
	SetTitle(value string)
	URL() foundation.INSURL
}

// Init initializes the instance.
func (p NSPathControlItem) Init() NSPathControlItem {
	rv := objc.Send[NSPathControlItem](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPathControlItem) Autorelease() NSPathControlItem {
	rv := objc.Send[NSPathControlItem](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPathControlItem creates a new NSPathControlItem instance.
func NewNSPathControlItem() NSPathControlItem {
	class := getNSPathControlItemClass()
	rv := objc.Send[NSPathControlItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSPathControlItem/attributedTitle
func (p NSPathControlItem) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (p NSPathControlItem) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](p.ID, objc.Sel("setAttributedTitle:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPathControlItem/image
func (p NSPathControlItem) Image() INSImage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (p NSPathControlItem) SetImage(value INSImage) {
	objc.Send[struct{}](p.ID, objc.Sel("setImage:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPathControlItem/title
func (p NSPathControlItem) Title() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPathControlItem) SetTitle(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setTitle:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/AppKit/NSPathControlItem/url
func (p NSPathControlItem) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

