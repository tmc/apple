# signpostweb

An HTTP server that emits one os_signpost interval per request, making a Go
server legible to Instruments and the unified logging system — with names
that decode. No cgo source is involved, but decoding does depend on the
build mode; see the matrix below.

## Run

    go generate   # writes names_darwin.syso (requires clang)
    go build
    ./signpostweb -load 20 &
    curl localhost:8077/work

`-load N` self-generates N requests/second across `/work` and `/flaky`, so
a recording is busy with no external load generator. `/flaky` emits nested
`parse`/`db`/`render` stage intervals, and 1 in 5 of its requests hits a
~100ms "query" marked by a `slow-query` point event — the latency tail to
look for in the trace.

Watch intervals live:

    log stream --signpost --predicate 'subsystem == "com.github.tmc.apple.signpostweb"'

Live stream, not `log show`: signposts are not persisted to the log store,
so `log show --last ...` finds nothing even while a stream of the same
predicate is busy (verified 2026-08-11).

Or record a trace and open it in Instruments:

    xctrace record --template Blank --instrument os_signpost --launch -- ./signpostweb

## How the names decode

The logging system records signpost names by reference: an offset into the
emitting image's `__TEXT,__oslogstring` section, resolved from the binary at
decode time. Go source cannot place strings in that section, so a name passed
from the Go heap pairs correctly but decodes as `<missing name>`.

`go generate` runs `signpostnames` (x/signpost/cmd/signpostnames), which scans
this package for the literal names passed to `handle` and `stage` (declared
with `-funcs handle=1,stage=0`, like go vet's printf `-funcs`; direct
literals passed to the signpost methods are pooled automatically),
assembling them into
`__oslogstring`, and writes `names_darwin.syso`; go build links `.syso`
files automatically, and `x/signpost` finds the strings at runtime and
passes in-image pointers. That is why each route is instrumented with a
string literal: a name built at run time cannot be pooled, and would emit
and pair but decode as missing.

`names_darwin.syso` is a generated binary and is not checked in; run
`go generate` after cloning.

## Build-mode matrix (measured)

| build | names decode? |
|---|---|
| `go build` (CGO_ENABLED=1, macOS default) | yes — the `.syso` forces external linking, which lays out `__oslogstring` |
| `CGO_ENABLED=0 go build` | **no** — the internal linker drops the section; intervals still emit and pair, and the process prints a one-time warning |

For builds that must be `CGO_ENABLED=0`, generate a runtime-loadable pool
instead and load it at startup:

    go run github.com/tmc/apple/x/signpost/cmd/signpostnames -dylib -funcs handle=1,stage=0
    // in main: signpost.LoadNames("names_darwin.dylib")

The dylib path decodes in every build mode, at the cost of shipping the
dylib beside the binary.
