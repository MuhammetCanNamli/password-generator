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

	for {
		length := getValidLength(reader)
		password, err := generatePassword(length)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println("Generated Password: ", password)

		strength := checkPasswordStrength(password)
		fmt.Println("Password Strength: ", strength)

		for {
			fmt.Print("Do you want to create another password? (Y / N): ")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToUpper(answer))

			if answer == "Y" || answer == "y" {
				fmt.Println("Let's create a new password")
				break
			} else if answer == "N" || answer == "n" {
				fmt.Println("Have a nice day :)")
				return
			} else {
				fmt.Println("You have entered incorrect data.")
			}
		}

	}
}

func getValidLength(reader *bufio.Reader) int {
	for {
		fmt.Print("Enter the length of the password you want to generate (at least 12): ")
		lengthStr, _ := reader.ReadString('\n')
		length, err := parseInt(lengthStr)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		if length < 12 {
			fmt.Println("Error: Password length must be at least 12.")
			continue
		}

		return length
	}
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

func checkPasswordStrength(password string) string {
	var (
		hasLower   bool
		hasUpper   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case 'a' <= char && char <= 'z':
			hasLower = true
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case '0' <= char && char <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
	}

	strength := "Weak"
	if len(password) >= 12 && hasLower && hasUpper && hasDigit && hasSpecial {
		strength = "Strong"
	} else if len(password) >= 8 && (hasLower || hasUpper) && hasDigit {
		strength = "Moderate"
	}

	return strength
}
