package bingo

import (
	"encoding/json"
	"math/rand"
)

type Card struct {
	B []int `json:"b"`
	I []int `json:"i"`
	N []int `json:"n"`
	G []int `json:"g"`
	O []int `json:"o"`
}

func NewCard() *Card {
	card := &Card{}

	card.B = generateColumn(5, 1, 15)
	card.I = generateColumn(5, 16, 30)
	card.N = generateColumn(4, 31, 45)
	card.G = generateColumn(5, 46, 60)
	card.O = generateColumn(5, 61, 75)

	return card
}

func (c *Card) GetJson() string {
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func generateColumn(size, min, max int) []int {
	var column []int

	for i := 0; i < size; i++ {
		num := generateNumber(column, min, max)
		column = append(column, num)
	}

	return column
}

func generateNumber(column []int, min, max int) int {
	num := rand.Intn(max-min) + min
	if contains(column, num) {
		num = generateNumber(column, min, max)
	}
	return num
}

func contains(data []int, value int) bool {
	for _, v := range data {
		if value == v {
			return true
		}
	}
	return false
}
