package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	c := NewChar('S', 5)

	fmt.Println("Hello, World", c)

	var e1 = FunCallExpr{
		SetFun{
			ValueExpr{StrValue{"Hello!"}},
			FunCallExpr{SetFun{
				ValueExpr{
					StrValue{"Hi!"},
				},
				ValueExpr{
					IntValue{30},
				},
			}},
		},
	}
	fmt.Println(e1.Repr())
	fmt.Println(e1.Code())

	chars := CharsFromSource(`
		a: "gi"
		bruh: lol -234 .g
`)

	tokenInfos, err := Tokenize(chars)

	if err != nil {
		fmt.Println("Error occurred while parsing")
		fmt.Println(err.Error())

		errorType, ok := err.(SyntaxErrorTokenizerError)

		if ok {
			fmt.Printf("SyntaxError at %d..%d in ", errorType.ErrorInfo.Start, errorType.ErrorInfo.End)
		}

	}

	fmt.Println(chars)

	fmt.Println(tokenInfos[3].Repr())

}
