package aes

import (
	"fmt"

	"github.com/as283-ua/crypto/bits"
)

type AesCipher struct {
}

type AesBlock struct {
}

// bytes to rounds
var keyRoundAssociation = map[int]int{
	16: 10,
	24: 12,
	32: 14,
}

var sBox = [][]byte{
	{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
	{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
	{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
	{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
	{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
	{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
	{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
	{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
	{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
	{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
	{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
	{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
	{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
	{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
	{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
	{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
}

var invSBox = [][]byte{
	{0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb},
	{0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87, 0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb},
	{0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d, 0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e},
	{0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2, 0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25},
	{0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92},
	{0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda, 0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84},
	{0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a, 0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06},
	{0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02, 0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b},
	{0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea, 0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73},
	{0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85, 0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e},
	{0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89, 0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b},
	{0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20, 0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4},
	{0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31, 0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f},
	{0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d, 0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef},
	{0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0, 0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61},
	{0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26, 0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d},
}

// constant for round 1 = 1, double for next. If overflow, XOR with 11B (x⁸+x⁴+x³+1)
var roundConstants = []byte{
	0x00, 0x01, 0x02, 0x04,
	0x08, 0x10, 0x20, 0x40,
	0x80, 0x1B, 0x36, 0x6C,
	0xD8, 0xAB, 0x4D,
}

var mixColsTable = [][]byte{
	{2, 3, 1, 1},
	{1, 2, 3, 1},
	{1, 1, 2, 3},
	{3, 1, 1, 2},
}

var invMixColsTable = [][]byte{
	{14, 11, 13, 9},
	{9, 14, 11, 13},
	{13, 9, 14, 11},
	{11, 13, 9, 14},
}

func applyRoundConstant(bytes []byte, round int) {
	bytes[len(bytes)-1] ^= roundConstants[round]
}

func KeyExpansion(key []byte) []uint32 {
	if len(key)%8 != 0 || len(key) > 32 || len(key) < 8 {
		panic("Key size must be 16, 24 or 32")
	}
	rounds := keyRoundAssociation[len(key)]
	wordsN := 4 * (rounds + 1)
	wordsInKey := len(key) / 4

	words := make([]uint32, wordsN)

	for i := 0; i < wordsInKey; i++ {
		words[i] = bits.BytesToUint32(key[i*4 : i*4+4])
	}

	var gWord uint32

	for i := wordsInKey; i < wordsN; i++ {
		if i%wordsInKey == 0 {
			lastWord := words[i-1]
			rotated := bits.RotateWord(lastWord, 8)
			g := bits.Uint32ToBytes(rotated)
			SubBytes(g)
			applyRoundConstant(g, i/wordsInKey)
			gWord = bits.BytesToUint32(g)
		} else if i%wordsInKey == 4 {
			lastWord := words[i-1]
			g := bits.Uint32ToBytes(lastWord)
			SubBytes(g)
			gWord = bits.BytesToUint32(g)
		} else {
			gWord = words[i-1]
		}

		words[i] = words[i-wordsInKey] ^ gWord
	}

	return words
}

func AddRoundKey(state []byte, rKey []uint32) {
	if len(state) != 16 || len(rKey) != 4 {
		panic(fmt.Sprintf("State must be 16 bytes long and rkey 4 words long: state=%v, rKey=%v", len(state), len(rKey)))
	}
	for i, word := range rKey {
		wordBytes := bits.Uint32ToBytes(word)
		for j, b := range wordBytes {
			state[i*4+j] ^= b
		}
	}
}

func SubBytes(bytes []byte) {
	for i, b := range bytes {
		x := b & 0x0f
		y := b & 0xf0 >> 4

		bytes[i] = sBox[y][x]
	}
}

func InvSubBytes(bytes []byte) {
	for i, b := range bytes {
		x := b & 0x0f
		y := b & 0xf0 >> 4

		bytes[i] = invSBox[y][x]
	}
}

func ShiftRows(state []byte) {
	row := make([]byte, 4)
	for i := 1; i < 4; i++ {
		row[0], row[1], row[2], row[3] = state[i], state[i+4], state[i+8], state[i+12]
		state[i], state[i+4], state[i+8], state[i+12] = row[(i)%4], row[(1+i)%4], row[(2+i)%4], row[(3+i)%4]
	}
}

func InvShiftRows(state []byte) {
	row := make([]byte, 4)
	for i := 1; i < 4; i++ {
		row[0], row[1], row[2], row[3] = state[i], state[i+4], state[i+8], state[i+12]
		state[i], state[i+4], state[i+8], state[i+12] = row[(-i+4)%4], row[(1-i+4)%4], row[(2-i+4)%4], row[(3-i+4)%4]
	}
}

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

func RowMatrixMult(a []byte, b []byte) byte {
	if len(a) != 4 || len(b) != 4 {
		panic("Matrix must be of size 4")
	}

	return GaloisMult(a[0], b[0]) ^ GaloisMult(a[1], b[1]) ^ GaloisMult(a[2], b[2]) ^ GaloisMult(a[3], b[3])
}

func MatrixMult(a []byte, b [][]byte) []byte {
	res := make([]byte, 16)
	for bRow := 0; bRow < 4; bRow++ {
		for aColumn := 0; aColumn < 4; aColumn++ {
			rowMult := RowMatrixMult(b[bRow], a[aColumn*4:aColumn*4+4])
			res[aColumn*4+aColumn] = rowMult
		}
	}
	return res
}

func MixColumns(state []byte) {
	res := MatrixMult(state, mixColsTable) //avoid modifying state while calculating matrix multiplication
	copy(res, state)
}

func InvMixColumns(state []byte) {
	res := MatrixMult(state, invMixColsTable)
	copy(res, state)
}

func EncryptBlock(data []byte, key []byte) []byte {
	if len(data) != 16 {
		panic("Invalid data block size")
	}

	state := make([]byte, len(data))
	copy(state, data)

	expKey := KeyExpansion(key)
	AddRoundKey(state, expKey[0:4])

	rounds := keyRoundAssociation[len(key)]

	for i := 1; i < rounds; i++ {
		SubBytes(state)
		ShiftRows(state)
		MixColumns(state)
		AddRoundKey(state, expKey[i*4:i*4+4])
	}

	SubBytes(state)
	ShiftRows(state)
	AddRoundKey(state, expKey[rounds*4:rounds*4+4])

	return state
}

func DecryptBlock(data []byte, key []byte) []byte {
	if len(data) != 16 {
		panic("Invalid data block size")
	}

	state := make([]byte, len(data))
	copy(state, data)

	expKey := KeyExpansion(key)
	rounds := keyRoundAssociation[len(key)]

	AddRoundKey(state, expKey[rounds*4:rounds*4+4])

	for i := rounds - 1; i >= 1; i-- {
		InvShiftRows(state)
		InvSubBytes(state)
		AddRoundKey(state, expKey[i*4:i*4+4])
		InvMixColumns(state)
	}

	InvShiftRows(state)
	InvSubBytes(state)
	AddRoundKey(state, expKey[0:4])

	return state
}
