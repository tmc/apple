package main

import (
	"fmt"
	"math"
)

// blockWeights are the parameters of one transformer block.
type blockWeights struct {
	rms1 []float32 // [dim]
	wq   []float32 // [dim][dim] OIHW
	wk   []float32
	wv   []float32
	wo   []float32
	rms2 []float32 // [dim]
	w1   []float32 // [hidden][dim]
	w2   []float32 // [dim][hidden]
}

// randomWeights returns n deterministic pseudo-random values in [-spread,
// spread), from a xorshift generator seeded by seed.
//
// Decorrelated values are the point. Correlated weights cancel under a dot
// product over hundreds of terms instead of accumulating like a random walk,
// and attention whose scores are all equal is an average over positions, which
// a wrong head split and a missing causal mask both reproduce exactly.
func randomWeights(n int, spread float64, seed uint64) []float32 {
	state := seed*2862933555777941757 + 3037000493
	next := func() float64 {
		state ^= state << 13
		state ^= state >> 7
		state ^= state << 17
		// The top 53 bits, mapped to [-1, 1).
		return float64(state>>11)/(1<<52) - 1
	}
	out := make([]float32, n)
	for i := range out {
		out[i] = float32(next() * spread)
	}
	return out
}

// newBlockWeights builds deterministic weights whose scales are chosen so that
// every stage of the block contributes.
//
// A projection sums over its input width, so a weight of scale s produces an
// output of roughly s*sqrt(width) for a unit-scale input; the scales below
// invert that. The query and key projections land near unit magnitude, which
// puts the scaled dot products over a range of a few and makes the softmax
// select rather than average. The output and second feed forward projections
// land near the residual's own magnitude, so neither result is lost against it.
func newBlockWeights(dim, hidden int) *blockWeights {
	rms := func(n int, seed uint64) []float32 {
		// RMSNorm weights vary around one: a constant would hide a channel misorder.
		out := randomWeights(n, 0.2, seed)
		for i := range out {
			out[i] += 1
		}
		return out
	}
	inv := func(width int) float64 { return 1 / math.Sqrt(float64(width)) }
	return &blockWeights{
		rms1: rms(dim, 11),
		wq:   randomWeights(dim*dim, inv(dim), 22),
		wk:   randomWeights(dim*dim, inv(dim), 33),
		wv:   randomWeights(dim*dim, inv(dim), 44),
		wo:   randomWeights(dim*dim, 0.5*inv(dim), 55),
		rms2: rms(dim, 66),
		w1:   randomWeights(hidden*dim, inv(dim), 77),
		w2:   randomWeights(dim*hidden, 0.5*inv(hidden), 88),
	}
}

// blockReference recomputes the block in float64.
func blockReference(x []float32, w *blockWeights, dim, heads, seq, hidden int) []float64 {
	headDim := dim / heads
	in := make([]float64, len(x))
	for i, v := range x {
		in[i] = float64(v)
	}

	// Attention: RMSNorm, QKV, scaled dot product with a causal mask, output
	// projection, residual.
	xn := rmsNorm(in, w.rms1, dim, seq)
	q := project(xn, w.wq, dim, dim, seq)
	k := project(xn, w.wk, dim, dim, seq)
	v := project(xn, w.wv, dim, dim, seq)

	// attn[c][s] in the same [dim][seq] layout the engine uses.
	attn := make([]float64, dim*seq)
	scale := 1 / math.Sqrt(float64(headDim))
	for h := range heads {
		for i := range seq {
			// Scores against every position at or before i.
			scores := make([]float64, i+1)
			maxScore := math.Inf(-1)
			for j := 0; j <= i; j++ {
				var dot float64
				for d := range headDim {
					c := h*headDim + d
					dot += q[c*seq+i] * k[c*seq+j]
				}
				scores[j] = dot * scale
				maxScore = math.Max(maxScore, scores[j])
			}
			var sum float64
			for j := range scores {
				scores[j] = math.Exp(scores[j] - maxScore)
				sum += scores[j]
			}
			for d := range headDim {
				c := h*headDim + d
				var acc float64
				for j := range scores {
					acc += scores[j] * v[c*seq+j]
				}
				attn[c*seq+i] = acc / sum
			}
		}
	}
	proj := project(attn, w.wo, dim, dim, seq)
	x2 := make([]float64, dim*seq)
	for i := range x2 {
		x2[i] = in[i] + proj[i]
	}

	// Feed forward: RMSNorm, W1, ReLU squared, W2. The residual is added here
	// because the generator leaves it out.
	fn := rmsNorm(x2, w.rms2, dim, seq)
	h1 := project(fn, w.w1, hidden, dim, seq)
	for i, v := range h1 {
		if v < 0 {
			h1[i] = 0
		} else {
			h1[i] = v * v
		}
	}
	f := project(h1, w.w2, dim, hidden, seq)
	out := make([]float64, dim*seq)
	for i := range out {
		out[i] = x2[i] + f[i]
	}
	return out
}

// checkDiscriminating reports whether the block, with these weights and this
// input, is actually exercising the parts it claims to.
//
// It is a check on the test, not on the engine. A block whose attention scores
// are all equal computes an average over positions, and then a wrong head
// split, a wrong query-key transpose and a missing causal mask all produce the
// same numbers. The same goes for a residual so much larger than what the block
// adds to it that the sum is the input again.
func checkDiscriminating(x []float32, w *blockWeights, dim, heads, seq, hidden int) error {
	headDim := dim / heads
	in := make([]float64, len(x))
	for i, v := range x {
		in[i] = float64(v)
	}
	xn := rmsNorm(in, w.rms1, dim, seq)
	q := project(xn, w.wq, dim, dim, seq)
	k := project(xn, w.wk, dim, dim, seq)

	scale := 1 / math.Sqrt(float64(headDim))
	lo, hi := math.Inf(1), math.Inf(-1)
	for h := range heads {
		for i := range seq {
			for j := 0; j <= i; j++ {
				var dot float64
				for d := range headDim {
					c := h*headDim + d
					dot += q[c*seq+i] * k[c*seq+j]
				}
				lo = math.Min(lo, dot*scale)
				hi = math.Max(hi, dot*scale)
			}
		}
	}
	spread := hi - lo
	fmt.Printf("  attention score spread %.3f", spread)
	if spread < 1 {
		fmt.Println()
		return fmt.Errorf("the softmax is nearly uniform, so attention degenerates to an average and this program cannot tell a correct block from several wrong ones")
	}

	// How much of the output the block contributes, against the residual.
	out := blockReference(x, w, dim, heads, seq, hidden)
	var deltaPeak, inPeak float64
	for i := range out {
		deltaPeak = math.Max(deltaPeak, math.Abs(out[i]-in[i]))
		inPeak = math.Max(inPeak, math.Abs(in[i]))
	}
	ratio := deltaPeak / inPeak
	fmt.Printf(", block contributes %.2fx the input magnitude\n", ratio)
	if ratio < 0.25 {
		return fmt.Errorf("the block barely moves its input, so the residual dominates and a broken block would still match")
	}
	return nil
}

// rmsNorm normalizes each position over the channel axis and applies the
// per-channel weight, matching the generators: mean of squares over channels,
// plus epsilon, raised to -1/2.
func rmsNorm(x []float64, weight []float32, dim, seq int) []float64 {
	out := make([]float64, dim*seq)
	for s := range seq {
		var sum float64
		for c := range dim {
			v := x[c*seq+s]
			sum += v * v
		}
		inv := math.Pow(sum/float64(dim)+1e-5, -0.5)
		for c := range dim {
			out[c*seq+s] = x[c*seq+s] * inv * float64(weight[c])
		}
	}
	return out
}

// project applies an OIHW 1x1 convolution: out[o][s] = sum_i w[o][i] * x[i][s].
func project(x []float64, weight []float32, outCh, inCh, seq int) []float64 {
	out := make([]float64, outCh*seq)
	for o := range outCh {
		row := weight[o*inCh : (o+1)*inCh]
		for s := range seq {
			var acc float64
			for i := range inCh {
				acc += float64(row[i]) * x[i*seq+s]
			}
			out[o*seq+s] = acc
		}
	}
	return out
}

// worstDiff returns the worst absolute difference and the peak magnitude of the
// reference, along with the tolerance the comparison uses.
//
// fp16 carries about three decimal digits and this block accumulates over four
// matrix products and a softmax, so the tolerance is proportional to the
// magnitude rather than absolute. A non-finite result is reported as an
// infinite difference rather than being silently skipped.
func worstDiff(got []float32, want []float64) (worst, peak, tol float64) {
	for i := range got {
		g := float64(got[i])
		if math.IsNaN(g) || math.IsInf(g, 0) {
			worst = math.Inf(1)
			continue
		}
		worst = math.Max(worst, math.Abs(g-want[i]))
	}
	for _, v := range want {
		peak = math.Max(peak, math.Abs(v))
	}
	return worst, peak, 0.02*peak + 0.01
}

// compare reports the worst difference against an fp16-scale tolerance. It
// prints the peak magnitude of the reference alongside, so a small difference
// against a small reference is not read as agreement.
func compare(label string, got []float32, want []float64) error {
	if len(got) != len(want) {
		return fmt.Errorf("%s: %d values, want %d", label, len(got), len(want))
	}
	worst, peak, tol := worstDiff(got, want)
	if worst > tol {
		return fmt.Errorf("%s: worst |diff| %.4f exceeds tolerance %.4f (reference peak |%.4f|)", label, worst, tol, peak)
	}
	fmt.Printf("  %s matches the float64 reference over %d values: worst |diff| %.4f, tolerance %.4f, reference peak |%.4f|\n",
		label, len(got), worst, tol, peak)
	return nil
}
