// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderDiassembly] class.
var (
	_GTAGX2ShaderDiassemblyClass     GTAGX2ShaderDiassemblyClass
	_GTAGX2ShaderDiassemblyClassOnce sync.Once
)

func getGTAGX2ShaderDiassemblyClass() GTAGX2ShaderDiassemblyClass {
	_GTAGX2ShaderDiassemblyClassOnce.Do(func() {
		_GTAGX2ShaderDiassemblyClass = GTAGX2ShaderDiassemblyClass{class: objc.GetClass("GTAGX2ShaderDiassembly")}
	})
	return _GTAGX2ShaderDiassemblyClass
}

// GetGTAGX2ShaderDiassemblyClass returns the class object for GTAGX2ShaderDiassembly.
func GetGTAGX2ShaderDiassemblyClass() GTAGX2ShaderDiassemblyClass {
	return getGTAGX2ShaderDiassemblyClass()
}

type GTAGX2ShaderDiassemblyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderDiassemblyClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderDiassemblyClass) Alloc() GTAGX2ShaderDiassembly {
	rv := objc.Send[GTAGX2ShaderDiassembly](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2ShaderDiassembly.Address]
//   - [GTAGX2ShaderDiassembly.SetAddress]
//   - [GTAGX2ShaderDiassembly.Binary]
//   - [GTAGX2ShaderDiassembly.Cost]
//   - [GTAGX2ShaderDiassembly.Diassembly]
//   - [GTAGX2ShaderDiassembly.SetDiassembly]
//   - [GTAGX2ShaderDiassembly.EncodeWithCoder]
//   - [GTAGX2ShaderDiassembly.Opcode]
//   - [GTAGX2ShaderDiassembly.SetOpcode]
//   - [GTAGX2ShaderDiassembly.OpcodeMask]
//   - [GTAGX2ShaderDiassembly.SetOpcodeMask]
//   - [GTAGX2ShaderDiassembly.OpcodeType]
//   - [GTAGX2ShaderDiassembly.SetOpcodeType]
//   - [GTAGX2ShaderDiassembly.InitWithCoder]
//   - [GTAGX2ShaderDiassembly.InitWithOpcodeOpcodeTypeOpcodeMaskAddressDiassemblyBinary]
//   - [GTAGX2ShaderDiassembly.DebugDescription]
//   - [GTAGX2ShaderDiassembly.Description]
//   - [GTAGX2ShaderDiassembly.Hash]
//   - [GTAGX2ShaderDiassembly.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly
type GTAGX2ShaderDiassembly struct {
	objectivec.Object
}

// GTAGX2ShaderDiassemblyFromID constructs a [GTAGX2ShaderDiassembly] from an objc.ID.
func GTAGX2ShaderDiassemblyFromID(id objc.ID) GTAGX2ShaderDiassembly {
	return GTAGX2ShaderDiassembly{objectivec.Object{ID: id}}
}

// Ensure GTAGX2ShaderDiassembly implements IGTAGX2ShaderDiassembly.
var _ IGTAGX2ShaderDiassembly = GTAGX2ShaderDiassembly{}

// An interface definition for the [GTAGX2ShaderDiassembly] class.
//
// # Methods
//
//   - [IGTAGX2ShaderDiassembly.Address]
//   - [IGTAGX2ShaderDiassembly.SetAddress]
//   - [IGTAGX2ShaderDiassembly.Binary]
//   - [IGTAGX2ShaderDiassembly.Cost]
//   - [IGTAGX2ShaderDiassembly.Diassembly]
//   - [IGTAGX2ShaderDiassembly.SetDiassembly]
//   - [IGTAGX2ShaderDiassembly.EncodeWithCoder]
//   - [IGTAGX2ShaderDiassembly.Opcode]
//   - [IGTAGX2ShaderDiassembly.SetOpcode]
//   - [IGTAGX2ShaderDiassembly.OpcodeMask]
//   - [IGTAGX2ShaderDiassembly.SetOpcodeMask]
//   - [IGTAGX2ShaderDiassembly.OpcodeType]
//   - [IGTAGX2ShaderDiassembly.SetOpcodeType]
//   - [IGTAGX2ShaderDiassembly.InitWithCoder]
//   - [IGTAGX2ShaderDiassembly.InitWithOpcodeOpcodeTypeOpcodeMaskAddressDiassemblyBinary]
//   - [IGTAGX2ShaderDiassembly.DebugDescription]
//   - [IGTAGX2ShaderDiassembly.Description]
//   - [IGTAGX2ShaderDiassembly.Hash]
//   - [IGTAGX2ShaderDiassembly.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly
type IGTAGX2ShaderDiassembly interface {
	objectivec.IObject

	// Topic: Methods

	Address() uint32
	SetAddress(value uint32)
	Binary() objectivec.IObject
	Cost() float64
	Diassembly() string
	SetDiassembly(value string)
	EncodeWithCoder(coder foundation.INSCoder)
	Opcode() uint32
	SetOpcode(value uint32)
	OpcodeMask() uint32
	SetOpcodeMask(value uint32)
	OpcodeType() uint32
	SetOpcodeType(value uint32)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderDiassembly
	InitWithOpcodeOpcodeTypeOpcodeMaskAddressDiassemblyBinary(opcode uint32, type_ uint32, mask uint32, address uint32, diassembly objectivec.IObject, binary objectivec.IObject) GTAGX2ShaderDiassembly
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderDiassembly) Init() GTAGX2ShaderDiassembly {
	rv := objc.Send[GTAGX2ShaderDiassembly](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderDiassembly) Autorelease() GTAGX2ShaderDiassembly {
	rv := objc.Send[GTAGX2ShaderDiassembly](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderDiassembly creates a new GTAGX2ShaderDiassembly instance.
func NewGTAGX2ShaderDiassembly() GTAGX2ShaderDiassembly {
	class := getGTAGX2ShaderDiassemblyClass()
	rv := objc.Send[GTAGX2ShaderDiassembly](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/initWithCoder:
func NewGTAGX2ShaderDiassemblyWithCoder(coder objectivec.IObject) GTAGX2ShaderDiassembly {
	instance := getGTAGX2ShaderDiassemblyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderDiassemblyFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/initWithOpcode:opcodeType:opcodeMask:address:diassembly:binary:
func NewGTAGX2ShaderDiassemblyWithOpcodeOpcodeTypeOpcodeMaskAddressDiassemblyBinary(opcode uint32, type_ uint32, mask uint32, address uint32, diassembly objectivec.IObject, binary objectivec.IObject) GTAGX2ShaderDiassembly {
	instance := getGTAGX2ShaderDiassemblyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOpcode:opcodeType:opcodeMask:address:diassembly:binary:"), opcode, type_, mask, address, diassembly, binary)
	return GTAGX2ShaderDiassemblyFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/encodeWithCoder:
func (g GTAGX2ShaderDiassembly) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/initWithCoder:
func (g GTAGX2ShaderDiassembly) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderDiassembly {
	rv := objc.Send[GTAGX2ShaderDiassembly](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/initWithOpcode:opcodeType:opcodeMask:address:diassembly:binary:
func (g GTAGX2ShaderDiassembly) InitWithOpcodeOpcodeTypeOpcodeMaskAddressDiassemblyBinary(opcode uint32, type_ uint32, mask uint32, address uint32, diassembly objectivec.IObject, binary objectivec.IObject) GTAGX2ShaderDiassembly {
	rv := objc.Send[GTAGX2ShaderDiassembly](g.ID, objc.Sel("initWithOpcode:opcodeType:opcodeMask:address:diassembly:binary:"), opcode, type_, mask, address, diassembly, binary)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/supportsSecureCoding
func (_GTAGX2ShaderDiassemblyClass GTAGX2ShaderDiassemblyClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderDiassemblyClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/address
func (g GTAGX2ShaderDiassembly) Address() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("address"))
	return rv
}
func (g GTAGX2ShaderDiassembly) SetAddress(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setAddress:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/binary
func (g GTAGX2ShaderDiassembly) Binary() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/cost
func (g GTAGX2ShaderDiassembly) Cost() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("cost"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/debugDescription
func (g GTAGX2ShaderDiassembly) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/description
func (g GTAGX2ShaderDiassembly) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/diassembly
func (g GTAGX2ShaderDiassembly) Diassembly() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("diassembly"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderDiassembly) SetDiassembly(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setDiassembly:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/hash
func (g GTAGX2ShaderDiassembly) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/opcode
func (g GTAGX2ShaderDiassembly) Opcode() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("opcode"))
	return rv
}
func (g GTAGX2ShaderDiassembly) SetOpcode(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setOpcode:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/opcodeMask
func (g GTAGX2ShaderDiassembly) OpcodeMask() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("opcodeMask"))
	return rv
}
func (g GTAGX2ShaderDiassembly) SetOpcodeMask(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setOpcodeMask:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/opcodeType
func (g GTAGX2ShaderDiassembly) OpcodeType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("opcodeType"))
	return rv
}
func (g GTAGX2ShaderDiassembly) SetOpcodeType(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setOpcodeType:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderDiassembly/superclass
func (g GTAGX2ShaderDiassembly) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
