package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Hump", "Ingrid"}},
		{Title: "sablanca", Year: 1922, Color: true, Actors: []string{"sasHump", "Iasfngrid"}},
		{Title: "aßdadCasablanca", Year: 2942, Color: false, Actors: []string{"Humpasdasdads", "Isangasrasdasfid"}},
	}

	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed %s", err)
	}
	fmt.Printf("%s\n", data)

}
