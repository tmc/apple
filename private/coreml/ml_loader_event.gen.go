// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLoaderEvent] class.
var (
	_MLLoaderEventClass     MLLoaderEventClass
	_MLLoaderEventClassOnce sync.Once
)

func getMLLoaderEventClass() MLLoaderEventClass {
	_MLLoaderEventClassOnce.Do(func() {
		_MLLoaderEventClass = MLLoaderEventClass{class: objc.GetClass("MLLoaderEvent")}
	})
	return _MLLoaderEventClass
}

// GetMLLoaderEventClass returns the class object for MLLoaderEvent.
func GetMLLoaderEventClass() MLLoaderEventClass {
	return getMLLoaderEventClass()
}

type MLLoaderEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLoaderEventClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLoaderEventClass) Alloc() MLLoaderEvent {
	rv := objc.Send[MLLoaderEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLLoaderEvent.BundleIdentifier]
//   - [MLLoaderEvent.SetBundleIdentifier]
//   - [MLLoaderEvent.CompilerVersion]
//   - [MLLoaderEvent.SetCompilerVersion]
//   - [MLLoaderEvent.ComputeUnits]
//   - [MLLoaderEvent.SetComputeUnits]
//   - [MLLoaderEvent.ContainsCustomLayer]
//   - [MLLoaderEvent.SetContainsCustomLayer]
//   - [MLLoaderEvent.DictionaryRepresentation]
//   - [MLLoaderEvent.ExtractAndSetModelDetailsFromArchive]
//   - [MLLoaderEvent.FirstPartyExecutable]
//   - [MLLoaderEvent.SetFirstPartyExecutable]
//   - [MLLoaderEvent.ModelDimension]
//   - [MLLoaderEvent.SetModelDimension]
//   - [MLLoaderEvent.ModelEngineType]
//   - [MLLoaderEvent.SetModelEngineType]
//   - [MLLoaderEvent.ModelHash]
//   - [MLLoaderEvent.SetModelHash]
//   - [MLLoaderEvent.ModelIsEncrypted]
//   - [MLLoaderEvent.SetModelIsEncrypted]
//   - [MLLoaderEvent.ModelLoadError]
//   - [MLLoaderEvent.SetModelLoadError]
//   - [MLLoaderEvent.ModelLoadTime]
//   - [MLLoaderEvent.SetModelLoadTime]
//   - [MLLoaderEvent.ModelName]
//   - [MLLoaderEvent.SetModelName]
//   - [MLLoaderEvent.ModelOrigin]
//   - [MLLoaderEvent.SetModelOrigin]
//   - [MLLoaderEvent.ModelProgramParsingError]
//   - [MLLoaderEvent.SetModelProgramParsingError]
//   - [MLLoaderEvent.ModelProgramValidationError]
//   - [MLLoaderEvent.SetModelProgramValidationError]
//   - [MLLoaderEvent.ModelType]
//   - [MLLoaderEvent.SetModelType]
//   - [MLLoaderEvent.ModelVersion]
//   - [MLLoaderEvent.SetModelVersion]
//   - [MLLoaderEvent.Name]
//   - [MLLoaderEvent.NnModelNetHash]
//   - [MLLoaderEvent.SetNnModelNetHash]
//   - [MLLoaderEvent.NnModelShapeHash]
//   - [MLLoaderEvent.SetNnModelShapeHash]
//   - [MLLoaderEvent.NnModelWeightsHash]
//   - [MLLoaderEvent.SetNnModelWeightsHash]
//   - [MLLoaderEvent.NumberFromCString]
//   - [MLLoaderEvent.ProcessName]
//   - [MLLoaderEvent.SetProcessName]
//   - [MLLoaderEvent.DebugDescription]
//   - [MLLoaderEvent.Description]
//   - [MLLoaderEvent.Hash]
//   - [MLLoaderEvent.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent
type MLLoaderEvent struct {
	objectivec.Object
}

// MLLoaderEventFromID constructs a [MLLoaderEvent] from an objc.ID.
func MLLoaderEventFromID(id objc.ID) MLLoaderEvent {
	return MLLoaderEvent{objectivec.Object{ID: id}}
}

// Ensure MLLoaderEvent implements IMLLoaderEvent.
var _ IMLLoaderEvent = MLLoaderEvent{}

// An interface definition for the [MLLoaderEvent] class.
//
// # Methods
//
//   - [IMLLoaderEvent.BundleIdentifier]
//   - [IMLLoaderEvent.SetBundleIdentifier]
//   - [IMLLoaderEvent.CompilerVersion]
//   - [IMLLoaderEvent.SetCompilerVersion]
//   - [IMLLoaderEvent.ComputeUnits]
//   - [IMLLoaderEvent.SetComputeUnits]
//   - [IMLLoaderEvent.ContainsCustomLayer]
//   - [IMLLoaderEvent.SetContainsCustomLayer]
//   - [IMLLoaderEvent.DictionaryRepresentation]
//   - [IMLLoaderEvent.ExtractAndSetModelDetailsFromArchive]
//   - [IMLLoaderEvent.FirstPartyExecutable]
//   - [IMLLoaderEvent.SetFirstPartyExecutable]
//   - [IMLLoaderEvent.ModelDimension]
//   - [IMLLoaderEvent.SetModelDimension]
//   - [IMLLoaderEvent.ModelEngineType]
//   - [IMLLoaderEvent.SetModelEngineType]
//   - [IMLLoaderEvent.ModelHash]
//   - [IMLLoaderEvent.SetModelHash]
//   - [IMLLoaderEvent.ModelIsEncrypted]
//   - [IMLLoaderEvent.SetModelIsEncrypted]
//   - [IMLLoaderEvent.ModelLoadError]
//   - [IMLLoaderEvent.SetModelLoadError]
//   - [IMLLoaderEvent.ModelLoadTime]
//   - [IMLLoaderEvent.SetModelLoadTime]
//   - [IMLLoaderEvent.ModelName]
//   - [IMLLoaderEvent.SetModelName]
//   - [IMLLoaderEvent.ModelOrigin]
//   - [IMLLoaderEvent.SetModelOrigin]
//   - [IMLLoaderEvent.ModelProgramParsingError]
//   - [IMLLoaderEvent.SetModelProgramParsingError]
//   - [IMLLoaderEvent.ModelProgramValidationError]
//   - [IMLLoaderEvent.SetModelProgramValidationError]
//   - [IMLLoaderEvent.ModelType]
//   - [IMLLoaderEvent.SetModelType]
//   - [IMLLoaderEvent.ModelVersion]
//   - [IMLLoaderEvent.SetModelVersion]
//   - [IMLLoaderEvent.Name]
//   - [IMLLoaderEvent.NnModelNetHash]
//   - [IMLLoaderEvent.SetNnModelNetHash]
//   - [IMLLoaderEvent.NnModelShapeHash]
//   - [IMLLoaderEvent.SetNnModelShapeHash]
//   - [IMLLoaderEvent.NnModelWeightsHash]
//   - [IMLLoaderEvent.SetNnModelWeightsHash]
//   - [IMLLoaderEvent.NumberFromCString]
//   - [IMLLoaderEvent.ProcessName]
//   - [IMLLoaderEvent.SetProcessName]
//   - [IMLLoaderEvent.DebugDescription]
//   - [IMLLoaderEvent.Description]
//   - [IMLLoaderEvent.Hash]
//   - [IMLLoaderEvent.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent
type IMLLoaderEvent interface {
	objectivec.IObject

	// Topic: Methods

	BundleIdentifier() string
	SetBundleIdentifier(value string)
	CompilerVersion() string
	SetCompilerVersion(value string)
	ComputeUnits() foundation.NSNumber
	SetComputeUnits(value foundation.NSNumber)
	ContainsCustomLayer() foundation.NSNumber
	SetContainsCustomLayer(value foundation.NSNumber)
	DictionaryRepresentation() foundation.INSDictionary
	ExtractAndSetModelDetailsFromArchive(archive unsafe.Pointer)
	FirstPartyExecutable() foundation.NSNumber
	SetFirstPartyExecutable(value foundation.NSNumber)
	ModelDimension() foundation.NSNumber
	SetModelDimension(value foundation.NSNumber)
	ModelEngineType() foundation.NSNumber
	SetModelEngineType(value foundation.NSNumber)
	ModelHash() string
	SetModelHash(value string)
	ModelIsEncrypted() foundation.NSNumber
	SetModelIsEncrypted(value foundation.NSNumber)
	ModelLoadError() foundation.NSNumber
	SetModelLoadError(value foundation.NSNumber)
	ModelLoadTime() foundation.NSNumber
	SetModelLoadTime(value foundation.NSNumber)
	ModelName() string
	SetModelName(value string)
	ModelOrigin() foundation.NSNumber
	SetModelOrigin(value foundation.NSNumber)
	ModelProgramParsingError() foundation.NSNumber
	SetModelProgramParsingError(value foundation.NSNumber)
	ModelProgramValidationError() foundation.NSNumber
	SetModelProgramValidationError(value foundation.NSNumber)
	ModelType() foundation.NSNumber
	SetModelType(value foundation.NSNumber)
	ModelVersion() string
	SetModelVersion(value string)
	Name() string
	NnModelNetHash() string
	SetNnModelNetHash(value string)
	NnModelShapeHash() string
	SetNnModelShapeHash(value string)
	NnModelWeightsHash() string
	SetNnModelWeightsHash(value string)
	NumberFromCString(cString string) objectivec.IObject
	ProcessName() string
	SetProcessName(value string)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (l MLLoaderEvent) Init() MLLoaderEvent {
	rv := objc.Send[MLLoaderEvent](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLoaderEvent) Autorelease() MLLoaderEvent {
	rv := objc.Send[MLLoaderEvent](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLoaderEvent creates a new MLLoaderEvent instance.
func NewMLLoaderEvent() MLLoaderEvent {
	class := getMLLoaderEventClass()
	rv := objc.Send[MLLoaderEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/extractAndSetModelDetailsFromArchive:
func (l MLLoaderEvent) ExtractAndSetModelDetailsFromArchive(archive unsafe.Pointer) {
	objc.Send[objc.ID](l.ID, objc.Sel("extractAndSetModelDetailsFromArchive:"), archive)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/numberFromCString:
func (l MLLoaderEvent) NumberFromCString(cString string) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("numberFromCString:"), unsafe.Pointer(unsafe.StringData(cString+"\x00")))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/bundleIdentifier
func (l MLLoaderEvent) BundleIdentifier() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetBundleIdentifier(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/compilerVersion
func (l MLLoaderEvent) CompilerVersion() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("compilerVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetCompilerVersion(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setCompilerVersion:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/computeUnits
func (l MLLoaderEvent) ComputeUnits() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("computeUnits"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetComputeUnits(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setComputeUnits:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/containsCustomLayer
func (l MLLoaderEvent) ContainsCustomLayer() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("containsCustomLayer"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetContainsCustomLayer(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setContainsCustomLayer:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/debugDescription
func (l MLLoaderEvent) DebugDescription() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/description
func (l MLLoaderEvent) Description() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/dictionaryRepresentation
func (l MLLoaderEvent) DictionaryRepresentation() foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("dictionaryRepresentation"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/firstPartyExecutable
func (l MLLoaderEvent) FirstPartyExecutable() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("firstPartyExecutable"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetFirstPartyExecutable(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setFirstPartyExecutable:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/hash
func (l MLLoaderEvent) Hash() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelDimension
func (l MLLoaderEvent) ModelDimension() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelDimension"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelDimension(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelDimension:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelEngineType
func (l MLLoaderEvent) ModelEngineType() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelEngineType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelEngineType(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelEngineType:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelHash
func (l MLLoaderEvent) ModelHash() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelHash"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetModelHash(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelHash:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelIsEncrypted
func (l MLLoaderEvent) ModelIsEncrypted() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelIsEncrypted"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelIsEncrypted(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelIsEncrypted:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelLoadError
func (l MLLoaderEvent) ModelLoadError() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelLoadError"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelLoadError(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelLoadError:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelLoadTime
func (l MLLoaderEvent) ModelLoadTime() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelLoadTime"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelLoadTime(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelLoadTime:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelName
func (l MLLoaderEvent) ModelName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetModelName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelOrigin
func (l MLLoaderEvent) ModelOrigin() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelOrigin"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelOrigin(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelOrigin:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelProgramParsingError
func (l MLLoaderEvent) ModelProgramParsingError() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelProgramParsingError"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelProgramParsingError(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelProgramParsingError:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelProgramValidationError
func (l MLLoaderEvent) ModelProgramValidationError() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelProgramValidationError"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelProgramValidationError(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelProgramValidationError:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelType
func (l MLLoaderEvent) ModelType() foundation.NSNumber {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (l MLLoaderEvent) SetModelType(value foundation.NSNumber) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelType:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/modelVersion
func (l MLLoaderEvent) ModelVersion() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetModelVersion(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelVersion:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/name
func (l MLLoaderEvent) Name() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/nnModelNetHash
func (l MLLoaderEvent) NnModelNetHash() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("nnModelNetHash"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetNnModelNetHash(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setNnModelNetHash:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/nnModelShapeHash
func (l MLLoaderEvent) NnModelShapeHash() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("nnModelShapeHash"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetNnModelShapeHash(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setNnModelShapeHash:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/nnModelWeightsHash
func (l MLLoaderEvent) NnModelWeightsHash() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("nnModelWeightsHash"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetNnModelWeightsHash(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setNnModelWeightsHash:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/processName
func (l MLLoaderEvent) ProcessName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("processName"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLoaderEvent) SetProcessName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setProcessName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLoaderEvent/superclass
func (l MLLoaderEvent) Superclass() objc.Class {
	rv := objc.Send[objc.Class](l.ID, objc.Sel("superclass"))
	return rv
}
