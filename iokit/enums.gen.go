// Code generated from Apple documentation for iokit. DO NOT EDIT.

package iokit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/iokit/iourlerror
type IOURLError int32

const (
	KIOURLImproperArgumentsError       IOURLError = -15
	KIOURLPropertyKeyUnavailableError  IOURLError = -17
	KIOURLRemoteHostUnavailableError   IOURLError = -14
	KIOURLResourceAccessViolationError IOURLError = -13
	KIOURLResourceNotFoundError        IOURLError = -12
	KIOURLTimeoutError                 IOURLError = -18
	KIOURLUnknownError                 IOURLError = -10
	KIOURLUnknownPropertyKeyError      IOURLError = -16
	KIOURLUnknownSchemeError           IOURLError = -11
)

func (e IOURLError) String() string {
	switch e {
	case KIOURLImproperArgumentsError:
		return "KIOURLImproperArgumentsError"
	case KIOURLPropertyKeyUnavailableError:
		return "KIOURLPropertyKeyUnavailableError"
	case KIOURLRemoteHostUnavailableError:
		return "KIOURLRemoteHostUnavailableError"
	case KIOURLResourceAccessViolationError:
		return "KIOURLResourceAccessViolationError"
	case KIOURLResourceNotFoundError:
		return "KIOURLResourceNotFoundError"
	case KIOURLTimeoutError:
		return "KIOURLTimeoutError"
	case KIOURLUnknownError:
		return "KIOURLUnknownError"
	case KIOURLUnknownPropertyKeyError:
		return "KIOURLUnknownPropertyKeyError"
	case KIOURLUnknownSchemeError:
		return "KIOURLUnknownSchemeError"
	default:
		return fmt.Sprintf("IOURLError(%d)", e)
	}
}

type KDisplayVendorID uint32

const (
	KDisplayVendorIDUnknown KDisplayVendorID = 'u'<<24 | 'n'<<16 | 'k'<<8 | 'n' // 'unkn'
)

func (e KDisplayVendorID) String() string {
	switch e {
	case KDisplayVendorIDUnknown:
		return "KDisplayVendorIDUnknown"
	default:
		return fmt.Sprintf("KDisplayVendorID(%d)", e)
	}
}

type KFirstIOKitNotificationType uint32

const (
	KFirstIOKitNotificationTypeValue     KFirstIOKitNotificationType = 100
	KIOAsyncCompletionNotificationType   KFirstIOKitNotificationType = 150
	KIOKitNoticationMsgSizeMask          KFirstIOKitNotificationType = 3
	KIOKitNoticationTypeMask             KFirstIOKitNotificationType = 0xfff
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

type KIOAsyncReservedIndex uint32

const (
	KIOAsyncCalloutCount           KIOAsyncReservedIndex = 3
	KIOAsyncCalloutFuncIndex       KIOAsyncReservedIndex = 1
	KIOAsyncCalloutRefconIndex     KIOAsyncReservedIndex = 2
	KIOAsyncReservedCount          KIOAsyncReservedIndex = 1
	KIOAsyncReservedIndexValue     KIOAsyncReservedIndex = 0
	KIOInterestCalloutCount        KIOAsyncReservedIndex = 4
	KIOInterestCalloutFuncIndex    KIOAsyncReservedIndex = 1
	KIOInterestCalloutRefconIndex  KIOAsyncReservedIndex = 2
	KIOInterestCalloutServiceIndex KIOAsyncReservedIndex = 3
	KIOMatchingCalloutCount        KIOAsyncReservedIndex = 3
	KIOMatchingCalloutFuncIndex    KIOAsyncReservedIndex = 1
	KIOMatchingCalloutRefconIndex  KIOAsyncReservedIndex = 2
)

func (e KIOAsyncReservedIndex) String() string {
	switch e {
	case KIOAsyncCalloutCount:
		return "KIOAsyncCalloutCount"
	case KIOAsyncCalloutFuncIndex:
		return "KIOAsyncCalloutFuncIndex"
	case KIOAsyncCalloutRefconIndex:
		return "KIOAsyncCalloutRefconIndex"
	case KIOAsyncReservedIndexValue:
		return "KIOAsyncReservedIndexValue"
	case KIOInterestCalloutCount:
		return "KIOInterestCalloutCount"
	default:
		return fmt.Sprintf("KIOAsyncReservedIndex(%d)", e)
	}
}

type KIOCFSerializeTo uint

const (
	KIOCFSerializeToBinary KIOCFSerializeTo = 0
)

func (e KIOCFSerializeTo) String() string {
	switch e {
	case KIOCFSerializeToBinary:
		return "KIOCFSerializeToBinary"
	default:
		return fmt.Sprintf("KIOCFSerializeTo(%d)", e)
	}
}

type KIOConnectMethodVarOutput int32

const (
	KIOConnectMethodVarOutputSize KIOConnectMethodVarOutput = -3
)

func (e KIOConnectMethodVarOutput) String() string {
	switch e {
	case KIOConnectMethodVarOutputSize:
		return "KIOConnectMethodVarOutputSize"
	default:
		return fmt.Sprintf("KIOConnectMethodVarOutput(%d)", e)
	}
}

type KIODefaultCache uint32

const (
	KIOCopybackCache      KIODefaultCache = 3
	KIOCopybackInnerCache KIODefaultCache = 5
	KIODefaultCacheValue  KIODefaultCache = 0
	KIOInhibitCache       KIODefaultCache = 1
	KIOWriteCombineCache  KIODefaultCache = 4
	KIOWriteThruCache     KIODefaultCache = 2
)

func (e KIODefaultCache) String() string {
	switch e {
	case KIOCopybackCache:
		return "KIOCopybackCache"
	case KIOCopybackInnerCache:
		return "KIOCopybackInnerCache"
	case KIODefaultCacheValue:
		return "KIODefaultCacheValue"
	case KIOInhibitCache:
		return "KIOInhibitCache"
	case KIOWriteCombineCache:
		return "KIOWriteCombineCache"
	case KIOWriteThruCache:
		return "KIOWriteThruCache"
	default:
		return fmt.Sprintf("KIODefaultCache(%d)", e)
	}
}

type KIODefaultMemory uint32

const (
	KIODefaultMemoryType KIODefaultMemory = 0
)

func (e KIODefaultMemory) String() string {
	switch e {
	case KIODefaultMemoryType:
		return "KIODefaultMemoryType"
	default:
		return fmt.Sprintf("KIODefaultMemory(%d)", e)
	}
}

type KIOMap uint32

const (
	KIOMapAnywhere           KIOMap = 0x1
	KIOMapCacheMask          KIOMap = 0xf00
	KIOMapCacheShift         KIOMap = 8
	KIOMapCopybackCache      KIOMap = 768
	KIOMapCopybackInnerCache KIOMap = 1280
	KIOMapDefaultCache       KIOMap = 0
	KIOMapInhibitCache       KIOMap = 256
	KIOMapOverwrite          KIOMap = 0x20000000
	KIOMapPrefault           KIOMap = 0x10000000
	KIOMapReadOnly           KIOMap = 0x1000
	KIOMapReference          KIOMap = 0x2000000
	KIOMapStatic             KIOMap = 0x1000000
	KIOMapUnique             KIOMap = 0x4000000
	KIOMapUserOptionsMask    KIOMap = 0xfff
	KIOMapWriteCombineCache  KIOMap = 1024
	KIOMapWriteThruCache     KIOMap = 512
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
	case KIOMapInhibitCache:
		return "KIOMapInhibitCache"
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

type KIORegistryIterate uint32

const (
	KIORegistryIterateParents     KIORegistryIterate = 0x2
	KIORegistryIterateRecursively KIORegistryIterate = 0x1
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

type KIOServiceInteraction uint32

const (
	KIOServiceInteractionAllowed KIOServiceInteraction = 0x1
)

func (e KIOServiceInteraction) String() string {
	switch e {
	case KIOServiceInteractionAllowed:
		return "KIOServiceInteractionAllowed"
	default:
		return fmt.Sprintf("KIOServiceInteraction(%d)", e)
	}
}

type KNanosecondScale uint32

const (
	// KMicrosecondScale: # Discussion
	KMicrosecondScale KNanosecondScale = 1000
	// KMillisecondScale: # Discussion
	KMillisecondScale KNanosecondScale = 1000000
	// KNanosecondScaleValue: # Discussion
	KNanosecondScaleValue KNanosecondScale = 1
	// KSecondScale: # Discussion
	KSecondScale KNanosecondScale = 1000000000
	// KTickScale: # Discussion
	KTickScale KNanosecondScale = 10000000
)

func (e KNanosecondScale) String() string {
	switch e {
	case KMicrosecondScale:
		return "KMicrosecondScale"
	case KMillisecondScale:
		return "KMillisecondScale"
	case KNanosecondScaleValue:
		return "KNanosecondScaleValue"
	case KSecondScale:
		return "KSecondScale"
	case KTickScale:
		return "KTickScale"
	default:
		return fmt.Sprintf("KNanosecondScale(%d)", e)
	}
}

type KOSAsyncRef uint32

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

type KOSAsyncRef64 uint32

const (
	KOSAsyncRef64Count KOSAsyncRef64 = 8
	KOSAsyncRef64Size  KOSAsyncRef64 = 64
)

func (e KOSAsyncRef64) String() string {
	switch e {
	case KOSAsyncRef64Count:
		return "KOSAsyncRef64Count"
	case KOSAsyncRef64Size:
		return "KOSAsyncRef64Size"
	default:
		return fmt.Sprintf("KOSAsyncRef64(%d)", e)
	}
}

type KOSNotificationMessageID uint32

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
