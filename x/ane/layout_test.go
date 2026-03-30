//go:build darwin

package ane

import (
	"errors"
	"testing"

	"github.com/tmc/apple/x/ane/mil"
)

func TestTensorLayoutSizes(t *testing.T) {
	tests := []struct {
		name     string
		layout   TensorLayout
		alloc    int
		logical  int
		elements int
	}{
		{
			name:     "1ch_1sp_fp16",
			layout:   TensorLayout{Channels: 1, Width: 1, Height: 1, ElemSize: 2, RowStride: 64, PlaneStride: 64},
			alloc:    64,
			logical:  2,
			elements: 1,
		},
		{
			name:     "4ch_1sp_fp16",
			layout:   TensorLayout{Channels: 4, Width: 1, Height: 1, ElemSize: 2, RowStride: 64, PlaneStride: 64},
			alloc:    256,
			logical:  8,
			elements: 4,
		},
		{
			name:     "4ch_16sp_fp16",
			layout:   TensorLayout{Channels: 4, Width: 16, Height: 1, ElemSize: 2, RowStride: 64, PlaneStride: 64},
			alloc:    256,
			logical:  128,
			elements: 64,
		},
		{
			name:     "32ch_32sp_fp16",
			layout:   TensorLayout{Channels: 32, Width: 32, Height: 1, ElemSize: 2, RowStride: 64, PlaneStride: 64},
			alloc:    2048,
			logical:  2048,
			elements: 1024,
		},
		{
			name:     "1ch_1sp_fp32",
			layout:   TensorLayout{Channels: 1, Width: 1, Height: 1, ElemSize: 4, RowStride: 64, PlaneStride: 64},
			alloc:    64,
			logical:  4,
			elements: 1,
		},
		{
			name:     "4ch_1sp_fp32",
			layout:   TensorLayout{Channels: 4, Width: 1, Height: 1, ElemSize: 4, RowStride: 64, PlaneStride: 64},
			alloc:    256,
			logical:  16,
			elements: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.layout.AllocSize(); got != tt.alloc {
				t.Errorf("AllocSize() = %d, want %d", got, tt.alloc)
			}
			if got := tt.layout.LogicalBytes(); got != tt.logical {
				t.Errorf("LogicalBytes() = %d, want %d", got, tt.logical)
			}
			if got := tt.layout.LogicalElements(); got != tt.elements {
				t.Errorf("LogicalElements() = %d, want %d", got, tt.elements)
			}
			if tt.layout.AllocSize() >= tt.layout.LogicalBytes() {
				// Expected: alloc >= logical due to stride padding.
			} else {
				t.Errorf("AllocSize() %d < LogicalBytes() %d", tt.layout.AllocSize(), tt.layout.LogicalBytes())
			}
		})
	}
}

func TestLayoutFromCompiledKernel(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const inCh, outCh, spatial = 4, 4, 16
	milText := mil.GenConvFP16IO(inCh, outCh, spatial)
	blob, err := mil.BuildIdentityWeightBlob(inCh)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	il := k.InputLayout(0)
	ol := k.OutputLayout(0)

	if il.Channels != inCh {
		t.Errorf("input channels = %d, want %d", il.Channels, inCh)
	}
	if ol.Channels != outCh {
		t.Errorf("output channels = %d, want %d", ol.Channels, outCh)
	}
	if il.Width != spatial {
		t.Errorf("input width = %d, want %d", il.Width, spatial)
	}
	if il.Height != 1 {
		t.Errorf("input height = %d, want 1", il.Height)
	}
	if il.ElemSize != 2 {
		t.Errorf("input elem size = %d, want 2 (fp16)", il.ElemSize)
	}
	if il.RowStride < il.Width*il.ElemSize {
		t.Errorf("row stride %d < width*elemSize %d", il.RowStride, il.Width*il.ElemSize)
	}
	if il.RowStride%64 != 0 {
		t.Errorf("row stride %d not 64-byte aligned", il.RowStride)
	}

	// AllocSize should be >= LogicalBytes.
	if il.AllocSize() < il.LogicalBytes() {
		t.Errorf("input AllocSize %d < LogicalBytes %d", il.AllocSize(), il.LogicalBytes())
	}
	if ol.AllocSize() < ol.LogicalBytes() {
		t.Errorf("output AllocSize %d < LogicalBytes %d", ol.AllocSize(), ol.LogicalBytes())
	}

	// InputAllocSize should match layout.
	if got := k.InputAllocSize(0); got != il.AllocSize() {
		t.Errorf("InputAllocSize = %d, want %d", got, il.AllocSize())
	}
	if got := k.OutputAllocSize(0); got != ol.AllocSize() {
		t.Errorf("OutputAllocSize = %d, want %d", got, ol.AllocSize())
	}

	t.Logf("input layout: %+v (alloc=%d logical=%d)", il, il.AllocSize(), il.LogicalBytes())
	t.Logf("output layout: %+v (alloc=%d logical=%d)", ol, ol.AllocSize(), ol.LogicalBytes())
}

func TestRawIOValidation(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	milText := mil.GenConvFP16IO(1, 1, 1)
	blob, err := mil.BuildIdentityWeightBlob(1)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	allocSize := k.InputAllocSize(0)

	// Wrong sizes should fail.
	if err := k.WriteInput(0, make([]byte, 0)); err == nil {
		t.Error("WriteInput with empty data should fail")
	}
	if err := k.WriteInput(0, make([]byte, allocSize+1)); err == nil {
		t.Error("WriteInput with oversized data should fail")
	}
	if err := k.WriteInput(0, make([]byte, allocSize-1)); err == nil {
		t.Error("WriteInput with undersized data should fail")
	}

	// Correct size should succeed.
	if err := k.WriteInput(0, make([]byte, allocSize)); err != nil {
		t.Errorf("WriteInput with correct size failed: %v", err)
	}

	// Out-of-range index.
	if err := k.WriteInput(1, make([]byte, allocSize)); err == nil {
		t.Error("WriteInput with out-of-range index should fail")
	}

	// ReadOutput validation.
	if err := k.ReadOutput(0, make([]byte, 0)); err == nil {
		t.Error("ReadOutput with empty data should fail")
	}
	if err := k.ReadOutput(0, make([]byte, allocSize)); err != nil {
		t.Errorf("ReadOutput with correct size failed: %v", err)
	}
}

func TestTypedIOValidation(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const inCh, outCh, spatial = 2, 2, 4
	milText := mil.GenConvFP16IO(inCh, outCh, spatial)
	blob, err := mil.BuildIdentityWeightBlob(inCh)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	want := inCh * spatial

	// Wrong element count should fail.
	if err := k.WriteInputFP16(0, make([]float32, want-1)); err == nil {
		t.Error("WriteInputFP16 with wrong count should fail")
	}
	if err := k.WriteInputFP16(0, make([]float32, want+1)); err == nil {
		t.Error("WriteInputFP16 with wrong count should fail")
	}

	// Correct count should succeed.
	if err := k.WriteInputFP16(0, make([]float32, want)); err != nil {
		t.Errorf("WriteInputFP16 with correct count failed: %v", err)
	}

	// ReadOutputFP16 with wrong count.
	if err := k.ReadOutputFP16(0, make([]float32, want-1)); err == nil {
		t.Error("ReadOutputFP16 with wrong count should fail")
	}
}

func TestRowStrideFor(t *testing.T) {
	tests := []struct {
		width, elemSize, want int
	}{
		{1, 2, 64},   // 2 bytes → 64 (minimum)
		{32, 2, 64},  // 64 bytes → 64
		{33, 2, 128}, // 66 bytes → 128
		{1, 4, 64},   // 4 bytes → 64
		{16, 4, 64},  // 64 bytes → 64
		{17, 4, 128}, // 68 bytes → 128
	}
	for _, tt := range tests {
		got := RowStrideFor(tt.width, tt.elemSize)
		if got != tt.want {
			t.Errorf("RowStrideFor(%d, %d) = %d, want %d", tt.width, tt.elemSize, got, tt.want)
		}
		if got%64 != 0 {
			t.Errorf("RowStrideFor(%d, %d) = %d, not 64-byte aligned", tt.width, tt.elemSize, got)
		}
	}
}

func TestTensorLayoutMultiHeight(t *testing.T) {
	// [1, 8, 4, 32] tensor: 8 channels, height=4, width=32, fp16
	l := TensorLayout{
		Channels:    8,
		Width:       32,
		Height:      4,
		ElemSize:    2,
		RowStride:   64,  // 32*2=64, already aligned
		PlaneStride: 256, // 4 rows * 64 stride
	}

	if l.AllocSize() != 8*256 {
		t.Errorf("AllocSize() = %d, want %d", l.AllocSize(), 8*256)
	}
	if l.LogicalBytes() != 8*4*32*2 {
		t.Errorf("LogicalBytes() = %d, want %d", l.LogicalBytes(), 8*4*32*2)
	}
	if l.LogicalElements() != 8*4*32 {
		t.Errorf("LogicalElements() = %d, want %d", l.LogicalElements(), 8*4*32)
	}
}

func TestUnsupportedLayoutError(t *testing.T) {
	// Verify ErrUnsupportedLayout is a distinct sentinel.
	if errors.Is(ErrUnsupportedLayout, ErrNoANE) {
		t.Error("ErrUnsupportedLayout should not match ErrNoANE")
	}
	if !errors.Is(ErrUnsupportedLayout, ErrUnsupportedLayout) {
		t.Error("ErrUnsupportedLayout should match itself")
	}
}
