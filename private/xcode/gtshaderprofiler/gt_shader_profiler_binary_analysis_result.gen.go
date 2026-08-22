// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerBinaryAnalysisResult] class.
var (
	_GTShaderProfilerBinaryAnalysisResultClass     GTShaderProfilerBinaryAnalysisResultClass
	_GTShaderProfilerBinaryAnalysisResultClassOnce sync.Once
)

func getGTShaderProfilerBinaryAnalysisResultClass() GTShaderProfilerBinaryAnalysisResultClass {
	_GTShaderProfilerBinaryAnalysisResultClassOnce.Do(func() {
		_GTShaderProfilerBinaryAnalysisResultClass = GTShaderProfilerBinaryAnalysisResultClass{class: objc.GetClass("GTShaderProfilerBinaryAnalysisResult")}
	})
	return _GTShaderProfilerBinaryAnalysisResultClass
}

// GetGTShaderProfilerBinaryAnalysisResultClass returns the class object for GTShaderProfilerBinaryAnalysisResult.
func GetGTShaderProfilerBinaryAnalysisResultClass() GTShaderProfilerBinaryAnalysisResultClass {
	return getGTShaderProfilerBinaryAnalysisResultClass()
}

type GTShaderProfilerBinaryAnalysisResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerBinaryAnalysisResultClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerBinaryAnalysisResultClass) Alloc() GTShaderProfilerBinaryAnalysisResult {
	rv := objc.SendIfResponds[GTShaderProfilerBinaryAnalysisResult](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerBinaryAnalysisResult.BinaryInfo]
//   - [GTShaderProfilerBinaryAnalysisResult.BinaryLocationCount]
//   - [GTShaderProfilerBinaryAnalysisResult.BinaryLocationData]
//   - [GTShaderProfilerBinaryAnalysisResult.BinaryLocations]
//   - [GTShaderProfilerBinaryAnalysisResult.BinaryRangeCount]
//   - [GTShaderProfilerBinaryAnalysisResult.BinaryRangeData]
//   - [GTShaderProfilerBinaryAnalysisResult.BinaryRanges]
//   - [GTShaderProfilerBinaryAnalysisResult.BranchTargetCount]
//   - [GTShaderProfilerBinaryAnalysisResult.BranchTargetData]
//   - [GTShaderProfilerBinaryAnalysisResult.BranchTargets]
//   - [GTShaderProfilerBinaryAnalysisResult.ClauseCount]
//   - [GTShaderProfilerBinaryAnalysisResult.ClauseData]
//   - [GTShaderProfilerBinaryAnalysisResult.Clauses]
//   - [GTShaderProfilerBinaryAnalysisResult.EncodeWithCoder]
//   - [GTShaderProfilerBinaryAnalysisResult.InstructionCount]
//   - [GTShaderProfilerBinaryAnalysisResult.InstructionData]
//   - [GTShaderProfilerBinaryAnalysisResult.Instructions]
//   - [GTShaderProfilerBinaryAnalysisResult.LastBinaryLocation]
//   - [GTShaderProfilerBinaryAnalysisResult.LastBinaryRange]
//   - [GTShaderProfilerBinaryAnalysisResult.LastBranchTarget]
//   - [GTShaderProfilerBinaryAnalysisResult.LastClause]
//   - [GTShaderProfilerBinaryAnalysisResult.LastInstruction]
//   - [GTShaderProfilerBinaryAnalysisResult.LastRegisterInfo]
//   - [GTShaderProfilerBinaryAnalysisResult.MaxOffset]
//   - [GTShaderProfilerBinaryAnalysisResult.RegisterInfo]
//   - [GTShaderProfilerBinaryAnalysisResult.RegisterInfoCount]
//   - [GTShaderProfilerBinaryAnalysisResult.RegisterInfoData]
//   - [GTShaderProfilerBinaryAnalysisResult.RegisterInfoOffsetForInstructionIndex]
//   - [GTShaderProfilerBinaryAnalysisResult.SetStrings]
//   - [GTShaderProfilerBinaryAnalysisResult.StringAtIndex]
//   - [GTShaderProfilerBinaryAnalysisResult.InitWithCoder]
//   - [GTShaderProfilerBinaryAnalysisResult.Version]
type GTShaderProfilerBinaryAnalysisResult struct {
	objectivec.Object
}

// GTShaderProfilerBinaryAnalysisResultFromID constructs a [GTShaderProfilerBinaryAnalysisResult] from an objc.ID.
func GTShaderProfilerBinaryAnalysisResultFromID(id objc.ID) GTShaderProfilerBinaryAnalysisResult {
	return GTShaderProfilerBinaryAnalysisResult{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerBinaryAnalysisResult implements IGTShaderProfilerBinaryAnalysisResult.
var _ IGTShaderProfilerBinaryAnalysisResult = GTShaderProfilerBinaryAnalysisResult{}

// An interface definition for the [GTShaderProfilerBinaryAnalysisResult] class.
//
// # Methods
//
//   - [IGTShaderProfilerBinaryAnalysisResult.BinaryInfo]
//   - [IGTShaderProfilerBinaryAnalysisResult.BinaryLocationCount]
//   - [IGTShaderProfilerBinaryAnalysisResult.BinaryLocationData]
//   - [IGTShaderProfilerBinaryAnalysisResult.BinaryLocations]
//   - [IGTShaderProfilerBinaryAnalysisResult.BinaryRangeCount]
//   - [IGTShaderProfilerBinaryAnalysisResult.BinaryRangeData]
//   - [IGTShaderProfilerBinaryAnalysisResult.BinaryRanges]
//   - [IGTShaderProfilerBinaryAnalysisResult.BranchTargetCount]
//   - [IGTShaderProfilerBinaryAnalysisResult.BranchTargetData]
//   - [IGTShaderProfilerBinaryAnalysisResult.BranchTargets]
//   - [IGTShaderProfilerBinaryAnalysisResult.ClauseCount]
//   - [IGTShaderProfilerBinaryAnalysisResult.ClauseData]
//   - [IGTShaderProfilerBinaryAnalysisResult.Clauses]
//   - [IGTShaderProfilerBinaryAnalysisResult.EncodeWithCoder]
//   - [IGTShaderProfilerBinaryAnalysisResult.InstructionCount]
//   - [IGTShaderProfilerBinaryAnalysisResult.InstructionData]
//   - [IGTShaderProfilerBinaryAnalysisResult.Instructions]
//   - [IGTShaderProfilerBinaryAnalysisResult.LastBinaryLocation]
//   - [IGTShaderProfilerBinaryAnalysisResult.LastBinaryRange]
//   - [IGTShaderProfilerBinaryAnalysisResult.LastBranchTarget]
//   - [IGTShaderProfilerBinaryAnalysisResult.LastClause]
//   - [IGTShaderProfilerBinaryAnalysisResult.LastInstruction]
//   - [IGTShaderProfilerBinaryAnalysisResult.LastRegisterInfo]
//   - [IGTShaderProfilerBinaryAnalysisResult.MaxOffset]
//   - [IGTShaderProfilerBinaryAnalysisResult.RegisterInfo]
//   - [IGTShaderProfilerBinaryAnalysisResult.RegisterInfoCount]
//   - [IGTShaderProfilerBinaryAnalysisResult.RegisterInfoData]
//   - [IGTShaderProfilerBinaryAnalysisResult.RegisterInfoOffsetForInstructionIndex]
//   - [IGTShaderProfilerBinaryAnalysisResult.SetStrings]
//   - [IGTShaderProfilerBinaryAnalysisResult.StringAtIndex]
//   - [IGTShaderProfilerBinaryAnalysisResult.InitWithCoder]
//   - [IGTShaderProfilerBinaryAnalysisResult.Version]
type IGTShaderProfilerBinaryAnalysisResult interface {
	objectivec.IObject

	// Topic: Methods

	BinaryInfo() unsafe.Pointer
	BinaryLocationCount() uint64
	BinaryLocationData() foundation.NSData
	BinaryLocations() objectivec.IObject
	BinaryRangeCount() uint64
	BinaryRangeData() foundation.NSData
	BinaryRanges() objectivec.IObject
	BranchTargetCount() uint64
	BranchTargetData() foundation.NSData
	BranchTargets() objectivec.IObject
	ClauseCount() uint64
	ClauseData() foundation.NSData
	Clauses() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	InstructionCount() uint64
	InstructionData() foundation.NSData
	Instructions() objectivec.IObject
	LastBinaryLocation() objectivec.IObject
	LastBinaryRange() objectivec.IObject
	LastBranchTarget() objectivec.IObject
	LastClause() objectivec.IObject
	LastInstruction() objectivec.IObject
	LastRegisterInfo() objectivec.IObject
	MaxOffset() uint64
	RegisterInfo() objectivec.IObject
	RegisterInfoCount() uint64
	RegisterInfoData() foundation.NSData
	RegisterInfoOffsetForInstructionIndex(index uint64) uint64
	SetStrings(strings objectivec.IObject)
	StringAtIndex(index uint64) objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) GTShaderProfilerBinaryAnalysisResult
	Version() uint32
}

// Init initializes the instance.
func (g GTShaderProfilerBinaryAnalysisResult) Init() GTShaderProfilerBinaryAnalysisResult {
	rv := objc.SendIfResponds[GTShaderProfilerBinaryAnalysisResult](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerBinaryAnalysisResult) Autorelease() GTShaderProfilerBinaryAnalysisResult {
	rv := objc.SendIfResponds[GTShaderProfilerBinaryAnalysisResult](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerBinaryAnalysisResult creates a new GTShaderProfilerBinaryAnalysisResult instance.
func NewGTShaderProfilerBinaryAnalysisResult() GTShaderProfilerBinaryAnalysisResult {
	class := getGTShaderProfilerBinaryAnalysisResultClass()
	rv := objc.SendIfResponds[GTShaderProfilerBinaryAnalysisResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTShaderProfilerBinaryAnalysisResultWithCoder(coder objectivec.IObject) GTShaderProfilerBinaryAnalysisResult {
	instance := getGTShaderProfilerBinaryAnalysisResultClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerBinaryAnalysisResultFromID(rv)
}

func (g GTShaderProfilerBinaryAnalysisResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (g GTShaderProfilerBinaryAnalysisResult) RegisterInfoOffsetForInstructionIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("registerInfoOffsetForInstructionIndex:"), index)
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) SetStrings(strings objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("setStrings:"), strings)
}
func (g GTShaderProfilerBinaryAnalysisResult) StringAtIndex(index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("stringAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) InitWithCoder(coder foundation.INSCoder) GTShaderProfilerBinaryAnalysisResult {
	rv := objc.SendIfResponds[GTShaderProfilerBinaryAnalysisResult](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_GTShaderProfilerBinaryAnalysisResultClass GTShaderProfilerBinaryAnalysisResultClass) AnalyzeBinaryTargetIndexIsaPrinter(binary objectivec.IObject, index int, printer objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_GTShaderProfilerBinaryAnalysisResultClass.class), objc.Sel("analyzeBinary:targetIndex:isaPrinter:"), binary, index, printer)
	return objectivec.Object{ID: rv}
}
func (_GTShaderProfilerBinaryAnalysisResultClass GTShaderProfilerBinaryAnalysisResultClass) CreateWithAnalysisResult(result unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_GTShaderProfilerBinaryAnalysisResultClass.class), objc.Sel("createWithAnalysisResult:"), result)
	return objectivec.Object{ID: rv}
}
func (_GTShaderProfilerBinaryAnalysisResultClass GTShaderProfilerBinaryAnalysisResultClass) DataFromArchivedDataURL(url foundation.NSURL) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_GTShaderProfilerBinaryAnalysisResultClass.class), objc.Sel("dataFromArchivedDataURL:"), url)
	return objectivec.Object{ID: rv}
}
func (_GTShaderProfilerBinaryAnalysisResultClass GTShaderProfilerBinaryAnalysisResultClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_GTShaderProfilerBinaryAnalysisResultClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (g GTShaderProfilerBinaryAnalysisResult) BinaryInfo() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("binaryInfo"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) BinaryLocationCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("binaryLocationCount"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) BinaryLocationData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaryLocationData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerBinaryAnalysisResult) BinaryLocations() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaryLocations"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) BinaryRangeCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("binaryRangeCount"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) BinaryRangeData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaryRangeData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerBinaryAnalysisResult) BinaryRanges() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binaryRanges"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) BranchTargetCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("branchTargetCount"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) BranchTargetData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("branchTargetData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerBinaryAnalysisResult) BranchTargets() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("branchTargets"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) ClauseCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("clauseCount"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) ClauseData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("clauseData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerBinaryAnalysisResult) Clauses() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("clauses"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) InstructionCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("instructionCount"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) InstructionData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("instructionData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerBinaryAnalysisResult) Instructions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("instructions"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) LastBinaryLocation() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("lastBinaryLocation"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) LastBinaryRange() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("lastBinaryRange"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) LastBranchTarget() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("lastBranchTarget"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) LastClause() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("lastClause"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) LastInstruction() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("lastInstruction"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) LastRegisterInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("lastRegisterInfo"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) MaxOffset() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("maxOffset"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) RegisterInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("registerInfo"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerBinaryAnalysisResult) RegisterInfoCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("registerInfoCount"))
	return rv
}
func (g GTShaderProfilerBinaryAnalysisResult) RegisterInfoData() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("registerInfoData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerBinaryAnalysisResult) Version() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("version"))
	return rv
}
