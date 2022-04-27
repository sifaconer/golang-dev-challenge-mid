package models

type TFIDF struct {
	FileName     string         // File name
	Word         string         // Word to search
	CountWord    int            //
	ContainWord  bool           // if exist in the document
	DocumentWord map[string]int // count total word
	TF           float64
	TFIDF        float64
	AllWord      int // total all word in document
}

// TF calculation of tf for the formula tf(t,d) = count of t in d / number of words in d
func (t *TFIDF) Tf() {
	if t.AllWord != 0 {
		t.TF = float64(t.DocumentWord[t.Word]) / float64(t.AllWord)
	}
}

// TfIdf: calculates the tf-idf receiving the idf
// as a parameter respecting the form tf-idf = tf*idf
func (t *TFIDF) TfIdf(idf float64) {
	t.TFIDF = t.TF * idf
}

// ExistWord: valid if exits word in document
func (t *TFIDF) ExistWord() {
	_, t.ContainWord = t.DocumentWord[t.Word]
}

// TotalWord: Count total word in document
func (t *TFIDF) TotalWord() {
	t.CountWord = t.DocumentWord[t.Word]
}
