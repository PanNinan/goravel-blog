package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type ReplyController struct {
	// Dependent services
}

func NewReplyController() *ReplyController {
	return &ReplyController{
		// Inject services
	}
}

func (r *ReplyController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *ReplyController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *ReplyController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *ReplyController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *ReplyController) Destroy(ctx http.Context) http.Response {
	return nil
}
