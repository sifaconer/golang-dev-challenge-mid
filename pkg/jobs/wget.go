package jobs

import (
	"bufio"
	"challenge/pkg/models"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Wget(w models.WGET) string {
	resp, err := getRequest(w.URI)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pathFile := w.PathFileName()
	err = writeFile(pathFile, resp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return fmt.Sprintf("File saved to path: %s", pathFile)
}

// getRequest: makes a get request to a given url and returns the response
func getRequest(url string) (*http.Response, error) {
	tr := new(http.Transport)
	client := http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// writeFile: write the response given from the request to the url into a file
func writeFile(pathFileName string, resp *http.Response) error {
	file, err := os.OpenFile(pathFileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	buff := bufio.NewWriterSize(file, 1024*8)
	_, err = io.Copy(buff, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
