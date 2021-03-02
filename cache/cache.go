package cache

import "graphql-srv/graph/model"

type VideoCache interface {
	Set(key string, value *model.Video) error
	Get(key string) (*model.Video, error)
}
