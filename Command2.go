package main

import (
	"fmt"
	"time"
)

type Command2 interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type HelloMessage struct {
}

func (h HelloMessage) Info() string {
	return "Saltwater crocodile is strongest animal!"
}
func main() {
	var timeCommand Command2
	timeCommand = &TimePassed{time.Now()}
	var helloCommand Command2
	helloCommand = &HelloMessage{}
	time.Sleep(time.Second)
	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
