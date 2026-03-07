// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// An interface for managing content for the macOS share sheet.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerDelegate
type NSSharingServicePickerDelegate interface {
	objectivec.IObject
}



// NSSharingServicePickerDelegateObject wraps an existing Objective-C object that conforms to the NSSharingServicePickerDelegate protocol.
type NSSharingServicePickerDelegateObject struct {
	objectivec.Object
}
func (o NSSharingServicePickerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSSharingServicePickerDelegateObjectFromID constructs a [NSSharingServicePickerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSharingServicePickerDelegateObjectFromID(id objc.ID) NSSharingServicePickerDelegateObject {
	return NSSharingServicePickerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Asks the delegate to specify which services to make available from the
// sharing service picker.
//
// sharingServicePicker: The sharing service picker.
//
// items: The items to share. Use the set of items to determine which services are
// relevant.
//
// proposedServices: The proposed services to include in the sharing service picker.
//
// # Return Value
// 
// An array of services to include in the sharing service picker.
//
// # Discussion
// 
// Use this method to remove default services, add custom services, or reorder
// the existing services before the picker appears onscreen. Unless you
// don’t intend to change the proposed services, create a new mutable array
// and fill it with the services that are appropriate for the specified set of
// items. The following example appends a custom [NSSharingService] object to
// the proposed list of services.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerDelegate/sharingServicePicker(_:sharingServicesForItems:proposedSharingServices:)

func (o NSSharingServicePickerDelegateObject) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker INSSharingServicePicker, items foundation.INSArray, proposedServices []NSSharingService) []NSSharingService {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"), sharingServicePicker, items, objectivec.IObjectSliceToNSArray(proposedServices))
	return objc.ConvertSlice(rv, func(id objc.ID) NSSharingService {
		return NSSharingServiceFromID(id)
	})
	}

// Tells the delegate that the person selected a sharing service for the
// current item.
//
// sharingServicePicker: The sharing service picker.
//
// service: The selected sharing service. Invoked to give the delegate to the sharing
// service that is about to be executed.
//
// # Discussion
// 
// After someone chooses a service, the sharing service picker calls this
// method to let you know which service they picked. The sharing service
// receives the item sometime after this method returns.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerDelegate/sharingServicePicker(_:didChoose:)

func (o NSSharingServicePickerDelegateObject) SharingServicePickerDidChooseSharingService(sharingServicePicker INSSharingServicePicker, service INSSharingService) {
	
	objc.Send[struct{}](o.ID, objc.Sel("sharingServicePicker:didChooseSharingService:"), sharingServicePicker, service)
	}

// Asks your delegate to provide an object that the selected sharing service
// can use as its delegate.
//
// sharingServicePicker: The sharing service picker.
//
// sharingService: The selected sharing service.
//
// # Return Value
// 
// An object that adopts the [NSSharingServiceDelegate] protocol.
//
// # Discussion
// 
// The sharing service assigns the returned object to its [Delegate] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerDelegate/sharingServicePicker(_:delegateFor:)

func (o NSSharingServicePickerDelegateObject) SharingServicePickerDelegateForSharingService(sharingServicePicker INSSharingServicePicker, sharingService INSSharingService) NSSharingServiceDelegate {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sharingServicePicker:delegateForSharingService:"), sharingServicePicker, sharingService)
	return NSSharingServiceDelegateObjectFromID(rv)
	}

// Used to specify the case where the share picker should not support some
// modes of sharing even if they are supported by the items being shared.
// Disabling all possible modes at the same time is not supported behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerDelegate/sharingServicePickerCollaborationModeRestrictions(_:)

func (o NSSharingServicePickerDelegateObject) SharingServicePickerCollaborationModeRestrictions(sharingServicePicker INSSharingServicePicker) []NSSharingCollaborationModeRestriction {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("sharingServicePickerCollaborationModeRestrictions:"), sharingServicePicker)
	return objc.ConvertSlice(rv, func(id objc.ID) NSSharingCollaborationModeRestriction {
		return NSSharingCollaborationModeRestrictionFromID(id)
	})
	}





// NSSharingServicePickerDelegateConfig holds optional typed callbacks for [NSSharingServicePickerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate
type NSSharingServicePickerDelegateConfig struct {

	// Other Methods
	// DidChooseSharingService — Tells the delegate that the person selected a sharing service for the current item.
	DidChooseSharingService func(sharingServicePicker NSSharingServicePicker, service NSSharingService)
}

// NewNSSharingServicePickerDelegate creates an Objective-C object implementing the [NSSharingServicePickerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSSharingServicePickerDelegateObject] satisfies the [NSSharingServicePickerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssharingservicepickerdelegate
func NewNSSharingServicePickerDelegate(config NSSharingServicePickerDelegateConfig) NSSharingServicePickerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSSharingServicePickerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DidChooseSharingService != nil {
		fn := config.DidChooseSharingService
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sharingServicePicker:didChooseSharingService:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sharingServicePickerID objc.ID, serviceID objc.ID) {
				sharingServicePicker := NSSharingServicePickerFromID(sharingServicePickerID)
				service := NSSharingServiceFromID(serviceID)
				fn(sharingServicePicker, service)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSSharingServicePickerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSSharingServicePickerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSSharingServicePickerDelegateObjectFromID(instance)
}





