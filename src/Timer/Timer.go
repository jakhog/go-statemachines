package Timer

import "time"

type Timer struct {
	Tick time.Duration
}

func (this *Timer) initialState() TimerStatefunc { return Tick }
