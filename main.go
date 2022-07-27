package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Game struct {
	Name string `json:"description"`
}

func main() {
	allGames := make([]Game, 0)

	newCol := colly.NewCollector(
		colly.AllowedDomains("store.steampowered.com", "https://store.steampowered.com/"),
	)

	newCol.OnHTML(".tab_item_name", func(element *colly.HTMLElement){
		gameName := element.Text

		game := Game{
			Name: gameName,
		}

		allGames = append(allGames, game)
	})
	newCol.OnRequest(func(request *colly.Request){
		fmt.Println("Visiting", request.URL.String())
	})

	newCol.Visit("https://store.steampowered.com/")

	encode := json.NewEncoder(os.Stdout)
	encode.SetIndent("", " ")
	encode.Encode(allGames)
}

