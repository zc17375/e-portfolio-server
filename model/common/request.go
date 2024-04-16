package common

// Pagination 包含分頁結果
type Pagination struct {
	CurrentPage  int    `json:"currentPage" form:"currentPage" example:"1"` // 當前頁碼
	TotalPages   int    // 總頁數
	TotalRecords int    // 總記錄數
	PageSize     int    `json:"pageSize" form:"pageSize" example:"10"`       // 每頁大小
	KeyWords     string `json:"keywords" form:"keywords" example:"software"` // 關鍵字
	SortBy       string `json:"sortBy" form:"sortBy" example:"name"`
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
