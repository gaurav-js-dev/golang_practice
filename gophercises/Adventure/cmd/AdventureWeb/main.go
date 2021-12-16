package main

import (
	"flag"
	"fmt"
	adventure "golang_practice/gophercises/Adventure"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 5500, "The Port to start Choose Your own adventure")
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := adventure.JsonStory(f)

	if err != nil {
		panic(err)
	}

	// d := json.NewDecoder(f)
	// var story adventure.Story
	// if err := d.Decode(&story); err != nil {
	// 	panic(err)
	// }

	h := adventure.NewHandler(story)

	fmt.Printf("starting server at %d\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
