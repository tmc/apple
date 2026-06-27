package ext4

import (
	"bytes"
	"context"
	"encoding/binary"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"
)

// buildTree writes a synthetic rootfs and returns its path.
func buildTree(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	mustWrite := func(rel string, data []byte, mode os.FileMode) {
		p := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, data, mode); err != nil {
			t.Fatal(err)
		}
	}

	mustWrite("hello.txt", []byte("hello world\n"), 0o644)
	mustWrite("empty", nil, 0o600)
	mustWrite("dir/nested.txt", []byte("nested\n"), 0o644)
	mustWrite("dir/deep/leaf", bytes.Repeat([]byte("x"), 100), 0o644)

	// >48KiB file to force single-indirect (12 direct blocks = 48KiB).
	big := bytes.Repeat([]byte("ABCD"), 20000) // 80000 bytes
	mustWrite("big.bin", big, 0o644)

	// Fast symlink (<60).
	if err := os.Symlink("hello.txt", filepath.Join(dir, "shortlink")); err != nil {
		t.Fatal(err)
	}
	// Slow symlink (>60).
	longTarget := string(bytes.Repeat([]byte("a/"), 40)) + "end"
	if err := os.Symlink(longTarget, filepath.Join(dir, "longlink")); err != nil {
		t.Fatal(err)
	}
	// Hardlink pair.
	if err := os.Link(filepath.Join(dir, "hello.txt"), filepath.Join(dir, "hardlink.txt")); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestBuildRoundTrip(t *testing.T) {
	src := buildTree(t)
	img := filepath.Join(t.TempDir(), "rootfs.ext4")
	if err := Build(context.Background(), src, img); err != nil {
		t.Fatal(err)
	}

	r, err := openReader(img)
	if err != nil {
		t.Fatal(err)
	}
	defer r.close()

	if r.blockSize != 4096 {
		t.Errorf("block size = %d, want 4096", r.blockSize)
	}
	if r.inodeSize != 256 {
		t.Errorf("inode size = %d, want 256", r.inodeSize)
	}

	sb, _ := r.superRaw()
	le := binary.LittleEndian
	if got := le.Uint16(sb[0x38:]); got != superMagic {
		t.Errorf("magic = %#x", got)
	}
	if got := le.Uint32(sb[0x60:]); got != featureIncompatFiletype {
		t.Errorf("feature_incompat = %#x, want %#x", got, featureIncompatFiletype)
	}
	if got := le.Uint32(sb[0x64:]); got != featureRoCompatLargeFile {
		t.Errorf("feature_ro_compat = %#x, want %#x", got, featureRoCompatLargeFile)
	}
	if got := le.Uint32(sb[0x5C:]); got != 0 {
		t.Errorf("feature_compat = %#x, want 0", got)
	}

	// Walk the tree from inode 2 and compare to source.
	found := map[string]bool{}
	var walkImg func(ino uint32, prefix string)
	walkImg = func(ino uint32, prefix string) {
		in, err := r.readInode(ino)
		if err != nil {
			t.Fatal(err)
		}
		entries, err := r.readDir(in)
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range entries {
			if e.name == "." || e.name == ".." || e.name == "lost+found" {
				continue
			}
			rel := filepath.Join(prefix, e.name)
			found[rel] = true
			ci, err := r.readInode(e.ino)
			if err != nil {
				t.Fatal(err)
			}
			srcPath := filepath.Join(src, rel)
			sfi, err := os.Lstat(srcPath)
			if err != nil {
				t.Fatal(err)
			}
			switch e.fileType {
			case ftDir:
				if ci.mode&sIFMASK != sIFDIR {
					t.Errorf("%s: not dir mode %#o", rel, ci.mode)
				}
				walkImg(e.ino, rel)
			case ftSymlink:
				target, err := r.symlinkTarget(ci)
				if err != nil {
					t.Fatal(err)
				}
				want, _ := os.Readlink(srcPath)
				if target != want {
					t.Errorf("%s: link = %q, want %q", rel, target, want)
				}
			case ftReg:
				content, err := r.readContent(ci)
				if err != nil {
					t.Fatal(err)
				}
				want, _ := os.ReadFile(srcPath)
				if !bytes.Equal(content, want) {
					t.Errorf("%s: content mismatch (%d vs %d bytes)", rel, len(content), len(want))
				}
				if ci.size != sfi.Size() {
					t.Errorf("%s: size = %d, want %d", rel, ci.size, sfi.Size())
				}
				// mtime seconds.
				if ci.mtime.Unix() != sfi.ModTime().Unix() {
					t.Errorf("%s: mtime = %v, want %v", rel, ci.mtime.Unix(), sfi.ModTime().Unix())
				}
				// permission bits.
				if ci.mode&0o777 != uint16(sfi.Mode().Perm()) {
					t.Errorf("%s: perm = %#o, want %#o", rel, ci.mode&0o777, sfi.Mode().Perm())
				}
			}
		}
	}
	walkImg(2, "")

	for _, want := range []string{"hello.txt", "empty", "dir", "dir/nested.txt", "dir/deep/leaf", "big.bin", "shortlink", "longlink", "hardlink.txt"} {
		if !found[want] {
			t.Errorf("missing entry %q", want)
		}
	}

	// Hardlink: hello.txt and hardlink.txt share an inode with links==2.
	in2, _ := r.readInode(2)
	rootEntries, _ := r.readDir(in2)
	var helloIno, hardIno uint32
	for _, e := range rootEntries {
		if e.name == "hello.txt" {
			helloIno = e.ino
		}
		if e.name == "hardlink.txt" {
			hardIno = e.ino
		}
	}
	if helloIno == 0 || helloIno != hardIno {
		t.Errorf("hardlink: hello=%d hard=%d, want equal nonzero", helloIno, hardIno)
	}
	hl, _ := r.readInode(helloIno)
	if hl.links != 2 {
		t.Errorf("hardlink links = %d, want 2", hl.links)
	}

	// big.bin must use single indirect (>12 blocks).
	var bigIno uint32
	for _, e := range rootEntries {
		if e.name == "big.bin" {
			bigIno = e.ino
		}
	}
	bin, _ := r.readInode(bigIno)
	if bin.block[12] == 0 {
		t.Error("big.bin: expected single-indirect block")
	}
}

func TestFreeCountsMatchBitmaps(t *testing.T) {
	src := buildTree(t)
	img := filepath.Join(t.TempDir(), "fs.ext4")
	if err := Build(context.Background(), src, img); err != nil {
		t.Fatal(err)
	}
	r, err := openReader(img)
	if err != nil {
		t.Fatal(err)
	}
	defer r.close()

	sb, _ := r.superRaw()
	le := binary.LittleEndian
	sbFreeBlocks := le.Uint32(sb[0x0C:])
	sbFreeInodes := le.Uint32(sb[0x10:])

	// Recount free blocks/inodes from the per-group bitmaps.
	freeBlocks, freeInodes := 0, 0
	gdt := make([]byte, r.groupCount*groupDescSize)
	if _, err := r.f.ReadAt(gdt, int64(r.blockSize)); err != nil {
		t.Fatal(err)
	}
	for g := 0; g < r.groupCount; g++ {
		bbBlk := le.Uint32(gdt[g*groupDescSize+0x00:])
		ibBlk := le.Uint32(gdt[g*groupDescSize+0x04:])
		bb, _ := r.readBlock(bbBlk)
		ib, _ := r.readBlock(ibBlk)
		for i := 0; i < r.blocksPerGroup; i++ {
			if bb[i/8]&(1<<uint(i%8)) == 0 {
				freeBlocks++
			}
		}
		for i := 0; i < r.inodesPerGroup; i++ {
			if ib[i/8]&(1<<uint(i%8)) == 0 {
				freeInodes++
			}
		}
	}
	if uint32(freeBlocks) != sbFreeBlocks {
		t.Errorf("free blocks: sb=%d bitmaps=%d", sbFreeBlocks, freeBlocks)
	}
	if uint32(freeInodes) != sbFreeInodes {
		t.Errorf("free inodes: sb=%d bitmaps=%d", sbFreeInodes, freeInodes)
	}
}

func TestDeterministic(t *testing.T) {
	src := buildTree(t)
	d := t.TempDir()
	a := filepath.Join(d, "a.ext4")
	b := filepath.Join(d, "b.ext4")
	fixed := time.Unix(1700000000, 0)
	var uuid [16]byte
	for i := range uuid {
		uuid[i] = byte(i)
	}
	// Pin atime/mtime on every entry so read-induced atime drift between the
	// two builds does not make the inputs differ.
	pinned := time.Unix(1600000000, 0)
	pin := func() {
		filepath.Walk(src, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			os.Chtimes(p, pinned, pinned)
			return nil
		})
	}
	bld := func(p string) {
		pin()
		bb := &Builder{now: func() time.Time { return fixed }, uuid: &uuid}
		if err := bb.Build(context.Background(), src, p); err != nil {
			t.Fatal(err)
		}
	}
	bld(a)
	bld(b)
	da, _ := os.ReadFile(a)
	db, _ := os.ReadFile(b)
	if !bytes.Equal(da, db) {
		t.Error("builds not byte-identical")
	}
}

func TestMultiGroup(t *testing.T) {
	src := buildTree(t)
	// Add enough one-block files that the tree spans several small groups.
	for i := 0; i < 40; i++ {
		p := filepath.Join(src, "fill"+pad(i))
		if err := os.WriteFile(p, []byte("fill"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	img := filepath.Join(t.TempDir(), "mg.ext4")
	// Tiny blocks_per_group to force several groups.
	bld := &Builder{BlocksPerGroup: 32}
	if err := bld.Build(context.Background(), src, img); err != nil {
		t.Fatal(err)
	}
	r, err := openReader(img)
	if err != nil {
		t.Fatal(err)
	}
	defer r.close()
	if r.groupCount < 3 {
		t.Errorf("group count = %d, want >= 3", r.groupCount)
	}

	// Backup superblock at the start of each group g>=1 must have magic.
	le := binary.LittleEndian
	for g := 1; g < r.groupCount; g++ {
		blk, err := r.readBlock(uint32(g * r.blocksPerGroup))
		if err != nil {
			t.Fatal(err)
		}
		if mag := le.Uint16(blk[0x38:]); mag != superMagic {
			t.Errorf("group %d backup sb magic = %#x", g, mag)
		}
		if gnr := le.Uint16(blk[0x5A:]); int(gnr) != g {
			t.Errorf("group %d block_group_nr = %d", g, gnr)
		}
	}

	// Content still readable.
	in2, _ := r.readInode(2)
	entries, _ := r.readDir(in2)
	var bigIno uint32
	for _, e := range entries {
		if e.name == "big.bin" {
			bigIno = e.ino
		}
	}
	bin, _ := r.readInode(bigIno)
	content, err := r.readContent(bin)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := os.ReadFile(filepath.Join(src, "big.bin"))
	if !bytes.Equal(content, want) {
		t.Error("big.bin content mismatch in multi-group image")
	}
}

func TestTimeExtra(t *testing.T) {
	for _, ns := range []int{0, 1, 123456789, 999999999} {
		tm := time.Unix(100, int64(ns))
		enc := encodeTimeExtra(tm)
		if enc&0x3 != 0 {
			t.Errorf("ns=%d: low epoch bits set", ns)
		}
		if got := int(enc >> 2); got != ns {
			t.Errorf("ns=%d: decoded %d", ns, got)
		}
	}
}

func TestBitmapLSB(t *testing.T) {
	b := newBitmap(16)
	b.set(0)
	b.set(9)
	if b.bits[0] != 0x01 {
		t.Errorf("byte0 = %#x, want 0x01", b.bits[0])
	}
	if b.bits[1] != 0x02 {
		t.Errorf("byte1 = %#x, want 0x02", b.bits[1])
	}
	if !b.get(0) || !b.get(9) || b.get(1) {
		t.Error("get mismatch")
	}
}

func TestDirEntryPacking(t *testing.T) {
	entries := []dirEntry{
		{ino: 2, name: ".", fileType: ftDir},
		{ino: 2, name: "..", fileType: ftDir},
		{ino: 12, name: "file", fileType: ftReg},
	}
	blocks := packDirBlocks(entries, 4096)
	if len(blocks) != 1 {
		t.Fatalf("blocks = %d, want 1", len(blocks))
	}
	le := binary.LittleEndian
	blk := blocks[0]
	// "." entry rec_len = 12 (8+1 -> round 12).
	if rl := le.Uint16(blk[4:]); rl != 12 {
		t.Errorf(". rec_len = %d, want 12", rl)
	}
	// ".." at offset 12, rec_len 12.
	if rl := le.Uint16(blk[12+4:]); rl != 12 {
		t.Errorf(".. rec_len = %d, want 12", rl)
	}
	// "file" final entry rec_len extends to block end: 4096 - 24 = 4072.
	if rl := le.Uint16(blk[24+4:]); rl != 4072 {
		t.Errorf("file rec_len = %d, want 4072", rl)
	}
}

func TestDirEntrySpansBlocks(t *testing.T) {
	// Many entries to force a second directory block.
	var entries []dirEntry
	entries = append(entries,
		dirEntry{ino: 2, name: ".", fileType: ftDir},
		dirEntry{ino: 2, name: "..", fileType: ftDir})
	names := map[string]bool{}
	for i := 0; i < 300; i++ {
		name := "entry" + pad(i)
		names[name] = true
		entries = append(entries, dirEntry{ino: uint32(12 + i), name: name, fileType: ftReg})
	}
	blocks := packDirBlocks(entries, 4096)
	if len(blocks) < 2 {
		t.Fatalf("expected multiple blocks, got %d", len(blocks))
	}
	// Each block's final entry must extend to the block end.
	le := binary.LittleEndian
	for bi, blk := range blocks {
		off := 0
		for {
			rl := int(le.Uint16(blk[off+4:]))
			if off+rl >= 4096 {
				if off+rl != 4096 {
					t.Errorf("block %d: final entry overruns: off=%d rl=%d", bi, off, rl)
				}
				break
			}
			off += rl
		}
	}
}

func pad(i int) string {
	s := []byte{'0', '0', '0'}
	s[0] = byte('0' + (i/100)%10)
	s[1] = byte('0' + (i/10)%10)
	s[2] = byte('0' + i%10)
	return string(s)
}

func TestBlockMapBoundaries(t *testing.T) {
	ptrs := 1024 // 4096/4
	cases := []int{12, 13, 12 + 1024, 12 + 1024 + 1}
	for _, n := range cases {
		// Allocate sequential data block numbers 100..; indirect from 100000.
		data := make([]uint32, n)
		for i := range data {
			data[i] = uint32(100 + i)
		}
		var indirectNext uint32 = 100000
		alloc := func() (uint32, error) { b := indirectNext; indirectNext++; return b, nil }
		m, err := buildBlockMap(data, ptrs, alloc)
		if err != nil {
			t.Fatalf("buildBlockMap: %v", err)
		}

		// Reconstruct flat data list by replaying the indirect structure.
		got := reconstruct(t, m, ptrs, n)
		if len(got) != n {
			t.Fatalf("n=%d: reconstructed %d blocks", n, len(got))
		}
		for i := range got {
			if got[i] != data[i] {
				t.Fatalf("n=%d: block %d = %d, want %d", n, i, got[i], data[i])
			}
		}
	}
}

// reconstruct replays a blockMap into a flat data-block slice using the
// in-memory indirect blocks.
func reconstruct(t *testing.T, m blockMap, ptrs, count int) []uint32 {
	t.Helper()
	byBlock := map[uint32][]uint32{}
	for _, ind := range m.indirects {
		byBlock[ind.block] = ind.data
	}
	var out []uint32
	for i := 0; i < 12 && len(out) < count; i++ {
		out = append(out, m.iBlock[i])
	}
	var follow func(blk uint32, level int)
	follow = func(blk uint32, level int) {
		if blk == 0 || len(out) >= count {
			return
		}
		ptrSlice := byBlock[blk]
		for _, p := range ptrSlice {
			if len(out) >= count {
				return
			}
			if level == 1 {
				out = append(out, p)
			} else {
				follow(p, level-1)
			}
		}
	}
	follow(m.iBlock[12], 1)
	follow(m.iBlock[13], 2)
	follow(m.iBlock[14], 3)
	if len(out) > count {
		out = out[:count]
	}
	return out
}

func TestEntriesSorted(t *testing.T) {
	// Directory entries (excluding . ..) should appear in lexical order for
	// determinism.
	src := buildTree(t)
	img := filepath.Join(t.TempDir(), "s.ext4")
	if err := Build(context.Background(), src, img); err != nil {
		t.Fatal(err)
	}
	r, _ := openReader(img)
	defer r.close()
	in2, _ := r.readInode(2)
	entries, _ := r.readDir(in2)
	var names []string
	for _, e := range entries {
		if e.name == "." || e.name == ".." || e.name == "lost+found" {
			continue
		}
		names = append(names, e.name)
	}
	if !sort.StringsAreSorted(names) {
		t.Errorf("entries not sorted: %v", names)
	}
}
