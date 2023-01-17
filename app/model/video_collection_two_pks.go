package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoCollectionTwoPks 数据对象
type VideoCollectionTwoPks struct {
	Id1         uint64      `json:"id1"`         // 视频集ID1，长整数格式
	Id2         string      `json:"id2"`         // 视频集ID2，字符串格式
	ContentType int         `json:"contentType"` // 内容类型
	FilterType  int         `json:"filterType"`  // 筛选类型
	Count       uint32      `json:"count"`       // 集合内视频数量
	IsOnline    bool        `json:"isOnline"`    // 是否上线：0 未上线|1 已上线
	CreatedAt   *gtime.Time `json:"createdAt"`   // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"`   // 更新时间
}

// VideoCollectionTwoPksInput 用于Insert、Update、Upsert的输入数据对象结构
type VideoCollectionTwoPksInput struct {
	Id1         uint64      `p:"id1" v:"required#视频集ID1不能为空" json:"id1"`               // 视频集ID1，长整数格式
	Id2         string      `p:"id2" v:"required#视频集ID2不能为空" json:"id2"`               // 视频集ID2，字符串格式
	ContentType int         `p:"contentType" v:"required#内容类型不能为空" json:"contentType"` // 内容类型
	FilterType  int         `p:"filterType" json:"filterType"`                         // 筛选类型
	Count       uint32      `p:"count" json:"count"`                                   // 集合内视频数量
	IsOnline    bool        `p:"isOnline" json:"isOnline"`                             // 是否上线：0 未上线|1 已上线
	CreatedAt   *gtime.Time `p:"createdAt" json:"createdAt"`                           // 创建时间
	UpdatedAt   *gtime.Time `p:"updatedAt" json:"updatedAt"`                           // 更新时间
}

// VideoCollectionTwoPksQuery 用于 Query By Example 模式的查询条件数据结构
type VideoCollectionTwoPksQuery struct {
	Id1         interface{} `orm:"id1,primary" json:"id1"`                                // 视频集ID1，长整数格式
	Id2         interface{} `orm:"id2,primary" wildcard:"none" json:"id2,omitempty"`      // 视频集ID2，字符串格式
	ContentType interface{} `orm:"content_type" json:"contentType,omitempty"`             // 内容类型
	FilterType  interface{} `orm:"filter_type" json:"filterType,omitempty"`               // 筛选类型
	Count       interface{} `orm:"count" json:"count,omitempty"`                          // 集合内视频数量
	IsOnline    interface{} `orm:"is_online" json:"isOnline,omitempty"`                   // 是否上线：0 未上线|1 已上线
	CreatedAt   interface{} `orm:"created_at" multi:"between" json:"createdAt,omitempty"` // 创建时间
	UpdatedAt   interface{} `orm:"updated_at" multi:"between" json:"updatedAt,omitempty"` // 更新时间
}

// VideoCollectionTwoPksCountReq 查询记录总条数的条件数据结构
type VideoCollectionTwoPksCountReq struct {
	g.Meta `json:"-" path:"/count" method:"get"`
	VideoCollectionTwoPksQuery
}

// VideoCollectionTwoPksCountRes 查询记录总条数的返回结果
type VideoCollectionTwoPksCountRes struct {
	Total int `json:"total"`
}

// VideoCollectionTwoPksOneReq 查询单一记录的条件数据结构
type VideoCollectionTwoPksOneReq struct {
	g.Meta `json:"-" path:"/one" method:"get"`
	VideoCollectionTwoPksQuery
	OrderBy string `json:"orderBy,omitempty"` // 排序方式
}

// VideoCollectionTwoPksOneRes 查询单一记录的返回结果
type VideoCollectionTwoPksOneRes struct {
	VideoCollectionTwoPks
}

// VideoCollectionTwoPksListReq 用于列表查询的查询条件数据结构，支持翻页和排序参数，支持查询条件参数类型自动转换
type VideoCollectionTwoPksListReq struct {
	g.Meta `json:"-" path:"/list" method:"get"`
	VideoCollectionTwoPksQuery
	Page     uint32 `d:"1" v:"min:0#分页号码错误" json:"page,omitempty"`          // 当前页码
	PageSize uint32 `d:"10" v:"max:50#分页数量最大50条" json:"pageSize,omitempty"` // 每页记录数
	OrderBy  string `json:"orderBy,omitempty"`                              // 排序方式
}

// VideoCollectionTwoPksListRes 分页返回结果
type VideoCollectionTwoPksListRes struct {
	Total   uint64                   `json:"total"`   // 记录总数
	Current uint32                   `json:"current"` // 当前页码
	Items   []*VideoCollectionTwoPks `json:"items"`   // 当前页记录列表
}

// VideoCollectionTwoPksCreateReq 插入操作请求参数
type VideoCollectionTwoPksCreateReq struct {
	g.Meta `json:"-" path:"/" method:"post"`
	VideoCollectionTwoPksInput
}

// VideoCollectionTwoPksCreateRes 插入操作返回结果
type VideoCollectionTwoPksCreateRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionTwoPksUpdateReq 更新操作请求参数
type VideoCollectionTwoPksUpdateReq struct {
	g.Meta `json:"-" path:"/:id1/:id2" method:"patch"`
	VideoCollectionTwoPksInput
}

// VideoCollectionTwoPksUpdateRes 更新操作返回结果
type VideoCollectionTwoPksUpdateRes struct {
	Message      string `json:"message"` // 提示信息
	RowsAffected int64  `json:"rowsAffected"`
}

// VideoCollectionTwoPksUpsertReq 更新插入操作请求参数
type VideoCollectionTwoPksUpsertReq struct {
	g.Meta `json:"-" path:"/" method:"put"`
	VideoCollectionTwoPksInput
}

// VideoCollectionTwoPksUpsertRes 更新插入操作返回结果
type VideoCollectionTwoPksUpsertRes struct {
	Message      string `json:"message"`      // 提示信息
	LastInsertId int64  `json:"lastInsertId"` // 上一条INSERT插入的记录主键，当主键为自增长时有效
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// VideoCollectionTwoPksDeleteReq 删除操作请求参数
type VideoCollectionTwoPksDeleteReq struct {
	g.Meta `json:"-" path:"/" method:"delete"`
	VideoCollectionTwoPksQuery
}

// VideoCollectionTwoPksDeleteRes 删除操作返回结果
type VideoCollectionTwoPksDeleteRes struct {
	Message      string `json:"message"`      // 提示信息
	RowsAffected int64  `json:"rowsAffected"` // 影响的条数
}

// IsEmpty 判断删除请求参数是否为空
func (q *VideoCollectionTwoPksDeleteReq) IsEmpty() bool {
	return q.Id1 == nil &&
		g.IsEmpty(q.Id2) &&
		q.ContentType == nil &&
		q.FilterType == nil &&
		q.Count == nil &&
		q.IsOnline == nil &&
		q.CreatedAt == nil &&
		q.UpdatedAt == nil
}
