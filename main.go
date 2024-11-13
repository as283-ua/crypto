package main

import (
	"crypto/rand"
	"fmt"
	"slices"

	"github.com/as283-ua/crypto/a5"
	"github.com/as283-ua/crypto/aes"
	"github.com/as283-ua/crypto/arc4"
	"github.com/as283-ua/crypto/bits"
	"github.com/as283-ua/crypto/e0"
)

func Arc4main() {
	fmt.Println("ARC4 demo")
	k := make([]byte, 256)
	rand.Read(k)
	cipher := arc4.MakeCypheRC4(k)

	text := "hola buenas esto es un mensaje"
	fmt.Printf("Original message: %v\n", text)

	data := []byte(text)
	fmt.Printf("Data:\n%v\n", data)

	enc := cipher.Encrypt(data)
	fmt.Printf("Encrypted:\n%v\n\n", enc)

	kBad := make([]byte, 256)
	rand.Read(kBad)
	cipher = arc4.MakeCypheRC4(kBad)

	decBad := cipher.Encrypt(enc)
	fmt.Printf("Unencrypted with bad key:\n%v\n", decBad)
	fmt.Printf("Equal? %v\n", slices.Equal(decBad, data))
	fmt.Printf("Text:\n%v\n\n", string(decBad))

	cipher = arc4.MakeCypheRC4(k)

	dec := cipher.Encrypt(enc)
	fmt.Printf("Unencrypted with correct key:\n%v\n", dec)
	fmt.Printf("Equal? %v\n", slices.Equal(dec, data))
	fmt.Printf("Text:\n%v\n", string(dec))
}

func A5main() {
	fmt.Println("\n\033[1mA5 demo\033[0m")
	k := make([]byte, 8)

	rand.Read(k)

	cipher, err := a5.MakeA5(k)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	text := "hola buenas esto es un mensaje"
	fmt.Printf("Original message: %v\n", text)

	data := []byte(text)
	fmt.Printf("Data:\n%v\n", data)

	enc := cipher.Encrypt(data)
	fmt.Printf("Encrypted:\n%v\n\n", enc)

	kBad := make([]byte, 8)
	rand.Read(kBad)
	cipherOther, err := a5.MakeA5(kBad)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	decBad := cipherOther.Encrypt(enc)
	fmt.Printf("Unencrypted with bad key:\n%v\n", decBad)
	fmt.Printf("Equal? %v\n", slices.Equal(decBad, data))
	fmt.Printf("Text:\n%v\n\n", string(decBad))

	cipher, err = a5.MakeA5(k)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	dec := cipher.Encrypt(enc)
	fmt.Printf("Unencrypted with correct key:\n%v\n", dec)
	fmt.Printf("Equal? %v\n", slices.Equal(dec, data))
	fmt.Printf("Text:\n%v\n", string(dec))

	cipher, err = a5.MakeA5(k)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	encAgain := cipher.Encrypt(data)
	fmt.Printf("Enc again with correct key:\n%v\n", encAgain)
}

func E0main() {
	fmt.Println("\n\033[1mE0 demo\033[0m")
	k := make([]byte, 16)

	rand.Read(k)

	cipher, err := e0.MakeE0(k)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	text := "hola buenas esto es un mensaje"
	fmt.Printf("Original message: %v\n", text)

	data := []byte(text)
	fmt.Printf("Data:\n%v\n", data)

	enc := cipher.Encrypt(data)
	fmt.Printf("Encrypted:\n%v\n\n", enc)

	kBad := make([]byte, 16)
	rand.Read(kBad)
	cipherOther, err := e0.MakeE0(kBad)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	decBad := cipherOther.Encrypt(enc)
	fmt.Printf("Unencrypted with bad key:\n%v\n", decBad)
	fmt.Printf("Equal? %v\n", slices.Equal(decBad, data))
	fmt.Printf("Text:\n%v\n\n", string(decBad))

	cipher, err = e0.MakeE0(k)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	dec := cipher.Encrypt(enc)
	fmt.Printf("Unencrypted with correct key:\n%v\n", dec)
	fmt.Printf("Equal? %v\n", slices.Equal(dec, data))
	fmt.Printf("Text:\n%v\n", string(dec))

	cipher, err = e0.MakeE0(k)

	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	encAgain := cipher.Encrypt(data)
	fmt.Printf("Enc again with correct key:\n%v\n", encAgain)
}

func AesKeyExp() {
	n := 32
	key := make([]byte, n)
	for i := 0; i < n; i++ {
		key[i] = byte(i)
	}

	fmt.Println(key)
	fmt.Println(aes.KeyExpansion(key))
}

func BitsFunc() {
	a := bits.Uint32ToBytes(0xffeebbaa)
	b := bits.BytesToUint32(a)
	c := bits.Uint32ToBytes(b)
	r := bits.RotateWord(b, 8)
	o := bits.Uint32ToBytes(r)
	fmt.Println(a)
	fmt.Printf("%x\n", b)
	fmt.Println(c)
	fmt.Println(r)
	fmt.Println(o)
}

func getState() []byte {
	return []byte{
		0x01, 0x02, 0x03, 0x04,
		0x11, 0x12, 0x13, 0x14,
		0x21, 0x22, 0x23, 0x24,
		0x31, 0x32, 0x33, 0x34,
	}
}

func getKey() []byte {
	return []byte{
		0xaa, 0xab, 0xac, 0xad,
		0xba, 0xbb, 0xbc, 0xbd,
		0xca, 0xcb, 0xcc, 0xcd,
		0xda, 0xdb, 0xdc, 0xdd,
	}
}

func StateString(state []byte) string {
	res := ""
	for i := 0; i < 4; i++ {
		res += fmt.Sprintf("%02x\t%02x\t%02x\t%02x\n", state[i], state[i+4], state[i+8], state[i+12])
	}
	return res
}

func AesInverseTest() {
	state := getState()
	key := getKey()

	expKey := aes.KeyExpansion(key)

	roundKey := expKey[0:4]
	fmt.Println("round key:", roundKey)

	fmt.Println("init state:")
	fmt.Print(StateString(state))

	aes.AddRoundKey(state, roundKey)
	fmt.Println("\nstate after addroundkey:")
	fmt.Print(StateString(state))

	aes.AddRoundKey(state, roundKey)
	fmt.Println("state after addroundkey:")
	fmt.Print(StateString(state))

	aes.SubBytes(state)
	fmt.Println("\nSubBytes:")
	fmt.Print(StateString(state))

	aes.InvSubBytes(state)
	fmt.Println("InvSubBytes:")
	fmt.Print(StateString(state))

	aes.ShiftRows(state)
	fmt.Println("\nShiftRows:")
	fmt.Print(StateString(state))

	aes.InvShiftRows(state)
	fmt.Println("InvShiftRows:")
	fmt.Print(StateString(state))

	aes.MixColumns(state)
	fmt.Println("\nMixColumns:")
	fmt.Print(StateString(state))

	aes.InvMixColumns(state)
	fmt.Println("InvMixColumns:")
	fmt.Print(StateString(state))
}

func Galois() {
	var a byte = 0x1a
	var b byte = 0x16
	fmt.Printf("0x%02x⊗0x%02x=0x%02x\n", a, b, aes.GaloisMult(a, b))
	fmt.Printf("0x%b⊗0x%b=0x%b\n", a, b, aes.GaloisMult(a, b))
	fmt.Printf("%v⊗%v=%v\n", a, b, aes.GaloisMult(a, b))
}

func main() {
	// Arc4main()
	// A5main()
	// E0main()
	AesInverseTest()
	// Galois()
}
