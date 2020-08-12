// compute_hash.go
// Some cryptographic hash functions

package main

import (
    "os"
    "fmt"
    "crypto/sha256"
    "crypto/md5"
)

func main() {
    if len(os.Args) != 2 {
        return
    }
    input := os.Args[1]

    hash_sha256 := sha256.Sum256([]byte(input))
    hash_md5 := md5.Sum([]byte(input))

    fmt.Printf("The sha256 hash is %x\n", hash_sha256)
    fmt.Printf("The md5 hash is %x\n", hash_md5)
}
