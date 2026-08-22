// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKTokenWatcherTokenInfo] class.
var (
	_TKTokenWatcherTokenInfoClass     TKTokenWatcherTokenInfoClass
	_TKTokenWatcherTokenInfoClassOnce sync.Once
)

func getTKTokenWatcherTokenInfoClass() TKTokenWatcherTokenInfoClass {
	_TKTokenWatcherTokenInfoClassOnce.Do(func() {
		_TKTokenWatcherTokenInfoClass = TKTokenWatcherTokenInfoClass{class: objc.GetClass("TKTokenWatcherTokenInfo")}
	})
	return _TKTokenWatcherTokenInfoClass
}

// GetTKTokenWatcherTokenInfoClass returns the class object for TKTokenWatcherTokenInfo.
func GetTKTokenWatcherTokenInfoClass() TKTokenWatcherTokenInfoClass {
	return getTKTokenWatcherTokenInfoClass()
}

type TKTokenWatcherTokenInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKTokenWatcherTokenInfoClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKTokenWatcherTokenInfoClass) Alloc() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [TKTokenWatcherTokenInfo.DriverName]
//   - [TKTokenWatcherTokenInfo.SlotName]
//   - [TKTokenWatcherTokenInfo.TokenID]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo
type TKTokenWatcherTokenInfo struct {
	objectivec.Object
}

// TKTokenWatcherTokenInfoFromID constructs a [TKTokenWatcherTokenInfo] from an objc.ID.
func TKTokenWatcherTokenInfoFromID(id objc.ID) TKTokenWatcherTokenInfo {
	return TKTokenWatcherTokenInfo{objectivec.Object{ID: id}}
}

// NOTE: TKTokenWatcherTokenInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKTokenWatcherTokenInfo] class.
//
// # Instance Properties
//
//   - [ITKTokenWatcherTokenInfo.DriverName]
//   - [ITKTokenWatcherTokenInfo.SlotName]
//   - [ITKTokenWatcherTokenInfo.TokenID]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo
type ITKTokenWatcherTokenInfo interface {
	objectivec.IObject

	// Topic: Instance Properties

	DriverName() string
	SlotName() string
	TokenID() string
}

// Init initializes the instance.
func (t TKTokenWatcherTokenInfo) Init() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKTokenWatcherTokenInfo) Autorelease() TKTokenWatcherTokenInfo {
	rv := objc.Send[TKTokenWatcherTokenInfo](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenWatcherTokenInfo creates a new TKTokenWatcherTokenInfo instance.
func NewTKTokenWatcherTokenInfo() TKTokenWatcherTokenInfo {
	class := getTKTokenWatcherTokenInfoClass()
	rv := objc.Send[TKTokenWatcherTokenInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/driverName
func (t TKTokenWatcherTokenInfo) DriverName() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("driverName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/slotName
func (t TKTokenWatcherTokenInfo) SlotName() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("slotName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenWatcher/TokenInfo/tokenID
func (t TKTokenWatcherTokenInfo) TokenID() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tokenID"))
	return foundation.NSStringFromID(rv).String()
}
