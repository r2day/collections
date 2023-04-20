package card

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
	modelName = "card"
)

// Assets 资产信息
type Assets struct {
	Balance float64 `json:"balance"  bson:"balance"`
	// 现金卡值
	CashCharge float64 `json:"cash_charge" bson:"cash_charge"`
	// 冻结卡值
	Freezing float64 `json:"freezing" bson:"freezing"`
	// 赠送卡值
	Gift float64 `json:"gift" bson:"gift"`
	// 积分余额
	Integral uint64 `json:"integral" bson:"integral"`
	// 累计储值总额
	TotalBalance float64 `json:"total_balance" bson:"total_balance"`
	//累计储值次数
	TotalBalanceCounter uint64 `json:"total_balance_counter" bson:"total_balance_counter"`
	// 累计消费总额
	TotalCumulativeConsumption float64 `json:"total_cumulative_consumption" bson:"total_cumulative_consumption"`
	// 累计消费总额
	TotalCumulativeConsumptionCounter uint64 `json:"total_cumulative_consumption_counter" bson:"total_cumulative_consumption_counter"`
	// 挂帐总额度
	DebitTotalLimit float64 `json:"debit_total_limit" bson:"debit_total_limit"`
	// 挂帐剩余额度
	DebitLeftLimit float64 `json:"debit_left_limit" bson:"debit_left_limit"`

	// 已用额度
	DebitUsedLimit float64 `json:"debit_used_limit" bson:"debit_used_limit"`
}

// 序号	卡号	客户编号	手机号	姓名	性别	生日
// 卡类别	卡状态	卡等级	开卡店铺
// 卡余额	现金卡值	冻结卡值	赠送卡值	积分余额	累计储值总额	累计储值次数	累计消费总额	累计消费次数	挂账额度	剩余额度	已用额度
// 603.9	596.88	0	7.02	0	1700	60	1483.9	39	0	0	0
// 实体卡号	开卡日期	有效期至	最后交易时间

// 每一个应用表示一个大的模块，通常其子模块是一个个接口
// 是有系统默认设定，用户无需修改
// 用户只需要在创建角色的时候选择好需要的应用即可
// 用户选择所需要的应用后->完成角色创建->系统自动拷贝应用具体信息到角色下
// 此时用户可以针对当前的角色中具体的项再自行选择是否移除部分接口，从而进行更精细的权限管理

// UserInformation 用户信息
// 序号	卡号	客户编号	手机号	姓名	性别	生日
type UserInformation struct {
	// 客户编号
	CustomerID string `json:"customer_id" bson:"customer_id"`
	// 姓名
	Name string `json:"name" bson:"name"`
	// 性别
	Gender string `json:"gender" bson:"gender"`
	// 手机号
	Phone string `json:"phone" bson:"phone"`
	// 生日
	BirthDay string `json:"birth_day" bson:"birth_day"`
}

// BasicInformation 卡信息
// 卡类别	卡状态	卡等级	开卡店铺
type BasicInformation struct {
	// Type 卡类别
	Type string `json:"type"  bson:"type"`
	// 卡状态
	CardStatus string `json:"card_status" bson:"card_status"`
	// 等级
	Level string `json:"level" bson:"level"`
	// 开卡店铺
	CardFrom string `json:"card_from" bson:"card_from"`
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

	// 用户信息
	UserInfo UserInformation `json:"user_info" bson:"user_info"`

	// 资产信息
	CardInfo BasicInformation `json:"card_info" bson:"card_info"`

	// 资产信息
	AssetsInfo Assets `json:"assets" bson:"assets"`

	// 实体卡号	开卡日期	有效期至	最后交易时间
	// 来源方式
	From string `json:"from" bson:"from"`
	// 卡号
	Number string `json:"number" bson:"number"`
	// OpeningDate 开卡日期
	// 可用的数量
	OpeningDate string `json:"opening_date" bson:"opening_date"`
	// 手机号验证
	// 会员迁移后需要进行短信验证完成数据与账号的绑定
	Verify bool `json:"verify"`
}
