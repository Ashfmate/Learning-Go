package main

import (
	"fmt"
	"unicode/utf8"
)

type Node struct {
	val   int
	other map[rune]*Node
}

func traverse(start *Node, input string) bool {
	if start == nil {
		return false
	}
	if start.val == 3 {
		return len(input) == 0
	}
	if len(input) == 0 {
		return false
	}
	symbol, width := utf8.DecodeRuneInString(input)
	if len(start.other) == 0 {
		return false
	}
	next := start.other[symbol]
	return traverse(next, input[width:])
}

func main() {
	node3 := Node{3, map[rune]*Node{}}
	node1 := Node{1, map[rune]*Node{'b': &node3}}
	node2 := Node{2, map[rune]*Node{'a': &node1}}
	node0 := Node{0, map[rune]*Node{'a': &node1, 'b': &node2}}
	fmt.Println(traverse(&node0, "bb"))
}
