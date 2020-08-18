// encrypt_decrypt.go
// How to encrypt and decrypt
// Enter input to have it encrypted
//
// go run encryot_decrypt.go


package main
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)


func main() {
	fmt.Println("Enter a string to encrypt it ")
    var input string
    fmt.Scanln(&input)
	cipher_text := encrypt([]byte(input), "pass_phrase")
	fmt.Printf("Encrypted: %x\n", cipher_text)
	plain_text := decrypt(cipher_text, "pass_phrase")
	fmt.Printf("Decrypted: %s\n", plain_text)
}


func hash_from_str(key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}


func encrypt(data []byte, pass_phrase string) []byte {
	// create a new block cipher based on the hashed passphrase
	block_cipher, _ := aes.NewCipher([]byte(hash_from_str(pass_phrase)))
	gcm, err := cipher.NewGCM(block_cipher)
	if err != nil {
		panic(err.Error())
	}
	// create a nonce of length specified by Galois Counter Mode
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	// make sure the decryption nonce matches the encryption nonce
	cipher_text := gcm.Seal(nonce, nonce, data, nil)
	return cipher_text
}


func decrypt(data []byte, pass_phrase string) []byte {
	key := []byte(hash_from_str(pass_phrase))
	block_cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block_cipher)
	if err != nil {
		panic(err.Error())
	}
	len_nonce := gcm.NonceSize()
	nonce, cipher_text := data[:len_nonce], data[len_nonce:]
	plain_text, err := gcm.Open(nil, nonce, cipher_text, nil)
	if err != nil {
		fmt.Println("Passphrase was incorrect!")
		panic(err.Error())
	}
	return plain_text
}


func encryptFile(file_name string, data []byte, pass_phrase string) {
	f, _ := os.Create(file_name)
	defer f.Close()
	f.Write(encrypt(data, pass_phrase))
}


func decryptFile(file_name string, pass_phrase string) []byte {
	data, _ := ioutil.ReadFile(file_name)
	return decrypt(data, pass_phrase)
}


