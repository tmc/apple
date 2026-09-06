# JACCL parity matrix

This matrix tracks behavioral parity with the standalone JACCL library in
`/Users/tmc/mlx/mlx/distributed/jaccl/lib` at
`d142de6a200eb2f8577aa828834c410056819901`. It does not claim C++ ABI or
implementation equivalence.

| Upstream behavior | `x/jaccl` status | Evidence |
| --- | --- | --- |
| Rank and size | implemented | `Group.Rank`, `Group.Size` tests |
| Send and receive bytes | implemented | point-to-point broker and strict-UC tests |
| All-gather, rank-major output | implemented for mesh and ring; ring stripes all configured wires | group, concurrent reconstruction, strict-UC ring tests, and matched two-host Thunderbolt mesh and ring receipts |
| Sum, min, and max | implemented for all upstream dtype encodings | live standalone reduction oracle, source-derived ring reduce-scatter order tests, strict-UC 14-dtype matrix, and matched two-host Thunderbolt receipts |
| Bool, signed/unsigned integers | implemented | live upstream safe-boundary reduction corpus and local fixtures |
| Float16 and BFloat16 | implemented with native arm64 JACCL byte encodings | live upstream edge corpus covers subnormals, signed zero, infinities, overflow, and quiet/signaling-NaN operand handling |
| Float32, Float64, Complex64 | implemented | live upstream edge corpus covers signed zero, infinities, quiet/signaling-NaN Sum handling, a non-associative three-rank Float32 sum, and complex lexical ordering |
| Mesh topology | implemented | `Mesh`, strict-UC, source-derived large-reduction dispatch tests, and a two-host Thunderbolt receipt |
| Ring topology | implemented for one to four wires per direction | topology, wire-identity, stripe, concurrent reconstruction, strict-UC ring tests, and a two-host Thunderbolt receipt |
| Device connectivity matrix | implemented for setup | meshes select the upstream first directed device; rings preserve every directed wire; the two-host receipt used `rdma_en2` to `rdma_en3` |
| Environment configuration | implemented | `ConfigFromEnv`, `OpenEnv`, and environment alias tests |
| JACCL C++ ABI and MLX `GroupImpl` integration | out of scope | this package is pure Go and does not link libjaccl or MLX |

The acceptance claim requires a differential behavioral result for every row
marked implemented, using matched ranks, topology, dtype, operation, bytes,
and payload hashes. That live JACCL-versus-Go fabric result has not yet been
recorded. A green local test suite is not hardware evidence.

On 2026-09-01, a configured Thunderbolt `/30` route between `rdma_en2` on the
local host (`192.168.0.1`, GID index 1) and `rdma_en3` on `tmc2`
(`192.168.0.2`, GID index 2) passed the full two-rank mesh and ring corpus.
Both used binary SHA-256
`98fda69f429ff37facd8df8dd8854ad5263682f6f9a1bc6af15f2e661d306659` and
the all-dtype receipt
`69f2838d5dda21b003a0199aadd7a85100316b646c1ec64c51f394c6485d8f95`.

The same standalone upstream C++ fixture did not produce a comparable fabric
receipt. An instrumented local diagnostic proved its initial TCP rank and
destination-vector exchanges complete. During native initialization, an
upstream QP-transition error begins destruction and `tbt_destroy_qp` then
blocks uninterruptibly in the Thunderbolt RDMA provider before C++ can report
the original exception. Its RTR path selects an IPv4-mapped destination GID
but hard-codes source GID index 1; that index is link-local on both hosts,
whereas the remote Thunderbolt route is GID index 2. The live result is
therefore a Go two-host hardware pass and an upstream-versus-Go hardware
comparison still `UNMEASURED`, not a differential parity claim.

For the standalone arm, build `testdata/upstream_hardware_oracle.cpp` outside
the repository with `cmake -S x/jaccl/testdata -B ~/tmp/jaccl-upstream-hardware
-DJACCL_SOURCE=/Users/tmc/mlx/mlx/distributed/jaccl/lib`, then run every rank
through `TestUpstreamHardwareEvidence`. Run the Go arm separately with
`TestHardwareEvidence` using the same device matrix, `JACCL_HARDWARE_SIZE`,
topology, payloads, and coordinator. Both receipts must retain their binary
hashes and matching SHA-256 values; these independently initialized
implementations are not claimed to interoperate on one RDMA wire protocol.
In addition to the point-to-point and all-gather receipts, each arm emits one
digest for 14 dtype encodings times Sum, Max, and Min using exact small values.

Hardware runs require two distinct hosts. Never use a loopback coordinator or
the local `devices-local.json` fixture: the test harness rejects both before
opening RDMA because that setup can wedge AppleThunderboltRDMA. Use an external
60-second watchdog, then inspect `ps` for related processes in state `U` before
another run. Prefer `mac2` for experiments when practical.
