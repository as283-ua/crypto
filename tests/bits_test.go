package tests

import (
	"bytes"
	"testing"

	"github.com/as283-ua/crypto/bits"
)

func TestBitsString(t *testing.T) {
	bitData := []bool{
		true, false, true, false, // 0101, first bit is right-most bit
		true, true, true, true, //f
	}

	str := bits.BitsString(bitData)
	expected := "11110101"
	if str != expected {
		t.Errorf("expected bits to be \"%v\" but got \"%v\"", expected, str)
	}
}

func bitsEqual(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetBits(t *testing.T) {
	data := []byte{0xf5, 0x49} // == 0x49f5
	bitData := bits.GetBits(data)
	expected := []bool{
		true, false, true, false, //5
		true, true, true, true, //f
		true, false, false, true, //9
		false, false, true, false, //4
	}

	if !bitsEqual(bitData, expected) {
		t.Errorf("expected bits to be \"%v\" but got \"%v\"", bits.BitsString(expected), bits.BitsString(bitData))
	}
}

/*
Groups of 8 bools form a byte
First bool is least significant bit
f f t f t f f t would be 10010100 == 0x94
*/
func TestGetBytes(t *testing.T) {
	data := []bool{ //0x94f5
		true, false, true, false, //5
		true, true, true, true, //f
		true, false, false, true, //9
		false, false, true, false, //4
	}
	byteData := bits.GetBytes(data)
	expected := []byte{0xf5, 0x49}

	if !bytes.Equal(byteData, expected) {
		t.Errorf("expected bits to be \"%v\" but got \"%v\"", expected, byteData)
	}
}

func TestBytesToUint32(t *testing.T) {
	data := []byte{0x01, 0x23, 0x45, 0x67}
	expected := 0x67452301

	res := bits.BytesToUint32(data)

	if res != uint32(expected) {
		t.Errorf("expected uint32 to be 0x%x but got 0x%x", expected, res)
	}
}

func TestUint32ToBytes(t *testing.T) {
	data := 0x67452301
	expected := []byte{0x01, 0x23, 0x45, 0x67}

	res := bits.Uint32ToBytes(uint32(data))

	if !bytes.Equal(res, expected) {
		t.Errorf("expected bytes to be \"%v\" but got \"%v\"", expected, res)
	}
}

func TestRotateLeft8(t *testing.T) {
	data := uint32(0x01234567)
	res := bits.RotateWord(data, 8)
	expected := uint32(0x23456701)

	if res != expected {
		t.Errorf("expected uint32 to be 0x%x but got 0x%x", expected, res)
	}
}

func TestRotateLeft16(t *testing.T) {
	data := uint32(0x01234567)
	res := bits.RotateWord(data, 16)
	expected := uint32(0x45670123)

	if res != expected {
		t.Errorf("expected uint32 to be 0x%x but got 0x%x", expected, res)
	}
}

func TestRotateRight8(t *testing.T) {
	data := uint32(0x01234567)
	res := bits.RotateWord(data, -8)
	expected := uint32(0x67012345)

	if res != expected {
		t.Errorf("expected uint32 to be 0x%x but got 0x%x", expected, res)
	}
}
