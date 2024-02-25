package main

import "fmt"

const (
	BoxerType = iota
	KickerType
)

type Fighter interface {
	Say()
	Interact(other *Fighter)
}

type Boxer struct{}
type Kicker struct{}
type Whimp struct{}

func (Boxer) Say() {
	fmt.Println("Omai wa mou shindero")
}

func (Boxer) Interact(other *Fighter) {
	switch (*other).(type) {
	case Boxer:
		fmt.Println("A true master of the arts, you are truely a man who can box")
	case Kicker:
		fmt.Println("Using your legs to fight, such a cowardly move, ought to be expected from you")
	}
}

func (Kicker) Say() {
	fmt.Println("Bro get out before I kick your balls")
}

func (Kicker) Interact(other *Fighter) {
	switch (*other).(type) {
	case Boxer:
		fmt.Println("You better not use them hands, I will explain with great detail as to why a spear is better than a sword")
	case Kicker:
		fmt.Println("Aight bro, imma let you off the hook since you aight but i don't want to see you in my block again")
	}
}

func CreateFighter(fighter_type int) Fighter {
	switch fighter_type {
	case BoxerType:
		return Boxer{}
	case KickerType:
		return Kicker{}
	default:
		return nil
	}
}
