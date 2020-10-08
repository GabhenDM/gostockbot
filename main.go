package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/bwmarrin/discordgo"
	"github.com/gabhendm/gostockbot/service"

	"github.com/gabhendm/gostockbot/config"
)

func main() {
	config.LoadConfig()

	dg, err := discordgo.New("Bot " + viper.GetString("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(service.Ready)

	dg.AddHandler(service.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("GoStockBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}
