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
	rv := objc.Send[MLCompilerEvent](objc.ID(mc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent
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
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent
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
	Superclass() objc.Class
}

// Init initializes the instance.
func (c MLCompilerEvent) Init() MLCompilerEvent {
	rv := objc.Send[MLCompilerEvent](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCompilerEvent) Autorelease() MLCompilerEvent {
	rv := objc.Send[MLCompilerEvent](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompilerEvent creates a new MLCompilerEvent instance.
func NewMLCompilerEvent() MLCompilerEvent {
	class := getMLCompilerEventClass()
	rv := objc.Send[MLCompilerEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/compilerVersion
func (c MLCompilerEvent) CompilerVersion() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("compilerVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerEvent) SetCompilerVersion(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCompilerVersion:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/debugDescription
func (c MLCompilerEvent) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/description
func (c MLCompilerEvent) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/dictionaryRepresentation
func (c MLCompilerEvent) DictionaryRepresentation() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dictionaryRepresentation"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/hash
func (c MLCompilerEvent) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/milUpgradeFailureReason
func (c MLCompilerEvent) MilUpgradeFailureReason() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("milUpgradeFailureReason"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerEvent) SetMilUpgradeFailureReason(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setMilUpgradeFailureReason:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/milUpgradeStatus
func (c MLCompilerEvent) MilUpgradeStatus() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("milUpgradeStatus"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c MLCompilerEvent) SetMilUpgradeStatus(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setMilUpgradeStatus:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/modelCompiledWithVersion
func (c MLCompilerEvent) ModelCompiledWithVersion() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelCompiledWithVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerEvent) SetModelCompiledWithVersion(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelCompiledWithVersion:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/modelHash
func (c MLCompilerEvent) ModelHash() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelHash"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerEvent) SetModelHash(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelHash:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/modelName
func (c MLCompilerEvent) ModelName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerEvent) SetModelName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/modelOrigin
func (c MLCompilerEvent) ModelOrigin() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelOrigin"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c MLCompilerEvent) SetModelOrigin(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelOrigin:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/modelType
func (c MLCompilerEvent) ModelType() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (c MLCompilerEvent) SetModelType(value foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelType:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/modelVersion
func (c MLCompilerEvent) ModelVersion() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (c MLCompilerEvent) SetModelVersion(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelVersion:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/name
func (c MLCompilerEvent) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerEvent/superclass
func (c MLCompilerEvent) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
