//go:build darwin

package ane

import (
	"testing"

	"github.com/tmc/apple/x/ane/mil"
)

// TestMultiChannelIdentity tests multi-channel identity conv round-trips
// across various channel and spatial dimensions.
func TestMultiChannelIdentity(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	tests := []struct {
		name    string
		inCh    int
		outCh   int
		spatial int
	}{
		{"1ch_1sp", 1, 1, 1},
		{"1ch_4sp", 1, 1, 4},
		{"2ch_1sp", 2, 2, 1},
		{"2ch_2sp", 2, 2, 2},
		{"2ch_16sp", 2, 2, 16},
		{"4ch_1sp", 4, 4, 1},
		{"4ch_16sp", 4, 4, 16},
		{"4ch_64sp", 4, 4, 64},
		{"16ch_16sp", 16, 16, 16},
		{"32ch_1sp", 32, 32, 1},
		{"32ch_32sp", 32, 32, 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			milText := mil.GenConvFP16IO(tt.inCh, tt.outCh, tt.spatial)
			blob, err := mil.BuildIdentityWeightBlob(tt.inCh)
			if err != nil {
				t.Fatal(err)
			}

			k, err := rt.Compile(CompileOptions{
				ModelType:  ModelTypeMIL,
				MILText:    []byte(milText),
				WeightBlob: blob,
			})
			if err != nil {
				t.Fatalf("compile: %v", err)
			}
			defer k.Close()

			// Log parsed layout.
			t.Logf("input layout: %+v", k.InputLayout(0))
			t.Logf("output layout: %+v", k.OutputLayout(0))

			nIn := tt.inCh * tt.spatial
			input := make([]float32, nIn)
			for i := range input {
				input[i] = float32(i%10 + 1)
			}

			if err := k.WriteInputFP16(0, input); err != nil {
				t.Fatal(err)
			}
			if err := k.Eval(); err != nil {
				t.Fatal(err)
			}

			nOut := tt.outCh * tt.spatial
			output := make([]float32, nOut)
			if err := k.ReadOutputFP16(0, output); err != nil {
				t.Fatal(err)
			}

			for i := range input {
				if i >= len(output) {
					break
				}
				diff := input[i] - output[i]
				if diff < -0.5 || diff > 0.5 {
					showN := 16
					if showN > len(output) {
						showN = len(output)
					}
					t.Logf("output[0:%d] = %v", showN, output[:showN])
					t.Logf("input[0:%d]  = %v", showN, input[:showN])
					t.Fatalf("mismatch at [%d]: got %v want %v", i, output[i], input[i])
				}
			}
		})
	}
}
