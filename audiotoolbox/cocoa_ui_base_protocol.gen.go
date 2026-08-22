// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AUCocoaUIBase protocol.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUCocoaUIBase
type AUCocoaUIBase interface {
	objectivec.IObject

	// InterfaceVersion protocol.
	//
	// See: https://developer.apple.com/documentation/AudioToolbox/AUCocoaUIBase/interfaceVersion
	InterfaceVersion() uint32

	// UiViewForAudioUnitWithSize protocol.
	//
	// See: https://developer.apple.com/documentation/AudioToolbox/AUCocoaUIBase/uiViewForAudioUnit:withSize:
	UiViewForAudioUnitWithSize(inAudioUnit AudioUnit, inPreferredSize corefoundation.CGSize) objectivec.IObject
}

// AUCocoaUIBaseObject wraps an existing Objective-C object that conforms to the AUCocoaUIBase protocol.
type AUCocoaUIBaseObject struct {
	objectivec.Object
}

func (o AUCocoaUIBaseObject) BaseObject() objectivec.Object {
	return o.Object
}

// AUCocoaUIBaseObjectFromID constructs a [AUCocoaUIBaseObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AUCocoaUIBaseObjectFromID(id objc.ID) AUCocoaUIBaseObject {
	return AUCocoaUIBaseObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUCocoaUIBase/interfaceVersion
func (o AUCocoaUIBaseObject) InterfaceVersion() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("interfaceVersion"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUCocoaUIBase/uiViewForAudioUnit:withSize:
func (o AUCocoaUIBaseObject) UiViewForAudioUnitWithSize(inAudioUnit AudioUnit, inPreferredSize corefoundation.CGSize) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uiViewForAudioUnit:withSize:"), inAudioUnit, inPreferredSize)
	return objectivec.Object{ID: rv}
}
