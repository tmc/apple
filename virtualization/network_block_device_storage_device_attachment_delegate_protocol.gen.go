// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Methods you implement to respond to changes to a network block device attachment.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachmentDelegate
type VZNetworkBlockDeviceStorageDeviceAttachmentDelegate interface {
	objectivec.IObject
}

// VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject wraps an existing Objective-C object that conforms to the VZNetworkBlockDeviceStorageDeviceAttachmentDelegate protocol.
type VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject struct {
	objectivec.Object
}
func (o VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObjectFromID constructs a [VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObjectFromID(id objc.ID) VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject {
	return VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The method the attachment object calls when the NBD client encounters a
// nonrecoverable error.
//
// attachment: The attachment object calling the delegate method.
//
// error: An [NSError] that describes the nonrecoverable error.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// After the attachment object calls this method, the NBD client is in a
// nonfunctional state.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachmentDelegate/attachment(_:didEncounterError:)
func (o VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject) AttachmentDidEncounterError(attachment IVZNetworkBlockDeviceStorageDeviceAttachment, error_ foundation.INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("attachment:didEncounterError:"), attachment, error_)
	}
// The method the attachment object calls when the NBD client successfully
// connects or reconnects with the server.
//
// attachment: The attachment object calling the delegate method.
//
// # Discussion
// 
// Connection with the server takes place when the VM is first started, and
// reconnection attempts take place when the connection times out or when the
// NBD client has encountered a recoverable error, such as an I/O error from
// the server. The Virtualization framework may call this method multiple
// times during a VM’s life cycle. Reconnections are transparent to the
// guest.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachmentDelegate/attachmentWasConnected(_:)
func (o VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject) AttachmentWasConnected(attachment IVZNetworkBlockDeviceStorageDeviceAttachment) {
	
	objc.Send[struct{}](o.ID, objc.Sel("attachmentWasConnected:"), attachment)
	}

// VZNetworkBlockDeviceStorageDeviceAttachmentDelegateConfig holds optional typed callbacks for [VZNetworkBlockDeviceStorageDeviceAttachmentDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachmentdelegate
type VZNetworkBlockDeviceStorageDeviceAttachmentDelegateConfig struct {

	// Responding to connectivity changes
	// AttachmentDidEncounterError — The method the attachment object calls when the NBD client encounters a nonrecoverable error.
	AttachmentDidEncounterError func(attachment VZNetworkBlockDeviceStorageDeviceAttachment, error_ foundation.NSError)
	// AttachmentWasConnected — The method the attachment object calls when the NBD client successfully connects or reconnects with the server.
	AttachmentWasConnected func(attachment VZNetworkBlockDeviceStorageDeviceAttachment)
}

// NewVZNetworkBlockDeviceStorageDeviceAttachmentDelegate creates an Objective-C object implementing the [VZNetworkBlockDeviceStorageDeviceAttachmentDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject] satisfies the [VZNetworkBlockDeviceStorageDeviceAttachmentDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/virtualization/vznetworkblockdevicestoragedeviceattachmentdelegate
func NewVZNetworkBlockDeviceStorageDeviceAttachmentDelegate(config VZNetworkBlockDeviceStorageDeviceAttachmentDelegateConfig) VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoVZNetworkBlockDeviceStorageDeviceAttachmentDelegate_%d", n)

	var methods []objc.MethodDef

	if config.AttachmentDidEncounterError != nil {
		fn := config.AttachmentDidEncounterError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("attachment:didEncounterError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, attachmentID objc.ID, error_ID objc.ID) {
				attachment := VZNetworkBlockDeviceStorageDeviceAttachmentFromID(attachmentID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(attachment, error_)
			},
		})
	}

	if config.AttachmentWasConnected != nil {
		fn := config.AttachmentWasConnected
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("attachmentWasConnected:"),
			Fn: func(self objc.ID, _cmd objc.SEL, attachmentID objc.ID) {
				attachment := VZNetworkBlockDeviceStorageDeviceAttachmentFromID(attachmentID)
				fn(attachment)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("VZNetworkBlockDeviceStorageDeviceAttachmentDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewVZNetworkBlockDeviceStorageDeviceAttachmentDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return VZNetworkBlockDeviceStorageDeviceAttachmentDelegateObjectFromID(instance)
}

