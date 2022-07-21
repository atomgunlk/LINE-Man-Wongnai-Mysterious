package main

import (
	"encoding/base64"
	"fmt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Printf("LINE-Man-Wongnai-Mysterious Secret error : %s", err.Error())
	} else {
		whatIsIt = Reverse(string(sd))
		fmt.Printf("%s", whatIsIt)
	}
}
