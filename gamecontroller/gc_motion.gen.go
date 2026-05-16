// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCMotion] class.
var (
	_GCMotionClass     GCMotionClass
	_GCMotionClassOnce sync.Once
)

func getGCMotionClass() GCMotionClass {
	_GCMotionClassOnce.Do(func() {
		_GCMotionClass = GCMotionClass{class: objc.GetClass("GCMotion")}
	})
	return _GCMotionClass
}

// GetGCMotionClass returns the class object for GCMotion.
func GetGCMotionClass() GCMotionClass {
	return getGCMotionClass()
}

type GCMotionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCMotionClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCMotionClass) Alloc() GCMotion {
	rv := objc.Send[GCMotion](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that supports orientation and motion.
//
// # Overview
//
// The motion controller profile provides attitude and rotation data, as well
// as acceleration and sensor information. Use this profile to get motion
// input from a controller that measures acceleration and rotation rate. If
// the controller’s [GCMotion.Motion] property is a [GCMotion] object, the controller
// supports motion.
//
// This illustration shows the direction of the x, y, and z axes of an iPhone
// when held upright.
//
// [media-2930224]
//
// # Getting the Controller
//
//   - [GCMotion.Controller]: The controller for the profile.
//
// # Receiving a Callback When Input Values Change
//
//   - [GCMotion.ValueChangedHandler]: The block that the profile calls when an element’s value changes.
//   - [GCMotion.SetValueChangedHandler]
//
// # Verifying Capabilities
//
//   - [GCMotion.HasAttitude]: A Boolean value that indicates whether the controller provides attitude data.
//   - [GCMotion.HasRotationRate]: A Boolean value that indicates whether the controller provides rotation data.
//   - [GCMotion.HasGravityAndUserAcceleration]: A Boolean value that indicates whether the controller provides gravity and user acceleration data.
//
// # Accessing Attitude and Rotation Data
//
//   - [GCMotion.Attitude]: The attitude of the controller.
//   - [GCMotion.RotationRate]: The rotation rate of the controller.
//
// # Accessing Gravity and Acceleration Data
//
//   - [GCMotion.Acceleration]: The total acceleration of the controller that includes gravity and the acceleration the user applies to the controller.
//   - [GCMotion.Gravity]: The gravity acceleration vector from the controller’s reference frame.
//   - [GCMotion.UserAcceleration]: The acceleration that the user applies to the controller.
//
// # Accessing Sensor Data
//
//   - [GCMotion.SensorsRequireManualActivation]: A Boolean value that indicates whether the sensors that compute the motion data require manual activation.
//   - [GCMotion.SensorsActive]: A Boolean value that indicates whether the sensors that compute the motion data are active.
//   - [GCMotion.SetSensorsActive]
//
// # Setting Snapshot Values
//
//   - [GCMotion.SetStateFromMotion]: Copies the input values from a specified motion profile to a snapshot of a motion profile.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion
type GCMotion struct {
	objectivec.Object
}

// GCMotionFromID constructs a [GCMotion] from an objc.ID.
//
// A controller profile that supports orientation and motion.
func GCMotionFromID(id objc.ID) GCMotion {
	return GCMotion{objectivec.Object{ID: id}}
}

// NOTE: GCMotion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCMotion] class.
//
// # Getting the Controller
//
//   - [IGCMotion.Controller]: The controller for the profile.
//
// # Receiving a Callback When Input Values Change
//
//   - [IGCMotion.ValueChangedHandler]: The block that the profile calls when an element’s value changes.
//   - [IGCMotion.SetValueChangedHandler]
//
// # Verifying Capabilities
//
//   - [IGCMotion.HasAttitude]: A Boolean value that indicates whether the controller provides attitude data.
//   - [IGCMotion.HasRotationRate]: A Boolean value that indicates whether the controller provides rotation data.
//   - [IGCMotion.HasGravityAndUserAcceleration]: A Boolean value that indicates whether the controller provides gravity and user acceleration data.
//
// # Accessing Attitude and Rotation Data
//
//   - [IGCMotion.Attitude]: The attitude of the controller.
//   - [IGCMotion.RotationRate]: The rotation rate of the controller.
//
// # Accessing Gravity and Acceleration Data
//
//   - [IGCMotion.Acceleration]: The total acceleration of the controller that includes gravity and the acceleration the user applies to the controller.
//   - [IGCMotion.Gravity]: The gravity acceleration vector from the controller’s reference frame.
//   - [IGCMotion.UserAcceleration]: The acceleration that the user applies to the controller.
//
// # Accessing Sensor Data
//
//   - [IGCMotion.SensorsRequireManualActivation]: A Boolean value that indicates whether the sensors that compute the motion data require manual activation.
//   - [IGCMotion.SensorsActive]: A Boolean value that indicates whether the sensors that compute the motion data are active.
//   - [IGCMotion.SetSensorsActive]
//
// # Setting Snapshot Values
//
//   - [IGCMotion.SetStateFromMotion]: Copies the input values from a specified motion profile to a snapshot of a motion profile.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion
type IGCMotion interface {
	objectivec.IObject

	// Topic: Getting the Controller

	// The controller for the profile.
	Controller() IGCController

	// Topic: Receiving a Callback When Input Values Change

	// The block that the profile calls when an element’s value changes.
	ValueChangedHandler() GCMotionValueChangedHandler
	SetValueChangedHandler(value GCMotionValueChangedHandler)

	// Topic: Verifying Capabilities

	// A Boolean value that indicates whether the controller provides attitude data.
	HasAttitude() bool
	// A Boolean value that indicates whether the controller provides rotation data.
	HasRotationRate() bool
	// A Boolean value that indicates whether the controller provides gravity and user acceleration data.
	HasGravityAndUserAcceleration() bool

	// Topic: Accessing Attitude and Rotation Data

	// The attitude of the controller.
	Attitude() GCQuaternion
	// The rotation rate of the controller.
	RotationRate() GCRotationRate

	// Topic: Accessing Gravity and Acceleration Data

	// The total acceleration of the controller that includes gravity and the acceleration the user applies to the controller.
	Acceleration() GCAcceleration
	// The gravity acceleration vector from the controller’s reference frame.
	Gravity() GCAcceleration
	// The acceleration that the user applies to the controller.
	UserAcceleration() GCAcceleration

	// Topic: Accessing Sensor Data

	// A Boolean value that indicates whether the sensors that compute the motion data require manual activation.
	SensorsRequireManualActivation() bool
	// A Boolean value that indicates whether the sensors that compute the motion data are active.
	SensorsActive() bool
	SetSensorsActive(value bool)

	// Topic: Setting Snapshot Values

	// Copies the input values from a specified motion profile to a snapshot of a motion profile.
	SetStateFromMotion(motion IGCMotion)

	// The motion input profile.
	Motion() IGCMotion
	SetMotion(value IGCMotion)
}

// Init initializes the instance.
func (g GCMotion) Init() GCMotion {
	rv := objc.Send[GCMotion](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCMotion) Autorelease() GCMotion {
	rv := objc.Send[GCMotion](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMotion creates a new GCMotion instance.
func NewGCMotion() GCMotion {
	class := getGCMotionClass()
	rv := objc.Send[GCMotion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Copies the input values from a specified motion profile to a snapshot of a
// motion profile.
//
// motion: The motion profile to copy the input values from.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/setStateFrom(_:)
func (g GCMotion) SetStateFromMotion(motion IGCMotion) {
	objc.Send[objc.ID](g.ID, objc.Sel("setStateFromMotion:"), motion)
}

// The controller for the profile.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/controller
func (g GCMotion) Controller() IGCController {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("controller"))
	return GCControllerFromID(objc.ID(rv))
}

// The block that the profile calls when an element’s value changes.
//
// # Discussion
//
// If multiple elements change values at the same time, the profile calls this
// block once for each element that changes. If the value of a subelement
// changes, the profile only calls the block for the containing element.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/valueChangedHandler
func (g GCMotion) ValueChangedHandler() GCMotionValueChangedHandler {
	rv := objc.Send[GCMotionValueChangedHandler](g.ID, objc.Sel("valueChangedHandler"))
	return GCMotionValueChangedHandler(rv)
}
func (g GCMotion) SetValueChangedHandler(value GCMotionValueChangedHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setValueChangedHandler:"), value)
}

// A Boolean value that indicates whether the controller provides attitude
// data.
//
// # Discussion
//
// true if the controller provides attitude data; otherwise, false.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/hasAttitude
func (g GCMotion) HasAttitude() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasAttitude"))
	return rv
}

// A Boolean value that indicates whether the controller provides rotation
// data.
//
// # Discussion
//
// true if the controller provides rotation data; otherwise, false.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/hasRotationRate
func (g GCMotion) HasRotationRate() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasRotationRate"))
	return rv
}

// A Boolean value that indicates whether the controller provides gravity and
// user acceleration data.
//
// # Discussion
//
// true if the controller provides both gravity and user acceleration data;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/hasGravityAndUserAcceleration
func (g GCMotion) HasGravityAndUserAcceleration() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasGravityAndUserAcceleration"))
	return rv
}

// The attitude of the controller.
//
// # Discussion
//
// The is the orientation of a body relative to the controller’s reference
// frame.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/attitude
func (g GCMotion) Attitude() GCQuaternion {
	rv := objc.Send[GCQuaternion](g.ID, objc.Sel("attitude"))
	return GCQuaternion(rv)
}

// The rotation rate of the controller.
//
// # Discussion
//
// The is a gyroscopic measurement of the controller’s rotation around the
// x, y, and z axes.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/rotationRate
func (g GCMotion) RotationRate() GCRotationRate {
	rv := objc.Send[GCRotationRate](g.ID, objc.Sel("rotationRate"))
	return GCRotationRate(rv)
}

// The total acceleration of the controller that includes gravity and the
// acceleration the user applies to the controller.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/acceleration
func (g GCMotion) Acceleration() GCAcceleration {
	rv := objc.Send[GCAcceleration](g.ID, objc.Sel("acceleration"))
	return GCAcceleration(rv)
}

// The gravity acceleration vector from the controller’s reference frame.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/gravity
func (g GCMotion) Gravity() GCAcceleration {
	rv := objc.Send[GCAcceleration](g.ID, objc.Sel("gravity"))
	return GCAcceleration(rv)
}

// The acceleration that the user applies to the controller.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/userAcceleration
func (g GCMotion) UserAcceleration() GCAcceleration {
	rv := objc.Send[GCAcceleration](g.ID, objc.Sel("userAcceleration"))
	return GCAcceleration(rv)
}

// A Boolean value that indicates whether the sensors that compute the motion
// data require manual activation.
//
// # Discussion
//
// true if the sensors require manual activation; otherwise, false.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/sensorsRequireManualActivation
func (g GCMotion) SensorsRequireManualActivation() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("sensorsRequireManualActivation"))
	return rv
}

// A Boolean value that indicates whether the sensors that compute the motion
// data are active.
//
// # Discussion
//
// true if the sensors are active; otherwise, false.
//
// See: https://developer.apple.com/documentation/GameController/GCMotion/sensorsActive
func (g GCMotion) SensorsActive() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("sensorsActive"))
	return rv
}
func (g GCMotion) SetSensorsActive(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setSensorsActive:"), value)
}

// The motion input profile.
//
// See: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g GCMotion) Motion() IGCMotion {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("motion"))
	return GCMotionFromID(objc.ID(rv))
}
func (g GCMotion) SetMotion(value IGCMotion) {
	objc.Send[struct{}](g.ID, objc.Sel("setMotion:"), value)
}
