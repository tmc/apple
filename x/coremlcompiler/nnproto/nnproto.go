// Package nnproto encodes CoreML NeuralNetwork models (specification version 1)
// as .mlmodel protobuf bytes. It provides the protobuf construction required
// to build ANE-executable models without external dependencies.
package nnproto

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Layer is an encoded NeuralNetworkLayer message.
type Layer []byte

// Feature names a model input or output and specifies its shape.
type Feature struct {
	Name  string
	Shape []int64
}

// Model describes a CoreML NeuralNetwork model.
type Model struct {
	Inputs  []Feature
	Outputs []Feature
	Layers  []Layer
}

// Minimal protobuf wire-format encoders (proto2/3)

func pvarint(buf []byte, v uint64) []byte {
	for v >= 0x80 {
		buf = append(buf, byte(v)|0x80)
		v >>= 7
	}
	return append(buf, byte(v))
}

func pfield(buf []byte, field int, wire int) []byte {
	return pvarint(buf, uint64(field)<<3|uint64(wire))
}

func pmsg(buf []byte, field int, msg []byte) []byte {
	buf = pfield(buf, field, 2)
	buf = pvarint(buf, uint64(len(msg)))
	return append(buf, msg...)
}

func pstr(buf []byte, field int, s string) []byte { return pmsg(buf, field, []byte(s)) }

func pint(buf []byte, field int, v uint64) []byte {
	buf = pfield(buf, field, 0)
	return pvarint(buf, v)
}

func ppackedInt64(buf []byte, field int, vals []int64) ([]byte, error) {
	var inner []byte
	for _, v := range vals {
		if v < 0 {
			return nil, fmt.Errorf("nnproto: shape dimensions cannot be negative (got %d)", v)
		}
		inner = pvarint(inner, uint64(v))
	}
	return pmsg(buf, field, inner), nil
}

func ppackedFloat(buf []byte, field int, vals []float32) []byte {
	inner := make([]byte, 0, 4*len(vals))
	for _, v := range vals {
		var b [4]byte
		binary.LittleEndian.PutUint32(b[:], math.Float32bits(v))
		inner = append(inner, b[:]...)
	}
	return pmsg(buf, field, inner)
}

// nnMultiArrayFeature encodes a FeatureDescription message for a float32 multi-array.
func nnMultiArrayFeature(name string, shape []int64) ([]byte, error) {
	arr, err := ppackedInt64(nil, 1, shape)
	if err != nil {
		return nil, err
	}
	arr = pint(arr, 2, 65568) // ArrayDataType.FLOAT32 (65568 = 0x10020)
	var ft []byte
	ft = pmsg(ft, 5, arr) // FeatureType.multiArrayType
	var fd []byte
	fd = pstr(fd, 1, name)
	fd = pmsg(fd, 3, ft)
	return fd, nil
}

// Conv1x1 encodes a 1x1 Convolution NeuralNetworkLayer with argument validation.
func Conv1x1(name, input, output string, inC, outC int, weights []float32) (Layer, error) {
	if inC <= 0 || outC <= 0 {
		return nil, fmt.Errorf("nnproto: conv1x1 channels must be positive (inC=%d, outC=%d)", inC, outC)
	}
	if len(weights) != inC*outC {
		return nil, fmt.Errorf("nnproto: conv1x1 weights length mismatch (got %d, want %d)", len(weights), inC*outC)
	}

	var wp []byte
	wp = ppackedFloat(wp, 1, weights) // WeightParams.floatValue

	var conv []byte
	conv = pint(conv, 1, uint64(outC)) // outputChannels
	conv = pint(conv, 2, uint64(inC))  // kernelChannels
	conv = pint(conv, 10, 1)           // nGroups
	shapePacked, _ := ppackedInt64(nil, 20, []int64{1, 1})
	conv = append(conv, shapePacked...) // kernelSize
	stridePacked, _ := ppackedInt64(nil, 30, []int64{1, 1})
	conv = append(conv, stridePacked...) // stride
	dilationPacked, _ := ppackedInt64(nil, 40, []int64{1, 1})
	conv = append(conv, dilationPacked...) // dilationFactor
	conv = pmsg(conv, 50, nil)             // ValidPadding{}
	conv = pmsg(conv, 90, wp)              // weights

	var l []byte
	l = pstr(l, 1, name)
	l = pstr(l, 2, input)
	l = pstr(l, 3, output)
	l = pmsg(l, 100, conv) // NeuralNetworkLayer.convolution
	return l, nil
}

// ReLU encodes a ReLU activation NeuralNetworkLayer.
func ReLU(name, input, output string) Layer {
	var act []byte
	act = pmsg(act, 10, nil) // ActivationParams.ReLU = ActivationReLU{}
	var l []byte
	l = pstr(l, 1, name)
	l = pstr(l, 2, input)
	l = pstr(l, 3, output)
	l = pmsg(l, 130, act) // NeuralNetworkLayer.activation
	return l
}

// Multiply encodes an elementwise multiply NeuralNetworkLayer.
func Multiply(name string, inputs []string, output string) Layer {
	var l []byte
	l = pstr(l, 1, name)
	for _, in := range inputs {
		l = pstr(l, 2, in)
	}
	l = pstr(l, 3, output)
	l = pmsg(l, 231, nil) // NeuralNetworkLayer.multiply = MultiplyLayerParams{}
	return l
}

// Encode serializes the Model into .mlmodel protobuf bytes with validation.
func (m Model) Encode() ([]byte, error) {
	var desc []byte
	for _, in := range m.Inputs {
		fDesc, err := nnMultiArrayFeature(in.Name, in.Shape)
		if err != nil {
			return nil, fmt.Errorf("nnproto: input feature %q: %w", in.Name, err)
		}
		desc = pmsg(desc, 1, fDesc)
	}
	for _, out := range m.Outputs {
		fDesc, err := nnMultiArrayFeature(out.Name, out.Shape)
		if err != nil {
			return nil, fmt.Errorf("nnproto: output feature %q: %w", out.Name, err)
		}
		desc = pmsg(desc, 10, fDesc)
	}

	var nn []byte
	for _, l := range m.Layers {
		nn = pmsg(nn, 1, l)
	}

	var proto []byte
	proto = pint(proto, 1, 1) // specificationVersion = 1
	proto = pmsg(proto, 2, desc)
	proto = pmsg(proto, 500, nn) // Model.neuralNetwork
	return proto, nil
}

// FFNReLU2 builds a Model computing out = W2·relu(W1·x)² over input [d, 1, seq] with argument validation.
func FFNReLU2(d, f, seq int, w1, w2 []float32) ([]byte, error) {
	if d <= 0 || f <= 0 || seq <= 0 {
		return nil, fmt.Errorf("nnproto: dimensions must be positive (d=%d, f=%d, seq=%d)", d, f, seq)
	}
	if len(w1) != f*d {
		return nil, fmt.Errorf("nnproto: w1 length mismatch (got %d, want %d)", len(w1), f*d)
	}
	if len(w2) != d*f {
		return nil, fmt.Errorf("nnproto: w2 length mismatch (got %d, want %d)", len(w2), d*f)
	}

	fc1, err := Conv1x1("fc1", "x", "h1", d, f, w1)
	if err != nil {
		return nil, err
	}
	fc2, err := Conv1x1("fc2", "g", "out", f, d, w2)
	if err != nil {
		return nil, err
	}

	m := Model{
		Inputs:  []Feature{{Name: "x", Shape: []int64{int64(d), 1, int64(seq)}}},
		Outputs: []Feature{{Name: "out", Shape: []int64{int64(d), 1, int64(seq)}}},
		Layers: []Layer{
			fc1,
			ReLU("relu1", "h1", "r1"),
			Multiply("sq", []string{"r1", "r1"}, "g"),
			fc2,
		},
	}
	return m.Encode()
}
