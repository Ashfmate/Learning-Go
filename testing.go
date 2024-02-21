package main

import (
	"fmt"
)

type dude struct {
	num int
	msg string
}

func getMsg(num int) (string, bool) {
	found := false
	choices := []dude{{3, "Fizz"}, {5, "Buzz"}, {7, "bazz"}}
	buffer := ""
	for _, val := range choices {
		if num%val.num == 0 {
			found = true
			buffer = fmt.Sprint(buffer, val.msg)
		}
	}
	return buffer, found
}

func main() {
	for num := 1; num <= 105; num++ {
		ans, found := getMsg(num)
		if found {
			fmt.Println(ans)
		} else {
			fmt.Println(num)
		}
	}
}
