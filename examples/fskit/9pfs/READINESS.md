# 9pfs Readiness

This example mounts a 9P server through macOS FSKit. The extension uses a Swift
`@main` ExtensionFoundation shell with Go-owned filesystem operations linked as
a c-archive (see `README.md`).

## Status

Working:

- Classic 9P2000 client operations against a disposable live server.
- 9P2000.L client operations against a disposable `p9ufs` server.
- FSKit callback wiring through the in-memory smoke path.
- Swift entrypoint shell and Go operation attachment.
- Installed FSKit mounts, including read, write, rename, chmod, mtime, truncate,
  symlink, hardlink, xattr, and remove.
- Special-file attribute mapping for FIFO, character device, block device, and
  socket nodes.

Not implemented:

- Device node creation.
- Advisory locking.
- Authentication beyond local or anonymous attach defaults.

See `FEATURES.md` for the full operation matrix.

## Local verification

```sh
./verify-local.sh
```

This checks shell/plist lint, `go test`, the in-memory FSKit smoke path, the
Swift entrypoint probe, classic and `.L` live 9P operations, and the default
bundle assembly.

## Installed verification

The installed gate requires a signed `NinePFSHost.app` in `/Applications`, the
FSKit extension enabled in System Settings, and no active 9pfs mounts.

```sh
./preflight-installed.sh
./test-installed-live-9p2000l.sh "$HOME/9pfs-mnt-$(date +%s)"
```

Expected final line:

```text
9pfs: installed FSKit mount read/write/rename/chmod/mtime/truncate/link/xattr/remove ok
```

The installed scripts refuse to run while another 9pfs mount is active. Override
only for an explicit probe window:

```sh
NINEPFS_ALLOW_ACTIVE_MOUNTS=yes ./test-installed-live-9p2000l.sh
```

For a read-only demonstration of an already-active mount:

```sh
./show-live-mount.sh
```

## Research

An experimental no-cgo, no-Swift-`@main` entrypoint and its Swift-metadata
synthesis tooling live in `../9pfs-research/`. That path does not achieve a
strict pure-Go FSKit entrypoint; see `../9pfs-research/RESEARCH.md`.
