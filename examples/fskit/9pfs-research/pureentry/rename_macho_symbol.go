//go:build ignore

package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "usage: %s in out old new\n", os.Args[0])
		os.Exit(2)
	}
	in, out, old, new := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	b, err := os.ReadFile(in)
	check(err)
	if len(b) < 32 {
		fatalf("short mach-o")
	}
	bo := binary.LittleEndian
	if magic := bo.Uint32(b[0:4]); magic != 0xfeedfacf {
		fatalf("unexpected magic %#x", magic)
	}
	ncmds := bo.Uint32(b[16:20])
	cmdOff := 32
	symcmd := -1
	var symoff, nsyms, stroff, strsize uint32
	for i := uint32(0); i < ncmds; i++ {
		if cmdOff+8 > len(b) {
			fatalf("bad load command")
		}
		cmd := bo.Uint32(b[cmdOff : cmdOff+4])
		cmdsize := bo.Uint32(b[cmdOff+4 : cmdOff+8])
		if cmd == 0x2 {
			symcmd = cmdOff
			symoff = bo.Uint32(b[cmdOff+8 : cmdOff+12])
			nsyms = bo.Uint32(b[cmdOff+12 : cmdOff+16])
			stroff = bo.Uint32(b[cmdOff+16 : cmdOff+20])
			strsize = bo.Uint32(b[cmdOff+20 : cmdOff+24])
		}
		cmdOff += int(cmdsize)
	}
	if symcmd < 0 {
		fatalf("no LC_SYMTAB")
	}
	if int(stroff+strsize) != len(b) {
		fatalf("string table not at EOF: end=%d len=%d", stroff+strsize, len(b))
	}
	newOff := strsize
	found := 0
	for i := uint32(0); i < nsyms; i++ {
		off := int(symoff + i*16)
		strx := bo.Uint32(b[off : off+4])
		if strx >= strsize {
			continue
		}
		nameStart := int(stroff + strx)
		nameEnd := nameStart
		for nameEnd < int(stroff+strsize) && b[nameEnd] != 0 {
			nameEnd++
		}
		if string(b[nameStart:nameEnd]) == old {
			bo.PutUint32(b[off:off+4], newOff)
			found++
		}
	}
	if found == 0 {
		fatalf("symbol %q not found", old)
	}
	b = append(b, []byte(new)...)
	b = append(b, 0)
	bo.PutUint32(b[symcmd+20:symcmd+24], strsize+uint32(len(new)+1))
	check(os.WriteFile(out, b, 0755))
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
