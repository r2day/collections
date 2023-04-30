package brand

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	// affiliate_ 会员营销
	collectionNamePrefix = "system_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "brand"
)

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

	// 名称
	Name string `json:"name" bson:"name"`
	// 业态ID
	CategoryID string `json:"category_id" bson:"category_id"`
	// 业态类型
	CategoryName string `json:"category_name" bson:"category_name"`
	// 所属品牌
	LogoURL string `json:"logo_url" bson:"logo_url"`
	// 描述
	Desc string `json:"desc" bson:"desc"`
}
