# gomodel — a CoreML model authored, compiled, and run entirely from Go

No Python, no coremltools, no Xcode. One Go binary:

1. defines an MIL program (`y = relu(x·Wᵀ + b)`) as a plain Go value,
2. writes its weights as a MILBlob Storage v2 `weight.bin`,
3. encodes the model to CoreML's protobuf (`x/coremlcompiler.EncodeModel`),
4. writes an Apple-compatible `.mlpackage` (`WriteMLPackage`),
5. compiles it to `.mlmodelc` — by default with the repo's **pure-Go
   compiler** (no Apple tool invoked, ~1ms); `-applecompile` routes through
   `[MLModel compileModelAtURL:]` instead, proving Apple's own toolchain
   accepts the Go-authored package,
6. loads it with CoreML and runs a prediction, verified against the same
   math computed in Go.

```
go run ./examples/coreml/gomodel
go run ./examples/coreml/gomodel -applecompile
go run ./examples/coreml/gomodel -keep ~/tmp/gomodel   # inspect the products
```

Both compile arms verified 2026-08-11 (macOS 26.x, go1.26.3).

## Spec gotchas this example encodes (found by running both arms)

- **`const` ops carry `val` as an attribute, not an input.** The MIL-text
  emitter accepts either (it merges const inputs into attributes), so the
  pure-Go arm masks the mistake; Apple's protobuf parser rejects it.
- **The protobuf opset is `CoreMLn`, not `iosNN`.** `CoreML7` corresponds
  to ios18 in MIL text — the pure-Go compiler performs that rewrite when
  emitting `model.mil`. Apple's parser also requires the
  `block_specializations` key to equal the function's opset string exactly.
- **Spec version must match the opset**: 9 for CoreML7. A mismatch fails
  Apple's validator with "not written in an opset that is supported".
- **Multi-dimensional immediate tensors didn't survive MIL text** — the
  serializer emitted flat literals, which CoreML's parser reads as rank-1
  ("Declared shape [3, 4] but expected shape 12"). FIXED same day: the
  emitter now nests literals to match the declared shape
  (`TestCompileRank2ImmediateLoadable` proves a rank-2 immediate loads).
  This example still routes weights through `weight.bin`
  (`WriteMILBlob` + `BlobFileValue`) because that is what coremltools
  does and what real models need.
- **`WriteMLPackage`'s directory branch dropped the `weights/` prefix**
  (contents landed at `com.apple.CoreML/` and the compiled model couldn't
  find them). FIXED same day: a directory `weightSrc` is now the weights
  directory itself and its contents land under `weights/`.
