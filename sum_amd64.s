// +build !gccgo

#include "textflag.h"

// func Sum(v []int16) int16
TEXT ·Sum(SB), NOSPLIT, $0-26
	MOVQ v_base+0(FP), R8
	MOVQ v_len+8(FP), R9
	SHLQ $1, R9
	ADDQ R8, R9
	MOVQ $0, R10

loop:
	CMPQ R8, R9
	JE   end
	ADDW (R8), R10
	ADDQ $2, R8
	JMP  loop

end:
	MOVW R10, ret+24(FP)
	RET
