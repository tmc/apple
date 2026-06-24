//go:build darwin && arm64 && pureentry

#include "textflag.h"

TEXT ·nsextMainEntry(SB), NOSPLIT, $0-0
	JMP nsext_main(SB)
