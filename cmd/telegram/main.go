package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/config"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/tg"
	homework_2 "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.New("config.yaml")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(cfg.Server.Host+":"+cfg.Server.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := homework_2.NewMathHelperClient(conn)
	tgClient := tg.New(cfg.Telegram.BotAPI)

	for update := range tgClient.GetUpdatesChan() {
		chatID := update.Message.Chat.ID

		ok, cmd, params := update.IsCommand()
		if !ok {
			continue
		}

		switch cmd {
		case "/start":
			tgClient.SendMessage(chatID, "hello!")
		case "/random":
			resp, err := client.GetRandom(ctx, &homework_2.GetRandomRequest{ChatId: int64(chatID)})
			if err != nil {
				log.Printf("failed to get random problem, err: %v", err)
			}

			for _, p := range resp.Problem.Description {
				tgClient.SendMessage(chatID, p)
			}
		case "/problem":
			number, err := strconv.Atoi(params)
			if err != nil {
				tgClient.SendMessage(chatID, "do not think that you put number as parameter")
				continue
			}
			resp, err := client.GetProblem(ctx, &homework_2.GetProblemRequest{ChatId: int64(chatID), TaskNumber: int64(number)})
			if err != nil {
				log.Printf("failed to get random problem, err: %v", err)
				continue
			}

			for _, p := range resp.Problem.Description {
				tgClient.SendMessage(chatID, p)
			}
		case "/check":
			answer := params
			resp, err := client.CheckAnswer(ctx, &homework_2.CheckAnswerRequest{ChatId: int64(chatID), Answer: answer})
			if err != nil {
				log.Printf("ERROR: %v", err)
				continue
			}

			tgClient.SendMessage(update.Message.Chat.ID, fmt.Sprintf("problem: %v\ncorrect answer: %v\nyour result: %v\n", resp.ProblemId, resp.Answer, resp.Result))
		case "/stat":
			// TODO: implement me
		case "/rating":
			// TODO: implement me
		}
	}
}
