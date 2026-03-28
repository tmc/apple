// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderBinary] class.
var (
	_GTAGX2ShaderBinaryClass     GTAGX2ShaderBinaryClass
	_GTAGX2ShaderBinaryClassOnce sync.Once
)

func getGTAGX2ShaderBinaryClass() GTAGX2ShaderBinaryClass {
	_GTAGX2ShaderBinaryClassOnce.Do(func() {
		_GTAGX2ShaderBinaryClass = GTAGX2ShaderBinaryClass{class: objc.GetClass("GTAGX2ShaderBinary")}
	})
	return _GTAGX2ShaderBinaryClass
}

// GetGTAGX2ShaderBinaryClass returns the class object for GTAGX2ShaderBinary.
func GetGTAGX2ShaderBinaryClass() GTAGX2ShaderBinaryClass {
	return getGTAGX2ShaderBinaryClass()
}

type GTAGX2ShaderBinaryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderBinaryClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderBinaryClass) Alloc() GTAGX2ShaderBinary {
	rv := objc.Send[GTAGX2ShaderBinary](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2ShaderBinary.AddBinaryRangeColumnAddrStartAddrEndFunctionNameFullPath]
//   - [GTAGX2ShaderBinary.AddCostForAddrCostDrawIdxIsALU]
//   - [GTAGX2ShaderBinary.AddString]
//   - [GTAGX2ShaderBinary.AddrEnd]
//   - [GTAGX2ShaderBinary.AddrStart]
//   - [GTAGX2ShaderBinary.AdjustLatencyForALUBlocksCount]
//   - [GTAGX2ShaderBinary.AnalysisResult]
//   - [GTAGX2ShaderBinary.SetAnalysisResult]
//   - [GTAGX2ShaderBinary.BinaryRanges]
//   - [GTAGX2ShaderBinary.CostForAddress]
//   - [GTAGX2ShaderBinary.CostForDrawAtIndex]
//   - [GTAGX2ShaderBinary.CostPercentageForDrawAtIndex]
//   - [GTAGX2ShaderBinary.Diassemblies]
//   - [GTAGX2ShaderBinary.DiassemblyAtAddress]
//   - [GTAGX2ShaderBinary.EncodeWithCoder]
//   - [GTAGX2ShaderBinary.FinalizeLineEntriesAndFixOffsetsAddrMin]
//   - [GTAGX2ShaderBinary.FullPath]
//   - [GTAGX2ShaderBinary.SetFullPath]
//   - [GTAGX2ShaderBinary.IsDylib]
//   - [GTAGX2ShaderBinary.SetIsDylib]
//   - [GTAGX2ShaderBinary.Key]
//   - [GTAGX2ShaderBinary.SetKey]
//   - [GTAGX2ShaderBinary.LineMax]
//   - [GTAGX2ShaderBinary.SetLineMax]
//   - [GTAGX2ShaderBinary.LineMin]
//   - [GTAGX2ShaderBinary.SetLineMin]
//   - [GTAGX2ShaderBinary.NumSamples]
//   - [GTAGX2ShaderBinary.StringFromIndex]
//   - [GTAGX2ShaderBinary.TotalCost]
//   - [GTAGX2ShaderBinary.Type]
//   - [GTAGX2ShaderBinary.SetType]
//   - [GTAGX2ShaderBinary.TypeName]
//   - [GTAGX2ShaderBinary.SetTypeName]
//   - [GTAGX2ShaderBinary.InitWithCoder]
//   - [GTAGX2ShaderBinary.InitWithKeyTypeTypeNameDylibAnalysisResult]
//   - [GTAGX2ShaderBinary.DebugDescription]
//   - [GTAGX2ShaderBinary.Description]
//   - [GTAGX2ShaderBinary.Hash]
//   - [GTAGX2ShaderBinary.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary
type GTAGX2ShaderBinary struct {
	objectivec.Object
}

// GTAGX2ShaderBinaryFromID constructs a [GTAGX2ShaderBinary] from an objc.ID.
func GTAGX2ShaderBinaryFromID(id objc.ID) GTAGX2ShaderBinary {
	return GTAGX2ShaderBinary{objectivec.Object{ID: id}}
}
// Ensure GTAGX2ShaderBinary implements IGTAGX2ShaderBinary.
var _ IGTAGX2ShaderBinary = GTAGX2ShaderBinary{}

// An interface definition for the [GTAGX2ShaderBinary] class.
//
// # Methods
//
//   - [IGTAGX2ShaderBinary.AddBinaryRangeColumnAddrStartAddrEndFunctionNameFullPath]
//   - [IGTAGX2ShaderBinary.AddCostForAddrCostDrawIdxIsALU]
//   - [IGTAGX2ShaderBinary.AddString]
//   - [IGTAGX2ShaderBinary.AddrEnd]
//   - [IGTAGX2ShaderBinary.AddrStart]
//   - [IGTAGX2ShaderBinary.AdjustLatencyForALUBlocksCount]
//   - [IGTAGX2ShaderBinary.AnalysisResult]
//   - [IGTAGX2ShaderBinary.SetAnalysisResult]
//   - [IGTAGX2ShaderBinary.BinaryRanges]
//   - [IGTAGX2ShaderBinary.CostForAddress]
//   - [IGTAGX2ShaderBinary.CostForDrawAtIndex]
//   - [IGTAGX2ShaderBinary.CostPercentageForDrawAtIndex]
//   - [IGTAGX2ShaderBinary.Diassemblies]
//   - [IGTAGX2ShaderBinary.DiassemblyAtAddress]
//   - [IGTAGX2ShaderBinary.EncodeWithCoder]
//   - [IGTAGX2ShaderBinary.FinalizeLineEntriesAndFixOffsetsAddrMin]
//   - [IGTAGX2ShaderBinary.FullPath]
//   - [IGTAGX2ShaderBinary.SetFullPath]
//   - [IGTAGX2ShaderBinary.IsDylib]
//   - [IGTAGX2ShaderBinary.SetIsDylib]
//   - [IGTAGX2ShaderBinary.Key]
//   - [IGTAGX2ShaderBinary.SetKey]
//   - [IGTAGX2ShaderBinary.LineMax]
//   - [IGTAGX2ShaderBinary.SetLineMax]
//   - [IGTAGX2ShaderBinary.LineMin]
//   - [IGTAGX2ShaderBinary.SetLineMin]
//   - [IGTAGX2ShaderBinary.NumSamples]
//   - [IGTAGX2ShaderBinary.StringFromIndex]
//   - [IGTAGX2ShaderBinary.TotalCost]
//   - [IGTAGX2ShaderBinary.Type]
//   - [IGTAGX2ShaderBinary.SetType]
//   - [IGTAGX2ShaderBinary.TypeName]
//   - [IGTAGX2ShaderBinary.SetTypeName]
//   - [IGTAGX2ShaderBinary.InitWithCoder]
//   - [IGTAGX2ShaderBinary.InitWithKeyTypeTypeNameDylibAnalysisResult]
//   - [IGTAGX2ShaderBinary.DebugDescription]
//   - [IGTAGX2ShaderBinary.Description]
//   - [IGTAGX2ShaderBinary.Hash]
//   - [IGTAGX2ShaderBinary.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary
type IGTAGX2ShaderBinary interface {
	objectivec.IObject

	// Topic: Methods

	AddBinaryRangeColumnAddrStartAddrEndFunctionNameFullPath(range_ int, column int, start uint32, end uint32, name objectivec.IObject, path objectivec.IObject) objectivec.IObject
	AddCostForAddrCostDrawIdxIsALU(addr uint32, cost float64, idx uint32, alu bool)
	AddString(string_ objectivec.IObject) uint64
	AddrEnd() uint32
	AddrStart() uint32
	AdjustLatencyForALUBlocksCount(aLUBlocks unsafe.Pointer, count uint64)
	AnalysisResult() IGTShaderProfilerBinaryAnalysisResult
	SetAnalysisResult(value IGTShaderProfilerBinaryAnalysisResult)
	BinaryRanges() foundation.INSArray
	CostForAddress(address uint32) float64
	CostForDrawAtIndex(index uint32) float64
	CostPercentageForDrawAtIndex(index uint32) float64
	Diassemblies() foundation.INSArray
	DiassemblyAtAddress(address uint32) objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	FinalizeLineEntriesAndFixOffsetsAddrMin(offsets uint32, min uint32)
	FullPath() string
	SetFullPath(value string)
	IsDylib() bool
	SetIsDylib(value bool)
	Key() string
	SetKey(value string)
	LineMax() int
	SetLineMax(value int)
	LineMin() int
	SetLineMin(value int)
	NumSamples() uint64
	StringFromIndex(index uint64) objectivec.IObject
	TotalCost() float64
	Type() uint32
	SetType(value uint32)
	TypeName() string
	SetTypeName(value string)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderBinary
	InitWithKeyTypeTypeNameDylibAnalysisResult(key objectivec.IObject, type_ uint32, name objectivec.IObject, dylib bool, result objectivec.IObject) GTAGX2ShaderBinary
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderBinary) Init() GTAGX2ShaderBinary {
	rv := objc.Send[GTAGX2ShaderBinary](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderBinary) Autorelease() GTAGX2ShaderBinary {
	rv := objc.Send[GTAGX2ShaderBinary](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderBinary creates a new GTAGX2ShaderBinary instance.
func NewGTAGX2ShaderBinary() GTAGX2ShaderBinary {
	class := getGTAGX2ShaderBinaryClass()
	rv := objc.Send[GTAGX2ShaderBinary](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/initWithCoder:
func NewGTAGX2ShaderBinaryWithCoder(coder objectivec.IObject) GTAGX2ShaderBinary {
	instance := getGTAGX2ShaderBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderBinaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/initWithKey:type:typeName:dylib:analysisResult:
func NewGTAGX2ShaderBinaryWithKeyTypeTypeNameDylibAnalysisResult(key objectivec.IObject, type_ uint32, name objectivec.IObject, dylib bool, result objectivec.IObject) GTAGX2ShaderBinary {
	instance := getGTAGX2ShaderBinaryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKey:type:typeName:dylib:analysisResult:"), key, type_, name, dylib, result)
	return GTAGX2ShaderBinaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/addBinaryRange:column:addrStart:addrEnd:functionName:fullPath:
func (g GTAGX2ShaderBinary) AddBinaryRangeColumnAddrStartAddrEndFunctionNameFullPath(range_ int, column int, start uint32, end uint32, name objectivec.IObject, path objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("addBinaryRange:column:addrStart:addrEnd:functionName:fullPath:"), range_, column, start, end, name, path)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/addCostForAddr:cost:drawIdx:isALU:
func (g GTAGX2ShaderBinary) AddCostForAddrCostDrawIdxIsALU(addr uint32, cost float64, idx uint32, alu bool) {
	objc.Send[objc.ID](g.ID, objc.Sel("addCostForAddr:cost:drawIdx:isALU:"), addr, cost, idx, alu)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/addString:
func (g GTAGX2ShaderBinary) AddString(string_ objectivec.IObject) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("addString:"), string_)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/adjustLatencyForALUBlocks:count:
func (g GTAGX2ShaderBinary) AdjustLatencyForALUBlocksCount(aLUBlocks unsafe.Pointer, count uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("adjustLatencyForALUBlocks:count:"), objc.CArray(aLUBlocks), count)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/costForAddress:
func (g GTAGX2ShaderBinary) CostForAddress(address uint32) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("costForAddress:"), address)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/costForDrawAtIndex:
func (g GTAGX2ShaderBinary) CostForDrawAtIndex(index uint32) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("costForDrawAtIndex:"), index)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/costPercentageForDrawAtIndex:
func (g GTAGX2ShaderBinary) CostPercentageForDrawAtIndex(index uint32) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("costPercentageForDrawAtIndex:"), index)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/diassemblyAtAddress:
func (g GTAGX2ShaderBinary) DiassemblyAtAddress(address uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("diassemblyAtAddress:"), address)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/encodeWithCoder:
func (g GTAGX2ShaderBinary) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/finalizeLineEntriesAndFixOffsets:addrMin:
func (g GTAGX2ShaderBinary) FinalizeLineEntriesAndFixOffsetsAddrMin(offsets uint32, min uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("finalizeLineEntriesAndFixOffsets:addrMin:"), offsets, min)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/setDiassemblies:
func (g GTAGX2ShaderBinary) SetDiassemblies(diassemblies objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("setDiassemblies:"), diassemblies)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/stringFromIndex:
func (g GTAGX2ShaderBinary) StringFromIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("stringFromIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/initWithCoder:
func (g GTAGX2ShaderBinary) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderBinary {
	rv := objc.Send[GTAGX2ShaderBinary](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/initWithKey:type:typeName:dylib:analysisResult:
func (g GTAGX2ShaderBinary) InitWithKeyTypeTypeNameDylibAnalysisResult(key objectivec.IObject, type_ uint32, name objectivec.IObject, dylib bool, result objectivec.IObject) GTAGX2ShaderBinary {
	rv := objc.Send[GTAGX2ShaderBinary](g.ID, objc.Sel("initWithKey:type:typeName:dylib:analysisResult:"), key, type_, name, dylib, result)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/supportsSecureCoding
func (_GTAGX2ShaderBinaryClass GTAGX2ShaderBinaryClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderBinaryClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/addrEnd
func (g GTAGX2ShaderBinary) AddrEnd() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("addrEnd"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/addrStart
func (g GTAGX2ShaderBinary) AddrStart() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("addrStart"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/analysisResult
func (g GTAGX2ShaderBinary) AnalysisResult() IGTShaderProfilerBinaryAnalysisResult {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("analysisResult"))
	return GTShaderProfilerBinaryAnalysisResultFromID(objc.ID(rv))
}
func (g GTAGX2ShaderBinary) SetAnalysisResult(value IGTShaderProfilerBinaryAnalysisResult) {
	objc.Send[struct{}](g.ID, objc.Sel("setAnalysisResult:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/binaryRanges
func (g GTAGX2ShaderBinary) BinaryRanges() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryRanges"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/debugDescription
func (g GTAGX2ShaderBinary) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/description
func (g GTAGX2ShaderBinary) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/diassemblies
func (g GTAGX2ShaderBinary) Diassemblies() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("diassemblies"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/fullPath
func (g GTAGX2ShaderBinary) FullPath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("fullPath"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinary) SetFullPath(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setFullPath:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/hash
func (g GTAGX2ShaderBinary) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/isDylib
func (g GTAGX2ShaderBinary) IsDylib() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isDylib"))
	return rv
}
func (g GTAGX2ShaderBinary) SetIsDylib(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setIsDylib:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/key
func (g GTAGX2ShaderBinary) Key() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("key"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinary) SetKey(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setKey:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/lineMax
func (g GTAGX2ShaderBinary) LineMax() int {
	rv := objc.Send[int](g.ID, objc.Sel("lineMax"))
	return rv
}
func (g GTAGX2ShaderBinary) SetLineMax(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setLineMax:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/lineMin
func (g GTAGX2ShaderBinary) LineMin() int {
	rv := objc.Send[int](g.ID, objc.Sel("lineMin"))
	return rv
}
func (g GTAGX2ShaderBinary) SetLineMin(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setLineMin:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/numSamples
func (g GTAGX2ShaderBinary) NumSamples() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numSamples"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/superclass
func (g GTAGX2ShaderBinary) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/totalCost
func (g GTAGX2ShaderBinary) TotalCost() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCost"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/type
func (g GTAGX2ShaderBinary) Type() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("type"))
	return rv
}
func (g GTAGX2ShaderBinary) SetType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setType:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinary/typeName
func (g GTAGX2ShaderBinary) TypeName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("typeName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinary) SetTypeName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setTypeName:"), objc.String(value))
}

