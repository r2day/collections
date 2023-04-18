package crole

import (
	"context"
	"time"

	rtime "github.com/r2day/base/time"
	"github.com/r2day/db"
	"github.com/r2day/rest"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSubffix
}

// Create 创建
// create	POST http://my.api.url/posts
func (m *Model) Create(ctx context.Context) (string, error) {
	coll := db.MDB.Collection(m.CollectionName())

	// 保存时间设定
	m.CreatedAt = rtime.FomratTimeAsReader(time.Now().Unix())
	// 更新时间设定
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

// Delete 删除
// delete	DELETE http://my.api.url/posts/123
func (m *Model) Delete(ctx context.Context, id string) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	logCtx := log.WithField("id", id)
	coll := db.MDB.Collection(m.CollectionName())
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	// 执行删除
	result, err := coll.DeleteOne(ctx, filter)

	if err != nil {
		logCtx.Error(err)
		return err
	}

	if result.DeletedCount < 1 {
		logCtx.Warning("result.DeletedCount < 1")
		return nil
	}
	return nil
}

// GetOne 详情
// getOne	GET http://my.api.url/posts/123
func (m *Model) GetOne(ctx context.Context, id string) (*Model, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(m.CollectionName())
	// 绑定查询结果
	result := &Model{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	logCtx := log.WithField("filter", filter)

	err = coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		logCtx.Error(err)
		return nil, err
	}
	return result, nil
}

// GetMany 获取条件查询的结果
// getMany	GET http://my.api.url/posts?filter={"ids":[123,456,789]}
func (m *Model) GetMany(ctx context.Context, ids []string) ([]*Model, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(m.CollectionName())
	// 绑定查询结果
	results := make([]*Model, 0)
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
// update	PUT http://my.api.url/posts/123
func (m *Model) Update(ctx context.Context, id string) error {
	coll := db.MDB.Collection(m.CollectionName())
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	// 设定更新时间
	m.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())

	result, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: m}})
	if err != nil {
		log.WithField("id", id).Error(err)
		return err
	}

	if result.MatchedCount < 1 {
		log.WithField("id", id).Warning("no matched record")
		return nil
	}

	return nil
}

// GetList 获取列表
// getList	GET http://my.api.url/posts?sort=["title","ASC"]&range=[0, 24]&filter={"title":"bar"}
func (m *Model) GetList(ctx context.Context, merchantID string, accountID string, urlParams *rest.UrlParams) ([]*Model, int64, error) {
	coll := db.MDB.Collection(m.CollectionName())
	// 声明需要返回的列表
	results := make([]*Model, 0)
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
			results, err := m.GetMany(ctx, val)
			if err != nil {
				logCtx.Error(err)
				return nil, 0, err
			}
			logCtx.WithField("results", results).Warning("is reference request")
			return results, int64(len(results)), nil
		}

		// 用户可以指定accountId
		bm := bson.E{Key: key, Value: val}
		filters = append(filters, bm)

	}

	// 添加状态过滤器
	if urlParams.HasFilter {
		filterByStatus := bson.E{Key: "status", Value: urlParams.FilterCommon.Status}
		filters = append(filters, filterByStatus)
	}

	logCtx.WithField("filters", filters).Debug("final filters has been combine")
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
