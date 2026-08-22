# surfaceffn — the Neural Engine computes over pages another process wrote

The ANE arm of the surfacecompute family (`examples/iosurface/tensorshare`,
`examples/metal/surfacecompute`). Two Go processes, zero tensor copies:

- The **consumer** compiles a dynamic matmul (`x/ane/mil`) on the ANE. The
  compile allocates IOSurfaces in exactly the strided layout the NPU DMAs
  from — so instead of receiving a surface, the consumer *sends* its
  kernel's own input and output surfaces to the producer as mach port
  rights, then evaluates on command. It never reads or writes one element.
- The **producer** writes activations and weights straight into the NPU's
  input pages (`x/ane.WrapIOSurfaceFloat32WithLayout`), triggers an eval,
  and reads the product from the NPU's output pages.

Proofs, mirroring surfacecompute:

1. The result read from the ANE's output surface matches the producer's
   float64 CPU reference (fp16-scale tolerance 2e-2 — the ANE accumulates
   in reduced precision and the demo says so).
2. The producer mutates one activation in place; the consumer re-evaluates
   with no new writes on its side, and the recomputed product appears — the
   kernel is bound to live pages, not a snapshot.

```
go run ./examples/ane/surfaceffn
go run ./examples/ane/surfaceffn -in 1024 -out 1024 -batch 128
```

Measured 2026-08-11 (M-series, macOS 26.x): 512×512×128 — 124µs wall per
eval including the pipe round trip (543 GFLOPS), 81µs hardware-reported
on-device (828 GFLOPS). 1024×1024×128 — 204µs wall (1318 GFLOPS), 196µs
on-device (1370 GFLOPS).

The ANE compile path is private-framework territory (`_ANEClient`) and is
presence-checked: where unavailable the consumer reports `unavail` and the
demo exits cleanly.
