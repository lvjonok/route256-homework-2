package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateImage(t *testing.T) {
	client, ctx := Prepare(t)

	img := []byte{0, 1, 2, 3, 4, 5, 6}
	url := "gitlab.com/someimage.svg"

	_, err := client.CreateImage(ctx, img, url)
	require.NoError(t, err)
}
