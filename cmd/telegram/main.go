package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/config"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/tg"
	homework_2 "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getImageID(raw string) int {
	imageID, _ := strconv.Atoi(raw[1 : len(raw)-1])
	return imageID
}

func isImageID(raw string) bool {
	// log.Printf("raw is image id %v", raw)
	return strings.HasSuffix(raw, "}") && strings.HasPrefix(raw, "{")
}

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
				continue
			}

			handleProblemSend(ctx, chatID, tgClient, client, resp.Problem)
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

			handleProblemSend(ctx, chatID, tgClient, client, resp.Problem)
		case "/check":
			answer := params
			resp, err := client.CheckAnswer(ctx, &homework_2.CheckAnswerRequest{ChatId: int64(chatID), Answer: answer})
			if err != nil {
				log.Printf("failed to query check answer, err: %v", err)
				continue
			}

			tgClient.SendMessage(chatID, fmt.Sprintf("problem: %v\ncorrect answer: %v\nyour result: %v\n", resp.ProblemId, resp.Answer, resp.Result))
		case "/stat":
			resp, err := client.GetStat(ctx, &homework_2.GetStatRequest{ChatId: int64(chatID)})
			if err != nil {
				log.Printf("failed to get query stat, err: <%v>", err)
				continue
			}

			msg := "Your statistics (correct/all):\n"
			for _, el := range resp.Stat {
				msg += fmt.Sprintf("Task %d: %d/%d\n", el.TaskNumber, el.Correct, el.All)
			}

			tgClient.SendMessage(chatID, msg)
		case "/rating":
			resp, err := client.GetRating(ctx, &homework_2.GetRatingRequest{ChatId: int64(chatID)})
			if err != nil {
				log.Printf("failed to query rating, err: <%v>", err)
				continue
			}

			msg := "You rating among all users\n"
			msg += fmt.Sprintf("Position: %d\nAmong: %d", resp.Position, resp.All)

			tgClient.SendMessage(chatID, msg)
		}
	}
}

func handleProblemSend(ctx context.Context,
	chatID int,
	tgClient *tg.Client,
	client homework_2.MathHelperClient,
	problem *homework_2.Problem) error {

	if problem.Image != "" {
		imageID := getImageID(problem.Image)
		im, err := client.GetImage(ctx, &homework_2.GetImageRequest{ImageId: int64(imageID)})
		if err != nil {
			log.Printf("tried to access image: %v", int64(imageID))
			log.Printf("errored: <%v>", err)
			return err
		}
		tgClient.SendPhoto(chatID, im.Image, "Problem image")
	}

	dl := len(problem.Description)
	for idx := 0; idx < dl; idx++ {
		curpart := problem.Description[idx]
		if isImageID(curpart) {
			im, err := client.GetImage(ctx, &homework_2.GetImageRequest{ImageId: int64(getImageID(curpart))})
			if err != nil {
				log.Printf("tried to access image: %v", int64(getImageID(curpart)))
				log.Printf("errored: <%v>", err)
			}
			tgClient.SendPhoto(chatID, im.Image, "additional image")
		} else if !isImageID(curpart) {
			if idx+1 < dl && isImageID(problem.Description[idx+1]) {
				im, err := client.GetImage(ctx, &homework_2.GetImageRequest{ImageId: int64(getImageID(problem.Description[idx+1]))})
				if err != nil {
					log.Printf("tried to access image: %v", int64(getImageID(curpart)))
					log.Printf("errored: <%v>", err)
					return err
				}
				tgClient.SendPhoto(chatID, im.Image, curpart)
				idx += 1
			} else {
				tgClient.SendMessage(chatID, curpart)
			}
		}
	}

	return nil
}
