//go:build darwin

package oslogabi

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// The os_log argument buffer is produced by __builtin_os_log_format, a clang
// builtin implemented in the compiler rather than declared in any header. There
// is nothing to parse and nothing to generate from: the only authority for the
// encoding is clang itself.
//
// So this test asks clang. For each case it emits a C program that calls the
// builtin, compiles it, runs it, and compares the bytes clang produced against
// the bytes encode produces. A hand-transcribed table (see TestEncode) records
// what somebody once read; this records what the compiler does now.
//
// The test needs clang at test time. It skips when clang is absent.

// oracleCase is one format string measured on both sides. cArgs is the argument
// list spelled for C; goArgs is the same values for encode. They must denote the
// same values, and the test cannot check that for you: a mismatch between them
// shows up as a difference in the payload bytes, which is the failure you want.
//
// knownWrong records a case encode gets wrong today. The test asserts such a
// case DOES differ. A known-wrong entry that starts matching is a failure, so
// this list cannot go stale silently: fixing the defect forces you to remove the
// entry, and an entry can never quietly convert a defect into a pass.
type oracleCase struct {
	name       string
	format     string
	cArgs      string
	goArgs     []any
	knownWrong string
}

var oracleCases = []oracleCase{
	{name: "int", format: "count=%d", cArgs: "1234", goArgs: []any{int32(1234)}},
	{name: "private int", format: "n=%{private}d", cArgs: "5", goArgs: []any{int32(5)}},
	{name: "public uint", format: "u=%{public}u", cArgs: "3u", goArgs: []any{uint32(3)}},
	{name: "long is 8 bytes", format: "big=%ld", cArgs: "(long)(1L<<40)", goArgs: []any{int64(1) << 40}},
	{name: "long long", format: "ll=%lld", cArgs: "(long long)-2", goArgs: []any{int64(-2)}},
	{name: "size_t", format: "z=%zu", cArgs: "(size_t)99", goArgs: []any{uint64(99)}},
	{name: "char stays 4 bytes", format: "c=%hhd", cArgs: "(signed char)5", goArgs: []any{int8(5)}},
	{name: "short stays 4 bytes", format: "h=%hd", cArgs: "(short)-7", goArgs: []any{int16(-7)}},
	{name: "hex", format: "x=%x", cArgs: "0xdeadu", goArgs: []any{uint32(0xdead)}},
	{name: "two scalars", format: "a=%d b=%ld", cArgs: "7, (long)9", goArgs: []any{int32(7), int64(9)}},
	{name: "literal percent takes no argument", format: "100%% done", cArgs: "", goArgs: nil},
	{name: "no specifiers", format: "plain message", cArgs: "", goArgs: nil},

	// Floating point. In C varargs a float promotes to double, so clang always
	// emits 8 bytes of IEEE754. encode sized these at 4 and had no float case in
	// scalarBits at all, so %f logged four zero bytes.
	{name: "double", format: "f=%f", cArgs: "1.5", goArgs: []any{1.5}},
	{name: "double negative", format: "f=%f", cArgs: "-0.25", goArgs: []any{-0.25}},
	{name: "float promotes to double", format: "f=%f", cArgs: "(float)0.5f", goArgs: []any{float32(0.5)}},
	{name: "exponent form", format: "e=%e", cArgs: "1024.0", goArgs: []any{1024.0}},
	{name: "general form", format: "g=%g", cArgs: "0.125", goArgs: []any{0.125}},

	// Strings. The payload is an address and cannot agree across the two
	// programs; maskPointerPayloads zeroes those bytes on both sides.
	{name: "string", format: "s=%s", cArgs: `"hi"`, goArgs: []any{"hi"}},
	{name: "private string", format: "s=%{private}s", cArgs: `"secret"`, goArgs: []any{"secret"}},
	{name: "public string", format: "host=%{public}s", cArgs: `"apple.com"`, goArgs: []any{"apple.com"}},
	{name: "scalar then string", format: "a=%d s=%s", cArgs: `7, "hi"`, goArgs: []any{int32(7), "hi"}},
	{name: "string then scalar", format: "s=%s a=%ld", cArgs: `"hi", (long)9`, goArgs: []any{"hi", int64(9)}},
	{name: "pointer", format: "p=%p", cArgs: "(void *)0x1234", goArgs: []any{uintptr(0x1234)}},

	{
		name:       "dynamic-width binary data",
		format:     "d=%.*P",
		cArgs:      `4, "abcd"`,
		goArgs:     []any{int32(4), "abcd"},
		knownWrong: "%P needs descriptor kinds 0x1 (count) and 0x3 (data); encode knows only 0x0 and 0x2, and isConv does not list 'P', so parseSpecs walks past the conversion",
	},
	{
		name:       "errno",
		format:     "e=%m",
		cArgs:      "",
		goArgs:     nil,
		knownWrong: "%m takes no argument and is expanded by the logging system; parseSpecs does not recognise 'm' as a conversion",
	},
}

func TestEncodeAgainstClang(t *testing.T) {
	clang, err := exec.LookPath("clang")
	if err != nil {
		t.Skip("clang not found; the oracle needs a C compiler at test time")
	}
	want := runOracle(t, clang)
	if len(want) != len(oracleCases) {
		t.Fatalf("oracle returned %d buffers for %d cases", len(want), len(oracleCases))
	}

	// The oracle must be able to disagree. If every comparison passes because
	// the comparison itself is inert, this whole file proves nothing.
	t.Run("oracle can fail", func(t *testing.T) {
		got, _ := Encode(oracleCases[0].format, oracleCases[0].goArgs)
		corrupt := append([]byte(nil), want[0]...)
		corrupt[len(corrupt)-1]++
		if a, b := maskPointerPayloads(got, corrupt); bytes.Equal(a, b) {
			t.Fatal("comparison did not detect a one-byte difference; it cannot fail")
		}
	})

	for i, tc := range oracleCases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := Encode(tc.format, tc.goArgs)
			a, b := maskPointerPayloads(got, want[i])
			equal := bytes.Equal(a, b)

			if tc.knownWrong != "" {
				if equal {
					t.Fatalf("%q now agrees with clang, but it is listed as known wrong.\n"+
						"Remove the knownWrong entry. Reason recorded was: %s", tc.format, tc.knownWrong)
				}
				t.Logf("known wrong, still wrong: %s\n  clang: % x\n  encode: % x", tc.knownWrong, b, a)
				return
			}
			if !equal {
				t.Errorf("Encode(%q) disagrees with __builtin_os_log_format\n  clang:  % x\n  encode: % x", tc.format, b, a)
			}
		})
	}
}

// maskPointerPayloads zeroes the payload of every non-scalar argument in both
// buffers, because those payloads are addresses that cannot agree across two
// programs. It walks using the ORACLE's descriptors, not ours: if encode
// produced a different structure, the walk does not find the same byte ranges
// and the comparison fails, which is the intended outcome.
func maskPointerPayloads(got, oracle []byte) ([]byte, []byte) {
	a := append([]byte(nil), got...)
	b := append([]byte(nil), oracle...)
	if len(b) < 2 {
		return a, b
	}
	pos := 2
	for range int(b[1]) {
		if pos+2 > len(b) {
			break
		}
		desc, size := b[pos], int(b[pos+1])
		payload := pos + 2
		// Kind is the high nibble. Anything other than a scalar carries an
		// address or an address-derived value.
		if desc&0xf0 != kindScalar {
			for j := payload; j < payload+size; j++ {
				if j < len(a) {
					a[j] = 0
				}
				if j < len(b) {
					b[j] = 0
				}
			}
		}
		pos = payload + size
	}
	return a, b
}

// runOracle writes one C program covering every case, compiles it, runs it, and
// returns the buffers in case order. One program rather than one per case keeps
// the test to a single compile.
func runOracle(t *testing.T, clang string) [][]byte {
	t.Helper()
	dir := t.TempDir()

	var src strings.Builder
	src.WriteString("#include <os/log.h>\n#include <stdio.h>\n#include <stdint.h>\n#include <stddef.h>\n\n")
	src.WriteString("#define P(idx, fmt, ...) do { \\\n" +
		"  char b[512]; \\\n" +
		"  uint32_t n = __builtin_os_log_format_buffer_size(fmt, ##__VA_ARGS__); \\\n" +
		"  __builtin_os_log_format(b, fmt, ##__VA_ARGS__); \\\n" +
		"  printf(\"%d\", idx); \\\n" +
		"  for (uint32_t i = 0; i < n; i++) printf(\" %02x\", (unsigned char)b[i]); \\\n" +
		"  printf(\"\\n\"); \\\n" +
		"} while (0)\n\nint main(void) {\n")
	for i, tc := range oracleCases {
		if tc.cArgs == "" {
			fmt.Fprintf(&src, "  P(%d, %s);\n", i, strconv.Quote(tc.format))
		} else {
			fmt.Fprintf(&src, "  P(%d, %s, %s);\n", i, strconv.Quote(tc.format), tc.cArgs)
		}
	}
	src.WriteString("  return 0;\n}\n")

	csrc := filepath.Join(dir, "oracle.c")
	if err := os.WriteFile(csrc, []byte(src.String()), 0o644); err != nil {
		t.Fatalf("write probe: %v", err)
	}
	bin := filepath.Join(dir, "oracle")
	// -Wno-format-invalid-specifier: the corpus deliberately contains specifiers
	// that clang warns about but still encodes, and those are the interesting
	// ones.
	build := exec.Command(clang, "-o", bin, csrc, "-Wno-format", "-Wno-format-invalid-specifier")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("compile probe: %v\n%s\n--- source ---\n%s", err, out, src.String())
	}
	out, err := exec.Command(bin).Output()
	if err != nil {
		t.Fatalf("run probe: %v", err)
	}

	bufs := make([][]byte, len(oracleCases))
	for line := range strings.SplitSeq(strings.TrimSpace(string(out)), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		idx, err := strconv.Atoi(fields[0])
		if err != nil || idx < 0 || idx >= len(bufs) {
			t.Fatalf("probe emitted an unusable index: %q", line)
		}
		buf := make([]byte, 0, len(fields)-1)
		for _, f := range fields[1:] {
			v, err := strconv.ParseUint(f, 16, 8)
			if err != nil {
				t.Fatalf("probe emitted an unusable byte %q in %q", f, line)
			}
			buf = append(buf, byte(v))
		}
		bufs[idx] = buf
	}
	for i, b := range bufs {
		if b == nil {
			t.Fatalf("probe produced no output for case %d (%s)", i, oracleCases[i].name)
		}
	}
	return bufs
}
