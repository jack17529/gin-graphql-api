package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-srv/graph/generated"
	"graphql-srv/graph/model"
	"math/rand"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	// panic(fmt.Errorf("not implemented"))
	video := &model.Video{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Title: input.Title,
		URL:   input.URL,
		Author: &model.User{
			ID:   input.UserID,
			Name: "user " + input.UserID,
		},
	}
	r.videos = append(r.videos, video)
	return video, nil
}

func (r *queryResolver) GetVideos(ctx context.Context) ([]*model.Video, error) {
	// panic(fmt.Errorf("not implemented"))
	return r.videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
