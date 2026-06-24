# 9pfs

`9pfs` is a Go FSKit example that mounts a real 9P server through macOS FSKit.
The filesystem operations are implemented in Go; a small Swift `@main` shell
supplies the ExtensionFoundation entrypoint. It is a separate Go module so it
can depend on 9P clients without changing the repository root module.

It speaks both classic 9P2000 (via `9fans.net/go`) and 9P2000.L (via
`github.com/hugelgupf/p9`).

## Architecture

Three files implement the extension:

  - `backend.go` — the `backend` interface and its 9P implementations
    (`ninePBackend` for 9P2000, `p9LBackend` for 9P2000.L). One concern: turn a
    9P connection into stat/read/write/create/remove/rename/setattr/xattr calls.
  - `fskit_bridge.go` — implements the `x/fskitbridge` volume interfaces on
    top of the backend. The shared `fskitbridge.Server` owns the FSKit side:
    class registration, operation selectors, item identity, reply blocks, and
    errno reporting.
  - `extension.go` / `cshared.go` — startup: create the bridge server against
    the Swift-provided `NinePFileSystem` class and route its exported entry
    points to the server.

The Swift entrypoint is nine lines (`appex/NinePFSExtension.swift`): a
`UnaryFileSystemExtension` whose `fileSystem` is the `NinePFileSystem` class.
The Go side builds as a c-archive that exports `NinePFSInit` and
`NinePFSConfigureFileSystem`; the Swift executable links it and calls them
before `UnaryFileSystemExtension.main()`.

## Simplest path to a mount

```sh
# 1. Build the extension, the .fs mount-helper bundle, and the host app.
./build-appex.sh /tmp/9pfs-build

# 2. Sign and install (see "Installing" below), then enable the extension in
#    System Settings > General > Login Items & Extensions > File System Extensions.

# 3. Mount a 9P2000.L server through FSKit.
/sbin/mount -F -t 9pfs 'ninep://127.0.0.1:5640?dialect=9p2000l' /path/to/mountpoint
```

`build-appex.sh` produces:

```text
/tmp/9pfs-build/NinePFSExtension.appex
/tmp/9pfs-build/9pfs.fs
/tmp/9pfs-build/NinePFSHost.app
```

Bundle structure:

```text
NinePFSExtension.appex/Contents/
  Info.plist                       NSExtension keys + principal class
  MacOS/NinePFSExtension           Swift @main + Go c-archive
  embedded.provisionprofile        (when signed)
9pfs.fs/Contents/
  Info.plist
  Resources/mount_9pfs             helper for plain `mount -t 9pfs`
NinePFSHost.app/Contents/
  MacOS/NinePFSHost                enables/lists the extension
  Extensions/NinePFSExtension.appex
```

## Mount URLs

The Go bridge parses the resource URL in `loadResource`:

```text
ninep://host[:port][/aname][?dialect=9p2000]
tcp://host[:port][/aname][?dialect=9p2000l]
unix:///path/to/socket?dialect=9p2000l
```

The default TCP port is 5640; the default dialect is classic 9P2000. Use
`ninep://` for installed FSKit mounts (`9p://` is accepted by the Go parser but
is not a valid FSKit resource scheme).

## CLI: verify the 9P side

The same command exercises the 9P client directly, without FSKit:

```sh
GOWORK=off go run . -dialect 9p2000 -net tcp -addr 127.0.0.1:5640 -ls /
GOWORK=off go run . -dialect 9p2000 -addr 127.0.0.1:5640 -cat /README
GOWORK=off go run . -dialect 9p2000l -addr 127.0.0.1:5640 -ls /
```

Use `-aname` when the server exports a named tree.

## Local verification

```sh
./verify-local.sh
```

This runs shell/plist lint, `go test`, the in-memory FSKit smoke path
(`-fskit-smoke`), the Swift entrypoint probe, classic and `.L` live 9P checks,
and the default bundle assembly.

The in-memory smoke path checks the FSKit callback wiring without a live server:

```sh
GOWORK=off go run . -fskit-smoke
```

The live checks start disposable servers and exercise the real client:

```sh
./test-live-9p2000.sh     # classic 9P2000 via 9fans.net/go
./test-live-9p2000l.sh    # 9P2000.L via github.com/hugelgupf/p9
```

These scripts patch a temporary copy of `github.com/hugelgupf/p9@v0.4.0` because
that release's `p9ufs` localfs has a Darwin compile issue and its client xattr
methods return `ENOSYS`; the patch maps chmod/mtime to `os.Chmod`/`os.Chtimes`
so the FSKit path is exercised rather than stopping at the test server. The
patch is local to the build and does not change module dependencies.

See `FEATURES.md` for the supported/unsupported operation matrix and
`READINESS.md` for verification state.

## Installing

For an installed test, sign the bundles and install the host app:

```sh
CODESIGN_IDENTITY='Apple Development: Your Name (TEAMID)' \
NINEPFS_REQUIRE_PROFILES=yes \
./build-appex.sh /tmp/9pfs-build
./verify-signed-build.sh /tmp/9pfs-build
./install-local.sh /tmp/9pfs-build      # prints commands; copies only with CONFIRM_9PFS_INSTALL=yes
```

Signing needs development provisioning profiles whose application identifiers
match `dev.tmc.apple.examples.fskit.9pfs` and
`dev.tmc.apple.examples.fskit.9pfs.extension`; the extension profile must grant
`com.apple.developer.fskit.fsmodule`. The script auto-discovers matching
profiles from `~/Library/MobileDevice/Provisioning Profiles`.

`NinePFSHost.app` must be copied to `/Applications` and enabled through System
Settings. Direct FSKit mounts use `/sbin/mount -F -t 9pfs` and do not need the
`.fs` bundle; install it under `/Library/Filesystems/9pfs.fs` only for plain
`mount -t 9pfs`.

## Installed mount gate

```sh
./preflight-installed.sh
./test-installed-live-9p2000l.sh "$HOME/9pfs-mnt-$(date +%s)"
```

The gate starts a disposable 9P2000.L server, mounts it with
`/sbin/mount -F -t 9pfs`, and verifies mounted directory listing, read, write,
rename, truncate, chmod, mtime, symlink, hardlink, xattr, and remove. It
refuses to run while another 9pfs mount is active unless
`NINEPFS_ALLOW_ACTIVE_MOUNTS=yes` is set.

To inspect an already-mounted volume read-only:

```sh
./show-live-mount.sh /path/to/active/9pfs-mount
```

## Notes

macOS ships `/sbin/mount_9p`, but it only accepts `mount_9p [-r] fs_tag`, where
the tag is an `AppleVirtIO9P` IORegistry device property — useful for
VM-provided virtio 9p, not arbitrary 9p servers, and it does not exercise this
bridge.

An experimental no-cgo / no-Swift-`@main` entrypoint lives in
`../9pfs-research/`. It is research scaffolding and is not part of this example;
the default Swift `@main` path above is the supported way to build and mount.
