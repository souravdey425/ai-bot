// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"github.com/shomali11/slacker"
// 	// witai "github.com/wit-ai/wit-go/v2"
// )

// func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
// 	for event := range analyticsChannel {
// 		fmt.Println("Command event")
// 		fmt.Println(event.Timestamp)
// 		fmt.Println(event.Command)
// 		fmt.Println(event.Parameters)
// 		fmt.Println(event.Event)
// 		fmt.Println()
// 	}
// }

// func main() {
// 	godotenv.Load(".env")
// 	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

// 	// client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
// 	go printCommandEvents(bot.CommandEvents())

// 	bot.Command("query for bot:<message>", &slacker.CommandDefinition{
// 		Description: "Send any question to wolframe",
// 		Examples:    []string{"who is president of india"},
// 		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
// 			query := r.Param("message")
// 			fmt.Printf("/my data %v/", query)

// 			// if query != "" {
// 			// 	msg, _ := client.Parse(&witai.MessageRequest{
// 			// 		Query: query,
// 			// 	})

// 			// }

// 			// if params != "" {

// 			// 	msg, error := client.Parse(&witai.MessageRequest{
// 			// 		Query: params,
// 			// 	})

// 			// 	if error != nil {
// 			// 		fmt.Println(error)

// 			// 	}
// 			// 	fmt.Println(msg)
// 			// }
// 			err := w.Reply("Recived")
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		},
// 	})

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	err := bot.Listen(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"github.com/shomali11/slacker"
// 	// witai "github.com/wit-ai/wit-go/v2"
// )

// func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
// 	for event := range analyticsChannel {
// 		fmt.Println("Command event")
// 		fmt.Println(event.Timestamp)
// 		fmt.Println(event.Command)
// 		fmt.Println(event.Parameters)
// 		fmt.Println(event.Event)
// 		fmt.Println()
// 	}
// }

// func main() {
// 	godotenv.Load(".env")
// 	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

// 	// client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
// 	go printCommandEvents(bot.CommandEvents())

// 	bot.Command("query for bot:<message>", &slacker.CommandDefinition{
// 		Description: "Send any question to wolframe",
// 		Examples:    []string{"who is president of india"},
// 		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
// 			query := r.Param("message")
// 			fmt.Printf("|Received message: %v|\n", query)

// 			// if query != "" {
// 			// 	msg, _ := client.Parse(&witai.MessageRequest{
// 			// 		Query: query,
// 			// 	})

// 			// }

// 			// if params != "" {

// 			// 	msg, error := client.Parse(&witai.MessageRequest{
// 			// 		Query: params,
// 			// 	})

// 			// 	if error != nil {
// 			// 		fmt.Println(error)

// 			// 	}
// 			// 	fmt.Println(msg)
// 			// }
// 			err := w.Reply("Received")
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		},
// 	})

//		ctx, cancel := context.WithCancel(context.Background())
//		defer cancel()
//		err := bot.Listen(ctx)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package main

import (
	"context"
	"encoding/json"
	"fmt"

	// "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

func main() {
	godotenv.Load(".env")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
	bot.Command("<hello>", &slacker.CommandDefinition{
		Description: "Responds to hello command",
		Handler: func(botContext slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			message := request.Param("hello")
			msg, _ := client.Parse(&witai.MessageRequest{
				Query: message,
			})
			data, _ := json.MarshalIndent(msg, "", "    ")
			rough := string(data[:])
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			answer := value.String()
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)

			response.Reply("Received: " + res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
