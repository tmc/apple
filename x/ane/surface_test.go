//go:build darwin

package ane

import (
	"math"
	"testing"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

func TestStridedFP16SurfaceLayoutIO(t *testing.T) {
	layout := TensorLayout{
		Channels:    2,
		Width:       5,
		Height:      2,
		ElemSize:    2,
		RowStride:   64,
		PlaneStride: 128,
	}
	surf, err := createSurfaceForLayout(layout)
	if err != nil {
		t.Fatal(err)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(surf))

	if err := writeSurface(surf, make([]byte, layout.AllocSize())); err != nil {
		t.Fatal(err)
	}

	input := []float32{
		0.0, 1.0, -2.5, 3.25, -4.5,
		5.0, -6.5, 7.75, -8.0, 9.5,
		-10.25, 11.5, -12.0, 13.25, -14.5,
		15.0, -16.5, 17.75, -18.0, 19.5,
	}
	if err := writeStridedFP16WithLayout(surf, input, layout); err != nil {
		t.Fatal(err)
	}

	raw := make([]byte, layout.AllocSize())
	if err := readSurface(surf, raw); err != nil {
		t.Fatal(err)
	}
	for c := range layout.Channels {
		for h := range layout.Height {
			rowStart := c*layout.PlaneStride + h*layout.RowStride
			for w := range layout.Width {
				got := uint16(raw[rowStart+w*2]) | uint16(raw[rowStart+w*2+1])<<8
				want := float32ToFP16(input[(c*layout.Height+h)*layout.Width+w])
				if got != want {
					t.Fatalf("raw fp16 at c=%d h=%d w=%d = 0x%04x, want 0x%04x", c, h, w, got, want)
				}
			}
			for _, b := range raw[rowStart+layout.Width*2 : rowStart+layout.RowStride] {
				if b != 0 {
					t.Fatalf("padding at c=%d h=%d was not preserved", c, h)
				}
			}
		}
	}

	got := make([]float32, len(input))
	if err := readStridedFP16WithLayout(surf, got, layout); err != nil {
		t.Fatal(err)
	}
	for i, want := range input {
		if diff := math.Abs(float64(got[i] - fp16ToFloat32(float32ToFP16(want)))); diff > 1e-4 {
			t.Fatalf("round trip[%d] = %g, want %g", i, got[i], fp16ToFloat32(float32ToFP16(want)))
		}
	}
}

func TestStridedFP16SurfaceChannelIO(t *testing.T) {
	layout := TensorLayout{
		Channels:    4,
		Width:       5,
		Height:      1,
		ElemSize:    2,
		RowStride:   64,
		PlaneStride: 64,
	}
	surf, err := createSurfaceForLayout(layout)
	if err != nil {
		t.Fatal(err)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(surf))

	if err := writeSurface(surf, make([]byte, layout.AllocSize())); err != nil {
		t.Fatal(err)
	}

	input := []float32{
		1.0, 2.0, 3.0, 4.0, 5.0,
		6.0, 7.0, 8.0, 9.0, 10.0,
	}
	if err := writeStridedFP16ChannelsWithLayout(surf, input, layout, 1); err != nil {
		t.Fatal(err)
	}

	raw := make([]byte, layout.AllocSize())
	if err := readSurface(surf, raw); err != nil {
		t.Fatal(err)
	}
	for w := range layout.Width {
		bits := uint16(raw[layout.PlaneStride+w*2]) | uint16(raw[layout.PlaneStride+w*2+1])<<8
		if bits != float32ToFP16(input[w]) {
			t.Fatalf("channel 1 element %d = 0x%04x, want 0x%04x", w, bits, float32ToFP16(input[w]))
		}
		bits = uint16(raw[2*layout.PlaneStride+w*2]) | uint16(raw[2*layout.PlaneStride+w*2+1])<<8
		if bits != float32ToFP16(input[layout.Width+w]) {
			t.Fatalf("channel 2 element %d = 0x%04x, want 0x%04x", w, bits, float32ToFP16(input[layout.Width+w]))
		}
	}

	got := make([]float32, len(input))
	if err := readStridedFP16ChannelsWithLayout(surf, got, layout, 1); err != nil {
		t.Fatal(err)
	}
	for i, want := range input {
		want = fp16ToFloat32(float32ToFP16(want))
		if diff := math.Abs(float64(got[i] - want)); diff > 1e-4 {
			t.Fatalf("channel round trip[%d] = %g, want %g", i, got[i], want)
		}
	}
}

func TestStridedF32SurfaceLayoutIO(t *testing.T) {
	layout := TensorLayout{
		Channels:    2,
		Width:       7,
		Height:      2,
		ElemSize:    4,
		RowStride:   64,
		PlaneStride: 128,
	}
	surf, err := createSurfaceForLayout(layout)
	if err != nil {
		t.Fatal(err)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(surf))

	input := []float32{
		0.5, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5,
		7.5, 8.5, 9.5, 10.5, 11.5, 12.5, 13.5,
		14.5, 15.5, 16.5, 17.5, 18.5, 19.5, 20.5,
		21.5, 22.5, 23.5, 24.5, 25.5, 26.5, 27.5,
	}
	if err := writeStridedF32WithLayout(surf, input, layout); err != nil {
		t.Fatal(err)
	}

	raw := make([]byte, layout.AllocSize())
	if err := readSurface(surf, raw); err != nil {
		t.Fatal(err)
	}
	for c := range layout.Channels {
		for h := range layout.Height {
			rowStart := c*layout.PlaneStride + h*layout.RowStride
			for w := range layout.Width {
				got := *(*float32)(unsafe.Pointer(&raw[rowStart+w*4]))
				want := input[(c*layout.Height+h)*layout.Width+w]
				if got != want {
					t.Fatalf("raw f32 at c=%d h=%d w=%d = %g, want %g", c, h, w, got, want)
				}
			}
		}
	}

	got := make([]float32, len(input))
	if err := readStridedF32WithLayout(surf, got, layout); err != nil {
		t.Fatal(err)
	}
	for i, want := range input {
		if got[i] != want {
			t.Fatalf("round trip[%d] = %g, want %g", i, got[i], want)
		}
	}
}
