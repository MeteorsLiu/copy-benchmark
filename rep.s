#include "textflag.h"
// func copy_rep(dst, src *byte, n int) 
TEXT ·copy_rep(SB),NOSPLIT,$0
    MOVQ dst+0(FP), DI
    MOVQ src+8(FP), SI
    MOVQ n+16(FP), CX 
    REP; MOVSB
    RET

