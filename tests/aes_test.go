package tests

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/as283-ua/crypto/aes"
)

func getKey(length int) []byte {
	res := make([]byte, length)
	rand.Read(res)
	return res
}

func getState() []byte {
	res := make([]byte, 16)
	for i := 0; i < len(res); i++ {
		res[i] = byte(i)
	}
	return res
}

func TestKeyExpansion128(t *testing.T) {
	key := getKey(16)
	expKey := aes.KeyExpansion(key)

	if len(expKey) != 44 {
		t.Errorf("expected expanded key length to be %v but got %v", 44, len(expKey))
	}
}

func TestKeyExpansion192(t *testing.T) {
	key := getKey(24)
	expKey := aes.KeyExpansion(key)

	if len(expKey) != 52 {
		t.Errorf("expected expanded key length to be %v but got %v", 60, len(expKey))
	}
}

func TestKeyExpansion256(t *testing.T) {
	key := getKey(32)
	expKey := aes.KeyExpansion(key)

	if len(expKey) != 60 {
		t.Errorf("expected expanded key length to be %v but got %v", 60, len(expKey))
	}
}

func TestAddRoundKey(t *testing.T) {
	key := getKey(16)
	expKey := aes.KeyExpansion(key)

	rKey := expKey[:4]

	initialState := getState()
	state := make([]byte, 16)
	copy(initialState, state)

	aes.AddRoundKey(state, rKey)
	aes.AddRoundKey(state, rKey)

	if !bytes.Equal(initialState, state) {
		t.Errorf("expected state to be %v but got %v", initialState, state)
	}
}

func TestSubBytes(t *testing.T) {
	initialState := getState()
	state := make([]byte, 16)
	copy(initialState, state)

	aes.SubBytes(state)
	aes.InvSubBytes(state)

	if !bytes.Equal(initialState, state) {
		t.Errorf("expected state to be %v but got %v", initialState, state)
	}
}

func TestShiftRows(t *testing.T) {
	initialState := getState()
	state := make([]byte, 16)
	copy(initialState, state)

	aes.ShiftRows(state)
	aes.InvShiftRows(state)

	if !bytes.Equal(initialState, state) {
		t.Errorf("expected state to be %v but got %v", initialState, state)
	}
}

func TestGaloisMult(t *testing.T) {
	cases := []struct {
		a, b, expected byte
	}{
		{0x1a, 0x16, 0xe7},
		{0x96, 0x6f, 0xf4},
	}

	for _, v := range cases {
		res := aes.GaloisMult(v.a, v.b)

		if res != v.expected {
			t.Errorf("expected result of galois multiplication G(2^8) to be %v but got %v", v.expected, res)
		}
	}
}

func TestGaloisMultCommutative(t *testing.T) {
	if aes.GaloisMult(13, 8) != aes.GaloisMult(8, 13) {
		t.Errorf("Galois multiplication should be commutative")
	}
}

func TestMatrixMult(t *testing.T) {
	a := aes.BMatrix{Array: []byte{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 0,
	}, Dimension: 4}

	var b = aes.GetMatrix([][]byte{
		{2, 3, 1, 1},
		{1, 2, 3, 1},
		{1, 1, 2, 3},
		{3, 1, 1, 2},
	})

	c := aes.MatrixMult(b, a)
	expected := []byte{
		aes.GaloisMult(1, 2) ^ aes.GaloisMult(2, 3) ^ aes.GaloisMult(3, 1) ^ aes.GaloisMult(4, 1),
		aes.GaloisMult(5, 2) ^ aes.GaloisMult(6, 3) ^ aes.GaloisMult(7, 1) ^ aes.GaloisMult(8, 1),
		aes.GaloisMult(9, 2) ^ aes.GaloisMult(10, 3) ^ aes.GaloisMult(11, 1) ^ aes.GaloisMult(12, 1),
		aes.GaloisMult(13, 2) ^ aes.GaloisMult(14, 3) ^ aes.GaloisMult(15, 1) ^ aes.GaloisMult(0, 1),

		aes.GaloisMult(1, 1) ^ aes.GaloisMult(2, 2) ^ aes.GaloisMult(3, 3) ^ aes.GaloisMult(4, 1),
		aes.GaloisMult(5, 1) ^ aes.GaloisMult(6, 2) ^ aes.GaloisMult(7, 3) ^ aes.GaloisMult(8, 1),
		aes.GaloisMult(9, 1) ^ aes.GaloisMult(10, 2) ^ aes.GaloisMult(11, 3) ^ aes.GaloisMult(12, 1),
		aes.GaloisMult(13, 1) ^ aes.GaloisMult(14, 2) ^ aes.GaloisMult(15, 3) ^ aes.GaloisMult(0, 1),

		aes.GaloisMult(1, 1) ^ aes.GaloisMult(2, 1) ^ aes.GaloisMult(3, 2) ^ aes.GaloisMult(4, 3),
		aes.GaloisMult(5, 1) ^ aes.GaloisMult(6, 1) ^ aes.GaloisMult(7, 2) ^ aes.GaloisMult(8, 3),
		aes.GaloisMult(9, 1) ^ aes.GaloisMult(10, 1) ^ aes.GaloisMult(11, 2) ^ aes.GaloisMult(12, 3),
		aes.GaloisMult(13, 1) ^ aes.GaloisMult(14, 1) ^ aes.GaloisMult(15, 2) ^ aes.GaloisMult(0, 3),

		aes.GaloisMult(1, 3) ^ aes.GaloisMult(2, 1) ^ aes.GaloisMult(3, 1) ^ aes.GaloisMult(4, 2),
		aes.GaloisMult(5, 3) ^ aes.GaloisMult(6, 1) ^ aes.GaloisMult(7, 1) ^ aes.GaloisMult(8, 2),
		aes.GaloisMult(9, 3) ^ aes.GaloisMult(10, 1) ^ aes.GaloisMult(11, 1) ^ aes.GaloisMult(12, 2),
		aes.GaloisMult(13, 3) ^ aes.GaloisMult(14, 1) ^ aes.GaloisMult(15, 1) ^ aes.GaloisMult(0, 2),
	}

	if !bytes.Equal(c.Array, expected) {
		t.Errorf("expected state to be %v but got %v", expected, c)
	}
}

func TestInvMatrixMult(t *testing.T) {
	var mixColsTable = aes.GetMatrix([][]byte{
		{2, 3, 1, 1},
		{1, 2, 3, 1},
		{1, 1, 2, 3},
		{3, 1, 1, 2},
	})

	var invMixColsTable = aes.GetMatrix([][]byte{
		{14, 11, 13, 9},
		{9, 14, 11, 13},
		{13, 9, 14, 11},
		{11, 13, 9, 14},
	})

	a := getState()

	b := aes.MatrixMult(aes.BMatrix{Array: a, Dimension: 4}, mixColsTable)

	c := aes.MatrixMult(b, invMixColsTable)

	if !bytes.Equal(a, c.Array) {
		t.Errorf("expected state to be %v but got %v", a, c)
	}
}

func TestMixColumns(t *testing.T) {
	initialState := getState()
	state := make([]byte, 16)
	copy(state, initialState)

	aes.MixColumns(state)
	aes.InvMixColumns(state)

	if !bytes.Equal(initialState, state) {
		t.Errorf("expected state to be %v but got %v", initialState, state)
	}
}

func TestAesEncryptBlock(t *testing.T) {
	message := "hi this is a msg" //16 byte block
	block := []byte(message)
	key := getKey(16)

	enc := aes.EncryptBlock(block, key)
	dec := aes.DecryptBlock(enc, key)

	if !bytes.Equal(block, dec) {
		t.Errorf("expected decrypted block to be \"%v\" but got \"%v\"", string(block), string(dec))
	}
}