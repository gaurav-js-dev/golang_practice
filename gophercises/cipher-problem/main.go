package main

import (
	"fmt"
)

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, cipher(ch, delta))
		// switch {
		// case strings.IndexRune(alphabetLower, ch) >= 0:
		// 	ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
		// case strings.IndexRune(alphabetUpper, ch) >= 0:
		// 	ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
		// default:
		// 	ret = ret + string(ch)
		// }

		// if strings.IndexRune(alphabetStr, ch) >= 0 {
		// 	ret = ret + string(rotate(ch, delta, alphabet))
		// } else {
		// 	ret = ret + string(rotate(ch, delta, alphabet))
		// }
	}
	fmt.Println(string(ret))
	// fmt.Println('a')
	// fmt.Println('b')
	// fmt.Println('z')

}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

// func rotate(s rune, delta int, key []rune) rune {
// 	idx := strings.IndexRune(string(key), s)
// 	if idx < 0 {
// 		panic("idx < 0")
// 	}
// 	idx = (idx + delta) % len(key)
// 	return key[idx]
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
// }
