package coremlcompiler

import "fmt"

// This file builds the [ModelDescription] a generated MIL program needs from
// the same vocabulary the MIL text itself is written in: an element type
// spelled the way MIL spells it ("fp16", "int32") and a static shape.
//
// Generators such as github.com/tmc/modelir describe their tensors that way
// already, so they can reach [CompileMILText] without this package taking a
// dependency on any of them, and without every generator rewriting the dtype
// mapping below. Getting that mapping wrong is silent: a model built with the
// wrong element type compiles and loads, and only the numbers come out wrong.
// Every conversion here therefore reports an error rather than falling back to
// a default type.

// ParseDataType returns the MIL data type spelled name, the inverse of
// [DataType.String]. It reports an error for a name MIL does not define.
func ParseDataType(name string) (DataType, error) {
	dt, err := parseDataType(name)
	if err != nil {
		return 0, fmt.Errorf("coremlcompiler: %w", err)
	}
	return dt, nil
}

func parseDataType(name string) (DataType, error) {
	for _, dt := range []DataType{
		DataTypeFloat16, DataTypeFloat32, DataTypeFloat64, DataTypeBFloat16,
		DataTypeInt8, DataTypeInt16, DataTypeInt32, DataTypeInt64, DataTypeInt4,
		DataTypeUInt8, DataTypeUInt16, DataTypeUInt32, DataTypeUInt64,
		DataTypeUInt1, DataTypeUInt2, DataTypeUInt3, DataTypeUInt4, DataTypeUInt6,
		DataTypeFloat8E4M3FN, DataTypeFloat8E5M2,
		DataTypeBool, DataTypeString,
	} {
		if dt.String() == name {
			return dt, nil
		}
	}
	return 0, fmt.Errorf("%q is not a MIL data type", name)
}

// ArrayDataType returns the Core ML multi-array element type corresponding to
// dt.
//
// A multi-array carries far fewer element types than MIL has: MIL's integer
// widths other than int32 and int8, its sub-byte types, and its float8 types
// have no multi-array spelling, and dt is rejected rather than widened. A
// program may still use those types internally; they just cannot appear on a
// model input, output, or state.
func (dt DataType) ArrayDataType() (ArrayDataType, error) {
	adt, err := dt.arrayDataType()
	if err != nil {
		return ArrayDataTypeInvalid, fmt.Errorf("coremlcompiler: %w", err)
	}
	return adt, nil
}

func (dt DataType) arrayDataType() (ArrayDataType, error) {
	switch dt {
	case DataTypeFloat16:
		return ArrayDataTypeFloat16, nil
	case DataTypeFloat32:
		return ArrayDataTypeFloat32, nil
	case DataTypeFloat64:
		return ArrayDataTypeDouble, nil
	case DataTypeInt32:
		return ArrayDataTypeInt32, nil
	case DataTypeInt8:
		return ArrayDataTypeInt8, nil
	}
	return ArrayDataTypeInvalid, fmt.Errorf("MIL type %s has no Core ML multi-array element type", dt)
}

// TensorFeatureDescription describes a model input or output holding a tensor
// of the given static shape whose element type is named the way MIL names it.
func TensorFeatureDescription(name, dataType string, shape []int64) (FeatureDescription, error) {
	arr, err := arrayFeatureType(name, dataType, shape)
	if err != nil {
		return FeatureDescription{}, err
	}
	return FeatureDescription{Name: name, Type: &FeatureType{MultiArrayType: arr}}, nil
}

// StateFeatureDescription describes a model state, the tensor a program reads
// with read_state and writes with write_state.
//
// Core ML accepts only fp16 and int8 states; a state of any other element type
// is rejected here rather than at load time.
func StateFeatureDescription(name, dataType string, shape []int64) (FeatureDescription, error) {
	arr, err := arrayFeatureType(name, dataType, shape)
	if err != nil {
		return FeatureDescription{}, err
	}
	if arr.DataType != ArrayDataTypeFloat16 && arr.DataType != ArrayDataTypeInt8 {
		return FeatureDescription{}, fmt.Errorf("coremlcompiler: state %q has type %s, Core ML states must be fp16 or int8", name, dataType)
	}
	return FeatureDescription{Name: name, Type: &FeatureType{StateArrayType: arr}}, nil
}

func arrayFeatureType(name, dataType string, shape []int64) (*ArrayFeatureType, error) {
	if name == "" {
		return nil, fmt.Errorf("coremlcompiler: feature has no name")
	}
	dt, err := parseDataType(dataType)
	if err != nil {
		return nil, fmt.Errorf("coremlcompiler: feature %q: %w", name, err)
	}
	adt, err := dt.arrayDataType()
	if err != nil {
		return nil, fmt.Errorf("coremlcompiler: feature %q: %w", name, err)
	}
	if len(shape) == 0 {
		return nil, fmt.Errorf("coremlcompiler: feature %q has no shape", name)
	}
	for _, d := range shape {
		if d < 1 {
			return nil, fmt.Errorf("coremlcompiler: feature %q has dimension %d, shapes must be static and positive", name, d)
		}
	}
	return &ArrayFeatureType{Shape: append([]int64(nil), shape...), DataType: adt}, nil
}
