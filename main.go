package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	data := "hello"
	difficulty := 4
	counter := 0
	for {
		counter = counter + 1
		newdata := data + strconv.Itoa(counter)

		hasher := sha512.New()
		hasher.Write([]byte(newdata))
		hashBytes := hasher.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)

		hashStrip := hashString[:difficulty]

		if hashStrip == strings.Repeat("0", difficulty) {
			fmt.Println("Data :", data)
			fmt.Println("newData :", newdata)
			fmt.Println("Difficulty :", difficulty)
			fmt.Printf("sha512 hash: %s\n", hashString)
			break
		}
	}
}
