package db

import (
	"time"
)

type Author struct {
	Name string
}

type NewPoem struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Poem        []string `json:"poem"`
	Categories  []categoryReference `json:"categories"`
	ReleaseDate time.Time `firestore:"releaseDate" json:"release_date"`
}

type Poem struct {
	NewPoem
	ID string `firestore:"-"`
}

type NewCategory struct {
	Name        string
	Description string
}

type Category struct {
	NewCategory
	ID string `firestore:"-"`
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
	ID            string `firestore:"-"`
	Name          string
	FavoritePoems []poemReference `firestore:"favoritePoems"`
	FavoriteLines []favoriteLine  `firestore:"favoriteLines"`
	Lists         []PoemList
	MostRecent    poemReaderReference `firestore:"mostRecent"`
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
	ParentWork string `firestore:"parentWork"`
}

type lineReference struct {
	Text  string
	Index int8
}

type favoriteLine struct {
	PoemID string
	Line   lineReference
}

type categoryReference struct {
	ID   string `json:"id"`
	Name string	`json:"name"`
}
