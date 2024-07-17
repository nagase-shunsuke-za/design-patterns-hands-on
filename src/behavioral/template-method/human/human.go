package human

import "fmt"

type IHuman interface {
	StartDay(string) //一連の流れ
	getUp()          //朝の行動
	mainAction()     //日中の行動
	eveningAction()  //夕方の行動
	sleep()          //夜の行動
}

type Human struct {
	IHuman IHuman
}

// static step method
func (human *Human) getUp() {
	fmt.Println("起床")
}
func (human *Human) mainAction() {
	fmt.Println("mainAction")
}

// static step method
func (human *Human) sleep() {
	fmt.Println("就寝")
}

// default step method
func (human *Human) eveningAction() {
	fmt.Println("自由に過ごす")
}

// template method
func (human *Human) StartDay(h string) {
	fmt.Println(h)
	human.getUp()
	human.IHuman.mainAction()
	human.IHuman.eveningAction()
	human.sleep()
	fmt.Println()
}
