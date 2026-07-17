# rdmaperf

`rdmaperf` compares TCP over ordinary network links, TCP over Thunderbolt
Bridge, and RDMA userspace readiness.

Run a server on the receiving Mac:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf serve -listen 169.254.x.y:9000
```

Run clients from the other Mac:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf tcp -addr 169.254.x.y:9000 -pattern stream -size 1M -duration 30s
GOWORK=off go run ./examples/rdma/rdmaperf tcp -addr 169.254.x.y:9000 -pattern pingpong -size 64 -duration 30s
GOWORK=off go run ./examples/rdma/rdmaperf sweep -addr 169.254.x.y:9000 -pattern stream -duration 10s
```

Use `interfaces` to find local addresses:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf interfaces
```

Use the same commands against Wi-Fi, Ethernet, and Thunderbolt Bridge addresses.
The TCP tests measure the chosen IP path; they do not prove RDMA datapath use.

Check RDMA userspace readiness separately:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf rdma-probe -timeout 10s -json
```

The `rdma-rc-capability` command is a separate, one-shot provider capability
experiment. It opens one device, allocates one PD and CQ, creates exactly one
RC QP, and immediately destroys it. It does not query a port, transition a QP,
register memory, or post work. Both `-allow-rc-probe` and
`CONFIRM_RDMA_RC_CAPABILITY=one-shot-qp-create` are required. Do not run it
without a human-approved plan; a successful creation means only that RC QP
creation is supported, not that RDMA READ or WRITE is supported.

The `rdma-rkey-capability` command is a separate, one-shot provider capability
experiment. It opens one device, allocates one PD, and registers exactly one
4096-byte buffer with local-write and remote-read/write access flags. It reports
the local address, lkey, and rkey, then deregisters the buffer. It creates no
QP, does not query a port, and does not post work. Both `-allow-rkey-probe` and
`CONFIRM_RDMA_RKEY_CAPABILITY=one-shot-mr-register` are required. Do not run it
without a human-approved plan; a zero or nonzero rkey reports only that
registration's key value, not whether a one-sided operation will succeed.

Measured on `rdma_en3` on 2026-07-17: the guarded RC create returned a nil QP
with Darwin `EOPNOTSUPP` (errno 102; `strerror` reports "Operation not
supported on socket"), and the guarded memory registration returned `rkey=0x0`
after requesting remote-read/write access. These are observations for that
device, configuration, and run. The zero rkey is not a recorded rejection of a
one-sided operation; neither READ nor WRITE has been attempted.

True RDMA datapath benchmarking requires successful protection-domain,
completion-queue, memory-region, queue-pair lifecycle, QP transition, and work
completion polling. When provider calls block, the probe watchdog exits 124
instead of leaving the command hung.

The `rdma-pingpong` command drives QP `INIT->RTR`. On the Apple Thunderbolt
RDMA provider, repeated failed RTR attempts have been observed to tear down the
kernel transmit path and wedge the port until reboot. The command therefore
requires `-allow-rtr`. Treat each run as one bounded experiment, not as a retry
loop.

AppleThunderboltRDMA also appears to have tight resource limits. Community
reports against JACCL/exo describe practical ceilings around dozens of
protection domains and about one hundred memory regions per device, with
resource exhaustion sometimes requiring reboot. `rdma-pingpong` opens one
context, one protection domain, one completion queue, one memory region, and
one queue pair per role; do not wrap it in a retry loop or launch many
instances in parallel.

Long-lived applications should not leave successful QPs idle for long periods.
EXO/JACCL reports describe idle connections degrading after tens of minutes.
This tool runs its TCP post-RTS barrier and immediately starts traffic; code
that keeps QPs open should add an application-level heartbeat or tear down idle
QPs.

To test the RDMA datapath, run `rdma-pingpong` on two Macs connected by
Thunderbolt. TCP is used only to exchange LID/QPN/PSN/GID setup data; measured
traffic uses RDMA UC SEND/RECV.

By default `rdma-pingpong` chooses the source GID by preferring IPv4-mapped
GIDs, then Apple Thunderbolt GID index 1, then the first non-zero GID on
non-Thunderbolt link layers. On Apple Thunderbolt, auto-selection never uses
index 0, even when index 0 is non-zero or IPv4-mapped. This matches the JACCL
Thunderbolt fallback and avoids the observed index-0 route that can fail at
`INIT->RTR` with `errno 60`. Use `-gid-index 0` only for an explicit diagnostic
run.

Server:

```sh
perl -e 'alarm shift; exec @ARGV' 30 env GOWORK=off \
  go run ./examples/rdma/rdmaperf rdma-pingpong \
  -listen 169.254.x.y:19100 -size 64 -iters 10000 \
  -setup-timeout 12s -allow-rtr -json
```

Client:

```sh
perl -e 'alarm shift; exec @ARGV' 30 env GOWORK=off \
  go run ./examples/rdma/rdmaperf rdma-pingpong \
  -addr 169.254.x.y:19100 -size 64 -iters 10000 \
  -setup-timeout 12s -allow-rtr -json
```

Save JSON, stderr, and exit status from both roles. If QP setup fails after
resources open, the output includes the local and remote LID, QPN, PSN, GID
index, and GID needed for the next QP transition fix. If local resource setup
times out first, the JSON reports the setup timeout instead.

`-setup-timeout` lets the command write a structured error when a provider call
blocks, but it does not recover a wedged provider goroutine. Keep the outer
timeout wrapper while testing ports or machines that may wedge in the RDMA
provider. The `perl` wrapper above is available on stock macOS.
If the remote side is launched by non-login SSH, ensure `PATH` includes the Go
toolchain before the script runs its preflight.

If both roles reach QP `INIT->RTR` and fail with `modify qp RTR: errno 60`,
save both JSON files, stderr, exit status, and recent AppleThunderboltRDMA logs,
then stop. Do not retry the same topology in a loop; repeated failed RTR
attempts can wedge the provider even when one bounded attempt cleans up.

Omit `-name` and `-device` after Thunderbolt cable or port changes so
`rdma-pingpong` can auto-select a `PORT_ACTIVE` RDMA device. Use explicit
selection only after confirming the RDMA device's port is active. With
`-setup-timeout`, auto-selection probes each candidate in a child process so a
wedged port query does not block the parent command. A wedged provider may
still leave the isolated child process behind until the kernel releases it.
On Apple Thunderbolt RDMA, a nonzero GID at index 0 is not auto-selected; use
`-gid-index 0` only for a deliberate one-shot diagnostic.

If RDMA setup starts failing with protection-domain allocation errors, or if a
provider probe leaves uninterruptible child processes behind, reboot the
affected node before another performance run. After reboot, confirm RDMA is
still enabled and verify the selected `rdma_en*` port is `PORT_ACTIVE`. Prefer
assigning explicit IP addresses to the underlying per-port Thunderbolt
interfaces over destroying `bridge0`, especially on headless machines or setups
that also use USB Ethernet adapters.
