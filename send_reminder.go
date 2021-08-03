package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func sendReminder(s *discordgo.Session, ch <-chan time.Time) {
	for range ch {
		// Add logic to check time and do
	}
}
