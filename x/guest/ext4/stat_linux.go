//go:build linux

package ext4

import (
	"os"
	"syscall"
	"time"
)

func statOf(fi os.FileInfo) (statData, bool) {
	st, ok := fi.Sys().(*syscall.Stat_t)
	if !ok || st == nil {
		return statData{}, false
	}
	return statData{
		uid:   st.Uid,
		gid:   st.Gid,
		atime: time.Unix(st.Atim.Unix()),
		nlink: uint64(st.Nlink),
		dev:   uint64(st.Dev),
		ino:   uint64(st.Ino),
	}, true
}
