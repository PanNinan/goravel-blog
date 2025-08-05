package controllers

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel-blog/app/models"
	"goravel-blog/app/models/common/request"
	"goravel-blog/app/models/common/response"
)

type CategoryController struct {
	// Dependent services
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		// Inject services
	}
}

func (r *CategoryController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *CategoryController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *CategoryController) Store(ctx http.Context) http.Response {
	input := request.CategoryStoreReq{}
	err := ctx.Request().Bind(&input)
	if err != nil {
		return response.Result(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	var category models.Category
	err = mapstructure.Decode(input, &category)
	if err != nil {
		facades.Log().Error("category convert failed: %v", err)
		return response.InternalServerError(ctx, "failed")
	}
	if err = facades.Orm().WithContext(ctx).Query().Create(&category); err != nil {
		facades.Log().Error("Create category failed: %v", err)
		return response.InternalServerError(ctx, "Create category failed")
	}
	return response.Ok(ctx)
}

func (r *CategoryController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	return nil
}
