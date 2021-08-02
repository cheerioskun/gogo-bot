package main

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
	// Check if someone is asking for the link
	if m.Content == "!link" {

		guildRoles, err := s.GuildRoles(m.GuildID)
		if err != nil {
			log.Printf("could not get list of server roles: %v", err)
			return
		}
		roleIDs := m.Member.Roles
		roles := make(map[string]bool)
		for _, id := range roleIDs {
			idint, err := strconv.Atoi(id)
			if err != nil {
				log.Printf("could not convert string to integer: %v", err)
				return
			}
			roles[guildRoles[idint].Name] = true
		}
		// Now we can check against the roles map for existence
	}
}
