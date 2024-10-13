loop_init:
    mvk -6, a2              ; arr first elem
    mvk 404, a3             ; start address
    mvk 12, a1              ; counter
    mvk 3, a4               ; value step
    mvk 48, a5              ; address step

loop:
    addk -1, a1
    stw a2, *a3
    ldb *a3--[a5], a0   
    nop 4

    cmpgt a1, 0, b2         ; counter > 0
    cmpgt a3, 0, b0         ; address > 0
    and b0, b2, b1          ; counter && address
    [b1] b loop             ; true -> repeat
    [b1] add a2, a4, a2     
    nop 5

exit:
    mvk 0x1, a10
    nop 1