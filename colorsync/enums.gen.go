// Code generated from Apple documentation for ColorSync. DO NOT EDIT.

package colorsync

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/ColorSync/ColorSyncAlphaInfo
type ColorSyncAlphaInfo uint32

const (
	KColorSyncAlphaFirst              ColorSyncAlphaInfo = 4
	KColorSyncAlphaLast               ColorSyncAlphaInfo = 3
	KColorSyncAlphaNone               ColorSyncAlphaInfo = 0
	KColorSyncAlphaNoneSkipFirst      ColorSyncAlphaInfo = 6
	KColorSyncAlphaNoneSkipLast       ColorSyncAlphaInfo = 5
	KColorSyncAlphaPremultipliedFirst ColorSyncAlphaInfo = 2
	KColorSyncAlphaPremultipliedLast  ColorSyncAlphaInfo = 1
)

func (e ColorSyncAlphaInfo) String() string {
	switch e {
	case KColorSyncAlphaFirst:
		return "KColorSyncAlphaFirst"
	case KColorSyncAlphaLast:
		return "KColorSyncAlphaLast"
	case KColorSyncAlphaNone:
		return "KColorSyncAlphaNone"
	case KColorSyncAlphaNoneSkipFirst:
		return "KColorSyncAlphaNoneSkipFirst"
	case KColorSyncAlphaNoneSkipLast:
		return "KColorSyncAlphaNoneSkipLast"
	case KColorSyncAlphaPremultipliedFirst:
		return "KColorSyncAlphaPremultipliedFirst"
	case KColorSyncAlphaPremultipliedLast:
		return "KColorSyncAlphaPremultipliedLast"
	default:
		return fmt.Sprintf("ColorSyncAlphaInfo(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ColorSync/ColorSyncDataDepth
type ColorSyncDataDepth uint32

const (
	KColorSync10BitInteger         ColorSyncDataDepth = 8
	KColorSync16BitFloat           ColorSyncDataDepth = 4
	KColorSync16BitInteger         ColorSyncDataDepth = 3
	KColorSync1BitGamut            ColorSyncDataDepth = 1
	KColorSync32BitFloat           ColorSyncDataDepth = 7
	KColorSync32BitInteger         ColorSyncDataDepth = 5
	KColorSync32BitNamedColorIndex ColorSyncDataDepth = 6
	KColorSync8BitInteger          ColorSyncDataDepth = 2
)

func (e ColorSyncDataDepth) String() string {
	switch e {
	case KColorSync10BitInteger:
		return "KColorSync10BitInteger"
	case KColorSync16BitFloat:
		return "KColorSync16BitFloat"
	case KColorSync16BitInteger:
		return "KColorSync16BitInteger"
	case KColorSync1BitGamut:
		return "KColorSync1BitGamut"
	case KColorSync32BitFloat:
		return "KColorSync32BitFloat"
	case KColorSync32BitInteger:
		return "KColorSync32BitInteger"
	case KColorSync32BitNamedColorIndex:
		return "KColorSync32BitNamedColorIndex"
	case KColorSync8BitInteger:
		return "KColorSync8BitInteger"
	default:
		return fmt.Sprintf("ColorSyncDataDepth(%d)", e)
	}
}

type KColorSync uint32

const (
	KColorSyncAlphaInfoMask     KColorSync = 0x1f
	KColorSyncByteOrder16Big    KColorSync = 12288
	KColorSyncByteOrder16Little KColorSync = 4096
	KColorSyncByteOrder32Big    KColorSync = 16384
	KColorSyncByteOrder32Little KColorSync = 8192
	KColorSyncByteOrderDefault  KColorSync = 0
	KColorSyncByteOrderMask     KColorSync = 0x7000
)

func (e KColorSync) String() string {
	switch e {
	case KColorSyncAlphaInfoMask:
		return "KColorSyncAlphaInfoMask"
	case KColorSyncByteOrder16Big:
		return "KColorSyncByteOrder16Big"
	case KColorSyncByteOrder16Little:
		return "KColorSyncByteOrder16Little"
	case KColorSyncByteOrder32Big:
		return "KColorSyncByteOrder32Big"
	case KColorSyncByteOrder32Little:
		return "KColorSyncByteOrder32Little"
	case KColorSyncByteOrderDefault:
		return "KColorSyncByteOrderDefault"
	case KColorSyncByteOrderMask:
		return "KColorSyncByteOrderMask"
	default:
		return fmt.Sprintf("KColorSync(%d)", e)
	}
}
