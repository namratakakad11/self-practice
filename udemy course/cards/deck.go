package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "heart", "clubs"}

	cardValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {

	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {

	bs, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)

}


func (d deck) shuffle(){

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
	newpos :=	r.Intn(len(d)-1)
	d[i], d[newpos] = d[newpos], d[i]
	}

}
