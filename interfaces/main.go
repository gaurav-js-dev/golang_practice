package main

import "fmt"

type bot interface {
	getGreeting() string // to qualify as a bot type there should be a function with getGreeting() name and string return type
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// Very custom logic for generating an English reply
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating an Spanish reply
	return "Hola El Paso El Camino"
}
