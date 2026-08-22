//go:build darwin

package mil

import (
	"strings"
	"testing"
)

func TestGenConv(t *testing.T) {
	tests := []struct {
		name            string
		inCh, outCh, sp int
		wantContains    []string
	}{
		{
			name: "4x8 spatial 1",
			inCh: 4, outCh: 8, sp: 1,
			wantContains: []string{
				"tensor<fp32, [1, 4, 1, 1]> x",
				"tensor<fp16, [8, 4, 1, 1]> W",
				"BLOBFILE(",
				"func main<ios18>",
				"conv(",
			},
		},
		{
			name: "16x16 spatial 4",
			inCh: 16, outCh: 16, sp: 4,
			wantContains: []string{
				"tensor<fp32, [1, 16, 1, 4]> x",
				"tensor<fp16, [16, 16, 1, 1]> W",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenConv(tt.inCh, tt.outCh, tt.sp)
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("GenConv output missing %q\ngot:\n%s", want, got)
				}
			}
		})
	}
}

func TestBuildWeightBlob(t *testing.T) {
	weights := []float32{1, 0, 0, 1}
	blob, err := BuildWeightBlob(weights, 2, 2)
	if err != nil {
		t.Fatal(err)
	}

	// File header: 64 bytes, chunk header: 64 bytes, data: 4*2 = 8 bytes.
	if len(blob) != 64+64+8 {
		t.Errorf("blob size = %d, want %d", len(blob), 64+64+8)
	}

	// Check file header.
	if blob[0] != 0x01 {
		t.Errorf("file magic byte 1 = 0x%02X, want 0x01", blob[0])
	}
	if blob[4] != 0x02 {
		t.Errorf("file magic byte 2 = 0x%02X, want 0x02", blob[4])
	}

	// Check chunk header at offset 64.
	magic := uint32(blob[64]) | uint32(blob[65])<<8 | uint32(blob[66])<<16 | uint32(blob[67])<<24
	if magic != 0xDEADBEEF {
		t.Errorf("chunk magic = 0x%X, want 0xDEADBEEF", magic)
	}
}

func TestBuildWeightBlobWrongSize(t *testing.T) {
	_, err := BuildWeightBlob([]float32{1, 2, 3}, 2, 2)
	if err == nil {
		t.Error("expected error for wrong weight count")
	}
}

func TestBuildIdentityWeightBlob(t *testing.T) {
	blob, err := BuildIdentityWeightBlob(4)
	if err != nil {
		t.Fatal(err)
	}
	// 4*4 = 16 weights, 16*2 = 32 bytes of fp16 data + 128 header.
	if len(blob) != 64+64+32 {
		t.Errorf("blob size = %d, want %d", len(blob), 64+64+32)
	}
}

func TestBuildRoPECosSinBlobsWithTheta(t *testing.T) {
	// theta=10000 should match the default BuildRoPECosSinBlobs output.
	cos1, sin1, err := BuildRoPECosSinBlobs(4, 8)
	if err != nil {
		t.Fatal(err)
	}
	cos2, sin2, err := BuildRoPECosSinBlobsWithTheta(4, 8, 10000)
	if err != nil {
		t.Fatal(err)
	}
	if string(cos1) != string(cos2) {
		t.Error("cos blobs differ for theta=10000 vs default")
	}
	if string(sin1) != string(sin2) {
		t.Error("sin blobs differ for theta=10000 vs default")
	}

	// theta=100000 should produce different output.
	cos3, sin3, err := BuildRoPECosSinBlobsWithTheta(4, 8, 100000)
	if err != nil {
		t.Fatal(err)
	}
	if string(cos1) == string(cos3) {
		t.Error("cos blobs should differ for theta=100000 vs 10000")
	}
	if string(sin1) == string(sin3) {
		t.Error("sin blobs should differ for theta=100000 vs 10000")
	}
}

func TestBuildRoPECosSinBlobsWithThetaErrors(t *testing.T) {
	if _, _, err := BuildRoPECosSinBlobsWithTheta(0, 8, 10000); err == nil {
		t.Error("expected error for seq=0")
	}
	if _, _, err := BuildRoPECosSinBlobsWithTheta(4, 3, 10000); err == nil {
		t.Error("expected error for odd headDim")
	}
}

func TestGenFFNForwardReLU2(t *testing.T) {
	got := GenFFNForwardReLU2(64, 128, 4)
	for _, want := range []string{"relu(", "mul(x=r1,y=r1)", "w1.bin", "w2.bin"} {
		if !strings.Contains(got, want) {
			t.Errorf("GenFFNForwardReLU2 missing %q", want)
		}
	}
	// Should NOT contain W3 or silu.
	for _, bad := range []string{"w3.bin", "sigmoid(", "silu"} {
		if strings.Contains(got, bad) {
			t.Errorf("GenFFNForwardReLU2 should not contain %q", bad)
		}
	}
}

func TestGenFFNBackwardReLU2(t *testing.T) {
	got := GenFFNBackwardReLU2(64, 128, 4)
	for _, want := range []string{"relu(", "val=fp16(2.0)", "w1t.bin", "w2t.bin"} {
		if !strings.Contains(got, want) {
			t.Errorf("GenFFNBackwardReLU2 missing %q", want)
		}
	}
	for _, bad := range []string{"w3t.bin", "sigmoid("} {
		if strings.Contains(got, bad) {
			t.Errorf("GenFFNBackwardReLU2 should not contain %q", bad)
		}
	}
}

func TestFP16RoundTrip(t *testing.T) {
	tests := []float32{0, 1, -1, 0.5, 3.14, 65504, -65504}
	for _, v := range tests {
		h := float32ToFP16(v)
		got := fp16ToFloat32(h)
		diff := v - got
		if diff < 0 {
			diff = -diff
		}
		limit := float32(0.01)
		if v != 0 {
			rel := diff / v
			if rel < 0 {
				rel = -rel
			}
			if rel > limit && diff > 0.001 {
				t.Errorf("FP16 round-trip(%v): got %v, diff %v", v, got, diff)
			}
		}
	}
}

// TestGenSDPANoScaleArg guards the invariant that the emitted
// scaled_dot_product_attention call carries no scale argument. The ios18
// operator has no such parameter and the ANE compiler rejects the entire
// program when one is present, so a reintroduction breaks every caller.
// This is a text check so that it still runs where no ANE is present;
// TestCompileGenerators/GenSDPA proves the program actually compiles.
func TestGenSDPANoScaleArg(t *testing.T) {
	text := GenSDPA(64, 4, 32)
	// Slice the argument list only: the operator's own name contains
	// "scale" as a substring, so matching the whole call always fires.
	call := text[strings.Index(text, "scaled_dot_product_attention("):]
	call = call[strings.Index(call, "(")+1 : strings.Index(call, ")")]
	if strings.Contains(call, "scale") {
		t.Errorf("scaled_dot_product_attention call has a scale argument: %s", call)
	}
	if !strings.Contains(text, "mul(x = Q, y = c_scale)") {
		t.Error("scale is no longer folded into the query")
	}
}

// TestGenStateOps checks the state ops against the coremltools op definitions:
// read_state takes a single required "input" bound to a declared state var
// (iOS18/states.py), and the coreml_update_state dialect op must appear in a
// serialized program only in its decomposed write_state/read_state form
// (backend/mil/load.py translate_coreml_update_state_op).
func TestGenStateOps(t *testing.T) {
	shape := [4]int{1, 1, 4, 64}

	read := GenReadState("kv", shape)
	for _, want := range []string{
		"state<tensor<fp16, [1, 1, 4, 64]>> kv",
		"read_state(input = kv)",
	} {
		if !strings.Contains(read, want) {
			t.Errorf("GenReadState missing %q\n%s", want, read)
		}
	}
	if strings.Contains(read, "read_state(name") {
		t.Errorf("GenReadState binds read_state's name attribute as an input\n%s", read)
	}

	update := GenUpdateState("kv", shape)
	for _, want := range []string{
		"state<tensor<fp16, [1, 1, 4, 64]>> kv",
		"write_state(input = kv, data = value)",
		"read_state(input = kv)",
	} {
		if !strings.Contains(update, want) {
			t.Errorf("GenUpdateState missing %q\n%s", want, update)
		}
	}
	if strings.Contains(update, "coreml_update_state") {
		t.Errorf("GenUpdateState emits an undecomposed dialect op\n%s", update)
	}
}
