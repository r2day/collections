package collections

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	rtime "github.com/r2day/base/time"
	"github.com/r2day/db"
	"github.com/r2day/rest"

	log "github.com/sirupsen/logrus"
)

const (
	// 管理员角色
	ManagerRoleCollection = "sys_manage_role"
)

// ApiInfo 接口信息
type ApiInfo struct {
	// 路径
	Path string `json:"path" bson:"path"`
	// 名称
	Name string `json:"name" bson:"name"`
}

// RoleModel 角色模型
type RoleModel struct {
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	// 账号状态
	Status bool `json:"status"  bson:"status"`

	// 商户号
	MerchantId string `json:"merchant_id"  bson:"merchant_id"`
	// 更多信息
	// 角色名称
	Name string `json:"name"`
	// AccessApi 可访问的api列表
	AccessApi []ApiInfo `json:"access_api"  bson:"access_api"`
}

// 定义类型名称（别名）
// 新接口只需要修改这里即可
// 以下代码可以复用
type UniversalModel = RoleModel

// ResourceName 返回资源名称
func (m *UniversalModel) ResourceName() string {
	return "roles"
}

// SimpleSave 快速保存
func (m *UniversalModel) SimpleSave(ctx context.Context) (string, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)

	// 保存时间设定
	m.CreatedAt = rtime.FomratTimeAsReader(time.Now().Unix())
	m.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	// 插入记录
	result, err := coll.InsertOne(ctx, m)
	if err != nil {
		return "", err
	}
	stringObjectID := result.InsertedID.(primitive.ObjectID).Hex()
	return stringObjectID, nil
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

func (m *UniversalModel) List(ctx context.Context, merchantId string, urlParams *rest.UrlParams) ([]*UniversalModel, int64, error) {
	coll := db.MDB.Collection(ManagerRoleCollection)
	// 绑定查询结果
	results := make([]*UniversalModel, 0)
	// 声明数据库过滤器
	filters := bson.D{{Key: "merchant_id", Value: merchantId}}
	for key, val := range urlParams.FilterMap {
		// 判断是否是通过id查询
		if m.ResourceName() == key {
			// string to array
			// idSlice := make([]string, 0)
			// err := json.Unmarshal([]byte(val), &idSlice)

			results, err := m.GetMany(ctx, val)
			if err != nil {
				log.Error(err)
				return nil, 0, err
			}
			// results = append(results, ...result)
			return results, int64(len(results)), nil
		} else {
			bm := bson.E{Key: key, Value: val}
			filters = append(filters, bm)
		}
	}
	// filter := bson.D{{Key: "merchant_id", Value: merchantId}}
	log.WithField("filters", filters).Info("############")
	// 获取总数（含过滤规则）
	totalCounter, err := coll.CountDocuments(context.TODO(), filters)
	if err == mongo.ErrNoDocuments {
		return nil, 0, err
	}
	if err != nil {
		return nil, 0, err
	}

	// 进行必要分页处理
	opt := options.Find()
	// 排序方式
	// opt.SetSort(bson.M{urlParams.Sort.Key: urlParams.Sort.SortType})

	opt.SetSkip(int64(urlParams.Range.Offset))
	opt.SetLimit(int64(urlParams.Range.Limit))

	// 获取数据列表
	cursor, err := coll.Find(ctx, filters, opt)
	if err == mongo.ErrNoDocuments {
		return nil, totalCounter, err
	}

	if err != nil {
		return nil, totalCounter, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, totalCounter, err
	}
	return results, totalCounter, nil

}

// Detail 详情
func (m *UniversalModel) Detail(ctx context.Context, id string) (*UniversalModel, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)
	// 绑定查询结果
	result := &UniversalModel{}
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.WithField("id", id).Error(err)
		return nil, err
	}
	return result, nil
}

// GetMany 获取条件查询的结果
func (m *UniversalModel) GetMany(ctx context.Context, ids []string) ([]*UniversalModel, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)
	// 绑定查询结果
	results := make([]*UniversalModel, 0)
	objIds := make([]*primitive.ObjectID, 0)

	for _, i := range ids {
		objId, _ := primitive.ObjectIDFromHex(i)
		objIds = append(objIds, &objId)
	}
	// objId, _ := primitive.ObjectIDFromHex(id)
	// filter := bson.D{{Key: "_id", Value: objId}}

	// err := coll.FindOne(context.TODO(), filter).Decode(&result)
	cursor, err := coll.Find(ctx, bson.M{"_id": bson.M{"$in": objIds}})

	if err != nil {
		log.WithField("ids", ids).Error(err)
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// Update 更新
func (m *UniversalModel) Update(ctx context.Context, id string) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}
	m.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	result, err := coll.UpdateOne(ctx, filter,
		bson.D{{Key: "$set", Value: m}})
	if err != nil {
		log.WithField("id", id).Error(err)
		return err
	}

	if result.MatchedCount < 1 {
		return errors.New("no matched record")
	}

	return nil
}
