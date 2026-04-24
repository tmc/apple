// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPKGPrivacyIndicatorWindowContainer] class.
var (
	_SkyLightPKGPrivacyIndicatorWindowContainerClass     SkyLightPKGPrivacyIndicatorWindowContainerClass
	_SkyLightPKGPrivacyIndicatorWindowContainerClassOnce sync.Once
)

func getSkyLightPKGPrivacyIndicatorWindowContainerClass() SkyLightPKGPrivacyIndicatorWindowContainerClass {
	_SkyLightPKGPrivacyIndicatorWindowContainerClassOnce.Do(func() {
		_SkyLightPKGPrivacyIndicatorWindowContainerClass = SkyLightPKGPrivacyIndicatorWindowContainerClass{class: objc.GetClass("SkyLight.PKGPrivacyIndicatorWindowContainer")}
	})
	return _SkyLightPKGPrivacyIndicatorWindowContainerClass
}

// GetSkyLightPKGPrivacyIndicatorWindowContainerClass returns the class object for SkyLight.PKGPrivacyIndicatorWindowContainer.
func GetSkyLightPKGPrivacyIndicatorWindowContainerClass() SkyLightPKGPrivacyIndicatorWindowContainerClass {
	return getSkyLightPKGPrivacyIndicatorWindowContainerClass()
}

type SkyLightPKGPrivacyIndicatorWindowContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPKGPrivacyIndicatorWindowContainerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPKGPrivacyIndicatorWindowContainerClass) Alloc() SkyLightPKGPrivacyIndicatorWindowContainer {
	rv := objc.Send[SkyLightPKGPrivacyIndicatorWindowContainer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGPrivacyIndicatorWindowContainer
type SkyLightPKGPrivacyIndicatorWindowContainer struct {
	objectivec.Object
}

// SkyLightPKGPrivacyIndicatorWindowContainerFromID constructs a [SkyLightPKGPrivacyIndicatorWindowContainer] from an objc.ID.
func SkyLightPKGPrivacyIndicatorWindowContainerFromID(id objc.ID) SkyLightPKGPrivacyIndicatorWindowContainer {
	return SkyLightPKGPrivacyIndicatorWindowContainer{objectivec.Object{ID: id}}
}

// NOTE: SkyLightPKGPrivacyIndicatorWindowContainer struct embeds objectivec.Object (parent type unavailable) but
// ISkyLightPKGPrivacyIndicatorWindowContainer embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SkyLightPKGPrivacyIndicatorWindowContainer] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGPrivacyIndicatorWindowContainer
type ISkyLightPKGPrivacyIndicatorWindowContainer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPKGPrivacyIndicatorWindowContainer) Init() SkyLightPKGPrivacyIndicatorWindowContainer {
	rv := objc.Send[SkyLightPKGPrivacyIndicatorWindowContainer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPKGPrivacyIndicatorWindowContainer) Autorelease() SkyLightPKGPrivacyIndicatorWindowContainer {
	rv := objc.Send[SkyLightPKGPrivacyIndicatorWindowContainer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPKGPrivacyIndicatorWindowContainer creates a new SkyLightPKGPrivacyIndicatorWindowContainer instance.
func NewSkyLightPKGPrivacyIndicatorWindowContainer() SkyLightPKGPrivacyIndicatorWindowContainer {
	class := getSkyLightPKGPrivacyIndicatorWindowContainerClass()
	rv := objc.Send[SkyLightPKGPrivacyIndicatorWindowContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
