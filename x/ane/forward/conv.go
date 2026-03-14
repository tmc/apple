package forward

import (
	"fmt"

	"github.com/tmc/apple/x/ane/model"
)

// ConvEval mirrors training/forward.h ane_conv_eval:
// input row-major [S, inDim] -> transpose to channel-first [inDim, S], eval, transpose back.
func ConvEval(k *model.Kernel, x []float32, s, inDim, outDim int) ([]float32, error) {
	out, _, err := ConvEvalWithStats(k, x, s, inDim, outDim)
	return out, err
}

// ConvEvalWithStats is ConvEval plus kernel eval stats from EvalWithStats.
func ConvEvalWithStats(k *model.Kernel, x []float32, s, inDim, outDim int) ([]float32, model.EvalStats, error) {
	var est model.EvalStats
	if len(x) != s*inDim {
		return nil, est, fmt.Errorf("ConvEval input len=%d want=%d", len(x), s*inDim)
	}
	in := make([]float32, len(x))
	for t := 0; t < s; t++ {
		for i := 0; i < inDim; i++ {
			in[i*s+t] = x[t*inDim+i]
		}
	}

	if err := k.WriteInputF32(0, in); err != nil {
		return nil, est, fmt.Errorf("write input: %w", err)
	}
	ev, err := k.EvalWithStats()
	if err != nil {
		return nil, est, fmt.Errorf("eval: %w", err)
	}
	est = ev

	outT := make([]float32, s*outDim)
	if err := k.ReadOutputF32(0, outT); err != nil {
		return nil, est, fmt.Errorf("read output: %w", err)
	}

	out := make([]float32, s*outDim)
	for t := 0; t < s; t++ {
		for i := 0; i < outDim; i++ {
			out[t*outDim+i] = outT[i*s+t]
		}
	}
	return out, est, nil
}

// ToChannelFirst transposes row-major [s, d] to channel-first [d, s].
func ToChannelFirst(x []float32, s, d int) ([]float32, error) {
	if len(x) != s*d {
		return nil, fmt.Errorf("ToChannelFirst input len=%d want=%d", len(x), s*d)
	}
	out := make([]float32, len(x))
	for t := 0; t < s; t++ {
		for i := 0; i < d; i++ {
			out[i*s+t] = x[t*d+i]
		}
	}
	return out, nil
}

// FromChannelFirst transposes channel-first [d, s] to row-major [s, d].
func FromChannelFirst(x []float32, s, d int) ([]float32, error) {
	if len(x) != s*d {
		return nil, fmt.Errorf("FromChannelFirst input len=%d want=%d", len(x), s*d)
	}
	out := make([]float32, len(x))
	for t := 0; t < s; t++ {
		for i := 0; i < d; i++ {
			out[t*d+i] = x[i*s+t]
		}
	}
	return out, nil
}
