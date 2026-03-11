package coremlcompiler

// Minimal protobuf wire format reader for CoreML model files.
// Reads only the fields needed for compilation, avoiding any
// external protobuf dependency.

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Wire types.
const (
	wireVarint  = 0
	wireFixed64 = 1
	wireBytes   = 2
	wireFixed32 = 5
)

// protoReader reads protobuf wire format.
type protoReader struct {
	data []byte
	pos  int
}

func newProtoReader(data []byte) *protoReader {
	return &protoReader{data: data}
}

func (r *protoReader) remaining() int { return len(r.data) - r.pos }
func (r *protoReader) done() bool     { return r.pos >= len(r.data) }

func (r *protoReader) readVarint() (uint64, error) {
	var x uint64
	var shift uint
	for r.pos < len(r.data) {
		b := r.data[r.pos]
		r.pos++
		x |= uint64(b&0x7F) << shift
		if b < 0x80 {
			return x, nil
		}
		shift += 7
		if shift >= 64 {
			return 0, fmt.Errorf("proto: varint overflow")
		}
	}
	return 0, fmt.Errorf("proto: truncated varint")
}

func (r *protoReader) readTag() (field int, wire int, err error) {
	v, err := r.readVarint()
	if err != nil {
		return 0, 0, err
	}
	return int(v >> 3), int(v & 7), nil
}

func (r *protoReader) readBytes() ([]byte, error) {
	n, err := r.readVarint()
	if err != nil {
		return nil, err
	}
	if int(n) > r.remaining() {
		return nil, fmt.Errorf("proto: bytes length %d exceeds remaining %d", n, r.remaining())
	}
	data := r.data[r.pos : r.pos+int(n)]
	r.pos += int(n)
	return data, nil
}

func (r *protoReader) readString() (string, error) {
	b, err := r.readBytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (r *protoReader) readFixed32() (uint32, error) {
	if r.remaining() < 4 {
		return 0, fmt.Errorf("proto: truncated fixed32")
	}
	v := binary.LittleEndian.Uint32(r.data[r.pos:])
	r.pos += 4
	return v, nil
}

func (r *protoReader) readFixed64() (uint64, error) {
	if r.remaining() < 8 {
		return 0, fmt.Errorf("proto: truncated fixed64")
	}
	v := binary.LittleEndian.Uint64(r.data[r.pos:])
	r.pos += 8
	return v, nil
}

func (r *protoReader) skip(wire int) error {
	switch wire {
	case wireVarint:
		_, err := r.readVarint()
		return err
	case wireFixed64:
		if r.remaining() < 8 {
			return fmt.Errorf("proto: truncated fixed64")
		}
		r.pos += 8
		return nil
	case wireBytes:
		_, err := r.readBytes()
		return err
	case wireFixed32:
		if r.remaining() < 4 {
			return fmt.Errorf("proto: truncated fixed32")
		}
		r.pos += 4
		return nil
	default:
		return fmt.Errorf("proto: unknown wire type %d", wire)
	}
}

// readPackedFloat32 reads a packed repeated float field.
func (r *protoReader) readPackedFloat32() ([]float32, error) {
	data, err := r.readBytes()
	if err != nil {
		return nil, err
	}
	if len(data)%4 != 0 {
		return nil, fmt.Errorf("proto: packed float32 length %d not multiple of 4", len(data))
	}
	out := make([]float32, len(data)/4)
	for i := range out {
		out[i] = math.Float32frombits(binary.LittleEndian.Uint32(data[i*4:]))
	}
	return out, nil
}

// readPackedFloat64 reads a packed repeated double field.
func (r *protoReader) readPackedFloat64() ([]float64, error) {
	data, err := r.readBytes()
	if err != nil {
		return nil, err
	}
	if len(data)%8 != 0 {
		return nil, fmt.Errorf("proto: packed float64 length %d not multiple of 8", len(data))
	}
	out := make([]float64, len(data)/8)
	for i := range out {
		out[i] = math.Float64frombits(binary.LittleEndian.Uint64(data[i*8:]))
	}
	return out, nil
}

// readPackedInt32 reads a packed repeated int32 field (varint encoding).
func readPackedInt32(data []byte) ([]int32, error) {
	r := newProtoReader(data)
	var out []int32
	for !r.done() {
		v, err := r.readVarint()
		if err != nil {
			return nil, err
		}
		out = append(out, int32(v))
	}
	return out, nil
}

// readPackedInt64 reads a packed repeated int64 field (varint encoding).
func readPackedInt64(data []byte) ([]int64, error) {
	r := newProtoReader(data)
	var out []int64
	for !r.done() {
		v, err := r.readVarint()
		if err != nil {
			return nil, err
		}
		out = append(out, int64(v))
	}
	return out, nil
}
