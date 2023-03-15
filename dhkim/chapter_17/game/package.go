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
}

func NewGame() *Game {
	seed := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(seed).Intn(100)
	return &Game{randomNumber}
}

func (g *Game) Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("숫자를 입력하세요: ")
		s, _ := reader.ReadString('\n')
		n, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 0)
		r := g.Check(int(n))
		if r == 0 {
			fmt.Println("정답입니다!")
			break
		} else if r == -1 {
			fmt.Println("입력한 숫자보다 큽니다.")
		} else {
			fmt.Println("입력한 숫자보다 작습니다.")
		}
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
