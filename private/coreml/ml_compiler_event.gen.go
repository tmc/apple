// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCompilerEvent] class.
var (
	_MLCompilerEventClass     MLCompilerEventClass
	_MLCompilerEventClassOnce sync.Once
)

func getMLCompilerEventClass() MLCompilerEventClass {
	_MLCompilerEventClassOnce.Do(func() {
		_MLCompilerEventClass = MLCompilerEventClass{class: objc.GetClass("MLCompilerEvent")}
	})
	return _MLCompilerEventClass
}

// GetMLCompilerEventClass returns the class object for MLCompilerEvent.
func GetMLCompilerEventClass() MLCompilerEventClass {
	return getMLCompilerEventClass()
}

type MLCompilerEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCompilerEventClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCompilerEventClass) Alloc() MLCompilerEvent {
	rv := objc.SendIfResponds[MLCompilerEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLCompilerEvent.CompilerVersion]
//   - [MLCompilerEvent.SetCompilerVersion]
//   - [MLCompilerEvent.DictionaryRepresentation]
//   - [MLCompilerEvent.MilUpgradeFailureReason]
//   - [MLCompilerEvent.SetMilUpgradeFailureReason]
//   - [MLCompilerEvent.MilUpgradeStatus]
//   - [MLCompilerEvent.SetMilUpgradeStatus]
//   - [MLCompilerEvent.ModelCompiledWithVersion]
//   - [MLCompilerEvent.SetModelCompiledWithVersion]
//   - [MLCompilerEvent.ModelHash]
//   - [MLCompilerEvent.SetModelHash]
//   - [MLCompilerEvent.ModelName]
//   - [MLCompilerEvent.SetModelName]
//   - [MLCompilerEvent.ModelOrigin]
//   - [MLCompilerEvent.SetModelOrigin]
//   - [MLCompilerEvent.ModelType]
//   - [MLCompilerEvent.SetModelType]
//   - [MLCompilerEvent.ModelVersion]
//   - [MLCompilerEvent.SetModelVersion]
//   - [MLCompilerEvent.Name]
//   - [MLCompilerEvent.DebugDescription]
//   - [MLCompilerEvent.Description]
//   - [MLCompilerEvent.Hash]
//   - [MLCompilerEvent.Superclass]
type MLCompilerEvent struct {
	objectivec.Object
}

// MLCompilerEventFromID constructs a [MLCompilerEvent] from an objc.ID.
func MLCompilerEventFromID(id objc.ID) MLCompilerEvent {
	return MLCompilerEvent{objectivec.Object{ID: id}}
}

// Ensure MLCompilerEvent implements IMLCompilerEvent.
var _ IMLCompilerEvent = MLCompilerEvent{}

// An interface definition for the [MLCompilerEvent] class.
//
// # Methods
//
//   - [IMLCompilerEvent.CompilerVersion]
//   - [IMLCompilerEvent.SetCompilerVersion]
//   - [IMLCompilerEvent.DictionaryRepresentation]
//   - [IMLCompilerEvent.MilUpgradeFailureReason]
//   - [IMLCompilerEvent.SetMilUpgradeFailureReason]
//   - [IMLCompilerEvent.MilUpgradeStatus]
//   - [IMLCompilerEvent.SetMilUpgradeStatus]
//   - [IMLCompilerEvent.ModelCompiledWithVersion]
//   - [IMLCompilerEvent.SetModelCompiledWithVersion]
//   - [IMLCompilerEvent.ModelHash]
//   - [IMLCompilerEvent.SetModelHash]
//   - [IMLCompilerEvent.ModelName]
//   - [IMLCompilerEvent.SetModelName]
//   - [IMLCompilerEvent.ModelOrigin]
//   - [IMLCompilerEvent.SetModelOrigin]
//   - [IMLCompilerEvent.ModelType]
//   - [IMLCompilerEvent.SetModelType]
//   - [IMLCompilerEvent.ModelVersion]
//   - [IMLCompilerEvent.SetModelVersion]
//   - [IMLCompilerEvent.Name]
//   - [IMLCompilerEvent.DebugDescription]
//   - [IMLCompilerEvent.Description]
//   - [IMLCompilerEvent.Hash]
//   - [IMLCompilerEvent.Superclass]
type IMLCompilerEvent interface {
	objectivec.IObject

	// Topic: Methods

	CompilerVersion() string
	SetCompilerVersion(value string)
	DictionaryRepresentation() foundation.INSDictionary
	MilUpgradeFailureReason() string
	SetMilUpgradeFailureReason(value string)
	MilUpgradeStatus() foundation.NSNumber
	SetMilUpgradeStatus(value foundation.NSNumber)
	ModelCompiledWithVersion() string
	SetModelCompiledWithVersion(value string)
	ModelHash() string
	SetModelHash(value string)
	ModelName() string
	SetModelName(value string)
	ModelOrigin() foundation.NSNumber
	SetModelOrigin(value foundation.NSNumber)
	ModelType() foundation.NSNumber
	SetModelType(value foundation.NSNumber)
	ModelVersion() string
	SetModelVersion(value string)
	Name() string
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLCompilerEvent) Init() MLCompilerEvent {
	rv := objc.SendIfResponds[MLCompilerEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLCompilerEvent) Autorelease() MLCompilerEvent {
	rv := objc.SendIfResponds[MLCompilerEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompilerEvent creates a new MLCompilerEvent instance.
func NewMLCompilerEvent() MLCompilerEvent {
	class := getMLCompilerEventClass()
	rv := objc.SendIfResponds[MLCompilerEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLCompilerEvent) CompilerVersion() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("compilerVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) SetCompilerVersion(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setCompilerVersion:"), objc.String(value))
}
func (m MLCompilerEvent) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) DictionaryRepresentation() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLCompilerEvent) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLCompilerEvent) MilUpgradeFailureReason() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("milUpgradeFailureReason"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) SetMilUpgradeFailureReason(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMilUpgradeFailureReason:"), objc.String(value))
}
func (m MLCompilerEvent) MilUpgradeStatus() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("milUpgradeStatus"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLCompilerEvent) SetMilUpgradeStatus(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMilUpgradeStatus:"), value)
}
func (m MLCompilerEvent) ModelCompiledWithVersion() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelCompiledWithVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) SetModelCompiledWithVersion(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelCompiledWithVersion:"), objc.String(value))
}
func (m MLCompilerEvent) ModelHash() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelHash"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) SetModelHash(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelHash:"), objc.String(value))
}
func (m MLCompilerEvent) ModelName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) SetModelName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelName:"), objc.String(value))
}
func (m MLCompilerEvent) ModelOrigin() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelOrigin"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLCompilerEvent) SetModelOrigin(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelOrigin:"), value)
}
func (m MLCompilerEvent) ModelType() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLCompilerEvent) SetModelType(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelType:"), value)
}
func (m MLCompilerEvent) ModelVersion() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) SetModelVersion(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelVersion:"), objc.String(value))
}
func (m MLCompilerEvent) Name() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCompilerEvent) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
