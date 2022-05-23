package app_test

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/app"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/db"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/models"
)

func TestAddProblemExists(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetProblemByProblemIDMock.Return(&models.Problem{
		ID:           1,
		ProblemID:    12345,
		CategoryID:   54321,
		ProblemImage: "some_image",
		Parts:        []string{"hello there"},
		Answer:       "ans",
	}, nil)

	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddProblem(ctx, &models.Problem{
		ProblemID:    12345,
		CategoryID:   54321,
		ProblemImage: "some_image",
		Parts:        []string{"hello there"},
		Answer:       "ans",
	})
	require.NoError(t, err)
	require.Equal(t, models.ID(1), *newID)
}

func TestAddProblemNew(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetProblemByProblemIDMock.Return(&models.Problem{
		ID:           1,
		ProblemID:    12345,
		CategoryID:   54321,
		ProblemImage: "some_image",
		Parts:        []string{"hello there"},
		Answer:       "ans",
	}, nil)

	newIDexp := models.ID(1)
	mockDB.CreateProblemMock.Return(&newIDexp, nil)

	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddProblem(ctx, &models.Problem{
		ProblemID:    12345,
		CategoryID:   54321,
		ProblemImage: "new image",
		Parts:        []string{"new part"},
		Answer:       "new ans",
	})
	require.NoError(t, err)
	require.Equal(t, newIDexp, *newID)
}

func TestAddProblemNotFound(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetProblemByProblemIDMock.Return(nil, db.ErrNotFound)

	newIDexp := models.ID(1)
	mockDB.CreateProblemMock.Return(&newIDexp, nil)

	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddProblem(ctx, &models.Problem{
		ProblemID:    12345,
		CategoryID:   54321,
		ProblemImage: "new image",
		Parts:        []string{"new part"},
		Answer:       "new ans",
	})
	require.NoError(t, err)
	require.Equal(t, newIDexp, *newID)
}

func TestAddImageExists(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetImageByHrefMock.Return(&models.Image{ID: models.ID(1), Content: []byte{0, 1, 2, 3}, Href: "gitlab.com"}, nil)

	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddImage(ctx, []byte{0, 1, 2, 3}, "gitlab.com")
	require.NoError(t, err)
	require.Equal(t, models.ID(1), *newID)
}

func TestAddImageNew(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetImageByHrefMock.Return(&models.Image{ID: models.ID(1), Content: []byte{0, 1, 2, 3, 4}, Href: "gitlab.com"}, nil)

	newIDexp := models.ID(1)
	mockDB.CreateImageMock.Return(&newIDexp, nil)

	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddImage(ctx, []byte{0, 1, 2, 3}, "gitlab.com")
	require.NoError(t, err)
	require.Equal(t, newIDexp, *newID)
}

func TestAddImageNotFound(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetImageByHrefMock.Return(nil, db.ErrNotFound)

	newIDexp := models.ID(1)
	mockDB.CreateImageMock.Return(&newIDexp, nil)

	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddImage(ctx, []byte{0, 1, 2, 3}, "gitlab.com")
	require.NoError(t, err)
	require.Equal(t, newIDexp, *newID)
}

func TestAddCategoryExists(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetCategoryByIDMock.Return(&models.Category{ID: models.ID(1), CategoryID: 12345, TaskNumber: 1, Title: "some title"}, nil)
	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddCategory(ctx, &models.Category{ID: models.ID(1), CategoryID: 12345, TaskNumber: 1, Title: "some title"})
	require.NoError(t, err)
	require.Equal(t, models.ID(1), *newID)
}

func TestAddCategoryNew(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetCategoryByIDMock.Return(&models.Category{ID: models.ID(1), CategoryID: 12345, TaskNumber: 1, Title: "some title"}, nil)
	newIDexp := models.ID(1)
	mockDB.CreateCategoryMock.Return(&newIDexp, nil)
	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddCategory(ctx, &models.Category{ID: models.ID(1), CategoryID: 54321, TaskNumber: 1, Title: "some title"})
	require.NoError(t, err)
	require.Equal(t, models.ID(1), *newID)
}

func TestAddCategoryNotFound(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockDB := app.NewDBMock(mc)
	mockDB.GetCategoryByIDMock.Return(nil, db.ErrNotFound)
	newIDexp := models.ID(1)
	mockDB.CreateCategoryMock.Return(&newIDexp, nil)
	srv := app.New(mockDB)

	ctx := context.Background()
	newID, err := srv.AddCategory(ctx, &models.Category{ID: models.ID(1), CategoryID: 54321, TaskNumber: 1, Title: "some title"})
	require.NoError(t, err)
	require.Equal(t, models.ID(1), *newID)
}
