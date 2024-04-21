package db

import (
	"time"
)

type Author struct {
	Name string
}

type NewPoem struct {
	Title        string    `json:"title"`
	Author       Author    `json:"author"`
	Poem         []string  `json:"poem"`
	SampleLength int8      `json:"sampleLength"`
	Categories   []string  `json:"categories"`
	ReleaseDate  time.Time `json:"releaseDate"`
}

type Poem struct {
	NewPoem
	ID string `json:"id"`
}

type Category struct {
	ID          int `json:"id"`
	Name        string
	Description string
}

type NewWork struct {
	Name         string
	Poems        []poemReference
	CoverPoem    poemReference
	NextPoemDate time.Time
}

type Work struct {
	NewWork
	ID string
}

type Reader struct {
	ID            string `json:"id"`
	Name          string
	FavoritePoems []poemReference `json:"favoritePoems"`
	FavoriteLines []favoriteLine  `json:"favoriteLines"`
	Lists         []PoemList
	MostRecent    poemReaderReference `json:"mostRecent"`
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
	ParentWork string `json:"parentWork"`
}

type lineReference struct {
	Text  string
	Index int8
}

type favoriteLine struct {
	PoemID string
	Line   lineReference
}
