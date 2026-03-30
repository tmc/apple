// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVFragmentedAssetMinder] class.
var (
	_AVFragmentedAssetMinderClass     AVFragmentedAssetMinderClass
	_AVFragmentedAssetMinderClassOnce sync.Once
)

func getAVFragmentedAssetMinderClass() AVFragmentedAssetMinderClass {
	_AVFragmentedAssetMinderClassOnce.Do(func() {
		_AVFragmentedAssetMinderClass = AVFragmentedAssetMinderClass{class: objc.GetClass("AVFragmentedAssetMinder")}
	})
	return _AVFragmentedAssetMinderClass
}

// GetAVFragmentedAssetMinderClass returns the class object for AVFragmentedAssetMinder.
func GetAVFragmentedAssetMinderClass() AVFragmentedAssetMinderClass {
	return getAVFragmentedAssetMinderClass()
}

type AVFragmentedAssetMinderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVFragmentedAssetMinderClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVFragmentedAssetMinderClass) Alloc() AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that periodically checks whether the system adds new fragments to
// a fragmented asset.
//
// # Creating an asset minder
//
//   - [AVFragmentedAssetMinder.InitWithAssetMindingInterval]: Creates a fragmented asset minder that monitors the specified asset at the indicated minding interval.
//
// # Configuring the minding interval
//
//   - [AVFragmentedAssetMinder.MindingInterval]: An interval that specifies when to perform a check for additional fragments.
//   - [AVFragmentedAssetMinder.SetMindingInterval]
//
// # Inspecting a fragment asset
//
//   - [AVFragmentedAssetMinder.Assets]: The minded array of fragmented assets.
//
// # Adding and removing fragmented assets
//
//   - [AVFragmentedAssetMinder.AddFragmentedAsset]: Adds a fragmented asset to the array of minded assets.
//   - [AVFragmentedAssetMinder.RemoveFragmentedAsset]: Removes a fragmented asset from the array of minded assets.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder
type AVFragmentedAssetMinder struct {
	objectivec.Object
}

// AVFragmentedAssetMinderFromID constructs a [AVFragmentedAssetMinder] from an objc.ID.
//
// An object that periodically checks whether the system adds new fragments to
// a fragmented asset.
func AVFragmentedAssetMinderFromID(id objc.ID) AVFragmentedAssetMinder {
	return AVFragmentedAssetMinder{objectivec.Object{ID: id}}
}

// NOTE: AVFragmentedAssetMinder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVFragmentedAssetMinder] class.
//
// # Creating an asset minder
//
//   - [IAVFragmentedAssetMinder.InitWithAssetMindingInterval]: Creates a fragmented asset minder that monitors the specified asset at the indicated minding interval.
//
// # Configuring the minding interval
//
//   - [IAVFragmentedAssetMinder.MindingInterval]: An interval that specifies when to perform a check for additional fragments.
//   - [IAVFragmentedAssetMinder.SetMindingInterval]
//
// # Inspecting a fragment asset
//
//   - [IAVFragmentedAssetMinder.Assets]: The minded array of fragmented assets.
//
// # Adding and removing fragmented assets
//
//   - [IAVFragmentedAssetMinder.AddFragmentedAsset]: Adds a fragmented asset to the array of minded assets.
//   - [IAVFragmentedAssetMinder.RemoveFragmentedAsset]: Removes a fragmented asset from the array of minded assets.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder
type IAVFragmentedAssetMinder interface {
	objectivec.IObject

	// Topic: Creating an asset minder

	// Creates a fragmented asset minder that monitors the specified asset at the indicated minding interval.
	InitWithAssetMindingInterval(asset IAVAsset, mindingInterval float64) AVFragmentedAssetMinder

	// Topic: Configuring the minding interval

	// An interval that specifies when to perform a check for additional fragments.
	MindingInterval() float64
	SetMindingInterval(value float64)

	// Topic: Inspecting a fragment asset

	// The minded array of fragmented assets.
	Assets() []AVAsset

	// Topic: Adding and removing fragmented assets

	// Adds a fragmented asset to the array of minded assets.
	AddFragmentedAsset(asset IAVAsset)
	// Removes a fragmented asset from the array of minded assets.
	RemoveFragmentedAsset(asset IAVAsset)
}

// Init initializes the instance.
func (f AVFragmentedAssetMinder) Init() AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f AVFragmentedAssetMinder) Autorelease() AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedAssetMinder creates a new AVFragmentedAssetMinder instance.
func NewAVFragmentedAssetMinder() AVFragmentedAssetMinder {
	class := getAVFragmentedAssetMinderClass()
	rv := objc.Send[AVFragmentedAssetMinder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a fragmented asset minder that monitors the specified asset at the
// indicated minding interval.
//
// asset: The fragmented asset added to the fragmented asset minder.
//
// mindingInterval: The amount of time between checking to see if the system appended
// additional fragments to the minded asset.
//
// # Return Value
//
// The new fragmented asset minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/init(asset:mindingInterval:)
func NewFragmentedAssetMinderWithAssetMindingInterval(asset IAVAsset, mindingInterval float64) AVFragmentedAssetMinder {
	instance := getAVFragmentedAssetMinderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:mindingInterval:"), asset, mindingInterval)
	return AVFragmentedAssetMinderFromID(rv)
}

// Creates a fragmented asset minder that monitors the specified asset at the
// indicated minding interval.
//
// asset: The fragmented asset added to the fragmented asset minder.
//
// mindingInterval: The amount of time between checking to see if the system appended
// additional fragments to the minded asset.
//
// # Return Value
//
// The new fragmented asset minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/init(asset:mindingInterval:)
func (f AVFragmentedAssetMinder) InitWithAssetMindingInterval(asset IAVAsset, mindingInterval float64) AVFragmentedAssetMinder {
	rv := objc.Send[AVFragmentedAssetMinder](f.ID, objc.Sel("initWithAsset:mindingInterval:"), asset, mindingInterval)
	return rv
}

// Adds a fragmented asset to the array of minded assets.
//
// asset: The fragmented asset to add to the minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/addFragmentedAsset(_:)
func (f AVFragmentedAssetMinder) AddFragmentedAsset(asset IAVAsset) {
	objc.Send[objc.ID](f.ID, objc.Sel("addFragmentedAsset:"), asset)
}

// Removes a fragmented asset from the array of minded assets.
//
// asset: The fragmented asset to remove from the minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/removeFragmentedAsset(_:)
func (f AVFragmentedAssetMinder) RemoveFragmentedAsset(asset IAVAsset) {
	objc.Send[objc.ID](f.ID, objc.Sel("removeFragmentedAsset:"), asset)
}

// Creates a fragmented asset minder containing the specified asset and
// minding interval.
//
// asset: The fragmented asset added to the fragmented asset minder.
//
// mindingInterval: The amount of time between checking to see if the system appended
// additional fragments to the minded asset.
//
// # Return Value
//
// The new fragmented asset minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/fragmentedAssetMinderWithAsset:mindingInterval:
func (_AVFragmentedAssetMinderClass AVFragmentedAssetMinderClass) FragmentedAssetMinderWithAssetMindingInterval(asset IAVAsset, mindingInterval float64) AVFragmentedAssetMinder {
	rv := objc.Send[objc.ID](objc.ID(_AVFragmentedAssetMinderClass.class), objc.Sel("fragmentedAssetMinderWithAsset:mindingInterval:"), asset, mindingInterval)
	return AVFragmentedAssetMinderFromID(rv)
}

// An interval that specifies when to perform a check for additional
// fragments.
//
// # Discussion
//
// The default interval in `10.0` seconds.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/mindingInterval
func (f AVFragmentedAssetMinder) MindingInterval() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("mindingInterval"))
	return rv
}
func (f AVFragmentedAssetMinder) SetMindingInterval(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setMindingInterval:"), value)
}

// The minded array of fragmented assets.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetMinder/assets
func (f AVFragmentedAssetMinder) Assets() []AVAsset {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("assets"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAsset {
		return AVAssetFromID(id)
	})
}
