package main

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, scheduleCommand) || strings.HasPrefix(m.Content, schCommand) {
		res := strings.Split(m.Content, " ")
		var day string
		if len(res) > 2 {
			s.ChannelMessageSend(m.ChannelID, scheduleCommandUsage)
			return
		} else if len(res) == 2 {
			day = res[1]
		} else {
			day = time.Now().Weekday().String()
		}
		s.ChannelMessageSend(m.ChannelID, makeScheduleClassString(m.Member.Roles, day))
	}
}
