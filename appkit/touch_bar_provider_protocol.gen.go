// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that an object adopts to create a bar object in your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarProvider
type NSTouchBarProvider interface {
	objectivec.IObject

	// The property you implement to provide a Touch Bar object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTouchBarProvider/touchBar
	TouchBar() INSTouchBar
}

// NSTouchBarProviderObject wraps an existing Objective-C object that conforms to the NSTouchBarProvider protocol.
type NSTouchBarProviderObject struct {
	objectivec.Object
}
func (o NSTouchBarProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTouchBarProviderObjectFromID constructs a [NSTouchBarProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTouchBarProviderObjectFromID(id objc.ID) NSTouchBarProviderObject {
	return NSTouchBarProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The property you implement to provide a Touch Bar object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarProvider/touchBar

func (o NSTouchBarProviderObject) TouchBar() INSTouchBar {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("touchBar"))
	return NSTouchBarFromID(rv)
	}

