// Generates modelir_transformer.mil: real MIL text emitted by
// github.com/tmc/modelir for a 1-layer stateful transformer with a linear FFN
// and an LM head, the shape mlx-go-ane compiles.
//
// It is not built as part of this module; modelir is not a dependency of
// github.com/tmc/apple, and the point of the bridge this file exercises is
// that it does not need to be. To regenerate, in a scratch module requiring
// github.com/tmc/modelir:
//
//	go run gen_modelir_transformer.go <outdir>
//
// then copy <outdir>/model.mil here. The weight blobs are not committed: the
// test synthesizes them from the BLOBFILE paths in the MIL text, which is what
// the weight-root path consumes.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ane "github.com/tmc/modelir/target/mil/ane"
)

func fill(n int, seed float32) []float32 {
	v := make([]float32, n)
	for i := range v {
		v[i] = seed + float32(i%7)*0.125
	}
	return v
}

func main() {
	outDir := os.Args[1]
	cfg := ane.NormalizeMILTransformerConfig(ane.MILTransformerConfig{
		NumLayers: 1, Dim: 8, NumHeads: 2, HeadDim: 4, HiddenDim: 16,
		VocabSize: 12, RMSNormEps: 1e-5, MaxSeqLen: 8,
		KVCache: ane.KVCacheModeState, KVCacheMaxLen: 4,
		LinearFFN: true, IncludeLMHead: true,
	})
	d, ad, hd, hid, V := cfg.Dim, ane.TransformerAttentionDim(cfg), cfg.HeadDim, cfg.HiddenDim, cfg.VocabSize
	w := ane.MILTransformerWeights{
		FinalNorm:   fill(d, 1),
		RopeCos:     fill(cfg.MaxSeqLen*hd, 0.5),
		RopeSin:     fill(cfg.MaxSeqLen*hd, 0.25),
		RopeRotateW: ane.BuildRoPERotateMatrix(ad, hd),
		RopeRotateB: fill(ad, 0),
		LMHeadW:     fill(V*d, 0.1),
		LMHeadB:     fill(V, 0),
	}
	for i := 0; i < cfg.NumLayers; i++ {
		w.Layers = append(w.Layers, ane.MILTransformerLayerWeights{
			QW: fill(ad*d, 0.1), QB: fill(ad, 0),
			KW: fill(ad*d, 0.2), KB: fill(ad, 0),
			VW: fill(ad*d, 0.3), VB: fill(ad, 0),
			OW: fill(d*ad, 0.4), OB: fill(d, 0),
			W1: fill(hid*d, 0.5), B1: fill(hid, 0),
			W3: fill(hid*d, 0.6), B3: fill(hid, 0),
			W2: fill(d*hid, 0.7), B2: fill(d, 0),
			InputNorm: fill(d, 1), PostAttentionNorm: fill(d, 1),
			QNorm: fill(hd, 1), KNorm: fill(hd, 1),
		})
	}
	w = ane.EnsureMILTransformerWeights(cfg, w)
	if err := ane.ValidateMILTransformerWeights(cfg, w); err != nil {
		fmt.Fprintln(os.Stderr, "validate weights:", err)
		os.Exit(1)
	}
	txt, files, err := ane.BuildMILTransformerArtifacts(cfg, w)
	if err != nil {
		fmt.Fprintln(os.Stderr, "build:", err)
		os.Exit(1)
	}
	os.MkdirAll(outDir, 0o755)
	if err := os.WriteFile(filepath.Join(outDir, "model.mil"), []byte(txt), 0o644); err != nil {
		panic(err)
	}
	var total int
	for _, f := range files {
		rel := strings.TrimPrefix(f.Path, "@model_path/")
		dst := filepath.Join(outDir, "weightroot", filepath.FromSlash(rel))
		os.MkdirAll(filepath.Dir(dst), 0o755)
		if err := os.WriteFile(dst, f.Blob, 0o644); err != nil {
			panic(err)
		}
		total += len(f.Blob)
	}
	fmt.Printf("mil=%d bytes, %d weight files, %d weight bytes\n", len(txt), len(files), total)
}
