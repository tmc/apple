// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSATSTypesetter] class.
var (
	_NSATSTypesetterClass     NSATSTypesetterClass
	_NSATSTypesetterClassOnce sync.Once
)

func getNSATSTypesetterClass() NSATSTypesetterClass {
	_NSATSTypesetterClassOnce.Do(func() {
		_NSATSTypesetterClass = NSATSTypesetterClass{class: objc.GetClass("NSATSTypesetter")}
	})
	return _NSATSTypesetterClass
}

// GetNSATSTypesetterClass returns the class object for NSATSTypesetter.
func GetNSATSTypesetterClass() NSATSTypesetterClass {
	return getNSATSTypesetterClass()
}

type NSATSTypesetterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSATSTypesetterClass) Alloc() NSATSTypesetter {
	rv := objc.Send[NSATSTypesetter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A concrete typesetter object that places glyphs during the text layout
// process.
//
// # Overview
// 
// An [NSATSTypesetter] object creates line fragment rectangles, positions
// glyphs within the line fragments, determines line breaks by word wrapping
// and hyphenation, and handles tab positioning. This object encapsulates the
// advanced typesetting capabilities of Core Text. [NSATSTypesetter] provides
// line and character spacing accuracy and supports many languages, including
// bidirectional languages.
//
// See: https://developer.apple.com/documentation/AppKit/NSATSTypesetter
type NSATSTypesetter struct {
	NSTypesetter
}

// NSATSTypesetterFromID constructs a [NSATSTypesetter] from an objc.ID.
//
// A concrete typesetter object that places glyphs during the text layout
// process.
func NSATSTypesetterFromID(id objc.ID) NSATSTypesetter {
	return NSATSTypesetter{NSTypesetter: NSTypesetterFromID(id)}
}
// NOTE: NSATSTypesetter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSATSTypesetter] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSATSTypesetter
type INSATSTypesetter interface {
	INSTypesetter
}





// Init initializes the instance.
func (a NSATSTypesetter) Init() NSATSTypesetter {
	rv := objc.Send[NSATSTypesetter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSATSTypesetter) Autorelease() NSATSTypesetter {
	rv := objc.Send[NSATSTypesetter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSATSTypesetter creates a new NSATSTypesetter instance.
func NewNSATSTypesetter() NSATSTypesetter {
	class := getNSATSTypesetterClass()
	rv := objc.Send[NSATSTypesetter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

























// Returns a shared instance of the typesetter.
//
// See: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/shared
func (_NSATSTypesetterClass NSATSTypesetterClass) SharedTypesetter() NSATSTypesetter {
	rv := objc.Send[objc.ID](objc.ID(_NSATSTypesetterClass.class), objc.Sel("sharedTypesetter"))
	return NSATSTypesetterFromID(objc.ID(rv))
}





















