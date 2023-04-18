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
