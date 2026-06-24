//go:build darwin && arm64

#include "textflag.h"

TEXT ·callSwiftSharedInstance(SB), NOSPLIT, $0-16
	MOVD fn+0(FP), R16
	BL (R16)
	MOVD R0, ret+8(FP)
	RET

TEXT ·callSwiftResume(SB), NOSPLIT, $0-16
	MOVD self+0(FP), R20
	MOVD fn+8(FP), R16
	BL (R16)
	RET
