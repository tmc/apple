// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAEmitterLayer] class.
var (
	_CAEmitterLayerClass     CAEmitterLayerClass
	_CAEmitterLayerClassOnce sync.Once
)

func getCAEmitterLayerClass() CAEmitterLayerClass {
	_CAEmitterLayerClassOnce.Do(func() {
		_CAEmitterLayerClass = CAEmitterLayerClass{class: objc.GetClass("CAEmitterLayer")}
	})
	return _CAEmitterLayerClass
}

// GetCAEmitterLayerClass returns the class object for CAEmitterLayer.
func GetCAEmitterLayerClass() CAEmitterLayerClass {
	return getCAEmitterLayerClass()
}

type CAEmitterLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAEmitterLayerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAEmitterLayerClass) Alloc() CAEmitterLayer {
	rv := objc.Send[CAEmitterLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that emits, animates, and renders a particle system.
//
// # Overview
// 
// The particles, defined by instances of [CAEmitterCell], are drawn above the
// layer’s background color and border.
// 
// The following code shows how to set up a simple point (the default
// [CAEmitterLayer.EmitterShape] is [point]) particle emitter. It uses an image named
// `RadialGradient.Png()` as the cell contents and, by setting the emitter
// cell’s [CAEmitterLayer.EmissionRange] to `2` × [CAEmitterLayer.Pi], the particles are emitted in all
// directions.
//
// [point]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/point
//
// # Specifying Particle Emitter Cells
//
//   - [CAEmitterLayer.EmitterCells]: The array emitter cells attached to the layer.
//   - [CAEmitterLayer.SetEmitterCells]
//
// # Emitter Geometry
//
//   - [CAEmitterLayer.RenderMode]: Defines how particle cells are rendered into the layer.
//   - [CAEmitterLayer.SetRenderMode]
//   - [CAEmitterLayer.EmitterPosition]: The position of the center of the particle emitter. Animatable.
//   - [CAEmitterLayer.SetEmitterPosition]
//   - [CAEmitterLayer.EmitterShape]: Specifies the emitter shape.
//   - [CAEmitterLayer.SetEmitterShape]
//   - [CAEmitterLayer.EmitterZPosition]: Specifies the center of the particle emitter shape along the z-axis. Animatable.
//   - [CAEmitterLayer.SetEmitterZPosition]
//   - [CAEmitterLayer.EmitterDepth]: Determines the depth of the emitter shape.
//   - [CAEmitterLayer.SetEmitterDepth]
//   - [CAEmitterLayer.EmitterSize]: Determines the size of the particle emitter shape. Animatable.
//   - [CAEmitterLayer.SetEmitterSize]
//
// # Emitter Cell Attribute Multipliers
//
//   - [CAEmitterLayer.Scale]: Defines a multiplier applied to the cell-defined particle scale.
//   - [CAEmitterLayer.SetScale]
//   - [CAEmitterLayer.Seed]: Specifies the seed used to initialize the random number generator.
//   - [CAEmitterLayer.SetSeed]
//   - [CAEmitterLayer.Spin]: Defines a multiplier applied to the cell-defined particle spin. Animatable.
//   - [CAEmitterLayer.SetSpin]
//   - [CAEmitterLayer.Velocity]: Defines a multiplier applied to the cell-defined particle velocity. Animatable.
//   - [CAEmitterLayer.SetVelocity]
//   - [CAEmitterLayer.BirthRate]: Defines a multiplier that is applied to the cell-defined birth rate. Animatable
//   - [CAEmitterLayer.SetBirthRate]
//   - [CAEmitterLayer.EmitterMode]: Specifies the emitter mode.
//   - [CAEmitterLayer.SetEmitterMode]
//   - [CAEmitterLayer.Lifetime]: Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable.
//   - [CAEmitterLayer.SetLifetime]
//   - [CAEmitterLayer.PreservesDepth]: Defines whether the layer flattens the particles into its plane.
//   - [CAEmitterLayer.SetPreservesDepth]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer
type CAEmitterLayer struct {
	CALayer
}

// CAEmitterLayerFromID constructs a [CAEmitterLayer] from an objc.ID.
//
// A layer that emits, animates, and renders a particle system.
func CAEmitterLayerFromID(id objc.ID) CAEmitterLayer {
	return CAEmitterLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CAEmitterLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAEmitterLayer] class.
//
// # Specifying Particle Emitter Cells
//
//   - [ICAEmitterLayer.EmitterCells]: The array emitter cells attached to the layer.
//   - [ICAEmitterLayer.SetEmitterCells]
//
// # Emitter Geometry
//
//   - [ICAEmitterLayer.RenderMode]: Defines how particle cells are rendered into the layer.
//   - [ICAEmitterLayer.SetRenderMode]
//   - [ICAEmitterLayer.EmitterPosition]: The position of the center of the particle emitter. Animatable.
//   - [ICAEmitterLayer.SetEmitterPosition]
//   - [ICAEmitterLayer.EmitterShape]: Specifies the emitter shape.
//   - [ICAEmitterLayer.SetEmitterShape]
//   - [ICAEmitterLayer.EmitterZPosition]: Specifies the center of the particle emitter shape along the z-axis. Animatable.
//   - [ICAEmitterLayer.SetEmitterZPosition]
//   - [ICAEmitterLayer.EmitterDepth]: Determines the depth of the emitter shape.
//   - [ICAEmitterLayer.SetEmitterDepth]
//   - [ICAEmitterLayer.EmitterSize]: Determines the size of the particle emitter shape. Animatable.
//   - [ICAEmitterLayer.SetEmitterSize]
//
// # Emitter Cell Attribute Multipliers
//
//   - [ICAEmitterLayer.Scale]: Defines a multiplier applied to the cell-defined particle scale.
//   - [ICAEmitterLayer.SetScale]
//   - [ICAEmitterLayer.Seed]: Specifies the seed used to initialize the random number generator.
//   - [ICAEmitterLayer.SetSeed]
//   - [ICAEmitterLayer.Spin]: Defines a multiplier applied to the cell-defined particle spin. Animatable.
//   - [ICAEmitterLayer.SetSpin]
//   - [ICAEmitterLayer.Velocity]: Defines a multiplier applied to the cell-defined particle velocity. Animatable.
//   - [ICAEmitterLayer.SetVelocity]
//   - [ICAEmitterLayer.BirthRate]: Defines a multiplier that is applied to the cell-defined birth rate. Animatable
//   - [ICAEmitterLayer.SetBirthRate]
//   - [ICAEmitterLayer.EmitterMode]: Specifies the emitter mode.
//   - [ICAEmitterLayer.SetEmitterMode]
//   - [ICAEmitterLayer.Lifetime]: Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable.
//   - [ICAEmitterLayer.SetLifetime]
//   - [ICAEmitterLayer.PreservesDepth]: Defines whether the layer flattens the particles into its plane.
//   - [ICAEmitterLayer.SetPreservesDepth]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer
type ICAEmitterLayer interface {
	ICALayer

	// Topic: Specifying Particle Emitter Cells

	// The array emitter cells attached to the layer.
	EmitterCells() []CAEmitterCell
	SetEmitterCells(value []CAEmitterCell)

	// Topic: Emitter Geometry

	// Defines how particle cells are rendered into the layer.
	RenderMode() CAEmitterLayerRenderMode
	SetRenderMode(value CAEmitterLayerRenderMode)
	// The position of the center of the particle emitter. Animatable.
	EmitterPosition() corefoundation.CGPoint
	SetEmitterPosition(value corefoundation.CGPoint)
	// Specifies the emitter shape.
	EmitterShape() CAEmitterLayerEmitterShape
	SetEmitterShape(value CAEmitterLayerEmitterShape)
	// Specifies the center of the particle emitter shape along the z-axis. Animatable.
	EmitterZPosition() float64
	SetEmitterZPosition(value float64)
	// Determines the depth of the emitter shape.
	EmitterDepth() float64
	SetEmitterDepth(value float64)
	// Determines the size of the particle emitter shape. Animatable.
	EmitterSize() corefoundation.CGSize
	SetEmitterSize(value corefoundation.CGSize)

	// Topic: Emitter Cell Attribute Multipliers

	// Defines a multiplier applied to the cell-defined particle scale.
	Scale() float32
	SetScale(value float32)
	// Specifies the seed used to initialize the random number generator.
	Seed() uint32
	SetSeed(value uint32)
	// Defines a multiplier applied to the cell-defined particle spin. Animatable.
	Spin() float32
	SetSpin(value float32)
	// Defines a multiplier applied to the cell-defined particle velocity. Animatable.
	Velocity() float32
	SetVelocity(value float32)
	// Defines a multiplier that is applied to the cell-defined birth rate. Animatable
	BirthRate() float32
	SetBirthRate(value float32)
	// Specifies the emitter mode.
	EmitterMode() CAEmitterLayerEmitterMode
	SetEmitterMode(value CAEmitterLayerEmitterMode)
	// Defines a multiplier applied to the cell-defined lifetime range when particles are created. Animatable.
	Lifetime() float32
	SetLifetime(value float32)
	// Defines whether the layer flattens the particles into its plane.
	PreservesDepth() bool
	SetPreservesDepth(value bool)

	// The angle, in radians, defining a cone around the emission angle. Animatable.
	EmissionRange() float64
	SetEmissionRange(value float64)
	// The mathematical constant pi (π), approximately equal to 3.14159.
	Pi() objectivec.IObject
	SetPi(value objectivec.IObject)
	// Particles are emitted from a single point at (
	Point() CAEmitterLayerEmitterShape
}

// Init initializes the instance.
func (e CAEmitterLayer) Init() CAEmitterLayer {
	rv := objc.Send[CAEmitterLayer](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e CAEmitterLayer) Autorelease() CAEmitterLayer {
	rv := objc.Send[CAEmitterLayer](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAEmitterLayer creates a new CAEmitterLayer instance.
func NewCAEmitterLayer() CAEmitterLayer {
	class := getCAEmitterLayerClass()
	rv := objc.Send[CAEmitterLayer](objc.ID(class.class), objc.Sel("new"))
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
func NewEmitterLayerWithLayer(layer objectivec.IObject) CAEmitterLayer {
	instance := getCAEmitterLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CAEmitterLayerFromID(rv)
}

// The array emitter cells attached to the layer.
//
// # Discussion
// 
// Each object in the array must be an instance of the [CAEmitterCell] class.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterCells
func (e CAEmitterLayer) EmitterCells() []CAEmitterCell {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("emitterCells"))
	return objc.ConvertSlice(rv, func(id objc.ID) CAEmitterCell {
		return CAEmitterCellFromID(id)
	})
}
func (e CAEmitterLayer) SetEmitterCells(value []CAEmitterCell) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterCells:"), objectivec.IObjectSliceToNSArray(value))
}
// Defines how particle cells are rendered into the layer.
//
// # Discussion
// 
// The possible values for render modes are shown in [Emitter Modes]. The
// default value is [unordered].
//
// [Emitter Modes]: https://developer.apple.com/documentation/QuartzCore/emitter-modes
// [unordered]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerRenderMode/unordered
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/renderMode
func (e CAEmitterLayer) RenderMode() CAEmitterLayerRenderMode {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("renderMode"))
	return CAEmitterLayerRenderMode(foundation.NSStringFromID(rv).String())
}
func (e CAEmitterLayer) SetRenderMode(value CAEmitterLayerRenderMode) {
	objc.Send[struct{}](e.ID, objc.Sel("setRenderMode:"), objc.String(string(value)))
}
// The position of the center of the particle emitter. Animatable.
//
// # Discussion
// 
// See [Emitter Shape] for details of how the `emitterPosition` relates to the
// possible emitter shapes.
// 
// Default is `(0.0,0.0)`.
//
// [Emitter Shape]: https://developer.apple.com/documentation/QuartzCore/emitter-shape
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterPosition
func (e CAEmitterLayer) EmitterPosition() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](e.ID, objc.Sel("emitterPosition"))
	return corefoundation.CGPoint(rv)
}
func (e CAEmitterLayer) SetEmitterPosition(value corefoundation.CGPoint) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterPosition:"), value)
}
// Specifies the emitter shape.
//
// # Discussion
// 
// The possible values for emitterMode are shown in [Emitter Shape]. The
// default value is [point].
//
// [Emitter Shape]: https://developer.apple.com/documentation/QuartzCore/emitter-shape
// [point]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/point
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterShape
func (e CAEmitterLayer) EmitterShape() CAEmitterLayerEmitterShape {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("emitterShape"))
	return CAEmitterLayerEmitterShape(foundation.NSStringFromID(rv).String())
}
func (e CAEmitterLayer) SetEmitterShape(value CAEmitterLayerEmitterShape) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterShape:"), objc.String(string(value)))
}
// Specifies the center of the particle emitter shape along the z-axis.
// Animatable.
//
// # Discussion
// 
// See [Emitter Shape] for details of how the emitterZPosition relates to the
// possible emitter shapes.
// 
// Default is `0.0`.
//
// [Emitter Shape]: https://developer.apple.com/documentation/QuartzCore/emitter-shape
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterZPosition
func (e CAEmitterLayer) EmitterZPosition() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emitterZPosition"))
	return rv
}
func (e CAEmitterLayer) SetEmitterZPosition(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterZPosition:"), value)
}
// Determines the depth of the emitter shape.
//
// # Discussion
// 
// How the emitter depth is applied depends on the emitter shape. See [Emitter
// Shape] for details. Depending on the value of [EmitterShape], this value
// may be ignored.
// 
// Default is `0.0`.
//
// [Emitter Shape]: https://developer.apple.com/documentation/QuartzCore/emitter-shape
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterDepth
func (e CAEmitterLayer) EmitterDepth() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emitterDepth"))
	return rv
}
func (e CAEmitterLayer) SetEmitterDepth(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterDepth:"), value)
}
// Determines the size of the particle emitter shape. Animatable.
//
// # Discussion
// 
// How the emitter size is applied depends on the emitter shape. See [Emitter
// Shape] for details. Depending on the value of [EmitterShape], this value
// may be ignored.
// 
// Default is `0.0`.
//
// [Emitter Shape]: https://developer.apple.com/documentation/QuartzCore/emitter-shape
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterSize
func (e CAEmitterLayer) EmitterSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](e.ID, objc.Sel("emitterSize"))
	return corefoundation.CGSize(rv)
}
func (e CAEmitterLayer) SetEmitterSize(value corefoundation.CGSize) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterSize:"), value)
}
// Defines a multiplier applied to the cell-defined particle scale.
//
// # Discussion
// 
// Default value is `1.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/scale
func (e CAEmitterLayer) Scale() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("scale"))
	return rv
}
func (e CAEmitterLayer) SetScale(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setScale:"), value)
}
// Specifies the seed used to initialize the random number generator.
//
// # Discussion
// 
// Each layer has its own random number generator state. Emitter cell
// properties that are defined as a mean and a range, such as a cell’s
// `speed`, the value of the properties are uniformly distributed in the
// interval [M - R/2, M + R/2].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/seed
func (e CAEmitterLayer) Seed() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("seed"))
	return rv
}
func (e CAEmitterLayer) SetSeed(value uint32) {
	objc.Send[struct{}](e.ID, objc.Sel("setSeed:"), value)
}
// Defines a multiplier applied to the cell-defined particle spin. Animatable.
//
// # Discussion
// 
// Default value is `1.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/spin
func (e CAEmitterLayer) Spin() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("spin"))
	return rv
}
func (e CAEmitterLayer) SetSpin(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setSpin:"), value)
}
// Defines a multiplier applied to the cell-defined particle velocity.
// Animatable.
//
// # Discussion
// 
// Default value is `1.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/velocity
func (e CAEmitterLayer) Velocity() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("velocity"))
	return rv
}
func (e CAEmitterLayer) SetVelocity(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setVelocity:"), value)
}
// Defines a multiplier that is applied to the cell-defined birth rate.
// Animatable
//
// # Discussion
// 
// The birth rate of each cell is multiplied by this number to give the actual
// number of particles created every second. Default value is `1.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/birthRate
func (e CAEmitterLayer) BirthRate() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("birthRate"))
	return rv
}
func (e CAEmitterLayer) SetBirthRate(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setBirthRate:"), value)
}
// Specifies the emitter mode.
//
// # Discussion
// 
// The possible values for emitterMode are shown in [Emitter Modes]. The
// default value is [volume].
//
// [Emitter Modes]: https://developer.apple.com/documentation/QuartzCore/emitter-modes
// [volume]: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterMode/volume
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/emitterMode
func (e CAEmitterLayer) EmitterMode() CAEmitterLayerEmitterMode {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("emitterMode"))
	return CAEmitterLayerEmitterMode(foundation.NSStringFromID(rv).String())
}
func (e CAEmitterLayer) SetEmitterMode(value CAEmitterLayerEmitterMode) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterMode:"), objc.String(string(value)))
}
// Defines a multiplier applied to the cell-defined lifetime range when
// particles are created. Animatable.
//
// # Discussion
// 
// Default value is `1.0`.
// 
// By setting an emitter’s [Lifetime] to `0`, you effectively stop particle
// emission: all new particles created have their [Lifetime] set to `0` and
// are never rendered.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/lifetime
func (e CAEmitterLayer) Lifetime() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("lifetime"))
	return rv
}
func (e CAEmitterLayer) SetLifetime(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setLifetime:"), value)
}
// Defines whether the layer flattens the particles into its plane.
//
// # Discussion
// 
// If [true], the layer renders its particles as if they directly inhabit the
// three-dimensional coordinate space of the layer’s superlayer. When
// enabled, the effect of the layer’s `filters`, `backgroundFilters`, and
// shadow related properties is undefined.
// 
// Default is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayer/preservesDepth
func (e CAEmitterLayer) PreservesDepth() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("preservesDepth"))
	return rv
}
func (e CAEmitterLayer) SetPreservesDepth(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setPreservesDepth:"), value)
}
// The angle, in radians, defining a cone around the emission angle.
// Animatable.
//
// See: https://developer.apple.com/documentation/quartzcore/caemittercell/emissionrange
func (e CAEmitterLayer) EmissionRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emissionRange"))
	return rv
}
func (e CAEmitterLayer) SetEmissionRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmissionRange:"), value)
}
// The mathematical constant pi (π), approximately equal to 3.14159.
//
// See: https://developer.apple.com/documentation/Swift/FloatingPoint/pi
func (e CAEmitterLayer) Pi() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pi"))
	return objectivec.Object{ID: rv}
}
func (e CAEmitterLayer) SetPi(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setPi:"), value)
}
// Particles are emitted from a single point at (
//
// See: https://developer.apple.com/documentation/quartzcore/caemitterlayeremittershape/point
func (e CAEmitterLayer) Point() CAEmitterLayerEmitterShape {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("kCAEmitterLayerPoint"))
	return CAEmitterLayerEmitterShape(foundation.NSStringFromID(rv).String())
}

