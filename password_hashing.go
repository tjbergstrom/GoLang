// pass_hash.go
// How to salt and hash a password
//
// go run password_hashing.go

package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func hash_from_pwd(pwd string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
    return string(bytes), err
}

func check_hash(pwd, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
    return err == nil
}

func get_input() (string, error) {
    var pwd string
    fmt.Scanln(&pwd)
    return pwd, nil
}

func check_input(pwd string) bool {
    if len(pwd) < 6 {
        fmt.Println("Failed, pwd must be > 6 chars")
        return false
    }
    return true
}

func main() {
    var pwd string
    fails := 0
    for true && (fails < 3) {
        fmt.Println("Enter a password: ")
        pwd, _ := get_input()
        if !check_input(pwd) {
            fails += 1
            continue
        } else {
            break
        }
    }
    if fails == 3 {
        fmt.Println("Failed to make a password")
        return
    }

    hash, _ := hash_from_pwd(pwd)

    fmt.Println("Password:", pwd)
    fmt.Println("Salted hash: ", hash)
    match := check_hash(pwd, hash)
    fmt.Println("Match: ", match)
}



