package payflow

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
	modelName = "pay"
)

// Model 模型
// 交易时间	交易通道	交易来源	账务主体	交易信息	交易类型	交易子类型	交易金额	交易状态	交易号	订单号	备注	店铺组织编码
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
	// 交易时间
	TradeTime string `json:"trade_time" bson:"trade_time"`
	// 交易通道
	TradeChannel string `json:"trade_channel" bson:"trade_channel"`
	// 交易来源
	TradeFrom string `json:"trade_from" bson:"trade_from"`
	// 账务主体
	FinancialEntity string `json:"financial_entity" bson:"financial_entity"`
	// 交易信息 Information
	Information string `json:"information" bson:"information"`
	// 交易类型
	TradeCategory string `json:"trade_category" bson:"trade_category"`
	// 交易子类型
	TradeSubCategory string `json:"trade_sub_category" bson:"trade_sub_category"`
	// 交易金额
	Amount string `json:"amount" bson:"amount"`
	// 交易状态
	TradeStatus string `json:"trade_status" bson:"trade_status"`
	// 订单号
	OrderID string `json:"order_id" bson:"order_id"`
	// 备注
	Remark string `json:"remark" bson:"remark"`
	// 店铺组织编码
	StoreOrgID string `json:"store_org_id" bson:"store_org_id"`
}
