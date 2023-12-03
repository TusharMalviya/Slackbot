package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/shomali11/slacker"
)

var jokes = []string{
	"Why don't scientists trust atoms? Because they make up everything.",
	"What do you call fake spaghetti? An impasta.",
	"I told my wife she was drawing her eyebrows too high. She looked surprised.",
	"How do you organize a space party? You planet.",
	"I only know 25 letters of the alphabet. I don't know y.",
	"Why don't skeletons fight each other? They don't have the guts.",
}

func getRandomJoke() string {
	rand.Seed(time.Now().UnixNano())
	return jokes[rand.Intn(len(jokes))]
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6278421305318-6278445983078-EyA8Ug5pY1x1c1J4YVUFlIYL")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A068D29MCTU-6297710091361-70fafa83822687a1340790c2032a3c79b9ce5a5ad386314a80c3ae907da2ba7d")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	bot.Command("nam batiye", &slacker.CommandDefinition{
		Description: "Get information about the bot",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			botName := "Bhupendra Jogi"
			response.Reply(fmt.Sprintf("%s", botName))
		},
	})

	bot.Command("time?", &slacker.CommandDefinition{
		Description: "Get the current time",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			currentTime := time.Now().Format("15:04:05")
			response.Reply(fmt.Sprintf("Current Time: %s", currentTime))
		},
	})

	bot.Command("date?", &slacker.CommandDefinition{
		Description: "Get the current date",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			currentDate := time.Now().Format("2006-01-02")
			response.Reply(fmt.Sprintf("Current Date: %s", currentDate))
		},
	})

	bot.Command("joke", &slacker.CommandDefinition{
		Description: "Tell a random joke",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			randomJoke := getRandomJoke()
			response.Reply(randomJoke)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
