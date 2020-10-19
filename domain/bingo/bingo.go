package bingo

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const DRAW = "D"
const RESTART = "R"
const STOP = "S"

type Bingo struct {
	globeNumbers []int
	drawnNumbers []int
}

func NewBingo() *Bingo {
	return &Bingo{
		globeNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
			41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
			51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
			61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
			71, 72, 73, 74, 75},
	}
}

func (b Bingo) Start() {
	fmt.Println("Enter letter 'D' to Draw a number")
	fmt.Println("Enter letter 'R' to Restart the draw")
	fmt.Println("Enter letter 'S' to Stop the system")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	command := strings.Replace(text, "\n", "", 1)

	b.Command(command)
}

func (b Bingo) Command(command string) {
	switch command {
	case DRAW:
		b.Draw()
		b.ShowDrawnNumbers()
		b.Start()
	case RESTART:
		b.Restart()
	case STOP:
		b.Stop()
	default:
		err := errors.New("invalid command")
		panic(err)
	}
}

func (b *Bingo) Draw() {
	// Do the draw
	num := rand.Intn(len(b.globeNumbers))

	// Show the drawn number
	fmt.Printf("The number drawn was: %d\n", b.globeNumbers[num])

	// Add number to Numbers Drawn List
	b.drawnNumbers = append(b.drawnNumbers, b.globeNumbers[num])

	// Remove the number from Globe Numbers
	b.globeNumbers = append(b.globeNumbers[:num], b.globeNumbers[num+1:]...)
}

func (b Bingo) ShowDrawnNumbers() {
	fmt.Printf("Numbers Drawn: %v\n", b.drawnNumbers)
}

func (b *Bingo) Restart() {
	fmt.Println("Restarting the draw...")
	bingo := NewBingo()
	b.globeNumbers = bingo.globeNumbers
	b.drawnNumbers = bingo.drawnNumbers
	b.Start()
}

func (b Bingo) Stop() {
	fmt.Println("Bye bye")
	os.Exit(0)
}
