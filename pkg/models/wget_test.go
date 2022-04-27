package models

import (
	"os"
	"testing"
)

var model = WGET{
	URI: "https://unec.edu.az/application/uploads/2014/12/pdf-sample.pdf",
}

func TestPathFileName(t *testing.T) {
	got := model.PathFileName()
	want := "pdf-sample.pdf"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGetNameFileFromURL(t *testing.T) {
	got := model.getNameFileFromURL()
	want := "pdf-sample.pdf"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestCreateDirectory(t *testing.T) {
	model.Directory = "/tmp/files/inside"
	model.createDirectory()
	if _, err := os.Stat(model.Directory); os.IsNotExist(err) {
		t.Errorf("directory not found: %s", model.Directory)
	}
}

func TestPathFileNameWithDir(t *testing.T) {
	model.Directory = "/temp/pdf"
	got := model.PathFileName()
	want := "/temp/pdf/pdf-sample.pdf"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestPathFileNameWithFileName(t *testing.T) {
	model.Directory = ""
	model.FileName = "sample.pdf"
	got := model.PathFileName()
	want := "sample.pdf"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
