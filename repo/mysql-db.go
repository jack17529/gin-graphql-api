package repo

import (
	"database/sql"
	"graphql-srv/graph/model"
	"log"
)

const dbsource = "<USER>:<PASSWORD>@tcp(localhost:3306)/<DATABASE_NAME>"

var db *sql.DB

func getSession() (*sql.DB, error) {
	var err error

	db, err = sql.Open("backend-mysql", dbsource)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully got connected to mysql")
	return db, nil
}

type DBservice struct{}

func New() VideoDB {
	return &DBservice{}
}

func (d *DBservice) Save(video *model.Video) error {
	db, err := getSession()
	if err != nil {
		return err
	}

	stmt, prepareErr := db.Prepare(`
		INSERT INTO videos(id, title, url, author_id, author_name)
		VALUES (?,?,?,?,?)
	`)
	if prepareErr != nil {
		return prepareErr
	}

	_, execErr := stmt.Exec(video.ID, video.Title, video.URL, video.Author.ID, video.Author.Name)
	if execErr != nil {
		return execErr
	}

	log.Println("Successfully added the video")
	return nil
}

func (d *DBservice) FindAll() ([]*model.Video, error) {
	db, err := getSession()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT * FROM videos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		videos                               = []*model.Video{}
		id, authorID, title, url, authorName string
	)

	for rows.Next() {
		if err = rows.Scan(&id, &title, &url, &authorID, &authorName); err != nil {
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

func (d *DBservice) FindVideoById(id string) (*model.Video, error) {

	db, err := getSession()
	if err != nil {
		return nil, err
	}

	var title, url, authorID, authorName string

	err = db.QueryRow(`
		SELECT title, url, author_id, author_name
		FROM videos
		WHERE id=?
		`, id).Scan(&title, &url, &authorID, &authorName)
	if err != nil {
		return nil, err
	}

	return &model.Video{
		ID:     id,
		Title:  title,
		URL:    url,
		Author: &model.User{ID: authorID, Name: authorName},
	}, nil
}
