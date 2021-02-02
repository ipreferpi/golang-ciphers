package main

import (
	"strings"
)
import "fmt"

var alphabet = `abcdefghijklmnopqrstuvwxyz`

func main() {
	fmt.Println(rot13(alphabet))
	fmt.Println(caesar(alphabet, 1))
	fmt.Println(substitution(alphabet, alphabet))
	fmt.Println(vigenere(`defendtheeastwallofthecastle`, `fortification`))
}

func shift(r rune, i rune) rune {
	return 'a' + (r - 'a' + i) % 26
}

func substitution(clear string, key string) string {
	var out = []rune(``)
	for _, c := range clear{
		sub := rune(key[strings.IndexRune(alphabet, c)])
		out = append(out, sub)
	}

	return string(out)
}

func sanitize(input string) string {
	output := strings.ToLower(input)
	output = strings.TrimSpace(input)
	return output
}

func rot13(clear string) string{
	out := strings.Map(func(r rune) rune {return shift(r, 13)}, clear)
	return out
}

func caesar(clear string, shiftval rune) string{
	out := strings.Map(func(r rune) rune {return shift(r, shiftval)}, clear)
	return out
}

func vigenere (clear string, key string) string {
	keyindex := 0
	out := strings.Map(func(r rune) rune {
		shiftval := rune(strings.IndexRune(alphabet, []rune(key)[keyindex % len(key)]))

		out := shift(r, shiftval)
		keyindex++

		return out
	}, clear)

	return out
}

