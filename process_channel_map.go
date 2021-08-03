package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var channelFromName = make(map[string]*discordgo.Channel)

func processChannelMap(s *discordgo.Session) error {
	channels, err := s.GuildChannels(IT_SERVER_GUILDID)
	if err != nil {
		log.Printf("could not get guild channels: %q", err)
		return err
	}
	for _, channel := range channels {
		// Only add channels and not channel dividers like "Text Channels"
		if strings.ToLower(channel.Name) == channel.Name {
			channelFromName[channel.Name] = channel
		}
	}
	return nil
}
