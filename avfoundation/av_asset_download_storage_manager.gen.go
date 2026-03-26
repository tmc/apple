// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetDownloadStorageManager] class.
var (
	_AVAssetDownloadStorageManagerClass     AVAssetDownloadStorageManagerClass
	_AVAssetDownloadStorageManagerClassOnce sync.Once
)

func getAVAssetDownloadStorageManagerClass() AVAssetDownloadStorageManagerClass {
	_AVAssetDownloadStorageManagerClassOnce.Do(func() {
		_AVAssetDownloadStorageManagerClass = AVAssetDownloadStorageManagerClass{class: objc.GetClass("AVAssetDownloadStorageManager")}
	})
	return _AVAssetDownloadStorageManagerClass
}

// GetAVAssetDownloadStorageManagerClass returns the class object for AVAssetDownloadStorageManager.
func GetAVAssetDownloadStorageManagerClass() AVAssetDownloadStorageManagerClass {
	return getAVAssetDownloadStorageManagerClass()
}

type AVAssetDownloadStorageManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetDownloadStorageManagerClass) Alloc() AVAssetDownloadStorageManager {
	rv := objc.Send[AVAssetDownloadStorageManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that manages policies to automatically purge downloaded assets.
//
// # Setting the storage policy
//
//   - [AVAssetDownloadStorageManager.StorageManagementPolicyForURL]: Returns the storage management policy for a downloaded asset.
//   - [AVAssetDownloadStorageManager.SetStorageManagementPolicyForURL]: Sets a storage policy for the downloaded asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager
type AVAssetDownloadStorageManager struct {
	objectivec.Object
}

// AVAssetDownloadStorageManagerFromID constructs a [AVAssetDownloadStorageManager] from an objc.ID.
//
// An object that manages policies to automatically purge downloaded assets.
func AVAssetDownloadStorageManagerFromID(id objc.ID) AVAssetDownloadStorageManager {
	return AVAssetDownloadStorageManager{objectivec.Object{ID: id}}
}
// NOTE: AVAssetDownloadStorageManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetDownloadStorageManager] class.
//
// # Setting the storage policy
//
//   - [IAVAssetDownloadStorageManager.StorageManagementPolicyForURL]: Returns the storage management policy for a downloaded asset.
//   - [IAVAssetDownloadStorageManager.SetStorageManagementPolicyForURL]: Sets a storage policy for the downloaded asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager
type IAVAssetDownloadStorageManager interface {
	objectivec.IObject

	// Topic: Setting the storage policy

	// Returns the storage management policy for a downloaded asset.
	StorageManagementPolicyForURL(downloadStorageURL foundation.INSURL) IAVAssetDownloadStorageManagementPolicy
	// Sets a storage policy for the downloaded asset.
	SetStorageManagementPolicyForURL(storageManagementPolicy IAVAssetDownloadStorageManagementPolicy, downloadStorageURL foundation.INSURL)
}

// Init initializes the instance.
func (a AVAssetDownloadStorageManager) Init() AVAssetDownloadStorageManager {
	rv := objc.Send[AVAssetDownloadStorageManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetDownloadStorageManager) Autorelease() AVAssetDownloadStorageManager {
	rv := objc.Send[AVAssetDownloadStorageManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetDownloadStorageManager creates a new AVAssetDownloadStorageManager instance.
func NewAVAssetDownloadStorageManager() AVAssetDownloadStorageManager {
	class := getAVAssetDownloadStorageManagerClass()
	rv := objc.Send[AVAssetDownloadStorageManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the storage management policy for a downloaded asset.
//
// downloadStorageURL: The location of the downloaded asset.
//
// # Return Value
// 
// The storage management policy for the asset, or `nil` if one isn’t set.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/storageManagementPolicy(for:)
func (a AVAssetDownloadStorageManager) StorageManagementPolicyForURL(downloadStorageURL foundation.INSURL) IAVAssetDownloadStorageManagementPolicy {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("storageManagementPolicyForURL:"), downloadStorageURL)
	return AVAssetDownloadStorageManagementPolicyFromID(rv)
}
// Sets a storage policy for the downloaded asset.
//
// storageManagementPolicy: The policy to set for the downloaded asset.
//
// downloadStorageURL: The location of the downloaded asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/setStorageManagementPolicy(_:for:)
func (a AVAssetDownloadStorageManager) SetStorageManagementPolicyForURL(storageManagementPolicy IAVAssetDownloadStorageManagementPolicy, downloadStorageURL foundation.INSURL) {
	objc.Send[objc.ID](a.ID, objc.Sel("setStorageManagementPolicy:forURL:"), storageManagementPolicy, downloadStorageURL)
}

// Returns the shared storage manager instance.
//
// # Return Value
// 
// The download storage manager.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/shared()
func (_AVAssetDownloadStorageManagerClass AVAssetDownloadStorageManagerClass) SharedDownloadStorageManager() AVAssetDownloadStorageManager {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetDownloadStorageManagerClass.class), objc.Sel("sharedDownloadStorageManager"))
	return AVAssetDownloadStorageManagerFromID(rv)
}

