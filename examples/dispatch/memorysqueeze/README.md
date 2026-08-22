# memorysqueeze — memory-pressure-aware caching from Go

A stand-in for an inference server holding a large KV cache: the process
serves tokens while growing an mmap'd cache, and with `-shed` a
`DISPATCH_SOURCE_TYPE_MEMORYPRESSURE` source (dispatch.NewMemoryPressureSource)
drops half the cache on WARN and all but the newest chunk on CRITICAL.
Without `-shed` the footprint only grows — under a real squeeze that twin
is the reclaim/jetsam target.

Every heartbeat carries its own evidence: the process reads its
phys_footprint from `task_info(TASK_VM_INFO)` — the kernel's ledgered
number, the one Activity Monitor shows and jetsam uses.

## Run

```
go run ./examples/dispatch/memorysqueeze -shed    # terminal 1
go run ./examples/dispatch/memorysqueeze          # terminal 2
```

Then squeeze, either way:

```
sudo memory_pressure -S -l warn      # simulate WARN (no real pressure)
memory_pressure -l warn              # apply REAL pressure (no sudo; ^C to release)
```

Real pressure allocates until the system-wide WARN notification fires
(on a mostly-free machine that is most of RAM and takes a minute or
two), then holds. Kill it to release.

## Verified 2026-08-11 (macOS 26.x, 128 GiB machine, real pressure)

```
[shed pid=43850] serving tokens=2050 cache=1024 MiB footprint=1029 MiB
[shed pid=43850] SHED level=warn dropped=512 MiB footprint 1029 MiB -> 517 MiB
[shed pid=43850] serving tokens=2060 cache=512 MiB footprint=517 MiB
```

WARN arrived when the squeezer stabilized around 47% free; the shed
handler ran on its dispatch queue, munmap returned the pages
immediately (footprint 1029 → 517 MiB on the same line), and token
serving never paused. The cache then regrew — refill cost is the honest
price of shedding, same as a KV cache repopulating.

## Notes

- Footprint tracks the cache byte-for-byte (~5 MiB Go runtime overhead
  at steady state), which is what makes munmap-backed chunks the right
  cache shape for pressure response: freeing them is visible to the
  kernel's ledger on the next read.
- `task_info` constants (flavor 22, count 93, phys_footprint at offset
  144) were verified against `<mach/task_info.h>` with a C probe;
  `kernel.Task_vm_info_data_t` is opaque storage, so the field is read
  by offset.
- The real integration this example stands in for: an LLM server (e.g.
  mlx-go-lm's lmserve) dropping speculative/prefix caches on WARN and
  KV beyond active sequences on CRITICAL. See
  `design/mlx-demo-ideas.md` §1.
