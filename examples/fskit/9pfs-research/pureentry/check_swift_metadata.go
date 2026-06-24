//go:build ignore

package main

import (
	"debug/macho"
	"flag"
	"fmt"
	"os"
	"strings"
)

var requiredSections = []struct {
	seg  string
	name string
}{
	{"__TEXT", "__swift5_typeref"},
	{"__TEXT", "__swift5_assocty"},
	{"__TEXT", "__swift5_entry"},
	{"__TEXT", "__swift5_fieldmd"},
	{"__TEXT", "__swift5_proto"},
	{"__TEXT", "__swift5_types"},
}

var requiredSymbols = []string{
	"_OBJC_CLASS_$_NinePFileSystem",
	"_NinePFSDelayedLoadResource",
	"_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAMc",
	"_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAMc",
	"_$s9swiftmeta16NinePFSExtensionVAC19ExtensionFoundation03AppD0AAWL",
	"_$s9swiftmeta16NinePFSExtensionVAC5FSKit24UnaryFileSystemExtensionAAWL",
}

var baseForbiddenSymbols = []string{
	"_BorrowedExtensionMetadata",
	"_NinePFSExtensionMainEntry",
}

var reducedForbiddenSymbols = []string{
	"__swift_FORCE_LOAD_$_swiftFoundation_$_swiftmeta",
	"__OBJC_LABEL_PROTOCOL_$_FSUnaryFileSystemOperations",
	"__OBJC_LABEL_PROTOCOL_$_NSObject",
}

func main() {
	reduced := flag.Bool("reduced", false, "check reduced bundled Swift metadata symbol surface")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: check_swift_metadata [-reduced] macho")
		os.Exit(2)
	}
	path := flag.Arg(0)
	f, err := macho.Open(path)
	check(err)
	defer f.Close()

	var failed bool
	for _, want := range requiredSections {
		sec := section(f, want.seg, want.name)
		if sec == nil || sec.Size == 0 {
			fmt.Fprintf(os.Stderr, "missing non-empty section %s,%s\n", want.seg, want.name)
			failed = true
			continue
		}
		fmt.Printf("section %s,%s size=%#x\n", want.seg, want.name, sec.Size)
	}

	syms := map[string]bool{}
	if f.Symtab != nil {
		for _, sym := range f.Symtab.Syms {
			syms[sym.Name] = true
		}
	}
	for _, name := range requiredSymbols {
		if !syms[name] && !syms[strings.TrimPrefix(name, "_")] {
			fmt.Fprintf(os.Stderr, "missing symbol %s\n", name)
			failed = true
			continue
		}
		fmt.Printf("symbol %s present\n", name)
	}
	for _, name := range baseForbiddenSymbols {
		if syms[name] || syms[strings.TrimPrefix(name, "_")] {
			fmt.Fprintf(os.Stderr, "forbidden symbol %s present\n", name)
			failed = true
		} else {
			fmt.Printf("symbol %s absent\n", name)
		}
	}
	if *reduced {
		for _, name := range reducedForbiddenSymbols {
			if syms[name] || syms[strings.TrimPrefix(name, "_")] {
				fmt.Fprintf(os.Stderr, "forbidden reduced metadata symbol %s present\n", name)
				failed = true
			} else {
				fmt.Printf("symbol %s absent\n", name)
			}
		}
	}
	if failed {
		os.Exit(1)
	}
}

func section(f *macho.File, seg, name string) *macho.Section {
	for _, sec := range f.Sections {
		if sec.Seg == seg && sec.Name == name {
			return sec
		}
	}
	return nil
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
