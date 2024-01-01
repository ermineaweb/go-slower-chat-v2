package main

import (
	messageslower "chat-slower/src/MessageSlower"
	"flag"
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

const (
	INITIAL_MESSAGE_COUNTER = 8
	DELAY_MAX_ORIGINAL_CHAT = 1000
)

var streamer string

func init() {
	// todo accept an array of string
	flag.StringVar(&streamer, "streamer", "crocodyletv", "the twitch channel to view")
	flag.StringVar(&streamer, "s", "crocodyletv", "the twitch channel to view")
}

func main() {
	flag.Parse()

	ms := messageslower.MessageSlower{
		MessageChan:           make(chan string),
		DisplayChan:           make(chan string, 10),
		Speed:                 1 / float64(DELAY_MAX_ORIGINAL_CHAT) * 1000,
		Delay:                 DELAY_MAX_ORIGINAL_CHAT,
		MessageCounter:        INITIAL_MESSAGE_COUNTER,
		InitialMessageCounter: INITIAL_MESSAGE_COUNTER,
	}

	go ms.Slow()

	go func() {
		for {
			select {
			case msg := <-ms.DisplayChan:
				fmt.Printf("Delay:%vms Speed:%.1fmsg/s %v\n", ms.Delay, ms.Speed, msg)
				ms.ResetMessageCounter()
			default:
				ms.DecreaseMessageCounter()
				if ms.MessageCounter == 0 {
					fmt.Println(messageslower.Format(fmt.Sprintf("%v void messages", INITIAL_MESSAGE_COUNTER), "red"))
					ms.Slower()
				}
			}
			time.Sleep(time.Duration(ms.Delay) * time.Millisecond)
		}
	}()

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		streamer := messageslower.Format(fmt.Sprintf("%v", msg.Channel), "cyan")
		username := messageslower.Format(fmt.Sprintf("%v", msg.User.Name), "magenta")
		ms.MessageChan <- fmt.Sprintf("%v %v: %v", streamer, username, msg.Message)
	})

	client.Join(streamer)

	err := client.Connect()
	if err != nil {
		panic(err)
	}

}
