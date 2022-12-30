package clock

import "fmt"
// Define the Clock type here.
type Clock struct{
    hours, minutes int
}

func fixHours(h int) int {
    if h < 0{
        h = h %24 + 24
    }
    return h % 24
}

func fixMinutes(m int) int{
    if m< 0{
        m = m %60 + 60
    }
	return m %60
}

func New(h, m int) Clock {
	c := Clock{}
    hours := fixHours(h)
	c.hours = hours

    if m < 0{
        return c.Subtract(-1 * m)
    }
	return c.Add(m)
}

func (c Clock) Add(m int) Clock {
	hoursToAdd := (m + c.minutes) / 60
    newMinutes := fixMinutes(m + c.minutes)

    newHours := fixHours(c.hours + hoursToAdd)
    
    return Clock{hours:newHours, minutes:newMinutes}
}

func (c Clock) Subtract(m int) Clock {
	var hoursToSub int
    minuteDiff := (c.minutes - m)
    if minuteDiff < 0{
        hoursToSub = -1 * (minuteDiff / 60) 
        if minuteDiff %60 != 0{
            hoursToSub +=1
        }
    }
	
    newMinutes := fixMinutes(c.minutes - m)
    newHours := fixHours(c.hours - hoursToSub)

    return Clock{hours: newHours, minutes:newMinutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
