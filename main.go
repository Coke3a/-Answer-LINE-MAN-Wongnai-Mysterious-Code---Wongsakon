package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 	Coding Quiz
	// var whatIsIt string
	// secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	// sd, _ := base64.StdEncoding.DecodeString(secret)

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt := fmt.Sprintf("%q", reverseByte(sd))

	fmt.Println(whatIsIt)

	// Answer: "Join:us:at:LINE:MAN:Wongnai"
}

func reverseByte(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(reverseByte(input[1:]), input[0])
}
