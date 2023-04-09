package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/r2day/db"
	log "github.com/sirupsen/logrus"
)

const (
	// 管理员角色
	ManagerRoleCollection = "sys_manage_role"
)

// RoleModel 角色模型
type RoleModel struct {
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	// 账号状态
	Status int `json:"status"  bson:"status"`

	// 更多信息
	// 角色名称
	Name string `json:"name"`
	// AccessApi 可访问的api列表
	AccessApi []string `json:"access_api"  bson:"access_api"`
}

// 定义类型名称（别名）
// 新接口只需要修改这里即可
// 以下代码可以复用
type UniversalModel = RoleModel

// SimpleSave 快速保存
func (m *UniversalModel) SimpleSave(ctx context.Context) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)

	// 保存时间设定
	m.CreatedAt = rtime.FomratTimeAsReader(time.Now().Unix())
	m.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	// 插入记录
	_, err := coll.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

// Delete 快速删除
func (m *UniversalModel) Delete(ctx context.Context, id string) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}

	// 执行删除
	result, err := coll.DeleteOne(ctx, filter)

	if err != nil {
		log.WithField("id", id).Error(err)
		return err
	}

	if result.DeletedCount < 1 {
		log.WithField("id", id).Error("delete failed")
		return err
	}

	return nil
}

func (m *UniversalModel) List (ctx context.Context, merchantId string, offset int64, limit int64) ([]*UniversalModel, error) {
	coll := db.MDB.Collection(ManagerRoleCollection)
	// 声明数据库过滤器
	// var filter bson.D
	filter := bson.D{{Key: "merchant_id", Value: merchantId}}
	// 获取总数（含过滤规则）
	totalCounter, err := coll.CountDocuments(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// 进行必要分页处理
	opt := options.Find()
	// 排序方式
	opt.SetSort(bson.M{"name": 1})
	opt.SetSkip(offset)
	opt.SetLimit(limit)

	// 获取数据列表
	cursor, err := coll.Find(ctx, filter, opt)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// 绑定查询结果
	results := make([]*UniversalModel, 0)
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err;
	}

	fmt.Println("--results->", results)
	return results, nil

}