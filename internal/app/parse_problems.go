package app

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/parser"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/svgconv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ParseProblems(ctx context.Context) error {
	categories, err := parser.ParseCategories()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to parse categories of problems: <%v>", err)
	}

	problemsChan := make(chan *models.Problem)

	var wg sync.WaitGroup

	for _, cat := range categories {
		categoryDbID, err := s.AddCategory(ctx, &models.Category{
			CategoryID: models.ID(cat.CategotyId),
			TaskNumber: cat.Problem,
			Title:      cat.Title,
		})
		if err != nil {
			log.Printf("failed to create category, err: <%v>", err)
			continue
		}

		wg.Add(1)
		go func(cat *parser.ProblemCategory, catDBIdx models.ID) {
			defer wg.Done()

			ids, err := parser.ParseProblemsIds(cat.CategotyId)
			if err != nil {
				log.Printf("failed to parse problem ids, err: <%v>", err)
				return
			}

			for _, id := range ids {
				wg.Add(1)
				go func(problemId int) {
					defer wg.Done()

					problem, err := parser.ParseProblem(problemId)
					if err != nil {
						log.Printf("failed to parse problem err: <%v>", err)
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
					log.Printf("failed to convert image part to bytes, err: <%v>", err)
					broke = true
					break
				}

				imageId, err := s.AddImage(ctx, imgbytes, part)
				if err != nil {
					log.Printf("failed to update image in database, <%v>", err)
					broke = true
					break
				}

				part = fmt.Sprintf("{%d}", *imageId)
			}
			dbproblem.Parts = append(dbproblem.Parts, part)
		}
		if broke {
			log.Printf("failed to prepare problem: %v, continuing", pproblem)
			continue
		}

		if dbproblem.ProblemImage != "" {
			imgbytes, err := svgconv.ImagePartToBytes(dbproblem.ProblemImage)
			if err != nil {
				log.Printf("failed to convert image part to bytes, err: <%v>, continuing", err)
				continue
			}

			imageId, err := s.AddImage(ctx, imgbytes, dbproblem.ProblemImage)
			if err != nil {
				log.Printf("failed to update image in database, err: <%v>, continuing", err)
				continue
			}
			dbproblem.ProblemImage = fmt.Sprintf("{%d}", *imageId)
		}

		if _, err := s.AddProblem(ctx, &dbproblem); err != nil {
			log.Printf("failed to add new problem %v, err: <%v>", pproblem, err)
		}
	}

	return nil
}

func (s *Service) AddProblem(ctx context.Context, dbproblem *models.Problem) (*models.ID, error) {
	oldProblem, err := s.DB.GetProblemByProblemID(ctx, dbproblem.ProblemID)
	if err != nil && err != db.ErrNotFound {
		return nil, fmt.Errorf("failed to check old problem, err: <%v>", err)
	}

	// check if we do not have to create new problem in database
	if err != db.ErrNotFound && oldProblem.Equal(dbproblem) {
		return &oldProblem.ID, nil
	}

	pID, err := s.DB.CreateProblem(ctx, *dbproblem)
	if err != nil {
		return nil, fmt.Errorf("failed to create problem, err: <%v>", err)
	}

	return pID, nil
}

func (s *Service) AddImage(ctx context.Context, img []byte, link string) (*models.ID, error) {
	// try to get old image
	oldImage, err := s.DB.GetImageByHref(ctx, link)
	if err != nil && err != db.ErrNotFound {
		return nil, fmt.Errorf("failed to check old image, err: <%v>", err)
	}

	// check if we do not have to create new image in database
	if err != db.ErrNotFound && oldImage.Equal(&models.Image{Content: img, Href: link}) {
		return &oldImage.ID, nil
	}

	imgID, cerr := s.DB.CreateImage(ctx, img, link)
	if cerr != nil {
		return nil, fmt.Errorf("failed to create image, err: <%v>", err)
	}

	return imgID, nil
}

func (s *Service) AddCategory(ctx context.Context, cat *models.Category) (*models.ID, error) {
	// try to get old category
	category, err := s.DB.GetCategoryByID(ctx, cat.CategoryID)
	if err != nil && err != db.ErrNotFound {
		return nil, fmt.Errorf("failed to check old category, err: <%v>", err)
	}

	// check if we do not have to add new category
	if err != db.ErrNotFound && category.Equal(cat) {
		return &category.ID, nil
	}

	catID, err := s.DB.CreateCategory(ctx, *cat)
	if err != nil {
		return nil, fmt.Errorf("failed to create new category, err: <%v>", err)
	}

	return catID, nil
}
