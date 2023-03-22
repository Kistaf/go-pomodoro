package main

import (
	"fmt"
)

type Countdown struct {
	left  int
	total int
    mode string
}

func (c *Countdown) prettyStatus() string {
	hrs := c.left / 3600
	mins := (c.left % 3600) / 60
	secs := (c.left % 3600) % 60
	if hrs == 0 && mins != 0{
		return fmt.Sprintf("%dm %ds", mins, secs)
	}
	if mins == 0 && hrs == 0 {
		return fmt.Sprintf("%ds", secs)
	}
	return fmt.Sprintf("%dh %dm %ds", hrs, mins, secs)
}

func (c *Countdown) next() float64 {
	c.left--
	return 1.0 / float64(c.total)
}

