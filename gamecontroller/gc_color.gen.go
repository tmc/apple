// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCColor] class.
var (
	_GCColorClass     GCColorClass
	_GCColorClassOnce sync.Once
)

func getGCColorClass() GCColorClass {
	_GCColorClassOnce.Do(func() {
		_GCColorClass = GCColorClass{class: objc.GetClass("GCColor")}
	})
	return _GCColorClass
}

// GetGCColorClass returns the class object for GCColor.
func GetGCColorClass() GCColorClass {
	return getGCColorClass()
}

type GCColorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCColorClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCColorClass) Alloc() GCColor {
	rv := objc.Send[GCColor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The color of a device light.
//
// # Creating colors
//
//   - [GCColor.InitWithRedGreenBlue]: Creates a color with the specified red, green, and blue values.
//
// # Setting color values
//
//   - [GCColor.Red]: The normalized value of the red component ranging from 0 to 1.
//   - [GCColor.Green]: The normalized value of the green component ranging from 0 to 1.
//   - [GCColor.Blue]: The normalized value of the blue component ranging from 0 to 1.
//
// # Initializers
//
//   - [GCColor.InitWithCoder]
//
// See: https://developer.apple.com/documentation/GameController/GCColor
type GCColor struct {
	objectivec.Object
}

// GCColorFromID constructs a [GCColor] from an objc.ID.
//
// The color of a device light.
func GCColorFromID(id objc.ID) GCColor {
	return GCColor{objectivec.Object{ID: id}}
}

// NOTE: GCColor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCColor] class.
//
// # Creating colors
//
//   - [IGCColor.InitWithRedGreenBlue]: Creates a color with the specified red, green, and blue values.
//
// # Setting color values
//
//   - [IGCColor.Red]: The normalized value of the red component ranging from 0 to 1.
//   - [IGCColor.Green]: The normalized value of the green component ranging from 0 to 1.
//   - [IGCColor.Blue]: The normalized value of the blue component ranging from 0 to 1.
//
// # Initializers
//
//   - [IGCColor.InitWithCoder]
//
// See: https://developer.apple.com/documentation/GameController/GCColor
type IGCColor interface {
	objectivec.IObject

	// Topic: Creating colors

	// Creates a color with the specified red, green, and blue values.
	InitWithRedGreenBlue(red float32, green float32, blue float32) GCColor

	// Topic: Setting color values

	// The normalized value of the red component ranging from 0 to 1.
	Red() float32
	// The normalized value of the green component ranging from 0 to 1.
	Green() float32
	// The normalized value of the blue component ranging from 0 to 1.
	Blue() float32

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) GCColor

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (g GCColor) Init() GCColor {
	rv := objc.Send[GCColor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCColor) Autorelease() GCColor {
	rv := objc.Send[GCColor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCColor creates a new GCColor instance.
func NewGCColor() GCColor {
	class := getGCColorClass()
	rv := objc.Send[GCColor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GameController/GCColor/init(coder:)
func NewGCColorWithCoder(coder foundation.INSCoder) GCColor {
	instance := getGCColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GCColorFromID(rv)
}

// Creates a color with the specified red, green, and blue values.
//
// red: The normalized value of the red component ranging from 0 to 1.
//
// green: The normalized value of the green component ranging from 0 to 1.
//
// blue: The normalized value of the blue component ranging from 0 to 1.
//
// See: https://developer.apple.com/documentation/GameController/GCColor/init(red:green:blue:)
func NewGCColorWithRedGreenBlue(red float32, green float32, blue float32) GCColor {
	instance := getGCColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	return GCColorFromID(rv)
}

// Creates a color with the specified red, green, and blue values.
//
// red: The normalized value of the red component ranging from 0 to 1.
//
// green: The normalized value of the green component ranging from 0 to 1.
//
// blue: The normalized value of the blue component ranging from 0 to 1.
//
// See: https://developer.apple.com/documentation/GameController/GCColor/init(red:green:blue:)
func (g GCColor) InitWithRedGreenBlue(red float32, green float32, blue float32) GCColor {
	rv := objc.Send[GCColor](g.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	return rv
}

// See: https://developer.apple.com/documentation/GameController/GCColor/init(coder:)
func (g GCColor) InitWithCoder(coder foundation.INSCoder) GCColor {
	rv := objc.Send[GCColor](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (g GCColor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The normalized value of the red component ranging from 0 to 1.
//
// See: https://developer.apple.com/documentation/GameController/GCColor/red
func (g GCColor) Red() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("red"))
	return rv
}

// The normalized value of the green component ranging from 0 to 1.
//
// See: https://developer.apple.com/documentation/GameController/GCColor/green
func (g GCColor) Green() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("green"))
	return rv
}

// The normalized value of the blue component ranging from 0 to 1.
//
// See: https://developer.apple.com/documentation/GameController/GCColor/blue
func (g GCColor) Blue() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("blue"))
	return rv
}
