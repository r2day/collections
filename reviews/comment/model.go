package comment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	// affiliate_ 会员营销
	collectionNamePrefix = "reviews_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "comment"
)

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

	// 客户id
	CustomerID string `json:"customer_id" bson:"customer_id"`
	// 产品id
	ProductID string `json:"product_id" bson:"product_id"`
	// 评星
	Rating string `json:"rating" bson:"rating"`
	// 评论内容
	Content string `json:"content" bson:"content"`
	// 评论状态
	ContentStatus string `json:"content_status" bson:"content_status"`
	// 图片
	Pictures []string `json:"pictures" bson:"pictures"`
}
