package coremlcompiler

import (
	"encoding/binary"
	"math"
)

func encodeVarintVal(v uint64) []byte {
	var buf [10]byte
	n := binary.PutUvarint(buf[:], v)
	return buf[:n]
}

func encodeVarint(tag uint64, val []byte) []byte {
	return concatBytes(encodeVarintVal(tag), val)
}

func encodeBytes(field int, data []byte) []byte {
	tag := encodeVarintVal(uint64(field)<<3 | wireBytes)
	length := encodeVarintVal(uint64(len(data)))
	return concatBytes(tag, length, data)
}

func concatBytes(parts ...[]byte) []byte {
	total := 0
	for _, p := range parts {
		total += len(p)
	}
	out := make([]byte, 0, total)
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}

func encodeModelDescription(desc ModelDescription) []byte {
	var fields [][]byte
	for _, in := range desc.Inputs {
		fields = append(fields, encodeBytes(1, encodeFeatureDescription(in, false)))
	}
	for _, out := range desc.Outputs {
		fields = append(fields, encodeBytes(10, encodeFeatureDescription(out, false)))
	}
	for _, state := range desc.States {
		fields = append(fields, encodeBytes(13, encodeFeatureDescription(state, true)))
	}
	return concatBytes(fields...)
}

func encodeFeatureDescription(fd FeatureDescription, isState bool) []byte {
	parts := [][]byte{encodeBytes(1, []byte(fd.Name))}
	if fd.Type != nil {
		parts = append(parts, encodeBytes(3, encodeFeatureType(fd.Type, isState)))
	}
	return concatBytes(parts...)
}

func encodeFeatureType(ft *FeatureType, isState bool) []byte {
	var parts [][]byte
	if ft.MultiArrayType != nil {
		if isState {
			parts = append(parts, encodeBytes(8, encodeStateFeatureType(ft.MultiArrayType)))
		} else {
			parts = append(parts, encodeBytes(5, encodeArrayFeatureType(ft.MultiArrayType)))
		}
	}
	if ft.IsOptional {
		parts = append(parts, encodeVarint(uint64(1000)<<3|wireVarint, encodeVarintVal(1)))
	}
	return concatBytes(parts...)
}

func encodeArrayFeatureType(arr *ArrayFeatureType) []byte {
	var parts [][]byte
	if len(arr.Shape) > 0 {
		shape := make([]byte, 0, len(arr.Shape)*2)
		for _, dim := range arr.Shape {
			shape = append(shape, encodeVarintVal(uint64(dim))...)
		}
		parts = append(parts, encodeBytes(1, shape))
	}
	if arr.DataType != 0 {
		parts = append(parts, encodeVarint(uint64(2)<<3|wireVarint, encodeVarintVal(uint64(arr.DataType))))
	}
	return concatBytes(parts...)
}

func encodeStateFeatureType(arr *ArrayFeatureType) []byte {
	return encodeBytes(1, encodeArrayFeatureType(arr))
}

// EncodeModel encodes a Model to protobuf wire format.
func EncodeModel(m *Model) []byte {
	var parts [][]byte
	if m.SpecVersion != 0 {
		parts = append(parts, encodeVarint(uint64(1)<<3|wireVarint, encodeVarintVal(uint64(m.SpecVersion))))
	}
	desc := encodeModelDescription(m.Description)
	if len(desc) > 0 {
		parts = append(parts, encodeBytes(2, desc))
	}
	if m.MLProgram != nil {
		parts = append(parts, encodeBytes(502, encodeProgram(m.MLProgram)))
	}
	return concatBytes(parts...)
}

func encodeProgram(p *program) []byte {
	var parts [][]byte
	if p.Version != 0 {
		parts = append(parts, encodeVarint(uint64(1)<<3|wireVarint, encodeVarintVal(uint64(p.Version))))
	}
	for _, key := range sortedKeys(p.Functions) {
		entry := encodeMapEntry(key, encodeFunction(p.Functions[key]))
		parts = append(parts, encodeBytes(2, entry))
	}
	for _, key := range sortedKeys(p.Attributes) {
		entry := encodeMapEntry(key, encodeValue(p.Attributes[key]))
		parts = append(parts, encodeBytes(4, entry))
	}
	return concatBytes(parts...)
}

func encodeFunction(f *function) []byte {
	var parts [][]byte
	for _, nvt := range f.Inputs {
		parts = append(parts, encodeBytes(1, encodeNamedValueType(nvt)))
	}
	if f.OpSet != "" {
		parts = append(parts, encodeBytes(2, []byte(f.OpSet)))
	}
	for _, key := range sortedKeys(f.BlockSpecializations) {
		entry := encodeMapEntry(key, encodeBlock(f.BlockSpecializations[key]))
		parts = append(parts, encodeBytes(3, entry))
	}
	for _, key := range sortedKeys(f.Attributes) {
		entry := encodeMapEntry(key, encodeValue(f.Attributes[key]))
		parts = append(parts, encodeBytes(4, entry))
	}
	return concatBytes(parts...)
}

func encodeBlock(b *block) []byte {
	var parts [][]byte
	for _, nvt := range b.Inputs {
		parts = append(parts, encodeBytes(1, encodeNamedValueType(nvt)))
	}
	for _, out := range b.Outputs {
		parts = append(parts, encodeBytes(2, []byte(out)))
	}
	for _, op := range b.Operations {
		parts = append(parts, encodeBytes(3, encodeOperation(op)))
	}
	return concatBytes(parts...)
}

func encodeOperation(op *operation) []byte {
	var parts [][]byte
	if op.Type != "" {
		parts = append(parts, encodeBytes(1, []byte(op.Type)))
	}
	for _, key := range sortedKeys(op.Inputs) {
		entry := encodeMapEntry(key, encodeArgument(op.Inputs[key]))
		parts = append(parts, encodeBytes(2, entry))
	}
	for _, nvt := range op.Outputs {
		parts = append(parts, encodeBytes(3, encodeNamedValueType(nvt)))
	}
	for _, blk := range op.Blocks {
		parts = append(parts, encodeBytes(4, encodeBlock(blk)))
	}
	for _, key := range sortedKeys(op.Attributes) {
		entry := encodeMapEntry(key, encodeValue(op.Attributes[key]))
		parts = append(parts, encodeBytes(5, entry))
	}
	return concatBytes(parts...)
}

func encodeNamedValueType(nvt namedValueType) []byte {
	var parts [][]byte
	if nvt.Name != "" {
		parts = append(parts, encodeBytes(1, []byte(nvt.Name)))
	}
	if nvt.Type != nil {
		parts = append(parts, encodeBytes(2, encodeValueType(nvt.Type)))
	}
	return concatBytes(parts...)
}

func encodeValueType(vt *valueType) []byte {
	var parts [][]byte
	if vt.TensorType != nil {
		parts = append(parts, encodeBytes(1, encodeTensorType(vt.TensorType)))
	}
	if vt.StateType != nil {
		parts = append(parts, encodeBytes(5, encodeStateType(vt.StateType)))
	}
	return concatBytes(parts...)
}

func encodeTensorType(tt *tensorType) []byte {
	var parts [][]byte
	if tt.DataType != 0 {
		parts = append(parts, encodeVarint(uint64(1)<<3|wireVarint, encodeVarintVal(uint64(tt.DataType))))
	}
	if tt.Rank != 0 {
		parts = append(parts, encodeVarint(uint64(2)<<3|wireVarint, encodeVarintVal(uint64(tt.Rank))))
	}
	for _, dim := range tt.Dimensions {
		parts = append(parts, encodeBytes(3, encodeDimension(dim)))
	}
	return concatBytes(parts...)
}

func encodeStateType(st *stateType) []byte {
	var parts [][]byte
	if st.WrappedType != nil {
		parts = append(parts, encodeBytes(1, encodeValueType(st.WrappedType)))
	}
	return concatBytes(parts...)
}

func encodeDimension(d dimension) []byte {
	if d.Unknown {
		// Dimension { field 2 = UnknownDimension {} }
		return encodeBytes(2, nil)
	}
	// Dimension { field 1 = ConstantDimension { field 1 = size } }
	inner := encodeVarint(uint64(1)<<3|wireVarint, encodeVarintVal(d.Constant))
	return encodeBytes(1, inner)
}

func encodeValue(v *value) []byte {
	var parts [][]byte
	if v.Type != nil {
		parts = append(parts, encodeBytes(2, encodeValueType(v.Type)))
	}
	if v.Immediate != nil {
		parts = append(parts, encodeBytes(3, encodeImmediateValue(v.Immediate)))
	}
	if v.BlobFile != nil {
		parts = append(parts, encodeBytes(5, encodeBlobFileValue(v.BlobFile)))
	}
	return concatBytes(parts...)
}

func encodeImmediateValue(iv *immediateValue) []byte {
	var parts [][]byte
	if iv.Tensor != nil {
		parts = append(parts, encodeBytes(1, encodeTensorValue(iv.Tensor)))
	}
	return concatBytes(parts...)
}

func encodeTensorValue(tv *tensorValue) []byte {
	var parts [][]byte
	if len(tv.Floats) > 0 {
		// RepeatedFloats { field 1 = packed float32 }
		inner := encodeBytes(1, encodePackedFloat32(tv.Floats))
		parts = append(parts, encodeBytes(1, inner))
	}
	if len(tv.Ints) > 0 {
		// RepeatedInts { field 1 = packed varint int32 }
		inner := encodeBytes(1, encodePackedVarintInt32(tv.Ints))
		parts = append(parts, encodeBytes(2, inner))
	}
	if len(tv.Bools) > 0 {
		// RepeatedBools { field 1 = packed varint bool }
		inner := encodeBytes(1, encodePackedVarintBool(tv.Bools))
		parts = append(parts, encodeBytes(3, inner))
	}
	if len(tv.Strings) > 0 {
		// RepeatedStrings { field 1 = repeated string (NOT packed) }
		var innerParts [][]byte
		for _, s := range tv.Strings {
			innerParts = append(innerParts, encodeBytes(1, []byte(s)))
		}
		parts = append(parts, encodeBytes(4, concatBytes(innerParts...)))
	}
	if len(tv.Longs) > 0 {
		// RepeatedLongInts { field 1 = packed varint int64 }
		inner := encodeBytes(1, encodePackedVarintInt64(tv.Longs))
		parts = append(parts, encodeBytes(5, inner))
	}
	if len(tv.Doubles) > 0 {
		// RepeatedDoubles { field 1 = packed float64 }
		inner := encodeBytes(1, encodePackedFloat64(tv.Doubles))
		parts = append(parts, encodeBytes(6, inner))
	}
	if len(tv.Bytes) > 0 {
		// RepeatedBytes { field 1 = bytes }
		inner := encodeBytes(1, tv.Bytes)
		parts = append(parts, encodeBytes(7, inner))
	}
	return concatBytes(parts...)
}

func encodeBlobFileValue(bf *blobFileValue) []byte {
	var parts [][]byte
	if bf.FileName != "" {
		parts = append(parts, encodeBytes(1, []byte(bf.FileName)))
	}
	if bf.Offset != 0 {
		parts = append(parts, encodeVarint(uint64(2)<<3|wireVarint, encodeVarintVal(bf.Offset)))
	}
	return concatBytes(parts...)
}

func encodeArgument(a *argument) []byte {
	var parts [][]byte
	for _, b := range a.Bindings {
		parts = append(parts, encodeBytes(1, encodeBinding(b)))
	}
	return concatBytes(parts...)
}

func encodeBinding(b binding) []byte {
	var parts [][]byte
	if b.Name != "" {
		parts = append(parts, encodeBytes(1, []byte(b.Name)))
	}
	if b.Value != nil {
		parts = append(parts, encodeBytes(2, encodeValue(b.Value)))
	}
	return concatBytes(parts...)
}

func encodeMapEntry(key string, val []byte) []byte {
	return concatBytes(
		encodeBytes(1, []byte(key)),
		encodeBytes(2, val),
	)
}

// Packed array helpers.

func encodePackedFloat32(vals []float32) []byte {
	out := make([]byte, len(vals)*4)
	for i, v := range vals {
		binary.LittleEndian.PutUint32(out[i*4:], math.Float32bits(v))
	}
	return out
}

func encodePackedFloat64(vals []float64) []byte {
	out := make([]byte, len(vals)*8)
	for i, v := range vals {
		binary.LittleEndian.PutUint64(out[i*8:], math.Float64bits(v))
	}
	return out
}

func encodePackedVarintInt32(vals []int32) []byte {
	out := make([]byte, 0, len(vals)*2)
	for _, v := range vals {
		out = append(out, encodeVarintVal(uint64(v))...)
	}
	return out
}

func encodePackedVarintInt64(vals []int64) []byte {
	out := make([]byte, 0, len(vals)*2)
	for _, v := range vals {
		out = append(out, encodeVarintVal(uint64(v))...)
	}
	return out
}

func encodePackedVarintBool(vals []bool) []byte {
	out := make([]byte, 0, len(vals))
	for _, v := range vals {
		if v {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	return out
}

