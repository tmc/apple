//go:build ignore

package main

import (
	"bytes"
	"crypto/sha256"
	"debug/macho"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type manifest struct {
	Path                   string         `json:"path"`
	Header                 header         `json:"header,omitempty"`
	Dylibs                 []dylib        `json:"dylibs,omitempty"`
	BuildVersion           *buildVersion  `json:"buildVersion,omitempty"`
	DataInCode             *linkeditData  `json:"dataInCode,omitempty"`
	LinkerOptimizationHint *linkeditData  `json:"linkerOptimizationHint,omitempty"`
	LinkerOptions          []linkerOption `json:"linkerOptions,omitempty"`
	Sections               []section      `json:"sections"`
	Symbols                []symbol       `json:"symbols"`
}

type header struct {
	Magic      uint32 `json:"magic,omitempty"`
	CPUType    uint32 `json:"cputype,omitempty"`
	CPUSubtype uint32 `json:"cpusubtype,omitempty"`
	FileType   uint32 `json:"filetype,omitempty"`
	Ncmds      uint32 `json:"ncmds,omitempty"`
	Sizeofcmds uint32 `json:"sizeofcmds,omitempty"`
	Flags      uint32 `json:"flags,omitempty"`
}

type dylib struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
	Weak bool   `json:"weak,omitempty"`
}

type buildVersion struct {
	Platform uint32      `json:"platform"`
	MinOS    uint32      `json:"minos"`
	SDK      uint32      `json:"sdk"`
	Tools    []buildTool `json:"tools,omitempty"`
}

type buildTool struct {
	Tool    uint32 `json:"tool"`
	Version uint32 `json:"version"`
}

type linkeditData struct {
	Offset uint32 `json:"offset"`
	Size   uint32 `json:"size"`
	SHA256 string `json:"sha256,omitempty"`
	Data   string `json:"data,omitempty"`
	Bytes  []byte `json:"bytes,omitempty"`
}

type linkerOption struct {
	Args []string `json:"args"`
}

type section struct {
	Segment       string          `json:"segment"`
	Name          string          `json:"name"`
	Addr          uint64          `json:"addr"`
	Size          uint64          `json:"size"`
	Offset        uint32          `json:"offset,omitempty"`
	Align         uint32          `json:"align,omitempty"`
	Reloff        uint32          `json:"reloff,omitempty"`
	Nreloc        uint32          `json:"nreloc,omitempty"`
	Flags         uint32          `json:"flags,omitempty"`
	Res1          uint32          `json:"reserved1,omitempty"`
	Res2          uint32          `json:"reserved2,omitempty"`
	Header        bool            `json:"header,omitempty"`
	SHA256        string          `json:"sha256"`
	Data          string          `json:"data,omitempty"`
	Strings       []strlit        `json:"strings,omitempty"`
	ByteSpans     []byteSpan      `json:"byteSpans,omitempty"`
	Words         []uint32        `json:"words,omitempty"`
	Bytes         []byte          `json:"bytes,omitempty"`
	CFStrings     []cfstring      `json:"cfstrings,omitempty"`
	CompactUnwind []compactUnwind `json:"compactUnwind,omitempty"`
	Relocs        []reloc         `json:"relocs,omitempty"`
}

type strlit struct {
	Offset uint32 `json:"offset"`
	Value  string `json:"value"`
}

type byteSpan struct {
	Offset uint32 `json:"offset"`
	Hex    string `json:"hex"`
}

type cfstring struct {
	Offset uint32 `json:"offset"`
	Flags  uint64 `json:"flags"`
	Length uint64 `json:"length"`
}

type compactUnwind struct {
	Offset uint32 `json:"offset"`
	Info   uint64 `json:"info"`
}

type reloc struct {
	Addr      uint32 `json:"addr"`
	Value     uint32 `json:"value"`
	Symbol    string `json:"symbol,omitempty"`
	Type      uint8  `json:"type"`
	Len       uint8  `json:"len"`
	Pcrel     bool   `json:"pcrel,omitempty"`
	Extern    bool   `json:"extern,omitempty"`
	Scattered bool   `json:"scattered,omitempty"`
}

type symbol struct {
	Name    string `json:"name"`
	RawName string `json:"rawName,omitempty"`
	Value   uint64 `json:"value"`
	Sect    uint8  `json:"sect,omitempty"`
	Type    uint8  `json:"type,omitempty"`
	Desc    uint16 `json:"desc,omitempty"`
	Entry   bool   `json:"entry,omitempty"`
}

func main() {
	checkPath := flag.String("check", "", "check against a metadata manifest")
	checkDataDir := flag.String("check-data-dir", "", "check extracted section payloads against a manifest")
	extractDir := flag.String("extract-dir", "", "write filtered section payloads to dir")
	includeRelocs := flag.Bool("relocs", false, "include section relocation records")
	includeAll := flag.Bool("all", false, "include all sections and symbols")
	ignoreLayout := flag.Bool("ignore-layout", false, "ignore file-layout offsets when checking a manifest")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: dump_swift_metadata [-all] [-relocs] [-ignore-layout] [-check manifest.json] [-check-data-dir dir] [-extract-dir dir] macho")
		os.Exit(2)
	}
	path := flag.Arg(0)
	f, err := macho.Open(path)
	check(err)
	defer f.Close()

	m := manifest{Path: path}
	info, err := loadInfo(path)
	check(err)
	m.Header = info.Header
	m.BuildVersion = info.BuildVersion
	m.DataInCode = info.DataInCode
	m.LinkerOptimizationHint = info.LinkerOptimizationHint
	m.LinkerOptions = info.LinkerOptions
	if *extractDir != "" {
		source, err := os.ReadFile(path)
		check(err)
		writeLinkeditData(*extractDir, source, "LC_DATA_IN_CODE", m.DataInCode)
		writeLinkeditData(*extractDir, source, "LC_LINKER_OPTIMIZATION_HINT", m.LinkerOptimizationHint)
	}
	for _, lib := range info.Dylibs {
		if isMetadataDylib(lib.Name) {
			m.Dylibs = append(m.Dylibs, lib)
		}
	}
	for _, sec := range f.Sections {
		if *includeAll || isMetadataSection(sec) || strings.HasPrefix(sec.Name, "__objc_") || sec.Name == "__constg_swiftt" {
			data, err := sectionData(sec)
			check(err)
			out := section{
				Segment: sec.Seg,
				Name:    sec.Name,
				Addr:    sec.Addr,
				Size:    sec.Size,
				Offset:  sec.Offset,
				Align:   sec.Align,
				Reloff:  sec.Reloff,
				Nreloc:  sec.Nreloc,
				Flags:   sec.Flags,
				Header:  true,
				SHA256:  hash(data),
			}
			if h, ok := info.Sections[sectionKey(sec.Seg, sec.Name)]; ok {
				out.Res1 = h.Res1
				out.Res2 = h.Res2
			}
			if *includeRelocs {
				out.Relocs = relocs(f, sec.Relocs)
			}
			out.Strings = stringLiterals(sec, data)
			out.ByteSpans = byteSpanLiterals(sec, data)
			out.Words = wordLiterals(sec, data)
			out.Bytes = byteLiterals(sec, data)
			out.CFStrings = cfstringLiterals(sec, data)
			out.CompactUnwind = compactUnwindLiterals(sec, data)
			if *extractDir != "" && !hasStructuredSectionData(out) {
				out.Data = writeSectionData(*extractDir, out, data)
			}
			m.Sections = append(m.Sections, out)
		}
	}
	if f.Symtab != nil {
		for i, sym := range f.Symtab.Syms {
			if *includeAll || isMetadataSymbol(sym.Name) {
				rawName := ""
				if i < len(info.RawSymbolNames) && info.RawSymbolNames[i] != sym.Name {
					rawName = info.RawSymbolNames[i]
				}
				m.Symbols = append(m.Symbols, symbol{
					Name:    sym.Name,
					RawName: rawName,
					Value:   sym.Value,
					Sect:    sym.Sect,
					Type:    sym.Type,
					Desc:    sym.Desc,
					Entry:   true,
				})
			}
		}
	}
	if *checkPath != "" {
		checkManifest(*checkPath, *checkDataDir, *ignoreLayout, m)
		return
	}
	writeManifest(os.Stdout, m)
}

func relocs(f *macho.File, in []macho.Reloc) []reloc {
	out := make([]reloc, 0, len(in))
	for _, r := range in {
		out = append(out, reloc{
			Addr:      r.Addr,
			Value:     r.Value,
			Symbol:    relocSymbol(f, r),
			Type:      r.Type,
			Len:       r.Len,
			Pcrel:     r.Pcrel,
			Extern:    r.Extern,
			Scattered: r.Scattered,
		})
	}
	return out
}

func relocSymbol(f *macho.File, r macho.Reloc) string {
	if !r.Extern || f.Symtab == nil || int(r.Value) >= len(f.Symtab.Syms) {
		return ""
	}
	return f.Symtab.Syms[r.Value].Name
}

func isMetadataDylib(name string) bool {
	return strings.Contains(name, "/usr/lib/swift/") ||
		strings.Contains(name, "ExtensionFoundation.framework") ||
		strings.Contains(name, "FSKit.framework") ||
		strings.Contains(name, "Foundation.framework")
}

type objectInfo struct {
	Header                 header
	Dylibs                 []dylib
	BuildVersion           *buildVersion
	DataInCode             *linkeditData
	LinkerOptimizationHint *linkeditData
	LinkerOptions          []linkerOption
	Sections               map[string]rawSection
	RawSymbolNames         []string
}

type rawSection struct {
	Res1 uint32
	Res2 uint32
}

func loadInfo(path string) (objectInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return objectInfo{}, err
	}
	if len(data) < 32 || binary.LittleEndian.Uint32(data[:4]) != 0xfeedfacf {
		return objectInfo{}, fmt.Errorf("%s is not a little-endian 64-bit Mach-O", path)
	}
	info := objectInfo{
		Header: header{
			Magic:      binary.LittleEndian.Uint32(data[0:4]),
			CPUType:    binary.LittleEndian.Uint32(data[4:8]),
			CPUSubtype: binary.LittleEndian.Uint32(data[8:12]),
			FileType:   binary.LittleEndian.Uint32(data[12:16]),
			Ncmds:      binary.LittleEndian.Uint32(data[16:20]),
			Sizeofcmds: binary.LittleEndian.Uint32(data[20:24]),
			Flags:      binary.LittleEndian.Uint32(data[24:28]),
		},
		Sections: map[string]rawSection{},
	}
	off := 32
	var symoff, nsyms, stroff, strsize uint32
	for i := uint32(0); i < info.Header.Ncmds; i++ {
		if off+8 > len(data) {
			return objectInfo{}, fmt.Errorf("load command %d out of range", i)
		}
		cmd := binary.LittleEndian.Uint32(data[off : off+4])
		size := binary.LittleEndian.Uint32(data[off+4 : off+8])
		if size < 8 || off+int(size) > len(data) {
			return objectInfo{}, fmt.Errorf("load command %d has invalid size", i)
		}
		if isDylibCommand(cmd) {
			if size < 24 {
				return objectInfo{}, fmt.Errorf("dylib load command %d too small", i)
			}
			nameOff := binary.LittleEndian.Uint32(data[off+8 : off+12])
			if nameOff >= size {
				return objectInfo{}, fmt.Errorf("dylib load command %d has invalid name offset", i)
			}
			name := cstring(data[off+int(nameOff) : off+int(size)])
			info.Dylibs = append(info.Dylibs, dylib{Name: name, Cmd: loadCommandName(cmd), Weak: cmd == 0x80000018})
		}
		switch cmd {
		case 0x2: // LC_SYMTAB
			if size < 24 {
				return objectInfo{}, fmt.Errorf("symtab command %d too small", i)
			}
			symoff = binary.LittleEndian.Uint32(data[off+8 : off+12])
			nsyms = binary.LittleEndian.Uint32(data[off+12 : off+16])
			stroff = binary.LittleEndian.Uint32(data[off+16 : off+20])
			strsize = binary.LittleEndian.Uint32(data[off+20 : off+24])
		case 0x19: // LC_SEGMENT_64
			if size < 72 {
				return objectInfo{}, fmt.Errorf("segment command %d too small", i)
			}
			nsects := binary.LittleEndian.Uint32(data[off+64 : off+68])
			if size < 72+nsects*80 {
				return objectInfo{}, fmt.Errorf("segment command %d has invalid section count", i)
			}
			secOff := off + 72
			for j := uint32(0); j < nsects; j++ {
				name := cstring(data[secOff : secOff+16])
				seg := cstring(data[secOff+16 : secOff+32])
				info.Sections[sectionKey(seg, name)] = rawSection{
					Res1: binary.LittleEndian.Uint32(data[secOff+72 : secOff+76]),
					Res2: binary.LittleEndian.Uint32(data[secOff+76 : secOff+80]),
				}
				secOff += 80
			}
		case 0x32: // LC_BUILD_VERSION
			if size < 24 {
				return objectInfo{}, fmt.Errorf("build version command %d too small", i)
			}
			n := binary.LittleEndian.Uint32(data[off+20 : off+24])
			if size < 24+n*8 {
				return objectInfo{}, fmt.Errorf("build version command %d has invalid tool count", i)
			}
			bv := buildVersion{
				Platform: binary.LittleEndian.Uint32(data[off+8 : off+12]),
				MinOS:    binary.LittleEndian.Uint32(data[off+12 : off+16]),
				SDK:      binary.LittleEndian.Uint32(data[off+16 : off+20]),
			}
			toolOff := off + 24
			for j := uint32(0); j < n; j++ {
				bv.Tools = append(bv.Tools, buildTool{
					Tool:    binary.LittleEndian.Uint32(data[toolOff : toolOff+4]),
					Version: binary.LittleEndian.Uint32(data[toolOff+4 : toolOff+8]),
				})
				toolOff += 8
			}
			info.BuildVersion = &bv
		case 0x29: // LC_DATA_IN_CODE
			if size < 16 {
				return objectInfo{}, fmt.Errorf("data-in-code command %d too small", i)
			}
			offset := binary.LittleEndian.Uint32(data[off+8 : off+12])
			dataSize := binary.LittleEndian.Uint32(data[off+12 : off+16])
			sum, err := dataHash(data, offset, dataSize)
			if err != nil {
				return objectInfo{}, err
			}
			info.DataInCode = &linkeditData{
				Offset: offset,
				Size:   dataSize,
				SHA256: sum,
			}
		case 0x2e: // LC_LINKER_OPTIMIZATION_HINT
			if size < 16 {
				return objectInfo{}, fmt.Errorf("linker optimization hint command %d too small", i)
			}
			offset := binary.LittleEndian.Uint32(data[off+8 : off+12])
			dataSize := binary.LittleEndian.Uint32(data[off+12 : off+16])
			sum, err := dataHash(data, offset, dataSize)
			if err != nil {
				return objectInfo{}, err
			}
			info.LinkerOptimizationHint = &linkeditData{
				Offset: offset,
				Size:   dataSize,
				SHA256: sum,
			}
		case 0x2d: // LC_LINKER_OPTION
			if size < 12 {
				return objectInfo{}, fmt.Errorf("linker option command %d too small", i)
			}
			count := binary.LittleEndian.Uint32(data[off+8 : off+12])
			args := cstrings(data[off+12 : off+int(size)])
			if uint32(len(args)) != count {
				return objectInfo{}, fmt.Errorf("linker option command %d has %d strings, want %d", i, len(args), count)
			}
			info.LinkerOptions = append(info.LinkerOptions, linkerOption{Args: args})
		}
		off += int(size)
	}
	names, err := rawSymbolNames(data, symoff, nsyms, stroff, strsize)
	if err != nil {
		return objectInfo{}, err
	}
	info.RawSymbolNames = names
	return info, nil
}

func rawSymbolNames(data []byte, symoff, nsyms, stroff, strsize uint32) ([]string, error) {
	if nsyms == 0 {
		return nil, nil
	}
	if uint64(symoff)+uint64(nsyms)*16 > uint64(len(data)) {
		return nil, fmt.Errorf("symbol table range %#x+%#x out of file range", symoff, nsyms*16)
	}
	if uint64(stroff)+uint64(strsize) > uint64(len(data)) {
		return nil, fmt.Errorf("string table range %#x+%#x out of file range", stroff, strsize)
	}
	names := make([]string, nsyms)
	for i := uint32(0); i < nsyms; i++ {
		strx := binary.LittleEndian.Uint32(data[symoff+i*16 : symoff+i*16+4])
		if strx == 0 {
			continue
		}
		if strx >= strsize {
			return nil, fmt.Errorf("symbol %d has string index %#x outside string table size %#x", i, strx, strsize)
		}
		names[i] = cstring(data[stroff+strx : stroff+strsize])
	}
	return names, nil
}

func dataHash(data []byte, off, size uint32) (string, error) {
	if uint64(off)+uint64(size) > uint64(len(data)) {
		return "", fmt.Errorf("linkedit data range %#x+%#x out of file range", off, size)
	}
	return hash(data[off : off+size]), nil
}

func sectionKey(seg, name string) string {
	return seg + "," + name
}

func isDylibCommand(cmd uint32) bool {
	switch cmd {
	case 0xc, 0x80000018, 0x8000001f, 0x80000023:
		return true
	default:
		return false
	}
}

func loadCommandName(cmd uint32) string {
	switch cmd {
	case 0xc:
		return "LC_LOAD_DYLIB"
	case 0x80000018:
		return "LC_LOAD_WEAK_DYLIB"
	case 0x8000001f:
		return "LC_REEXPORT_DYLIB"
	case 0x80000023:
		return "LC_LOAD_UPWARD_DYLIB"
	default:
		return fmt.Sprintf("LC_%#x", cmd)
	}
}

func cstring(b []byte) string {
	if i := strings.IndexByte(string(b), 0); i >= 0 {
		return string(b[:i])
	}
	return string(b)
}

func cstrings(b []byte) []string {
	var out []string
	for len(b) > 0 {
		i := strings.IndexByte(string(b), 0)
		if i < 0 {
			break
		}
		if i > 0 {
			out = append(out, string(b[:i]))
		}
		b = b[i+1:]
	}
	return out
}

func isMetadataSection(sec *macho.Section) bool {
	return sec.Seg == "__TEXT" && strings.HasPrefix(sec.Name, "__swift")
}

func isMetadataSymbol(name string) bool {
	return strings.Contains(name, "swiftmeta") ||
		strings.Contains(name, "NinePFS") ||
		strings.Contains(name, "NinePFileSystem") ||
		strings.Contains(name, "ExtensionFoundation") ||
		strings.Contains(name, "UnaryFileSystemExtension") ||
		strings.Contains(name, "FSKit")
}

func hash(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func sectionData(sec *macho.Section) ([]byte, error) {
	if sec.Flags&0xff == 1 { // S_ZEROFILL
		return make([]byte, sec.Size), nil
	}
	return sec.Data()
}

func stringLiterals(sec *macho.Section, data []byte) []strlit {
	if sec.Seg == "__TEXT" && sec.Name == "__swift5_typeref" {
		strings, _ := typerefLiterals(data)
		return strings
	}
	if !isStringSection(sec) {
		return nil
	}
	var out []strlit
	for off := 0; off < len(data); {
		for off < len(data) && data[off] == 0 {
			off++
		}
		start := off
		for off < len(data) && data[off] != 0 {
			if data[off] < 0x20 || data[off] > 0x7e {
				return nil
			}
			off++
		}
		if start != off {
			out = append(out, strlit{Offset: uint32(start), Value: string(data[start:off])})
		}
	}
	return out
}

func byteSpanLiterals(sec *macho.Section, data []byte) []byteSpan {
	if sec.Seg != "__TEXT" || sec.Name != "__swift5_typeref" {
		return nil
	}
	_, spans := typerefLiterals(data)
	return spans
}

func typerefLiterals(data []byte) ([]strlit, []byteSpan) {
	covered := make([]bool, len(data))
	var strings []strlit
	for off := 0; off < len(data); {
		for off < len(data) && !isPrintableTyperefByte(data[off]) {
			off++
		}
		start := off
		for off < len(data) && isPrintableTyperefByte(data[off]) {
			off++
		}
		if off < len(data) && data[off] == 0 && off-start >= 2 {
			strings = append(strings, strlit{Offset: uint32(start), Value: string(data[start:off])})
			for i := start; i <= off; i++ {
				covered[i] = true
			}
		}
	}
	var spans []byteSpan
	for off := 0; off < len(data); {
		for off < len(data) && (covered[off] || data[off] == 0) {
			off++
		}
		start := off
		for off < len(data) && !covered[off] && data[off] != 0 {
			off++
		}
		if start != off {
			spans = append(spans, byteSpan{Offset: uint32(start), Hex: hex.EncodeToString(data[start:off])})
		}
	}
	return strings, spans
}

func isPrintableTyperefByte(b byte) bool {
	return b >= 0x20 && b <= 0x7e
}

func isStringSection(sec *macho.Section) bool {
	switch sec.Seg + "," + sec.Name {
	case "__TEXT,__cstring",
		"__TEXT,__objc_classname",
		"__TEXT,__objc_methname",
		"__TEXT,__objc_methtype",
		"__TEXT,__swift5_reflstr":
		return len(sec.Relocs) == 0
	default:
		return false
	}
}

func wordLiterals(sec *macho.Section, data []byte) []uint32 {
	if !isWordSection(sec) || len(data)%4 != 0 {
		return nil
	}
	words := make([]uint32, len(data)/4)
	for i := range words {
		words[i] = binary.LittleEndian.Uint32(data[i*4:])
	}
	return words
}

func isWordSection(sec *macho.Section) bool {
	switch sec.Seg + "," + sec.Name {
	case "__TEXT,__swift5_capture",
		"__TEXT,__text",
		"__TEXT,__swift5_assocty",
		"__TEXT,__swift5_fieldmd",
		"__TEXT,__const",
		"__TEXT,__constg_swiftt",
		"__DATA,__const",
		"__DATA,__objc_const",
		"__DATA,__objc_data",
		"__DATA,__data",
		"__LD,__compact_unwind":
		return len(sec.Relocs) > 0
	default:
		return false
	}
}

func byteLiterals(sec *macho.Section, data []byte) []byte {
	if sec.Seg == "__DATA" && sec.Name == "__objc_imageinfo" {
		return append([]byte(nil), data...)
	}
	return nil
}

func cfstringLiterals(sec *macho.Section, data []byte) []cfstring {
	if sec.Seg != "__DATA" || sec.Name != "__cfstring" || len(data)%32 != 0 {
		return nil
	}
	var out []cfstring
	for off := 0; off < len(data); off += 32 {
		if binary.LittleEndian.Uint64(data[off:]) != 0 || binary.LittleEndian.Uint64(data[off+16:]) != 0 {
			return nil
		}
		out = append(out, cfstring{
			Offset: uint32(off),
			Flags:  binary.LittleEndian.Uint64(data[off+8:]),
			Length: binary.LittleEndian.Uint64(data[off+24:]),
		})
	}
	return out
}

func compactUnwindLiterals(sec *macho.Section, data []byte) []compactUnwind {
	if sec.Seg != "__LD" || sec.Name != "__compact_unwind" || len(data)%32 != 0 {
		return nil
	}
	var out []compactUnwind
	for off := 0; off < len(data); off += 32 {
		if binary.LittleEndian.Uint64(data[off:]) != 0 ||
			binary.LittleEndian.Uint64(data[off+16:]) != 0 ||
			binary.LittleEndian.Uint64(data[off+24:]) != 0 {
			return nil
		}
		out = append(out, compactUnwind{
			Offset: uint32(off),
			Info:   binary.LittleEndian.Uint64(data[off+8:]),
		})
	}
	return out
}

func writeSectionData(dir string, sec section, data []byte) string {
	check(os.MkdirAll(dir, 0755))
	name := safeName(sec.Segment + "." + sec.Name + ".bin")
	path := filepath.Join(dir, name)
	check(os.WriteFile(path, data, 0644))
	return name
}

func hasStructuredSectionData(sec section) bool {
	return len(sec.Strings) > 0 ||
		len(sec.ByteSpans) > 0 ||
		len(sec.Words) > 0 ||
		len(sec.Bytes) > 0 ||
		len(sec.CFStrings) > 0 ||
		len(sec.CompactUnwind) > 0 ||
		isZeroRelocTableManifest(sec) ||
		(sec.Segment == "__TEXT" && sec.Name == "__swift5_entry") ||
		(sec.Segment == "__TEXT" && sec.Name == "__literal8") ||
		(sec.Segment == "__DATA" && sec.Name == "__objc_imageinfo") ||
		sec.Flags&1 != 0
}

func isZeroRelocTableManifest(sec section) bool {
	switch sec.Segment + "," + sec.Name {
	case "__TEXT,__swift5_proto",
		"__TEXT,__swift5_types",
		"__DATA,__objc_classlist",
		"__DATA,__objc_protolist",
		"__DATA,__objc_selrefs",
		"__DATA,__objc_protorefs",
		"__DATA,__objc_classrefs":
		return sec.Size%8 == 0 && len(sec.Relocs) == int(sec.Size/8)
	default:
		return false
	}
}

func writeLinkeditData(dir string, source []byte, name string, data *linkeditData) {
	if data == nil || data.Size == 0 {
		return
	}
	check(os.MkdirAll(dir, 0755))
	file := safeName(name + ".bin")
	path := filepath.Join(dir, file)
	if uint64(data.Offset)+uint64(data.Size) > uint64(len(source)) {
		check(fmt.Errorf("linkedit data range %#x+%#x out of file range", data.Offset, data.Size))
	}
	payload := source[data.Offset : data.Offset+data.Size]
	_ = path
	_ = file
	data.Bytes = append(data.Bytes[:0], payload...)
}

func safeName(name string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return r
		case r >= 'A' && r <= 'Z':
			return r
		case r >= '0' && r <= '9':
			return r
		case r == '.' || r == '_' || r == '-':
			return r
		default:
			return '_'
		}
	}, name)
}

func checkManifest(path, dataDir string, ignoreLayout bool, got manifest) {
	data, err := os.ReadFile(path)
	check(err)
	var want manifest
	check(json.Unmarshal(data, &want))
	failed := false
	if want.Header.Magic != 0 && !sameHeader(got.Header, want.Header, ignoreLayout) {
		fmt.Fprintf(os.Stderr, "header mismatch: got %+v want %+v\n", got.Header, want.Header)
		failed = true
	}
	if want.BuildVersion != nil && !sameBuildVersion(got.BuildVersion, want.BuildVersion) {
		fmt.Fprintf(os.Stderr, "build version mismatch: got %+v want %+v\n", got.BuildVersion, want.BuildVersion)
		failed = true
	}
	if want.DataInCode != nil && !sameLinkeditData(got.DataInCode, want.DataInCode, ignoreLayout) {
		fmt.Fprintf(os.Stderr, "data-in-code mismatch: got %+v want %+v\n", got.DataInCode, want.DataInCode)
		failed = true
	}
	if want.LinkerOptimizationHint != nil && !sameLinkeditData(got.LinkerOptimizationHint, want.LinkerOptimizationHint, ignoreLayout) {
		fmt.Fprintf(os.Stderr, "linker optimization hint mismatch: got %+v want %+v\n", got.LinkerOptimizationHint, want.LinkerOptimizationHint)
		failed = true
	}
	if len(want.LinkerOptions) > 0 && !sameLinkerOptions(got.LinkerOptions, want.LinkerOptions) {
		fmt.Fprintf(os.Stderr, "linker options mismatch: got %d want %d\n", len(got.LinkerOptions), len(want.LinkerOptions))
		failed = true
	}
	if dataDir != "" && !checkDataFiles(dataDir, want) {
		failed = true
	}
	wantSections := map[string]section{}
	wantSectionCount := map[string]int{}
	for _, sec := range want.Sections {
		name := sec.Segment + "," + sec.Name
		wantSections[name] = sec
		wantSectionCount[name]++
	}
	gotSections := map[string]section{}
	gotSectionCount := map[string]int{}
	for _, sec := range got.Sections {
		name := sec.Segment + "," + sec.Name
		gotSections[name] = sec
		gotSectionCount[name]++
	}
	for name, want := range wantSections {
		got, ok := gotSections[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "missing section %s\n", name)
			failed = true
			continue
		}
		relocMismatch := got.Reloff != want.Reloff && !ignoreLayout
		headerMismatch := want.Header && (relocMismatch || got.Nreloc != want.Nreloc ||
			got.Flags != want.Flags || got.Res1 != want.Res1 || got.Res2 != want.Res2)
		if got.Size != want.Size || got.Align != want.Align || headerMismatch || got.SHA256 != want.SHA256 {
			fmt.Fprintf(os.Stderr, "section %s mismatch: size %#x/%#x align %d/%d reloff %#x/%#x nreloc %d/%d flags %#x/%#x reserved1 %d/%d reserved2 %d/%d sha256 %s/%s\n",
				name, got.Size, want.Size, got.Align, want.Align, got.Reloff, want.Reloff,
				got.Nreloc, want.Nreloc, got.Flags, want.Flags, got.Res1, want.Res1,
				got.Res2, want.Res2, got.SHA256, want.SHA256)
			failed = true
		}
		if len(want.Relocs) > 0 || len(got.Relocs) > 0 {
			if !sameRelocs(got.Relocs, want.Relocs) {
				fmt.Fprintf(os.Stderr, "section %s relocations mismatch: %d/%d\n",
					name, len(got.Relocs), len(want.Relocs))
				failed = true
			}
		}
		if gotSectionCount[name] != wantSectionCount[name] {
			fmt.Fprintf(os.Stderr, "section %s count mismatch: %d/%d\n", name, gotSectionCount[name], wantSectionCount[name])
			failed = true
		}
	}
	for name := range gotSections {
		if _, ok := wantSections[name]; !ok {
			fmt.Fprintf(os.Stderr, "unexpected section %s\n", name)
			failed = true
		}
	}
	wantSymbols := map[string]bool{}
	wantSymbolCount := map[string]int{}
	for _, sym := range want.Symbols {
		wantSymbols[sym.Name] = true
		wantSymbolCount[sym.Name]++
	}
	gotSymbols := map[string]bool{}
	gotSymbolCount := map[string]int{}
	for _, sym := range got.Symbols {
		gotSymbols[sym.Name] = true
		gotSymbolCount[sym.Name]++
	}
	for name := range wantSymbols {
		if !gotSymbols[name] {
			fmt.Fprintf(os.Stderr, "missing symbol %s\n", name)
			failed = true
		}
		if gotSymbols[name] && !sameSymbolEntries(symbolEntries(got.Symbols, name), symbolEntries(want.Symbols, name)) {
			fmt.Fprintf(os.Stderr, "symbol %s entries mismatch\n", name)
			failed = true
		}
		if gotSymbolCount[name] != wantSymbolCount[name] {
			fmt.Fprintf(os.Stderr, "symbol %s count mismatch: %d/%d\n", name, gotSymbolCount[name], wantSymbolCount[name])
			failed = true
		}
	}
	for name := range gotSymbols {
		if !wantSymbols[name] {
			fmt.Fprintf(os.Stderr, "unexpected symbol %s\n", name)
			failed = true
		}
	}
	wantDylibs := map[string]dylib{}
	wantDylibCount := map[string]int{}
	for _, lib := range want.Dylibs {
		wantDylibs[lib.Name] = lib
		wantDylibCount[lib.Name]++
	}
	gotDylibs := map[string]dylib{}
	gotDylibCount := map[string]int{}
	for _, lib := range got.Dylibs {
		gotDylibs[lib.Name] = lib
		gotDylibCount[lib.Name]++
	}
	for name, want := range wantDylibs {
		got, ok := gotDylibs[name]
		if !ok {
			fmt.Fprintf(os.Stderr, "missing dylib %s\n", name)
			failed = true
			continue
		}
		if got.Cmd != want.Cmd || got.Weak != want.Weak {
			fmt.Fprintf(os.Stderr, "dylib %s mismatch: cmd %s/%s weak %v/%v\n",
				name, got.Cmd, want.Cmd, got.Weak, want.Weak)
			failed = true
		}
		if gotDylibCount[name] != wantDylibCount[name] {
			fmt.Fprintf(os.Stderr, "dylib %s count mismatch: %d/%d\n", name, gotDylibCount[name], wantDylibCount[name])
			failed = true
		}
	}
	for name := range gotDylibs {
		if _, ok := wantDylibs[name]; !ok {
			fmt.Fprintf(os.Stderr, "unexpected dylib %s\n", name)
			failed = true
		}
	}
	if failed {
		os.Exit(1)
	}
	fmt.Printf("metadata manifest ok: sections=%d symbols=%d dylibs=%d\n", len(want.Sections), len(want.Symbols), len(want.Dylibs))
}

func sameHeader(a, b header, ignoreLayout bool) bool {
	if a.Magic != b.Magic || a.CPUType != b.CPUType || a.CPUSubtype != b.CPUSubtype ||
		a.FileType != b.FileType || a.Flags != b.Flags {
		return false
	}
	if ignoreLayout {
		return true
	}
	return a.Ncmds == b.Ncmds && a.Sizeofcmds == b.Sizeofcmds
}

func checkDataFiles(dir string, m manifest) bool {
	ok := true
	if !checkLinkeditDataFile(dir, "LC_DATA_IN_CODE", m.DataInCode) {
		ok = false
	}
	if !checkLinkeditDataFile(dir, "LC_LINKER_OPTIMIZATION_HINT", m.LinkerOptimizationHint) {
		ok = false
	}
	for _, sec := range m.Sections {
		if sec.Data == "" {
			continue
		}
		path := filepath.Join(dir, sec.Data)
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "section %s,%s data: %v\n", sec.Segment, sec.Name, err)
			ok = false
			continue
		}
		if uint64(len(data)) != sec.Size || hash(data) != sec.SHA256 {
			fmt.Fprintf(os.Stderr, "section %s,%s data mismatch: size %#x/%#x sha256 %s/%s\n",
				sec.Segment, sec.Name, len(data), sec.Size, hash(data), sec.SHA256)
			ok = false
		}
	}
	return ok
}

func checkLinkeditDataFile(dir, name string, d *linkeditData) bool {
	if d == nil || d.Data == "" {
		return true
	}
	path := filepath.Join(dir, d.Data)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s data: %v\n", name, err)
		return false
	}
	if uint32(len(data)) != d.Size || hash(data) != d.SHA256 {
		fmt.Fprintf(os.Stderr, "%s data mismatch: size %#x/%#x sha256 %s/%s\n",
			name, len(data), d.Size, hash(data), d.SHA256)
		return false
	}
	return true
}

func sameRelocs(a, b []reloc) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sameBuildVersion(a, b *buildVersion) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Platform != b.Platform || a.MinOS != b.MinOS || a.SDK != b.SDK || len(a.Tools) != len(b.Tools) {
		return false
	}
	for i := range a.Tools {
		if a.Tools[i] != b.Tools[i] {
			return false
		}
	}
	return true
}

func sameLinkeditData(a, b *linkeditData, ignoreLayout bool) bool {
	if a == nil || b == nil {
		return a == b
	}
	if !ignoreLayout && a.Offset != b.Offset {
		return false
	}
	if a.Size != b.Size {
		return false
	}
	if len(a.Bytes) != 0 && len(b.Bytes) != 0 && !bytes.Equal(a.Bytes, b.Bytes) {
		return false
	}
	return b.SHA256 == "" || a.SHA256 == b.SHA256
}

func sameLinkerOptions(a, b []linkerOption) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i].Args) != len(b[i].Args) {
			return false
		}
		for j := range a[i].Args {
			if a[i].Args[j] != b[i].Args[j] {
				return false
			}
		}
	}
	return true
}

func symbolEntries(syms []symbol, name string) []symbol {
	var out []symbol
	for _, sym := range syms {
		if sym.Name == name {
			out = append(out, sym)
		}
	}
	return out
}

func sameSymbolEntries(a, b []symbol) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if b[i].Entry {
			if !sameSymbolEntry(a[i], b[i]) {
				return false
			}
			continue
		}
		if a[i].Name != b[i].Name || a[i].Value != b[i].Value || a[i].Sect != b[i].Sect || a[i].Type != b[i].Type {
			return false
		}
	}
	return true
}

func sameSymbolEntry(got, want symbol) bool {
	if got.Name != want.Name || got.Value != want.Value || got.Sect != want.Sect ||
		got.Type != want.Type || got.Desc != want.Desc || got.Entry != want.Entry {
		return false
	}
	return want.RawName == "" || got.RawName == want.RawName
}

func writeManifest(f *os.File, m manifest) {
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	check(enc.Encode(m))
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
