// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that a sharing service picker item delegate uses to provide a list of items eligible for sharing.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItemDelegate
type NSSharingServicePickerTouchBarItemDelegate interface {
	objectivec.IObject
	NSSharingServicePickerDelegate

	// Asks the delegate for items that represent the objects to be shared.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItemDelegate/items(for:)
	ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem INSSharingServicePickerTouchBarItem) foundation.INSArray
}

// NSSharingServicePickerTouchBarItemDelegateObject wraps an existing Objective-C object that conforms to the NSSharingServicePickerTouchBarItemDelegate protocol.
type NSSharingServicePickerTouchBarItemDelegateObject struct {
	objectivec.Object
}
func (o NSSharingServicePickerTouchBarItemDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSharingServicePickerTouchBarItemDelegateObjectFromID constructs a [NSSharingServicePickerTouchBarItemDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSharingServicePickerTouchBarItemDelegateObjectFromID(id objc.ID) NSSharingServicePickerTouchBarItemDelegateObject {
	return NSSharingServicePickerTouchBarItemDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate for items that represent the objects to be shared.
//
// pickerTouchBarItem: The sharing service picker item that is requesting the items to be shared.
//
// # Return Value
// 
// An array of items that represents the objects to be shared. Each element of
// the array must either conform to the [NSPasteboardWriting] protocol, or be
// an [NSItemProvider].
//
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItemDelegate/items(for:)
func (o NSSharingServicePickerTouchBarItemDelegateObject) ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem INSSharingServicePickerTouchBarItem) foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemsForSharingServicePickerTouchBarItem:"), pickerTouchBarItem)
	return foundation.NSArrayFromID(rv)
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
func (o NSSharingServicePickerTouchBarItemDelegateObject) SharingServicePickerSharingServicesForItemsProposedSharingServices(sharingServicePicker INSSharingServicePicker, items foundation.INSArray, proposedServices []NSSharingService) []NSSharingService {
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
func (o NSSharingServicePickerTouchBarItemDelegateObject) SharingServicePickerDidChooseSharingService(sharingServicePicker INSSharingServicePicker, service INSSharingService) {
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
func (o NSSharingServicePickerTouchBarItemDelegateObject) SharingServicePickerDelegateForSharingService(sharingServicePicker INSSharingServicePicker, sharingService INSSharingService) NSSharingServiceDelegate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sharingServicePicker:delegateForSharingService:"), sharingServicePicker, sharingService)
	return NSSharingServiceDelegateObjectFromID(rv)
	}
// Used to specify the case where the share picker should not support some
// modes of sharing even if they are supported by the items being shared.
// Disabling all possible modes at the same time is not supported behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerDelegate/sharingServicePickerCollaborationModeRestrictions(_:)
func (o NSSharingServicePickerTouchBarItemDelegateObject) SharingServicePickerCollaborationModeRestrictions(sharingServicePicker INSSharingServicePicker) []NSSharingCollaborationModeRestriction {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("sharingServicePickerCollaborationModeRestrictions:"), sharingServicePicker)
	return objc.ConvertSlice(rv, func(id objc.ID) NSSharingCollaborationModeRestriction {
		return NSSharingCollaborationModeRestrictionFromID(id)
	})
	}

// NSSharingServicePickerTouchBarItemDelegateConfig holds optional typed callbacks for [NSSharingServicePickerTouchBarItemDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritemdelegate
type NSSharingServicePickerTouchBarItemDelegateConfig struct {

	// Other Methods
	// ItemsForSharingServicePickerTouchBarItem — Asks the delegate for items that represent the objects to be shared.
	ItemsForSharingServicePickerTouchBarItem func(pickerTouchBarItem NSSharingServicePickerTouchBarItem) foundation.INSArray
}

// NewNSSharingServicePickerTouchBarItemDelegate creates an Objective-C object implementing the [NSSharingServicePickerTouchBarItemDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSSharingServicePickerTouchBarItemDelegateObject] satisfies the [NSSharingServicePickerTouchBarItemDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssharingservicepickertouchbaritemdelegate
func NewNSSharingServicePickerTouchBarItemDelegate(config NSSharingServicePickerTouchBarItemDelegateConfig) NSSharingServicePickerTouchBarItemDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSSharingServicePickerTouchBarItemDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ItemsForSharingServicePickerTouchBarItem != nil {
		fn := config.ItemsForSharingServicePickerTouchBarItem
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("itemsForSharingServicePickerTouchBarItem:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pickerTouchBarItemID objc.ID) objc.ID {
				pickerTouchBarItem := NSSharingServicePickerTouchBarItemFromID(pickerTouchBarItemID)
				return fn(pickerTouchBarItem).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSSharingServicePickerTouchBarItemDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSSharingServicePickerTouchBarItemDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSSharingServicePickerTouchBarItemDelegateObjectFromID(instance)
}

