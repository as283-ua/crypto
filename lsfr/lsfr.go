package lsfr

type LSFR struct {
	Slots []bool
	Taps  []int
}

func (lsfr *LSFR) Next() bool {
	res := lsfr.Slots[0]

	var val bool = false

	for _, tap := range lsfr.Taps {
		val = lsfr.Slots[tap] != val
	}

	for i, slot := range lsfr.Slots {
		if i == 0 {
			continue
		}
		lsfr.Slots[i-1] = slot
	}

	lsfr.Slots[len(lsfr.Slots)-1] = val
	return res
}
