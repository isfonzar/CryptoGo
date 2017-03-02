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
)

func main() {

	encrypt()

	decrypt()

}

func encrypt() {
	plaintext, err := ioutil.ReadFile("teste.txt")

	if err != nil {
		panic(err.Error())
	}

	key := []byte("teste123")
	salt := []byte("salt")

	dk := pbkdf2.Key(key, salt, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

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
	plaintext, err := ioutil.ReadFile("teste.enc")

	if err != nil {
		panic(err.Error())
	}

	key := []byte("teste123")
	salt := []byte("salt")

	dk := pbkdf2.Key(key, salt, 4096, 32, sha1.New)

	fmt.Println(dk)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// create a new file for saving the encrypted data.
	f, err := os.Create("teste.dec")
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(f, bytes.NewReader(ciphertext))
	if err != nil {
		panic(err.Error())
	}
}