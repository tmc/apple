# pureentry install + mount probe — runbook

The pureentry research reached a verified verdict (`RESEARCH.md`): the strict
no-swiftc/no-cgo/EF-loaded goal is **not reached** — the metadata is a captured
swiftc blob, not synthesis. (Note: the remaining no-swiftc work is hand-authoring
the ABI structs, *not* recovering values from closed metadata — the conformance
targets are public in `FSKit.tbd`; see `RESEARCH.md`. And "A better path" there
makes this whole probe lower-priority: Stub-A2/B reaches the real goal without
synthesis.) One empirical slice remains open: *does the captured-blob pureentry
binary actually load under ExtensionFoundation and mount?* Observing that does
not change the "not reached" verdict on synthesis, but it isolates whether the
only unproven leg (EF reflective instantiation → FSKit probe/load via the
swapped-in Go methods) works at all.

This step is **gated on you**: it needs an interactive admin password (copy into
`/Applications`) and replaces the currently-installed Swift-`@main` extension.
An agent on a timer cannot drive the password prompt, so the steps are here for
you to run when ready.

## State

> **Not staged — recoverable.** As of 2026-06-23 the `/tmp` artifacts this
> runbook references (`/tmp/pureentry-build`, `/tmp/orig-build-appex.sh`,
> `/tmp/9pfs-profiles`) have been **wiped** (volatile `/tmp`). This probe is not
> ready-to-run; it needs a full rebuild + re-extract first. The installed Swift
> extension and the signing profiles make it *recoverable* (see the recovery
> commands below), not staged. Treat the paths below as where artifacts go once
> rebuilt, not as existing files.

As originally captured (2026-05-30):

- Verified pureentry binary (built + LC_MAIN/metadata-checked by the workflow):
  `/tmp/pureentry-build/Contents/MacOS/NinePFSExtension` (8.1 MB, arm64) — **now
  wiped; rebuild required.** A **bare binary**, not yet wrapped in a signed
  `NinePFSHost.app`.
- Codesigning identity present: `Apple Development: travis.cline@gmail.com`.
- Currently installed (working) extension:
  `dev.tmc.apple.examples.fskit.9pfs.extension` (the default Swift `@main` build).
- `build-appex.sh` no longer has a `pureentry` mode (it was extracted). The
  original pureentry build pipeline is preserved in this dir's git history (and
  was copied to `/tmp/orig-build-appex.sh`, now wiped — recover it with the
  `git show` command in the recovery section).

## Caveat before you start

Installing the pureentry bundle **replaces the known-good Swift-`@main`
extension**. To restore the working one afterward, rebuild + reinstall the
default extension (with the same identity + profiles — see below):
```sh
cd ../9pfs
CODESIGN_IDENTITY="Apple Development: Travis Cline (A6VCBGEHSE)" \
  NINEPFS_EXTENSION_PROFILE=/tmp/9pfs-profiles/extension.provisionprofile \
  NINEPFS_HOST_PROFILE=/tmp/9pfs-profiles/host.provisionprofile \
  ./build-appex.sh /tmp/9pfs-appex
CONFIRM_9PFS_INSTALL=yes ./install-local.sh /tmp/9pfs-appex
```

## Signing requirement (why the bundle build needs profiles)

An FSKit extension must be signed with Team `LJ98655CHY` and carry the matching
provisioning profiles, or ExtensionFoundation will refuse to load it — which
would make the probe fail for a *signing* reason and produce a false negative on
the research question. The required profiles are **not** in the standard
`~/Library/MobileDevice/Provisioning Profiles/` dir on this machine; they were
extracted from the installed working extension to `/tmp/9pfs-profiles/`
(`extension.provisionprofile` = `…9pfs.extension`, fskit=true, exp 2027;
`host.provisionprofile` = `…9pfs`, exp 2027). `build-pureentry-appex.sh` defaults
to those. If `/tmp/9pfs-profiles/` is gone, re-extract:
```sh
mkdir -p /tmp/9pfs-profiles
cp /Applications/NinePFSHost.app/Contents/Extensions/NinePFSExtension.appex/Contents/embedded.provisionprofile /tmp/9pfs-profiles/extension.provisionprofile
cp /Applications/NinePFSHost.app/Contents/embedded.provisionprofile /tmp/9pfs-profiles/host.provisionprofile
```

## Steps

1. **Build a signed pureentry host-app bundle.** The bare binary in
   `/tmp/pureentry-build` is not a bundle. Re-run the original pureentry build,
   which assembles + signs `NinePFSHost.app` around the Go-entry binary:

   ```sh
   cd /Volumes/tmc/go/src/github.com/tmc/apple-worktrees-fskit/examples/fskit/9pfs
   # drive the extracted pipeline via the saved original script:
   NINEPFS_GO_PUREENTRY=yes NINEPFS_SWIFTMETA_BUNDLED=yes \
     bash /tmp/orig-build-appex.sh /tmp/9pfs-pureentry-appex
   ```

   (If `/tmp/orig-build-appex.sh` is gone, recover it:
   `git show b37f9648^:examples/fskit/9pfs/build-appex.sh > /tmp/orig-build-appex.sh`.
   Note `../9pfs/main_darwin.go` is now guarded `darwin && !pureentry`, so the
   overlay's `pureentry_darwin.go` is the sole `main` — the duplicate-`main`
   regression is fixed.)

2. **Watch the FSKit + extension logs** (separate terminal, leave running):

   ```sh
   log stream --style compact \
     --predicate 'subsystem == "com.apple.fskit" OR eventMessage CONTAINS "9pfs-nocgo-entry"'
   ```

   The decisive evidence the Go entry ran is the binary's own line:
   `9pfs-nocgo-entry: installed Go file system methods`
   (from `pureentry_darwin.go` after `class_replaceMethod` swaps in the Go ops).

3. **Install (needs admin password):**

   ```sh
   cd /Volumes/tmc/go/src/github.com/tmc/apple-worktrees-fskit/examples/fskit/9pfs
   CONFIRM_9PFS_INSTALL=yes ./install-local.sh /tmp/9pfs-pureentry-appex
   pluginkit -m | grep 9pfs   # confirm registration
   ```

4. **Start a 9P server** the extension can dial (defaults to `127.0.0.1:5640`,
   dialect `9p2000`), then **mount**:

   ```sh
   mkdir -p "$HOME/9pfs-pureentry-mnt"
   /sbin/mount -F -t 9pfs 'ninep://127.0.0.1:5640' "$HOME/9pfs-pureentry-mnt"
   ls -la "$HOME/9pfs-pureentry-mnt"
   ```

   (The existing `./test-installed-live-9p2000l.sh "$HOME/mnt"` spins up a
   9P2000.L server + mounts + exercises the op matrix; it targets whatever
   extension is installed, so it works against the pureentry build too.)

## How to read the result

- **You see `9pfs-nocgo-entry: installed Go file system methods` AND the mount
  lists files** → ExtensionFoundation *does* load the Go entrypoint and FSKit
  drives the swapped-in Go methods. The only remaining gap to the strict goal is
  then provably the metadata *origin* (captured vs. synthesized) — experiment #2
  in `RESEARCH.md`.
- **EF never instantiates the principal / no `9pfs-nocgo-entry` line / mount
  fails at load** → the strict criterion is unreachable even *with* swiftc-
  captured metadata; the project should stop on the synthesis front.

Record whichever happens back into `RESEARCH.md`.
