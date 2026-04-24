// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGSystemStatusIndicatorsShim] class.
var (
	_PKGSystemStatusIndicatorsShimClass     PKGSystemStatusIndicatorsShimClass
	_PKGSystemStatusIndicatorsShimClassOnce sync.Once
)

func getPKGSystemStatusIndicatorsShimClass() PKGSystemStatusIndicatorsShimClass {
	_PKGSystemStatusIndicatorsShimClassOnce.Do(func() {
		_PKGSystemStatusIndicatorsShimClass = PKGSystemStatusIndicatorsShimClass{class: objc.GetClass("PKGSystemStatusIndicatorsShim")}
	})
	return _PKGSystemStatusIndicatorsShimClass
}

// GetPKGSystemStatusIndicatorsShimClass returns the class object for PKGSystemStatusIndicatorsShim.
func GetPKGSystemStatusIndicatorsShimClass() PKGSystemStatusIndicatorsShimClass {
	return getPKGSystemStatusIndicatorsShimClass()
}

type PKGSystemStatusIndicatorsShimClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGSystemStatusIndicatorsShimClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGSystemStatusIndicatorsShimClass) Alloc() PKGSystemStatusIndicatorsShim {
	rv := objc.Send[PKGSystemStatusIndicatorsShim](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGSystemStatusIndicatorsShim.DisplayNeedsSecureIndicatorFlush]
//   - [PKGSystemStatusIndicatorsShim.GlobalEnableSecureIndicators]
//   - [PKGSystemStatusIndicatorsShim.SetGlobalEnableSecureIndicators]
//   - [PKGSystemStatusIndicatorsShim.HasIndicatorWindows]
//   - [PKGSystemStatusIndicatorsShim.OverrideIndicatorSize]
//   - [PKGSystemStatusIndicatorsShim.RebuildDisplays]
//   - [PKGSystemStatusIndicatorsShim.RebuildLayers]
//   - [PKGSystemStatusIndicatorsShim.RebuildOcclusionMetadata]
//   - [PKGSystemStatusIndicatorsShim.SetDisplayNeedsSecureIndicatorFlush]
//   - [PKGSystemStatusIndicatorsShim.SuspendUpdates]
//   - [PKGSystemStatusIndicatorsShim.SetSuspendUpdates]
//   - [PKGSystemStatusIndicatorsShim.UpdateForIndicators]
//   - [PKGSystemStatusIndicatorsShim.InitWithWindowSize]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim
type PKGSystemStatusIndicatorsShim struct {
	objectivec.Object
}

// PKGSystemStatusIndicatorsShimFromID constructs a [PKGSystemStatusIndicatorsShim] from an objc.ID.
func PKGSystemStatusIndicatorsShimFromID(id objc.ID) PKGSystemStatusIndicatorsShim {
	return PKGSystemStatusIndicatorsShim{objectivec.Object{ID: id}}
}

// Ensure PKGSystemStatusIndicatorsShim implements IPKGSystemStatusIndicatorsShim.
var _ IPKGSystemStatusIndicatorsShim = PKGSystemStatusIndicatorsShim{}

// An interface definition for the [PKGSystemStatusIndicatorsShim] class.
//
// # Methods
//
//   - [IPKGSystemStatusIndicatorsShim.DisplayNeedsSecureIndicatorFlush]
//   - [IPKGSystemStatusIndicatorsShim.GlobalEnableSecureIndicators]
//   - [IPKGSystemStatusIndicatorsShim.SetGlobalEnableSecureIndicators]
//   - [IPKGSystemStatusIndicatorsShim.HasIndicatorWindows]
//   - [IPKGSystemStatusIndicatorsShim.OverrideIndicatorSize]
//   - [IPKGSystemStatusIndicatorsShim.RebuildDisplays]
//   - [IPKGSystemStatusIndicatorsShim.RebuildLayers]
//   - [IPKGSystemStatusIndicatorsShim.RebuildOcclusionMetadata]
//   - [IPKGSystemStatusIndicatorsShim.SetDisplayNeedsSecureIndicatorFlush]
//   - [IPKGSystemStatusIndicatorsShim.SuspendUpdates]
//   - [IPKGSystemStatusIndicatorsShim.SetSuspendUpdates]
//   - [IPKGSystemStatusIndicatorsShim.UpdateForIndicators]
//   - [IPKGSystemStatusIndicatorsShim.InitWithWindowSize]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim
type IPKGSystemStatusIndicatorsShim interface {
	objectivec.IObject

	// Topic: Methods

	DisplayNeedsSecureIndicatorFlush(flush objectivec.IObject) bool
	GlobalEnableSecureIndicators() bool
	SetGlobalEnableSecureIndicators(value bool)
	HasIndicatorWindows() bool
	OverrideIndicatorSize(size float32)
	RebuildDisplays(displays bool)
	RebuildLayers()
	RebuildOcclusionMetadata() uint32
	SetDisplayNeedsSecureIndicatorFlush(display objectivec.IObject, flush bool)
	SuspendUpdates() bool
	SetSuspendUpdates(value bool)
	UpdateForIndicators(indicators uint32)
	InitWithWindowSize(size corefoundation.CGSize) PKGSystemStatusIndicatorsShim
}

// Init initializes the instance.
func (g PKGSystemStatusIndicatorsShim) Init() PKGSystemStatusIndicatorsShim {
	rv := objc.Send[PKGSystemStatusIndicatorsShim](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g PKGSystemStatusIndicatorsShim) Autorelease() PKGSystemStatusIndicatorsShim {
	rv := objc.Send[PKGSystemStatusIndicatorsShim](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGSystemStatusIndicatorsShim creates a new PKGSystemStatusIndicatorsShim instance.
func NewPKGSystemStatusIndicatorsShim() PKGSystemStatusIndicatorsShim {
	class := getPKGSystemStatusIndicatorsShimClass()
	rv := objc.Send[PKGSystemStatusIndicatorsShim](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/initWithWindowSize:
func NewGSystemStatusIndicatorsShimWithWindowSize(size corefoundation.CGSize) PKGSystemStatusIndicatorsShim {
	instance := getPKGSystemStatusIndicatorsShimClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowSize:"), size)
	return PKGSystemStatusIndicatorsShimFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/displayNeedsSecureIndicatorFlush:
func (g PKGSystemStatusIndicatorsShim) DisplayNeedsSecureIndicatorFlush(flush objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("displayNeedsSecureIndicatorFlush:"), flush)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/hasIndicatorWindows
func (g PKGSystemStatusIndicatorsShim) HasIndicatorWindows() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasIndicatorWindows"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/overrideIndicatorSize:
func (g PKGSystemStatusIndicatorsShim) OverrideIndicatorSize(size float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("overrideIndicatorSize:"), size)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/rebuildDisplays:
func (g PKGSystemStatusIndicatorsShim) RebuildDisplays(displays bool) {
	objc.Send[objc.ID](g.ID, objc.Sel("rebuildDisplays:"), displays)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/rebuildLayers
func (g PKGSystemStatusIndicatorsShim) RebuildLayers() {
	objc.Send[objc.ID](g.ID, objc.Sel("rebuildLayers"))
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/rebuildOcclusionMetadata
func (g PKGSystemStatusIndicatorsShim) RebuildOcclusionMetadata() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("rebuildOcclusionMetadata"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/setDisplay:needsSecureIndicatorFlush:
func (g PKGSystemStatusIndicatorsShim) SetDisplayNeedsSecureIndicatorFlush(display objectivec.IObject, flush bool) {
	objc.Send[objc.ID](g.ID, objc.Sel("setDisplay:needsSecureIndicatorFlush:"), display, flush)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/updateForIndicators:
func (g PKGSystemStatusIndicatorsShim) UpdateForIndicators(indicators uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("updateForIndicators:"), indicators)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/initWithWindowSize:
func (g PKGSystemStatusIndicatorsShim) InitWithWindowSize(size corefoundation.CGSize) PKGSystemStatusIndicatorsShim {
	rv := objc.Send[PKGSystemStatusIndicatorsShim](g.ID, objc.Sel("initWithWindowSize:"), size)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/additionalBarInsetForModePixelWidth:pixelHeight:pointWidth:pointHeight:scale:physicalSize:
func (_PKGSystemStatusIndicatorsShimClass PKGSystemStatusIndicatorsShimClass) AdditionalBarInsetForModePixelWidthPixelHeightPointWidthPointHeightScalePhysicalSize(width float64, height float64, width2 float64, height2 float64, scale float64, size corefoundation.CGSize) float64 {
	rv := objc.Send[float64](objc.ID(_PKGSystemStatusIndicatorsShimClass.class), objc.Sel("additionalBarInsetForModePixelWidth:pixelHeight:pointWidth:pointHeight:scale:physicalSize:"), width, height, width2, height2, scale, size)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/globalEnableSecureIndicators
func (g PKGSystemStatusIndicatorsShim) GlobalEnableSecureIndicators() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("globalEnableSecureIndicators"))
	return rv
}
func (g PKGSystemStatusIndicatorsShim) SetGlobalEnableSecureIndicators(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setGlobalEnableSecureIndicators:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSystemStatusIndicatorsShim/suspendUpdates
func (g PKGSystemStatusIndicatorsShim) SuspendUpdates() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("suspendUpdates"))
	return rv
}
func (g PKGSystemStatusIndicatorsShim) SetSuspendUpdates(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setSuspendUpdates:"), value)
}
