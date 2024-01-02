package main

import (
	"chat-slower/src/slower"
	"chat-slower/src/utils"
	"flag"
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
)

var (
	args                   string
	channels               []string
	initSpeedMsgSec        float64
	initVoidMessageCounter int
	initDelayMs            int
)

func init() {
	flag.StringVar(&args, "c", "shroud,ninja,pokimane,xqc", "list of channels to view")
	flag.Float64Var(&initSpeedMsgSec, "s", 1, "initial chats speed in msg/s")
	flag.IntVar(&initVoidMessageCounter, "m", 8, "number of void messages before decrease the chat speed")
	flag.Parse()

	channels = strings.Split(args, ",")
	// speed = msg / delay <---> delay = 1 / speed
	initDelayMs = int(1 / initSpeedMsgSec * 1000)
}

func main() {
	ms := slower.MessageSlower{
		MessageChan:            make(chan string),
		SlowChan:               make(chan string, 25),
		DisplayChan:            make(chan string),
		Speed:                  initSpeedMsgSec,
		Delay:                  initDelayMs,
		VoidMessageCounter:     initVoidMessageCounter,
		InitVoidMessageCounter: initVoidMessageCounter,
	}

	go ms.Funnel()
	go ms.Slow()
	go ms.Display()

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		streamer := utils.Format(fmt.Sprintf("%v", msg.Channel), "cyan")
		username := utils.Format(fmt.Sprintf("%v", msg.User.Name), "magenta")
		ms.MessageChan <- fmt.Sprintf("%-35v %40v: %v", streamer, username, msg.Message)
	})

	for _, channel := range channels {
		client.Join(channel)
	}

	err := client.Connect()
	if err != nil {
		panic(err)
	}

}
