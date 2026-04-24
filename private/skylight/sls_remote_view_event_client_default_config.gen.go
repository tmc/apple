// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSRemoteViewEventClientDefaultConfig] class.
var (
	_SLSRemoteViewEventClientDefaultConfigClass     SLSRemoteViewEventClientDefaultConfigClass
	_SLSRemoteViewEventClientDefaultConfigClassOnce sync.Once
)

func getSLSRemoteViewEventClientDefaultConfigClass() SLSRemoteViewEventClientDefaultConfigClass {
	_SLSRemoteViewEventClientDefaultConfigClassOnce.Do(func() {
		_SLSRemoteViewEventClientDefaultConfigClass = SLSRemoteViewEventClientDefaultConfigClass{class: objc.GetClass("_SLSRemoteViewEventClientDefaultConfig")}
	})
	return _SLSRemoteViewEventClientDefaultConfigClass
}

// GetSLSRemoteViewEventClientDefaultConfigClass returns the class object for _SLSRemoteViewEventClientDefaultConfig.
func GetSLSRemoteViewEventClientDefaultConfigClass() SLSRemoteViewEventClientDefaultConfigClass {
	return getSLSRemoteViewEventClientDefaultConfigClass()
}

type SLSRemoteViewEventClientDefaultConfigClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSRemoteViewEventClientDefaultConfigClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSRemoteViewEventClientDefaultConfigClass) Alloc() SLSRemoteViewEventClientDefaultConfig {
	rv := objc.Send[SLSRemoteViewEventClientDefaultConfig](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSRemoteViewEventClientDefaultConfig.Connection]
//   - [SLSRemoteViewEventClientDefaultConfig.ServiceInterface]
//   - [SLSRemoteViewEventClientDefaultConfig.DebugDescription]
//   - [SLSRemoteViewEventClientDefaultConfig.Description]
//   - [SLSRemoteViewEventClientDefaultConfig.Hash]
//   - [SLSRemoteViewEventClientDefaultConfig.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig
type SLSRemoteViewEventClientDefaultConfig struct {
	objectivec.Object
}

// SLSRemoteViewEventClientDefaultConfigFromID constructs a [SLSRemoteViewEventClientDefaultConfig] from an objc.ID.
func SLSRemoteViewEventClientDefaultConfigFromID(id objc.ID) SLSRemoteViewEventClientDefaultConfig {
	return SLSRemoteViewEventClientDefaultConfig{objectivec.Object{ID: id}}
}

// Ensure SLSRemoteViewEventClientDefaultConfig implements ISLSRemoteViewEventClientDefaultConfig.
var _ ISLSRemoteViewEventClientDefaultConfig = SLSRemoteViewEventClientDefaultConfig{}

// An interface definition for the [SLSRemoteViewEventClientDefaultConfig] class.
//
// # Methods
//
//   - [ISLSRemoteViewEventClientDefaultConfig.Connection]
//   - [ISLSRemoteViewEventClientDefaultConfig.ServiceInterface]
//   - [ISLSRemoteViewEventClientDefaultConfig.DebugDescription]
//   - [ISLSRemoteViewEventClientDefaultConfig.Description]
//   - [ISLSRemoteViewEventClientDefaultConfig.Hash]
//   - [ISLSRemoteViewEventClientDefaultConfig.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig
type ISLSRemoteViewEventClientDefaultConfig interface {
	objectivec.IObject

	// Topic: Methods

	Connection() objectivec.IObject
	ServiceInterface() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s SLSRemoteViewEventClientDefaultConfig) Init() SLSRemoteViewEventClientDefaultConfig {
	rv := objc.Send[SLSRemoteViewEventClientDefaultConfig](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSRemoteViewEventClientDefaultConfig) Autorelease() SLSRemoteViewEventClientDefaultConfig {
	rv := objc.Send[SLSRemoteViewEventClientDefaultConfig](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSRemoteViewEventClientDefaultConfig creates a new SLSRemoteViewEventClientDefaultConfig instance.
func NewSLSRemoteViewEventClientDefaultConfig() SLSRemoteViewEventClientDefaultConfig {
	class := getSLSRemoteViewEventClientDefaultConfigClass()
	rv := objc.Send[SLSRemoteViewEventClientDefaultConfig](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig/connection
func (s SLSRemoteViewEventClientDefaultConfig) Connection() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("connection"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig/serviceInterface
func (s SLSRemoteViewEventClientDefaultConfig) ServiceInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("serviceInterface"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig/debugDescription
func (s SLSRemoteViewEventClientDefaultConfig) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig/description
func (s SLSRemoteViewEventClientDefaultConfig) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig/hash
func (s SLSRemoteViewEventClientDefaultConfig) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_SLSRemoteViewEventClientDefaultConfig/superclass
func (s SLSRemoteViewEventClientDefaultConfig) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
