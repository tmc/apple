// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableAssetDownloadStorageManagementPolicy] class.
var (
	_AVMutableAssetDownloadStorageManagementPolicyClass     AVMutableAssetDownloadStorageManagementPolicyClass
	_AVMutableAssetDownloadStorageManagementPolicyClassOnce sync.Once
)

func getAVMutableAssetDownloadStorageManagementPolicyClass() AVMutableAssetDownloadStorageManagementPolicyClass {
	_AVMutableAssetDownloadStorageManagementPolicyClassOnce.Do(func() {
		_AVMutableAssetDownloadStorageManagementPolicyClass = AVMutableAssetDownloadStorageManagementPolicyClass{class: objc.GetClass("AVMutableAssetDownloadStorageManagementPolicy")}
	})
	return _AVMutableAssetDownloadStorageManagementPolicyClass
}

// GetAVMutableAssetDownloadStorageManagementPolicyClass returns the class object for AVMutableAssetDownloadStorageManagementPolicy.
func GetAVMutableAssetDownloadStorageManagementPolicyClass() AVMutableAssetDownloadStorageManagementPolicyClass {
	return getAVMutableAssetDownloadStorageManagementPolicyClass()
}

type AVMutableAssetDownloadStorageManagementPolicyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableAssetDownloadStorageManagementPolicyClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableAssetDownloadStorageManagementPolicyClass) Alloc() AVMutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AVMutableAssetDownloadStorageManagementPolicy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable object that you use to create a new storage management policy.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy
type AVMutableAssetDownloadStorageManagementPolicy struct {
	AVAssetDownloadStorageManagementPolicy
}

// AVMutableAssetDownloadStorageManagementPolicyFromID constructs a [AVMutableAssetDownloadStorageManagementPolicy] from an objc.ID.
//
// A mutable object that you use to create a new storage management policy.
func AVMutableAssetDownloadStorageManagementPolicyFromID(id objc.ID) AVMutableAssetDownloadStorageManagementPolicy {
	return AVMutableAssetDownloadStorageManagementPolicy{AVAssetDownloadStorageManagementPolicy: AVAssetDownloadStorageManagementPolicyFromID(id)}
}

// NOTE: AVMutableAssetDownloadStorageManagementPolicy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableAssetDownloadStorageManagementPolicy] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy
type IAVMutableAssetDownloadStorageManagementPolicy interface {
	IAVAssetDownloadStorageManagementPolicy
}

// Init initializes the instance.
func (m AVMutableAssetDownloadStorageManagementPolicy) Init() AVMutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AVMutableAssetDownloadStorageManagementPolicy](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableAssetDownloadStorageManagementPolicy) Autorelease() AVMutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AVMutableAssetDownloadStorageManagementPolicy](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableAssetDownloadStorageManagementPolicy creates a new AVMutableAssetDownloadStorageManagementPolicy instance.
func NewAVMutableAssetDownloadStorageManagementPolicy() AVMutableAssetDownloadStorageManagementPolicy {
	class := getAVMutableAssetDownloadStorageManagementPolicyClass()
	rv := objc.Send[AVMutableAssetDownloadStorageManagementPolicy](objc.ID(class.class), objc.Sel("new"))
	return rv
}
