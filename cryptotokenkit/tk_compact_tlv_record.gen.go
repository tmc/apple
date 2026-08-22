// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [TKCompactTLVRecord] class.
var (
	_TKCompactTLVRecordClass     TKCompactTLVRecordClass
	_TKCompactTLVRecordClassOnce sync.Once
)

func getTKCompactTLVRecordClass() TKCompactTLVRecordClass {
	_TKCompactTLVRecordClassOnce.Do(func() {
		_TKCompactTLVRecordClass = TKCompactTLVRecordClass{class: objc.GetClass("TKCompactTLVRecord")}
	})
	return _TKCompactTLVRecordClass
}

// GetTKCompactTLVRecordClass returns the class object for TKCompactTLVRecord.
func GetTKCompactTLVRecordClass() TKCompactTLVRecordClass {
	return getTKCompactTLVRecordClass()
}

type TKCompactTLVRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKCompactTLVRecordClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKCompactTLVRecordClass) Alloc() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// An object that implements encoding using Compact-TLV encoding according to
// ISO 7816-4.
//
// # Creating TLV Records
//
//   - [TKCompactTLVRecord.InitWithTagValue]: Initializes a TLV record with the specified tag and value.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord
type TKCompactTLVRecord struct {
	TKTLVRecord
}

// TKCompactTLVRecordFromID constructs a [TKCompactTLVRecord] from an objc.ID.
//
// An object that implements encoding using Compact-TLV encoding according to
// ISO 7816-4.
func TKCompactTLVRecordFromID(id objc.ID) TKCompactTLVRecord {
	return TKCompactTLVRecord{TKTLVRecord: TKTLVRecordFromID(id)}
}

// NOTE: TKCompactTLVRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKCompactTLVRecord] class.
//
// # Creating TLV Records
//
//   - [ITKCompactTLVRecord.InitWithTagValue]: Initializes a TLV record with the specified tag and value.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord
type ITKCompactTLVRecord interface {
	ITKTLVRecord

	// Topic: Creating TLV Records

	// Initializes a TLV record with the specified tag and value.
	InitWithTagValue(tag uint8, value foundation.NSData) TKCompactTLVRecord
}

// Init initializes the instance.
func (t TKCompactTLVRecord) Init() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKCompactTLVRecord) Autorelease() TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKCompactTLVRecord creates a new TKCompactTLVRecord instance.
func NewTKCompactTLVRecord() TKCompactTLVRecord {
	class := getTKCompactTLVRecordClass()
	rv := objc.Send[TKCompactTLVRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a TLV record from by parsing the specified data.
//
// data: A data object containing the serialized representation of a TLV record.
//
// # Return Value
//
// A TLV record, or `nil` if `data` does not specify a valid record.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVRecord/init(from:)
func NewTKCompactTLVRecordFromData(data foundation.NSData) TKCompactTLVRecord {
	rv := objc.Send[objc.ID](objc.ID(getTKCompactTLVRecordClass().class), objc.Sel("recordFromData:"), data)
	return TKCompactTLVRecordFromID(rv)
}

// Initializes a TLV record with the specified tag and value.
//
// tag: The tag field of the record.
//
// value: The value field of the record.
//
// # Return Value
//
// A new TLV record containing the specified tag and value fields.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord/init(tag:value:)
func NewTKCompactTLVRecordWithTagValue(tag uint8, value foundation.NSData) TKCompactTLVRecord {
	instance := getTKCompactTLVRecordClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:value:"), tag, value)
	return TKCompactTLVRecordFromID(rv)
}

// Initializes a TLV record with the specified tag and value.
//
// tag: The tag field of the record.
//
// value: The value field of the record.
//
// # Return Value
//
// A new TLV record containing the specified tag and value fields.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKCompactTLVRecord/init(tag:value:)
func (t TKCompactTLVRecord) InitWithTagValue(tag uint8, value foundation.NSData) TKCompactTLVRecord {
	rv := objc.Send[TKCompactTLVRecord](t.ID, objc.Sel("initWithTag:value:"), tag, value)
	return rv
}
