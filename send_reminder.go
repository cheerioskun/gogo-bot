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
		linkChannelID = channelFromName[CHANNEL_NAME_BOTCMDS].ID
	}
	for range ch {
		log.Printf("Checking classes at %s", time.Now().String())
		for _, class := range classes {
			if class.Weekday != time.Now().Weekday().String() {
				continue
			}
			toRemind := false
			for idx, ts := range class.Timeslots {
				// Only remind on the start of first timeslot
				if ts.isComingUp() && idx == 0 {
					toRemind = true
					break
				}
			}
			if toRemind {
				log.Printf("Time to remind!")
				contentstring := makeRemindString(class)
				s.ChannelMessageSend(linkChannelID, contentstring)
			}
		}
	}
}
