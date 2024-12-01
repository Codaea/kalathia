package kalathia

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// event handler
	discord.AddHandler(newMessage)

	// open session
	discord.Open()

	// keep bot running until os interupt
	fmt.Println("Bot is running. Press CTRL-C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	discord.Close()
}

func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// prevent bot from responding to itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "PONG!")
	}
	switch m.Content {
		case "ping": {
			s.ChannelMessageSend(m.ChannelID, "PONG!")
		}
	}
}

