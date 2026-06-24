//go:build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type delayedManifest struct {
	Path                   string           `json:"path,omitempty"`
	Header                 json.RawMessage  `json:"header,omitempty"`
	BuildVersion           json.RawMessage  `json:"buildVersion,omitempty"`
	LinkerOptimizationHint linkeditData     `json:"linkerOptimizationHint"`
	Sections               []delayedSection `json:"sections"`
	Symbols                []symbol         `json:"symbols"`
}

type delayedSection struct {
	Segment string   `json:"segment"`
	Name    string   `json:"name"`
	Addr    uint64   `json:"addr,omitempty"`
	Size    uint64   `json:"size"`
	Offset  uint32   `json:"offset,omitempty"`
	Align   uint32   `json:"align,omitempty"`
	Reloff  uint32   `json:"reloff,omitempty"`
	Nreloc  uint32   `json:"nreloc,omitempty"`
	Flags   uint32   `json:"flags,omitempty"`
	Header  bool     `json:"header,omitempty"`
	SHA256  string   `json:"sha256,omitempty"`
	Strings []strlit `json:"strings,omitempty"`
	Words   []uint32 `json:"words"`
	Bytes   []byte   `json:"bytes,omitempty"`
	Relocs  []reloc  `json:"relocs"`
}

type strlit struct {
	Offset uint32 `json:"offset"`
	Value  string `json:"value"`
}

type reloc struct {
	Addr   uint32 `json:"addr"`
	Value  uint32 `json:"value"`
	Symbol string `json:"symbol,omitempty"`
	Type   uint8  `json:"type"`
	Len    uint8  `json:"len"`
	PCRel  bool   `json:"pcrel,omitempty"`
	Extern bool   `json:"extern,omitempty"`
}

type symbol struct {
	Name    string `json:"name"`
	RawName string `json:"rawName,omitempty"`
	Value   uint64 `json:"value"`
	Sect    uint8  `json:"sect,omitempty"`
	Type    uint8  `json:"type"`
	Entry   bool   `json:"entry,omitempty"`
}

type linkeditData struct {
	Offset uint32 `json:"offset"`
	Size   uint32 `json:"size"`
	SHA256 string `json:"sha256,omitempty"`
	Bytes  []byte `json:"bytes,omitempty"`
}

func main() {
	check := flag.String("check", "", "check generated helper code against manifest")
	write := flag.String("write", "", "write generated helper manifest")
	flag.Parse()
	if *check == "" && *write == "" {
		fmt.Fprintln(os.Stderr, "usage: encode_delayed_helper -check manifest.json [-write out.json]")
		os.Exit(2)
	}
	m := layoutManifest()
	if *check != "" {
		var err error
		m, err = readManifest(*check)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if *write != "" {
		generated := generatedManifest(m)
		out, err := json.MarshalIndent(generated, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		out = append(out, '\n')
		if err := os.WriteFile(*write, out, 0644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if *check != "" {
		if err := checkManifest(m); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if *check == "" && *write != "" {
		if err := checkManifest(generatedManifest(m)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func readManifest(path string) (delayedManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return delayedManifest{}, fmt.Errorf("read manifest: %w", err)
	}
	var m delayedManifest
	if err := json.Unmarshal(data, &m); err != nil {
		return delayedManifest{}, fmt.Errorf("decode manifest: %w", err)
	}
	return m, nil
}

func checkManifest(m delayedManifest) error {
	var text delayedSection
	for _, sec := range m.Sections {
		if sec.Segment == "__TEXT" && sec.Name == "__text" {
			text = sec
			break
		}
	}
	if len(text.Words) == 0 {
		return fmt.Errorf("missing __TEXT,__text words")
	}
	if err := checkWords(text.Words, 0, encodeDelayedLoadResource(), "_NinePFSDelayedLoadResource"); err != nil {
		return err
	}
	const goThreadStart = 0x104 / 4
	if err := checkWords(text.Words, goThreadStart, encodeGoThread(), "_go_thread"); err != nil {
		return err
	}
	const fileoffStart = 0x218 / 4
	if err := checkWords(text.Words, fileoffStart, encodeFileoffToVMAddr(), "_fileoff_to_vmaddr"); err != nil {
		return err
	}
	if err := checkRelocs(text.Relocs, textRelocs()); err != nil {
		return err
	}
	if err := checkSymbols(m.Symbols, helperSymbols()); err != nil {
		return err
	}
	if err := checkNonTextSections(m.Sections, helperNonTextSections()); err != nil {
		return err
	}
	if err := checkLinkerOptimizationHint(m.LinkerOptimizationHint, helperLinkerOptimizationHint()); err != nil {
		return err
	}
	fmt.Printf("delayed helper encoder ok: _NinePFSDelayedLoadResource words=%d _go_thread words=%d _fileoff_to_vmaddr words=%d textRelocs=%d symbols=%d nonTextSections=%d linkerHint=%d\n",
		len(encodeDelayedLoadResource()), len(encodeGoThread()), len(encodeFileoffToVMAddr()), len(textRelocs()), len(helperSymbols()), len(helperNonTextSections()), len(helperLinkerOptimizationHint().Bytes))
	return nil
}

func generatedManifest(template delayedManifest) delayedManifest {
	m := template
	m.LinkerOptimizationHint = helperLinkerOptimizationHint()
	m.Symbols = helperSymbols()
	generatedSections := map[string]delayedSection{}
	text := delayedSection{
		Segment: "__TEXT",
		Name:    "__text",
		Size:    628,
		Flags:   2147484672,
		Words:   append(append(encodeDelayedLoadResource(), encodeGoThread()...), encodeFileoffToVMAddr()...),
		Relocs:  textRelocs(),
	}
	generatedSections[text.Segment+","+text.Name] = text
	for _, sec := range helperNonTextSections() {
		generatedSections[sec.Segment+","+sec.Name] = sec
	}
	for i, sec := range m.Sections {
		gen, ok := generatedSections[sec.Segment+","+sec.Name]
		if !ok {
			continue
		}
		sec.Strings = gen.Strings
		sec.Words = gen.Words
		sec.Bytes = gen.Bytes
		sec.Relocs = gen.Relocs
		sec.Flags = gen.Flags
		m.Sections[i] = sec
	}
	return m
}

func layoutManifest() delayedManifest {
	return delayedManifest{
		Path:         "/tmp/9pfs-delayed-nodata-build/obj/delayed_load.o",
		Header:       json.RawMessage(`{"magic":4277009103,"cputype":16777228,"filetype":1,"ncmds":5,"sizeofcmds":696,"flags":8192}`),
		BuildVersion: json.RawMessage(`{"platform":1,"minos":984064,"sdk":1705216}`),
		Sections: []delayedSection{
			{
				Segment: "__TEXT", Name: "__text",
				Addr: 0, Size: 628, Offset: 728, Align: 2, Reloff: 1576, Nreloc: 24,
				Flags: 2147484672, Header: true,
				SHA256: "cdab17062eac1c36db6aa8a9ac10405953c0449e5b5ddfa1c5b03afa0b300aab",
			},
			{
				Segment: "__TEXT", Name: "__cstring",
				Addr: 628, Size: 70, Offset: 1356,
				Flags: 2, Header: true,
				SHA256: "d0fd00faa3f3946d4ac158fae3c32c78932003f7c71d570d3722187586f27963",
			},
			{
				Segment: "__DATA", Name: "__bss",
				Addr: 848, Size: 4, Align: 2,
				Flags: 1, Header: true,
				SHA256: "df3f619804a92fdb4057192dc43dd748ea778adc52bc498ce80524c014b81119",
			},
			{
				Segment: "__DATA", Name: "__data",
				Addr: 704, Size: 40, Offset: 1432, Align: 3, Reloff: 1768, Nreloc: 1,
				Header: true,
				SHA256: "7bb716041739bb8cc9d68a50deccefc46f07b9cecf9fed6ec26f217ab748e5d8",
			},
			{
				Segment: "__DATA", Name: "__objc_imageinfo",
				Addr: 744, Size: 8, Offset: 1472,
				Flags: 268435456, Header: true,
				SHA256: "94039884329d3732c577860f40d0dfdf883ff89dac332f87216e692c08892fa2",
			},
			{
				Segment: "__LD", Name: "__compact_unwind",
				Addr: 752, Size: 96, Offset: 1480, Align: 3, Reloff: 1776, Nreloc: 3,
				Flags: 33554432, Header: true,
				SHA256: "97030d2ca466634062b022374b76cd7597899037c3ecefec2cbf55f0bdf43cd1",
			},
		},
	}
}

func checkWords(text []uint32, start int, got []uint32, name string) error {
	if len(text) < start+len(got) {
		return fmt.Errorf("__TEXT,__text has %d words, need at least %d for %s", len(text), start+len(got), name)
	}
	for i, word := range got {
		if text[start+i] != word {
			return fmt.Errorf("%s word %d mismatch: got %#08x want %#08x", name, i, word, text[start+i])
		}
	}
	return nil
}

func checkRelocs(got, want []reloc) error {
	if len(got) != len(want) {
		return fmt.Errorf("__TEXT,__text reloc count mismatch: got %d want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			return fmt.Errorf("__TEXT,__text reloc %d mismatch: got %+v want %+v", i, got[i], want[i])
		}
	}
	return nil
}

func checkSymbols(got, want []symbol) error {
	if len(got) != len(want) {
		return fmt.Errorf("symbol count mismatch: got %d want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			return fmt.Errorf("symbol %d mismatch: got %+v want %+v", i, got[i], want[i])
		}
	}
	return nil
}

func checkNonTextSections(got, want []delayedSection) error {
	byName := map[string]delayedSection{}
	for _, sec := range got {
		if sec.Segment != "__TEXT" || sec.Name != "__text" {
			byName[sec.Segment+","+sec.Name] = sec
		}
	}
	for _, wantSec := range want {
		name := wantSec.Segment + "," + wantSec.Name
		gotSec, ok := byName[name]
		if !ok {
			return fmt.Errorf("missing section %s", name)
		}
		gotSec = comparableSection(gotSec)
		wantSec = comparableSection(wantSec)
		if fmt.Sprintf("%#v", gotSec) != fmt.Sprintf("%#v", wantSec) {
			return fmt.Errorf("section %s mismatch: got %#v want %#v", name, gotSec, wantSec)
		}
	}
	return nil
}

func checkLinkerOptimizationHint(got, want linkeditData) error {
	got.SHA256 = ""
	want.SHA256 = ""
	if fmt.Sprintf("%#v", got) != fmt.Sprintf("%#v", want) {
		return fmt.Errorf("linker optimization hint mismatch: got %#v want %#v", got, want)
	}
	return nil
}

func helperLinkerOptimizationHint() linkeditData {
	return linkeditData{
		Offset: 1800,
		Size:   32,
		Bytes: []byte{
			0x07, 0x02, 0x60, 0x64, 0x07, 0x02, 0x3c, 0x40,
			0x07, 0x02, 0x2c, 0x30, 0x07, 0x02, 0x74, 0x78,
			0x07, 0x02, 0x98, 0x03, 0x9c, 0x03, 0x07, 0x02,
			0xf4, 0x03, 0xf8, 0x03, 0x00, 0x00, 0x00, 0x00,
		},
	}
}

func comparableSection(sec delayedSection) delayedSection {
	return delayedSection{
		Segment: sec.Segment,
		Name:    sec.Name,
		Size:    sec.Size,
		Flags:   sec.Flags,
		Strings: sec.Strings,
		Words:   sec.Words,
		Bytes:   sec.Bytes,
		Relocs:  sec.Relocs,
	}
}

func helperNonTextSections() []delayedSection {
	return []delayedSection{
		{
			Segment: "__TEXT",
			Name:    "__cstring",
			Size:    70,
			Flags:   2,
			Strings: []strlit{
				{Offset: 0, Value: "NinePFileSystem"},
				{Offset: 16, Value: "loadResource:options:replyHandler:"},
				{Offset: 51, Value: "__rt0_arm64_darwin"},
			},
		},
		{
			Segment: "__DATA",
			Name:    "__bss",
			Size:    4,
			Flags:   1,
		},
		{
			Segment: "__DATA",
			Name:    "__data",
			Size:    40,
			Words:   []uint32{1701734734, 1163085392, 1852142712, 1852795251, 0, 0, 0, 0, 0, 0},
			Relocs: []reloc{
				{Addr: 24, Value: 11, Symbol: "go_thread.name", Type: 0, Len: 3, Extern: true},
			},
		},
		{
			Segment: "__DATA",
			Name:    "__objc_imageinfo",
			Size:    8,
			Flags:   268435456,
			Bytes:   []byte{0, 0, 0, 0, 64, 0, 0, 0},
		},
		{
			Segment: "__LD",
			Name:    "__compact_unwind",
			Size:    96,
			Flags:   33554432,
			Words: []uint32{
				0, 0, 260, 67108879, 0, 0, 0, 0,
				260, 0, 276, 67108867, 0, 0, 0, 0,
				536, 0, 92, 33554432, 0, 0, 0, 0,
			},
			Relocs: []reloc{
				{Addr: 64, Value: 1, Type: 0, Len: 3},
				{Addr: 32, Value: 1, Type: 0, Len: 3},
				{Addr: 0, Value: 1, Type: 0, Len: 3},
			},
		},
	}
}

func helperSymbols() []symbol {
	return []symbol{
		sym("ltmp0", "", 0, 1, 14),
		sym("l_.str", "", 628, 2, 14),
		sym("l_.str.1", "", 644, 2, 14),
		sym("start_go_once.started", "_start_go_once.started", 848, 3, 14),
		sym("_go_thread", "", 260, 1, 14),
		sym("_fileoff_to_vmaddr", "", 536, 1, 14),
		sym("l_.str.2", "", 679, 2, 14),
		sym("go_thread.argv", "_go_thread.argv", 728, 4, 14),
		sym("ltmp1", "", 628, 2, 14),
		sym("ltmp2", "", 848, 3, 14),
		sym("ltmp3", "", 704, 4, 14),
		sym("go_thread.name", "_go_thread.name", 704, 4, 14),
		sym("ltmp4", "", 744, 5, 14),
		sym("ltmp5", "", 752, 6, 14),
		sym("_NinePFSDelayedLoadResource", "", 0, 1, 15),
		sym("__dyld_get_image_header", "", 0, 0, 1),
		sym("__dyld_get_image_vmaddr_slide", "", 0, 0, 1),
		sym("_class_getMethodImplementation", "", 0, 0, 1),
		sym("_objc_getClass", "", 0, 0, 1),
		sym("_pthread_create", "", 0, 0, 1),
		sym("_sel_registerName", "", 0, 0, 1),
		sym("_usleep", "", 0, 0, 1),
	}
}

func sym(name, raw string, value uint64, sect, typ uint8) symbol {
	return symbol{Name: name, RawName: raw, Value: value, Sect: sect, Type: typ, Entry: true}
}

func textRelocs() []reloc {
	return []reloc{
		pageoff(0x1f8, 7, "go_thread.argv"),
		page21(0x1f4, 7, "go_thread.argv"),
		addend(0x19c, 1),
		pageoff(0x19c, 6, "l_.str.2"),
		addend(0x198, 1),
		page21(0x198, 6, "l_.str.2"),
		branch26(0x170, 5, "_fileoff_to_vmaddr"),
		branch26(0x160, 5, "_fileoff_to_vmaddr"),
		branch26(0x124, 16, "__dyld_get_image_vmaddr_slide"),
		branch26(0x118, 15, "__dyld_get_image_header"),
		branch26(0x0a8, 21, "_usleep"),
		branch26(0x098, 17, "_class_getMethodImplementation"),
		branch26(0x088, 19, "_pthread_create"),
		pageoff(0x078, 4, "_go_thread"),
		page21(0x074, 4, "_go_thread"),
		pageoff(0x064, 3, "start_go_once.started"),
		page21(0x060, 3, "start_go_once.started"),
		branch26(0x054, 17, "_class_getMethodImplementation"),
		branch26(0x044, 20, "_sel_registerName"),
		pageoff(0x040, 2, "l_.str.1"),
		page21(0x03c, 2, "l_.str.1"),
		branch26(0x034, 18, "_objc_getClass"),
		pageoff(0x030, 1, "l_.str"),
		page21(0x02c, 1, "l_.str"),
	}
}

func branch26(addr, value uint32, symbol string) reloc {
	return reloc{Addr: addr, Value: value, Symbol: symbol, Type: 2, Len: 2, PCRel: true, Extern: true}
}

func page21(addr, value uint32, symbol string) reloc {
	return reloc{Addr: addr, Value: value, Symbol: symbol, Type: 3, Len: 2, PCRel: true, Extern: true}
}

func pageoff(addr, value uint32, symbol string) reloc {
	return reloc{Addr: addr, Value: value, Symbol: symbol, Type: 4, Len: 2, Extern: true}
}

func addend(addr, value uint32) reloc {
	return reloc{Addr: addr, Value: value, Type: 10, Len: 2}
}

func encodeDelayedLoadResource() []uint32 {
	at := func(off uint32) uint32 { return off }
	return []uint32{
		subImm(true, 31, 31, 0x60),
		stp(true, 26, 25, 31, 0x10),
		stp(true, 24, 23, 31, 0x20),
		stp(true, 22, 21, 31, 0x30),
		stp(true, 20, 19, 31, 0x40),
		stp(true, 29, 30, 31, 0x50),
		addImm(true, 29, 31, 0x50),
		movReg(true, 19, 3),
		movReg(true, 20, 2),
		movReg(true, 21, 1),
		movReg(true, 22, 0),
		adrp(0),
		addImm(true, 0, 0, 0),
		bl(at(0x34), at(0x34)),
		movReg(true, 23, 0),
		adrp(0),
		addImm(true, 0, 0, 0),
		bl(at(0x44), at(0x44)),
		movReg(true, 24, 0),
		movReg(true, 0, 23),
		movReg(true, 1, 24),
		bl(at(0x54), at(0x54)),
		movReg(true, 25, 0),
		movz(false, 8, 0),
		adrp(9),
		addImm(true, 9, 9, 0),
		movz(false, 10, 1),
		casal32(8, 10, 9),
		cbnz(false, 8, at(0x70), at(0x8c)),
		adrp(2),
		addImm(true, 2, 2, 0),
		addImm(true, 0, 31, 8),
		movz(true, 1, 0),
		movz(true, 3, 0),
		bl(at(0x88), at(0x88)),
		movz(false, 26, 0xc8),
		movReg(true, 0, 23),
		movReg(true, 1, 24),
		bl(at(0x98), at(0x98)),
		cmpReg(true, 0, 25),
		bCond(condNE, at(0xa0), at(0xcc)),
		movz(false, 0, 0x2710),
		bl(at(0xa8), at(0xa8)),
		subsImm(false, 26, 26, 1),
		bCond(condNE, at(0xb0), at(0x90)),
		ldrImm(true, 8, 19, 0x10),
		movReg(true, 0, 19),
		movz(true, 1, 0),
		movz(true, 2, 0),
		blr(8),
		b(at(0xc8), at(0xe8)),
		movReg(true, 8, 0),
		movReg(true, 0, 22),
		movReg(true, 1, 24),
		movReg(true, 2, 21),
		movReg(true, 3, 20),
		movReg(true, 4, 19),
		blr(8),
		ldp(true, 29, 30, 31, 0x50),
		ldp(true, 20, 19, 31, 0x40),
		ldp(true, 22, 21, 31, 0x30),
		ldp(true, 24, 23, 31, 0x20),
		ldp(true, 26, 25, 31, 0x10),
		addImm(true, 31, 31, 0x60),
		ret(),
	}
}

func encodeGoThread() []uint32 {
	const base = 0x104
	at := func(off uint32) uint32 { return off - base }
	return []uint32{
		stpPre(true, 22, 21, 31, -0x30),
		stp(true, 20, 19, 31, 0x10),
		stp(true, 29, 30, 31, 0x20),
		addImm(true, 29, 31, 0x20),
		movz(false, 0, 0),
		bl(at(0x118), at(0x118)),
		movReg(true, 20, 0),
		movz(false, 0, 0),
		bl(at(0x124), at(0x124)),
		ldrImm(false, 8, 20, 0x10),
		cbz(false, 8, at(0x12c), at(0x204)),
		movReg(true, 19, 0),
		addImm(true, 22, 20, 0x20),
		ldrImm(false, 9, 22, 0),
		cmpImm(false, 9, 2),
		bCond(condEQ, at(0x140), at(0x158)),
		ldrImm(false, 9, 22, 4),
		addReg(true, 22, 22, 9),
		subsImm(false, 8, 8, 1),
		bCond(condNE, at(0x150), at(0x138)),
		b(at(0x154), at(0x204)),
		ldrImm(false, 1, 22, 8),
		movReg(true, 0, 20),
		bl(at(0x160), at(0x160)),
		movReg(true, 21, 0),
		ldrImm(false, 1, 22, 0x10),
		movReg(true, 0, 20),
		bl(at(0x170), at(0x170)),
		cmpImm(true, 21, 0),
		ccmpImm(true, 0, 0, 4, condNE),
		bCond(condEQ, at(0x17c), at(0x204)),
		ldrImm(false, 8, 22, 0xc),
		cbz(false, 8, at(0x184), at(0x204)),
		movz(true, 9, 0),
		addReg(true, 10, 21, 19),
		addReg(true, 11, 0, 19),
		addImm(true, 12, 11, 1),
		adrp(13),
		addImm(true, 13, 13, 0),
		lslImm(true, 14, 9, 4),
		ldrReg(false, 14, 10, 14),
		cbz(false, 14, at(0x1a8), at(0x1d4)),
		ldrbReg(15, 11, 14),
		cmpImm(false, 15, 0x5f),
		bCond(condNE, at(0x1b4), at(0x1d4)),
		addReg(true, 14, 12, 14),
		movReg(true, 16, 13),
		cbz(false, 15, at(0x1c0), at(0x1e4)),
		ldrbPost(15, 14, 1),
		ldrbPost(17, 16, 1),
		cmpReg(false, 15, 17),
		bCond(condEQ, at(0x1d0), at(0x1c0)),
		addImm(true, 9, 9, 1),
		cmpReg(true, 9, 8),
		bCond(condNE, at(0x1dc), at(0x1a0)),
		b(at(0x1e0), at(0x204)),
		addRegShift(true, 8, 10, 9, 4),
		ldrImm(true, 8, 8, 8),
		addsReg(true, 8, 8, 19),
		bCond(condEQ, at(0x1f0), at(0x204)),
		adrp(1),
		addImm(true, 1, 1, 0),
		movz(false, 0, 1),
		blr(8),
		movz(true, 0, 0),
		ldp(true, 29, 30, 31, 0x20),
		ldp(true, 20, 19, 31, 0x10),
		ldpPost(true, 22, 21, 31, 0x30),
		ret(),
	}
}

func encodeFileoffToVMAddr() []uint32 {
	const base = 0x218
	at := func(off uint32) uint32 { return off - base }
	return []uint32{
		ldrImm(false, 8, 0, 0x10),
		cbz(false, 8, at(0x21c), at(0x260)),
		addImm(true, 9, 0, 0x20),
		movReg(false, 10, 1),
		ldrImm(false, 11, 9, 0),
		cmpImm(false, 11, 0x19),
		bCond(condNE, at(0x230), at(0x250)),
		ldrImm(true, 12, 9, 0x28),
		subsReg(true, 11, 10, 12),
		bCond(condLO, at(0x23c), at(0x250)),
		ldrImm(true, 13, 9, 0x30),
		addReg(true, 12, 13, 12),
		cmpReg(true, 12, 10),
		bCond(condHI, at(0x24c), at(0x268)),
		ldrImm(false, 11, 9, 4),
		addReg(true, 9, 9, 11),
		subsImm(false, 8, 8, 1),
		bCond(condNE, at(0x25c), at(0x228)),
		movz(true, 0, 0),
		ret(),
		ldrImm(true, 8, 9, 0x18),
		addReg(true, 0, 11, 8),
		ret(),
	}
}

const (
	condEQ = 0
	condNE = 1
	condLO = 3
	condHI = 8
)

func ldrImm(x bool, rt, rn, imm uint32) uint32 {
	if x {
		return 0xf9400000 | ((imm / 8) << 10) | (rn << 5) | rt
	}
	return 0xb9400000 | ((imm / 4) << 10) | (rn << 5) | rt
}

func stp(x bool, rt, rt2, rn, imm uint32) uint32 {
	op := uint32(0x29000000)
	scale := uint32(4)
	if x {
		op = 0xa9000000
		scale = 8
	}
	return op | ((imm / scale) << 15) | (rt2 << 10) | (rn << 5) | rt
}

func stpPre(x bool, rt, rt2, rn uint32, imm int32) uint32 {
	op := uint32(0x29800000)
	scale := int32(4)
	if x {
		op = 0xa9800000
		scale = 8
	}
	return op | (uint32((imm/scale)&0x7f) << 15) | (rt2 << 10) | (rn << 5) | rt
}

func ldp(x bool, rt, rt2, rn, imm uint32) uint32 {
	op := uint32(0x29400000)
	scale := uint32(4)
	if x {
		op = 0xa9400000
		scale = 8
	}
	return op | ((imm / scale) << 15) | (rt2 << 10) | (rn << 5) | rt
}

func ldpPost(x bool, rt, rt2, rn, imm uint32) uint32 {
	op := uint32(0x28c00000)
	scale := uint32(4)
	if x {
		op = 0xa8c00000
		scale = 8
	}
	return op | ((imm / scale) << 15) | (rt2 << 10) | (rn << 5) | rt
}

func cbz(x bool, rt, pc, target uint32) uint32 {
	op := uint32(0x34000000)
	if x {
		op = 0xb4000000
	}
	return op | (branchImm19(pc, target) << 5) | rt
}

func bCond(cond, pc, target uint32) uint32 {
	return 0x54000000 | (branchImm19(pc, target) << 5) | cond
}

func b(pc, target uint32) uint32 {
	return 0x14000000 | branchImm26(pc, target)
}

func branchImm19(pc, target uint32) uint32 {
	return ((target - pc) / 4) & 0x7ffff
}

func branchImm26(pc, target uint32) uint32 {
	return ((target - pc) / 4) & 0x3ffffff
}

func addImm(x bool, rd, rn, imm uint32) uint32 {
	op := uint32(0x11000000)
	if x {
		op = 0x91000000
	}
	return op | (imm << 10) | (rn << 5) | rd
}

func subImm(x bool, rd, rn, imm uint32) uint32 {
	op := uint32(0x51000000)
	if x {
		op = 0xd1000000
	}
	return op | (imm << 10) | (rn << 5) | rd
}

func subsImm(x bool, rd, rn, imm uint32) uint32 {
	op := uint32(0x71000000)
	if x {
		op = 0xf1000000
	}
	return op | (imm << 10) | (rn << 5) | rd
}

func movReg(x bool, rd, rm uint32) uint32 {
	op := uint32(0x2a0003e0)
	if x {
		op = 0xaa0003e0
	}
	return op | (rm << 16) | rd
}

func addReg(x bool, rd, rn, rm uint32) uint32 {
	op := uint32(0x0b000000)
	if x {
		op = 0x8b000000
	}
	return op | (rm << 16) | (rn << 5) | rd
}

func subsReg(x bool, rd, rn, rm uint32) uint32 {
	op := uint32(0x6b000000)
	if x {
		op = 0xeb000000
	}
	return op | (rm << 16) | (rn << 5) | rd
}

func cmpImm(x bool, rn, imm uint32) uint32 {
	return subsImm(x, 31, rn, imm)
}

func cmpReg(x bool, rn, rm uint32) uint32 {
	return subsReg(x, 31, rn, rm)
}

func ccmpImm(x bool, rn, imm, nzcv, cond uint32) uint32 {
	if x && rn == 0 && imm == 0 && nzcv == 4 && cond == condNE {
		return 0xfa401804
	}
	panic("unsupported ccmpImm")
}

func movz(x bool, rd, imm uint32) uint32 {
	op := uint32(0x52800000)
	if x {
		op = 0xd2800000
	}
	return op | (imm << 5) | rd
}

func ret() uint32 {
	return 0xd65f03c0
}

func adrp(rd uint32) uint32 {
	return 0x90000000 | rd
}

func bl(pc, target uint32) uint32 {
	return 0x94000000 | branchImm26(pc, target)
}

func blr(rn uint32) uint32 {
	return 0xd63f0000 | (rn << 5)
}

func cbnz(x bool, rt, pc, target uint32) uint32 {
	return cbz(x, rt, pc, target) | 0x01000000
}

func casal32(old, new, rn uint32) uint32 {
	if old != 8 || new != 10 || rn != 9 {
		panic("unsupported casal32")
	}
	return 0x88e8fd2a
}

func lslImm(x bool, rd, rn, shift uint32) uint32 {
	if x && rd == 14 && rn == 9 && shift == 4 {
		return 0xd37ced2e
	}
	panic("unsupported lslImm")
}

func ldrReg(x bool, rt, rn, rm uint32) uint32 {
	if !x && rt == 14 && rn == 10 && rm == 14 {
		return 0xb86e694e
	}
	panic("unsupported ldrReg")
}

func ldrbReg(rt, rn, rm uint32) uint32 {
	if rt == 15 && rn == 11 && rm == 14 {
		return 0x386e696f
	}
	panic("unsupported ldrbReg")
}

func ldrbPost(rt, rn, imm uint32) uint32 {
	switch {
	case rt == 15 && rn == 14 && imm == 1:
		return 0x384015cf
	case rt == 17 && rn == 16 && imm == 1:
		return 0x38401611
	default:
		panic("unsupported ldrbPost")
	}
}

func addRegShift(x bool, rd, rn, rm, shift uint32) uint32 {
	op := uint32(0x0b000000)
	if x {
		op = 0x8b000000
	}
	return op | (rm << 16) | (shift << 10) | (rn << 5) | rd
}

func addsReg(x bool, rd, rn, rm uint32) uint32 {
	op := uint32(0x2b000000)
	if x {
		op = 0xab000000
	}
	return op | (rm << 16) | (rn << 5) | rd
}
