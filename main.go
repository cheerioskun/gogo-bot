package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/profclems/go-dotenv"
)

var (
	bot                  *discordgo.Session
	TOKEN                string
	prefixChar           string
	scheduleCommand      string
	schCommand           string
	scheduleCommandUsage string
)

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
	unmarshalTimeSlots()
	unmarshalClasses()
	prepareCommandsAndUsages()
	addHandlers()
	addIntents()
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
	// Populate channel name to channel id map for the guild
	err := processChannelMap(s)
	if err != nil {
		return
	}
	prefixChar = dotenv.GetString("PREFIX")
	err = mapRoles(s)
	if err != nil {
		return
	}
	// Check every 5 minutes
	reminderChan := time.Tick(time.Minute * 5)
	go sendReminder(bot, reminderChan)
}

func addHandlers() {
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
}

func addIntents() {
	bot.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers
}
