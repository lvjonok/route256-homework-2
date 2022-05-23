package svgconv

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

func ImagePartToBytes(url string) ([]byte, error) {
	// download image and add it to the database
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to get req: %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		return nil, err
	}

	encurl := strings.Join(strings.Split(url, "/"), "-")

	svgFilename := path.Join("tmp", encurl)
	log.Printf("path: %v", svgFilename)

	err = os.WriteFile(svgFilename, raw, 0777)
	if err != nil {
		log.Printf("failed to create file with new image, %v", err)
		return nil, err
	}

	imgbytes, err := svg2png(svgFilename)
	if err != nil {
		log.Printf("error in svg2path, %v", err)
		return nil, err
	}
	return imgbytes, nil
}

func svg2png(svgfpath string) ([]byte, error) {
	app := "inkscape"

	variants := [][]string{
		{
			"-b",
			"FFFFFF",
			"-d",
			"300",
			fmt.Sprintf("--export-filename=%s.png", svgfpath),
			svgfpath,
		},
		{
			"-b",
			"FFFFFF",
			"-d",
			"300",
			fmt.Sprintf("--export-png=%s.png", svgfpath),
			svgfpath,
		},
	}

	for _, v := range variants {
		cmd := exec.Command(app, v...)
		_, err := cmd.Output()
		if err != nil {
			log.Printf("got error, while converting, err: <%v>, trying next", err)
			continue
		}

		raw, err := ioutil.ReadFile(fmt.Sprintf("%s.png", svgfpath))
		if err != nil {
			return nil, fmt.Errorf("converted, but read error: <%v>", err)
		}

		return raw, nil
	}

	return nil, fmt.Errorf("failed to convert using each variant")
}
