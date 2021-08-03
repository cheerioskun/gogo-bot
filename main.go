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

var bot *discordgo.Session
var TOKEN string
var timeslotMap map[string]*Timeslot
var classes []*Class

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
	err := processChannelMap(s)
	if err != nil {
		return
	}
	reminderChan := time.Tick(time.Second * 10)
	go sendReminder(bot, channelFromName[CHANNEL_NAME_BOTCMDS].ID, reminderChan)
}

func addHandlers() {
	bot.AddHandler(ready)
	bot.AddHandler(messageCreate)
}

func addIntents() {
	bot.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers
}

func sendReminder(s *discordgo.Session, channelID string, ch <-chan time.Time) {
	for range ch {
		s.ChannelMessageSend(channelID, "Namaste")
	}
}
