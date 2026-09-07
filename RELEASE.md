# v0.7.0 preparation

Status: not ready to tag.

The release branch includes the reviewed debugger and binding work, optional
Objective-C checked dispatch, Core ML program and weight construction, affine
weight quantization, and method/protocol attachment for reused FSKit classes.
The Core ML integration retains the newer fp16 corrections and rejects shape
overflow, unrepresentable fp16 scales, and weight paths that escape through a
symlink. Checked dispatch is opt-in and checks type families; it does not prove
integer widths, struct ABI compatibility, or receiver lifetime.

## Open gates

- SDK selection and provenance: the generator's release/beta targeting spec
  selects SDK build `25F70` and explicitly makes mixed-SDK detection (G2) a tag
  blocker. A single SDK resolver, conflicting `SDKROOT` rejection, readable
  per-package SDK identity, and enforcement of one SDK base per release tree
  remain incomplete. Existing input hashes do not replace these checks.
- API availability: the generator has an optional `--target-macos` introduction
  ceiling, but none of this tree's 104 tracked `generate.go` files sets it.
  Unknown introduction metadata is retained. This flag does not define the
  minimum supported runtime.
- Signature comparison: the gate reports 21 unadjudicated findings in nine
  private generated files against baseline `7b66b9767d1b39c5565beaed0c06361b586a8103`.
  They predate this preparation. Review the generator and ownership evidence
  before classifying them as regressions or accepted corrections; do not add
  blanket allowances or repair generated files by hand.
- Anonymous union identity (Class C): the generator loses anonymous-union
  encodings before ownership arbitration and misattributes unrelated fields to
  `virtualization.Union`. The resulting named fields are not a resolution of
  `gtshaderprofiler.GTAGX2ShaderProfilerProgramAddress`. Its `allOpaqueRecords`
  entry is restored: the expected stale-entry failure exposes this unresolved
  cross-framework misattribution. Removing it in `bde66274f` was incorrect.
  Fix Class C in the generator before re-evaluating the entry; do not broaden
  allowances or treat the opaque-record gate as green.
- macOS 27 runtime qualification: `TestConstexprPlacementRuntime` crashes while
  loading its deliberately invalid parameters-as-inputs model on macOS 27.0
  build `26A5425a`. On macOS 26.6.2 build `25G83`, the runtime rejects that model
  normally. The macOS 27 crash remains unresolved.
- Support policy: explicitly select the minimum runtime and qualify it, macOS
  26, and macOS 27 before advertising that range. Include missing symbol,
  class, and selector controls, enum checks, and struct ABI probes. Public SDK
  availability and private framework compatibility need separate evidence.

## Validation scope

Local validation used Darwin arm64, macOS 26.6.2 (`25G83`), SDK `25F70`, and
Go 1.25.0 with `GOWORK=off`. The full module builds. Checked dispatch and the
Objective-C/FSKit bridge tests pass under the race detector. Core ML package
tests pass, including runtime prediction and invalid-model controls. The
signature gate is not green for the reasons above.

On Darwin arm64, macOS 27.0 (`26A5425a`), SDK `26A5419a`, and Go 1.25.0, the
checked-dispatch and bridge tests pass. The full module builds and the Core ML
suite passes with the crashing runtime probe excluded. That exclusion is a
qualification gap, not a passing runtime-compatibility result.

## Preserved work

The macOS 27 private inventory remains a historical snapshot, not a replacement
for the shared binding tree. ANE shared-event and inference experiments retain
HOLD evidence and need separate completion/lifetime qualification. Installed
FSKit experiments, parked next-release tools, and old generated worktrees are
preserved separately. GPU counter slice helpers still need bounds and loader
review before inclusion. No broad hardware qualification, tag, or publication
is part of this preparation.
