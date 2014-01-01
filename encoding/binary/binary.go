/*
Package binary is a simple set of constants and functions for approximating binary literals.

Example:

	package main

	import . "github.com/iNamik/go_pkg/encoding/binary"

	func main() {
		println(B8(B1111, B1111)) // 255
	}

*/
package binary

type b uint64

const (
	B0000 b = 0x00
	B0001 b = 0x01
	B0010 b = 0x02
	B0011 b = 0x03
	B0100 b = 0x04
	B0101 b = 0x05
	B0110 b = 0x06
	B0111 b = 0x07
	B1000 b = 0x08
	B1001 b = 0x09
	B1010 b = 0x0A
	B1011 b = 0x0B
	B1100 b = 0x0C
	B1101 b = 0x0D
	B1110 b = 0x0E
	B1111 b = 0x0F
)

func B8(b0, b1 b) uint8 {
	return uint8((b0 << 4) | b1)
}

func B16(b0, b1, b2, b3 b) uint16 {
	return uint16((uint16(B8(b0, b1)) << 8) | uint16(B8(b2, b3)))
}

func B32(b0, b1, b2, b3, b4, b5, b6, b7 b) uint32 {
	return uint32((uint32(B16(b0, b1, b2, b3)) << 16) | uint32(B16(b4, b5, b6, b7)))
}

func B64(b0, b1, b2, b3, b4, b5, b6, b7, b8, b9, ba, bb, bc, bd, be, bf b) uint64 {
	return uint64((uint64(B32(b0, b1, b2, b3, b4, b5, b6, b7)) << 32) | uint64(B32(b8, b9, ba, bb, bc, bd, be, bf)))
}
