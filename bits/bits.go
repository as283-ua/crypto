package bits

import "math"

func GetBits(data []byte) []bool {
	res := make([]bool, len(data)*8)

	for i, v := range data {
		for j := 0; j < 8; j++ {
			res[i*8+j] = v&(1<<j) != 0
		}
	}

	return res
}

func GetBytes(data []bool) []byte {
	res := make([]byte, int(math.Ceil(float64(len(data))/8)))

	for i, bit := range data {
		idx := i / 8

		v := res[idx]
		if bit {
			v |= (1 << (i % 8))
		}

		res[idx] = v
	}

	return res
}

func BitsString(data []bool) string {
	res := ""

	for _, vbool := range data {
		var v string = "1"
		if !vbool {
			v = "0"
		}

		res = v + res
	}

	return res
}

func BytesToUint32(bytes []byte) uint32 {
	var res uint32 = 0

	for i, v := range bytes {
		res |= uint32(v) << (i * 8)
	}

	return res
}

func Uint32ToBytes(value uint32) []byte {
	return []byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24)}
}

func RotateWord(word uint32, bits int) uint32 {
	if bits < 0 {
		bits = -bits
		var mask uint32 = 0xffffffff >> bits
		return word>>bits | (word & mask << (32 - bits))
	}
	var mask uint32 = 0xffffffff << bits
	return word<<bits | (word & mask >> (32 - bits))
}
