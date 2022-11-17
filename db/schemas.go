package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Poetry struct {
	ID          primitive.ObjectID
	Title       string
	Author      string
	Categories  []string
	Collections []Collection
	ReleaseDate primitive.DateTime
}

type Collection struct {
	Name         string
	Poems        []poemReference
	MostRecent   poemReference
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
	ID          primitive.ObjectID
	Title       string
	ReleaseDate primitive.DateTime 
}
