package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	const text = "Hello World"
	const salt = "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"

	fmt.Println("Hello, playground")
	h1 := sha256.New()
	h1.Write([]byte(salt + text))

	fmt.Printf("%x\n", h1.Sum(nil))
	fmt.Printf("%x\n", sha256.Sum256([]byte(salt+text)))
	//fmt.Printf("%x\n", h1.Sum([]byte(text)))
	//fmt.Printf("%x\n", sha256.Sum256([]byte(salt + text)))

	//h2 := sha256.New()
	//h2.Write([]byte(salt))
	//h2.

	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(text))
	fmt.Printf("%x\n", mac.Sum(nil))
	fmt.Printf("%x\n", HMac256([]byte(text), []byte(salt)))
}

func HMac256(msg []byte, key []byte) []byte {
	mac:= hmac.New(sha256.New, key)
	mac.Write([]byte(msg))
	return mac.Sum(nil)
}
