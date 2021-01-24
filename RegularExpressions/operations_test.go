package RegularExpressions

import (
	"fmt"
	"testing"
)

func TestLiteral(t *testing.T) {
	//pattern := Repeat{
	//	Choose{
	//		Concatenate {
	//				Left:  Literal{Character: 'a'},
	//				Right: Literal{Character: 'b'},
	//		},
	//		Literal{Character: 'a'},
	//	}}
	//
	//fmt.Println(pattern)
	handler := Handler{}
	fmt.Println(handler.Matches(Empty{}, ""))
	fmt.Println(handler.Matches(Empty{}, "a"))

	fmt.Println(handler.Matches(Literal{'a'}, ""))
	fmt.Println(handler.Matches(Literal{'a'}, "a"))
	fmt.Println(handler.Matches(Literal{'a'}, "b"))
}

func TestConcatenate(t *testing.T) {
	handler := Handler{}
	fmt.Println(handler.Matches(Concatenate{
		Left:  Literal{Character: 'a'},
		Right: Literal{Character: 'b'},
	}, "a"))
	fmt.Println(handler.Matches(Concatenate{
		Left:  Literal{Character: 'a'},
		Right: Literal{Character: 'b'},
	}, "ab"))
	fmt.Println(handler.Matches(Concatenate{
		Left:  Literal{Character: 'a'},
		Right: Literal{Character: 'b'},
	}, "abc"))

	fmt.Println(handler.Matches(Concatenate{
		Left: Concatenate{
			Left:  Literal{Character: 'a'},
			Right: Literal{Character: 'b'},
		},
		Right: Literal{Character: 'c'},
	}, "abc"))
	fmt.Println(handler.Matches(Concatenate{
		Left: Concatenate{
			Left:  Literal{Character: 'a'},
			Right: Literal{Character: 'b'},
		},
		Right: Literal{Character: 'c'},
	}, "ab"))
}

func TestChoose(t *testing.T) {
	handler := Handler{}
	fmt.Println(handler.Matches(Choose{
		Left:  Literal{Character: 'a'},
		Right: Literal{Character: 'b'},
	}, "a"))
	fmt.Println(handler.Matches(Choose{
		Left:  Literal{Character: 'a'},
		Right: Literal{Character: 'b'},
	}, "b"))
	fmt.Println(handler.Matches(Choose{
		Left:  Literal{Character: 'a'},
		Right: Literal{Character: 'b'},
	}, "c"))
	fmt.Println(handler.Matches(Choose{
		Left:  Literal{Character: 'a'},
		Right: Literal{Character: 'b'},
	}, "ab"))
}

func TestRepeat(t *testing.T) {
	handler := Handler{}
	fmt.Println(handler.Matches(Repeat{Literal{Character: 'a'}}, "a"))
	fmt.Println(handler.Matches(Repeat{Literal{Character: 'a'}}, "b"))
	fmt.Println(handler.Matches(Repeat{Literal{Character: 'a'}}, "aaa"))
	fmt.Println(handler.Matches(Repeat{Literal{Character: 'a'}}, "aa"))
}
