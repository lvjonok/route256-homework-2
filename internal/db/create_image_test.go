package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateImage(t *testing.T) {
	client, ctx := Prepare(t)

	img := []byte{0, 1, 2, 3, 4, 5, 6}
	url := "gitlab.com/someimage.svg"

	resID, err := client.CreateImage(ctx, img, url)
	require.NoError(t, err)

	resID2, err := client.CreateImage(ctx, img, url)
	require.NoError(t, err)

	// they should be the same, because we didn't update image on this url
	require.Equal(t, resID, resID2)

	img = []byte{0, 1, 2, 3, 4, 5, 6, 8}
	resID3, err := client.CreateImage(ctx, img, url)
	require.NoError(t, err)

	// they should not be equal, because we updated image -> newID
	require.NotEqual(t, resID, resID3)
}
