package command

import (
	"fmt"

	"github.com/ChaseBrand/basebot/internal/handler"
	"github.com/bwmarrin/discordgo"
)

func PingCommand() *handler.Command {
	return &handler.Command{
		Name:        "ping",
		Description: "Pong!",
		Options:     []*discordgo.ApplicationCommandOption{},
		Executor:    handlePing,
	}
}

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong",
		},
	})

	if err != nil {
		fmt.Printf("error responding to ping command: %v\n", err)
	}
}
