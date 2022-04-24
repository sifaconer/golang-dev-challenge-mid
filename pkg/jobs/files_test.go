package jobs

import (
	"bytes"
	"challenge/pkg/models"
	"testing"
)

var corpus = []models.TFIDF{
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

func TestIDF(t *testing.T) {
	got := idf(corpus)
	want := 0.0
	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestAllWordCount(t *testing.T) {
	got, _ := wordCount(bytes.NewBufferString("a quick brown fox jumps over the lazy fox. What a fox!"))
	want := 12
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestWordCount(t *testing.T) {
	_, got := wordCount(bytes.NewBufferString("a quick brown fox jumps over the lazy fox. What a fox!"))
	want := 9
	if len(got) != want {
		t.Errorf("got %d, wanted %d", len(got), want)
	}
}
