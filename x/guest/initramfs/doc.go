// Package initramfs builds deterministic SVR4 newc cpio archives (Linux
// initramfs images) from a directory tree or a tar/tar.gz rootfs.
//
// The archive is byte-for-byte deterministic: entries are written in the order
// given, with a zero mtime and nlink 1, so the same input always yields the
// same image. [PackTree] walks a directory; [PackTar] streams a (possibly
// gzipped) tar; both accept extra in-memory [Entry] values (config files,
// injected guest binaries) that are appended before the archive trailer.
//
// The package is format-only and has no notion of any particular guest's
// layout: callers decide what extra entries to inject and under what names.
package initramfs
