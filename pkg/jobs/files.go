package jobs

import (
	"bufio"
	"challenge/pkg/models"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strings"
)

// All files in directory
func AllFiles(pathFile string) ([]string, error) {
	dir := path.Dir(pathFile)
	return listFilesInDir(dir)
}

func listFilesInDir(pathDir string) ([]string, error) {
	var totalFiles []string
	files, err := ioutil.ReadDir(pathDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			totalFiles = append(totalFiles, path.Join(pathDir, file.Name()))
		}
	}

	return totalFiles, nil
}

// Open File
func openFile(pathFile string) (*os.File, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Count all Word in file
func wordCount(read io.Reader) (int, map[string]int) {
	counts := map[string]int{}
	total := 0
	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		total++
		word := scanner.Text()
		word = strings.ToLower(word)
		word = strings.Replace(word, ",", "", -1)
		word = strings.Replace(word, ".", "", -1)
		word = strings.Replace(word, "!", "", -1)
		word = strings.Replace(word, "?", "", -1)

		counts[word]++
	}

	return total, counts
}

func idf(corpus []models.TFIDF) float64 {
	n := 0.0  // total documents
	df := 0.0 // total number of documents containing the word

	for _, v := range corpus {
		n += 1
		if v.ContainWord {
			df += 1
		}
	}

	if df != 0.0 {
		idf := math.Log(n / df)
		return idf
	}
	return 0.0
}

func header() {
	sep()
	fmt.Printf("%20s%16s%13s%16s%16s%16s%16s\n",
		"File Name", "Word", "T", "F", "TF", "IDF", "TF-IDF")
	sep()
}

func row(idf float64, v models.TFIDF) {
	fmt.Printf("%20s%16s%13d%16d%16f%16f%16f\n",
		path.Base(v.FileName), v.Word, v.DocumentWord[v.Word], len(v.DocumentWord), v.TF, idf, v.TFIDF)
	sep()
}

func sep() {
	fmt.Printf("    %s", "--------------------------------------------------------------------------------------------------------------\n")
}
