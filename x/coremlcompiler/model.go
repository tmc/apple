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

	// Wire bytes of fields this package does not model, kept so a decoded
	// model re-encodes without losing them. See EncodeModel.
	unknown []byte
}

// ModelDescription describes model inputs, outputs, and metadata.
// Proto: CoreML.Specification.ModelDescription
type ModelDescription struct {
	// Functions and DefaultFunctionName describe multi-function models.
	// When Functions is non-empty the model-level Inputs/Outputs/States
	// must be empty: the proto reserves them for the single-function case.
	Functions           []FunctionDescription // field 20
	DefaultFunctionName string                // field 21

	Inputs  []FeatureDescription // field 1
	Outputs []FeatureDescription // field 10
	States  []FeatureDescription // field 13

	Metadata *ModelMetadata // field 100

	unknown []byte // unmodeled fields, preserved verbatim
}

// ModelMetadata carries informational model metadata.
// Proto: CoreML.Specification.Metadata
type ModelMetadata struct {
	ShortDescription string            // field 1
	VersionString    string            // field 2
	Author           string            // field 3
	License          string            // field 4
	UserDefined      map[string]string // field 100
}

// FunctionDescription describes one entry point of a multi-function model.
// Proto: CoreML.Specification.FunctionDescription
type FunctionDescription struct {
	Name    string               // field 1
	Inputs  []FeatureDescription // field 2
	Outputs []FeatureDescription // field 3
	States  []FeatureDescription // field 6

	unknown []byte // unmodeled fields, preserved verbatim
}

// FeatureDescription is a named, typed model feature.
// Proto: CoreML.Specification.FeatureDescription
type FeatureDescription struct {
	Name string       // field 1
	Type *FeatureType // field 3

	unknown []byte // unmodeled fields, preserved verbatim
}

// FeatureType describes a Core ML model feature type.
type FeatureType struct {
	MultiArrayType *ArrayFeatureType
	ImageType      *ImageFeatureType
	StringType     bool
	Int64Type      bool
	DoubleType     bool
	DictionaryType *DictionaryFeatureType
	SequenceType   *SequenceFeatureType
	StateArrayType *ArrayFeatureType
	IsOptional     bool

	unknown []byte // unmodeled fields, preserved verbatim
}

// ImageFeatureType describes an image boundary feature.
type ImageFeatureType struct {
	Width      int64      // field 1
	Height     int64      // field 2
	ColorSpace ColorSpace // field 3

	unknown []byte // unmodeled fields, preserved verbatim
}

// ColorSpace identifies an image feature's pixel format.
// Proto: CoreML.Specification.ImageFeatureType.ColorSpace
type ColorSpace int32

const (
	ColorSpaceInvalid          ColorSpace = 0
	ColorSpaceGrayscale        ColorSpace = 10
	ColorSpaceRGB              ColorSpace = 20
	ColorSpaceBGR              ColorSpace = 30
	ColorSpaceGrayscaleFloat16 ColorSpace = 40
)

// DictionaryFeatureType describes a dictionary boundary feature.
type DictionaryFeatureType struct {
	KeyType string
}

// SequenceFeatureType describes a sequence boundary feature.
type SequenceFeatureType struct {
	ElementType *FeatureType
}

// ArrayFeatureType describes a Core ML multi-array feature.
type ArrayFeatureType struct {
	Shape    []int64
	DataType ArrayDataType

	unknown []byte // unmodeled fields, preserved verbatim
}

// ArrayDataType identifies Core ML multi-array element types.
type ArrayDataType int32

const (
	ArrayDataTypeInvalid ArrayDataType = 0
	ArrayDataTypeFloat16 ArrayDataType = 65552
	ArrayDataTypeFloat32 ArrayDataType = 65568
	ArrayDataTypeDouble  ArrayDataType = 65600
	ArrayDataTypeInt32   ArrayDataType = 131104
	ArrayDataTypeInt8    ArrayDataType = 131080
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
	TensorType     *TensorType     // field 1
	ListType       *ListType       // field 2
	TupleType      *TupleType      // field 3
	DictionaryType *DictionaryType // field 4
	StateType      *StateType      // field 5
}

// ListType describes a MIL list type.
type ListType struct {
	ElementType *ValueType // field 1
	Length      int64      // field 2
}

// TupleType describes a MIL tuple type.
type TupleType struct {
	Types []*ValueType // field 1
}

// DictionaryType describes a MIL dictionary type.
type DictionaryType struct {
	KeyType   *ValueType // field 1
	ValueType *ValueType // field 2
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

// Dimension is a tensor Dimension (constant, unknown, or variadic).
// Proto: MILSpec.Dimension
type Dimension struct {
	Constant uint64 // from ConstantDimension.size (field 1.1)
	Unknown  bool   // true if UnknownDimension
	Variadic bool   // true if UnknownDimension.variadic is true
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
	// The sub-byte block is deliberately non-monotonic in MIL.proto; these
	// numbers are the wire contract, not a sorted sequence.
	DataTypeUInt4        DataType = 35
	DataTypeUInt2        DataType = 36
	DataTypeUInt1        DataType = 37
	DataTypeUInt6        DataType = 38
	DataTypeUInt3        DataType = 39
	DataTypeFloat8E4M3FN DataType = 40
	DataTypeFloat8E5M2   DataType = 41
	DataTypeBool         DataType = 1
	DataTypeString       DataType = 2
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
	case DataTypeUInt1:
		return "uint1"
	case DataTypeUInt2:
		return "uint2"
	case DataTypeUInt3:
		return "uint3"
	case DataTypeUInt4:
		return "uint4"
	case DataTypeUInt6:
		return "uint6"
	case DataTypeFloat8E4M3FN:
		return "fp8_e4m3fn"
	case DataTypeFloat8E5M2:
		return "fp8_e5m2"
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
	Tensor     *TensorValue     // field 1
	Tuple      *TupleValue      // field 2
	List       *ListValue       // field 3
	Dictionary *DictionaryValue // field 4
}

// TupleValue holds structured tuple constants.
type TupleValue struct {
	Values []*Value // field 1
}

// ListValue holds structured list constants.
type ListValue struct {
	Values []*Value // field 1
}

// DictionaryMapEntry represents a key/value pair in a dictionary constant.
type DictionaryMapEntry struct {
	Key   *Value
	Value *Value
}

// DictionaryValue holds structured dictionary constants.
type DictionaryValue struct {
	Entries []DictionaryMapEntry
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

// TensorValueField identifies the TensorValue oneof field a MIL immediate
// constant of a given element type must use. Which field carries the data is
// fixed by the element type, not free choice: the MIL reader picks the field
// from the tensor's declared dtype, so a mismatch is read as an empty value.
type TensorValueField int

const (
	TensorValueFloats TensorValueField = iota + 1
	TensorValueInts
	TensorValueBools
	TensorValueStrings
	TensorValueLongs
	TensorValueDoubles
	TensorValueBytes
)

// FieldForDataType reports which TensorValue field an immediate constant of
// this element type must use, mirroring coremltools' _tensor_field_by_type.
// Note the asymmetry: uint32 goes to Bytes while int32 goes to Ints, and
// int16/uint16 go to Ints widened to int32.
func (dt DataType) FieldForDataType() (TensorValueField, error) {
	switch dt {
	case DataTypeBool:
		return TensorValueBools, nil
	case DataTypeString:
		return TensorValueStrings, nil
	case DataTypeInt64, DataTypeUInt64:
		return TensorValueLongs, nil
	case DataTypeFloat64:
		return TensorValueDoubles, nil
	case DataTypeFloat32:
		return TensorValueFloats, nil
	case DataTypeFloat16, DataTypeInt4, DataTypeInt8, DataTypeUInt1, DataTypeUInt2,
		DataTypeUInt3, DataTypeUInt4, DataTypeUInt6, DataTypeUInt8, DataTypeUInt32:
		return TensorValueBytes, nil
	case DataTypeInt16, DataTypeUInt16, DataTypeInt32:
		return TensorValueInts, nil
	}
	return 0, fmt.Errorf("no MIL immediate encoding for data type %v", dt)
}

// SetField reports which field of tv holds data, or 0 if none does.
func (tv *TensorValue) SetField() TensorValueField {
	switch {
	case tv.Floats != nil:
		return TensorValueFloats
	case tv.Ints != nil:
		return TensorValueInts
	case tv.Bools != nil:
		return TensorValueBools
	case tv.Strings != nil:
		return TensorValueStrings
	case tv.Longs != nil:
		return TensorValueLongs
	case tv.Doubles != nil:
		return TensorValueDoubles
	case tv.Bytes != nil:
		return TensorValueBytes
	}
	return 0
}

// ValidateTensorValue reports whether tv stores its data in the one field the
// MIL reader consults for element type dt.
func ValidateTensorValue(dt DataType, tv *TensorValue) error {
	want, err := dt.FieldForDataType()
	if err != nil {
		return err
	}
	got := tv.SetField()
	if got == 0 {
		return fmt.Errorf("tensor value for %v sets no field", dt)
	}
	if got != want {
		return fmt.Errorf("tensor value for %v uses field %d, want field %d", dt, got, want)
	}
	return nil
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
// Exactly one of Name and Value is set; if both are, Value wins.
// Name has no presence bit, so an empty name is indistinguishable from
// an unset one and is never emitted.
// Proto: MILSpec.Argument.Binding
type Binding struct {
	Name  string // field 1 (oneof)
	Value *Value // field 2 (oneof)
}
