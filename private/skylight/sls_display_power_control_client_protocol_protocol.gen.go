// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSDisplayPowerControlClientProtocol protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClientProtocol
type SLSDisplayPowerControlClientProtocol interface {
	objectivec.IObject

	// TerminateConnection protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClientProtocol/terminateConnection:
	TerminateConnection(connection []objectivec.IObject)
}

// SLSDisplayPowerControlClientProtocolObject wraps an existing Objective-C object that conforms to the SLSDisplayPowerControlClientProtocol protocol.
type SLSDisplayPowerControlClientProtocolObject struct {
	objectivec.Object
}

func (o SLSDisplayPowerControlClientProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSDisplayPowerControlClientProtocolObjectFromID constructs a [SLSDisplayPowerControlClientProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSDisplayPowerControlClientProtocolObjectFromID(id objc.ID) SLSDisplayPowerControlClientProtocolObject {
	return SLSDisplayPowerControlClientProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClientProtocol/requestStateChange:error:
func (o SLSDisplayPowerControlClientProtocolObject) RequestStateChangeError(change objectivec.IObject) (uint64, error) {
	rv, err := objc.SendWithError[uint64](o.ID, objc.Sel("requestStateChange:error:"), change)
	if err != nil {
		return 0, err
	}
	return rv, nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayPowerControlClientProtocol/terminateConnection:
func (o SLSDisplayPowerControlClientProtocolObject) TerminateConnection(connection []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("terminateConnection:"), objectivec.IObjectSliceToNSArray(connection))
}
