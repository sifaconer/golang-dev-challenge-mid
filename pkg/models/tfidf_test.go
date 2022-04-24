package models

import "testing"

var corpus = []TFIDF{
	{
		FileName:     "test.txt",
		Word:         "fox",
		CountWord:    2,
		AllWord:      12,
		ContainWord:  true,
		DocumentWord: map[string]int{"a": 2, "quick": 1, "brown": 1, "fox": 2, "jumps": 1, "over": 1, "the": 1, "lazy": 1, "What": 1},
	},
	{
		FileName:     "test.txt",
		Word:         "fox",
		CountWord:    3,
		ContainWord:  true,
		AllWord:      12,
		DocumentWord: map[string]int{"a": 2, "quick": 1, "brown": 1, "fox": 3, "jumps": 1, "over": 1, "the": 1, "lazy": 1, "What": 1},
	},
}

func TestTF(t *testing.T) {
	corpus[0].Tf()

	want := float64(2) / float64(12)
	if corpus[0].TF != want {
		t.Errorf("got %f, wanted %f", corpus[0].TF, want)
	}
}

func TestTFIDF(t *testing.T) {
	corpus[0].TfIdf(0)
	want := 0.0

	if corpus[0].TFIDF != want {
		t.Errorf("got %f, wanted %f", corpus[0].TFIDF, want)
	}
}

func TestExistWord(t *testing.T) {
	corpus[0].ExistWord()
	want := true

	if !corpus[0].ContainWord {
		t.Errorf("got %v, wanted %v", corpus[0].ContainWord, want)
	}
}

func TestTotalWord(t *testing.T) {
	corpus[0].TotalWord()
	want := 2

	if corpus[0].CountWord != want {
		t.Errorf("got %d, wanted %d", corpus[0].CountWord, want)
	}
}
