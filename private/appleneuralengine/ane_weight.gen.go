// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEWeight] class.
var (
	_ANEWeightClass     ANEWeightClass
	_ANEWeightClassOnce sync.Once
)

func getANEWeightClass() ANEWeightClass {
	_ANEWeightClassOnce.Do(func() {
		_ANEWeightClass = ANEWeightClass{class: objc.GetClass("_ANEWeight")}
	})
	return _ANEWeightClass
}

// GetANEWeightClass returns the class object for _ANEWeight.
func GetANEWeightClass() ANEWeightClass {
	return getANEWeightClass()
}

type ANEWeightClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEWeightClass) Alloc() ANEWeight {
	rv := objc.Send[ANEWeight](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [ANEWeight.SHACode]
//   - [ANEWeight.EncodeWithCoder]
//   - [ANEWeight.SandboxExtension]
//   - [ANEWeight.SetSandboxExtension]
//   - [ANEWeight.UpdateWeightURL]
//   - [ANEWeight.WeightSymbol]
//   - [ANEWeight.WeightURL]
//   - [ANEWeight.SetWeightURL]
//   - [ANEWeight.InitWithCoder]
//   - [ANEWeight.InitWithWeightSymbolAndURLWeightURL]
//   - [ANEWeight.InitWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight
type ANEWeight struct {
	objectivec.Object
}

// ANEWeightFromID constructs a [ANEWeight] from an objc.ID.
func ANEWeightFromID(id objc.ID) ANEWeight {
	return ANEWeight{objectivec.Object{id}}
}
// Ensure ANEWeight implements IANEWeight.
var _ IANEWeight = ANEWeight{}





// An interface definition for the [ANEWeight] class.
//
// # Methods
//
//   - [IANEWeight.SHACode]
//   - [IANEWeight.EncodeWithCoder]
//   - [IANEWeight.SandboxExtension]
//   - [IANEWeight.SetSandboxExtension]
//   - [IANEWeight.UpdateWeightURL]
//   - [IANEWeight.WeightSymbol]
//   - [IANEWeight.WeightURL]
//   - [IANEWeight.SetWeightURL]
//   - [IANEWeight.InitWithCoder]
//   - [IANEWeight.InitWithWeightSymbolAndURLWeightURL]
//   - [IANEWeight.InitWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight
type IANEWeight interface {
	objectivec.IObject

	// Topic: Methods

	SHACode() foundation.INSData
	EncodeWithCoder(coder objectivec.IObject)
	SandboxExtension() string
	SetSandboxExtension(value string)
	UpdateWeightURL(url foundation.INSURL)
	WeightSymbol() string
	WeightURL() foundation.INSURL
	SetWeightURL(value foundation.INSURL)
	InitWithCoder(coder objectivec.IObject) ANEWeight
	InitWithWeightSymbolAndURLWeightURL(url foundation.INSURL, url2 foundation.INSURL) ANEWeight
	InitWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension(urlsha objectivec.IObject, url foundation.INSURL, hACode objectivec.IObject, extension objectivec.IObject) ANEWeight
}





// Init initializes the instance.
func (a ANEWeight) Init() ANEWeight {
	rv := objc.Send[ANEWeight](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEWeight) Autorelease() ANEWeight {
	rv := objc.Send[ANEWeight](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEWeight creates a new ANEWeight instance.
func NewANEWeight() ANEWeight {
	class := getANEWeightClass()
	rv := objc.Send[ANEWeight](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/initWithCoder:
func NewANEWeightWithCoder(coder objectivec.IObject) ANEWeight {
	instance := getANEWeightClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEWeightFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/initWithWeightSymbolAndURLSHA:weightURL:SHACode:sandboxExtension:
func NewANEWeightWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension(urlsha objectivec.IObject, url foundation.INSURL, hACode objectivec.IObject, extension objectivec.IObject) ANEWeight {
	instance := getANEWeightClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWeightSymbolAndURLSHA:weightURL:SHACode:sandboxExtension:"), urlsha, url, hACode, extension)
	return ANEWeightFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/initWithWeightSymbolAndURL:weightURL:
func NewANEWeightWithWeightSymbolAndURLWeightURL(url foundation.INSURL, url2 foundation.INSURL) ANEWeight {
	instance := getANEWeightClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWeightSymbolAndURL:weightURL:"), url, url2)
	return ANEWeightFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/encodeWithCoder:
func (a ANEWeight) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/updateWeightURL:
func (a ANEWeight) UpdateWeightURL(url foundation.INSURL) {
	objc.Send[objc.ID](a.ID, objc.Sel("updateWeightURL:"), url)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/initWithCoder:
func (a ANEWeight) InitWithCoder(coder objectivec.IObject) ANEWeight {
	rv := objc.Send[ANEWeight](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/initWithWeightSymbolAndURL:weightURL:
func (a ANEWeight) InitWithWeightSymbolAndURLWeightURL(url foundation.INSURL, url2 foundation.INSURL) ANEWeight {
	rv := objc.Send[ANEWeight](a.ID, objc.Sel("initWithWeightSymbolAndURL:weightURL:"), url, url2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/initWithWeightSymbolAndURLSHA:weightURL:SHACode:sandboxExtension:
func (a ANEWeight) InitWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension(urlsha objectivec.IObject, url foundation.INSURL, hACode objectivec.IObject, extension objectivec.IObject) ANEWeight {
	rv := objc.Send[ANEWeight](a.ID, objc.Sel("initWithWeightSymbolAndURLSHA:weightURL:SHACode:sandboxExtension:"), urlsha, url, hACode, extension)
	return rv
}





// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/supportsSecureCoding
func (_ANEWeightClass ANEWeightClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEWeightClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/weightWithSymbolAndURL:weightURL:
func (_ANEWeightClass ANEWeightClass) WeightWithSymbolAndURLWeightURL(url foundation.INSURL, url2 foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEWeightClass.class), objc.Sel("weightWithSymbolAndURL:weightURL:"), url, url2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/weightWithSymbolAndURLSHA:weightURL:SHACode:
func (_ANEWeightClass ANEWeightClass) WeightWithSymbolAndURLSHAWeightURLSHACode(urlsha objectivec.IObject, url foundation.INSURL, hACode objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEWeightClass.class), objc.Sel("weightWithSymbolAndURLSHA:weightURL:SHACode:"), urlsha, url, hACode)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/SHACode
func (a ANEWeight) SHACode() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("SHACode"))
	return foundation.NSDataFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/sandboxExtension
func (a ANEWeight) SandboxExtension() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sandboxExtension"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEWeight) SetSandboxExtension(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSandboxExtension:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/weightSymbol
func (a ANEWeight) WeightSymbol() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("weightSymbol"))
	return foundation.NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEWeight/weightURL
func (a ANEWeight) WeightURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("weightURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a ANEWeight) SetWeightURL(value foundation.INSURL) {
	objc.Send[struct{}](a.ID, objc.Sel("setWeightURL:"), value)
}

















