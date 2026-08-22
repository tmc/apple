// Command searchfs searches a volume catalog with searchfs(2).
//
// It is intentionally a small example of the searchfs session protocol. The
// syscall is not generated yet, so the unexported mirrors below are kept next
// to the example and resolved from libSystem with purego.
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/ebitengine/purego"
)

const (
	attrName      = 1
	attrFSID      = 4
	attrObjID     = 0x20
	start         = 1
	partialNames  = 2
	matchDirs     = 4
	matchFiles    = 8
	skipInvisible = 0x20
	skipPackages  = 0x40
	utf8          = 0x08000103
)

type attrlist struct {
	BitmapCount                           uint16
	Reserved                              uint16
	Common, Volume, Directory, File, Fork uint32
}
type attrreference struct {
	Offset int32
	Length uint32
}
type timeval struct{ Sec, Usec int64 }
type fssearchblock struct {
	ReturnAttrs, ReturnBuffer unsafe.Pointer
	ReturnBufferSize          uintptr
	MaxMatches                uint32
	TimeLimit                 timeval
	SearchParams1             unsafe.Pointer
	SizeofSearchParams1       uintptr
	SearchParams2             unsafe.Pointer
	SizeofSearchParams2       uintptr
	SearchAttrs               attrlist
}
type fsid struct{ Dev, Val int32 }

type api struct {
	search  func(*byte, *fssearchblock, *uint32, uint32, uint32, *byte) int32
	getpath func(*byte, uintptr, *fsid, uint64) int64
	errno   func() *int32
}

func loadAPI() (api, error) {
	lib, err := purego.Dlopen("/usr/lib/libSystem.B.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return api{}, err
	}
	var a api
	purego.RegisterLibFunc(&a.search, lib, "searchfs")
	purego.RegisterLibFunc(&a.getpath, lib, "fsgetpath")
	purego.RegisterLibFunc(&a.errno, lib, "__error")
	return a, nil
}

func main() {
	volume := flag.String("v", ".", "volume or directory on the volume")
	dirs := flag.Bool("d", false, "match directories")
	files := flag.Bool("f", false, "match files")
	exact := flag.Bool("e", false, "match the complete name")
	skipHidden := flag.Bool("i", false, "skip invisible files")
	skipPackages := flag.Bool("p", false, "skip packages")
	name := flag.String("name", "", "name or substring to search for")
	flag.Parse()
	if *name == "" || (*dirs && *files) {
		fmt.Fprintln(os.Stderr, "usage: searchfs -name name [-e] [-d|-f] [-v volume]")
		os.Exit(2)
	}
	api, err := loadAPI()
	if err != nil {
		fmt.Fprintln(os.Stderr, "searchfs:", err)
		os.Exit(1)
	}
	if err := run(api, *volume, *name, *dirs, *files, *exact, *skipHidden, *skipPackages); err != nil {
		fmt.Fprintln(os.Stderr, "searchfs:", err)
		os.Exit(1)
	}
}

func run(a api, volume, name string, dirs, files, exact, skipHidden, omitPackages bool) error {
	vol, err := filepath.Abs(volume)
	if err != nil {
		return err
	}
	volBytes := append([]byte(vol), 0)
	query := append([]byte(name), 0)
	param := struct {
		Ref  attrreference
		Name [1]byte
	}{}
	param.Ref = attrreference{Offset: int32(unsafe.Sizeof(param.Ref)), Length: uint32(len(query))}
	paramBytes := make([]byte, 4+int(unsafe.Sizeof(param.Ref))+len(query))
	*(*uint32)(unsafe.Pointer(&paramBytes[0])) = uint32(len(paramBytes))
	*(*attrreference)(unsafe.Pointer(&paramBytes[4])) = param.Ref
	copy(paramBytes[4+unsafe.Sizeof(param.Ref):], query)
	returnAttrs := attrlist{BitmapCount: 5, Common: attrFSID | attrObjID | attrName}
	searchAttrs := attrlist{BitmapCount: 5, Common: attrName}
	buf := make([]byte, 64<<10)
	var state [560]byte // sizeof(struct searchstate), packed: 12-byte header + 548 private bytes
	block := fssearchblock{
		ReturnAttrs: unsafe.Pointer(&returnAttrs), ReturnBuffer: unsafe.Pointer(&buf[0]), ReturnBufferSize: uintptr(len(buf)), MaxMatches: 256,
		TimeLimit: timeval{Usec: 500000}, SearchParams1: unsafe.Pointer(&paramBytes[0]), SizeofSearchParams1: uintptr(len(paramBytes)),
		SearchParams2: unsafe.Pointer(&paramBytes[0]), SizeofSearchParams2: uintptr(len(paramBytes)), SearchAttrs: searchAttrs,
	}
	opts := uint32(start | matchDirs | matchFiles)
	if dirs {
		opts &^= matchFiles
	}
	if files {
		opts &^= matchDirs
	}
	if exact {
		opts &^= partialNames
	} else {
		opts |= partialNames
	}
	if skipHidden {
		opts |= skipInvisible
	}
	if omitPackages {
		opts |= skipPackages
	}
	for {
		n := uint32(256)
		block.MaxMatches = n
		result := a.search(&volBytes[0], &block, &n, utf8, opts, &state[0])
		for _, path := range parseResults(buf, uintptr(n), a) {
			fmt.Println(path)
		}
		opts &^= start
		if result == 0 {
			return nil
		}
		errno := syscall.Errno(*a.errno())
		if errno == syscall.EAGAIN {
			continue
		}
		if errno == syscall.EBUSY {
			opts |= start
			continue
		}
		if errno == syscall.ENOBUFS {
			buf = append(buf, make([]byte, len(buf))...)
			block.ReturnBuffer = unsafe.Pointer(&buf[0])
			block.ReturnBufferSize = uintptr(len(buf))
			continue
		}
		if errno == syscall.ENOTSUP || errno == syscall.EOPNOTSUPP {
			fmt.Printf("searchfs: volume does not support searchfs (errno %d)\n", errno)
			return nil
		}
		return fmt.Errorf("searchfs: errno %d", errno)
	}
}

func parseResults(buf []byte, n uintptr, a api) []string {
	var out []string
	for i := uintptr(0); i < n; i++ {
		fs, id, length, ok := nextRecord(buf)
		if !ok {
			break
		}
		buf = buf[length:]
		pathbuf := make([]byte, 4096)
		rv := a.getpath(&pathbuf[0], uintptr(len(pathbuf)), &fs, id)
		if rv > 0 {
			length := int(rv)
			if length > len(pathbuf) {
				continue
			}
			if length > 0 && pathbuf[length-1] == 0 {
				length--
			}
			out = append(out, string(pathbuf[:length]))
		}
	}
	return out
}

// nextRecord decodes the fixed prefix documented by fssearchblock. The
// returned record length is bounded by buf before any field is read.
func nextRecord(buf []byte) (fsid, uint64, int, bool) {
	if len(buf) < 20 {
		return fsid{}, 0, 0, false
	}
	length := int(binary.LittleEndian.Uint32(buf[:4]))
	if length < 20 || length > len(buf) {
		return fsid{}, 0, 0, false
	}
	return fsid{
		Dev: int32(binary.LittleEndian.Uint32(buf[4:8])),
		Val: int32(binary.LittleEndian.Uint32(buf[8:12])),
	}, binary.LittleEndian.Uint64(buf[12:20]), length, true
}
