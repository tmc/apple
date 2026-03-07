// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSFontChanging protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontChanging
type NSFontChanging interface {
	objectivec.IObject
}



// NSFontChangingObject wraps an existing Objective-C object that conforms to the NSFontChanging protocol.
type NSFontChangingObject struct {
	objectivec.Object
}
func (o NSFontChangingObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSFontChangingObjectFromID constructs a [NSFontChangingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFontChangingObjectFromID(id objc.ID) NSFontChangingObject {
	return NSFontChangingObject{
		Object: objectivec.ObjectFromID(id),
	}
}




//
// See: https://developer.apple.com/documentation/AppKit/NSFontChanging/changeFont(_:)

func (o NSFontChangingObject) ChangeFont(sender INSFontManager) {
	
	objc.Send[struct{}](o.ID, objc.Sel("changeFont:"), sender)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSFontChanging/validModesForFontPanel(_:)

func (o NSFontChangingObject) ValidModesForFontPanel(fontPanel NSFontPanel) NSFontPanelModeMask {
	
	rv := objc.Send[NSFontPanelModeMask](o.ID, objc.Sel("validModesForFontPanel:"), fontPanel)
	return rv
	}







