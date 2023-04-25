package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	// affiliate_ 会员营销
	collectionNamePrefix = "scm_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "supplier"
)

// Model 模型
// 货主名称	供货商编码(必填）	供货商名称(必填)	供应商助记码	供货商类别编码(必填)	供货商类别名称(必填)	供货商联系人	电话(必填)	邮箱	地址	状态
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
	// 货主名称
	CargoOwner string `json:"cargo_owner" bson:"cargo_owner"`
	// 供货商编码(必填）
	SupplierID string `json:"supplier_id" bson:"supplier_id"`
	// 供货商名称(必填)
	SupplierName string `json:"supplier_name" bson:"supplier_name"`
	// 供应商助记码
	Mnemonic string `json:"mnemonic" bson:"mnemonic"`
	// 供货商类别名称(必填)
	SupplierCategory string `json:"supplier_category" bson:"supplier_category"`
	// 供货商联系人
	SupplierContact string `json:"supplier_contact" bson:"supplier_contact"`
	// 电话(必填)
	Phone string `json:"phone" bson:"phone"`
	// 邮箱
	Email string `json:"email" bson:"email"`
	// 地址
	Address string `json:"address" bson:"address"`
}
