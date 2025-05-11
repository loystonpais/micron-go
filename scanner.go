package main

type Char struct {
	Char rune
	N    uint
}

func NewChar(char rune, n uint) Char {
	return Char{char, n}
}

func CharsFromSource(source string) []Char {
	runes := []rune(source)

	var chars []Char

	for n, char := range runes {
		chars = append(chars, NewChar(char, uint(n)))
	}

	return chars
}
