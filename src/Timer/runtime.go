package Timer

import "fmt"

type TimerStatefunc func(*Timer) TimerStatefunc

func run(this *Timer) {
	fmt.Println("Starting Timer")
	for state := this.initialState(); state != nil; {
		state = state(this)
	}
	fmt.Println("Ending Timer")
}

func (this *Timer) Run() {
	go run(this)
}
