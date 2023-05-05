package store

import (
	"github.com/r2day/collections"
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
	modelName = "store"
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

	// 门店名称
	Name string `json:"name" bson:"name"`
	// 门店ID
	StoreID string `json:"store_id" bson:"store_id"`
	// 组织ID
	OrgID string `json:"org_id" bson:"org_id"`
	// 财务主体
	FinancialSubject string `json:"financial_subject" bson:"financial_subject"`
	// 标签
	Tags string `json:"tags" bson:"tags"`
	// 集团ID
	GroupID string `json:"group_id" bson:"group_id"`
	// 类型ID
	CategoryName string `json:"category_name" bson:"category_name"`
	// 所属品牌
	BrandID string `json:"brand_id" bson:"brand_id"`
	// 所属品牌
	BrandName string `json:"brand_name" bson:"brand_name"`
	// 描述
	Desc string `json:"desc" bson:"desc"`
	// 描述
	Phone string `json:"phone" bson:"phone"`
	// 门店公告
	BBS string `json:"bbs" bson:"bbs"`
	// 运营模式
	Mode string `json:"mode" bson:"mode"`
	// 营业时间
	WorkingTime string `json:"working_time" bson:"working_time"`
	// 地址信息
	AddressInfo collections.Address `json:"address_info" bson:"address_info"`
}
