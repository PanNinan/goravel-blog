package request

type CategoryIndexReq struct {
	PageInfo
}

type CategoryStoreReq struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}
