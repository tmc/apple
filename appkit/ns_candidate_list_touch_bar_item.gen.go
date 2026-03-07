// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCandidateListTouchBarItem] class.
var (
	_NSCandidateListTouchBarItemClass     NSCandidateListTouchBarItemClass
	_NSCandidateListTouchBarItemClassOnce sync.Once
)

func getNSCandidateListTouchBarItemClass() NSCandidateListTouchBarItemClass {
	_NSCandidateListTouchBarItemClassOnce.Do(func() {
		_NSCandidateListTouchBarItemClass = NSCandidateListTouchBarItemClass{class: objc.GetClass("NSCandidateListTouchBarItem")}
	})
	return _NSCandidateListTouchBarItemClass
}

// GetNSCandidateListTouchBarItemClass returns the class object for NSCandidateListTouchBarItem.
func GetNSCandidateListTouchBarItemClass() NSCandidateListTouchBarItemClass {
	return getNSCandidateListTouchBarItemClass()
}

type NSCandidateListTouchBarItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCandidateListTouchBarItemClass) Alloc() NSCandidateListTouchBarItem {
	rv := objc.Send[NSCandidateListTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A bar item that, along with its delegate, provides a list of textual
// suggestions for the current text view.
//
// # Providing a client and a delegate
//
//   - [NSCandidateListTouchBarItem.Client]: The client object for the candidate list item.
//   - [NSCandidateListTouchBarItem.SetClient]
//   - [NSCandidateListTouchBarItem.Delegate]: The delegate of the candidate list item.
//   - [NSCandidateListTouchBarItem.SetDelegate]
//
// # Populating the candidate list
//
//   - [NSCandidateListTouchBarItem.SetCandidatesForSelectedRangeInString]: Sets an array of candidate objects to be displayed in the candidate list bar item.
//   - [NSCandidateListTouchBarItem.Candidates]: The array of candidate objects previously set by [setCandidates(_:forSelectedRange:in:)](<doc://com.apple.appkit/documentation/AppKit/NSCandidateListTouchBarItem/setCandidates(_:forSelectedRange:in:)>).
//   - [NSCandidateListTouchBarItem.AttributedStringForCandidate]: A block that converts a candidate object into an attributed string for display in the candidate list item.
//   - [NSCandidateListTouchBarItem.SetAttributedStringForCandidate]
//   - [NSCandidateListTouchBarItem.AllowsTextInputContextCandidates]: A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
//   - [NSCandidateListTouchBarItem.SetAllowsTextInputContextCandidates]
//
// # Handling collapsible behavior
//
//   - [NSCandidateListTouchBarItem.AllowsCollapsing]: A Boolean value that specifies whether the item can be collapsed.
//   - [NSCandidateListTouchBarItem.SetAllowsCollapsing]
//   - [NSCandidateListTouchBarItem.Collapsed]: A Boolean value that controls the visibility of the candidate list.
//   - [NSCandidateListTouchBarItem.SetCollapsed]
//
// # Managing candidate list visibility
//
//   - [NSCandidateListTouchBarItem.CandidateListVisible]: A Boolean value that represents the visibility of this item’s candidate list.
//   - [NSCandidateListTouchBarItem.UpdateWithInsertionPointVisibility]: Updates the candidate list visibility configuration based on the client’s insertion point state.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem
type NSCandidateListTouchBarItem struct {
	NSTouchBarItem
}

// NSCandidateListTouchBarItemFromID constructs a [NSCandidateListTouchBarItem] from an objc.ID.
//
// A bar item that, along with its delegate, provides a list of textual
// suggestions for the current text view.
func NSCandidateListTouchBarItemFromID(id objc.ID) NSCandidateListTouchBarItem {
	return NSCandidateListTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}
// NOTE: NSCandidateListTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCandidateListTouchBarItem] class.
//
// # Providing a client and a delegate
//
//   - [INSCandidateListTouchBarItem.Client]: The client object for the candidate list item.
//   - [INSCandidateListTouchBarItem.SetClient]
//   - [INSCandidateListTouchBarItem.Delegate]: The delegate of the candidate list item.
//   - [INSCandidateListTouchBarItem.SetDelegate]
//
// # Populating the candidate list
//
//   - [INSCandidateListTouchBarItem.SetCandidatesForSelectedRangeInString]: Sets an array of candidate objects to be displayed in the candidate list bar item.
//   - [INSCandidateListTouchBarItem.Candidates]: The array of candidate objects previously set by [setCandidates(_:forSelectedRange:in:)](<doc://com.apple.appkit/documentation/AppKit/NSCandidateListTouchBarItem/setCandidates(_:forSelectedRange:in:)>).
//   - [INSCandidateListTouchBarItem.AttributedStringForCandidate]: A block that converts a candidate object into an attributed string for display in the candidate list item.
//   - [INSCandidateListTouchBarItem.SetAttributedStringForCandidate]
//   - [INSCandidateListTouchBarItem.AllowsTextInputContextCandidates]: A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
//   - [INSCandidateListTouchBarItem.SetAllowsTextInputContextCandidates]
//
// # Handling collapsible behavior
//
//   - [INSCandidateListTouchBarItem.AllowsCollapsing]: A Boolean value that specifies whether the item can be collapsed.
//   - [INSCandidateListTouchBarItem.SetAllowsCollapsing]
//   - [INSCandidateListTouchBarItem.Collapsed]: A Boolean value that controls the visibility of the candidate list.
//   - [INSCandidateListTouchBarItem.SetCollapsed]
//
// # Managing candidate list visibility
//
//   - [INSCandidateListTouchBarItem.CandidateListVisible]: A Boolean value that represents the visibility of this item’s candidate list.
//   - [INSCandidateListTouchBarItem.UpdateWithInsertionPointVisibility]: Updates the candidate list visibility configuration based on the client’s insertion point state.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem
type INSCandidateListTouchBarItem interface {
	INSTouchBarItem

	// Topic: Providing a client and a delegate

	// The client object for the candidate list item.
	Client() INSView
	SetClient(value INSView)
	// The delegate of the candidate list item.
	Delegate() NSCandidateListTouchBarItemDelegate
	SetDelegate(value NSCandidateListTouchBarItemDelegate)

	// Topic: Populating the candidate list

	// Sets an array of candidate objects to be displayed in the candidate list bar item.
	SetCandidatesForSelectedRangeInString(candidates []objectivec.IObject, selectedRange foundation.NSRange, originalString string)
	// The array of candidate objects previously set by [setCandidates(_:forSelectedRange:in:)](<doc://com.apple.appkit/documentation/AppKit/NSCandidateListTouchBarItem/setCandidates(_:forSelectedRange:in:)>).
	Candidates() []objectivec.IObject
	// A block that converts a candidate object into an attributed string for display in the candidate list item.
	AttributedStringForCandidate() ObjectHandler
	SetAttributedStringForCandidate(value ObjectHandler)
	// A Boolean value that specifies whether a candidate list item displays candidates from text input providers.
	AllowsTextInputContextCandidates() bool
	SetAllowsTextInputContextCandidates(value bool)

	// Topic: Handling collapsible behavior

	// A Boolean value that specifies whether the item can be collapsed.
	AllowsCollapsing() bool
	SetAllowsCollapsing(value bool)
	// A Boolean value that controls the visibility of the candidate list.
	Collapsed() bool
	SetCollapsed(value bool)

	// Topic: Managing candidate list visibility

	// A Boolean value that represents the visibility of this item’s candidate list.
	CandidateListVisible() bool
	// Updates the candidate list visibility configuration based on the client’s insertion point state.
	UpdateWithInsertionPointVisibility(isVisible bool)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c NSCandidateListTouchBarItem) Init() NSCandidateListTouchBarItem {
	rv := objc.Send[NSCandidateListTouchBarItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCandidateListTouchBarItem) Autorelease() NSCandidateListTouchBarItem {
	rv := objc.Send[NSCandidateListTouchBarItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCandidateListTouchBarItem creates a new NSCandidateListTouchBarItem instance.
func NewNSCandidateListTouchBarItem() NSCandidateListTouchBarItem {
	class := getNSCandidateListTouchBarItemClass()
	rv := objc.Send[NSCandidateListTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewCandidateListTouchBarItemWithCoder(coder foundation.INSCoder) NSCandidateListTouchBarItem {
	instance := getNSCandidateListTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCandidateListTouchBarItemFromID(rv)
}


// Creates a new item with the specified identifier.
//
// # Discussion
// 
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewCandidateListTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSCandidateListTouchBarItem {
	instance := getNSCandidateListTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSCandidateListTouchBarItemFromID(rv)
}







// Sets an array of candidate objects to be displayed in the candidate list
// bar item.
//
// candidates: The array of candidates you wish to display in the candidate list item.
//
// selectedRange: A range ([NSRange]) within the string that the candidates represent.
// //
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
//
// originalString: The original string from which the candidate list was derived.
//
// # Discussion
// 
// The item uses the block in the [AttributedStringForCandidate] property to
// convert each candidate in the array into an attributed string. If the value
// of the [AttributedStringForCandidate] property is `nil` then
// [NSCandidateListTouchBarItem] can format candidates of type [NSString],
// [NSAttributedString], and [NSTextCheckingResult].
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/setCandidates(_:forSelectedRange:in:)
func (c NSCandidateListTouchBarItem) SetCandidatesForSelectedRangeInString(candidates []objectivec.IObject, selectedRange foundation.NSRange, originalString string) {
	objc.Send[objc.ID](c.ID, objc.Sel("setCandidates:forSelectedRange:inString:"), objectivec.IObjectSliceToNSArray(candidates), selectedRange, objc.String(originalString))
}

// Updates the candidate list visibility configuration based on the client’s
// insertion point state.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/update(withInsertionPointVisibility:)
func (c NSCandidateListTouchBarItem) UpdateWithInsertionPointVisibility(isVisible bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("updateWithInsertionPointVisibility:"), isVisible)
}
func (c NSCandidateListTouchBarItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The client object for the candidate list item.
//
// # Discussion
// 
// This object must be a subclass of [NSView] and adopt the
// [NSTextInputClient] protocol.
// 
// The candidate list item uses this property to show completion candidates as
// users enter text. You can disable this behavior with the
// [AllowsTextInputContextCandidates] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/client
func (c NSCandidateListTouchBarItem) Client() INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("client"))
	return NSViewFromID(objc.ID(rv))
}
func (c NSCandidateListTouchBarItem) SetClient(value INSView) {
	objc.Send[struct{}](c.ID, objc.Sel("setClient:"), value)
}



// The delegate of the candidate list item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/delegate
func (c NSCandidateListTouchBarItem) Delegate() NSCandidateListTouchBarItemDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return NSCandidateListTouchBarItemDelegateObjectFromID(rv)
}
func (c NSCandidateListTouchBarItem) SetDelegate(value NSCandidateListTouchBarItemDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}



// The array of candidate objects previously set by
// [SetCandidatesForSelectedRangeInString].
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/candidates
func (c NSCandidateListTouchBarItem) Candidates() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("candidates"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}



// A block that converts a candidate object into an attributed string for
// display in the candidate list item.
//
// # Discussion
// 
// This property is not required if the object type of your candidates is
// [NSString], [NSAttributedString], or [NSTextCheckingResult]. The default
// value of this property is `nil`.
// 
// If the attributed string you return does not specify [font] or
// [foregroundColor] then the candidate is displayed with the standard
// appearance font and color.
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
// [font]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/font
// [foregroundColor]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/foregroundColor
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/attributedStringForCandidate
func (c NSCandidateListTouchBarItem) AttributedStringForCandidate() ObjectHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("attributedStringForCandidate"))
	_ = rv
	return nil
}
func (c NSCandidateListTouchBarItem) SetAttributedStringForCandidate(value ObjectHandler) {
	block, cleanup := NewObjectBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setAttributedStringForCandidate:"), block)
}



// A Boolean value that specifies whether a candidate list item displays
// candidates from text input providers.
//
// # Discussion
// 
// When [true], the candidate list item shows candidates from the text input
// client provided in the [Client] property, before those in the [Candidates]
// property.
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/allowsTextInputContextCandidates
func (c NSCandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsTextInputContextCandidates"))
	return rv
}
func (c NSCandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsTextInputContextCandidates:"), value)
}



// A Boolean value that specifies whether the item can be collapsed.
//
// # Discussion
// 
// When [true], the candidate list item can be collapsed, [false] otherwise.
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/allowsCollapsing
func (c NSCandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsCollapsing"))
	return rv
}
func (c NSCandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsCollapsing:"), value)
}



// A Boolean value that controls the visibility of the candidate list.
//
// # Discussion
// 
// When [true], the candidate list is collapsed and not visible to the user.
// 
// The default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCollapsed
func (c NSCandidateListTouchBarItem) Collapsed() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCollapsed"))
	return rv
}
func (c NSCandidateListTouchBarItem) SetCollapsed(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCollapsed:"), value)
}



// A Boolean value that represents the visibility of this item’s candidate
// list.
//
// # Discussion
// 
// If [true], then the candidate list is currently visible, [false] otherwise.
// 
// When [Collapsed] is [false], and the item is not obscured by UI then this
// property is true.
// 
// This property is KVO compliant, and you should supply a candidate list when
// its value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItem/isCandidateListVisible
func (c NSCandidateListTouchBarItem) CandidateListVisible() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCandidateListVisible"))
	return rv
}


























