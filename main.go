package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // math.rand -> rand.Seed() 함수는  1.20 에서 Deprecated 되었다.

	n := rand.Intn(100)
	fmt.Println(n)
}
