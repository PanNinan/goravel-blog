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
	input := request.CategoryIndexReq{}
	if err := ctx.Request().Bind(&input); err != nil {
		return response.Result(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	query := facades.Orm().WithContext(ctx).Query().Model(&models.Category{}).OrderBy("id", "asc")
	if input.Keyword != "" {
		query = query.Where("name like ?", "%"+input.Keyword+"%")
	}
	var total int64
	list := make([]models.Category, 0)
	if err := input.Paginate(query, &list, &total); err != nil {
		facades.Log().Error("Get category list failed: %v", err)
		return response.InternalServerError(ctx, "Get category list failed")
	}
	res := response.PageResult{
		List:     list,
		Page:     input.Page,
		PageSize: input.PageSize,
		Total:    total,
	}
	return response.OkWithData(ctx, res)
}

func (r *CategoryController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	var category models.Category
	if err := facades.Orm().WithContext(ctx).Query().Where("id = ?", id).First(&category); err != nil {
		facades.Log().Error("Get category failed: %v", err)
		return response.InternalServerError(ctx, "Get category failed")
	}
	return response.OkWithData(ctx, category)
}

func (r *CategoryController) Store(ctx http.Context) http.Response {
	input := request.CategoryStoreReq{}
	if err := ctx.Request().Bind(&input); err != nil {
		return response.Result(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	var category models.Category
	err := mapstructure.Decode(input, &category)
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
	id := ctx.Request().RouteInt("id")
	var category models.Category
	if err := facades.Orm().WithContext(ctx).Query().Where("id = ?", id).First(&category); err != nil {
		facades.Log().Error("Get category failed: %v", err)
		return response.InternalServerError(ctx, "Get category failed")
	}
	input := request.CategoryStoreReq{}
	if err := ctx.Request().Bind(&input); err != nil {
		return response.Result(ctx, http.StatusBadRequest, nil, "Invalid input")
	}
	err := mapstructure.Decode(input, &category)
	if err != nil {
		facades.Log().Error("category convert failed: %v", err)
	}
	category.ID = uint(id)
	if err = facades.Orm().WithContext(ctx).Query().Save(&category); err != nil {
		facades.Log().Error("Update category failed: %v", err)
		return response.InternalServerError(ctx, "Update category failed")
	}
	return response.OkWithData(ctx, category)
}

func (r *CategoryController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	if _, err := facades.Orm().WithContext(ctx).Query().Where("id = ?", id).Delete(&models.Category{}); err != nil {
		facades.Log().Error("Delete category failed: %v", err)
		return response.InternalServerError(ctx, "Delete category failed")
	}
	return response.Ok(ctx)
}
