package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateSignature(sig string) (string) {
	bytes := []byte(sig)
	hash := crypto.Keccak256(bytes)
	hex := hex.EncodeToString(hash)
	return hex[0:8] // first 4 bytes only
}

func main() {
	var input string
	fmt.Println("enter function: ")
	fmt.Scanln(&input)
	sig := GenerateSignature(input)
	fmt.Println(sig)
}