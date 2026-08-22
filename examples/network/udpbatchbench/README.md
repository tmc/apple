# udpbatchbench

Measures the Darwin batched UDP datapath (`sendmsg_x`/`recvmsg_x` via
`x/udpbatch`) against the one-syscall-per-datagram path every Go program on
macOS uses today.

## Sweep mode (single machine, loopback)

```
go run . -sweep
```

Measured 2026-08-11, macOS 26.x, arm64 (M-series), go1.26.3, 1200-byte
payloads, 2s per row:

| batch | send pps | vs batch=1 | fill/call |
|------:|---------:|-----------:|----------:|
| 1     | 318k     | 1.00x      | 1.0       |
| 8     | 434k     | 1.36x      | 8.0       |
| 16    | 453k     | 1.42x      | 16.0      |
| 32    | 474k     | 1.49x      | 32.0      |
| 64    | 466k     | 1.46x      | 64.0      |
| 128   | 499k     | 1.57x      | 127.9     |

Read the table with its caveats:

- **Loopback is not the wire.** This measures the syscall boundary and
  nothing else (see `design/high-performance-darwin.md`). The curve's shape —
  flattening around batch 32–64 at roughly 1.5x — is the derivation the
  design docs asked for in place of Rust's uncited hardcoded 32, but the
  magnitude on a real interface must be measured across two machines with
  the split modes below.
- **The recv pps column (printed by the tool) is drop-bound in this rig.**
  The sender floods faster than the receiver drains no matter the batch
  size, so received pps reflects loopback drop behavior, not receive-path
  cost. What the fill/call column *does* prove: the kernel delivers
  completely full batches under load — `recvmsg_x` batching works.
- Send-side errors on loopback overrun (ENOBUFS) are skipped, not counted.

## Split modes (two machines — the wire numbers)

```
udpbatchbench -recv -addr :9999
udpbatchbench -send -addr other-host:9999 -batch 32
```

Both modes print live pps once per second and report (via `x/sockbuf`) any
silent clamp the kernel applied to the requested socket buffers.

## Observed kernel behavior recorded while building this

Truncation is silent on macOS 26.x: the header promises MSG_TRUNC when a
datagram exceeds the buffer, but the kernel clamps `msg_datalen` and sets no
flag. Size receive payloads above the largest expected datagram. Details in
the `x/udpbatch` package documentation.
