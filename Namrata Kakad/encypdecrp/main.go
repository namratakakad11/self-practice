package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const mysecret = "abc&1*~#^2^#s0^=)^^7%b34"

func main() {

	strtoencrypt := "hello world"

	estr, err := Encryp(strtoencrypt, mysecret)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(estr)

	dstr, err := Decrypt("AyVL7wcMVPcTbss=", mysecret)
	if err != nil {
		panic(err)
	}

	fmt.Println(dstr)

}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encryp(text, mysecret string) (string, error) {


	//creates aipher block
	block, err := aes.NewCipher([]byte(mysecret))
	if err != nil {
		fmt.Println(err)
	}
	//converts text to byte
	plaintext := []byte(text)

	//return cipher stream  which encrypts with cipher feedback mode, using the given Block. The iv must be the same length as the Block's block size

	cfb := cipher.NewCFBEncrypter(block, bytes)
//make byte slice of len text
	ciphertext := make([]byte, len(plaintext))

	//XORKeyStream XORs each byte in the given slice with a byte from the cipher's key stream. Dst and src must overlap entirely or not at all
	cfb.XORKeyStream(ciphertext, plaintext)

	return Encode(ciphertext), nil

}

func Decrypt(text, mysecret string) (string, error) {
	block, err := aes.NewCipher([]byte(mysecret))
	if err != nil {
		panic(err)
	}
	ciphertext := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plaintext := make([]byte, len(text))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil

}
