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
	n := 16
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

func main() {
	// Arc4main()
	// A5main()
	// E0main()
	AesKeyExp()
}
