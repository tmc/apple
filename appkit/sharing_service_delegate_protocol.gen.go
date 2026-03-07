// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A set of methods that you use to customize the position and animation of a share sheet, and to be notified whether the item is successfully shared.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate
type NSSharingServiceDelegate interface {
	objectivec.IObject
}



// NSSharingServiceDelegateObject wraps an existing Objective-C object that conforms to the NSSharingServiceDelegate protocol.
type NSSharingServiceDelegateObject struct {
	objectivec.Object
}
func (o NSSharingServiceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSSharingServiceDelegateObjectFromID constructs a [NSSharingServiceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSharingServiceDelegateObjectFromID(id objc.ID) NSSharingServiceDelegateObject {
	return NSSharingServiceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Invoked when the sharing service will share the specified items.
//
// sharingService: The sharing service.
//
// items: The items being shared.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:willShareItems:)

func (o NSSharingServiceDelegateObject) SharingServiceWillShareItems(sharingService INSSharingService, items foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("sharingService:willShareItems:"), sharingService, items)
	}

// Invoked when the sharing service has finished sharing the items.
//
// sharingService: The sharing service.
//
// items: The items being shared.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:didShareItems:)

func (o NSSharingServiceDelegateObject) SharingServiceDidShareItems(sharingService INSSharingService, items foundation.INSArray) {
	
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
// //
// [NSUserCancelledError]: https://developer.apple.com/documentation/Foundation/NSUserCancelledError-swift.var
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/sharingService(_:didFailToShareItems:error:)

func (o NSSharingServiceDelegateObject) SharingServiceDidFailToShareItemsError(sharingService INSSharingService, items foundation.INSArray, error_ foundation.INSError) {
	
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

func (o NSSharingServiceDelegateObject) SharingServiceSourceFrameOnScreenForShareItem(sharingService INSSharingService, item objectivec.IObject) corefoundation.CGRect {
	
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

func (o NSSharingServiceDelegateObject) SharingServiceTransitionImageForShareItemContentRect(sharingService INSSharingService, item objectivec.IObject, contentRect corefoundation.CGRect) INSImage {
	
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
// default value of [NSSharingService.SharingContentScope.item] by setting a
// different value in the out parameter `sharingContentScope`. See
// [NSSharingService.SharingContentScope] for supported values.
// //
// [NSSharingService.SharingContentScope.item]: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope/item
// [NSSharingService.SharingContentScope]: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope
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

func (o NSSharingServiceDelegateObject) SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService INSSharingService, items foundation.INSArray, sharingContentScope NSSharingContentScope) INSWindow {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sharingService:sourceWindowForShareItems:sharingContentScope:"), sharingService, items, sharingContentScope)
	return NSWindowFromID(rv)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServiceDelegate/anchoringView(for:showRelativeTo:preferredEdge:)

func (o NSSharingServiceDelegateObject) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService INSSharingService, positioningRect corefoundation.CGRect, preferredEdge foundation.NSRectEdge) INSView {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"), sharingService, positioningRect, preferredEdge)
	return NSViewFromID(rv)
	}





// NSSharingServiceDelegateConfig holds optional typed callbacks for [NSSharingServiceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate
type NSSharingServiceDelegateConfig struct {

	// Sharing Items
	// WillShareItems — Invoked when the sharing service will share the specified items.
	WillShareItems func(sharingService NSSharingService, items foundation.INSArray)
	// DidShareItems — Invoked when the sharing service has finished sharing the items.
	DidShareItems func(sharingService NSSharingService, items foundation.INSArray)
	// DidFailToShareItemsError — Invoked when the sharing service encountered an error when sharing items.
	DidFailToShareItemsError func(sharingService NSSharingService, items foundation.INSArray, error_ foundation.NSError)

	// Customizing Transition Animation
	// SourceWindowForShareItemsSharingContentScope — Returns the window that contained the share items.
	SourceWindowForShareItemsSharingContentScope func(sharingService NSSharingService, items foundation.INSArray, sharingContentScope NSSharingContentScope) NSWindow
}

// NewNSSharingServiceDelegate creates an Objective-C object implementing the [NSSharingServiceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSSharingServiceDelegateObject] satisfies the [NSSharingServiceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate
func NewNSSharingServiceDelegate(config NSSharingServiceDelegateConfig) NSSharingServiceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSSharingServiceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WillShareItems != nil {
		fn := config.WillShareItems
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sharingService:willShareItems:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sharingServiceID objc.ID, itemsID objc.ID) {
				sharingService := NSSharingServiceFromID(sharingServiceID)
				items := foundation.NSArrayFromID(itemsID)
				fn(sharingService, items)
			},
		})
	}

	if config.DidShareItems != nil {
		fn := config.DidShareItems
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sharingService:didShareItems:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sharingServiceID objc.ID, itemsID objc.ID) {
				sharingService := NSSharingServiceFromID(sharingServiceID)
				items := foundation.NSArrayFromID(itemsID)
				fn(sharingService, items)
			},
		})
	}

	if config.DidFailToShareItemsError != nil {
		fn := config.DidFailToShareItemsError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sharingService:didFailToShareItems:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sharingServiceID objc.ID, itemsID objc.ID, error_ID objc.ID) {
				sharingService := NSSharingServiceFromID(sharingServiceID)
				items := foundation.NSArrayFromID(itemsID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(sharingService, items, error_)
			},
		})
	}

	if config.SourceWindowForShareItemsSharingContentScope != nil {
		fn := config.SourceWindowForShareItemsSharingContentScope
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("sharingService:sourceWindowForShareItems:sharingContentScope:"),
			Fn: func(self objc.ID, _cmd objc.SEL, sharingServiceID objc.ID, itemsID objc.ID, sharingContentScope NSSharingContentScope) objc.ID {
				sharingService := NSSharingServiceFromID(sharingServiceID)
				items := foundation.NSArrayFromID(itemsID)
				return fn(sharingService, items, sharingContentScope).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSSharingServiceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSSharingServiceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSSharingServiceDelegateObjectFromID(instance)
}





