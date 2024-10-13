loop_init:
    mvk -6, a0          ; arr first elem
    mvk -6, a1          ; checked value
    mvk 3, a2           ; value step
    mvk 0, a3           ; stop value

loop:
    cmpeq a0, 0, a1
    [!a1] b loop
    [!a1] add a0, a2, a0
    nop 4

exit:
    mvk 0x1, a10
    nop 1