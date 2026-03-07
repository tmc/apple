// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextFinder] class.
var (
	_NSTextFinderClass     NSTextFinderClass
	_NSTextFinderClassOnce sync.Once
)

func getNSTextFinderClass() NSTextFinderClass {
	_NSTextFinderClassOnce.Do(func() {
		_NSTextFinderClass = NSTextFinderClass{class: objc.GetClass("NSTextFinder")}
	})
	return _NSTextFinderClass
}

// GetNSTextFinderClass returns the class object for NSTextFinder.
func GetNSTextFinderClass() NSTextFinderClass {
	return getNSTextFinderClass()
}

type NSTextFinderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextFinderClass) Alloc() NSTextFinder {
	rv := objc.Send[NSTextFinder](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An optional search-and-replace find interface inside a view, usually a
// scroll view.
//
// # Overview
// 
// The class serves as a controller for the standard Cocoa find bar. The
// [NSTextFinder] class interacts heavily with a `client` object which
// supports the [NSTextFinderClient] protocol. The client object provides
// access to the content being searched and provides visual feedback for a
// search operation.
// 
// All menu items related to finding (Find…, Find Next, Find Previous, Use
// Selection for Find, etc.) should have the same action,
// [PerformTextFinderAction], which gets sent down the responder chain in the
// standard method.
// 
// # Implementing a Find Bar
// 
// A responder of [PerformTextFinderAction] is responsible for creating and
// owning an instance of [NSTextFinder]. Before any other messages are sent to
// the [NSTextFinder], you should set its [NSTextFinder.Client] property to an object which
// implements the [NSTextFinderClient] protocol.
// 
// Each user interface item used for finding has a different tag value , which
// corresponds to the appropriate value in [NSTextFinder.Action]. Upon receipt
// of the [PerformTextFinderAction] message, the responder should call the
// following on its [NSTextFinder] instance:
// 
// This method will determine the desired action from the tag and make various
// callbacks to [NSTextFinder.Client] to perform that action. These callbacks are defined
// in the [NSTextFinderClient] protocol.
// 
// Validation occurs by a similar pattern. The responder should implement
// [ValidateUserInterfaceItem] and, when the item’s action is the
// [PerformTextFinderAction] method, it should pass the item’s tag to the
// following method and return the result:
// 
// This method will invoke the required client methods to determine what the
// appropriate response should be. These callbacks are also defined in the
// [NSTextFinderClient] protocol.
// 
// When an action is performed, the text finder will ask its client for the
// string to search. There are two ways the client can provide this string. It
// can either implement the `string` method, and simply return an [NSString]
// instance, or it can implement the following methods:
// [StringAtIndexEffectiveRangeEndsWithSearchBoundary] and [StringLength].
// 
// These methods make it possible to use the text finder with data that cannot
// be easily or efficiently flattened into a single string. In the first
// method, the client should return the string which contains the character at
// the given index in as conceptual concatenation of all strings that are to
// be searched.
// 
// For example, if a client had the strings “The quick”, “brown fox” ,
// “jumped over the lazy”, and “dog”, and a [NSTextFinder] instance
// requests the string at index 17, then the client should return “brown
// fox”. If all the strings were concatenated together, this would be the
// ‘x’ in “fox”. Additionally, the client should return, by-reference
// through the `outRange` parameter, the effective range of the string that is
// returned. In this example, the range of “brown fox” is {9, 9}, because,
// in the imagined concatenation, the substring starts at index 9 and is 9
// characters long.
// 
// In some cases, it is not desirable for a match to be found which overlaps
// multiple strings returned by the
// [StringAtIndexEffectiveRangeEndsWithSearchBoundary] method. For example,
// suppose a client has a list of names like “John Doe”, and “Jane
// Doe”, etc. Normally, if the string is concatenated together like so:
// “John Doe Jane Doe”, then a search for “Doe Jane” would succeed. To
// prevent this often undesirable behavior, the client should return [true],
// by-reference, through the `outFlag` parameter, when returning each
// individual string. This indicates that there is a boundary between the
// current string and the next string that should not be crossed when
// searching for a match.
// 
// One of the two approaches (implementing the `string` method the
// [StringAtIndexEffectiveRangeEndsWithSearchBoundary] and [StringLength]
// methods) must be implemented by the client or the [NSTextFinder] instance
// will not function. If all three methods are implemented,
// [StringAtIndexEffectiveRangeEndsWithSearchBoundary] will be used by
// default.
// 
// The text finder may require additional information from the [NSTextFinder.Client] object
// to perform an action, or it may require the `client` to perform some parts
// of the action on its behalf. The methods and properties described in
// [NSTextFinder] describes the hooks the text finder requires for each of the
// actions it supports. If the client does not implement one of these methods
// or properties, then the action that requires it will be disabled.
// 
// # Text Finder Client Implementation Requirements
// 
// The text finder’s client may implement the following properties for by
// the [NSTextFinder.ValidateAction] method: [NSTextFinder.IsSelectable], [NSTextFinder.AllowsMultipleSelection], and
// `editable`. If any of these properties is not implemented, a value of
// [true] is assumed. Returning [false] from any of these methods will disable
// the actions that require them. For more information about these properties
// see [NSTextFinderClient].
// 
// The following implementation’s are required by the text client to support
// the specified actions:
// 
// - The [NSTextFinderClient] protocol must implement the [NSTextFinder.FirstSelectedRange]
// property to support the following functionality:
// [NSTextFinder.Action.nextMatch], [NSTextFinder.Action.previousMatch],
// [NSTextFinder.Action.replace], [NSTextFinder.Action.replaceAndFind], and
// [NSTextFinder.Action.setSearchString] actions. The `client` implementation
// of [NSTextFinder.FirstSelectedRange] needs to return its first selected range, or
// {index, 0} to indicate the location of the insertion point if there is no
// selection. - The [NSTextFinder.SelectedRanges] property is required for the
// [NSTextFinder.Action.replaceAllInSelection],
// [NSTextFinder.Action.selectAll], and
// [NSTextFinder.Action.selectAllInSelection] actions. The array should
// contain [NSRange] structures wrapped in [NSValue] instances. - The
// [ScrollRangeToVisible] method is used by many actions, but is not strictly
// required by any to perform correctly. - The client protocol’s
// [ShouldReplaceCharactersInRangesWithStrings],
// [ReplaceCharactersInRangeWithString] and [DidReplaceCharacters] methods are
// used by the [NSTextFinder.Action.replace],
// [NSTextFinder.Action.replaceAll],
// [NSTextFinder.Action.replaceAllInSelection], and
// [NSTextFinder.Action.replaceAndFind] actions. The NSTextFinder instance
// does not have the ability to directly make changes to the [NSTextFinder.Client] content,
// so the `client` is responsible for performing replace operations in these
// methods. - Before a replace operation is performed, the [NSTextFinder]
// instance calls the [ShouldReplaceCharactersInRangesWithStrings] method to
// determine if a replacement should take place. If it returns [false], then
// the given range will not be replaced. If the method returns [true], or is
// not implemented, then the [NSTextFinder] instance will call the
// [ReplaceCharactersInRangeWithString] method, instructing the `client` to
// carry out the replacement. Finally, the [DidReplaceCharacters] method will
// be called, if implemented, to indicate that the replacement was completed.
// For [NSTextFinder.Action.replaceAll] actions, these methods may be called
// multiple times, starting from the end of the string and moving toward the
// beginning, in order to preserve the indexes of the matches which precede
// the current one.
// 
// # Incremental Search Support
// 
// A common feature used in conjunction with the find bar is incremental
// search. This feature allows users to find matches immediately as they are
// typing. In OS X v10.7, the text finder provides this feature for its
// clients with minimal additional API.
// 
// Incremental searching can be enabled by setting the
// [NSTextFinder.IncrementalSearchingEnabled] property [true]. This property alone is
// sufficient to start searching incrementally.
// 
// For improved responsiveness, when a document is sufficiently long, the text
// finder will search the document in the background. To ensure thread-safety,
// a client using incremental search must call the
// [NSTextFinder.NoteClientStringWillChange] method before any changes are made to the
// string provided to the text finder.
// 
// During incremental search, all visible matches are highlighted. If the
// `client` object that conforms to the [NSTextFinderClient] protocol
// implements the read-only [NSTextFinder.VisibleCharacterRanges] property, then by default
// a gray overlay will appear over your find bar container’s content view
// with transparent areas for each match. This view should be a superview of
// all subviews reported by the text finder client. The [NSScrollView] class
// already implements this property, so you only need to implement this
// property when using a different container class.
// 
// Finally, this mode also requires the same two [NSTextFinderClient] protocol
// methods that are used to display the find indicator:
// [ContentViewAtIndexEffectiveCharacterRange] and [RectsForCharacterRange] be
// implemented. However, the implementation of these methods is identical for
// both purposes, so no additional work is required to support incremental
// search.
// 
// # Replacing Text
// 
// The text replacement methods:[NSTextFinder.Action],
// [NSTextFinder.FindIndicatorNeedsUpdate], and [NSTextFinder.ValidateAction] are used by the replace,
// replace all, replace all in section, and replace and find actions.
// 
// Before a replace operation is performed, the [NSTextFinder] instance calls
// the `client` object’s [NSTextFinder.Action] method to determine if a
// replacement should take place. If it returns [false], then the characters
// in the given ranges will not be replaced. If the method returns [true], or
// is not implemented, then the second method, [NSTextFinder.FindIndicatorNeedsUpdate],
// instructing the client to carry out the replacement. Finally,
// [NSTextFinder.ValidateAction], if implemented, is invoked to indicate that the
// replacement was completed.
// 
// For replace all actions, these methods will be called multiple times,
// starting from the last match and moving toward the first, in order to
// preserve the indexes of the matches which precede the current one.
// 
// # The Text Finder Container
// 
// In order to display the find bar, a container for the find bar must be
// specified. The container is an object that conforms to the
// [NSTextFinderBarContainer] protocol. You specify a find bar container using
// the following NSTextFinder class’s [NSTextFinder.FindBarContainer] property.
// 
// When a new NSTextFinder instance is created and instructed to display the
// find bar, it will create a view for the find bar and assign that to the
// container via the [NSTextFinderBarContainer] class’s [NSTextFinder.FindBarView]
// property. The container then owns that view and should release it when the
// container is deallocated. The container should make the find bar visible
// when the container’s [NSTextFinder.IsFindBarVisible] property is set to [true]. The
// container should implement the [FindBarViewDidChangeHeight] method so it
// can reposition the find bar when its height changes, usually in response to
// user action.
// 
// # Implementation by AppKit Classes
// 
// Two AppKit classes already provide support for the [NSTextFinder] class,
// including: the [NSScrollView] and [NSTextView] classes.
// 
// # Scroll View Support for the Find Bar
// 
// The [NSScrollView] class conforms to [NSTextFinderBarContainer] protocol in
// order to support the find bar for any document view searched by the find
// bar. The find bar can be positioned either above or below the document view
// by assigning one of the values of the [NSScrollView.FindBarPosition]
// constants to the [NSTextFinder.FindBarPosition] property.
// 
// # Text View Support for the Find Bar
// 
// The [NSTextView] class also supports the find bar. The find bar can be
// enabled or disabled on a text view with the [NSTextFinder.UsesFindBar] property.
// 
// Since OS X v10.5, the [NSTextView] class has used an animated find
// indicator to give feedback to the user about a successful find action. The
// find indicator could be activated manually on an text view via the
// [ShowFindIndicatorForRange] method.
// 
// To provide this functionality for text finder clients in OS X v10.7, the
// [NSTextFinder] class shows the find indicator at the appropriate time
// automatically when performing a find. However, text finder needs to know
// where to show the indicator and it needs assistance in drawing the text
// contents of the find indicator. During incremental search the
// [NSTextFinder] class needs also to know the locations of all the visible
// matches so it can highlight them as well. The following method additions to
// the [NSTextFinderClient] protocol provide these capabilities:
// [ContentViewAtIndexEffectiveCharacterRange] and [RectsForCharacterRange].
// 
// The [ContentViewAtIndexEffectiveCharacterRange] method is called to
// determine what view the content is displayed in, that is, over which the
// find indicator will be displayed. Since content may potentially be spread
// over multiple views (like the [NSLayoutManager] class , which supports
// multiple text views), the method asks for the view at a given index, and
// the full range of the string that is contained by that view.
// 
// In the [RectsForCharacterRange] method, the client should determine and
// return the rectangles in which the content with the given range is
// displayed in its `contentView`. The given range is guaranteed not to span
// multiple content views. The returned rectangles tell the [NSTextFinder]
// instance where the matched range is found, so it can show the find
// indicator there.
// 
// The [NSTextFinder] and [NSView] classes will handle the find indicator
// correctly when the [ContentView] is resized, moved, or removed from the
// view hierarchy. If your content view’s scrolling is done by an
// [NSScrollView] object, the find indicator will also be handled for you
// during scrolling. Beyond these cases, there may be some circumstances where
// the find indicator should be immediately cancelled and hidden, such as when
// the view’s content or selection is changed without the knowledge of the
// text finder . The find indicator can be cancelled using the [NSTextFinder]
// [NSTextFinder.CancelFindIndicator] method. If your document is not scrolled by
// [NSScrollView], then you should set the [NSTextFinder.FindIndicatorNeedsUpdate] property
// to [true].
// 
// [NSTextFinder] is responsible for drawing the yellow find indicator
// background bezel, but the view must provide the contents. [NSTextFinder]
// also sets up a drawing context and causes the content view to draw into it.
// There are two ways this can happen.
// 
// - The [NSTextFinder] instance invokes the `client` object’s
// [DrawCharactersInRangeForContentView] method, if implemented. The client
// should then draw the requested characters (or ask the containing view to
// draw the characters, which is provided as a convenience) at the origin of
// the current context. If the requested character range partially intersects
// a glyph range, the client may draw the entire glyph, if necessary for
// performance. - If [DrawCharactersInRangeForContentView] method is not
// implemented, then the [NSTextFinder] instance uses the normal view drawing
// mechanisms. The `contentView` does not need take any additional action to
// make this happen
// 
// In OS X v10.7, the [NSTextView] class also provides incremental search
// support. It is disabled by default, but can be enabled by setting the
// [NSTextFinder.IncrementalSearchingEnabled] property to [true]. Also, because incremental
// searching requires the find bar, [NSTextFinder.UsesFindBar] must be set to [true] for
// incremental searching to be occur.
//
// [NSScrollView.FindBarPosition]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum
// [NSTextFinder.Action.nextMatch]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/nextMatch
// [NSTextFinder.Action.previousMatch]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/previousMatch
// [NSTextFinder.Action.replaceAllInSelection]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replaceAllInSelection
// [NSTextFinder.Action.replaceAll]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replaceAll
// [NSTextFinder.Action.replaceAndFind]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replaceAndFind
// [NSTextFinder.Action.replace]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replace
// [NSTextFinder.Action.selectAllInSelection]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/selectAllInSelection
// [NSTextFinder.Action.selectAll]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/selectAll
// [NSTextFinder.Action.setSearchString]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/setSearchString
// [NSTextFinder.Action]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Validating and Performing Text Finding
//
//   - [NSTextFinder.PerformAction]: Performs the specified text finding action.
//   - [NSTextFinder.ValidateAction]: Allows validation of the find action before performing.
//   - [NSTextFinder.CancelFindIndicator]: Cancels the find indicator immediately.
//
// # Getting the Find Bar Container
//
//   - [NSTextFinder.FindBarContainer]: Specifies the find bar container.
//   - [NSTextFinder.SetFindBarContainer]
//
// # Getting and Setting the Find Bar Client
//
//   - [NSTextFinder.Client]: The object that provides the target search string, find bar location, and feedback methods.
//   - [NSTextFinder.SetClient]
//
// # Noting Changes in the Original Content
//
//   - [NSTextFinder.NoteClientStringWillChange]: Invoke this method when the searched content will change.
//
// # Updating the Find Indicator
//
//   - [NSTextFinder.FindIndicatorNeedsUpdate]: Invoke to specify that the find indicator needs updating when not contained within a scroll view.
//   - [NSTextFinder.SetFindIndicatorNeedsUpdate]
//
// # Incremental Search Configuration
//
//   - [NSTextFinder.IncrementalMatchRanges]: Array of incremental search matches posted on the main queue, which have been found during a background search.
//   - [NSTextFinder.IncrementalSearchingEnabled]: Determines if incremental searching is enabled.
//   - [NSTextFinder.SetIncrementalSearchingEnabled]
//   - [NSTextFinder.IncrementalSearchingShouldDimContentView]: Determines the type of incremental search feedback to be presented
//   - [NSTextFinder.SetIncrementalSearchingShouldDimContentView]
//
// # Initializers
//
//   - [NSTextFinder.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder
type NSTextFinder struct {
	objectivec.Object
}

// NSTextFinderFromID constructs a [NSTextFinder] from an objc.ID.
//
// An optional search-and-replace find interface inside a view, usually a
// scroll view.
func NSTextFinderFromID(id objc.ID) NSTextFinder {
	return NSTextFinder{objectivec.Object{id}}
}
// NOTE: NSTextFinder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextFinder] class.
//
// # Validating and Performing Text Finding
//
//   - [INSTextFinder.PerformAction]: Performs the specified text finding action.
//   - [INSTextFinder.ValidateAction]: Allows validation of the find action before performing.
//   - [INSTextFinder.CancelFindIndicator]: Cancels the find indicator immediately.
//
// # Getting the Find Bar Container
//
//   - [INSTextFinder.FindBarContainer]: Specifies the find bar container.
//   - [INSTextFinder.SetFindBarContainer]
//
// # Getting and Setting the Find Bar Client
//
//   - [INSTextFinder.Client]: The object that provides the target search string, find bar location, and feedback methods.
//   - [INSTextFinder.SetClient]
//
// # Noting Changes in the Original Content
//
//   - [INSTextFinder.NoteClientStringWillChange]: Invoke this method when the searched content will change.
//
// # Updating the Find Indicator
//
//   - [INSTextFinder.FindIndicatorNeedsUpdate]: Invoke to specify that the find indicator needs updating when not contained within a scroll view.
//   - [INSTextFinder.SetFindIndicatorNeedsUpdate]
//
// # Incremental Search Configuration
//
//   - [INSTextFinder.IncrementalMatchRanges]: Array of incremental search matches posted on the main queue, which have been found during a background search.
//   - [INSTextFinder.IncrementalSearchingEnabled]: Determines if incremental searching is enabled.
//   - [INSTextFinder.SetIncrementalSearchingEnabled]
//   - [INSTextFinder.IncrementalSearchingShouldDimContentView]: Determines the type of incremental search feedback to be presented
//   - [INSTextFinder.SetIncrementalSearchingShouldDimContentView]
//
// # Initializers
//
//   - [INSTextFinder.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder
type INSTextFinder interface {
	objectivec.IObject

	// Topic: Validating and Performing Text Finding

	// Performs the specified text finding action.
	PerformAction(op NSTextFinderAction)
	// Allows validation of the find action before performing.
	ValidateAction(op NSTextFinderAction) bool
	// Cancels the find indicator immediately.
	CancelFindIndicator()

	// Topic: Getting the Find Bar Container

	// Specifies the find bar container.
	FindBarContainer() NSTextFinderBarContainer
	SetFindBarContainer(value NSTextFinderBarContainer)

	// Topic: Getting and Setting the Find Bar Client

	// The object that provides the target search string, find bar location, and feedback methods.
	Client() NSTextFinderClient
	SetClient(value NSTextFinderClient)

	// Topic: Noting Changes in the Original Content

	// Invoke this method when the searched content will change.
	NoteClientStringWillChange()

	// Topic: Updating the Find Indicator

	// Invoke to specify that the find indicator needs updating when not contained within a scroll view.
	FindIndicatorNeedsUpdate() bool
	SetFindIndicatorNeedsUpdate(value bool)

	// Topic: Incremental Search Configuration

	// Array of incremental search matches posted on the main queue, which have been found during a background search.
	IncrementalMatchRanges() []foundation.NSValue
	// Determines if incremental searching is enabled.
	IncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	// Determines the type of incremental search feedback to be presented
	IncrementalSearchingShouldDimContentView() bool
	SetIncrementalSearchingShouldDimContentView(value bool)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) NSTextFinder

	// Type for the Find panel metadata property list.
	TextFinderOptions() NSPasteboardType
	// The position of the find bar.
	FindBarPosition() NSScrollViewFindBarPosition
	SetFindBarPosition(value NSScrollViewFindBarPosition)
	// The view assigned by the text bar as the find bar view for the container.
	FindBarView() INSView
	SetFindBarView(value INSView)
	// Returns whether the container should display its find bar.
	IsFindBarVisible() bool
	SetIsFindBarVisible(value bool)
	// Returns whether multiple items can be selected.
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	// Returns the currently selected range.
	FirstSelectedRange() foundation.NSRange
	SetFirstSelectedRange(value foundation.NSRange)
	// Returns whether the text is selectable.
	IsSelectable() bool
	SetIsSelectable(value bool)
	// Returns an array of selected ranges.
	SelectedRanges() foundation.NSValue
	SetSelectedRanges(value foundation.NSValue)
	// An array of visible character ranges.
	VisibleCharacterRanges() foundation.NSValue
	SetVisibleCharacterRanges(value foundation.NSValue)
	// A Boolean value that indicates whether to use the find bar for this text view.
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTextFinder) Init() NSTextFinder {
	rv := objc.Send[NSTextFinder](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextFinder) Autorelease() NSTextFinder {
	rv := objc.Send[NSTextFinder](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextFinder creates a new NSTextFinder instance.
func NewNSTextFinder() NSTextFinder {
	class := getNSTextFinderClass()
	rv := objc.Send[NSTextFinder](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/init(coder:)
func NewTextFinderWithCoder(coder foundation.INSCoder) NSTextFinder {
	instance := getNSTextFinderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextFinderFromID(rv)
}







// Performs the specified text finding action.
//
// op: The text finding action. See [NSTextFinder.Action] for the possible values.
// //
// [NSTextFinder.Action]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action
//
// # Discussion
// 
// Objects that respond to [PerformTextFinderAction] typically call
// [ValidateAction] to ensure that the action is valid and then invoke
// [PerformAction] if validation is successful.
// 
// When invoking the [ValidateAction] and [PerformAction] the item or
// sender’s tag should be passed as the parameter. By convention, the
// `sender` parameter for this method will have an [NSTextFinder.Action] set
// as its tag. The responder that receives this method should pass the tag as
// the action for this method:
//
// [NSTextFinder.Action]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/performAction(_:)
func (t NSTextFinder) PerformAction(op NSTextFinderAction) {
	objc.Send[objc.ID](t.ID, objc.Sel("performAction:"), op)
}

// Allows validation of the find action before performing.
//
// op: The sender’s tag.
//
// # Return Value
// 
// [true] if the operation is valid; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Responders the [NSResponder] method [PerformTextFinderAction] should call
// this method.
// 
// This method should be called by an implementation of
// [ValidateUserInterfaceItem] to properly validate items with an action of
// [PerformTextFinderAction]. The responder’s [ValidateUserInterfaceItem] or
// [validateMenuItem:] implementation should pass the tag as the action for
// this method..
//
// [validateMenuItem:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validateMenuItem:
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/validateAction(_:)
func (t NSTextFinder) ValidateAction(op NSTextFinderAction) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("validateAction:"), op)
	return rv
}

// Cancels the find indicator immediately.
//
// # Discussion
// 
// There may be some circumstances where the find indicator should be
// immediately cancelled or hidden, such as when the view’s content or
// selection is changed without the knowledge of the text finder. This method
// will immediately cancel the current find indicator.
// 
// The [NSTextFinder] and [NSView] classes will handle the find indicator
// correctly when a content view is resized, moved, or removed from the view
// hierarchy. If your content view’s scrolling is done by an [NSScrollView],
// the find indicator will also be handled for you during scrolling.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/cancelFindIndicator()
func (t NSTextFinder) CancelFindIndicator() {
	objc.Send[objc.ID](t.ID, objc.Sel("cancelFindIndicator"))
}

// Invoke this method when the searched content will change.
//
// # Discussion
// 
// When incremental search is enabled, this method must be called when it is
// known that the client’s string will be modified. This method must be
// called before the client string modification takes place.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/noteClientStringWillChange()
func (t NSTextFinder) NoteClientStringWillChange() {
	objc.Send[objc.ID](t.ID, objc.Sel("noteClientStringWillChange"))
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/init(coder:)
func (t NSTextFinder) InitWithCoder(coder foundation.INSCoder) NSTextFinder {
	rv := objc.Send[NSTextFinder](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (t NSTextFinder) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Override this method to draw custom highlighting.
//
// rect: The rectangle that needs to be drawn highlighted in the current coordinate
// system.
//
// # Discussion
// 
// If [IncrementalSearchingShouldDimContentView] is [false], it is recommended
// to highlight incremental matches in your own view. However, some
// applications may choose to show incremental search values in a different
// manner.
// 
// This method is not recommended to be overridden. The text finder never
// calls it. The view calls it to get the standard highlight behavior. It is
// recommended that views use this method do draw the highlight for
// consistency and to allow Application Kit to tweak the behavior in the
// future. If the view wants custom drawing, then it should be implemented by
// the view.
// 
// The common usage pattern for this is: draw the background, draw the
// incremental match highlights for the [IncrementalMatchRanges], and then
// draw the text.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/drawIncrementalMatchHighlight(in:)
func (_NSTextFinderClass NSTextFinderClass) DrawIncrementalMatchHighlightInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](objc.ID(_NSTextFinderClass.class), objc.Sel("drawIncrementalMatchHighlightInRect:"), rect)
}








// Specifies the find bar container.
//
// # Discussion
// 
// This property must be set to support the find bar.
// 
// When an instance of [NSTextFinder] receives a request to display the find
// bar, it creates a view for the find bar and assigns it to the [FindBarView]
// property of its `findBarContainer`. This container owns that view, and
// it’s responsible for making the view visible when the container’s
// [IsFindBarVisible] property is `true`. Implement the container’s
// [FindBarViewDidChangeHeight] method to reposition the find bar when its
// height changes, which typically occurs in response to a user interaction.
// 
// The [NSScrollView] class implements [NSTextFinderBarContainer] protocol and
// is the correct place to display the find bar in most circumstances. The
// container may freely modify the find bar view’s width and origin, but not
// its height.
// 
// If this property is not set, then the find bar cannot be shown.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/findBarContainer
func (t NSTextFinder) FindBarContainer() NSTextFinderBarContainer {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("findBarContainer"))
	return NSTextFinderBarContainerObjectFromID(rv)
}
func (t NSTextFinder) SetFindBarContainer(value NSTextFinderBarContainer) {
	objc.Send[struct{}](t.ID, objc.Sel("setFindBarContainer:"), value)
}



// The object that provides the target search string, find bar location, and
// feedback methods.
//
// # Discussion
// 
// The [NSTextFinder] instance class must be associated with a client object
// that implements the NSTextFinderClient protocol in order to function. The
// client is responsible for providing the string to be searched, the location
// for the find bar, and the methods which control feedback to the user about
// the search results.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/client
func (t NSTextFinder) Client() NSTextFinderClient {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("client"))
	return NSTextFinderClientObjectFromID(rv)
}
func (t NSTextFinder) SetClient(value NSTextFinderClient) {
	objc.Send[struct{}](t.ID, objc.Sel("setClient:"), value)
}



// Invoke to specify that the find indicator needs updating when not contained
// within a scroll view.
//
// # Discussion
// 
// If the [Client] object’s document is not scrolled by an instance of
// [NSScrollView], then set this property to [true] when scrolling occurs to
// cause the find indicator to be updated appropriately.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/findIndicatorNeedsUpdate
func (t NSTextFinder) FindIndicatorNeedsUpdate() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("findIndicatorNeedsUpdate"))
	return rv
}
func (t NSTextFinder) SetFindIndicatorNeedsUpdate(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setFindIndicatorNeedsUpdate:"), value)
}



// Array of incremental search matches posted on the main queue, which have
// been found during a background search.
//
// # Discussion
// 
// This array is updated periodically on the main queue as the incremental
// search operation on a background queue finds matches. You can use this
// property when incrementalSearchingShouldDimContentView is [false] to know
// where to draw highlights for incremental matches.
// 
// If no incremental search is active, or there are no matches found, this
// array will be empty. If an incremental search is currently in progress, but
// not yet complete, this will return all the ranges found so far.
// 
// This array is key-value observing compliant and can be observed to know
// when to update your highlights. When [new] and [old] options are used, the
// key-value observing change dictionary provides the ranges (and their
// indexes) that are added or removed. This allows the invalidation of the
// minimal region necessary to sync highlights with the receiver’s results.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [new]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/new
// [old]: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions/old
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalMatchRanges
func (t NSTextFinder) IncrementalMatchRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("incrementalMatchRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}



// Determines if incremental searching is enabled.
//
// # Discussion
// 
// If [true], then the find bar will do incremental searches. If it returns
// [false], then the find bar will behave regularly.
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/isIncrementalSearchingEnabled
func (t NSTextFinder) IncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}
func (t NSTextFinder) SetIncrementalSearchingEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIncrementalSearchingEnabled:"), value)
}



// Determines the type of incremental search feedback to be presented
//
// # Discussion
// 
// If [true], then when an incremental search begins, the `findBarContainer`
// instance’s parent `contentView` will be dimmed, except for the locations
// of the incremental matches. If [false], then the incremental matches will
// not be highlighted automatically, but you can use incrementalMatchRanges to
// highlight the matches yourself.
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinder/incrementalSearchingShouldDimContentView
func (t NSTextFinder) IncrementalSearchingShouldDimContentView() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("incrementalSearchingShouldDimContentView"))
	return rv
}
func (t NSTextFinder) SetIncrementalSearchingShouldDimContentView(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIncrementalSearchingShouldDimContentView:"), value)
}



// Type for the Find panel metadata property list.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/pasteboardtype/textfinderoptions
func (t NSTextFinder) TextFinderOptions() NSPasteboardType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NSPasteboardTypeTextFinderOptions"))
	return NSPasteboardType(foundation.NSStringFromID(rv).String())
}



// The position of the find bar.
//
// See: https://developer.apple.com/documentation/appkit/nsscrollview/findbarposition-swift.property
func (t NSTextFinder) FindBarPosition() NSScrollViewFindBarPosition {
	rv := objc.Send[NSScrollViewFindBarPosition](t.ID, objc.Sel("findBarPosition"))
	return NSScrollViewFindBarPosition(rv)
}
func (t NSTextFinder) SetFindBarPosition(value NSScrollViewFindBarPosition) {
	objc.Send[struct{}](t.ID, objc.Sel("setFindBarPosition:"), value)
}



// The view assigned by the text bar as the find bar view for the container.
//
// See: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/findbarview
func (t NSTextFinder) FindBarView() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("findBarView"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSTextFinder) SetFindBarView(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setFindBarView:"), value)
}



// Returns whether the container should display its find bar.
//
// See: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/isfindbarvisible
func (t NSTextFinder) IsFindBarVisible() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("findBarVisible"))
	return rv
}
func (t NSTextFinder) SetIsFindBarVisible(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setFindBarVisible:"), value)
}



// Returns whether multiple items can be selected.
//
// See: https://developer.apple.com/documentation/appkit/nstextfinderclient/allowsmultipleselection
func (t NSTextFinder) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}
func (t NSTextFinder) SetAllowsMultipleSelection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}



// Returns the currently selected range.
//
// See: https://developer.apple.com/documentation/appkit/nstextfinderclient/firstselectedrange
func (t NSTextFinder) FirstSelectedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("firstSelectedRange"))
	return foundation.NSRange(rv)
}
func (t NSTextFinder) SetFirstSelectedRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setFirstSelectedRange:"), value)
}



// Returns whether the text is selectable.
//
// See: https://developer.apple.com/documentation/appkit/nstextfinderclient/isselectable
func (t NSTextFinder) IsSelectable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("selectable"))
	return rv
}
func (t NSTextFinder) SetIsSelectable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectable:"), value)
}



// Returns an array of selected ranges.
//
// See: https://developer.apple.com/documentation/appkit/nstextfinderclient/selectedranges
func (t NSTextFinder) SelectedRanges() foundation.NSValue {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedRanges"))
	return foundation.NSValueFromID(objc.ID(rv))
}
func (t NSTextFinder) SetSelectedRanges(value foundation.NSValue) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedRanges:"), value)
}



// An array of visible character ranges.
//
// See: https://developer.apple.com/documentation/appkit/nstextfinderclient/visiblecharacterranges
func (t NSTextFinder) VisibleCharacterRanges() foundation.NSValue {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("visibleCharacterRanges"))
	return foundation.NSValueFromID(objc.ID(rv))
}
func (t NSTextFinder) SetVisibleCharacterRanges(value foundation.NSValue) {
	objc.Send[struct{}](t.ID, objc.Sel("setVisibleCharacterRanges:"), value)
}



// A Boolean value that indicates whether to use the find bar for this text
// view.
//
// See: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t NSTextFinder) UsesFindBar() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesFindBar"))
	return rv
}
func (t NSTextFinder) SetUsesFindBar(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesFindBar:"), value)
}
























