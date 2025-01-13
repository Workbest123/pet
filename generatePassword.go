package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

// go run generatePassword.go -length=16 -output=my_password.txt запуск программы через флаги можно сохранять в текстовый файл пароль от любого документа Например "пароль_от_одноклассников"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_+=<>?"

// generatePassword создает случайный пароль заданной длины
func generatePassword(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}
		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}

func saveToFile(filename, password string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(password)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func main() {
	// Определяем длину пароля через флаг командной строки
	length := flag.Int("length", 12, "Length of the generated password")
	output := flag.String("output", "password.txt", "File to save the generated password")
	flag.Parse()

	password, err := generatePassword(*length)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated password:", password)

	err = saveToFile(*output, password)
	if err != nil {
		fmt.Println("Error saving to file:", err)
		return
	}

	fmt.Println("Password saved to file:", *output)
}
