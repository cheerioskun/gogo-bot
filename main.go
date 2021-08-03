package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func unmarshalTimeSlots() {
	var timeslotPath = "data/timeslots.json"
	jsonFile, err := os.Open(timeslotPath)

	if err != nil {
		log.Fatalf("Error in opening timeslot json file, %v.", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &timeslotMap)

	// Testing, will remove before going into production
	for _, v := range timeslotMap {
		// Before Parsing for Testing => fmt.Println(v)
		v.parseTime()
		// After Parsing for Testing => fmt.Println(v)
	}
}

func unmarshalClasses() {
	sections := [2]string{"it-a", "it-b"}
	days := [7]string{"monday.json", "tuesday.json", "wednesday.json", "thursday.json", "friday.json"}
	basePath := "data/classes"
	slash := "/"
	for _, section := range sections {
		for _, day := range days {
			path := basePath + slash + section + slash + day
			jsonFile, err := os.Open(path)
			if err != nil {
				log.Fatalf("Error in opening timeslot json file, %v.", err)
			}
			var classesForOneDayForOneSection []*Class
			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &classesForOneDayForOneSection)
			for _, class := range classesForOneDayForOneSection {
				// fmt.Println(class)
				for _, timeslotNo := range class.TimeslotNos {
					class.Timeslots = append(class.Timeslots, timeslotMap[timeslotNo])
				}
				// fmt.Println(class, class.Timeslots)
			}
			classes = append(classes, classesForOneDayForOneSection...)
		}
	}
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
