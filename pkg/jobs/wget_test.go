package jobs

import (
	"testing"
)

var url string = "https://unec.edu.az/application/uploads/2014/12/pdf-sample.pdf"

func TestGetRequest(t *testing.T) {
	_, err := getRequest(url)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestWriteFile(t *testing.T) {
	resp, _ := getRequest(url)
	err := writeFile("/tmp/pdf-sample.pdf", resp)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}
