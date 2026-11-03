// filename: get-imdb-info.go
// from: https://github.com/sraz001/goIMDB

package main

import (
	"fmt"
	"github.com/sraz001/goIMDB"
	"log"
)

func main() {
	m := goIMDB.InitIMDBPie()
	// search for The Matrix title:
	result, err := m.SearchTitle("The Matrix", 1999)
	if err != nil {
		log.Fatal(err)
	}
	// print the first result
	fmt.Printf("Got %d Search Results, the first one is %s - %s (%d)\n", len(result), result[0].ID, result[0].Title, result[0].Year)
	// Got 8 Search Results, the first one is tt0133093 - The Matrix (1999)

	// get Basic title information
	title, err := m.GetTitle(result[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Plot of %s is [ %s ]\n", title.ID, title.PlotOutline)
	// Plot of tt0133093 is [ When a beautiful stranger leads computer hacker Neo to a forbidding underworld, he discovers the shocking truth--the life he knows is the elaborate deception of an evil cyber-intelligence. ]

	// get the cast of the title
	credits, _ := title.GetTitleCredits() // skip error check for readability
	actor := credits["cast"][0]           // get first actor
	fmt.Printf("Actor %s is %s with Image url: %s", actor.ID, actor.Name, actor.Image.Url)
	// Actor nm0000206 is Keanu Reeves with Image url: https://m.media-amazon.com/images/M/MV5BNGJmMWEzOGQtMWZkNS00MGNiLTk5NGEtYzg1YzAyZTgzZTZmXkEyXkFqcGdeQXVyMTE1MTYxNDAw._V1_.jpg

}

