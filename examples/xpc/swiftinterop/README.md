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

**Everything the Go codec can build survives in both directions.** bool,
int64, uint64, double, string, data, null, array, and dictionary all arrive
with the same type and value. `uint64(1) << 63` arrives as
`9223372036854775808`, so the signed/unsigned split is preserved rather than
collapsed.

**Five XPC types travel only one way, and lossily.** Swift can put a date,
uuid, fd, or shmem in a message; the Go codec has no case for any of them, so
`rawObjectToValue` falls through to `xpc_copy_description` and Go receives a
Go `string` such as

    date   "<date: 0x102b83770> Tue Nov 14 14:13:20 2023 PST (approx)"
    uuid   "<uuid: 0x102b837c0> 01234567-89AB-CDEF-FEDC-BA9876543210"
    fd     "<fd: 0x102b82a50> { type = (invalid descriptor), path = /private/tmp/swiftinterop-fd-probe.txt }"
    shmem  "<shmem: 0x102b83450>: 16384 bytes (1 page)"

Go cannot build any of them, so the reverse direction cannot even be
attempted. A descriptor and a shared page are *resources*, and a description
of one is not one: this is where the binding stops being usable, not merely
inconvenient.

**Endpoints are the exception, and they round-trip.** `xpc_type_endpoint` has
a case in `rawObjectToValue`, and `Endpoint` has a case on the encoding side,
so Go receives a real `xpc.Endpoint` and can put it back in a dictionary.
`endpointrelay` proves the copy is live: Swift builds an `XPCSession` from the
endpoint that made the round trip through Go and gets a reply through it. Go
is a working endpoint *courier*. What Go cannot do is originate one (no
listener endpoint accessor) or consume one (no session-from-endpoint
initializer); both are recorded as omissions in `xpc/xpc.omissions.gen.go`.

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

See `~/tmp/agent-collab/apple/20260815-xpcinterop-gaps.md` for the full
findings with literal output.

## A note on shmem

An early version of these examples built the shared region with
`posix_memalign`. `xpc_shmem_create` treats malloc memory as API misuse and
`SIGTRAP`s the *sending* process, which reaches the peer as a connection
interruption with no indication of the cause. The region must be `mmap`ed. That
is a Swift-side authoring hazard, not a finding about the Go binding, and it is
recorded here so the next reader does not rediscover it as one.
