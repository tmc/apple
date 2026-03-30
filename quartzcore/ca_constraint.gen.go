// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAConstraint] class.
var (
	_CAConstraintClass     CAConstraintClass
	_CAConstraintClassOnce sync.Once
)

func getCAConstraintClass() CAConstraintClass {
	_CAConstraintClassOnce.Do(func() {
		_CAConstraintClass = CAConstraintClass{class: objc.GetClass("CAConstraint")}
	})
	return _CAConstraintClass
}

// GetCAConstraintClass returns the class object for CAConstraint.
func GetCAConstraintClass() CAConstraintClass {
	return getCAConstraintClass()
}

type CAConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAConstraintClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAConstraintClass) Alloc() CAConstraint {
	rv := objc.Send[CAConstraint](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a single layout constraint between two layers.
//
// # Overview
//
// Each [CAConstraint] instance encapsulates one geometry relationship between
// two layers on the same axis.
//
// Sibling layers are referenced by name, using the name property of each
// layer. The special name `superlayer` is used to refer to the layer’s
// superlayer.
//
// For example, to specify that a layer should be horizontally centered in its
// superview you would use the following:
//
// A minimum of two relationships must be specified per axis. If you specify
// constraints for the left and right edges of a layer, the width will vary.
// If you specify constraints for the left edge and the width, the right edge
// of the layer will move relative to the superlayer’s frame. Often you’ll
// specify only a single edge constraint, the layer’s size in the same axis
// will be used as the second relationship.
//
// # Create a New Constraint
//
//   - [CAConstraint.InitWithAttributeRelativeToAttributeScaleOffset]: Returns an [CAConstraint] object with the specified parameters. Designated initializer.
//
// # Accessing Constraint Values
//
//   - [CAConstraint.Attribute]: The attribute the constraint affects.
//   - [CAConstraint.Offset]: Offset value of the constraint attribute.
//   - [CAConstraint.Scale]: Scale factor of the constraint attribute.
//   - [CAConstraint.SourceAttribute]: The constraint attribute of the layer the receiver is calculated relative to
//   - [CAConstraint.SourceName]: Name of the layer that the constraint is calculated relative to.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint
type CAConstraint struct {
	objectivec.Object
}

// CAConstraintFromID constructs a [CAConstraint] from an objc.ID.
//
// A representation of a single layout constraint between two layers.
func CAConstraintFromID(id objc.ID) CAConstraint {
	return CAConstraint{objectivec.Object{ID: id}}
}

// NOTE: CAConstraint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAConstraint] class.
//
// # Create a New Constraint
//
//   - [ICAConstraint.InitWithAttributeRelativeToAttributeScaleOffset]: Returns an [CAConstraint] object with the specified parameters. Designated initializer.
//
// # Accessing Constraint Values
//
//   - [ICAConstraint.Attribute]: The attribute the constraint affects.
//   - [ICAConstraint.Offset]: Offset value of the constraint attribute.
//   - [ICAConstraint.Scale]: Scale factor of the constraint attribute.
//   - [ICAConstraint.SourceAttribute]: The constraint attribute of the layer the receiver is calculated relative to
//   - [ICAConstraint.SourceName]: Name of the layer that the constraint is calculated relative to.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint
type ICAConstraint interface {
	objectivec.IObject

	// Topic: Create a New Constraint

	// Returns an [CAConstraint] object with the specified parameters. Designated initializer.
	InitWithAttributeRelativeToAttributeScaleOffset(attr CAConstraintAttribute, srcId string, srcAttr CAConstraintAttribute, m float64, c float64) CAConstraint

	// Topic: Accessing Constraint Values

	// The attribute the constraint affects.
	Attribute() CAConstraintAttribute
	// Offset value of the constraint attribute.
	Offset() float64
	// Scale factor of the constraint attribute.
	Scale() float64
	// The constraint attribute of the layer the receiver is calculated relative to
	SourceAttribute() CAConstraintAttribute
	// Name of the layer that the constraint is calculated relative to.
	SourceName() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CAConstraint) Init() CAConstraint {
	rv := objc.Send[CAConstraint](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CAConstraint) Autorelease() CAConstraint {
	rv := objc.Send[CAConstraint](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAConstraint creates a new CAConstraint instance.
func NewCAConstraint() CAConstraint {
	class := getCAConstraintClass()
	rv := objc.Send[CAConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns an [CAConstraint] object with the specified parameters.
//
// attr: The attribute of the layer for which to create a new constraint.
//
// srcId: The name of the layer that this constraint is calculated relative to.
//
// srcAttr: The attribute of `srcLayer` the constraint is calculated relative to.
//
// # Return Value
//
// A new [CAConstraint] object with the specified parameters. The scale of the
// constraint is set to 1.0. The offset of the constraint is set to 0.0.
//
// # Discussion
//
// The value for the constraint is calculated is `srcAttr`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/init(attribute:relativeTo:attribute:)
func NewConstraintWithAttributeRelativeToAttribute(attr CAConstraintAttribute, srcId string, srcAttr CAConstraintAttribute) CAConstraint {
	rv := objc.Send[objc.ID](objc.ID(getCAConstraintClass().class), objc.Sel("constraintWithAttribute:relativeTo:attribute:"), attr, objc.String(srcId), srcAttr)
	return CAConstraintFromID(rv)
}

// Creates and returns an [CAConstraint] object with the specified parameters.
//
// attr: The attribute of the layer for which to create a new constraint.
//
// srcId: The name of the layer that this constraint is calculated relative to.
//
// srcAttr: The attribute of `srcLayer` the constraint is calculated relative to.
//
// c: The offset added to the value of `srcAttr`.
//
// # Return Value
//
// A new [CAConstraint] object with the specified parameters. The scale of the
// constraint is set to 1.0.
//
// # Discussion
//
// The value for the constraint is calculated as (`srcAttr` + `offset`).
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/init(attribute:relativeTo:attribute:offset:)
func NewConstraintWithAttributeRelativeToAttributeOffset(attr CAConstraintAttribute, srcId string, srcAttr CAConstraintAttribute, c float64) CAConstraint {
	rv := objc.Send[objc.ID](objc.ID(getCAConstraintClass().class), objc.Sel("constraintWithAttribute:relativeTo:attribute:offset:"), attr, objc.String(srcId), srcAttr, c)
	return CAConstraintFromID(rv)
}

// Returns an [CAConstraint] object with the specified parameters. Designated
// initializer.
//
// attr: The attribute of the layer for which to create a new constraint.
//
// srcId: The name of the layer that this constraint is calculated relative to.
//
// srcAttr: The attribute of `srcLayer` the constraint is calculated relative to.
//
// m: The amount to scale the value of `srcAttr`.
//
// c: The offset added to the value of `srcAttr`.
//
// # Return Value
//
// An initialized constraint object using the specified parameters.
//
// # Discussion
//
// The value for the constraint is calculated as (`srcAttr` * `scale`) +
// `offset`).
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/init(attribute:relativeTo:attribute:scale:offset:)
func NewConstraintWithAttributeRelativeToAttributeScaleOffset(attr CAConstraintAttribute, srcId string, srcAttr CAConstraintAttribute, m float64, c float64) CAConstraint {
	instance := getCAConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAttribute:relativeTo:attribute:scale:offset:"), attr, objc.String(srcId), srcAttr, m, c)
	return CAConstraintFromID(rv)
}

// Returns an [CAConstraint] object with the specified parameters. Designated
// initializer.
//
// attr: The attribute of the layer for which to create a new constraint.
//
// srcId: The name of the layer that this constraint is calculated relative to.
//
// srcAttr: The attribute of `srcLayer` the constraint is calculated relative to.
//
// m: The amount to scale the value of `srcAttr`.
//
// c: The offset added to the value of `srcAttr`.
//
// # Return Value
//
// An initialized constraint object using the specified parameters.
//
// # Discussion
//
// The value for the constraint is calculated as (`srcAttr` * `scale`) +
// `offset`).
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/init(attribute:relativeTo:attribute:scale:offset:)
func (c_ CAConstraint) InitWithAttributeRelativeToAttributeScaleOffset(attr CAConstraintAttribute, srcId string, srcAttr CAConstraintAttribute, m float64, c float64) CAConstraint {
	rv := objc.Send[CAConstraint](c_.ID, objc.Sel("initWithAttribute:relativeTo:attribute:scale:offset:"), attr, objc.String(srcId), srcAttr, m, c)
	return rv
}
func (c CAConstraint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns an [CAConstraint] object with the specified parameters.
//
// attr: The attribute of the layer for which to create a new constraint.
//
// srcId: The name of the layer that this constraint is calculated relative to.
//
// srcAttr: The attribute of `srcLayer` the constraint is calculated relative to.
//
// m: The amount to scale the value of `srcAttr`.
//
// c: The offset from the `srcAttr`.
//
// # Return Value
//
// A new [CAConstraint] object with the specified parameters.
//
// # Discussion
//
// The value for the constraint is calculated as ((`srcAttr` * scale) +
// offset).
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/constraintWithAttribute:relativeTo:attribute:scale:offset:
func (_CAConstraintClass CAConstraintClass) ConstraintWithAttributeRelativeToAttributeScaleOffset(attr CAConstraintAttribute, srcId string, srcAttr CAConstraintAttribute, m float64, c float64) CAConstraint {
	rv := objc.Send[objc.ID](objc.ID(_CAConstraintClass.class), objc.Sel("constraintWithAttribute:relativeTo:attribute:scale:offset:"), attr, objc.String(srcId), srcAttr, m, c)
	return CAConstraintFromID(rv)
}

// The attribute the constraint affects.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/attribute
func (c CAConstraint) Attribute() CAConstraintAttribute {
	rv := objc.Send[CAConstraintAttribute](c.ID, objc.Sel("attribute"))
	return CAConstraintAttribute(rv)
}

// Offset value of the constraint attribute.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/offset
func (c CAConstraint) Offset() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("offset"))
	return rv
}

// Scale factor of the constraint attribute.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/scale
func (c CAConstraint) Scale() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("scale"))
	return rv
}

// The constraint attribute of the layer the receiver is calculated relative
// to
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/sourceAttribute
func (c CAConstraint) SourceAttribute() CAConstraintAttribute {
	rv := objc.Send[CAConstraintAttribute](c.ID, objc.Sel("sourceAttribute"))
	return CAConstraintAttribute(rv)
}

// Name of the layer that the constraint is calculated relative to.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraint/sourceName
func (c CAConstraint) SourceName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sourceName"))
	return foundation.NSStringFromID(rv).String()
}
