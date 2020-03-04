package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("Experiment #1*************")
	const msg1 = "My Secret Data"

	//We're splitting the salt because we'll redistribute this later.



	var salt = "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"

	makeOutputs([]byte(msg1), []byte(salt))

	fmt.Println("\n\n\nExperiment #2************")

	const salt_pt1 = "a591a6d40bf420"
	
	const salt_pt2 = "404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
	msg2 := msg1+salt_pt1

	makeOutputs([]byte(msg2), []byte(salt_pt2))

}

func makeOutputs(msg []byte, salt []byte) {
	fmt.Printf("\nMessage: %v\n", string(msg))
	fmt.Printf("Salt/Key: %v\n\n", string(salt))

	h1 := sha256.New()
	h1.Write(append(msg, salt...))

	out1 := h1.Sum(nil)
	fmt.Printf("Hash with Salt:\t%x\n", out1)

	mac := hmac.New(sha256.New, salt)
	mac.Write([]byte(msg))
	out2 := mac.Sum(nil)
	fmt.Printf("HMAC with key:\t%x\n", out2)
}
