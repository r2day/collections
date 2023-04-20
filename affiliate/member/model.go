package member

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
	modelName = "account"
)

// 每一个应用表示一个大的模块，通常其子模块是一个个接口
// 是有系统默认设定，用户无需修改
// 用户只需要在创建角色的时候选择好需要的应用即可
// 用户选择所需要的应用后->完成角色创建->系统自动拷贝应用具体信息到角色下
// 此时用户可以针对当前的角色中具体的项再自行选择是否移除部分接口，从而进行更精细的权限管理

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

	// 客户编号
	CustomerID string `json:"customer_id" bson:"customer_id"`
	// 姓名
	Name string `json:"name" bson:"name"`
	// 性别
	Gender string `json:"gender" bson:"gender"`
	// 手机号
	Phone string `json:"phone" bson:"phone"`
	// 生日类型
	BirthType string `json:"birth_type" bson:"birth_type"`
	// 生日
	BirthDay string `json:"birth_day" bson:"birth_day"`
	// 来源方式
	From string `json:"from" bson:"from"`
	// 注册时间
	RegisterDate string `json:"register_date" bson:"register_date"`
	// 优惠券
	// 可用的数量
	Coupon int `json:"coupon" bson:"coupon"`
	// 手机号验证
	// 会员迁移后需要进行短信验证完成数据与账号的绑定
	Verify bool `json:"verify"`
}
