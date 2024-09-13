package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	// create new dal
	dal, _ := newDal("db.db", os.Getpagesize())

	// create new page
	p := dal.allocateEmptyPage()
	p.num = dal.getNextPage()

	//copy(p.data[:], "This is my initial test string. Want to see it!")
	//copy(p.data[:], generateRandomString(dal.pageSize+100))
	//p.data += "This is my initial test string. Want to see it!"

	p.data = generateRandomString(dal.pageSize + 100)
	_ = dal.writePage(p)
}

func generateRandomString(length int) []byte {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
		"!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	rand.Seed(time.Now().UnixNano())
	randomString := make([]byte, length)
	var counter uint64 = 0

	for i := range randomString {
		randomString[i] = byte(counter)
		counter++
	}
	return randomString
}
