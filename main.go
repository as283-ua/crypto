package main

import (
	"crypto/rand"
	"fmt"
	"slices"

	"github.com/as283-ua/crypto/a5"
	"github.com/as283-ua/crypto/arc4"
)

func arc4main() {
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

func a5main() {
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

func main() {
	arc4main()
	a5main()
}
