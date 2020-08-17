/*

Convert hex to base64

*/


package main
import (
	"fmt"
	"encoding/base64"
	"encoding/hex"
)


func main() {
	test_check()
	fmt.Println("Input hex to convert it ")
    var input []byte
    fmt.Scanln(&input)
    result, err := hex_to_base64(input)
	if err != nil {
		fmt.Println("err: %s", err)
	}
	fmt.Println("\nHere is the result in base64: ")
	fmt.Println(string(result))
}


func test_check() {
	fmt.Println("\nTesting hex_to_base64()...")
    given_input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    expected_output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
    result, err := hex_to_base64(given_input)
	if err != nil {
		fmt.Println("Error: %s", err)
	}
	if string(result) != expected_output {
		fmt.Println("expected: %s, got: %s", expected_output, string(result))
	} else {
        fmt.Println("Passed test.\n")
    }
}


func hex_to_base64(hexBytes []byte) ([]byte, error) {
	decoded, err := decode_hex_bytes(hexBytes)
	if err != nil {
		return nil, err
	}
	ret := encode_to_base64(decoded)
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


func encode_to_base64(input []byte) []byte {
	ret := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(ret, input)
	return ret
}



//
