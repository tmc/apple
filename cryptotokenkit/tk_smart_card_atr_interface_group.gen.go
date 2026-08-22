// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCardATRInterfaceGroup] class.
var (
	_TKSmartCardATRInterfaceGroupClass     TKSmartCardATRInterfaceGroupClass
	_TKSmartCardATRInterfaceGroupClassOnce sync.Once
)

func getTKSmartCardATRInterfaceGroupClass() TKSmartCardATRInterfaceGroupClass {
	_TKSmartCardATRInterfaceGroupClassOnce.Do(func() {
		_TKSmartCardATRInterfaceGroupClass = TKSmartCardATRInterfaceGroupClass{class: objc.GetClass("TKSmartCardATRInterfaceGroup")}
	})
	return _TKSmartCardATRInterfaceGroupClass
}

// GetTKSmartCardATRInterfaceGroupClass returns the class object for TKSmartCardATRInterfaceGroup.
func GetTKSmartCardATRInterfaceGroupClass() TKSmartCardATRInterfaceGroupClass {
	return getTKSmartCardATRInterfaceGroupClass()
}

type TKSmartCardATRInterfaceGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardATRInterfaceGroupClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardATRInterfaceGroupClass) Alloc() TKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A single interface-bytes group for a Smart Card ATR (Answer to Reset).
//
// # Overview
//
// You access instances of this class by calling the
// [TKSmartCardATR.InterfaceGroupAtIndex] and
// [TKSmartCardATR.InterfaceGroupForProtocol] methods on an [TKSmartCardATR]
// object.
//
// # Accessing Interface Group Protocol and Bytes
//
//   - [TKSmartCardATRInterfaceGroup.Protocol]: The protocol for this group.
//   - [TKSmartCardATRInterfaceGroup.TA]: The TA interface byte of ATR group, or `nil` if TA is not present.
//   - [TKSmartCardATRInterfaceGroup.TB]: The TB interface byte of ATR group, or `nil` if TB is not present.
//   - [TKSmartCardATRInterfaceGroup.TC]: The TC interface byte of ATR group, or `nil` if TC is not present.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup
type TKSmartCardATRInterfaceGroup struct {
	objectivec.Object
}

// TKSmartCardATRInterfaceGroupFromID constructs a [TKSmartCardATRInterfaceGroup] from an objc.ID.
//
// A single interface-bytes group for a Smart Card ATR (Answer to Reset).
func TKSmartCardATRInterfaceGroupFromID(id objc.ID) TKSmartCardATRInterfaceGroup {
	return TKSmartCardATRInterfaceGroup{objectivec.Object{ID: id}}
}

// NOTE: TKSmartCardATRInterfaceGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardATRInterfaceGroup] class.
//
// # Accessing Interface Group Protocol and Bytes
//
//   - [ITKSmartCardATRInterfaceGroup.Protocol]: The protocol for this group.
//   - [ITKSmartCardATRInterfaceGroup.TA]: The TA interface byte of ATR group, or `nil` if TA is not present.
//   - [ITKSmartCardATRInterfaceGroup.TB]: The TB interface byte of ATR group, or `nil` if TB is not present.
//   - [ITKSmartCardATRInterfaceGroup.TC]: The TC interface byte of ATR group, or `nil` if TC is not present.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup
type ITKSmartCardATRInterfaceGroup interface {
	objectivec.IObject

	// Topic: Accessing Interface Group Protocol and Bytes

	// The protocol for this group.
	Protocol() foundation.NSNumber
	// The TA interface byte of ATR group, or `nil` if TA is not present.
	TA() foundation.NSNumber
	// The TB interface byte of ATR group, or `nil` if TB is not present.
	TB() foundation.NSNumber
	// The TC interface byte of ATR group, or `nil` if TC is not present.
	TC() foundation.NSNumber
}

// Init initializes the instance.
func (t TKSmartCardATRInterfaceGroup) Init() TKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardATRInterfaceGroup) Autorelease() TKSmartCardATRInterfaceGroup {
	rv := objc.Send[TKSmartCardATRInterfaceGroup](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardATRInterfaceGroup creates a new TKSmartCardATRInterfaceGroup instance.
func NewTKSmartCardATRInterfaceGroup() TKSmartCardATRInterfaceGroup {
	class := getTKSmartCardATRInterfaceGroupClass()
	rv := objc.Send[TKSmartCardATRInterfaceGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The protocol for this group.
//
// # Discussion
//
// This property returns an [NSNumber] object containing an [NSUInteger] value
// corresponding to a member of the [TKSmartCardProtocol] enumeration.
//
// This property is `nil` for the first interface group (global), as it has no
// assigned protocol.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/protocol
//
// [TKSmartCardProtocol]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol
func (t TKSmartCardATRInterfaceGroup) Protocol() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("protocol"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The TA interface byte of ATR group, or `nil` if TA is not present.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/ta
func (t TKSmartCardATRInterfaceGroup) TA() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("TA"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The TB interface byte of ATR group, or `nil` if TB is not present.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/tb
func (t TKSmartCardATRInterfaceGroup) TB() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("TB"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The TC interface byte of ATR group, or `nil` if TC is not present.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/InterfaceGroup/tc
func (t TKSmartCardATRInterfaceGroup) TC() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("TC"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
