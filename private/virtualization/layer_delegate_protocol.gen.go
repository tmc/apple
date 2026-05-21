// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CALayerDelegate protocol.
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

func (o CALayerDelegateObject) ActionForLayerForKey(layer objectivec.IObject, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("actionForLayer:forKey:"), layer, key)
	return objectivec.Object{ID: rv}
}
func (o CALayerDelegateObject) DisplayLayer(layer objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("displayLayer:"), layer)
}
func (o CALayerDelegateObject) DrawLayerInContext(layer objectivec.IObject, context coregraphics.CGContextRef) {
	objc.Send[struct{}](o.ID, objc.Sel("drawLayer:inContext:"), layer, context)
}
func (o CALayerDelegateObject) LayerWillDraw(draw objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("layerWillDraw:"), draw)
}
func (o CALayerDelegateObject) LayoutSublayersOfLayer(layer objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("layoutSublayersOfLayer:"), layer)
}
