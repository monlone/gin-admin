package model

import (
	"context"

	"github.com/goodcorn/src/schema"
)

// IMerchant merchant存储接口
type IMerchant interface {
	// 查询数据
	Query(ctx context.Context, params schema.MerchantQueryParam, opts ...schema.MerchantQueryOptions) (*schema.MerchantQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.MerchantQueryOptions) (*schema.Merchant, error)
	// 创建数据
	Create(ctx context.Context, item schema.Merchant) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Merchant) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	// 更新状态
	UpdateStatus(ctx context.Context, recordID string, status int) error
}
