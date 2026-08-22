package coremlcompiler

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

// TestFeatureTypeOneofFieldNumbers pins the oneof member numbering from
// FeatureTypes.proto: DictionaryFeatureType {int64KeyType=1, stringKeyType=2}
// and SequenceFeatureType {int64Type=1, stringType=3}, each an empty message.
func TestFeatureTypeOneofFieldNumbers(t *testing.T) {
	tests := []struct {
		name string
		ft   *FeatureType
		want []byte
	}{
		{
			name: "dict int64 key",
			ft:   &FeatureType{DictionaryType: &DictionaryFeatureType{KeyType: "int64"}},
			want: encodeBytes(6, encodeBytes(1, nil)),
		},
		{
			name: "dict string key",
			ft:   &FeatureType{DictionaryType: &DictionaryFeatureType{KeyType: "string"}},
			want: encodeBytes(6, encodeBytes(2, nil)),
		},
		{
			name: "sequence of int64",
			ft:   &FeatureType{SequenceType: &SequenceFeatureType{ElementType: &FeatureType{Int64Type: true}}},
			want: encodeBytes(7, encodeBytes(1, nil)),
		},
		{
			name: "sequence of string",
			ft:   &FeatureType{SequenceType: &SequenceFeatureType{ElementType: &FeatureType{StringType: true}}},
			want: encodeBytes(7, encodeBytes(3, nil)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeFeatureType(tt.ft, false)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeFeatureType = % x, want % x", got, tt.want)
			}
			back, err := decodeFeatureType(got)
			if err != nil {
				t.Fatalf("decodeFeatureType: %v", err)
			}
			if !reflect.DeepEqual(back, tt.ft) {
				t.Errorf("round trip = %+v, want %+v", back, tt.ft)
			}
		})
	}
}

// TestImageFeatureTypeColorSpace checks that colorSpace (field 3) is emitted;
// Core ML rejects a model whose image feature has INVALID_COLOR_SPACE.
func TestImageFeatureTypeColorSpace(t *testing.T) {
	ft := &FeatureType{ImageType: &ImageFeatureType{Width: 8, Height: 4, ColorSpace: ColorSpaceBGR}}
	enc := encodeFeatureType(ft, false)
	want := encodeBytes(4, concatBytes(
		encodeVarint(1<<3|wireVarint, encodeVarintVal(8)),
		encodeVarint(2<<3|wireVarint, encodeVarintVal(4)),
		encodeVarint(3<<3|wireVarint, encodeVarintVal(30)),
	))
	if !bytes.Equal(enc, want) {
		t.Errorf("encode = % x, want % x", enc, want)
	}
	back, err := decodeFeatureType(enc)
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !reflect.DeepEqual(back, ft) {
		t.Errorf("round trip = %+v, want %+v", back.ImageType, ft.ImageType)
	}
}

// TestListTypeLengthIsDimension checks that ListType.length is a Dimension
// message (wire type 2), not a bare varint.
func TestListTypeLengthIsDimension(t *testing.T) {
	lt := &ListType{
		ElementType: &ValueType{TensorType: &TensorType{DataType: DataTypeFloat16, Rank: 1}},
		Length:      7,
	}
	enc := encodeListType(lt)
	wantLen := encodeBytes(2, encodeDimension(Dimension{Constant: 7}))
	if !bytes.Contains(enc, wantLen) {
		t.Errorf("encodeListType = % x, missing length dimension % x", enc, wantLen)
	}
	back, err := decodeListType(enc)
	if err != nil {
		t.Fatalf("decodeListType: %v", err)
	}
	if back.Length != 7 {
		t.Errorf("round trip length = %d, want 7", back.Length)
	}
}

// TestReadBytesRejectsOverlongLength checks that an untrusted length is
// compared without a signed conversion: a length >= 2^63 must be an error,
// not a slice-bounds panic.
// TestDecodeRepeatedAcceptsUnpackedAndSplit checks the three legal wire forms
// of a proto3 packed repeated scalar: one packed chunk, several packed chunks
// (which concatenate), and the unpacked form.
func TestDecodeRepeatedAcceptsUnpackedAndSplit(t *testing.T) {
	f32 := func(v float32) []byte {
		b := encodeVarintVal(uint64(1)<<3 | wireFixed32)
		var buf [4]byte
		bits := math.Float32bits(v)
		buf[0], buf[1], buf[2], buf[3] = byte(bits), byte(bits>>8), byte(bits>>16), byte(bits>>24)
		return concatBytes(b, buf[:])
	}
	tests := []struct {
		name string
		data []byte
	}{
		{"single packed chunk", encodeBytes(1, encodePackedFloat32([]float32{1, 2, 3}))},
		{"split packed chunks", concatBytes(
			encodeBytes(1, encodePackedFloat32([]float32{1, 2})),
			encodeBytes(1, encodePackedFloat32([]float32{3})),
		)},
		{"unpacked", concatBytes(f32(1), f32(2), f32(3))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeRepeatedFloats(tt.data)
			if err != nil {
				t.Fatalf("decodeRepeatedFloats: %v", err)
			}
			if want := []float32{1, 2, 3}; !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}

	longs := concatBytes(
		encodeBytes(1, encodePackedVarintInt64([]int64{5})),
		encodeVarint(1<<3|wireVarint, encodeVarintVal(6)),
	)
	gotLongs, err := decodeRepeatedLongs(longs)
	if err != nil {
		t.Fatalf("decodeRepeatedLongs: %v", err)
	}
	if want := []int64{5, 6}; !reflect.DeepEqual(gotLongs, want) {
		t.Errorf("longs = %v, want %v", gotLongs, want)
	}
}

func TestEncodeEmptyTensorValueSetsField(t *testing.T) {
	tests := []struct {
		name string
		tv   TensorValue
		want []byte // field tag + zero length
	}{
		{"floats", TensorValue{Floats: []float32{}}, []byte{0x0A, 0x00}},
		{"ints", TensorValue{Ints: []int32{}}, []byte{0x12, 0x00}},
		{"bools", TensorValue{Bools: []bool{}}, []byte{0x1A, 0x00}},
		{"strings", TensorValue{Strings: []string{}}, []byte{0x22, 0x00}},
		{"longs", TensorValue{Longs: []int64{}}, []byte{0x2A, 0x00}},
		{"doubles", TensorValue{Doubles: []float64{}}, []byte{0x32, 0x00}},
		{"bytes", TensorValue{Bytes: []byte{}}, []byte{0x3A, 0x00}},
		{"unset stays unset", TensorValue{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeTensorValue(&tt.tv)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encodeTensorValue = % x, want % x", got, tt.want)
			}
		})
	}
}

// topLevelFields returns the field numbers present at the top level of a
// serialized message, in wire order.
func topLevelFields(t *testing.T, data []byte) []int {
	t.Helper()
	var fields []int
	r := newProtoReader(data)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil {
			t.Fatalf("readTag: %v", err)
		}
		if err := r.skip(wire); err != nil {
			t.Fatalf("skip: %v", err)
		}
		fields = append(fields, field)
	}
	return fields
}

// TestEncodeOneofEmitsOneMember checks that the MIL oneofs emit at most one
// member even when the Go struct has several set: MIL.proto ValueType.type,
// Value.value, ImmediateValue.value and Argument.Binding.binding are oneofs,
// but each member is an independent Go field.
func TestEncodeOneofEmitsOneMember(t *testing.T) {
	tensor := &ValueType{TensorType: &TensorType{DataType: DataTypeFloat32, Rank: 1}}
	tests := []struct {
		name string
		enc  []byte
		want []int
	}{
		{
			name: "ValueType tensor and list",
			enc: encodeValueType(&ValueType{
				TensorType: &TensorType{DataType: DataTypeFloat32, Rank: 1},
				ListType:   &ListType{ElementType: tensor, Length: 2},
			}),
			want: []int{1},
		},
		{
			name: "Value immediate and blob file",
			enc: encodeValue(&Value{
				Type:      tensor,
				Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: []float32{1}}},
				BlobFile:  &BlobFileValue{FileName: "weights.bin", Offset: 64},
			}),
			want: []int{2, 3},
		},
		{
			name: "ImmediateValue tensor and tuple",
			enc: encodeImmediateValue(&ImmediateValue{
				Tensor: &TensorValue{Floats: []float32{1}},
				Tuple:  &TupleValue{Values: []*Value{{Type: tensor}}},
			}),
			want: []int{1},
		},
		{
			name: "Binding name and value",
			enc: encodeBinding(Binding{
				Name:  "x",
				Value: &Value{Type: tensor},
			}),
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topLevelFields(t, tt.enc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fields = %v, want %v", got, tt.want)
			}
		})
	}
}
