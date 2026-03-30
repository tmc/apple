// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"math"
	"math/bits"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CF bridging (pointer casts in purego — no ARC)

func cfBridgingRelease(X corefoundation.CFTypeRef) objectivec.Object {
	return objectivec.Object{ID: objc.ID(X)}
}

func cfBridgingRetain(X objectivec.Object) corefoundation.CFTypeRef {
	return corefoundation.CFTypeRef(X.ID)
}

func nsMakeCollectable(cf corefoundation.CFTypeRef) objectivec.Object {
	return objectivec.Object{ID: objc.ID(cf)}
}

// Decimal

func nsDecimalIsNotANumber(dcm Decimal) bool {
	// NSDecimalIsNotANumber checks internal fields not exposed in Go bindings.
	// Return false (not NaN) as a safe default.
	return false
}

// Geometry constructors

func nsMakePoint(x, y float64) NSPoint {
	return NSPoint{X: x, Y: y}
}

func nsMakeSize(width, height float64) NSSize {
	return NSSize{Width: width, Height: height}
}

func nsMakeRect(x, y, w, h float64) NSRect {
	return NSRect{
		Origin: corefoundation.CGPoint{X: x, Y: y},
		Size:   corefoundation.CGSize{Width: w, Height: h},
	}
}

// Range constructors

func nsMakeRange(loc, len uint) NSRange {
	return NSRange{Location: loc, Length: len}
}

// Geometry queries

func nsMaxX(aRect NSRect) float64   { return aRect.Origin.X + aRect.Size.Width }
func nsMaxY(aRect NSRect) float64   { return aRect.Origin.Y + aRect.Size.Height }
func nsMidX(aRect NSRect) float64   { return aRect.Origin.X + aRect.Size.Width/2 }
func nsMidY(aRect NSRect) float64   { return aRect.Origin.Y + aRect.Size.Height/2 }
func nsMinX(aRect NSRect) float64   { return aRect.Origin.X }
func nsMinY(aRect NSRect) float64   { return aRect.Origin.Y }
func nsWidth(aRect NSRect) float64  { return aRect.Size.Width }
func nsHeight(aRect NSRect) float64 { return aRect.Size.Height }

// Range queries

func nsMaxRange(range_ NSRange) uint { return range_.Location + range_.Length }
func nsLocationInRange(loc uint, range_ NSRange) bool {
	return loc >= range_.Location && (loc-range_.Location) < range_.Length
}

// Equality

func nsEqualPoints(aPoint, bPoint NSPoint) bool {
	return aPoint.X == bPoint.X && aPoint.Y == bPoint.Y
}

func nsEqualSizes(aSize, bSize NSSize) bool {
	return aSize.Width == bSize.Width && aSize.Height == bSize.Height
}

func nsEqualRects(aRect, bRect NSRect) bool {
	return nsEqualPoints(aRect.Origin, bRect.Origin) && nsEqualSizes(aRect.Size, bRect.Size)
}

func nsEqualRanges(range1, range2 NSRange) bool {
	return range1.Location == range2.Location && range1.Length == range2.Length
}

func nsEdgeInsetsEqual(aInsets, bInsets NSEdgeInsets) bool {
	return aInsets.Top == bInsets.Top && aInsets.Left == bInsets.Left &&
		aInsets.Bottom == bInsets.Bottom && aInsets.Right == bInsets.Right
}

// Rect operations

func nsIsEmptyRect(aRect NSRect) bool {
	return aRect.Size.Width <= 0 || aRect.Size.Height <= 0
}

func nsInsetRect(aRect NSRect, dX, dY float64) NSRect {
	return NSRect{
		Origin: corefoundation.CGPoint{X: aRect.Origin.X + dX, Y: aRect.Origin.Y + dY},
		Size:   corefoundation.CGSize{Width: aRect.Size.Width - 2*dX, Height: aRect.Size.Height - 2*dY},
	}
}

func nsOffsetRect(aRect NSRect, dX, dY float64) NSRect {
	return NSRect{
		Origin: corefoundation.CGPoint{X: aRect.Origin.X + dX, Y: aRect.Origin.Y + dY},
		Size:   aRect.Size,
	}
}

func nsIntegralRect(aRect NSRect) NSRect {
	if nsIsEmptyRect(aRect) {
		return NSRect{}
	}
	x := math.Floor(aRect.Origin.X)
	y := math.Floor(aRect.Origin.Y)
	maxX := math.Ceil(aRect.Origin.X + aRect.Size.Width)
	maxY := math.Ceil(aRect.Origin.Y + aRect.Size.Height)
	return NSRect{
		Origin: corefoundation.CGPoint{X: x, Y: y},
		Size:   corefoundation.CGSize{Width: maxX - x, Height: maxY - y},
	}
}

func nsUnionRect(aRect, bRect NSRect) NSRect {
	if nsIsEmptyRect(aRect) {
		return bRect
	}
	if nsIsEmptyRect(bRect) {
		return aRect
	}
	x := math.Min(nsMinX(aRect), nsMinX(bRect))
	y := math.Min(nsMinY(aRect), nsMinY(bRect))
	maxX := math.Max(nsMaxX(aRect), nsMaxX(bRect))
	maxY := math.Max(nsMaxY(aRect), nsMaxY(bRect))
	return NSRect{
		Origin: corefoundation.CGPoint{X: x, Y: y},
		Size:   corefoundation.CGSize{Width: maxX - x, Height: maxY - y},
	}
}

func nsIntersectionRect(aRect, bRect NSRect) NSRect {
	if nsIsEmptyRect(aRect) || nsIsEmptyRect(bRect) {
		return NSRect{}
	}
	x := math.Max(nsMinX(aRect), nsMinX(bRect))
	y := math.Max(nsMinY(aRect), nsMinY(bRect))
	maxX := math.Min(nsMaxX(aRect), nsMaxX(bRect))
	maxY := math.Min(nsMaxY(aRect), nsMaxY(bRect))
	if maxX <= x || maxY <= y {
		return NSRect{}
	}
	return NSRect{
		Origin: corefoundation.CGPoint{X: x, Y: y},
		Size:   corefoundation.CGSize{Width: maxX - x, Height: maxY - y},
	}
}

func nsContainsRect(aRect, bRect NSRect) bool {
	if nsIsEmptyRect(bRect) {
		return false
	}
	return nsMinX(aRect) <= nsMinX(bRect) && nsMaxX(aRect) >= nsMaxX(bRect) &&
		nsMinY(aRect) <= nsMinY(bRect) && nsMaxY(aRect) >= nsMaxY(bRect)
}

func nsIntersectsRect(aRect, bRect NSRect) bool {
	if nsIsEmptyRect(aRect) || nsIsEmptyRect(bRect) {
		return false
	}
	return nsMaxX(aRect) > nsMinX(bRect) && nsMinX(aRect) < nsMaxX(bRect) &&
		nsMaxY(aRect) > nsMinY(bRect) && nsMinY(aRect) < nsMaxY(bRect)
}

func nsPointInRect(aPoint NSPoint, aRect NSRect) bool {
	return aPoint.X >= nsMinX(aRect) && aPoint.X < nsMaxX(aRect) &&
		aPoint.Y >= nsMinY(aRect) && aPoint.Y < nsMaxY(aRect)
}

func nsMouseInRect(aPoint NSPoint, aRect NSRect, flipped bool) bool {
	if flipped {
		return aPoint.X >= nsMinX(aRect) && aPoint.X < nsMaxX(aRect) &&
			aPoint.Y >= nsMinY(aRect) && aPoint.Y < nsMaxY(aRect)
	}
	return aPoint.X >= nsMinX(aRect) && aPoint.X < nsMaxX(aRect) &&
		aPoint.Y > nsMinY(aRect) && aPoint.Y <= nsMaxY(aRect)
}

func nsDivideRect(inRect NSRect, slice, rem *NSRect, amount float64, edge NSRectEdge) {
	if slice == nil || rem == nil {
		return
	}
	if nsIsEmptyRect(inRect) {
		*slice = NSRect{}
		*rem = NSRect{}
		return
	}
	switch edge {
	case NSMinXEdge:
		if amount > inRect.Size.Width {
			amount = inRect.Size.Width
		}
		*slice = NSRect{
			Origin: inRect.Origin,
			Size:   corefoundation.CGSize{Width: amount, Height: inRect.Size.Height},
		}
		*rem = NSRect{
			Origin: corefoundation.CGPoint{X: inRect.Origin.X + amount, Y: inRect.Origin.Y},
			Size:   corefoundation.CGSize{Width: inRect.Size.Width - amount, Height: inRect.Size.Height},
		}
	case NSMinYEdge:
		if amount > inRect.Size.Height {
			amount = inRect.Size.Height
		}
		*slice = NSRect{
			Origin: inRect.Origin,
			Size:   corefoundation.CGSize{Width: inRect.Size.Width, Height: amount},
		}
		*rem = NSRect{
			Origin: corefoundation.CGPoint{X: inRect.Origin.X, Y: inRect.Origin.Y + amount},
			Size:   corefoundation.CGSize{Width: inRect.Size.Width, Height: inRect.Size.Height - amount},
		}
	case NSMaxXEdge:
		if amount > inRect.Size.Width {
			amount = inRect.Size.Width
		}
		*slice = NSRect{
			Origin: corefoundation.CGPoint{X: inRect.Origin.X + inRect.Size.Width - amount, Y: inRect.Origin.Y},
			Size:   corefoundation.CGSize{Width: amount, Height: inRect.Size.Height},
		}
		*rem = NSRect{
			Origin: inRect.Origin,
			Size:   corefoundation.CGSize{Width: inRect.Size.Width - amount, Height: inRect.Size.Height},
		}
	case NSMaxYEdge:
		if amount > inRect.Size.Height {
			amount = inRect.Size.Height
		}
		*slice = NSRect{
			Origin: corefoundation.CGPoint{X: inRect.Origin.X, Y: inRect.Origin.Y + inRect.Size.Height - amount},
			Size:   corefoundation.CGSize{Width: inRect.Size.Width, Height: amount},
		}
		*rem = NSRect{
			Origin: inRect.Origin,
			Size:   corefoundation.CGSize{Width: inRect.Size.Width, Height: inRect.Size.Height - amount},
		}
	}
}

func nsIntegralRectWithOptions(aRect NSRect, opts AlignmentOptions) NSRect {
	// Simplified: use standard integral rect behavior
	return nsIntegralRect(aRect)
}

func nsEdgeInsetsMake(top, left, bottom, right float64) NSEdgeInsets {
	return NSEdgeInsets{Top: top, Left: left, Bottom: bottom, Right: right}
}

// Range operations

func nsUnionRange(range1, range2 NSRange) NSRange {
	var start, end uint
	if range1.Location < range2.Location {
		start = range1.Location
	} else {
		start = range2.Location
	}
	end1 := range1.Location + range1.Length
	end2 := range2.Location + range2.Length
	if end1 > end2 {
		end = end1
	} else {
		end = end2
	}
	return NSRange{Location: start, Length: end - start}
}

func nsIntersectionRange(range1, range2 NSRange) NSRange {
	var start, end1, end2 uint
	if range1.Location > range2.Location {
		start = range1.Location
	} else {
		start = range2.Location
	}
	end1 = range1.Location + range1.Length
	end2 = range2.Location + range2.Length
	var end uint
	if end1 < end2 {
		end = end1
	} else {
		end = end2
	}
	if end <= start {
		return NSRange{Location: 0, Length: 0}
	}
	return NSRange{Location: start, Length: end - start}
}

// Type conversions (NSPoint/NSRect/NSSize ↔ CGPoint/CGRect/CGSize)

func nsPointFromCGPoint(point corefoundation.CGPoint) NSPoint { return point }
func nsPointToCGPoint(point NSPoint) corefoundation.CGPoint   { return point }
func nsSizeFromCGSize(size corefoundation.CGSize) NSSize      { return size }
func nsSizeToCGSize(size NSSize) corefoundation.CGSize        { return size }
func nsRectFromCGRect(rect corefoundation.CGRect) NSRect      { return rect }
func nsRectToCGRect(rect NSRect) corefoundation.CGRect        { return rect }

// Byte order

func nsHostByteOrder() int {
	return 1 // NS_LittleEndian on ARM64 and x86_64
}

// Integer byte swapping

func nsSwapShort(inv uint16) uint16    { return bits.ReverseBytes16(inv) }
func nsSwapInt(inv uint) uint          { return uint(bits.ReverseBytes32(uint32(inv))) }
func nsSwapLong(inv uint) uint         { return uint(bits.ReverseBytes64(uint64(inv))) }
func nsSwapLongLong(inv uint64) uint64 { return bits.ReverseBytes64(inv) }

func nsSwapBigShortToHost(x uint16) uint16    { return nsSwapShort(x) }
func nsSwapBigIntToHost(x uint) uint          { return nsSwapInt(x) }
func nsSwapBigLongToHost(x uint) uint         { return nsSwapLong(x) }
func nsSwapBigLongLongToHost(x uint64) uint64 { return nsSwapLongLong(x) }

func nsSwapHostShortToBig(x uint16) uint16    { return nsSwapShort(x) }
func nsSwapHostIntToBig(x uint) uint          { return nsSwapInt(x) }
func nsSwapHostLongToBig(x uint) uint         { return nsSwapLong(x) }
func nsSwapHostLongLongToBig(x uint64) uint64 { return nsSwapLongLong(x) }

func nsSwapLittleShortToHost(x uint16) uint16    { return x }
func nsSwapLittleIntToHost(x uint) uint          { return x }
func nsSwapLittleLongToHost(x uint) uint         { return x }
func nsSwapLittleLongLongToHost(x uint64) uint64 { return x }

func nsSwapHostShortToLittle(x uint16) uint16    { return x }
func nsSwapHostIntToLittle(x uint) uint          { return x }
func nsSwapHostLongToLittle(x uint) uint         { return x }
func nsSwapHostLongLongToLittle(x uint64) uint64 { return x }

// Float byte swapping

func nsSwapFloat(x NSSwappedFloat) NSSwappedFloat {
	return NSSwappedFloat{V: uint(nsSwapInt(uint(x.V)))}
}

func nsSwapDouble(x NSSwappedDouble) NSSwappedDouble {
	return NSSwappedDouble{V: nsSwapLongLong(x.V)}
}

// Float conversions

func nsConvertHostFloatToSwapped(x float32) NSSwappedFloat {
	return NSSwappedFloat{V: uint(nsSwapHostIntToBig(uint(math.Float32bits(x))))}
}

func nsConvertSwappedFloatToHost(x NSSwappedFloat) float32 {
	return math.Float32frombits(uint32(nsSwapBigIntToHost(uint(x.V))))
}

func nsConvertHostDoubleToSwapped(x float64) NSSwappedDouble {
	return NSSwappedDouble{V: nsSwapHostLongLongToBig(math.Float64bits(x))}
}

func nsConvertSwappedDoubleToHost(x NSSwappedDouble) float64 {
	return math.Float64frombits(nsSwapBigLongLongToHost(x.V))
}

func nsSwapBigFloatToHost(x NSSwappedFloat) float32 {
	return nsConvertSwappedFloatToHost(x)
}

func nsSwapBigDoubleToHost(x NSSwappedDouble) float64 {
	return nsConvertSwappedDoubleToHost(x)
}

func nsSwapHostFloatToBig(x float32) NSSwappedFloat {
	return nsConvertHostFloatToSwapped(x)
}

func nsSwapHostDoubleToBig(x float64) NSSwappedDouble {
	return nsConvertHostDoubleToSwapped(x)
}

func nsSwapLittleFloatToHost(x NSSwappedFloat) float32 {
	return math.Float32frombits(uint32(x.V))
}

func nsSwapLittleDoubleToHost(x NSSwappedDouble) float64 {
	return math.Float64frombits(x.V)
}

func nsSwapHostFloatToLittle(x float32) NSSwappedFloat {
	return NSSwappedFloat{V: uint(math.Float32bits(x))}
}

func nsSwapHostDoubleToLittle(x float64) NSSwappedDouble {
	return NSSwappedDouble{V: math.Float64bits(x)}
}
