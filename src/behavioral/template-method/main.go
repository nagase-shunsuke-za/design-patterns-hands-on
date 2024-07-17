package main

import (
	"design-patterns-hands-on/behavioral/template-method/human"
)

func main() {
	child := human.Human{
		IHuman: &human.Child{},
	}

	adult := human.Human{
		IHuman: &human.Adult{},
	}

	child.StartDay("child")
	adult.StartDay("adult")
}
