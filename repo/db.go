package repo

import (
	"database/sql"
	"graphql-srv/graph/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dbsource = "<USERNAME>:<PASSWORD>@tcp(localhost:3306)/videoDB"

var db *sql.DB

func GetSession() (*sql.DB, error) {
	var err error

	db, err = sql.Open("mysql", dbsource)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully got connected to mysql")
	return db, nil
}

type VideoDB interface {
	Save(video *model.Video) error
	FindAll() ([]*model.Video, error)
}

type DBservice struct{}

func New() VideoDB {
	// log.Println("Connected to MySQL")
	return &DBservice{}
}

func (d *DBservice) Save(video *model.Video) error {
	db, err := GetSession()
	if err != nil {
		log.Fatal(err)
	}

	stmt, prepareErr := db.Prepare(`
		INSERT INTO videos(id, title, url, author_id, author_name)
		VALUES (?,?,?,?,?)
	`)
	if prepareErr != nil {
		// log.Fatal(prepareErr)
		return prepareErr
	}

	_, execErr := stmt.Exec(video.ID, video.Title, video.URL, video.Author.ID, video.Author.Name)
	if execErr != nil {
		// log.Fatal(execErr)
		return execErr
	}

	log.Println("Successfully added the video")
	return nil
}

func (d *DBservice) FindAll() ([]*model.Video, error) {
	db, err := GetSession()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT * FROM videos`)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var (
		videos                               = []*model.Video{}
		id, authorID, title, url, authorName string
	)

	for rows.Next() {
		if err = rows.Scan(&id, &title, &url, &authorID, &authorName); err != nil {
			// log.Fatal(err)
			return nil, err
		}
		videos = append(videos, &model.Video{
			ID:     id,
			Title:  title,
			URL:    url,
			Author: &model.User{ID: authorID, Name: authorName},
		})
	}

	log.Println("Successfully found all videos")
	return videos, nil
}
