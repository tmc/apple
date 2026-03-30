// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of methods for responding to the life cycle events of the cloud-sharing service.
//
// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingServiceDelegate
type NSCloudSharingServiceDelegate interface {
	objectivec.IObject
	NSSharingServiceDelegate
}

// NSCloudSharingServiceDelegateObject wraps an existing Objective-C object that conforms to the NSCloudSharingServiceDelegate protocol.
type NSCloudSharingServiceDelegateObject struct {
	objectivec.Object
}

func (o NSCloudSharingServiceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCloudSharingServiceDelegateObjectFromID constructs a [NSCloudSharingServiceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCloudSharingServiceDelegateObjectFromID(id objc.ID) NSCloudSharingServiceDelegateObject {
	return NSCloudSharingServiceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate when the cloud-sharing service completes.
//
// sharingService: The cloud-sharing service that invokes this delegate method.
//
// items: The items the service is sharing.
//
// error: If the service can’t share the items, an error that provides information
// about the failure; otherwise, nil.
//
// # Discussion
//
// The cloud-sharing service invokes this method when the user finishes
// sharing or dismisses the service’s view controller. If you implement this
// method, the service calls it instead of the
// [SharingServiceDidFailToShareItemsError] and [SharingServiceDidShareItems]
// methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingServiceDelegate/sharingService(_:didCompleteForItems:error:)
func (o NSCloudSharingServiceDelegateObject) SharingServiceDidCompleteForItemsError(sharingService INSSharingService, items foundation.INSArray, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingService:didCompleteForItems:error:"), sharingService, items, error_)
}

// Tells the delegate when the cloud-sharing service saves the CloudKit share.
//
// sharingService: The cloud-sharing service that invokes this delegate method.
//
// share: The saved CloudKit share.
//
// # Discussion
//
// The cloud-sharing service invokes this method when it saves changes to the
// share. The `share` parameter is the most recent state of the share on the
// server.
//
// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingServiceDelegate/sharingService(_:didSave:)
func (o NSCloudSharingServiceDelegateObject) SharingServiceDidSaveShare(sharingService INSSharingService, share objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingService:didSaveShare:"), sharingService, share)
}

// Tells the delegate when the user stops sharing the CloudKit share.
//
// sharingService: The cloud-sharing service that invokes this delegate method.
//
// share: The share the user is no longer sharing.
//
// # Discussion
//
// The cloud-sharing service invokes this method after it deletes the share on
// the server. The `share` parameter is the most recent state of the share
// before the service deletes it.
//
// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingServiceDelegate/sharingService(_:didStopSharing:)
func (o NSCloudSharingServiceDelegateObject) SharingServiceDidStopSharing(sharingService INSSharingService, share objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingService:didStopSharing:"), sharingService, share)
}

// Asks the delegate for the participant options for the cloud-sharing
// service.
//
// cloudKitSharingService: The cloud-sharing service that invokes this delegate method.
//
// provider: The item provider that supplies the share to the service.
//
// # Discussion
//
// Use this method to specify whether the share is public or private. The
// options you return also determine any permissions that the share’s
// participants have. If you don’t implement this method, the cloud-sharing
// service uses the [NSCloudKitSharingServiceStandard] options.
//
// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingServiceDelegate/options(for:share:)
func (o NSCloudSharingServiceDelegateObject) OptionsForSharingServiceShareProvider(cloudKitSharingService INSSharingService, provider foundation.NSItemProvider) NSCloudKitSharingServiceOptions {
	rv := objc.Send[NSCloudKitSharingServiceOptions](o.ID, objc.Sel("optionsForSharingService:shareProvider:"), cloudKitSharingService, provider)
	return rv
}

// Invoked when the sharing service will share the specified items.
//
// sharingService: The sharing service.
//
// items: The items being shared.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:willShareItems:)
func (o NSCloudSharingServiceDelegateObject) SharingServiceWillShareItems(sharingService INSSharingService, items foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingService:willShareItems:"), sharingService, items)
}

// Invoked when the sharing service has finished sharing the items.
//
// sharingService: The sharing service.
//
// items: The items being shared.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:didShareItems:)
func (o NSCloudSharingServiceDelegateObject) SharingServiceDidShareItems(sharingService INSSharingService, items foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingService:didShareItems:"), sharingService, items)
}

// Invoked when the sharing service encountered an error when sharing items.
//
// sharingService: The sharing service.
//
// items: The items being shared.
//
// error: The error that was encountered when trying to share the item. If the error
// is [NSUserCancelledError], the user simply cancelled the error.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:didFailToShareItems:error:)
//
// [NSUserCancelledError]: https://developer.apple.com/documentation/Foundation/NSUserCancelledError-swift.var
func (o NSCloudSharingServiceDelegateObject) SharingServiceDidFailToShareItemsError(sharingService INSSharingService, items foundation.INSArray, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingService:didFailToShareItems:error:"), sharingService, items, error_)
}

// Invoked when the sharing service is performed and the sharing window is
// displayed, to present a transition between the original items and the
// sharing window.
//
// sharingService: The sharing service.
//
// item: The item being shared.
//
// # Return Value
//
// The rectangle, in screen coordinates, to display the transition.
//
// # Discussion
//
// The following is an example implementation of this method:
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:sourceFrameOnScreenForShareItem:)
func (o NSCloudSharingServiceDelegateObject) SharingServiceSourceFrameOnScreenForShareItem(sharingService INSSharingService, item objectivec.IObject) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("sharingService:sourceFrameOnScreenForShareItem:"), sharingService, item)
	return rv
}

// Invoked to allow returning a custom transition image when sharing an item.
//
// sharingService: The sharing service.
//
// item: The shared item.
//
// contentRect: The content rectangle is the frame of the actual content inside the
// transition image, excluding all decorations. For example, if the transition
// image is a QuickLook thumbnail, the value would be
// [QLThumbnailGetContentRect].
//
// # Return Value
//
// The image to display for the sharing transition. Its size should exactly
// match that of the original image.
//
// # Discussion
//
// A sample implementation of this method:
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:transitionImageForShareItem:contentRect:)
func (o NSCloudSharingServiceDelegateObject) SharingServiceTransitionImageForShareItemContentRect(sharingService INSSharingService, item objectivec.IObject, contentRect *corefoundation.CGRect) INSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sharingService:transitionImageForShareItem:contentRect:"), sharingService, item, contentRect)
	return NSImageFromID(rv)
}

// Returns the window that contained the share items.
//
// sharingService: The sharing service.
//
// items: The items being shared.
//
// sharingContentScope: The sharing content scope. The sharing scope can be modified from the
// default value of [NSSharingContentScopeItem] by setting a different value
// in the out parameter `sharingContentScope`. See
// [NSSharingService.SharingContentScope] for supported values.
//
// # Return Value
//
// The window of the shared items.
//
// # Discussion
//
// The following is an example implementation of this method. It changes the
// item scope, and returns the window the source image view is contained
// within.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:sourceWindowForShareItems:sharingContentScope:)
//
// [NSSharingService.SharingContentScope]: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope
func (o NSCloudSharingServiceDelegateObject) SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService INSSharingService, items foundation.INSArray, sharingContentScope NSSharingContentScope) INSWindow {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sharingService:sourceWindowForShareItems:sharingContentScope:"), sharingService, items, sharingContentScope)
	return NSWindowFromID(rv)
}

// The method invoked when the service is performed and wants to display its
// contents in a popover.
//
// # Discussion
//
// The delegate should return the view that will act as the anchor of the
// popover, along with the target rectangle within the bounds of that view and
// preferred edge of that rectangle for the popover to appear. The delegate
// may also return `nil`, indicating that there is no anchoring view currently
// available, in which case the service may attempt to display the service via
// some other means.
//
// The service named [NSSharingServiceNameCloudSharing] prefers to display
// itself using a popover anchored to an “Add People” or “Share”
// button. If no such button is available or visible, return `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/anchoringView(for:showRelativeTo:preferredEdge:)
func (o NSCloudSharingServiceDelegateObject) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService INSSharingService, positioningRect *corefoundation.CGRect, preferredEdge foundation.NSRectEdge) INSView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"), sharingService, positioningRect, preferredEdge)
	return NSViewFromID(rv)
}

// NSCloudSharingServiceDelegateConfig holds optional typed callbacks for [NSCloudSharingServiceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate
type NSCloudSharingServiceDelegateConfig struct {

	// Managing the Cloud-Sharing Service
	// SharingServiceDidCompleteForItemsError — Tells the delegate when the cloud-sharing service completes.
	SharingServiceDidCompleteForItemsError func(sharingService NSSharingService, items foundation.INSArray, error_ foundation.NSError)

	// Other Methods
	// OptionsForSharingServiceShareProvider — Asks the delegate for the participant options for the cloud-sharing service.
	OptionsForSharingServiceShareProvider func(cloudKitSharingService NSSharingService, provider foundation.NSItemProvider) NSCloudKitSharingServiceOptions
}

// NewNSCloudSharingServiceDelegate creates an Objective-C object implementing the [NSCloudSharingServiceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSCloudSharingServiceDelegateObject] satisfies the [NSCloudSharingServiceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nscloudsharingservicedelegate
func NewNSCloudSharingServiceDelegate(config NSCloudSharingServiceDelegateConfig) NSCloudSharingServiceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSCloudSharingServiceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SharingServiceDidCompleteForItemsError != nil {
		fn := config.SharingServiceDidCompleteForItemsError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sharingService:didCompleteForItems:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sharingServiceID objc.ID, itemsID objc.ID, error_ID objc.ID) {
				sharingService := NSSharingServiceFromID(sharingServiceID)
				items := foundation.NSArrayFromID(itemsID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(sharingService, items, error_)
			},
		})
	}

	if config.OptionsForSharingServiceShareProvider != nil {
		fn := config.OptionsForSharingServiceShareProvider
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("optionsForSharingService:shareProvider:"),
			Fn: func(self objc.ID, _cmd objc.SEL, cloudKitSharingServiceID objc.ID, providerID objc.ID) NSCloudKitSharingServiceOptions {
				cloudKitSharingService := NSSharingServiceFromID(cloudKitSharingServiceID)
				provider := foundation.NSItemProviderFromID(providerID)
				return fn(cloudKitSharingService, provider)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSCloudSharingServiceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSCloudSharingServiceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSCloudSharingServiceDelegateObjectFromID(instance)
}
