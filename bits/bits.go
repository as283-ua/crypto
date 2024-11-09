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
