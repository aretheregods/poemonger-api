package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	Name string
}

type NewPoem struct {
	Title       string
	Author      Author
	Text        []string
	Categories  []categoryReference
	ReleaseDate primitive.DateTime `bson:"release_date"`
}

type Poem struct {
	NewPoem
	ID primitive.ObjectID `bson:"_id" json:"id"`
}

type Category struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string
	Description string
}

type NewWork struct {
	Name         string
	Poems        []poemReference
	CoverPoem    poemReference      `bson:"cover_poem"`
	NextPoemDate primitive.DateTime `bson:"next_poem_date"`
}

type Work struct {
	NewWork
	ID primitive.ObjectID `bson:"_id" json:"id"`
}

type Reader struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Name          string
	FavoritePoems []poemReference
	FavoriteLines []favoriteLine
	Lists         []PoemList
	MostRecent    poemReaderReference `bson:"most_recent"`
}

type PoemList struct {
	Name  string
	Poems []poemReaderReference
}

type poemReference struct {
	ID     primitive.ObjectID
	Title  string
	Sample []lineReference
}

type poemReaderReference struct {
	poemReference
	ParentWork primitive.ObjectID
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
