package model

import (
	"context"
	"fmt"

	"github.com/goodcorn/src/errors"
	"github.com/goodcorn/src/logger"
	"github.com/goodcorn/src/model/gorm/entity"
	"github.com/goodcorn/src/schema"
	"github.com/goodcorn/src/service/gormplus"
)

// Merchant merchant存储
type Merchant struct {
	db *gormplus.DB
}

// Init 初始化
func (a *Merchant) Init(db *gormplus.DB) *Merchant {
	db.AutoMigrate(new(entity.Merchant))
	a.db = db
	return a
}

func (a *Merchant) getFuncName(name string) string {
	return fmt.Sprintf("gorm.model.Merchant.%s", name)
}

func (a *Merchant) getQueryOption(opts ...schema.MerchantQueryOptions) schema.MerchantQueryOptions {
	var opt schema.MerchantQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *Merchant) Query(ctx context.Context, params schema.MerchantQueryParam, opts ...schema.MerchantQueryOptions) (*schema.MerchantQueryResult, error) {
	span := logger.StartSpan(ctx, "查询数据", a.getFuncName("Query"))
	defer span.Finish()

	db := entity.GetMerchantDB(ctx, a.db).DB
	if v := params.Code; v != "" {
		db = db.Where("code=?", v)
	}
	if v := params.LikeCode; v != "" {
		db = db.Where("code LIKE ?", "%"+v+"%")
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Merchants
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询数据发生错误")
	}
	qr := &schema.MerchantQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaMerchants(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *Merchant) Get(ctx context.Context, recordID string, opts ...schema.MerchantQueryOptions) (*schema.Merchant, error) {
	span := logger.StartSpan(ctx, "查询指定数据", a.getFuncName("Get"))
	defer span.Finish()

	db := entity.GetMerchantDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.Merchant
	ok, err := a.db.FindOne(db, &item)
	if err != nil {
		span.Errorf(err.Error())
		return nil, errors.New("查询指定数据发生错误")
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaMerchant(), nil
}

// Create 创建数据
func (a *Merchant) Create(ctx context.Context, item schema.Merchant) error {
	span := logger.StartSpan(ctx, "创建数据", a.getFuncName("Create"))
	defer span.Finish()

	merchant := entity.SchemaMerchant(item).ToMerchant()
	result := entity.GetMerchantDB(ctx, a.db).Create(merchant)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("创建数据发生错误")
	}
	return nil
}

// Update 更新数据
func (a *Merchant) Update(ctx context.Context, recordID string, item schema.Merchant) error {
	span := logger.StartSpan(ctx, "更新数据", a.getFuncName("Update"))
	defer span.Finish()

	merchant := entity.SchemaMerchant(item).ToMerchant()
	result := entity.GetMerchantDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(merchant)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新数据发生错误")
	}
	return nil
}

// Delete 删除数据
func (a *Merchant) Delete(ctx context.Context, recordID string) error {
	span := logger.StartSpan(ctx, "删除数据", a.getFuncName("Delete"))
	defer span.Finish()

	result := entity.GetMerchantDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.Merchant{})
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("删除数据发生错误")
	}
	return nil
}

// UpdateStatus 更新状态
func (a *Merchant) UpdateStatus(ctx context.Context, recordID string, status int) error {
	span := logger.StartSpan(ctx, "更新状态", a.getFuncName("UpdateStatus"))
	defer span.Finish()

	result := entity.GetMerchantDB(ctx, a.db).Where("record_id=?", recordID).Update("status", status)
	if err := result.Error; err != nil {
		span.Errorf(err.Error())
		return errors.New("更新状态发生错误")
	}
	return nil
}
