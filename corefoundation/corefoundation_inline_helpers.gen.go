// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

import (
	"math"
	"math/bits"
	"unsafe"
)

// Byte order

func cfByteOrderGetCurrent() CFByteOrder {
	return CFByteOrderLittleEndian
}

// Byte swaps

func cfSwapInt16(v uint16) uint16 { return bits.ReverseBytes16(v) }
func cfSwapInt32(v uint32) uint32 { return bits.ReverseBytes32(v) }
func cfSwapInt64(v uint64) uint64 { return bits.ReverseBytes64(v) }

func cfSwapInt16BigToHost(v uint16) uint16    { return cfSwapInt16(v) }
func cfSwapInt16HostToBig(v uint16) uint16    { return cfSwapInt16(v) }
func cfSwapInt16HostToLittle(v uint16) uint16 { return v }
func cfSwapInt16LittleToHost(v uint16) uint16 { return v }

func cfSwapInt32BigToHost(v uint32) uint32    { return cfSwapInt32(v) }
func cfSwapInt32HostToBig(v uint32) uint32    { return cfSwapInt32(v) }
func cfSwapInt32HostToLittle(v uint32) uint32 { return v }
func cfSwapInt32LittleToHost(v uint32) uint32 { return v }

func cfSwapInt64BigToHost(v uint64) uint64    { return cfSwapInt64(v) }
func cfSwapInt64HostToBig(v uint64) uint64    { return cfSwapInt64(v) }
func cfSwapInt64HostToLittle(v uint64) uint64 { return v }
func cfSwapInt64LittleToHost(v uint64) uint64 { return v }

// Float conversions

func cfConvertFloatHostToSwapped(v float32) CFSwappedFloat32 {
	return CFSwappedFloat32{V: cfSwapInt32HostToBig(math.Float32bits(v))}
}

func cfConvertFloatSwappedToHost(v CFSwappedFloat32) float32 {
	return math.Float32frombits(cfSwapInt32BigToHost(v.V))
}

func cfConvertFloat32HostToSwapped(v float32) CFSwappedFloat32 {
	return cfConvertFloatHostToSwapped(v)
}

func cfConvertFloat32SwappedToHost(v CFSwappedFloat32) float32 {
	return cfConvertFloatSwappedToHost(v)
}

func cfConvertDoubleHostToSwapped(v float64) CFSwappedFloat64 {
	return CFSwappedFloat64{V: cfSwapInt64HostToBig(math.Float64bits(v))}
}

func cfConvertDoubleSwappedToHost(v CFSwappedFloat64) float64 {
	return math.Float64frombits(cfSwapInt64BigToHost(v.V))
}

func cfConvertFloat64HostToSwapped(v float64) CFSwappedFloat64 {
	return cfConvertDoubleHostToSwapped(v)
}

func cfConvertFloat64SwappedToHost(v CFSwappedFloat64) float64 {
	return cfConvertDoubleSwappedToHost(v)
}

// Range

func cfRangeMake(loc int, length int) CFRange {
	return CFRange{Location: loc, Length: length}
}

// String inline buffer

func cfStringInitInlineBuffer(str CFStringRef, buf *CFStringInlineBuffer, range_ CFRange) {
	if buf == nil {
		return
	}
	buf.TheString = str
	buf.RangeToBuffer = range_
	buf.BufferedRangeStart = 0
	buf.BufferedRangeEnd = 0
	buf.DirectCStringBuffer = nil
	buf.DirectUniCharBuffer = nil
}

func cfStringGetCharacterFromInlineBuffer(buf *CFStringInlineBuffer, idx int) uint16 {
	if buf == nil {
		return 0
	}
	if idx < 0 || idx >= buf.RangeToBuffer.Length {
		return 0
	}
	return 0
}

func cfStringIsSurrogateHighCharacter(character uint16) bool {
	return character >= 0xD800 && character <= 0xDBFF
}

func cfStringIsSurrogateLowCharacter(character uint16) bool {
	return character >= 0xDC00 && character <= 0xDFFF
}

func cfStringGetLongCharacterForSurrogatePair(surrogateHigh uint16, surrogateLow uint16) uint32 {
	return (uint32(surrogateHigh)-0xD800)<<10 + (uint32(surrogateLow) - 0xDC00) + 0x10000
}

func cfStringGetSurrogatePairForLongCharacter(character uint32, surrogates *byte) bool {
	if character > 0x10FFFF || (character >= 0xD800 && character <= 0xDFFF) {
		return false
	}
	if character < 0x10000 {
		return false
	}
	if surrogates == nil {
		return false
	}
	p := (*[2]uint16)(unsafe.Pointer(surrogates))
	p[0] = uint16((character-0x10000)>>10) + 0xD800
	p[1] = uint16((character-0x10000)&0x3FF) + 0xDC00
	return true
}

// User notification

func cfUserNotificationCheckBoxChecked(i CFIndex) CFOptionFlags {
	return 1 << uint(8+i)
}

func cfUserNotificationSecureTextField(i CFIndex) CFOptionFlags {
	return 1 << uint(16+i)
}

func cfUserNotificationPopUpSelection(n CFIndex) CFOptionFlags {
	return CFOptionFlags(n) << 24
}

