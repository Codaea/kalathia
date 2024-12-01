package kalathia

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"codaea.com/kalathia/Bot/utils"
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

	if strings.Contains(content, "dog") {
		s.ChannelMessageSend(m.ChannelID, "Nice try. i'm a cat person")
		return
	}

	if strings.Contains(content, "isdown") {
		content = strings.Replace(content, "isdown", "", -1)
		content = strings.TrimSpace(content)

		statusCode, err := utils.Ping(content)
		checkNilErr(err)

		req, err := http.Get("https://http.cat/" + fmt.Sprint(statusCode) + ".jpg")
		checkNilErr(err)

		s.ChannelFileSend(m.ChannelID, "resp.jpg", req.Body)
	}

	switch m.Content {
		case "ping": {
			s.ChannelMessageSend(m.ChannelID, "Pong! :)")
		} 
	}
}

