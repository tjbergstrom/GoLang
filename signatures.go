// signatures.go
// Digital signatures show that a hash hasn't been modified
// When you make a hash, you can sign it with a private key,
// And then verify it with your public key
//
// go run signatures.go


package main
import (
	"os"
    "fmt"
	"bufio"
    "crypto"
    "crypto/rsa"
    "crypto/sha1"
    //"crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "encoding/base64"
)


var rawPubKey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvtjdLkS+FP+0fPC09j25\ny/PiuYDDivIT86COVedvlElk99BBYTrqNaJybxjXbIZ1Q6xFNhOY+iTcBr4E1zJu\ntizF3Xi0V9tOuP/M8Wn4Y/1lCWbQKlWrNQuqNBmhovF4K3mDCYswVbpgTmp+JQYu\nBm9QMdieZMNry5s6aiMA9aSjDlNyedvSENYo18F+NYg1J0C0JiPYTxheCb4optr1\n5xNzFKhAkuGs4XTOA5C7Q06GCKtDNf44s/CVE30KODUxBi0MCKaxiXw/yy55zxX2\n/YdGphIyQiA5iO1986ZmZCLLW8udz9uhW5jUr3Jlp9LbmphAC61bVSf4ou2YsJaN\n0QIDAQAB\n-----END PUBLIC KEY-----"
var rawSignature = "c2pkYWpuY2sgZmphbm9panF3b2lqYWRvbmFzbWQgc2EsbWMgc2FuZHBvZHA5cTN1cjA5M3Vyajg4OUoocHEqaDlIUkZKU0ZLQkZPSDk4"
var message = "authenticmessage"

func main() {
    fmt.Println("Enter your public key")
    public_key := get_input()
    raw_public_key := &pem.Block {
        Type: "MESSAGE",
        Bytes: []byte(public_key),
    }
	if err := pem.Encode(os.Stdout, raw_public_key); err != nil {
		panic(err.Error())
	}

    block, _ := pem.Decode([]byte(rawPubKey))
    if block == nil {
        fmt.Println("Invalid PEM Block")
        return
    }

    block = raw_public_key

    key, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        fmt.Println(err)
        return
    }

    pubKey := key.(*rsa.PublicKey)

    signature, err := base64.StdEncoding.DecodeString(rawSignature)
    if err != nil {
        fmt.Println(err)
        return
    }

    hash := sha1.Sum([]byte(message))

    err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, hash[:], signature)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Successfully verified message with signature and public key")
    return

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
