//go:build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type manifest struct {
	Sections []section `json:"sections"`
	Symbols  []symbol  `json:"symbols"`
}

type section struct {
	Segment   string     `json:"segment"`
	Name      string     `json:"name"`
	Size      uint64     `json:"size"`
	Align     uint32     `json:"align,omitempty"`
	Strings   []strlit   `json:"strings,omitempty"`
	ByteSpans []byteSpan `json:"byteSpans,omitempty"`
	Words     []uint32   `json:"words,omitempty"`
	Relocs    []reloc    `json:"relocs,omitempty"`
}

type strlit struct {
	Offset uint32 `json:"offset"`
	Value  string `json:"value"`
}

type byteSpan struct {
	Offset uint32 `json:"offset"`
	Hex    string `json:"hex"`
}

type reloc struct {
	Addr   uint32 `json:"addr"`
	Symbol string `json:"symbol,omitempty"`
	Type   uint8  `json:"type"`
	Len    uint8  `json:"len"`
	Pcrel  bool   `json:"pcrel,omitempty"`
	Extern bool   `json:"extern,omitempty"`
}

type symbol struct {
	Name  string `json:"name"`
	Value uint64 `json:"value"`
	Sect  uint8  `json:"sect,omitempty"`
	Type  uint8  `json:"type,omitempty"`
}

func main() {
	sectionPrefix := flag.String("section-prefix", "__swift5_", "only explain sections whose name has this prefix; empty explains all")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: explain_swiftmeta_manifest [-section-prefix prefix] manifest.json")
		os.Exit(2)
	}
	data, err := os.ReadFile(flag.Arg(0))
	check(err)
	var m manifest
	check(json.Unmarshal(data, &m))

	refs := map[string]int{}
	for _, sec := range m.Sections {
		for _, r := range sec.Relocs {
			if r.Symbol != "" {
				refs[r.Symbol]++
			}
		}
	}

	fmt.Printf("# Swift metadata manifest explanation\n\n")
	fmt.Printf("Sections: %d\n\n", len(m.Sections))
	for i, sec := range m.Sections {
		if *sectionPrefix != "" && !strings.HasPrefix(sec.Name, *sectionPrefix) {
			continue
		}
		fmt.Printf("## %s,%s\n\n", sec.Segment, sec.Name)
		fmt.Printf("- index: %d\n", i+1)
		fmt.Printf("- size: %#x\n", sec.Size)
		fmt.Printf("- align: %d\n", sec.Align)
		if len(sec.Words) > 0 {
			fmt.Printf("- words: %d\n", len(sec.Words))
			relocsByAddr := map[uint32][]reloc{}
			for _, r := range sec.Relocs {
				relocsByAddr[r.Addr] = append(relocsByAddr[r.Addr], r)
			}
			for i, word := range sec.Words {
				off := uint32(i * 4)
				fmt.Printf("  - %#x: %#08x", off, word)
				if relocs := relocsByAddr[off]; len(relocs) > 0 {
					var parts []string
					for _, r := range relocs {
						name := r.Symbol
						if name == "" {
							name = "<section>"
						}
						parts = append(parts, fmt.Sprintf("type=%d extern=%v pcrel=%v %s", r.Type, r.Extern, r.Pcrel, name))
					}
					fmt.Printf(" // %s", strings.Join(parts, "; "))
				}
				fmt.Println()
			}
		}
		if len(sec.Strings) > 0 {
			fmt.Printf("- strings: %d\n", len(sec.Strings))
			for _, s := range sec.Strings {
				fmt.Printf("  - %#x: %q\n", s.Offset, s.Value)
			}
		}
		if len(sec.ByteSpans) > 0 {
			fmt.Printf("- byte spans: %d\n", len(sec.ByteSpans))
			for _, span := range sec.ByteSpans {
				fmt.Printf("  - %#x: %s\n", span.Offset, span.Hex)
			}
		}
		if len(sec.Relocs) > 0 {
			fmt.Printf("- relocations: %d\n", len(sec.Relocs))
			for _, r := range sec.Relocs {
				name := r.Symbol
				if name == "" {
					name = "<section>"
				}
				fmt.Printf("  - addr=%#x type=%d len=%d extern=%v pcrel=%v symbol=%s\n",
					r.Addr, r.Type, r.Len, r.Extern, r.Pcrel, name)
			}
		}
		fmt.Println()
	}

	var unreferenced []symbol
	for _, sym := range m.Symbols {
		if sym.Sect == 0 || refs[sym.Name] > 0 {
			continue
		}
		unreferenced = append(unreferenced, sym)
	}
	sort.Slice(unreferenced, func(i, j int) bool {
		return unreferenced[i].Name < unreferenced[j].Name
	})
	fmt.Printf("## Unreferenced section-defined symbols\n\n")
	for _, sym := range unreferenced {
		fmt.Printf("- %s sect=%d type=%#x value=%#x\n", sym.Name, sym.Sect, sym.Type, sym.Value)
	}
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
