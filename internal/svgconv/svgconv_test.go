package svgconv_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/lvjonok/homework-2/internal/svgconv"
)

func TestImagePartToBytes(t *testing.T) {
	// test preparation
	_, err := os.ReadDir("tmp")
	if err != nil {
		os.Mkdir("tmp", 0777)
	}

	img, err := svgconv.ImagePartToBytes("https://ege.sdamgia.ru/get_file?id=29490")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(img))
}
