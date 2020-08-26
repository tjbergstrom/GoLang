// hash_checksum.go
// See if your checksum matches the file


package main

import (
    "io"
    "os"
    "fmt"
    "crypto/sha256"
)

var filename = "decrypt_this_file.txt"

func main() {
    fmt.Printf("Enter a checksum to check it with the file %s\n", filename)
    input_checksum, _ := get_input()
    fmt.Printf("\n")

    file, err := os.Open(filename)
    if err != nil {
        panic(err.Error)
    }
    defer file.Close()

    hash := sha256.New()
    if _, err := io.Copy(hash, file); err != nil {
        panic(err.Error)
    }
    sum := hash.Sum(nil)

    fmt.Printf("The input hash is %s\n", input_checksum)
    fmt.Printf("The file hash is %x\n", sum)

    if []byte(input_checksum) == []byte(sum) {
        fmt.Printf("Match")
    } else {
        fmt.Printf("No match\n")
    }
}

func get_input() (string, error) {
    var input string
    fmt.Scanln(&input)
    return input, nil
}



// 1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa
