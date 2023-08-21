package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Poem struct {
	Title       string
	Author      string
	Text        []lineReference
	Categories  []categoryReference
	ReleaseDate primitive.DateTime `bson:"release_date"`
}

type Work struct {
	Name         string
	Poems        []poemReference
	CoverPoem    poemReference      `bson:"cover_poem"`
	NextPoemDate primitive.DateTime `bson:"next_poem_date"`
}

type Reader struct {
	Name       string
	Favorites  []poemReference
	Lists      []PoemList
	MostRecent poemReference `bson:"most_recent"`
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

type lineReference struct {
	PoemID primitive.ObjectID `bson:"poem_id"`
	Text   string
	Index  int8
}

type categoryReference struct {
	ID   primitive.ObjectID
	Name string
}
