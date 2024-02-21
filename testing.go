package main

import (
	"fmt"
)

struct dude {
  int num
  string msg
}

func main() {
  choices := []dude{dude(3,"Fizz"), dude(5,"Buzz")}
	for num := range 101 {
    found := false
	  for _, val := range choices {
      if num % val.num == 0 {
        found = true
        fmt.Print(val.msg)
      }
    }
    if found {
      fmt.Println()
    } else {
      fmt.Println(num)
    }
  }
}
