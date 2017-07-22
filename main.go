package magic

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"math/rand"
)


var AllCards MagicCards
var AvailableCards MagicCards

type Player struct {
	Name string
	Cards *[]MagicCards
}

func FetchCards() (*MagicCards, error) {
	response, err := http.Get("https://api.deckbrew.com/mtg/cards")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(contents, &AllCards)
		if err != nil {
			return nil, err
		}
	}
	return &AllCards, nil

}


func (p *Player) GiveCards() {

	//var cardArray []MagicCards

	for i:=0; i<7; i++ {
		//
		//random := AllCards[rand.Intn(len(AllCards))]
		//
		//cardArray = append(cardArray, random)

	}
}


type MagicCards []struct {
	Name string `json:"name"`
	ID string `json:"id"`
	URL string `json:"url"`
	StoreURL string `json:"store_url"`
	Types []string `json:"types"`
	Subtypes []string `json:"subtypes,omitempty"`
	Colors []string `json:"colors,omitempty"`
	Cmc int `json:"cmc"`
	Cost string `json:"cost"`
	Text string `json:"text"`
	Power string `json:"power,omitempty"`
	Toughness string `json:"toughness,omitempty"`
	Formats struct {
	} `json:"formats"`
	Editions []struct {
		Set string `json:"set"`
		SetID string `json:"set_id"`
		Rarity string `json:"rarity"`
		Artist string `json:"artist"`
		MultiverseID int `json:"multiverse_id"`
		Flavor string `json:"flavor"`
		Number string `json:"number"`
		Layout string `json:"layout"`
		Price struct {
			Low int `json:"low"`
			Median int `json:"median"`
			High int `json:"high"`
		} `json:"price"`
		URL string `json:"url"`
		ImageURL string `json:"image_url"`
		SetURL string `json:"set_url"`
		StoreURL string `json:"store_url"`
		HTMLURL string `json:"html_url"`
	} `json:"editions"`
	Supertypes []string `json:"supertypes,omitempty"`
}