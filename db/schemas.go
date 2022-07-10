package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Poetry struct {
	ID         primitive.ObjectID
	Title      string
	Author     string
	Categories []string
	Books      []Book
}

type Book struct {
	Poems        []poemReference
	PreviousPoem poemReference
	NextPoemDate primitive.DateTime
}

type Reader struct {
	Name string
}

type poemReference struct {
	ID          primitive.ObjectID
	Title       string
	ReleaseDate primitive.DateTime
}
