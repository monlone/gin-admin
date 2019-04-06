package bll

import (
	"context"
	"github.com/goodcorn/src/errors"
	"github.com/goodcorn/src/model"
	"github.com/goodcorn/src/schema"
	"github.com/goodcorn/src/util"
)

// Merchant 示例程序
type Merchant struct {
	MerchantModel model.IMerchant `inject:"IMerchant"`
	CommonBll     *Common         `inject:""`
}

// QueryPage 查询分页数据
func (a *Merchant) QueryPage(ctx context.Context, params schema.MerchantQueryParam, pp *schema.PaginationParam) ([]*schema.Merchant, *schema.PaginationResult, error) {
	result, err := a.MerchantModel.Query(ctx, params, schema.MerchantQueryOptions{PageParam: pp})
	if err != nil {
		return nil, nil, err
	}
	return result.Data, result.PageResult, nil
}

// Get 查询指定数据
func (a *Merchant) Get(ctx context.Context, recordID string) (*schema.Merchant, error) {
	item, err := a.MerchantModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *Merchant) checkCode(ctx context.Context, code string) error {
	result, err := a.MerchantModel.Query(ctx, schema.MerchantQueryParam{
		Code: code,
	}, schema.MerchantQueryOptions{
		PageParam: &schema.PaginationParam{PageSize: -1},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.NewBadRequestError("编号已经存在")
	}
	return nil
}

// Create 创建数据
func (a *Merchant) Create(ctx context.Context, item schema.Merchant) (*schema.Merchant, error) {
	err := a.checkCode(ctx, item.Code)
	if err != nil {
		return nil, err
	}

	item.RecordID = util.MustUUID()
	item.Creator = a.CommonBll.GetUserID(ctx)
	err = a.MerchantModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, item.RecordID)
}

// Update 更新数据
func (a *Merchant) Update(ctx context.Context, recordID string, item schema.Merchant) (*schema.Merchant, error) {
	oldItem, err := a.MerchantModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.Code != item.Code {
		err := a.checkCode(ctx, item.Code)
		if err != nil {
			return nil, err
		}
	}

	err = a.MerchantModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.Get(ctx, recordID)
}

// Delete 删除数据
func (a *Merchant) Delete(ctx context.Context, recordID string) error {
	err := a.MerchantModel.Delete(ctx, recordID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus 更新状态
func (a *Merchant) UpdateStatus(ctx context.Context, recordID string, status int) error {
	return a.MerchantModel.UpdateStatus(ctx, recordID, status)
}
