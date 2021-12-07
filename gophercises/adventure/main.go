package main

import (
	"encoding/json"
	"flag"
	"fmt"
	adventure "go/gophercises/adventure"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the choose your own adventure story")

	flag.Parse()

	fmt.Printf("Using the story in %s.\n ", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	d := json.NewDecoder(f)
	var story adventure.Story

	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", story)
}
