//go:build ignore

package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

type manifest struct {
	Path                   string         `json:"path,omitempty"`
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
	Offset uint32 `json:"offset,omitempty"`
	Size   uint32 `json:"size,omitempty"`
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

type stringList []string

func (l *stringList) String() string {
	return strings.Join(*l, ",")
}

func (l *stringList) Set(s string) error {
	if s == "" {
		return fmt.Errorf("empty value")
	}
	*l = append(*l, s)
	return nil
}

func main() {
	out := flag.String("dir", "", "directory to write the Swift metadata fixture")
	compact := flag.Bool("compact", false, "remove payload files already represented in the manifest")
	stripUnwind := flag.Bool("strip-unwind", false, "remove compact unwind and linker optimization hint records")
	dropSectionSymbols := flag.Bool("drop-section-symbols", false, "drop symbols that belong to dropped sections")
	fixtureOut := flag.String("fixture", "", "write a base64 gzip tar fixture from dir")
	var drop stringList
	flag.Var(&drop, "drop-section", "remove a section by segment,name and remap symbol section indexes")
	flag.Parse()
	if *out == "" || flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "usage: write_swiftmeta_fixture -dir dir [-compact] [-strip-unwind] [-drop-section segment,name] [-drop-section-symbols] [-fixture file]")
		os.Exit(2)
	}
	_, source, _, ok := runtime.Caller(0)
	if !ok {
		check(fmt.Errorf("cannot locate source directory"))
	}
	loadFixture(filepath.Dir(source), *out)
	if *compact {
		compactFixture(*out)
	}
	if *stripUnwind {
		stripUnwindRecords(*out, *dropSectionSymbols)
	}
	if len(drop) > 0 {
		dropManifestSections(*out, drop, *dropSectionSymbols)
	}
	if *fixtureOut != "" {
		writeFixtureArchive(*out, *fixtureOut)
	}
}

func loadFixture(sourceDir, out string) {
	path := filepath.Join(sourceDir, "swiftmeta_manifest.json")
	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		check(err)
		var m manifest
		check(json.Unmarshal(data, &m))
		m.Path = ""
		encoded, err := json.MarshalIndent(m, "", "\t")
		check(err)
		encoded = append(encoded, '\n')
		check(os.MkdirAll(filepath.Join(out, "sections"), 0755))
		check(os.WriteFile(filepath.Join(out, "manifest.json"), encoded, 0644))
		return
	} else if !os.IsNotExist(err) {
		check(err)
	}

	data, err := os.ReadFile(filepath.Join(sourceDir, "swiftmeta_fixture.tgz.b64"))
	check(err)
	raw, err := base64.StdEncoding.DecodeString(strings.Join(strings.Fields(string(data)), ""))
	check(err)
	zr, err := gzip.NewReader(bytes.NewReader(raw))
	check(err)
	defer zr.Close()
	tr := tar.NewReader(zr)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		check(err)
		if h.Name == "" || filepath.IsAbs(h.Name) || strings.Contains(h.Name, "..") {
			check(fmt.Errorf("unsafe fixture path %q", h.Name))
		}
		path := filepath.Join(out, h.Name)
		switch h.Typeflag {
		case tar.TypeDir:
			check(os.MkdirAll(path, 0755))
		case tar.TypeReg:
			check(os.MkdirAll(filepath.Dir(path), 0755))
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			check(err)
			_, copyErr := io.Copy(f, tr)
			closeErr := f.Close()
			check(copyErr)
			check(closeErr)
		default:
			check(fmt.Errorf("unsupported fixture entry %q type %d", h.Name, h.Typeflag))
		}
	}
}

func compactFixture(dir string) {
	path := filepath.Join(dir, "manifest.json")
	data, err := os.ReadFile(path)
	check(err)
	var m manifest
	check(json.Unmarshal(data, &m))
	if m.LinkerOptimizationHint != nil && m.LinkerOptimizationHint.Data != "" && len(m.LinkerOptimizationHint.Bytes) > 0 {
		removeFixtureFile(dir, m.LinkerOptimizationHint.Data)
		m.LinkerOptimizationHint.Data = ""
	}
	for i := range m.Sections {
		sec := &m.Sections[i]
		if sec.Segment == "__TEXT" && sec.Name == "__swift5_typeref" && len(sec.Bytes) > 0 {
			sec.Strings, sec.ByteSpans = typerefLiterals(sec.Bytes)
			sec.Bytes = nil
		}
		if isWordSection(*sec) && len(sec.Bytes) > 0 {
			sec.Words = bytesToWords(*sec)
			sec.Bytes = nil
		}
		if sec.Segment == "__DATA" && sec.Name == "__objc_imageinfo" && len(sec.Bytes) > 0 {
			sec.Bytes = nil
		}
		if sec.Data == "" || !hasStructuredSectionData(*sec) {
			continue
		}
		removeFixtureFile(dir, sec.Data)
		sec.Data = ""
	}
	if m.DataInCode != nil && m.DataInCode.Size == 0 {
		m.DataInCode = nil
	}
	m.Header.Ncmds = 0
	m.Header.Sizeofcmds = 0
	for i := range m.Sections {
		sec := &m.Sections[i]
		sec.Offset = 0
		sec.Reloff = 0
		sec.Nreloc = 0
		sec.Header = false
	}
	compactSymbols(&m)
	m.LinkerOptions = compactLinkerOptions(m.LinkerOptions)
	out, err := json.MarshalIndent(m, "", "\t")
	check(err)
	out = append(out, '\n')
	check(os.WriteFile(path, out, 0644))
}

func stripUnwindRecords(dir string, dropSectionSymbols bool) {
	dropManifestSections(dir, []string{"__LD,__compact_unwind"}, dropSectionSymbols)
	path := filepath.Join(dir, "manifest.json")
	data, err := os.ReadFile(path)
	check(err)
	var m manifest
	check(json.Unmarshal(data, &m))
	m.LinkerOptimizationHint = nil
	out, err := json.MarshalIndent(m, "", "\t")
	check(err)
	out = append(out, '\n')
	check(os.WriteFile(path, out, 0644))
}

func dropManifestSections(dir string, names []string, dropSectionSymbols bool) {
	path := filepath.Join(dir, "manifest.json")
	data, err := os.ReadFile(path)
	check(err)
	var m manifest
	check(json.Unmarshal(data, &m))
	drop := map[string]bool{}
	for _, name := range names {
		parts := strings.Split(name, ",")
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			check(fmt.Errorf("section name must be segment,name: %q", name))
		}
		drop[name] = true
	}
	remap := make([]uint8, len(m.Sections)+1)
	outSections := m.Sections[:0]
	for i, sec := range m.Sections {
		old := i + 1
		if drop[sec.Segment+","+sec.Name] {
			if sec.Data != "" {
				removeFixtureFile(dir, sec.Data)
			}
			continue
		}
		outSections = append(outSections, sec)
		remap[old] = uint8(len(outSections))
	}
	droppedSymbols := map[string]bool{}
	symRemap := make([]uint32, len(m.Symbols))
	outSymbols := m.Symbols[:0]
	for i, sym := range m.Symbols {
		sect := sym.Sect
		if sect == 0 {
			symRemap[i] = uint32(len(outSymbols))
			outSymbols = append(outSymbols, sym)
			continue
		}
		if int(sect) >= len(remap) || remap[sect] == 0 {
			if !dropSectionSymbols {
				check(fmt.Errorf("symbol %s refers to removed section %d", sym.Name, sect))
			}
			droppedSymbols[sym.Name] = true
			if sym.RawName != "" {
				droppedSymbols[sym.RawName] = true
			}
			continue
		}
		sym.Sect = remap[sect]
		symRemap[i] = uint32(len(outSymbols))
		outSymbols = append(outSymbols, sym)
	}
	if len(droppedSymbols) > 0 {
		for _, sec := range outSections {
			for _, r := range sec.Relocs {
				if droppedSymbols[r.Symbol] {
					check(fmt.Errorf("section %s,%s relocation references dropped symbol %s", sec.Segment, sec.Name, r.Symbol))
				}
			}
		}
	}
	for i := range outSections {
		for j := range outSections[i].Relocs {
			r := &outSections[i].Relocs[j]
			if r.Scattered {
				continue
			}
			if r.Extern {
				if int(r.Value) >= len(symRemap) {
					check(fmt.Errorf("section %s,%s relocation references missing symbol index %d", outSections[i].Segment, outSections[i].Name, r.Value))
				}
				r.Value = symRemap[r.Value]
				continue
			}
			if r.Value == 0 {
				continue
			}
			if int(r.Value) >= len(remap) || remap[r.Value] == 0 {
				check(fmt.Errorf("section %s,%s relocation references removed section %d", outSections[i].Segment, outSections[i].Name, r.Value))
			}
			r.Value = uint32(remap[r.Value])
		}
	}
	m.Sections = outSections
	m.Symbols = outSymbols
	out, err := json.MarshalIndent(m, "", "\t")
	check(err)
	out = append(out, '\n')
	check(os.WriteFile(path, out, 0644))
}

func removeFixtureFile(dir, name string) {
	if filepath.IsAbs(name) || strings.Contains(name, "..") {
		check(fmt.Errorf("unsafe fixture data path %q", name))
	}
	err := os.Remove(filepath.Join(dir, "sections", name))
	if err != nil && !os.IsNotExist(err) {
		check(err)
	}
}

func hasStructuredSectionData(sec section) bool {
	return len(sec.Strings) > 0 ||
		len(sec.ByteSpans) > 0 ||
		len(sec.Words) > 0 ||
		len(sec.Bytes) > 0 ||
		len(sec.CFStrings) > 0 ||
		len(sec.CompactUnwind) > 0 ||
		isZeroRelocTable(sec) ||
		(sec.Segment == "__TEXT" && sec.Name == "__swift5_entry") ||
		(sec.Segment == "__TEXT" && sec.Name == "__literal8") ||
		(sec.Segment == "__DATA" && sec.Name == "__objc_imageinfo") ||
		sec.Flags&1 != 0
}

func isZeroRelocTable(sec section) bool {
	switch sec.Segment + "," + sec.Name {
	case "__TEXT,__swift5_proto",
		"__TEXT,__swift5_types",
		"__DATA,__objc_classlist",
		"__DATA,__objc_protolist",
		"__DATA,__objc_selrefs",
		"__DATA,__objc_protorefs",
		"__DATA,__objc_classrefs":
		return sec.Size%8 == 0
	default:
		return false
	}
}

func isWordSection(sec section) bool {
	switch sec.Segment + "," + sec.Name {
	case "__TEXT,__text",
		"__TEXT,__swift5_capture",
		"__TEXT,__swift5_assocty",
		"__TEXT,__swift5_fieldmd",
		"__TEXT,__const",
		"__TEXT,__constg_swiftt",
		"__DATA,__const",
		"__DATA,__objc_const",
		"__DATA,__objc_data",
		"__DATA,__data",
		"__LD,__compact_unwind":
		return sec.Size%4 == 0
	default:
		return false
	}
}

func bytesToWords(sec section) []uint32 {
	if uint64(len(sec.Bytes)) != sec.Size || len(sec.Bytes)%4 != 0 {
		check(fmt.Errorf("section %s,%s byte count %#x cannot convert to %#x bytes of words", sec.Segment, sec.Name, len(sec.Bytes), sec.Size))
	}
	words := make([]uint32, len(sec.Bytes)/4)
	for i := range words {
		words[i] = binary.LittleEndian.Uint32(sec.Bytes[i*4:])
	}
	return words
}

func compactLinkerOptions(options []linkerOption) []linkerOption {
	out := options[:0]
	for _, opt := range options {
		if keepLinkerOption(opt) {
			out = append(out, opt)
		}
	}
	return out
}

func compactSymbols(m *manifest) {
	refs := map[string]bool{}
	for _, sec := range m.Sections {
		for _, r := range sec.Relocs {
			if r.Symbol != "" {
				refs[r.Symbol] = true
			}
		}
	}
	keep := make([]bool, len(m.Symbols))
	for i, sym := range m.Symbols {
		if refs[sym.Name] || (sym.RawName != "" && refs[sym.RawName]) ||
			sym.Sect == 0 ||
			isRequiredMetadataSymbol(sym.Name) ||
			isPublicMetadataSymbol(sym.Name) {
			keep[i] = true
		}
	}
	remap := make([]uint32, len(m.Symbols))
	out := m.Symbols[:0]
	for i, sym := range m.Symbols {
		if !keep[i] {
			continue
		}
		remap[i] = uint32(len(out))
		out = append(out, sym)
	}
	for i := range m.Sections {
		for j := range m.Sections[i].Relocs {
			r := &m.Sections[i].Relocs[j]
			if !r.Extern {
				continue
			}
			if int(r.Value) >= len(keep) || !keep[r.Value] {
				check(fmt.Errorf("section %s,%s relocation references pruned symbol index %d", m.Sections[i].Segment, m.Sections[i].Name, r.Value))
			}
			r.Value = remap[r.Value]
		}
	}
	m.Symbols = out
}

func isPublicMetadataSymbol(name string) bool {
	return strings.HasPrefix(name, "_OBJC_CLASS_$_") ||
		strings.HasPrefix(name, "_OBJC_METACLASS_$_") ||
		strings.HasPrefix(name, "__OBJC_PROTOCOL_$_") ||
		strings.HasPrefix(name, "__OBJC_METACLASS_$_") ||
		strings.HasPrefix(name, "__OBJC_CLASS_$_") ||
		strings.HasPrefix(name, "_NinePFS") ||
		strings.HasPrefix(name, "_symbolic") ||
		strings.HasPrefix(name, "_associated conformance")
}

func isRequiredMetadataSymbol(name string) bool {
	switch name {
	case "_$s9swiftmeta16NinePFSExtensionV19ExtensionFoundation03AppD0AAMc",
		"_$s9swiftmeta16NinePFSExtensionV5FSKit24UnaryFileSystemExtensionAAMc",
		"_$s9swiftmeta16NinePFSExtensionVAC19ExtensionFoundation03AppD0AAWL",
		"_$s9swiftmeta16NinePFSExtensionVAC5FSKit24UnaryFileSystemExtensionAAWL",
		"__swift_FORCE_LOAD_$_swiftCoreFoundation_$_swiftmeta",
		"__swift_FORCE_LOAD_$_swiftDispatch_$_swiftmeta",
		"__swift_FORCE_LOAD_$_swiftIOKit_$_swiftmeta",
		"__swift_FORCE_LOAD_$_swiftOSLog_$_swiftmeta",
		"__swift_FORCE_LOAD_$_swiftObjectiveC_$_swiftmeta",
		"__swift_FORCE_LOAD_$_swiftXPC_$_swiftmeta",
		"__swift_FORCE_LOAD_$_swift_Builtin_float_$_swiftmeta",
		"__swift_FORCE_LOAD_$_swiftos_$_swiftmeta":
		return true
	default:
		return false
	}
}

func keepLinkerOption(opt linkerOption) bool {
	if len(opt.Args) != 1 {
		return false
	}
	switch opt.Args[0] {
	case "-lswiftCoreFoundation",
		"-lswiftDispatch",
		"-lswiftIOKit",
		"-lswiftOSLog",
		"-lswiftObjectiveC",
		"-lswiftXPC",
		"-lswift_Builtin_float",
		"-lswiftos":
		return true
	default:
		return false
	}
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

func writeFixtureArchive(dir, path string) {
	var files []string
	check(filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		files = append(files, filepath.ToSlash(rel))
		return nil
	}))
	sort.Strings(files)

	var raw bytes.Buffer
	gz := gzip.NewWriter(&raw)
	tw := tar.NewWriter(gz)
	for _, name := range files {
		info, err := os.Stat(filepath.Join(dir, filepath.FromSlash(name)))
		check(err)
		h := &tar.Header{
			Name: name,
			Mode: 0644,
			Size: info.Size(),
		}
		if info.IsDir() {
			h.Typeflag = tar.TypeDir
			h.Mode = 0755
			h.Size = 0
			check(tw.WriteHeader(h))
			continue
		}
		h.Typeflag = tar.TypeReg
		check(tw.WriteHeader(h))
		f, err := os.Open(filepath.Join(dir, filepath.FromSlash(name)))
		check(err)
		_, copyErr := io.Copy(tw, f)
		closeErr := f.Close()
		check(copyErr)
		check(closeErr)
	}
	check(tw.Close())
	check(gz.Close())

	enc := base64.StdEncoding.EncodeToString(raw.Bytes())
	var out strings.Builder
	for len(enc) > 0 {
		n := 76
		if len(enc) < n {
			n = len(enc)
		}
		out.WriteString(enc[:n])
		out.WriteByte('\n')
		enc = enc[n:]
	}
	check(os.WriteFile(path, []byte(out.String()), 0644))
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
