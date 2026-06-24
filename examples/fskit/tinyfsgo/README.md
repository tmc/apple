# tinyfsgo

`tinyfsgo` is a pure Go smoke test for implementing the TinyFS filesystem
objects with the repository's FSKit and Objective-C runtime bindings.

It registers Go subclasses of `FSUnaryFileSystem`, `FSVolume`, and `FSItem`,
implements the key FSKit operation selectors, and invokes FSKit reply blocks
directly from Go. It uses no cgo.

Run it with:

```sh
go run ./examples/fskit/tinyfsgo --smoke
```

Expected output:

```text
tinyfsgo: smoke ok
tinyfsgo: registered FSUnaryFileSystem, FSVolume, and FSItem subclasses from pure Go
tinyfsgo: exercised FSKit reply blocks without cgo
```

The example also exposes FSKit client probes:

```sh
go run ./examples/fskit/tinyfsgo --fskit-list
go run ./examples/fskit/tinyfsgo --fskit-enable dev.tmc.apple.examples.fskit.tinyfs.extension
go run ./examples/fskit/tinyfsgo --fskit-disable dev.tmc.apple.examples.fskit.tinyfs.extension
```

This does not replace the Swift `@main` extension entrypoint in `../tinyfs`.
Apple's public FSKit extension template uses `UnaryFileSystemExtension`, a
Swift `ExtensionFoundation.AppExtension` protocol. The Go code here is the
filesystem runtime slice that can sit behind that entrypoint once a public
non-Swift ExtensionFoundation bootstrap is available or a tiny Swift shim is
accepted.

An experimental signed bundle with this Go binary calling `EXExtensionMain`
launched under ExtensionFoundation, but FSKit logged:

```text
FSKit Module delegate class '(null)' is not a valid FSModule delegate. This module will not function
```

Adding `EXExtensionPrincipalClass` and `NSExtensionPrincipalClass` keys did not
change that result. Current evidence points to Swift protocol conformance and
witness-table metadata as the required public entrypoint surface.

A second private-key experiment with `EXAppExtensionClass=FSModuleExtension`
and `EXAppExtensionDelegateClass=GoTinyFSFileSystem`, both at the root and
under `EXAppExtensionAttributes`, did not produce a working pure Go appex.
After forcing a LaunchServices and pluginkit refresh, `pluginkit` reported the
new installed appex UUID and the installed plist contained both private keys,
but the host probe no longer listed the TinyFS module. A mount attempt failed
before launching the Go extension:

```text
mount_status=69
mount: Unable to invoke task
```

Without the forced identity refresh, FSKit reached the module but still left
the delegate unresolved:

```text
Delegate class (null) doesn't support probeResource:replyHandler:
```

The remaining private surface appears to be FSKit module enablement, not
filesystem callbacks. `FSClient.sharedInstance` responds to
`fetchInstalledExtensionsWithCompletionHandler:` and
`setEnabledStateForIdentifier:newState:replyHandler:`, and the Go wrappers can
list modules through `com.apple.filesystems.fskitd`. The enable call reaches
`fskitd` only after using an `NSError`-typed reply block signature, but an
ordinary `go run` process and a temporarily substituted signed host executable
both return:

```text
The operation couldn’t be completed. Operation not permitted
```

In both cases `fskitd` logs the incoming client as `entitled 0`. The
extension executable has the FSKit module entitlement, but without the
mach-lookup exception it cannot communicate with `fskitd`. Re-signing the
extension with both `com.apple.developer.fskit.fsmodule` and the
`com.apple.filesystems.fskitd` mach-lookup exception allows the call to reach
`fskitd`, but it still returns EPERM.

Reversing `fskitd` shows the management gate uses the Apple-private
`com.apple.private.LiveFS.connection` entitlement. The connection accept path
reads that entitlement with `valueForEntitlement:`, logs `Incomming connection,
entitled %d`, and stores the result in the XPC server. The
`setEnabledStateForIdentifier:newState:replyHandler:` implementation returns
`fs_errorForPOSIXError(1)` before forwarding when that bit is false. Apple's
`FSKitModuleManagement.appex` has this private entitlement and acts as the
System Settings broker; third-party host and module binaries do not.

The
`ExtensionKitSettings` dyld-cache image and Apple's
`FSKitModuleManagement.appex` contain the same enablement selector, but the
settings framework has no `.swiftinterface`, `.swiftmodule`, or `.abi.json` on
disk, the public Swift toolchain cannot import it with
`-F /System/Library/PrivateFrameworks`, and the FSKit-specific metadata in the
settings appex is non-external. Driving that path requires either the System
Settings UI or deeper private enablement entitlement/XPC work.

The repository can generate private bindings for `ExtensionKitSettings`:

```sh
APPLEGEN_KEEP_TMP=1 go run github.com/tmc/appledocs/cmd/applegen generate \
  --output /tmp/apple-eksettings-gen \
  --framework ExtensionKitSettings \
  --module github.com/tmc/apple \
  --private --private-fallback-discover --fix-iterations=0
```

That discovers the Swift-backed Objective-C scene and XPC classes used by the
Settings extension, but it does not expose a new enablement API. In particular,
`_EKIsFSExtensionEnabled` and `_EKEnableFSExtension` are visible in the dyld
shared cache as private text symbols, but they are not exported through
`dlsym`, and private generation emits no functions for them. The generated
scene classes are useful for documenting the broker path, but calls made from
Go still run with the Go process entitlement set.

The existing Swift helper pattern in this repository is still useful for a
different problem: a tiny Swift `@main` shim can satisfy
`ExtensionFoundation.AppExtension` and `UnaryFileSystemExtension` witness-table
requirements, then delegate filesystem operations to the pure Go FSKit objects.
That is an entrypoint shim, not an enablement bypass; direct module enablement
still requires Apple's private LiveFS entitlement or Apple's Settings broker.

This example includes a small proof of that entrypoint shape in `swiftshim`.
The shim defines an otherwise empty Objective-C `GoTinyFSFileSystem` class so
Swift has a concrete `UnaryFileSystemExtension.FileSystem` associated type, and
exports `TinyFSRunExtensionMain`. The Go process loads the dylib, attaches the
FSKit operation methods to that class with `objc.AddMethod`, and verifies the
same probe/load/volume callbacks used by `--smoke`.

Build and probe it with:

```sh
(cd examples/fskit/tinyfsgo/swiftshim && swift build -c release --product TinyFSShim)
go run ./examples/fskit/tinyfsgo --extension-main-probe
```

Expected additional output:

```text
tinyfsgo: swift shim loaded
tinyfsgo: attached Go FSKit operation IMPs to GoTinyFSFileSystem
tinyfsgo: resolved TinyFSRunExtensionMain without entering ExtensionFoundation main
```

The probe intentionally does not call `TinyFSRunExtensionMain`; that function
enters ExtensionFoundation's extension main loop and is for the appex binary
path. The proof only verifies that the Swift witness-table shell and Go-owned
FSKit operation methods can coexist in one process.
