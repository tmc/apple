// Command modelchain runs a transformer block on the Neural Engine through the
// public x/ane API, chaining two compiled models by sharing an IOSurface, and
// checks the result against a float64 CPU reference.
//
// The other ANE examples all drive the private e5rt route directly. This one
// uses what an outside caller of the package has: [ane.Probe] for the device,
// [ane.Open] and [ane.Client.Compile] to build models from MIL text and weight
// blobs, [ane.ShareSurface] to hand one model's output to the next without a
// copy, and [ane.Model.Eval] to run them.
//
// # What sharing a surface means here
//
// Compile gives every model its own input and output IOSurfaces. ShareSurface
// replaces the second model's input surface with the first model's output
// surface, so the feed forward network reads the attention result exactly where
// the attention model left it. Nothing is copied and the host does not touch
// the intermediate activation.
//
// That is worth a control of its own, because a chain that carried the
// activation some other way would produce the same numbers. So the block is run
// twice: once before the two models are chained, where the feed forward network
// reads its own untouched input and the result must not match the reference,
// and once after, where it must. Checking only that the surfaces are now the
// same object would show that the call assigned something, not that the
// assignment is what carries the data.
//
// # The residual the generator does not add
//
// GenFFNForwardRMSReLU2 computes W2(relu(W1(rms_norm(x)))^2) and stops. It has
// no residual add, so the block output is the attention result plus the feed
// forward result, and this program does that last add on the host. The e5rt
// examples close it with a third program on the device; through the public API
// there are two models and two Eval calls, so the host is already in the loop
// between them.
//
// # What is checked
//
// The reference recomputes the whole block in float64 from the same weights.
// Everything on the engine is fp16, so the comparison is to an fp16-scale
// tolerance, printed next to the magnitude of the reference rather than
// asserted as agreement. The weights are decorrelated so that the softmax
// selects rather than averages: an attention that averages over positions is
// reproduced exactly by a wrong head split and by a missing causal mask, which
// is how an earlier version of the e5rtblock example passed while verifying
// nothing. One input element is then perturbed and the block re-evaluated,
// which must move the output and must still match its own reference.
//
//	go run ./examples/ane/modelchain
//	go run ./examples/ane/modelchain -dim 128 -heads 8 -seq 32 -hidden 256
package main

import (
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

func main() {
	log.SetFlags(0)
	dim := flag.Int("dim", 64, "model width")
	heads := flag.Int("heads", 4, "attention heads")
	seq := flag.Int("seq", 16, "sequence length")
	hidden := flag.Int("hidden", 128, "feed forward width")
	flag.Parse()

	if *dim%*heads != 0 {
		log.Fatalf("dim %d is not a multiple of heads %d", *dim, *heads)
	}
	if err := run(*dim, *heads, *seq, *hidden); err != nil {
		log.Fatal(err)
	}
}

func run(dim, heads, seq, hidden int) error {
	info, err := ane.Probe()
	if err != nil {
		return fmt.Errorf("probe: %w", err)
	}
	if !info.HasANE {
		return fmt.Errorf("this machine reports no Neural Engine")
	}
	fmt.Printf("%s, %s, %d ANE core(s)\n\n", info.Product, info.Architecture, info.NumCores)

	c, err := ane.Open()
	if err != nil {
		return err
	}
	defer c.Close()

	w := newBlockWeights(dim, hidden)
	fmt.Printf("block: dim %d, heads %d, seq %d, hidden %d\n", dim, heads, seq, hidden)

	attnOpts, err := attentionOptions(dim, heads, seq, w)
	if err != nil {
		return err
	}
	attn, err := c.Compile(attnOpts)
	if err != nil {
		return fmt.Errorf("compile attention: %w", err)
	}
	defer attn.Close()

	ffnOpts, err := ffnOptions(dim, hidden, seq, w)
	if err != nil {
		return err
	}
	ffn, err := c.Compile(ffnOpts)
	if err != nil {
		return fmt.Errorf("compile feed forward: %w", err)
	}
	defer ffn.Close()

	fmt.Printf("  attention: %d input(s) %d output(s), output %q is %d channels over %d positions\n",
		attn.NumInputs(), attn.NumOutputs(), attn.OutputName(0), attn.OutputChannels(0), attn.Spatial(0))
	fmt.Printf("  feed forward: %d input(s) %d output(s)\n", ffn.NumInputs(), ffn.NumOutputs())

	acts := dim * seq
	input := make([]float32, acts)
	for i := range input {
		input[i] = float32(0.5 * math.Sin(float64(i)*0.21))
	}
	if err := checkDiscriminating(input, w, dim, heads, seq, hidden); err != nil {
		return err
	}
	want := blockReference(input, w, dim, heads, seq, hidden)

	// Negative control, before the chain exists. The two models own separate
	// surfaces, so the feed forward network reads whatever its own input
	// happens to hold rather than the attention result, and the block must not
	// match. Running this first is what makes the positive result below
	// evidence that ShareSurface is load-bearing: without it, a chain that
	// quietly copied somewhere else would look identical.
	unshared, err := evalBlock(attn, ffn, input, acts)
	if err != nil {
		return err
	}
	worst, _, tol := worstDiff(unshared, want)
	if worst <= tol {
		return fmt.Errorf("the block matched the reference before the two models were chained (worst |diff| %.4f within tolerance %.4f), so this program cannot show that sharing the surface is what carries the activation", worst, tol)
	}
	fmt.Printf("  unchained: worst |diff| %.4f against a tolerance of %.4f, as it must be\n", worst, tol)

	// The chain. After this the feed forward network reads the attention
	// model's output in place, with nothing copied.
	before := ffn.InputSurface(0)
	if err := ane.ShareSurface(attn, 0, ffn, 0); err != nil {
		return fmt.Errorf("share surface: %w", err)
	}
	if ffn.InputSurface(0) != attn.OutputSurface(0) {
		return fmt.Errorf("after ShareSurface the feed forward input is not the attention output")
	}
	if ffn.InputSurface(0) == before {
		return fmt.Errorf("ShareSurface left the feed forward input surface unchanged")
	}
	fmt.Printf("  chained: the feed forward input is now the attention output surface\n")

	got, err := evalBlock(attn, ffn, input, acts)
	if err != nil {
		return err
	}
	if err := compare("block output", got, want); err != nil {
		return err
	}

	// Control: move one input element and require the output to follow.
	perturbed := append([]float32(nil), input...)
	perturbed[0] += 1
	got2, err := evalBlock(attn, ffn, perturbed, acts)
	if err != nil {
		return err
	}
	moved := 0
	for i := range got {
		if got[i] != got2[i] {
			moved++
		}
	}
	if moved == 0 {
		return fmt.Errorf("perturbing an input changed nothing; the models are not reading their input surfaces")
	}
	want2 := blockReference(perturbed, w, dim, heads, seq, hidden)
	if err := compare("perturbed block output", got2, want2); err != nil {
		return err
	}
	fmt.Printf("  control: perturbing one input element moved %d of %d outputs, and the new output still matches its reference\n",
		moved, len(got))
	return nil
}

// evalBlock runs both models and returns the block output.
//
// There is no write between the two Eval calls. The feed forward network's
// input surface is the attention model's output surface, so evaluating the
// first leaves its result where the second reads it.
func evalBlock(attn, ffn *ane.Model, input []float32, acts int) ([]float32, error) {
	if err := attn.WriteInputFP16(0, input); err != nil {
		return nil, fmt.Errorf("write input: %w", err)
	}
	if err := attn.Eval(); err != nil {
		return nil, fmt.Errorf("eval attention: %w", err)
	}
	if err := ffn.Eval(); err != nil {
		return nil, fmt.Errorf("eval feed forward: %w", err)
	}

	// The block's second residual, which the feed forward generator omits.
	x2 := make([]float32, acts)
	if err := attn.ReadOutputFP16(0, x2); err != nil {
		return nil, fmt.Errorf("read attention output: %w", err)
	}
	f := make([]float32, acts)
	if err := ffn.ReadOutputFP16(0, f); err != nil {
		return nil, fmt.Errorf("read feed forward output: %w", err)
	}
	out := make([]float32, acts)
	for i := range out {
		out[i] = x2[i] + f[i]
	}
	return out, nil
}

// attentionOptions builds the compile options for the attention model and its
// six weight files.
func attentionOptions(dim, heads, seq int, w *blockWeights) (ane.CompileOptions, error) {
	var files []ane.WeightFile
	add := func(path string, blob []byte, err error) error {
		if err != nil {
			return err
		}
		files = append(files, ane.WeightFile{Path: path, Blob: blob})
		return nil
	}
	rms1, err := mil.BuildVectorWeightBlob(w.rms1)
	if err := firstErr(err, add("@model_path/weights/rms1.bin", rms1, nil)); err != nil {
		return ane.CompileOptions{}, err
	}
	for _, spec := range []struct {
		path        string
		data        []float32
		outCh, inCh int
	}{
		{"@model_path/weights/wq.bin", w.wq, dim, dim},
		{"@model_path/weights/wk.bin", w.wk, dim, dim},
		{"@model_path/weights/wv.bin", w.wv, dim, dim},
		{"@model_path/weights/wo.bin", w.wo, dim, dim},
	} {
		blob, err := mil.BuildWeightBlob(spec.data, spec.outCh, spec.inCh)
		if err := firstErr(err, add(spec.path, blob, nil)); err != nil {
			return ane.CompileOptions{}, err
		}
	}
	mask, err := mil.BuildCausalMaskBlob(seq)
	if err := firstErr(err, add("@model_path/weights/mask.bin", mask, nil)); err != nil {
		return ane.CompileOptions{}, err
	}
	return ane.CompileOptions{
		ModelType:   ane.ModelTypeMIL,
		MILText:     []byte(mil.GenSDPAForward(dim, heads, seq)),
		WeightFiles: files,
	}, nil
}

// ffnOptions builds the compile options for the feed forward model and its
// three weight files.
func ffnOptions(dim, hidden, seq int, w *blockWeights) (ane.CompileOptions, error) {
	rms2, err := mil.BuildVectorWeightBlob(w.rms2)
	if err != nil {
		return ane.CompileOptions{}, err
	}
	w1, err := mil.BuildWeightBlob(w.w1, hidden, dim)
	if err != nil {
		return ane.CompileOptions{}, err
	}
	w2, err := mil.BuildWeightBlob(w.w2, dim, hidden)
	if err != nil {
		return ane.CompileOptions{}, err
	}
	return ane.CompileOptions{
		ModelType: ane.ModelTypeMIL,
		MILText:   []byte(mil.GenFFNForwardRMSReLU2(dim, hidden, seq)),
		WeightFiles: []ane.WeightFile{
			{Path: "@model_path/weights/rms2.bin", Blob: rms2},
			{Path: "@model_path/weights/w1.bin", Blob: w1},
			{Path: "@model_path/weights/w2.bin", Blob: w2},
		},
	}, nil
}

func firstErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
