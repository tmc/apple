// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that enables the Ignore button in the Spelling panel to function properly.
//
// See: https://developer.apple.com/documentation/AppKit/NSIgnoreMisspelledWords
type NSIgnoreMisspelledWords interface {
	objectivec.IObject
}



// NSIgnoreMisspelledWordsObject wraps an existing Objective-C object that conforms to the NSIgnoreMisspelledWords protocol.
type NSIgnoreMisspelledWordsObject struct {
	objectivec.Object
}
func (o NSIgnoreMisspelledWordsObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSIgnoreMisspelledWordsObjectFromID constructs a [NSIgnoreMisspelledWordsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSIgnoreMisspelledWordsObjectFromID(id objc.ID) NSIgnoreMisspelledWordsObject {
	return NSIgnoreMisspelledWordsObject{
		Object: objectivec.ObjectFromID(id),
	}
}




//
// # Discussion
// 
// Implement this action method to allow an application to ignore misspelled
// words on a document-by-document basis. This message is sent by the
// NSSpellChecker instance to the object whose text is being checked.
// 
// Implement this method by using the code shown in the protocol description.
//
// See: https://developer.apple.com/documentation/AppKit/NSIgnoreMisspelledWords/ignoreSpelling(_:)

func (o NSIgnoreMisspelledWordsObject) IgnoreSpelling(sender objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("ignoreSpelling:"), sender)
	}







