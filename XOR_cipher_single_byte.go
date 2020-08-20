/*

Single-byte XOR cipher

The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext.
Character frequency is a good metric.
Evaluate each output and choose the one with the best score.

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
	fmt.Println("\nTesting single_byte_xor_cipher()...")
	hex_bytes := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	result, _, err := single_byte_xor_cipher(hex_bytes)
	if err != nil {
		fmt.Println("Error: %s", err)
	}
	fmt.Println("Decryption test:", string(result))
}


// Take an encoded message and decode it into a string
func single_byte_xor_cipher(coded_msg []byte) ([]byte, int, error) {
	bytes, err := decode_hex_bytes(coded_msg)
	if err != nil {
		return nil, 0,  err
	}
	var result []byte
	var score int
	for i := 0; i < 256; i++ {
		tmp_result := make([]byte, len(bytes))
		var tmp_score int
		for j := 0; j < len(bytes); j++ {
			ch := bytes[j] ^ byte(i)
			tmp_score += char_weight(ch)
			tmp_result[j] = ch
		}
		if tmp_score > score {
			result = tmp_result
			score = tmp_score
		}
		tmp_score = 0
	}
	return result, score, nil
}


func char_weight(ch byte) int {
	weight_map := map[byte]int{
		byte('U'): 2, byte('u'): 2,
		byte('L'): 3, byte('l'): 3,
		byte('D'): 4, byte('d'): 4,
		byte('R'): 5, byte('r'): 5,
		byte('H'): 6, byte('h'): 6,
		byte('S'): 7, byte('s'): 7,
		byte(' '): 8,
		byte('N'): 9, byte('n'): 9,
		byte('I'): 10, byte('i'): 10,
		byte('O'): 11, byte('o'): 11,
		byte('A'): 12, byte('a'): 12,
		byte('T'): 13, byte('t'): 13,
		byte('E'): 14, byte('e'): 14,
	}
	return weight_map[ch]
}


func decode_hex_bytes(hexBytes []byte) ([]byte, error) {
	ret := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(ret, hexBytes)
	if err != nil {
		return nil, err
	}
	return ret, nil
}



//
