// Code generated from Apple documentation for WebKit. DO NOT EDIT.
//go:build ios
// +build ios

package webkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Performs the command associated with the given key command.
//
// keyCommand: The key command received by the first responder.
//
// keyCommand is a [*uikit.UIKeyCommand].
//
// # Return Value
//
// Returns [YES] if a command corresponding to the UIKeyCommand was found and
// performed, [NO] otherwise.
//
// # Discussion
//
// This method checks for a command corresponding to the provided
// [UIKeyCommand] and performs it, if available. The app should use this
// method to perform any extension commands at an appropriate time in the
// app’s responder object that handles the [PerformCommandForKeyCommand]
// action.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext/performCommand(for:)-25rd1
//
// [UIKeyCommand]: https://developer.apple.com/documentation/UIKit/UIKeyCommand
func (w WKWebExtensionContext) PerformCommandForKeyCommand(keyCommand objectivec.IObject) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("performCommandForKeyCommand:"), keyCommand)
	return rv
}
