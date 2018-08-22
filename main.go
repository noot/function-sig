package main

import (
	"encoding/hex"
	"fmt"
	"flag"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateSignature(sig string) (string) {
	bytes := []byte(sig)
	hash := crypto.Keccak256(bytes)
	hex := hex.EncodeToString(hash)
	return hex[0:8] // first 4 bytes only
}

func GenerateTopic(sig string) (string) {
	bytes := []byte(sig)
	hash := crypto.Keccak256(bytes)
	hex := hex.EncodeToString(hash)
	return hex // first 4 bytes only
}

func main() {
	sigPtr := flag.Bool("sig", false, "a bool representing whether to prompt to generate a function signature")
	topicPtr := flag.Bool("topic", false, "a bool representing whether to prompt for a topic")

	flag.Parse()
	if *sigPtr || !*topicPtr {
		var input string
		fmt.Println("enter function: ")
		fmt.Scanln(&input)
		sig := GenerateSignature(input)
		fmt.Println(sig)
	} else if *topicPtr {
		var input string
		fmt.Println("enter event: ")
		fmt.Scanln(&input)
		topic := GenerateTopic(input)
		fmt.Println(topic)
	}
}