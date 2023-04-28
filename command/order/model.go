package order

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	// affiliate_ 会员营销
	collectionNamePrefix = "command_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "order"
)

// Customer 顾客信息
// 顾客姓名	顾客电话	顾客地址
type Customer struct {
	// 顾客姓名
	Name string `json:"name" bson:"name"`
	// 顾客电话
	Phone string `json:"phone" bson:"phone"`
	// 顾客地址
	Address string `json:"address" bson:"address"`
}

// Amounter 金额
// 订单总额	已支付金额	菜品总额	会员卡支付	会员卡积分抵扣金额	已退订/已退款
type Amounter struct {
	// 订单总额
	Amount float64 `json:"amount" bson:"amount"`
	// 已经支付金额
	Paid float64 `json:"paid" bson:"paid"`
	// 菜品总额
	Total float64 `json:"total" bson:"total"`
	// 会员卡支付
	VIP float64 `json:"vip" bson:"vip"`
	// 会员卡积分抵扣金额
	Deduction float64 `json:"deduction" bson:"deduction"`
	// 已退订/已退款
	Refund float64 `json:"refund" bson:"refund"`
}

// Model 模型
// 序号	城市	店铺名称	订单号	账单号	三方单号	支付流水号	顾客姓名	顾客电话	顾客地址
// 订单时间	订单状态	订单类型	渠道	支付通道	支付方式(第三方支付方式)	订单总额
// 已支付金额	菜品总额	会员卡支付	会员卡积分抵扣金额	已退订/已退款	账单备注
type Model struct {
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 商户号
	MerchantID string `json:"merchant_id" bson:"merchant_id"`
	// 创建者
	AccountID string `json:"account_id" bson:"account_id"`
	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	// 状态
	Status bool `json:"status"`
	// 根据角色的最低级别写入
	AccessLevel uint `json:"access_level" bson:"access_level"`

	// 序号
	SerialNumber uint `json:"serial_number" bson:"serial_number"`
	// 城市
	City string `json:"city" bson:"city"`
	// 店铺名称
	StoreName string `json:"store_name" bson:"store_name"`
	// 订单号
	OrderID string `json:"order_id" bson:"order_id"`
	// 账单号
	BillID string `json:"bill_id" bson:"bill_id"`
	// 三方单号
	ExternalID string `json:"external_id" bson:"external_id"`
	// 支付流水号
	FlowID string `json:"flow_id" bson:"flow_id"`

	// 顾客信息
	CustomerInfo Customer `json:"customer_info" bson:"customer_info"`
	// 支付信息
	AmountInfo Amounter `json:"amount_info" bson:"amount_info"`

	// 订单时间
	OrderTime string `json:"order_time" bson:"order_time"`
	// 订单状态
	OrderStatus string `json:"order_status" bson:"order_status"`
	// 订单类型
	OrderCategory string `json:"order_category" bson:"order_category"`
	// 渠道
	Channel string `json:"channel" bson:"channel"`
	// 支付通道
	PayChannel string `json:"pay_channel" bson:"pay_channel"`
	// 支付方式(第三方支付方式)
	ExternalPayMethod string `json:"external_pay_method" bson:"external_pay_method"`
	// 备注
	Remark string `json:"remark" bson:"remark"`
}
