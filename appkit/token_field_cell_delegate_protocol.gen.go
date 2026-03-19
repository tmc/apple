// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of optional methods implemented by delegates of [NSTokenFieldCell](<doc://com.apple.appkit/documentation/AppKit/NSTokenFieldCell>) objects to work with tokenized strings.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate
type NSTokenFieldCellDelegate interface {
	objectivec.IObject
}

// NSTokenFieldCellDelegateObject wraps an existing Objective-C object that conforms to the NSTokenFieldCellDelegate protocol.
type NSTokenFieldCellDelegateObject struct {
	objectivec.Object
}
func (o NSTokenFieldCellDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTokenFieldCellDelegateObjectFromID constructs a [NSTokenFieldCellDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTokenFieldCellDelegateObjectFromID(id objc.ID) NSTokenFieldCellDelegateObject {
	return NSTokenFieldCellDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Allows the delegate to provide a string to be displayed as a proxy for the
// represented object.
//
// tokenFieldCell: The token field cell that sent the message.
//
// representedObject: A represented object of the token field cell.
//
// # Return Value
// 
// The string to be used as a proxy for `representedObject`. If you return
// `nil` or do not implement this method, then `representedObject` is
// displayed as the string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:displayStringForRepresentedObject:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell INSTokenFieldCell, representedObject objectivec.IObject) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenFieldCell:displayStringForRepresentedObject:"), tokenFieldCell, representedObject)
	return foundation.NSStringFromID(rv).String()
	}
// Allows the delegate to return the token style for editing the specified
// represented object.
//
// tokenFieldCell: The token field cell that sent the message.
//
// representedObject: A represented object of the token field cell.
//
// # Return Value
// 
// The style that should be used to display the representedObject. Possible
// values are shown in NSTokenStyle_Values.
//
// # Discussion
// 
// If the delegate implements this method and returns an
// [NSTokenField.TokenStyle] that differs from the style set by [TokenStyle],
// the value the delegate returns is preferred.
// 
// If the delegate does not implement this method, the token field cell’s
// [TokenStyle] is used.
//
// [NSTokenField.TokenStyle]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:styleForRepresentedObject:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellStyleForRepresentedObject(tokenFieldCell INSTokenFieldCell, representedObject objectivec.IObject) NSTokenStyle {
	
	rv := objc.Send[NSTokenStyle](o.ID, objc.Sel("tokenFieldCell:styleForRepresentedObject:"), tokenFieldCell, representedObject)
	return rv
	}
// Allows the delegate to provide an array of appropriate completions for the
// contents of the receiver.
//
// tokenFieldCell: The token field cell that sent the message.
//
// substring: The partial string that is to be completed.
//
// tokenIndex: The index of the token being edited.
//
// selectedIndex: Optionally, you can return by-reference an index into the returned array
// that specifies which of the completions should be initially selected. If
// none are to be selected, return by reference `-1`.
//
// # Return Value
// 
// An array of strings that are possible completions.
//
// # Discussion
// 
// If the delegate does not implement this method, no completions are
// provided.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:completionsForSubstring:indexOfToken:indexOfSelectedItem:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell INSTokenFieldCell, substring string, tokenIndex int, selectedIndex unsafe.Pointer) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), tokenFieldCell, objc.String(substring), tokenIndex, selectedIndex)
	return foundation.NSArrayFromID(rv)
	}
// Allows the delegate to provide a string to be edited as a proxy for the
// represented object.
//
// tokenFieldCell: The token field cell that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// A string that’s an editable proxy of the represented object, or `nil` if
// the token should not be editable.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:editingStringForRepresentedObject:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell INSTokenFieldCell, representedObject objectivec.IObject) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenFieldCell:editingStringForRepresentedObject:"), tokenFieldCell, representedObject)
	return foundation.NSStringFromID(rv).String()
	}
// Allows the delegate to provide a represented object for the string being
// edited.
//
// tokenFieldCell: The token field cell that sent the message.
//
// editingString: The edited string representation of a represented object.
//
// # Return Value
// 
// A represented object that is displayed rather than the editing string.
//
// # Discussion
// 
// If your application uses some object other than an [NSString] for their
// represented objects, you should return a new, autoreleased instance of that
// object from this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:representedObjectForEditing:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell INSTokenFieldCell, editingString string) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenFieldCell:representedObjectForEditingString:"), tokenFieldCell, objc.String(editingString))
	return objectivec.Object{ID: rv}
	}
// Allows the delegate to validate the tokens to be added to the receiver at a
// given index.
//
// tokenFieldCell: The token field cell that sent the message.
//
// tokens: An array of tokens to be inserted in the receiver at `index`.
//
// index: The index of the receiver in which the array of tokens to be validated
// (`tokens`) will be inserted.
//
// # Return Value
// 
// An array of validated tokens.
//
// # Discussion
// 
// The delegate can return the array unchanged or return a modified array of
// tokens. To reject the add completely, return an empty array. Returning
// `nil` causes an error.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:shouldAdd:at:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell INSTokenFieldCell, tokens foundation.INSArray, index uint) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenFieldCell:shouldAddObjects:atIndex:"), tokenFieldCell, tokens, index)
	return foundation.NSArrayFromID(rv)
	}
// Allows the delegate to return an array of objects representing the data
// read from `pboard`.
//
// tokenFieldCell: The token field cell that sent the message.
//
// pboard: The pasteboard from which to read the represented objects.
//
// # Return Value
// 
// An array of represented objects created from the pasteboard data.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:readFrom:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellReadFromPasteboard(tokenFieldCell INSTokenFieldCell, pboard INSPasteboard) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenFieldCell:readFromPasteboard:"), tokenFieldCell, pboard)
	return foundation.NSArrayFromID(rv)
	}
// Allows the delegate the opportunity to write custom pasteboard types to the
// pasteboard for the represented objects in `objects`.
//
// tokenFieldCell: The token field cell that sent the message.
//
// objects: An array of represented objects associated with the token field cell.
//
// pboard: The pasteboard to which to write the represented objects.
//
// # Return Value
// 
// [true] if the delegate writes the represented objects to the pasteboard,
// [false] otherwise. If [false], the token field writes the display strings
// to the [NSStringPboardType] pasteboard.
//
// [NSStringPboardType]: https://developer.apple.com/documentation/AppKit/NSStringPboardType
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:writeRepresentedObjects:to:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell INSTokenFieldCell, objects foundation.INSArray, pboard INSPasteboard) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tokenFieldCell:writeRepresentedObjects:toPasteboard:"), tokenFieldCell, objects, pboard)
	return rv
	}
// Allows the delegate to specify whether the represented object provides a
// menu.
//
// tokenFieldCell: The token field cell that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// [true] if the represented object has a menu, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// By default tokens have no menus.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:hasMenuForRepresentedObject:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell INSTokenFieldCell, representedObject objectivec.IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("tokenFieldCell:hasMenuForRepresentedObject:"), tokenFieldCell, representedObject)
	return rv
	}
// Allows the delegate to provide a menu for the specified represented object.
//
// tokenFieldCell: The token field cell that sent the message.
//
// representedObject: A represented object of the token field.
//
// # Return Value
// 
// The menu associated with the represented object.
//
// # Discussion
// 
// The returned menu should be autoreleased. By default tokens in a token
// field cell do not return menus.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCellDelegate/tokenFieldCell(_:menuForRepresentedObject:)
func (o NSTokenFieldCellDelegateObject) TokenFieldCellMenuForRepresentedObject(tokenFieldCell INSTokenFieldCell, representedObject objectivec.IObject) INSMenu {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("tokenFieldCell:menuForRepresentedObject:"), tokenFieldCell, representedObject)
	return NSMenuFromID(rv)
	}

// NSTokenFieldCellDelegateConfig holds optional typed callbacks for [NSTokenFieldCellDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstokenfieldcelldelegate
type NSTokenFieldCellDelegateConfig struct {

	// Other Methods
	// TokenFieldCellShouldAddObjectsAtIndex — Allows the delegate to validate the tokens to be added to the receiver at a given index.
	TokenFieldCellShouldAddObjectsAtIndex func(tokenFieldCell NSTokenFieldCell, tokens foundation.INSArray, index uint) foundation.INSArray
	// TokenFieldCellReadFromPasteboard — Allows the delegate to return an array of objects representing the data read from `pboard`.
	TokenFieldCellReadFromPasteboard func(tokenFieldCell NSTokenFieldCell, pboard NSPasteboard) foundation.INSArray
	// TokenFieldCellWriteRepresentedObjectsToPasteboard — Allows the delegate the opportunity to write custom pasteboard types to the pasteboard for the represented objects in `objects`.
	TokenFieldCellWriteRepresentedObjectsToPasteboard func(tokenFieldCell NSTokenFieldCell, objects foundation.INSArray, pboard NSPasteboard) bool
}

// NewNSTokenFieldCellDelegate creates an Objective-C object implementing the [NSTokenFieldCellDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTokenFieldCellDelegateObject] satisfies the [NSTokenFieldCellDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstokenfieldcelldelegate
func NewNSTokenFieldCellDelegate(config NSTokenFieldCellDelegateConfig) NSTokenFieldCellDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTokenFieldCellDelegate_%d", n)

	var methods []objc.MethodDef

	if config.TokenFieldCellShouldAddObjectsAtIndex != nil {
		fn := config.TokenFieldCellShouldAddObjectsAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenFieldCell:shouldAddObjects:atIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenFieldCellID objc.ID, tokensID objc.ID, index uint) objc.ID {
				tokenFieldCell := NSTokenFieldCellFromID(tokenFieldCellID)
				tokens := foundation.NSArrayFromID(tokensID)
				return fn(tokenFieldCell, tokens, index).GetID()
			},
		})
	}

	if config.TokenFieldCellReadFromPasteboard != nil {
		fn := config.TokenFieldCellReadFromPasteboard
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenFieldCell:readFromPasteboard:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenFieldCellID objc.ID, pboardID objc.ID) objc.ID {
				tokenFieldCell := NSTokenFieldCellFromID(tokenFieldCellID)
				pboard := NSPasteboardFromID(pboardID)
				return fn(tokenFieldCell, pboard).GetID()
			},
		})
	}

	if config.TokenFieldCellWriteRepresentedObjectsToPasteboard != nil {
		fn := config.TokenFieldCellWriteRepresentedObjectsToPasteboard
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("tokenFieldCell:writeRepresentedObjects:toPasteboard:"),
			Fn: func(self objc.ID, _cmd objc.SEL, tokenFieldCellID objc.ID, objectsID objc.ID, pboardID objc.ID) bool {
				tokenFieldCell := NSTokenFieldCellFromID(tokenFieldCellID)
				objects := foundation.NSArrayFromID(objectsID)
				pboard := NSPasteboardFromID(pboardID)
				return fn(tokenFieldCell, objects, pboard)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTokenFieldCellDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTokenFieldCellDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTokenFieldCellDelegateObjectFromID(instance)
}

