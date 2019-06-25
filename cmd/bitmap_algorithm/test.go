package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

const SHIFT = 5
const MASK uint32 = 0x1F

type x struct {
	m map[uint32]uint32
}

func (x *x) set(n uint32)  {
	idx := n>>SHIFT
	x.m[idx] = x.m[idx] | 1<<(n&MASK)
}

func (x *x) isPresent(n uint32) bool {
	idx := n>>SHIFT

	value, isPresent := x.m[idx]
	if !isPresent {
		return false
	}

	num := (value >> (n&MASK)) & 1
	if num == 0 {
		return  false
	}

	return true
}

func main()  {
	newX := new(x)
	newX.m = map[uint32]uint32{
	}

	i := 0
	newX.set(uint32(12))
	unixtime1 := time.Now().Second()
	for i < 10000000 {
		i = i+1

		n :=rand.Intn(1000000000)
		newX.set(uint32(n))
	}
	unixtime2 := time.Now().Second()

	bo := newX.isPresent(uint32(12))

	//fmt.Println(newX);
	fmt.Println(unixtime2-unixtime1,bo,unsafe.Sizeof(*newX));


}
