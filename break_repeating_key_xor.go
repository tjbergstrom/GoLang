/*

There's a file here. It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

*/


package main
import (
	"fmt"
	//"bytes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	//"sort"
)


const decrypt_this_file = "./decrypt_this_file.txt"


func main() {
	test_check()
}


func test_check() {
	fmt.Println("\nTesting decrypt_file()...")
	result, err := decrypt_file(decrypt_this_file)
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
	unencrypted_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(f)))
	_, err = base64.StdEncoding.Decode(unencrypted_bytes, f)
	if err != nil {
		return nil, err
	}
	var ret []byte

	/*
	// todo
	*/

	return ret, nil
}


func hamming_dist(input_1, input_2 []byte) int {
	var dist int
	for i := 0; i < len(input_1); i ++ {
		for j := 0; j < 8; j++ {
			if input_1[i] & (1<<uint(j)) != input_2[i] & (1<<uint(j)) {
				dist++
			}
		}
	}
	return dist
}


func xor_cipher(key_len int, chunks [][]byte) ([][]byte, []byte, error) {
	unencoded := make([][]byte, key_len)
	for i, ch := range chunks {
		hex_bytes := make([]byte, hex.EncodedLen(len(ch)))
		hex.Encode(hex_bytes, ch)
		result, _, err := single_byte_xor_cipher(hex_bytes)
		if err != nil {
			return nil, nil, err
		}
		unencoded[i] = result
	}
	return unencoded, nil, nil
}


type key_length struct {
	length int
	strength float64
}


func get_key_length(input []byte) []key_length {
	key_len := make([]key_length, 0)

	/*
	// todo
	*/

	return key_len
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


func decode_hex_bytes(hexBytes []byte) ([]byte, error) {
	ret := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(ret, hexBytes)
	if err != nil {
		return nil, err
	}
	return ret, nil
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



//
