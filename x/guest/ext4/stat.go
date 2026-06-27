package ext4

import "time"

// statData is the portable subset of stat(2) the builder needs: ownership,
// access time, and the link-count/device/inode used to detect hardlinks.
// statOf, defined per-GOOS, fills it from an os.FileInfo's underlying Sys()
// value; on platforms where that value is unavailable it returns ok=false and
// the builder falls back to ModTime and zero ownership.
type statData struct {
	uid, gid uint32
	atime    time.Time
	nlink    uint64
	dev, ino uint64
}
