// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEInMemoryModelDescriptor] class.
var (
	_ANEInMemoryModelDescriptorClass     ANEInMemoryModelDescriptorClass
	_ANEInMemoryModelDescriptorClassOnce sync.Once
)

func getANEInMemoryModelDescriptorClass() ANEInMemoryModelDescriptorClass {
	_ANEInMemoryModelDescriptorClassOnce.Do(func() {
		_ANEInMemoryModelDescriptorClass = ANEInMemoryModelDescriptorClass{class: objc.GetClass("_ANEInMemoryModelDescriptor")}
	})
	return _ANEInMemoryModelDescriptorClass
}

// GetANEInMemoryModelDescriptorClass returns the class object for _ANEInMemoryModelDescriptor.
func GetANEInMemoryModelDescriptorClass() ANEInMemoryModelDescriptorClass {
	return getANEInMemoryModelDescriptorClass()
}

type ANEInMemoryModelDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEInMemoryModelDescriptorClass) Alloc() ANEInMemoryModelDescriptor {
	rv := objc.Send[ANEInMemoryModelDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEInMemoryModelDescriptor.HexStringIdentifier]
//   - [ANEInMemoryModelDescriptor.IsEqualToInMemoryModelDescriptor]
//   - [ANEInMemoryModelDescriptor.IsMILModel]
//   - [ANEInMemoryModelDescriptor.NetworkText]
//   - [ANEInMemoryModelDescriptor.NetworkTextHash]
//   - [ANEInMemoryModelDescriptor.OptionsPlist]
//   - [ANEInMemoryModelDescriptor.OptionsPlistHash]
//   - [ANEInMemoryModelDescriptor.Weights]
//   - [ANEInMemoryModelDescriptor.WeightsHash]
//   - [ANEInMemoryModelDescriptor.InitWithNetworkTextWeightsOptionsPlistIsMILModel]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor
type ANEInMemoryModelDescriptor struct {
	objectivec.Object
}

// ANEInMemoryModelDescriptorFromID constructs a [ANEInMemoryModelDescriptor] from an objc.ID.
func ANEInMemoryModelDescriptorFromID(id objc.ID) ANEInMemoryModelDescriptor {
	return ANEInMemoryModelDescriptor{objectivec.Object{id}}
}
// Ensure ANEInMemoryModelDescriptor implements IANEInMemoryModelDescriptor.
var _ IANEInMemoryModelDescriptor = ANEInMemoryModelDescriptor{}

// An interface definition for the [ANEInMemoryModelDescriptor] class.
//
// # Methods
//
//   - [IANEInMemoryModelDescriptor.HexStringIdentifier]
//   - [IANEInMemoryModelDescriptor.IsEqualToInMemoryModelDescriptor]
//   - [IANEInMemoryModelDescriptor.IsMILModel]
//   - [IANEInMemoryModelDescriptor.NetworkText]
//   - [IANEInMemoryModelDescriptor.NetworkTextHash]
//   - [IANEInMemoryModelDescriptor.OptionsPlist]
//   - [IANEInMemoryModelDescriptor.OptionsPlistHash]
//   - [IANEInMemoryModelDescriptor.Weights]
//   - [IANEInMemoryModelDescriptor.WeightsHash]
//   - [IANEInMemoryModelDescriptor.InitWithNetworkTextWeightsOptionsPlistIsMILModel]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor
type IANEInMemoryModelDescriptor interface {
	objectivec.IObject

	// Topic: Methods

	HexStringIdentifier() objectivec.IObject
	IsEqualToInMemoryModelDescriptor(descriptor objectivec.IObject) bool
	IsMILModel() bool
	NetworkText() foundation.INSData
	NetworkTextHash() string
	OptionsPlist() foundation.INSData
	OptionsPlistHash() string
	Weights() foundation.INSDictionary
	WeightsHash() string
	InitWithNetworkTextWeightsOptionsPlistIsMILModel(text objectivec.IObject, weights objectivec.IObject, plist objectivec.IObject, mILModel bool) ANEInMemoryModelDescriptor
}

// Init initializes the instance.
func (a ANEInMemoryModelDescriptor) Init() ANEInMemoryModelDescriptor {
	rv := objc.Send[ANEInMemoryModelDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEInMemoryModelDescriptor) Autorelease() ANEInMemoryModelDescriptor {
	rv := objc.Send[ANEInMemoryModelDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEInMemoryModelDescriptor creates a new ANEInMemoryModelDescriptor instance.
func NewANEInMemoryModelDescriptor() ANEInMemoryModelDescriptor {
	class := getANEInMemoryModelDescriptorClass()
	rv := objc.Send[ANEInMemoryModelDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/initWithNetworkText:weights:optionsPlist:isMILModel:
func NewANEInMemoryModelDescriptorWithNetworkTextWeightsOptionsPlistIsMILModel(text objectivec.IObject, weights objectivec.IObject, plist objectivec.IObject, mILModel bool) ANEInMemoryModelDescriptor {
	instance := getANEInMemoryModelDescriptorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkText:weights:optionsPlist:isMILModel:"), text, weights, plist, mILModel)
	return ANEInMemoryModelDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/hexStringIdentifier
func (a ANEInMemoryModelDescriptor) HexStringIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("hexStringIdentifier"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/isEqualToInMemoryModelDescriptor:
func (a ANEInMemoryModelDescriptor) IsEqualToInMemoryModelDescriptor(descriptor objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEqualToInMemoryModelDescriptor:"), descriptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/initWithNetworkText:weights:optionsPlist:isMILModel:
func (a ANEInMemoryModelDescriptor) InitWithNetworkTextWeightsOptionsPlistIsMILModel(text objectivec.IObject, weights objectivec.IObject, plist objectivec.IObject, mILModel bool) ANEInMemoryModelDescriptor {
	rv := objc.Send[ANEInMemoryModelDescriptor](a.ID, objc.Sel("initWithNetworkText:weights:optionsPlist:isMILModel:"), text, weights, plist, mILModel)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/modelWithMILText:weights:optionsPlist:
func (_ANEInMemoryModelDescriptorClass ANEInMemoryModelDescriptorClass) ModelWithMILTextWeightsOptionsPlist(mILText objectivec.IObject, weights objectivec.IObject, plist objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEInMemoryModelDescriptorClass.class), objc.Sel("modelWithMILText:weights:optionsPlist:"), mILText, weights, plist)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/modelWithNetworkDescription:weights:optionsPlist:
func (_ANEInMemoryModelDescriptorClass ANEInMemoryModelDescriptorClass) ModelWithNetworkDescriptionWeightsOptionsPlist(description objectivec.IObject, weights objectivec.IObject, plist objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEInMemoryModelDescriptorClass.class), objc.Sel("modelWithNetworkDescription:weights:optionsPlist:"), description, weights, plist)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/isMILModel
func (a ANEInMemoryModelDescriptor) IsMILModel() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isMILModel"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/networkText
func (a ANEInMemoryModelDescriptor) NetworkText() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("networkText"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/networkTextHash
func (a ANEInMemoryModelDescriptor) NetworkTextHash() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("networkTextHash"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/optionsPlist
func (a ANEInMemoryModelDescriptor) OptionsPlist() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("optionsPlist"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/optionsPlistHash
func (a ANEInMemoryModelDescriptor) OptionsPlistHash() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("optionsPlistHash"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/weights
func (a ANEInMemoryModelDescriptor) Weights() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("weights"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModelDescriptor/weightsHash
func (a ANEInMemoryModelDescriptor) WeightsHash() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("weightsHash"))
	return foundation.NSStringFromID(rv).String()
}

