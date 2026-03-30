package kernel

import "syscall"

// Attrlist matches struct attrlist from sys/attr.h.
type Attrlist struct {
	Bitmapcount uint16
	Reserved    uint16
	Commonattr  Attrgroup_t
	Volattr     Attrgroup_t
	Dirattr     Attrgroup_t
	Fileattr    Attrgroup_t
	Forkattr    Attrgroup_t
}

// Stat64 matches struct stat64 from sys/stat.h.
type Stat64 struct {
	St_dev           int32
	St_mode          uint16
	St_nlink         uint16
	St_ino           uint64
	St_uid           uint32
	St_gid           uint32
	St_rdev          int32
	St_atimespec     syscall.Timespec
	St_mtimespec     syscall.Timespec
	St_ctimespec     syscall.Timespec
	St_birthtimespec syscall.Timespec
	St_size          int64
	St_blocks        int64
	St_blksize       int32
	St_flags         uint32
	St_gen           uint32
	St_lspare        int32
	St_qspare        [2]int64
}
