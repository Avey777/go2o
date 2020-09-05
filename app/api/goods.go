package api

import (
	"github.com/ixre/gof/api"
	"go2o/core/service/thrift/rsi"
)

/**
 * Copyright 2009-2019 @ to2.net
 * name : goods.go
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2019-12-06 10:13
 * description :
 * history :
 */

var _ api.Handler = new(goodsApi)

type goodsApi struct {
	utils
}

func NewGoodsApi() *goodsApi {
	return &goodsApi{}
}

func (g goodsApi) Process(fn string, ctx api.Context) *api.Response {
	return api.HandleMultiFunc(fn, ctx, map[string]api.HandlerFunc{
		"get_new_goods":       g.newGoods,
		"get_hot_sales_goods": g.hotSalesGoods,
		"get_label_goods":     g.saleLabelGoods,
		"favorite":            g.Favorite,
	})
}

/**
 * @api {post} /goods/get_new_goods 获取指定数量的商品
 * @apiGroup goods
 * @apiParam {Int} shop_id 店铺编号, 如果不指定店铺则为所有店铺
 * @apiParam {Int} begin 开始记录数,默认为:0
 * @apiParam {Int} size 数量
 * @apiSuccessExample Success-Response
 * {}
 * @apiSuccessExample Error-Response
 * {"err_code":1,"err_msg":"access denied"}
 */
func (g goodsApi) newGoods(ctx api.Context) interface{} {
	shopId := ctx.Form().GetInt("shop_id")
	begin := ctx.Form().GetInt("begin")
	size := ctx.Form().GetInt("size")
	if size <= 0 {
		size = 10
	}
	_, ss := rsi.ItemService.GetPagedOnShelvesGoods__(int32(shopId), -1,
		begin, begin+size, "it.id DESC")
	return ss
}

/**
 * @api {post} /goods/get_hot_sales_goods 获取指定数量的热销商品
 * @apiGroup goods
 * @apiParam {Int} shop_id 店铺编号, 如果不指定店铺则为所有店铺
 * @apiParam {Int} begin 开始记录数,默认为:0
 * @apiParam {Int} size 数量
 * @apiSuccessExample Success-Response
 * {}
 * @apiSuccessExample Error-Response
 * {"err_code":1,"err_msg":"access denied"}
 */
func (g goodsApi) hotSalesGoods(ctx api.Context) interface{} {
	shopId := ctx.Form().GetInt("shop_id")
	begin := ctx.Form().GetInt("begin")
	size := ctx.Form().GetInt("size")
	if size <= 0 {
		size = 10
	}
	_, ss := rsi.ItemService.GetPagedOnShelvesGoods__(int32(shopId), -1,
		begin, begin+size, "it.sale_num DESC")
	return ss
}

/**
 * @api {post} /goods/get_label_goods 获取指定标签和数量的商品
 * @apiGroup goods
 * @apiParam {string} label_code 销售标签代码
 * @apiParam {int} begin 开始记录数,默认为:0
 * @apiParam {int} size 数量
 * @apiSuccessExample Success-Response
 * {}
 * @apiSuccessExample Error-Response
 * {"err_code":1,"err_msg":"access denied"}
 */
func (g goodsApi) saleLabelGoods(ctx api.Context) interface{} {
	code := ctx.Form().GetString("label_code")
	begin := ctx.Form().GetInt("begin")
	size := ctx.Form().GetInt("size")
	if size <= 0 {
		size = 10
	}
	list := rsi.ItemService.GetValueGoodsBySaleLabel(
		code, "", begin, begin+size)
	return list
}

/**
 * @api {post} /goods/favorite 收藏商品
 * @apiGroup goods
 * @apiParam {int} item_id 商品编号
 * @apiSuccessExample Success-Response
 * {}
 * @apiSuccessExample Error-Response
 * {"err_code":1,"err_msg":"access denied"}
 */
func (g goodsApi) Favorite(ctx api.Context) interface{} {
	memberId := getMemberId(ctx)
	id := ctx.Form().GetInt("item_id")
	err := rsi.MemberService.FavoriteGoods(int64(memberId), int32(id))
	if err != nil {
		return api.ResponseWithCode(1, err.Error())
	}
	return api.NewResponse(nil)
}

//
//func (j *JsonC) Mch_goods(c *echox.Context) error {
//	typeParams := strings.TrimSpace(c.FormValue("params"))
//	types := strings.Split(typeParams, "|")
//	mchId, _ := util.I32Err(strconv.Atoi(c.FormValue("mch_id")))
//	result := make(map[string]interface{}, len(types))
//	key := fmt.Sprint("go2o:repo:sg:front:%d_%s", mchId, typeParams)
//	sto := c.App.Storage()
//	if err := sto.Get(key, &result); err != nil {
//		//从缓存中读取
//		ss := rsi.ItemService
//		for _, t := range types {
//			p, size, begin := j.getMultiParams(t)
//			switch p {
//			case "new-goods":
//				_, result[p] = ss.GetShopPagedOnShelvesGoods(mchId,
//					-1, begin, begin+size, "it.id DESC")
//			case "hot-sales":
//				_, result[p] = ss.GetShopPagedOnShelvesGoods(mchId,
//					-1, begin, begin+size, "it.sale_num DESC")
//			}
//		}
//		sto.SetExpire(key, result, maxSeconds)
//	}
//	return c.Debug(c.JSON(http.StatusOK, result))
//}

//
//func (j *JsonC) Get_shop(c *echox.Context) error {
//	typeParams := strings.TrimSpace(c.FormValue("params"))
//	types := strings.Split(typeParams, "|")
//	result := make(map[string]interface{}, len(types))
//	key := fmt.Sprint("go2o:repo:shop:front:glob_%s", typeParams)
//	sto := c.App.Storage()
//	//从缓存中读取
//	if err := sto.Get(key, &result); err != nil {
//		ss := rsi.ShopService
//		for _, t := range types {
//			p, size, begin := j.getMultiParams(t)
//			switch p {
//			case "new-shop":
//				_, result[p] = ss.PagedOnBusinessOnlineShops(
//					begin, begin+size, "", "sp.create_time DESC")
//			case "hot-shop":
//				_, result[p] = ss.PagedOnBusinessOnlineShops(
//					begin, begin+size, "", "")
//			}
//		}
//		sto.SetExpire(key, result, maxSeconds)
//	}
//	return c.Debug(c.JSON(http.StatusOK, result))
//}

//
//// 最新商品
//func (j *JsonC) Get_NewShop2(c *echox.Context) error {
//	begin, _ := strconv.Atoi(c.FormValue("begin"))
//	size, _ := strconv.Atoi(c.FormValue("size"))
//	ss := rsi.ShopService
//	_, result := ss.PagedOnBusinessOnlineShops(
//		begin, begin+size, "", "sp.create_time DESC")
//
//	return c.Debug(c.JSON(http.StatusOK, result))
//}
