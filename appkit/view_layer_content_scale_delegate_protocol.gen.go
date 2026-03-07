// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An optional layer delegate method for handling resolution changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewLayerContentScaleDelegate
type NSViewLayerContentScaleDelegate interface {
	objectivec.IObject
}



// NSViewLayerContentScaleDelegateObject wraps an existing Objective-C object that conforms to the NSViewLayerContentScaleDelegate protocol.
type NSViewLayerContentScaleDelegateObject struct {
	objectivec.Object
}
func (o NSViewLayerContentScaleDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSViewLayerContentScaleDelegateObjectFromID constructs a [NSViewLayerContentScaleDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSViewLayerContentScaleDelegateObjectFromID(id objc.ID) NSViewLayerContentScaleDelegateObject {
	return NSViewLayerContentScaleDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Notifies you when a resolution changes occurs for the window that hosts the
// layer.
//
// layer: The layer whose scale and content might need updating.
//
// newScale: The new scale of the window.
//
// window: The window that hosts the layer.
//
// # Return Value
// 
// A Boolean value that specifies whether to change the layer’s
// `contentsScale` property.
//
// # Discussion
// 
// When a resolution change occurs for a given window, the system traverses
// the layer trees in that window to decide what action, if any, to take for
// each layer. The system queries the layer’s delegate to determine whether
// to change the layer’s `contentsScale` property to the new scale (either
// `2.0` or `1.0`).
// 
// Note that you don’t need to manage [NSImage] contents and that this
// method is not called on the delegate of a layer whose content is an
// [NSImage] object.
// 
// If the delegate returns [true], it should make any corresponding changes to
// the layer’s properties, as required by the resolution change. For
// example, a layer whose contents contain a CGImage object needs to determine
// whether an alternate CGImage object is available for the new scale factor.
// If the delegate finds a suitable CGImage object, then in addition to
// returning [true], it should set the appropriate CGImage object as the
// layer’s new contents.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSViewLayerContentScaleDelegate/layer(_:shouldInheritContentsScale:from:)

func (o NSViewLayerContentScaleDelegateObject) LayerShouldInheritContentsScaleFromWindow(layer objectivec.IObject, newScale float64, window INSWindow) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("layer:shouldInheritContentsScale:fromWindow:"), layer, newScale, window)
	return rv
	}







