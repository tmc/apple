//go:build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type manifest struct {
	Sections []section `json:"sections"`
}

type section struct {
	Segment   string      `json:"segment"`
	Name      string      `json:"name"`
	Size      uint64      `json:"size"`
	Align     uint32      `json:"align,omitempty"`
	SHA256    string      `json:"sha256,omitempty"`
	Words     []uint32    `json:"words,omitempty"`
	Strings   []stringRec `json:"strings,omitempty"`
	ByteSpans []byteSpan  `json:"byteSpans,omitempty"`
	Relocs    []reloc     `json:"relocs,omitempty"`
}

type stringRec struct {
	Offset uint32 `json:"offset"`
	Value  string `json:"value"`
}

type byteSpan struct {
	Offset uint32 `json:"offset"`
	Hex    string `json:"hex"`
}

type reloc struct {
	Addr   uint32 `json:"addr"`
	Value  uint32 `json:"value"`
	Symbol string `json:"symbol,omitempty"`
	Type   uint8  `json:"type"`
	Len    uint8  `json:"len"`
	Pcrel  bool   `json:"pcrel,omitempty"`
	Extern bool   `json:"extern,omitempty"`
}

func main() {
	checkPath := flag.String("check", "", "manifest to check generated Swift metadata records against")
	flag.Parse()
	if *checkPath == "" || flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "usage: encode_swift_metadata -check manifest.json")
		os.Exit(2)
	}
	data, err := os.ReadFile(*checkPath)
	check(err)
	var m manifest
	check(json.Unmarshal(data, &m))
	want := generatedSections()
	for _, gen := range want {
		got, ok := findSection(m.Sections, gen.Segment, gen.Name)
		if !ok {
			check(fmt.Errorf("missing section %s,%s", gen.Segment, gen.Name))
		}
		got.SHA256 = ""
		if !sameSection(got, gen) {
			check(fmt.Errorf("section %s,%s differs from generated record", gen.Segment, gen.Name))
		}
	}
	fmt.Printf("swift metadata encoder ok: sections=%d\n", len(want))
}

func generatedSections() []section {
	return []section{
		{
			Segment: "__TEXT",
			Name:    "__swift5_capture",
			Size:    48,
			Align:   2,
			Words: []uint32{
				0x00000001, 0x00000000, 0x00000000, 0xfffffff4,
				0x00000001, 0x00000000, 0x00000000, 0xfffffff4,
				0x00000001, 0x00000000, 0x00000000, 0xfffffff4,
			},
			Relocs: []reloc{
				{Addr: 0xc, Symbol: "l__swift5_reflection_descriptor", Type: 1, Len: 2, Extern: true},
				{Addr: 0xc, Symbol: "_symbolic So8FSVolumeCSg______pSgIeghgg_ s5ErrorP", Type: 0, Len: 2, Extern: true},
				{Addr: 0x1c, Symbol: "l__swift5_reflection_descriptor.3", Type: 1, Len: 2, Extern: true},
				{Addr: 0x1c, Symbol: "_symbolic So8FSVolumeCSgSo7NSErrorCSgIeyBhyy_", Type: 0, Len: 2, Extern: true},
				{Addr: 0x2c, Symbol: "l__swift5_reflection_descriptor.6", Type: 1, Len: 2, Extern: true},
				{Addr: 0x2c, Symbol: "_symbolic So8FSVolumeCSg______pSgIeghgg_ s5ErrorP", Type: 0, Len: 2, Extern: true},
			},
		},
		{
			Segment: "__TEXT",
			Name:    "__swift5_entry",
			Size:    8,
			Align:   2,
			Relocs: []reloc{
				{Addr: 0, Symbol: "l_entry_point", Type: 1, Len: 2, Extern: true},
				{Addr: 0, Symbol: "_main", Type: 0, Len: 2, Extern: true},
			},
		},
		{
			Segment: "__TEXT",
			Name:    "__swift5_typeref",
			Size:    204,
			Align:   1,
			Strings: []stringRec{
				{Offset: 0x13, Value: "_pSgIeghgg_"},
				{Offset: 0x34, Value: "$s5FSKit24UnaryFileSystemExtensionP"},
				{Offset: 0x6b, Value: "_Qo_"},
				{Offset: 0x70, Value: "$s19ExtensionFoundation03AppA0P"},
				{Offset: 0x90, Value: "So17FSUnaryFileSystemC"},
				{Offset: 0xa8, Value: "So8FSVolumeCSgSo7NSErrorCSgIeyBhyy_"},
			},
			ByteSpans: []byteSpan{
				{Offset: 0x0, Hex: "536f384653566f6c756d6543536702"},
				{Offset: 0x20, Hex: "ff07feffffff"},
				{Offset: 0x28, Hex: "01ffffffff"},
				{Offset: 0x2e, Hex: "01ffffffff"},
				{Offset: 0x58, Hex: "ff07feffffff"},
				{Offset: 0x60, Hex: "02"},
				{Offset: 0x65, Hex: "7901f9ffffff"},
			},
			Relocs: []reloc{
				{Addr: 0xf, Symbol: "_$ss5ErrorMp", Type: 7, Len: 2, Pcrel: true, Extern: true},
				{Addr: 0x22, Symbol: "_associated conformance 9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAA0H10Foundation03AppH0", Type: 1, Len: 2, Extern: true},
				{Addr: 0x22, Symbol: "_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAA0H10Foundation03AppH0PWb", Type: 0, Len: 2, Extern: true},
				{Addr: 0x29, Symbol: "_symbolic _____ 9swiftmeta15NinePFileSystemC", Type: 1, Len: 2, Extern: true},
				{Addr: 0x29, Symbol: "_$s9swiftmeta15NinePFileSystemCMn", Type: 0, Len: 2, Extern: true},
				{Addr: 0x2f, Symbol: "_symbolic _____ 9swiftmeta16NinePFSExtensionV", Type: 1, Len: 2, Extern: true},
				{Addr: 0x2f, Symbol: "_$s9swiftmeta16NinePFSExtensionVMn", Type: 0, Len: 2, Extern: true},
				{Addr: 0x5a, Symbol: "_associated conformance 9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AA13ConfigurationAdEP_AD0fdG0", Type: 1, Len: 2, Extern: true},
				{Addr: 0x5a, Symbol: "_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AA13ConfigurationAdEP_AD0fdG0PWT", Type: 0, Len: 2, Extern: true},
				{Addr: 0x67, Symbol: "_symbolic _____y______Qo_ 5FSKit24UnaryFileSystemExtensionPAAE13configurationQrvpQO 9swiftmeta16NinePFSExtensionV", Type: 1, Len: 2, Extern: true},
				{Addr: 0x67, Symbol: "_$s9swiftmeta16NinePFSExtensionVMn", Type: 0, Len: 2, Extern: true},
				{Addr: 0x61, Symbol: "_$s5FSKit24UnaryFileSystemExtensionPAAE13configurationQrvpQOMQ", Type: 7, Len: 2, Pcrel: true, Extern: true},
			},
		},
		{
			Segment: "__TEXT",
			Name:    "__swift5_reflstr",
			Size:    36,
			Strings: []stringRec{
				{Offset: 0x0, Value: "FileSystem"},
				{Offset: 0xb, Value: "Configuration"},
				{Offset: 0x19, Value: "fileSystem"},
			},
		},
		{
			Segment: "__TEXT",
			Name:    "__swift5_fieldmd",
			Size:    44,
			Align:   2,
			Words: []uint32{
				0x00000000, 0xfffffffc, 0x000c0007, 0x00000000,
				0x00000000, 0x00000000, 0x000c0000, 0x00000001,
				0x00000000, 0xffffffec, 0xffffffe8,
			},
			Relocs: []reloc{
				{Addr: 0x4, Symbol: "_$s9swiftmeta15NinePFileSystemCMF", Type: 1, Len: 2, Extern: true},
				{Addr: 0x4, Symbol: "_symbolic So17FSUnaryFileSystemC", Type: 0, Len: 2, Extern: true},
				{Addr: 0x0, Symbol: "_$s9swiftmeta15NinePFileSystemCMF", Type: 1, Len: 2, Extern: true},
				{Addr: 0x0, Symbol: "_symbolic _____ 9swiftmeta15NinePFileSystemC", Type: 0, Len: 2, Extern: true},
				{Addr: 0x28, Symbol: "_$s9swiftmeta16NinePFSExtensionVMF", Type: 1, Len: 2, Extern: true},
				{Addr: 0x28, Symbol: "l___unnamed_3", Type: 0, Len: 2, Extern: true},
				{Addr: 0x24, Symbol: "_$s9swiftmeta16NinePFSExtensionVMF", Type: 1, Len: 2, Extern: true},
				{Addr: 0x24, Symbol: "_symbolic _____ 9swiftmeta15NinePFileSystemC", Type: 0, Len: 2, Extern: true},
				{Addr: 0x10, Symbol: "_$s9swiftmeta16NinePFSExtensionVMF", Type: 1, Len: 2, Extern: true},
				{Addr: 0x10, Symbol: "_symbolic _____ 9swiftmeta16NinePFSExtensionV", Type: 0, Len: 2, Extern: true},
			},
		},
		{
			Segment: "__TEXT",
			Name:    "__swift5_assocty",
			Size:    48,
			Align:   2,
			Words: []uint32{
				0x00000000, 0xfffffffc, 0x00000001, 0x00000008,
				0xfffffff0, 0xffffffec, 0x00000000, 0xfffffffc,
				0x00000001, 0x00000008, 0xfffffff0, 0xffffffec,
			},
			Relocs: []reloc{
				{Addr: 0x14, Symbol: "_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x14, Symbol: "_symbolic _____ 9swiftmeta15NinePFileSystemC", Type: 0, Len: 2, Extern: true},
				{Addr: 0x10, Symbol: "_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x10, Symbol: "l___unnamed_1", Type: 0, Len: 2, Extern: true},
				{Addr: 0x4, Symbol: "_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x4, Symbol: "_symbolic $s5FSKit24UnaryFileSystemExtensionP", Type: 0, Len: 2, Extern: true},
				{Addr: 0x0, Symbol: "_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x0, Symbol: "_symbolic _____ 9swiftmeta16NinePFSExtensionV", Type: 0, Len: 2, Extern: true},
				{Addr: 0x2c, Symbol: "_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x2c, Symbol: "_symbolic _____y______Qo_ 5FSKit24UnaryFileSystemExtensionPAAE13configurationQrvpQO 9swiftmeta16NinePFSExtensionV", Type: 0, Len: 2, Extern: true},
				{Addr: 0x28, Symbol: "_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x28, Symbol: "l___unnamed_2", Type: 0, Len: 2, Extern: true},
				{Addr: 0x1c, Symbol: "_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x1c, Symbol: "_symbolic $s19ExtensionFoundation03AppA0P", Type: 0, Len: 2, Extern: true},
				{Addr: 0x18, Symbol: "_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAMA", Type: 1, Len: 2, Extern: true},
				{Addr: 0x18, Symbol: "_symbolic _____ 9swiftmeta16NinePFSExtensionV", Type: 0, Len: 2, Extern: true},
			},
		},
		{
			Segment: "__TEXT",
			Name:    "__swift5_proto",
			Size:    8,
			Align:   2,
			Relocs: []reloc{
				{Addr: 0, Symbol: "l_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAHc", Type: 1, Len: 2, Extern: true},
				{Addr: 0, Symbol: "_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAMc", Type: 0, Len: 2, Extern: true},
				{Addr: 4, Symbol: "l_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAHc", Type: 1, Len: 2, Extern: true},
				{Addr: 4, Symbol: "_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAMc", Type: 0, Len: 2, Extern: true},
			},
		},
		{
			Segment: "__TEXT",
			Name:    "__swift5_types",
			Size:    8,
			Align:   2,
			Relocs: []reloc{
				{Addr: 0, Symbol: "l_$s9swiftmeta15NinePFileSystemCHn", Type: 1, Len: 2, Extern: true},
				{Addr: 0, Symbol: "_$s9swiftmeta15NinePFileSystemCMn", Type: 0, Len: 2, Extern: true},
				{Addr: 4, Symbol: "l_$s9swiftmeta16NinePFSExtensionVHn", Type: 1, Len: 2, Extern: true},
				{Addr: 4, Symbol: "_$s9swiftmeta16NinePFSExtensionVMn", Type: 0, Len: 2, Extern: true},
			},
		},
	}
}

func findSection(sections []section, segment, name string) (section, bool) {
	for _, sec := range sections {
		if sec.Segment == segment && sec.Name == name {
			return sec, true
		}
	}
	return section{}, false
}

func sameSection(a, b section) bool {
	if a.Segment != b.Segment || a.Name != b.Name || a.Size != b.Size || a.Align != b.Align ||
		len(a.Words) != len(b.Words) || len(a.Strings) != len(b.Strings) ||
		len(a.ByteSpans) != len(b.ByteSpans) || len(a.Relocs) != len(b.Relocs) {
		return false
	}
	for i := range a.Words {
		if a.Words[i] != b.Words[i] {
			return false
		}
	}
	for i := range a.Strings {
		if a.Strings[i] != b.Strings[i] {
			return false
		}
	}
	for i := range a.ByteSpans {
		if a.ByteSpans[i] != b.ByteSpans[i] {
			return false
		}
	}
	for i := range a.Relocs {
		if !sameReloc(a.Relocs[i], b.Relocs[i]) {
			return false
		}
	}
	return true
}

func sameReloc(a, b reloc) bool {
	a.Value = 0
	b.Value = 0
	return a == b
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
