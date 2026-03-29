// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETImageFolderDataProvider] class.
var (
	_ETImageFolderDataProviderClass     ETImageFolderDataProviderClass
	_ETImageFolderDataProviderClassOnce sync.Once
)

func getETImageFolderDataProviderClass() ETImageFolderDataProviderClass {
	_ETImageFolderDataProviderClassOnce.Do(func() {
		_ETImageFolderDataProviderClass = ETImageFolderDataProviderClass{class: objc.GetClass("ETImageFolderDataProvider")}
	})
	return _ETImageFolderDataProviderClass
}

// GetETImageFolderDataProviderClass returns the class object for ETImageFolderDataProvider.
func GetETImageFolderDataProviderClass() ETImageFolderDataProviderClass {
	return getETImageFolderDataProviderClass()
}

type ETImageFolderDataProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETImageFolderDataProviderClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETImageFolderDataProviderClass) Alloc() ETImageFolderDataProvider {
	rv := objc.Send[ETImageFolderDataProvider](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETImageFolderDataProvider.DataPointAtIndexError]
//   - [ETImageFolderDataProvider.NumberOfDataPoints]
//   - [ETImageFolderDataProvider.PrepareForEpoch]
//   - [ETImageFolderDataProvider.InitWithFolderForImageTensorAndLabelTensorShuffleBeforeEachEpochShuffleRandomSeedWithImagePreprocessParams]
//   - [ETImageFolderDataProvider.DebugDescription]
//   - [ETImageFolderDataProvider.Description]
//   - [ETImageFolderDataProvider.Hash]
//   - [ETImageFolderDataProvider.Superclass]
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider
type ETImageFolderDataProvider struct {
	objectivec.Object
}

// ETImageFolderDataProviderFromID constructs a [ETImageFolderDataProvider] from an objc.ID.
func ETImageFolderDataProviderFromID(id objc.ID) ETImageFolderDataProvider {
	return ETImageFolderDataProvider{objectivec.Object{ID: id}}
}
// Ensure ETImageFolderDataProvider implements IETImageFolderDataProvider.
var _ IETImageFolderDataProvider = ETImageFolderDataProvider{}

// An interface definition for the [ETImageFolderDataProvider] class.
//
// # Methods
//
//   - [IETImageFolderDataProvider.DataPointAtIndexError]
//   - [IETImageFolderDataProvider.NumberOfDataPoints]
//   - [IETImageFolderDataProvider.PrepareForEpoch]
//   - [IETImageFolderDataProvider.InitWithFolderForImageTensorAndLabelTensorShuffleBeforeEachEpochShuffleRandomSeedWithImagePreprocessParams]
//   - [IETImageFolderDataProvider.DebugDescription]
//   - [IETImageFolderDataProvider.Description]
//   - [IETImageFolderDataProvider.Hash]
//   - [IETImageFolderDataProvider.Superclass]
//
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider
type IETImageFolderDataProvider interface {
	objectivec.IObject

	// Topic: Methods

	DataPointAtIndexError(index uint64) (objectivec.IObject, error)
	NumberOfDataPoints() uint64
	PrepareForEpoch()
	InitWithFolderForImageTensorAndLabelTensorShuffleBeforeEachEpochShuffleRandomSeedWithImagePreprocessParams(folder objectivec.IObject, tensor objectivec.IObject, tensor2 objectivec.IObject, epoch bool, seed objectivec.IObject, params objectivec.IObject) ETImageFolderDataProvider
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e ETImageFolderDataProvider) Init() ETImageFolderDataProvider {
	rv := objc.Send[ETImageFolderDataProvider](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETImageFolderDataProvider) Autorelease() ETImageFolderDataProvider {
	rv := objc.Send[ETImageFolderDataProvider](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETImageFolderDataProvider creates a new ETImageFolderDataProvider instance.
func NewETImageFolderDataProvider() ETImageFolderDataProvider {
	class := getETImageFolderDataProviderClass()
	rv := objc.Send[ETImageFolderDataProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/initWithFolder:forImageTensor:andLabelTensor:shuffleBeforeEachEpoch:shuffleRandomSeed:withImagePreprocessParams:
func NewETImageFolderDataProviderWithFolderForImageTensorAndLabelTensorShuffleBeforeEachEpochShuffleRandomSeedWithImagePreprocessParams(folder objectivec.IObject, tensor objectivec.IObject, tensor2 objectivec.IObject, epoch bool, seed objectivec.IObject, params objectivec.IObject) ETImageFolderDataProvider {
	instance := getETImageFolderDataProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFolder:forImageTensor:andLabelTensor:shuffleBeforeEachEpoch:shuffleRandomSeed:withImagePreprocessParams:"), folder, tensor, tensor2, epoch, seed, params)
	return ETImageFolderDataProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/dataPointAtIndex:error:
func (e ETImageFolderDataProvider) DataPointAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataPointAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/numberOfDataPoints
func (e ETImageFolderDataProvider) NumberOfDataPoints() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/prepareForEpoch
func (e ETImageFolderDataProvider) PrepareForEpoch() {
	objc.Send[objc.ID](e.ID, objc.Sel("prepareForEpoch"))
}
//
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/initWithFolder:forImageTensor:andLabelTensor:shuffleBeforeEachEpoch:shuffleRandomSeed:withImagePreprocessParams:
func (e ETImageFolderDataProvider) InitWithFolderForImageTensorAndLabelTensorShuffleBeforeEachEpochShuffleRandomSeedWithImagePreprocessParams(folder objectivec.IObject, tensor objectivec.IObject, tensor2 objectivec.IObject, epoch bool, seed objectivec.IObject, params objectivec.IObject) ETImageFolderDataProvider {
	rv := objc.Send[ETImageFolderDataProvider](e.ID, objc.Sel("initWithFolder:forImageTensor:andLabelTensor:shuffleBeforeEachEpoch:shuffleRandomSeed:withImagePreprocessParams:"), folder, tensor, tensor2, epoch, seed, params)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/debugDescription
func (e ETImageFolderDataProvider) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/description
func (e ETImageFolderDataProvider) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/hash
func (e ETImageFolderDataProvider) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETImageFolderDataProvider/superclass
func (e ETImageFolderDataProvider) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}

