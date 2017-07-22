package magic

import (
	"testing"
	"math/rand"
	"time"
)


var Player1 Player
var Player2 Player
var Player3 Player
var Player4 Player

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestFetchCards(t *testing.T) {

	AllCards, err := FetchCards()
	if err != nil {
		panic(err)
	}

	for _, v := range *AllCards {
		t.Log(v.Name)
	}

}


func CreatePlayers(t *testing.T) {

	Player1 := Player{
		"Hunter",
		nil,
	}

	Player2 := Player{
		"Billy",
		nil,
	}

	Player3 := Player{
		"Jimmy",
		nil,
	}

	Player4 := Player{
		"Bobby",
		nil,
	}

	Player1.GiveCards()
	Player2.GiveCards()
	Player3.GiveCards()
	Player4.GiveCards()


}

func TestPassCardsOut(t *testing.T) {

}