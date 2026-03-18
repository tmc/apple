// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentName
type IOSurfaceComponentName int

const (
	KIOSurfaceComponentNameAlpha IOSurfaceComponentName = 1
	KIOSurfaceComponentNameBlue IOSurfaceComponentName = 4
	KIOSurfaceComponentNameChromaBlue IOSurfaceComponentName = 7
	KIOSurfaceComponentNameChromaRed IOSurfaceComponentName = 6
	KIOSurfaceComponentNameGreen IOSurfaceComponentName = 3
	KIOSurfaceComponentNameLuma IOSurfaceComponentName = 5
	KIOSurfaceComponentNameRed IOSurfaceComponentName = 2
	KIOSurfaceComponentNameUnknown IOSurfaceComponentName = 0
)

func (e IOSurfaceComponentName) String() string {
	switch e {
	case KIOSurfaceComponentNameAlpha:
		return "KIOSurfaceComponentNameAlpha"
	case KIOSurfaceComponentNameBlue:
		return "KIOSurfaceComponentNameBlue"
	case KIOSurfaceComponentNameChromaBlue:
		return "KIOSurfaceComponentNameChromaBlue"
	case KIOSurfaceComponentNameChromaRed:
		return "KIOSurfaceComponentNameChromaRed"
	case KIOSurfaceComponentNameGreen:
		return "KIOSurfaceComponentNameGreen"
	case KIOSurfaceComponentNameLuma:
		return "KIOSurfaceComponentNameLuma"
	case KIOSurfaceComponentNameRed:
		return "KIOSurfaceComponentNameRed"
	case KIOSurfaceComponentNameUnknown:
		return "KIOSurfaceComponentNameUnknown"
	default:
		return fmt.Sprintf("IOSurfaceComponentName(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentRange
type IOSurfaceComponentRange int

const (
	KIOSurfaceComponentRangeFullRange IOSurfaceComponentRange = 1
	KIOSurfaceComponentRangeUnknown IOSurfaceComponentRange = 0
	KIOSurfaceComponentRangeVideoRange IOSurfaceComponentRange = 2
	KIOSurfaceComponentRangeWideRange IOSurfaceComponentRange = 3
)

func (e IOSurfaceComponentRange) String() string {
	switch e {
	case KIOSurfaceComponentRangeFullRange:
		return "KIOSurfaceComponentRangeFullRange"
	case KIOSurfaceComponentRangeUnknown:
		return "KIOSurfaceComponentRangeUnknown"
	case KIOSurfaceComponentRangeVideoRange:
		return "KIOSurfaceComponentRangeVideoRange"
	case KIOSurfaceComponentRangeWideRange:
		return "KIOSurfaceComponentRangeWideRange"
	default:
		return fmt.Sprintf("IOSurfaceComponentRange(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceComponentType
type IOSurfaceComponentType int

const (
	KIOSurfaceComponentTypeFloat IOSurfaceComponentType = 3
	KIOSurfaceComponentTypeSignedInteger IOSurfaceComponentType = 2
	KIOSurfaceComponentTypeSignedNormalized IOSurfaceComponentType = 4
	KIOSurfaceComponentTypeUnknown IOSurfaceComponentType = 0
	KIOSurfaceComponentTypeUnsignedInteger IOSurfaceComponentType = 1
)

func (e IOSurfaceComponentType) String() string {
	switch e {
	case KIOSurfaceComponentTypeFloat:
		return "KIOSurfaceComponentTypeFloat"
	case KIOSurfaceComponentTypeSignedInteger:
		return "KIOSurfaceComponentTypeSignedInteger"
	case KIOSurfaceComponentTypeSignedNormalized:
		return "KIOSurfaceComponentTypeSignedNormalized"
	case KIOSurfaceComponentTypeUnknown:
		return "KIOSurfaceComponentTypeUnknown"
	case KIOSurfaceComponentTypeUnsignedInteger:
		return "KIOSurfaceComponentTypeUnsignedInteger"
	default:
		return fmt.Sprintf("IOSurfaceComponentType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceLockOptions
type IOSurfaceLockOptions int

const (
	// KIOSurfaceLockAvoidSync: # Discussion
	KIOSurfaceLockAvoidSync IOSurfaceLockOptions = 2
	// KIOSurfaceLockReadOnly: # Discussion
	KIOSurfaceLockReadOnly IOSurfaceLockOptions = 1
)

func (e IOSurfaceLockOptions) String() string {
	switch e {
	case KIOSurfaceLockAvoidSync:
		return "KIOSurfaceLockAvoidSync"
	case KIOSurfaceLockReadOnly:
		return "KIOSurfaceLockReadOnly"
	default:
		return fmt.Sprintf("IOSurfaceLockOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerFlags
type IOSurfaceMemoryLedgerFlags int

const (
	KIOSurfaceMemoryLedgerFlagNoFootprint IOSurfaceMemoryLedgerFlags = 1
)

func (e IOSurfaceMemoryLedgerFlags) String() string {
	switch e {
	case KIOSurfaceMemoryLedgerFlagNoFootprint:
		return "KIOSurfaceMemoryLedgerFlagNoFootprint"
	default:
		return fmt.Sprintf("IOSurfaceMemoryLedgerFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceMemoryLedgerTags
type IOSurfaceMemoryLedgerTags int

const (
	KIOSurfaceMemoryLedgerTagDefault IOSurfaceMemoryLedgerTags = 1
	KIOSurfaceMemoryLedgerTagGraphics IOSurfaceMemoryLedgerTags = 4
	KIOSurfaceMemoryLedgerTagMedia IOSurfaceMemoryLedgerTags = 3
	KIOSurfaceMemoryLedgerTagNetwork IOSurfaceMemoryLedgerTags = 2
	KIOSurfaceMemoryLedgerTagNeural IOSurfaceMemoryLedgerTags = 5
)

func (e IOSurfaceMemoryLedgerTags) String() string {
	switch e {
	case KIOSurfaceMemoryLedgerTagDefault:
		return "KIOSurfaceMemoryLedgerTagDefault"
	case KIOSurfaceMemoryLedgerTagGraphics:
		return "KIOSurfaceMemoryLedgerTagGraphics"
	case KIOSurfaceMemoryLedgerTagMedia:
		return "KIOSurfaceMemoryLedgerTagMedia"
	case KIOSurfaceMemoryLedgerTagNetwork:
		return "KIOSurfaceMemoryLedgerTagNetwork"
	case KIOSurfaceMemoryLedgerTagNeural:
		return "KIOSurfaceMemoryLedgerTagNeural"
	default:
		return fmt.Sprintf("IOSurfaceMemoryLedgerTags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurfacePurgeabilityState
type IOSurfacePurgeabilityState int

const (
	KIOSurfacePurgeableEmpty IOSurfacePurgeabilityState = 2
	KIOSurfacePurgeableKeepCurrent IOSurfacePurgeabilityState = 3
	KIOSurfacePurgeableNonVolatile IOSurfacePurgeabilityState = 0
	KIOSurfacePurgeableVolatile IOSurfacePurgeabilityState = 1
)

func (e IOSurfacePurgeabilityState) String() string {
	switch e {
	case KIOSurfacePurgeableEmpty:
		return "KIOSurfacePurgeableEmpty"
	case KIOSurfacePurgeableKeepCurrent:
		return "KIOSurfacePurgeableKeepCurrent"
	case KIOSurfacePurgeableNonVolatile:
		return "KIOSurfacePurgeableNonVolatile"
	case KIOSurfacePurgeableVolatile:
		return "KIOSurfacePurgeableVolatile"
	default:
		return fmt.Sprintf("IOSurfacePurgeabilityState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurfaceSubsampling
type IOSurfaceSubsampling int

const (
	KIOSurfaceSubsampling411 IOSurfaceSubsampling = 4
	KIOSurfaceSubsampling420 IOSurfaceSubsampling = 3
	KIOSurfaceSubsampling422 IOSurfaceSubsampling = 2
	KIOSurfaceSubsamplingNone IOSurfaceSubsampling = 1
	KIOSurfaceSubsamplingUnknown IOSurfaceSubsampling = 0
)

func (e IOSurfaceSubsampling) String() string {
	switch e {
	case KIOSurfaceSubsampling411:
		return "KIOSurfaceSubsampling411"
	case KIOSurfaceSubsampling420:
		return "KIOSurfaceSubsampling420"
	case KIOSurfaceSubsampling422:
		return "KIOSurfaceSubsampling422"
	case KIOSurfaceSubsamplingNone:
		return "KIOSurfaceSubsamplingNone"
	case KIOSurfaceSubsamplingUnknown:
		return "KIOSurfaceSubsamplingUnknown"
	default:
		return fmt.Sprintf("IOSurfaceSubsampling(%d)", e)
	}
}

type KIOSurface uint

const (
	KIOSurfaceCopybackCache KIOSurface = 3
	KIOSurfaceCopybackInnerCache KIOSurface = 5
	KIOSurfaceDefaultCache KIOSurface = 0
	KIOSurfaceInhibitCache KIOSurface = 1
	KIOSurfaceWriteCombineCache KIOSurface = 4
	KIOSurfaceWriteThruCache KIOSurface = 2
)

func (e KIOSurface) String() string {
	switch e {
	case KIOSurfaceCopybackCache:
		return "KIOSurfaceCopybackCache"
	case KIOSurfaceCopybackInnerCache:
		return "KIOSurfaceCopybackInnerCache"
	case KIOSurfaceDefaultCache:
		return "KIOSurfaceDefaultCache"
	case KIOSurfaceInhibitCache:
		return "KIOSurfaceInhibitCache"
	case KIOSurfaceWriteCombineCache:
		return "KIOSurfaceWriteCombineCache"
	case KIOSurfaceWriteThruCache:
		return "KIOSurfaceWriteThruCache"
	default:
		return fmt.Sprintf("KIOSurface(%d)", e)
	}
}

type KIOSurfaceMap uint

const (
	KIOSurfaceMapCacheShift KIOSurfaceMap = 8
	KIOSurfaceMapCopybackCache KIOSurfaceMap = 3
	KIOSurfaceMapCopybackInnerCache KIOSurfaceMap = 5
	KIOSurfaceMapDefaultCache KIOSurfaceMap = 0
	KIOSurfaceMapInhibitCache KIOSurfaceMap = 1
	KIOSurfaceMapWriteCombineCache KIOSurfaceMap = 4
	KIOSurfaceMapWriteThruCache KIOSurfaceMap = 2
)

func (e KIOSurfaceMap) String() string {
	switch e {
	case KIOSurfaceMapCacheShift:
		return "KIOSurfaceMapCacheShift"
	case KIOSurfaceMapCopybackCache:
		return "KIOSurfaceMapCopybackCache"
	case KIOSurfaceMapCopybackInnerCache:
		return "KIOSurfaceMapCopybackInnerCache"
	case KIOSurfaceMapDefaultCache:
		return "KIOSurfaceMapDefaultCache"
	case KIOSurfaceMapInhibitCache:
		return "KIOSurfaceMapInhibitCache"
	case KIOSurfaceMapWriteCombineCache:
		return "KIOSurfaceMapWriteCombineCache"
	case KIOSurfaceMapWriteThruCache:
		return "KIOSurfaceMapWriteThruCache"
	default:
		return fmt.Sprintf("KIOSurfaceMap(%d)", e)
	}
}

