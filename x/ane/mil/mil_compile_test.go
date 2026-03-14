//go:build darwin

package mil_test

import (
	"errors"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

func openOrSkip(t *testing.T) *ane.Client {
	t.Helper()
	c, err := ane.Open()
	if errors.Is(err, ane.ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestCompileGenerators(t *testing.T) {
	c := openOrSkip(t)
	defer c.Close()

	tests := []struct {
		name string
		opts ane.CompileOptions
	}{
		{
			name: "GenIdentity",
			opts: func() ane.CompileOptions {
				blob, err := mil.BuildIdentityWeightBlob(4)
				if err != nil {
					t.Fatal(err)
				}
				return ane.CompileOptions{
					ModelType:  ane.ModelTypeMIL,
					MILText:    []byte(mil.GenIdentity(4, 1)),
					WeightBlob: blob,
				}
			}(),
		},
		{
			name: "GenIdentityFP16IO",
			opts: func() ane.CompileOptions {
				blob, err := mil.BuildIdentityWeightBlob(4)
				if err != nil {
					t.Fatal(err)
				}
				return ane.CompileOptions{
					ModelType:  ane.ModelTypeMIL,
					MILText:    []byte(mil.GenIdentityFP16IO(4, 1)),
					WeightBlob: blob,
				}
			}(),
		},
		{
			name: "GenConvFP32",
			opts: func() ane.CompileOptions {
				blob, err := mil.BuildWeightBlob(identityWeights(4), 4, 4)
				if err != nil {
					t.Fatal(err)
				}
				return ane.CompileOptions{
					ModelType:  ane.ModelTypeMIL,
					MILText:    []byte(mil.GenConvFP32(4, 4, 1)),
					WeightBlob: blob,
				}
			}(),
		},
		{
			name: "GenScaleFP16IO",
			opts: func() ane.CompileOptions {
				blob, err := mil.BuildWeightBlob([]float32{2.0}, 1, 1)
				if err != nil {
					t.Fatal(err)
				}
				return ane.CompileOptions{
					ModelType:  ane.ModelTypeMIL,
					MILText:    []byte(mil.GenScaleFP16IO(4)),
					WeightBlob: blob,
				}
			}(),
		},
		{
			name: "GenRMSNorm",
			opts: func() ane.CompileOptions {
				blob, err := mil.BuildWeightBlobV1(onesWeights(4))
				if err != nil {
					t.Fatal(err)
				}
				return ane.CompileOptions{
					ModelType:  ane.ModelTypeMIL,
					MILText:    []byte(mil.GenRMSNorm(4, 1, 1e-6)),
					WeightBlob: blob,
				}
			}(),
		},
		{
			name: "GenGQAExpand",
			opts: ane.CompileOptions{
				ModelType: ane.ModelTypeMIL,
				MILText:   []byte(mil.GenGQAExpand(2, 8, 64, 1)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := c.Compile(tt.opts)
			if err != nil {
				t.Fatal(err)
			}
			defer m.Close()

			if err := m.Eval(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func identityWeights(n int) []float32 {
	w := make([]float32, n*n)
	for i := range n {
		w[i*n+i] = 1.0
	}
	return w
}

func onesWeights(n int) []float32 {
	w := make([]float32, n)
	for i := range w {
		w[i] = 1.0
	}
	return w
}
