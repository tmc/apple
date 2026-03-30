// Code generated from Apple documentation for iokit. DO NOT EDIT.

package iokit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/iokit/iourlerror
type IOURLError int32

const (
	KIOURLImproperArgumentsError       IOURLError = 0
	KIOURLPropertyKeyUnavailableError  IOURLError = 0
	KIOURLRemoteHostUnavailableError   IOURLError = 0
	KIOURLResourceAccessViolationError IOURLError = 0
	KIOURLResourceNotFoundError        IOURLError = 0
	KIOURLTimeoutError                 IOURLError = 0
	KIOURLUnknownError                 IOURLError = 0
	KIOURLUnknownPropertyKeyError      IOURLError = 0
	KIOURLUnknownSchemeError           IOURLError = 0
)

func (e IOURLError) String() string {
	switch e {
	case KIOURLImproperArgumentsError:
		return "KIOURLImproperArgumentsError"
	default:
		return fmt.Sprintf("IOURLError(%d)", e)
	}
}

type KDisplayVendorID uint

const (
	KDisplayVendorIDUnknown KDisplayVendorID = 0
)

func (e KDisplayVendorID) String() string {
	switch e {
	case KDisplayVendorIDUnknown:
		return "KDisplayVendorIDUnknown"
	default:
		return fmt.Sprintf("KDisplayVendorID(%d)", e)
	}
}

type KFirstIOKitNotificationType uint

const (
	KFirstIOKitNotificationTypeValue     KFirstIOKitNotificationType = 100
	KIOAsyncCompletionNotificationType   KFirstIOKitNotificationType = 150
	KIOKitNoticationMsgSizeMask          KFirstIOKitNotificationType = 3
	KIOKitNoticationTypeMask             KFirstIOKitNotificationType = 4095
	KIOKitNoticationTypeSizeAdjShift     KFirstIOKitNotificationType = 30
	KIOServiceMatchedNotificationType    KFirstIOKitNotificationType = 101
	KIOServiceMessageNotificationType    KFirstIOKitNotificationType = 160
	KIOServicePublishNotificationType    KFirstIOKitNotificationType = 100
	KIOServiceTerminatedNotificationType KFirstIOKitNotificationType = 102
	KLastIOKitNotificationType           KFirstIOKitNotificationType = 199
)

func (e KFirstIOKitNotificationType) String() string {
	switch e {
	case KFirstIOKitNotificationTypeValue:
		return "KFirstIOKitNotificationTypeValue"
	case KIOAsyncCompletionNotificationType:
		return "KIOAsyncCompletionNotificationType"
	case KIOKitNoticationMsgSizeMask:
		return "KIOKitNoticationMsgSizeMask"
	case KIOKitNoticationTypeMask:
		return "KIOKitNoticationTypeMask"
	case KIOKitNoticationTypeSizeAdjShift:
		return "KIOKitNoticationTypeSizeAdjShift"
	case KIOServiceMatchedNotificationType:
		return "KIOServiceMatchedNotificationType"
	case KIOServiceMessageNotificationType:
		return "KIOServiceMessageNotificationType"
	case KIOServiceTerminatedNotificationType:
		return "KIOServiceTerminatedNotificationType"
	case KLastIOKitNotificationType:
		return "KLastIOKitNotificationType"
	default:
		return fmt.Sprintf("KFirstIOKitNotificationType(%d)", e)
	}
}

type KIO uint

const (
	KIOAsyncCalloutCount           KIO = 3
	KIOAsyncCalloutFuncIndex       KIO = 1
	KIOAsyncCalloutRefconIndex     KIO = 2
	KIOAsyncReservedCount          KIO = 1
	KIOAsyncReservedIndex          KIO = 0
	KIOCopybackCache               KIO = 3
	KIOCopybackInnerCache          KIO = 5
	KIODefaultCache                KIO = 0
	KIOInhibitCache                KIO = 1
	KIOInterestCalloutCount        KIO = 4
	KIOInterestCalloutFuncIndex    KIO = 1
	KIOInterestCalloutRefconIndex  KIO = 2
	KIOInterestCalloutServiceIndex KIO = 3
	KIOMatchingCalloutCount        KIO = 3
	KIOMatchingCalloutFuncIndex    KIO = 1
	KIOMatchingCalloutRefconIndex  KIO = 2
	KIOWriteCombineCache           KIO = 4
	KIOWriteThruCache              KIO = 2
)

func (e KIO) String() string {
	switch e {
	case KIOAsyncCalloutCount:
		return "KIOAsyncCalloutCount"
	case KIOAsyncCalloutFuncIndex:
		return "KIOAsyncCalloutFuncIndex"
	case KIOAsyncCalloutRefconIndex:
		return "KIOAsyncCalloutRefconIndex"
	case KIOAsyncReservedIndex:
		return "KIOAsyncReservedIndex"
	case KIOCopybackInnerCache:
		return "KIOCopybackInnerCache"
	case KIOInterestCalloutCount:
		return "KIOInterestCalloutCount"
	default:
		return fmt.Sprintf("KIO(%d)", e)
	}
}

const KIOCFSerializeToBinary uint = 0

const KIOConnectMethodVarOutputSize int = -3

const KIODefaultMemoryType uint = 0

type KIOMap uint

const (
	KIOMapAnywhere           KIOMap = 1
	KIOMapCacheMask          KIOMap = 3840
	KIOMapCacheShift         KIOMap = 8
	KIOMapCopybackCache      KIOMap = 3
	KIOMapCopybackInnerCache KIOMap = 5
	KIOMapDefaultCache       KIOMap = 0
	KIOMapInhibitCache       KIOMap = 1
	KIOMapOverwrite          KIOMap = 536870912
	KIOMapPrefault           KIOMap = 268435456
	KIOMapReadOnly           KIOMap = 4096
	KIOMapReference          KIOMap = 33554432
	KIOMapStatic             KIOMap = 16777216
	KIOMapUnique             KIOMap = 67108864
	KIOMapUserOptionsMask    KIOMap = 4095
	KIOMapWriteCombineCache  KIOMap = 4
	KIOMapWriteThruCache     KIOMap = 2
)

func (e KIOMap) String() string {
	switch e {
	case KIOMapAnywhere:
		return "KIOMapAnywhere"
	case KIOMapCacheMask:
		return "KIOMapCacheMask"
	case KIOMapCacheShift:
		return "KIOMapCacheShift"
	case KIOMapCopybackCache:
		return "KIOMapCopybackCache"
	case KIOMapCopybackInnerCache:
		return "KIOMapCopybackInnerCache"
	case KIOMapDefaultCache:
		return "KIOMapDefaultCache"
	case KIOMapOverwrite:
		return "KIOMapOverwrite"
	case KIOMapPrefault:
		return "KIOMapPrefault"
	case KIOMapReadOnly:
		return "KIOMapReadOnly"
	case KIOMapReference:
		return "KIOMapReference"
	case KIOMapStatic:
		return "KIOMapStatic"
	case KIOMapUnique:
		return "KIOMapUnique"
	case KIOMapUserOptionsMask:
		return "KIOMapUserOptionsMask"
	case KIOMapWriteCombineCache:
		return "KIOMapWriteCombineCache"
	case KIOMapWriteThruCache:
		return "KIOMapWriteThruCache"
	default:
		return fmt.Sprintf("KIOMap(%d)", e)
	}
}

type KIORegistryIterate uint

const (
	KIORegistryIterateParents     KIORegistryIterate = 2
	KIORegistryIterateRecursively KIORegistryIterate = 1
)

func (e KIORegistryIterate) String() string {
	switch e {
	case KIORegistryIterateParents:
		return "KIORegistryIterateParents"
	case KIORegistryIterateRecursively:
		return "KIORegistryIterateRecursively"
	default:
		return fmt.Sprintf("KIORegistryIterate(%d)", e)
	}
}

const KIOServiceInteractionAllowed uint = 1

type KNanosecondScale uint

const (
	// KMicrosecondScale: # Discussion
	KMicrosecondScale KNanosecondScale = 1000
	// KMillisecondScale: # Discussion
	KMillisecondScale KNanosecondScale = 0
	// KNanosecondScaleValue: # Discussion
	KNanosecondScaleValue KNanosecondScale = 1
	// KSecondScale: # Discussion
	KSecondScale KNanosecondScale = 0
	// KTickScale: # Discussion
	KTickScale KNanosecondScale = 0
)

func (e KNanosecondScale) String() string {
	switch e {
	case KMicrosecondScale:
		return "KMicrosecondScale"
	case KMillisecondScale:
		return "KMillisecondScale"
	case KNanosecondScaleValue:
		return "KNanosecondScaleValue"
	default:
		return fmt.Sprintf("KNanosecondScale(%d)", e)
	}
}

type KOSAsyncRef uint

const (
	KOSAsyncRefCount KOSAsyncRef = 8
	KOSAsyncRefSize  KOSAsyncRef = 32
)

func (e KOSAsyncRef) String() string {
	switch e {
	case KOSAsyncRefCount:
		return "KOSAsyncRefCount"
	case KOSAsyncRefSize:
		return "KOSAsyncRefSize"
	default:
		return fmt.Sprintf("KOSAsyncRef(%d)", e)
	}
}

type KOSAsyncRef64 uint

const (
	KOSAsyncRef64Count KOSAsyncRef64 = 8
	KOSAsyncRef64Size  KOSAsyncRef64 = 8
)

func (e KOSAsyncRef64) String() string {
	switch e {
	case KOSAsyncRef64Count:
		return "KOSAsyncRef64Count"
	default:
		return fmt.Sprintf("KOSAsyncRef64(%d)", e)
	}
}

type KOSNotificationMessageID uint

const (
	KMaxAsyncArgs                 KOSNotificationMessageID = 16
	KOSAsyncCompleteMessageID     KOSNotificationMessageID = 57
	KOSNotificationMessageIDValue KOSNotificationMessageID = 53
)

func (e KOSNotificationMessageID) String() string {
	switch e {
	case KMaxAsyncArgs:
		return "KMaxAsyncArgs"
	case KOSAsyncCompleteMessageID:
		return "KOSAsyncCompleteMessageID"
	case KOSNotificationMessageIDValue:
		return "KOSNotificationMessageIDValue"
	default:
		return fmt.Sprintf("KOSNotificationMessageID(%d)", e)
	}
}
