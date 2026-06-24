# 9pfs no-Swift entry research

This directory holds the experimental "pure entry" build path for the `../9pfs`
FSKit example. It is research scaffolding, not part of the example. The example
itself ships and mounts with the default Swift `@main` entrypoint documented in
`../9pfs/README.md`; nothing here is required to build, sign, install, or mount
the 9pfs filesystem.

## Goal and current state

The goal was a macOS FSKit extension whose entrypoint is a Go executable with no
Swift `@main` and no cgo. The `pureentry` build links Swift ExtensionFoundation
and FSKit metadata into a Go executable and starts the Go runtime from a delayed
`loadResource` callback, then replaces the FSKit operation methods with Go
implementations.

### Verdict: the strict goal is unreached, but it is unfinished labor, not a wall

(Verified 2026-05-30; the blocker framing was corrected 2026-06-23 — see
"The remaining work" below and "A better path," which makes the labor moot.)

A strict success criterion — *no swiftc, no cgo, and ExtensionFoundation
actually loads and mounts the Go entrypoint* — was tested empirically on this
stack (macOS 26.5 arm64, Swift 6.3.2, Go 1.26.3, FSKit + ExtensionFoundation
present). Results:

- **The swiftc pipeline builds clean.** The full pipeline completes and was
  independently verified: `swiftc` compiles `NinePFSExtension.swift` →
  `dump`/`emit` round-trip proving the Go emitter reproduces the metadata →
  rename `_main`→`_swift_unused_main` → `CGO_ENABLED=0 go build -tags pureentry`
  → `go tool link -linkmode=external` (EF/FSKit/Foundation/swiftCore/libobjc) →
  `patch_lcmain` repoints `LC_MAIN` to `main.nsextMainEntry.abi0` →
  `check_swift_metadata`. Final binary is a valid Mach-O arm64; its `LC_MAIN`
  entry maps exactly to `main.nsextMainEntry.abi0`, which `nsext_arm64.s` JMPs to
  `NSExtensionMain`.
- **It is genuinely no-cgo.** `CGO_ENABLED=0`; the entry glue uses `purego` plus
  `//go:cgo_import_dynamic`, and there is no Swift `@main` in the build (the
  entry is the Go `nsextMainEntry`, not a swiftc-synthesized `main`).
- **"No swiftc" holds only at build time — the metadata is a captured swiftc
  blob, not synthesis.** The `swiftmeta_bundled` path runs with no swiftc,
  `os/exec`, or compiler at build time (`write_swiftmeta_fixture` →
  `emit_macho_object` → `dump_swift_metadata -check` all pass; the extracted
  `sections/` dir is empty after `-compact`). But the bytes it links are a
  byte-for-byte capture of one prior `swiftc` run, stored inline in the
  133 KB `pureentry/swiftmeta_manifest.json` (27 sections, 239 symbols): e.g.
  `__TEXT,__text` is 660 inline `words` replayed verbatim (captured ARM64 method
  bodies), and the conformance descriptors carry relocations into the opaque-type
  descriptor `…UnaryFileSystemExtensionPAAE13configurationQrvpQOMQ`.
  `git log --follow` on the manifest shows it has only ever been *captured*, not
  regenerated.
- **The "actually loads and mounts" leg was never observed.** The build stops at
  static metadata checks; no probe in this tree shows ExtensionFoundation
  reflectively instantiating the principal struct and FSKit completing a mount
  through the swapped-in Go methods.

### The remaining work: hand-authoring ABI structs, not recovering hidden values

The remaining no-swiftc work is the Swift protocol-conformance and reflection
metadata for `ExtensionFoundation.AppExtension` and
`FSKit.UnaryFileSystemExtension` — the `__swift5_proto` / `__swift5_types` /
`__swift5_typeref` / `__swift5_assocty` / `__swift5_capture` records plus their
conformance descriptors and witness tables. `check_swift_metadata.go` requires
exactly these six non-empty `__swift5_*` sections and the four
conformance/witness symbols, because they are what ExtensionFoundation reads to
instantiate the `@main` principal.

An earlier version of this file called these values "resolved by swiftc against
the closed-source `.swiftmodule`" and "not recoverable from the open Swift ABI."
**That was wrong, and verifying it is what reframed this research.** The targets
the conformance descriptors point at are *public, linkable symbols* in the SDK's
TBD stub:

```
$ grep -F 'QOMQ' $(xcrun --show-sdk-path)/System/Library/Frameworks/FSKit.framework/FSKit.tbd
  _$s5FSKit19FileSystemExtensionPAAE13configurationQrvpQOMQ
  _$s5FSKit24UnaryFileSystemExtensionPAAE13configurationQrvpQOMQ   ← the symbol the old verdict called SDK-private
$ grep -F 'UnaryFileSystemExtensionMp' .../FSKit.tbd               ← protocol descriptor, also exported
```

All 47 external FSKit/ExtensionFoundation symbol references in
`swiftmeta_manifest.json` are public mangled symbols. Swift conformance
descriptors use *relative* references whose offsets the **linker** computes at
link time against the symbol's final address — the synthesizer only emits a
relocation to the named public symbol, exactly like any C `extern`. The protocol
*requirements* (the witness-table slot list, in declared order) are spelled out
in the textual `.swiftinterface`, which swiftc itself consumes. Record *shape*
comes from the open Swift ABI (`swift/include/swift/ABI/Metadata.h`,
`GenProto.cpp`). So all three inputs — requirement list, slot layout, and link
targets — are in hand.

The honest blocker is therefore **labor, not secrecy**: hand-authoring the four
conformance descriptors + witness tables and the `__swift5_*` records to the ABI
spec, byte-for-byte correct, and keeping them correct across OS/toolchain
revisions. That is real, fragile work — but it is *unfinished*, not *impossible*.
See "A better path" below: it makes that labor unnecessary.

## A better path: don't synthesize metadata at all

The metadata synthesis above is only required if the goal is *zero swiftc, even
for one fixed stub*. It isn't worth that — and there is a path that achieves the
real goal ("no Swift source and no Go static-linking in a filesystem author's
build") without synthesizing a single ABI byte. This is the recommended
direction.

FSKit discovers the extension's principal from **Swift conformance metadata, not
from a plist class** — `appex/Info.plist` has no `NSExtensionPrincipalClass`, and
`grep -r PrincipalClass` across `FSKit.framework`/`ExtensionFoundation.framework`
finds nothing; discovery is via `EXExtensionPointIdentifier =
com.apple.fskit.fsmodule` plus the `@main` conformance records. So a single
prebuilt `.appex` can't be repointed at a different filesystem by editing a
plist. But that conformance can be **generic**:

- **Stub-A2 (shippable now):** a generic ObjC `FSUnaryFileSystem` subclass whose
  selectors forward to `dlsym`'d C functions, and a 9-line Swift `@main` whose
  `fileSystem` is that generic class. Build the Go side as
  `-buildmode=c-shared` (a `.dylib`) instead of `c-archive`, and `dlopen` it from
  the ObjC `+load`. The native code becomes filesystem-agnostic; the only
  per-filesystem artifact is the Go dylib plus its `FSPersonalities` plist. swiftc
  still runs once for the 9-line conformance, but never compiles filesystem logic.
- **Stub-B (one universal bundle):** promote A2's generic type into a
  once-compiled, pre-signed `GoFSKitHost.appex` whose `fileSystem` getter
  resolves the Go dylib named in its Info.plist at runtime. Reused across many
  filesystems; the conformance is for a fixed generic type compiled normally, so
  **no metadata synthesis is needed**. Caveats: an FSKit extension must be signed
  by an fskit-entitled team, so this is "drop-in within one signing authority,"
  not an open plugin market; and `FSPersonalities` are static in the plist, so
  each filesystem needs a plist templating step (not swiftc).

Either path makes the `__swift5_*` synthesis above unnecessary. Pursue Stub-A2/B
instead of metadata synthesis.

### If you still want to settle the synthesis questions (lower priority)

1. **Settle "actually loads and mounts" for the captured-blob binary.** Install
   the pureentry appex, attempt a real `mount -F -t 9pfs`, and capture `log
   stream --predicate 'subsystem == "com.apple.fskit"'` plus the binary's own
   syslog line `9pfs-nocgo-entry: installed Go file system methods`. If EF never
   instantiates the principal, the strict criterion is unreachable even *with*
   captured metadata; if it mounts, the sole remaining gap is the metadata
   origin. Requires interactive sudo. **Not staged:** the build artifacts under
   `/tmp` (`/tmp/pureentry-build`, `/tmp/orig-build-appex.sh`,
   `/tmp/9pfs-profiles`) have been wiped; this needs a full rebuild + re-extract
   first (the installed Swift extension and signing profiles make it
   *recoverable*, not ready-to-run). See `MOUNT-PROBE-RUNBOOK.md`.
2. **Settle "genuine synthesis."** Hand-author the four conformance descriptors +
   witness tables and the `__swift5_proto`/`__swift5_types` records from the Swift
   ABI spec *without* copying the captured `words`/`byteSpans`, emitting
   relocations to the public FSKit/EF symbols (in `FSKit.tbd`) and reading the
   requirement list from the framework `.swiftinterface`, then diff the emitted
   object's `__swift5_*` against the captured manifest. A byte-divergent object
   that EF still loads would be the first genuine no-swiftc result.

This was driven by an adversarially-verified multi-agent workflow
(`pureentry-research.workflow.js` in this directory); the conclusions above are
the verified findings.

## Contents

- `pureentry/` — standalone `go run` tools (`//go:build ignore`): a Mach-O
  dumper, object emitter, symbol renamer, delayed-load helper encoder, Swift
  metadata fixture writer, and the `NinePFSExtension.swift` witness shell.
- `pureentry/swiftmeta_manifest.json`, `pureentry/delayed_helper_manifest.json`
  — Swift-derived metadata fixtures used by the manifest-only emitter paths.
- `pureentry_darwin.go`, `dynimport_pureentry_darwin.go`, `nsext_arm64.s` — the
  Go entrypoint glue (build tag `pureentry`). These are fragments of the `9pfs`
  `main` package; the `pureentry` build overlays the `9pfs` core source and adds
  these files. They do not compile on their own.
- `pureentry-research.workflow.js` — the multi-agent workflow that drove the
  verification above (re-runnable via the Workflow tooling).

## Rebuilding the experiment

The pureentry build mode was removed from `../9pfs/build-appex.sh` when this
tree was extracted. To resurrect it, overlay these files onto a copy of the
`../9pfs` source and drive the `pureentry/` tools as the prior build script did:
compile the Swift metadata object, extract/re-emit it with
`emit_macho_object.go`, rename `_main`, link a `-tags pureentry` Go build with
`linkmode=external`, and patch `LC_MAIN` with `patch_lcmain.go`. See this file's
git history (`git log --follow`) for the original `build-appex.sh` branches.

Note when overlaying: `../9pfs/main_darwin.go` is guarded `//go:build darwin &&
!pureentry` so the overlay's `pureentry_darwin.go` provides the sole `func main`
under `-tags pureentry`. (An earlier extraction dropped that `!pureentry` guard,
which caused a `main redeclared` build error in the overlay; it is restored.)
