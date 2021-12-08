package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%d\n", &input)
	fmt.Scanf("%d\n", &delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	// alphabet := []rune(alphabetStr)
	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}

		// if strings.IndexRune(alphabetStr, ch) >= 0 {
		// 	ret = ret + string(rotate(ch, delta, alphabet))
		// } else {
		// 	ret = ret + string(rotate(ch, delta, alphabet))
		// }
	}
	fmt.Println(ret)
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)
	return key[idx]
	// for i, r := range key {
	// 	if r == s {
	// 		idx = i
	// 		break
	// 	}
	// }
	// idx := -1
	// for i, r := range key {
	// 	if r == s {
	// 		idx = i
	// 		break
	// 	}
	// }

	// for i := 0; i < delta; i++ {
	// 	idx++
	// 	if idx >= len(key) {
	// 		idx = 0
	// 	}
	// }

}
