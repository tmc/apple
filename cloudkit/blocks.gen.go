// Code generated from Apple documentation. DO NOT EDIT.

package cloudkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// CKAccountStatusErrorHandler handles The handler to execute when the call completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.AccountStatusWithCompletionHandler]
type CKAccountStatusErrorHandler = func(CKAccountStatus, error)

// NewCKAccountStatusErrorBlock wraps a Go [CKAccountStatusErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.AccountStatusWithCompletionHandler]
func NewCKAccountStatusErrorBlock(handler CKAccountStatusErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal CKAccountStatus, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKApplicationPermissionBlock handles A closure that processes the outcome of a permissions request.

// NewCKApplicationPermissionBlock wraps a Go [CKApplicationPermissionBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKApplicationPermissionBlock(handler CKApplicationPermissionBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive CKApplicationPermissionStatus, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKApplicationPermissionStatusErrorHandler handles The handler to execute with the outcome.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.RequestApplicationPermissionCompletionHandler]
//   - [CKContainer.StatusForApplicationPermissionCompletionHandler]
type CKApplicationPermissionStatusErrorHandler = func(int, error)

// NewCKApplicationPermissionStatusErrorBlock wraps a Go [CKApplicationPermissionStatusErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.RequestApplicationPermissionCompletionHandler]
//   - [CKContainer.StatusForApplicationPermissionCompletionHandler]
func NewCKApplicationPermissionStatusErrorBlock(handler CKApplicationPermissionStatusErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal int, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordErrorHandler handles The closure to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKDatabase.FetchRecordWithIDCompletionHandler]
//   - [CKDatabase.SaveRecordCompletionHandler]
type CKRecordErrorHandler = func(*CKRecord, error)

// NewCKRecordErrorBlock wraps a Go [CKRecordErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKDatabase.FetchRecordWithIDCompletionHandler]
//   - [CKDatabase.SaveRecordCompletionHandler]
func NewCKRecordErrorBlock(handler CKRecordErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKRecord
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordFloat64Handler is the signature for a completion handler block.
type CKRecordFloat64Handler = func(*CKRecord, float64)

// NewCKRecordFloat64Block wraps a Go [CKRecordFloat64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKRecordFloat64Block(handler CKRecordFloat64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float64) {
		var result *CKRecord
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordIDErrorHandler handles The handler to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.FetchUserRecordIDWithCompletionHandler]
//   - [CKDatabase.DeleteRecordWithIDCompletionHandler]
type CKRecordIDErrorHandler = func(*CKRecordID, error)

// NewCKRecordIDErrorBlock wraps a Go [CKRecordIDErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.FetchUserRecordIDWithCompletionHandler]
//   - [CKDatabase.DeleteRecordWithIDCompletionHandler]
func NewCKRecordIDErrorBlock(handler CKRecordIDErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKRecordID
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordIDFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordIDFloat64Handler is the signature for a completion handler block.
type CKRecordIDFloat64Handler = func(*CKRecordID, float64)

// NewCKRecordIDFloat64Block wraps a Go [CKRecordIDFloat64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKRecordIDFloat64Block(handler CKRecordIDFloat64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 float64) {
		var result *CKRecordID
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordIDFromID(resultID)
			result = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordIDHandler handles A block that returns the record for the specified record identifier.
//
// Used by:
//   - [CKSyncEngineRecordZoneChangeBatch.InitWithPendingChangesRecordProvider]
type CKRecordIDHandler = func(*CKRecordID) CKRecord

// NewCKRecordIDBlock wraps a Go [CKRecordIDHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKSyncEngineRecordZoneChangeBatch.InitWithPendingChangesRecordProvider]
func NewCKRecordIDBlock(handler CKRecordIDHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) objc.ID {
		var result *CKRecordID
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordIDFromID(resultID)
			result = &v
		}
		return handler(result).ID
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordZoneArrayErrorHandler handles The closure to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKDatabase.FetchAllRecordZonesWithCompletionHandler]
type CKRecordZoneArrayErrorHandler = func(*[]CKRecordZone, error)

// NewCKRecordZoneArrayErrorBlock wraps a Go [CKRecordZoneArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKDatabase.FetchAllRecordZonesWithCompletionHandler]
func NewCKRecordZoneArrayErrorBlock(handler CKRecordZoneArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]CKRecordZone
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]CKRecordZone, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = CKRecordZoneFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordZoneErrorHandler handles The closure to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKDatabase.FetchRecordZoneWithIDCompletionHandler]
//   - [CKDatabase.SaveRecordZoneCompletionHandler]
type CKRecordZoneErrorHandler = func(*CKRecordZone, error)

// NewCKRecordZoneErrorBlock wraps a Go [CKRecordZoneErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKDatabase.FetchRecordZoneWithIDCompletionHandler]
//   - [CKDatabase.SaveRecordZoneCompletionHandler]
func NewCKRecordZoneErrorBlock(handler CKRecordZoneErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKRecordZone
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordZoneFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordZoneIDCKServerChangeTokenDataHandler is the signature for a completion handler block.
type CKRecordZoneIDCKServerChangeTokenDataHandler = func(*CKRecordZoneID, *CKServerChangeToken, *foundation.NSData)

// NewCKRecordZoneIDCKServerChangeTokenDataBlock wraps a Go [CKRecordZoneIDCKServerChangeTokenDataHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKRecordZoneIDCKServerChangeTokenDataBlock(handler CKRecordZoneIDCKServerChangeTokenDataHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1ID objc.ID) {
		var result *CKRecordZoneID
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordZoneIDFromID(resultID)
			result = &v
		}
		var extra0 *CKServerChangeToken
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := CKServerChangeTokenFromID(extra0ID)
			extra0 = &v
		}
		var extra1 *foundation.NSData
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := foundation.NSDataFromID(extra1ID)
			extra1 = &v
		}
		handler(result, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordZoneIDErrorHandler handles The closure to execute after CloudKit deletes the record zone.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKDatabase.DeleteRecordZoneWithIDCompletionHandler]
type CKRecordZoneIDErrorHandler = func(*CKRecordZoneID, error)

// NewCKRecordZoneIDErrorBlock wraps a Go [CKRecordZoneIDErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKDatabase.DeleteRecordZoneWithIDCompletionHandler]
func NewCKRecordZoneIDErrorBlock(handler CKRecordZoneIDErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKRecordZoneID
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordZoneIDFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKRecordZoneIDHandler is the signature for a completion handler block.
type CKRecordZoneIDHandler = func(*CKRecordZoneID)

// NewCKRecordZoneIDBlock wraps a Go [CKRecordZoneIDHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKRecordZoneIDBlock(handler CKRecordZoneIDHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *CKRecordZoneID
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKRecordZoneIDFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKServerChangeTokenHandler is the signature for a completion handler block.
type CKServerChangeTokenHandler = func(*CKServerChangeToken)

// NewCKServerChangeTokenBlock wraps a Go [CKServerChangeTokenHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKServerChangeTokenBlock(handler CKServerChangeTokenHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *CKServerChangeToken
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKServerChangeTokenFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKShareErrorHandler handles The handler to execute when the process finishes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.AcceptShareMetadataCompletionHandler]
type CKShareErrorHandler = func(*CKShare, error)

// NewCKShareErrorBlock wraps a Go [CKShareErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.AcceptShareMetadataCompletionHandler]
func NewCKShareErrorBlock(handler CKShareErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKShare
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKShareFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKShareMetadataErrorHandler handles The handler to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.FetchShareMetadataWithURLCompletionHandler]
type CKShareMetadataErrorHandler = func(*CKShareMetadata, error)

// NewCKShareMetadataErrorBlock wraps a Go [CKShareMetadataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.FetchShareMetadataWithURLCompletionHandler]
func NewCKShareMetadataErrorBlock(handler CKShareMetadataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKShareMetadata
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKShareMetadataFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKShareParticipantErrorHandler handles The handler to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.FetchShareParticipantWithEmailAddressCompletionHandler]
//   - [CKContainer.FetchShareParticipantWithPhoneNumberCompletionHandler]
//   - [CKContainer.FetchShareParticipantWithUserRecordIDCompletionHandler]
type CKShareParticipantErrorHandler = func(*CKShareParticipant, error)

// NewCKShareParticipantErrorBlock wraps a Go [CKShareParticipantErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.FetchShareParticipantWithEmailAddressCompletionHandler]
//   - [CKContainer.FetchShareParticipantWithPhoneNumberCompletionHandler]
//   - [CKContainer.FetchShareParticipantWithUserRecordIDCompletionHandler]
func NewCKShareParticipantErrorBlock(handler CKShareParticipantErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKShareParticipant
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKShareParticipantFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKSharePreparationCompletionHandler handles completion with primitive and object results.

// NewCKSharePreparationCompletionHandlerBlock wraps a Go [CKSharePreparationCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKSharePreparationCompletionHandlerBlock(handler CKSharePreparationCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive CKShare, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKSharePreparationHandler is the signature for a completion handler block.

// CKSubscriptionArrayErrorHandler handles The closure to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKDatabase.FetchAllSubscriptionsWithCompletionHandler]
type CKSubscriptionArrayErrorHandler = func(*[]CKSubscription, error)

// NewCKSubscriptionArrayErrorBlock wraps a Go [CKSubscriptionArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKDatabase.FetchAllSubscriptionsWithCompletionHandler]
func NewCKSubscriptionArrayErrorBlock(handler CKSubscriptionArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]CKSubscription
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]CKSubscription, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = CKSubscriptionFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKSubscriptionErrorHandler handles The closure to execute after CloudKit saves the subscription.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKDatabase.SaveSubscriptionCompletionHandler]
type CKSubscriptionErrorHandler = func(*CKSubscription, error)

// NewCKSubscriptionErrorBlock wraps a Go [CKSubscriptionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKDatabase.SaveSubscriptionCompletionHandler]
func NewCKSubscriptionErrorBlock(handler CKSubscriptionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKSubscription
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKSubscriptionFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKUserIdentityArrayErrorHandler handles The handler to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.DiscoverAllIdentitiesWithCompletionHandler]
type CKUserIdentityArrayErrorHandler = func(*[]CKUserIdentity, error)

// NewCKUserIdentityArrayErrorBlock wraps a Go [CKUserIdentityArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.DiscoverAllIdentitiesWithCompletionHandler]
func NewCKUserIdentityArrayErrorBlock(handler CKUserIdentityArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]CKUserIdentity
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]CKUserIdentity, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = CKUserIdentityFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKUserIdentityCKUserIdentityLookupInfoHandler is the signature for a completion handler block.
type CKUserIdentityCKUserIdentityLookupInfoHandler = func(*CKUserIdentity, *CKUserIdentityLookupInfo)

// NewCKUserIdentityCKUserIdentityLookupInfoBlock wraps a Go [CKUserIdentityCKUserIdentityLookupInfoHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKUserIdentityCKUserIdentityLookupInfoBlock(handler CKUserIdentityCKUserIdentityLookupInfoHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *CKUserIdentity
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKUserIdentityFromID(resultID)
			result = &v
		}
		var extra0 *CKUserIdentityLookupInfo
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := CKUserIdentityLookupInfoFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CKUserIdentityErrorHandler handles The handler to execute with the fetch results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKContainer.DiscoverUserIdentityWithEmailAddressCompletionHandler]
//   - [CKContainer.DiscoverUserIdentityWithPhoneNumberCompletionHandler]
//   - [CKContainer.DiscoverUserIdentityWithUserRecordIDCompletionHandler]
type CKUserIdentityErrorHandler = func(*CKUserIdentity, error)

// NewCKUserIdentityErrorBlock wraps a Go [CKUserIdentityErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKContainer.DiscoverUserIdentityWithEmailAddressCompletionHandler]
//   - [CKContainer.DiscoverUserIdentityWithPhoneNumberCompletionHandler]
//   - [CKContainer.DiscoverUserIdentityWithUserRecordIDCompletionHandler]
func NewCKUserIdentityErrorBlock(handler CKUserIdentityErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *CKUserIdentity
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKUserIdentityFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CKUserIdentityHandler is the signature for a completion handler block.
type CKUserIdentityHandler = func(*CKUserIdentity)

// NewCKUserIdentityBlock wraps a Go [CKUserIdentityHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewCKUserIdentityBlock(handler CKUserIdentityHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *CKUserIdentity
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CKUserIdentityFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles The block to execute when the fetch completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CKSyncEngine.FetchChangesWithCompletionHandler]
//   - [CKSyncEngine.FetchChangesWithOptionsCompletionHandler]
//   - [CKSyncEngine.SendChangesWithCompletionHandler]
//   - [CKSyncEngine.SendChangesWithOptionsCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKSyncEngine.FetchChangesWithCompletionHandler]
//   - [CKSyncEngine.FetchChangesWithOptionsCompletionHandler]
//   - [CKSyncEngine.SendChangesWithCompletionHandler]
//   - [CKSyncEngine.SendChangesWithOptionsCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [CKSyncEngine.CancelOperationsWithCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CKSyncEngine.CancelOperationsWithCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
