package app

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/parser"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/svgconv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ParseProblems(ctx context.Context) error {

	// start parser

	categories, err := parser.ParseCategories()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to parse categories of problems: %v", err)
	}

	problemsChan := make(chan *models.Problem)

	var wg sync.WaitGroup

	// categories := []*parser.ProblemCategory{{Problem: 3, CategotyId: 112, Title: "some freaky good title"}}

	for _, cat := range categories {
		categoryDbID, err := s.DB.CreateCategory(ctx, models.Category{
			CategoryID: models.ID(cat.CategotyId),
			TaskNumber: cat.Problem,
			Title:      cat.Title,
		})
		if err != nil {
			log.Printf("failed to create category, err: %v", err)
			continue
		}

		wg.Add(1)
		go func(cat *parser.ProblemCategory, catDBIdx models.ID) {
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
							CategoryID:   catDBIdx,
							ProblemImage: problem.ProblemImage,
							Parts:        problem.Parts,
							Answer:       problem.Answer,
						}
					}
				}(id)
			}

		}(cat, *categoryDbID)
	}

	go func() {
		wg.Wait()
		close(problemsChan)
	}()

	for pproblem := range problemsChan {
		dbproblem := *pproblem
		dbproblem.Parts = []string{}

		broke := false

		for _, part := range pproblem.Parts {
			if strings.HasPrefix(part, "https://") {
				imgbytes, err := svgconv.ImagePartToBytes(part)
				if err != nil {
					log.Printf("failed to convert image part to bytes, err: %v", err)
					broke = true
					break
					// continue
				}

				imageId, err := s.DB.CreateImage(ctx, imgbytes, part)
				if err != nil {
					log.Printf("failed to update image in database, %v", err)
					broke = true
					break
					// return err
				}

				part = fmt.Sprintf("{%d}", *imageId)
			}
			dbproblem.Parts = append(dbproblem.Parts, part)
		}
		if broke {
			log.Printf("failed: %v", pproblem)
			continue
		}

		if dbproblem.ProblemImage != "" {
			imgbytes, err := svgconv.ImagePartToBytes(dbproblem.ProblemImage)
			if err != nil {
				log.Printf("failed to convert image part to bytes, err: %v", err)
				continue
				return err
			}

			imageId, err := s.DB.CreateImage(ctx, imgbytes, dbproblem.ProblemImage)
			if err != nil {
				log.Printf("failed to update image in database, %v", err)
				continue
				return err
			}
			dbproblem.ProblemImage = fmt.Sprintf("{%d}", *imageId)
		}

		if _, err := s.DB.CreateProblem(ctx, dbproblem); err != nil {
			log.Printf("failed to add new problem %v, err: %v", pproblem, err)
		}
	}

	return nil
}
