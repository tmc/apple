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

// IOBluetoothL2CAPChannelDelegate protocol.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate
type IOBluetoothL2CAPChannelDelegate interface {
	objectivec.IObject
}

// IOBluetoothL2CAPChannelDelegateObject wraps an existing Objective-C object that conforms to the IOBluetoothL2CAPChannelDelegate protocol.
type IOBluetoothL2CAPChannelDelegateObject struct {
	objectivec.Object
}

func (o IOBluetoothL2CAPChannelDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothL2CAPChannelDelegateObjectFromID constructs a [IOBluetoothL2CAPChannelDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothL2CAPChannelDelegateObjectFromID(id objc.ID) IOBluetoothL2CAPChannelDelegateObject {
	return IOBluetoothL2CAPChannelDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate/l2capChannelClosed(_:)
func (o IOBluetoothL2CAPChannelDelegateObject) L2capChannelClosed(l2capChannel IIOBluetoothL2CAPChannel) {
	objc.Send[struct{}](o.ID, objc.Sel("l2capChannelClosed:"), l2capChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate/l2capChannelData(_:data:length:)
func (o IOBluetoothL2CAPChannelDelegateObject) L2capChannelDataDataLength(l2capChannel IIOBluetoothL2CAPChannel, dataPointer unsafe.Pointer, dataLength uintptr) {
	objc.Send[struct{}](o.ID, objc.Sel("l2capChannelData:data:length:"), l2capChannel, dataPointer, dataLength)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate/l2capChannelOpenComplete(_:status:)
func (o IOBluetoothL2CAPChannelDelegateObject) L2capChannelOpenCompleteStatus(l2capChannel IIOBluetoothL2CAPChannel, error_ kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("l2capChannelOpenComplete:status:"), l2capChannel, error_)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate/l2capChannelQueueSpaceAvailable(_:)
func (o IOBluetoothL2CAPChannelDelegateObject) L2capChannelQueueSpaceAvailable(l2capChannel IIOBluetoothL2CAPChannel) {
	objc.Send[struct{}](o.ID, objc.Sel("l2capChannelQueueSpaceAvailable:"), l2capChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate/l2capChannelReconfigured(_:)
func (o IOBluetoothL2CAPChannelDelegateObject) L2capChannelReconfigured(l2capChannel IIOBluetoothL2CAPChannel) {
	objc.Send[struct{}](o.ID, objc.Sel("l2capChannelReconfigured:"), l2capChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDelegate/l2capChannelWriteComplete(_:refcon:status:)
func (o IOBluetoothL2CAPChannelDelegateObject) L2capChannelWriteCompleteRefconStatus(l2capChannel IIOBluetoothL2CAPChannel, refcon uintptr, error_ kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("l2capChannelWriteComplete:refcon:status:"), l2capChannel, refcon, error_)
}

// IOBluetoothL2CAPChannelDelegateConfig holds optional typed callbacks for [IOBluetoothL2CAPChannelDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothl2capchanneldelegate
type IOBluetoothL2CAPChannelDelegateConfig struct {

	// Instance Methods
	L2capChannelClosed              func(l2capChannel IOBluetoothL2CAPChannel)
	L2capChannelOpenCompleteStatus  func(l2capChannel IOBluetoothL2CAPChannel, error_ kernel.IOReturn)
	L2capChannelQueueSpaceAvailable func(l2capChannel IOBluetoothL2CAPChannel)
	L2capChannelReconfigured        func(l2capChannel IOBluetoothL2CAPChannel)
}

// NewIOBluetoothL2CAPChannelDelegate creates an Objective-C object implementing the [IOBluetoothL2CAPChannelDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [IOBluetoothL2CAPChannelDelegateObject] satisfies the [IOBluetoothL2CAPChannelDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothl2capchanneldelegate
func NewIOBluetoothL2CAPChannelDelegate(config IOBluetoothL2CAPChannelDelegateConfig) IOBluetoothL2CAPChannelDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoIOBluetoothL2CAPChannelDelegate_%d", n)

	var methods []objc.MethodDef

	if config.L2capChannelClosed != nil {
		fn := config.L2capChannelClosed
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("l2capChannelClosed:"),
			Fn: func(self objc.ID, _cmd objc.SEL, l2capChannelID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothL2CAPChannelDelegate", "l2capChannelClosed:")
					}
				}()
				l2capChannel := IOBluetoothL2CAPChannelFromID(l2capChannelID)
				fn(l2capChannel)
				_delegateDone = true
			},
		})
	}

	if config.L2capChannelOpenCompleteStatus != nil {
		fn := config.L2capChannelOpenCompleteStatus
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("l2capChannelOpenComplete:status:"),
			Fn: func(self objc.ID, _cmd objc.SEL, l2capChannelID objc.ID, error_ kernel.IOReturn) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothL2CAPChannelDelegate", "l2capChannelOpenComplete:status:")
					}
				}()
				l2capChannel := IOBluetoothL2CAPChannelFromID(l2capChannelID)
				fn(l2capChannel, error_)
				_delegateDone = true
			},
		})
	}

	if config.L2capChannelQueueSpaceAvailable != nil {
		fn := config.L2capChannelQueueSpaceAvailable
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("l2capChannelQueueSpaceAvailable:"),
			Fn: func(self objc.ID, _cmd objc.SEL, l2capChannelID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothL2CAPChannelDelegate", "l2capChannelQueueSpaceAvailable:")
					}
				}()
				l2capChannel := IOBluetoothL2CAPChannelFromID(l2capChannelID)
				fn(l2capChannel)
				_delegateDone = true
			},
		})
	}

	if config.L2capChannelReconfigured != nil {
		fn := config.L2capChannelReconfigured
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("l2capChannelReconfigured:"),
			Fn: func(self objc.ID, _cmd objc.SEL, l2capChannelID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothL2CAPChannelDelegate", "l2capChannelReconfigured:")
					}
				}()
				l2capChannel := IOBluetoothL2CAPChannelFromID(l2capChannelID)
				fn(l2capChannel)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("IOBluetoothL2CAPChannelDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewIOBluetoothL2CAPChannelDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return IOBluetoothL2CAPChannelDelegateObjectFromID(instance)
}
