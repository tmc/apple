// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [TKSimpleTLVRecord] class.
var (
	_TKSimpleTLVRecordClass     TKSimpleTLVRecordClass
	_TKSimpleTLVRecordClassOnce sync.Once
)

func getTKSimpleTLVRecordClass() TKSimpleTLVRecordClass {
	_TKSimpleTLVRecordClassOnce.Do(func() {
		_TKSimpleTLVRecordClass = TKSimpleTLVRecordClass{class: objc.GetClass("TKSimpleTLVRecord")}
	})
	return _TKSimpleTLVRecordClass
}

// GetTKSimpleTLVRecordClass returns the class object for TKSimpleTLVRecord.
func GetTKSimpleTLVRecordClass() TKSimpleTLVRecordClass {
	return getTKSimpleTLVRecordClass()
}

type TKSimpleTLVRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSimpleTLVRecordClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSimpleTLVRecordClass) Alloc() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// An object that implements encoding using Simple-TLV encoding according to
// ISO 7816-4.
//
// # Creating TLV Records
//
//   - [TKSimpleTLVRecord.InitWithTagValue]: Initializes a TLV record with the specified tag and value.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSimpleTLVRecord
type TKSimpleTLVRecord struct {
	TKTLVRecord
}

// TKSimpleTLVRecordFromID constructs a [TKSimpleTLVRecord] from an objc.ID.
//
// An object that implements encoding using Simple-TLV encoding according to
// ISO 7816-4.
func TKSimpleTLVRecordFromID(id objc.ID) TKSimpleTLVRecord {
	return TKSimpleTLVRecord{TKTLVRecord: TKTLVRecordFromID(id)}
}

// NOTE: TKSimpleTLVRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSimpleTLVRecord] class.
//
// # Creating TLV Records
//
//   - [ITKSimpleTLVRecord.InitWithTagValue]: Initializes a TLV record with the specified tag and value.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSimpleTLVRecord
type ITKSimpleTLVRecord interface {
	ITKTLVRecord

	// Topic: Creating TLV Records

	// Initializes a TLV record with the specified tag and value.
	InitWithTagValue(tag uint8, value foundation.NSData) TKSimpleTLVRecord
}

// Init initializes the instance.
func (t TKSimpleTLVRecord) Init() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSimpleTLVRecord) Autorelease() TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSimpleTLVRecord creates a new TKSimpleTLVRecord instance.
func NewTKSimpleTLVRecord() TKSimpleTLVRecord {
	class := getTKSimpleTLVRecordClass()
	rv := objc.Send[TKSimpleTLVRecord](objc.ID(class.class), objc.Sel("new"))
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
func NewTKSimpleTLVRecordFromData(data foundation.NSData) TKSimpleTLVRecord {
	rv := objc.Send[objc.ID](objc.ID(getTKSimpleTLVRecordClass().class), objc.Sel("recordFromData:"), data)
	return TKSimpleTLVRecordFromID(rv)
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
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSimpleTLVRecord/init(tag:value:)
func NewTKSimpleTLVRecordWithTagValue(tag uint8, value foundation.NSData) TKSimpleTLVRecord {
	instance := getTKSimpleTLVRecordClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTag:value:"), tag, value)
	return TKSimpleTLVRecordFromID(rv)
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
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSimpleTLVRecord/init(tag:value:)
func (t TKSimpleTLVRecord) InitWithTagValue(tag uint8, value foundation.NSData) TKSimpleTLVRecord {
	rv := objc.Send[TKSimpleTLVRecord](t.ID, objc.Sel("initWithTag:value:"), tag, value)
	return rv
}
