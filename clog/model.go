package clog

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "sys_"
	// CollectionNameSubffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSubffix = "_log"
	// 这个需要用户根据具体业务完成设定
	modelName = "signin"
)

// Model 模型
type Model struct {
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 商户号
	MerchantID string `json:"-" bson:"merchant_id"`
	// 创建者
	AccountID string `json:"account_id" bson:"account_id"`
	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	// 状态
	Status bool `json:"status"`

	// 用户根据业务需求定义的字段
	// 客户IP
	ClientIP string `json:"client_ip" bson:"client_ip"`
	// 远程IP
	RemoteIP string `json:"remote_ip"  bson:"remote_ip"`
	// 路径
	FullPath string `json:"full_path"  bson:"full_path"`
	// 请求方法/操作
	Method string `json:"method"  bson:"method"`
	// 相应代码
	RespCode int `json:"resp_code"  bson:"resp_code"`
	// 操作对象id
	TargetID string `json:"target_id"  bson:"target_id"`
}
