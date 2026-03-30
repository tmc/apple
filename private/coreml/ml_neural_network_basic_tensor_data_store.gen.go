// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkBasicTensorDataStore] class.
var (
	_MLNeuralNetworkBasicTensorDataStoreClass     MLNeuralNetworkBasicTensorDataStoreClass
	_MLNeuralNetworkBasicTensorDataStoreClassOnce sync.Once
)

func getMLNeuralNetworkBasicTensorDataStoreClass() MLNeuralNetworkBasicTensorDataStoreClass {
	_MLNeuralNetworkBasicTensorDataStoreClassOnce.Do(func() {
		_MLNeuralNetworkBasicTensorDataStoreClass = MLNeuralNetworkBasicTensorDataStoreClass{class: objc.GetClass("_MLNeuralNetworkBasicTensorDataStore")}
	})
	return _MLNeuralNetworkBasicTensorDataStoreClass
}

// GetMLNeuralNetworkBasicTensorDataStoreClass returns the class object for _MLNeuralNetworkBasicTensorDataStore.
func GetMLNeuralNetworkBasicTensorDataStoreClass() MLNeuralNetworkBasicTensorDataStoreClass {
	return getMLNeuralNetworkBasicTensorDataStoreClass()
}

type MLNeuralNetworkBasicTensorDataStoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkBasicTensorDataStoreClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkBasicTensorDataStoreClass) Alloc() MLNeuralNetworkBasicTensorDataStore {
	rv := objc.Send[MLNeuralNetworkBasicTensorDataStore](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNeuralNetworkBasicTensorDataStore.Data]
//   - [MLNeuralNetworkBasicTensorDataStore.TensorDataForOffsetExpectedLength]
//   - [MLNeuralNetworkBasicTensorDataStore.WriteToFileError]
//   - [MLNeuralNetworkBasicTensorDataStore.InitWithContentsOfFileError]
//   - [MLNeuralNetworkBasicTensorDataStore.InitWithData]
//   - [MLNeuralNetworkBasicTensorDataStore.DebugDescription]
//   - [MLNeuralNetworkBasicTensorDataStore.Description]
//   - [MLNeuralNetworkBasicTensorDataStore.Hash]
//   - [MLNeuralNetworkBasicTensorDataStore.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore
type MLNeuralNetworkBasicTensorDataStore struct {
	objectivec.Object
}

// MLNeuralNetworkBasicTensorDataStoreFromID constructs a [MLNeuralNetworkBasicTensorDataStore] from an objc.ID.
func MLNeuralNetworkBasicTensorDataStoreFromID(id objc.ID) MLNeuralNetworkBasicTensorDataStore {
	return MLNeuralNetworkBasicTensorDataStore{objectivec.Object{ID: id}}
}

// Ensure MLNeuralNetworkBasicTensorDataStore implements IMLNeuralNetworkBasicTensorDataStore.
var _ IMLNeuralNetworkBasicTensorDataStore = MLNeuralNetworkBasicTensorDataStore{}

// An interface definition for the [MLNeuralNetworkBasicTensorDataStore] class.
//
// # Methods
//
//   - [IMLNeuralNetworkBasicTensorDataStore.Data]
//   - [IMLNeuralNetworkBasicTensorDataStore.TensorDataForOffsetExpectedLength]
//   - [IMLNeuralNetworkBasicTensorDataStore.WriteToFileError]
//   - [IMLNeuralNetworkBasicTensorDataStore.InitWithContentsOfFileError]
//   - [IMLNeuralNetworkBasicTensorDataStore.InitWithData]
//   - [IMLNeuralNetworkBasicTensorDataStore.DebugDescription]
//   - [IMLNeuralNetworkBasicTensorDataStore.Description]
//   - [IMLNeuralNetworkBasicTensorDataStore.Hash]
//   - [IMLNeuralNetworkBasicTensorDataStore.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore
type IMLNeuralNetworkBasicTensorDataStore interface {
	objectivec.IObject

	// Topic: Methods

	Data() foundation.NSMutableData
	TensorDataForOffsetExpectedLength(offset uint64, length uint64) objectivec.IObject
	WriteToFileError(file objectivec.IObject) (bool, error)
	InitWithContentsOfFileError(file objectivec.IObject) (MLNeuralNetworkBasicTensorDataStore, error)
	InitWithData(data objectivec.IObject) MLNeuralNetworkBasicTensorDataStore
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLNeuralNetworkBasicTensorDataStore) Init() MLNeuralNetworkBasicTensorDataStore {
	rv := objc.Send[MLNeuralNetworkBasicTensorDataStore](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralNetworkBasicTensorDataStore) Autorelease() MLNeuralNetworkBasicTensorDataStore {
	rv := objc.Send[MLNeuralNetworkBasicTensorDataStore](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkBasicTensorDataStore creates a new MLNeuralNetworkBasicTensorDataStore instance.
func NewMLNeuralNetworkBasicTensorDataStore() MLNeuralNetworkBasicTensorDataStore {
	class := getMLNeuralNetworkBasicTensorDataStoreClass()
	rv := objc.Send[MLNeuralNetworkBasicTensorDataStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/initWithContentsOfFile:error:
func NewMLNeuralNetworkBasicTensorDataStoreWithContentsOfFileError(file objectivec.IObject) (MLNeuralNetworkBasicTensorDataStore, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkBasicTensorDataStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkBasicTensorDataStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkBasicTensorDataStoreFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/initWithData:
func NewMLNeuralNetworkBasicTensorDataStoreWithData(data objectivec.IObject) MLNeuralNetworkBasicTensorDataStore {
	instance := getMLNeuralNetworkBasicTensorDataStoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return MLNeuralNetworkBasicTensorDataStoreFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/tensorDataForOffset:expectedLength:
func (m MLNeuralNetworkBasicTensorDataStore) TensorDataForOffsetExpectedLength(offset uint64, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("tensorDataForOffset:expectedLength:"), offset, length)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/writeToFile:error:
func (m MLNeuralNetworkBasicTensorDataStore) WriteToFileError(file objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("writeToFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToFile:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/initWithContentsOfFile:error:
func (m MLNeuralNetworkBasicTensorDataStore) InitWithContentsOfFileError(file objectivec.IObject) (MLNeuralNetworkBasicTensorDataStore, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithContentsOfFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkBasicTensorDataStore{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkBasicTensorDataStoreFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/initWithData:
func (m MLNeuralNetworkBasicTensorDataStore) InitWithData(data objectivec.IObject) MLNeuralNetworkBasicTensorDataStore {
	rv := objc.Send[MLNeuralNetworkBasicTensorDataStore](m.ID, objc.Sel("initWithData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/data
func (m MLNeuralNetworkBasicTensorDataStore) Data() foundation.NSMutableData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("data"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/debugDescription
func (m MLNeuralNetworkBasicTensorDataStore) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/description
func (m MLNeuralNetworkBasicTensorDataStore) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/hash
func (m MLNeuralNetworkBasicTensorDataStore) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLNeuralNetworkBasicTensorDataStore/superclass
func (m MLNeuralNetworkBasicTensorDataStore) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
