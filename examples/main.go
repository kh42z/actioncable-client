package main

import (
	"github/kh42z/actioncable"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/cable", nil)
	if err != nil {
		log.Fatalln("unable to connect ", err)
	}
	defer ws.Close()
	c := actioncable.NewClient(ws, actioncable.WithLogger(&log.Logger{}))
	c.AddChannelHandler("UserChannel", &UserChannel{})
	go func() {
		c.Subscribe("UserChannel", 1)
	}()
	if err := c.Run(); err != nil {
		log.Fatal("Actioncable: ", err)
	}
}

func (u *UserChannel) SubscriptionHandler(c *actioncable.Client, id int) {
	log.Println("UserChannel, successfully subscribed to id: ", id)
	c.SendMessage("UserChannel", id, "Hello!")
}

func (u *UserChannel) MessageHandler(_ *actioncable.Client, content []byte, id int) {
	log.Printf("UserChannel_%d: <%s>", id, content)
}

type UserChannel struct{}
