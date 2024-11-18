package aes

func GaloisMult(a byte, b byte) byte {
	var p byte = 0
	for b != 0 {
		if b&1 != 0 {
			p ^= a
		}
		highBit := a & 0x80
		a <<= 1
		if highBit != 0 {
			a ^= 0x1b
		}
		b >>= 1
	}
	return p
}
