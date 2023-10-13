package db

import (
	"time"
)

type Author struct {
	Name string
}

type NewPoem struct {
	Title        string    `firestore:"title" json:"title"`
	Author       string    `firestore:"author" json:"author"`
	Poem         []string  `firestore:"poem" json:"poem"`
	SampleLength int8      `firestore:"sampleLength" json:"sampleLength"`
	Categories   []string  `firestore:"categories" json:"categories"`
	ReleaseDate  time.Time `firestore:"releaseDate" json:"releaseDate"`
}

type Poem struct {
	NewPoem
	ID string `firestore:"-" json:"id"`
}

type Category struct {
	Name        string
	Description string
}

type NewWork struct {
	Name         string
	Poems        []poemReference
	CoverPoem    poemReference `firestore:"coverPoem"`
	NextPoemDate time.Time     `firestore:"nextPoemDate"`
}

type Work struct {
	NewWork
	ID string `firestore:"-"`
}

type Reader struct {
	ID            string `firestore:"-" json:"id"`
	Name          string
	FavoritePoems []poemReference `firestore:"favoritePoems" json:"favoritePoems"`
	FavoriteLines []favoriteLine  `firestore:"favoriteLines" json:"favoriteLines"`
	Lists         []PoemList
	MostRecent    poemReaderReference `firestore:"mostRecent" json:"mostRecent"`
}

type PoemList struct {
	Name  string
	Poems []poemReaderReference
}

type poemReference struct {
	ID     string
	Title  string
	Sample []lineReference
}

type poemReaderReference struct {
	poemReference
	ParentWork string `firestore:"parentWork" json:"parentWork"`
}

type lineReference struct {
	Text  string
	Index int8
}

type favoriteLine struct {
	PoemID string
	Line   lineReference
}
