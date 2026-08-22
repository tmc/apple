// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PGDisplayMode] class.
var (
	_PGDisplayModeClass     PGDisplayModeClass
	_PGDisplayModeClassOnce sync.Once
)

func getPGDisplayModeClass() PGDisplayModeClass {
	_PGDisplayModeClassOnce.Do(func() {
		_PGDisplayModeClass = PGDisplayModeClass{class: objc.GetClass("PGDisplayMode")}
	})
	return _PGDisplayModeClass
}

// GetPGDisplayModeClass returns the class object for PGDisplayMode.
func GetPGDisplayModeClass() PGDisplayModeClass {
	return getPGDisplayModeClass()
}

type PGDisplayModeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PGDisplayModeClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PGDisplayModeClass) Alloc() PGDisplayMode {
	rv := objc.Send[PGDisplayMode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A description of a supported display mode.
//
// # Creating a Display Mode
//
//   - [PGDisplayMode.InitWithSizeInPixelsRefreshRateInHz]: Creates a new display mode.
//
// # Inspecting Mode Properties
//
//   - [PGDisplayMode.SizeInPixels]: The display mode’s dimensions in pixels.
//   - [PGDisplayMode.RefreshRate]: The mode’s refresh rate.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode
type PGDisplayMode struct {
	objectivec.Object
}

// PGDisplayModeFromID constructs a [PGDisplayMode] from an objc.ID.
//
// A description of a supported display mode.
func PGDisplayModeFromID(id objc.ID) PGDisplayMode {
	return PGDisplayMode{objectivec.Object{ID: id}}
}

// NOTE: PGDisplayMode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PGDisplayMode] class.
//
// # Creating a Display Mode
//
//   - [IPGDisplayMode.InitWithSizeInPixelsRefreshRateInHz]: Creates a new display mode.
//
// # Inspecting Mode Properties
//
//   - [IPGDisplayMode.SizeInPixels]: The display mode’s dimensions in pixels.
//   - [IPGDisplayMode.RefreshRate]: The mode’s refresh rate.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode
type IPGDisplayMode interface {
	objectivec.IObject

	// Topic: Creating a Display Mode

	// Creates a new display mode.
	InitWithSizeInPixelsRefreshRateInHz(sizeInPixels PGDisplayCoord_t, refreshRateInHz float64) PGDisplayMode

	// Topic: Inspecting Mode Properties

	// The display mode’s dimensions in pixels.
	SizeInPixels() PGDisplayCoord_t
	// The mode’s refresh rate.
	RefreshRate() float64
}

// Init initializes the instance.
func (p PGDisplayMode) Init() PGDisplayMode {
	rv := objc.Send[PGDisplayMode](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PGDisplayMode) Autorelease() PGDisplayMode {
	rv := objc.Send[PGDisplayMode](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPGDisplayMode creates a new PGDisplayMode instance.
func NewPGDisplayMode() PGDisplayMode {
	class := getPGDisplayModeClass()
	rv := objc.Send[PGDisplayMode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new display mode.
//
// sizeInPixels: The display mode’s dimensions in pixels.
//
// refreshRateInHz: The mode’s refresh rate.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode/init(sizeInPixels:refreshRateInHz:)
func NewPGDisplayModeWithSizeInPixelsRefreshRateInHz(sizeInPixels PGDisplayCoord_t, refreshRateInHz float64) PGDisplayMode {
	instance := getPGDisplayModeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSizeInPixels:refreshRateInHz:"), sizeInPixels, refreshRateInHz)
	return PGDisplayModeFromID(rv)
}

// Creates a new display mode.
//
// sizeInPixels: The display mode’s dimensions in pixels.
//
// refreshRateInHz: The mode’s refresh rate.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode/init(sizeInPixels:refreshRateInHz:)
func (p PGDisplayMode) InitWithSizeInPixelsRefreshRateInHz(sizeInPixels PGDisplayCoord_t, refreshRateInHz float64) PGDisplayMode {
	rv := objc.Send[PGDisplayMode](p.ID, objc.Sel("initWithSizeInPixels:refreshRateInHz:"), sizeInPixels, refreshRateInHz)
	return rv
}

// The display mode’s dimensions in pixels.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode/sizeInPixels
func (p PGDisplayMode) SizeInPixels() PGDisplayCoord_t {
	rv := objc.Send[PGDisplayCoord_t](p.ID, objc.Sel("sizeInPixels"))
	return PGDisplayCoord_t(rv)
}

// The mode’s refresh rate.
//
// # Discussion
//
// Consider supplying only modes that have a refresh rate equal to that of the
// host environment’s physical display.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayMode/refreshRate
func (p PGDisplayMode) RefreshRate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("refreshRate"))
	return rv
}
