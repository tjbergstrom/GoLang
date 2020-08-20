// rsa.go


package main
import (
	"bufio"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "os"
)


func main() {
    fmt.Println("Enter a message")
    secret_msg := get_input()

    // Generate a private key
    priv_key, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println(err.Error)
        os.Exit(1)
    }

    // extract a public key from a private key
    pub_key := priv_key.PublicKey
    encrypted_msg := EncryptOAEP(secret_msg, pub_key);
    fmt.Println("cipher text:", encrypted_msg)
    DecryptOAEP(encrypted_msg, *priv_key)
}


func EncryptOAEP(secret_msg string, pub_key rsa.PublicKey)  (string) {
    label := []byte("OAEP Encrypted")
    encrypt := rand.Reader
    cipher_text, err := rsa.EncryptOAEP(sha256.New(), encrypt, &pub_key, []byte(secret_msg), label)
    if err != nil {
        fmt.Fprintf(os.Stderr, "encryption error: %s\n", err)
        return "encryption error";
    }
    return base64.StdEncoding.EncodeToString(cipher_text)
}


func DecryptOAEP(cipherText string, privKey rsa.PrivateKey)  (string) {
    ct,_ := base64.StdEncoding.DecodeString(cipherText)
    label := []byte("OAEP Encrypted")
    rng := rand.Reader
    plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
    if err != nil {
        fmt.Fprintf(os.Stderr, "dencryption error: %s\n", err)
        return "dencryption error";
    }
    fmt.Printf("original message: %s\n", string(plaintext))
    return string(plaintext)
}


func get_input() string {
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
    return input
}



//
