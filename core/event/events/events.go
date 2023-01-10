package events

import (
	"github.com/ixre/go2o/core/domain/interface/order"
	"github.com/ixre/go2o/core/domain/interface/wallet"
)

// 钱包更新事件
type WalletLogClickhouseUpdateEvent struct {
	Data *wallet.WalletLog
}

// 订单分销事件
type OrderAffiliateRebateEvent struct {
	// 订单号
	OrderNo string
	// 子订单
	SubOrder bool
	// 买家编号
	BuyerId int64
	// 订单金额
	OrderAmount int64
	// 分销商品
	AffiliateItems []*order.SubOrderItem
}

// 发送短信事件
type SendSmsEvent struct {
	// 短信服务商
	Provider int
	// 手机号
	Phone string
	// 短信内容
	Template string
	// 短信模板
	TemplateId string
	// 数据
	Data []string
}
