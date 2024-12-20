package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"

)

func main() {

	urls := make(map[string]string)

	var longUrl string
	fmt.Scan(&longUrl)
	hash := sha256.New()
	hash.Write([]byte(longUrl))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	urls[longUrl] = hashString[:6]

	fmt.Print(longUrl,":",urls[longUrl])

}