// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

const (
	NumCards = 52
)

const (
	Ace = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	Heart = iota
	Spade
	Diamond
	Club
)

type Deck struct {
	cards []Card
}

type Face int
type Suit int

type Card struct {
	Face Face
	Suit Suit
}


func (f Face) String() string {
	switch f {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "Invalid face card"
	}
}

func (s Suit) String() string {
	switch s {
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Club:
		return "Club"
	case Spade:
		return "Spade"
	default:
		return "Invalid suit name"
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%v %q", c.Face.String(), c.Suit.String())
}

func (d *Deck) deal() Card {
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
}

/*

slice from n to the end
a[n:]
slice from the beginning to n (exclusive, i think)
a[:n]

Append syntax:
[]T = append([]T, T...)
*/
func (d *Deck) returnCard(card Card) {
	cards := make([]Card, 1)
	cards[0] = card
	d.cards = append(cards, d.cards...)
}

/*
func (d* Deck) shuffle() {

}
*/

func (d *Deck) String() string {
	s := ""
	for _, card := range d.cards {
		s = fmt.Sprintf("%s %v\n", s, card)
	}
	return s
}

func getSuit(index int) Suit {
	if index > 51 {
		return -1
	}
	switch (index) / 13 {
	case 0:
		return Heart
	case 1:
		return Diamond
	case 2:
		return Spade
	case 3:
		return Club
	default:
		return -1
	}
}

func NewDeck() *Deck {
	d := &Deck{
		cards: make([]Card, NumCards),
	}
	for index, card := range d.cards {
		card.Face = Face((index)%13)
		card.Suit = getSuit(index)
		d.cards[index] = card
		fmt.Printf("%v\n", card)
	}
	return d
}

func main() {
	fmt.Println("Creating a new deck")
	d := NewDeck()
	card := d.deal()
	fmt.Println()
	fmt.Printf("Drew card: %v\n", card)
	card = d.deal()
	fmt.Println()
	fmt.Printf("Drew card: %v\n", card)
	d.returnCard(card)
	fmt.Printf("%v", d)
}

/*
Your previous Plain Text content is preserved below:

This is just a simple shared plaintext pad, with no execution capabilities.

When you know what language you'd like to use for your interview,
simply choose it from the dropdown in the top bar.

You can also change the default language your pads are created with
in your account settings: https://coderpad.io/settings

Enjoy your interview!

Playing Cards

Suit: Clubs, Diamonds, Spades, Hearts
Face: Ace-Jack, 2-10
Deck: All 52 cards

Card deal(); // removes a card from the front/head of the deck
void returnCard(Card card); // adds a card to the bottom/end of the deck
void shuffle(); // permute cards in the deck such that their order is randomized relative to the previous deck order


*/
