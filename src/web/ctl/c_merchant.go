package ctl

import (
	"github.com/goodcorn/src/bll"
	"github.com/goodcorn/src/errors"
	"github.com/goodcorn/src/schema"
	"github.com/goodcorn/src/util"
	"github.com/goodcorn/src/web/context"
)

// Merchant demo
// @Name Merchant
// @Description demo接口
type Merchant struct {
	MerchantBll *bll.Merchant `inject:""`
}

// Query 查询数据
func (a *Merchant) Query(ctx *context.Context) {
	switch ctx.Query("q") {
	case "page":
		a.QueryPage(ctx)
	default:
		ctx.ResError(errors.NewBadRequestError("未知的查询类型"))
	}
}

// QueryPage 查询分页数据
// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Param code query string false "编号"
// @Param name query string false "名称"
// @Param status query int false "状态(1:启用 2:停用)"
// @Success 200 []schema.Merchant "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/demos?q=page
func (a *Merchant) QueryPage(ctx *context.Context) {
	var params schema.MerchantQueryParam
	//params.LikeCode = ctx.Query("code")
	params.Status = util.S(ctx.Query("status")).Int()

	items, pr, err := a.MerchantBll.QueryPage(ctx.GetContext(), params, ctx.GetPaginationParam())
	if err != nil {
		ctx.ResError(err)
		return
	}

	ctx.ResPage(items, pr)
}

// Get 查询指定数据
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.Merchant
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/demos/{id}
func (a *Merchant) Get(ctx *context.Context) {
	item, err := a.MerchantBll.Get(ctx.GetContext(), ctx.Param("id"))
	if err != nil {
		ctx.ResError(err)
		return
	}
	ctx.ResSuccess(item)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.Merchant true
// @Success 200 schema.Merchant
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/demos
func (a *Merchant) Create(ctx *context.Context) {
	var item schema.Merchant
	if err := ctx.ParseJSON(&item); err != nil {
		ctx.ResError(err)
		return
	}

	nitem, err := a.MerchantBll.Create(ctx.GetContext(), item)
	if err != nil {
		ctx.ResError(err)
		return
	}
	ctx.ResSuccess(nitem)
}

// Update 更新数据
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Param body body schema.Merchant true
// @Success 200 schema.Merchant
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/demos/{id}
func (a *Merchant) Update(ctx *context.Context) {
	var item schema.Merchant
	if err := ctx.ParseJSON(&item); err != nil {
		ctx.ResError(err)
		return
	}

	nitem, err := a.MerchantBll.Update(ctx.GetContext(), ctx.Param("id"), item)
	if err != nil {
		ctx.ResError(err)
		return
	}
	ctx.ResSuccess(nitem)
}

// Delete 删除数据
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router DELETE /api/v1/demos/{id}
func (a *Merchant) Delete(ctx *context.Context) {
	err := a.MerchantBll.Delete(ctx.GetContext(), ctx.Param("id"))
	if err != nil {
		ctx.ResError(err)
		return
	}
	ctx.ResOK()
}

// Enable 启用数据
// @Summary 启用数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PATCH /api/v1/demos/{id}/enable
func (a *Merchant) Enable(ctx *context.Context) {
	err := a.MerchantBll.UpdateStatus(ctx.GetContext(), ctx.Param("id"), 1)
	if err != nil {
		ctx.ResError(err)
		return
	}
	ctx.ResOK()
}

// Disable 禁用数据
// @Summary 禁用数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PATCH /api/v1/demos/{id}/disable
func (a *Merchant) Disable(ctx *context.Context) {
	err := a.MerchantBll.UpdateStatus(ctx.GetContext(), ctx.Param("id"), 2)
	if err != nil {
		ctx.ResError(err)
		return
	}
	ctx.ResOK()
}
