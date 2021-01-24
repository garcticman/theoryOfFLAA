package RegularExpressions

import (
	"fmt"
	"testing"
)

func TestTask4(t *testing.T) {
	// syntax abstract tree of this regular expression: (a(|b))*
	regularTree := Repeat{Pattern: Concatenate{
		Left: Literal{Character: 'a'},
		Right: Choose{
			Left:  Empty{},
			Right: Literal{Character: 'b'},
		},
	}}

	handler := Handler{}
	fmt.Println(handler.Matches(regularTree, ""))      //true
	fmt.Println(handler.Matches(regularTree, "a"))     //true
	fmt.Println(handler.Matches(regularTree, "ab"))    //true
	fmt.Println(handler.Matches(regularTree, "aba"))   //true
	fmt.Println(handler.Matches(regularTree, "abab"))  //true
	fmt.Println(handler.Matches(regularTree, "abaab")) //true
	fmt.Println(handler.Matches(regularTree, "abba"))  //false
	fmt.Println(handler.Matches(regularTree, "ba"))    //false
}
