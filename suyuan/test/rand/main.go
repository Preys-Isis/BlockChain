package main

import (
	"crypto/sha1"
	"fmt"
	createrand "test/rand/createRand"
)

func main() {
	Teststr := createrand.RandomString(256)
	Test_hash := sha1.New().Sum([]byte(Teststr))

	fmt.Printf("随机哈希值：%s\n", string(Test_hash))
}
