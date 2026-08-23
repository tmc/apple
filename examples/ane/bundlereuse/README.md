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

  compiled in 268ms
  backend: the compiler emitted [ane]
  freshly compiled vs float64 CPU reference: max diff 0.000431 (tolerance 0.05)

  source MIL and weights deleted
  reopened in 3.912ms, 68x faster than compiling
  reopened vs freshly compiled: identical
  mutation control: one opened program, second execution with a negated input moved the output by 0.555 and still matches its own reference
  second process vs freshly compiled: identical over all 256 values
  parser controls: a well-formed result parses, and 6 malformed ones are refused

  refused, as it must be: a function name the bundle does not contain
    retain function "no_such_function": e5rt: e5rt_program_library_retain_program_function: status 1
  refused, as it must be: a bundle with H16C.bundle/H16C.e5 truncated from 3880 bytes to 0
    open program bundle: e5rt: e5rt_program_library_create: status 1

OK
```

## Why the source is deleted

A program that reopened a bundle and silently recompiled from source would look
identical to one that reused it — same answer, same call from the caller's side.
Deleting the model directory between compiling and reopening removes that
reading.

The claim is narrower than "no recompilation happened": it is that neither
`OpenBundle` nor the second process fell back to the caller-side MIL and
weights, because those no longer exist. What `aned` does internally when it
materializes a compiled bundle is not observed here.

## What each check would catch

| check | what it would catch |
| --- | --- |
| every arm against a float64 CPU evaluation | an arm that loaded a different or stale program |
| reopened arms against the freshly compiled arm, requiring *exact* agreement | a subtly different program; same fp16 code on same hardware should not merely agree closely |
| negating the input, through **one** opened program on its second execution, requiring the result to match its own reference | a stale output buffer, a rebinding bug, and a program that reacts but reacts wrongly |
| a declared value count from the second process, re-checked here | a truncated or empty child result compared over its prefix and called identical |
| a function name not in the bundle | an open that succeeds without resolving anything |
| a bundle with a selected retained file truncated, named and sized in the output | an open that does not validate what it loaded |

The two refusals fail at different native entry points —
`retain_program_function` and `program_library_create` — which is what makes them
two controls rather than one repeated. They establish that a wrong function name
and a damaged selected file are refused, not that every file in the bundle is
validated.

The comparator rejects NaN and unequal lengths. It previously did neither, so a
second process that printed one of its 256 values was compared over that one
element and reported **identical** — a false accept on the cross-process claim,
which is the entire point of that arm. An all-NaN result was worse: `NaN > max`
is false, so it reported a maximum difference of zero. Six malformed child
results are now fed to the parser on every run to prove it still refuses them.

The corruption control reports which file it truncated and how large it was
(`H16C.bundle/H16C.e5`, 3880 bytes) rather than assuming the largest file is the
program payload. Bundle selection also fails closed: the compiler nests a
per-hardware `H16C.bundle` inside the digest-named bundle, so "the first
`.bundle` found" was right only by directory-walk ordering.

## Scope of the cross-process claim

The second process shares this one's code-signing identity and has it as a
parent. That is the case in which E5RT bundle reuse has been observed on macOS
26.x. Reuse across an `aned` restart is measured separately, in
`examples/ane/internal/anedrestart`, and **survives**: a bundle compiled under
one daemon instance reopens and computes correctly under the next, with the
compiling instance provably gone.

Reuse after a reboot or from an unrelated process is **UNMEASURED**, and this
example does not claim it.

Bundles are also **not durable**. Two bundles roughly a day old failed to create
an operation at all (`status 13`), while bundles minutes old succeeded in the
same processes. Whatever bounds a bundle's life, it is not the daemon's
lifetime. The boundary is UNMEASURED, so treat reuse as a within-session
optimisation rather than a persistent artifact cache.

The timing ratio is a single measurement on one small program, not a benchmark.
Compile time varies substantially run to run (93–268ms observed), so the ratio
has been seen anywhere from 22x to 68x on the same machine. Treat it as
"reopening is orders of magnitude cheaper", not as a figure.
It is printed because it is the number that motivates the interface, not as a
performance claim.
