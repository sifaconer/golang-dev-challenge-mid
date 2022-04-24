package jobs

import (
	"challenge/pkg/models"
	"fmt"
	"strings"
)

func ETL(word string, files []string) {

	fileChannel := make(chan string)
	dataChannel := make(chan models.Transport)
	modelChannel := make(chan models.TFIDF)
	done := make(chan bool)

	go extract(fileChannel, dataChannel)
	go transform(word, dataChannel, modelChannel)
	go load(modelChannel, done)

	for _, v := range files {
		fileChannel <- v
	}

	close(fileChannel)
	<-done
}

func extract(fileChannel chan string, dataChannel chan models.Transport) {
	for pathFile := range fileChannel {
		var doc models.Transport
		file, err := openFile(pathFile)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			continue
		}
		doc.TotalWord, doc.Data = wordCount(file)
		doc.Info = pathFile
		dataChannel <- doc
	}

	close(dataChannel)
}

func transform(word string, dataChannel chan models.Transport, modelChannel chan models.TFIDF) {
	for doc := range dataChannel {
		var model models.TFIDF
		model.FileName = doc.Info
		model.Word = strings.ToLower(word)
		model.DocumentWord = doc.Data
		model.AllWord = doc.TotalWord
		model.ExistWord()
		model.TotalWord()
		model.Tf()

		modelChannel <- model
	}

	close(modelChannel)
}

func load(modelChannel chan models.TFIDF, done chan bool) {
	var corpus []models.TFIDF

	for model := range modelChannel {
		corpus = append(corpus, model)
	}

	valIDF := idf(corpus)

	header()
	for _, v := range corpus {
		v.TfIdf(valIDF)
		row(valIDF, v)
	}

	done <- true
}
