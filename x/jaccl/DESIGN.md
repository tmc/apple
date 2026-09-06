# Pure-Go JACCL-like RDMA collectives

Status: proposed

This document specifies `github.com/tmc/apple/x/jaccl`, a Go collective
communication package built on `github.com/tmc/apple/rdma`. It is an
independent implementation inspired by JACCL's operational model. It neither
links `libjaccl` nor calls the MLX-C JACCL C++ wrapper. It uses no cgo; the
existing `rdma` package supplies the purego-backed Darwin provider calls.

The first supported platform is `darwin/arm64` with Apple's Thunderbolt RDMA
provider. Other platforms fail closed with `ErrUnsupported`; there is no silent
TCP collective fallback.

## Goals

- Provide a small Go API for a communicator, point-to-point bytes,
  `Barrier`, `AllGather`, and `AllReduce`.
- Establish an RDMA connection with an explicit, versioned Go TCP control
  plane.
- Make unreliable-connected (UC) queue-pair flow control correct before adding
  throughput optimizations.
- Keep provider handles, registered memory, and work-request descriptors alive
  for exactly the duration required by the verbs calls.
- Make every operation cancelable and group teardown safe with operations in
  flight.

## Non-goals

- Source or ABI compatibility with `libjaccl`.
- MLX arrays, Metal buffers, GPU-direct memory, CUDA, NCCL, remote RDMA read,
  remote write, or atomics.
- Reliable-connected (RC) queue pairs, adaptive topology selection, or a
  performance claim in the initial release.
- A public environment-variable API. An optional adapter may be added after
  the explicit API is stable.

## Package boundary and public API

The implementation lives in `x/jaccl`. `rdma` remains the ABI binding layer;
`x/rdma` remains the home for generic Apple route and resource policy. No
JACCL-specific state belongs in either package.

The first API is deliberately byte-oriented:

```go
package jaccl

type Config struct {
	Rank        int
	Size        int
	GroupID     string
	Coordinator string
	Device      string
	Devices     [][][]string
	PreferRing  bool
	Port        uint8
	Topology    Topology
}

type Topology interface {
	Peers(rank int) []int
}

type DType uint8
type ReduceOp uint8

const (
	Bool DType = iota
	Int8
	Int16
	Int32
	Int64
	UInt8
	UInt16
	UInt32
	UInt64
	Float16
	BFloat16
	Float32
	Float64
	Complex64
)


const (
	Sum ReduceOp = iota + 1
	Min
	Max
)

type Group struct { /* unexported */ }

func Available() bool
func ConfigFromEnv() (Config, error)
func Open(context.Context, Config) (*Group, error)
func OpenEnv(context.Context) (*Group, error)
func (g *Group) Rank() int
func (g *Group) Size() int
func (g *Group) Close() error
func (g *Group) Send(context.Context, dst int, src []byte) error
func (g *Group) Recv(context.Context, src int, dst []byte) error
func (g *Group) Barrier(context.Context) error
func (g *Group) AllGather(context.Context, dst, src []byte) error
func (g *Group) AllReduce(context.Context, dst, src []byte, DType, ReduceOp) error
```

`Config` is valid only when `0 <= Rank < Size`, `Size >= 2`, `GroupID` is
nonempty, `Port` is nonzero, and `Topology` supplies no duplicate, self, or
out-of-range peers. `Device` is either an exact RDMA device name or empty to
request the single available safe device. When `Devices` is present, it is the
JACCL directed connectivity matrix: meshes use the first device per directed
peer and rings retain one to four wires in each direction. `Open` validates
reciprocal links through the control plane before creating any QP.

The communicator is safe for concurrent calls but executes collective and
point-to-point operations one at a time per group. Each admitted operation has
a monotonically increasing epoch and all peers must enter the same operation,
with the same shape, in the same epoch. A mismatch is a protocol error that
poisons the group. This is simpler and more honest than advertising unordered
collectives before an ordering protocol exists.

`dst` and `src` may be exactly the same slice for in-place `AllReduce`; partial
overlap is rejected. `AllGather` requires `len(dst) == Size*len(src)` and emits
rank-major segments. `AllReduce` requires equal source and destination lengths,
length divisible by the selected dtype width, and supports all JACCL dtypes
from Bool through Complex64. The checked-in upstream oracle verifies the
reduction byte encodings used by the live standalone arm64 build.

Typed wrappers are a later package-level convenience, not a second transport
API. They must call this byte API after proving slice representation and dtype
mapping.

## Errors and availability

The package exports sentinels suitable for `errors.Is`:

- `ErrUnsupported`: the platform or RDMA provider is unavailable.
- `ErrInvalidConfig`: static configuration is invalid.
- `ErrProtocol`: a control-plane peer, epoch, shape, or destination is invalid.
- `ErrClosed`: the group has started closing or is closed.
- `ErrPoisoned`: a previous RDMA or protocol failure made further operations
  unsafe.

All returned errors add operation, rank, peer, epoch, and (when applicable)
slot and completion identifiers. A context cancellation reports the context
error wrapped with operation context. A failed provider call or non-success
work completion poisons the group: automatic recovery of a UC queue pair is
out of scope.

## Control plane

The control plane is ordinary Go TCP, used only for setup, shape agreement,
credits, and failure propagation. It is not a payload fallback.

`Open` has these stages:

1. Validate `Config`, topology reciprocity, group ID, and a protocol version.
2. Each rank dials or accepts according to deterministic rank ordering. The
   coordinator distributes a membership record; all ranks compare a hash of
   rank, size, group ID, topology, device choice, and protocol version.
3. Each directed link exchanges a fixed-size queue-pair destination record:
   rank, peer, wire number, QPN, PSN, LID, selected local GID index, GID,
   active MTU, link layer, and a random group incarnation.
4. Each rank validates all records, creates its resources, transitions its QPs
   through INIT, RTR, and RTS, then performs a bidirectional control-plane
   readiness acknowledgement.
5. A failed stage tells every reachable peer to abort. `Open` then tears down
   only resources it created and returns an error; it never returns a partial
   group.

Every frame has a magic value, version, kind, group ID, incarnation, source
rank, destination rank, epoch, and bounded length. The decoder rejects an
oversize frame before allocation and rejects version or identity mismatches.
The maximum frame size is a constant covered by tests. The coordinator is a
rendezvous service, not a trusted source of unvalidated route data.

Each collective uses a matching operation agreement before and after its
payload phase. The second agreement prevents a fast rank from closing its
point-to-point credit sessions while a peer is completing that same epoch.

## RDMA setup and ownership

The Darwin backend is private to `x/jaccl` and adapts only exported
`apple/rdma` calls. Its production chain is:

```text
device context -> protection domain -> completion queue -> UC queue pair
                                             \-> send and receive memory regions
```

For each directed peer wire it creates one UC QP, one CQ, and separately
registered fixed send and receive buffers. It requests one SGE per work
request and signal-all sends. Startup builds INIT attributes locally, uses
`x/rdma.RTRAttr` for route construction, then moves the QP to RTS. It must
query port attributes and a bounded GID prefix, choose a GID with
`x/rdma.SelectRouteGID`, and retain the exact selected index in the exchanged
destination.

Thunderbolt route selection must not silently choose GID index zero. The normal
automatic policy uses the existing `x/rdma` policy (prefer nonzero IPv4-mapped
GIDs, then index one). If it has no safe route, `Open` fails. A diagnostic
override, if ever added, is separate from `Config`, explicitly named unsafe,
and cannot be the default.

Page-backed buffers registered as memory regions are fields of the link object.
The Apple provider rejects Go-heap registrations, so each link owns anonymous
`mmap` send and receive buffers and unmaps them after deregistration. Every
verbs call receiving a Go pointer uses the typed helper and keeps the referent
alive through the call. Work-request, SGE, bad-WR, and completion arrays are
long-lived link fields, not per-operation stack values. This is both a lifetime
requirement and a prerequisite for later allocation measurements.

## UC flow-control invariant

The target reference uses UC, not a reliable queue pair. A local send
completion only proves that the local provider accepted or emitted a send; it
does not prove that a remote receive was armed. Consequently, a bounded slot
ring alone is not flow control.

Each link has `depth` send slots and `depth` receive slots. A slot is owned by
exactly one state at a time:

```text
recv: idle -> posted -> completed -> consumed -> posted
send: idle -> staged -> posted -> completed -> idle
```

The required invariant is:

> A receive slot is not reposted, and its credit is not returned to the peer,
> until the operation has consumed the completed bytes from that slot.

At operation start the receiver posts all receive slots before the sender may
post a payload. It advertises one credit only after each receive post succeeds.
The sender posts at most one payload per received credit. On a receive
completion, the collective copies or reduces the bytes while the slot remains
owned by that completion; it then reposts the slot and returns one credit. A
send completion releases only its local send slot. Credits and payloads carry
the group incarnation, operation epoch, direction, slot, generation, length,
and dtype/op tag as applicable.

This MVP uses operation-scoped slot arming. Keeping receives armed across
operation boundaries is a later optimization and needs an additional header,
epoch isolation, and the same consume-before-repost proof. It is not an
implementation shortcut for the first release.

The CQ poller accepts only the exact outstanding work ID, opcode, byte length,
and success status. A stale or duplicate completion, unexpected opcode, short
completion, or non-success status poisons the link and group. Work IDs include
the operation generation so a delayed completion cannot alias a reused slot.

## Algorithms

Milestone algorithms favor a small correctness surface:

- `Send` and `Recv`: a mesh uses one directed-link byte transfer. A ring
  stripes bytes across its directed wires with JACCL's ceiling-based ranges;
  the two-rank case biases send left and receive right. Exact lengths are
  agreed on the control plane before receivers post buffers.
- `Barrier`: a coordinator-mediated control-plane barrier in the first
  release. It establishes operation agreement; it is not counted as RDMA
  payload performance.
- `AllGather`: direct pairwise exchange on a full mesh. A ring circulates
  rank-major blocks clockwise and stripes each block across its configured
  wires.
- `AllReduce`: a full mesh exchanges then reduces in ascending global rank
  order, matching the standalone mesh implementation even for non-associative
  floating-point sums. Above standalone JACCL's dtype-specific mesh thresholds,
  it instead uses the same internal one-wire ring reduce-scatter. A ring
  gathers rank-major blocks, then applies the standalone ring reduce-scatter's
  chunk-specific, right-associated cyclic order: the
  one-direction path is counter-clockwise, while the two-direction path uses
  counter-clockwise order for the first per-chunk region and clockwise order
  for the remainder.

Topology is part of correctness. `AllGather` and `AllReduce` fail rather than
silently route through undeclared peers. The supported sparse topology is a
reciprocal ring; arbitrary graph routing remains unsupported.

## Shutdown

`Close` is idempotent and has a defined order:

1. Mark the group closing, reject new calls, and signal every active poller and
   control-plane read.
2. Wait for every admitted operation to leave the group. Poll loops must observe
   both their caller context and the close signal.
3. Close reachable control-plane connections, destroy QPs directly, deregister
   MRs, destroy CQs, deallocate PDs, then close device contexts.
4. Join all teardown errors and leave every handle cleared.

The default teardown does not first force QPs to ERR. The historical Apple
provider reference observed teardown diagnostics after an explicit ERR
transition; direct QP destruction is the specified default until an isolated
provider test proves another sequence necessary. No memory region is
deregistered while a QP may still access it.

## Tests and evidence gates

The package ships a private `verbs` interface and deterministic fake provider.
The production adapter is the only implementation allowed to import
`apple/rdma`. Unit tests cover config, framed control traffic, route selection,
QP transition arguments, resource order, work-ID decoding, completion matching,
dtype arithmetic, aliasing, and every public error path.

The fake transport has two modes:

- `reliable`: a convenience transport that may buffer unmatched sends.
- `strictUC`: an adjudicating transport that drops a send with no posted receive
  and records drops, slot reuse before consumption, stale completions, and
  credit violations.

The following are mandatory, non-vacuous gates:

1. The shipping driver completes 2- and 3-rank byte transfer, gather, and every
   supported reduction under randomized peer skew in `strictUC`, with zero
   drops, zero premature reposts, and exact results.
2. A frozen credit-free driver is run against `strictUC` and is required to
   lose or stall for a schedule with more chunks than slot depth. If it stops
   failing, the adjudicator is defective, not the implementation validated.
3. A frozen repost-before-consume driver is required to corrupt data under the
   fake transport. The shipping driver must remain clean under the same
   accounting.
4. A stale-completion injection must be rejected by generation matching. An
   unexpected completion status and short completion must poison the group.
5. `Close` while a receive waits must return, wait for the operation, and prove
   the dependency order QP, MR, CQ, PD, context without a provider.
6. Race-enabled tests cover concurrent `Close`, cancellation, and operations;
   no test skip or `[no tests to run]` is acceptance evidence.

Hardware tests are opt-in and require separately started ranks with one shared
size/topology, an exact binary/source receipt, selected device/port/GID/MTU
records, the destination exchange transcript, completion/error records,
payload hashes, and teardown output. They report `UNMEASURED` when the fabric
is absent. The full 14-dtype corpus has a one-minute deadline to include both
the admission and completion agreements for every operation. These tests do
not turn a green fake transport suite into a hardware or performance claim.

Performance work starts only after the correctness gates pass and captures the
loaded `apple` revision, executable hash, `MLX_LIB_PATH`-like override state if
present, provider/device identity, CPU host state, and raw samples. It compares
interleaved arms and keeps setup, payload time, and teardown diagnostics
separate. No benchmark gate may use a loopback transport to claim RDMA poller
allocation behavior.

## Delivery slices

| Slice | Deliverable | Acceptance |
| --- | --- | --- |
| 0 | `x/jaccl` package docs, errors, config, topology, fake verbs | Unit config and protocol validation are exhaustive and cross-platform builds return `ErrUnsupported`. |
| 1 | Darwin device discovery, safe GID selection, resource chain, destination exchange | Fake transition/order tests pass; real hardware remains opt-in and reports a receipt or `UNMEASURED`. |
| 2 | UC slot machine and `Send`/`Recv` | Strict-UC shipping and two frozen negative-control drivers prove the invariant. |
| 3 | Group admission/close, barrier, two-rank `AllGather` and `AllReduce` | Cancellation, stale completion, poison, in-place, and close tests pass. |
| 4 | Full-mesh multi-rank collectives | Three-rank strict-UC randomized-skew matrix passes with exact values. |
| 5 | Ring algorithms and persistent cross-operation receives | A separate design amendment defines headers, epoch isolation, and a new consume-before-repost proof. |
| 6 | Hardware and performance qualification | Retained two-rank receipts and provenance-qualified measurements; absence is reported as `UNMEASURED`. |

No later slice changes the earlier public semantics without a separately
reviewed API proposal.

## Evidence inventory, 2026-08-31

This proposal was derived from a read-only scan of the Apple repository's
branches and worktrees and `/Users/tmc/ml-explore/mlx-c*` worktrees.

| Location/ref | Head inspected | Relevant finding |
| --- | --- | --- |
| `apple-rdma` worktree / `apple-rdma` | `3bb655576` | Generated verbs expose provider calls, QP ERR/RESET helpers, CQ polling, and `x/rdma` route/resource policy. |
| Apple `integration/jaccl-call3-main` | `50a7164b` | Preserved the RDMA/purego Call3 integration line. |
| Apple `upstream-tag` / `rdma-path-mtu-cap` | `f27477d7` / `440a3b18e` | Record the named-port-attribute and negotiated-MTU lineage. |
| MLX-C `jaccl/per-slot-invariant` | `cbfc2b59` | Native Go reference and tests state the per-slot safety rule. |
| MLX-C `codegen/jaccl-purego` | `7bfb0ff9` | Separate generated purego bindings are for the JACCL C API, not a replacement for a Go transport. |
| MLX-C `codegen/jaccl-c-api` | `f7e64542` | Optional C API and generated `mlx/c/jaccl.*` surface. |
| MLX-C `explore-pureffi-jacclc` | `02b9de66` | Existing FFI experiment; its worktree had unrelated dirty generated/module files and was not modified. |
| MLX-C profile harness | detached `9ca14637` | Historical wire-benchmark worktree with uncommitted benchmark edits; treated as evidence only. |

The MLX-C reference documents that a batch-reap wave loop can starve a UC
receiver and that a send completion is not remote receive proof. It also
contains stricter negative controls and lifetime/teardown tests. Those findings
are carried into this design as requirements, not copied as an implementation
dependency.
