package repo

import (
	"graphql-srv/graph/model"

	_ "github.com/go-sql-driver/mysql"
)

type VideoDB interface {
	Save(video *model.Video) error
	FindAll() ([]*model.Video, error)
	FindVideoById(id string) (*model.Video, error)
}
