// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEAnalyticsProcedure] class.
var (
	_ANEAnalyticsProcedureClass     ANEAnalyticsProcedureClass
	_ANEAnalyticsProcedureClassOnce sync.Once
)

func getANEAnalyticsProcedureClass() ANEAnalyticsProcedureClass {
	_ANEAnalyticsProcedureClassOnce.Do(func() {
		_ANEAnalyticsProcedureClass = ANEAnalyticsProcedureClass{class: objc.GetClass("_ANEAnalyticsProcedure")}
	})
	return _ANEAnalyticsProcedureClass
}

// GetANEAnalyticsProcedureClass returns the class object for _ANEAnalyticsProcedure.
func GetANEAnalyticsProcedureClass() ANEAnalyticsProcedureClass {
	return getANEAnalyticsProcedureClass()
}

type ANEAnalyticsProcedureClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEAnalyticsProcedureClass) Alloc() ANEAnalyticsProcedure {
	rv := objc.Send[ANEAnalyticsProcedure](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEAnalyticsProcedure.GroupInfo]
//   - [ANEAnalyticsProcedure.Identifier]
//   - [ANEAnalyticsProcedure.ProcedureMetrics]
//   - [ANEAnalyticsProcedure.Serialize]
//   - [ANEAnalyticsProcedure.InitWithGroupsProcedureMetricsIndentifier]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure
type ANEAnalyticsProcedure struct {
	objectivec.Object
}

// ANEAnalyticsProcedureFromID constructs a [ANEAnalyticsProcedure] from an objc.ID.
func ANEAnalyticsProcedureFromID(id objc.ID) ANEAnalyticsProcedure {
	return ANEAnalyticsProcedure{objectivec.Object{ID: id}}
}
// Ensure ANEAnalyticsProcedure implements IANEAnalyticsProcedure.
var _ IANEAnalyticsProcedure = ANEAnalyticsProcedure{}

// An interface definition for the [ANEAnalyticsProcedure] class.
//
// # Methods
//
//   - [IANEAnalyticsProcedure.GroupInfo]
//   - [IANEAnalyticsProcedure.Identifier]
//   - [IANEAnalyticsProcedure.ProcedureMetrics]
//   - [IANEAnalyticsProcedure.Serialize]
//   - [IANEAnalyticsProcedure.InitWithGroupsProcedureMetricsIndentifier]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure
type IANEAnalyticsProcedure interface {
	objectivec.IObject

	// Topic: Methods

	GroupInfo() foundation.INSArray
	Identifier() string
	ProcedureMetrics() foundation.INSDictionary
	Serialize() objectivec.IObject
	InitWithGroupsProcedureMetricsIndentifier(groups objectivec.IObject, metrics objectivec.IObject, indentifier objectivec.IObject) ANEAnalyticsProcedure
}

// Init initializes the instance.
func (a ANEAnalyticsProcedure) Init() ANEAnalyticsProcedure {
	rv := objc.Send[ANEAnalyticsProcedure](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEAnalyticsProcedure) Autorelease() ANEAnalyticsProcedure {
	rv := objc.Send[ANEAnalyticsProcedure](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEAnalyticsProcedure creates a new ANEAnalyticsProcedure instance.
func NewANEAnalyticsProcedure() ANEAnalyticsProcedure {
	class := getANEAnalyticsProcedureClass()
	rv := objc.Send[ANEAnalyticsProcedure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure/initWithGroups:procedureMetrics:indentifier:
func NewANEAnalyticsProcedureWithGroupsProcedureMetricsIndentifier(groups objectivec.IObject, metrics objectivec.IObject, indentifier objectivec.IObject) ANEAnalyticsProcedure {
	instance := getANEAnalyticsProcedureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGroups:procedureMetrics:indentifier:"), groups, metrics, indentifier)
	return ANEAnalyticsProcedureFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure/serialize
func (a ANEAnalyticsProcedure) Serialize() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("serialize"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure/initWithGroups:procedureMetrics:indentifier:
func (a ANEAnalyticsProcedure) InitWithGroupsProcedureMetricsIndentifier(groups objectivec.IObject, metrics objectivec.IObject, indentifier objectivec.IObject) ANEAnalyticsProcedure {
	rv := objc.Send[ANEAnalyticsProcedure](a.ID, objc.Sel("initWithGroups:procedureMetrics:indentifier:"), groups, metrics, indentifier)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure/objectWithGroups:procedureMetrics:indentifier:
func (_ANEAnalyticsProcedureClass ANEAnalyticsProcedureClass) ObjectWithGroupsProcedureMetricsIndentifier(groups objectivec.IObject, metrics objectivec.IObject, indentifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEAnalyticsProcedureClass.class), objc.Sel("objectWithGroups:procedureMetrics:indentifier:"), groups, metrics, indentifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure/groupInfo
func (a ANEAnalyticsProcedure) GroupInfo() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("groupInfo"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure/identifier
func (a ANEAnalyticsProcedure) Identifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsProcedure/procedureMetrics
func (a ANEAnalyticsProcedure) ProcedureMetrics() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureMetrics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

