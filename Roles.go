package collections

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	collections "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	rtime "github.com/r2day/base/time"
	"github.com/r2day/db"
	"github.com/r2day/rest"

	log "github.com/sirupsen/logrus"
)

const (
	// ManagerRoleCollection 角色数据表
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
	// 状态
	Status bool `json:"status"  bson:"status"`

	// 商户号
	MerchantID string `json:"merchant_id"  bson:"merchant_id"`
	// 更多信息
	// 角色名称
	Name string `json:"name"`
	// AccessApi 可访问的api列表
	AccessAPI []collections.APIInfo `json:"access_api"  bson:"access_api"`
}

// UniversalModel 定义类型名称（别名）
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
		log.WithField("m", m).Error(err)
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
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}

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

// List 获取列表
func (m *UniversalModel) List(ctx context.Context, merchantID string, urlParams *rest.UrlParams) ([]*UniversalModel, int64, error) {
	coll := db.MDB.Collection(ManagerRoleCollection)
	// 声明需要返回的列表
	results := make([]*UniversalModel, 0)
	// 声明日志基本信息
	logCtx := log.WithField("merchantID", merchantID).WithField("urlParams.FilterMap", urlParams.FilterMap)
	// 声明数据库过滤器
	// 定义基本过滤规则
	// 以商户id为基本命名空间
	filters := bson.D{{Key: "merchant_id", Value: merchantID}}
	// 添加更多过滤器
	// 根据用户规则进行筛选
	for key, val := range urlParams.FilterMap {
		// 判断是否是通过id查询
		// 则进行转换
		// 一般对应于 ReferenceArrayInput 和 ReferenceManyField
		if m.ResourceName() == key || key == "id" {
			// string to array
			results, err := m.GetManyInIds(ctx, val)
			if err != nil {
				logCtx.Error(err)
				return nil, 0, err
			}
			logCtx.WithField("results", results).Warning("is reference request")
			return results, int64(len(results)), nil
		} else {
			bm := bson.E{Key: key, Value: val}
			filters = append(filters, bm)
		}
	}

	// 添加状态过滤器
	if urlParams.HasFilter {
		filterByStatus := bson.E{Key: "status", Value: urlParams.FilterCommon.Status}
		filters = append(filters, filterByStatus)
	}

	logCtx.WithField("filters", filters).Info("final filters has been combine")
	// 获取总数（含过滤规则）
	totalCounter, err := coll.CountDocuments(context.TODO(), filters)
	if err == mongo.ErrNoDocuments {
		logCtx.Error(err)
		return nil, 0, err
	}
	if err != nil {
		logCtx.Error(err)
		return nil, 0, err
	}

	// 进行必要分页处理
	opt := options.Find()
	// 排序方式
	if urlParams.Sort.SortType == rest.AES {
		opt.SetSort(bson.M{urlParams.Sort.Key: -1})
	} else {
		opt.SetSort(bson.M{urlParams.Sort.Key: 1})
	}

	opt.SetSkip(int64(urlParams.Range.Offset))
	opt.SetLimit(int64(urlParams.Range.Limit))

	// 获取数据列表
	cursor, err := coll.Find(ctx, filters, opt)
	if err == mongo.ErrNoDocuments {
		logCtx.Error(err)
		return nil, totalCounter, err
	}

	if err != nil {
		logCtx.Error(err)
		return nil, totalCounter, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		logCtx.Error(err)
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
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	logCtx := log.WithField("filter", filter)
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		logCtx.Error(err)
		return nil, err
	}
	return result, nil
}

// GetManyInIds 获取条件查询的结果
func (m *UniversalModel) GetManyInIds(ctx context.Context, ids []string) ([]*UniversalModel, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)
	// 绑定查询结果
	results := make([]*UniversalModel, 0)
	objIds := make([]*primitive.ObjectID, 0)
	logCtx := log.WithField("ids", ids)

	for _, i := range ids {
		objID, _ := primitive.ObjectIDFromHex(i)
		objIds = append(objIds, &objID)
	}
	cursor, err := coll.Find(ctx, bson.M{"_id": bson.M{"$in": objIds}})

	if err != nil {
		logCtx.Error(err)
		return nil, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		logCtx.Error(err)
		return nil, err
	}
	return results, nil
}

// Update 更新
func (m *UniversalModel) Update(ctx context.Context, id string) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerRoleCollection)
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
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
