package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/prometheus/common/log"

	"github.com/gabhendm/gostockbot/routes"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the stock status.
	s.UpdateStatus(0, "Fire Sale Time!!")
}

// TODO - Refactor Function into separate ones.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check if Message has stockPrice command
	if strings.HasPrefix(m.Content, "!stockPrice") {
		words := strings.Fields(m.Content)
		//get price
		quoteResponse, err := routes.GetStockPrice(words[1])
		if err != nil {
			log.Error(err.Error())
			return
		}
		if quoteResponse.Quote.Price == "" {
			s.ChannelMessageSend(m.ChannelID, "Stock not found!")
			return
		}
		// return price message
		price, err := strconv.ParseFloat(quoteResponse.Quote.Price, 2)
		if err != nil {
			log.Error(err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The price of Stock: %s is: R$%.2f BRL", words[1], price))
	}
	if strings.HasPrefix(m.Content, "!stockOverview") {
		words := strings.Fields(m.Content)
		overviewResponse, err := routes.GetStockOverview(words[1])

		if err != nil {
			log.Error(err.Error())
			return
		}
		if overviewResponse.Symbol == "" {
			s.ChannelMessageSend(m.ChannelID, "Stock not Found!")
			return
		}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Symbol: %s | Name: %s | Exchange: %s | Currency: %s | Country: %s | Sector: %s | Analyst Target Price: %s",
			overviewResponse.Symbol, overviewResponse.Name, overviewResponse.Exchange, overviewResponse.Currency, overviewResponse.Country, overviewResponse.Sector, overviewResponse.AnalystTargetPrice))

	}
	if strings.HasPrefix(m.Content, "!stockSearch") {
		words := strings.Fields(m.Content)
		overviewResponse, err := routes.SymbolSearch(words[1])

		if err != nil {
			log.Error(err.Error())
			return
		}
		if overviewResponse.BestMatches == nil {
			s.ChannelMessageSend(m.ChannelID, "Stock not Found!")
			return
		}
		for _, stock := range overviewResponse.BestMatches {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Symbol: %s | Name: %s | Type: %s | Region: %s | Currency: %s | Timezone: %s", stock.Symbol, stock.Name, stock.Type, stock.Region, stock.Currency, stock.Timezone))
		}

	}
}
