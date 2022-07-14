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

	for i, j := 0, len(sd)-1; i < j; i, j = i+1, j-1 {
		sd[i], sd[j] = sd[j], sd[i]
	}

	whatIsIt := fmt.Sprintf("%s", sd)
	fmt.Println(whatIsIt)

	// Answer: "Join:us:at:LINE:MAN:Wongnai"
}
