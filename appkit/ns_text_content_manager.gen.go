// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextContentManager] class.
var (
	_NSTextContentManagerClass     NSTextContentManagerClass
	_NSTextContentManagerClassOnce sync.Once
)

func getNSTextContentManagerClass() NSTextContentManagerClass {
	_NSTextContentManagerClassOnce.Do(func() {
		_NSTextContentManagerClass = NSTextContentManagerClass{class: objc.GetClass("NSTextContentManager")}
	})
	return _NSTextContentManagerClass
}

// GetNSTextContentManagerClass returns the class object for NSTextContentManager.
func GetNSTextContentManagerClass() NSTextContentManagerClass {
	return getNSTextContentManagerClass()
}

type NSTextContentManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextContentManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextContentManagerClass) Alloc() NSTextContentManager {
	rv := objc.Send[NSTextContentManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that defines the interface and a default implementation
// for managing the text document contents.
//
// # Creating a content manager
//
//   - [NSTextContentManager.InitWithCoder]: Creates a new content manager object from data in an unarchiver.
//
// # Controlling backing store synchronization
//
//   - [NSTextContentManager.AutomaticallySynchronizesToBackingStore]: Determines whether to automatically synchronize with the backing store when an editing transaction finishes.
//   - [NSTextContentManager.SetAutomaticallySynchronizesToBackingStore]
//
// # Performing transactions
//
//   - [NSTextContentManager.HasEditingTransaction]: Indicates there’s an active editing transaction from the primary text layout manager.
//   - [NSTextContentManager.PerformEditingTransactionUsingBlock]: Performs an editing transaction and invokes a block upon completion.
//   - [NSTextContentManager.RecordEditActionInRangeNewTextRange]: Records information about an edit action to the transaction.
//
// # Working with layout managers
//
//   - [NSTextContentManager.PrimaryTextLayoutManager]: The primary text layout manager for this content.
//   - [NSTextContentManager.SetPrimaryTextLayoutManager]
//   - [NSTextContentManager.TextLayoutManagers]: The array of text layout managers associated with this text content manager.
//   - [NSTextContentManager.AutomaticallySynchronizesTextLayoutManagers]: Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
//   - [NSTextContentManager.SetAutomaticallySynchronizesTextLayoutManagers]
//   - [NSTextContentManager.AddTextLayoutManager]: Adds the layout manager you provide to the list of layout managers.
//   - [NSTextContentManager.RemoveTextLayoutManager]: Removes the layout manager you specifiy from the list of layout managers.
//   - [NSTextContentManager.SynchronizeTextLayoutManagers]: Synchronizes changes to all nonprimary text layout managers.
//
// # Customizing and validating text elements
//
//   - [NSTextContentManager.Delegate]: The delegate for the content manager object.
//   - [NSTextContentManager.SetDelegate]
//
// # Finding a specific text element
//
//   - [NSTextContentManager.TextElementsForRange]: Returns an array of text elements that intersect with the range you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager
type NSTextContentManager struct {
	objectivec.Object
}

// NSTextContentManagerFromID constructs a [NSTextContentManager] from an objc.ID.
//
// An abstract class that defines the interface and a default implementation
// for managing the text document contents.
func NSTextContentManagerFromID(id objc.ID) NSTextContentManager {
	return NSTextContentManager{objectivec.Object{ID: id}}
}
// NOTE: NSTextContentManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextContentManager] class.
//
// # Creating a content manager
//
//   - [INSTextContentManager.InitWithCoder]: Creates a new content manager object from data in an unarchiver.
//
// # Controlling backing store synchronization
//
//   - [INSTextContentManager.AutomaticallySynchronizesToBackingStore]: Determines whether to automatically synchronize with the backing store when an editing transaction finishes.
//   - [INSTextContentManager.SetAutomaticallySynchronizesToBackingStore]
//
// # Performing transactions
//
//   - [INSTextContentManager.HasEditingTransaction]: Indicates there’s an active editing transaction from the primary text layout manager.
//   - [INSTextContentManager.PerformEditingTransactionUsingBlock]: Performs an editing transaction and invokes a block upon completion.
//   - [INSTextContentManager.RecordEditActionInRangeNewTextRange]: Records information about an edit action to the transaction.
//
// # Working with layout managers
//
//   - [INSTextContentManager.PrimaryTextLayoutManager]: The primary text layout manager for this content.
//   - [INSTextContentManager.SetPrimaryTextLayoutManager]
//   - [INSTextContentManager.TextLayoutManagers]: The array of text layout managers associated with this text content manager.
//   - [INSTextContentManager.AutomaticallySynchronizesTextLayoutManagers]: Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
//   - [INSTextContentManager.SetAutomaticallySynchronizesTextLayoutManagers]
//   - [INSTextContentManager.AddTextLayoutManager]: Adds the layout manager you provide to the list of layout managers.
//   - [INSTextContentManager.RemoveTextLayoutManager]: Removes the layout manager you specifiy from the list of layout managers.
//   - [INSTextContentManager.SynchronizeTextLayoutManagers]: Synchronizes changes to all nonprimary text layout managers.
//
// # Customizing and validating text elements
//
//   - [INSTextContentManager.Delegate]: The delegate for the content manager object.
//   - [INSTextContentManager.SetDelegate]
//
// # Finding a specific text element
//
//   - [INSTextContentManager.TextElementsForRange]: Returns an array of text elements that intersect with the range you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager
type INSTextContentManager interface {
	objectivec.IObject
	NSTextElementProvider

	// Topic: Creating a content manager

	// Creates a new content manager object from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSTextContentManager

	// Topic: Controlling backing store synchronization

	// Determines whether to automatically synchronize with the backing store when an editing transaction finishes.
	AutomaticallySynchronizesToBackingStore() bool
	SetAutomaticallySynchronizesToBackingStore(value bool)

	// Topic: Performing transactions

	// Indicates there’s an active editing transaction from the primary text layout manager.
	HasEditingTransaction() bool
	// Performs an editing transaction and invokes a block upon completion.
	PerformEditingTransactionUsingBlock(transaction VoidHandler)
	// Records information about an edit action to the transaction.
	RecordEditActionInRangeNewTextRange(originalTextRange INSTextRange, newTextRange INSTextRange)

	// Topic: Working with layout managers

	// The primary text layout manager for this content.
	PrimaryTextLayoutManager() INSTextLayoutManager
	SetPrimaryTextLayoutManager(value INSTextLayoutManager)
	// The array of text layout managers associated with this text content manager.
	TextLayoutManagers() []NSTextLayoutManager
	// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
	AutomaticallySynchronizesTextLayoutManagers() bool
	SetAutomaticallySynchronizesTextLayoutManagers(value bool)
	// Adds the layout manager you provide to the list of layout managers.
	AddTextLayoutManager(textLayoutManager INSTextLayoutManager)
	// Removes the layout manager you specifiy from the list of layout managers.
	RemoveTextLayoutManager(textLayoutManager INSTextLayoutManager)
	// Synchronizes changes to all nonprimary text layout managers.
	SynchronizeTextLayoutManagers(completionHandler ErrorHandler)

	// Topic: Customizing and validating text elements

	// The delegate for the content manager object.
	Delegate() NSTextContentManagerDelegate
	SetDelegate(value NSTextContentManagerDelegate)

	// Topic: Finding a specific text element

	// Returns an array of text elements that intersect with the range you specify.
	TextElementsForRange(range_ INSTextRange) []NSTextElement

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextContentManager) Init() NSTextContentManager {
	rv := objc.Send[NSTextContentManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextContentManager) Autorelease() NSTextContentManager {
	rv := objc.Send[NSTextContentManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextContentManager creates a new NSTextContentManager instance.
func NewNSTextContentManager() NSTextContentManager {
	class := getNSTextContentManagerClass()
	rv := objc.Send[NSTextContentManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new content manager object from data in an unarchiver.
//
// coder: An unachiver that conforms to the [NSCoder] class.
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/init(coder:)
func NewTextContentManagerWithCoder(coder foundation.INSCoder) NSTextContentManager {
	instance := getNSTextContentManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextContentManagerFromID(rv)
}

// Creates a new content manager object from data in an unarchiver.
//
// coder: An unachiver that conforms to the [NSCoder] class.
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/init(coder:)
func (t NSTextContentManager) InitWithCoder(coder foundation.INSCoder) NSTextContentManager {
	rv := objc.Send[NSTextContentManager](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Performs an editing transaction and invokes a block upon completion.
//
// transaction: The editing transaction.
//
// # Discussion
// 
// The primary [NSTextLayoutManager] controlling the active editing
// transaction invokes this method. It’s possible to nest multiple editing
// transactions. The outer most transaction toggles `hasEditingTransaction`
// and sends synchronization messages if enabled after invoking a transaction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/performEditingTransaction(_:)
func (t NSTextContentManager) PerformEditingTransactionUsingBlock(transaction VoidHandler) {
_block0, _ := NewVoidBlock(transaction)
	objc.Send[objc.ID](t.ID, objc.Sel("performEditingTransactionUsingBlock:"), _block0)
}
// Records information about an edit action to the transaction.
//
// originalTextRange: The range before the action.
//
// newTextRange: The corresponding range after the action.
//
// # Discussion
// 
// The concrete subclass invokes this method for each edit action.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/recordEditAction(in:newTextRange:)
func (t NSTextContentManager) RecordEditActionInRangeNewTextRange(originalTextRange INSTextRange, newTextRange INSTextRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("recordEditActionInRange:newTextRange:"), originalTextRange, newTextRange)
}
// Adds the layout manager you provide to the list of layout managers.
//
// textLayoutManager: The layout manager to add.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/addTextLayoutManager(_:)
func (t NSTextContentManager) AddTextLayoutManager(textLayoutManager INSTextLayoutManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("addTextLayoutManager:"), textLayoutManager)
}
// Removes the layout manager you specifiy from the list of layout managers.
//
// textLayoutManager: The layout manager to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/removeTextLayoutManager(_:)
func (t NSTextContentManager) RemoveTextLayoutManager(textLayoutManager INSTextLayoutManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeTextLayoutManager:"), textLayoutManager)
}
// Synchronizes changes to all nonprimary text layout managers.
//
// completionHandler: A completion handler that runs on success, or to handle error conditions.
//
// # Discussion
// 
// If `completionHandler` is `nil`, this method performs the operation
// synchronously. The framework passes any error to the `completionHandler`.
// The method blocks (or fails, if synchronous) when there’s an active
// transaction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/synchronizeTextLayoutManagers(_:)
func (t NSTextContentManager) SynchronizeTextLayoutManagers(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("synchronizeTextLayoutManagers:"), _block0)
}
// Returns an array of text elements that intersect with the range you
// specify.
//
// range: An [NSTextRange] that describes the range of text to process.
//
// # Return Value
// 
// An array of [NSTextElement].
//
// # Discussion
// 
// This method can return a set of elements that don’t fill the entire range
// if the entire range isn’t synchronously available. Uses
// [EnumerateTextElementsFromLocationOptionsUsingBlock] to fill the array.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/textElements(for:)
func (t NSTextContentManager) TextElementsForRange(range_ INSTextRange) []NSTextElement {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textElementsForRange:"), range_)
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextElement {
		return NSTextElementFromID(id)
	})
}
// A method you implement if the location backing store requires manual
// adjustment after editing.
//
// textRange: An [NSTextRange] that the method adjusts.
//
// forEditingTextSelection: A Boolean value that indicates if `textRange` is for the text selection
// associated with the edit session.
//
// # Return Value
// 
// When `textRange` is intersecting or following the current edited range, the
// method returns the range adjusted for the modification in the editing
// session. Returns `nil`, when no adjustment necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/adjustedRange(from:forEditingTextSelection:)
func (t NSTextContentManager) AdjustedRangeFromRangeForEditingTextSelection(textRange INSTextRange, forEditingTextSelection bool) INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("adjustedRangeFromRange:forEditingTextSelection:"), textRange, forEditingTextSelection)
	return NSTextRangeFromID(rv)
}
// Enumerates text elements starting at the text location you provide.
//
// textLocation: The [NSTextLocation] at which to start the enumeration.
//
// options: One of the possible [NSTextElementProviderEnumerationOptions] directions.
//
// block: A block you use to evaluate whether to continue the enumeration or tell the
// method to stop. Return `false` to end the enumeration process.
//
// # Return Value
// 
// An [NSTextLocation].
//
// # Discussion
// 
// If `textLocation` is `nil`, the method uses `documentRange.Location()` for
// forward enumeration and `documentRange.EndLocation()` for reverse
// enumeration. When enumerating backward, the method starts with the element
// preceding the one containing `textLocation`. If enumerated at least one
// element, it returns the edge of the enumerated range.
// 
// The enumerated range might not match the range of the last element
// returned. It enumerates the elements in the sequence, but it can skip a
// range (it can limit the maximum number of text elements enumerated for a
// single invocation or hide some elements from the layout).
// 
// Returning [NO] or `false` from block breaks out of the enumeration.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/enumerateTextElements(from:options:using:)
func (t NSTextContentManager) EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation NSTextLocation, options NSTextContentManagerEnumerationOptions, block TextElementHandler) NSTextLocation {
_block2, _ := NewTextElementBlock(block)
	rv := objc.Send[objc.ID](t.ID, objc.Sel("enumerateTextElementsFromLocation:options:usingBlock:"), textLocation, options, _block2)
	return NSTextLocationObjectFromID(rv)
}
// Returns a new location from location with offset you provide.
//
// location: An [NSTextLocation] in the text element.
//
// offset: An offset of the number of characters to or from `location`.
//
// # Return Value
// 
// An new [NSTextLocation], or `nil` of the offset exceeds the bounds of the
// text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/location(_:offsetBy:)
func (t NSTextContentManager) LocationFromLocationWithOffset(location NSTextLocation, offset int) NSTextLocation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("locationFromLocation:withOffset:"), location, offset)
	return NSTextLocationObjectFromID(rv)
}
// Returns the offset between the two specified locations.
//
// from: A starting location.
//
// to: An ending location.
//
// # Return Value
// 
// An [Integer] that represents the offset between the starting and ending
// locations.
//
// # Discussion
// 
// The return value could be positive or negative. This method can return
// [NSNotFound] when the method can’t represent an offset as an integer
// value. This can occur, for example, if the locations aren’t in the same
// document).
//
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-4qp9h
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/offset(from:to:)
func (t NSTextContentManager) OffsetFromLocationToLocation(from NSTextLocation, to NSTextLocation) int {
	rv := objc.Send[int](t.ID, objc.Sel("offsetFromLocation:toLocation:"), from, to)
	return rv
}
// Replaces the characters specified by range with the text elements you
// provide.
//
// range: An [NSTextRange].
//
// textElements: The elements to replace that characters at `range`.
//
// # Discussion
// 
// If the edges of `range` aren’t at existing element range boundaries, the
// method either splits the element if it allows the operation (for example,
// [NSTextParagraph]), or the adjusts the replacement range.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/replaceContents(in:with:)
func (t NSTextContentManager) ReplaceContentsInRangeWithTextElements(range_ INSTextRange, textElements []NSTextElement) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceContentsInRange:withTextElements:"), range_, objectivec.IObjectSliceToNSArray(textElements))
}
// Synchronizes changes to the backing store.
//
// completionHandler: A completion handler to run upon successful completion, or to process an
// error upon failure.
//
// # Discussion
// 
// If `completionHandler` is `nil`, performs the operation synchronously. The
// `completionHandler` gets passed `error` if the synchronization fails. It
// should block (or fails if synchronous) when there’s an active
// transaction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/synchronizeToBackingStore(_:)
func (t NSTextContentManager) SynchronizeToBackingStore(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](t.ID, objc.Sel("synchronizeToBackingStore:"), _block0)
}
func (t NSTextContentManager) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Determines whether to automatically synchronize with the backing store when
// an editing transaction finishes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesToBackingStore
func (t NSTextContentManager) AutomaticallySynchronizesToBackingStore() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("automaticallySynchronizesToBackingStore"))
	return rv
}
func (t NSTextContentManager) SetAutomaticallySynchronizesToBackingStore(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticallySynchronizesToBackingStore:"), value)
}
// Indicates there’s an active editing transaction from the primary text
// layout manager.
//
// # Discussion
// 
// When this property is `true`, there’s an active editing transaction from
// the `primaryTextLayoutManager`. The synchronization operations to
// nonprimary text layout managers and the backing store block (or fail when
// synchronous) while this property is `true`. Avoid accessing the elements
// from a nonprimary text layout manager while this values is `true`.
// 
// This property is KVO-compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/hasEditingTransaction
func (t NSTextContentManager) HasEditingTransaction() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("hasEditingTransaction"))
	return rv
}
// The primary text layout manager for this content.
//
// # Discussion
// 
// Setting this property to an [NSTextLayoutManager] not in
// `textLayoutManagers` resets it to `nil`. It automatically synchronizes
// pending edits before switching to a new primary object. The operation is
// synchronous.
// 
// This property is KVO-compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/primaryTextLayoutManager
func (t NSTextContentManager) PrimaryTextLayoutManager() INSTextLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("primaryTextLayoutManager"))
	return NSTextLayoutManagerFromID(objc.ID(rv))
}
func (t NSTextContentManager) SetPrimaryTextLayoutManager(value INSTextLayoutManager) {
	objc.Send[struct{}](t.ID, objc.Sel("setPrimaryTextLayoutManager:"), value)
}
// The array of text layout managers associated with this text content
// manager.
//
// # Discussion
// 
// This property is KVO-compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/textLayoutManagers
func (t NSTextContentManager) TextLayoutManagers() []NSTextLayoutManager {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textLayoutManagers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextLayoutManager {
		return NSTextLayoutManagerFromID(id)
	})
}
// Determines if the framework should automatically synchronize all text
// layout managers when exiting an editing transaction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesTextLayoutManagers
func (t NSTextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("automaticallySynchronizesTextLayoutManagers"))
	return rv
}
func (t NSTextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticallySynchronizesTextLayoutManagers:"), value)
}
// The delegate for the content manager object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContentManager/delegate
func (t NSTextContentManager) Delegate() NSTextContentManagerDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTextContentManagerDelegateObjectFromID(rv)
}
func (t NSTextContentManager) SetDelegate(value NSTextContentManagerDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
// Describes the starting and ending locations for the document.
//
// # Discussion
// 
// The subclass could use its own implementation of a location object
// conforming to [NSTextRange].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextElementProvider/documentRange
func (t NSTextContentManager) DocumentRange() INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("documentRange"))
	return NSTextRangeFromID(objc.ID(rv))
}

			// Protocol methods for NSTextElementProvider
			

// PerformEditingTransactionUsingBlockSync is a synchronous wrapper around [NSTextContentManager.PerformEditingTransactionUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTextContentManager) PerformEditingTransactionUsingBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.PerformEditingTransactionUsingBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SynchronizeTextLayoutManagersSync is a synchronous wrapper around [NSTextContentManager.SynchronizeTextLayoutManagers].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTextContentManager) SynchronizeTextLayoutManagersSync(ctx context.Context) error {
	done := make(chan error, 1)
	t.SynchronizeTextLayoutManagers(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateTextElementsFromLocationOptionsUsingBlockSync is a synchronous wrapper around [NSTextContentManager.EnumerateTextElementsFromLocationOptionsUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTextContentManager) EnumerateTextElementsFromLocationOptionsUsingBlockSync(ctx context.Context, textLocation NSTextLocation, options NSTextContentManagerEnumerationOptions) (*NSTextElement, error) {
	done := make(chan *NSTextElement, 1)
	t.EnumerateTextElementsFromLocationOptionsUsingBlock(textLocation, options, func(val *NSTextElement) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SynchronizeToBackingStoreSync is a synchronous wrapper around [NSTextContentManager.SynchronizeToBackingStore].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTextContentManager) SynchronizeToBackingStoreSync(ctx context.Context) error {
	done := make(chan error, 1)
	t.SynchronizeToBackingStore(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

