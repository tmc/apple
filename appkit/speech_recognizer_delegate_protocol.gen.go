// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of optional methods implemented by delegates of [NSSpeechRecognizer](<doc://com.apple.appkit/documentation/AppKit/NSSpeechRecognizer>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizerDelegate
type NSSpeechRecognizerDelegate interface {
	objectivec.IObject
}

// NSSpeechRecognizerDelegateObject wraps an existing Objective-C object that conforms to the NSSpeechRecognizerDelegate protocol.
type NSSpeechRecognizerDelegateObject struct {
	objectivec.Object
}
func (o NSSpeechRecognizerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSpeechRecognizerDelegateObjectFromID constructs a [NSSpeechRecognizerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSpeechRecognizerDelegateObjectFromID(id objc.ID) NSSpeechRecognizerDelegateObject {
	return NSSpeechRecognizerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Invoked when the recognition engine has recognized the application command
// `command`.
//
// # Discussion
// 
// `command` is one of the strings from the array passed to [Commands]. The
// delegate typically evaluates which command was recognized and performs the
// related action.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpeechRecognizerDelegate/speechRecognizer(_:didRecognizeCommand:)

func (o NSSpeechRecognizerDelegateObject) SpeechRecognizerDidRecognizeCommand(sender INSSpeechRecognizer, command string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognizer:didRecognizeCommand:"), sender, objc.String(command))
	}

