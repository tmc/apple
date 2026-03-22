// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// A set of methods for managing interactions with an open or save panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate
type NSOpenSavePanelDelegate interface {
	objectivec.IObject
}

// NSOpenSavePanelDelegateObject wraps an existing Objective-C object that conforms to the NSOpenSavePanelDelegate protocol.
type NSOpenSavePanelDelegateObject struct {
	objectivec.Object
}
func (o NSOpenSavePanelDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSOpenSavePanelDelegateObjectFromID constructs a [NSOpenSavePanelDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSOpenSavePanelDelegateObjectFromID(id objc.ID) NSOpenSavePanelDelegateObject {
	return NSOpenSavePanelDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the user confirmed a filename choice by clicking
// Save in a Save panel.
//
// sender: The panel that reports the user’s confirmation of a filename choice.
//
// filename: The user’s filename choice.
//
// okFlag: If [true], the user clicked the Save button; if [false], the user did not.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The filename selected by the user, or `nil` if you want to cancel the save
// operation and leave the Save panel onscreen.
//
// # Discussion
// 
// The Save panel calls this method before appending any required filename
// extension information, and before it asks the user whether to replace an
// existing file, if a file with the specified name already exists in the
// given location.
// 
// The panel may call this method multiple times as the user types. When it
// does, the `okFlag` parameter is [false]. When the use confirms their
// choice, the value in the `okFlag` is [true]. If your delegate does
// extensive validation or puts up alerts, do so only when `okFlag` is [true].
// 
// In macOS 10.15 and later, you cannot change the filename that the user
// selects. Prior to macOS 10.15, you could sanitize the app’s filename to
// remove undesirable characters or limit its length only if your app wasn’t
// running in a sandbox.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:userEnteredFilename:confirmed:)
func (o NSOpenSavePanelDelegateObject) PanelUserEnteredFilenameConfirmed(sender objectivec.IObject, filename string, okFlag bool) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("panel:userEnteredFilename:confirmed:"), sender, objc.String(filename), okFlag)
	return foundation.NSStringFromID(rv).String()
	}
// Tells the delegate that the user changed the selection in the specified
// Save panel.
//
// sender: The panel whose selection changed.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panelSelectionDidChange(_:)
func (o NSOpenSavePanelDelegateObject) PanelSelectionDidChange(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("panelSelectionDidChange:"), sender)
	}
// Tells the delegate that the user changed the selected directory to the
// directory located at the specified URL.
//
// sender: The panel whose directory changed.
//
// url: The URL of the new directory, or `nil` if it can’t be represented by an
// [NSURL] object.
// //
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:didChangeToDirectoryURL:)
func (o NSOpenSavePanelDelegateObject) PanelDidChangeToDirectoryURL(sender objectivec.IObject, url foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("panel:didChangeToDirectoryURL:"), sender, url)
	}
// Tells the delegate that the Save panel is about to expand or collapse
// because the user clicked the disclosure triangle that displays or hides the
// file browser.
//
// sender: The panel that is about to expand or collapse.
//
// expanding: [true] specifies that the panel is expanding; [false] specifies that it is
// collapsing.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:willExpand:)
func (o NSOpenSavePanelDelegateObject) PanelWillExpand(sender objectivec.IObject, expanding bool) {
	objc.Send[struct{}](o.ID, objc.Sel("panel:willExpand:"), sender, expanding)
	}
// Asks the delegate whether the specified URL should be enabled in the Open
// panel.
//
// sender: The panel that asks whether the URL should be enabled.
//
// url: The URL for you to check.
//
// # Return Value
// 
// [true] if you want the panel to display the item at the specifed URL as
// enabled, or [false] to display it as disabled.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Save panels do not call this method; they always disable URLs.
// Implementations of this method should be fast to avoid stalling the user
// interface. Use [PanelValidateURLError] instead if processing will take a
// long time.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:shouldEnable:)
func (o NSOpenSavePanelDelegateObject) PanelShouldEnableURL(sender objectivec.IObject, url foundation.INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("panel:shouldEnableURL:"), sender, url)
	return rv
	}
// Asks the delegate to validate the URL for a file that the user selected.
//
// sender: The panel that requests URL validation.
//
// url: The URL for you to validate.
//
// # Discussion
// 
// Save panels call this method when the user clicks the Save button. Open
// panels call it when the user clicks the Open button. An Open panel calls
// this method once for each selected filename or directory.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:validate:)
func (o NSOpenSavePanelDelegateObject) PanelValidateURLError(sender objectivec.IObject, url foundation.INSURL) (bool, error) {
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("panel:validateURL:error:"), sender, url)
	if err != nil {
		return false, err
	}
	return rv, nil
	}
// [NSSavePanel]: Optional — Sent when the user changes the current type.
// [NSOpenPanel]: Not sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:didSelect:)
func (o NSOpenSavePanelDelegateObject) PanelDidSelectType(sender objectivec.IObject, type_ uniformtypeidentifiers.UTType) {
	objc.Send[struct{}](o.ID, objc.Sel("panel:didSelectType:"), sender, type_)
	}
// [NSSavePanel]: Optional — Sent when the content type popup is displayed
// and the save panel needs the display name for a type. If `nil` is returned,
// the save panel will display type’s `localizedDescription`. [NSOpenPanel]:
// Not sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenSavePanelDelegate/panel(_:displayNameFor:)
func (o NSOpenSavePanelDelegateObject) PanelDisplayNameForType(sender objectivec.IObject, type_ uniformtypeidentifiers.UTType) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("panel:displayNameForType:"), sender, type_)
	return foundation.NSStringFromID(rv).String()
	}

