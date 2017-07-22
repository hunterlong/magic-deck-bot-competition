package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)


var AllCards MagicCards

func main() {
	response, err := http.Get("https://api.deckbrew.com/mtg/cards")
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(contents, &AllCards)
		if err != nil {
			panic(err)
		}
		fmt.Println("Total Magic Cards: ",len(AllCards))
	}

	for _, v := range AllCards {
		fmt.Println(v.Name)
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