package lfsr

type LSFR struct {
	Slots []bool
	Taps  []int
}

func (lfsr *LSFR) Next() bool {
	res := lfsr.Slots[0]

	var val bool = false

	for _, tap := range lfsr.Taps {
		val = lfsr.Slots[tap] != val
	}

	for i, slot := range lfsr.Slots {
		if i == 0 {
			continue
		}
		lfsr.Slots[i-1] = slot
	}

	lfsr.Slots[len(lfsr.Slots)-1] = val
	return res
}
