// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetDownloadStorageManagementPolicy] class.
var (
	_AVAssetDownloadStorageManagementPolicyClass     AVAssetDownloadStorageManagementPolicyClass
	_AVAssetDownloadStorageManagementPolicyClassOnce sync.Once
)

func getAVAssetDownloadStorageManagementPolicyClass() AVAssetDownloadStorageManagementPolicyClass {
	_AVAssetDownloadStorageManagementPolicyClassOnce.Do(func() {
		_AVAssetDownloadStorageManagementPolicyClass = AVAssetDownloadStorageManagementPolicyClass{class: objc.GetClass("AVAssetDownloadStorageManagementPolicy")}
	})
	return _AVAssetDownloadStorageManagementPolicyClass
}

// GetAVAssetDownloadStorageManagementPolicyClass returns the class object for AVAssetDownloadStorageManagementPolicy.
func GetAVAssetDownloadStorageManagementPolicyClass() AVAssetDownloadStorageManagementPolicyClass {
	return getAVAssetDownloadStorageManagementPolicyClass()
}

type AVAssetDownloadStorageManagementPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetDownloadStorageManagementPolicyClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetDownloadStorageManagementPolicyClass) Alloc() AVAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AVAssetDownloadStorageManagementPolicy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that defines a policy to automatically manage the storage of
// downloaded assets.
//
// # Inspecting a policy
//
//   - [AVAssetDownloadStorageManagementPolicy.ExpirationDate]: The expiration date for an asset.
//   - [AVAssetDownloadStorageManagementPolicy.Priority]: The eviction priority for an asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy
type AVAssetDownloadStorageManagementPolicy struct {
	objectivec.Object
}

// AVAssetDownloadStorageManagementPolicyFromID constructs a [AVAssetDownloadStorageManagementPolicy] from an objc.ID.
//
// An object that defines a policy to automatically manage the storage of
// downloaded assets.
func AVAssetDownloadStorageManagementPolicyFromID(id objc.ID) AVAssetDownloadStorageManagementPolicy {
	return AVAssetDownloadStorageManagementPolicy{objectivec.Object{ID: id}}
}
// NOTE: AVAssetDownloadStorageManagementPolicy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetDownloadStorageManagementPolicy] class.
//
// # Inspecting a policy
//
//   - [IAVAssetDownloadStorageManagementPolicy.ExpirationDate]: The expiration date for an asset.
//   - [IAVAssetDownloadStorageManagementPolicy.Priority]: The eviction priority for an asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy
type IAVAssetDownloadStorageManagementPolicy interface {
	objectivec.IObject

	// Topic: Inspecting a policy

	// The expiration date for an asset.
	ExpirationDate() foundation.INSDate
	// The eviction priority for an asset.
	Priority() AVAssetDownloadedAssetEvictionPriority
}

// Init initializes the instance.
func (a AVAssetDownloadStorageManagementPolicy) Init() AVAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AVAssetDownloadStorageManagementPolicy](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetDownloadStorageManagementPolicy) Autorelease() AVAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AVAssetDownloadStorageManagementPolicy](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadStorageManagementPolicy creates a new AVAssetDownloadStorageManagementPolicy instance.
func NewAVAssetDownloadStorageManagementPolicy() AVAssetDownloadStorageManagementPolicy {
	class := getAVAssetDownloadStorageManagementPolicyClass()
	rv := objc.Send[AVAssetDownloadStorageManagementPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The expiration date for an asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy/expirationDate
func (a AVAssetDownloadStorageManagementPolicy) ExpirationDate() foundation.INSDate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("expirationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// The eviction priority for an asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy/priority
func (a AVAssetDownloadStorageManagementPolicy) Priority() AVAssetDownloadedAssetEvictionPriority {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("priority"))
	return AVAssetDownloadedAssetEvictionPriority(foundation.NSStringFromID(rv).String())
}

