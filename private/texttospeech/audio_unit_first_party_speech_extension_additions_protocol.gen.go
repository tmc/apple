// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AUAudioUnit_FirstPartySpeechExtensionAdditions protocol.
type AUAudioUnit_FirstPartySpeechExtensionAdditions interface {
	objectivec.IObject

	// _hostAuditToken protocol.
	_hostAuditToken() unsafe.Pointer

	// Set_hostAuditToken protocol.
	Set_hostAuditToken(token unsafe.Pointer)
}

// AUAudioUnit_FirstPartySpeechExtensionAdditionsObject wraps an existing Objective-C object that conforms to the AUAudioUnit_FirstPartySpeechExtensionAdditions protocol.
type AUAudioUnit_FirstPartySpeechExtensionAdditionsObject struct {
	objectivec.Object
}

func (o AUAudioUnit_FirstPartySpeechExtensionAdditionsObject) BaseObject() objectivec.Object {
	return o.Object
}

// AUAudioUnit_FirstPartySpeechExtensionAdditionsObjectFromID constructs a [AUAudioUnit_FirstPartySpeechExtensionAdditionsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AUAudioUnit_FirstPartySpeechExtensionAdditionsObjectFromID(id objc.ID) AUAudioUnit_FirstPartySpeechExtensionAdditionsObject {
	return AUAudioUnit_FirstPartySpeechExtensionAdditionsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o AUAudioUnit_FirstPartySpeechExtensionAdditionsObject) _hostAuditToken() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("_hostAuditToken"))
	return rv
}
func (o AUAudioUnit_FirstPartySpeechExtensionAdditionsObject) Set_hostAuditToken(token unsafe.Pointer) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("set_hostAuditToken:"), token)
}
