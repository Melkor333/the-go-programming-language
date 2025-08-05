// count amount of bits
package main

func mypop(b uint64) int {
	var c int
	for b > 0 {
		c += int(b)&1
		b >>= 1
	}
	return c
}

var c [256]int
func init() {
	//m = make(map[byte]int)
	for i := range c {
		c[i] = int(byte(c[i/2] + i&1))
	}

}

func preload1(x uint64) int {
	return c[byte(x)] +
	c[byte(x>>(1*8))] +
	c[byte(x>>(2*8))] +
	c[byte(x>>(3*8))] +
	c[byte(x>>(4*8))] +
	c[byte(x>>(5*8))] +
	c[byte(x>>(6*8))] +
	c[byte(x>>(7*8))]
}

func preload2(x uint64) int {
	var r int
	for i := range 8 {
		r += c[byte(x>>(i*8))]
	}
	return r
}

func clearbit(x uint64) int {
	var r int
	for x > 0 {
		r++
		x &= (x-1);
	}
	return r
}
