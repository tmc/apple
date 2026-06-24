//go:build ignore

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type manifest struct {
	Header                 header         `json:"header"`
	BuildVersion           *buildVersion  `json:"buildVersion,omitempty"`
	DataInCode             *linkeditData  `json:"dataInCode,omitempty"`
	LinkerOptimizationHint *linkeditData  `json:"linkerOptimizationHint,omitempty"`
	LinkerOptions          []linkerOption `json:"linkerOptions,omitempty"`
	Sections               []section      `json:"sections"`
	Symbols                []symbol       `json:"symbols"`
}

type header struct {
	Magic      uint32 `json:"magic"`
	CPUType    uint32 `json:"cputype"`
	CPUSubtype uint32 `json:"cpusubtype"`
	FileType   uint32 `json:"filetype"`
	Flags      uint32 `json:"flags"`
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
	Size  uint32 `json:"size"`
	Data  string `json:"data,omitempty"`
	Bytes []byte `json:"bytes,omitempty"`
}

type linkerOption struct {
	Args []string `json:"args"`
}

type section struct {
	Segment       string          `json:"segment"`
	Name          string          `json:"name"`
	Addr          uint64          `json:"addr"`
	Size          uint64          `json:"size"`
	Align         uint32          `json:"align,omitempty"`
	Flags         uint32          `json:"flags,omitempty"`
	Res1          uint32          `json:"reserved1,omitempty"`
	Res2          uint32          `json:"reserved2,omitempty"`
	Data          string          `json:"data,omitempty"`
	Strings       []strlit        `json:"strings,omitempty"`
	ByteSpans     []byteSpan      `json:"byteSpans,omitempty"`
	Words         []uint32        `json:"words,omitempty"`
	Bytes         []byte          `json:"bytes,omitempty"`
	CFStrings     []cfstring      `json:"cfstrings,omitempty"`
	CompactUnwind []compactUnwind `json:"compactUnwind,omitempty"`
	Relocs        []reloc         `json:"relocs,omitempty"`

	offset uint32
	reloff uint32
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
}

func main() {
	manifestPath := flag.String("manifest", "", "metadata manifest produced by dump_swift_metadata -all -relocs")
	dataDir := flag.String("data-dir", "", "directory containing extracted section and linkedit payloads")
	outPath := flag.String("o", "", "output Mach-O object")
	flag.Parse()
	if *manifestPath == "" || *dataDir == "" || *outPath == "" || flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "usage: emit_macho_object -manifest manifest.json -data-dir dir -o object.syso")
		os.Exit(2)
	}

	data, err := os.ReadFile(*manifestPath)
	check(err)
	var m manifest
	check(json.Unmarshal(data, &m))
	out, err := emit(m, *dataDir)
	check(err)
	check(os.WriteFile(*outPath, out, 0644))
}

func emit(m manifest, dir string) ([]byte, error) {
	linkOptions := make([][]byte, len(m.LinkerOptions))
	for i, opt := range m.LinkerOptions {
		linkOptions[i] = linkerOptionCommand(opt.Args)
	}

	segCmdSize := uint32(72 + 80*len(m.Sections))
	symCmdSize := uint32(24)
	buildCmdSize := uint32(0)
	if m.BuildVersion != nil {
		buildCmdSize = uint32(24 + 8*len(m.BuildVersion.Tools))
	}
	dataInCodeCmdSize := uint32(0)
	if m.DataInCode != nil {
		dataInCodeCmdSize = 16
	}
	lohCmdSize := uint32(0)
	if m.LinkerOptimizationHint != nil {
		lohCmdSize = 16
	}
	ncmds := uint32(2)
	if m.BuildVersion != nil {
		ncmds++
	}
	if m.DataInCode != nil {
		ncmds++
	}
	ncmds += uint32(len(linkOptions))
	if m.LinkerOptimizationHint != nil {
		ncmds++
	}
	sizeofcmds := segCmdSize + symCmdSize + buildCmdSize + dataInCodeCmdSize + lohCmdSize
	for _, cmd := range linkOptions {
		sizeofcmds += uint32(len(cmd))
	}

	fileoff := uint32(32) + sizeofcmds
	pos := fileoff
	payloads := make([][]byte, len(m.Sections))
	for i := range m.Sections {
		sec := &m.Sections[i]
		if isZerofill(sec.Flags) {
			continue
		}
		pos = align(pos, sec.Align)
		sec.offset = pos
		payload, err := sectionPayload(dir, *sec)
		if err != nil {
			return nil, err
		}
		if uint64(len(payload)) != sec.Size {
			return nil, fmt.Errorf("section %s,%s data size %#x, want %#x", sec.Segment, sec.Name, len(payload), sec.Size)
		}
		payloads[i] = payload
		pos += uint32(len(payload))
	}

	relocStart := align(pos, 3)
	pos = relocStart
	for i := range m.Sections {
		sec := &m.Sections[i]
		if len(sec.Relocs) == 0 {
			continue
		}
		sec.reloff = pos
		pos += uint32(8 * len(sec.Relocs))
	}

	var dataInCodeOff, lohOff uint32
	if m.DataInCode != nil && m.DataInCode.Size != 0 {
		pos = align(pos, 2)
		dataInCodeOff = pos
		pos += m.DataInCode.Size
	}
	if m.LinkerOptimizationHint != nil && m.LinkerOptimizationHint.Size != 0 {
		pos = align(pos, 2)
		lohOff = pos
		pos += m.LinkerOptimizationHint.Size
	}

	symoff := align(pos, 3)
	pos = symoff + uint32(16*len(m.Symbols))
	strtab, strx := stringTable(m.Symbols)
	stroff := pos
	pos += uint32(len(strtab))

	var b bytes.Buffer
	write(&b, m.Header.Magic)
	write(&b, int32(m.Header.CPUType))
	write(&b, int32(m.Header.CPUSubtype))
	write(&b, m.Header.FileType)
	write(&b, ncmds)
	write(&b, sizeofcmds)
	write(&b, m.Header.Flags)
	write(&b, uint32(0))

	writeSegment(&b, segCmdSize, m.Sections, fileoff, relocStart-fileoff)
	write(&b, uint32(0x2))
	write(&b, symCmdSize)
	write(&b, symoff)
	write(&b, uint32(len(m.Symbols)))
	write(&b, stroff)
	write(&b, uint32(len(strtab)))
	if m.BuildVersion != nil {
		writeBuildVersion(&b, buildCmdSize, m.BuildVersion)
	}
	if m.DataInCode != nil {
		writeLinkeditCommand(&b, 0x29, dataInCodeCmdSize, dataInCodeOff, m.DataInCode.Size)
	}
	for _, cmd := range linkOptions {
		b.Write(cmd)
	}
	if m.LinkerOptimizationHint != nil {
		writeLinkeditCommand(&b, 0x2e, lohCmdSize, lohOff, m.LinkerOptimizationHint.Size)
	}
	if uint32(b.Len()) != fileoff {
		return nil, fmt.Errorf("load commands size %#x, want %#x", b.Len(), fileoff)
	}

	buf := b.Bytes()
	for i, sec := range m.Sections {
		if len(payloads[i]) == 0 {
			continue
		}
		buf = padTo(buf, int(sec.offset))
		buf = append(buf, payloads[i]...)
	}
	buf = padTo(buf, int(relocStart))
	for _, sec := range m.Sections {
		for _, r := range sec.Relocs {
			if r.Scattered {
				return nil, fmt.Errorf("section %s,%s has unsupported scattered relocation", sec.Segment, sec.Name)
			}
			writeReloc(&buf, r)
		}
	}
	if m.DataInCode != nil && m.DataInCode.Size != 0 {
		buf = appendLinkedit(buf, dir, dataInCodeOff, m.DataInCode)
	}
	if m.LinkerOptimizationHint != nil && m.LinkerOptimizationHint.Size != 0 {
		buf = appendLinkedit(buf, dir, lohOff, m.LinkerOptimizationHint)
	}
	buf = padTo(buf, int(symoff))
	for i, sym := range m.Symbols {
		writeSymbol(&buf, strx[i], sym)
	}
	buf = padTo(buf, int(stroff))
	buf = append(buf, strtab...)
	return buf[:pos], nil
}

func sectionPayload(dir string, sec section) ([]byte, error) {
	if len(sec.Strings) > 0 || len(sec.ByteSpans) > 0 {
		return stringByteSpanPayload(sec)
	}
	if len(sec.Words) > 0 {
		return wordPayload(sec)
	}
	if len(sec.Bytes) > 0 {
		return bytePayload(sec)
	}
	if len(sec.CFStrings) > 0 {
		return cfstringPayload(sec)
	}
	if len(sec.CompactUnwind) > 0 {
		return compactUnwindPayload(sec)
	}
	if isZeroRelocTable(sec) {
		return make([]byte, sec.Size), nil
	}
	if sec.Flags&1 != 0 {
		return make([]byte, sec.Size), nil
	}
	if sec.Segment == "__TEXT" && sec.Name == "__swift5_entry" {
		var data [8]byte
		binary.LittleEndian.PutUint32(data[4:], 1)
		return data[:], nil
	}
	if sec.Segment == "__TEXT" && sec.Name == "__literal8" {
		var data [8]byte
		binary.LittleEndian.PutUint64(data[:], 0x42000000)
		return data[:], nil
	}
	if sec.Segment == "__DATA" && sec.Name == "__objc_imageinfo" {
		// version=0, flags=Objective-C 2, optimized-by-dyld, Swift stable ABI.
		var data [8]byte
		binary.LittleEndian.PutUint32(data[4:], 0x06030740)
		return data[:], nil
	}
	payload, err := os.ReadFile(filepath.Join(dir, sec.Data))
	if err != nil {
		return nil, fmt.Errorf("read section %s,%s: %w", sec.Segment, sec.Name, err)
	}
	return payload, nil
}

func wordPayload(sec section) ([]byte, error) {
	if uint64(len(sec.Words))*4 != sec.Size {
		return nil, fmt.Errorf("section %s,%s word count %#x, want size %#x", sec.Segment, sec.Name, len(sec.Words), sec.Size)
	}
	data := make([]byte, sec.Size)
	for i, word := range sec.Words {
		binary.LittleEndian.PutUint32(data[i*4:], word)
	}
	return data, nil
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

func stringByteSpanPayload(sec section) ([]byte, error) {
	data := make([]byte, sec.Size)
	for _, s := range sec.Strings {
		if uint64(s.Offset)+uint64(len(s.Value))+1 > sec.Size {
			return nil, fmt.Errorf("section %s,%s string at %#x exceeds section size %#x", sec.Segment, sec.Name, s.Offset, sec.Size)
		}
		copy(data[s.Offset:], s.Value)
		data[int(s.Offset)+len(s.Value)] = 0
	}
	for _, span := range sec.ByteSpans {
		bytes, err := hex.DecodeString(span.Hex)
		if err != nil {
			return nil, fmt.Errorf("section %s,%s byte span at %#x: %w", sec.Segment, sec.Name, span.Offset, err)
		}
		if uint64(span.Offset)+uint64(len(bytes)) > sec.Size {
			return nil, fmt.Errorf("section %s,%s byte span at %#x exceeds section size %#x", sec.Segment, sec.Name, span.Offset, sec.Size)
		}
		copy(data[span.Offset:], bytes)
	}
	return data, nil
}

func bytePayload(sec section) ([]byte, error) {
	if uint64(len(sec.Bytes)) != sec.Size {
		return nil, fmt.Errorf("section %s,%s byte count %#x, want size %#x", sec.Segment, sec.Name, len(sec.Bytes), sec.Size)
	}
	return append([]byte(nil), sec.Bytes...), nil
}

func compactUnwindPayload(sec section) ([]byte, error) {
	data := make([]byte, sec.Size)
	for _, entry := range sec.CompactUnwind {
		if uint64(entry.Offset)+32 > sec.Size {
			return nil, fmt.Errorf("section %s,%s compact unwind at %#x exceeds section size %#x", sec.Segment, sec.Name, entry.Offset, sec.Size)
		}
		binary.LittleEndian.PutUint64(data[entry.Offset+8:], entry.Info)
	}
	return data, nil
}

func cfstringPayload(sec section) ([]byte, error) {
	data := make([]byte, sec.Size)
	for _, s := range sec.CFStrings {
		if uint64(s.Offset)+32 > sec.Size {
			return nil, fmt.Errorf("section %s,%s cfstring at %#x exceeds section size %#x", sec.Segment, sec.Name, s.Offset, sec.Size)
		}
		binary.LittleEndian.PutUint64(data[s.Offset+8:], s.Flags)
		binary.LittleEndian.PutUint64(data[s.Offset+24:], s.Length)
	}
	return data, nil
}

func writeSegment(b *bytes.Buffer, cmdSize uint32, sections []section, fileoff, filesize uint32) {
	write(b, uint32(0x19))
	write(b, cmdSize)
	b.Write(make([]byte, 16))
	write(b, uint64(0))
	var vmsize uint64
	for _, sec := range sections {
		if end := sec.Addr + sec.Size; end > vmsize {
			vmsize = end
		}
	}
	if uint64(filesize) > vmsize {
		vmsize = uint64(filesize)
	}
	write(b, vmsize)
	write(b, uint64(fileoff))
	write(b, uint64(filesize))
	write(b, int32(7))
	write(b, int32(7))
	write(b, uint32(len(sections)))
	write(b, uint32(0))
	for _, sec := range sections {
		writeName(b, sec.Name)
		writeName(b, sec.Segment)
		write(b, sec.Addr)
		write(b, sec.Size)
		write(b, sec.offset)
		write(b, sec.Align)
		write(b, sec.reloff)
		write(b, uint32(len(sec.Relocs)))
		write(b, sec.Flags)
		write(b, sec.Res1)
		write(b, sec.Res2)
		write(b, uint32(0))
	}
}

func writeBuildVersion(b *bytes.Buffer, cmdSize uint32, v *buildVersion) {
	write(b, uint32(0x32))
	write(b, cmdSize)
	write(b, v.Platform)
	write(b, v.MinOS)
	write(b, v.SDK)
	write(b, uint32(len(v.Tools)))
	for _, tool := range v.Tools {
		write(b, tool.Tool)
		write(b, tool.Version)
	}
}

func writeLinkeditCommand(b *bytes.Buffer, cmd, size, off, dataSize uint32) {
	write(b, cmd)
	write(b, size)
	write(b, off)
	write(b, dataSize)
}

func linkerOptionCommand(args []string) []byte {
	var b bytes.Buffer
	write(&b, uint32(0x2d))
	write(&b, uint32(0))
	write(&b, uint32(len(args)))
	for _, arg := range args {
		b.WriteString(arg)
		b.WriteByte(0)
	}
	for b.Len()%8 != 0 {
		b.WriteByte(0)
	}
	out := b.Bytes()
	binary.LittleEndian.PutUint32(out[4:8], uint32(len(out)))
	return out
}

func stringTable(symbols []symbol) ([]byte, []uint32) {
	var b bytes.Buffer
	b.WriteByte(0)
	index := make([]uint32, len(symbols))
	for i, sym := range symbols {
		index[i] = uint32(b.Len())
		name := sym.Name
		if sym.RawName != "" {
			name = sym.RawName
		}
		b.WriteString(name)
		b.WriteByte(0)
	}
	return b.Bytes(), index
}

func writeReloc(buf *[]byte, r reloc) {
	var tmp [8]byte
	binary.LittleEndian.PutUint32(tmp[0:4], r.Addr)
	word := r.Value & 0x00ffffff
	if r.Pcrel {
		word |= 1 << 24
	}
	word |= uint32(r.Len&3) << 25
	if r.Extern {
		word |= 1 << 27
	}
	word |= uint32(r.Type&0xf) << 28
	binary.LittleEndian.PutUint32(tmp[4:8], word)
	*buf = append(*buf, tmp[:]...)
}

func writeSymbol(buf *[]byte, strx uint32, sym symbol) {
	var tmp [16]byte
	binary.LittleEndian.PutUint32(tmp[0:4], strx)
	tmp[4] = sym.Type
	tmp[5] = sym.Sect
	binary.LittleEndian.PutUint16(tmp[6:8], sym.Desc)
	binary.LittleEndian.PutUint64(tmp[8:16], sym.Value)
	*buf = append(*buf, tmp[:]...)
}

func appendLinkedit(buf []byte, dir string, off uint32, d *linkeditData) []byte {
	data := d.Bytes
	if len(data) == 0 {
		var err error
		data, err = os.ReadFile(filepath.Join(dir, d.Data))
		check(err)
	}
	if uint32(len(data)) != d.Size {
		check(fmt.Errorf("linkedit data %s size %#x, want %#x", d.Data, len(data), d.Size))
	}
	buf = padTo(buf, int(off))
	return append(buf, data...)
}

func writeName(b *bytes.Buffer, name string) {
	var buf [16]byte
	copy(buf[:], name)
	b.Write(buf[:])
}

func write[T ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int32](b *bytes.Buffer, v T) {
	check(binary.Write(b, binary.LittleEndian, v))
}

func padTo(buf []byte, n int) []byte {
	if len(buf) >= n {
		return buf
	}
	return append(buf, make([]byte, n-len(buf))...)
}

func align(x, pow uint32) uint32 {
	if pow == 0 {
		return x
	}
	n := uint32(1) << pow
	return (x + n - 1) &^ (n - 1)
}

func isZerofill(flags uint32) bool {
	return flags&0xff == 1
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
