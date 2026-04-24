// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSIconAppearanceConfiguration] class.
var (
	_SLSIconAppearanceConfigurationClass     SLSIconAppearanceConfigurationClass
	_SLSIconAppearanceConfigurationClassOnce sync.Once
)

func getSLSIconAppearanceConfigurationClass() SLSIconAppearanceConfigurationClass {
	_SLSIconAppearanceConfigurationClassOnce.Do(func() {
		_SLSIconAppearanceConfigurationClass = SLSIconAppearanceConfigurationClass{class: objc.GetClass("SLSIconAppearanceConfiguration")}
	})
	return _SLSIconAppearanceConfigurationClass
}

// GetSLSIconAppearanceConfigurationClass returns the class object for SLSIconAppearanceConfiguration.
func GetSLSIconAppearanceConfigurationClass() SLSIconAppearanceConfigurationClass {
	return getSLSIconAppearanceConfigurationClass()
}

type SLSIconAppearanceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSIconAppearanceConfigurationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSIconAppearanceConfigurationClass) Alloc() SLSIconAppearanceConfiguration {
	rv := objc.Send[SLSIconAppearanceConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSIconAppearanceConfiguration._initWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme]
//   - [SLSIconAppearanceConfiguration.AppearanceTheme]
//   - [SLSIconAppearanceConfiguration.SetAppearanceTheme]
//   - [SLSIconAppearanceConfiguration.IconAppearanceTheme]
//   - [SLSIconAppearanceConfiguration.SetIconAppearanceTheme]
//   - [SLSIconAppearanceConfiguration.IconTintColorName]
//   - [SLSIconAppearanceConfiguration.SetIconTintColorName]
//   - [SLSIconAppearanceConfiguration.OtherIconTintColor]
//   - [SLSIconAppearanceConfiguration.SetOtherIconTintColor]
//   - [SLSIconAppearanceConfiguration.Save]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration
type SLSIconAppearanceConfiguration struct {
	objectivec.Object
}

// SLSIconAppearanceConfigurationFromID constructs a [SLSIconAppearanceConfiguration] from an objc.ID.
func SLSIconAppearanceConfigurationFromID(id objc.ID) SLSIconAppearanceConfiguration {
	return SLSIconAppearanceConfiguration{objectivec.Object{ID: id}}
}

// Ensure SLSIconAppearanceConfiguration implements ISLSIconAppearanceConfiguration.
var _ ISLSIconAppearanceConfiguration = SLSIconAppearanceConfiguration{}

// An interface definition for the [SLSIconAppearanceConfiguration] class.
//
// # Methods
//
//   - [ISLSIconAppearanceConfiguration._initWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme]
//   - [ISLSIconAppearanceConfiguration.AppearanceTheme]
//   - [ISLSIconAppearanceConfiguration.SetAppearanceTheme]
//   - [ISLSIconAppearanceConfiguration.IconAppearanceTheme]
//   - [ISLSIconAppearanceConfiguration.SetIconAppearanceTheme]
//   - [ISLSIconAppearanceConfiguration.IconTintColorName]
//   - [ISLSIconAppearanceConfiguration.SetIconTintColorName]
//   - [ISLSIconAppearanceConfiguration.OtherIconTintColor]
//   - [ISLSIconAppearanceConfiguration.SetOtherIconTintColor]
//   - [ISLSIconAppearanceConfiguration.Save]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration
type ISLSIconAppearanceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_initWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme(theme uint32, name uint32, color coregraphics.CGColorRef, theme2 uint32) objectivec.IObject
	AppearanceTheme() uint32
	SetAppearanceTheme(value uint32)
	IconAppearanceTheme() uint32
	SetIconAppearanceTheme(value uint32)
	IconTintColorName() uint32
	SetIconTintColorName(value uint32)
	OtherIconTintColor() coregraphics.CGColorRef
	SetOtherIconTintColor(value coregraphics.CGColorRef)
	Save()
}

// Init initializes the instance.
func (s SLSIconAppearanceConfiguration) Init() SLSIconAppearanceConfiguration {
	rv := objc.Send[SLSIconAppearanceConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSIconAppearanceConfiguration) Autorelease() SLSIconAppearanceConfiguration {
	rv := objc.Send[SLSIconAppearanceConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSIconAppearanceConfiguration creates a new SLSIconAppearanceConfiguration instance.
func NewSLSIconAppearanceConfiguration() SLSIconAppearanceConfiguration {
	class := getSLSIconAppearanceConfigurationClass()
	rv := objc.Send[SLSIconAppearanceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration/_initWithIconAppearanceTheme:iconTintColorName:otherIconTintColor:appearanceTheme:
func (s SLSIconAppearanceConfiguration) _initWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme(theme uint32, name uint32, color coregraphics.CGColorRef, theme2 uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_initWithIconAppearanceTheme:iconTintColorName:otherIconTintColor:appearanceTheme:"), theme, name, color, theme2)
	return objectivec.Object{ID: rv}
}

// InitWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme is an exported wrapper for the private method _initWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme.
func (s SLSIconAppearanceConfiguration) InitWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme(theme uint32, name uint32, color coregraphics.CGColorRef, theme2 uint32) objectivec.IObject {
	return s._initWithIconAppearanceThemeIconTintColorNameOtherIconTintColorAppearanceTheme(theme, name, color, theme2)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration/save
func (s SLSIconAppearanceConfiguration) Save() {
	objc.Send[objc.ID](s.ID, objc.Sel("save"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration/fetchCurrentIconAppearanceConfiguration
func (_SLSIconAppearanceConfigurationClass SLSIconAppearanceConfigurationClass) FetchCurrentIconAppearanceConfiguration() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSIconAppearanceConfigurationClass.class), objc.Sel("fetchCurrentIconAppearanceConfiguration"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration/appearanceTheme
func (s SLSIconAppearanceConfiguration) AppearanceTheme() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appearanceTheme"))
	return rv
}
func (s SLSIconAppearanceConfiguration) SetAppearanceTheme(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setAppearanceTheme:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration/iconAppearanceTheme
func (s SLSIconAppearanceConfiguration) IconAppearanceTheme() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("iconAppearanceTheme"))
	return rv
}
func (s SLSIconAppearanceConfiguration) SetIconAppearanceTheme(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setIconAppearanceTheme:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration/iconTintColorName
func (s SLSIconAppearanceConfiguration) IconTintColorName() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("iconTintColorName"))
	return rv
}
func (s SLSIconAppearanceConfiguration) SetIconTintColorName(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setIconTintColorName:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSIconAppearanceConfiguration/otherIconTintColor
func (s SLSIconAppearanceConfiguration) OtherIconTintColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](s.ID, objc.Sel("otherIconTintColor"))
	return coregraphics.CGColorRef(rv)
}
func (s SLSIconAppearanceConfiguration) SetOtherIconTintColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setOtherIconTintColor:"), value)
}
