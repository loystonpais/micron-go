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

}
