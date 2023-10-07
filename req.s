#include "textflag.h"

// func copy_req(dst, src *byte, n int) 
TEXT ·copy_req(SB),NOSPLIT,$0
    MOVQ dst+0(FP), DI
    MOVQ src+8(FP), SI
    MOVQ n+16(FP), BX 
    MOVQ BX, CX 
    SHRQ $3, CX
    ANDQ $7, BX 
    REP; MOVSQ
    RET

