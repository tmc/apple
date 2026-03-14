package espresso

import (
	"encoding/json"
	"testing"
)

func TestBuilderBasic(t *testing.T) {
	b := NewBuilder()
	b.InnerProduct("fc1", 128,
		WithBottom("data"),
		WithWeights(make([]float32, 128*64)),
		WithBias(make([]float32, 128)),
	)
	b.Activation("relu1", "relu")
	b.InnerProduct("fc2", 10,
		WithWeights(make([]float32, 10*128)),
	)
	b.Softmax("softmax1")

	netJSON, shapeJSON, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}

	var net map[string]any
	if err := json.Unmarshal(netJSON, &net); err != nil {
		t.Fatalf("invalid net JSON: %v", err)
	}
	if v := net["format_version"]; v != float64(200) {
		t.Errorf("format_version = %v, want 200", v)
	}
	layers, ok := net["layers"].(map[string]any)
	if !ok {
		t.Fatal("layers not a map")
	}
	if len(layers) != 4 {
		t.Errorf("got %d layers, want 4", len(layers))
	}

	// Check fc1 has weight and bias indices.
	fc1 := layers["fc1"].(map[string]any)
	if _, ok := fc1["blob_weights"]; !ok {
		t.Error("fc1 missing blob_weights")
	}
	if _, ok := fc1["blob_biases"]; !ok {
		t.Error("fc1 missing blob_biases")
	}

	// Check auto-connect: relu1 bottom should be fc1's top.
	relu1 := layers["relu1"].(map[string]any)
	bottom := relu1["bottom"].([]any)
	if len(bottom) != 1 || bottom[0] != "fc1" {
		t.Errorf("relu1 bottom = %v, want [fc1]", bottom)
	}

	// Shape JSON should parse.
	var shape map[string]any
	if err := json.Unmarshal(shapeJSON, &shape); err != nil {
		t.Fatalf("invalid shape JSON: %v", err)
	}
}

func TestBuilderWithWeights(t *testing.T) {
	b := NewBuilder()
	weights := []float32{1.0, 2.0, 3.0, 4.0}
	b.InnerProduct("fc1", 2,
		WithBottom("data"),
		WithWeights(weights),
	)

	_, _, weightsBin, err := b.BuildWithWeights()
	if err != nil {
		t.Fatal(err)
	}
	if len(weightsBin) != 4*4 { // 4 float32s
		t.Errorf("weights length = %d, want 16", len(weightsBin))
	}
}

func TestBuilderConvolution(t *testing.T) {
	b := NewBuilder()
	b.Convolution("conv1", 32, 3, 3,
		WithBottom("data"),
		WithPad(1, 1, 1, 1),
		WithStride(1, 1),
		WithGroups(1),
		WithHasReLU(true),
	)

	netJSON, _, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}

	var net map[string]any
	if err := json.Unmarshal(netJSON, &net); err != nil {
		t.Fatal(err)
	}
	layers := net["layers"].(map[string]any)
	conv := layers["conv1"].(map[string]any)

	if conv["type"] != "convolution" {
		t.Errorf("type = %v, want convolution", conv["type"])
	}
	if conv["Nx"] != float64(3) {
		t.Errorf("Nx = %v, want 3", conv["Nx"])
	}
	if conv["has_relu"] != float64(1) {
		t.Errorf("has_relu = %v, want 1", conv["has_relu"])
	}
}

func TestBuilderEmpty(t *testing.T) {
	b := NewBuilder()
	_, _, err := b.Build()
	if err == nil {
		t.Error("expected error for empty builder")
	}
}
