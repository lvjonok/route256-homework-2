package svgconv_test

import (
	"io/ioutil"
	"os"
	"reflect"
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

	exp, err := ioutil.ReadFile("test/expected.png")
	require.NoError(t, err)

	bytes, err := svgconv.ImagePartToBytes("https://ege.sdamgia.ru/get_file?id=29490")
	require.NoError(t, err)

	require.Truef(t, reflect.DeepEqual(exp, bytes), "resulted image is not the same")
}
