#include "textflag.h"

// func cvtF32ToF16(dst *byte, src *float32, n int)
//
// Converts n float32 values to float16 using NEON FCVTN.
TEXT ·cvtF32ToF16(SB), NOSPLIT, $0-24
	MOVD	dst+0(FP), R0
	MOVD	src+8(FP), R1
	MOVD	n+16(FP), R2

	CBZ	R2, done_w

	// Process 8 floats at a time.
	CMP	$8, R2
	BLT	tail_w

loop8_w:
	// Load 8 float32 into V0, V1.
	VLD1.P	32(R1), [V0.S4, V1.S4]
	// FCVTN  V2.4H, V0.4S   — narrow low 4 floats to fp16 in V2 lower half
	WORD	$0x0e216802
	// FCVTN2 V2.8H, V1.4S   — narrow high 4 floats into V2 upper half
	WORD	$0x4e216822
	// Store 8 fp16 (16 bytes) from V2.
	VST1.P	[V2.H8], 16(R0)
	SUB	$8, R2, R2
	CMP	$8, R2
	BGE	loop8_w

tail_w:
	CBZ	R2, done_w

loop1_w:
	// Load one float32.
	FMOVS	(R1), F0
	// FCVT H0, S0  (float32 to float16 scalar)
	WORD	$0x1E23C000
	// Store the low 16 bits.
	FMOVD	F0, R3
	MOVH	R3, (R0)
	ADD	$4, R1, R1
	ADD	$2, R0, R0
	SUB	$1, R2, R2
	CBNZ	R2, loop1_w

done_w:
	RET

// func cvtF16ToF32(dst *float32, src *byte, n int)
//
// Converts n float16 values to float32 using NEON FCVTL.
TEXT ·cvtF16ToF32(SB), NOSPLIT, $0-24
	MOVD	dst+0(FP), R0
	MOVD	src+8(FP), R1
	MOVD	n+16(FP), R2

	CBZ	R2, done_r

	// Process 8 at a time.
	CMP	$8, R2
	BLT	tail_r

loop8_r:
	// Load 8 fp16 (16 bytes) into V0.
	VLD1.P	16(R1), [V0.H8]
	// We need both halves, but FCVTL will widen the lower 4 halfs
	// and clobber the full 128-bit register. So copy upper half first.
	// FCVTL2 V1.4S, V0.8H  — widen upper 4 fp16 to fp32 in V1
	WORD	$0x4E217801
	// FCVTL  V0.4S, V0.4H  — widen lower 4 fp16 to fp32 in V0
	WORD	$0x0E217800
	// Store 8 float32 (32 bytes) from V0, V1.
	VST1.P	[V0.S4, V1.S4], 32(R0)
	SUB	$8, R2, R2
	CMP	$8, R2
	BGE	loop8_r

tail_r:
	CBZ	R2, done_r

loop1_r:
	// Load one fp16.
	MOVHU	(R1), R3
	FMOVD	R3, F0
	// FCVT S0, H0  (float16 to float32 scalar)
	WORD	$0x1EE24000
	// Store one float32.
	FMOVS	F0, (R0)
	ADD	$2, R1, R1
	ADD	$4, R0, R0
	SUB	$1, R2, R2
	CBNZ	R2, loop1_r

done_r:
	RET
