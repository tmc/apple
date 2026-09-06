//go:build darwin && arm64

package jaccl

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestUpstreamReductionOracle compiles the checked-in oracle against the live
// standalone JACCL headers. It is opt-in because upstream source and clang++
// are external test inputs, not package dependencies.
func TestUpstreamReductionOracle(t *testing.T) {
	if os.Getenv("JACCL_UPSTREAM_ORACLE") != "1" {
		t.Log("UNMEASURED: set JACCL_UPSTREAM_ORACLE=1 to compare native JACCL reduction encodings")
		return
	}
	upstream := os.Getenv("JACCL_UPSTREAM_SOURCE")
	if upstream == "" {
		upstream = "/Users/tmc/mlx/mlx"
	}
	include := filepath.Join(upstream, "distributed", "jaccl", "lib")
	if info, err := os.Stat(filepath.Join(include, "jaccl", "reduction_ops.h")); err != nil || info.IsDir() {
		t.Fatalf("upstream JACCL headers at %s: %v", include, err)
	}
	root, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	dir, err := os.MkdirTemp(filepath.Join(root, "tmp"), "jaccl-oracle-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	binaryPath := filepath.Join(dir, "reduction-oracle")
	compile := exec.Command("clang++", "-std=c++20", "-I", include, filepath.Join("testdata", "reduction_oracle.cpp"), "-o", binaryPath)
	if output, err := compile.CombinedOutput(); err != nil {
		t.Fatalf("compile upstream reduction oracle: %v\n%s", err, output)
	}
	output, err := exec.Command(binaryPath).Output()
	if err != nil {
		t.Fatalf("run upstream reduction oracle: %v", err)
	}
	dst := uint16Bytes(0x3c00)
	if err := reduce(dst, uint16Bytes(0x7e00), Float16, Sum); err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf(
		"%04x %04x %04x %04x\n",
		float32ToHalf(float32(math.NaN())),
		float32ToHalf(65520),
		binary.LittleEndian.Uint16(dst),
		float32ToBFloat16(float32(math.NaN())),
	)
	want += reductionOracleOutput(t)
	rankMajor := make([]byte, 4)
	if err := reduceRankMajor(rankMajor, float32Bytes(1e20, -1e20, 1), 3, 4, Float32, Sum); err != nil {
		t.Fatal(err)
	}
	want += hex.EncodeToString(rankMajor) + "\n"
	want += reductionRawEdgeOracleOutput(t, Bool, [][]byte{{0}, {1}})
	want += reductionRawEdgeOracleOutput(t, Int8, [][]byte{{0xc0}, {0xff}, {0}, {1}, {0x3f}})
	want += reductionRawEdgeOracleOutput(t, Int16, [][]byte{
		int16Bytes(-16_384), int16Bytes(-1), int16Bytes(0), int16Bytes(1), int16Bytes(16_383),
	})
	want += reductionRawEdgeOracleOutput(t, Int32, [][]byte{
		int32Bytes(-1 << 29), int32Bytes(-1), int32Bytes(0), int32Bytes(1), int32Bytes(1<<29 - 1),
	})
	want += reductionRawEdgeOracleOutput(t, Int64, [][]byte{
		int64Bytes(-1 << 61), int64Bytes(-1), int64Bytes(0), int64Bytes(1), int64Bytes(1<<61 - 1),
	})
	want += reductionRawEdgeOracleOutput(t, UInt8, [][]byte{{0}, {1}, {2}, {0x3f}, {0x7f}})
	want += reductionRawEdgeOracleOutput(t, UInt16, [][]byte{
		uint16Bytes(0), uint16Bytes(1), uint16Bytes(2), uint16Bytes(0x3fff), uint16Bytes(0x7fff),
	})
	want += reductionRawEdgeOracleOutput(t, UInt32, [][]byte{
		uint32Bytes(0), uint32Bytes(1), uint32Bytes(2), uint32Bytes(0x3fffffff), uint32Bytes(0x7fffffff),
	})
	want += reductionRawEdgeOracleOutput(t, UInt64, [][]byte{
		uint64Bytes(0), uint64Bytes(1), uint64Bytes(2), uint64Bytes(^uint64(0) / 4), uint64Bytes(^uint64(0) / 2),
	})
	want += reductionEdgeOracleOutput(t, Float16, []uint16{
		0x0000, 0x8000, 0x0001, 0x03ff, 0x0400, 0x3555, 0x3c00,
		0x3c01, 0x7bff, 0x7c00, 0xfc00, 0x7d00, 0x7e00, 0xfe00,
	})
	want += reductionEdgeOracleOutput(t, BFloat16, []uint16{
		0x0000, 0x8000, 0x0001, 0x3f80, 0x4000, 0x7f7f, 0x7f80, 0xff80,
		0x7f81, 0x7fc0,
	})
	want += reductionRawEdgeOracleOutput(t, Float32, [][]byte{
		uint32Bytes(0x00000000), uint32Bytes(0x80000000), uint32Bytes(0x3f800000), uint32Bytes(0xbf800000),
		uint32Bytes(0x7f800000), uint32Bytes(0xff800000), uint32Bytes(0x7f800001), uint32Bytes(0x7fc00000), uint32Bytes(0xffc00000),
	})
	want += reductionRawEdgeOracleOutput(t, Float64, [][]byte{
		uint64Bytes(0x0000000000000000), uint64Bytes(0x8000000000000000), uint64Bytes(0x3ff0000000000000), uint64Bytes(0xbff0000000000000),
		uint64Bytes(0x7ff0000000000000), uint64Bytes(0xfff0000000000000), uint64Bytes(0x7ff0000000000001), uint64Bytes(0x7ff8000000000000), uint64Bytes(0xfff8000000000000),
	})
	want += reductionRawEdgeOracleOutput(t, Complex64, [][]byte{
		complex64Raw(0x00000000, 0x80000000),
		complex64Raw(0x3f800000, 0xbf800000),
		complex64Raw(0x7f800001, 0xff800001),
		complex64Raw(0x7fc00000, 0xffc00000),
		complex64Raw(0xffc00000, 0x7fc00000),
	})
	if got := string(output); got != want {
		gotLines := strings.Split(strings.TrimSuffix(got, "\n"), "\n")
		wantLines := strings.Split(strings.TrimSuffix(want, "\n"), "\n")
		for line := 0; line < len(gotLines) && line < len(wantLines); line++ {
			if gotLines[line] != wantLines[line] {
				t.Fatalf("upstream reduction oracle line %d = %q, want %q", line+1, gotLines[line], wantLines[line])
			}
		}
		t.Fatalf("upstream reduction oracle emitted %d lines, want %d", len(gotLines), len(wantLines))
	}
	if strings.TrimSpace(string(output)) == "" {
		t.Fatal("upstream reduction oracle produced no encodings")
	}
}

func reductionEdgeOracleOutput(t *testing.T, dtype DType, values []uint16) string {
	t.Helper()
	raw := make([][]byte, len(values))
	for i, value := range values {
		raw[i] = uint16Bytes(value)
	}
	return reductionRawEdgeOracleOutput(t, dtype, raw)
}

func reductionRawEdgeOracleOutput(t *testing.T, dtype DType, values [][]byte) string {
	t.Helper()
	var output strings.Builder
	for _, a := range values {
		for _, b := range values {
			results := [3][]byte{append([]byte(nil), a...), append([]byte(nil), a...), append([]byte(nil), a...)}
			for i, op := range []ReduceOp{Sum, Max, Min} {
				if err := reduce(results[i], b, dtype, op); err != nil {
					t.Fatal(err)
				}
				if i > 0 {
					output.WriteByte(' ')
				}
				output.WriteString(hex.EncodeToString(results[i]))
			}
			output.WriteByte('\n')
		}
	}
	return output.String()
}

func complex64Raw(realBits, imagBits uint32) []byte {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint32(data, realBits)
	binary.LittleEndian.PutUint32(data[4:], imagBits)
	return data
}

func reductionOracleOutput(t *testing.T) string {
	t.Helper()
	tests := []struct {
		dtype DType
		dst   []byte
		src   []byte
	}{
		{Bool, []byte{1}, []byte{0}},
		{Int8, []byte{0xfe}, []byte{5}},
		{Int16, int16Bytes(-2), int16Bytes(5)},
		{Int32, int32Bytes(-2), int32Bytes(5)},
		{Int64, int64Bytes(-2), int64Bytes(5)},
		{UInt8, []byte{2}, []byte{5}},
		{UInt16, uint16Bytes(2), uint16Bytes(5)},
		{UInt32, uint32Bytes(2), uint32Bytes(5)},
		{UInt64, uint64Bytes(2), uint64Bytes(5)},
		{Float16, uint16Bytes(float32ToHalf(2.25)), uint16Bytes(float32ToHalf(1.5))},
		{BFloat16, uint16Bytes(float32ToBFloat16(2.25)), uint16Bytes(float32ToBFloat16(1.5))},
		{Float32, float32Bytes(2.25), float32Bytes(1.5)},
		{Float64, float64Bytes(2.25), float64Bytes(1.5)},
		{Complex64, complex64Bytes(complex(2, 5)), complex64Bytes(complex(1, 8))},
	}
	var output strings.Builder
	for _, test := range tests {
		values := [3][]byte{append([]byte(nil), test.dst...), append([]byte(nil), test.dst...), append([]byte(nil), test.dst...)}
		for i, op := range []ReduceOp{Sum, Max, Min} {
			if err := reduce(values[i], test.src, test.dtype, op); err != nil {
				t.Fatal(err)
			}
			if i > 0 {
				output.WriteByte(' ')
			}
			output.WriteString(hex.EncodeToString(values[i]))
		}
		output.WriteByte('\n')
	}
	return output.String()
}
