package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/zaz600/brute-force-protector/internal/bruteforceprotector"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	bfProtector := bruteforceprotector.NewBruteForceProtector(10, 100, 1000)
	for i := 0; i < 90000; i++ {
		result := bfProtector.Verify("foo", "password", "127.0.0.1")
		log.Println(i, result)
		time.Sleep(time.Duration(rand.Intn(5000-100+1)+100) * time.Millisecond)
		//if result {
		//	time.Sleep(1 * time.Second)
		//} else {
		//	time.Sleep(1 * time.Second)
		//	//time.Sleep(100 * time.Millisecond)
		//}
	}
}
