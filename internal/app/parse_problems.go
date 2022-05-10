package app

import (
	"context"
	"log"
	"sync"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/parser"
	pb "gitlab.ozon.dev/lvjonok/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) ParseProblems(ctx context.Context, req *pb.ParseProblemsRequest) (*pb.ParseProblemsResponse, error) {

	// start parser

	categories, err := parser.ParseCategories()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse categories of problems: %v", err)
	}

	problemsChan := make(chan *models.Problem)

	var wg sync.WaitGroup

	for _, cat := range categories {
		if err := s.DB.CreateCategory(ctx, models.Category{
			CategoryID: models.ID(cat.CategotyId),
			TaskNumber: cat.Problem,
			Title:      cat.Title,
		}); err != nil {
			log.Printf("failed to create category, err: %v", err)
			continue
		}

		wg.Add(1)
		go func(cat *parser.ProblemCategory) {
			defer wg.Done()

			ids, err := parser.ParseProblemsIds(cat.CategotyId)
			if err != nil {
				log.Printf("err: %v", err)
			}

			for _, id := range ids {
				wg.Add(1)
				go func(problemId int) {
					defer wg.Done()

					problem, err := parser.ParseProblem(problemId)
					if err != nil {
						log.Printf("problem err: %v", err)
					}

					if err == nil {
						problemsChan <- &models.Problem{
							ProblemID:    models.ID(problem.ProblemId),
							CategoryID:   models.ID(cat.CategotyId),
							ProblemImage: problem.ProblemImage,
							Parts:        problem.Parts,
							Answer:       problem.Answer,
						}
					}
				}(id)
			}

		}(cat)
	}

	go func() {
		wg.Wait()
		close(problemsChan)
	}()

	for pproblem := range problemsChan {
		// log.Printf("new %v", pproblem.ProblemId)
		if err := s.DB.CreateProblem(ctx, *pproblem); err != nil {
			log.Printf("failed to add new problem %v, err: %v", pproblem, err)
		}
	}

	return &pb.ParseProblemsResponse{}, nil
}
