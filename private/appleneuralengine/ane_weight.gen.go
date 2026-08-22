// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (ac ANEWeightClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEWeightClass) Alloc() ANEWeight {
	rv := objc.SendIfResponds[ANEWeight](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

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
type ANEWeight struct {
	objectivec.Object
}

// ANEWeightFromID constructs a [ANEWeight] from an objc.ID.
func ANEWeightFromID(id objc.ID) ANEWeight {
	return ANEWeight{objectivec.Object{ID: id}}
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
type IANEWeight interface {
	objectivec.IObject

	// Topic: Methods

	SHACode() foundation.NSData
	EncodeWithCoder(coder foundation.INSCoder)
	SandboxExtension() string
	SetSandboxExtension(value string)
	UpdateWeightURL(url foundation.NSURL)
	WeightSymbol() string
	WeightURL() foundation.NSURL
	SetWeightURL(value foundation.NSURL)
	InitWithCoder(coder foundation.INSCoder) ANEWeight
	InitWithWeightSymbolAndURLWeightURL(url foundation.NSURL, url2 foundation.NSURL) ANEWeight
	InitWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension(urlsha objectivec.IObject, url foundation.NSURL, hACode objectivec.IObject, extension objectivec.IObject) ANEWeight
}

// Init initializes the instance.
func (a ANEWeight) Init() ANEWeight {
	rv := objc.SendIfResponds[ANEWeight](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEWeight) Autorelease() ANEWeight {
	rv := objc.SendIfResponds[ANEWeight](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEWeight creates a new ANEWeight instance.
func NewANEWeight() ANEWeight {
	class := getANEWeightClass()
	rv := objc.SendIfResponds[ANEWeight](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEWeightWithCoder(coder objectivec.IObject) ANEWeight {
	instance := getANEWeightClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEWeightFromID(rv)
}

func NewANEWeightWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension(urlsha objectivec.IObject, url foundation.NSURL, hACode objectivec.IObject, extension objectivec.IObject) ANEWeight {
	instance := getANEWeightClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithWeightSymbolAndURLSHA:weightURL:SHACode:sandboxExtension:"), urlsha, url, hACode, extension)
	return ANEWeightFromID(rv)
}

func NewANEWeightWithWeightSymbolAndURLWeightURL(url foundation.NSURL, url2 foundation.NSURL) ANEWeight {
	instance := getANEWeightClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithWeightSymbolAndURL:weightURL:"), url, url2)
	return ANEWeightFromID(rv)
}

func (a ANEWeight) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (a ANEWeight) UpdateWeightURL(url foundation.NSURL) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("updateWeightURL:"), url)
}
func (a ANEWeight) InitWithCoder(coder foundation.INSCoder) ANEWeight {
	rv := objc.SendIfResponds[ANEWeight](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a ANEWeight) InitWithWeightSymbolAndURLWeightURL(url foundation.NSURL, url2 foundation.NSURL) ANEWeight {
	rv := objc.SendIfResponds[ANEWeight](a.ID, objc.Sel("initWithWeightSymbolAndURL:weightURL:"), url, url2)
	return rv
}
func (a ANEWeight) InitWithWeightSymbolAndURLSHAWeightURLSHACodeSandboxExtension(urlsha objectivec.IObject, url foundation.NSURL, hACode objectivec.IObject, extension objectivec.IObject) ANEWeight {
	rv := objc.SendIfResponds[ANEWeight](a.ID, objc.Sel("initWithWeightSymbolAndURLSHA:weightURL:SHACode:sandboxExtension:"), urlsha, url, hACode, extension)
	return rv
}

func (_ANEWeightClass ANEWeightClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANEWeightClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_ANEWeightClass ANEWeightClass) WeightWithSymbolAndURLWeightURL(url foundation.NSURL, url2 foundation.NSURL) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEWeightClass.class), objc.Sel("weightWithSymbolAndURL:weightURL:"), url, url2)
	return objectivec.Object{ID: rv}
}
func (_ANEWeightClass ANEWeightClass) WeightWithSymbolAndURLSHAWeightURLSHACode(urlsha objectivec.IObject, url foundation.NSURL, hACode objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEWeightClass.class), objc.Sel("weightWithSymbolAndURLSHA:weightURL:SHACode:"), urlsha, url, hACode)
	return objectivec.Object{ID: rv}
}

func (a ANEWeight) SHACode() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("SHACode"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a ANEWeight) SandboxExtension() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sandboxExtension"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEWeight) SetSandboxExtension(value string) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setSandboxExtension:"), objc.String(value))
}
func (a ANEWeight) WeightSymbol() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("weightSymbol"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEWeight) WeightURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("weightURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a ANEWeight) SetWeightURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setWeightURL:"), value)
}
