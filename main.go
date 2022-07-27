package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Game struct {
	Name       string `json:"name"`
	Price      string `json:"price"`
}

func main() {

	game := Game{}
	allGames := make([]Game, 0)

	newCol := colly.NewCollector(
		colly.AllowedDomains("store.steampowered.com", "https://store.steampowered.com/"),
	)

	newCol.OnHTML(".tab_item_name", func(element *colly.HTMLElement) {
		game.Name = element.Text
		allGames = append(allGames, game)
	})

	newCol.OnHTML(".discount_final_price", func(element *colly.HTMLElement) {
		game.Price = element.Text
		allGames = append(allGames, game)
	})

	newCol.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	newCol.Visit("https://store.steampowered.com/")

	newCol.OnScraped(func(r *colly.Response){
		allGames = append(allGames, game)
		game = Game{}
	})

	encode := json.NewEncoder(os.Stdout)
	encode.SetIndent("", " ")
	encode.Encode(allGames)
}
