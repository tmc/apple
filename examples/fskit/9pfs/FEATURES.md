# 9pfs Feature Matrix

`9pfs` maps a 9P client connection into FSKit volume callbacks.

## Implemented

| Operation | 9P2000 | 9P2000.L | FSKit callback |
| --- | --- | --- | --- |
| attach | yes | yes | `loadResource` |
| stat | yes | yes | `getAttributes` |
| readdir | yes | yes | `enumerateDirectory` |
| lookup | yes | yes | `lookupItemNamed` |
| read | yes | yes | `readFromFile` |
| write | yes | yes | `writeContents` |
| create file | yes | yes | `createItemNamed` |
| create directory | yes | yes | `createItemNamed` |
| remove | yes | yes | `removeItem` |
| rename | same-directory | yes | `renameItem` |
| chmod/truncate | yes | yes | `setAttributes` |
| mtime | client request; server-dependent | yes | `setAttributes` |
| symlink creation/readlink | no | yes | `createSymbolicLinkNamed`, `readSymbolicLink` |
| hard links | no | yes | `createLink` |
| device/FIFO/socket attributes | stat only | stat only | `getAttributes` |
| extended attributes | no | yes | `getXattr`, `setXattr`, `listXattrs`, `supportedXattrNames` |
| xattr create/replace/delete policies | no | yes | `setXattr` policy argument |
| open/close notifications | yes | yes | `openItem`, `closeItem` |
| access checks | local allow | local allow | `checkAccessToItem` |
| open-unlink emulation | yes | yes | `enableOpenUnlinkEmulation` |
| preallocation by size extension | yes | yes | `preallocateSpaceForItem` |
| volume statistics | synthetic | synthetic | `volumeStatistics` |

## Not Yet Implemented

| Operation | Reason |
| --- | --- |
| device node creation | FSKit's generic item creation callback is documented for files and directories; existing device, FIFO, and socket nodes are reported through attributes. |
| advisory locking | 9P2000.L exposes locks, but the current FSKit operation set does not advertise a lock mapping. |
| authentication | The current examples attach with the local user or anonymous defaults. |

## Verification

Local gates:

```sh
./verify-local.sh
```

Installed macOS mount gate:

```sh
./preflight-installed.sh
./test-installed-live-9p2000l.sh "$HOME/9pfs-mnt-$(date +%s)"
```

The installed gate starts a disposable 9P2000.L server and mounts it through
FSKit using `/sbin/mount -F -t 9pfs`, then verifies mounted read, write,
rename, chmod, mtime, truncate, symlink, hardlink, xattr, and remove operations. It
requires a signed host app with the 9pfs extension enabled. The
`/Library/Filesystems/9pfs.fs` helper is optional for plain `mount -t 9pfs`.
