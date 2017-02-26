package Timer

import "fmt"
import "time"

func Tick(this *Timer) TimerStatefunc {
	time.Sleep(this.Tick)
	fmt.Println("Tick!")

	return Tack
}

func Tack(this *Timer) TimerStatefunc {
	time.Sleep(this.Tick)
	fmt.Println("Tack!")

	return nil
}
