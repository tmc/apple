# bundlereuse

Compiles a MIL program to a Neural Engine bundle once, then reopens that bundle
— in this process and in a second process, with the source MIL and its weights
deleted — and checks that every arm matches a float64 CPU reference.

Compilation is the expensive part of the private `e5rt` route; binding and
dispatch are microseconds. `e5rt.OpenBundle` is the interface for paying it
once, and nothing outside the package's own tests called it before this example.

```sh
go run ./examples/ane/bundlereuse
go run ./examples/ane/bundlereuse -inch 128 -outch 128 -spatial 32
```

```
bundle reuse: inch=32 outch=32 spatial=8

  compiled in 120ms
  backend: the compiler emitted [ane]
  freshly compiled vs float64 CPU reference: max diff 0.000431 (tolerance 0.05)

  source MIL and weights deleted
  reopened in 5.573ms, 22x faster than compiling
  reopened vs freshly compiled: identical
  mutation control: negating the input moved the output by 0.555
  second process vs freshly compiled: identical

  refused, as it must be: a function name the bundle does not contain
    retain function "no_such_function": e5rt: e5rt_program_library_retain_program_function: status 1
  refused, as it must be: a bundle with a truncated file
    open program bundle: e5rt: e5rt_program_library_create: status 1

OK
```

## Why the source is deleted

A program that reopened a bundle and silently recompiled from source would look
identical to one that reused it — same answer, same call from the caller's side.
Deleting the model directory between compiling and reopening removes that
reading: anything that still runs cannot have recompiled, because there is
nothing left to compile from.

## What each check would catch

| check | what it would catch |
| --- | --- |
| every arm against a float64 CPU evaluation | an arm that loaded a different or stale program |
| reopened arms against the freshly compiled arm, requiring *exact* agreement | a subtly different program; same fp16 code on same hardware should not merely agree closely |
| negating the input | a bundle returning a constant, or replaying the previous output buffer |
| a function name not in the bundle | an open that succeeds without resolving anything |
| a bundle with its largest file truncated | an open that does not validate what it loaded |

The two refusals fail at different points — `retain_program_function` and
`program_library_create` — which is what makes them two controls rather than one
repeated.

## Scope of the cross-process claim

The second process shares this one's code-signing identity and has it as a
parent. That is the case in which E5RT bundle reuse has been observed on macOS
26.x. Reuse after an `aned` restart, after a reboot, or from an unrelated
process is **UNMEASURED**, and this example does not claim it.

The timing ratio is a single measurement on one small program, not a benchmark.
It is printed because it is the number that motivates the interface, not as a
performance claim.
