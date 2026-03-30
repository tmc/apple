// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CALayerDelegate protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/CALayerDelegate
type CALayerDelegate interface {
	objectivec.IObject
}

// CALayerDelegateObject wraps an existing Objective-C object that conforms to the CALayerDelegate protocol.
type CALayerDelegateObject struct {
	objectivec.Object
}

func (o CALayerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CALayerDelegateObjectFromID constructs a [CALayerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CALayerDelegateObjectFromID(id objc.ID) CALayerDelegateObject {
	return CALayerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Virtualization/CALayerDelegate/actionForLayer:forKey:
func (o CALayerDelegateObject) ActionForLayerForKey(layer objectivec.IObject, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("actionForLayer:forKey:"), layer, key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/CALayerDelegate/displayLayer:
func (o CALayerDelegateObject) DisplayLayer(layer objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("displayLayer:"), layer)
}

// See: https://developer.apple.com/documentation/Virtualization/CALayerDelegate/drawLayer:inContext:
func (o CALayerDelegateObject) DrawLayerInContext(layer objectivec.IObject, context *coregraphics.CGContextRef) {
	objc.Send[struct{}](o.ID, objc.Sel("drawLayer:inContext:"), layer, context)
}

// See: https://developer.apple.com/documentation/Virtualization/CALayerDelegate/layerWillDraw:
func (o CALayerDelegateObject) LayerWillDraw(draw objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("layerWillDraw:"), draw)
}

// See: https://developer.apple.com/documentation/Virtualization/CALayerDelegate/layoutSublayersOfLayer:
func (o CALayerDelegateObject) LayoutSublayersOfLayer(layer objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("layoutSublayersOfLayer:"), layer)
}
