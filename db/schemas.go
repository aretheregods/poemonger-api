package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Poem struct {
	Title       string
	Author      string
	Text        []lineReference
	Categories  []categoryReference
	ReleaseDate primitive.DateTime
}

type Work struct {
	Name         string
	Poems        []poemReference
	CoverPoem    poemReference
	NextPoemDate primitive.DateTime
}

type Reader struct {
	Name       string
	Favorites  []poemReference
	Lists      []PoemList
	MostRecent poemReference
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
	ID   primitive.ObjectID
	Text string
	Poem poemReference
}

type categoryReference struct {
	ID   primitive.ObjectID
	Name string
}
