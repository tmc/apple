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
	for _, fn := range desc.Functions {
		fields = append(fields, encodeBytes(20, encodeFunctionDescription(fn)))
	}
	if desc.DefaultFunctionName != "" {
		fields = append(fields, encodeBytes(21, []byte(desc.DefaultFunctionName)))
	}
	for _, in := range desc.Inputs {
		fields = append(fields, encodeBytes(1, encodeFeatureDescription(in, false)))
	}
	for _, out := range desc.Outputs {
		fields = append(fields, encodeBytes(10, encodeFeatureDescription(out, false)))
	}
	for _, state := range desc.States {
		fields = append(fields, encodeBytes(13, encodeFeatureDescription(state, true)))
	}
	if desc.Metadata != nil {
		fields = append(fields, encodeBytes(100, encodeModelMetadata(desc.Metadata)))
	}
	fields = append(fields, desc.unknown)
	return concatBytes(fields...)
}

func encodeModelMetadata(md *ModelMetadata) []byte {
	var parts [][]byte
	if md.ShortDescription != "" {
		parts = append(parts, encodeBytes(1, []byte(md.ShortDescription)))
	}
	if md.VersionString != "" {
		parts = append(parts, encodeBytes(2, []byte(md.VersionString)))
	}
	if md.Author != "" {
		parts = append(parts, encodeBytes(3, []byte(md.Author)))
	}
	if md.License != "" {
		parts = append(parts, encodeBytes(4, []byte(md.License)))
	}
	// Sorted keys keep encoding deterministic, as elsewhere for proto maps.
	for _, key := range sortedKeys(md.UserDefined) {
		parts = append(parts, encodeBytes(100, encodeMapEntry(key, []byte(md.UserDefined[key]))))
	}
	return concatBytes(parts...)
}

func encodeFunctionDescription(fn FunctionDescription) []byte {
	parts := [][]byte{encodeBytes(1, []byte(fn.Name))}
	for _, in := range fn.Inputs {
		parts = append(parts, encodeBytes(2, encodeFeatureDescription(in, false)))
	}
	for _, out := range fn.Outputs {
		parts = append(parts, encodeBytes(3, encodeFeatureDescription(out, false)))
	}
	for _, state := range fn.States {
		parts = append(parts, encodeBytes(6, encodeFeatureDescription(state, true)))
	}
	parts = append(parts, fn.unknown)
	return concatBytes(parts...)
}

func encodeFeatureDescription(fd FeatureDescription, isState bool) []byte {
	parts := [][]byte{encodeBytes(1, []byte(fd.Name))}
	if fd.Type != nil {
		parts = append(parts, encodeBytes(3, encodeFeatureType(fd.Type, isState)))
	}
	parts = append(parts, fd.unknown)
	return concatBytes(parts...)
}

func encodeFeatureType(ft *FeatureType, isState bool) []byte {
	var parts [][]byte
	if ft.Int64Type {
		parts = append(parts, encodeBytes(1, nil))
	}
	if ft.DoubleType {
		parts = append(parts, encodeBytes(2, nil))
	}
	if ft.StringType {
		parts = append(parts, encodeBytes(3, nil))
	}
	if ft.ImageType != nil {
		var imgParts [][]byte
		if ft.ImageType.Width != 0 {
			imgParts = append(imgParts, encodeVarint(uint64(1)<<3|wireVarint, encodeVarintVal(uint64(ft.ImageType.Width))))
		}
		if ft.ImageType.Height != 0 {
			imgParts = append(imgParts, encodeVarint(uint64(2)<<3|wireVarint, encodeVarintVal(uint64(ft.ImageType.Height))))
		}
		// colorSpace is field 3; Core ML rejects INVALID_COLOR_SPACE (0).
		if ft.ImageType.ColorSpace != 0 {
			imgParts = append(imgParts, encodeVarint(uint64(3)<<3|wireVarint, encodeVarintVal(uint64(ft.ImageType.ColorSpace))))
		}
		imgParts = append(imgParts, ft.ImageType.unknown)
		parts = append(parts, encodeBytes(4, concatBytes(imgParts...)))
	}
	if ft.MultiArrayType != nil {
		if isState {
			parts = append(parts, encodeBytes(8, encodeStateFeatureType(ft.MultiArrayType)))
		} else {
			parts = append(parts, encodeBytes(5, encodeArrayFeatureType(ft.MultiArrayType)))
		}
	}
	if ft.DictionaryType != nil {
		// DictionaryFeatureType.KeyType: int64KeyType = 1, stringKeyType = 2.
		var dictParts [][]byte
		if ft.DictionaryType.KeyType == "int64" {
			dictParts = append(dictParts, encodeBytes(1, nil))
		} else if ft.DictionaryType.KeyType == "string" {
			dictParts = append(dictParts, encodeBytes(2, nil))
		}
		parts = append(parts, encodeBytes(6, concatBytes(dictParts...)))
	}
	if ft.SequenceType != nil {
		// SequenceFeatureType.Type selects an empty message by field number:
		// int64Type = 1, stringType = 3. The element type is not nested.
		if et := ft.SequenceType.ElementType; et != nil {
			switch {
			case et.Int64Type:
				parts = append(parts, encodeBytes(7, encodeBytes(1, nil)))
			case et.StringType:
				parts = append(parts, encodeBytes(7, encodeBytes(3, nil)))
			}
		}
	}
	if ft.StateArrayType != nil {
		parts = append(parts, encodeBytes(8, encodeStateFeatureType(ft.StateArrayType)))
	}
	if ft.IsOptional {
		parts = append(parts, encodeVarint(uint64(1000)<<3|wireVarint, encodeVarintVal(1)))
	}
	parts = append(parts, ft.unknown)
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
	parts = append(parts, arr.unknown)
	return concatBytes(parts...)
}

func encodeStateFeatureType(arr *ArrayFeatureType) []byte {
	return encodeBytes(1, encodeArrayFeatureType(arr))
}

// EncodeModel encodes a Model to protobuf wire format.
//
// The Model type models the mlprogram subset of the CoreML spec. Fields
// outside that subset cannot be built from Go, but are preserved verbatim on a
// model that was decoded from wire bytes and are re-emitted after the modeled
// fields of the message they belong to.
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
	parts = append(parts, m.unknown)
	return concatBytes(parts...)
}

func encodeProgram(p *Program) []byte {
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

func encodeFunction(f *Function) []byte {
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

func encodeBlock(b *Block) []byte {
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

func encodeOperation(op *Operation) []byte {
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

func encodeNamedValueType(nvt NamedValueType) []byte {
	var parts [][]byte
	if nvt.Name != "" {
		parts = append(parts, encodeBytes(1, []byte(nvt.Name)))
	}
	if nvt.Type != nil {
		parts = append(parts, encodeBytes(2, encodeValueType(nvt.Type)))
	}
	return concatBytes(parts...)
}

func encodeValueType(vt *ValueType) []byte {
	// ValueType.type is a oneof: emit at most one member, lowest proto
	// field number first.
	switch {
	case vt.TensorType != nil:
		return encodeBytes(1, encodeTensorType(vt.TensorType))
	case vt.ListType != nil:
		return encodeBytes(2, encodeListType(vt.ListType))
	case vt.TupleType != nil:
		return encodeBytes(3, encodeTupleType(vt.TupleType))
	case vt.DictionaryType != nil:
		return encodeBytes(4, encodeDictionaryType(vt.DictionaryType))
	case vt.StateType != nil:
		return encodeBytes(5, encodeStateType(vt.StateType))
	}
	return nil
}

func encodeListType(lt *ListType) []byte {
	var parts [][]byte
	if lt.ElementType != nil {
		parts = append(parts, encodeBytes(1, encodeValueType(lt.ElementType)))
	}
	// ListType.length is a Dimension message, not an integer.
	if lt.Length != 0 {
		parts = append(parts, encodeBytes(2, encodeDimension(Dimension{Constant: uint64(lt.Length)})))
	}
	return concatBytes(parts...)
}

func encodeTupleType(tt *TupleType) []byte {
	var parts [][]byte
	for _, vt := range tt.Types {
		parts = append(parts, encodeBytes(1, encodeValueType(vt)))
	}
	return concatBytes(parts...)
}

func encodeDictionaryType(dt *DictionaryType) []byte {
	var parts [][]byte
	if dt.KeyType != nil {
		parts = append(parts, encodeBytes(1, encodeValueType(dt.KeyType)))
	}
	if dt.ValueType != nil {
		parts = append(parts, encodeBytes(2, encodeValueType(dt.ValueType)))
	}
	return concatBytes(parts...)
}

func encodeTensorType(tt *TensorType) []byte {
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

func encodeStateType(st *StateType) []byte {
	var parts [][]byte
	if st.WrappedType != nil {
		parts = append(parts, encodeBytes(1, encodeValueType(st.WrappedType)))
	}
	return concatBytes(parts...)
}

func encodeDimension(d Dimension) []byte {
	if d.Unknown {
		// Dimension { field 2 = UnknownDimension { field 1 = variadic } }
		if d.Variadic {
			inner := encodeVarint(uint64(1)<<3|wireVarint, encodeVarintVal(1))
			return encodeBytes(2, inner)
		}
		return encodeBytes(2, nil)
	}
	// Dimension { field 1 = ConstantDimension { field 1 = size } }
	inner := encodeVarint(uint64(1)<<3|wireVarint, encodeVarintVal(d.Constant))
	return encodeBytes(1, inner)
}

func encodeValue(v *Value) []byte {
	var parts [][]byte
	if v.Type != nil {
		parts = append(parts, encodeBytes(2, encodeValueType(v.Type)))
	}
	// Value.value is a oneof over immediateValue and blobFileValue; type is
	// outside it.
	switch {
	case v.Immediate != nil:
		parts = append(parts, encodeBytes(3, encodeImmediateValue(v.Immediate)))
	case v.BlobFile != nil:
		parts = append(parts, encodeBytes(5, encodeBlobFileValue(v.BlobFile)))
	}
	return concatBytes(parts...)
}

func encodeImmediateValue(iv *ImmediateValue) []byte {
	var parts [][]byte
	// ImmediateValue.value is a oneof: emit at most one member.
	switch {
	case iv.Tensor != nil:
		parts = append(parts, encodeBytes(1, encodeTensorValue(iv.Tensor)))
	case iv.Tuple != nil:
		var tupleParts [][]byte
		for _, v := range iv.Tuple.Values {
			tupleParts = append(tupleParts, encodeBytes(1, encodeValue(v)))
		}
		parts = append(parts, encodeBytes(2, concatBytes(tupleParts...)))
	case iv.List != nil:
		var listParts [][]byte
		for _, v := range iv.List.Values {
			listParts = append(listParts, encodeBytes(1, encodeValue(v)))
		}
		parts = append(parts, encodeBytes(3, concatBytes(listParts...)))
	case iv.Dictionary != nil:
		var dictParts [][]byte
		for _, entry := range iv.Dictionary.Entries {
			var entryParts [][]byte
			if entry.Key != nil {
				entryParts = append(entryParts, encodeBytes(1, encodeValue(entry.Key)))
			}
			if entry.Value != nil {
				entryParts = append(entryParts, encodeBytes(2, encodeValue(entry.Value)))
			}
			dictParts = append(dictParts, encodeBytes(1, concatBytes(entryParts...)))
		}
		parts = append(parts, encodeBytes(4, concatBytes(dictParts...)))
	}
	return concatBytes(parts...)
}

func encodeTensorValue(tv *TensorValue) []byte {
	var parts [][]byte
	// Presence, not length, selects the field: an empty tensor const (a
	// dimension of size 0) must still set its oneof field so the reader can
	// tell which field was intended. coremltools does this via SetInParent.
	emit := func(field int, inner []byte) {
		if len(inner) == 0 {
			// Empty submessage: field present, no values.
			parts = append(parts, encodeBytes(field, nil))
			return
		}
		parts = append(parts, encodeBytes(field, encodeBytes(1, inner)))
	}
	switch {
	case tv.Floats != nil:
		// RepeatedFloats { field 1 = packed float32 }
		emit(1, encodePackedFloat32(tv.Floats))
	case tv.Ints != nil:
		// RepeatedInts { field 1 = packed varint int32 }
		emit(2, encodePackedVarintInt32(tv.Ints))
	case tv.Bools != nil:
		// RepeatedBools { field 1 = packed varint bool }
		emit(3, encodePackedVarintBool(tv.Bools))
	case tv.Strings != nil:
		// RepeatedStrings { field 1 = repeated string (NOT packed) }
		var innerParts [][]byte
		for _, s := range tv.Strings {
			innerParts = append(innerParts, encodeBytes(1, []byte(s)))
		}
		parts = append(parts, encodeBytes(4, concatBytes(innerParts...)))
	case tv.Longs != nil:
		// RepeatedLongInts { field 1 = packed varint int64 }
		emit(5, encodePackedVarintInt64(tv.Longs))
	case tv.Doubles != nil:
		// RepeatedDoubles { field 1 = packed float64 }
		emit(6, encodePackedFloat64(tv.Doubles))
	case tv.Bytes != nil:
		// RepeatedBytes { field 1 = bytes }
		emit(7, tv.Bytes)
	}
	return concatBytes(parts...)
}

func encodeBlobFileValue(bf *BlobFileValue) []byte {
	var parts [][]byte
	if bf.FileName != "" {
		parts = append(parts, encodeBytes(1, []byte(bf.FileName)))
	}
	if bf.Offset != 0 {
		parts = append(parts, encodeVarint(uint64(2)<<3|wireVarint, encodeVarintVal(bf.Offset)))
	}
	return concatBytes(parts...)
}

func encodeArgument(a *Argument) []byte {
	var parts [][]byte
	for _, b := range a.Bindings {
		parts = append(parts, encodeBytes(1, encodeBinding(b)))
	}
	return concatBytes(parts...)
}

func encodeBinding(b Binding) []byte {
	// Binding.binding is a oneof: an inline value wins over a name, since a
	// non-nil Value is unambiguous intent while Name has no presence bit.
	switch {
	case b.Value != nil:
		return encodeBytes(2, encodeValue(b.Value))
	case b.Name != "":
		return encodeBytes(1, []byte(b.Name))
	}
	return nil
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
