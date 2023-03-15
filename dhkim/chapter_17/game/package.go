package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	randomNumber int
	cnt          int
}

func NewGame() *Game {
	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed).Intn(100)
	cnt := 1
	return &Game{randomNumber, cnt}
}

func (g *Game) Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("숫자를 입력하세요: ")
		s, _ := reader.ReadString('\n')
		n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 0)
		if err != nil {
			fmt.Println("숫자가 아닌 값을 입력하셨습니다!!")
			continue
		}

		r := g.Check(int(n))
		if r == 0 {
			fmt.Println("정답입니다!", g.cnt, "번 만에 맞추셨습니다!")
			break
		} else if r == -1 {
			fmt.Println("입력한 숫자보다 큽니다.")
		} else {
			fmt.Println("입력한 숫자보다 작습니다.")
		}
		g.cnt += 1
	}
}

func (g *Game) Check(number int) int {
	if number == g.randomNumber {
		return 0
	} else if number < g.randomNumber {
		return -1
	}
	return 1
}
