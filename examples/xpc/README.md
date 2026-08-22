# XPC examples

Go equivalents of Apple's XPC service templates, using the `xpc` package.

- `addservice` — the service. Mirrors the low-level template's `main.c`.
- `addclient` — the caller. Mirrors the client sketch in that template's
  trailing comment.
- `machdemo` — the same add service over a Mach service, which needs no
  application bundle.
- `build-bundle.sh` — assembles the host application the first two need.

`addservice` and `addclient` use the template's service name and message keys,
so the Go service is a drop-in replacement inside the same `.xpc` bundle:

    dev.tmc.sample-xpc-service-ll
    {"firstNumber": int64, "secondNumber": int64} -> {"result": int64}

## Mapping from the C template

| C | Go |
| --- | --- |
| `xpc_listener_create(name, NULL, 0, handler, &err)` | `xpc.NewServiceListener(name, xpc.ListenerOptions{}, accept)` |
| `xpc_listener_activate` | `(*xpc.Listener).Activate` |
| `xpc_session_set_incoming_message_handler` | returned by `IncomingSessionRequest.Accept` |
| `xpc_dictionary_get_int64(msg, "firstNumber")` | `ReceivedMessage.Decode(&req)` with an `xpc:"firstNumber"` tag |
| `xpc_dictionary_create_reply` + `xpc_session_send_message` | return a value from the handler |
| `xpc_session_create_xpc_service` | `xpc.DialXPCService` |
| `xpc_session_send_message_with_reply_sync` | `(*xpc.Session).Call` with a context that cannot end |
| `xpc_session_send_message_with_reply_async` | `(*xpc.Session).Call` with a cancellable context |
| `xpc_session_send_message` (no reply) | `(*xpc.Session).Notify` |
| `xpc_session_cancel` | `(*xpc.Session).Cancel` |
| `dispatch_main()` | `select {}` |

## Activation

A listener or session is live as soon as it is created. `Activate` exists for
the `Inactive: true` case, which lets a handler be installed before the first
message arrives. Calling it on an already-active object is API misuse: XPC
traps the process rather than returning an error, so the program dies with
`SIGTRAP` inside `xpc_listener_activate` and no Go error to catch.

Replies are a return value rather than a second call. A handler that returns a
value replies with it, and one that returns an error replies with
`{"error": ...}`.

A handler that returns `nil, nil` sends no reply, but that does not mean the
peer waits. If the message asked for a reply, the peer's reply handler fires
promptly with `Underlying connection interrupted` and the session is dead
afterwards:

    silent op: timedOut=false reply callback fired with error: Underlying connection interrupted

This is XPC's behaviour, not the binding's. Swift's own handler returning a nil
`Output` does exactly the same thing to a Go peer. Measured five times in each
direction between a Go and a Swift peer, ten runs identical, against warmed
services so no probe hit a starting instance; see
`examples/xpc/swiftinterop/`. Do not file it as a Go bug.

The practical consequence is that "no reply" is only usable for messages the
sender never expected a reply to. It is not a way to ignore a request.

Message bodies can be `xpc.Dictionary` or any struct; the codec reads `xpc`
tags and falls back to `json` tags. `addclient -raw` sends the dictionary form
to show both.

## Running

An XPC service is launched on demand by launchd out of a `.xpc` bundle in
`Contents/XPCServices` of a host application, and only a process inside that
application can look it up. `build-bundle.sh` puts `addservice` in the bundle
and `addclient` in `Contents/MacOS`, so the client *is* the host application:

    $ ./examples/xpc/build-bundle.sh ~/tmp/AddDemo.app
    /Users/tmc/tmp/AddDemo.app
    $ ~/tmp/AddDemo.app/Contents/MacOS/AddDemo -first 23 -second 19
    42
    $ ~/tmp/AddDemo.app/Contents/MacOS/AddDemo -raw -first 100 -second 5
    105

No code signing is needed. The ad-hoc signature the Go linker writes draws an
`AMFI: … has no CMS blob?` line in the unified log, but launchd spawns the
service anyway. The app does not have to be launched through LaunchServices
either; running its executable directly is enough.

Run outside the bundle, both fail promptly rather than hanging. Observed on
macOS 26:

    $ go run ./examples/xpc/addservice
    create listener: Unable to activate listener: Connection init failed at
    listener activation with error 1 - Operation not permitted

    $ go run ./examples/xpc/addclient
    dial dev.tmc.sample-xpc-service-ll: Underlying connection was invalidated.
    Reason: Connection init failed at lookup with error 3 - No such process

The service fails inside `NewServiceListener`, not at `Activate`, even though
the message names activation: a process that launchd did not start as this
service may not create the listener at all. Both messages come from
`xpc.RichError`, which carries the underlying XPC description.

A service that dies during startup is reported to the client as `send:
Underlying connection interrupted`, with no detail about the cause. launchd
gives the service no terminal, so diagnosing it means logging to a file or
reading `log show --predicate 'processID == <pid>'`.

## Running machdemo

    go build -o ~/tmp/machdemo ./examples/xpc/machdemo
    ~/tmp/machdemo -install        # writes ~/Library/LaunchAgents and bootstraps it
    ~/tmp/machdemo -call 23 19     # 42
    ~/tmp/machdemo -uninstall

launchd starts the service on demand. The name must appear in the agent's
`MachServices` dictionary, and the listener needs `ForceMach`, or creation
fails. The agent sends the service's output to a log under `$TMPDIR`, since a
launchd job has no terminal and its errors are otherwise invisible.

### Target queues

`-target-queue` configures a private dispatch queue on both sides. The listener
passes it in `ListenerOptions`; the client creates an inactive session, calls
`SetTargetQueue`, and then activates it:

    ~/tmp/machdemo -uninstall
    ~/tmp/machdemo -install -target-queue
    ~/tmp/machdemo -call -target-queue 23 19
    ~/tmp/machdemo -uninstall

This demonstrates configuration and lifecycle. The `xpclive` test suite also
asserts inside the service that listener callbacks execute on the selected
queue.

### Peer requirements

`-require-peer-entitlement NAME` installs an entitlement-exists requirement on
the listener or session. Requirements are owned values: constructors return a
`*xpc.PeerRequirement`, the caller closes its reference, and XPC retains the
requirement when it is installed. Installation is inactive-only, so the client
creates an inactive session, calls `SetPeerRequirement`, and then activates it.

This option is intentionally not a zero-setup command. XPC evaluates the code
signature of the remote process, so each binary must be signed with the named
entitlement. The following uses `get-task-allow`, a development entitlement
that macOS permits in an ad-hoc signature:

    cat >~/tmp/machdemo-entitlements.plist <<'EOF'
    <?xml version="1.0" encoding="UTF-8"?>
    <plist version="1.0"><dict>
      <key>com.apple.security.get-task-allow</key><true/>
    </dict></plist>
    EOF
    go build -o ~/tmp/machdemo-service ./examples/xpc/machdemo
    go build -o ~/tmp/machdemo-client ./examples/xpc/machdemo
    codesign --force --sign - --entitlements ~/tmp/machdemo-entitlements.plist ~/tmp/machdemo-service
    codesign --force --sign - --entitlements ~/tmp/machdemo-entitlements.plist ~/tmp/machdemo-client
    ~/tmp/machdemo-service -install -require-peer-entitlement com.apple.security.get-task-allow
    ~/tmp/machdemo-client -call -require-peer-entitlement com.apple.security.get-task-allow 23 19
    ~/tmp/machdemo-service -uninstall

Signing only one binary is a useful rejection control: the other side refuses
it. Arbitrary application entitlements require an appropriate signing identity
and provisioning; substituting an invented name into an ad-hoc signature is
not a valid test. The same API also provides same-team, platform-binary,
entitlement-value, and lightweight-code-requirement constructors; those require
identities or LWCR data specific to the application and are therefore not
fabricated here. Peer-requirement constructors require macOS 26; on an older
runtime they return an availability error.

## What crosses the boundary

Measured against a real Swift `XPCListener` and `XPCSession` in both
directions on macOS 26.6.1 (25G76), Swift 6.3.3, SDK 26.5. The harness is
`examples/xpc/swiftinterop/`; every value below was observed, not inferred.
Swift built each value with the raw C constructors, which is the widest
surface a sender has.

| Swift sent | Go received | Observed value |
| --- | --- | --- |
| `bool` | `bool` | `true` |
| `int64` | `int64` | `-42` |
| `uint64` | `uint64` | `9223372036854775808` (not collapsed to int64) |
| `double` | `float64` | `3.5` |
| `string` | `string` | `"héllo"` |
| `data` | `[]byte` | `4 bytes deadbeef` |
| `null` | `nil` | key present with a nil value |
| `array` | `[]any` | `[int64 1, string "two", bool false]` |
| `dictionary` | `xpc.Dictionary` | `nested int64 7`, recursive |
| `endpoint` | `xpc.Endpoint` | `handle=0x102b83820`, and re-sendable |
| `date` | **`string`** | `"<date: 0x102b83770> Tue Nov 14 14:13:20 2023 PST (approx)"` |
| `uuid` | **`string`** | `"<uuid: 0x102b837c0> 01234567-89AB-CDEF-FEDC-BA9876543210"` |
| `fd` | **`string`** | `"<fd: 0x102b82a50> { type = (invalid descriptor), path = ... }"` |
| `shmem` | **`string`** | `"<shmem: 0x102b83450>: 16384 bytes (1 page)"` |

The last four rows are one mechanism: `rawObjectToValue` switches on
`xpc_get_type` and any type with no case becomes its `xpc_copy_description`
string. There is no error and no flag, and the result decodes into a `string`
struct field without complaint.

Going the other way, Go can build only the first nine rows, and all nine
arrive on the Swift side with the correct XPC type. Go cannot build a date,
uuid, fd, or shmem at all, so those have no reverse direction to measure.

The date string is local-time text stamped "(approx)", not the nanosecond
value that was sent; the exact instant is not recoverable. The uuid string
does contain the UUID in readable form, but only by parsing a description
format Apple never promised.

### Descriptors and shared memory

Neither can be sent or received through the high-level API. Sending fails with
`xpc: unsupported dictionary value type`, because there is no `xpc` value type
to hand it; receiving yields the description string above rather than a
descriptor or a mapping.

Do not go looking for the missing function. `raw_xpc_fd_create`,
`raw_xpc_fd_dup`, `raw_xpc_dictionary_set_fd`, `raw_xpc_dictionary_dup_fd`,
`raw_xpc_shmem_create`, and `raw_xpc_shmem_map` all exist in
`xpc/xpc.raw.gen.go` already — they are unexported, so nothing outside package
`xpc` can reach them. This is a missing high-level surface over machinery that
is already generated, not missing machinery.

### Endpoints relay through Go intact

`xpc_type_endpoint` has a case on the way in and `Endpoint` has a case on the
way out, so Go can receive an endpoint, hold it, and put it back into an
outgoing dictionary. The relayed copy is still live. In `swiftinterop`, Go
fetched a Swift anonymous listener's endpoint, handed it straight back, and
Swift built an `XPCSession` from the round-tripped copy and got a reply
through it:

    Go holds an xpc.Endpoint (handle=0x105353690) and has no API to dial it
    relayed endpoint is live; reply through it:
    pong	string	"reached through a relayed endpoint"

So Go works as an endpoint courier: it can hand endpoints between two other
processes. The asymmetry is that it can be neither end of that handoff. It
cannot produce an endpoint from its own listener and cannot dial one it holds,
which is what the omissions below record.

One limitation, noted from the source and **not tested at runtime**: the
`Endpoint` encoding case is on `setDictionaryValue` only. Array elements go
through `scalarToRawObject`, which has no `Endpoint` case, so an endpoint
inside a slice looks like it should fail to encode. Confirm before relying on
either behaviour.

## Not covered

The high-level template uses `NSXPCConnection`, `NSXPCListener`, and a
protocol vended through `NSXPCInterface`. That is Foundation's
proxy-based API, not this package; the bindings live in `foundation`. The
`xpc` package covers the `XPCSession`/`XPCListener` API that replaced it.

Anonymous listeners cannot be reached from Go today: exporting a listener's
endpoint and constructing a session from one are both recorded as omissions in
`xpc/xpc.omissions.gen.go`, so these examples use a named service.
