start:
    cmpgt a8, a9, a1    ; a8 > a9
    cmpeq a8, a9, a2    ; a8 == a9
    and a1, a2, a1
    
    [a1] b if_greater
    [!a1] b else_less
    nop 5

if_greater:            ; operation 1
    mvk 0xc3e4, a0
    mvklh 0x0f67, a0
    abs a0, a1
    b end
    nop 5

else_less:              ; operation 2-3
    mvk 0x8600, a0      ; operation 2
    mvklh 0x7c5a, a0
    mvk 0x5d03, a1
    mvklh 0x48b0, a1

    sub a0, a1, a2

    abs a2, a3          ; operation 3

    b end
    nop 5

end:                    ; end program
    nop 1
