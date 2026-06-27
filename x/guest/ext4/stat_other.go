//go:build !darwin && !linux

package ext4

import "os"

// statOf has no portable way to read uid/gid/atime/hardlink data on platforms
// other than darwin and linux, so it reports ok=false. The builder then uses
// ModTime for atime, zero ownership, and treats every file as unlinked. The
// package still builds and produces a valid image; only host ownership and
// hardlink coalescing are unavailable off darwin/linux.
func statOf(os.FileInfo) (statData, bool) {
	return statData{}, false
}
