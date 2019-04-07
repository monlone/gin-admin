package schema

import "time"

// Merchant merchant对象
type Merchant struct {
	RecordID           string    `json:"record_id" swaggo:"false,记录ID"`
	Code               string    `json:"code" binding:"required" swaggo:"true,编号"`
	Description        string    `json:"description" swaggo:"false, 商家描述"`
	AddressDescription string    `json:"address_description" swaggo:"false, 地址描述"`
	Address            string    `json:"address" swaggo:"false, 商家地址"`
	Name               string    `json:"name" binding:"required" swaggo:"true,名称"`
	Phone              string    `json:"phone" swaggo:"true,手机号"`
	Status             int       `json:"status" binding:"required,max=2,min=1" swaggo:"true,状态(1:启用 2:停用)"`
	Creator            string    `json:"creator" swaggo:"false,创建者"`
	CreatedAt          time.Time `json:"created_at" swaggo:"false,创建时间"`
}

// MerchantQueryParam 查询条件
type MerchantQueryParam struct {
	Code     string // 编号
	Name     string
	Status   int    // 状态(1:启用 2:停用)
	LikeCode string // 编号(模糊查询)
	LikeName string // 名称(模糊查询)
}

// MerchantQueryOptions merchant对象查询可选参数项
type MerchantQueryOptions struct {
	PageParam *PaginationParam // 分页参数
}

// MerchantQueryResult merchant对象查询结果
type MerchantQueryResult struct {
	Data       []*Merchant
	PageResult *PaginationResult
}
