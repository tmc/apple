// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCardATR] class.
var (
	_TKSmartCardATRClass     TKSmartCardATRClass
	_TKSmartCardATRClassOnce sync.Once
)

func getTKSmartCardATRClass() TKSmartCardATRClass {
	_TKSmartCardATRClassOnce.Do(func() {
		_TKSmartCardATRClass = TKSmartCardATRClass{class: objc.GetClass("TKSmartCardATR")}
	})
	return _TKSmartCardATRClass
}

// GetTKSmartCardATRClass returns the class object for TKSmartCardATR.
func GetTKSmartCardATRClass() TKSmartCardATRClass {
	return getTKSmartCardATRClass()
}

type TKSmartCardATRClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardATRClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardATRClass) Alloc() TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// A parsed ATR (Answer To Reset) message from a Smart Card.
//
// # Overview
//
// This class declares a programmatic interface to parsing an ATR from data or
// a byte stream, and accessing the individual parts.
//
// # Creating a Smart Card ATR
//
//   - [TKSmartCardATR.InitWithBytes]: Initializes a [TKSmartCardATR] object from a provided data object.
//   - [TKSmartCardATR.InitWithSource]: Initializes a [TKSmartCardATR] object from a provided data source.
//
// # Accessing ATR Attributes
//
//   - [TKSmartCardATR.Protocols]: An array of protocols indicated in the ATR
//   - [TKSmartCardATR.Bytes]: The ATR message data.
//   - [TKSmartCardATR.HistoricalBytes]: The ATR historical bytes, not including interface bytes or the TCK (check byte).
//   - [TKSmartCardATR.HistoricalRecords]: A list of compact TLV records parsed from historical bytes.
//
// # Retrieving Interface Groups
//
//   - [TKSmartCardATR.InterfaceGroupAtIndex]: Returns the interface group at the specified index.
//   - [TKSmartCardATR.InterfaceGroupForProtocol]: Returns the interface group with the specified protocol.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR
type TKSmartCardATR struct {
	objectivec.Object
}

// TKSmartCardATRFromID constructs a [TKSmartCardATR] from an objc.ID.
//
// A parsed ATR (Answer To Reset) message from a Smart Card.
func TKSmartCardATRFromID(id objc.ID) TKSmartCardATR {
	return TKSmartCardATR{objectivec.Object{ID: id}}
}

// NOTE: TKSmartCardATR adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardATR] class.
//
// # Creating a Smart Card ATR
//
//   - [ITKSmartCardATR.InitWithBytes]: Initializes a [TKSmartCardATR] object from a provided data object.
//   - [ITKSmartCardATR.InitWithSource]: Initializes a [TKSmartCardATR] object from a provided data source.
//
// # Accessing ATR Attributes
//
//   - [ITKSmartCardATR.Protocols]: An array of protocols indicated in the ATR
//   - [ITKSmartCardATR.Bytes]: The ATR message data.
//   - [ITKSmartCardATR.HistoricalBytes]: The ATR historical bytes, not including interface bytes or the TCK (check byte).
//   - [ITKSmartCardATR.HistoricalRecords]: A list of compact TLV records parsed from historical bytes.
//
// # Retrieving Interface Groups
//
//   - [ITKSmartCardATR.InterfaceGroupAtIndex]: Returns the interface group at the specified index.
//   - [ITKSmartCardATR.InterfaceGroupForProtocol]: Returns the interface group with the specified protocol.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR
type ITKSmartCardATR interface {
	objectivec.IObject

	// Topic: Creating a Smart Card ATR

	// Initializes a [TKSmartCardATR] object from a provided data object.
	InitWithBytes(bytes foundation.NSData) TKSmartCardATR
	// Initializes a [TKSmartCardATR] object from a provided data source.
	InitWithSource(source IntVoidHandler) TKSmartCardATR

	// Topic: Accessing ATR Attributes

	// An array of protocols indicated in the ATR
	Protocols() []foundation.NSNumber
	// The ATR message data.
	Bytes() foundation.NSData
	// The ATR historical bytes, not including interface bytes or the TCK (check byte).
	HistoricalBytes() foundation.NSData
	// A list of compact TLV records parsed from historical bytes.
	HistoricalRecords() []TKCompactTLVRecord

	// Topic: Retrieving Interface Groups

	// Returns the interface group at the specified index.
	InterfaceGroupAtIndex(index int) ITKSmartCardATRInterfaceGroup
	// Returns the interface group with the specified protocol.
	InterfaceGroupForProtocol(protocol_ TKSmartCardProtocol) ITKSmartCardATRInterfaceGroup
}

// Init initializes the instance.
func (t TKSmartCardATR) Init() TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardATR) Autorelease() TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardATR creates a new TKSmartCardATR instance.
func NewTKSmartCardATR() TKSmartCardATR {
	class := getTKSmartCardATRClass()
	rv := objc.Send[TKSmartCardATR](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a [TKSmartCardATR] object from a provided data object.
//
// bytes: The ATR data to be parsed.
//
// # Return Value
//
// A [TKSmartCardATR] object initialized with the parsed data. If `bytes` does
// not contain a valid ATR, returns `nil`.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/init(bytes:)
func NewTKSmartCardATRWithBytes(bytes foundation.NSData) TKSmartCardATR {
	instance := getTKSmartCardATRClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:"), bytes)
	return TKSmartCardATRFromID(rv)
}

// Initializes a [TKSmartCardATR] object from a provided data object.
//
// bytes: The ATR data to be parsed.
//
// # Return Value
//
// A [TKSmartCardATR] object initialized with the parsed data. If `bytes` does
// not contain a valid ATR, returns `nil`.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/init(bytes:)
func (t TKSmartCardATR) InitWithBytes(bytes foundation.NSData) TKSmartCardATR {
	rv := objc.Send[TKSmartCardATR](t.ID, objc.Sel("initWithBytes:"), bytes)
	return rv
}

// Initializes a [TKSmartCardATR] object from a provided data source.
//
// source: The block providing a stream of data for an ATR.
//
// The block takes no arguments and returns one byte. To indicate that an
// error occured, the block returns `-1`.
//
// # Return Value
//
// A [TKSmartCardATR] object initialized with the parsed data. If the byte
// stream produces an error or does not contain a valid ATR, returns `nil`.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/init(source:)
func (t TKSmartCardATR) InitWithSource(source IntVoidHandler) TKSmartCardATR {
	_block0, _ := NewIntVoidBlock(source)
	rv := objc.Send[TKSmartCardATR](t.ID, objc.Sel("initWithSource:"), _block0)
	return rv
}

// Returns the interface group at the specified index.
//
// index: The index of the desired interface group.
//
// # Return Value
//
// The interface group at the specified index, or `nil` if not present.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/interfaceGroup(at:)
func (t TKSmartCardATR) InterfaceGroupAtIndex(index int) ITKSmartCardATRInterfaceGroup {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("interfaceGroupAtIndex:"), index)
	return TKSmartCardATRInterfaceGroupFromID(rv)
}

// Returns the interface group with the specified protocol.
//
// protocol: The protocol used by the desired interface group.
//
// # Return Value
//
// The interface group with the specified protocol, or `nil` if none exists.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/interfaceGroup(for:)
func (t TKSmartCardATR) InterfaceGroupForProtocol(protocol_ TKSmartCardProtocol) ITKSmartCardATRInterfaceGroup {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("interfaceGroupForProtocol:"), protocol_)
	return TKSmartCardATRInterfaceGroupFromID(rv)
}

// An array of protocols indicated in the ATR
//
// # Discussion
//
// Each element in the returned array is an [NSNumber] object containing an
// [NSUInteger] value corresponding to a member of the [TKSmartCardProtocol]
// enumeration.
//
// The returned protocols are ordered such that the default protocol is at
// index `0`, and any duplicate values are removed.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/protocols
//
// [TKSmartCardProtocol]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardProtocol
func (t TKSmartCardATR) Protocols() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("protocols"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The ATR message data.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/bytes
func (t TKSmartCardATR) Bytes() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("bytes"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The ATR historical bytes, not including interface bytes or the TCK (check
// byte).
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/historicalBytes
func (t TKSmartCardATR) HistoricalBytes() foundation.NSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("historicalBytes"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A list of compact TLV records parsed from historical bytes.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardATR/historicalRecords
func (t TKSmartCardATR) HistoricalRecords() []TKCompactTLVRecord {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("historicalRecords"))
	return objc.ConvertSlice(rv, func(id objc.ID) TKCompactTLVRecord {
		return TKCompactTLVRecordFromID(id)
	})
}
