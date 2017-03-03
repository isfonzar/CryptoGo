package main

import (
	"os"
	"bytes"
	"io"
	"io/ioutil"
	"crypto/cipher"
	"crypto/aes"
	"crypto/rand"
	"golang.org/x/crypto/pbkdf2"
	"fmt"
	"crypto/sha1"
	"encoding/hex"
)

func main() {

	fmt.Println("Encrypting")
	encrypt()

	fmt.Println("Decrypting")
	decrypt()

}

func encrypt() {

	if _, err := os.Stat("teste.txt"); os.IsNotExist(err) {
		panic("File not found")
	}

	plaintext, err := ioutil.ReadFile("teste.txt")

	if err != nil {
		panic(err.Error())
	}

	key := []byte("teste123")
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
	f, err := os.Create("teste.enc")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}

func decrypt() {
	ciphertext, err := ioutil.ReadFile("teste.enc")

	if err != nil {
		panic(err.Error())
	}

	key := []byte("teste123")
	salt := ciphertext[len(ciphertext)-12:]
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

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	// create a new file for saving the encrypted data.
	f, err := os.Create("teste.dec")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(plaintext))
	if err != nil {
		panic(err.Error())
	}
}