package membership

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	// affiliate_ 会员营销
	collectionNamePrefix = "affiliate_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "membership"
)

// 每一个应用表示一个大的模块，通常其子模块是一个个接口
// 是有系统默认设定，用户无需修改
// 用户只需要在创建角色的时候选择好需要的应用即可
// 用户选择所需要的应用后->完成角色创建->系统自动拷贝应用具体信息到角色下
// 此时用户可以针对当前的角色中具体的项再自行选择是否移除部分接口，从而进行更精细的权限管理

// Settlement 结算模型
type Settlement struct {
	// 会员方案结算主体
	SettlementObject string `json:"settlement_object" bson:"settlement_object"`
	// 会员方案结算主体
	StoredValueSettlement string `json:"stored_value_settlement" bson:"stored_value_settlement"`
}

// FeeSetting 费用设置
type FeeSetting struct {
	// 线下开卡工本费
	OfflineOpenCard float64 `json:"offline_open_card" bson:"offline_open_card"`
	// 线上开卡工本费
	OnlineOpenCard float64 `json:"online_open_card" bson:"online_open_card"`
	// 线下补办收取工本费
	IsRequireFeeSecond bool `json:"is_require_fee_second" bson:"is_require_fee_second"`
	// 开卡押金
	Deposit float64 `json:"deposit" bson:"deposit"`
	// 储值并开卡
	ChargerAndOpen float64 `json:"charger_and_open" bson:"charger_and_open"`
}

// TradeSetting 交易设置
type TradeSetting struct {
	// 卡值消费方式
	ConsumptionPatterns string `json:"consumption_patterns" bson:"consumption_patterns"`
	// 开通线上充值
	IsOnlineCharger bool `json:"is_online_charger" bson:"is_online_charger"`
	// 交易是否限制
	TradeLimit string `json:"trade_limit" bson:"trade_limit"`
	// 储值限额金额 -1 表示不限制
	StoredMaxLimit float64 `json:"stored_max_limit" bson:"stored_max_limit"`
	// 消息推送
	MsgPush string `json:"msg_push" bson:"msg_push"`
	// 交易校验
	Verify bool `json:"verify" bson:"verify"`
	// 会员卡支付适用业务类型
	SupportBusinessCategory []string `json:"support_business_category" bson:"support_business_category"`
	// 卡值消费金额限制 -1 表示不限制
	ConsumptionMaxValueLimit float64 `json:"consumption_max_value_limit" bson:"consumption_max_value_limit"`
	// 卡值消费次数限制 -1 表示不限制
	ConsumptionMaxTimesLimit int `json:"consumption_max_times_limit" bson:"consumption_max_times_limit"`
	// 是否可注销
	IsCanLogOut bool `json:"is_can_log_out" bson:"is_can_log_out"`
}

// ExpireSetting 期限设置
type ExpireSetting struct {
	// 发票有效期 -1 表示永久
	Invoice int `json:"invoice" bson:"invoice"`
	// 会员有效期 -1 表示永久
	Member int `json:"member" bson:"member"`
}

// ParamsSetting 开卡参数设置
type ParamsSetting struct {
	// 生日性别 是否必填
	IsBirthDayRequire bool `json:"is_birth_day_require" bson:"is_birth_day_require"`
	// 身份证 是否必填
	IsIDCardRequire bool `json:"is_id_card_require" bson:"is_id_card_require"`
}

// Model 模型
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

	// 姓名
	Name string `json:"name" bson:"name"`
	// 会员服务电话
	Phone string `json:"phone" bson:"phone"`
	// 会员卡LOGO
	CardLOGO string `json:"card_logo" bson:"card_logo"`
	// 卡背景图
	CardBackgroud string `json:"card_background" bson:"card_background"`
	// 字体颜色
	FrontColor string `json:"front_color" bson:"front_color"`
	// 会员卡背景颜色
	BackgroudColor string `json:"background_color" bson:"background_color"`
	// 会员服务说明
	Desc string `json:"desc" bson:"desc"`
	// 会员注册协议
	Protocol string `json:"protocol"`
	// 结算设置
	SettlementInfo Settlement `json:"settlement_info" bson:"background_color"`
	// 费用设置
	FeeSettingInfo FeeSetting `json:"fee_setting_info" bson:"fee_setting_info"`
	// 交易设置
	TradeSettingInfo TradeSetting `json:"trade_setting_info" bson:"trade_setting_info"`
	// 业务设置
	ExpireSettingInfo ExpireSetting `json:"expire_setting_info" bson:"expire_setting_info"`
	// 开卡参数设置
	ParamsSettingInfo ParamsSetting `json:"params_setting_info" bson:"params_setting_info"`
	// 会员方案适用店铺
	SupportStoreList []string `json:"support_store_list" bson:"support_store_list"`
}
