package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"io/ioutil"
	"os"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	function := os.Args[1]

	switch function {
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run CryptoGo encrypt to encrypt a file, and CryptoGo decrypt to decrypt a file.")
		os.Exit(1)
	}

}

func encryptHandle() {

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password2, _ := terminal.ReadPassword(0)

	if !validatePassword(password, password2) {
		panic("Passwords do not match")
	}

	fmt.Println("\nEncrypting...")
	encrypt(file, password)
	fmt.Println("\nFile successfully protected")

}

func decryptHandle() {

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)

	fmt.Println("\nDecrypting...")
	decrypt(file, password)
	fmt.Println("\nFile successfully decrypted.")

}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

func encrypt(file string, password []byte) {

	plaintext, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err.Error())
	}

	key := password
	nonce := make([]byte, 12)

	// Randomizing the nonce
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(ciphertext, nonce...)

	// create a new file for saving the encrypted data.
	f, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}

func decrypt(file string, password []byte) {
	ciphertext, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err.Error())
	}

	key := password
	salt := ciphertext[len(ciphertext) - 12:]
	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)

	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext) - 12], nil)
	if err != nil {
		panic(err.Error())
	}

	// create a new file for saving the encrypted data.
	f, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		panic(err.Error())
	}
}
