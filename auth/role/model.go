package role

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "sys_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "role"
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
	// 名称
	Name string `json:"name" bson:"name"`
	// 描述
	Desc string `json:"desc" bson:"desc"`
	// 图片
	Image string `json:"image" bson:"image"`
	// 应用列表
	// 存储应用的id
	// 通过应用id 快速获得应用列表
	Apps []string `json:"apps" bson:"apps"`
	// AccessApi 可访问的api列表
	// 即将废弃
	// AccessAPI []collections.APIInfo `json:"access_api"  bson:"access_api"`
	// 根据角色的最低级别写入
	AccessLevel uint `json:"access_level" bson:"access_level"`

	// 图片
	Permissions PermissionsModel `json:"permissions" bson:"permissions"`
}

// PermissionsModel 模型
// 记录角色对接口的操作细节
type PermissionsModel struct {
	// 商户号
	RoleID string `json:"role_id" bson:"role_id"`
	// 商户号
	AppID string `json:"app_id" bson:"app_id"`
	// 读
	Read bool `json:"read" bson:"read"`
	// 写
	Write bool `json:"write" bson:"write"`
	// 改
	Update bool `json:"update" bson:"update"`
	// 详情
	Detail bool `json:"detail" bson:"detail"`
	// 删除
	Delete bool `json:"delete" bson:"delete"`
}
