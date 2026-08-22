# tensorshare

Zero-copy tensor handoff between two Go processes via IOSurface.

The parent allocates a float32 IOSurface, fills it, and spawns this same
binary as a consumer. The consumer maps the surface, and verifies the
checksum — no bytes cross a pipe or socket. The parent then mutates one
element in place and the consumer observes the new value, proving both
processes see the same physical pages. Finally both routes are benchmarked:
reading the shared mapping in place versus streaming the same buffer
through a pipe.

The surface handoff itself is a mach port-right descriptor, the
production route: the consumer registers a receive right under a
bootstrap name, the producer sends the port from `IOSurfaceCreateMachPort`
in a mach message with `MoveSend`, and the consumer recovers the surface
via `IOSurfaceLookupFromMachPort` — all through `x/mach` (ports, messages,
bootstrap). `-transport=global` keeps the original deprecated
`IOSurfaceIsGlobal`/ID-lookup path for comparison.

    go run ./examples/iosurface/tensorshare

Measured on an arm64 Mac (macOS 26.x, go1.26.3), 64 MiB tensor, best of 5:

    bench: mapped in-place read      15.04 GiB/s  (64 MiB in 4.154ms)
    bench: pipe stream + read         2.65 GiB/s  (64 MiB in 23.562ms)

Both arms run the identical scalar float32 sum, so the ~5x gap is the
copy and syscall cost of the pipe, not a compute difference. Both
transports read the mapping at the same rate — the handoff mechanism
changes setup cost only. On a loaded machine raise -reps so best-of
finds quiet windows, and compare arms at equal reps.

## Seeing the shared mapping

    go run ./examples/iosurface/tensorshare -hold

pauses after verification and prints `vmmap` invocations for both live
processes. Each shows the surface as an `IOSurface` region with `SM=SHM` —
the same kernel-shared memory mapped into two address spaces.

## Notes

- The port right crosses with `MoveSend`, so the producer has nothing to
  balance after the send — the right lives in the message, then in the
  consumer. Getting the disposition wrong leaks a right or produces a dead
  name; `x/mach`'s tests assert both polarities.
- `bootstrap_register` is deprecated but remains the only route for an
  unprivileged process to publish a port at runtime; the rendezvous name is
  PID-scoped to avoid collisions between concurrent runs.
- With `-transport=global` the surface is created `IOSurfaceIsGlobal`
  (also deprecated) and found by bare ID — kept as the plumbing-free
  comparison arm.
- Locking discipline: the producer takes the full lock around writes; the
  consumer takes `kIOSurfaceLockReadOnly` around reads. Skipping the lock
  can read stale data on surfaces that ever touch the GPU.
- A subtle aliasing hazard this demo hit during development: a Go slice
  over the surface base address aliases the shared pages, so "save the old
  value, then mutate" must read the value before the in-place write — the
  slice does not snapshot.
