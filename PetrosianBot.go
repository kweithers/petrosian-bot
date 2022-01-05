package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

func main() {
	twitch_channels := []string{"gothamchess"}
	pastas := []string{`Are you kidding ??? What the **** are you talking about man ?`,
		`You are a biggest looser i ever seen in my life !`,
		`You was doing PIPI in your pampers when i was beating players much more stronger then you!`,
		`You are not proffesional, because proffesionals knew how to lose and congratulate opponents, you are like a girl crying after i beat you!`,
		`Be brave, be honest to yourself and stop this trush talkings!!!`,
		`Everybody know that i am very good blitz player, i can win anyone in the world in single game!`,
		`And "w"esley "s"o is nobody for me, just a player who are crying every single time when loosing, ( remember what you say about Firouzja ) !!!`,
		`Stop playing with my name, i deserve to have a good name during whole my chess carrier,`,
		`I am Officially inviting you to OTB blitz match with the Prize fund!`,
		`Both of us will invest 5000$ and winner takes it all!`}

	rand.Seed(time.Now().Unix())
	client := twitch.NewClient("petrosianbot", "oauth:abc123")
	//Flag used to prevent the bot from spamming
	ready_to_post := true

	//Whenever a twitch chat message is posted
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		//Check if it contains one of the keywords AND the bot hasn't posted too recently
		if checkForKeywords(strings.ToLower(message.Message)) && ready_to_post {
			//See the message that contained a keyword
			fmt.Println(message.Channel, message.User.Name, message.Message)
			//Craft the response
			reply := "@" + message.User.Name + " " + pastas[rand.Intn(10)]
			//Say the response
			client.Say(message.Channel, reply)
			//Start a countdown until the bot can post again
			ready_to_post = false
			go countdown(&ready_to_post, "1m")
		}
	})
	//Join all of the twitch chats specified above
	for _, ttv := range twitch_channels {
		client.Join(ttv)
	}
	//Connect to the IRC server
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

//Make the bot wait before it posts again to prevent spamming
func countdown(ready_to_post *bool, str string) {
	m, _ := time.ParseDuration(str)
	time.Sleep(m)
	*ready_to_post = true
}

//Check if the twitch chat message contains one of the keywords
func checkForKeywords(str string) bool {
	return strings.Contains(str, "pipi") || strings.Contains(str, "pampers") || strings.Contains(str, "tigran") || strings.Contains(str, "petrosian")
}
