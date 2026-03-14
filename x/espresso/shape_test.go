package espresso

import "testing"

func TestReadEspressoShape(t *testing.T) {
	data := []byte(`{
  "layer_shapes": {
    "input": {"n": 1, "k": 3, "h": 224, "w": 224},
    "conv1": {"n": 1, "k": 64, "h": 112, "w": 112},
    "output": {}
  }
}`)

	m, err := ReadEspressoShape(data)
	if err != nil {
		t.Fatal(err)
	}

	input, ok := m.Get("input")
	if !ok {
		t.Fatal("missing input shape")
	}
	if input.C != 3 || input.H != 224 || input.W != 224 || input.N != 1 {
		t.Errorf("input shape = %+v, want N=1 C=3 H=224 W=224", input)
	}
	if input.Elements() != 1*3*224*224 {
		t.Errorf("input elements = %d, want %d", input.Elements(), 1*3*224*224)
	}

	conv, ok := m.Get("conv1")
	if !ok {
		t.Fatal("missing conv1 shape")
	}
	if conv.C != 64 {
		t.Errorf("conv1 channels = %d, want 64", conv.C)
	}

	// Empty shape entry should parse without error.
	output, ok := m.Get("output")
	if !ok {
		t.Fatal("missing output shape")
	}
	if output.Elements() != 0 {
		t.Errorf("empty output elements = %d, want 0", output.Elements())
	}

	names := m.Names()
	if len(names) != 3 {
		t.Errorf("names count = %d, want 3", len(names))
	}
}

func TestReadEspressoShapeInvalid(t *testing.T) {
	_, err := ReadEspressoShape([]byte(`not json`))
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}
