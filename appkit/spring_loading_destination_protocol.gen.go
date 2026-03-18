// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that the destination object (or recipient) of a dragged object can implement to support spring-loading.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination
type NSSpringLoadingDestination interface {
	objectivec.IObject

	// Responds to the activation or deactivation of spring-loading on a destination.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/springLoadingActivated(_:draggingInfo:)
	SpringLoadingActivatedDraggingInfo(activated bool, draggingInfo NSDraggingInfo)

	// Updates the destination’s user interface to display a new highlighting style during a spring-loading operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/springLoadingHighlightChanged(_:)
	SpringLoadingHighlightChanged(draggingInfo NSDraggingInfo)
}

// NSSpringLoadingDestinationObject wraps an existing Objective-C object that conforms to the NSSpringLoadingDestination protocol.
type NSSpringLoadingDestinationObject struct {
	objectivec.Object
}
func (o NSSpringLoadingDestinationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSpringLoadingDestinationObjectFromID constructs a [NSSpringLoadingDestinationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSpringLoadingDestinationObjectFromID(id objc.ID) NSSpringLoadingDestinationObject {
	return NSSpringLoadingDestinationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Responds to the activation or deactivation of spring-loading on a
// destination.
//
// activated: A Boolean value indicating whether spring-loading has been activated on the
// destination. [true] indicates that spring-loading has been activated.
// [false] indicates that spring-loading has been deactivated.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// draggingInfo: An [NSDraggingInfo] object, which provides information about the drag
// event, including the dragged data.
//
// # Discussion
// 
// Typically, spring-loading is fully activated when a hover timeout occurs or
// the user finishes force clicking on a destination object to initiate
// spring-loading. In these cases, the `` method is only called once with an
// `activated` parameter value of [true].
// 
// However, if the destination is configured with continuous activation
// ([NSSpringLoadingOptions] was set to
// [NSSpringLoadingContinuousActivation]), then the `` method is called twice.
// First, it’s called with an `activated` parameter value of [true] when a
// hover timeout occurs or the user begins force clicking on a destination
// object to initiate spring-loading. Then, it’s called again with an
// `activated` parameter value of [false] when the hover exits the
// destination’s bounds or the user finishes force clicking on the
// destination object.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/springLoadingActivated(_:draggingInfo:)

func (o NSSpringLoadingDestinationObject) SpringLoadingActivatedDraggingInfo(activated bool, draggingInfo NSDraggingInfo) {
	
	objc.Send[struct{}](o.ID, objc.Sel("springLoadingActivated:draggingInfo:"), activated, draggingInfo)
	}

// Updates the destination’s user interface to display a new highlighting
// style during a spring-loading operation.
//
// draggingInfo: An object of type [NSDraggingInfo], which provides information about the
// drag event, including a highlighting style to apply.
//
// # Discussion
// 
// During a spring-loaded operation, a destination may initiate animated
// highlighting to visually cue the user that spring-loading has been engaged
// or disengaged. This method is called as different stages of this animation
// are reached, providing an opportunity to change the highlighting style.
// Check the `springLoadingHighlight` property of the `draggingInfo` object to
// determine the style of highlighting to apply. Then, update the
// destination’s user interface accordingly.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/springLoadingHighlightChanged(_:)

func (o NSSpringLoadingDestinationObject) SpringLoadingHighlightChanged(draggingInfo NSDraggingInfo) {
	
	objc.Send[struct{}](o.ID, objc.Sel("springLoadingHighlightChanged:"), draggingInfo)
	}

// Returns whether to enable or disable spring-loading when a drag enters the
// bounds of the spring-loading destination.
//
// draggingInfo: An object of type [NSDraggingInfo], which provides information about the
// drag event, including the dragged data.
//
// # Return Value
// 
// A value of type [NSSpringLoadingOptions] to enable or disable
// spring-loading. Returning a value of [NSSpringLoadingEnabled] enables
// typical spring-loading behavior that is appropriate in most cases.
//
// # Discussion
// 
// This method is called when a drag enters the bounds of the spring-loading
// destination. It returns a value of type [NSSpringLoadingOptions] to enable
// or disable spring-loading for the destination.
// 
// This method provides an opportunity to perform work in preparation for
// spring-loading becoming engaged.
// 
// Note that you implement either this method or [SpringLoadingUpdated] to
// enable spring-loading.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/springLoadingEntered(_:)

func (o NSSpringLoadingDestinationObject) SpringLoadingEntered(draggingInfo NSDraggingInfo) NSSpringLoadingOptions {
	
	rv := objc.Send[NSSpringLoadingOptions](o.ID, objc.Sel("springLoadingEntered:"), draggingInfo)
	return rv
	}

// Returns whether to enable or disable spring-loading as a drag moves within
// the bounds of the spring-loading destination or `draggingInfo` changes
// during the drag.
//
// draggingInfo: An object of type [NSDraggingInfo], which provides information about the
// drag event, including the dragged data.
//
// # Return Value
// 
// A value of type [NSSpringLoadingOptions] to enable or disable
// spring-loading. A value of [NSSpringLoadingEnabled] enables typical
// spring-loading behavior.
//
// # Discussion
// 
// This method is called periodically as a drag changes position within the
// bounds of a spring-loaded destination or the `draggingInfo` changes during
// the drag. It returns a value of type [NSSpringLoadingOptions] to enable or
// disable spring-loading for the destination. If this method is not
// implemented, then spring-loading is enabled or disabled for the destination
// based on the return value of the `` method.
// 
// Note that you implement either this method or [SpringLoadingEntered] to
// enable spring-loading.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/springLoadingUpdated(_:)

func (o NSSpringLoadingDestinationObject) SpringLoadingUpdated(draggingInfo NSDraggingInfo) NSSpringLoadingOptions {
	
	rv := objc.Send[NSSpringLoadingOptions](o.ID, objc.Sel("springLoadingUpdated:"), draggingInfo)
	return rv
	}

// Responds when a drag exits the bounds of the spring-loading destination.
//
// draggingInfo: An object of type [NSDraggingInfo], which provides information about the
// drag event, including the dragged data.
//
// # Discussion
// 
// This method is called when a drag exits the bounds of a spring-loaded
// destination.
// 
// This is a good place to clean up any initialization work that may have been
// performed during [SpringLoadingEntered].
//
// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/springLoadingExited(_:)

func (o NSSpringLoadingDestinationObject) SpringLoadingExited(draggingInfo NSDraggingInfo) {
	
	objc.Send[struct{}](o.ID, objc.Sel("springLoadingExited:"), draggingInfo)
	}

// Responds to the end of a drag operation.
//
// draggingInfo: An object of type [NSDraggingInfo], which provides information about the
// drag event, including the dragged data.
//
// # Discussion
// 
// This method is called when a drag operation has ended. If the destination
// object is both a dragging destination (class [NSDraggingDestination]) and a
// spring-loading destination (class [NSSpringLoadingDestination]), note that
// this method is only called once.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpringLoadingDestination/draggingEnded(_:)

func (o NSSpringLoadingDestinationObject) DraggingEnded(draggingInfo NSDraggingInfo) {
	
	objc.Send[struct{}](o.ID, objc.Sel("draggingEnded:"), draggingInfo)
	}

