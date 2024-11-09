package a5

type LSFR struct {
	slots []bool
	taps  []int
}

func (lsfr *LSFR) Next() bool {
	res := lsfr.slots[0]

	var val bool = false

	for _, tap := range lsfr.taps {
		val = lsfr.slots[tap] != val
	}

	for i, slot := range lsfr.slots {
		if i == 0 {
			continue
		}
		lsfr.slots[i-1] = slot
	}

	lsfr.slots[len(lsfr.slots)-1] = val
	return res
}
