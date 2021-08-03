package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func sendReminder(s *discordgo.Session, ch <-chan time.Time) {
	//linkChannelID := channelFromName[CHANNEL_NAME_LINKS]
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
				//contentstring = makeRemindStringFromClass(class)
				//s.ChannelMessageSend(linkChannelID, contentstring)
			}
		}
	}
}
