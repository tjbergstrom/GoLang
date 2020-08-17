/*

There's a file here. It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

*/


import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"sort"
)


const decrypt_this_file = "./decrypt_this_file.txt"


func main() {
	test_check()
}


func test_check() {
	fmt.Println("\nTesting decrypt_file()...")
	result, err := XorDecryptFile(decrypt_this_file)
	if err != nil {
		fmt.Println("Error %s", err)
	}
	fmt.Println("Decryption test:", string(result))
}


func decrypt_file(file string) ([]byte, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// todo

	var ret []byte

	// todo

	return ret, nil
}



//
