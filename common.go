package collections

// APIInfo 接口信息
type APIInfo struct {
	// 路径
	Path string `json:"path" bson:"path"`
	// 名称
	Name string `json:"name" bson:"name"`
	// 描述
	Desc string `json:"desc" bson:"desc"`
	// 是否禁用
	Disable bool `json:"disable" bson:"disable"`
	// 是否可以访问详情
	CanViewDetail bool `json:"can_view_detail" bson:"can_view_detail"`
	// 是否在sidebar中隐藏
	// 默认false， 表示默认不隐藏
	HideOnSidebar bool `json:"hide_on_sidebar" bson:"hide_on_sidebar"`
}

// Address 地址
// Country Province City County District
type Address struct {
	// 国家
	Country string `json:"country" bson:"country"`
	// 省
	Province string `json:"province" bson:"province"`
	// 市
	City string `json:"city" bson:"city"`
	// 区
	District string `json:"district" bson:"district"`
	// 街道
	Street string `json:"street" bson:"street"`
	// 详情地址
	Detail string `json:"detail" bson:"detail"`
	// 详情地址

	// 金纬度
	LatitudeAndLongitude string `json:"latitude_and_longitude" bson:"latitude_and_longitude"`
}
