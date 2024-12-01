package kalathia

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

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
	content := strings.ToLower(m.Content)

	if strings.Contains(content, "cat") {
		resp, err := http.Get("https://cataas.com/cat")
		checkNilErr(err)
		_, err = s.ChannelFileSend(m.ChannelID, "kitteh.png", resp.Body)
		checkNilErr(err)
		s.ChannelMessageSend(m.ChannelID, "Meow! :3")
		return
	}

	switch m.Content {
		case "ping": {
			s.ChannelMessageSend(m.ChannelID, "Pong! :)")
		} 
	}
}

