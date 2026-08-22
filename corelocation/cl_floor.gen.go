// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLFloor] class.
var (
	_CLFloorClass     CLFloorClass
	_CLFloorClassOnce sync.Once
)

func getCLFloorClass() CLFloorClass {
	_CLFloorClassOnce.Do(func() {
		_CLFloorClass = CLFloorClass{class: objc.GetClass("CLFloor")}
	})
	return _CLFloorClass
}

// GetCLFloorClass returns the class object for CLFloor.
func GetCLFloorClass() CLFloorClass {
	return getCLFloorClass()
}

type CLFloorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLFloorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLFloorClass) Alloc() CLFloor {
	rv := objc.Send[CLFloor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The floor of a building on which the user’s device is located.
//
// # Overview
//
// A [CLFloor] object specifies the floor of the building on which the device
// is located. In places where floor information can be determined, a
// [CLLocation] object may include a floor object along with the regular
// location data.
//
// You do not create instances of this class directly, nor should you subclass
// it.
//
// # Getting the floor level
//
//   - [CLFloor.Level]: The logical floor of the building.
//
// # Initializers
//
//   - [CLFloor.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLFloor
type CLFloor struct {
	objectivec.Object
}

// CLFloorFromID constructs a [CLFloor] from an objc.ID.
//
// The floor of a building on which the user’s device is located.
func CLFloorFromID(id objc.ID) CLFloor {
	return CLFloor{objectivec.Object{ID: id}}
}

// NOTE: CLFloor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLFloor] class.
//
// # Getting the floor level
//
//   - [ICLFloor.Level]: The logical floor of the building.
//
// # Initializers
//
//   - [ICLFloor.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLFloor
type ICLFloor interface {
	objectivec.IObject

	// Topic: Getting the floor level

	// The logical floor of the building.
	Level() int

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CLFloor

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f CLFloor) Init() CLFloor {
	rv := objc.Send[CLFloor](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f CLFloor) Autorelease() CLFloor {
	rv := objc.Send[CLFloor](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLFloor creates a new CLFloor instance.
func NewCLFloor() CLFloor {
	class := getCLFloorClass()
	rv := objc.Send[CLFloor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLFloor/init(coder:)
func NewFloorWithCoder(coder foundation.INSCoder) CLFloor {
	instance := getCLFloorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLFloorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreLocation/CLFloor/init(coder:)
func (f CLFloor) InitWithCoder(coder foundation.INSCoder) CLFloor {
	rv := objc.Send[CLFloor](f.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (f CLFloor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The logical floor of the building.
//
// # Discussion
//
// Level values represent logical levels above or below ground level and are
// not intended to correspond to any numbering scheme in use by the building
// itself. The ground floor of a building is always represented by the value
// `0`. Floors above the ground floor are represented by positive integers, so
// a value of `1` represents the floor above ground level, a value of `2`
// represents two floors above ground level, and so on. Floors below the
// ground floor are represented by corresponding negative integers, with a
// value of -1 representing the floor immediately below ground level and so
// on.
//
// It is erroneous to use the user’s level in a building as an estimate of
// altitude.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLFloor/level
func (f CLFloor) Level() int {
	rv := objc.Send[int](f.ID, objc.Sel("level"))
	return rv
}
