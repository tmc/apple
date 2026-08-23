# statecache

Keeps mutable state resident on the Neural Engine across executions, and across
two separately compiled programs, through one retained E5RT inout port.

Every other ANE example here is a pure function. A decoder is not: it carries
state between steps, and if that state has to be staged back to the host after
every token then the state, not the arithmetic, sets the pace.
`Lib.OperationRetainInoutPort` is the interface that avoids it. The package's
own probes proved it works; nothing had used it.

```sh
go run ./examples/ane/statecache
go run ./examples/ane/statecache -layout
go run ./examples/ane/statecache -chan 8 -dim 32 -steps 8
```

```
retained state: chan=4 dim=4 (16 values) steps=6

  state buffer: 256 bytes, 4 channels of 64 bytes — a packed 32 would be too small by 224
  a freshly allocated buffer object read back as all zeros (observed, not relied on)

  a host write to the bound state buffer is what the engine reads back

  6 updates match a float64 recurrence exactly, checked after every step
  the engine's writes are visible through the state buffer's own host pointer

  a separately compiled reader on the same buffer returns the accumulated state
  the same reader program bound to a different buffer returns zeros, as it must

  updating 8 values in the upper half left the lower half untouched

  two states advanced in alternation each followed their own recurrence

  refused, as it must be: an inout port the program does not declare
    e5rt: e5rt_execution_stream_operation_retain_inout_port: status 11
  and the declared port "kv" still resolves on the same operation

OK
```

## The state buffer is not packed

The most useful thing this example produced is a fact about buffer layout that
nothing in the MIL text tells you.

Each channel of a `[1, channels, 1, dim]` fp16 state starts on a 64-byte
boundary. The buffer a caller must allocate is
`channels * round_up(dim*2, 64)` bytes, not `channels * dim * 2`.

| shape | packed size | required size |
| --- | --- | --- |
| `[1, 4, 1, 4]` | 32 | 256 |
| `[1, 4, 1, 32]` | 256 | 256 |
| `[1, 4, 1, 40]` | 320 | 512 |
| `[1, 2, 1, 64]` | 256 | 256 |
| `[1, 8, 1, 2]` | 32 | 512 |

Shapes whose row already fills whole 64-byte lines are packed by coincidence,
which is a good way to write code that works on `dim=32` and corrupts memory on
`dim=4`.

Allocating the packed size produces no error. The engine reads and writes the
padded offsets anyway, so for an accumulate program it writes past the end of
the allocation — 224 bytes past, at `dim=4`. The first version of this example
did exactly that; the symptom was that channel 0 round-tripped and every other
channel came back zero.

`e5rt.StateLayout` owns this arithmetic now, and `Lib.BindStatePort` is a bind
that asks the runtime how large a buffer is — via `e5rt_buffer_object_get_size`,
which the library exports and which had no binding — and refuses one too small
for the shape. The raw retain-and-bind pair accepts an undersized buffer without
complaint; the package test asserts that as a mutation control, so if the raw
path ever started checking, the new one would be flagged as redundant rather
than quietly duplicated.

`-layout` measures it instead of citing it. It binds an oversized state buffer,
fills every 2-byte slot with its own index, and prints the slot each tensor
element was read from:

```
  state shape [1, 4, 1, 4] over a 1024-byte buffer
  each value below is the host 2-byte slot the engine read that element from:
    channel 0: [0 1 2 3]
    channel 1: [32 33 34 35]
    channel 2: [64 65 66 67]
    channel 3: [96 97 98 99]
  the state tensor is NOT host-packed; a caller must write through the layout above
```

The ordinary `value` and `y` ports are packed — the layout probe reads its
16-element result out of a 32-byte output buffer, which would be impossible if
outputs carried the same padding. Only the state has it.

## What each check would catch

| check | what it would catch |
| --- | --- |
| a float64 recurrence, compared after **every** step | a dropped, doubled, or misordered update; a correct total arrived at wrongly |
| a host-written pattern read back through the engine, before anything depends on it | a wrong buffer layout, and a host pointer the engine does not observe — without which the zeroing between arms is not a reset |
| a **separately compiled** reader on the same buffer | an update program whose tensor output merely echoes what it just computed rather than reflecting stored state |
| the same reader program bound to a **different** buffer, required to return zeros | the reader obtaining the right answer by any route other than the shared buffer |
| two states advanced in alternation, each against its own recurrence | a single process-wide state behind the inout port, which every single-state check would pass |
| disjoint halves updated in separate steps | an update that replaces rather than accumulates, or that lands at the wrong offset |
| an inout port name the program does not declare | a retain call that returns non-nil without resolving anything |

The refusal has a positive polarity in the same place: after `no_such_state` is
refused, `kv` is retained again on the same operation and must still succeed, so
the refusal is evidence about the name rather than about the call having stopped
working.

## Why the comparison is exact

The state is fp16 and the update is an fp16 add, so an accumulation would
normally drift and need a tolerance. A tolerance is a weaker check, and here it
is avoidable: every value is a multiple of 0.5 with magnitude at most 2, and at
most 8 steps are allowed, so no partial sum passes 16. Multiples of 0.5 in that
range are exactly representable in binary16, and a correctly rounded add of two
exactly representable values with a representable sum is exact.

So the whole recurrence is exact and any disagreement at all is a defect.
`-steps` outside 2..8 is refused rather than silently switching to a weaker
comparison.

## A control that fired

`recurrenceArm` refuses to continue if the accumulated state comes out all
zeros, because the fresh-buffer control immediately below it requires a reader
on an untouched buffer to return zeros — which would pass against a completely
broken implementation if the real answer were zeros too.

It fired on the first run. The step generator was `(7t + 3i) mod 9 - 4`, whose
values sum to exactly zero over six steps for every `i`. The generator is now
`(3t + 5i) mod 7 - 2`, which is asymmetric about zero.

## Limits

This is **not** a KV cache. A cache is position-indexed: token *t* writes row
*t* and leaves rows 0..*t*-1 alone. Nothing here indexes by position. The
disjoint-halves control shows that updates to separate regions do not disturb
each other, which is necessary for a cache and not sufficient. A per-layer,
position-indexed cache proved against a full-prefix oracle is separate, larger
work.

Rebinding an inout port to a different buffer object after the operation is
encoded is **UNMEASURED**. Each program here binds its state once at open, and
sharing is arranged by binding several programs to the same buffer rather than
by rebinding one.

Whether the 64-byte channel stride is a property of this hardware generation,
this macOS version, or the E5RT ABI is **UNMEASURED**. It was observed on one
machine across five shapes. Callers should compute it with the rule above rather
than hard-coding a size, and `-layout` is there to re-check it.

The device mask is a permission, not a placement, so the `main_ane` directory
the compiler leaves is the only local evidence the state lives on the engine.
Every program here fails to open if that directory is absent.
