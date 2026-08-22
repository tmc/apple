# Swift/Go XPC interoperability

Two XPC services and two clients, one of each in Swift and Go, talking to each
other over Mach services. They exist to answer one question: what survives the
trip between Swift's `XPCListener`/`XPCSession` and this repository's `xpc`
package, which binds the same API.

- `swiftservice.swift` — Swift `XPCListener`; a Go `xpc.Session` calls it.
- `goclient/` — Go `xpc.Session`; calls the Swift service.
- `goservice/` — Go `xpc.Listener`; a Swift `XPCSession` calls it.
- `swiftclient.swift` — Swift `XPCSession`; calls the Go service.
- `run.sh` — builds all four, installs both services as LaunchAgents, runs
  every probe in both directions, and boots the jobs out again.

## Running

    ./examples/xpc/swiftinterop/run.sh

It needs a logged-in GUI session, because it bootstraps LaunchAgents into
`gui/$(id -u)`. Everything it creates is named after the shell's pid, so
concurrent runs cannot collide, and its cleanup trap asserts on the way out
that no job and no plist was left behind. No code signing is needed: neither
side installs a peer requirement, so no entitlement has to be carried.

## The ops

Both services answer the same vocabulary, so each probe can be run in either
direction.

| op | meaning |
| --- | --- |
| `describe` | the sender puts one value of every type it can build in the message; the receiver reports what it decoded |
| `typezoo` | the receiver replies with one value of every type it can build; the sender reports what it decoded |
| `typezoo:KEY` | as `typezoo`, but only `KEY` beyond the base scalars — how a single value is isolated |
| `describe:KEY` | as `describe`, but only `KEY` beyond the base scalars |
| `silent` | the handler declines to reply (Go `nil, nil`; Swift `nil`) |
| `fail` | the Go handler returns an error, which the binding turns into `{"error": string}` |
| `errorkey` | a legitimate payload that happens to contain an `"error"` key |
| `endpointrelay` | Go fetches the Swift listener's endpoint, hands it straight back, and Swift dials the copy |

## What it shows

Measured on macOS 26.6.1 (25G76), Swift 6.3.3, SDK 26.5.

**Every XPC value type round-trips in both directions.** bool, int64, uint64,
double, string, data, null, array, dictionary, date, and uuid all arrive with
the same type and value whichever side sent them. `uint64(1) << 63` arrives as
`9223372036854775808`, so the signed/unsigned split is preserved rather than
collapsed.

**Descriptors and shared pages round-trip as resources, not descriptions.**
A `fd` decodes to an `*xpc.FileDescriptor` and a `shmem` to an
`*xpc.SharedMemory`, and each probe proves the result is usable by reading
through it rather than by printing its type. Both directions carry a known
payload:

    Go   -> Swift   fd     dup'd to fd 3, contents="GO FD PAYLOAD\n"
                    shmem  16384 bytes mapped, prefix="GO SHMEM PAYLOAD"
    Swift -> Go     fd     *xpc.FileDescriptor   dup=3 contents="FD PAYLOAD\n"
                    shmem  *xpc.SharedMemory     16384 bytes mapped, prefix="SHMEM PAYLOAD"

Reading the bytes is the point. A descriptor that duplicates is not yet a
descriptor that works, and the number alone cannot tell the two apart.

Both are handles rather than values, because that is what XPC makes them.
`xpc_fd_dup` returns a *different* descriptor on every call, so there is no
one descriptor for the Go value to be; `FileDescriptor.Dup` performs the
duplication and the caller closes the result. `SharedMemory.Map` places the
region in this process and the caller closes the `Mapping`. Both retain the
boxed object on decode, so they outlive the handler that received them, and
both have an idempotent `Close`.

**Endpoints round-trip too, but Go is only a courier.** `endpointrelay`
proves the copy is live: Swift builds an `XPCSession` from the endpoint that
made the round trip through Go and gets a reply through it. What Go cannot do
is originate one (no listener endpoint accessor) or consume one (no
session-from-endpoint initializer); both are recorded as omissions in
`xpc/xpc.omissions.gen.go`.

**Anything still unmodelled decodes to `xpc.Unsupported`.** That type exists
so decoding is never silently lossy: an earlier version of the codec returned
the description *string* for anything it did not model, which made a file
descriptor arrive as ordinary Go data with no error and nothing to test.
`Unsupported` is that distinction, and encoding one is an error rather than a
re-send of the description.

Which types those are is UNMEASURED here. `xpc_type_connection` is the
obvious candidate and the fallthrough in `rawObjectToValue` would catch it,
but `typeZoo` does not put a connection in the message, so this harness has
never exercised the path and cannot report on it. An earlier version of this
file listed connection among the types the zoo sends; it never did.

**Declining to reply cancels the connection, on both sides.** A handler that
returns no reply to a message that asked for one is observed by the peer as
`Underlying connection interrupted`, not as silence. This is symmetric —
Swift returning `nil` and Go returning `nil, nil` behave identically, 5 runs
each — so it is XPC's behaviour, not the binding's. The `xpc` package README's
"one that returns `nil, nil` sends no reply" is true of the local process and
misleading about the peer.

**The `{"error": ...}` convention is invisible on the wire.** A Go handler
returning an error and a Go handler returning `Dictionary{"error": ...}` as
real data produce dictionaries a Swift client cannot tell apart; `errorkey`
and `fail` are the same probe run twice to show it.

## A note on shmem

An early version of these examples built the shared region with
`posix_memalign`. `xpc_shmem_create` treats malloc memory as API misuse and
`SIGTRAP`s the *sending* process, which reaches the peer as a connection
interruption with no indication of the cause. The region must be `mmap`ed with
`MAP_SHARED`.

This is why `xpc.NewSharedMemory` allocates the region itself and hands back
the mapping, rather than taking a caller's buffer. Go memory and malloc memory
are owned by an allocator rather than by the caller, so an API that accepted
either would be one that traps on correct-looking code. The hazard is real on
the Swift side, where the region is the author's to get right; on the Go side
it is closed by construction.
