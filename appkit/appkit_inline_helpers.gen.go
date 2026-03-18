// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// Geometry constructors

func nsDirectionalEdgeInsetsMake(top, leading, bottom, trailing float64) NSDirectionalEdgeInsets {
	return NSDirectionalEdgeInsets{Top: top, Leading: leading, Bottom: bottom, Trailing: trailing}
}

// Event mask constructors

func nsEventMaskFromType(type_ NSEventType) NSEventMask {
	return 1 << NSEventMask(type_)
}

func nsTouchTypeMaskFromType(type_ NSTouchType) NSTouchTypeMask {
	return 1 << NSTouchTypeMask(type_)
}

