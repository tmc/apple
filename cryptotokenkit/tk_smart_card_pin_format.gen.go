// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TKSmartCardPINFormat] class.
var (
	_TKSmartCardPINFormatClass     TKSmartCardPINFormatClass
	_TKSmartCardPINFormatClassOnce sync.Once
)

func getTKSmartCardPINFormatClass() TKSmartCardPINFormatClass {
	_TKSmartCardPINFormatClassOnce.Do(func() {
		_TKSmartCardPINFormatClass = TKSmartCardPINFormatClass{class: objc.GetClass("TKSmartCardPINFormat")}
	})
	return _TKSmartCardPINFormatClass
}

// GetTKSmartCardPINFormatClass returns the class object for TKSmartCardPINFormat.
func GetTKSmartCardPINFormatClass() TKSmartCardPINFormatClass {
	return getTKSmartCardPINFormatClass()
}

type TKSmartCardPINFormatClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TKSmartCardPINFormatClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TKSmartCardPINFormatClass) Alloc() TKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// The formatting properties for a PIN, such as character encoding and length
// constraints.
//
// # Overview
//
// You typically interact with [TKSmartCardPINFormat] objects when calling the
// [TKSmartCard.UserInteractionForSecurePINChangeWithPINFormatAPDUCurrentPINByteOffsetNewPINByteOffset]
// and
// [TKSmartCard.UserInteractionForSecurePINVerificationWithPINFormatAPDUPINByteOffset]
// methods on an instance of [TKSmartCard].
//
// # Configuring PIN Formatting
//
//   - [TKSmartCardPINFormat.Charset]: The format of PIN characters. [TKSmartCardPINCharsetNumeric] by default.
//   - [TKSmartCardPINFormat.SetCharset]
//   - [TKSmartCardPINFormat.Encoding]: The encoding of PIN characters. [TKSmartCardPINEncodingASCII] by default.
//   - [TKSmartCardPINFormat.SetEncoding]
//   - [TKSmartCardPINFormat.MinPINLength]: The minimum number of characters to form a valid PIN. `4` by default.
//   - [TKSmartCardPINFormat.SetMinPINLength]
//   - [TKSmartCardPINFormat.MaxPINLength]: The maximum number of characters to form a valid PIN. `8` by default.
//   - [TKSmartCardPINFormat.SetMaxPINLength]
//   - [TKSmartCardPINFormat.PINBlockByteLength]: The total length of the PIN block in bytes. `8` by default.
//   - [TKSmartCardPINFormat.SetPINBlockByteLength]
//   - [TKSmartCardPINFormat.PINJustification]: The justification within the PIN block. [TKSmartCardPINJustificationLeft] by default.
//   - [TKSmartCardPINFormat.SetPINJustification]
//   - [TKSmartCardPINFormat.PINBitOffset]: The offset, in bits, within the PIN block to mark a location for filling in the formatted PIN, which is justified with respect to the [pinJustification](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinJustification>) property value. `0` by default.
//   - [TKSmartCardPINFormat.SetPINBitOffset]
//   - [TKSmartCardPINFormat.PINLengthBitOffset]: The offset, in bits, within the PIN block to mark a location for filling in the PIN length, which is always left justified. `0` by default.
//   - [TKSmartCardPINFormat.SetPINLengthBitOffset]
//   - [TKSmartCardPINFormat.PINLengthBitSize]: The size, in bits, of the PIN length field. If set to `0`, PIN length is not written. `0` by default.
//   - [TKSmartCardPINFormat.SetPINLengthBitSize]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat
type TKSmartCardPINFormat struct {
	objectivec.Object
}

// TKSmartCardPINFormatFromID constructs a [TKSmartCardPINFormat] from an objc.ID.
//
// The formatting properties for a PIN, such as character encoding and length
// constraints.
func TKSmartCardPINFormatFromID(id objc.ID) TKSmartCardPINFormat {
	return TKSmartCardPINFormat{objectivec.Object{ID: id}}
}

// NOTE: TKSmartCardPINFormat adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [TKSmartCardPINFormat] class.
//
// # Configuring PIN Formatting
//
//   - [ITKSmartCardPINFormat.Charset]: The format of PIN characters. [TKSmartCardPINCharsetNumeric] by default.
//   - [ITKSmartCardPINFormat.SetCharset]
//   - [ITKSmartCardPINFormat.Encoding]: The encoding of PIN characters. [TKSmartCardPINEncodingASCII] by default.
//   - [ITKSmartCardPINFormat.SetEncoding]
//   - [ITKSmartCardPINFormat.MinPINLength]: The minimum number of characters to form a valid PIN. `4` by default.
//   - [ITKSmartCardPINFormat.SetMinPINLength]
//   - [ITKSmartCardPINFormat.MaxPINLength]: The maximum number of characters to form a valid PIN. `8` by default.
//   - [ITKSmartCardPINFormat.SetMaxPINLength]
//   - [ITKSmartCardPINFormat.PINBlockByteLength]: The total length of the PIN block in bytes. `8` by default.
//   - [ITKSmartCardPINFormat.SetPINBlockByteLength]
//   - [ITKSmartCardPINFormat.PINJustification]: The justification within the PIN block. [TKSmartCardPINJustificationLeft] by default.
//   - [ITKSmartCardPINFormat.SetPINJustification]
//   - [ITKSmartCardPINFormat.PINBitOffset]: The offset, in bits, within the PIN block to mark a location for filling in the formatted PIN, which is justified with respect to the [pinJustification](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinJustification>) property value. `0` by default.
//   - [ITKSmartCardPINFormat.SetPINBitOffset]
//   - [ITKSmartCardPINFormat.PINLengthBitOffset]: The offset, in bits, within the PIN block to mark a location for filling in the PIN length, which is always left justified. `0` by default.
//   - [ITKSmartCardPINFormat.SetPINLengthBitOffset]
//   - [ITKSmartCardPINFormat.PINLengthBitSize]: The size, in bits, of the PIN length field. If set to `0`, PIN length is not written. `0` by default.
//   - [ITKSmartCardPINFormat.SetPINLengthBitSize]
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat
type ITKSmartCardPINFormat interface {
	objectivec.IObject

	// Topic: Configuring PIN Formatting

	// The format of PIN characters. [TKSmartCardPINCharsetNumeric] by default.
	Charset() TKSmartCardPINCharset
	SetCharset(value TKSmartCardPINCharset)
	// The encoding of PIN characters. [TKSmartCardPINEncodingASCII] by default.
	Encoding() TKSmartCardPINEncoding
	SetEncoding(value TKSmartCardPINEncoding)
	// The minimum number of characters to form a valid PIN. `4` by default.
	MinPINLength() int
	SetMinPINLength(value int)
	// The maximum number of characters to form a valid PIN. `8` by default.
	MaxPINLength() int
	SetMaxPINLength(value int)
	// The total length of the PIN block in bytes. `8` by default.
	PINBlockByteLength() int
	SetPINBlockByteLength(value int)
	// The justification within the PIN block. [TKSmartCardPINJustificationLeft] by default.
	PINJustification() TKSmartCardPINJustification
	SetPINJustification(value TKSmartCardPINJustification)
	// The offset, in bits, within the PIN block to mark a location for filling in the formatted PIN, which is justified with respect to the [pinJustification](<https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinJustification>) property value. `0` by default.
	PINBitOffset() int
	SetPINBitOffset(value int)
	// The offset, in bits, within the PIN block to mark a location for filling in the PIN length, which is always left justified. `0` by default.
	PINLengthBitOffset() int
	SetPINLengthBitOffset(value int)
	// The size, in bits, of the PIN length field. If set to `0`, PIN length is not written. `0` by default.
	PINLengthBitSize() int
	SetPINLengthBitSize(value int)
}

// Init initializes the instance.
func (t TKSmartCardPINFormat) Init() TKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TKSmartCardPINFormat) Autorelease() TKSmartCardPINFormat {
	rv := objc.Send[TKSmartCardPINFormat](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardPINFormat creates a new TKSmartCardPINFormat instance.
func NewTKSmartCardPINFormat() TKSmartCardPINFormat {
	class := getTKSmartCardPINFormatClass()
	rv := objc.Send[TKSmartCardPINFormat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The format of PIN characters. [TKSmartCardPINCharsetNumeric] by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/charset-swift.property
func (t TKSmartCardPINFormat) Charset() TKSmartCardPINCharset {
	rv := objc.Send[TKSmartCardPINCharset](t.ID, objc.Sel("charset"))
	return TKSmartCardPINCharset(rv)
}
func (t TKSmartCardPINFormat) SetCharset(value TKSmartCardPINCharset) {
	objc.Send[struct{}](t.ID, objc.Sel("setCharset:"), value)
}

// The encoding of PIN characters. [TKSmartCardPINEncodingASCII] by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/encoding-swift.property
func (t TKSmartCardPINFormat) Encoding() TKSmartCardPINEncoding {
	rv := objc.Send[TKSmartCardPINEncoding](t.ID, objc.Sel("encoding"))
	return TKSmartCardPINEncoding(rv)
}
func (t TKSmartCardPINFormat) SetEncoding(value TKSmartCardPINEncoding) {
	objc.Send[struct{}](t.ID, objc.Sel("setEncoding:"), value)
}

// The minimum number of characters to form a valid PIN. `4` by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/minPINLength
func (t TKSmartCardPINFormat) MinPINLength() int {
	rv := objc.Send[int](t.ID, objc.Sel("minPINLength"))
	return rv
}
func (t TKSmartCardPINFormat) SetMinPINLength(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setMinPINLength:"), value)
}

// The maximum number of characters to form a valid PIN. `8` by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/maxPINLength
func (t TKSmartCardPINFormat) MaxPINLength() int {
	rv := objc.Send[int](t.ID, objc.Sel("maxPINLength"))
	return rv
}
func (t TKSmartCardPINFormat) SetMaxPINLength(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaxPINLength:"), value)
}

// The total length of the PIN block in bytes. `8` by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinBlockByteLength
func (t TKSmartCardPINFormat) PINBlockByteLength() int {
	rv := objc.Send[int](t.ID, objc.Sel("PINBlockByteLength"))
	return rv
}
func (t TKSmartCardPINFormat) SetPINBlockByteLength(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINBlockByteLength:"), value)
}

// The justification within the PIN block. [TKSmartCardPINJustificationLeft]
// by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinJustification
func (t TKSmartCardPINFormat) PINJustification() TKSmartCardPINJustification {
	rv := objc.Send[TKSmartCardPINJustification](t.ID, objc.Sel("PINJustification"))
	return TKSmartCardPINJustification(rv)
}
func (t TKSmartCardPINFormat) SetPINJustification(value TKSmartCardPINJustification) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINJustification:"), value)
}

// The offset, in bits, within the PIN block to mark a location for filling in
// the formatted PIN, which is justified with respect to the
// [TKSmartCardPINFormat.PINJustification] property value. `0` by default.
//
// # Discussion
//
// The value of [PINBitOffset] indirectly controls the internal system units
// indicator. If [PINBitOffset] is byte aligned (that is, `PINBitOffset % 8 ==
// 0`), the internal representation of [PINBitOffset] gets converted from bits
// to bytes.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinBitOffset
func (t TKSmartCardPINFormat) PINBitOffset() int {
	rv := objc.Send[int](t.ID, objc.Sel("PINBitOffset"))
	return rv
}
func (t TKSmartCardPINFormat) SetPINBitOffset(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINBitOffset:"), value)
}

// The offset, in bits, within the PIN block to mark a location for filling in
// the PIN length, which is always left justified. `0` by default.
//
// # Discussion
//
// The value of [PINLengthBitOffset] indirectly controls the internal system
// units indicator. If [PINLengthBitOffset] is byte aligned (that is,
// `PINLengthBitOffset % 8 == 0`), the internal representation of
// [PINLengthBitOffset] gets converted from bits to bytes.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinLengthBitOffset
func (t TKSmartCardPINFormat) PINLengthBitOffset() int {
	rv := objc.Send[int](t.ID, objc.Sel("PINLengthBitOffset"))
	return rv
}
func (t TKSmartCardPINFormat) SetPINLengthBitOffset(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINLengthBitOffset:"), value)
}

// The size, in bits, of the PIN length field. If set to `0`, PIN length is
// not written. `0` by default.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardPINFormat/pinLengthBitSize
func (t TKSmartCardPINFormat) PINLengthBitSize() int {
	rv := objc.Send[int](t.ID, objc.Sel("PINLengthBitSize"))
	return rv
}
func (t TKSmartCardPINFormat) SetPINLengthBitSize(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setPINLengthBitSize:"), value)
}
