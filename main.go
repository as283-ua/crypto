package main

import (
	"crypto/rand"
	"fmt"
	"slices"

	"github.com/as283-ua/crypto/a5"
	"github.com/as283-ua/crypto/aes"
	"github.com/as283-ua/crypto/arc4"
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

// 16 byte key for aes128
func getKey() []byte {
	return []byte{
		0xaa, 0xab, 0xac, 0xad,
		0xba, 0xbb, 0xbc, 0xbd,
		0xca, 0xcb, 0xcc, 0xcd,
		0xda, 0xdb, 0xdc, 0xdd,
	}
}

func getEncrypted(msg string) []byte {
	block := []byte(msg)
	key := getKey()

	return aes.EncryptBlock(block, key)
}

func main() {
	message := "hi this is a msg" //16 byte block
	block := []byte(message)
	key := getKey()

	fmt.Printf("Original message: \"%v\"\nMessage as bytes: %v\nKey: %v\n\n", message, block, key)

	enc := aes.EncryptBlock(block, key)
	fmt.Printf("Encrypted message block: %v\nDecoded message: %v\n\n", enc, string(enc))

	dec := aes.DecryptBlock(enc, key)
	fmt.Printf("Decrypted message block: %v\nDecoded message: \"%v\"\n", dec, string(dec))

	fmt.Println("Comparing slight differences in message input:")
	fmt.Printf("aaaa %v\n\n", getEncrypted("hi this is a msg"))
	fmt.Printf("aaaa %v\n", getEncrypted("hi this is a msh"))
}
