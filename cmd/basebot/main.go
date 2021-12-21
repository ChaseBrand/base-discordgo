package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ChaseBrand/basebot/internal/command"
	"github.com/ChaseBrand/basebot/internal/handler"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	serverID := os.Getenv("SERVER_ID")
	lavalinkPassword := os.Getenv("LAVA_PASS")

	if token == "" || serverID == "" || lavalinkPassword == "" {
		panic(".env file not filled out correctly. Shutting down.")
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic("error starting Discord session.")
	}

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers

	if err = session.Open(); err != nil {
		panic("Error while opening session")
	}

	commandHandler := handler.NewHandler(session, serverID)

	// Register Events here.

	// Register Commands here.
	commandHandler.Register(command.PingCommand())

	fmt.Println("Bot is running successfully. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	commands := commandHandler.GetCommands()

	for _, command := range commands {
		err := commandHandler.Remove(command)
		if err != nil {
			panic("error removing command")
		}
	}

	session.Close()
}
