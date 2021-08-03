package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var channelFromName = make(map[string]*discordgo.Channel)
var itGuild *discordgo.Guild

func processChannelMap(s *discordgo.Session) error {
	var err error
	itGuild, err = s.Guild(IT_SERVER_GUILDID)
	if err != nil {
		log.Printf("could not get guild from id: %q", err)
		return err
	}
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
