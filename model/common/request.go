package common

type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 頁數
	PageSize int    `json:"pageSize" form:"pageSize"` // 每頁大小
	Keyword  string `json:"keyword" form:"keyword"`   // 關鍵字
}

type GetById struct {
	ID int `json:"id" form:"id"` // 主鍵 ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
