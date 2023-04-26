package refundflow

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	// affiliate_ 会员营销
	collectionNamePrefix = "trade_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_flow"
	// 这个需要用户根据具体业务完成设定
	modelName = "refund"
)

// Model 模型
// 退款时间	交易通道	原交易号	原订单号	账务主体	原交易金额(元)	退款金额(元)	交易来源	交易类型	交易子类型	退款状态	备注	店铺组织编码
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
	// 退款时间
	RefundTime string `json:"refund_time" bson:"refund_time"`
	// 交易通道
	TradeChannel string `json:"trade_channel" bson:"trade_channel"`
	// 原交易号
	OriginTradeID string `json:"origin_trade_id" bson:"origin_trade_id"`
	// 原订单号
	OriginOrderID string `json:"origin_order_id" bson:"origin_order_id"`
	// 账务主体
	FinancialEntity string `json:"financial_entity" bson:"financial_entity"`
	// 原交易金额(元)
	OriginAmount float64 `json:"origin_amount" bson:"origin_amount"`
	// 原交易金额(元)
	RefundAmount float64 `json:"refund_amount" bson:"refund_amount"`
	// 交易来源
	TradeFrom string `json:"trade_from" bson:"trade_from"`
	// 交易类型
	TradeCategory string `json:"trade_category" bson:"trade_category"`
	// 交易子类型
	TradeSubCategory string `json:"trade_sub_category" bson:"trade_sub_category"`
	// 退款状态
	RefundTradeStatus string `json:"refund_trade_status" bson:"refund_trade_status"`
	// 备注
	Remark string `json:"remark" bson:"remark"`
	// 店铺组织编码
	StoreOrgID string `json:"store_org_id" bson:"store_org_id"`
}
