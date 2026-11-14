package clock
import "fmt"
// Define the Clock type here.

type Clock struct {
	hora   int
	minuto int
}

const minuteInAnHour, minutesInADay = 60 , 1440

func New(h, m int) Clock {
	minutosTotales := (h * minuteInAnHour + m) % minutesInADay
	if minutosTotales < 0 {
		minutosTotales += minutesInADay
	}

	return Clock{
		hora: minutosTotales / minuteInAnHour,
		minuto: minutosTotales % minuteInAnHour,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hora , c.minuto + m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d",c.hora , c.minuto)
}