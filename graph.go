package main

import "fmt"

type Node struct {
	val   int
	other []*Node
}

func main() {
	node1 := Node{0, nil}
	node2 := Node{1, []*Node{&node1}}
	node1.other = []*Node{&node2}
	cur := &node1
	for len(cur.other) > 0 {
		fmt.Println(cur.val)
		cur = cur.other[0]
	}
}
