// Code generated from Apple documentation. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// MIDICIDiscoveredNodeArrayHandler handles A closure the system calls when a MIDI-CI node discovery request is complete.
//
// Used by:
//   - [MIDICIDiscoveryManager.DiscoverWithHandler]
type MIDICIDiscoveredNodeArrayHandler = func(*[]MIDICIDiscoveredNode)

// NewMIDICIDiscoveredNodeArrayBlock wraps a Go [MIDICIDiscoveredNodeArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MIDICIDiscoveryManager.DiscoverWithHandler]
func NewMIDICIDiscoveredNodeArrayBlock(handler MIDICIDiscoveredNodeArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *[]MIDICIDiscoveredNode
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]MIDICIDiscoveredNode, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = MIDICIDiscoveredNodeFromID(item.GetID())
			}
			result = &res
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDICIDiscoveryResponseBlock handles A block the system calls when a MIDI-CI node discovery request completes.

// MIDICIProfileChangedBlock handles A block the system calls to indicate it has enabled or disabled a profile.

// NewMIDICIProfileChangedBlock wraps a Go [MIDICIProfileChangedBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMIDICIProfileChangedBlock(handler MIDICIProfileChangedBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MIDICISession, extra0 uint32, extra1 MIDICIProfile, extra2 int32) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDICIProfileSpecificDataBlock handles A block the system calls when a MIDI-CI session or responder receives profile-specific data.

// NewMIDICIProfileSpecificDataBlock wraps a Go [MIDICIProfileSpecificDataBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMIDICIProfileSpecificDataBlock(handler MIDICIProfileSpecificDataBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MIDICISession, extra0 uint32, extra1 MIDICIProfile, extra2 foundation.NSData) {
		handler(primitive, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDICISessionDisconnectBlock handles A block the system calls when a MIDI-CI session disconnects.

// NewMIDICISessionDisconnectBlock wraps a Go [MIDICISessionDisconnectBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMIDICISessionDisconnectBlock(handler MIDICISessionDisconnectBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MIDICISession, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDICISessionErrorHandler handles A block the system calls when you disconnect from the session.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MIDICISession.InitWithDiscoveredNodeDataReadyHandlerDisconnectHandler]
type MIDICISessionErrorHandler = func(*MIDICISession, error)

// NewMIDICISessionErrorBlock wraps a Go [MIDICISessionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MIDICISession.InitWithDiscoveredNodeDataReadyHandlerDisconnectHandler]
func NewMIDICISessionErrorBlock(handler MIDICISessionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MIDICISession
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MIDICISessionFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDICISessionUint32MIDICIProfileDataHandler is the signature for a completion handler block.
type MIDICISessionUint32MIDICIProfileDataHandler = func(*MIDICISession, uint32, *MIDICIProfile, *foundation.NSData)

// NewMIDICISessionUint32MIDICIProfileDataBlock wraps a Go [MIDICISessionUint32MIDICIProfileDataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMIDICISessionUint32MIDICIProfileDataBlock(handler MIDICISessionUint32MIDICIProfileDataHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 uint32, extra1ID objc.ID, extra2ID objc.ID) {
		var result *MIDICISession
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MIDICISessionFromID(resultID)
			result = &v
		}
		var extra1 *MIDICIProfile
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := MIDICIProfileFromID(extra1ID)
			extra1 = &v
		}
		var extra2 *foundation.NSData
		if extra2ID != 0 {
			objc.Send[objc.ID](extra2ID, objc.Sel("retain"))
			v := foundation.NSDataFromID(extra2ID)
			extra2 = &v
		}
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDICISessionUint32MIDICIProfileInt32Handler is the signature for a completion handler block.
type MIDICISessionUint32MIDICIProfileInt32Handler = func(*MIDICISession, uint32, *MIDICIProfile, int32)

// NewMIDICISessionUint32MIDICIProfileInt32Block wraps a Go [MIDICISessionUint32MIDICIProfileInt32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMIDICISessionUint32MIDICIProfileInt32Block(handler MIDICISessionUint32MIDICIProfileInt32Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 uint32, extra1ID objc.ID, extra2 int32) {
		var result *MIDICISession
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MIDICISessionFromID(resultID)
			result = &v
		}
		var extra1 *MIDICIProfile
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := MIDICIProfileFromID(extra1ID)
			extra1 = &v
		}
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDIEventListUnsafePointerHandler handles completion with primitive and object results.
//
// Used by:
//   - [MIDIUMPMutableEndpoint.InitWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback]
type MIDIEventListUnsafePointerHandler = func(*MIDIEventList, unsafe.Pointer)

// NewMIDIEventListUnsafePointerBlock wraps a Go [MIDIEventListUnsafePointerHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MIDIUMPMutableEndpoint.InitWithNameDeviceInfoProductInstanceIDMIDIProtocolDestinationCallback]
func NewMIDIEventListUnsafePointerBlock(handler MIDIEventListUnsafePointerHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *MIDIEventList, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDINotifyBlock handles A callback block for notifying clients of state changes.

// NewMIDINotifyBlock wraps a Go [MIDINotifyBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMIDINotifyBlock(handler MIDINotifyBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal *MIDINotification) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// MIDIReceiveBlock handles A block receiving MIDI input that includes the incoming messages and a refCon to identify the source.

// NewMIDIReceiveBlock wraps a Go [MIDIReceiveBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMIDIReceiveBlock(handler MIDIReceiveBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *MIDIEventList, extra0 unsafe.Pointer) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A block the system calls when the session’s data is ready.
//
// Used by:
//   - [MIDICISession.InitWithDiscoveredNodeDataReadyHandlerDisconnectHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MIDICISession.InitWithDiscoveredNodeDataReadyHandlerDisconnectHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
