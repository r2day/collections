package collections

import (
	"context"
	"errors"
	"time"

	rtime "github.com/r2day/base/time"
	"github.com/r2day/db"
	"github.com/r2day/rest"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// ManagerAccountCollection 管理员账号表-名称
	ManagerAccountCollection = "sys_manage_account"
)

const (
	// DeployModeKey 是否为单机模式
	DeployModeKey = "DEPLOY_MODE"
	// StandaloneModeValue 部署模式为单机模式时需要使用该变量
	StandaloneModeValue = "standalone"
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
	// 状态
	Status bool `json:"status"  bson:"status"`
	// 商户id
	// 如果用户部署为单机模式，则商户号为固定值
	// 一般用于部署在私有化的时候启动
	// 在其他模式下，merchantId 起到命名空间的作用
	MerchantId string `json:"merchant_id"  bson:"merchant_id"`
	// 是否是管理员
	// 一般standalone 模式下，是通过读取部署时配置的 ADMIN_PHONE
	// 如果与配置的手机号匹配，那么就可以定义为管理员
	// 如果是其他的模式，一般需要超级商户平台授权后才能成为管理员
	IsAdmin bool `json:"is_admin"  bson:"is_admin"`
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
	// 角色名称列表
	Roles []string `json:"roles"  bson:"roles"`
}

// 定义类型名称（别名）
// 新接口只需要修改这里即可
// 以下代码可以复用
type universalModel = ManagerAccountModel

// ResourceName 返回资源名称
func (m *universalModel) ResourceName() string {
	return "account"
}

// SimpleSave 快速保存
func (m *universalModel) SimpleSave(ctx context.Context) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerAccountCollection)

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

// UpdateById 通过id更新数据库
func (m *universalModel) UpdateById(ctx context.Context) error {
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 更新数据库
	m.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())
	filter := bson.D{{Key: "_id", Value: m.ID}}
	_, err := coll.UpdateOne(ctx, filter,
		bson.D{{Key: "$set", Value: m}})
	if err != nil {
		return err
	}
	return nil
}

// FindByPhone 通过手机号查找到账号信息
func (m *universalModel) FindByPhone(ctx context.Context) error {
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 更新数据库
	filter := bson.D{{Key: "phone", Value: m.Phone}}
	err := coll.FindOne(ctx, filter).Decode(m)
	if err != nil {
		return err
	}
	return nil
}

// FindByAccountId 通过手机号查找到账号信息
func (m *universalModel) FindByAccountId(ctx context.Context) error {
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 更新数据库
	filter := bson.D{{Key: "account_id", Value: m.AccountId}}
	err := coll.FindOne(ctx, filter).Decode(m)
	if err != nil {
		return err
	}
	return nil
}

// Delete 快速删除
func (m *universalModel) Delete(ctx context.Context, id string) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerAccountCollection)
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

func (m *universalModel) List(ctx context.Context, merchantID string, urlParams *rest.UrlParams) ([]*universalModel, int64, error) {
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 声明需要返回的列表
	results := make([]*universalModel, 0)
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

// GetManyInIds 获取条件查询的结果
func (m *universalModel) GetManyInIds(ctx context.Context, ids []string) ([]*universalModel, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 绑定查询结果
	results := make([]*universalModel, 0)
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

// Detail 详情
func (m *universalModel) Detail(ctx context.Context, id string) (*universalModel, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerAccountCollection)
	// 绑定查询结果
	result := &universalModel{}
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}

	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.WithField("id", id).Error(err)
		return result, err
	}
	return result, nil
}

// Update 更新
func (m *universalModel) Update(ctx context.Context, id string) error {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ManagerAccountCollection)
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
