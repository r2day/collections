package collections

import (
	"context"

	"github.com/r2day/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// 管理员账号表-名称
	ManagerAccountCollection = "sys_manage_account"
)

// ManagerAccountModel 管理员账号
type ManagerAccountModel struct {
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	// 是否开启审核
	IsRequiredApprove bool `json:"is_required_approve" bson:"is_required_approve"`
	// 账号状态
	Status int `json:"status"  bson:"status"`

	// 关键信息
	// 手机号
	Phone string `json:"phone"`
	// 密码
	Password string `json:"password"  bson:"password"`

	// 更多信息
	// 账号名称
	Name string `json:"name"`
	// 账号Id (自动生成)
	AccountId string `json:"account_id"  bson:"account_id"`
	// 邮箱
	Email string `json:"email"  bson:"email"`
}

// FastSignUp 快速注册
func (m *ManagerAccountModel) SimpleSave(ctx context.Context) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerAccountCollection)

	// 插入记录
	_, err := coll.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

// UpdateById 通过id更新数据库
func (m *ManagerAccountModel) UpdateById(ctx context.Context) error {
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 更新数据库
	filter := bson.D{{Key: "_id", Value: m.ID}}
	_, err := coll.UpdateOne(ctx, filter,
		bson.D{{Key: "$set", Value: m}})
	if err != nil {
		return err
	}
	return nil
}

// FindByPhone 通过手机号查找到账号信息
func (m *ManagerAccountModel) FindByPhone(ctx context.Context) error {
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 更新数据库
	filter := bson.D{{Key: "phone", Value: m.Phone}}
	err := coll.FindOne(ctx, filter).Decode(m)
	if err != nil {
		return err
	}
	return nil
}
