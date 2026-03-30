// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"fmt"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods your app can implement to respond to layer-related events.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerDelegate
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

// Tells the delegate to implement the display process.
//
// layer: The layer whose contents need updating.
//
// # Discussion
//
// The [DisplayLayer] delegate method is called when the layer is marked for
// its content to be reloaded, typically initiated by the [SetNeedsDisplay]
// method. The typical technique for updating is to set the layer’s
// `contents` property.
//
// The following code shows how you can create a class named [LayerDelegate]
// that implements [CALayerDelegate] and sets it as a layer’s (named
// `sublayer`) delegate. When [SetNeedsDisplay] is called on `sublayer`, the
// delegate’s [DisplayLayer] replaces its contents with a specified image.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerDelegate/display(_:)
func (o CALayerDelegateObject) DisplayLayer(layer ICALayer) {
	objc.Send[struct{}](o.ID, objc.Sel("displayLayer:"), layer)
}

// Tells the delegate to implement the display process using the layer’s
// context.
//
// layer: The layer whose contents need to be drawn.
//
// ctx: The graphics context to use for drawing. The graphics context incorporates
// the appropriate scale factor for drawing to the target screen.
//
// # Discussion
//
// The [DrawLayerInContext] method is called when the layer is marked for its
// content to be reloaded, typically with the [SetNeedsDisplay] method. It is
// not called if the delegate implements the [DisplayLayer] method. You can
// use the context to draw vectors, such as curves and lines, or images with
// the [draw(_:in:byTiling:)] method.
//
// The following code shows how you can create a class named [LayerDelegate]
// that implements [CALayerDelegate] and sets it as a layer’s (named
// `sublayer`) delegate. When [SetNeedsDisplay] is called on `sublayer`, the
// delegate’s [DrawLayerInContext] method draws an ellipse fitting the
// bounding box of the layer using the [boundingBoxOfClipPath] function.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerDelegate/draw(_:in:)
//
// [boundingBoxOfClipPath]: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfClipPath
// [draw(_:in:byTiling:)]: https://developer.apple.com/documentation/CoreGraphics/CGContext/draw(_:in:byTiling:)
func (o CALayerDelegateObject) DrawLayerInContext(layer ICALayer, ctx coregraphics.CGContextRef) {
	objc.Send[struct{}](o.ID, objc.Sel("drawLayer:inContext:"), layer, ctx)
}

// Notifies the delegate of an imminent draw.
//
// layer: The layer whose contents will be drawn.
//
// # Discussion
//
// The [LayerWillDraw] method is called before [DrawLayerInContext]. You can
// use this method to configure any layer state affecting contents prior to
// [DrawLayerInContext] such as [ContentsFormat] and [Opaque].
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerDelegate/layerWillDraw(_:)
func (o CALayerDelegateObject) LayerWillDraw(layer ICALayer) {
	objc.Send[struct{}](o.ID, objc.Sel("layerWillDraw:"), layer)
}

// Tells the delegate a layer’s bounds have changed.
//
// layer: The layer that requires layout of its sublayers.
//
// # Discussion
//
// The [LayoutSublayersOfLayer] method is called when a layer’s bounds have
// changed and its sublayers may need rearranging, for example by changing its
// frame’s size. You can implement this method if you need precise control
// over the layout of your layer’s sublayers.
//
// The following code shows how you can create a class named [LayerDelegate]
// that implements [CALayerDelegate] and sets it as a layer’s (named
// `sublayer`) delegate. When the layer’s size changes, the delegate’s
// [LayoutSublayersOfLayer] iterates over all of the sublayers of `sublayer`
// and resizes them to fit within it.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerDelegate/layoutSublayers(of:)
func (o CALayerDelegateObject) LayoutSublayersOfLayer(layer ICALayer) {
	objc.Send[struct{}](o.ID, objc.Sel("layoutSublayersOfLayer:"), layer)
}

// Returns the default action of the [ActionForKey] method.
//
// layer: The layer that is the target of the action.
//
// event: The identifier of the action.
//
// # Return Value
//
// An object implementing the [CAAction] protocol or `nil` if the delegate
// does not specify a behavior for the specified `key`.
//
// # Discussion
//
// A layer’s delegate that implements this method returns an action for a
// specified key and stops any further searches (i.e. actions for the same key
// in the layer’s [Actions] dictionary or specified by [DefaultActionForKey]
// are not returned).
//
// The following code shows how you can create a class named [LayerDelegate]
// that implements [CALayerDelegate] and sets it as a layer’s (named
// `sublayer`) delegate. [LayerDelegate] returns a basic animation that moves
// an object from left to right. The `moveSublayer()` method searches
// `sublayer` for an action with the key `moveRight` and, if the action is
// found, runs it.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerDelegate/action(for:forKey:)
func (o CALayerDelegateObject) ActionForLayerForKey(layer ICALayer, event string) CAAction {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("actionForLayer:forKey:"), layer, objc.String(event))
	return CAActionObjectFromID(rv)
}

// CALayerDelegateConfig holds optional typed callbacks for [CALayerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/quartzcore/calayerdelegate
type CALayerDelegateConfig struct {

	// Providing the Layer’s Content
	// LayerWillDraw — Notifies the delegate of an imminent draw.
	LayerWillDraw func(layer CALayer)

	// Other Methods
	// DisplayLayer — Tells the delegate to implement the display process.
	DisplayLayer func(layer CALayer)
	// LayoutSublayersOfLayer — Tells the delegate a layer’s bounds have changed.
	LayoutSublayersOfLayer func(layer CALayer)
}

// NewCALayerDelegate creates an Objective-C object implementing the [CALayerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CALayerDelegateObject] satisfies the [CALayerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/quartzcore/calayerdelegate
func NewCALayerDelegate(config CALayerDelegateConfig) CALayerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCALayerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DisplayLayer != nil {
		fn := config.DisplayLayer
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("displayLayer:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layerID objc.ID) {
				layer := CALayerFromID(layerID)
				fn(layer)
			},
		})
	}

	if config.LayerWillDraw != nil {
		fn := config.LayerWillDraw
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layerWillDraw:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layerID objc.ID) {
				layer := CALayerFromID(layerID)
				fn(layer)
			},
		})
	}

	if config.LayoutSublayersOfLayer != nil {
		fn := config.LayoutSublayersOfLayer
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("layoutSublayersOfLayer:"),
			Fn: func(self objc.ID, _cmd objc.SEL, layerID objc.ID) {
				layer := CALayerFromID(layerID)
				fn(layer)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CALayerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCALayerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CALayerDelegateObjectFromID(instance)
}
