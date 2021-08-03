package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/profclems/go-dotenv"
)

func sendReminder(s *discordgo.Session, ch <-chan time.Time) {
	linkChannelID := channelFromName[CHANNEL_NAME_LINKS].ID
	flag := dotenv.GetBool("STAGE")
	if flag {
		linkChannelID = dotenv.GetString("STAGE_TEST_CHANNEL")
	}
	for range ch {
		log.Printf("Checking classes at %s", time.Now().String())
		for _, class := range classes {
			if class.Weekday != time.Now().Weekday().String() {
				continue
			}
			toRemind := false
			for _, ts := range class.Timeslots {
				if ts.isComingUp() {
					toRemind = true
					break
				}
			}
			if toRemind {
				log.Printf("Time to remind!")
				contentstring := makeRemindStringFromClass(class)
				s.ChannelMessageSend(linkChannelID, contentstring)
			}
		}
	}
}
