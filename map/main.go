package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4b745",
		"white": "fffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#fffff"

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
