package main

import "fmt"

type Node struct {
	val   int
	other []*Node
}

func traverse(start *Node) {
	for _, val := range start.other {
		traverse(val)
	}
	fmt.Println(start.val)
}

func main() {
	node1 := Node{0, nil}
	node2 := Node{1, []*Node{}}
	node3 := Node{2, []*Node{&node2}}
	node4 := Node{3, []*Node{}}
	node1.other = []*Node{&node2, &node3}
	node2.other = []*Node{&node4}
	traverse(&node1)
}
