package slower

import (
	"chat-slower/src/utils"
	"fmt"
	"time"
)

type MessageSlower struct {
	MessageChan            chan string
	SlowChan               chan string
	DisplayChan            chan string
	Speed                  float64
	Delay                  int
	VoidMessageCounter     int
	InitVoidMessageCounter int
}

func (m *MessageSlower) Funnel() {
	for {
		message := <-m.MessageChan

		select {
		case m.SlowChan <- message:
		default:
			fmt.Println(utils.Format("channel full, speed increase", "red"))
			m.Faster()
		}
	}
}

func (m *MessageSlower) Slow() {
	for {
		select {
		case msg := <-m.SlowChan:
			m.ResetMessageCounter()
			m.DisplayChan <- msg
		default:
			m.DecreaseMessageCounter()
			if m.VoidMessageCounter == 0 {
				fmt.Println(utils.Format(fmt.Sprintf("%v void messages, speed decrease", m.InitVoidMessageCounter), "red"))
				m.Slower()
			}
		}
		time.Sleep(time.Duration(m.Delay) * time.Millisecond)
	}
}

func (m *MessageSlower) Display() {
	for {
		fmt.Printf("delay:%vms speed:%.1fmsg/s %v\n", m.Delay, m.Speed, <-m.DisplayChan)
	}
}

func (m *MessageSlower) Slower() {
	m.Delay = m.Delay + int(m.Delay*15/100)
	m.Speed = 1 / float64(m.Delay) * 1000
}

func (m *MessageSlower) Faster() {
	m.Delay = m.Delay - int(m.Delay*25/100)
	m.Speed = 1 / float64(m.Delay) * 1000
}

func (m *MessageSlower) DecreaseMessageCounter() {
	m.VoidMessageCounter--
}

func (m *MessageSlower) ResetMessageCounter() {
	m.VoidMessageCounter = m.InitVoidMessageCounter
}
