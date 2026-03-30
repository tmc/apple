package coremlcompiler

// CoreML model protobuf structures. Only the fields needed for
// compilation are represented.

import "fmt"

// Model is the top-level CoreML model container.
// Proto: CoreML.Specification.Model
type Model struct {
	SpecVersion int32            // field 1
	Description ModelDescription // field 2
	MLProgram   *Program         // field 502 (oneof)

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

// MIL Program types.

// Program is an MIL Program.
// Proto: MILSpec.Program
type Program struct {
	Version    int64                // field 1
	Functions  map[string]*Function // field 2
	Attributes map[string]*Value    // field 4
}

// Function is an MIL Function.
// Proto: MILSpec.Function
type Function struct {
	Inputs               []NamedValueType  // field 1
	OpSet                string            // field 2
	BlockSpecializations map[string]*Block // field 3
	Attributes           map[string]*Value // field 4
}

// Block is a sequence of operations.
// Proto: MILSpec.Block
type Block struct {
	Inputs     []NamedValueType // field 1
	Outputs    []string         // field 2
	Operations []*Operation     // field 3
}

// Operation is a single MIL Operation.
// Proto: MILSpec.Operation
type Operation struct {
	Type       string               // field 1
	Inputs     map[string]*Argument // field 2
	Outputs    []NamedValueType     // field 3
	Blocks     []*Block             // field 4
	Attributes map[string]*Value    // field 5
}

// NamedValueType is a (name, type) pair.
// Proto: MILSpec.NamedValueType
type NamedValueType struct {
	Name string     // field 1
	Type *ValueType // field 2
}

// ValueType describes a MIL value's type.
// Proto: MILSpec.ValueType
type ValueType struct {
	// Exactly one of these is set.
	TensorType *TensorType // field 1
	StateType  *StateType  // field 5
	// ListType, TupleType, DictionaryType omitted for now.
}

// TensorType describes a tensor's element type and shape.
// Proto: MILSpec.TensorType
type TensorType struct {
	DataType   DataType    // field 1
	Rank       int64       // field 2
	Dimensions []Dimension // field 3
}

// StateType wraps a ValueType for stateful operations.
// Proto: MILSpec.StateType
type StateType struct {
	WrappedType *ValueType // field 1
}

// Dimension is a tensor Dimension (constant or unknown).
// Proto: MILSpec.Dimension
type Dimension struct {
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
type Value struct {
	Type *ValueType // field 2

	// Exactly one of these is set.
	Immediate *ImmediateValue // field 3
	BlobFile  *BlobFileValue  // field 5
}

// ImmediateValue holds inline constant data.
// Proto: MILSpec.ImmediateValue
type ImmediateValue struct {
	// Exactly one of these is set.
	Tensor *TensorValue // field 1
}

// TensorValue holds tensor data inline.
// Proto: MILSpec.TensorValue
type TensorValue struct {
	// Exactly one of these is set.
	Floats  []float32 // field 1
	Ints    []int32   // field 2
	Bools   []bool    // field 3
	Strings []string  // field 4
	Longs   []int64   // field 5
	Doubles []float64 // field 6
	Bytes   []byte    // field 7
}

// BlobFileValue references weight data in a blob file.
// Proto: MILSpec.BlobFileValue
type BlobFileValue struct {
	FileName string // field 1
	Offset   uint64 // field 2
}

// Argument is an Operation input (list of bindings).
// Proto: MILSpec.Argument
type Argument struct {
	Bindings []Binding // field 1
}

// Binding is either a name reference or an inline value.
// Proto: MILSpec.Argument.Binding
type Binding struct {
	Name  string // field 1 (oneof)
	Value *Value // field 2 (oneof)
}
