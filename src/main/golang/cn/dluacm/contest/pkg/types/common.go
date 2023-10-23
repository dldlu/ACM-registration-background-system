package types

// PageRequest
// @Description: 分页请求
type PageRequest struct {
	Size int `uri:"size" binding:"required"`
	Num  int `uri:"num" binding:"required"`
}

type YearPageRequest struct {
	Year     int `uri:"year" binding:"required"`
	HalfYear int `uri:"halfYear" binding:"required"`
	Size     int `uri:"size" binding:"required"`
	Num      int `uri:"num" binding:"required"`
}

type YearRequest struct {
	Year string `json:"year" binding:"required"`
}

// PageResp
// @Description: 分页响应
type PageResp struct {
	ItemTotal int64 `json:"item_total"` // 总元素数
	PageTotal int64 `json:"page_total"` // 总页数
	Array     any   `json:"array"`
}

// IdRequest
// @Description: 用户id
type IdRequest struct {
	Id int64 `uri:"id" binding:"required"`
}
