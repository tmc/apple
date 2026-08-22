# surfacecompute

GPU compute over an IOSurface shared between two Go processes, with zero
copies. The natural extension of `examples/iosurface/tensorshare`:
tensorshare proved two CPUs can share the pages; this proves the GPU reads
the same pages through the texture unit — producer writes from Go, Metal
kernel reads, nothing is blitted or uploaded.

```
go run ./examples/metal/surfacecompute
go run ./examples/metal/surfacecompute -width 8192 -height 8192
```

## What it proves, in order

1. **Zero-copy handoff to the GPU.** The producer fills a float32 IOSurface
   and passes its mach port to a spawned consumer (bootstrap rendezvous,
   port-right descriptor, `MoveSend` — tensorshare's production route). The
   consumer wraps the surface in a Metal texture with
   `newTextureWithDescriptor:iosurface:plane:` and dispatches a row-sum
   kernel. The GPU's checksum matches the sum the producer computed while
   writing — over pages the consumer process never copied.
2. **Live pages, not a snapshot.** The producer mutates one element in
   place; the consumer re-dispatches the same kernel and the GPU sum moves
   by exactly the delta. A texture backed by an upload would return the
   stale sum.

## Measured, 2026-08-11

macOS 26.x, Apple Silicon (arm64), go1.26.3. Best of 5–10 reps, same shared
pages in both arms:

| surface | GPU kernel sum | CPU mapped sum |
|---|---|---|
| 64 MiB (4096²) | 15.5 GiB/s | 9.7 GiB/s |
| 256 MiB (8192²) | 30.9 GiB/s | 10.6 GiB/s |

The GPU arm scales with surface height because the kernel is
row-parallel (one thread per row); the point is not the ratio but that the
texture unit is reading the producer's pages at memory-system rates with
no copy in the path.

## Native comparison — what the Go bindings cost

`compare/nativecompare.m` is the Objective-C twin of the consumer: same MSL
kernel, same IOSurface shape and pixel format, same best-of-reps timing.
The GPU work is identical by construction, so the difference is purely the
host-side cost of driving Metal from Go (purego + objc.Send) vs native
`objc_msgSend`.

```
clang -O2 -fobjc-arc -framework Metal -framework IOSurface \
      -framework CoreFoundation -o nativecompare compare/nativecompare.m
```

Measured 2026-08-11, best of 20 reps:

| surface | Go | ObjC | delta |
|---|---|---|---|
| 256 MiB (8192²) | 8.087 ms | 8.077 ms | +0.1% |
| 64 MiB (4096²) | 4.022 ms | 4.005 ms | +0.4% |
| 256 KiB (256²) | 282 µs | 150 µs | +132 µs |
| 16 KiB (64²) | 176 µs | 135 µs | +41 µs |

At workload sizes the arms are indistinguishable — the GPU is the clock.
The bindings' cost is visible only in the tiny-dispatch regime as tens of
microseconds per dispatch (a handful of objc.Send round trips plus
completion-wait plumbing), and it amortizes to noise as soon as the kernel
has real work. The Go and ObjC checksums agree in both arms, so the
comparison is validated, not assumed.

## Notes

- The IOSurface is created with pixel format `'L00f'`
  (one-component 32-bit float) to match `MTLPixelFormatR32Float`;
  `bytesPerRow` is `width*4`, which satisfies the linear-texture alignment
  for these widths — `newTextureWithDescriptor:iosurface:plane:` fails
  loudly if it does not.
- Checksums are compared under a tolerance: the kernel accumulates each
  row in float32 while the references accumulate in float64.
- The generated `NewTextureWithDescriptorIosurfacePlane` takes the
  `iosurface` package's `IOSurfaceRef` directly — no conversion from the
  `coregraphics` package's same-named typedef is needed on this path.
