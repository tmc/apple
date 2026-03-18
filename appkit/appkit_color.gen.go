// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "image/color"

// Compile-time interface assertion.
var _ color.Color = NSColor{}

// RGBA implements the [image/color.Color] interface.
// The color is converted to sRGB color space before extracting components,
// since calling component accessors on non-RGB colors (grayscale, CMYK)
// raises NSInternalInconsistencyException.
func (n NSColor) RGBA() (r, g, b, a uint32) {
	c := n.ColorUsingColorSpace(GetNSColorSpaceClass().GenericRGBColorSpace())
	r = uint32(c.RedComponent() * 0xffff)
	g = uint32(c.GreenComponent() * 0xffff)
	b = uint32(c.BlueComponent() * 0xffff)
	a = uint32(c.AlphaComponent() * 0xffff)
	return
}

