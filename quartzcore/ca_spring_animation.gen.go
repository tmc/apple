// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CASpringAnimation] class.
var (
	_CASpringAnimationClass     CASpringAnimationClass
	_CASpringAnimationClassOnce sync.Once
)

func getCASpringAnimationClass() CASpringAnimationClass {
	_CASpringAnimationClassOnce.Do(func() {
		_CASpringAnimationClass = CASpringAnimationClass{class: objc.GetClass("CASpringAnimation")}
	})
	return _CASpringAnimationClass
}

// GetCASpringAnimationClass returns the class object for CASpringAnimation.
func GetCASpringAnimationClass() CASpringAnimationClass {
	return getCASpringAnimationClass()
}

type CASpringAnimationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CASpringAnimationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CASpringAnimationClass) Alloc() CASpringAnimation {
	rv := objc.Send[CASpringAnimation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An animation that applies a spring-like force to a layer’s properties.
//
// # Overview
//
// You would typically use a spring animation to animate a layer’s position
// so that it appears to be pulled towards a target by a spring. The further
// the layer is from the target, the greater the acceleration towards it is.
//
// [CASpringAnimation] allows control over physically based attributes such as
// the spring’s damping and stiffness.
//
// You can use a spring animation to animation properties of a layer other
// than its position. The following code shows how to create a spring
// animation that bounces a layer into view by animating its scale from `0` to
// `1`. Because the spring animation can overshoot its [CASpringAnimation.ToValue], the animated
// layer may exceed its frame.
//
// # Configuring Physical Attributes
//
//   - [CASpringAnimation.Damping]: Defines how the spring’s motion should be damped due to the forces of friction.
//   - [CASpringAnimation.SetDamping]
//   - [CASpringAnimation.InitialVelocity]: The initial velocity of the object attached to the spring.
//   - [CASpringAnimation.SetInitialVelocity]
//   - [CASpringAnimation.Mass]: The mass of the object attached to the end of the spring.
//   - [CASpringAnimation.SetMass]
//   - [CASpringAnimation.SettlingDuration]: The estimated duration required for the spring system to be considered at rest.
//   - [CASpringAnimation.Stiffness]: The spring stiffness coefficient.
//   - [CASpringAnimation.SetStiffness]
//
// # Initializers
//
//   - [CASpringAnimation.InitWithPerceptualDurationBounce]
//
// # Instance Properties
//
//   - [CASpringAnimation.AllowsOverdamping]
//   - [CASpringAnimation.SetAllowsOverdamping]
//   - [CASpringAnimation.Bounce]
//   - [CASpringAnimation.PerceptualDuration]
//
// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation
type CASpringAnimation struct {
	CABasicAnimation
}

// CASpringAnimationFromID constructs a [CASpringAnimation] from an objc.ID.
//
// An animation that applies a spring-like force to a layer’s properties.
func CASpringAnimationFromID(id objc.ID) CASpringAnimation {
	return CASpringAnimation{CABasicAnimation: CABasicAnimationFromID(id)}
}

// NOTE: CASpringAnimation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CASpringAnimation] class.
//
// # Configuring Physical Attributes
//
//   - [ICASpringAnimation.Damping]: Defines how the spring’s motion should be damped due to the forces of friction.
//   - [ICASpringAnimation.SetDamping]
//   - [ICASpringAnimation.InitialVelocity]: The initial velocity of the object attached to the spring.
//   - [ICASpringAnimation.SetInitialVelocity]
//   - [ICASpringAnimation.Mass]: The mass of the object attached to the end of the spring.
//   - [ICASpringAnimation.SetMass]
//   - [ICASpringAnimation.SettlingDuration]: The estimated duration required for the spring system to be considered at rest.
//   - [ICASpringAnimation.Stiffness]: The spring stiffness coefficient.
//   - [ICASpringAnimation.SetStiffness]
//
// # Initializers
//
//   - [ICASpringAnimation.InitWithPerceptualDurationBounce]
//
// # Instance Properties
//
//   - [ICASpringAnimation.AllowsOverdamping]
//   - [ICASpringAnimation.SetAllowsOverdamping]
//   - [ICASpringAnimation.Bounce]
//   - [ICASpringAnimation.PerceptualDuration]
//
// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation
type ICASpringAnimation interface {
	ICABasicAnimation

	// Topic: Configuring Physical Attributes

	// Defines how the spring’s motion should be damped due to the forces of friction.
	Damping() float64
	SetDamping(value float64)
	// The initial velocity of the object attached to the spring.
	InitialVelocity() float64
	SetInitialVelocity(value float64)
	// The mass of the object attached to the end of the spring.
	Mass() float64
	SetMass(value float64)
	// The estimated duration required for the spring system to be considered at rest.
	SettlingDuration() float64
	// The spring stiffness coefficient.
	Stiffness() float64
	SetStiffness(value float64)

	// Topic: Initializers

	InitWithPerceptualDurationBounce(perceptualDuration float64, bounce float64) CASpringAnimation

	// Topic: Instance Properties

	AllowsOverdamping() bool
	SetAllowsOverdamping(value bool)
	Bounce() float64
	PerceptualDuration() float64
}

// Init initializes the instance.
func (s CASpringAnimation) Init() CASpringAnimation {
	rv := objc.Send[CASpringAnimation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s CASpringAnimation) Autorelease() CASpringAnimation {
	rv := objc.Send[CASpringAnimation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewCASpringAnimation creates a new CASpringAnimation instance.
func NewCASpringAnimation() CASpringAnimation {
	class := getCASpringAnimationClass()
	rv := objc.Send[CASpringAnimation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns an [CAPropertyAnimation] instance for the specified key
// path.
//
// path: The key path of the property to be animated.
//
// # Return Value
//
// A new instance of [CAPropertyAnimation] with the key path set to `keyPath`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func NewSpringAnimationWithKeyPath(path string) CASpringAnimation {
	rv := objc.Send[objc.ID](objc.ID(getCASpringAnimationClass().class), objc.Sel("animationWithKeyPath:"), objc.String(path))
	return CASpringAnimationFromID(rv)
}

// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/init(perceptualDuration:bounce:)
func NewSpringAnimationWithPerceptualDurationBounce(perceptualDuration float64, bounce float64) CASpringAnimation {
	instance := getCASpringAnimationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPerceptualDuration:bounce:"), perceptualDuration, bounce)
	return CASpringAnimationFromID(rv)
}

// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/init(perceptualDuration:bounce:)
func (s CASpringAnimation) InitWithPerceptualDurationBounce(perceptualDuration float64, bounce float64) CASpringAnimation {
	rv := objc.Send[CASpringAnimation](s.ID, objc.Sel("initWithPerceptualDuration:bounce:"), perceptualDuration, bounce)
	return rv
}

// Defines how the spring’s motion should be damped due to the forces of
// friction.
//
// # Discussion
//
// The default value of the [Damping] property is `10`. Reducing this value
// reduces the energy loss with each oscillation: the animated value will
// overshoot the [ToValue] and the [SettlingDuration] may be greater than the
// [Duration]. Increasing the value increases the energy loss with each
// duration: there will be fewer and smaller oscillations and the
// [SettlingDuration] may be smaller than the duration.
//
// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/damping
func (s CASpringAnimation) Damping() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("damping"))
	return rv
}
func (s CASpringAnimation) SetDamping(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDamping:"), value)
}

// The initial velocity of the object attached to the spring.
//
// # Discussion
//
// Defaults to `0`, which represents an unmoving object. Negative values
// represent the object moving away from the spring attachment point, positive
// values represent the object moving towards the spring attachment point.
//
// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/initialVelocity
func (s CASpringAnimation) InitialVelocity() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("initialVelocity"))
	return rv
}
func (s CASpringAnimation) SetInitialVelocity(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setInitialVelocity:"), value)
}

// The mass of the object attached to the end of the spring.
//
// # Discussion
//
// The default mass is `1`. Increasing this value will increase the spring
// effect: the attached object will be subject to more oscillations and
// greater overshoot, resulting in an increased [SettlingDuration]. Decreasing
// the mass will reduce the spring effect: there will be fewer oscillations
// and a reduced overshoot, resulting in a decreased [SettlingDuration].
//
// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/mass
func (s CASpringAnimation) Mass() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("mass"))
	return rv
}
func (s CASpringAnimation) SetMass(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMass:"), value)
}

// The estimated duration required for the spring system to be considered at
// rest.
//
// # Discussion
//
// The duration is evaluated for the current animation parameters and may not
// the same as the [Duration].
//
// The following code creates a spring animation with a [Duration] of 2
// seconds.
//
// With a damping coefficient of `5`, the settling duration is approximately
// 2.85 seconds: the animated layer bounces around the target position several
// times before settling. However, changing the [Damping] property to `15`
// reduces the settling duration to just over 1 second: the animated layer
// quickly comes to a stop as it reaches the target position.
//
// All of the spring animation’s physical attributes: [Damping],
// [InitialVelocity], [Mass] and [Stiffness], can affect the settling
// duration.
//
// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/settlingDuration
func (s CASpringAnimation) SettlingDuration() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("settlingDuration"))
	return rv
}

// The spring stiffness coefficient.
//
// # Discussion
//
// The default stiffness coefficient is `100`. Increasing the [Stiffness]
// reduces the number of oscillations and will reduce the settling duration.
// Decreasing the [Stiffness] increases the the number of oscillations and
// will increase the settling duration.
//
// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/stiffness
func (s CASpringAnimation) Stiffness() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("stiffness"))
	return rv
}
func (s CASpringAnimation) SetStiffness(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setStiffness:"), value)
}

// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/allowsOverdamping
func (s CASpringAnimation) AllowsOverdamping() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("allowsOverdamping"))
	return rv
}
func (s CASpringAnimation) SetAllowsOverdamping(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowsOverdamping:"), value)
}

// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/bounce
func (s CASpringAnimation) Bounce() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("bounce"))
	return rv
}

// See: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/perceptualDuration
func (s CASpringAnimation) PerceptualDuration() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("perceptualDuration"))
	return rv
}
