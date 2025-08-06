package request

import (
	"github.com/goravel/framework/contracts/database/orm"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   // 关键字
}

func (r *PageInfo) Paginate(q orm.Query, dest any, total *int64) error {
	if r.Page <= 0 {
		r.Page = 1
	}
	switch {
	case r.PageSize > 100:
		r.PageSize = 100
	case r.PageSize <= 0:
		r.PageSize = 10
	}
	return q.Paginate(r.Page, r.PageSize, dest, total)
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type Empty struct{}
