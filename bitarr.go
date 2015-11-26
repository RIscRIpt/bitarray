package bitarray

import (
	"fmt"
	"math"
)

var _ = math.Ceil
var _ = fmt.Printf

/**
Bitmap algorithm. (BitSet With mulit-choiced bit width.)

*/

type BitArray struct {
	b             []byte //we use bit
	valueBitWidth byte   // how many bit present one value. only value 1,2,4,8 is supported
	countPerByte  byte
	bitmapLen     uint32
}

/*
bitmapLen: how many bit we should save
*/
func NewBitArray(bitmapLen uint32, valueBitWidth byte) *BitArray {
	return new(BitArray).Init(bitmapLen, valueBitWidth)
}

func (s *BitArray) Init(bitmapLen uint32, valueBitWidth byte) *BitArray {
	if valueBitWidth == 0 || valueBitWidth&(valueBitWidth-1) != 0 || valueBitWidth > 8 {
		panic("BitArray validBitLen only 1,2,4,8 is supported")
	}
	s.countPerByte = 8 / valueBitWidth
	s.b = make([]byte, bitmapLen/uint32(s.countPerByte)+1)
	s.bitmapLen = bitmapLen
	s.valueBitWidth = valueBitWidth
	return s
}

func (s *BitArray) GetAllocLen() int {
	return len(s.b)
}

func (s *BitArray) SetB(pos uint32, val byte) {
	whichByte := pos / uint32(s.countPerByte)
	whichPos := pos % uint32(s.countPerByte)
	n := byte(whichPos)
	w := s.valueBitWidth
	nw := n * w
	oo := ^((byte(1)<<w - 1) << nw)
	zr := s.b[whichByte] & oo // [rr00 rrrr]
	sr := val << nw           // [00ss 0000]
	s.b[whichByte] = zr | sr
}

func (s *BitArray) GetBytes() []byte {
	return s.b
}

func (s *BitArray) GetB(pos uint32) byte {
	whichByte := pos / uint32(s.countPerByte)
	whichPos := pos % uint32(s.countPerByte)
	n := byte(whichPos)
	w := s.valueBitWidth
	return (s.b[whichByte] >> (n * w)) & (1<<w - 1)
}
