// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AUMessageChannel protocol.
type AUMessageChannel interface {
	objectivec.IObject

	// CallAudioUnit protocol.
	CallAudioUnit(unit objectivec.IObject) objectivec.IObject

	// CallHostBlock protocol.
	CallHostBlock() unsafe.Pointer

	// SetCallHostBlock protocol.
	SetCallHostBlock(block INSDictionaryINSDictionaryHandler)
}

// AUMessageChannelObject wraps an existing Objective-C object that conforms to the AUMessageChannel protocol.
type AUMessageChannelObject struct {
	objectivec.Object
}

func (o AUMessageChannelObject) BaseObject() objectivec.Object {
	return o.Object
}

// AUMessageChannelObjectFromID constructs a [AUMessageChannelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AUMessageChannelObjectFromID(id objc.ID) AUMessageChannelObject {
	return AUMessageChannelObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o AUMessageChannelObject) CallAudioUnit(unit objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("callAudioUnit:"), unit)
	return objectivec.Object{ID: rv}
}
func (o AUMessageChannelObject) CallHostBlock() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("callHostBlock"))
	return rv
}

var _aumessagechannelobject_setcallhostblock_p0_key byte

func (o AUMessageChannelObject) SetCallHostBlock(block INSDictionaryINSDictionaryHandler) {
	_block0, _cleanup0 := NewINSDictionaryINSDictionaryBlock(block)
	defer _cleanup0()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setCallHostBlock:"), objc.ID(_block0))
}
