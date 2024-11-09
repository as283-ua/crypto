package arc4

type CypheRC4 struct {
	s []byte
	i byte
	j byte
}

func MakeCypheRC4(k []byte) CypheRC4 {
	s := make([]byte, 256)
	t := make([]byte, 256)
	for i := 0; i < len(s); i++ {
		s[i] = byte(i)
		t[i] = k[i%len(k)] //repeat key if not len 256
	}

	var j byte = 0
	for i := 0; i < len(s); i++ {
		j += s[i] + t[i]
		s[i], s[j] = s[j], s[i]
	}

	return CypheRC4{s: s, i: 0, j: 0}
}

func (cyph *CypheRC4) next() byte {
	cyph.i++
	cyph.j += cyph.s[cyph.i]
	i, j := cyph.i, cyph.j
	cyph.s[i], cyph.s[j] = cyph.s[j], cyph.s[i]
	t := cyph.s[i] + cyph.s[j]
	return cyph.s[t]
}

func (cyph *CypheRC4) Encrypt(data []byte) []byte {
	res := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		res[i] = cyph.next() ^ data[i]
	}
	return res
}
