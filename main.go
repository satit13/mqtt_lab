package main

import (
	"fmt"
	//"time"

	emitter "github.com/emitter-io/go"
	"time"
	//"github.com/gin-gonic/gin/json"
)

func main() {
	// Create the options with default values
	o := emitter.NewClientOptions()

	// Set the message handler
	o.SetOnMessageHandler(func(client emitter.Emitter, msg emitter.Message) {
		fmt.Printf(" \n Received message: %s\n", msg.Payload())
		fmt.Printf("topic is : %s  \n", msg.Topic())

	})

	// Set the presence notification handler
	o.SetOnPresenceHandler(func(_ emitter.Emitter, p emitter.PresenceEvent) {
		//fmt.Printf("Occupancy: %v\n", p.Occupancy)
	})

	// Create a new emitter client and connect to the broker
	c := emitter.NewClient(o)
	sToken := c.Connect()
	if sToken.Wait() && sToken.Error() != nil {
		panic("Error on Client.Connect(): " + sToken.Error().Error())
	}

	mqKey := "ypwbEyi92rkl5z6dZTN1_KAj8GXA8ax1"
	mSubChanel := "satit13/+/"
	mPubChanel := "satit13/MOBILE/B001/"
	// Subscribe to the presence demo channel
	c.Subscribe(mqKey, mSubChanel)

	// Publish to the channel
	//c.Publish(mqKey, mChanel, "hello")
	//c.Publish(mqKey, mChanel, "hello")
	//c.Publish(mqKey, mChanel, "hello")
	//c.Publish(mqKey, mChanel, "hello")


	// Ask for presence
	//r := emitter.NewPresenceRequest()
	//r.Key = mqKey
	//r.Channel = "satit13/#/"
	//c.Presence(r)

	i := 60
	for i >= 0 {
		fmt.Printf("loop %v \n ", i)
		time.Sleep(1 * time.Second)
		c.Publish(mqKey, mPubChanel, `{"confirm":true}`)
		i--
	}
	fmt.Println("\n Finished")
	// stop after 10 seconds
}