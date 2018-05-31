package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(letterToRegional(strings.Join(os.Args[1:], " ")))
}

func letterToRegional(s string) string {
	val := []rune{}
	for _, r := range s {
		switch {
		case '?' == r:
			val = append(val, '❓')
		case r >= 'a' && r <= 'z':
			val = append(val, '🇦'+(r-'a'))
		case r >= 'A' && r <= 'Z':
			val = append(val, '🇦'+(r-'A'))
		default:
			val = append(val, r)
		}
		val = append(val, ' ')
	}
	return string(val)
}
