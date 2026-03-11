//go:build darwin && arm64 && cgo

package ane

/*
#cgo CFLAGS: -O3
#include <arm_neon.h>
#include <stddef.h>
#include <stdint.h>

static void ane_cvt_f16_f32(float *dst, const _Float16 *src, size_t n) {
	size_t i = 0;
	for (; i + 7 < n; i += 8) {
		float16x8_t h = vld1q_f16((const __fp16 *)(src + i));
		vst1q_f32(dst + i, vcvt_f32_f16(vget_low_f16(h)));
		vst1q_f32(dst + i + 4, vcvt_f32_f16(vget_high_f16(h)));
	}
	for (; i < n; i++) {
		dst[i] = (float)src[i];
	}
}

static void ane_cvt_f32_f16(_Float16 *dst, const float *src, size_t n) {
	size_t i = 0;
	for (; i + 7 < n; i += 8) {
		float16x8_t h = vcombine_f16(
			vcvt_f16_f32(vld1q_f32(src + i)),
			vcvt_f16_f32(vld1q_f32(src + i + 4)));
		vst1q_f16((__fp16 *)(dst + i), h);
	}
	for (; i < n; i++) {
		dst[i] = (_Float16)src[i];
	}
}

static void ane_write_fp16_layout(
	uint8_t *dst,
	size_t alloc,
	const float *src,
	size_t src_len,
	int channels,
	int height,
	int width,
	int row_stride,
	int plane_stride) {
	if (dst == NULL || src == NULL || channels <= 0 || height <= 0 || width <= 0) {
		return;
	}
	size_t hw = (size_t)height * (size_t)width;
	size_t limit = src_len;
	size_t logical = (size_t)channels * hw;
	if (limit > logical) {
		limit = logical;
	}
	for (int c = 0; c < channels; c++) {
		for (int h = 0; h < height; h++) {
			size_t idx = (size_t)c * hw + (size_t)h * (size_t)width;
			if (idx >= limit) {
				return;
			}
			size_t n = (size_t)width;
			if (limit - idx < n) {
				n = limit - idx;
			}
			size_t off = (size_t)c * (size_t)plane_stride + (size_t)h * (size_t)row_stride;
			if (off + n * sizeof(_Float16) > alloc) {
				return;
			}
			ane_cvt_f32_f16((_Float16 *)(dst + off), src + idx, n);
		}
	}
}

static void ane_read_fp16_layout(
	float *dst,
	size_t dst_len,
	const uint8_t *src,
	size_t alloc,
	int channels,
	int height,
	int width,
	int row_stride,
	int plane_stride) {
	if (dst == NULL || src == NULL || channels <= 0 || height <= 0 || width <= 0) {
		return;
	}
	size_t hw = (size_t)height * (size_t)width;
	size_t limit = dst_len;
	size_t logical = (size_t)channels * hw;
	if (limit > logical) {
		limit = logical;
	}
	for (int c = 0; c < channels; c++) {
		for (int h = 0; h < height; h++) {
			size_t idx = (size_t)c * hw + (size_t)h * (size_t)width;
			if (idx >= limit) {
				return;
			}
			size_t n = (size_t)width;
			if (limit - idx < n) {
				n = limit - idx;
			}
			size_t off = (size_t)c * (size_t)plane_stride + (size_t)h * (size_t)row_stride;
			if (off + n * sizeof(_Float16) > alloc) {
				return;
			}
			ane_cvt_f16_f32(dst + idx, (const _Float16 *)(src + off), n);
		}
	}
}
*/
import "C"

import "unsafe"

func writeStridedFP16Bytes(dst []byte, data []float32, layout TensorLayout) {
	if len(dst) == 0 || len(data) == 0 {
		return
	}
	C.ane_write_fp16_layout(
		(*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(dst))),
		C.size_t(len(dst)),
		(*C.float)(unsafe.Pointer(unsafe.SliceData(data))),
		C.size_t(len(data)),
		C.int(layout.Channels),
		C.int(layout.Height),
		C.int(layout.Width),
		C.int(layout.RowStride),
		C.int(layout.PlaneStride),
	)
}

func readStridedFP16Bytes(data []float32, src []byte, layout TensorLayout) {
	if len(data) == 0 || len(src) == 0 {
		return
	}
	C.ane_read_fp16_layout(
		(*C.float)(unsafe.Pointer(unsafe.SliceData(data))),
		C.size_t(len(data)),
		(*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(src))),
		C.size_t(len(src)),
		C.int(layout.Channels),
		C.int(layout.Height),
		C.int(layout.Width),
		C.int(layout.RowStride),
		C.int(layout.PlaneStride),
	)
}
