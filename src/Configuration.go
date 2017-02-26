package main

import "Timer"

func main() {
	var timerInst Timer.Timer

	timerInst.Tick = 1000000000

	timerInst.Run()

	for {}
}
