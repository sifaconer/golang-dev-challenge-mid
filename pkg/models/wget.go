package models

import (
	"os"
	"path"
	"strings"
)

type WGET struct {
	URI       string // download uri
	FileName  string // file name to save
	Directory string // directory to save
}

// PathFileName: returns the full path with the file name
func (w *WGET) PathFileName() string {
	if w.FileName == "" {
		w.FileName = w.getNameFileFromURL()
	}
	w.createDirectory()
	return path.Join(w.Directory, w.FileName) // join directory to file name
}

// getNameFileFromURL: extract the name from the given url of the last segment
func (w *WGET) getNameFileFromURL() string {
	name := strings.Split(w.URI, "/")
	return name[len(name)-1] // get last segment in url
}

// createDirectory: check if the directory does not exist to create it
func (w *WGET) createDirectory() {
	if _, err := os.Stat(w.Directory); os.IsNotExist(err) {
		os.MkdirAll(w.Directory, os.ModePerm) // create directory
	}
}
