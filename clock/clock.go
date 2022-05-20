package clock

import "time"

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	date := time.Date(1970, 0, 0, h, m, 0, 0, time.UTC)
	return Clock{
		hour:   date.Hour(),
		minute: date.Minute(),
	}
}

func (c Clock) Add(m int) Clock {
	date := c.time()
	date = date.Add(time.Minute * time.Duration(m))
	return Clock{minute: date.Minute(), hour: date.Hour()}
}

func (c Clock) Subtract(m int) Clock {
	date := c.time()
	date = date.Add(-time.Minute * time.Duration(m))
	return Clock{minute: date.Minute(), hour: date.Hour()}
}

func (c Clock) String() string {
	date := c.time()
	return date.Format("15:04")
}

func (c Clock) time() time.Time {
	return time.Date(1970, 0, 0, c.hour, c.minute, 0, 0, time.UTC)
}
