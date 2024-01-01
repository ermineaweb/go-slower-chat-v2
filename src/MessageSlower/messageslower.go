package messageslower

import (
	"fmt"
)

type MessageSlower struct {
	MessageChan           chan string
	DisplayChan           chan string
	Speed                 float64
	Delay                 int
	MessageCounter        int
	InitialMessageCounter int
}

func (m *MessageSlower) Slow() {
	for {
		message := <-m.MessageChan

		select {
		case m.DisplayChan <- message:
		default:
			fmt.Println(Format("Channel full", "red"))
			m.Faster()
		}
	}
}

func (m *MessageSlower) Slower() {
	m.Delay = m.Delay + int(m.Delay*10/100)
	m.Speed = 1 / float64(m.Delay) * 1000
	fmt.Println(Format("Decrease speed", "red"))
}

func (m *MessageSlower) Faster() {
	m.Delay = m.Delay - int(m.Delay*20/100)
	m.Speed = 1 / float64(m.Delay) * 1000
	fmt.Println(Format("Increase speed", "red"))
}

func (m *MessageSlower) DecreaseMessageCounter() { m.MessageCounter-- }

func (m *MessageSlower) ResetMessageCounter() { m.MessageCounter = m.InitialMessageCounter }

func Format(text, color string) string {
	var colors = map[string]string{
		"black":   "\x1b[30m",
		"red":     "\x1b[31m",
		"green":   "\x1b[32m",
		"yellow":  "\x1b[33m",
		"blue":    "\x1b[34m",
		"magenta": "\x1b[35m",
		"cyan":    "\x1b[36m",
		"white":   "\x1b[37m",
		"bold":    "\x1b[1m",
		"reset":   "\x1b[0m\x1b[39m",
	}
	return fmt.Sprintf("%v%v%v", colors[color], text, colors["reset"])
}
