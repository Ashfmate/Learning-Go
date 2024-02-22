package main

import (
	"fmt"
	"strings"
)

type dude struct {
	num int
	msg string
}

func getMsg(num int, choices ...dude) string {
	res := []string{}
	for _, val := range choices {
		if num%val.num == 0 {
			res = append(res, val.msg)
		}
	}
	return strings.Join(res, "")
}

func main() {
	for num := 1; num <= 105; num++ {
		ans := getMsg(num, dude{3, "Fizz"}, dude{5, "Buzz"}, dude{7, "Bazz"})
		if len(ans) > 0 {
			fmt.Println(ans)
		} else {
			fmt.Println(num)
		}
	}
}
