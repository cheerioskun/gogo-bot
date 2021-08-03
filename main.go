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

	mapRoles()
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
	// Populate channel name to channel id map for the guild
	err := processChannelMap(s)
	if err != nil {
		return
	}
	// Check every 5 minutes
	reminderChan := time.Tick(time.Minute * 5)
	go sendReminder(bot, reminderChan)
}

func addHandlers() {
	bot.AddHandler(ready)
}

func addIntents() {
	bot.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers
}

func makeRemindStringFromClass(class *Class) string {
	allRelevantRoles := ""
	for _, role := range class.Sections {
		allRelevantRoles += roleNameToRoleId[role] + "\n"
	}
	fmt.Println(allRelevantRoles)
	return fmt.Sprintf("%s%s for %s is starting soon!\nHere's the link to join: %s",
		allRelevantRoles,
		class.ClassType,
		class.SubjectName, class.MeetLink)
}

func mapRoles() {
	roleNameToRoleId = make(map[string]string)
	rolesPairs := [8](PairString){{"it-a", "ROLE_IT_A"}, {"it-b", "ROLE_IT_B"},
		{"it-1", "ROLE_IT_1"}, {"it-2", "ROLE_IT_2"},
		{"it-3", "ROLE_IT_3"}, {"it-4", "ROLE_IT_4"},
		{"it-5", "ROLE_IT_5"}, {"it-6", "ROLE_IT_6"}}
	for _, rolePair := range rolesPairs {
		fmt.Println(rolePair.roleEnvName)
		roleNameToRoleId[rolePair.roleName] = dotenv.GetString(rolePair.roleEnvName)
	}
}
