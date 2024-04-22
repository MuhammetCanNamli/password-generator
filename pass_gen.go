package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	length := 0

	for {
		fmt.Print("Enter the length of the password you want to generate: ")
		lengthStr, _ := reader.ReadString('\n')
		newLength, err := parseInt(lengthStr)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		if newLength < 12 {
			fmt.Println("Error: Password length must be at least 12.")
			continue
		}

		length = newLength
		break
	}

	// Generate passwords with different lengths
	password, err := generatePassword(length)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated Password 1:", password)

}

func generatePassword(length int) (string, error) {
	// Defining password characters
	const (
		lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
		uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digitChars     = "0123456789"
		specialChars   = "!@#$%^&*()-_=+/"
	)

	// Combine all character sets
	allChars := lowercaseChars + uppercaseChars + digitChars + specialChars

	// Initialize password string
	var password strings.Builder

	// Generate password characters randomly
	for i := 0; i < length; i++ {
		// Generate a random index within the charset length using crypto/rand
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		randomChar := allChars[randomIndex.Int64()]
		password.WriteByte(randomChar)
	}
	return password.String(), nil
}

func parseInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("An invalid number entry: %s", s)
	}
	return num, nil
}
