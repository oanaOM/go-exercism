//Package clock provides functionalities to calculate time
package clock

import (
	"fmt"
)

//Clock type defines the minutes
type Clock struct {
	minutes int
}

//String formats the time as string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)

}

//New implementes a constructor for our Clock type
func New(h, m int) Clock {
	const minutesInDay = 1440

	minutes := (h*60 + m) % minutesInDay

	if minutes < 0 {
		minutes += minutesInDay
	}

	return Clock{minutes}
}

//Add method that adds m minutes to a time
func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

//Subtract method that substracts m minutes from time
func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}
