package a5

import (
	"fmt"
	"math"

	"github.com/as283-ua/crypto/lsfr"
)

type A5 struct {
	l1, l2, l3    lsfr.LSFR
	ck1, ck2, ck3 byte
}

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

func MakeA5(key []byte) (*A5, error) {
	if len(key) != 8 {
		return nil, fmt.Errorf("wrong key size: %v. Must be 8 bytes", len(key))
	}

	keyBits := GetBits(key)

	return &A5{
		lsfr.LSFR{Slots: keyBits[:19], Taps: []int{18, 17, 16, 13}},
		lsfr.LSFR{Slots: keyBits[19:41], Taps: []int{21, 20}},
		lsfr.LSFR{Slots: keyBits[41:], Taps: []int{22, 21, 20, 7}},
		8, 10, 10,
	}, nil
}

func (cipher *A5) Next() bool {
	l1, l2, l3 := cipher.l1.Next(), cipher.l2.Next(), cipher.l3.Next()
	res := (l1 != l2) != l3

	return res
}

func (cipher A5) Encrypt(data []byte) []byte {
	dataBits := GetBits(data)
	resBits := make([]bool, len(dataBits))

	for i, v := range dataBits {
		resBits[i] = v != cipher.Next()
	}

	res := GetBytes(resBits)

	return res
}

func (cipher A5) String() string {
	return fmt.Sprintf("{\n\t%v,\n\t%v,\n\t%v\n}", GetBytes(cipher.l1.Slots), GetBytes(cipher.l2.Slots), GetBytes(cipher.l3.Slots))
}
