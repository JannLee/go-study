package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func GetNumber() (int, error) {
	var mynumber int
	_, error := fmt.Scanln(&mynumber)
	if error != nil {
		stdin.ReadString('\n')
	}
	return mynumber, error
}

func main() {

	rand.Seed(time.Now().UnixMicro())
	realNumber := rand.Intn(100)
	cnt := 1
	for {
		fmt.Printf("숫자를 입력해라")
		inputNumber,error := GetNumber()

		if error != nil{
			fmt.Println("에러가 발생. 숫자가 아닌거 같아.")
			//한줄 읽어 들이는걸 빼먹었네요! ㅜㅜ
		}else{
			if realNumber > inputNumber {
				fmt.Println("더 작아")
			}else if realNumber < inputNumber{
				fmt.Println("더 커")
			}else {
				fmt.Println("숫자를 맞췃어.:", inputNumber, realNumber, cnt)
				break
			}
		}

		cnt++

	}
}
