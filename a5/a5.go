package a5

import (
	"fmt"

	"github.com/as283-ua/crypto/bits"
	"github.com/as283-ua/crypto/lsfr"
)

type A5 struct {
	l1, l2, l3 lsfr.LSFR
}

func MakeA5(key []byte) (*A5, error) {
	if len(key) != 8 {
		return nil, fmt.Errorf("wrong key size: %v. Must be 8 bytes", len(key))
	}

	keyBits := bits.GetBits(key)

	return &A5{
		lsfr.LSFR{Slots: keyBits[:19], Taps: []int{18, 17, 16, 13}},
		lsfr.LSFR{Slots: keyBits[19:41], Taps: []int{21, 20}},
		lsfr.LSFR{Slots: keyBits[41:], Taps: []int{22, 21, 20, 7}},
	}, nil
}

func (cipher *A5) Next() bool {
	l1, l2, l3 := cipher.l1.Next(), cipher.l2.Next(), cipher.l3.Next()
	res := (l1 != l2) != l3

	return res
}

func (cipher A5) Encrypt(data []byte) []byte {
	dataBits := bits.GetBits(data)
	resBits := make([]bool, len(dataBits))

	for i, v := range dataBits {
		resBits[i] = v != cipher.Next()
	}

	res := bits.GetBytes(resBits)

	return res
}

func (cipher A5) String() string {
	return fmt.Sprintf("{\n\t%v,\n\t%v,\n\t%v\n}", bits.GetBytes(cipher.l1.Slots), bits.GetBytes(cipher.l2.Slots), bits.GetBytes(cipher.l3.Slots))
}
