package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a seed for randomness
	rand.Seed(time.Now().UnixNano())

	// Call the password generator function and create a 12-character password
	password := generatePassword(12)
	fmt.Println("Oluşturulan Parola:", password)
}

func generatePassword(length int) string {
	// Defining password characters
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

	// Select random characters by looping through password length
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}
