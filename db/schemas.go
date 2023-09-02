package db

import (
	"time"
)

type Author struct {
	Name string
}

type NewPoem struct {
	Title       string
	Author      Author
	Text        []string
	Categories  []categoryReference
	ReleaseDate time.Time `firestore:"release_date"`
}

type Poem struct {
	NewPoem
	ID string `firestore:"id"`
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
	CoverPoem    poemReference `firestore:"cover_poem"`
	NextPoemDate time.Time     `firestore:"next_poem_date"`
}

type Work struct {
	NewWork
	ID string `firestore:"id"`
}

type Reader struct {
	ID            string `json:"id"`
	Name          string
	FavoritePoems []poemReference `json:"favorite_poems"`
	FavoriteLines []favoriteLine  `json:"favorite_lines"`
	Lists         []PoemList
	MostRecent    poemReaderReference `json:"most_recent"`
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
	ParentWork string `json:"parent_work"`
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
	ID   string
	Name string
}
