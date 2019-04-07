package entity

import (
	"context"

	"github.com/goodcorn/src/schema"
	"github.com/goodcorn/src/service/gormplus"
)

// GetMerchantDB 获取merchant存储
func GetMerchantDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModel(ctx, defDB, Merchant{})
}

// SchemaMerchant merchant对象
type SchemaMerchant schema.Merchant

// ToMerchant 转换为merchant实体
func (a SchemaMerchant) ToMerchant() *Merchant {
	//todo 要添加经纬度
	item := &Merchant{
		RecordID:           a.RecordID,
		Code:               a.Code,
		Name:               a.Name,
		Description:        a.Description,
		AddressDescription: a.AddressDescription,
		Address:            a.Address,
		Status:             a.Status,
		Creator:            a.Creator,
	}
	return item
}

// Merchant merchant实体
type Merchant struct {
	Model
	RecordID           string `gorm:"column:record_id;size:36;index;"`      // 记录内码
	Code               string `gorm:"column:code;size:50;index;"`           // 编号
	Name               string `gorm:"column:name;size:100;index;"`          // 名称
	AddressDescription string `gorm:"column:address_description;size:200;"` // 地址描述
	Address            string `gorm:"column:address;size:200;"`             // 地址
	Description        string `gorm:"column:description;size:200;"`         // 备注
	Phone              string `gorm:"column:phone;size:20;"`                // 手机号
	Status             int    `gorm:"column:status;index;"`                 // 状态(1:启用 2:停用)
	Creator            string `gorm:"column:creator;size:36;"`              // 创建者
}

func (a Merchant) String() string {
	return toString(a)
}

// TableName 表名
func (a Merchant) TableName() string {
	return a.Model.TableName("merchant")
}

// ToSchemaMerchant 转换为merchant对象
func (a Merchant) ToSchemaMerchant() *schema.Merchant {
	item := &schema.Merchant{
		RecordID:           a.RecordID,
		Phone:              a.Phone,
		Code:               a.Code,
		Name:               a.Name,
		Status:             a.Status,
		Description:        a.Description,
		Address:            a.Address,
		AddressDescription: a.AddressDescription,
		Creator:            a.Creator,
		CreatedAt:          a.CreatedAt,
	}
	return item
}

// Merchants merchant列表
type Merchants []*Merchant

// ToSchemaMerchants 转换为merchant对象列表
func (a Merchants) ToSchemaMerchants() []*schema.Merchant {
	list := make([]*schema.Merchant, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaMerchant()
	}
	return list
}
