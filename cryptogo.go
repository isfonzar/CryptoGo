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
	salt := []byte("salt")

	dk := pbkdf2.Key(key, salt, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure. @todo

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
	salt := []byte("salt")

	dk := pbkdf2.Key(key, salt, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

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