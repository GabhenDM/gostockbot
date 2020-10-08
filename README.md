# GoStockBot

A Go Discord Bot for checking stock prices and creating price-based alerts


Utilizes discordgo for WebSocket interfacing with Discord and AlphaVantage API for Stock Quote references.


### Example .env file

``` shell
export AV_API_KEY= << INSERT AlphaVantage API KEY >>
export API_URL=https://www.alphavantage.co/query
export DISCORD_TOKEN=<< INSERT DISCORD TOKEN >>
```

### TODO 

- [X] Implement Configuration Params
- [X] Implement Basic REST API
- [X] Implement Requests to Third-Party Stock Price API
- [X] Implement Discord Integration
- [ ] Documenting/Logging
- [ ] Caching (?)
- [ ] Dockerize everything!!!


