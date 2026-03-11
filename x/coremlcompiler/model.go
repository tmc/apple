package coremlcompiler

// CoreML model protobuf structures. Only the fields needed for
// compilation are represented.

import "fmt"

// Model is the top-level CoreML model container.
// Proto: CoreML.Specification.Model
type Model struct {
	SpecVersion int32            // field 1
	Description ModelDescription // field 2
	MLProgram   *program         // field 502 (oneof)

	// Raw bytes of the description submessage, used for coremldata.bin.
	descriptionRaw []byte
}

// ModelDescription describes model inputs, outputs, and metadata.
// Proto: CoreML.Specification.ModelDescription
type ModelDescription struct {
	Inputs  []FeatureDescription // field 1
	Outputs []FeatureDescription // field 10
	States  []FeatureDescription // field 13
}

// FeatureDescription is a named, typed model feature.
// Proto: CoreML.Specification.FeatureDescription
type FeatureDescription struct {
	Name string       // field 1
	Type *FeatureType // field 3
}

// FeatureType describes a Core ML model feature type.
type FeatureType struct {
	MultiArrayType *ArrayFeatureType
	IsOptional     bool
}

// ArrayFeatureType describes a Core ML multi-array feature.
type ArrayFeatureType struct {
	Shape    []int64
	DataType ArrayDataType
}

// ArrayDataType identifies Core ML multi-array element types.
type ArrayDataType int32

const (
	ArrayDataTypeInvalid ArrayDataType = 0
	ArrayDataTypeFloat16 ArrayDataType = 65552
	ArrayDataTypeFloat32 ArrayDataType = 65568
	ArrayDataTypeDouble  ArrayDataType = 65600
	ArrayDataTypeInt32   ArrayDataType = 131104
)

// MIL program types.

// program is an MIL program.
// Proto: MILSpec.Program
type program struct {
	Version    int64                 // field 1
	Functions  map[string]*function  // field 2
	Attributes map[string]*value    // field 4
}

// function is an MIL function.
// Proto: MILSpec.Function
type function struct {
	Inputs               []namedValueType  // field 1
	OpSet                string            // field 2
	BlockSpecializations map[string]*block // field 3
	Attributes           map[string]*value // field 4
}

// block is a sequence of operations.
// Proto: MILSpec.Block
type block struct {
	Inputs     []namedValueType // field 1
	Outputs    []string         // field 2
	Operations []*operation     // field 3
}

// operation is a single MIL operation.
// Proto: MILSpec.Operation
type operation struct {
	Type       string               // field 1
	Inputs     map[string]*argument // field 2
	Outputs    []namedValueType     // field 3
	Blocks     []*block             // field 4
	Attributes map[string]*value    // field 5
}

// namedValueType is a (name, type) pair.
// Proto: MILSpec.NamedValueType
type namedValueType struct {
	Name string     // field 1
	Type *valueType // field 2
}

// valueType describes a MIL value's type.
// Proto: MILSpec.ValueType
type valueType struct {
	// Exactly one of these is set.
	TensorType *tensorType // field 1
	StateType  *stateType  // field 5
	// ListType, TupleType, DictionaryType omitted for now.
}

// tensorType describes a tensor's element type and shape.
// Proto: MILSpec.TensorType
type tensorType struct {
	DataType   DataType    // field 1
	Rank       int64       // field 2
	Dimensions []dimension // field 3
}

// stateType wraps a valueType for stateful operations.
// Proto: MILSpec.StateType
type stateType struct {
	WrappedType *valueType // field 1
}

// dimension is a tensor dimension (constant or unknown).
// Proto: MILSpec.Dimension
type dimension struct {
	Constant uint64 // from ConstantDimension.size (field 1.1)
	Unknown  bool   // true if UnknownDimension
}

// DataType identifies element types in MIL.
// Proto: MILSpec.DataType
type DataType int32

const (
	DataTypeFloat16  DataType = 10
	DataTypeFloat32  DataType = 11
	DataTypeFloat64  DataType = 12
	DataTypeBFloat16 DataType = 13
	DataTypeInt8     DataType = 21
	DataTypeInt16    DataType = 22
	DataTypeInt32    DataType = 23
	DataTypeInt64    DataType = 24
	DataTypeInt4     DataType = 25
	DataTypeUInt8    DataType = 31
	DataTypeUInt16   DataType = 32
	DataTypeUInt32   DataType = 33
	DataTypeUInt64   DataType = 34
	DataTypeBool     DataType = 1
	DataTypeString   DataType = 2
)

// String returns the MIL text name for the data type.
func (dt DataType) String() string {
	switch dt {
	case DataTypeFloat16:
		return "fp16"
	case DataTypeFloat32:
		return "fp32"
	case DataTypeFloat64:
		return "fp64"
	case DataTypeBFloat16:
		return "bf16"
	case DataTypeInt8:
		return "int8"
	case DataTypeInt16:
		return "int16"
	case DataTypeInt32:
		return "int32"
	case DataTypeInt64:
		return "int64"
	case DataTypeInt4:
		return "int4"
	case DataTypeUInt8:
		return "uint8"
	case DataTypeUInt16:
		return "uint16"
	case DataTypeUInt32:
		return "uint32"
	case DataTypeUInt64:
		return "uint64"
	case DataTypeBool:
		return "bool"
	case DataTypeString:
		return "string"
	default:
		return fmt.Sprintf("unknown(%d)", int(dt))
	}
}

// value is an MIL value (immediate or blob reference).
// Proto: MILSpec.Value
type value struct {
	Type *valueType // field 2

	// Exactly one of these is set.
	Immediate *immediateValue // field 3
	BlobFile  *blobFileValue  // field 5
}

// immediateValue holds inline constant data.
// Proto: MILSpec.ImmediateValue
type immediateValue struct {
	// Exactly one of these is set.
	Tensor *tensorValue // field 1
}

// tensorValue holds tensor data inline.
// Proto: MILSpec.TensorValue
type tensorValue struct {
	// Exactly one of these is set.
	Floats  []float32 // field 1
	Ints    []int32   // field 2
	Bools   []bool    // field 3
	Strings []string  // field 4
	Longs   []int64   // field 5
	Doubles []float64 // field 6
	Bytes   []byte    // field 7
}

// blobFileValue references weight data in a blob file.
// Proto: MILSpec.BlobFileValue
type blobFileValue struct {
	FileName string // field 1
	Offset   uint64 // field 2
}

// argument is an operation input (list of bindings).
// Proto: MILSpec.Argument
type argument struct {
	Bindings []binding // field 1
}

// binding is either a name reference or an inline value.
// Proto: MILSpec.Argument.Binding
type binding struct {
	Name  string // field 1 (oneof)
	Value *value // field 2 (oneof)
}
