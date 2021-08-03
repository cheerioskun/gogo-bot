package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func sendReminder(s *discordgo.Session, channelID string, ch <-chan time.Time) {
	for range ch {
		s.ChannelMessageSend(channelID, "Namaste")
	}
}
