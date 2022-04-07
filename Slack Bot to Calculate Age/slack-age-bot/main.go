package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main() {
	//SET AN ENVIRONMENT
	os.Setenv("SLACK-BOT-TOKEN", "xoxb-3362450863105-3362493827169-wMdYAdjjctkKGaaT9yYlEvOX")
	os.Setenv("SLACK-APP-TOEKN", "xapp-1-A03AC5QHWA0-3349851097411-7169b4333b7f181fbcb0e85be73ed5dd43ebf82ca59318695708f53c34320840")

	//NOW WE WILL CREATE A BOT
	bot := slacker.NewClient(os.Getenv("SLACK-BOT-TOKEN"), os.Getenv("SLACK-APP-TOKEN")) //Setting up an new client with credentials

	//Go routine for printing command events(Basically prints out given commands)
	go printCommandEvents(bot.CommandEvents()) // commandEvents is func which will handle commands given to bot

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob Calculator",
		Example:     "my yob is 2022",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Fatal(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
