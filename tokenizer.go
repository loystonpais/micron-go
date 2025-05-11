package main

import (
	"fmt"
	"strconv"
)

type Token interface {
	Display
	isToken()
}

type StrToken struct {
	Str string
}

func (StrToken) isToken() {}
func (t StrToken) Repr() string {
	return fmt.Sprintf("Str(%s)", t.Str)
}

type IntToken struct {
	Int int
}

func (IntToken) isToken() {}
func (t IntToken) Repr() string {
	return fmt.Sprintf("Int(%d)", t.Int)
}

type IdnToken struct {
	Str string
}

func (IdnToken) isToken() {}
func (t IdnToken) Repr() string {
	return fmt.Sprintf("Idn(%s)", t.Str)
}

type TilToken struct{}

func (TilToken) isToken() {}
func (t TilToken) Repr() string {
	return "Til"
}

type ColToken struct{}

func (ColToken) isToken() {}
func (t ColToken) Repr() string {
	return "Col"
}

type SmiToken struct{}

func (SmiToken) isToken() {}
func (t SmiToken) Repr() string {
	return "Smi"
}

type DotToken struct{}

func (DotToken) isToken() {}
func (t DotToken) Repr() string {
	return "Dot"
}

type EolToken struct{}

func (EolToken) isToken() {}
func (t EolToken) Repr() string {
	return "Eol"
}

type DolToken struct{}

func (DolToken) isToken() {}
func (t DolToken) Repr() string {
	return "Dol"
}

type QueToken struct{}

func (QueToken) isToken() {}
func (t QueToken) Repr() string {
	return "Que"
}

type EqlToken struct{}

func (EqlToken) isToken() {}
func (t EqlToken) Repr() string {
	return "Eql"
}

type NotToken struct{}

func (NotToken) isToken() {}
func (t NotToken) Repr() string {
	return "Not"
}

type HshToken struct{}

func (HshToken) isToken() {}
func (t HshToken) Repr() string {
	return "Hsh"
}

type TokenInfo struct {
	Token Token
	Start uint
	End   uint
}

func (t TokenInfo) Repr() string {
	return fmt.Sprintf("TokenInfo(Token=%s, Start=%d, End=%d)", t.Token.Repr(), t.Start, t.End)
}

type TokenizerErrorInfo struct {
	Start uint
	End   uint
	Msg   string
}

type TokenizerError interface {
	error
	isTokenizerError()
}

type SyntaxErrorTokenizerError struct {
	ErrorInfo TokenizerErrorInfo
}

func (SyntaxErrorTokenizerError) isTokenizerError() {}
func (e SyntaxErrorTokenizerError) Error() string {
	return e.ErrorInfo.Msg
}

func Tokenize(chars []Char) ([]TokenInfo, TokenizerError) {
	var tokens []TokenInfo
	var i uint = 0

	for i < uint(len(chars)) {
		switch chars[i].Char {
		case '~':
			tokens = append(tokens, TokenInfo{TilToken{}, i, i})
		case ':':
			tokens = append(tokens, TokenInfo{ColToken{}, i, i})
		case '.':
			tokens = append(tokens, TokenInfo{DotToken{}, i, i})
		case '\n':
			tokens = append(tokens, TokenInfo{EolToken{}, i, i})
		case '?':
			tokens = append(tokens, TokenInfo{DolToken{}, i, i})
		case '=':
			tokens = append(tokens, TokenInfo{EqlToken{}, i, i})
		case '!':
			tokens = append(tokens, TokenInfo{NotToken{}, i, i})
		case '#':
			tokens = append(tokens, TokenInfo{HshToken{}, i, i})
		case ' ':
		case '\t':
		case '"':
			j := i
			i += 1
			var temp []rune

			if !(i < uint(len(chars))) {
				return nil, SyntaxErrorTokenizerError{
					TokenizerErrorInfo{
						j,
						j,
						"EOL while scanning for the string literal",
					},
				}
			}

			for chars[i].Char != '"' {
				if chars[i].Char == '\\' {
					if i+1 < uint(len(chars)) {
						char := chars[i+1]
						var newChar rune
						switch char.Char {
						case 'n':
							newChar = '\n'
						case 't':
							newChar = '\t'
						default:
							newChar = char.Char
						}

						temp = append(temp, newChar)
					} else {
						return nil, SyntaxErrorTokenizerError{
							TokenizerErrorInfo{
								i,
								i,
								"EOF while scanning for the escape sequence",
							},
						}
					}

					i += 1
				} else {
					temp = append(temp, chars[i].Char)
				}
				i += 1

				if !(i < uint(len(chars))) {
					return nil, SyntaxErrorTokenizerError{
						TokenizerErrorInfo{
							j,
							j,
							"EOF while scanning for the string literal",
						},
					}
				}
			}

			collected := string(temp)
			tokens = append(tokens, TokenInfo{StrToken{collected}, j, i + 1})

		case '[':
			j := i
			i += 1

			if !(i < uint(len(chars))) {
				return nil, SyntaxErrorTokenizerError{
					TokenizerErrorInfo{
						i,
						j,
						"EOF while scanning for the comment literal",
					},
				}
			}

			for chars[i].Char != ']' {
				i += 1
				if !(i < uint(len(chars))) {
					return nil, SyntaxErrorTokenizerError{
						TokenizerErrorInfo{
							j,
							j,
							"EOF while scanning for the comment literal",
						},
					}
				}
			}

		default:
			char := chars[i].Char
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_' {

				j := i
				var temp []rune

				for i < uint(len(chars)) {
					char := chars[i].Char
					if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_' {
						temp = append(temp, char)
					} else {
						break
					}
					i += 1
				}

				tokens = append(tokens, TokenInfo{
					IdnToken{
						string(temp),
					},
					j,
					i,
				})

				i -= 1

			} else if (char >= '0' && char <= '9') || (char == '-') {
				j := i
				i += 1
				temp := []rune{chars[j].Char}

				for i < uint(len(chars)) {
					char := chars[i].Char
					if char >= '0' && char <= '9' {
						temp = append(temp, char)
					} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
						return nil, SyntaxErrorTokenizerError{
							TokenizerErrorInfo{
								j,
								i,
								"Invalid number literal",
							},
						}
					} else {
						break
					}

					i += 1
				}
				num, err := strconv.Atoi(string(temp))

				if err != nil {
					return nil, SyntaxErrorTokenizerError{
						TokenizerErrorInfo{
							j,
							i,
							"Invalid int",
						},
					}
				}

				tokens = append(tokens, TokenInfo{IntToken{num}, j, i})
				i -= 1
			} else {
				return nil, SyntaxErrorTokenizerError{
					TokenizerErrorInfo{
						i,
						i,
						"Invalid character",
					},
				}
			}
		}
		i += 1
	}

	return tokens, nil
}
