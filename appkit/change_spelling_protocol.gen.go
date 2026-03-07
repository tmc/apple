// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that responder objects can implement to correct a misspelled word.
//
// See: https://developer.apple.com/documentation/AppKit/NSChangeSpelling
type NSChangeSpelling interface {
	objectivec.IObject
}



// NSChangeSpellingObject wraps an existing Objective-C object that conforms to the NSChangeSpelling protocol.
type NSChangeSpellingObject struct {
	objectivec.Object
}
func (o NSChangeSpellingObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSChangeSpellingObjectFromID constructs a [NSChangeSpellingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSChangeSpellingObjectFromID(id objc.ID) NSChangeSpellingObject {
	return NSChangeSpellingObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Replaces the selected word in the receiver with a corrected version from
// the Spelling panel.
//
// # Discussion
// 
// This message is sent by the [NSSpellChecker] to the object whose text is
// being checked. To get the corrected spelling, ask `sender` for the string
// value of its selected cell (visible to the user as the text field in the
// Spelling panel). This method should replace the selected portion of the
// text with the string that it gets from the NSSpellChecker.
//
// See: https://developer.apple.com/documentation/AppKit/NSChangeSpelling/changeSpelling(_:)

func (o NSChangeSpellingObject) ChangeSpelling(sender objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("changeSpelling:"), sender)
	}







