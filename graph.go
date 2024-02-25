package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Node
type Node struct {
	id       int
	other    map[rune]*Node
	is_final bool
}

func NewNode(id int) *Node {
	return &Node{id, map[rune]*Node{}, false}
}

func PrintNode(self *Node) {
	buffer := fmt.Sprint("Node (", self.id, ") Points to: [")
	others := []string{}
	for _, val := range self.other {
		others = append(others, fmt.Sprint(" ", val.id, " "))
	}
	buffer += strings.Join(others, " , ") + "] "
	if self.is_final {
		buffer += "And it is a final Node"
	} else {
		buffer += "And it is not a final Node"
	}
	fmt.Println(buffer)
}

// Node

// Machine
type Machine struct {
	nodes []*Node
}

func NewMachine() *Machine {
	return &Machine{nodes: []*Node{}}
}

func AddNode(self *Machine) *Machine {
	self.nodes = append(self.nodes, NewNode(len(self.nodes)))
	return self
}

func NodeTransition(self *Machine, from int, to int, label rune) *Machine {
	if from > len(self.nodes) {
		fmt.Fprintf(os.Stderr, "The id %d does not exit", from)
		os.Exit(-1)
	}
	if to > len(self.nodes) {
		fmt.Fprintf(os.Stderr, "The id %d does not exit", to)
		os.Exit(-1)
	}
	self.nodes[from].other[label] = self.nodes[to]
	return self
}

func MakeStart(self *Machine, start int) *Machine {
	if start > len(self.nodes) {
		fmt.Fprintf(os.Stderr, "The id %d does not exit", start)
		os.Exit(-1)
	}
	temp := self.nodes[0]
	self.nodes[0] = self.nodes[start]
	self.nodes[start] = temp
	return self
}

func MakeEnd(self *Machine, end int) *Machine {
	if end > len(self.nodes) {
		fmt.Fprintf(os.Stderr, "The id %d does not exit", end)
		os.Exit(-1)
	}
	self.nodes[end].is_final = true
	return self
}

func TestInput(self *Machine, input string) bool {
	cur_node := self.nodes[0]
	for len(input) > 0 {
		sym, width := utf8.DecodeRuneInString(input)
		cur_node = cur_node.other[sym]
		if cur_node == nil {
			return false
		}
		input = input[width:]
	}
	return cur_node.is_final
}

// Machine
