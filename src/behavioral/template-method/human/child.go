package human

import (
	"fmt"
)

type Child struct {
	Human
}

// step method
func (child *Child) mainAction() {
	fmt.Println("遊ぶ")
}

// step method
func (child *Child) getUp() {
	fmt.Println("遊ぶ")
}
