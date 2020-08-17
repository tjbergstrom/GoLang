/*

Fixed XOR

Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c

... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965

... should produce:

746865206b696420646f6e277420706c6179

*/


package main
import (
	"fmt"
	"encoding/hex"
	"errors"
)


func main() {
	test_check()
}


func test_check() {
	fmt.Println("\nTesting fixed_xor()...")
	input_1 := []byte("1c0111001f010100061a024b53535009181c")
	input_2 := []byte("686974207468652062756c6c277320657965")
	expected := "746865206b696420646f6e277420706c6179"
	ret, err := fixed_xor(input_1, input_2)
	if err != nil {
		fmt.Println("Error: %s", err)
	}
	if string(ret) != expected {
		fmt.Println("expteced: %s, got: %s", expected, ret)
	} else {
        fmt.Println("Passed test.\n")
    }
}


// Take 2 hex encoded byte slices
// And return a hex encoded byte slice of their xor combinations
func fixed_xor(input_1, input_2 []byte) ([]byte, error) {
	fmt.Println("\nTesting fixed_xor()...")
	decode_1, err := decode_hex_bytes(input_1)
	if err != nil {
		return nil, err
	}
	decode_2, err := decode_hex_bytes(input_2)
	if err != nil {
		return nil, err
	}
	if len(decode_1) != len(decode_2) {
		return nil, errors.New("inputs' lengths are different")
	}
	tmp := make([]byte, len(decode_1))
	for i := 0; i < len(decode_1); i++ {
		tmp[i] = decode_1[i] ^ decode_2[i]
	}
	ret := make([]byte, hex.EncodedLen(len(tmp)))
	hex.Encode(ret, tmp)
	return ret, nil
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
