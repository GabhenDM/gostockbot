package service

import (
	"fmt"
	"strings"

	"github.com/gabhendm/gostockbot/models"
	"github.com/prometheus/common/log"

	"github.com/gabhendm/gostockbot/routes"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func Ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the stock status.
	s.UpdateStatus(0, "!stock")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// check if the message is "!airhorn"
	if strings.HasPrefix(m.Content, "!stockPrice") {

		// Find the channel that the message came from.
		//c, err := s.State.Channel(m.ChannelID)
		//f err != nil {
		// Could not find channel.
		//	return
		//}

		// If the message is "ping" reply with "Pong!"
		words := strings.Fields(m.Content)
		quoteResponse := models.AVQuoteResponse{}
		quoteResponse, err := routes.GetStockPrice(words[1])
		if err != nil {
			log.Error(err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The price of Stock: %s is: %s", words[1], quoteResponse.Quote.Price))
	}
}
