package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	Name string
}

type Poem struct {
	Title       string
	Author      Author
	Text        []string
	Categories  []categoryReference
	ReleaseDate primitive.DateTime `bson:"release_date"`
}

type Category struct {
	Name        string
	Description string
}

type Work struct {
	Name         string
	Poems        []poemReference
	CoverPoem    poemReference      `bson:"cover_poem"`
	NextPoemDate primitive.DateTime `bson:"next_poem_date"`
}

type Reader struct {
	Name          string
	FavoritePoems []poemReference
	FavoriteLines []favoriteLine
	Lists         []PoemList
	MostRecent    poemReference `bson:"most_recent"`
}

type PoemList struct {
	Name  string
	Poems []poemReference
}

type poemReference struct {
	ID     primitive.ObjectID
	Title  string
	Sample []lineReference
}

type favoriteLine struct {
	PoemID primitive.ObjectID
	Line   lineReference
}

type lineReference struct {
	Text  string
	Index int8
}

type categoryReference struct {
	ID   primitive.ObjectID
	Name string
}
