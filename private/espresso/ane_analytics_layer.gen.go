// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEAnalyticsLayer] class.
var (
	_ANEAnalyticsLayerClass     ANEAnalyticsLayerClass
	_ANEAnalyticsLayerClassOnce sync.Once
)

func getANEAnalyticsLayerClass() ANEAnalyticsLayerClass {
	_ANEAnalyticsLayerClassOnce.Do(func() {
		_ANEAnalyticsLayerClass = ANEAnalyticsLayerClass{class: objc.GetClass("_ANEAnalyticsLayer")}
	})
	return _ANEAnalyticsLayerClass
}

// GetANEAnalyticsLayerClass returns the class object for _ANEAnalyticsLayer.
func GetANEAnalyticsLayerClass() ANEAnalyticsLayerClass {
	return getANEAnalyticsLayerClass()
}

type ANEAnalyticsLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEAnalyticsLayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEAnalyticsLayerClass) Alloc() ANEAnalyticsLayer {
	rv := objc.SendIfResponds[ANEAnalyticsLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEAnalyticsLayer.LayerName]
//   - [ANEAnalyticsLayer.Serialize]
//   - [ANEAnalyticsLayer.Weight]
//   - [ANEAnalyticsLayer.InitWithNameWeight]
type ANEAnalyticsLayer struct {
	objectivec.Object
}

// ANEAnalyticsLayerFromID constructs a [ANEAnalyticsLayer] from an objc.ID.
func ANEAnalyticsLayerFromID(id objc.ID) ANEAnalyticsLayer {
	return ANEAnalyticsLayer{objectivec.Object{ID: id}}
}

// Ensure ANEAnalyticsLayer implements IANEAnalyticsLayer.
var _ IANEAnalyticsLayer = ANEAnalyticsLayer{}

// An interface definition for the [ANEAnalyticsLayer] class.
//
// # Methods
//
//   - [IANEAnalyticsLayer.LayerName]
//   - [IANEAnalyticsLayer.Serialize]
//   - [IANEAnalyticsLayer.Weight]
//   - [IANEAnalyticsLayer.InitWithNameWeight]
type IANEAnalyticsLayer interface {
	objectivec.IObject

	// Topic: Methods

	LayerName() string
	Serialize() objectivec.IObject
	Weight() foundation.NSNumber
	InitWithNameWeight(name objectivec.IObject, weight objectivec.IObject) ANEAnalyticsLayer
}

// Init initializes the instance.
func (a ANEAnalyticsLayer) Init() ANEAnalyticsLayer {
	rv := objc.SendIfResponds[ANEAnalyticsLayer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEAnalyticsLayer) Autorelease() ANEAnalyticsLayer {
	rv := objc.SendIfResponds[ANEAnalyticsLayer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEAnalyticsLayer creates a new ANEAnalyticsLayer instance.
func NewANEAnalyticsLayer() ANEAnalyticsLayer {
	class := getANEAnalyticsLayerClass()
	rv := objc.SendIfResponds[ANEAnalyticsLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEAnalyticsLayerWithNameWeight(name objectivec.IObject, weight objectivec.IObject) ANEAnalyticsLayer {
	instance := getANEAnalyticsLayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:weight:"), name, weight)
	return ANEAnalyticsLayerFromID(rv)
}

func (a ANEAnalyticsLayer) Serialize() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("serialize"))
	return objectivec.Object{ID: rv}
}
func (a ANEAnalyticsLayer) InitWithNameWeight(name objectivec.IObject, weight objectivec.IObject) ANEAnalyticsLayer {
	rv := objc.SendIfResponds[ANEAnalyticsLayer](a.ID, objc.Sel("initWithName:weight:"), name, weight)
	return rv
}

func (_ANEAnalyticsLayerClass ANEAnalyticsLayerClass) ObjectWithNameWeight(name objectivec.IObject, weight objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEAnalyticsLayerClass.class), objc.Sel("objectWithName:weight:"), name, weight)
	return objectivec.Object{ID: rv}
}

func (a ANEAnalyticsLayer) LayerName() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("layerName"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEAnalyticsLayer) Weight() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("weight"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
