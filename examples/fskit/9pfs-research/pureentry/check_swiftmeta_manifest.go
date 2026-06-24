//go:build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type manifest struct {
	Sections []section `json:"sections"`
	Symbols  []symbol  `json:"symbols"`
}

type section struct {
	Segment string  `json:"segment"`
	Name    string  `json:"name"`
	Relocs  []reloc `json:"relocs,omitempty"`
}

type reloc struct {
	Symbol string `json:"symbol,omitempty"`
}

type symbol struct {
	Name string `json:"name"`
	Sect uint8  `json:"sect,omitempty"`
	Type uint8  `json:"type,omitempty"`
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: check_swiftmeta_manifest manifest.json")
		os.Exit(2)
	}
	data, err := os.ReadFile(flag.Arg(0))
	check(err)
	var m manifest
	check(json.Unmarshal(data, &m))

	refs := map[string]bool{}
	for _, sec := range m.Sections {
		for _, r := range sec.Relocs {
			if r.Symbol != "" {
				refs[r.Symbol] = true
			}
		}
	}

	var failed bool
	for _, sym := range m.Symbols {
		if refs[sym.Name] || sym.Sect == 0 || allowedUnreferencedSymbol(sym.Name) {
			continue
		}
		fmt.Fprintf(os.Stderr, "unexpected unreferenced symbol %s\n", sym.Name)
		failed = true
	}
	if failed {
		os.Exit(1)
	}
	fmt.Printf("swift metadata manifest ok: sections=%d symbols=%d\n", len(m.Sections), len(m.Symbols))
}

func allowedUnreferencedSymbol(name string) bool {
	if name == "_OBJC_CLASS_$_NinePFileSystem" {
		return true
	}
	return strings.HasPrefix(name, "__swift_FORCE_LOAD_$_") &&
		strings.HasSuffix(name, "_$_swiftmeta") &&
		name != "__swift_FORCE_LOAD_$_swiftFoundation_$_swiftmeta"
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
