package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/profclems/go-dotenv"
)

var bot *discordgo.Session
var TOKEN string

// Initializes the bot on package initialization
func init() {
	err := dotenv.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	TOKEN = dotenv.GetString("AUTHTOKEN")
	bot, err = discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
	bot.Identify.Intents = discordgo.IntentsGuildMessages
}

func main() {
	// Open a websocket connection to Discord and begin listening.
	err := bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()

}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Println("Bot is up!")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
