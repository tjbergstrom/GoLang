/*

Implement repeating-key XOR

Here is the opening stanza of an important work of the English language:

Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal

Encrypt it, under the key "ICE", using repeating-key XOR.

In repeating-key XOR, you'll sequentially apply each byte of the key; the first byte of plaintext will be XOR'd against I, the next C, the next E, then I again for the 4th byte, and so on.

It should come out to:

0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

Encrypt a bunch of stuff using your repeating-key XOR function. 

*/


package main
import (
	"fmt"
	"encoding/hex"
)


func main() {
	test_check()
}


func test_check() {
	fmt.Println("\nTesting encrypt_repeating_xor()...")
	input := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	result, err := encrypt_repeating_xor(input, key)
	if err != nil {
		fmt.Println("Error: %s", err)
	}
	if string(result) != expected {
		fmt.Println("Expected: %s, got: %s", expected, string(result))
	} else {
        fmt.Println("Passed test.\n")
	}
}


// Encrypt a string with a key ~ return a hex encoded string
func encrypt_repeating_xor(input, key []byte) ([]byte, error) {
	encrypt_bytes := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		encrypt_bytes[i] = input[i] ^ key[i%len(key)]
	}
	ret := make([]byte, hex.EncodedLen(len(encrypt_bytes)))
	hex.Encode(ret, encrypt_bytes)
	return ret, nil
}



//
