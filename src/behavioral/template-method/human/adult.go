package human

import (
	"fmt"
)

type Adult struct {
	Human
}

// step method
func (adult *Adult) mainAction() {
	fmt.Println("仕事")
}

// step method
func (adult *Adult) eveningAction() {
	fmt.Println("仕事")
}
