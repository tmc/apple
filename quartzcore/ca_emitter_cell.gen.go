// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAEmitterCell] class.
var (
	_CAEmitterCellClass     CAEmitterCellClass
	_CAEmitterCellClassOnce sync.Once
)

func getCAEmitterCellClass() CAEmitterCellClass {
	_CAEmitterCellClassOnce.Do(func() {
		_CAEmitterCellClass = CAEmitterCellClass{class: objc.GetClass("CAEmitterCell")}
	})
	return _CAEmitterCellClass
}

// GetCAEmitterCellClass returns the class object for CAEmitterCell.
func GetCAEmitterCellClass() CAEmitterCellClass {
	return getCAEmitterCellClass()
}

type CAEmitterCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAEmitterCellClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAEmitterCellClass) Alloc() CAEmitterCell {
	rv := objc.Send[CAEmitterCell](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The definition of a particle emitted by a particle layer.
//
// # Overview
//
// The [CAEmitterCell] class represents one source of particles being emitted
// by a [CAEmitterLayer] object. An emitter cell defines the direction and
// properties of the emitted particles. Emitter cells can have an array of
// sub-cells, which lets the particles themselves emit particles.
//
// # Providing Emitter Cell Content
//
//   - [CAEmitterCell.Contents]: An object that provides the contents of the layer. Animatable.
//   - [CAEmitterCell.SetContents]
//   - [CAEmitterCell.ContentsRect]: A rectangle (in the unit coordinate space) that specifies the portion of [contents](<doc://com.apple.quartzcore/documentation/QuartzCore/CAEmitterCell/contents>) that the receiver should draw. Animatable.
//   - [CAEmitterCell.SetContentsRect]
//   - [CAEmitterCell.EmitterCells]: An optional array containing the sub-cells of this cell.
//   - [CAEmitterCell.SetEmitterCells]
//
// # Setting Emitter Cell Visual Attributes
//
//   - [CAEmitterCell.Enabled]: A Boolean value indicating whether or not cells from this emitter are rendered.
//   - [CAEmitterCell.SetEnabled]
//   - [CAEmitterCell.Color]: The color of each emitted object. Animatable.
//   - [CAEmitterCell.SetColor]
//   - [CAEmitterCell.RedRange]: The amount by which the red color component of the cell can vary. Animatable.
//   - [CAEmitterCell.SetRedRange]
//   - [CAEmitterCell.GreenRange]: The amount by which the green color component of the cell can vary. Animatable.
//   - [CAEmitterCell.SetGreenRange]
//   - [CAEmitterCell.BlueRange]: The amount by which the blue color component of the cell can vary. Animatable.
//   - [CAEmitterCell.SetBlueRange]
//   - [CAEmitterCell.AlphaRange]: The amount by which the alpha component of the cell can vary. Animatable.
//   - [CAEmitterCell.SetAlphaRange]
//   - [CAEmitterCell.RedSpeed]: The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
//   - [CAEmitterCell.SetRedSpeed]
//   - [CAEmitterCell.GreenSpeed]: The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
//   - [CAEmitterCell.SetGreenSpeed]
//   - [CAEmitterCell.BlueSpeed]: The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
//   - [CAEmitterCell.SetBlueSpeed]
//   - [CAEmitterCell.AlphaSpeed]: The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
//   - [CAEmitterCell.SetAlphaSpeed]
//   - [CAEmitterCell.MagnificationFilter]: The filter used when increasing the size of the content.
//   - [CAEmitterCell.SetMagnificationFilter]
//   - [CAEmitterCell.MinificationFilter]: The filter used when reducing the size of the content.
//   - [CAEmitterCell.SetMinificationFilter]
//   - [CAEmitterCell.MinificationFilterBias]: The bias factor used by the minification filter to determine the levels of detail.
//   - [CAEmitterCell.SetMinificationFilterBias]
//   - [CAEmitterCell.Scale]: Specifies the scale factor applied to the cell. Animatable.
//   - [CAEmitterCell.SetScale]
//   - [CAEmitterCell.ScaleRange]: Specifies the range over which the scale value can vary. Animatable.
//   - [CAEmitterCell.SetScaleRange]
//   - [CAEmitterCell.ContentsScale]: The scale factor of the cell contents.
//   - [CAEmitterCell.SetContentsScale]
//   - [CAEmitterCell.Name]: The name of the cell.
//   - [CAEmitterCell.SetName]
//   - [CAEmitterCell.Style]: An optional dictionary containing additional style values that are not explicitly defined by the receiver.
//   - [CAEmitterCell.SetStyle]
//
// # Setting Emitter Cell Motion Attributes
//
//   - [CAEmitterCell.Spin]: The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
//   - [CAEmitterCell.SetSpin]
//   - [CAEmitterCell.SpinRange]: The amount by which the spin of the cell can vary over its lifetime. Animatable.
//   - [CAEmitterCell.SetSpinRange]
//   - [CAEmitterCell.EmissionLatitude]: The latitudinal orientation of the emission angle. Animatable.
//   - [CAEmitterCell.SetEmissionLatitude]
//   - [CAEmitterCell.EmissionLongitude]: The longitudinal orientation of the emission angle. Animatable.
//   - [CAEmitterCell.SetEmissionLongitude]
//   - [CAEmitterCell.EmissionRange]: The angle, in radians, defining a cone around the emission angle. Animatable.
//   - [CAEmitterCell.SetEmissionRange]
//
// # Setting Emitter Cell Temporal Attributes
//
//   - [CAEmitterCell.Lifetime]: The lifetime of the cell, in seconds. Animatable.
//   - [CAEmitterCell.SetLifetime]
//   - [CAEmitterCell.LifetimeRange]: The mean value by which the [lifetime](<doc://com.apple.quartzcore/documentation/QuartzCore/CAEmitterCell/lifetime>) of the cell can vary. Animatable.
//   - [CAEmitterCell.SetLifetimeRange]
//   - [CAEmitterCell.BirthRate]: The number of emitted objects created every second. Animatable.
//   - [CAEmitterCell.SetBirthRate]
//   - [CAEmitterCell.ScaleSpeed]: The speed at which the scale changes over the lifetime of the cell. Animatable.
//   - [CAEmitterCell.SetScaleSpeed]
//   - [CAEmitterCell.Velocity]: The initial velocity of the cell. Animatable.
//   - [CAEmitterCell.SetVelocity]
//   - [CAEmitterCell.VelocityRange]: The amount by which the velocity of the cell can vary. Animatable.
//   - [CAEmitterCell.SetVelocityRange]
//   - [CAEmitterCell.XAcceleration]: The x component of an acceleration vector applied to cell.
//   - [CAEmitterCell.SetXAcceleration]
//   - [CAEmitterCell.YAcceleration]: The y component of an acceleration vector applied to cell.
//   - [CAEmitterCell.SetYAcceleration]
//   - [CAEmitterCell.ZAcceleration]: The z component of an acceleration vector applied to cell.
//   - [CAEmitterCell.SetZAcceleration]
//
// # Using Key-Value Coding Extensions
//
//   - [CAEmitterCell.ShouldArchiveValueForKey]: Returns a Boolean value indicating whether the value for a given key should be archived.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell
type CAEmitterCell struct {
	objectivec.Object
}

// CAEmitterCellFromID constructs a [CAEmitterCell] from an objc.ID.
//
// The definition of a particle emitted by a particle layer.
func CAEmitterCellFromID(id objc.ID) CAEmitterCell {
	return CAEmitterCell{objectivec.Object{ID: id}}
}

// NOTE: CAEmitterCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAEmitterCell] class.
//
// # Providing Emitter Cell Content
//
//   - [ICAEmitterCell.Contents]: An object that provides the contents of the layer. Animatable.
//   - [ICAEmitterCell.SetContents]
//   - [ICAEmitterCell.ContentsRect]: A rectangle (in the unit coordinate space) that specifies the portion of [contents](<doc://com.apple.quartzcore/documentation/QuartzCore/CAEmitterCell/contents>) that the receiver should draw. Animatable.
//   - [ICAEmitterCell.SetContentsRect]
//   - [ICAEmitterCell.EmitterCells]: An optional array containing the sub-cells of this cell.
//   - [ICAEmitterCell.SetEmitterCells]
//
// # Setting Emitter Cell Visual Attributes
//
//   - [ICAEmitterCell.Enabled]: A Boolean value indicating whether or not cells from this emitter are rendered.
//   - [ICAEmitterCell.SetEnabled]
//   - [ICAEmitterCell.Color]: The color of each emitted object. Animatable.
//   - [ICAEmitterCell.SetColor]
//   - [ICAEmitterCell.RedRange]: The amount by which the red color component of the cell can vary. Animatable.
//   - [ICAEmitterCell.SetRedRange]
//   - [ICAEmitterCell.GreenRange]: The amount by which the green color component of the cell can vary. Animatable.
//   - [ICAEmitterCell.SetGreenRange]
//   - [ICAEmitterCell.BlueRange]: The amount by which the blue color component of the cell can vary. Animatable.
//   - [ICAEmitterCell.SetBlueRange]
//   - [ICAEmitterCell.AlphaRange]: The amount by which the alpha component of the cell can vary. Animatable.
//   - [ICAEmitterCell.SetAlphaRange]
//   - [ICAEmitterCell.RedSpeed]: The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
//   - [ICAEmitterCell.SetRedSpeed]
//   - [ICAEmitterCell.GreenSpeed]: The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
//   - [ICAEmitterCell.SetGreenSpeed]
//   - [ICAEmitterCell.BlueSpeed]: The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
//   - [ICAEmitterCell.SetBlueSpeed]
//   - [ICAEmitterCell.AlphaSpeed]: The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
//   - [ICAEmitterCell.SetAlphaSpeed]
//   - [ICAEmitterCell.MagnificationFilter]: The filter used when increasing the size of the content.
//   - [ICAEmitterCell.SetMagnificationFilter]
//   - [ICAEmitterCell.MinificationFilter]: The filter used when reducing the size of the content.
//   - [ICAEmitterCell.SetMinificationFilter]
//   - [ICAEmitterCell.MinificationFilterBias]: The bias factor used by the minification filter to determine the levels of detail.
//   - [ICAEmitterCell.SetMinificationFilterBias]
//   - [ICAEmitterCell.Scale]: Specifies the scale factor applied to the cell. Animatable.
//   - [ICAEmitterCell.SetScale]
//   - [ICAEmitterCell.ScaleRange]: Specifies the range over which the scale value can vary. Animatable.
//   - [ICAEmitterCell.SetScaleRange]
//   - [ICAEmitterCell.ContentsScale]: The scale factor of the cell contents.
//   - [ICAEmitterCell.SetContentsScale]
//   - [ICAEmitterCell.Name]: The name of the cell.
//   - [ICAEmitterCell.SetName]
//   - [ICAEmitterCell.Style]: An optional dictionary containing additional style values that are not explicitly defined by the receiver.
//   - [ICAEmitterCell.SetStyle]
//
// # Setting Emitter Cell Motion Attributes
//
//   - [ICAEmitterCell.Spin]: The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
//   - [ICAEmitterCell.SetSpin]
//   - [ICAEmitterCell.SpinRange]: The amount by which the spin of the cell can vary over its lifetime. Animatable.
//   - [ICAEmitterCell.SetSpinRange]
//   - [ICAEmitterCell.EmissionLatitude]: The latitudinal orientation of the emission angle. Animatable.
//   - [ICAEmitterCell.SetEmissionLatitude]
//   - [ICAEmitterCell.EmissionLongitude]: The longitudinal orientation of the emission angle. Animatable.
//   - [ICAEmitterCell.SetEmissionLongitude]
//   - [ICAEmitterCell.EmissionRange]: The angle, in radians, defining a cone around the emission angle. Animatable.
//   - [ICAEmitterCell.SetEmissionRange]
//
// # Setting Emitter Cell Temporal Attributes
//
//   - [ICAEmitterCell.Lifetime]: The lifetime of the cell, in seconds. Animatable.
//   - [ICAEmitterCell.SetLifetime]
//   - [ICAEmitterCell.LifetimeRange]: The mean value by which the [lifetime](<doc://com.apple.quartzcore/documentation/QuartzCore/CAEmitterCell/lifetime>) of the cell can vary. Animatable.
//   - [ICAEmitterCell.SetLifetimeRange]
//   - [ICAEmitterCell.BirthRate]: The number of emitted objects created every second. Animatable.
//   - [ICAEmitterCell.SetBirthRate]
//   - [ICAEmitterCell.ScaleSpeed]: The speed at which the scale changes over the lifetime of the cell. Animatable.
//   - [ICAEmitterCell.SetScaleSpeed]
//   - [ICAEmitterCell.Velocity]: The initial velocity of the cell. Animatable.
//   - [ICAEmitterCell.SetVelocity]
//   - [ICAEmitterCell.VelocityRange]: The amount by which the velocity of the cell can vary. Animatable.
//   - [ICAEmitterCell.SetVelocityRange]
//   - [ICAEmitterCell.XAcceleration]: The x component of an acceleration vector applied to cell.
//   - [ICAEmitterCell.SetXAcceleration]
//   - [ICAEmitterCell.YAcceleration]: The y component of an acceleration vector applied to cell.
//   - [ICAEmitterCell.SetYAcceleration]
//   - [ICAEmitterCell.ZAcceleration]: The z component of an acceleration vector applied to cell.
//   - [ICAEmitterCell.SetZAcceleration]
//
// # Using Key-Value Coding Extensions
//
//   - [ICAEmitterCell.ShouldArchiveValueForKey]: Returns a Boolean value indicating whether the value for a given key should be archived.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell
type ICAEmitterCell interface {
	objectivec.IObject
	CAMediaTiming

	// Topic: Providing Emitter Cell Content

	// An object that provides the contents of the layer. Animatable.
	Contents() objectivec.IObject
	SetContents(value objectivec.IObject)
	// A rectangle (in the unit coordinate space) that specifies the portion of [contents](<doc://com.apple.quartzcore/documentation/QuartzCore/CAEmitterCell/contents>) that the receiver should draw. Animatable.
	ContentsRect() corefoundation.CGRect
	SetContentsRect(value corefoundation.CGRect)
	// An optional array containing the sub-cells of this cell.
	EmitterCells() []CAEmitterCell
	SetEmitterCells(value []CAEmitterCell)

	// Topic: Setting Emitter Cell Visual Attributes

	// A Boolean value indicating whether or not cells from this emitter are rendered.
	Enabled() bool
	SetEnabled(value bool)
	// The color of each emitted object. Animatable.
	Color() coregraphics.CGColorRef
	SetColor(value coregraphics.CGColorRef)
	// The amount by which the red color component of the cell can vary. Animatable.
	RedRange() float32
	SetRedRange(value float32)
	// The amount by which the green color component of the cell can vary. Animatable.
	GreenRange() float32
	SetGreenRange(value float32)
	// The amount by which the blue color component of the cell can vary. Animatable.
	BlueRange() float32
	SetBlueRange(value float32)
	// The amount by which the alpha component of the cell can vary. Animatable.
	AlphaRange() float32
	SetAlphaRange(value float32)
	// The speed, in seconds, at which the red color component changes over the lifetime of the cell. Animatable.
	RedSpeed() float32
	SetRedSpeed(value float32)
	// The speed, in seconds, at which the green color component changes over the lifetime of the cell. Animatable.
	GreenSpeed() float32
	SetGreenSpeed(value float32)
	// The speed, in seconds, at which the blue color component changes over the lifetime of the cell. Animatable.
	BlueSpeed() float32
	SetBlueSpeed(value float32)
	// The speed, in seconds, at which the alpha component changes over the lifetime of the cell. Animatable.
	AlphaSpeed() float32
	SetAlphaSpeed(value float32)
	// The filter used when increasing the size of the content.
	MagnificationFilter() string
	SetMagnificationFilter(value string)
	// The filter used when reducing the size of the content.
	MinificationFilter() string
	SetMinificationFilter(value string)
	// The bias factor used by the minification filter to determine the levels of detail.
	MinificationFilterBias() float32
	SetMinificationFilterBias(value float32)
	// Specifies the scale factor applied to the cell. Animatable.
	Scale() float64
	SetScale(value float64)
	// Specifies the range over which the scale value can vary. Animatable.
	ScaleRange() float64
	SetScaleRange(value float64)
	// The scale factor of the cell contents.
	ContentsScale() float64
	SetContentsScale(value float64)
	// The name of the cell.
	Name() string
	SetName(value string)
	// An optional dictionary containing additional style values that are not explicitly defined by the receiver.
	Style() foundation.INSDictionary
	SetStyle(value foundation.INSDictionary)

	// Topic: Setting Emitter Cell Motion Attributes

	// The rotational velocity, measured in radians per second, to apply to the cell. Animatable.
	Spin() float64
	SetSpin(value float64)
	// The amount by which the spin of the cell can vary over its lifetime. Animatable.
	SpinRange() float64
	SetSpinRange(value float64)
	// The latitudinal orientation of the emission angle. Animatable.
	EmissionLatitude() float64
	SetEmissionLatitude(value float64)
	// The longitudinal orientation of the emission angle. Animatable.
	EmissionLongitude() float64
	SetEmissionLongitude(value float64)
	// The angle, in radians, defining a cone around the emission angle. Animatable.
	EmissionRange() float64
	SetEmissionRange(value float64)

	// Topic: Setting Emitter Cell Temporal Attributes

	// The lifetime of the cell, in seconds. Animatable.
	Lifetime() float32
	SetLifetime(value float32)
	// The mean value by which the [lifetime](<doc://com.apple.quartzcore/documentation/QuartzCore/CAEmitterCell/lifetime>) of the cell can vary. Animatable.
	LifetimeRange() float32
	SetLifetimeRange(value float32)
	// The number of emitted objects created every second. Animatable.
	BirthRate() float32
	SetBirthRate(value float32)
	// The speed at which the scale changes over the lifetime of the cell. Animatable.
	ScaleSpeed() float64
	SetScaleSpeed(value float64)
	// The initial velocity of the cell. Animatable.
	Velocity() float64
	SetVelocity(value float64)
	// The amount by which the velocity of the cell can vary. Animatable.
	VelocityRange() float64
	SetVelocityRange(value float64)
	// The x component of an acceleration vector applied to cell.
	XAcceleration() float64
	SetXAcceleration(value float64)
	// The y component of an acceleration vector applied to cell.
	YAcceleration() float64
	SetYAcceleration(value float64)
	// The z component of an acceleration vector applied to cell.
	ZAcceleration() float64
	SetZAcceleration(value float64)

	// Topic: Using Key-Value Coding Extensions

	// Returns a Boolean value indicating whether the value for a given key should be archived.
	ShouldArchiveValueForKey(key string) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (e CAEmitterCell) Init() CAEmitterCell {
	rv := objc.Send[CAEmitterCell](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e CAEmitterCell) Autorelease() CAEmitterCell {
	rv := objc.Send[CAEmitterCell](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAEmitterCell creates a new CAEmitterCell instance.
func NewCAEmitterCell() CAEmitterCell {
	class := getCAEmitterCellClass()
	rv := objc.Send[CAEmitterCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Boolean value indicating whether the value for a given key should
// be archived.
//
// key: The name of one of the receiver’s properties.
//
// # Return Value
//
// true if the specified property should be archived, otherwise false.
//
// # Discussion
//
// The default implementation returns true. This method is called by the
// object’s implementation of “.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/shouldArchiveValue(forKey:)
func (e CAEmitterCell) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("shouldArchiveValueForKey:"), objc.String(key))
	return rv
}
func (e CAEmitterCell) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the default value of the property with the specified key.
//
// key: The name of one of the receiver’s properties.
//
// # Return Value
//
// The default value for the named property. Returns `nil` if no default value
// has been set.
//
// # Discussion
//
// If this method returns `nil` a suitable “zero” default value for the
// property is provided, based on the declared type of the `key`. For example,
// if `key` is a [CGSize] object, a size of (0.0,0.0) is returned. For a
// [CGRect] an empty rectangle is returned. For [CGAffineTransform] and
// [CATransform3D], the appropriate identity matrix is returned.
//
// # Special Considerations
//
// If `key` is not a known for property of the class, the result of the method
// is undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/defaultValue(forKey:)
func (_CAEmitterCellClass CAEmitterCellClass) DefaultValueForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_CAEmitterCellClass.class), objc.Sel("defaultValueForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Creates and returns an instance of [CAEmitterCell].
//
// # Return Value
//
// The initialized [CAEmitterCell] object or `nil` if initialization is not
// successful.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCell
func (_CAEmitterCellClass CAEmitterCellClass) EmitterCell() CAEmitterCell {
	rv := objc.Send[objc.ID](objc.ID(_CAEmitterCellClass.class), objc.Sel("emitterCell"))
	return CAEmitterCellFromID(rv)
}

// An object that provides the contents of the layer. Animatable.
//
// # Discussion
//
// A layer can set this property to a [CGImage] to display the image as its
// contents.
//
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contents
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
func (e CAEmitterCell) Contents() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("contents"))
	return objectivec.Object{ID: rv}
}
func (e CAEmitterCell) SetContents(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setContents:"), value)
}

// A rectangle (in the unit coordinate space) that specifies the portion of
// [Contents] that the receiver should draw. Animatable.
//
// # Discussion
//
// By default, this property is set to the unit rectangle (0.0,0.0,1.0,1.0),
// which results in all of the layer’s contents being drawn.
//
// If pixels outside the unit rectangle are requested, the edge pixels of the
// contents image are extended outwards.
//
// If you assign an empty rectangle to this property, the results are
// undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsRect
func (e CAEmitterCell) ContentsRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](e.ID, objc.Sel("contentsRect"))
	return corefoundation.CGRect(rv)
}
func (e CAEmitterCell) SetContentsRect(value corefoundation.CGRect) {
	objc.Send[struct{}](e.ID, objc.Sel("setContentsRect:"), value)
}

// An optional array containing the sub-cells of this cell.
//
// # Discussion
//
// When specified, each particle emitted by the cell acts as an emitter for
// each of the cell’s sub-cells. The emission point is the current particle
// position and the emission angle is relative to the current direction of the
// particle.
//
// The default value of this property is `nil`.
//
// The following code shows how you can create a firework style effect using
// sub-cells. The `fireworkCell` has an emission longitude of one quarter turn
// anti-clockwise to emit particles upwards. It emits `trailCell` instances
// which have a slight [YAcceleration] that simulates gravity.
//
// Note that the [Scale] and [Color] of `fireworkCell` are inherited by
// `trailCell`.
//
// Listing 1. Creating particle trails
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emitterCells
func (e CAEmitterCell) EmitterCells() []CAEmitterCell {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("emitterCells"))
	return objc.ConvertSlice(rv, func(id objc.ID) CAEmitterCell {
		return CAEmitterCellFromID(id)
	})
}
func (e CAEmitterCell) SetEmitterCells(value []CAEmitterCell) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmitterCells:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value indicating whether or not cells from this emitter are
// rendered.
//
// # Discussion
//
// The default value of this property is true.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/isEnabled
func (e CAEmitterCell) Enabled() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isEnabled"))
	return rv
}
func (e CAEmitterCell) SetEnabled(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setEnabled:"), value)
}

// The color of each emitted object. Animatable.
//
// # Discussion
//
// The specified color of the cell will vary by a random amount within the
// [RedRange], [GreenRange], [BlueRange] and [AlphaRange]values over the
// lifetime of the cell. The [RedSpeed], [GreenSpeed], [BlueSpeed], and
// [AlphaSpeed] determine the rate of change.
//
// The default value of this property is a color object set to opaque white.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/color
func (e CAEmitterCell) Color() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](e.ID, objc.Sel("color"))
	return coregraphics.CGColorRef(rv)
}
func (e CAEmitterCell) SetColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](e.ID, objc.Sel("setColor:"), value)
}

// The amount by which the red color component of the cell can vary.
// Animatable.
//
// # Discussion
//
// The range specifies the mean amount by which the red component of the
// [Color] property can vary for the cell.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redRange
func (e CAEmitterCell) RedRange() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("redRange"))
	return rv
}
func (e CAEmitterCell) SetRedRange(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setRedRange:"), value)
}

// The amount by which the green color component of the cell can vary.
// Animatable.
//
// # Discussion
//
// The range specifies the mean amount by which the green component of the
// [Color] property can vary for the cell.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenRange
func (e CAEmitterCell) GreenRange() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("greenRange"))
	return rv
}
func (e CAEmitterCell) SetGreenRange(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setGreenRange:"), value)
}

// The amount by which the blue color component of the cell can vary.
// Animatable.
//
// # Discussion
//
// The range specifies the mean amount by which the blue component of the
// [Color] property can vary for the cell.
//
// The default value of this property value is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueRange
func (e CAEmitterCell) BlueRange() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("blueRange"))
	return rv
}
func (e CAEmitterCell) SetBlueRange(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setBlueRange:"), value)
}

// The amount by which the alpha component of the cell can vary. Animatable.
//
// # Discussion
//
// The range specifies the mean amount by which the alpha component of the
// [Color] property can vary for the cell.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaRange
func (e CAEmitterCell) AlphaRange() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("alphaRange"))
	return rv
}
func (e CAEmitterCell) SetAlphaRange(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setAlphaRange:"), value)
}

// The speed, in seconds, at which the red color component changes over the
// lifetime of the cell. Animatable.
//
// # Discussion
//
// The speed change is defined as the rate of change per second.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/redSpeed
func (e CAEmitterCell) RedSpeed() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("redSpeed"))
	return rv
}
func (e CAEmitterCell) SetRedSpeed(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setRedSpeed:"), value)
}

// The speed, in seconds, at which the green color component changes over the
// lifetime of the cell. Animatable.
//
// # Discussion
//
// The speed change is defined as the rate of change per second.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/greenSpeed
func (e CAEmitterCell) GreenSpeed() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("greenSpeed"))
	return rv
}
func (e CAEmitterCell) SetGreenSpeed(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setGreenSpeed:"), value)
}

// The speed, in seconds, at which the blue color component changes over the
// lifetime of the cell. Animatable.
//
// # Discussion
//
// The speed change is defined as the rate of change per second.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/blueSpeed
func (e CAEmitterCell) BlueSpeed() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("blueSpeed"))
	return rv
}
func (e CAEmitterCell) SetBlueSpeed(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setBlueSpeed:"), value)
}

// The speed, in seconds, at which the alpha component changes over the
// lifetime of the cell. Animatable.
//
// # Discussion
//
// The speed change is defined as the rate of change per second.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/alphaSpeed
func (e CAEmitterCell) AlphaSpeed() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("alphaSpeed"))
	return rv
}
func (e CAEmitterCell) SetAlphaSpeed(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setAlphaSpeed:"), value)
}

// The filter used when increasing the size of the content.
//
// # Discussion
//
// The possible values for this property are listed in Scaling Filters. The
// default value is [linear].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/magnificationFilter
//
// [linear]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/linear
func (e CAEmitterCell) MagnificationFilter() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("magnificationFilter"))
	return foundation.NSStringFromID(rv).String()
}
func (e CAEmitterCell) SetMagnificationFilter(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setMagnificationFilter:"), objc.String(value))
}

// The filter used when reducing the size of the content.
//
// # Discussion
//
// The possible values for this property are listed in Scaling Filters. The
// default value is [linear].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilter
//
// [linear]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/linear
func (e CAEmitterCell) MinificationFilter() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("minificationFilter"))
	return foundation.NSStringFromID(rv).String()
}
func (e CAEmitterCell) SetMinificationFilter(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setMinificationFilter:"), objc.String(value))
}

// The bias factor used by the minification filter to determine the levels of
// detail.
//
// # Discussion
//
// This value is used by the [MinificationFilter] property when it is set to
// `kCAFilterTrilinear`.
//
// The default value of this property to `0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/minificationFilterBias
func (e CAEmitterCell) MinificationFilterBias() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("minificationFilterBias"))
	return rv
}
func (e CAEmitterCell) SetMinificationFilterBias(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setMinificationFilterBias:"), value)
}

// Specifies the scale factor applied to the cell. Animatable.
//
// # Discussion
//
// The scale of the cell will vary by a random amount within the range
// specified by [ScaleRange]. The [ScaleSpeed] property determines the rate of
// change.
//
// The default value of this property is `1.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scale
func (e CAEmitterCell) Scale() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("scale"))
	return rv
}
func (e CAEmitterCell) SetScale(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setScale:"), value)
}

// Specifies the range over which the scale value can vary. Animatable.
//
// # Discussion
//
// The range specifies the mean amount that the [Scale] value can vary for the
// cell over its lifetime.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleRange
func (e CAEmitterCell) ScaleRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("scaleRange"))
	return rv
}
func (e CAEmitterCell) SetScaleRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setScaleRange:"), value)
}

// The scale factor of the cell contents.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/contentsScale
func (e CAEmitterCell) ContentsScale() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("contentsScale"))
	return rv
}
func (e CAEmitterCell) SetContentsScale(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setContentsScale:"), value)
}

// The name of the cell.
//
// # Discussion
//
// The cell name is used when constructing animation key paths that reference
// the cell.
//
// For example, adding an animation to a cell’s enclosing layer with the a
// keypath such as `emitterCells.MyCellNameXCUIElementTypeRedRange()` would
// animate the `redRange` property of the cell in the layer’s emitterCells
// array with the name `myCellName`.
//
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/name
func (e CAEmitterCell) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (e CAEmitterCell) SetName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setName:"), objc.String(value))
}

// An optional dictionary containing additional style values that are not
// explicitly defined by the receiver.
//
// # Discussion
//
// This dictionary may in turn have a `style` key, forming a hierarchy of
// default values. In the case of hierarchical style dictionaries the
// shallowest value for a property is used. For example, the value for
// “style.someValue” takes precedence over “style.style.someValue”.
//
// If the style dictionary doesn’t define a value for an attribute, the
// cell’s [DefaultValueForKey] class method is called.
//
// The style dictionary is not consulted for the following keys: `bounds`,
// `frame`.
//
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/style
func (e CAEmitterCell) Style() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("style"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e CAEmitterCell) SetStyle(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setStyle:"), value)
}

// The rotational velocity, measured in radians per second, to apply to the
// cell. Animatable.
//
// # Discussion
//
// The spin of the cell will vary by a random amount with the range specified
// by [SpinRange].
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spin
func (e CAEmitterCell) Spin() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("spin"))
	return rv
}
func (e CAEmitterCell) SetSpin(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setSpin:"), value)
}

// The amount by which the spin of the cell can vary over its lifetime.
// Animatable.
//
// # Discussion
//
// The range specifies the mean amount the [Spin] value can vary over the
// cell’s lifetime.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/spinRange
func (e CAEmitterCell) SpinRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("spinRange"))
	return rv
}
func (e CAEmitterCell) SetSpinRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setSpinRange:"), value)
}

// The latitudinal orientation of the emission angle. Animatable.
//
// # Discussion
//
// The emission latitude is the orientation of the emission angle from the
// z-axis. It is also referred to as the colatitude.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLatitude
func (e CAEmitterCell) EmissionLatitude() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emissionLatitude"))
	return rv
}
func (e CAEmitterCell) SetEmissionLatitude(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmissionLatitude:"), value)
}

// The longitudinal orientation of the emission angle. Animatable.
//
// # Discussion
//
// The emission longitude is the orientation of the emission angle in the
// xy-plane. it is also often referred to as the azimuth.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionLongitude
func (e CAEmitterCell) EmissionLongitude() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emissionLongitude"))
	return rv
}
func (e CAEmitterCell) SetEmissionLongitude(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmissionLongitude:"), value)
}

// The angle, in radians, defining a cone around the emission angle.
// Animatable.
//
// # Discussion
//
// Cells are uniformly distributed across this cone.
//
// The default value of this property is `0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/emissionRange
func (e CAEmitterCell) EmissionRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emissionRange"))
	return rv
}
func (e CAEmitterCell) SetEmissionRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmissionRange:"), value)
}

// The lifetime of the cell, in seconds. Animatable.
//
// # Discussion
//
// The lifetime of the cell will vary by a random amount with the range
// specified by [LifetimeRange].
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetime
func (e CAEmitterCell) Lifetime() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("lifetime"))
	return rv
}
func (e CAEmitterCell) SetLifetime(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setLifetime:"), value)
}

// The mean value by which the [Lifetime] of the cell can vary. Animatable.
//
// # Discussion
//
// If the [LifetimeRange] is 3 seconds, and the [Lifetime] of the cell is 10
// seconds, the cell’s actual lifetime will be between 7 and 13 seconds.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/lifetimeRange
func (e CAEmitterCell) LifetimeRange() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("lifetimeRange"))
	return rv
}
func (e CAEmitterCell) SetLifetimeRange(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setLifetimeRange:"), value)
}

// The number of emitted objects created every second. Animatable.
//
// # Discussion
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/birthRate
func (e CAEmitterCell) BirthRate() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("birthRate"))
	return rv
}
func (e CAEmitterCell) SetBirthRate(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setBirthRate:"), value)
}

// The speed at which the scale changes over the lifetime of the cell.
// Animatable.
//
// # Discussion
//
// The speed change is defined as the rate of change per second.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/scaleSpeed
func (e CAEmitterCell) ScaleSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("scaleSpeed"))
	return rv
}
func (e CAEmitterCell) SetScaleSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setScaleSpeed:"), value)
}

// The initial velocity of the cell. Animatable.
//
// # Discussion
//
// The velocity of the cell will vary by a random amount within the range
// specified by [VelocityRange].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocity
func (e CAEmitterCell) Velocity() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("velocity"))
	return rv
}
func (e CAEmitterCell) SetVelocity(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setVelocity:"), value)
}

// The amount by which the velocity of the cell can vary. Animatable.
//
// # Discussion
//
// The range specifies the mean amount the initial [Velocity] value change.
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/velocityRange
func (e CAEmitterCell) VelocityRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("velocityRange"))
	return rv
}
func (e CAEmitterCell) SetVelocityRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setVelocityRange:"), value)
}

// The x component of an acceleration vector applied to cell.
//
// # Discussion
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/xAcceleration
func (e CAEmitterCell) XAcceleration() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("xAcceleration"))
	return rv
}
func (e CAEmitterCell) SetXAcceleration(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setXAcceleration:"), value)
}

// The y component of an acceleration vector applied to cell.
//
// # Discussion
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/yAcceleration
func (e CAEmitterCell) YAcceleration() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("yAcceleration"))
	return rv
}
func (e CAEmitterCell) SetYAcceleration(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setYAcceleration:"), value)
}

// The z component of an acceleration vector applied to cell.
//
// # Discussion
//
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterCell/zAcceleration
func (e CAEmitterCell) ZAcceleration() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("zAcceleration"))
	return rv
}
func (e CAEmitterCell) SetZAcceleration(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setZAcceleration:"), value)
}

// Determines if the receiver plays in the reverse upon completion.
//
// # Discussion
//
// When true, the receiver plays backwards after playing forwards. Defaults to
// false.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/autoreverses
func (e CAEmitterCell) Autoreverses() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("autoreverses"))
	return rv
}
func (e CAEmitterCell) SetAutoreverses(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setAutoreverses:"), value)
}

// Specifies the begin time of the receiver in relation to its parent object,
// if applicable.
//
// # Discussion
//
// Defaults to 0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (e CAEmitterCell) BeginTime() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("beginTime"))
	return rv
}
func (e CAEmitterCell) SetBeginTime(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setBeginTime:"), value)
}

// Specifies the basic duration of the animation, in seconds.
//
// # Discussion
//
// Defaults to 0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/duration
func (e CAEmitterCell) Duration() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("duration"))
	return rv
}
func (e CAEmitterCell) SetDuration(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setDuration:"), value)
}

// Determines if the receiver’s presentation is frozen or removed once its
// active duration has completed.
//
// # Discussion
//
// The possible values are described in [Fill Modes]. The default is
// [removed].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/fillMode
//
// [Fill Modes]: https://developer.apple.com/documentation/QuartzCore/fill-modes
// [removed]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFillMode/removed
func (e CAEmitterCell) FillMode() CAMediaTimingFillMode {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("fillMode"))
	return CAMediaTimingFillMode(foundation.NSStringFromID(rv).String())
}
func (e CAEmitterCell) SetFillMode(value CAMediaTimingFillMode) {
	objc.Send[struct{}](e.ID, objc.Sel("setFillMode:"), objc.String(string(value)))
}

// Determines the number of times the animation will repeat.
//
// # Discussion
//
// May be fractional. If the `repeatCount` is 0, it is ignored. Defaults to 0.
// If both [RepeatDuration] and [RepeatCount] are specified the behavior is
// undefined.
//
// Setting this property to [greatestFiniteMagnitude] will cause the animation
// to repeat forever.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatCount
//
// [greatestFiniteMagnitude]: https://developer.apple.com/documentation/Swift/Float/greatestFiniteMagnitude
func (e CAEmitterCell) RepeatCount() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("repeatCount"))
	return rv
}
func (e CAEmitterCell) SetRepeatCount(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setRepeatCount:"), value)
}

// Determines how many seconds the animation will repeat for.
//
// # Discussion
//
// Defaults to 0. If the `repeatDuration` is 0, it is ignored. If both
// [RepeatDuration] and [RepeatCount] are specified the behavior is undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatDuration
func (e CAEmitterCell) RepeatDuration() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("repeatDuration"))
	return rv
}
func (e CAEmitterCell) SetRepeatDuration(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setRepeatDuration:"), value)
}

// Specifies how time is mapped to receiver’s time space from the parent
// time space.
//
// # Discussion
//
// For example, if `speed` is 2.0 local time progresses twice as fast as
// parent time. Defaults to 1.0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/speed
func (e CAEmitterCell) Speed() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("speed"))
	return rv
}
func (e CAEmitterCell) SetSpeed(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setSpeed:"), value)
}

// Specifies an additional time offset in active local time.
//
// # Discussion
//
// Defaults to 0. .
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/timeOffset
func (e CAEmitterCell) TimeOffset() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("timeOffset"))
	return rv
}
func (e CAEmitterCell) SetTimeOffset(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setTimeOffset:"), value)
}

// Protocol methods for CAMediaTiming
