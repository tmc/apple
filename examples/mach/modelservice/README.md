# modelservice — one copy of the weights, many processes

The shape of Apple's on-device model daemon, built from this repo's
mach and iosurface bindings: a service loads "model weights" into an
IOSurface once, registers a bootstrap port, and any number of client
processes use them — token streams over raw mach messages, and the
weights themselves attachable zero-copy via a port-passed surface.

Transport note: this is `x/mach` bootstrap + mach messages (the route
`examples/iosurface/tensorshare` proved), not XPC — `apple/xpc` is not
generated into the module yet (see design/xpc.md's gates). If it ships,
this demo is its first adopter.

Two verbs:

- **generate** — client sends a prompt and a reply port; the service
  streams tokens back one mach message each. Each token folds a strided
  read of the weight pages into its seed, so the weights are
  load-bearing: unmap them and tokens stop.
- **attach** — the service mints a send right with
  `IOSurfaceCreateMachPort` and hands the whole weights surface to the
  client (`MoveSend`), which maps the same physical pages and verifies
  the checksum. Gigabytes cross the process boundary, zero bytes copied.

Every process prints its phys_footprint from `task_info(TASK_VM_INFO)`
— the kernel's ledger, the number Activity Monitor and jetsam use — so
the "one copy" claim carries its own evidence.

## Run

```
go run ./examples/mach/modelservice              # orchestrated demo
go run ./examples/mach/modelservice -mib 4096    # 4 GiB of weights
```

The orchestrated mode spawns the service and four clients: two
generating concurrently, one attaching the weights, and one SIGKILLed
mid-stream — the service logs the failed send and keeps serving.

Manual mode, separate terminals:

```
go run ./examples/mach/modelservice -serve -service com.tmc.modelservice
go run ./examples/mach/modelservice -client -service com.tmc.modelservice -prompt "hello"
go run ./examples/mach/modelservice -attach -service com.tmc.modelservice
```

## Verified 2026-08-11 (macOS 26.x, Apple silicon)

```
[service] service: weights loaded: 1024 MiB, checksum 16777215754, footprint 1032 MiB
[service] service: listening on bootstrap "com.tmc.modelservice.48303"
[attach-3] attach: verified 1024 MiB of weights in place at 4.65 GiB/s — zero bytes copied; footprint 6 MiB -> 7 MiB
demo: SIGKILLed victim-4 mid-stream
[service] service: generate #3: client vanished mid-stream (mach: mach_msg send: 0x10000003) — stream dropped, service continues
[client-1] client: "unified memory serves many masters" -> "…12 tokens…" (footprint 6 MiB)
demo: all surviving clients done; one weights copy (1024 MiB) served 2 generations + 1 attach + 1 dead client
```

The ledger line is the demo: the attach client reads the full 1 GiB of
weights in place and its footprint moves 6 → 7 MiB. The pages are
resident once, in the service, and every client sees them.

Client death is handled where it surfaces: the dead client's reply port
becomes a dead name, `mach_msg` returns `MACH_SEND_INVALID_DEST`
(0x10000003), and the service drops that stream and continues.

## What mlx-go-lm would keep

The protocol shape (bootstrap rendezvous, reply-port streaming, surface
handoff for large tensors) is exactly what an lmserve daemon needs for
"N CLI tools, one resident model": swap the synthetic weights for MLX
buffers and the stand-in decode for real generation. Pairs with
`examples/dispatch/memorysqueeze` (the service is the natural place for
the pressure hook) and row 12 of design/mlx-demo-ideas.md (a service
that forks sessions).
