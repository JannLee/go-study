package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func main() {
	//	rand.Seed(time.Now().UnixNano())
	randnum := rand.Intn(100)
	var inputnum int
	cnt := 1
	for {
		fmt.Println("숫자를 입력하세요~!")
		_, err := fmt.Scanln(&inputnum)
		if err != nil {
			fmt.Println("숫자만 입력하라고오~~!!!!")
			stdin.ReadString('\n')
		} else {
			if inputnum > randnum {
				fmt.Println("입력한 숫자가 더 큽니다~~ 다시 입력하세요")
			} else if inputnum < randnum {
				fmt.Println("입력한 숫자가 더 작아요~~ 다시 입력하세요")
			} else {
				fmt.Println("정답입니다.", cnt, "번 만에 맞췄습니다~~!")
				break
			}
			cnt++
		}
	}
}
