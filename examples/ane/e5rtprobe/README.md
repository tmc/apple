# e5rtprobe

`e5rtprobe` reports what `x/ane/e5rt` — the binding of Espresso's `e5rt_*`
direct-dispatch route to the Neural Engine — can be shown to do.

The tool has the shape it has because of where the package started: every
symbol resolved with `dlsym`, that was all that had been verified, the argument
lists came from Bryngelson, arXiv:2606.22283 chapter 6, and no function had
ever been called.

That has since changed — the route runs end to end and every wrapper the
package exports has been called — but the tool's discipline still holds,
because the hazard does. A call made with a wrong signature faults in C, where
the fault is fatal and `recover` cannot see it, and E5RT reports at least one
failure by throwing a C++ exception that a `purego` caller cannot catch either,
so the wrapper returns success and the process dies later somewhere else.

So the tool never reports success it did not observe, and it puts every call
into a child process regardless of how well established the signature is.

## Commands

`symbols` is the default and calls no `e5rt_*` function:

```sh
go run ./examples/ane/e5rtprobe
go run ./examples/ane/e5rtprobe symbols -json
```

It resolves every name in `e5rt.Symbols`, and compares them against the
`_e5rt_*` spellings in the SDK's `Espresso.tbd` stub. The stub's `symbols:`
list mixes functions and data with no marker, so its count is an upper bound
on `e5rt_*` functions and the reported fraction is not a function-coverage
figure; the tool says so every time it prints one.

`canary` runs the two controls:

```sh
go run ./examples/ane/e5rtprobe canary
```

The first jumps a child to an invalid address, to show the parent reports a
fault as CRASHED. The second runs the buffer readback probe against an address
nothing backs, to show that probe can fail. A crash report from a harness not
shown to detect crashes, and a passing probe not shown to be able to fail, are
both worth nothing.

`stages` calls the unverified wrappers. Each sequence runs in its own child and
streams one record per step it reaches, so a fault localizes to a named step
and the parent survives to print it. Both acknowledgements are required:

```sh
CONFIRM_E5RT_UNVERIFIED_CALLS=call-guessed-signatures \
  go run ./examples/ane/e5rtprobe stages -allow-unverified-calls
```

The two controls run first, so every report carries its own proof that it could
have reported a failure.

## Outcomes

No step is reported as OK.

| Outcome | Meaning |
| --- | --- |
| RESOLVED / MISSING | what `dlsym` found — the one verified operation |
| RETURNED | control came back, with the status code. Not evidence the argument list was right |
| ERROR | the wrapper reported an error before or from the call |
| CRASHED | the child died at this step |
| UNREACHED | not attempted. Never a success |

## What a run shows today

On macOS 26.x all 39 names in `e5rt.Symbols` resolve. The two things a run
establishes beyond that:

- `e5rt_buffer_object_get_data_ptr`, whose two-argument shape the package flags
  as a guess taken from a comment rather than a call, returns a pointer backed
  by writable memory of the requested size. The readback control shows the probe
  would have faulted otherwise. This bounds the address, not the buffer: the
  size came from the `alloc` request, not from the runtime.
- `e5rt_e5_compiler_create_with_config` returns status 1 for a null config —
  the value the source paper passes. Building a config first with
  `e5rt_e5_compiler_config_options_create` and passing it makes the same
  constructor return 0. So the status is a missing argument, not a broken entry
  point, and the compile phase is not blocked where it first appears to be.

When this tool was written, that second result needed `dlsym` directly:
`Lib.Sym` is a lookup over what `Open` resolved, not a live `dlsym`, so it
reached the names in `e5rt.Symbols` and nothing else — and
`e5rt_e5_compiler_config_options_create` was not one of them. Both gaps have
since closed. That symbol is now listed and typed, and `Lib.Lookup` resolves
names outside the list.

Bind and dispatch stay UNREACHED **from this tool**, which reports resolution
and does not drive the route. They are no longer unreached from the package:
see [`e5rtdispatch`](../e5rtdispatch), which compiles, binds, and dispatches a
MIL program on the Neural Engine and checks the result against a CPU
reference.

## Relation to ANEForge

The source paper's author publishes a reference implementation, ANEForge, whose
`aneforge/_lib/e5rt_api.h` and `ane_e5rt_dispatch.mm` are working C that drives
this route. It settles most of what the paper left open, including the argument
order the package's TODOs are blocked on — every `e5rt_*` constructor that takes
no owning object puts its out-parameter first, and methods on an existing object
take that object first and the out-parameter last. Note that the header's own
prose says the opposite ("out goes LAST"); its per-symbol comments and typedefs,
which record what was actually observed, are the ones to follow.

That was strong evidence rather than a verified fact when this tool was the only
thing calling into the package. It has since been verified: `e5rtdispatch`
drives the route end to end with these argument orders and gets correct results
out of the Neural Engine, which exercises the out-first rule at four
constructors and the object-first rule at two methods.

One ANEForge claim did NOT survive contact: it reports that resetting a
never-executed execution stream is rejected, and on macOS 26.x that call
succeeds. See `Lib.ExecutionStreamReset`.
