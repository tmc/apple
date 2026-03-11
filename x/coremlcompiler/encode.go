package coremlcompiler

import "encoding/binary"

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
