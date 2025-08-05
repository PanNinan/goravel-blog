package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type TopicController struct {
	// Dependent services
}

func NewTopicController() *TopicController {
	return &TopicController{
		// Inject services
	}
}

func (r *TopicController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *TopicController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *TopicController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *TopicController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *TopicController) Destroy(ctx http.Context) http.Response {
	return nil
}
