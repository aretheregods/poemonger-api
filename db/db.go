package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"poemonger/api/with_time"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func InitializeDB() *sql.DB {
	url := os.Getenv("DB_URL")
	tkn := os.Getenv("DB_TKN")
	str := fmt.Sprintf("%v?authToken=%v", url, tkn)

	db, err := sql.Open("libsql", str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", str, err)
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}

func InitializeTables(d *sql.DB) {
	ctxTimeout, cancel := withTime.WithTimeout(10)
	defer cancel()

	_, err := d.ExecContext(ctxTimeout, `create table if not exists works (
		id integer primary key, 
		name text not null,
		description text not null,
		categories blob not null,
		poems blob not null,
		cover_poem blob not null,
		release_date datetime
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = d.ExecContext(ctxTimeout, `create table if not exists poetry (
		id integer primary key, 
		title text not null, 
		author blob not null, 
		sample_length integer not null, 
		description text not null,
		poem blob not null,
		categories blob not null,
		release_date datetime not null,
		audio_url text not null,
		video_url text not null,
		work integer not null,
		foreign key(work) references works(rowid)
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = d.ExecContext(ctxTimeout, `create table if not exists categories (
		id integer primary key,
		name text not null,
		description text not null
	);`)
	if err != nil {
		log.Fatal(err)
	}

}
