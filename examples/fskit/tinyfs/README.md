# TinyFS FSKit extension

This example is a minimal macOS FSKit file system extension. It builds a host app
and an embedded `TinyFSExtension.appex`.

The file system is intentionally small: it probes as usable, loads one unary
volume, exposes an empty read-only root directory, and returns POSIX errors for
mutating operations. It is meant to show the real FSKit extension shape without
the size of a passthrough file system.

Build without signing to check the project structure:

```sh
cd examples/fskit/tinyfs
xcodegen generate
xcodebuild -project TinyFS.xcodeproj -scheme TinyFSHost -configuration Debug CODE_SIGNING_ALLOWED=NO build
```

To use it as a mountable extension, build and sign it with an Apple Developer
identity that has the `com.apple.developer.fskit.fsmodule` entitlement and a
matching provisioning profile. Install the app in `/Applications`, then verify
FSKit sees the module:

```sh
/Applications/TinyFSHost.app/Contents/MacOS/TinyFSHost --fskit-probe
```

Enable it in System Settings under:

```text
General > Login Items & Extensions > Extensions > By Category > File System Extensions
```

The per-app detail sheet can show the same row, but on current macOS builds it
may route through the generic Login Items controller and fail with EPERM. The
category view uses Apple's FSKit module-management controller.

Create a small block device and mount the empty read-only file system:

```sh
mkfile -n 32m /tmp/tinyfs-dummy.img
hdiutil attach -imagekey diskimage-class=CRawDiskImage -nomount /tmp/tinyfs-dummy.img
mkdir -p /tmp/tinyfs-mount
mount -F -t TinyFS /dev/diskN /tmp/tinyfs-mount
```

Replace `/dev/diskN` with the disk printed by `hdiutil attach`.

For the pure Go filesystem-runtime slice, see `../tinyfsgo`. The mountable
extension here still uses Swift for the public ExtensionFoundation `@main`
entrypoint.
