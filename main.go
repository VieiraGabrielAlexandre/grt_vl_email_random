package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomEmailInternational() {
	fmt.Println("International")
	fmt.Println(generateString(10) + "@" + generateString(5) + ".com")
}

func generateRandomEmailNational() {
	fmt.Println("National")
	fmt.Println(generateString(10) + "@" + generateString(5) + ".com.br")
}

func generateString(qt_chars int) string {
	sb := strings.Builder{}

	sb.Grow(qt_chars)
	for i := 0; i < qt_chars; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go generateRandomEmailInternational()
	go generateRandomEmailNational()
	time.Sleep(1 * time.Second)
	fmt.Println("Waiting... to finish")
}
