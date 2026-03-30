// Code generated from Apple documentation for WebKit. DO NOT EDIT.
//go:build ios
// +build ios

package webkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A key command representation of the web extension command for use in the
// responder chain.
//
// # Discussion
//
// Provides a [UIKeyCommand] instance representing the web extension command,
// ready for integration in the app.
//
// The property is `nil` if no shortcut is defined. Otherwise, the key command
// is fully configured with the necessary input key and modifier flags to
// perform the associated command upon activation. It can be included in a
// view controller or other responder’s [keyCommands] property, enabling
// keyboard activation and discoverability of the web extension command.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/keyCommand
//
// [UIKeyCommand]: https://developer.apple.com/documentation/UIKit/UIKeyCommand
// [keyCommands]: https://developer.apple.com/documentation/UIKit/UIResponder/keyCommands
func (w WKWebExtensionCommand) KeyCommand() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("keyCommand"))
	return objectivec.Object{ID: rv}
}
