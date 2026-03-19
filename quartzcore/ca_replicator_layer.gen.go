// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAReplicatorLayer] class.
var (
	_CAReplicatorLayerClass     CAReplicatorLayerClass
	_CAReplicatorLayerClassOnce sync.Once
)

func getCAReplicatorLayerClass() CAReplicatorLayerClass {
	_CAReplicatorLayerClassOnce.Do(func() {
		_CAReplicatorLayerClass = CAReplicatorLayerClass{class: objc.GetClass("CAReplicatorLayer")}
	})
	return _CAReplicatorLayerClass
}

// GetCAReplicatorLayerClass returns the class object for CAReplicatorLayer.
func GetCAReplicatorLayerClass() CAReplicatorLayerClass {
	return getCAReplicatorLayerClass()
}

type CAReplicatorLayerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAReplicatorLayerClass) Alloc() CAReplicatorLayer {
	rv := objc.Send[CAReplicatorLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that creates a specified number of sublayer copies with varying
// geometric, temporal, and color transformations.
//
// # Overview
// 
// You can use a [CAReplicatorLayer] object to build complex layouts based on
// a single source layer that is replicated with transformation rules that can
// affect the position, rotation color, and time.
// 
// The following shows a simple example: a red square is added to a replicator
// layer with an instance count of `5`. The position of each replicated
// instance is offset along the `x` axis so that it appears to the right of
// the previous instance. The blue and green color channels are offset so that
// their values reach `0` at the final instance.
// 
// The result of the code above is a row of five squares, with colors
// graduating from white to red.
// 
// [media-2776906]
// 
// Replicator layers can be nested. The following code adds `replicatorLayer`
// to a second replicator layer that offsets the position of each instance
// vertically and subtracts from the red channel.
// 
// The result of adding this code is to create a grid with the value of the
// red channel being reduced in the vertical direction.
// 
// [media-2776908]
//
// # Setting Instance Display Properties
//
//   - [CAReplicatorLayer.InstanceCount]: The number of copies to create, including the source layers.
//   - [CAReplicatorLayer.SetInstanceCount]
//   - [CAReplicatorLayer.InstanceDelay]: Specifies the delay, in seconds, between replicated copies. Animatable.
//   - [CAReplicatorLayer.SetInstanceDelay]
//   - [CAReplicatorLayer.InstanceTransform]: The transform matrix applied to the previous instance to produce the current instance. Animatable.
//   - [CAReplicatorLayer.SetInstanceTransform]
//
// # Modifying Instance Layer Geometry
//
//   - [CAReplicatorLayer.PreservesDepth]: Defines whether this layer flattens its sublayers into its plane.
//   - [CAReplicatorLayer.SetPreservesDepth]
//
// # Accessing Instance Color Values
//
//   - [CAReplicatorLayer.InstanceColor]: Defines the color used to multiply the source object. Animatable.
//   - [CAReplicatorLayer.SetInstanceColor]
//   - [CAReplicatorLayer.InstanceRedOffset]: Defines the offset added to the red component of the color for each replicated instance. Animatable.
//   - [CAReplicatorLayer.SetInstanceRedOffset]
//   - [CAReplicatorLayer.InstanceGreenOffset]: Defines the offset added to the green component of the color for each replicated instance. Animatable.
//   - [CAReplicatorLayer.SetInstanceGreenOffset]
//   - [CAReplicatorLayer.InstanceBlueOffset]: Defines the offset added to the blue component of the color for each replicated instance. Animatable.
//   - [CAReplicatorLayer.SetInstanceBlueOffset]
//   - [CAReplicatorLayer.InstanceAlphaOffset]: Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
//   - [CAReplicatorLayer.SetInstanceAlphaOffset]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer
type CAReplicatorLayer struct {
	CALayer
}

// CAReplicatorLayerFromID constructs a [CAReplicatorLayer] from an objc.ID.
//
// A layer that creates a specified number of sublayer copies with varying
// geometric, temporal, and color transformations.
func CAReplicatorLayerFromID(id objc.ID) CAReplicatorLayer {
	return CAReplicatorLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CAReplicatorLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAReplicatorLayer] class.
//
// # Setting Instance Display Properties
//
//   - [ICAReplicatorLayer.InstanceCount]: The number of copies to create, including the source layers.
//   - [ICAReplicatorLayer.SetInstanceCount]
//   - [ICAReplicatorLayer.InstanceDelay]: Specifies the delay, in seconds, between replicated copies. Animatable.
//   - [ICAReplicatorLayer.SetInstanceDelay]
//   - [ICAReplicatorLayer.InstanceTransform]: The transform matrix applied to the previous instance to produce the current instance. Animatable.
//   - [ICAReplicatorLayer.SetInstanceTransform]
//
// # Modifying Instance Layer Geometry
//
//   - [ICAReplicatorLayer.PreservesDepth]: Defines whether this layer flattens its sublayers into its plane.
//   - [ICAReplicatorLayer.SetPreservesDepth]
//
// # Accessing Instance Color Values
//
//   - [ICAReplicatorLayer.InstanceColor]: Defines the color used to multiply the source object. Animatable.
//   - [ICAReplicatorLayer.SetInstanceColor]
//   - [ICAReplicatorLayer.InstanceRedOffset]: Defines the offset added to the red component of the color for each replicated instance. Animatable.
//   - [ICAReplicatorLayer.SetInstanceRedOffset]
//   - [ICAReplicatorLayer.InstanceGreenOffset]: Defines the offset added to the green component of the color for each replicated instance. Animatable.
//   - [ICAReplicatorLayer.SetInstanceGreenOffset]
//   - [ICAReplicatorLayer.InstanceBlueOffset]: Defines the offset added to the blue component of the color for each replicated instance. Animatable.
//   - [ICAReplicatorLayer.SetInstanceBlueOffset]
//   - [ICAReplicatorLayer.InstanceAlphaOffset]: Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
//   - [ICAReplicatorLayer.SetInstanceAlphaOffset]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer
type ICAReplicatorLayer interface {
	ICALayer

	// Topic: Setting Instance Display Properties

	// The number of copies to create, including the source layers.
	InstanceCount() int
	SetInstanceCount(value int)
	// Specifies the delay, in seconds, between replicated copies. Animatable.
	InstanceDelay() float64
	SetInstanceDelay(value float64)
	// The transform matrix applied to the previous instance to produce the current instance. Animatable.
	InstanceTransform() CATransform3D
	SetInstanceTransform(value CATransform3D)

	// Topic: Modifying Instance Layer Geometry

	// Defines whether this layer flattens its sublayers into its plane.
	PreservesDepth() bool
	SetPreservesDepth(value bool)

	// Topic: Accessing Instance Color Values

	// Defines the color used to multiply the source object. Animatable.
	InstanceColor() coregraphics.CGColorRef
	SetInstanceColor(value coregraphics.CGColorRef)
	// Defines the offset added to the red component of the color for each replicated instance. Animatable.
	InstanceRedOffset() float32
	SetInstanceRedOffset(value float32)
	// Defines the offset added to the green component of the color for each replicated instance. Animatable.
	InstanceGreenOffset() float32
	SetInstanceGreenOffset(value float32)
	// Defines the offset added to the blue component of the color for each replicated instance. Animatable.
	InstanceBlueOffset() float32
	SetInstanceBlueOffset(value float32)
	// Defines the offset added to the alpha component of the color for each replicated instance. Animatable.
	InstanceAlphaOffset() float32
	SetInstanceAlphaOffset(value float32)
}

// Init initializes the instance.
func (r CAReplicatorLayer) Init() CAReplicatorLayer {
	rv := objc.Send[CAReplicatorLayer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CAReplicatorLayer) Autorelease() CAReplicatorLayer {
	rv := objc.Send[CAReplicatorLayer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAReplicatorLayer creates a new CAReplicatorLayer instance.
func NewCAReplicatorLayer() CAReplicatorLayer {
	class := getCAReplicatorLayerClass()
	rv := objc.Send[CAReplicatorLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Override to copy or initialize custom fields of the specified layer.
//
// layer: The layer from which custom fields should be copied.
//
// # Return Value
// 
// A layer instance with any custom instance variables copied from `layer`.
//
// # Discussion
// 
// This initializer is used to create shadow copies of layers, for example,
// for the [PresentationLayer] method. Using this method in any other
// situation will produce undefined behavior. For example, do not use this
// method to initialize a new layer with an existing layer’s content.
// 
// If you are implementing a custom layer subclass, you can override this
// method and use it to copy the values of instance variables into the new
// object. Subclasses should always invoke the superclass implementation.
// 
// This method is the designated initializer for layer objects in the
// presentation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewReplicatorLayerWithLayer(layer objectivec.IObject) CAReplicatorLayer {
	instance := getCAReplicatorLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CAReplicatorLayerFromID(rv)
}

// The number of copies to create, including the source layers.
//
// # Discussion
// 
// Default value is `1`, no extra copies are created.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceCount
func (r CAReplicatorLayer) InstanceCount() int {
	rv := objc.Send[int](r.ID, objc.Sel("instanceCount"))
	return rv
}
func (r CAReplicatorLayer) SetInstanceCount(value int) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceCount:"), value)
}
// Specifies the delay, in seconds, between replicated copies. Animatable.
//
// # Discussion
// 
// The default value is `0.0`, meaning that any animations added to replicated
// copies will be synchronized.
// 
// The following code shows a replicator layer being used to create an
// animated activity monitor. The replicator layer creates 30 small circles
// forming a larger circle. The source layer, `circle`, has a 1 second
// animated fade out and each of the copies offsets the time of the animation
// by 1 / 30 seconds.
// 
// The following illustration shows the result of the above code:
// 
// [media-2776911]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceDelay
func (r CAReplicatorLayer) InstanceDelay() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("instanceDelay"))
	return rv
}
func (r CAReplicatorLayer) SetInstanceDelay(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceDelay:"), value)
}
// The transform matrix applied to the previous instance to produce the
// current instance. Animatable.
//
// # Discussion
// 
// This transform matrix is applied to instance `k-1` to produce instance `k`.
// The matrix is applied relative to the center of this layer.
// 
// Defaults to the identity matrix.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceTransform
func (r CAReplicatorLayer) InstanceTransform() CATransform3D {
	rv := objc.Send[CATransform3D](r.ID, objc.Sel("instanceTransform"))
	return CATransform3D(rv)
}
func (r CAReplicatorLayer) SetInstanceTransform(value CATransform3D) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceTransform:"), value)
}
// Defines whether this layer flattens its sublayers into its plane.
//
// # Discussion
// 
// If [true], the layer acts similarly to the [CATransformLayer] and has the
// same restrictions.
// 
// Default is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/preservesDepth
func (r CAReplicatorLayer) PreservesDepth() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("preservesDepth"))
	return rv
}
func (r CAReplicatorLayer) SetPreservesDepth(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setPreservesDepth:"), value)
}
// Defines the color used to multiply the source object. Animatable.
//
// # Discussion
// 
// Defaults to opaque white.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceColor
func (r CAReplicatorLayer) InstanceColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](r.ID, objc.Sel("instanceColor"))
	return coregraphics.CGColorRef(rv)
}
func (r CAReplicatorLayer) SetInstanceColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceColor:"), value)
}
// Defines the offset added to the red component of the color for each
// replicated instance. Animatable.
//
// # Discussion
// 
// The `instanceRedOffset` is added to the red color component of instance
// `k-1` to produce the modulation color of instance k.
// 
// Default is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceRedOffset
func (r CAReplicatorLayer) InstanceRedOffset() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("instanceRedOffset"))
	return rv
}
func (r CAReplicatorLayer) SetInstanceRedOffset(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceRedOffset:"), value)
}
// Defines the offset added to the green component of the color for each
// replicated instance. Animatable.
//
// # Discussion
// 
// The `instanceGreenOffset` is added to the green color component of instance
// `k-1` to produce the modulation color of instance k.
// 
// Default is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceGreenOffset
func (r CAReplicatorLayer) InstanceGreenOffset() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("instanceGreenOffset"))
	return rv
}
func (r CAReplicatorLayer) SetInstanceGreenOffset(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceGreenOffset:"), value)
}
// Defines the offset added to the blue component of the color for each
// replicated instance. Animatable.
//
// # Discussion
// 
// The `instanceBlueOffset` is added to the blue color component of instance
// `k-1` to produce the modulation color of instance k.
// 
// Default is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceBlueOffset
func (r CAReplicatorLayer) InstanceBlueOffset() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("instanceBlueOffset"))
	return rv
}
func (r CAReplicatorLayer) SetInstanceBlueOffset(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceBlueOffset:"), value)
}
// Defines the offset added to the alpha component of the color for each
// replicated instance. Animatable.
//
// # Discussion
// 
// The `instanceAlphaOffset` is added to the alpha color component of instance
// `k-1` to produce the modulation color of instance k.
// 
// Default is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAReplicatorLayer/instanceAlphaOffset
func (r CAReplicatorLayer) InstanceAlphaOffset() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("instanceAlphaOffset"))
	return rv
}
func (r CAReplicatorLayer) SetInstanceAlphaOffset(value float32) {
	objc.Send[struct{}](r.ID, objc.Sel("setInstanceAlphaOffset:"), value)
}

