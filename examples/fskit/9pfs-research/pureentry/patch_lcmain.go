//go:build ignore

package main

import (
	"debug/macho"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	mhMagic64 = 0xfeedfacf
	lcMain    = 0x80000028
)

func main() {
	in := flag.String("in", "", "input Mach-O")
	out := flag.String("out", "", "output Mach-O")
	target := flag.String("target", "main.nsextMainEntry.abi0", "entry symbol")
	flag.Parse()
	if *in == "" || *out == "" {
		fmt.Fprintln(os.Stderr, "usage: patch_lcmain -in file -out file -target symbol")
		os.Exit(2)
	}
	data, err := os.ReadFile(*in)
	check(err)
	if len(data) < 32 || binary.LittleEndian.Uint32(data[:4]) != mhMagic64 {
		fatalf("%s is not a little-endian 64-bit Mach-O", *in)
	}
	ncmds := binary.LittleEndian.Uint32(data[16:20])
	textVM, textFile, err := textMapping(data, ncmds)
	check(err)
	targetAddr, err := symbolAddr(*in, *target)
	check(err)
	if targetAddr < textVM {
		fatalf("target %#x before __TEXT vmaddr %#x", targetAddr, textVM)
	}
	entryoff := targetAddr - textVM + textFile
	patched := append([]byte(nil), data...)
	off := 32
	found := false
	for i := uint32(0); i < ncmds; i++ {
		cmd := binary.LittleEndian.Uint32(patched[off : off+4])
		size := binary.LittleEndian.Uint32(patched[off+4 : off+8])
		if cmd == lcMain {
			binary.LittleEndian.PutUint64(patched[off+8:off+16], entryoff)
			found = true
			break
		}
		off += int(size)
	}
	if !found {
		fatalf("missing LC_MAIN")
	}
	check(os.WriteFile(*out, patched, 0755))
}

func textMapping(data []byte, ncmds uint32) (vmaddr, fileoff uint64, err error) {
	off := 32
	for i := uint32(0); i < ncmds; i++ {
		if off+72 > len(data) {
			return 0, 0, fmt.Errorf("load command out of range")
		}
		cmd := binary.LittleEndian.Uint32(data[off : off+4])
		size := binary.LittleEndian.Uint32(data[off+4 : off+8])
		if cmd == 0x19 && cstring(data[off+8:off+24]) == "__TEXT" {
			return binary.LittleEndian.Uint64(data[off+24 : off+32]), binary.LittleEndian.Uint64(data[off+40 : off+48]), nil
		}
		off += int(size)
	}
	return 0, 0, fmt.Errorf("missing __TEXT")
}

func symbolAddr(path, name string) (uint64, error) {
	f, err := macho.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	if f.Symtab == nil {
		return 0, fmt.Errorf("missing symtab")
	}
	for _, s := range f.Symtab.Syms {
		if s.Name == name || strings.TrimPrefix(s.Name, "_") == name {
			return s.Value, nil
		}
	}
	return 0, fmt.Errorf("symbol %q not found", name)
}

func cstring(b []byte) string {
	if i := strings.IndexByte(string(b), 0); i >= 0 {
		return string(b[:i])
	}
	return string(b)
}

func check(err error) {
	if err != nil {
		fatalf("%v", err)
	}
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
