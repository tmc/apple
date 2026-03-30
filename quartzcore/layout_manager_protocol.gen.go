// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods that allow an object to manage the layout of a layer and its sublayers.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayoutManager
type CALayoutManager interface {
	objectivec.IObject
}

// CALayoutManagerObject wraps an existing Objective-C object that conforms to the CALayoutManager protocol.
type CALayoutManagerObject struct {
	objectivec.Object
}

func (o CALayoutManagerObject) BaseObject() objectivec.Object {
	return o.Object
}

// CALayoutManagerObjectFromID constructs a [CALayoutManagerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CALayoutManagerObjectFromID(id objc.ID) CALayoutManagerObject {
	return CALayoutManagerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Invalidates the layout of a layer so it knows to refresh its content on the
// next frame.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayoutManager/invalidateLayout(of:)
func (o CALayoutManagerObject) InvalidateLayoutOfLayer(layer ICALayer) {
	objc.Send[struct{}](o.ID, objc.Sel("invalidateLayoutOfLayer:"), layer)
}

// Override to customize layout of sublayers whenever the layer needs
// redrawing.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayoutManager/layoutSublayers(of:)
func (o CALayoutManagerObject) LayoutSublayersOfLayer(layer ICALayer) {
	objc.Send[struct{}](o.ID, objc.Sel("layoutSublayersOfLayer:"), layer)
}

// Override to customize layer size.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayoutManager/preferredSize(of:)
func (o CALayoutManagerObject) PreferredSizeOfLayer(layer ICALayer) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("preferredSizeOfLayer:"), layer)
	return rv
}
