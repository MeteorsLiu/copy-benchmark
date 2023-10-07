#include "textflag.h"

// func copy_req(dst, src *byte, n int) (bx,cx  int)
TEXT ·copy_req(SB),NOSPLIT,$0
    MOVQ dst+0(FP), DI
    MOVQ src+8(FP), SI
    MOVQ n+16(FP), BX 
    MOVQ BX, CX 
    SHRQ $3, CX
    ANDQ $7, BX 
    MOVQ BX, bx+24(FP)
    MOVQ CX, cx+32(FP)
    REP; MOVSQ
    RET

