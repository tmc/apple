// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// IOBluetoothRFCOMMChannelDelegate protocol.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate
type IOBluetoothRFCOMMChannelDelegate interface {
	objectivec.IObject
}

// IOBluetoothRFCOMMChannelDelegateObject wraps an existing Objective-C object that conforms to the IOBluetoothRFCOMMChannelDelegate protocol.
type IOBluetoothRFCOMMChannelDelegateObject struct {
	objectivec.Object
}

func (o IOBluetoothRFCOMMChannelDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothRFCOMMChannelDelegateObjectFromID constructs a [IOBluetoothRFCOMMChannelDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothRFCOMMChannelDelegateObjectFromID(id objc.ID) IOBluetoothRFCOMMChannelDelegateObject {
	return IOBluetoothRFCOMMChannelDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelClosed(_:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelClosed(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelClosed:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelControlSignalsChanged(_:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelControlSignalsChanged(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelControlSignalsChanged:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelData(_:data:length:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelDataDataLength(rfcommChannel IIOBluetoothRFCOMMChannel, dataPointer unsafe.Pointer, dataLength uintptr) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelData:data:length:"), rfcommChannel, dataPointer, dataLength)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelFlowControlChanged(_:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelFlowControlChanged(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelFlowControlChanged:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelOpenComplete(_:status:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelOpenCompleteStatus(rfcommChannel IIOBluetoothRFCOMMChannel, error_ kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelOpenComplete:status:"), rfcommChannel, error_)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelQueueSpaceAvailable(_:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelQueueSpaceAvailable(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelQueueSpaceAvailable:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelWriteComplete(_:refcon:status:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelWriteCompleteRefconStatus(rfcommChannel IIOBluetoothRFCOMMChannel, refcon uintptr, error_ kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelWriteComplete:refcon:status:"), rfcommChannel, refcon, error_)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelWriteComplete(_:refcon:status:bytesWritten:)
func (o IOBluetoothRFCOMMChannelDelegateObject) RfcommChannelWriteCompleteRefconStatusBytesWritten(rfcommChannel IIOBluetoothRFCOMMChannel, refcon uintptr, error_ kernel.IOReturn, length uintptr) {
	objc.Send[struct{}](o.ID, objc.Sel("rfcommChannelWriteComplete:refcon:status:bytesWritten:"), rfcommChannel, refcon, error_, length)
}

// IOBluetoothRFCOMMChannelDelegateConfig holds optional typed callbacks for [IOBluetoothRFCOMMChannelDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothrfcommchanneldelegate
type IOBluetoothRFCOMMChannelDelegateConfig struct {

	// Instance Methods
	RfcommChannelClosed                func(rfcommChannel IOBluetoothRFCOMMChannel)
	RfcommChannelControlSignalsChanged func(rfcommChannel IOBluetoothRFCOMMChannel)
	RfcommChannelFlowControlChanged    func(rfcommChannel IOBluetoothRFCOMMChannel)
	RfcommChannelOpenCompleteStatus    func(rfcommChannel IOBluetoothRFCOMMChannel, error_ kernel.IOReturn)
	RfcommChannelQueueSpaceAvailable   func(rfcommChannel IOBluetoothRFCOMMChannel)
}

// NewIOBluetoothRFCOMMChannelDelegate creates an Objective-C object implementing the [IOBluetoothRFCOMMChannelDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [IOBluetoothRFCOMMChannelDelegateObject] satisfies the [IOBluetoothRFCOMMChannelDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothrfcommchanneldelegate
func NewIOBluetoothRFCOMMChannelDelegate(config IOBluetoothRFCOMMChannelDelegateConfig) IOBluetoothRFCOMMChannelDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoIOBluetoothRFCOMMChannelDelegate_%d", n)

	var methods []objc.MethodDef

	if config.RfcommChannelClosed != nil {
		fn := config.RfcommChannelClosed
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("rfcommChannelClosed:"),
			Fn: func(self objc.ID, _cmd objc.SEL, rfcommChannelID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothRFCOMMChannelDelegate", "rfcommChannelClosed:")
					}
				}()
				rfcommChannel := IOBluetoothRFCOMMChannelFromID(rfcommChannelID)
				fn(rfcommChannel)
				_delegateDone = true
			},
		})
	}

	if config.RfcommChannelControlSignalsChanged != nil {
		fn := config.RfcommChannelControlSignalsChanged
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("rfcommChannelControlSignalsChanged:"),
			Fn: func(self objc.ID, _cmd objc.SEL, rfcommChannelID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothRFCOMMChannelDelegate", "rfcommChannelControlSignalsChanged:")
					}
				}()
				rfcommChannel := IOBluetoothRFCOMMChannelFromID(rfcommChannelID)
				fn(rfcommChannel)
				_delegateDone = true
			},
		})
	}

	if config.RfcommChannelFlowControlChanged != nil {
		fn := config.RfcommChannelFlowControlChanged
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("rfcommChannelFlowControlChanged:"),
			Fn: func(self objc.ID, _cmd objc.SEL, rfcommChannelID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothRFCOMMChannelDelegate", "rfcommChannelFlowControlChanged:")
					}
				}()
				rfcommChannel := IOBluetoothRFCOMMChannelFromID(rfcommChannelID)
				fn(rfcommChannel)
				_delegateDone = true
			},
		})
	}

	if config.RfcommChannelOpenCompleteStatus != nil {
		fn := config.RfcommChannelOpenCompleteStatus
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("rfcommChannelOpenComplete:status:"),
			Fn: func(self objc.ID, _cmd objc.SEL, rfcommChannelID objc.ID, error_ kernel.IOReturn) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothRFCOMMChannelDelegate", "rfcommChannelOpenComplete:status:")
					}
				}()
				rfcommChannel := IOBluetoothRFCOMMChannelFromID(rfcommChannelID)
				fn(rfcommChannel, error_)
				_delegateDone = true
			},
		})
	}

	if config.RfcommChannelQueueSpaceAvailable != nil {
		fn := config.RfcommChannelQueueSpaceAvailable
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("rfcommChannelQueueSpaceAvailable:"),
			Fn: func(self objc.ID, _cmd objc.SEL, rfcommChannelID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothRFCOMMChannelDelegate", "rfcommChannelQueueSpaceAvailable:")
					}
				}()
				rfcommChannel := IOBluetoothRFCOMMChannelFromID(rfcommChannelID)
				fn(rfcommChannel)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("IOBluetoothRFCOMMChannelDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewIOBluetoothRFCOMMChannelDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return IOBluetoothRFCOMMChannelDelegateObjectFromID(instance)
}
