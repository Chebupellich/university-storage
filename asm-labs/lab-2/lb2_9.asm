    MVK 0xF1C0, A1
    MVKLH 0xA123, A1
    
    MVK 0x50, A2

    STW A1, *A2

    LDH *A2, A3
    NOP 4
    MVK 0x76, A2

    MV A3, A1

    