#include "textflag.h"

TEXT ·Add(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), BX
    MOVQ y+8(FP), AX
    ADDQ AX, BX
    MOVQ BX, ret+16(FP)
    RET

