package algorithmn

func reverseBits(num uint32) uint32 {
    var val uint32
    for i :=0 ; i < 32; i++ {
        val = val<<1 | num&1
        num >>= 1
    }
    return val
}