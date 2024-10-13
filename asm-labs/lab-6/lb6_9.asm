input_value:            ; A123F1C0 hex
    mvk 0xf1c0, A9      
    mvklh 0xa123, A9

condition_1:            ; b8 < 0 && b9 < 0
    cmplt b8, 0, b10
    cmplt b9, 0, b11
    and b10, b11, b11

    [b11] b operation_1
    [!b11] b condition_2
    nop 5

condition_2:            ; b8 ^ b9 -> res > 0
    xor b8, b9, b10
    cmpgt b10, 0, b11

    [b11] b operation_2
    [!b11] b operation_3

    nop 5

operation_1:            ; absolute source val
    abs A9, A8    
    b end
    nop 5

operation_2:            ; put '1' in range 0-14 of source val
    set A4, 0, 14, A5
    b end
    nop 5

operation_3:            ; move val on other side
    mv A9 B9

end:
    nop 1
