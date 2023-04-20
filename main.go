package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n') // 에러 발생시 입력 String 비우기
	}
	return n, err
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // math.rand -> rand.Seed() 함수는  1.20 에서 Deprecated 되었다.
	readingNumber := rand.Intn(100)

	cnt := 1
	for {
		fmt.Println("숫자값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			if n > readingNumber {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < readingNumber {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", cnt)
				break
			}
			cnt++
		}
	}

}
