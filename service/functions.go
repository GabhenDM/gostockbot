package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gabhendm/gostockbot/models"
	"github.com/prometheus/common/log"

	"github.com/gabhendm/gostockbot/routes"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the stock status.
	s.UpdateStatus(0, "Fire Sale Time!!")
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check if Message has stockPrice command
	if strings.HasPrefix(m.Content, "!stockPrice") {
		words := strings.Fields(m.Content)
		quoteResponse := models.AVQuoteResponse{}
		//get price
		quoteResponse, err := routes.GetStockPrice(words[1])
		if err != nil {
			log.Error(err.Error())
			return
		}
		// return price message
		price, err := strconv.ParseFloat(quoteResponse.Quote.Price, 2)
		if err != nil {
			log.Error(err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The price of Stock: %s is: $%.2f BRL", words[1], price))
	}
}
