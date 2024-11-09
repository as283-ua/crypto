package e0

import (
	"fmt"

	"github.com/as283-ua/crypto/bits"
	"github.com/as283-ua/crypto/lsfr"
)

type E0 struct {
	lsfrs []lsfr.LSFR
}

func MakeE0(key []byte) (*E0, error) {
	if len(key) != 16 {
		return nil, fmt.Errorf("wrong key size: %v. Must be 16 bytes", len(key))
	}

	keyBits := bits.GetBits(key)

	lsfrs := make([]lsfr.LSFR, 4)
	sizes := []int{25, 31, 33, 39}
	taps := [][]int{
		{18, 17, 16, 13},
		{21, 6, 4},
		{30, 28, 26, 14},
		{35, 31, 29, 17, 5},
	}

	bitsUsed := 0
	for i, v := range sizes {
		lsfrs[i] = lsfr.LSFR{Slots: keyBits[bitsUsed : bitsUsed+v], Taps: taps[i]}
		bitsUsed += v
	}

	return &E0{lsfrs}, nil
}

func (cipher *E0) ChangeStateLsfrs() {
	pattern := []bool{false, true, true, true, false, true, true, false, true, false, false}
	for i := 0; i < len(cipher.lsfrs); i++ {
		for j := 0; j < len(cipher.lsfrs[i].Slots); j++ {
			cipher.lsfrs[i].Slots[j] = cipher.lsfrs[i].Slots[j] != pattern[i%len(pattern)]
		}
	}
}

func (cipher *E0) Next() bool {
	res := false
	for _, v := range cipher.lsfrs {
		res = res != v.Next()
	}

	return res
}

func (cipher E0) Encrypt(data []byte) []byte {
	dataBits := bits.GetBits(data)
	resBits := make([]bool, len(dataBits))

	for i, v := range dataBits {
		if i%2746 == 0 {
			cipher.ChangeStateLsfrs()
		}
		resBits[i] = v != cipher.Next()
	}

	res := bits.GetBytes(resBits)

	return res
}
