// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask
type CAAutoresizingMask uint32

const (
	// KCALayerHeightSizable: The receiver’s height is flexible.
	KCALayerHeightSizable CAAutoresizingMask = 16
	// KCALayerMaxXMargin: The right margin between the receiver and its superview is flexible.
	KCALayerMaxXMargin CAAutoresizingMask = 4
	// KCALayerMaxYMargin: The top margin between the receiver and its superview is flexible.
	KCALayerMaxYMargin CAAutoresizingMask = 32
	// KCALayerMinXMargin: The left margin between the receiver and its superview is flexible.
	KCALayerMinXMargin CAAutoresizingMask = 1
	// KCALayerMinYMargin: The bottom margin between the receiver and its superview is flexible.
	KCALayerMinYMargin CAAutoresizingMask = 8
	// KCALayerNotSizable: The receiver cannot be resized.
	KCALayerNotSizable CAAutoresizingMask = 0
	// KCALayerWidthSizable: The receiver’s width is flexible.
	KCALayerWidthSizable CAAutoresizingMask = 2
)

func (e CAAutoresizingMask) String() string {
	switch e {
	case KCALayerHeightSizable:
		return "KCALayerHeightSizable"
	case KCALayerMaxXMargin:
		return "KCALayerMaxXMargin"
	case KCALayerMaxYMargin:
		return "KCALayerMaxYMargin"
	case KCALayerMinXMargin:
		return "KCALayerMinXMargin"
	case KCALayerMinYMargin:
		return "KCALayerMinYMargin"
	case KCALayerNotSizable:
		return "KCALayerNotSizable"
	case KCALayerWidthSizable:
		return "KCALayerWidthSizable"
	default:
		return fmt.Sprintf("CAAutoresizingMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute
type CAConstraintAttribute int32

const (
	// KCAConstraintHeight: The height of a layer.
	KCAConstraintHeight CAConstraintAttribute = 7
	// KCAConstraintMaxX: The right edge of a layer’s frame.
	KCAConstraintMaxX CAConstraintAttribute = 2
	// KCAConstraintMaxY: The top edge of a layer’s frame.
	KCAConstraintMaxY CAConstraintAttribute = 6
	// KCAConstraintMidX: The horizontal location of the center of a layer’s frame.
	KCAConstraintMidX CAConstraintAttribute = 1
	// KCAConstraintMidY: The vertical location of the center of a layer’s frame.
	KCAConstraintMidY CAConstraintAttribute = 5
	// KCAConstraintMinX: The left edge of a layer’s frame.
	KCAConstraintMinX CAConstraintAttribute = 0
	// KCAConstraintMinY: The bottom edge of a layer’s frame.
	KCAConstraintMinY CAConstraintAttribute = 4
	// KCAConstraintWidth: The width of a layer.
	KCAConstraintWidth CAConstraintAttribute = 3
)

func (e CAConstraintAttribute) String() string {
	switch e {
	case KCAConstraintHeight:
		return "KCAConstraintHeight"
	case KCAConstraintMaxX:
		return "KCAConstraintMaxX"
	case KCAConstraintMaxY:
		return "KCAConstraintMaxY"
	case KCAConstraintMidX:
		return "KCAConstraintMidX"
	case KCAConstraintMidY:
		return "KCAConstraintMidY"
	case KCAConstraintMinX:
		return "KCAConstraintMinX"
	case KCAConstraintMinY:
		return "KCAConstraintMinY"
	case KCAConstraintWidth:
		return "KCAConstraintWidth"
	default:
		return fmt.Sprintf("CAConstraintAttribute(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/QuartzCore/CACornerMask
type CACornerMask uint

const (
	KCALayerMaxXMaxYCorner CACornerMask = 8
	KCALayerMaxXMinYCorner CACornerMask = 2
	KCALayerMinXMaxYCorner CACornerMask = 4
	KCALayerMinXMinYCorner CACornerMask = 1
)

func (e CACornerMask) String() string {
	switch e {
	case KCALayerMaxXMaxYCorner:
		return "KCALayerMaxXMaxYCorner"
	case KCALayerMaxXMinYCorner:
		return "KCALayerMaxXMinYCorner"
	case KCALayerMinXMaxYCorner:
		return "KCALayerMinXMaxYCorner"
	case KCALayerMinXMinYCorner:
		return "KCALayerMinXMinYCorner"
	default:
		return fmt.Sprintf("CACornerMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/QuartzCore/CAEdgeAntialiasingMask
type CAEdgeAntialiasingMask uint32

const (
	// KCALayerBottomEdge: # Discussion
	KCALayerBottomEdge CAEdgeAntialiasingMask = 4
	// KCALayerLeftEdge: # Discussion
	KCALayerLeftEdge CAEdgeAntialiasingMask = 1
	// KCALayerRightEdge: # Discussion
	KCALayerRightEdge CAEdgeAntialiasingMask = 2
	// KCALayerTopEdge: # Discussion
	KCALayerTopEdge CAEdgeAntialiasingMask = 8
)

func (e CAEdgeAntialiasingMask) String() string {
	switch e {
	case KCALayerBottomEdge:
		return "KCALayerBottomEdge"
	case KCALayerLeftEdge:
		return "KCALayerLeftEdge"
	case KCALayerRightEdge:
		return "KCALayerRightEdge"
	case KCALayerTopEdge:
		return "KCALayerTopEdge"
	default:
		return fmt.Sprintf("CAEdgeAntialiasingMask(%d)", e)
	}
}
