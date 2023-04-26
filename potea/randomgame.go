package main

import (
	"fmt"
	"math/rand"
)

func randomGame() {
	var inNum int
	ranNum := rand.Intn(100)
	fmt.Println(ranNum)

MainRoop:
	for i := 1; true; i++ {
		_, err := fmt.Scanln(&inNum)
		if err != nil {
			fmt.Println("입력 오류")
		} else {
			switch {
			case inNum == ranNum:
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수", i)
				break MainRoop
			case inNum < ranNum:
				fmt.Println("입력하신 숫자가 더 작습니다.")
			case inNum > ranNum:
				fmt.Println("입력하신 숫자가 더 큽니다.")
			}
		}
	}

}

func main() {
	randomGame()
}
