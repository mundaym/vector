// +build !gccgo

#include "textflag.h"

// func Sum(v []int16) int16
TEXT ·Sum(SB), NOSPLIT, $0-26
	MOVD v_ptr+0(FP), R8
	MOVD v_len+8(FP), R9
	SLD  $1, R9
	ADD  R8, R9
	MOVH $0, R7

loop:
	CMP  R8, R9
	BEQ  end
	MOVH (R8), R6
	ADD  R6, R7
	ADD  $2, R8
	BR   loop

end:
	MOVH R7, ret+24(FP)
	RET
