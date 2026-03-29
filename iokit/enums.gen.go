// Code generated from Apple documentation for iokit. DO NOT EDIT.

package iokit

import (
	"fmt"
)

type KFirstIOKitNotificationType uint

const (
	KFirstIOKitNotificationTypeValue KFirstIOKitNotificationType = 0
	KIOAsyncCompletionNotificationType KFirstIOKitNotificationType = 0
	KIOKitNoticationMsgSizeMask KFirstIOKitNotificationType = 0
	KIOKitNoticationTypeMask KFirstIOKitNotificationType = 0
	KIOKitNoticationTypeSizeAdjShift KFirstIOKitNotificationType = 0
	KIOServiceMatchedNotificationType KFirstIOKitNotificationType = 0
	KIOServiceMessageNotificationType KFirstIOKitNotificationType = 0
	KIOServicePublishNotificationType KFirstIOKitNotificationType = 0
	KIOServiceTerminatedNotificationType KFirstIOKitNotificationType = 0
	KLastIOKitNotificationType KFirstIOKitNotificationType = 0
)

func (e KFirstIOKitNotificationType) String() string {
	switch e {
	case KFirstIOKitNotificationTypeValue:
		return "KFirstIOKitNotificationTypeValue"
	default:
		return fmt.Sprintf("KFirstIOKitNotificationType(%d)", e)
	}
}

type KIO uint

const (
	KIOAsyncCalloutCount KIO = 0
	KIOAsyncCalloutFuncIndex KIO = 0
	KIOAsyncCalloutRefconIndex KIO = 0
	KIOAsyncReservedCount KIO = 0
	KIOAsyncReservedIndex KIO = 0
	KIOCopybackCache KIO = 0
	KIOCopybackInnerCache KIO = 0
	KIODefaultCache KIO = 0
	KIOInhibitCache KIO = 0
	KIOInterestCalloutCount KIO = 0
	KIOInterestCalloutFuncIndex KIO = 0
	KIOInterestCalloutRefconIndex KIO = 0
	KIOInterestCalloutServiceIndex KIO = 0
	KIOMatchingCalloutCount KIO = 0
	KIOMatchingCalloutFuncIndex KIO = 0
	KIOMatchingCalloutRefconIndex KIO = 0
	KIOWriteCombineCache KIO = 0
	KIOWriteThruCache KIO = 0
)

func (e KIO) String() string {
	switch e {
	case KIOAsyncCalloutCount:
		return "KIOAsyncCalloutCount"
	default:
		return fmt.Sprintf("KIO(%d)", e)
	}
}

const KIOCFSerializeToBinary uint = 0

const KIOConnectMethodVarOutputSize uint = 0

const KIODefaultMemoryType uint = 0

type KIOMap uint

const (
	KIOMapAnywhere KIOMap = 0
	KIOMapCacheMask KIOMap = 0
	KIOMapCacheShift KIOMap = 0
	KIOMapCopybackCache KIOMap = 0
	KIOMapCopybackInnerCache KIOMap = 0
	KIOMapDefaultCache KIOMap = 0
	KIOMapInhibitCache KIOMap = 0
	KIOMapOverwrite KIOMap = 0
	KIOMapPrefault KIOMap = 0
	KIOMapReadOnly KIOMap = 0
	KIOMapReference KIOMap = 0
	KIOMapStatic KIOMap = 0
	KIOMapUnique KIOMap = 0
	KIOMapUserOptionsMask KIOMap = 0
	KIOMapWriteCombineCache KIOMap = 0
	KIOMapWriteThruCache KIOMap = 0
)

func (e KIOMap) String() string {
	switch e {
	case KIOMapAnywhere:
		return "KIOMapAnywhere"
	default:
		return fmt.Sprintf("KIOMap(%d)", e)
	}
}

type KIORegistryIterate uint

const (
	KIORegistryIterateParents KIORegistryIterate = 0
	KIORegistryIterateRecursively KIORegistryIterate = 0
)

func (e KIORegistryIterate) String() string {
	switch e {
	case KIORegistryIterateParents:
		return "KIORegistryIterateParents"
	default:
		return fmt.Sprintf("KIORegistryIterate(%d)", e)
	}
}

const KIOServiceInteractionAllowed uint = 0

type KNanosecondScale uint

const (
	// KMicrosecondScale: # Discussion
	KMicrosecondScale KNanosecondScale = 0
	// KMillisecondScale: # Discussion
	KMillisecondScale KNanosecondScale = 0
	// KNanosecondScaleValue: # Discussion
	KNanosecondScaleValue KNanosecondScale = 0
	// KSecondScale: # Discussion
	KSecondScale KNanosecondScale = 0
	// KTickScale: # Discussion
	KTickScale KNanosecondScale = 0
)

func (e KNanosecondScale) String() string {
	switch e {
	case KMicrosecondScale:
		return "KMicrosecondScale"
	default:
		return fmt.Sprintf("KNanosecondScale(%d)", e)
	}
}

type KOSAsyncRef uint

const (
	KOSAsyncRefCount KOSAsyncRef = 0
	KOSAsyncRefSize KOSAsyncRef = 0
)

func (e KOSAsyncRef) String() string {
	switch e {
	case KOSAsyncRefCount:
		return "KOSAsyncRefCount"
	default:
		return fmt.Sprintf("KOSAsyncRef(%d)", e)
	}
}

type KOSAsyncRef64 uint

const (
	KOSAsyncRef64Count KOSAsyncRef64 = 0
	KOSAsyncRef64Size KOSAsyncRef64 = 0
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
	KMaxAsyncArgs KOSNotificationMessageID = 0
	KOSAsyncCompleteMessageID KOSNotificationMessageID = 0
	KOSNotificationMessageIDValue KOSNotificationMessageID = 0
)

func (e KOSNotificationMessageID) String() string {
	switch e {
	case KMaxAsyncArgs:
		return "KMaxAsyncArgs"
	default:
		return fmt.Sprintf("KOSNotificationMessageID(%d)", e)
	}
}

