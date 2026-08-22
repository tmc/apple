// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXBrailleMap] class.
var (
	_AXBrailleMapClass     AXBrailleMapClass
	_AXBrailleMapClassOnce sync.Once
)

func getAXBrailleMapClass() AXBrailleMapClass {
	_AXBrailleMapClassOnce.Do(func() {
		_AXBrailleMapClass = AXBrailleMapClass{class: objc.GetClass("AXBrailleMap")}
	})
	return _AXBrailleMapClass
}

// GetAXBrailleMapClass returns the class object for AXBrailleMap.
func GetAXBrailleMapClass() AXBrailleMapClass {
	return getAXBrailleMapClass()
}

type AXBrailleMapClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXBrailleMapClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXBrailleMapClass) Alloc() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A representation of a two-dimensional braille display.
//
// # Overview
//
// A braille map object represents a two-dimensional braille display that’s
// connected to the current Apple device. By specifying the dot patterns in
// the braille map, you can change the content the user experiences. To render
// the data from the braille map to the display, implement
// [AXBrailleMapRenderer].
//
// # Creating a braille map
//
//   - [AXBrailleMap.InitWithCoder]
//
// # Getting display dimensions
//
//   - [AXBrailleMap.Dimensions]: The number of pins in each dimension of the braille display.
//
// # Accessing dots
//
//   - [AXBrailleMap.SetHeightAtPoint]: Sets the height of an individual pin on the braille display.
//   - [AXBrailleMap.HeightAtPoint]: Retrieves the height of an individual pin on the braille display.
//
// # Displaying images
//
//   - [AXBrailleMap.PresentImage]: Converts the data from the image you specify into the braille map.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap
type AXBrailleMap struct {
	objectivec.Object
}

// AXBrailleMapFromID constructs a [AXBrailleMap] from an objc.ID.
//
// A representation of a two-dimensional braille display.
func AXBrailleMapFromID(id objc.ID) AXBrailleMap {
	return AXBrailleMap{objectivec.Object{ID: id}}
}

// NOTE: AXBrailleMap adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXBrailleMap] class.
//
// # Creating a braille map
//
//   - [IAXBrailleMap.InitWithCoder]
//
// # Getting display dimensions
//
//   - [IAXBrailleMap.Dimensions]: The number of pins in each dimension of the braille display.
//
// # Accessing dots
//
//   - [IAXBrailleMap.SetHeightAtPoint]: Sets the height of an individual pin on the braille display.
//   - [IAXBrailleMap.HeightAtPoint]: Retrieves the height of an individual pin on the braille display.
//
// # Displaying images
//
//   - [IAXBrailleMap.PresentImage]: Converts the data from the image you specify into the braille map.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap
type IAXBrailleMap interface {
	objectivec.IObject

	// Topic: Creating a braille map

	InitWithCoder(coder foundation.INSCoder) AXBrailleMap

	// Topic: Getting display dimensions

	// The number of pins in each dimension of the braille display.
	Dimensions() corefoundation.CGSize

	// Topic: Accessing dots

	// Sets the height of an individual pin on the braille display.
	SetHeightAtPoint(status float32, point corefoundation.CGPoint)
	// Retrieves the height of an individual pin on the braille display.
	HeightAtPoint(point corefoundation.CGPoint) float32

	// Topic: Displaying images

	// Converts the data from the image you specify into the braille map.
	PresentImage(image coregraphics.CGImageRef)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AXBrailleMap) Init() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXBrailleMap) Autorelease() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleMap creates a new AXBrailleMap instance.
func NewAXBrailleMap() AXBrailleMap {
	class := getAXBrailleMapClass()
	rv := objc.Send[AXBrailleMap](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/init(coder:)
func NewAXBrailleMapWithCoder(coder foundation.INSCoder) AXBrailleMap {
	instance := getAXBrailleMapClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AXBrailleMapFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/init(coder:)
func (a AXBrailleMap) InitWithCoder(coder foundation.INSCoder) AXBrailleMap {
	rv := objc.Send[AXBrailleMap](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Sets the height of an individual pin on the braille display.
//
// status: A floating-point number between `0.0` and `1.0` that specifies the height
// of the pin. A value of `0.0` lowers the pin completely, and a value of
// `1.0` raises the pin completely.
//
// point: The location of the pin to adjust the height for. The bottom-left of the
// display is at `{ 0,0 }`, and the top-right of the display is at `{
// dimensions.Width() - 1, dimensions.Height() - 1}`.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/setHeight(_:at:)
func (a AXBrailleMap) SetHeightAtPoint(status float32, point corefoundation.CGPoint) {
	objc.Send[objc.ID](a.ID, objc.Sel("setHeight:atPoint:"), status, point)
}

// Retrieves the height of an individual pin on the braille display.
//
// point: The location of the pin to retrieve the height for. The bottom-left of the
// display is at `{ 0,0 }`, and the top-right of the display is at `{
// dimensions.Width() - 1, dimensions.Height() - 1}`.
//
// # Return Value
//
// A floating-point number between `0.0` and `1.0` that specifies the height
// of the pin. A value of `0.0` indicates that the pin is completely lowered,
// and a value of `1.0` indicates that the pin is completely raised.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/height(at:)
func (a AXBrailleMap) HeightAtPoint(point corefoundation.CGPoint) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("heightAtPoint:"), point)
	return rv
}

// Converts the data from the image you specify into the braille map.
//
// image: An image to convert into the braille map.
//
// # Discussion
//
// Use this method to convert image data into the braille map directly,
// without the need to modify the heights of individual pins using
// [AXBrailleMap.SetHeightAtPoint].
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/present(_:)
func (a AXBrailleMap) PresentImage(image coregraphics.CGImageRef) {
	objc.Send[objc.ID](a.ID, objc.Sel("presentImage:"), image)
}
func (a AXBrailleMap) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The number of pins in each dimension of the braille display.
//
// # Discussion
//
// The dimensions can change if the user zooms in on the content.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/dimensions
func (a AXBrailleMap) Dimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("dimensions"))
	return corefoundation.CGSize(rv)
}
