package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	// affiliate_ 会员营销
	collectionNamePrefix = "dishes_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "items"
)

// *菜品编码	*菜品名称	*POS分类	*线上分类	收入科目	出品部门
// *规格名称 *常规售价	常规会员价	外卖价格	外卖会员价	*规格名称	*常规售价	常规会员价	外卖价格	外卖会员价	*规格名称	*常规售价	常规会员价	外卖价格	外卖会员价
// 菜品启用	菜品上架	支持业务	是否招牌菜	是否新菜	是否推荐菜	参与打折	菜品数量支持小数	称重菜	参与开台自动加入	菜品可单独销售	菜品起售数量	销售提点	打包费	第二出品部门	税率	菜品标签

// BasicInformation 菜品基本信息
// 887633463001	二两螺蛳粉	粉类	粉类	商品销售收入	厨房
type BasicInformation struct {
	// DishesID *菜品编码
	DishesID string `json:"dishes_id"  bson:"dishes_id"`
	// *菜品名称
	Name string `json:"name" bson:"name"`
	// *POS分类
	PosCategory string `json:"pos_category" bson:"pos_category"`
	// *线上分类
	OnlineCategory string `json:"online_category" bson:"online_category"`
	// 收入科目
	Income string `json:"income" bson:"income"`
	// 出品部门
	ProductionDepartment string `json:"production_department" bson:"production_department"`
}

// SpecificationPrice 规格
// *规格名称 *常规售价	常规会员价	外卖价格	外卖会员价
// 生菜(底菜)	12.00	12.00
// 豆芽(底菜)	12.00	12.00
// 火筒菜(底菜)	12.00	12.00
type SpecificationPrice struct {
	// *规格名称
	Name string `json:"name" bson:"name"`
	// *常规售价
	Normal float64 `json:"normal" bson:"normal"`
	// 常规会员价
	NormalVIP float64 `json:"normal_vip" bson:"normal_vip"`
	// 外卖价格
	TakeOut float64 `json:"take_out" bson:"take_out"`
	// 外卖会员价
	TakeOutVIP float64 `json:"take_out_vip" bson:"take_out_vip"`
}

// EnablesSwitch 开关
// 菜品启用	菜品上架	支持业务	是否招牌菜	是否新菜	是否推荐菜	参与打折	菜品数量支持小数	称重菜	参与开台自动加入	菜品可单独销售
// 启用	是	堂食+外卖+自提	否	否	否	是	否	否	否	是
type EnablesSwitch struct {
	// Open 菜品启用
	IsOpen bool `json:"is_open"  bson:"is_open"`
	// 菜品上架
	IsOnShelves bool `json:"is_on_shelves" bson:"is_on_shelves"`
	// 是否招牌菜
	IsSupportSigns bool `json:"is_support_signs" bson:"is_support_signs"`
	// 是否新菜
	IsSupportNew bool `json:"is_support_new" bson:"is_support_new"`
	// 是否推荐菜
	IsSupportRecommend bool `json:"is_support_recommend" bson:"is_support_recommend"`
	// 是否参与打折
	IsSupportDiscount bool `json:"is_support_discount" bson:"is_support_discount"`
	// 是否新菜
	IsSupportDecimal bool `json:"is_support_decimal" bson:"is_support_decimal"`
	// 是否称重菜
	IsSupportWeighing bool `json:"is_support_weighing" bson:"is_support_weighing"`
	// 参与开台自动加入
	IsSupportAutoJoin bool `json:"is_support_auto_join" bson:"is_support_auto_join"`
	// 是否菜品可单独销售
	IsSupportSoldSeparately bool `json:"is_support_sold_separately" bson:"is_support_sold_separately"`
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

	// 基本信息
	BasicInfo BasicInformation `json:"basic_info" bson:"basic_info"`
	// 开启状态
	Enables EnablesSwitch `json:"enables" bson:"enables"`
	// 收入科目
	Income string `json:"income" bson:"income"`
	// 出品部门
	ProductionDepartment string `json:"production_department" bson:"production_department"`
	// 规格列表
	Specification []SpecificationPrice `json:"specification_list" bson:"specification_list"`
	// 支持业务
	SupportBusinessList []string `json:"support_business_list" bson:"support_business_list"`
}
