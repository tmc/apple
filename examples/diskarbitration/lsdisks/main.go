// Command lsdisks lists the mounted volumes and prints the Disk Arbitration
// description of each one.
//
// It enumerates the mounted filesystems with getfsstat, looks up the BSD device
// of each one through a Disk Arbitration session, and reports selected keys of
// the disk description dictionary.
//
// Usage: lsdisks
package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"syscall"
	"unsafe"

	cf "github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/diskarbitration"
	"github.com/tmc/apple/objc"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "lsdisks: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	session := diskarbitration.DASessionCreate(0)
	if session == 0 {
		return fmt.Errorf("DASessionCreate failed")
	}
	defer cf.CFRelease(unsafe.Pointer(session))

	names, err := bsdNames()
	if err != nil {
		return err
	}
	if len(names) == 0 {
		return fmt.Errorf("no BSD-backed volumes mounted")
	}

	keys := []string{
		diskarbitration.KDADiskDescriptionVolumeNameKey,
		diskarbitration.KDADiskDescriptionVolumeKindKey,
		diskarbitration.KDADiskDescriptionVolumePathKey,
		diskarbitration.KDADiskDescriptionMediaNameKey,
		diskarbitration.KDADiskDescriptionMediaSizeKey,
		diskarbitration.KDADiskDescriptionDeviceModelKey,
		diskarbitration.KDADiskDescriptionDeviceProtocolKey,
	}

	for _, name := range names {
		disk := diskarbitration.DADiskCreateFromBSDName(0, session, name)
		if disk == 0 {
			fmt.Printf("%s: no Disk Arbitration disk\n", name)
			continue
		}
		desc := diskarbitration.DADiskCopyDescription(disk)
		if desc == 0 {
			fmt.Printf("%s: no description\n", name)
			cf.CFRelease(unsafe.Pointer(disk))
			continue
		}
		fmt.Printf("%s\n", name)
		for _, key := range keys {
			if key == "" {
				continue
			}
			v, ok := describe(desc, key)
			if !ok {
				continue
			}
			fmt.Printf("\t%-24s %s\n", strings.TrimPrefix(key, "DA"), v)
		}
		cf.CFRelease(unsafe.Pointer(desc))
		cf.CFRelease(unsafe.Pointer(disk))
	}
	return nil
}

// bsdNames returns the BSD device names of the mounted filesystems, in mount
// order, without duplicates. Filesystems not backed by a device under /dev are
// skipped.
func bsdNames() ([]string, error) {
	fs, err := syscall.Getfsstat(nil, 1 /* MNT_WAIT */)
	if err != nil {
		return nil, fmt.Errorf("getfsstat: %w", err)
	}
	buf := make([]syscall.Statfs_t, fs)
	n, err := syscall.Getfsstat(buf, 1)
	if err != nil {
		return nil, fmt.Errorf("getfsstat: %w", err)
	}
	var names []string
	seen := make(map[string]bool)
	for _, st := range buf[:n] {
		from := gostring(st.Mntfromname[:])
		if !strings.HasPrefix(from, "/dev/") {
			continue
		}
		name := path.Base(from)
		if seen[name] {
			continue
		}
		seen[name] = true
		names = append(names, name)
	}
	return names, nil
}

func gostring(b []int8) string {
	s := make([]byte, 0, len(b))
	for _, c := range b {
		if c == 0 {
			break
		}
		s = append(s, byte(c))
	}
	return string(s)
}

// describe returns a printable rendering of the description value for key, and
// reports whether the key was present and of a type it knows how to render.
func describe(desc cf.CFDictionaryRef, key string) (string, bool) {
	k := cf.CFStringCreateWithCString(0, key, uint32(cf.KCFStringEncodingUTF8))
	if k == 0 {
		return "", false
	}
	defer cf.CFRelease(unsafe.Pointer(k))

	v := cf.CFDictionaryGetValue(desc, unsafe.Pointer(k))
	if v == nil {
		return "", false
	}
	switch cf.CFGetTypeID(v) {
	case cf.CFStringGetTypeID():
		return cfString(cf.CFStringRef(uintptr(v))), true
	case cf.CFNumberGetTypeID():
		var n int64
		if !cf.CFNumberGetValue(cf.CFNumberRef(uintptr(v)), cf.KCFNumberSInt64Type, unsafe.Pointer(&n)) {
			return "", false
		}
		return fmt.Sprintf("%d", n), true
	case cf.CFBooleanGetTypeID():
		return fmt.Sprintf("%t", cf.CFBooleanGetValue(cf.CFBooleanRef(uintptr(v)))), true
	case cf.CFURLGetTypeID():
		p := cf.CFURLCopyFileSystemPath(cf.CFURLRef(uintptr(v)), cf.KCFURLPOSIXPathStyle)
		if p == 0 {
			return "", false
		}
		defer cf.CFRelease(unsafe.Pointer(p))
		return cfString(p), true
	}
	return "", false
}

// cfString converts a CFString to a Go string.
func cfString(s cf.CFStringRef) string {
	if p := cf.CFStringGetCStringPtr(s, uint32(cf.KCFStringEncodingUTF8)); p != nil {
		return objc.GoString(p)
	}
	buf := make([]byte, 4*cf.CFStringGetLength(s)+1)
	if !cf.CFStringGetCString(s, &buf[0], len(buf), uint32(cf.KCFStringEncodingUTF8)) {
		return ""
	}
	return objc.GoString(&buf[0])
}
