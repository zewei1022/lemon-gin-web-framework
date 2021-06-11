package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type IdInfo struct {
	ID int `json:"id" form:"id"` // 主键ID
}

type IdsInfo struct {
	Ids []int `json:"ids" form:"ids"`
}
