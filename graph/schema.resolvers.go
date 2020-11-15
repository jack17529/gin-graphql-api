package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"graphql-srv/graph/generated"
	"graphql-srv/graph/model"
	"graphql-srv/repo"
)

var vr repo.VideoDB = repo.New()

func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	// panic(fmt.Errorf("not implemented"))

	video := &model.Video{
		ID:    getToken(10),
		Title: input.Title,
		URL:   input.URL,
		Author: &model.User{
			ID:   input.UserID,
			Name: "user " + input.UserID,
		},
	}

	err := vr.Save(video)
	// r.videos = append(r.videos, video)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (r *queryResolver) GetVideos(ctx context.Context) ([]*model.Video, error) {
	// panic(fmt.Errorf("not implemented"))
	v, err := vr.FindAll()
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
